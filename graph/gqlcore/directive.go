package gqlcore

import (
	"context"
	"demo-go/graph/generated"
	"demo-go/graph/model"

	"github.com/99designs/gqlgen/graphql"
)

func NewDirectiveRoot() generated.DirectiveRoot {
	return generated.DirectiveRoot{}
}

func HasRole(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (res interface{}, err error) {
	return next(ctx)
}
