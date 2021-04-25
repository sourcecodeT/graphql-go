package main

import (
	"demo-go/graph"
	"demo-go/graph/generated"
	"demo-go/graph/graphiql"
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

func main() {
	r := chi.NewRouter()

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8000"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	srv := handler.New(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{},
	}))

	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	r.Handle("/", graphiql.Handler("GraphQL playground", "/graphql"))
	r.Handle("/static/*", graphiql.Static())
	r.Handle("/graphql", srv)

	fmt.Println(`graphiql [http://localhost:8000]`)

	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatalln(err)
	}
}
