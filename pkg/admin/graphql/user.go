package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"

	"github.com/authgear/authgear-server/pkg/lib/api/model"
	"github.com/authgear/authgear-server/pkg/lib/authn/user"
)

const typeUser = "User"

var nodeUser = entity(
	graphql.NewObject(graphql.ObjectConfig{
		Name:        typeUser,
		Description: "Authgear user",
		Interfaces: []*graphql.Interface{
			nodeDefs.NodeInterface,
			entityInterface,
		},
		Fields: graphql.Fields{
			"id":        relay.GlobalIDField(typeUser, nil),
			"createdAt": entityCreatedAtField,
			"updatedAt": entityUpdatedAtField,
			"lastLoginAt": &graphql.Field{
				Type:        graphql.DateTime,
				Description: "The last login time of user",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					ref := p.Source.(*user.Ref)
					user, err := GQLContext(p.Context).Users.Get(ref.ID)
					if err != nil {
						return nil, err
					}
					return user.LastLoginAt, nil
				},
			},
		},
	}),
	&model.User{},
	func(ctx *Context, id string) (interface{}, error) {
		return ctx.Users.Get(id)
	},
)

var connUser = connection(nodeUser)
