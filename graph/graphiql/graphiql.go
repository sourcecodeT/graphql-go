package graphiql

import (
	"embed"
	"html/template"
	"net/http"
)

var page = template.Must(template.New("graphiql").Parse(`<!DOCTYPE html>
<!DOCTYPE html>
<html>
  <head>
    <style>
      html,
      body {
        height: 100%;
        margin: 0;
        overflow: hidden;
        width: 100%;
      }
      #graphiql {
        height: 100vh;
      }
    </style>

    <link
      rel="stylesheet"
      href="static/graphiqlWithExtensions.css"
    />
    <script src="static/fetch.min.txt"></script>
    <script src="static/react.production.min.txt"></script>
    <script src="static/react-dom.production.min.txt"></script>
    <script src="static/graphiqlWithExtensions.min.txt"></script>
    <script src="static/subscriptions-transport-ws.txt"></script>
    <script src="static/subscriptions-fetcher.txt"></script>
    <title>GraphiQL</title>
  </head>
  <body>
    <div id="graphiql"></div>
    <script>
      var fetchURL = location.protocol + '//' + location.host + '{{.endpoint}}';
      var wsProto = location.protocol == 'https:' ? 'wss:' : 'ws:';
      var wsURL = wsProto + '//' + location.host + '{{.endpoint }}';

      function graphQLFetcher(graphQLParams) {
        return fetch(fetchURL, {
          method: "post",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(graphQLParams),
          credentials: "omit",
        }).then((response) => response.json());
      } 

      var subscriptionsClient = new window.SubscriptionsTransportWs.SubscriptionClient(wsURL, { reconnect: true });
      var subscriptionsFetcher = window.GraphiQLSubscriptionsFetcher.graphQLFetcher(subscriptionsClient, graphQLFetcher);

      ReactDOM.render(
        React.createElement(GraphiQLWithExtensions.GraphiQLWithExtensions, {
          fetcher: subscriptionsFetcher,
        }),
        document.getElementById('graphiql'),
      );
    </script>
  </body>
</html>
`))

func Handler(title string, endpoint string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		err := page.Execute(w, map[string]string{
			"endpoint": endpoint,
		})
		if err != nil {
			panic(err)
		}
	}
}

//go:embed static/*
var content embed.FS

func Static() http.Handler {
	return http.FileServer(http.FS(content))
}
