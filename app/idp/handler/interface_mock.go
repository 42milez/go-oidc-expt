// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package handler is a generated GoMock package.
package handler

import (
	context "context"
	http "net/http"
	reflect "reflect"

	ent "github.com/42milez/go-oidc-server/app/idp/ent/ent"
	model "github.com/42milez/go-oidc-server/app/idp/model"
	gomock "github.com/golang/mock/gomock"
)

// MockHealthChecker is a mock of HealthChecker interface.
type MockHealthChecker struct {
	ctrl     *gomock.Controller
	recorder *MockHealthCheckerMockRecorder
}

// MockHealthCheckerMockRecorder is the mock recorder for MockHealthChecker.
type MockHealthCheckerMockRecorder struct {
	mock *MockHealthChecker
}

// NewMockHealthChecker creates a new mock instance.
func NewMockHealthChecker(ctrl *gomock.Controller) *MockHealthChecker {
	mock := &MockHealthChecker{ctrl: ctrl}
	mock.recorder = &MockHealthCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHealthChecker) EXPECT() *MockHealthCheckerMockRecorder {
	return m.recorder
}

// CheckCacheStatus mocks base method.
func (m *MockHealthChecker) CheckCacheStatus(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckCacheStatus", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckCacheStatus indicates an expected call of CheckCacheStatus.
func (mr *MockHealthCheckerMockRecorder) CheckCacheStatus(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckCacheStatus", reflect.TypeOf((*MockHealthChecker)(nil).CheckCacheStatus), ctx)
}

// CheckDBStatus mocks base method.
func (m *MockHealthChecker) CheckDBStatus(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckDBStatus", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckDBStatus indicates an expected call of CheckDBStatus.
func (mr *MockHealthCheckerMockRecorder) CheckDBStatus(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckDBStatus", reflect.TypeOf((*MockHealthChecker)(nil).CheckDBStatus), ctx)
}

// MockAuthorizer is a mock of Authorizer interface.
type MockAuthorizer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizerMockRecorder
}

// MockAuthorizerMockRecorder is the mock recorder for MockAuthorizer.
type MockAuthorizerMockRecorder struct {
	mock *MockAuthorizer
}

// NewMockAuthorizer creates a new mock instance.
func NewMockAuthorizer(ctrl *gomock.Controller) *MockAuthorizer {
	mock := &MockAuthorizer{ctrl: ctrl}
	mock.recorder = &MockAuthorizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorizer) EXPECT() *MockAuthorizerMockRecorder {
	return m.recorder
}

// Authorize mocks base method.
func (m *MockAuthorizer) Authorize(ctx context.Context, param *model.AuthorizeRequest) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorize", ctx, param)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authorize indicates an expected call of Authorize.
func (mr *MockAuthorizerMockRecorder) Authorize(ctx, param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorize", reflect.TypeOf((*MockAuthorizer)(nil).Authorize), ctx, param)
}

// MockAuthenticator is a mock of Authenticator interface.
type MockAuthenticator struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticatorMockRecorder
}

// MockAuthenticatorMockRecorder is the mock recorder for MockAuthenticator.
type MockAuthenticatorMockRecorder struct {
	mock *MockAuthenticator
}

// NewMockAuthenticator creates a new mock instance.
func NewMockAuthenticator(ctrl *gomock.Controller) *MockAuthenticator {
	mock := &MockAuthenticator{ctrl: ctrl}
	mock.recorder = &MockAuthenticatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthenticator) EXPECT() *MockAuthenticatorMockRecorder {
	return m.recorder
}

// Authenticate mocks base method.
func (m *MockAuthenticator) Authenticate(ctx context.Context, name, pw string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", ctx, name, pw)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockAuthenticatorMockRecorder) Authenticate(ctx, name, pw interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockAuthenticator)(nil).Authenticate), ctx, name, pw)
}

// MockUserCreator is a mock of UserCreator interface.
type MockUserCreator struct {
	ctrl     *gomock.Controller
	recorder *MockUserCreatorMockRecorder
}

// MockUserCreatorMockRecorder is the mock recorder for MockUserCreator.
type MockUserCreatorMockRecorder struct {
	mock *MockUserCreator
}

// NewMockUserCreator creates a new mock instance.
func NewMockUserCreator(ctrl *gomock.Controller) *MockUserCreator {
	mock := &MockUserCreator{ctrl: ctrl}
	mock.recorder = &MockUserCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserCreator) EXPECT() *MockUserCreatorMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserCreator) CreateUser(ctx context.Context, name, pw string) (*ent.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, name, pw)
	ret0, _ := ret[0].(*ent.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserCreatorMockRecorder) CreateUser(ctx, name, pw interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserCreator)(nil).CreateUser), ctx, name, pw)
}

// MockUserSelector is a mock of UserSelector interface.
type MockUserSelector struct {
	ctrl     *gomock.Controller
	recorder *MockUserSelectorMockRecorder
}

// MockUserSelectorMockRecorder is the mock recorder for MockUserSelector.
type MockUserSelectorMockRecorder struct {
	mock *MockUserSelector
}

// NewMockUserSelector creates a new mock instance.
func NewMockUserSelector(ctrl *gomock.Controller) *MockUserSelector {
	mock := &MockUserSelector{ctrl: ctrl}
	mock.recorder = &MockUserSelectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserSelector) EXPECT() *MockUserSelectorMockRecorder {
	return m.recorder
}

// SelectUser mocks base method.
func (m *MockUserSelector) SelectUser(ctx context.Context) (*ent.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectUser", ctx)
	ret0, _ := ret[0].(*ent.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectUser indicates an expected call of SelectUser.
func (mr *MockUserSelectorMockRecorder) SelectUser(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectUser", reflect.TypeOf((*MockUserSelector)(nil).SelectUser), ctx)
}

// MockSessionManager is a mock of SessionManager interface.
type MockSessionManager struct {
	ctrl     *gomock.Controller
	recorder *MockSessionManagerMockRecorder
}

// MockSessionManagerMockRecorder is the mock recorder for MockSessionManager.
type MockSessionManagerMockRecorder struct {
	mock *MockSessionManager
}

// NewMockSessionManager creates a new mock instance.
func NewMockSessionManager(ctrl *gomock.Controller) *MockSessionManager {
	mock := &MockSessionManager{ctrl: ctrl}
	mock.recorder = &MockSessionManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionManager) EXPECT() *MockSessionManagerMockRecorder {
	return m.recorder
}

// FillContext mocks base method.
func (m *MockSessionManager) FillContext(r *http.Request) (*http.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FillContext", r)
	ret0, _ := ret[0].(*http.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FillContext indicates an expected call of FillContext.
func (mr *MockSessionManagerMockRecorder) FillContext(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FillContext", reflect.TypeOf((*MockSessionManager)(nil).FillContext), r)
}
