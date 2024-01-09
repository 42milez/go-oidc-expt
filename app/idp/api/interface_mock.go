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
	url "net/url"
	reflect "reflect"
	time "time"

	entity "github.com/42milez/go-oidc-server/app/idp/entity"
	typedef "github.com/42milez/go-oidc-server/app/pkg/typedef"
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
func (m *MockConsentVerifier) VerifyConsent(ctx context.Context, userID typedef.UserID, clientID typedef.ClientID) (bool, error) {
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
func (m *MockAuthenticator) VerifyConsent(ctx context.Context, userID typedef.UserID, clientID typedef.ClientID) (bool, error) {
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

// MockRequestFingerprintSaver is a mock of RequestFingerprintSaver interface.
type MockRequestFingerprintSaver struct {
	ctrl     *gomock.Controller
	recorder *MockRequestFingerprintSaverMockRecorder
}

// MockRequestFingerprintSaverMockRecorder is the mock recorder for MockRequestFingerprintSaver.
type MockRequestFingerprintSaverMockRecorder struct {
	mock *MockRequestFingerprintSaver
}

// NewMockRequestFingerprintSaver creates a new mock instance.
func NewMockRequestFingerprintSaver(ctrl *gomock.Controller) *MockRequestFingerprintSaver {
	mock := &MockRequestFingerprintSaver{ctrl: ctrl}
	mock.recorder = &MockRequestFingerprintSaverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequestFingerprintSaver) EXPECT() *MockRequestFingerprintSaverMockRecorder {
	return m.recorder
}

// SaveAuthorizationRequestFingerprint mocks base method.
func (m *MockRequestFingerprintSaver) SaveAuthorizationRequestFingerprint(ctx context.Context, clientID typedef.ClientID, redirectURI, nonce, authCode string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveAuthorizationRequestFingerprint", ctx, clientID, redirectURI, nonce, authCode)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveAuthorizationRequestFingerprint indicates an expected call of SaveAuthorizationRequestFingerprint.
func (mr *MockRequestFingerprintSaverMockRecorder) SaveAuthorizationRequestFingerprint(ctx, clientID, redirectURI, nonce, authCode any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveAuthorizationRequestFingerprint", reflect.TypeOf((*MockRequestFingerprintSaver)(nil).SaveAuthorizationRequestFingerprint), ctx, clientID, redirectURI, nonce, authCode)
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
func (m *MockAuthorizer) Authorize(ctx context.Context, clientID typedef.ClientID, redirectURI, state string) (*url.URL, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorize", ctx, clientID, redirectURI, state)
	ret0, _ := ret[0].(*url.URL)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Authorize indicates an expected call of Authorize.
func (mr *MockAuthorizerMockRecorder) Authorize(ctx, clientID, redirectURI, state any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorize", reflect.TypeOf((*MockAuthorizer)(nil).Authorize), ctx, clientID, redirectURI, state)
}

// SaveAuthorizationRequestFingerprint mocks base method.
func (m *MockAuthorizer) SaveAuthorizationRequestFingerprint(ctx context.Context, clientID typedef.ClientID, redirectURI, nonce, authCode string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveAuthorizationRequestFingerprint", ctx, clientID, redirectURI, nonce, authCode)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveAuthorizationRequestFingerprint indicates an expected call of SaveAuthorizationRequestFingerprint.
func (mr *MockAuthorizerMockRecorder) SaveAuthorizationRequestFingerprint(ctx, clientID, redirectURI, nonce, authCode any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveAuthorizationRequestFingerprint", reflect.TypeOf((*MockAuthorizer)(nil).SaveAuthorizationRequestFingerprint), ctx, clientID, redirectURI, nonce, authCode)
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
func (m *MockConsentAcceptor) AcceptConsent(ctx context.Context, userID typedef.UserID, clientID typedef.ClientID) error {
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
func (m *MockCredentialValidator) ValidateCredential(ctx context.Context, clientID typedef.ClientID, clientSecret string) error {
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

// MockRefreshTokenVerifier is a mock of RefreshTokenVerifier interface.
type MockRefreshTokenVerifier struct {
	ctrl     *gomock.Controller
	recorder *MockRefreshTokenVerifierMockRecorder
}

// MockRefreshTokenVerifierMockRecorder is the mock recorder for MockRefreshTokenVerifier.
type MockRefreshTokenVerifierMockRecorder struct {
	mock *MockRefreshTokenVerifier
}

// NewMockRefreshTokenVerifier creates a new mock instance.
func NewMockRefreshTokenVerifier(ctrl *gomock.Controller) *MockRefreshTokenVerifier {
	mock := &MockRefreshTokenVerifier{ctrl: ctrl}
	mock.recorder = &MockRefreshTokenVerifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRefreshTokenVerifier) EXPECT() *MockRefreshTokenVerifierMockRecorder {
	return m.recorder
}

// VerifyRefreshToken mocks base method.
func (m *MockRefreshTokenVerifier) VerifyRefreshToken(ctx context.Context, token string, clientID typedef.ClientID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyRefreshToken", ctx, token, clientID)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyRefreshToken indicates an expected call of VerifyRefreshToken.
func (mr *MockRefreshTokenVerifierMockRecorder) VerifyRefreshToken(ctx, token, clientID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyRefreshToken", reflect.TypeOf((*MockRefreshTokenVerifier)(nil).VerifyRefreshToken), ctx, token, clientID)
}

// MockUserIDExtractor is a mock of UserIDExtractor interface.
type MockUserIDExtractor struct {
	ctrl     *gomock.Controller
	recorder *MockUserIDExtractorMockRecorder
}

// MockUserIDExtractorMockRecorder is the mock recorder for MockUserIDExtractor.
type MockUserIDExtractorMockRecorder struct {
	mock *MockUserIDExtractor
}

// NewMockUserIDExtractor creates a new mock instance.
func NewMockUserIDExtractor(ctrl *gomock.Controller) *MockUserIDExtractor {
	mock := &MockUserIDExtractor{ctrl: ctrl}
	mock.recorder = &MockUserIDExtractorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserIDExtractor) EXPECT() *MockUserIDExtractorMockRecorder {
	return m.recorder
}

// ExtractUserID mocks base method.
func (m *MockUserIDExtractor) ExtractUserID(refreshToken string) (typedef.UserID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractUserID", refreshToken)
	ret0, _ := ret[0].(typedef.UserID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExtractUserID indicates an expected call of ExtractUserID.
func (mr *MockUserIDExtractorMockRecorder) ExtractUserID(refreshToken any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractUserID", reflect.TypeOf((*MockUserIDExtractor)(nil).ExtractUserID), refreshToken)
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
func (m *MockAuthCodeRevoker) RevokeAuthCode(ctx context.Context, code string, clientID typedef.ClientID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeAuthCode", ctx, code, clientID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeAuthCode indicates an expected call of RevokeAuthCode.
func (mr *MockAuthCodeRevokerMockRecorder) RevokeAuthCode(ctx, code, clientID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeAuthCode", reflect.TypeOf((*MockAuthCodeRevoker)(nil).RevokeAuthCode), ctx, code, clientID)
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

// ReadAuthorizationRequestFingerprint mocks base method.
func (m *MockTokenCacheReadWriter) ReadAuthorizationRequestFingerprint(ctx context.Context, clientID typedef.ClientID, authCode string) (*typedef.AuthorizationRequestFingerprint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAuthorizationRequestFingerprint", ctx, clientID, authCode)
	ret0, _ := ret[0].(*typedef.AuthorizationRequestFingerprint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAuthorizationRequestFingerprint indicates an expected call of ReadAuthorizationRequestFingerprint.
func (mr *MockTokenCacheReadWriterMockRecorder) ReadAuthorizationRequestFingerprint(ctx, clientID, authCode any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAuthorizationRequestFingerprint", reflect.TypeOf((*MockTokenCacheReadWriter)(nil).ReadAuthorizationRequestFingerprint), ctx, clientID, authCode)
}

// WriteRefreshToken mocks base method.
func (m *MockTokenCacheReadWriter) WriteRefreshToken(ctx context.Context, token string, clientID typedef.ClientID, userID typedef.UserID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteRefreshToken", ctx, token, clientID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteRefreshToken indicates an expected call of WriteRefreshToken.
func (mr *MockTokenCacheReadWriterMockRecorder) WriteRefreshToken(ctx, token, clientID, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteRefreshToken", reflect.TypeOf((*MockTokenCacheReadWriter)(nil).WriteRefreshToken), ctx, token, clientID, userID)
}

// MockAuthCodeGrantAcceptor is a mock of AuthCodeGrantAcceptor interface.
type MockAuthCodeGrantAcceptor struct {
	ctrl     *gomock.Controller
	recorder *MockAuthCodeGrantAcceptorMockRecorder
}

// MockAuthCodeGrantAcceptorMockRecorder is the mock recorder for MockAuthCodeGrantAcceptor.
type MockAuthCodeGrantAcceptorMockRecorder struct {
	mock *MockAuthCodeGrantAcceptor
}

// NewMockAuthCodeGrantAcceptor creates a new mock instance.
func NewMockAuthCodeGrantAcceptor(ctrl *gomock.Controller) *MockAuthCodeGrantAcceptor {
	mock := &MockAuthCodeGrantAcceptor{ctrl: ctrl}
	mock.recorder = &MockAuthCodeGrantAcceptorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthCodeGrantAcceptor) EXPECT() *MockAuthCodeGrantAcceptorMockRecorder {
	return m.recorder
}

// GenerateAccessToken mocks base method.
func (m *MockAuthCodeGrantAcceptor) GenerateAccessToken(uid typedef.UserID, claims map[string]any) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAccessToken", uid, claims)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateAccessToken indicates an expected call of GenerateAccessToken.
func (mr *MockAuthCodeGrantAcceptorMockRecorder) GenerateAccessToken(uid, claims any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAccessToken", reflect.TypeOf((*MockAuthCodeGrantAcceptor)(nil).GenerateAccessToken), uid, claims)
}

// GenerateIDToken mocks base method.
func (m *MockAuthCodeGrantAcceptor) GenerateIDToken(uid typedef.UserID, audiences []string, authTime time.Time, nonce string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateIDToken", uid, audiences, authTime, nonce)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateIDToken indicates an expected call of GenerateIDToken.
func (mr *MockAuthCodeGrantAcceptorMockRecorder) GenerateIDToken(uid, audiences, authTime, nonce any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateIDToken", reflect.TypeOf((*MockAuthCodeGrantAcceptor)(nil).GenerateIDToken), uid, audiences, authTime, nonce)
}

// GenerateRefreshToken mocks base method.
func (m *MockAuthCodeGrantAcceptor) GenerateRefreshToken(uid typedef.UserID, claims map[string]any) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRefreshToken", uid, claims)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateRefreshToken indicates an expected call of GenerateRefreshToken.
func (mr *MockAuthCodeGrantAcceptorMockRecorder) GenerateRefreshToken(uid, claims any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRefreshToken", reflect.TypeOf((*MockAuthCodeGrantAcceptor)(nil).GenerateRefreshToken), uid, claims)
}

// RevokeAuthCode mocks base method.
func (m *MockAuthCodeGrantAcceptor) RevokeAuthCode(ctx context.Context, code string, clientID typedef.ClientID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeAuthCode", ctx, code, clientID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeAuthCode indicates an expected call of RevokeAuthCode.
func (mr *MockAuthCodeGrantAcceptorMockRecorder) RevokeAuthCode(ctx, code, clientID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeAuthCode", reflect.TypeOf((*MockAuthCodeGrantAcceptor)(nil).RevokeAuthCode), ctx, code, clientID)
}

// MockRefreshTokenGrantAcceptor is a mock of RefreshTokenGrantAcceptor interface.
type MockRefreshTokenGrantAcceptor struct {
	ctrl     *gomock.Controller
	recorder *MockRefreshTokenGrantAcceptorMockRecorder
}

// MockRefreshTokenGrantAcceptorMockRecorder is the mock recorder for MockRefreshTokenGrantAcceptor.
type MockRefreshTokenGrantAcceptorMockRecorder struct {
	mock *MockRefreshTokenGrantAcceptor
}

// NewMockRefreshTokenGrantAcceptor creates a new mock instance.
func NewMockRefreshTokenGrantAcceptor(ctrl *gomock.Controller) *MockRefreshTokenGrantAcceptor {
	mock := &MockRefreshTokenGrantAcceptor{ctrl: ctrl}
	mock.recorder = &MockRefreshTokenGrantAcceptorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRefreshTokenGrantAcceptor) EXPECT() *MockRefreshTokenGrantAcceptorMockRecorder {
	return m.recorder
}

// ExtractUserID mocks base method.
func (m *MockRefreshTokenGrantAcceptor) ExtractUserID(refreshToken string) (typedef.UserID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractUserID", refreshToken)
	ret0, _ := ret[0].(typedef.UserID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExtractUserID indicates an expected call of ExtractUserID.
func (mr *MockRefreshTokenGrantAcceptorMockRecorder) ExtractUserID(refreshToken any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractUserID", reflect.TypeOf((*MockRefreshTokenGrantAcceptor)(nil).ExtractUserID), refreshToken)
}

// GenerateAccessToken mocks base method.
func (m *MockRefreshTokenGrantAcceptor) GenerateAccessToken(uid typedef.UserID, claims map[string]any) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAccessToken", uid, claims)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateAccessToken indicates an expected call of GenerateAccessToken.
func (mr *MockRefreshTokenGrantAcceptorMockRecorder) GenerateAccessToken(uid, claims any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAccessToken", reflect.TypeOf((*MockRefreshTokenGrantAcceptor)(nil).GenerateAccessToken), uid, claims)
}

// GenerateRefreshToken mocks base method.
func (m *MockRefreshTokenGrantAcceptor) GenerateRefreshToken(uid typedef.UserID, claims map[string]any) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRefreshToken", uid, claims)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateRefreshToken indicates an expected call of GenerateRefreshToken.
func (mr *MockRefreshTokenGrantAcceptorMockRecorder) GenerateRefreshToken(uid, claims any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRefreshToken", reflect.TypeOf((*MockRefreshTokenGrantAcceptor)(nil).GenerateRefreshToken), uid, claims)
}

// VerifyRefreshToken mocks base method.
func (m *MockRefreshTokenGrantAcceptor) VerifyRefreshToken(ctx context.Context, token string, clientID typedef.ClientID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyRefreshToken", ctx, token, clientID)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyRefreshToken indicates an expected call of VerifyRefreshToken.
func (mr *MockRefreshTokenGrantAcceptorMockRecorder) VerifyRefreshToken(ctx, token, clientID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyRefreshToken", reflect.TypeOf((*MockRefreshTokenGrantAcceptor)(nil).VerifyRefreshToken), ctx, token, clientID)
}
