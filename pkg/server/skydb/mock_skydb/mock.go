// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/skygeario/skygear-server/pkg/server/skydb (interfaces: Conn,Database)

package mock_skydb

import (
	gomock "github.com/golang/mock/gomock"
	skydb "github.com/skygeario/skygear-server/pkg/server/skydb"
	time "time"
)

// Mock of Conn interface
type MockConn struct {
	ctrl     *gomock.Controller
	recorder *_MockConnRecorder
}

// Recorder for MockConn (not exported)
type _MockConnRecorder struct {
	mock *MockConn
}

func NewMockConn(ctrl *gomock.Controller) *MockConn {
	mock := &MockConn{ctrl: ctrl}
	mock.recorder = &_MockConnRecorder{mock}
	return mock
}

func (_m *MockConn) EXPECT() *_MockConnRecorder {
	return _m.recorder
}

func (_m *MockConn) AddRelation(_param0 string, _param1 string, _param2 string) error {
	ret := _m.ctrl.Call(_m, "AddRelation", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConnRecorder) AddRelation(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddRelation", arg0, arg1, arg2)
}

func (_m *MockConn) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConnRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockConn) CreateUser(_param0 *skydb.UserInfo) error {
	ret := _m.ctrl.Call(_m, "CreateUser", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConnRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateUser", arg0)
}

func (_m *MockConn) DeleteDevice(_param0 string) error {
	ret := _m.ctrl.Call(_m, "DeleteDevice", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConnRecorder) DeleteDevice(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteDevice", arg0)
}

func (_m *MockConn) DeleteDevicesByToken(_param0 string, _param1 time.Time) error {
	ret := _m.ctrl.Call(_m, "DeleteDevicesByToken", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConnRecorder) DeleteDevicesByToken(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteDevicesByToken", arg0, arg1)
}

func (_m *MockConn) DeleteEmptyDevicesByTime(_param0 time.Time) error {
	ret := _m.ctrl.Call(_m, "DeleteEmptyDevicesByTime", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConnRecorder) DeleteEmptyDevicesByTime(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteEmptyDevicesByTime", arg0)
}

func (_m *MockConn) DeleteUser(_param0 string) error {
	ret := _m.ctrl.Call(_m, "DeleteUser", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConnRecorder) DeleteUser(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteUser", arg0)
}

func (_m *MockConn) GetAdminRoles() ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetAdminRoles")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnRecorder) GetAdminRoles() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAdminRoles")
}

func (_m *MockConn) GetAsset(_param0 string, _param1 *skydb.Asset) error {
	ret := _m.ctrl.Call(_m, "GetAsset", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConnRecorder) GetAsset(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAsset", arg0, arg1)
}

func (_m *MockConn) GetAssets(_param0 []string) ([]skydb.Asset, error) {
	ret := _m.ctrl.Call(_m, "GetAssets", _param0)
	ret0, _ := ret[0].([]skydb.Asset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnRecorder) GetAssets(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAssets", arg0, arg1)
}

func (_m *MockConn) GetDefaultRoles() ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetDefaultRoles")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnRecorder) GetDefaultRoles() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetDefaultRoles")
}

func (_m *MockConn) GetDevice(_param0 string, _param1 *skydb.Device) error {
	ret := _m.ctrl.Call(_m, "GetDevice", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConnRecorder) GetDevice(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetDevice", arg0, arg1)
}

func (_m *MockConn) GetRecordAccess(_param0 string) (skydb.RecordACL, error) {
	ret := _m.ctrl.Call(_m, "GetRecordAccess", _param0)
	ret0, _ := ret[0].(skydb.RecordACL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_m *MockConn) GetRecordDefaultAccess(_param0 string) (skydb.RecordACL, error) {
	ret := _m.ctrl.Call(_m, "GetRecordDefaultAccess", _param0)
	ret0, _ := ret[0].(skydb.RecordACL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnRecorder) GetRecordAccess(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRecordAccess", arg0)
}

func (_m *MockConn) GetUser(_param0 string, _param1 *skydb.UserInfo) error {
	ret := _m.ctrl.Call(_m, "GetUser", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConnRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUser", arg0, arg1)
}

func (_m *MockConn) GetUserByPrincipalID(_param0 string, _param1 *skydb.UserInfo) error {
	ret := _m.ctrl.Call(_m, "GetUserByPrincipalID", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConnRecorder) GetUserByPrincipalID(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUserByPrincipalID", arg0, arg1)
}

func (_m *MockConn) GetUserByUsernameEmail(_param0 string, _param1 string, _param2 *skydb.UserInfo) error {
	ret := _m.ctrl.Call(_m, "GetUserByUsernameEmail", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConnRecorder) GetUserByUsernameEmail(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUserByUsernameEmail", arg0, arg1, arg2)
}

func (_m *MockConn) PrivateDB(_param0 string) skydb.Database {
	ret := _m.ctrl.Call(_m, "PrivateDB", _param0)
	ret0, _ := ret[0].(skydb.Database)
	return ret0
}

func (_mr *_MockConnRecorder) PrivateDB(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PrivateDB", arg0)
}

func (_m *MockConn) PublicDB() skydb.Database {
	ret := _m.ctrl.Call(_m, "PublicDB")
	ret0, _ := ret[0].(skydb.Database)
	return ret0
}

func (_mr *_MockConnRecorder) PublicDB() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PublicDB")
}

func (_m *MockConn) QueryDevicesByUser(_param0 string) ([]skydb.Device, error) {
	ret := _m.ctrl.Call(_m, "QueryDevicesByUser", _param0)
	ret0, _ := ret[0].([]skydb.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnRecorder) QueryDevicesByUser(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "QueryDevicesByUser", arg0)
}

func (_m *MockConn) QueryDevicesByUserAndTopic(_param0 string, _param1 string) ([]skydb.Device, error) {
	ret := _m.ctrl.Call(_m, "QueryDevicesByUserAndTopic", _param0, _param1)
	ret0, _ := ret[0].([]skydb.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnRecorder) QueryDevicesByUserAndTopic(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "QueryDevicesByUserAndTopic", arg0, arg1)
}

func (_m *MockConn) QueryRelation(_param0 string, _param1 string, _param2 string, _param3 skydb.QueryConfig) []skydb.UserInfo {
	ret := _m.ctrl.Call(_m, "QueryRelation", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].([]skydb.UserInfo)
	return ret0
}

func (_mr *_MockConnRecorder) QueryRelation(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "QueryRelation", arg0, arg1, arg2, arg3)
}

func (_m *MockConn) QueryRelationCount(_param0 string, _param1 string, _param2 string) (uint64, error) {
	ret := _m.ctrl.Call(_m, "QueryRelationCount", _param0, _param1, _param2)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnRecorder) QueryRelationCount(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "QueryRelationCount", arg0, arg1, arg2)
}

func (_m *MockConn) QueryUser(_param0 []string, _param1 []string) ([]skydb.UserInfo, error) {
	ret := _m.ctrl.Call(_m, "QueryUser", _param0, _param1)
	ret0, _ := ret[0].([]skydb.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnRecorder) QueryUser(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "QueryUser", arg0, arg1)
}

func (_m *MockConn) RemoveRelation(_param0 string, _param1 string, _param2 string) error {
	ret := _m.ctrl.Call(_m, "RemoveRelation", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConnRecorder) RemoveRelation(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveRelation", arg0, arg1, arg2)
}

func (_m *MockConn) SaveAsset(_param0 *skydb.Asset) error {
	ret := _m.ctrl.Call(_m, "SaveAsset", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConnRecorder) SaveAsset(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveAsset", arg0)
}

func (_m *MockConn) SaveDevice(_param0 *skydb.Device) error {
	ret := _m.ctrl.Call(_m, "SaveDevice", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConnRecorder) SaveDevice(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveDevice", arg0)
}

func (_m *MockConn) SetAdminRoles(_param0 []string) error {
	ret := _m.ctrl.Call(_m, "SetAdminRoles", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConnRecorder) SetAdminRoles(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetAdminRoles", arg0)
}

func (_m *MockConn) SetDefaultRoles(_param0 []string) error {
	ret := _m.ctrl.Call(_m, "SetDefaultRoles", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConnRecorder) SetDefaultRoles(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetDefaultRoles", arg0)
}

func (_m *MockConn) SetRecordAccess(_param0 string, _param1 skydb.RecordACL) error {
	ret := _m.ctrl.Call(_m, "SetRecordAccess", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_m *MockConn) SetRecordDefaultAccess(_param0 string, _param1 skydb.RecordACL) error {
	ret := _m.ctrl.Call(_m, "SetRecordDefaultAccess", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConnRecorder) SetRecordAccess(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetRecordAccess", arg0, arg1)
}

func (_m *MockConn) Subscribe(_param0 chan skydb.RecordEvent) error {
	ret := _m.ctrl.Call(_m, "Subscribe", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConnRecorder) Subscribe(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Subscribe", arg0)
}

func (_m *MockConn) UnionDB() skydb.Database {
	ret := _m.ctrl.Call(_m, "UnionDB")
	ret0, _ := ret[0].(skydb.Database)
	return ret0
}

func (_mr *_MockConnRecorder) UnionDB() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UnionDB")
}

func (_m *MockConn) UpdateUser(_param0 *skydb.UserInfo) error {
	ret := _m.ctrl.Call(_m, "UpdateUser", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConnRecorder) UpdateUser(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateUser", arg0)
}

// Mock of Database interface
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *_MockDatabaseRecorder
}

// Recorder for MockDatabase (not exported)
type _MockDatabaseRecorder struct {
	mock *MockDatabase
}

func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &_MockDatabaseRecorder{mock}
	return mock
}

func (_m *MockDatabase) EXPECT() *_MockDatabaseRecorder {
	return _m.recorder
}

func (_m *MockDatabase) Conn() skydb.Conn {
	ret := _m.ctrl.Call(_m, "Conn")
	ret0, _ := ret[0].(skydb.Conn)
	return ret0
}

func (_mr *_MockDatabaseRecorder) Conn() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Conn")
}

func (_m *MockDatabase) DatabaseType() skydb.DatabaseType {
	ret := _m.ctrl.Call(_m, "DatabaseType")
	ret0, _ := ret[0].(skydb.DatabaseType)
	return ret0
}

func (_mr *_MockDatabaseRecorder) DatabaseType() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DatabaseType")
}

func (_m *MockDatabase) Delete(_param0 skydb.RecordID) error {
	ret := _m.ctrl.Call(_m, "Delete", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDatabaseRecorder) Delete(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete", arg0)
}

func (_m *MockDatabase) DeleteSchema(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "DeleteSchema", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDatabaseRecorder) DeleteSchema(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteSchema", arg0, arg1)
}

func (_m *MockDatabase) DeleteSubscription(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "DeleteSubscription", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDatabaseRecorder) DeleteSubscription(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteSubscription", arg0, arg1)
}

func (_m *MockDatabase) Extend(_param0 string, _param1 skydb.RecordSchema) (bool, error) {
	ret := _m.ctrl.Call(_m, "Extend", _param0, _param1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDatabaseRecorder) Extend(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Extend", arg0, arg1)
}

func (_m *MockDatabase) Get(_param0 skydb.RecordID, _param1 *skydb.Record) error {
	ret := _m.ctrl.Call(_m, "Get", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDatabaseRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0, arg1)
}

func (_m *MockDatabase) GetByIDs(_param0 []skydb.RecordID) (*skydb.Rows, error) {
	ret := _m.ctrl.Call(_m, "GetByIDs", _param0)
	ret0, _ := ret[0].(*skydb.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDatabaseRecorder) GetByIDs(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetByIDs", arg0)
}

func (_m *MockDatabase) GetMatchingSubscriptions(_param0 *skydb.Record) []skydb.Subscription {
	ret := _m.ctrl.Call(_m, "GetMatchingSubscriptions", _param0)
	ret0, _ := ret[0].([]skydb.Subscription)
	return ret0
}

func (_mr *_MockDatabaseRecorder) GetMatchingSubscriptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMatchingSubscriptions", arg0)
}

func (_m *MockDatabase) GetRecordSchemas() (map[string]skydb.RecordSchema, error) {
	ret := _m.ctrl.Call(_m, "GetRecordSchemas")
	ret0, _ := ret[0].(map[string]skydb.RecordSchema)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDatabaseRecorder) GetRecordSchemas() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRecordSchemas")
}

func (_m *MockDatabase) GetSchema(_param0 string) (skydb.RecordSchema, error) {
	ret := _m.ctrl.Call(_m, "GetSchema", _param0)
	ret0, _ := ret[0].(skydb.RecordSchema)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDatabaseRecorder) GetSchema(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSchema", arg0)
}

func (_m *MockDatabase) GetSubscription(_param0 string, _param1 string, _param2 *skydb.Subscription) error {
	ret := _m.ctrl.Call(_m, "GetSubscription", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDatabaseRecorder) GetSubscription(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSubscription", arg0, arg1, arg2)
}

func (_m *MockDatabase) GetSubscriptionsByDeviceID(_param0 string) []skydb.Subscription {
	ret := _m.ctrl.Call(_m, "GetSubscriptionsByDeviceID", _param0)
	ret0, _ := ret[0].([]skydb.Subscription)
	return ret0
}

func (_mr *_MockDatabaseRecorder) GetSubscriptionsByDeviceID(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSubscriptionsByDeviceID", arg0)
}

func (_m *MockDatabase) ID() string {
	ret := _m.ctrl.Call(_m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockDatabaseRecorder) ID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ID")
}

func (_m *MockDatabase) IsReadOnly() bool {
	ret := _m.ctrl.Call(_m, "IsReadOnly")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockDatabaseRecorder) IsReadOnly() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsReadOnly")
}

func (_m *MockDatabase) Query(_param0 *skydb.Query) (*skydb.Rows, error) {
	ret := _m.ctrl.Call(_m, "Query", _param0)
	ret0, _ := ret[0].(*skydb.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDatabaseRecorder) Query(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Query", arg0)
}

func (_m *MockDatabase) QueryCount(_param0 *skydb.Query) (uint64, error) {
	ret := _m.ctrl.Call(_m, "QueryCount", _param0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDatabaseRecorder) QueryCount(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "QueryCount", arg0)
}

func (_m *MockDatabase) RenameSchema(_param0 string, _param1 string, _param2 string) error {
	ret := _m.ctrl.Call(_m, "RenameSchema", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDatabaseRecorder) RenameSchema(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RenameSchema", arg0, arg1, arg2)
}

func (_m *MockDatabase) Save(_param0 *skydb.Record) error {
	ret := _m.ctrl.Call(_m, "Save", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDatabaseRecorder) Save(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Save", arg0)
}

func (_m *MockDatabase) SaveSubscription(_param0 *skydb.Subscription) error {
	ret := _m.ctrl.Call(_m, "SaveSubscription", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDatabaseRecorder) SaveSubscription(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveSubscription", arg0)
}

func (_m *MockDatabase) UserRecordType() string {
	ret := _m.ctrl.Call(_m, "UserRecordType")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockDatabaseRecorder) UserRecordType() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UserRecordType")
}
