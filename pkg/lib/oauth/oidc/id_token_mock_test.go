// Code generated by MockGen. DO NOT EDIT.
// Source: id_token.go

// Package oidc is a generated GoMock package.
package oidc

import (
	context "context"
	url "net/url"
	reflect "reflect"

	model "github.com/authgear/authgear-server/pkg/api/model"
	oauth "github.com/authgear/authgear-server/pkg/lib/oauth"
	session "github.com/authgear/authgear-server/pkg/lib/session"
	idpsession "github.com/authgear/authgear-server/pkg/lib/session/idpsession"
	accesscontrol "github.com/authgear/authgear-server/pkg/util/accesscontrol"
	gomock "github.com/golang/mock/gomock"
	jwt "github.com/lestrrat-go/jwx/v2/jwt"
)

// MockUserProvider is a mock of UserProvider interface.
type MockUserProvider struct {
	ctrl     *gomock.Controller
	recorder *MockUserProviderMockRecorder
}

// MockUserProviderMockRecorder is the mock recorder for MockUserProvider.
type MockUserProviderMockRecorder struct {
	mock *MockUserProvider
}

// NewMockUserProvider creates a new mock instance.
func NewMockUserProvider(ctrl *gomock.Controller) *MockUserProvider {
	mock := &MockUserProvider{ctrl: ctrl}
	mock.recorder = &MockUserProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserProvider) EXPECT() *MockUserProviderMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockUserProvider) Get(ctx context.Context, id string, role accesscontrol.Role) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id, role)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUserProviderMockRecorder) Get(ctx, id, role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserProvider)(nil).Get), ctx, id, role)
}

// MockRolesAndGroupsProvider is a mock of RolesAndGroupsProvider interface.
type MockRolesAndGroupsProvider struct {
	ctrl     *gomock.Controller
	recorder *MockRolesAndGroupsProviderMockRecorder
}

// MockRolesAndGroupsProviderMockRecorder is the mock recorder for MockRolesAndGroupsProvider.
type MockRolesAndGroupsProviderMockRecorder struct {
	mock *MockRolesAndGroupsProvider
}

// NewMockRolesAndGroupsProvider creates a new mock instance.
func NewMockRolesAndGroupsProvider(ctrl *gomock.Controller) *MockRolesAndGroupsProvider {
	mock := &MockRolesAndGroupsProvider{ctrl: ctrl}
	mock.recorder = &MockRolesAndGroupsProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRolesAndGroupsProvider) EXPECT() *MockRolesAndGroupsProviderMockRecorder {
	return m.recorder
}

// ListEffectiveRolesByUserID mocks base method.
func (m *MockRolesAndGroupsProvider) ListEffectiveRolesByUserID(ctx context.Context, userID string) ([]*model.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEffectiveRolesByUserID", ctx, userID)
	ret0, _ := ret[0].([]*model.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEffectiveRolesByUserID indicates an expected call of ListEffectiveRolesByUserID.
func (mr *MockRolesAndGroupsProviderMockRecorder) ListEffectiveRolesByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEffectiveRolesByUserID", reflect.TypeOf((*MockRolesAndGroupsProvider)(nil).ListEffectiveRolesByUserID), ctx, userID)
}

// MockBaseURLProvider is a mock of BaseURLProvider interface.
type MockBaseURLProvider struct {
	ctrl     *gomock.Controller
	recorder *MockBaseURLProviderMockRecorder
}

// MockBaseURLProviderMockRecorder is the mock recorder for MockBaseURLProvider.
type MockBaseURLProviderMockRecorder struct {
	mock *MockBaseURLProvider
}

// NewMockBaseURLProvider creates a new mock instance.
func NewMockBaseURLProvider(ctrl *gomock.Controller) *MockBaseURLProvider {
	mock := &MockBaseURLProvider{ctrl: ctrl}
	mock.recorder = &MockBaseURLProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBaseURLProvider) EXPECT() *MockBaseURLProviderMockRecorder {
	return m.recorder
}

// Origin mocks base method.
func (m *MockBaseURLProvider) Origin() *url.URL {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Origin")
	ret0, _ := ret[0].(*url.URL)
	return ret0
}

// Origin indicates an expected call of Origin.
func (mr *MockBaseURLProviderMockRecorder) Origin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Origin", reflect.TypeOf((*MockBaseURLProvider)(nil).Origin))
}

// MockSessionLike is a mock of SessionLike interface.
type MockSessionLike struct {
	ctrl     *gomock.Controller
	recorder *MockSessionLikeMockRecorder
}

// MockSessionLikeMockRecorder is the mock recorder for MockSessionLike.
type MockSessionLikeMockRecorder struct {
	mock *MockSessionLike
}

// NewMockSessionLike creates a new mock instance.
func NewMockSessionLike(ctrl *gomock.Controller) *MockSessionLike {
	mock := &MockSessionLike{ctrl: ctrl}
	mock.recorder = &MockSessionLikeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionLike) EXPECT() *MockSessionLikeMockRecorder {
	return m.recorder
}

// SessionID mocks base method.
func (m *MockSessionLike) SessionID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SessionID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SessionID indicates an expected call of SessionID.
func (mr *MockSessionLikeMockRecorder) SessionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SessionID", reflect.TypeOf((*MockSessionLike)(nil).SessionID))
}

// SessionType mocks base method.
func (m *MockSessionLike) SessionType() session.Type {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SessionType")
	ret0, _ := ret[0].(session.Type)
	return ret0
}

// SessionType indicates an expected call of SessionType.
func (mr *MockSessionLikeMockRecorder) SessionType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SessionType", reflect.TypeOf((*MockSessionLike)(nil).SessionType))
}

// MockIDTokenHintResolverIssuer is a mock of IDTokenHintResolverIssuer interface.
type MockIDTokenHintResolverIssuer struct {
	ctrl     *gomock.Controller
	recorder *MockIDTokenHintResolverIssuerMockRecorder
}

// MockIDTokenHintResolverIssuerMockRecorder is the mock recorder for MockIDTokenHintResolverIssuer.
type MockIDTokenHintResolverIssuerMockRecorder struct {
	mock *MockIDTokenHintResolverIssuer
}

// NewMockIDTokenHintResolverIssuer creates a new mock instance.
func NewMockIDTokenHintResolverIssuer(ctrl *gomock.Controller) *MockIDTokenHintResolverIssuer {
	mock := &MockIDTokenHintResolverIssuer{ctrl: ctrl}
	mock.recorder = &MockIDTokenHintResolverIssuerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDTokenHintResolverIssuer) EXPECT() *MockIDTokenHintResolverIssuerMockRecorder {
	return m.recorder
}

// VerifyIDToken mocks base method.
func (m *MockIDTokenHintResolverIssuer) VerifyIDToken(idTokenHint string) (jwt.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyIDToken", idTokenHint)
	ret0, _ := ret[0].(jwt.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyIDToken indicates an expected call of VerifyIDToken.
func (mr *MockIDTokenHintResolverIssuerMockRecorder) VerifyIDToken(idTokenHint interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyIDToken", reflect.TypeOf((*MockIDTokenHintResolverIssuer)(nil).VerifyIDToken), idTokenHint)
}

// MockIDTokenHintResolverSessionProvider is a mock of IDTokenHintResolverSessionProvider interface.
type MockIDTokenHintResolverSessionProvider struct {
	ctrl     *gomock.Controller
	recorder *MockIDTokenHintResolverSessionProviderMockRecorder
}

// MockIDTokenHintResolverSessionProviderMockRecorder is the mock recorder for MockIDTokenHintResolverSessionProvider.
type MockIDTokenHintResolverSessionProviderMockRecorder struct {
	mock *MockIDTokenHintResolverSessionProvider
}

// NewMockIDTokenHintResolverSessionProvider creates a new mock instance.
func NewMockIDTokenHintResolverSessionProvider(ctrl *gomock.Controller) *MockIDTokenHintResolverSessionProvider {
	mock := &MockIDTokenHintResolverSessionProvider{ctrl: ctrl}
	mock.recorder = &MockIDTokenHintResolverSessionProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDTokenHintResolverSessionProvider) EXPECT() *MockIDTokenHintResolverSessionProviderMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockIDTokenHintResolverSessionProvider) Get(ctx context.Context, id string) (*idpsession.IDPSession, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*idpsession.IDPSession)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIDTokenHintResolverSessionProviderMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIDTokenHintResolverSessionProvider)(nil).Get), ctx, id)
}

// MockIDTokenHintResolverOfflineGrantService is a mock of IDTokenHintResolverOfflineGrantService interface.
type MockIDTokenHintResolverOfflineGrantService struct {
	ctrl     *gomock.Controller
	recorder *MockIDTokenHintResolverOfflineGrantServiceMockRecorder
}

// MockIDTokenHintResolverOfflineGrantServiceMockRecorder is the mock recorder for MockIDTokenHintResolverOfflineGrantService.
type MockIDTokenHintResolverOfflineGrantServiceMockRecorder struct {
	mock *MockIDTokenHintResolverOfflineGrantService
}

// NewMockIDTokenHintResolverOfflineGrantService creates a new mock instance.
func NewMockIDTokenHintResolverOfflineGrantService(ctrl *gomock.Controller) *MockIDTokenHintResolverOfflineGrantService {
	mock := &MockIDTokenHintResolverOfflineGrantService{ctrl: ctrl}
	mock.recorder = &MockIDTokenHintResolverOfflineGrantServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDTokenHintResolverOfflineGrantService) EXPECT() *MockIDTokenHintResolverOfflineGrantServiceMockRecorder {
	return m.recorder
}

// GetOfflineGrant mocks base method.
func (m *MockIDTokenHintResolverOfflineGrantService) GetOfflineGrant(ctx context.Context, id string) (*oauth.OfflineGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOfflineGrant", ctx, id)
	ret0, _ := ret[0].(*oauth.OfflineGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOfflineGrant indicates an expected call of GetOfflineGrant.
func (mr *MockIDTokenHintResolverOfflineGrantServiceMockRecorder) GetOfflineGrant(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOfflineGrant", reflect.TypeOf((*MockIDTokenHintResolverOfflineGrantService)(nil).GetOfflineGrant), ctx, id)
}
