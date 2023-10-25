// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	entity "github.com/42milez/go-oidc-server/app/entity"
	typedef "github.com/42milez/go-oidc-server/app/typedef"
	gomock "github.com/golang/mock/gomock"
)

// MockAccessTokenGenerator is a mock of AccessTokenGenerator interface.
type MockAccessTokenGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockAccessTokenGeneratorMockRecorder
}

// MockAccessTokenGeneratorMockRecorder is the mock recorder for MockAccessTokenGenerator.
type MockAccessTokenGeneratorMockRecorder struct {
	mock *MockAccessTokenGenerator
}

// NewMockAccessTokenGenerator creates a new mock instance.
func NewMockAccessTokenGenerator(ctrl *gomock.Controller) *MockAccessTokenGenerator {
	mock := &MockAccessTokenGenerator{ctrl: ctrl}
	mock.recorder = &MockAccessTokenGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessTokenGenerator) EXPECT() *MockAccessTokenGeneratorMockRecorder {
	return m.recorder
}

// GenerateAccessToken mocks base method.
func (m *MockAccessTokenGenerator) GenerateAccessToken(name string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAccessToken", name)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateAccessToken indicates an expected call of GenerateAccessToken.
func (mr *MockAccessTokenGeneratorMockRecorder) GenerateAccessToken(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAccessToken", reflect.TypeOf((*MockAccessTokenGenerator)(nil).GenerateAccessToken), name)
}

// MockRefreshTokenGenerator is a mock of RefreshTokenGenerator interface.
type MockRefreshTokenGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockRefreshTokenGeneratorMockRecorder
}

// MockRefreshTokenGeneratorMockRecorder is the mock recorder for MockRefreshTokenGenerator.
type MockRefreshTokenGeneratorMockRecorder struct {
	mock *MockRefreshTokenGenerator
}

// NewMockRefreshTokenGenerator creates a new mock instance.
func NewMockRefreshTokenGenerator(ctrl *gomock.Controller) *MockRefreshTokenGenerator {
	mock := &MockRefreshTokenGenerator{ctrl: ctrl}
	mock.recorder = &MockRefreshTokenGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRefreshTokenGenerator) EXPECT() *MockRefreshTokenGeneratorMockRecorder {
	return m.recorder
}

// GenerateRefreshToken mocks base method.
func (m *MockRefreshTokenGenerator) GenerateRefreshToken(name string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRefreshToken", name)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateRefreshToken indicates an expected call of GenerateRefreshToken.
func (mr *MockRefreshTokenGeneratorMockRecorder) GenerateRefreshToken(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRefreshToken", reflect.TypeOf((*MockRefreshTokenGenerator)(nil).GenerateRefreshToken), name)
}

// MockIdTokenGenerator is a mock of IdTokenGenerator interface.
type MockIdTokenGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockIdTokenGeneratorMockRecorder
}

// MockIdTokenGeneratorMockRecorder is the mock recorder for MockIdTokenGenerator.
type MockIdTokenGeneratorMockRecorder struct {
	mock *MockIdTokenGenerator
}

// NewMockIdTokenGenerator creates a new mock instance.
func NewMockIdTokenGenerator(ctrl *gomock.Controller) *MockIdTokenGenerator {
	mock := &MockIdTokenGenerator{ctrl: ctrl}
	mock.recorder = &MockIdTokenGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIdTokenGenerator) EXPECT() *MockIdTokenGeneratorMockRecorder {
	return m.recorder
}

// GenerateIdToken mocks base method.
func (m *MockIdTokenGenerator) GenerateIdToken(name string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateIdToken", name)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateIdToken indicates an expected call of GenerateIdToken.
func (mr *MockIdTokenGeneratorMockRecorder) GenerateIdToken(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateIdToken", reflect.TypeOf((*MockIdTokenGenerator)(nil).GenerateIdToken), name)
}

// MockTokenGenerator is a mock of TokenGenerator interface.
type MockTokenGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockTokenGeneratorMockRecorder
}

// MockTokenGeneratorMockRecorder is the mock recorder for MockTokenGenerator.
type MockTokenGeneratorMockRecorder struct {
	mock *MockTokenGenerator
}

// NewMockTokenGenerator creates a new mock instance.
func NewMockTokenGenerator(ctrl *gomock.Controller) *MockTokenGenerator {
	mock := &MockTokenGenerator{ctrl: ctrl}
	mock.recorder = &MockTokenGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenGenerator) EXPECT() *MockTokenGeneratorMockRecorder {
	return m.recorder
}

// GenerateAccessToken mocks base method.
func (m *MockTokenGenerator) GenerateAccessToken(name string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAccessToken", name)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateAccessToken indicates an expected call of GenerateAccessToken.
func (mr *MockTokenGeneratorMockRecorder) GenerateAccessToken(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAccessToken", reflect.TypeOf((*MockTokenGenerator)(nil).GenerateAccessToken), name)
}

// GenerateIdToken mocks base method.
func (m *MockTokenGenerator) GenerateIdToken(name string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateIdToken", name)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateIdToken indicates an expected call of GenerateIdToken.
func (mr *MockTokenGeneratorMockRecorder) GenerateIdToken(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateIdToken", reflect.TypeOf((*MockTokenGenerator)(nil).GenerateIdToken), name)
}

// GenerateRefreshToken mocks base method.
func (m *MockTokenGenerator) GenerateRefreshToken(name string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRefreshToken", name)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateRefreshToken indicates an expected call of GenerateRefreshToken.
func (mr *MockTokenGeneratorMockRecorder) GenerateRefreshToken(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRefreshToken", reflect.TypeOf((*MockTokenGenerator)(nil).GenerateRefreshToken), name)
}

// MockTokenValidator is a mock of TokenValidator interface.
type MockTokenValidator struct {
	ctrl     *gomock.Controller
	recorder *MockTokenValidatorMockRecorder
}

// MockTokenValidatorMockRecorder is the mock recorder for MockTokenValidator.
type MockTokenValidatorMockRecorder struct {
	mock *MockTokenValidator
}

// NewMockTokenValidator creates a new mock instance.
func NewMockTokenValidator(ctrl *gomock.Controller) *MockTokenValidator {
	mock := &MockTokenValidator{ctrl: ctrl}
	mock.recorder = &MockTokenValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenValidator) EXPECT() *MockTokenValidatorMockRecorder {
	return m.recorder
}

// Validate mocks base method.
func (m *MockTokenValidator) Validate(name *string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockTokenValidatorMockRecorder) Validate(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockTokenValidator)(nil).Validate), name)
}

// MockTokenGenerateValidator is a mock of TokenGenerateValidator interface.
type MockTokenGenerateValidator struct {
	ctrl     *gomock.Controller
	recorder *MockTokenGenerateValidatorMockRecorder
}

// MockTokenGenerateValidatorMockRecorder is the mock recorder for MockTokenGenerateValidator.
type MockTokenGenerateValidatorMockRecorder struct {
	mock *MockTokenGenerateValidator
}

// NewMockTokenGenerateValidator creates a new mock instance.
func NewMockTokenGenerateValidator(ctrl *gomock.Controller) *MockTokenGenerateValidator {
	mock := &MockTokenGenerateValidator{ctrl: ctrl}
	mock.recorder = &MockTokenGenerateValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenGenerateValidator) EXPECT() *MockTokenGenerateValidatorMockRecorder {
	return m.recorder
}

// GenerateAccessToken mocks base method.
func (m *MockTokenGenerateValidator) GenerateAccessToken(name string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAccessToken", name)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateAccessToken indicates an expected call of GenerateAccessToken.
func (mr *MockTokenGenerateValidatorMockRecorder) GenerateAccessToken(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAccessToken", reflect.TypeOf((*MockTokenGenerateValidator)(nil).GenerateAccessToken), name)
}

// GenerateIdToken mocks base method.
func (m *MockTokenGenerateValidator) GenerateIdToken(name string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateIdToken", name)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateIdToken indicates an expected call of GenerateIdToken.
func (mr *MockTokenGenerateValidatorMockRecorder) GenerateIdToken(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateIdToken", reflect.TypeOf((*MockTokenGenerateValidator)(nil).GenerateIdToken), name)
}

// GenerateRefreshToken mocks base method.
func (m *MockTokenGenerateValidator) GenerateRefreshToken(name string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRefreshToken", name)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateRefreshToken indicates an expected call of GenerateRefreshToken.
func (mr *MockTokenGenerateValidatorMockRecorder) GenerateRefreshToken(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRefreshToken", reflect.TypeOf((*MockTokenGenerateValidator)(nil).GenerateRefreshToken), name)
}

// Validate mocks base method.
func (m *MockTokenGenerateValidator) Validate(name *string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockTokenGenerateValidatorMockRecorder) Validate(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockTokenGenerateValidator)(nil).Validate), name)
}

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
func (mr *MockCachePingSenderMockRecorder) PingCache(ctx interface{}) *gomock.Call {
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
func (mr *MockDatabasePingSenderMockRecorder) PingDatabase(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PingDatabase", reflect.TypeOf((*MockDatabasePingSender)(nil).PingDatabase), ctx)
}

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

// PingCache mocks base method.
func (m *MockHealthChecker) PingCache(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PingCache", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// PingCache indicates an expected call of PingCache.
func (mr *MockHealthCheckerMockRecorder) PingCache(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PingCache", reflect.TypeOf((*MockHealthChecker)(nil).PingCache), ctx)
}

// PingDatabase mocks base method.
func (m *MockHealthChecker) PingDatabase(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PingDatabase", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// PingDatabase indicates an expected call of PingDatabase.
func (mr *MockHealthCheckerMockRecorder) PingDatabase(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PingDatabase", reflect.TypeOf((*MockHealthChecker)(nil).PingDatabase), ctx)
}

// MockContextReader is a mock of ContextReader interface.
type MockContextReader struct {
	ctrl     *gomock.Controller
	recorder *MockContextReaderMockRecorder
}

// MockContextReaderMockRecorder is the mock recorder for MockContextReader.
type MockContextReaderMockRecorder struct {
	mock *MockContextReader
}

// NewMockContextReader creates a new mock instance.
func NewMockContextReader(ctrl *gomock.Controller) *MockContextReader {
	mock := &MockContextReader{ctrl: ctrl}
	mock.recorder = &MockContextReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContextReader) EXPECT() *MockContextReaderMockRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockContextReader) Read(ctx context.Context, key any) any {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, key)
	ret0, _ := ret[0].(any)
	return ret0
}

// Read indicates an expected call of Read.
func (mr *MockContextReaderMockRecorder) Read(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockContextReader)(nil).Read), ctx, key)
}

// MockSessionReader is a mock of SessionReader interface.
type MockSessionReader struct {
	ctrl     *gomock.Controller
	recorder *MockSessionReaderMockRecorder
}

// MockSessionReaderMockRecorder is the mock recorder for MockSessionReader.
type MockSessionReaderMockRecorder struct {
	mock *MockSessionReader
}

// NewMockSessionReader creates a new mock instance.
func NewMockSessionReader(ctrl *gomock.Controller) *MockSessionReader {
	mock := &MockSessionReader{ctrl: ctrl}
	mock.recorder = &MockSessionReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionReader) EXPECT() *MockSessionReaderMockRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockSessionReader) Read(ctx context.Context, sid typedef.SessionID) (*entity.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, sid)
	ret0, _ := ret[0].(*entity.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockSessionReaderMockRecorder) Read(ctx, sid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockSessionReader)(nil).Read), ctx, sid)
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
func (mr *MockConsentReaderMockRecorder) ReadConsent(ctx, userID, clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadConsent", reflect.TypeOf((*MockConsentReader)(nil).ReadConsent), ctx, userID, clientID)
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
func (mr *MockUserCreatorMockRecorder) CreateUser(ctx, name, pw interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserCreator)(nil).CreateUser), ctx, name, pw)
}

// MockUserByNameReader is a mock of UserByNameReader interface.
type MockUserByNameReader struct {
	ctrl     *gomock.Controller
	recorder *MockUserByNameReaderMockRecorder
}

// MockUserByNameReaderMockRecorder is the mock recorder for MockUserByNameReader.
type MockUserByNameReaderMockRecorder struct {
	mock *MockUserByNameReader
}

// NewMockUserByNameReader creates a new mock instance.
func NewMockUserByNameReader(ctrl *gomock.Controller) *MockUserByNameReader {
	mock := &MockUserByNameReader{ctrl: ctrl}
	mock.recorder = &MockUserByNameReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserByNameReader) EXPECT() *MockUserByNameReaderMockRecorder {
	return m.recorder
}

// ReadUserByName mocks base method.
func (m *MockUserByNameReader) ReadUserByName(ctx context.Context, name string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadUserByName", ctx, name)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadUserByName indicates an expected call of ReadUserByName.
func (mr *MockUserByNameReaderMockRecorder) ReadUserByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadUserByName", reflect.TypeOf((*MockUserByNameReader)(nil).ReadUserByName), ctx, name)
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
func (mr *MockUserConsentReaderMockRecorder) ReadConsent(ctx, userID, clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadConsent", reflect.TypeOf((*MockUserConsentReader)(nil).ReadConsent), ctx, userID, clientID)
}

// ReadUserByName mocks base method.
func (m *MockUserConsentReader) ReadUserByName(ctx context.Context, name string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadUserByName", ctx, name)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadUserByName indicates an expected call of ReadUserByName.
func (mr *MockUserConsentReaderMockRecorder) ReadUserByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadUserByName", reflect.TypeOf((*MockUserConsentReader)(nil).ReadUserByName), ctx, name)
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
func (mr *MockCredentialReaderMockRecorder) ReadCredential(ctx, clientID, clientSecret interface{}) *gomock.Call {
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
func (mr *MockAuthCodeCreatorMockRecorder) CreateAuthCode(ctx, code, clientID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAuthCode", reflect.TypeOf((*MockAuthCodeCreator)(nil).CreateAuthCode), ctx, code, clientID, userID)
}

// MockRedirectUriByRelyingPartyIDReader is a mock of RedirectUriByRelyingPartyIDReader interface.
type MockRedirectUriByRelyingPartyIDReader struct {
	ctrl     *gomock.Controller
	recorder *MockRedirectUriByRelyingPartyIDReaderMockRecorder
}

// MockRedirectUriByRelyingPartyIDReaderMockRecorder is the mock recorder for MockRedirectUriByRelyingPartyIDReader.
type MockRedirectUriByRelyingPartyIDReaderMockRecorder struct {
	mock *MockRedirectUriByRelyingPartyIDReader
}

// NewMockRedirectUriByRelyingPartyIDReader creates a new mock instance.
func NewMockRedirectUriByRelyingPartyIDReader(ctrl *gomock.Controller) *MockRedirectUriByRelyingPartyIDReader {
	mock := &MockRedirectUriByRelyingPartyIDReader{ctrl: ctrl}
	mock.recorder = &MockRedirectUriByRelyingPartyIDReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedirectUriByRelyingPartyIDReader) EXPECT() *MockRedirectUriByRelyingPartyIDReaderMockRecorder {
	return m.recorder
}

// ReadRedirectUriByClientID mocks base method.
func (m *MockRedirectUriByRelyingPartyIDReader) ReadRedirectUriByClientID(ctx context.Context, clientID string) ([]*entity.RedirectUri, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadRedirectUriByClientID", ctx, clientID)
	ret0, _ := ret[0].([]*entity.RedirectUri)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadRedirectUriByClientID indicates an expected call of ReadRedirectUriByClientID.
func (mr *MockRedirectUriByRelyingPartyIDReaderMockRecorder) ReadRedirectUriByClientID(ctx, clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadRedirectUriByClientID", reflect.TypeOf((*MockRedirectUriByRelyingPartyIDReader)(nil).ReadRedirectUriByClientID), ctx, clientID)
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
func (mr *MockAuthorizerMockRecorder) CreateAuthCode(ctx, code, clientID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAuthCode", reflect.TypeOf((*MockAuthorizer)(nil).CreateAuthCode), ctx, code, clientID, userID)
}

// ReadRedirectUriByClientID mocks base method.
func (m *MockAuthorizer) ReadRedirectUriByClientID(ctx context.Context, clientID string) ([]*entity.RedirectUri, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadRedirectUriByClientID", ctx, clientID)
	ret0, _ := ret[0].([]*entity.RedirectUri)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadRedirectUriByClientID indicates an expected call of ReadRedirectUriByClientID.
func (mr *MockAuthorizerMockRecorder) ReadRedirectUriByClientID(ctx, clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadRedirectUriByClientID", reflect.TypeOf((*MockAuthorizer)(nil).ReadRedirectUriByClientID), ctx, clientID)
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
func (mr *MockConsentCreatorMockRecorder) CreateConsent(ctx, userID, clientID interface{}) *gomock.Call {
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
func (m *MockAuthCodeReader) ReadAuthCode(ctx context.Context, code, clientId string) (*entity.AuthCode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAuthCode", ctx, code, clientId)
	ret0, _ := ret[0].(*entity.AuthCode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAuthCode indicates an expected call of ReadAuthCode.
func (mr *MockAuthCodeReaderMockRecorder) ReadAuthCode(ctx, code, clientId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAuthCode", reflect.TypeOf((*MockAuthCodeReader)(nil).ReadAuthCode), ctx, code, clientId)
}

// MockAuthCodeMarker is a mock of AuthCodeMarker interface.
type MockAuthCodeMarker struct {
	ctrl     *gomock.Controller
	recorder *MockAuthCodeMarkerMockRecorder
}

// MockAuthCodeMarkerMockRecorder is the mock recorder for MockAuthCodeMarker.
type MockAuthCodeMarkerMockRecorder struct {
	mock *MockAuthCodeMarker
}

// NewMockAuthCodeMarker creates a new mock instance.
func NewMockAuthCodeMarker(ctrl *gomock.Controller) *MockAuthCodeMarker {
	mock := &MockAuthCodeMarker{ctrl: ctrl}
	mock.recorder = &MockAuthCodeMarkerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthCodeMarker) EXPECT() *MockAuthCodeMarkerMockRecorder {
	return m.recorder
}

// MarkAuthCodeUsed mocks base method.
func (m *MockAuthCodeMarker) MarkAuthCodeUsed(ctx context.Context, code, clientId string) (*entity.AuthCode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkAuthCodeUsed", ctx, code, clientId)
	ret0, _ := ret[0].(*entity.AuthCode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarkAuthCodeUsed indicates an expected call of MarkAuthCodeUsed.
func (mr *MockAuthCodeMarkerMockRecorder) MarkAuthCodeUsed(ctx, code, clientId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkAuthCodeUsed", reflect.TypeOf((*MockAuthCodeMarker)(nil).MarkAuthCodeUsed), ctx, code, clientId)
}

// MockAuthCodeReadMarker is a mock of AuthCodeReadMarker interface.
type MockAuthCodeReadMarker struct {
	ctrl     *gomock.Controller
	recorder *MockAuthCodeReadMarkerMockRecorder
}

// MockAuthCodeReadMarkerMockRecorder is the mock recorder for MockAuthCodeReadMarker.
type MockAuthCodeReadMarkerMockRecorder struct {
	mock *MockAuthCodeReadMarker
}

// NewMockAuthCodeReadMarker creates a new mock instance.
func NewMockAuthCodeReadMarker(ctrl *gomock.Controller) *MockAuthCodeReadMarker {
	mock := &MockAuthCodeReadMarker{ctrl: ctrl}
	mock.recorder = &MockAuthCodeReadMarkerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthCodeReadMarker) EXPECT() *MockAuthCodeReadMarkerMockRecorder {
	return m.recorder
}

// MarkAuthCodeUsed mocks base method.
func (m *MockAuthCodeReadMarker) MarkAuthCodeUsed(ctx context.Context, code, clientId string) (*entity.AuthCode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkAuthCodeUsed", ctx, code, clientId)
	ret0, _ := ret[0].(*entity.AuthCode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarkAuthCodeUsed indicates an expected call of MarkAuthCodeUsed.
func (mr *MockAuthCodeReadMarkerMockRecorder) MarkAuthCodeUsed(ctx, code, clientId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkAuthCodeUsed", reflect.TypeOf((*MockAuthCodeReadMarker)(nil).MarkAuthCodeUsed), ctx, code, clientId)
}

// ReadAuthCode mocks base method.
func (m *MockAuthCodeReadMarker) ReadAuthCode(ctx context.Context, code, clientId string) (*entity.AuthCode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAuthCode", ctx, code, clientId)
	ret0, _ := ret[0].(*entity.AuthCode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAuthCode indicates an expected call of ReadAuthCode.
func (mr *MockAuthCodeReadMarkerMockRecorder) ReadAuthCode(ctx, code, clientId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAuthCode", reflect.TypeOf((*MockAuthCodeReadMarker)(nil).ReadAuthCode), ctx, code, clientId)
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
func (m *MockRedirectUriReader) ReadRedirectUri(ctx context.Context, clientId string) (*entity.RedirectUri, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadRedirectUri", ctx, clientId)
	ret0, _ := ret[0].(*entity.RedirectUri)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadRedirectUri indicates an expected call of ReadRedirectUri.
func (mr *MockRedirectUriReaderMockRecorder) ReadRedirectUri(ctx, clientId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadRedirectUri", reflect.TypeOf((*MockRedirectUriReader)(nil).ReadRedirectUri), ctx, clientId)
}
