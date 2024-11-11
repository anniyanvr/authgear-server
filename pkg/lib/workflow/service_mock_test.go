// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package workflow is a generated GoMock package.
package workflow

import (
	context "context"
	reflect "reflect"

	authenticationinfo "github.com/authgear/authgear-server/pkg/lib/authn/authenticationinfo"
	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateSession mocks base method.
func (m *MockStore) CreateSession(ctx context.Context, session *Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", ctx, session)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockStoreMockRecorder) CreateSession(ctx, session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockStore)(nil).CreateSession), ctx, session)
}

// CreateWorkflow mocks base method.
func (m *MockStore) CreateWorkflow(ctx context.Context, workflow *Workflow) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWorkflow", ctx, workflow)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateWorkflow indicates an expected call of CreateWorkflow.
func (mr *MockStoreMockRecorder) CreateWorkflow(ctx, workflow interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWorkflow", reflect.TypeOf((*MockStore)(nil).CreateWorkflow), ctx, workflow)
}

// DeleteSession mocks base method.
func (m *MockStore) DeleteSession(ctx context.Context, session *Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", ctx, session)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockStoreMockRecorder) DeleteSession(ctx, session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockStore)(nil).DeleteSession), ctx, session)
}

// DeleteWorkflow mocks base method.
func (m *MockStore) DeleteWorkflow(ctx context.Context, workflow *Workflow) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWorkflow", ctx, workflow)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWorkflow indicates an expected call of DeleteWorkflow.
func (mr *MockStoreMockRecorder) DeleteWorkflow(ctx, workflow interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWorkflow", reflect.TypeOf((*MockStore)(nil).DeleteWorkflow), ctx, workflow)
}

// GetSession mocks base method.
func (m *MockStore) GetSession(ctx context.Context, workflowID string) (*Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", ctx, workflowID)
	ret0, _ := ret[0].(*Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockStoreMockRecorder) GetSession(ctx, workflowID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockStore)(nil).GetSession), ctx, workflowID)
}

// GetWorkflowByInstanceID mocks base method.
func (m *MockStore) GetWorkflowByInstanceID(ctx context.Context, instanceID string) (*Workflow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkflowByInstanceID", ctx, instanceID)
	ret0, _ := ret[0].(*Workflow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkflowByInstanceID indicates an expected call of GetWorkflowByInstanceID.
func (mr *MockStoreMockRecorder) GetWorkflowByInstanceID(ctx, instanceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflowByInstanceID", reflect.TypeOf((*MockStore)(nil).GetWorkflowByInstanceID), ctx, instanceID)
}

// MockServiceDatabase is a mock of ServiceDatabase interface.
type MockServiceDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockServiceDatabaseMockRecorder
}

// MockServiceDatabaseMockRecorder is the mock recorder for MockServiceDatabase.
type MockServiceDatabaseMockRecorder struct {
	mock *MockServiceDatabase
}

// NewMockServiceDatabase creates a new mock instance.
func NewMockServiceDatabase(ctrl *gomock.Controller) *MockServiceDatabase {
	mock := &MockServiceDatabase{ctrl: ctrl}
	mock.recorder = &MockServiceDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceDatabase) EXPECT() *MockServiceDatabaseMockRecorder {
	return m.recorder
}

// ReadOnly mocks base method.
func (m *MockServiceDatabase) ReadOnly(ctx context.Context, do func(context.Context) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadOnly", ctx, do)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadOnly indicates an expected call of ReadOnly.
func (mr *MockServiceDatabaseMockRecorder) ReadOnly(ctx, do interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadOnly", reflect.TypeOf((*MockServiceDatabase)(nil).ReadOnly), ctx, do)
}

// WithTx mocks base method.
func (m *MockServiceDatabase) WithTx(ctx context.Context, do func(context.Context) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTx", ctx, do)
	ret0, _ := ret[0].(error)
	return ret0
}

// WithTx indicates an expected call of WithTx.
func (mr *MockServiceDatabaseMockRecorder) WithTx(ctx, do interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTx", reflect.TypeOf((*MockServiceDatabase)(nil).WithTx), ctx, do)
}

// MockServiceUIInfoResolver is a mock of ServiceUIInfoResolver interface.
type MockServiceUIInfoResolver struct {
	ctrl     *gomock.Controller
	recorder *MockServiceUIInfoResolverMockRecorder
}

// MockServiceUIInfoResolverMockRecorder is the mock recorder for MockServiceUIInfoResolver.
type MockServiceUIInfoResolverMockRecorder struct {
	mock *MockServiceUIInfoResolver
}

// NewMockServiceUIInfoResolver creates a new mock instance.
func NewMockServiceUIInfoResolver(ctrl *gomock.Controller) *MockServiceUIInfoResolver {
	mock := &MockServiceUIInfoResolver{ctrl: ctrl}
	mock.recorder = &MockServiceUIInfoResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceUIInfoResolver) EXPECT() *MockServiceUIInfoResolverMockRecorder {
	return m.recorder
}

// SetAuthenticationInfoInQuery mocks base method.
func (m *MockServiceUIInfoResolver) SetAuthenticationInfoInQuery(redirectURI string, e *authenticationinfo.Entry) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAuthenticationInfoInQuery", redirectURI, e)
	ret0, _ := ret[0].(string)
	return ret0
}

// SetAuthenticationInfoInQuery indicates an expected call of SetAuthenticationInfoInQuery.
func (mr *MockServiceUIInfoResolverMockRecorder) SetAuthenticationInfoInQuery(redirectURI, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAuthenticationInfoInQuery", reflect.TypeOf((*MockServiceUIInfoResolver)(nil).SetAuthenticationInfoInQuery), redirectURI, e)
}
