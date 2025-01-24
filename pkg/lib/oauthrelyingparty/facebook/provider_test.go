package facebook

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/authgear/oauthrelyingparty/pkg/api/oauthrelyingparty"
)

func TestFacebook(t *testing.T) {
	Convey("Facebook", t, func() {
		deps := oauthrelyingparty.Dependencies{
			ProviderConfig: oauthrelyingparty.ProviderConfig{
				"client_id": "client_id",
				"type":      Type,
			},
		}
		g := Facebook{}

		ctx := context.Background()
		u, err := g.GetAuthorizationURL(ctx, deps, oauthrelyingparty.GetAuthorizationURLOptions{
			RedirectURI: "https://localhost/",
			Nonce:       "nonce",
			State:       "state",
			Prompt:      []string{"login"},
		})
		So(err, ShouldBeNil)
		So(u, ShouldEqual, "https://www.facebook.com/v11.0/dialog/oauth?client_id=client_id&redirect_uri=https%3A%2F%2Flocalhost%2F&response_type=code&scope=email+public_profile&state=state")
	})
}
