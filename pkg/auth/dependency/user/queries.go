package user

import (
	"github.com/authgear/authgear-server/pkg/auth/dependency/authenticator"
	"github.com/authgear/authgear-server/pkg/auth/dependency/identity"
	"github.com/authgear/authgear-server/pkg/auth/model"
)

type IdentityProvider interface {
	ListByUser(userID string) ([]*identity.Info, error)
}

type VerificationService interface {
	IsUserVerified(userID string) (bool, error)
	IsVerified(identities []*identity.Info, authenticators []*authenticator.Info) bool
}

type Queries struct {
	Store        store
	Identities   IdentityProvider
	Verification VerificationService
}

func (p *Queries) Get(id string) (*model.User, error) {
	user, err := p.Store.Get(id)
	if err != nil {
		return nil, err
	}

	identities, err := p.Identities.ListByUser(id)
	if err != nil {
		return nil, err
	}

	isVerified, err := p.Verification.IsUserVerified(id)
	if err != nil {
		return nil, err
	}

	return newUserModel(user, identities, isVerified), nil
}
