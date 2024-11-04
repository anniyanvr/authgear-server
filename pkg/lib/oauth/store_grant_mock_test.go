// Code generated by MockGen. DO NOT EDIT.
// Source: store_grant.go

// Package oauth is a generated GoMock package.
package oauth

import (
	context "context"
	reflect "reflect"
	time "time"

	access "github.com/authgear/authgear-server/pkg/lib/session/access"
	gomock "github.com/golang/mock/gomock"
)

// MockCodeGrantStore is a mock of CodeGrantStore interface.
type MockCodeGrantStore struct {
	ctrl     *gomock.Controller
	recorder *MockCodeGrantStoreMockRecorder
}

// MockCodeGrantStoreMockRecorder is the mock recorder for MockCodeGrantStore.
type MockCodeGrantStoreMockRecorder struct {
	mock *MockCodeGrantStore
}

// NewMockCodeGrantStore creates a new mock instance.
func NewMockCodeGrantStore(ctrl *gomock.Controller) *MockCodeGrantStore {
	mock := &MockCodeGrantStore{ctrl: ctrl}
	mock.recorder = &MockCodeGrantStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCodeGrantStore) EXPECT() *MockCodeGrantStoreMockRecorder {
	return m.recorder
}

// CreateCodeGrant mocks base method.
func (m *MockCodeGrantStore) CreateCodeGrant(ctx context.Context, g *CodeGrant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCodeGrant", ctx, g)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCodeGrant indicates an expected call of CreateCodeGrant.
func (mr *MockCodeGrantStoreMockRecorder) CreateCodeGrant(ctx, g interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCodeGrant", reflect.TypeOf((*MockCodeGrantStore)(nil).CreateCodeGrant), ctx, g)
}

// DeleteCodeGrant mocks base method.
func (m *MockCodeGrantStore) DeleteCodeGrant(ctx context.Context, g *CodeGrant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCodeGrant", ctx, g)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCodeGrant indicates an expected call of DeleteCodeGrant.
func (mr *MockCodeGrantStoreMockRecorder) DeleteCodeGrant(ctx, g interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCodeGrant", reflect.TypeOf((*MockCodeGrantStore)(nil).DeleteCodeGrant), ctx, g)
}

// GetCodeGrant mocks base method.
func (m *MockCodeGrantStore) GetCodeGrant(ctx context.Context, codeHash string) (*CodeGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCodeGrant", ctx, codeHash)
	ret0, _ := ret[0].(*CodeGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCodeGrant indicates an expected call of GetCodeGrant.
func (mr *MockCodeGrantStoreMockRecorder) GetCodeGrant(ctx, codeHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCodeGrant", reflect.TypeOf((*MockCodeGrantStore)(nil).GetCodeGrant), ctx, codeHash)
}

// MockSettingsActionGrantStore is a mock of SettingsActionGrantStore interface.
type MockSettingsActionGrantStore struct {
	ctrl     *gomock.Controller
	recorder *MockSettingsActionGrantStoreMockRecorder
}

// MockSettingsActionGrantStoreMockRecorder is the mock recorder for MockSettingsActionGrantStore.
type MockSettingsActionGrantStoreMockRecorder struct {
	mock *MockSettingsActionGrantStore
}

// NewMockSettingsActionGrantStore creates a new mock instance.
func NewMockSettingsActionGrantStore(ctrl *gomock.Controller) *MockSettingsActionGrantStore {
	mock := &MockSettingsActionGrantStore{ctrl: ctrl}
	mock.recorder = &MockSettingsActionGrantStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSettingsActionGrantStore) EXPECT() *MockSettingsActionGrantStoreMockRecorder {
	return m.recorder
}

// CreateSettingsActionGrant mocks base method.
func (m *MockSettingsActionGrantStore) CreateSettingsActionGrant(ctx context.Context, g *SettingsActionGrant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSettingsActionGrant", ctx, g)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSettingsActionGrant indicates an expected call of CreateSettingsActionGrant.
func (mr *MockSettingsActionGrantStoreMockRecorder) CreateSettingsActionGrant(ctx, g interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSettingsActionGrant", reflect.TypeOf((*MockSettingsActionGrantStore)(nil).CreateSettingsActionGrant), ctx, g)
}

// DeleteSettingsActionGrant mocks base method.
func (m *MockSettingsActionGrantStore) DeleteSettingsActionGrant(ctx context.Context, g *SettingsActionGrant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSettingsActionGrant", ctx, g)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSettingsActionGrant indicates an expected call of DeleteSettingsActionGrant.
func (mr *MockSettingsActionGrantStoreMockRecorder) DeleteSettingsActionGrant(ctx, g interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSettingsActionGrant", reflect.TypeOf((*MockSettingsActionGrantStore)(nil).DeleteSettingsActionGrant), ctx, g)
}

// GetSettingsActionGrant mocks base method.
func (m *MockSettingsActionGrantStore) GetSettingsActionGrant(ctx context.Context, codeHash string) (*SettingsActionGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSettingsActionGrant", ctx, codeHash)
	ret0, _ := ret[0].(*SettingsActionGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSettingsActionGrant indicates an expected call of GetSettingsActionGrant.
func (mr *MockSettingsActionGrantStoreMockRecorder) GetSettingsActionGrant(ctx, codeHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSettingsActionGrant", reflect.TypeOf((*MockSettingsActionGrantStore)(nil).GetSettingsActionGrant), ctx, codeHash)
}

// MockOfflineGrantStore is a mock of OfflineGrantStore interface.
type MockOfflineGrantStore struct {
	ctrl     *gomock.Controller
	recorder *MockOfflineGrantStoreMockRecorder
}

// MockOfflineGrantStoreMockRecorder is the mock recorder for MockOfflineGrantStore.
type MockOfflineGrantStoreMockRecorder struct {
	mock *MockOfflineGrantStore
}

// NewMockOfflineGrantStore creates a new mock instance.
func NewMockOfflineGrantStore(ctrl *gomock.Controller) *MockOfflineGrantStore {
	mock := &MockOfflineGrantStore{ctrl: ctrl}
	mock.recorder = &MockOfflineGrantStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOfflineGrantStore) EXPECT() *MockOfflineGrantStoreMockRecorder {
	return m.recorder
}

// AddOfflineGrantRefreshToken mocks base method.
func (m *MockOfflineGrantStore) AddOfflineGrantRefreshToken(ctx context.Context, grantID string, expireAt time.Time, tokenHash, clientID string, scopes []string, authorizationID, dpopJKT string) (*OfflineGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOfflineGrantRefreshToken", ctx, grantID, expireAt, tokenHash, clientID, scopes, authorizationID, dpopJKT)
	ret0, _ := ret[0].(*OfflineGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddOfflineGrantRefreshToken indicates an expected call of AddOfflineGrantRefreshToken.
func (mr *MockOfflineGrantStoreMockRecorder) AddOfflineGrantRefreshToken(ctx, grantID, expireAt, tokenHash, clientID, scopes, authorizationID, dpopJKT interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOfflineGrantRefreshToken", reflect.TypeOf((*MockOfflineGrantStore)(nil).AddOfflineGrantRefreshToken), ctx, grantID, expireAt, tokenHash, clientID, scopes, authorizationID, dpopJKT)
}

// AddOfflineGrantSAMLServiceProviderParticipant mocks base method.
func (m *MockOfflineGrantStore) AddOfflineGrantSAMLServiceProviderParticipant(ctx context.Context, grantID, newServiceProviderID string, expireAt time.Time) (*OfflineGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOfflineGrantSAMLServiceProviderParticipant", ctx, grantID, newServiceProviderID, expireAt)
	ret0, _ := ret[0].(*OfflineGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddOfflineGrantSAMLServiceProviderParticipant indicates an expected call of AddOfflineGrantSAMLServiceProviderParticipant.
func (mr *MockOfflineGrantStoreMockRecorder) AddOfflineGrantSAMLServiceProviderParticipant(ctx, grantID, newServiceProviderID, expireAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOfflineGrantSAMLServiceProviderParticipant", reflect.TypeOf((*MockOfflineGrantStore)(nil).AddOfflineGrantSAMLServiceProviderParticipant), ctx, grantID, newServiceProviderID, expireAt)
}

// CleanUpForDeletingUserID mocks base method.
func (m *MockOfflineGrantStore) CleanUpForDeletingUserID(ctx context.Context, userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanUpForDeletingUserID", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanUpForDeletingUserID indicates an expected call of CleanUpForDeletingUserID.
func (mr *MockOfflineGrantStoreMockRecorder) CleanUpForDeletingUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanUpForDeletingUserID", reflect.TypeOf((*MockOfflineGrantStore)(nil).CleanUpForDeletingUserID), ctx, userID)
}

// CreateOfflineGrant mocks base method.
func (m *MockOfflineGrantStore) CreateOfflineGrant(ctx context.Context, offlineGrant *OfflineGrant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOfflineGrant", ctx, offlineGrant)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOfflineGrant indicates an expected call of CreateOfflineGrant.
func (mr *MockOfflineGrantStoreMockRecorder) CreateOfflineGrant(ctx, offlineGrant interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOfflineGrant", reflect.TypeOf((*MockOfflineGrantStore)(nil).CreateOfflineGrant), ctx, offlineGrant)
}

// DeleteOfflineGrant mocks base method.
func (m *MockOfflineGrantStore) DeleteOfflineGrant(ctx context.Context, g *OfflineGrant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOfflineGrant", ctx, g)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOfflineGrant indicates an expected call of DeleteOfflineGrant.
func (mr *MockOfflineGrantStoreMockRecorder) DeleteOfflineGrant(ctx, g interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOfflineGrant", reflect.TypeOf((*MockOfflineGrantStore)(nil).DeleteOfflineGrant), ctx, g)
}

// GetOfflineGrantWithoutExpireAt mocks base method.
func (m *MockOfflineGrantStore) GetOfflineGrantWithoutExpireAt(ctx context.Context, id string) (*OfflineGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOfflineGrantWithoutExpireAt", ctx, id)
	ret0, _ := ret[0].(*OfflineGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOfflineGrantWithoutExpireAt indicates an expected call of GetOfflineGrantWithoutExpireAt.
func (mr *MockOfflineGrantStoreMockRecorder) GetOfflineGrantWithoutExpireAt(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOfflineGrantWithoutExpireAt", reflect.TypeOf((*MockOfflineGrantStore)(nil).GetOfflineGrantWithoutExpireAt), ctx, id)
}

// ListClientOfflineGrants mocks base method.
func (m *MockOfflineGrantStore) ListClientOfflineGrants(ctx context.Context, clientID, userID string) ([]*OfflineGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListClientOfflineGrants", ctx, clientID, userID)
	ret0, _ := ret[0].([]*OfflineGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClientOfflineGrants indicates an expected call of ListClientOfflineGrants.
func (mr *MockOfflineGrantStoreMockRecorder) ListClientOfflineGrants(ctx, clientID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClientOfflineGrants", reflect.TypeOf((*MockOfflineGrantStore)(nil).ListClientOfflineGrants), ctx, clientID, userID)
}

// ListOfflineGrants mocks base method.
func (m *MockOfflineGrantStore) ListOfflineGrants(ctx context.Context, userID string) ([]*OfflineGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOfflineGrants", ctx, userID)
	ret0, _ := ret[0].([]*OfflineGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOfflineGrants indicates an expected call of ListOfflineGrants.
func (mr *MockOfflineGrantStoreMockRecorder) ListOfflineGrants(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOfflineGrants", reflect.TypeOf((*MockOfflineGrantStore)(nil).ListOfflineGrants), ctx, userID)
}

// RemoveOfflineGrantRefreshTokens mocks base method.
func (m *MockOfflineGrantStore) RemoveOfflineGrantRefreshTokens(ctx context.Context, grantID string, tokenHashes []string, expireAt time.Time) (*OfflineGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveOfflineGrantRefreshTokens", ctx, grantID, tokenHashes, expireAt)
	ret0, _ := ret[0].(*OfflineGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveOfflineGrantRefreshTokens indicates an expected call of RemoveOfflineGrantRefreshTokens.
func (mr *MockOfflineGrantStoreMockRecorder) RemoveOfflineGrantRefreshTokens(ctx, grantID, tokenHashes, expireAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveOfflineGrantRefreshTokens", reflect.TypeOf((*MockOfflineGrantStore)(nil).RemoveOfflineGrantRefreshTokens), ctx, grantID, tokenHashes, expireAt)
}

// UpdateOfflineGrantApp2AppDeviceKey mocks base method.
func (m *MockOfflineGrantStore) UpdateOfflineGrantApp2AppDeviceKey(ctx context.Context, id, newKey string, expireAt time.Time) (*OfflineGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOfflineGrantApp2AppDeviceKey", ctx, id, newKey, expireAt)
	ret0, _ := ret[0].(*OfflineGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOfflineGrantApp2AppDeviceKey indicates an expected call of UpdateOfflineGrantApp2AppDeviceKey.
func (mr *MockOfflineGrantStoreMockRecorder) UpdateOfflineGrantApp2AppDeviceKey(ctx, id, newKey, expireAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOfflineGrantApp2AppDeviceKey", reflect.TypeOf((*MockOfflineGrantStore)(nil).UpdateOfflineGrantApp2AppDeviceKey), ctx, id, newKey, expireAt)
}

// UpdateOfflineGrantAuthenticatedAt mocks base method.
func (m *MockOfflineGrantStore) UpdateOfflineGrantAuthenticatedAt(ctx context.Context, id string, authenticatedAt, expireAt time.Time) (*OfflineGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOfflineGrantAuthenticatedAt", ctx, id, authenticatedAt, expireAt)
	ret0, _ := ret[0].(*OfflineGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOfflineGrantAuthenticatedAt indicates an expected call of UpdateOfflineGrantAuthenticatedAt.
func (mr *MockOfflineGrantStoreMockRecorder) UpdateOfflineGrantAuthenticatedAt(ctx, id, authenticatedAt, expireAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOfflineGrantAuthenticatedAt", reflect.TypeOf((*MockOfflineGrantStore)(nil).UpdateOfflineGrantAuthenticatedAt), ctx, id, authenticatedAt, expireAt)
}

// UpdateOfflineGrantDeviceInfo mocks base method.
func (m *MockOfflineGrantStore) UpdateOfflineGrantDeviceInfo(ctx context.Context, id string, deviceInfo map[string]interface{}, expireAt time.Time) (*OfflineGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOfflineGrantDeviceInfo", ctx, id, deviceInfo, expireAt)
	ret0, _ := ret[0].(*OfflineGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOfflineGrantDeviceInfo indicates an expected call of UpdateOfflineGrantDeviceInfo.
func (mr *MockOfflineGrantStoreMockRecorder) UpdateOfflineGrantDeviceInfo(ctx, id, deviceInfo, expireAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOfflineGrantDeviceInfo", reflect.TypeOf((*MockOfflineGrantStore)(nil).UpdateOfflineGrantDeviceInfo), ctx, id, deviceInfo, expireAt)
}

// UpdateOfflineGrantDeviceSecretHash mocks base method.
func (m *MockOfflineGrantStore) UpdateOfflineGrantDeviceSecretHash(ctx context.Context, grantID, newDeviceSecretHash, dpopJKT string, expireAt time.Time) (*OfflineGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOfflineGrantDeviceSecretHash", ctx, grantID, newDeviceSecretHash, dpopJKT, expireAt)
	ret0, _ := ret[0].(*OfflineGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOfflineGrantDeviceSecretHash indicates an expected call of UpdateOfflineGrantDeviceSecretHash.
func (mr *MockOfflineGrantStoreMockRecorder) UpdateOfflineGrantDeviceSecretHash(ctx, grantID, newDeviceSecretHash, dpopJKT, expireAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOfflineGrantDeviceSecretHash", reflect.TypeOf((*MockOfflineGrantStore)(nil).UpdateOfflineGrantDeviceSecretHash), ctx, grantID, newDeviceSecretHash, dpopJKT, expireAt)
}

// UpdateOfflineGrantLastAccess mocks base method.
func (m *MockOfflineGrantStore) UpdateOfflineGrantLastAccess(ctx context.Context, id string, accessEvent access.Event, expireAt time.Time) (*OfflineGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOfflineGrantLastAccess", ctx, id, accessEvent, expireAt)
	ret0, _ := ret[0].(*OfflineGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOfflineGrantLastAccess indicates an expected call of UpdateOfflineGrantLastAccess.
func (mr *MockOfflineGrantStoreMockRecorder) UpdateOfflineGrantLastAccess(ctx, id, accessEvent, expireAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOfflineGrantLastAccess", reflect.TypeOf((*MockOfflineGrantStore)(nil).UpdateOfflineGrantLastAccess), ctx, id, accessEvent, expireAt)
}

// MockAccessGrantStore is a mock of AccessGrantStore interface.
type MockAccessGrantStore struct {
	ctrl     *gomock.Controller
	recorder *MockAccessGrantStoreMockRecorder
}

// MockAccessGrantStoreMockRecorder is the mock recorder for MockAccessGrantStore.
type MockAccessGrantStoreMockRecorder struct {
	mock *MockAccessGrantStore
}

// NewMockAccessGrantStore creates a new mock instance.
func NewMockAccessGrantStore(ctrl *gomock.Controller) *MockAccessGrantStore {
	mock := &MockAccessGrantStore{ctrl: ctrl}
	mock.recorder = &MockAccessGrantStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessGrantStore) EXPECT() *MockAccessGrantStoreMockRecorder {
	return m.recorder
}

// CreateAccessGrant mocks base method.
func (m *MockAccessGrantStore) CreateAccessGrant(ctx context.Context, g *AccessGrant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccessGrant", ctx, g)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAccessGrant indicates an expected call of CreateAccessGrant.
func (mr *MockAccessGrantStoreMockRecorder) CreateAccessGrant(ctx, g interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessGrant", reflect.TypeOf((*MockAccessGrantStore)(nil).CreateAccessGrant), ctx, g)
}

// DeleteAccessGrant mocks base method.
func (m *MockAccessGrantStore) DeleteAccessGrant(ctx context.Context, g *AccessGrant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccessGrant", ctx, g)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccessGrant indicates an expected call of DeleteAccessGrant.
func (mr *MockAccessGrantStoreMockRecorder) DeleteAccessGrant(ctx, g interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccessGrant", reflect.TypeOf((*MockAccessGrantStore)(nil).DeleteAccessGrant), ctx, g)
}

// GetAccessGrant mocks base method.
func (m *MockAccessGrantStore) GetAccessGrant(ctx context.Context, tokenHash string) (*AccessGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccessGrant", ctx, tokenHash)
	ret0, _ := ret[0].(*AccessGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessGrant indicates an expected call of GetAccessGrant.
func (mr *MockAccessGrantStoreMockRecorder) GetAccessGrant(ctx, tokenHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessGrant", reflect.TypeOf((*MockAccessGrantStore)(nil).GetAccessGrant), ctx, tokenHash)
}

// MockAppSessionStore is a mock of AppSessionStore interface.
type MockAppSessionStore struct {
	ctrl     *gomock.Controller
	recorder *MockAppSessionStoreMockRecorder
}

// MockAppSessionStoreMockRecorder is the mock recorder for MockAppSessionStore.
type MockAppSessionStoreMockRecorder struct {
	mock *MockAppSessionStore
}

// NewMockAppSessionStore creates a new mock instance.
func NewMockAppSessionStore(ctrl *gomock.Controller) *MockAppSessionStore {
	mock := &MockAppSessionStore{ctrl: ctrl}
	mock.recorder = &MockAppSessionStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppSessionStore) EXPECT() *MockAppSessionStoreMockRecorder {
	return m.recorder
}

// CreateAppSession mocks base method.
func (m *MockAppSessionStore) CreateAppSession(ctx context.Context, s *AppSession) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAppSession", ctx, s)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAppSession indicates an expected call of CreateAppSession.
func (mr *MockAppSessionStoreMockRecorder) CreateAppSession(ctx, s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAppSession", reflect.TypeOf((*MockAppSessionStore)(nil).CreateAppSession), ctx, s)
}

// DeleteAppSession mocks base method.
func (m *MockAppSessionStore) DeleteAppSession(ctx context.Context, s *AppSession) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAppSession", ctx, s)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAppSession indicates an expected call of DeleteAppSession.
func (mr *MockAppSessionStoreMockRecorder) DeleteAppSession(ctx, s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAppSession", reflect.TypeOf((*MockAppSessionStore)(nil).DeleteAppSession), ctx, s)
}

// GetAppSession mocks base method.
func (m *MockAppSessionStore) GetAppSession(ctx context.Context, tokenHash string) (*AppSession, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppSession", ctx, tokenHash)
	ret0, _ := ret[0].(*AppSession)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAppSession indicates an expected call of GetAppSession.
func (mr *MockAppSessionStoreMockRecorder) GetAppSession(ctx, tokenHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppSession", reflect.TypeOf((*MockAppSessionStore)(nil).GetAppSession), ctx, tokenHash)
}

// MockAppSessionTokenStore is a mock of AppSessionTokenStore interface.
type MockAppSessionTokenStore struct {
	ctrl     *gomock.Controller
	recorder *MockAppSessionTokenStoreMockRecorder
}

// MockAppSessionTokenStoreMockRecorder is the mock recorder for MockAppSessionTokenStore.
type MockAppSessionTokenStoreMockRecorder struct {
	mock *MockAppSessionTokenStore
}

// NewMockAppSessionTokenStore creates a new mock instance.
func NewMockAppSessionTokenStore(ctrl *gomock.Controller) *MockAppSessionTokenStore {
	mock := &MockAppSessionTokenStore{ctrl: ctrl}
	mock.recorder = &MockAppSessionTokenStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppSessionTokenStore) EXPECT() *MockAppSessionTokenStoreMockRecorder {
	return m.recorder
}

// CreateAppSessionToken mocks base method.
func (m *MockAppSessionTokenStore) CreateAppSessionToken(ctx context.Context, t *AppSessionToken) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAppSessionToken", ctx, t)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAppSessionToken indicates an expected call of CreateAppSessionToken.
func (mr *MockAppSessionTokenStoreMockRecorder) CreateAppSessionToken(ctx, t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAppSessionToken", reflect.TypeOf((*MockAppSessionTokenStore)(nil).CreateAppSessionToken), ctx, t)
}

// DeleteAppSessionToken mocks base method.
func (m *MockAppSessionTokenStore) DeleteAppSessionToken(ctx context.Context, t *AppSessionToken) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAppSessionToken", ctx, t)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAppSessionToken indicates an expected call of DeleteAppSessionToken.
func (mr *MockAppSessionTokenStoreMockRecorder) DeleteAppSessionToken(ctx, t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAppSessionToken", reflect.TypeOf((*MockAppSessionTokenStore)(nil).DeleteAppSessionToken), ctx, t)
}

// GetAppSessionToken mocks base method.
func (m *MockAppSessionTokenStore) GetAppSessionToken(ctx context.Context, tokenHash string) (*AppSessionToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppSessionToken", ctx, tokenHash)
	ret0, _ := ret[0].(*AppSessionToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAppSessionToken indicates an expected call of GetAppSessionToken.
func (mr *MockAppSessionTokenStoreMockRecorder) GetAppSessionToken(ctx, tokenHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppSessionToken", reflect.TypeOf((*MockAppSessionTokenStore)(nil).GetAppSessionToken), ctx, tokenHash)
}

// MockPreAuthenticatedURLTokenStore is a mock of PreAuthenticatedURLTokenStore interface.
type MockPreAuthenticatedURLTokenStore struct {
	ctrl     *gomock.Controller
	recorder *MockPreAuthenticatedURLTokenStoreMockRecorder
}

// MockPreAuthenticatedURLTokenStoreMockRecorder is the mock recorder for MockPreAuthenticatedURLTokenStore.
type MockPreAuthenticatedURLTokenStoreMockRecorder struct {
	mock *MockPreAuthenticatedURLTokenStore
}

// NewMockPreAuthenticatedURLTokenStore creates a new mock instance.
func NewMockPreAuthenticatedURLTokenStore(ctrl *gomock.Controller) *MockPreAuthenticatedURLTokenStore {
	mock := &MockPreAuthenticatedURLTokenStore{ctrl: ctrl}
	mock.recorder = &MockPreAuthenticatedURLTokenStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPreAuthenticatedURLTokenStore) EXPECT() *MockPreAuthenticatedURLTokenStoreMockRecorder {
	return m.recorder
}

// ConsumePreAuthenticatedURLToken mocks base method.
func (m *MockPreAuthenticatedURLTokenStore) ConsumePreAuthenticatedURLToken(ctx context.Context, tokenHash string) (*PreAuthenticatedURLToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConsumePreAuthenticatedURLToken", ctx, tokenHash)
	ret0, _ := ret[0].(*PreAuthenticatedURLToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConsumePreAuthenticatedURLToken indicates an expected call of ConsumePreAuthenticatedURLToken.
func (mr *MockPreAuthenticatedURLTokenStoreMockRecorder) ConsumePreAuthenticatedURLToken(ctx, tokenHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumePreAuthenticatedURLToken", reflect.TypeOf((*MockPreAuthenticatedURLTokenStore)(nil).ConsumePreAuthenticatedURLToken), ctx, tokenHash)
}

// CreatePreAuthenticatedURLToken mocks base method.
func (m *MockPreAuthenticatedURLTokenStore) CreatePreAuthenticatedURLToken(ctx context.Context, t *PreAuthenticatedURLToken) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePreAuthenticatedURLToken", ctx, t)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePreAuthenticatedURLToken indicates an expected call of CreatePreAuthenticatedURLToken.
func (mr *MockPreAuthenticatedURLTokenStoreMockRecorder) CreatePreAuthenticatedURLToken(ctx, t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePreAuthenticatedURLToken", reflect.TypeOf((*MockPreAuthenticatedURLTokenStore)(nil).CreatePreAuthenticatedURLToken), ctx, t)
}
