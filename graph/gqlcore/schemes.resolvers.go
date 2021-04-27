package gqlcore

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"demo-go/graph/generated"
	"demo-go/graph/model"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	return nil, nil
}

func (r *mutationResolver) SingleUpload(ctx context.Context, file graphql.Upload) (bool, error) {
	f, err := os.Create(file.Filename)
	if err != nil {
		return false, err
	}
	defer f.Close()

	if _, err := io.Copy(f, file.File); err != nil {
		return false, err
	}

	return true, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return nil, nil
}

func (r *queryResolver) Baz(ctx context.Context) (bool, error) {
	return true, nil
}

func (r *subscriptionResolver) Baz(ctx context.Context, id string) (<-chan string, error) {
	todoc := make(chan string)

	go func() {
		defer fmt.Printf("done: %s\n", id)

		for {
			select {
			case <-ctx.Done():
				close(todoc)
				return
			default:
				todoc <- time.Now().String()
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return todoc, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
