package intents

import (
	"context"
	"fmt"

	"github.com/authgear/authgear-server/pkg/lib/authn"
	"github.com/authgear/authgear-server/pkg/lib/interaction"
	"github.com/authgear/authgear-server/pkg/lib/interaction/nodes"
)

func init() {
	interaction.RegisterIntent(&IntentAddIdentity{})
}

type IntentAddIdentity struct {
	UserID string `json:"user_id"`
}

func NewIntentAddIdentity(userID string) *IntentAddIdentity {
	return &IntentAddIdentity{
		UserID: userID,
	}
}

func (i *IntentAddIdentity) InstantiateRootNode(goCtx context.Context, ctx *interaction.Context, graph *interaction.Graph) (interaction.Node, error) {
	edge := nodes.EdgeDoUseUser{UseUserID: i.UserID}
	return edge.Instantiate(goCtx, ctx, graph, i)
}

func (i *IntentAddIdentity) DeriveEdgesForNode(goCtx context.Context, graph *interaction.Graph, node interaction.Node) ([]interaction.Edge, error) {
	switch node := node.(type) {
	case *nodes.NodeDoUseUser:
		return []interaction.Edge{
			&nodes.EdgeCreateIdentityBegin{},
		}, nil

	case *nodes.NodeCreateIdentityEnd:
		return []interaction.Edge{
			&nodes.EdgeDoCreateIdentity{
				Identity:   node.IdentityInfo,
				IsAddition: true,
			},
		}, nil
	case *nodes.NodeDoCreateIdentity:
		return []interaction.Edge{
			&nodes.EdgeEnsureVerificationBegin{
				Identity:        node.Identity,
				RequestedByUser: false,
			},
		}, nil

	case *nodes.NodeEnsureVerificationEnd:
		return []interaction.Edge{
			&nodes.EdgeDoVerifyIdentity{
				Identity:         node.Identity,
				NewVerifiedClaim: node.NewVerifiedClaim,
			},
		}, nil

	case *nodes.NodeDoVerifyIdentity:
		return []interaction.Edge{
			&nodes.EdgeDoUseIdentity{Identity: node.Identity},
		}, nil

	case *nodes.NodeDoUseIdentity:
		return []interaction.Edge{&nodes.EdgeEnsureRemoveAnonymousIdentity{}}, nil
	case *nodes.NodeEnsureRemoveAnonymousIdentity:
		return []interaction.Edge{
			&nodes.EdgeCreateAuthenticatorBegin{
				Stage: authn.AuthenticationStagePrimary,
			},
		}, nil

	case *nodes.NodeCreateAuthenticatorEnd:
		return []interaction.Edge{
			&nodes.EdgeDoCreateAuthenticator{
				Stage:          node.Stage,
				Authenticators: node.Authenticators,
			},
		}, nil

	case *nodes.NodeDoCreateAuthenticator:
		switch node.Stage {
		case authn.AuthenticationStagePrimary:
			return nil, nil

		case authn.AuthenticationStageSecondary:
			return []interaction.Edge{
				&nodes.EdgeGenerateRecoveryCode{},
			}, nil

		default:
			panic(fmt.Errorf("interaction: unexpected authentication stage: %v", node.Stage))
		}

	case *nodes.NodeGenerateRecoveryCodeEnd:
		return []interaction.Edge{
			&nodes.EdgeDoGenerateRecoveryCode{
				RecoveryCodes: node.RecoveryCodes,
			},
		}, nil

	case *nodes.NodeDoGenerateRecoveryCode:
		return nil, nil

	default:
		panic(fmt.Errorf("interaction: unexpected node: %T", node))
	}
}
