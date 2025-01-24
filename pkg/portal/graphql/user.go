package graphql

import (
	"context"

	"github.com/graphql-go/graphql"

	"github.com/authgear/authgear-server/pkg/graphqlgo/relay"

	"github.com/authgear/authgear-server/pkg/portal/model"
)

const typeUser = "User"

var nodeUser = node(
	graphql.NewObject(graphql.ObjectConfig{
		Name:        typeUser,
		Description: "Portal User",
		Interfaces: []*graphql.Interface{
			nodeDefs.NodeInterface,
		},
		Fields: graphql.Fields{
			"id": relay.GlobalIDField(typeUser, nil),
			"email": &graphql.Field{
				Type: graphql.String,
			},
		},
	}),
	&model.User{},
	func(ctx context.Context, id string) (interface{}, error) {
		gqlCtx := GQLContext(ctx)
		return gqlCtx.Users.Load(ctx, id).Value, nil
	},
)
