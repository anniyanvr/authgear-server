package graphql

import (
	"context"

	"github.com/graphql-go/graphql"

	relay "github.com/authgear/authgear-server/pkg/graphqlgo/relay"

	"github.com/authgear/authgear-server/pkg/portal/model"
	"github.com/authgear/authgear-server/pkg/portal/session"
	"github.com/authgear/authgear-server/pkg/util/geoip"
	"github.com/authgear/authgear-server/pkg/util/httputil"
)

const typeViewer = "Viewer"

var viewerSubresolver = func(ctx context.Context, gqlCtx *Context, id string) (interface{}, error) {
	userIface, err := gqlCtx.Users.Load(ctx, id).Value()
	if err != nil {
		return nil, err
	}

	user := userIface.(*model.User)

	requestIP := httputil.GetIP(gqlCtx.Request, bool(gqlCtx.TrustProxy))
	geoipInfo, ok := geoip.IPString(requestIP)
	if ok {
		user.GeoIPCountryCode = geoipInfo.CountryCode
	}

	return user, nil
}

var nodeViewer = node(
	graphql.NewObject(graphql.ObjectConfig{
		Name:        typeViewer,
		Description: "The viewer",
		Interfaces: []*graphql.Interface{
			nodeDefs.NodeInterface,
		},
		Fields: graphql.Fields{
			"id": relay.GlobalIDField(typeViewer, nil),
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"formattedName": &graphql.Field{
				Type: graphql.String,
			},
			"projectQuota": &graphql.Field{
				Type: graphql.Int,
			},
			"projectOwnerCount": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"geoIPCountryCode": &graphql.Field{
				Type: graphql.String,
			},
			"isOnboardingSurveyCompleted": &graphql.Field{
				Type: graphql.Boolean,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					user := p.Source.(*model.User)
					ctx := p.Context
					gqlCtx := GQLContext(ctx)
					isCompleted, err := gqlCtx.OnboardService.CheckOnboardingSurveyCompletion(ctx, user.ID)
					if err != nil {
						return nil, err
					}
					return isCompleted, nil
				},
			},
		},
	}),
	&model.User{},
	func(ctx context.Context, id string) (interface{}, error) {
		gqlCtx := GQLContext(ctx)

		// Ensure only the authenticated user can fetch their own viewer.
		sessionInfo := session.GetValidSessionInfo(ctx)
		if sessionInfo == nil {
			return nil, nil
		}
		if sessionInfo.UserID != id {
			return nil, nil
		}

		return viewerSubresolver(ctx, gqlCtx, id)
	},
)
