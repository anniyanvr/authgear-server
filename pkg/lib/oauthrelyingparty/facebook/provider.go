package facebook

import (
	"context"
	"net/url"

	"github.com/authgear/oauthrelyingparty/pkg/api/oauthrelyingparty"

	"github.com/authgear/authgear-server/pkg/lib/authn/stdattrs"
	liboauthrelyingparty "github.com/authgear/authgear-server/pkg/lib/oauthrelyingparty"
	"github.com/authgear/authgear-server/pkg/lib/oauthrelyingparty/oauthrelyingpartyutil"
	"github.com/authgear/authgear-server/pkg/util/crypto"
	"github.com/authgear/authgear-server/pkg/util/validation"
)

func init() {
	oauthrelyingparty.RegisterProvider(Type, Facebook{})
}

const Type = liboauthrelyingparty.TypeFacebook

var _ oauthrelyingparty.Provider = Facebook{}

const (
	facebookAuthorizationURL string = "https://www.facebook.com/v11.0/dialog/oauth"
	// nolint: gosec
	facebookTokenURL    string = "https://graph.facebook.com/v11.0/oauth/access_token"
	facebookUserInfoURL string = "https://graph.facebook.com/v11.0/me?fields=id,email,first_name,last_name,middle_name,name,name_format,picture,short_name"
)

type Facebook struct{}

func (Facebook) GetJSONSchema() map[string]interface{} {
	builder := validation.SchemaBuilder{}
	builder.Type(validation.TypeObject)
	builder.Properties().
		Property("type", validation.SchemaBuilder{}.Type(validation.TypeString)).
		Property("client_id", validation.SchemaBuilder{}.Type(validation.TypeString).MinLength(1)).
		Property("claims", validation.SchemaBuilder{}.Type(validation.TypeObject).
			AdditionalPropertiesFalse().
			Properties().
			Property("email", validation.SchemaBuilder{}.Type(validation.TypeObject).
				AdditionalPropertiesFalse().Properties().
				Property("assume_verified", validation.SchemaBuilder{}.Type(validation.TypeBoolean)).
				Property("required", validation.SchemaBuilder{}.Type(validation.TypeBoolean)),
			),
		)
	builder.Required("type", "client_id")
	return builder
}

func (Facebook) SetDefaults(cfg oauthrelyingparty.ProviderConfig) {
	cfg.SetDefaultsEmailClaimConfig(oauthrelyingpartyutil.Email_AssumeVerified_Required())
}

func (Facebook) ProviderID(cfg oauthrelyingparty.ProviderConfig) oauthrelyingparty.ProviderID {
	// Facebook does NOT support OIDC.
	// Facebook user ID is scoped to client_id.
	// Therefore, ProviderID is Type + client_id.
	//
	// Rotating the OAuth application is problematic.
	// But if email remains unchanged, the user can associate their account.
	keys := map[string]interface{}{
		"client_id": cfg.ClientID(),
	}
	return oauthrelyingparty.NewProviderID(cfg.Type(), keys)
}

func (Facebook) scope() []string {
	// https://developers.facebook.com/docs/permissions/reference
	return []string{"email", "public_profile"}
}

func (p Facebook) GetAuthorizationURL(ctx context.Context, deps oauthrelyingparty.Dependencies, param oauthrelyingparty.GetAuthorizationURLOptions) (string, error) {
	return oauthrelyingpartyutil.MakeAuthorizationURL(facebookAuthorizationURL, oauthrelyingpartyutil.AuthorizationURLParams{
		ClientID:     deps.ProviderConfig.ClientID(),
		RedirectURI:  param.RedirectURI,
		Scope:        p.scope(),
		ResponseType: oauthrelyingparty.ResponseTypeCode,
		// ResponseMode is unset
		State: param.State,
		// Prompt is unset.
		// Facebook doesn't support prompt parameter
		// https://developers.facebook.com/docs/facebook-login/manually-build-a-login-flow/

		// Nonce is unset
	}.Query()), nil
}

func (Facebook) GetUserProfile(ctx context.Context, deps oauthrelyingparty.Dependencies, param oauthrelyingparty.GetUserProfileOptions) (authInfo oauthrelyingparty.UserProfile, err error) {
	authInfo = oauthrelyingparty.UserProfile{}

	code, err := oauthrelyingpartyutil.GetCode(param.Query)
	if err != nil {
		return
	}

	accessTokenResp, err := oauthrelyingpartyutil.FetchAccessTokenResp(
		ctx,
		deps.HTTPClient,
		code,
		facebookTokenURL,
		param.RedirectURI,
		deps.ProviderConfig.ClientID(),
		deps.ClientSecret,
	)
	if err != nil {
		return
	}

	userProfileURL, err := url.Parse(facebookUserInfoURL)
	if err != nil {
		return
	}
	q := userProfileURL.Query()
	appSecretProof := crypto.HMACSHA256String([]byte(deps.ClientSecret), []byte(accessTokenResp.AccessToken()))
	q.Set("appsecret_proof", appSecretProof)
	userProfileURL.RawQuery = q.Encode()

	// Here is the refacted user profile of Louis' facebook account.
	// {
	//   "id": "redacted",
	//   "email": "redacted",
	//   "first_name": "Jonathan",
	//   "last_name": "Doe",
	//   "name": "Johnathan Doe",
	//   "name_format": "{first} {last}",
	//   "picture": {
	//     "data": {
	//       "height": 50,
	//       "is_silhouette": true,
	//       "url": "http://example.com",
	//       "width": 50
	//     }
	//   },
	//   "short_name": "John"
	// }

	userProfile, err := oauthrelyingpartyutil.FetchUserProfile(ctx, deps.HTTPClient, accessTokenResp, userProfileURL.String())
	if err != nil {
		return
	}
	authInfo.ProviderRawProfile = userProfile

	id, _ := userProfile["id"].(string)
	email, _ := userProfile["email"].(string)
	firstName, _ := userProfile["first_name"].(string)
	lastName, _ := userProfile["last_name"].(string)
	name, _ := userProfile["name"].(string)
	shortName, _ := userProfile["short_name"].(string)
	var picture string
	if pictureObj, ok := userProfile["picture"].(map[string]interface{}); ok {
		if data, ok := pictureObj["data"].(map[string]interface{}); ok {
			if url, ok := data["url"].(string); ok {
				picture = url
			}
		}
	}

	authInfo.ProviderUserID = id
	emailRequired := deps.ProviderConfig.EmailClaimConfig().Required()
	stdAttrs, err := stdattrs.Extract(map[string]interface{}{
		stdattrs.Email:      email,
		stdattrs.GivenName:  firstName,
		stdattrs.FamilyName: lastName,
		stdattrs.Name:       name,
		stdattrs.Nickname:   shortName,
		stdattrs.Picture:    picture,
	}, stdattrs.ExtractOptions{
		EmailRequired: emailRequired,
	})
	if err != nil {
		return
	}
	authInfo.StandardAttributes = stdAttrs

	return
}
