package queries

import (
	"context"
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/rodrwan/sicily/graph"
	"github.com/rodrwan/sicily/graph/types"
	"github.com/rodrwan/syracuse/citizens"
)

// GetUser fill graphql Field with data from postgres service.
func GetUser(ctx *graph.Context) *graphql.Field {
	return &graphql.Field{
		Type:        types.User,
		Description: "Get user by id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "return user information by id",
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id, ok := params.Args["id"].(string)
			if !ok {
				return nil, errors.New("Invalid params")
			}

			ctxb := context.Background()
			opts := &citizens.GetRequest{
				UserId: id,
			}
			u, err := ctx.UserService.Get(ctxb, opts)
			if err != nil {
				return nil, err
			}

			return u.Data, nil
		},
	}
}

// GetUsers get a collection of users
func GetUsers(ctx *graph.Context) *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(types.User),
		Description: "Get collection of users",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			ctxb := context.Background()
			opts := &citizens.SelectRequest{}
			uu, err := ctx.UserService.Select(ctxb, opts)
			if err != nil {
				return nil, err
			}

			return uu.Data, nil
		},
	}
}

// Users expose UserQuery
func Users(ctx *graph.Context) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "UserQueries",
		Fields: graphql.Fields{
			"getUser":  GetUser(ctx),
			"getUsers": GetUsers(ctx),
		},
	})
}
