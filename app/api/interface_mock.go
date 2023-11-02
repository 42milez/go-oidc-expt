// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go
//
// Generated by this command:
//
//	mockgen -source=interface.go -destination=interface_mock.go -package=api
//
// Package api is a generated GoMock package.
package api

import (
	context "context"
	reflect "reflect"

	entity "github.com/42milez/go-oidc-server/app/entity"
	typedef "github.com/42milez/go-oidc-server/app/typedef"
	gomock "go.uber.org/mock/gomock"
)

// MockCacheStatusChecker is a mock of CacheStatusChecker interface.
type MockCacheStatusChecker struct {
	ctrl     *gomock.Controller
	recorder *MockCacheStatusCheckerMockRecorder
}

// MockCacheStatusCheckerMockRecorder is the mock recorder for MockCacheStatusChecker.
type MockCacheStatusCheckerMockRecorder struct {
	mock *MockCacheStatusChecker
}

// NewMockCacheStatusChecker creates a new mock instance.
func NewMockCacheStatusChecker(ctrl *gomock.Controller) *MockCacheStatusChecker {
	mock := &MockCacheStatusChecker{ctrl: ctrl}
	mock.recorder = &MockCacheStatusCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacheStatusChecker) EXPECT() *MockCacheStatusCheckerMockRecorder {
	return m.recorder
}

// CheckCacheStatus mocks base method.
func (m *MockCacheStatusChecker) CheckCacheStatus(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckCacheStatus", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckCacheStatus indicates an expected call of CheckCacheStatus.
func (mr *MockCacheStatusCheckerMockRecorder) CheckCacheStatus(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckCacheStatus", reflect.TypeOf((*MockCacheStatusChecker)(nil).CheckCacheStatus), ctx)
}

// MockDBStatusChecker is a mock of DBStatusChecker interface.
type MockDBStatusChecker struct {
	ctrl     *gomock.Controller
	recorder *MockDBStatusCheckerMockRecorder
}

// MockDBStatusCheckerMockRecorder is the mock recorder for MockDBStatusChecker.
type MockDBStatusCheckerMockRecorder struct {
	mock *MockDBStatusChecker
}

// NewMockDBStatusChecker creates a new mock instance.
func NewMockDBStatusChecker(ctrl *gomock.Controller) *MockDBStatusChecker {
	mock := &MockDBStatusChecker{ctrl: ctrl}
	mock.recorder = &MockDBStatusCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBStatusChecker) EXPECT() *MockDBStatusCheckerMockRecorder {
	return m.recorder
}

// CheckDBStatus mocks base method.
func (m *MockDBStatusChecker) CheckDBStatus(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckDBStatus", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckDBStatus indicates an expected call of CheckDBStatus.
func (mr *MockDBStatusCheckerMockRecorder) CheckDBStatus(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckDBStatus", reflect.TypeOf((*MockDBStatusChecker)(nil).CheckDBStatus), ctx)
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

// CheckCacheStatus mocks base method.
func (m *MockHealthChecker) CheckCacheStatus(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckCacheStatus", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckCacheStatus indicates an expected call of CheckCacheStatus.
func (mr *MockHealthCheckerMockRecorder) CheckCacheStatus(ctx any) *gomock.Call {
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
func (mr *MockHealthCheckerMockRecorder) CheckDBStatus(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckDBStatus", reflect.TypeOf((*MockHealthChecker)(nil).CheckDBStatus), ctx)
}

// MockConsentVerifier is a mock of ConsentVerifier interface.
type MockConsentVerifier struct {
	ctrl     *gomock.Controller
	recorder *MockConsentVerifierMockRecorder
}

// MockConsentVerifierMockRecorder is the mock recorder for MockConsentVerifier.
type MockConsentVerifierMockRecorder struct {
	mock *MockConsentVerifier
}

// NewMockConsentVerifier creates a new mock instance.
func NewMockConsentVerifier(ctrl *gomock.Controller) *MockConsentVerifier {
	mock := &MockConsentVerifier{ctrl: ctrl}
	mock.recorder = &MockConsentVerifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConsentVerifier) EXPECT() *MockConsentVerifierMockRecorder {
	return m.recorder
}

// VerifyConsent mocks base method.
func (m *MockConsentVerifier) VerifyConsent(ctx context.Context, userID typedef.UserID, clientID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyConsent", ctx, userID, clientID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyConsent indicates an expected call of VerifyConsent.
func (mr *MockConsentVerifierMockRecorder) VerifyConsent(ctx, userID, clientID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyConsent", reflect.TypeOf((*MockConsentVerifier)(nil).VerifyConsent), ctx, userID, clientID)
}

// MockPasswordVerifier is a mock of PasswordVerifier interface.
type MockPasswordVerifier struct {
	ctrl     *gomock.Controller
	recorder *MockPasswordVerifierMockRecorder
}

// MockPasswordVerifierMockRecorder is the mock recorder for MockPasswordVerifier.
type MockPasswordVerifierMockRecorder struct {
	mock *MockPasswordVerifier
}

// NewMockPasswordVerifier creates a new mock instance.
func NewMockPasswordVerifier(ctrl *gomock.Controller) *MockPasswordVerifier {
	mock := &MockPasswordVerifier{ctrl: ctrl}
	mock.recorder = &MockPasswordVerifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPasswordVerifier) EXPECT() *MockPasswordVerifierMockRecorder {
	return m.recorder
}

// VerifyPassword mocks base method.
func (m *MockPasswordVerifier) VerifyPassword(ctx context.Context, name, pw string) (typedef.UserID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyPassword", ctx, name, pw)
	ret0, _ := ret[0].(typedef.UserID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyPassword indicates an expected call of VerifyPassword.
func (mr *MockPasswordVerifierMockRecorder) VerifyPassword(ctx, name, pw any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyPassword", reflect.TypeOf((*MockPasswordVerifier)(nil).VerifyPassword), ctx, name, pw)
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

// VerifyConsent mocks base method.
func (m *MockAuthenticator) VerifyConsent(ctx context.Context, userID typedef.UserID, clientID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyConsent", ctx, userID, clientID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyConsent indicates an expected call of VerifyConsent.
func (mr *MockAuthenticatorMockRecorder) VerifyConsent(ctx, userID, clientID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyConsent", reflect.TypeOf((*MockAuthenticator)(nil).VerifyConsent), ctx, userID, clientID)
}

// VerifyPassword mocks base method.
func (m *MockAuthenticator) VerifyPassword(ctx context.Context, name, pw string) (typedef.UserID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyPassword", ctx, name, pw)
	ret0, _ := ret[0].(typedef.UserID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyPassword indicates an expected call of VerifyPassword.
func (mr *MockAuthenticatorMockRecorder) VerifyPassword(ctx, name, pw any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyPassword", reflect.TypeOf((*MockAuthenticator)(nil).VerifyPassword), ctx, name, pw)
}

// MockUserRegisterer is a mock of UserRegisterer interface.
type MockUserRegisterer struct {
	ctrl     *gomock.Controller
	recorder *MockUserRegistererMockRecorder
}

// MockUserRegistererMockRecorder is the mock recorder for MockUserRegisterer.
type MockUserRegistererMockRecorder struct {
	mock *MockUserRegisterer
}

// NewMockUserRegisterer creates a new mock instance.
func NewMockUserRegisterer(ctrl *gomock.Controller) *MockUserRegisterer {
	mock := &MockUserRegisterer{ctrl: ctrl}
	mock.recorder = &MockUserRegistererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRegisterer) EXPECT() *MockUserRegistererMockRecorder {
	return m.recorder
}

// RegisterUser mocks base method.
func (m *MockUserRegisterer) RegisterUser(ctx context.Context, name, pw string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", ctx, name, pw)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockUserRegistererMockRecorder) RegisterUser(ctx, name, pw any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockUserRegisterer)(nil).RegisterUser), ctx, name, pw)
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
func (m *MockAuthorizer) Authorize(ctx context.Context, clientID, redirectURI, state string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorize", ctx, clientID, redirectURI, state)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Authorize indicates an expected call of Authorize.
func (mr *MockAuthorizerMockRecorder) Authorize(ctx, clientID, redirectURI, state any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorize", reflect.TypeOf((*MockAuthorizer)(nil).Authorize), ctx, clientID, redirectURI, state)
}

// MockConsentAcceptor is a mock of ConsentAcceptor interface.
type MockConsentAcceptor struct {
	ctrl     *gomock.Controller
	recorder *MockConsentAcceptorMockRecorder
}

// MockConsentAcceptorMockRecorder is the mock recorder for MockConsentAcceptor.
type MockConsentAcceptorMockRecorder struct {
	mock *MockConsentAcceptor
}

// NewMockConsentAcceptor creates a new mock instance.
func NewMockConsentAcceptor(ctrl *gomock.Controller) *MockConsentAcceptor {
	mock := &MockConsentAcceptor{ctrl: ctrl}
	mock.recorder = &MockConsentAcceptorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConsentAcceptor) EXPECT() *MockConsentAcceptorMockRecorder {
	return m.recorder
}

// AcceptConsent mocks base method.
func (m *MockConsentAcceptor) AcceptConsent(ctx context.Context, userID typedef.UserID, clientID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptConsent", ctx, userID, clientID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AcceptConsent indicates an expected call of AcceptConsent.
func (mr *MockConsentAcceptorMockRecorder) AcceptConsent(ctx, userID, clientID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptConsent", reflect.TypeOf((*MockConsentAcceptor)(nil).AcceptConsent), ctx, userID, clientID)
}

// MockCredentialValidator is a mock of CredentialValidator interface.
type MockCredentialValidator struct {
	ctrl     *gomock.Controller
	recorder *MockCredentialValidatorMockRecorder
}

// MockCredentialValidatorMockRecorder is the mock recorder for MockCredentialValidator.
type MockCredentialValidatorMockRecorder struct {
	mock *MockCredentialValidator
}

// NewMockCredentialValidator creates a new mock instance.
func NewMockCredentialValidator(ctrl *gomock.Controller) *MockCredentialValidator {
	mock := &MockCredentialValidator{ctrl: ctrl}
	mock.recorder = &MockCredentialValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCredentialValidator) EXPECT() *MockCredentialValidatorMockRecorder {
	return m.recorder
}

// ValidateCredential mocks base method.
func (m *MockCredentialValidator) ValidateCredential(ctx context.Context, clientID, clientSecret string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateCredential", ctx, clientID, clientSecret)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateCredential indicates an expected call of ValidateCredential.
func (mr *MockCredentialValidatorMockRecorder) ValidateCredential(ctx, clientID, clientSecret any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateCredential", reflect.TypeOf((*MockCredentialValidator)(nil).ValidateCredential), ctx, clientID, clientSecret)
}

// MockAuthCodeValidator is a mock of AuthCodeValidator interface.
type MockAuthCodeValidator struct {
	ctrl     *gomock.Controller
	recorder *MockAuthCodeValidatorMockRecorder
}

// MockAuthCodeValidatorMockRecorder is the mock recorder for MockAuthCodeValidator.
type MockAuthCodeValidatorMockRecorder struct {
	mock *MockAuthCodeValidator
}

// NewMockAuthCodeValidator creates a new mock instance.
func NewMockAuthCodeValidator(ctrl *gomock.Controller) *MockAuthCodeValidator {
	mock := &MockAuthCodeValidator{ctrl: ctrl}
	mock.recorder = &MockAuthCodeValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthCodeValidator) EXPECT() *MockAuthCodeValidatorMockRecorder {
	return m.recorder
}

// ValidateAuthCode mocks base method.
func (m *MockAuthCodeValidator) ValidateAuthCode(ctx context.Context, code, clientId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateAuthCode", ctx, code, clientId)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateAuthCode indicates an expected call of ValidateAuthCode.
func (mr *MockAuthCodeValidatorMockRecorder) ValidateAuthCode(ctx, code, clientId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateAuthCode", reflect.TypeOf((*MockAuthCodeValidator)(nil).ValidateAuthCode), ctx, code, clientId)
}

// MockRefreshTokenValidator is a mock of RefreshTokenValidator interface.
type MockRefreshTokenValidator struct {
	ctrl     *gomock.Controller
	recorder *MockRefreshTokenValidatorMockRecorder
}

// MockRefreshTokenValidatorMockRecorder is the mock recorder for MockRefreshTokenValidator.
type MockRefreshTokenValidatorMockRecorder struct {
	mock *MockRefreshTokenValidator
}

// NewMockRefreshTokenValidator creates a new mock instance.
func NewMockRefreshTokenValidator(ctrl *gomock.Controller) *MockRefreshTokenValidator {
	mock := &MockRefreshTokenValidator{ctrl: ctrl}
	mock.recorder = &MockRefreshTokenValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRefreshTokenValidator) EXPECT() *MockRefreshTokenValidatorMockRecorder {
	return m.recorder
}

// ValidateRefreshToken mocks base method.
func (m *MockRefreshTokenValidator) ValidateRefreshToken(ctx context.Context, token *string, clientId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateRefreshToken", ctx, token, clientId)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateRefreshToken indicates an expected call of ValidateRefreshToken.
func (mr *MockRefreshTokenValidatorMockRecorder) ValidateRefreshToken(ctx, token, clientId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateRefreshToken", reflect.TypeOf((*MockRefreshTokenValidator)(nil).ValidateRefreshToken), ctx, token, clientId)
}

// MockTokenRequestValidator is a mock of TokenRequestValidator interface.
type MockTokenRequestValidator struct {
	ctrl     *gomock.Controller
	recorder *MockTokenRequestValidatorMockRecorder
}

// MockTokenRequestValidatorMockRecorder is the mock recorder for MockTokenRequestValidator.
type MockTokenRequestValidatorMockRecorder struct {
	mock *MockTokenRequestValidator
}

// NewMockTokenRequestValidator creates a new mock instance.
func NewMockTokenRequestValidator(ctrl *gomock.Controller) *MockTokenRequestValidator {
	mock := &MockTokenRequestValidator{ctrl: ctrl}
	mock.recorder = &MockTokenRequestValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenRequestValidator) EXPECT() *MockTokenRequestValidatorMockRecorder {
	return m.recorder
}

// ValidateAuthCode mocks base method.
func (m *MockTokenRequestValidator) ValidateAuthCode(ctx context.Context, code, clientId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateAuthCode", ctx, code, clientId)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateAuthCode indicates an expected call of ValidateAuthCode.
func (mr *MockTokenRequestValidatorMockRecorder) ValidateAuthCode(ctx, code, clientId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateAuthCode", reflect.TypeOf((*MockTokenRequestValidator)(nil).ValidateAuthCode), ctx, code, clientId)
}

// ValidateRefreshToken mocks base method.
func (m *MockTokenRequestValidator) ValidateRefreshToken(ctx context.Context, token *string, clientId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateRefreshToken", ctx, token, clientId)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateRefreshToken indicates an expected call of ValidateRefreshToken.
func (mr *MockTokenRequestValidatorMockRecorder) ValidateRefreshToken(ctx, token, clientId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateRefreshToken", reflect.TypeOf((*MockTokenRequestValidator)(nil).ValidateRefreshToken), ctx, token, clientId)
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
func (m *MockAuthCodeRevoker) RevokeAuthCode(ctx context.Context, code, clientId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeAuthCode", ctx, code, clientId)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeAuthCode indicates an expected call of RevokeAuthCode.
func (mr *MockAuthCodeRevokerMockRecorder) RevokeAuthCode(ctx, code, clientId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeAuthCode", reflect.TypeOf((*MockAuthCodeRevoker)(nil).RevokeAuthCode), ctx, code, clientId)
}

// MockTokenCacheReadWriter is a mock of TokenCacheReadWriter interface.
type MockTokenCacheReadWriter struct {
	ctrl     *gomock.Controller
	recorder *MockTokenCacheReadWriterMockRecorder
}

// MockTokenCacheReadWriterMockRecorder is the mock recorder for MockTokenCacheReadWriter.
type MockTokenCacheReadWriterMockRecorder struct {
	mock *MockTokenCacheReadWriter
}

// NewMockTokenCacheReadWriter creates a new mock instance.
func NewMockTokenCacheReadWriter(ctrl *gomock.Controller) *MockTokenCacheReadWriter {
	mock := &MockTokenCacheReadWriter{ctrl: ctrl}
	mock.recorder = &MockTokenCacheReadWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenCacheReadWriter) EXPECT() *MockTokenCacheReadWriterMockRecorder {
	return m.recorder
}

// ReadOpenIdParam mocks base method.
func (m *MockTokenCacheReadWriter) ReadOpenIdParam(ctx context.Context, clientId, authCode string) (*typedef.OpenIdParam, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadOpenIdParam", ctx, clientId, authCode)
	ret0, _ := ret[0].(*typedef.OpenIdParam)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadOpenIdParam indicates an expected call of ReadOpenIdParam.
func (mr *MockTokenCacheReadWriterMockRecorder) ReadOpenIdParam(ctx, clientId, authCode any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadOpenIdParam", reflect.TypeOf((*MockTokenCacheReadWriter)(nil).ReadOpenIdParam), ctx, clientId, authCode)
}

// ReadRefreshTokenOwner mocks base method.
func (m *MockTokenCacheReadWriter) ReadRefreshTokenOwner(ctx context.Context, token string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadRefreshTokenOwner", ctx, token)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadRefreshTokenOwner indicates an expected call of ReadRefreshTokenOwner.
func (mr *MockTokenCacheReadWriterMockRecorder) ReadRefreshTokenOwner(ctx, token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadRefreshTokenOwner", reflect.TypeOf((*MockTokenCacheReadWriter)(nil).ReadRefreshTokenOwner), ctx, token)
}

// WriteRefreshTokenOwner mocks base method.
func (m *MockTokenCacheReadWriter) WriteRefreshTokenOwner(ctx context.Context, token, clientId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteRefreshTokenOwner", ctx, token, clientId)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteRefreshTokenOwner indicates an expected call of WriteRefreshTokenOwner.
func (mr *MockTokenCacheReadWriterMockRecorder) WriteRefreshTokenOwner(ctx, token, clientId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteRefreshTokenOwner", reflect.TypeOf((*MockTokenCacheReadWriter)(nil).WriteRefreshTokenOwner), ctx, token, clientId)
}

// MockTokenRequestAcceptor is a mock of TokenRequestAcceptor interface.
type MockTokenRequestAcceptor struct {
	ctrl     *gomock.Controller
	recorder *MockTokenRequestAcceptorMockRecorder
}

// MockTokenRequestAcceptorMockRecorder is the mock recorder for MockTokenRequestAcceptor.
type MockTokenRequestAcceptorMockRecorder struct {
	mock *MockTokenRequestAcceptor
}

// NewMockTokenRequestAcceptor creates a new mock instance.
func NewMockTokenRequestAcceptor(ctrl *gomock.Controller) *MockTokenRequestAcceptor {
	mock := &MockTokenRequestAcceptor{ctrl: ctrl}
	mock.recorder = &MockTokenRequestAcceptorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenRequestAcceptor) EXPECT() *MockTokenRequestAcceptorMockRecorder {
	return m.recorder
}

// GenerateAccessToken mocks base method.
func (m *MockTokenRequestAcceptor) GenerateAccessToken() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAccessToken")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateAccessToken indicates an expected call of GenerateAccessToken.
func (mr *MockTokenRequestAcceptorMockRecorder) GenerateAccessToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAccessToken", reflect.TypeOf((*MockTokenRequestAcceptor)(nil).GenerateAccessToken))
}

// GenerateIdToken mocks base method.
func (m *MockTokenRequestAcceptor) GenerateIdToken(uid typedef.UserID) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateIdToken", uid)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateIdToken indicates an expected call of GenerateIdToken.
func (mr *MockTokenRequestAcceptorMockRecorder) GenerateIdToken(uid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateIdToken", reflect.TypeOf((*MockTokenRequestAcceptor)(nil).GenerateIdToken), uid)
}

// GenerateRefreshToken mocks base method.
func (m *MockTokenRequestAcceptor) GenerateRefreshToken() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRefreshToken")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateRefreshToken indicates an expected call of GenerateRefreshToken.
func (mr *MockTokenRequestAcceptorMockRecorder) GenerateRefreshToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRefreshToken", reflect.TypeOf((*MockTokenRequestAcceptor)(nil).GenerateRefreshToken))
}

// RevokeAuthCode mocks base method.
func (m *MockTokenRequestAcceptor) RevokeAuthCode(ctx context.Context, code, clientId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeAuthCode", ctx, code, clientId)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeAuthCode indicates an expected call of RevokeAuthCode.
func (mr *MockTokenRequestAcceptorMockRecorder) RevokeAuthCode(ctx, code, clientId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeAuthCode", reflect.TypeOf((*MockTokenRequestAcceptor)(nil).RevokeAuthCode), ctx, code, clientId)
}

// ValidateAuthCode mocks base method.
func (m *MockTokenRequestAcceptor) ValidateAuthCode(ctx context.Context, code, clientId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateAuthCode", ctx, code, clientId)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateAuthCode indicates an expected call of ValidateAuthCode.
func (mr *MockTokenRequestAcceptorMockRecorder) ValidateAuthCode(ctx, code, clientId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateAuthCode", reflect.TypeOf((*MockTokenRequestAcceptor)(nil).ValidateAuthCode), ctx, code, clientId)
}

// ValidateRefreshToken mocks base method.
func (m *MockTokenRequestAcceptor) ValidateRefreshToken(ctx context.Context, token *string, clientId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateRefreshToken", ctx, token, clientId)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateRefreshToken indicates an expected call of ValidateRefreshToken.
func (mr *MockTokenRequestAcceptorMockRecorder) ValidateRefreshToken(ctx, token, clientId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateRefreshToken", reflect.TypeOf((*MockTokenRequestAcceptor)(nil).ValidateRefreshToken), ctx, token, clientId)
}
