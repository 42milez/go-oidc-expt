// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go
//
// Generated by this command:
//
//	mockgen -source=interface.go -destination=interface_mock.go -package=service
//

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	entity "github.com/42milez/go-oidc-server/app/idp/entity"
	typedef "github.com/42milez/go-oidc-server/app/pkg/typedef"
	gomock "go.uber.org/mock/gomock"
)

// MockCachePingSender is a mock of CachePingSender interface.
type MockCachePingSender struct {
	ctrl     *gomock.Controller
	recorder *MockCachePingSenderMockRecorder
}

// MockCachePingSenderMockRecorder is the mock recorder for MockCachePingSender.
type MockCachePingSenderMockRecorder struct {
	mock *MockCachePingSender
}

// NewMockCachePingSender creates a new mock instance.
func NewMockCachePingSender(ctrl *gomock.Controller) *MockCachePingSender {
	mock := &MockCachePingSender{ctrl: ctrl}
	mock.recorder = &MockCachePingSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCachePingSender) EXPECT() *MockCachePingSenderMockRecorder {
	return m.recorder
}

// PingCache mocks base method.
func (m *MockCachePingSender) PingCache(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PingCache", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// PingCache indicates an expected call of PingCache.
func (mr *MockCachePingSenderMockRecorder) PingCache(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PingCache", reflect.TypeOf((*MockCachePingSender)(nil).PingCache), ctx)
}

// MockDatabasePingSender is a mock of DatabasePingSender interface.
type MockDatabasePingSender struct {
	ctrl     *gomock.Controller
	recorder *MockDatabasePingSenderMockRecorder
}

// MockDatabasePingSenderMockRecorder is the mock recorder for MockDatabasePingSender.
type MockDatabasePingSenderMockRecorder struct {
	mock *MockDatabasePingSender
}

// NewMockDatabasePingSender creates a new mock instance.
func NewMockDatabasePingSender(ctrl *gomock.Controller) *MockDatabasePingSender {
	mock := &MockDatabasePingSender{ctrl: ctrl}
	mock.recorder = &MockDatabasePingSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabasePingSender) EXPECT() *MockDatabasePingSenderMockRecorder {
	return m.recorder
}

// PingDatabase mocks base method.
func (m *MockDatabasePingSender) PingDatabase(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PingDatabase", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// PingDatabase indicates an expected call of PingDatabase.
func (mr *MockDatabasePingSenderMockRecorder) PingDatabase(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PingDatabase", reflect.TypeOf((*MockDatabasePingSender)(nil).PingDatabase), ctx)
}

// MockPingSender is a mock of PingSender interface.
type MockPingSender struct {
	ctrl     *gomock.Controller
	recorder *MockPingSenderMockRecorder
}

// MockPingSenderMockRecorder is the mock recorder for MockPingSender.
type MockPingSenderMockRecorder struct {
	mock *MockPingSender
}

// NewMockPingSender creates a new mock instance.
func NewMockPingSender(ctrl *gomock.Controller) *MockPingSender {
	mock := &MockPingSender{ctrl: ctrl}
	mock.recorder = &MockPingSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPingSender) EXPECT() *MockPingSenderMockRecorder {
	return m.recorder
}

// PingCache mocks base method.
func (m *MockPingSender) PingCache(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PingCache", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// PingCache indicates an expected call of PingCache.
func (mr *MockPingSenderMockRecorder) PingCache(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PingCache", reflect.TypeOf((*MockPingSender)(nil).PingCache), ctx)
}

// PingDatabase mocks base method.
func (m *MockPingSender) PingDatabase(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PingDatabase", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// PingDatabase indicates an expected call of PingDatabase.
func (mr *MockPingSenderMockRecorder) PingDatabase(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PingDatabase", reflect.TypeOf((*MockPingSender)(nil).PingDatabase), ctx)
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
func (m *MockUserCreator) CreateUser(ctx context.Context, name, pw string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, name, pw)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserCreatorMockRecorder) CreateUser(ctx, name, pw any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserCreator)(nil).CreateUser), ctx, name, pw)
}

// MockConsentReader is a mock of ConsentReader interface.
type MockConsentReader struct {
	ctrl     *gomock.Controller
	recorder *MockConsentReaderMockRecorder
}

// MockConsentReaderMockRecorder is the mock recorder for MockConsentReader.
type MockConsentReaderMockRecorder struct {
	mock *MockConsentReader
}

// NewMockConsentReader creates a new mock instance.
func NewMockConsentReader(ctrl *gomock.Controller) *MockConsentReader {
	mock := &MockConsentReader{ctrl: ctrl}
	mock.recorder = &MockConsentReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConsentReader) EXPECT() *MockConsentReaderMockRecorder {
	return m.recorder
}

// ReadConsent mocks base method.
func (m *MockConsentReader) ReadConsent(ctx context.Context, userID typedef.UserID, clientID string) (*entity.Consent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadConsent", ctx, userID, clientID)
	ret0, _ := ret[0].(*entity.Consent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadConsent indicates an expected call of ReadConsent.
func (mr *MockConsentReaderMockRecorder) ReadConsent(ctx, userID, clientID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadConsent", reflect.TypeOf((*MockConsentReader)(nil).ReadConsent), ctx, userID, clientID)
}

// MockUserReader is a mock of UserReader interface.
type MockUserReader struct {
	ctrl     *gomock.Controller
	recorder *MockUserReaderMockRecorder
}

// MockUserReaderMockRecorder is the mock recorder for MockUserReader.
type MockUserReaderMockRecorder struct {
	mock *MockUserReader
}

// NewMockUserReader creates a new mock instance.
func NewMockUserReader(ctrl *gomock.Controller) *MockUserReader {
	mock := &MockUserReader{ctrl: ctrl}
	mock.recorder = &MockUserReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserReader) EXPECT() *MockUserReaderMockRecorder {
	return m.recorder
}

// ReadUser mocks base method.
func (m *MockUserReader) ReadUser(ctx context.Context, name string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadUser", ctx, name)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadUser indicates an expected call of ReadUser.
func (mr *MockUserReaderMockRecorder) ReadUser(ctx, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadUser", reflect.TypeOf((*MockUserReader)(nil).ReadUser), ctx, name)
}

// MockUserConsentReader is a mock of UserConsentReader interface.
type MockUserConsentReader struct {
	ctrl     *gomock.Controller
	recorder *MockUserConsentReaderMockRecorder
}

// MockUserConsentReaderMockRecorder is the mock recorder for MockUserConsentReader.
type MockUserConsentReaderMockRecorder struct {
	mock *MockUserConsentReader
}

// NewMockUserConsentReader creates a new mock instance.
func NewMockUserConsentReader(ctrl *gomock.Controller) *MockUserConsentReader {
	mock := &MockUserConsentReader{ctrl: ctrl}
	mock.recorder = &MockUserConsentReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserConsentReader) EXPECT() *MockUserConsentReaderMockRecorder {
	return m.recorder
}

// ReadConsent mocks base method.
func (m *MockUserConsentReader) ReadConsent(ctx context.Context, userID typedef.UserID, clientID string) (*entity.Consent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadConsent", ctx, userID, clientID)
	ret0, _ := ret[0].(*entity.Consent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadConsent indicates an expected call of ReadConsent.
func (mr *MockUserConsentReaderMockRecorder) ReadConsent(ctx, userID, clientID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadConsent", reflect.TypeOf((*MockUserConsentReader)(nil).ReadConsent), ctx, userID, clientID)
}

// ReadUser mocks base method.
func (m *MockUserConsentReader) ReadUser(ctx context.Context, name string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadUser", ctx, name)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadUser indicates an expected call of ReadUser.
func (mr *MockUserConsentReaderMockRecorder) ReadUser(ctx, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadUser", reflect.TypeOf((*MockUserConsentReader)(nil).ReadUser), ctx, name)
}

// MockCredentialReader is a mock of CredentialReader interface.
type MockCredentialReader struct {
	ctrl     *gomock.Controller
	recorder *MockCredentialReaderMockRecorder
}

// MockCredentialReaderMockRecorder is the mock recorder for MockCredentialReader.
type MockCredentialReaderMockRecorder struct {
	mock *MockCredentialReader
}

// NewMockCredentialReader creates a new mock instance.
func NewMockCredentialReader(ctrl *gomock.Controller) *MockCredentialReader {
	mock := &MockCredentialReader{ctrl: ctrl}
	mock.recorder = &MockCredentialReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCredentialReader) EXPECT() *MockCredentialReaderMockRecorder {
	return m.recorder
}

// ReadCredential mocks base method.
func (m *MockCredentialReader) ReadCredential(ctx context.Context, clientID, clientSecret string) (*entity.RelyingParty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadCredential", ctx, clientID, clientSecret)
	ret0, _ := ret[0].(*entity.RelyingParty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadCredential indicates an expected call of ReadCredential.
func (mr *MockCredentialReaderMockRecorder) ReadCredential(ctx, clientID, clientSecret any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadCredential", reflect.TypeOf((*MockCredentialReader)(nil).ReadCredential), ctx, clientID, clientSecret)
}

// MockAuthCodeCreator is a mock of AuthCodeCreator interface.
type MockAuthCodeCreator struct {
	ctrl     *gomock.Controller
	recorder *MockAuthCodeCreatorMockRecorder
}

// MockAuthCodeCreatorMockRecorder is the mock recorder for MockAuthCodeCreator.
type MockAuthCodeCreatorMockRecorder struct {
	mock *MockAuthCodeCreator
}

// NewMockAuthCodeCreator creates a new mock instance.
func NewMockAuthCodeCreator(ctrl *gomock.Controller) *MockAuthCodeCreator {
	mock := &MockAuthCodeCreator{ctrl: ctrl}
	mock.recorder = &MockAuthCodeCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthCodeCreator) EXPECT() *MockAuthCodeCreatorMockRecorder {
	return m.recorder
}

// CreateAuthCode mocks base method.
func (m *MockAuthCodeCreator) CreateAuthCode(ctx context.Context, code, clientID string, userID typedef.UserID) (*entity.AuthCode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAuthCode", ctx, code, clientID, userID)
	ret0, _ := ret[0].(*entity.AuthCode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAuthCode indicates an expected call of CreateAuthCode.
func (mr *MockAuthCodeCreatorMockRecorder) CreateAuthCode(ctx, code, clientID, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAuthCode", reflect.TypeOf((*MockAuthCodeCreator)(nil).CreateAuthCode), ctx, code, clientID, userID)
}

// MockRedirectUrisReader is a mock of RedirectUrisReader interface.
type MockRedirectUrisReader struct {
	ctrl     *gomock.Controller
	recorder *MockRedirectUrisReaderMockRecorder
}

// MockRedirectUrisReaderMockRecorder is the mock recorder for MockRedirectUrisReader.
type MockRedirectUrisReaderMockRecorder struct {
	mock *MockRedirectUrisReader
}

// NewMockRedirectUrisReader creates a new mock instance.
func NewMockRedirectUrisReader(ctrl *gomock.Controller) *MockRedirectUrisReader {
	mock := &MockRedirectUrisReader{ctrl: ctrl}
	mock.recorder = &MockRedirectUrisReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedirectUrisReader) EXPECT() *MockRedirectUrisReaderMockRecorder {
	return m.recorder
}

// ReadRedirectUris mocks base method.
func (m *MockRedirectUrisReader) ReadRedirectUris(ctx context.Context, clientID string) ([]*entity.RedirectUri, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadRedirectUris", ctx, clientID)
	ret0, _ := ret[0].([]*entity.RedirectUri)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadRedirectUris indicates an expected call of ReadRedirectUris.
func (mr *MockRedirectUrisReaderMockRecorder) ReadRedirectUris(ctx, clientID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadRedirectUris", reflect.TypeOf((*MockRedirectUrisReader)(nil).ReadRedirectUris), ctx, clientID)
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

// CreateAuthCode mocks base method.
func (m *MockAuthorizer) CreateAuthCode(ctx context.Context, code, clientID string, userID typedef.UserID) (*entity.AuthCode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAuthCode", ctx, code, clientID, userID)
	ret0, _ := ret[0].(*entity.AuthCode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAuthCode indicates an expected call of CreateAuthCode.
func (mr *MockAuthorizerMockRecorder) CreateAuthCode(ctx, code, clientID, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAuthCode", reflect.TypeOf((*MockAuthorizer)(nil).CreateAuthCode), ctx, code, clientID, userID)
}

// ReadRedirectUris mocks base method.
func (m *MockAuthorizer) ReadRedirectUris(ctx context.Context, clientID string) ([]*entity.RedirectUri, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadRedirectUris", ctx, clientID)
	ret0, _ := ret[0].([]*entity.RedirectUri)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadRedirectUris indicates an expected call of ReadRedirectUris.
func (mr *MockAuthorizerMockRecorder) ReadRedirectUris(ctx, clientID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadRedirectUris", reflect.TypeOf((*MockAuthorizer)(nil).ReadRedirectUris), ctx, clientID)
}

// MockConsentCreator is a mock of ConsentCreator interface.
type MockConsentCreator struct {
	ctrl     *gomock.Controller
	recorder *MockConsentCreatorMockRecorder
}

// MockConsentCreatorMockRecorder is the mock recorder for MockConsentCreator.
type MockConsentCreatorMockRecorder struct {
	mock *MockConsentCreator
}

// NewMockConsentCreator creates a new mock instance.
func NewMockConsentCreator(ctrl *gomock.Controller) *MockConsentCreator {
	mock := &MockConsentCreator{ctrl: ctrl}
	mock.recorder = &MockConsentCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConsentCreator) EXPECT() *MockConsentCreatorMockRecorder {
	return m.recorder
}

// CreateConsent mocks base method.
func (m *MockConsentCreator) CreateConsent(ctx context.Context, userID typedef.UserID, clientID string) (*entity.Consent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateConsent", ctx, userID, clientID)
	ret0, _ := ret[0].(*entity.Consent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateConsent indicates an expected call of CreateConsent.
func (mr *MockConsentCreatorMockRecorder) CreateConsent(ctx, userID, clientID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateConsent", reflect.TypeOf((*MockConsentCreator)(nil).CreateConsent), ctx, userID, clientID)
}

// MockAuthCodeReader is a mock of AuthCodeReader interface.
type MockAuthCodeReader struct {
	ctrl     *gomock.Controller
	recorder *MockAuthCodeReaderMockRecorder
}

// MockAuthCodeReaderMockRecorder is the mock recorder for MockAuthCodeReader.
type MockAuthCodeReaderMockRecorder struct {
	mock *MockAuthCodeReader
}

// NewMockAuthCodeReader creates a new mock instance.
func NewMockAuthCodeReader(ctrl *gomock.Controller) *MockAuthCodeReader {
	mock := &MockAuthCodeReader{ctrl: ctrl}
	mock.recorder = &MockAuthCodeReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthCodeReader) EXPECT() *MockAuthCodeReaderMockRecorder {
	return m.recorder
}

// ReadAuthCode mocks base method.
func (m *MockAuthCodeReader) ReadAuthCode(ctx context.Context, code, clientID string) (*entity.AuthCode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAuthCode", ctx, code, clientID)
	ret0, _ := ret[0].(*entity.AuthCode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAuthCode indicates an expected call of ReadAuthCode.
func (mr *MockAuthCodeReaderMockRecorder) ReadAuthCode(ctx, code, clientID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAuthCode", reflect.TypeOf((*MockAuthCodeReader)(nil).ReadAuthCode), ctx, code, clientID)
}

// MockAuthCodeRevoker is a mock of AuthCodeRevoker interface.
type MockAuthCodeRevoker struct {
	ctrl     *gomock.Controller
	recorder *MockAuthCodeRevokerMockRecorder
}

// MockAuthCodeRevokerMockRecorder is the mock recorder for MockAuthCodeRevoker.
type MockAuthCodeRevokerMockRecorder struct {
	mock *MockAuthCodeRevoker
}

// NewMockAuthCodeRevoker creates a new mock instance.
func NewMockAuthCodeRevoker(ctrl *gomock.Controller) *MockAuthCodeRevoker {
	mock := &MockAuthCodeRevoker{ctrl: ctrl}
	mock.recorder = &MockAuthCodeRevokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthCodeRevoker) EXPECT() *MockAuthCodeRevokerMockRecorder {
	return m.recorder
}

// RevokeAuthCode mocks base method.
func (m *MockAuthCodeRevoker) RevokeAuthCode(ctx context.Context, code, clientID string) (*entity.AuthCode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeAuthCode", ctx, code, clientID)
	ret0, _ := ret[0].(*entity.AuthCode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevokeAuthCode indicates an expected call of RevokeAuthCode.
func (mr *MockAuthCodeRevokerMockRecorder) RevokeAuthCode(ctx, code, clientID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeAuthCode", reflect.TypeOf((*MockAuthCodeRevoker)(nil).RevokeAuthCode), ctx, code, clientID)
}

// MockAuthCodeReadRevoker is a mock of AuthCodeReadRevoker interface.
type MockAuthCodeReadRevoker struct {
	ctrl     *gomock.Controller
	recorder *MockAuthCodeReadRevokerMockRecorder
}

// MockAuthCodeReadRevokerMockRecorder is the mock recorder for MockAuthCodeReadRevoker.
type MockAuthCodeReadRevokerMockRecorder struct {
	mock *MockAuthCodeReadRevoker
}

// NewMockAuthCodeReadRevoker creates a new mock instance.
func NewMockAuthCodeReadRevoker(ctrl *gomock.Controller) *MockAuthCodeReadRevoker {
	mock := &MockAuthCodeReadRevoker{ctrl: ctrl}
	mock.recorder = &MockAuthCodeReadRevokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthCodeReadRevoker) EXPECT() *MockAuthCodeReadRevokerMockRecorder {
	return m.recorder
}

// ReadAuthCode mocks base method.
func (m *MockAuthCodeReadRevoker) ReadAuthCode(ctx context.Context, code, clientID string) (*entity.AuthCode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAuthCode", ctx, code, clientID)
	ret0, _ := ret[0].(*entity.AuthCode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAuthCode indicates an expected call of ReadAuthCode.
func (mr *MockAuthCodeReadRevokerMockRecorder) ReadAuthCode(ctx, code, clientID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAuthCode", reflect.TypeOf((*MockAuthCodeReadRevoker)(nil).ReadAuthCode), ctx, code, clientID)
}

// RevokeAuthCode mocks base method.
func (m *MockAuthCodeReadRevoker) RevokeAuthCode(ctx context.Context, code, clientID string) (*entity.AuthCode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeAuthCode", ctx, code, clientID)
	ret0, _ := ret[0].(*entity.AuthCode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevokeAuthCode indicates an expected call of RevokeAuthCode.
func (mr *MockAuthCodeReadRevokerMockRecorder) RevokeAuthCode(ctx, code, clientID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeAuthCode", reflect.TypeOf((*MockAuthCodeReadRevoker)(nil).RevokeAuthCode), ctx, code, clientID)
}

// MockRedirectUriReader is a mock of RedirectUriReader interface.
type MockRedirectUriReader struct {
	ctrl     *gomock.Controller
	recorder *MockRedirectUriReaderMockRecorder
}

// MockRedirectUriReaderMockRecorder is the mock recorder for MockRedirectUriReader.
type MockRedirectUriReaderMockRecorder struct {
	mock *MockRedirectUriReader
}

// NewMockRedirectUriReader creates a new mock instance.
func NewMockRedirectUriReader(ctrl *gomock.Controller) *MockRedirectUriReader {
	mock := &MockRedirectUriReader{ctrl: ctrl}
	mock.recorder = &MockRedirectUriReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedirectUriReader) EXPECT() *MockRedirectUriReaderMockRecorder {
	return m.recorder
}

// ReadRedirectUri mocks base method.
func (m *MockRedirectUriReader) ReadRedirectUri(ctx context.Context, clientID string) (*entity.RedirectUri, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadRedirectUri", ctx, clientID)
	ret0, _ := ret[0].(*entity.RedirectUri)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadRedirectUri indicates an expected call of ReadRedirectUri.
func (mr *MockRedirectUriReaderMockRecorder) ReadRedirectUri(ctx, clientID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadRedirectUri", reflect.TypeOf((*MockRedirectUriReader)(nil).ReadRedirectUri), ctx, clientID)
}
