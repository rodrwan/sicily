package mutation

import (
	"context"
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/rodrwan/sicily/graph"
	"github.com/rodrwan/sicily/graph/types"
	"github.com/rodrwan/syracuse/citizens"
)

// CreateUser create a user in remote service.
func CreateUser(ctx *graph.Context) *graphql.Field {
	return &graphql.Field{
		Type:        types.User,
		Description: "Create user",
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"fullname": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			email, ok := params.Args["email"].(string)
			if !ok {
				return nil, errors.New("Invalid params")
			}

			fullname, ok := params.Args["fullname"].(string)
			if !ok {
				return nil, errors.New("Invalid params")
			}

			ctxb := context.Background()
			opts := &citizens.CreateRequest{
				Data: &citizens.Citizen{
					Email:    email,
					Fullname: fullname,
				},
			}

			u, err := ctx.UserService.Create(ctxb, opts)
			if err != nil {
				return nil, err
			}

			return u.Data, nil
		},
	}
}

// Users expose UserQuery
func Users(ctx *graph.Context) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "UserMutations",
		Fields: graphql.Fields{
			"createUser": CreateUser(ctx),
		},
	})
}
