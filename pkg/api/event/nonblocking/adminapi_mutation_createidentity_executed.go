package nonblocking

import (
	"github.com/authgear/authgear-server/pkg/api/event"
	"github.com/authgear/authgear-server/pkg/api/model"
)

const (
	AdminAPIMutationCreateIdentityExecuted event.Type = "admin_api.mutation.create_identity.executed"
)

type AdminAPIMutationCreateIdentityExecutedEventPayload struct {
	UserRef   model.UserRef  `json:"-" resolve:"user"`
	UserModel model.User     `json:"user"`
	Identity  model.Identity `json:"identity"`
}

func (e *AdminAPIMutationCreateIdentityExecutedEventPayload) NonBlockingEventType() event.Type {
	return AdminAPIMutationCreateIdentityExecuted
}

func (e *AdminAPIMutationCreateIdentityExecutedEventPayload) UserID() string {
	return e.UserModel.ID
}

func (e *AdminAPIMutationCreateIdentityExecutedEventPayload) GetTriggeredBy() event.TriggeredByType {
	return event.TriggeredByTypeAdminAPI
}

func (e *AdminAPIMutationCreateIdentityExecutedEventPayload) FillContext(ctx *event.Context) {
}

func (e *AdminAPIMutationCreateIdentityExecutedEventPayload) ForHook() bool {
	return false
}

func (e *AdminAPIMutationCreateIdentityExecutedEventPayload) ForAudit() bool {
	return true
}

func (e *AdminAPIMutationCreateIdentityExecutedEventPayload) ReindexUserNeeded() bool {
	return false
}

func (e *AdminAPIMutationCreateIdentityExecutedEventPayload) IsUserDeleted() bool {
	return false
}

var _ event.NonBlockingPayload = &AdminAPIMutationCreateIdentityExecutedEventPayload{}
