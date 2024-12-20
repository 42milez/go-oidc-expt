// Code generated by MockGen. DO NOT EDIT.
// Source: iface.go
//
// Generated by this command:
//
//	mockgen -source=iface.go -destination=iface_mock.go -package=iface
//

// Package iface is a generated GoMock package.
package iface

import (
	context "context"
	http "net/http"
	reflect "reflect"
	time "time"

	typedef "github.com/42milez/go-oidc-expt/pkg/typedef"
	jwt "github.com/lestrrat-go/jwx/v2/jwt"
	gomock "go.uber.org/mock/gomock"
)

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
func (mr *MockContextReaderMockRecorder) Read(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockContextReader)(nil).Read), ctx, key)
}

// MockClocker is a mock of Clocker interface.
type MockClocker struct {
	ctrl     *gomock.Controller
	recorder *MockClockerMockRecorder
}

// MockClockerMockRecorder is the mock recorder for MockClocker.
type MockClockerMockRecorder struct {
	mock *MockClocker
}

// NewMockClocker creates a new mock instance.
func NewMockClocker(ctrl *gomock.Controller) *MockClocker {
	mock := &MockClocker{ctrl: ctrl}
	mock.recorder = &MockClockerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClocker) EXPECT() *MockClockerMockRecorder {
	return m.recorder
}

// Now mocks base method.
func (m *MockClocker) Now() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Now")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Now indicates an expected call of Now.
func (mr *MockClockerMockRecorder) Now() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Now", reflect.TypeOf((*MockClocker)(nil).Now))
}

// MockCookieReader is a mock of CookieReader interface.
type MockCookieReader struct {
	ctrl     *gomock.Controller
	recorder *MockCookieReaderMockRecorder
}

// MockCookieReaderMockRecorder is the mock recorder for MockCookieReader.
type MockCookieReaderMockRecorder struct {
	mock *MockCookieReader
}

// NewMockCookieReader creates a new mock instance.
func NewMockCookieReader(ctrl *gomock.Controller) *MockCookieReader {
	mock := &MockCookieReader{ctrl: ctrl}
	mock.recorder = &MockCookieReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCookieReader) EXPECT() *MockCookieReaderMockRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockCookieReader) Read(r *http.Request, name string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", r, name)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockCookieReaderMockRecorder) Read(r, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockCookieReader)(nil).Read), r, name)
}

// MockCookieWriter is a mock of CookieWriter interface.
type MockCookieWriter struct {
	ctrl     *gomock.Controller
	recorder *MockCookieWriterMockRecorder
}

// MockCookieWriterMockRecorder is the mock recorder for MockCookieWriter.
type MockCookieWriterMockRecorder struct {
	mock *MockCookieWriter
}

// NewMockCookieWriter creates a new mock instance.
func NewMockCookieWriter(ctrl *gomock.Controller) *MockCookieWriter {
	mock := &MockCookieWriter{ctrl: ctrl}
	mock.recorder = &MockCookieWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCookieWriter) EXPECT() *MockCookieWriterMockRecorder {
	return m.recorder
}

// Write mocks base method.
func (m *MockCookieWriter) Write(w http.ResponseWriter, name, val string, ttl time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", w, name, val, ttl)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockCookieWriterMockRecorder) Write(w, name, val, ttl any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockCookieWriter)(nil).Write), w, name, val, ttl)
}

// MockCookieReadWriter is a mock of CookieReadWriter interface.
type MockCookieReadWriter struct {
	ctrl     *gomock.Controller
	recorder *MockCookieReadWriterMockRecorder
}

// MockCookieReadWriterMockRecorder is the mock recorder for MockCookieReadWriter.
type MockCookieReadWriterMockRecorder struct {
	mock *MockCookieReadWriter
}

// NewMockCookieReadWriter creates a new mock instance.
func NewMockCookieReadWriter(ctrl *gomock.Controller) *MockCookieReadWriter {
	mock := &MockCookieReadWriter{ctrl: ctrl}
	mock.recorder = &MockCookieReadWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCookieReadWriter) EXPECT() *MockCookieReadWriterMockRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockCookieReadWriter) Read(r *http.Request, name string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", r, name)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockCookieReadWriterMockRecorder) Read(r, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockCookieReadWriter)(nil).Read), r, name)
}

// Write mocks base method.
func (m *MockCookieReadWriter) Write(w http.ResponseWriter, name, val string, ttl time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", w, name, val, ttl)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockCookieReadWriterMockRecorder) Write(w, name, val, ttl any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockCookieReadWriter)(nil).Write), w, name, val, ttl)
}

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
func (m *MockAccessTokenGenerator) GenerateAccessToken(uid typedef.UserID, claims map[string]any) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAccessToken", uid, claims)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateAccessToken indicates an expected call of GenerateAccessToken.
func (mr *MockAccessTokenGeneratorMockRecorder) GenerateAccessToken(uid, claims any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAccessToken", reflect.TypeOf((*MockAccessTokenGenerator)(nil).GenerateAccessToken), uid, claims)
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
func (m *MockRefreshTokenGenerator) GenerateRefreshToken(uid typedef.UserID, claims map[string]any) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRefreshToken", uid, claims)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateRefreshToken indicates an expected call of GenerateRefreshToken.
func (mr *MockRefreshTokenGeneratorMockRecorder) GenerateRefreshToken(uid, claims any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRefreshToken", reflect.TypeOf((*MockRefreshTokenGenerator)(nil).GenerateRefreshToken), uid, claims)
}

// MockIDTokenGenerator is a mock of IDTokenGenerator interface.
type MockIDTokenGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockIDTokenGeneratorMockRecorder
}

// MockIDTokenGeneratorMockRecorder is the mock recorder for MockIDTokenGenerator.
type MockIDTokenGeneratorMockRecorder struct {
	mock *MockIDTokenGenerator
}

// NewMockIDTokenGenerator creates a new mock instance.
func NewMockIDTokenGenerator(ctrl *gomock.Controller) *MockIDTokenGenerator {
	mock := &MockIDTokenGenerator{ctrl: ctrl}
	mock.recorder = &MockIDTokenGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDTokenGenerator) EXPECT() *MockIDTokenGeneratorMockRecorder {
	return m.recorder
}

// GenerateIDToken mocks base method.
func (m *MockIDTokenGenerator) GenerateIDToken(uid typedef.UserID, audiences []string, authTime time.Time, nonce string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateIDToken", uid, audiences, authTime, nonce)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateIDToken indicates an expected call of GenerateIDToken.
func (mr *MockIDTokenGeneratorMockRecorder) GenerateIDToken(uid, audiences, authTime, nonce any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateIDToken", reflect.TypeOf((*MockIDTokenGenerator)(nil).GenerateIDToken), uid, audiences, authTime, nonce)
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
func (m *MockTokenGenerator) GenerateAccessToken(uid typedef.UserID, claims map[string]any) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAccessToken", uid, claims)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateAccessToken indicates an expected call of GenerateAccessToken.
func (mr *MockTokenGeneratorMockRecorder) GenerateAccessToken(uid, claims any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAccessToken", reflect.TypeOf((*MockTokenGenerator)(nil).GenerateAccessToken), uid, claims)
}

// GenerateIDToken mocks base method.
func (m *MockTokenGenerator) GenerateIDToken(uid typedef.UserID, audiences []string, authTime time.Time, nonce string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateIDToken", uid, audiences, authTime, nonce)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateIDToken indicates an expected call of GenerateIDToken.
func (mr *MockTokenGeneratorMockRecorder) GenerateIDToken(uid, audiences, authTime, nonce any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateIDToken", reflect.TypeOf((*MockTokenGenerator)(nil).GenerateIDToken), uid, audiences, authTime, nonce)
}

// GenerateRefreshToken mocks base method.
func (m *MockTokenGenerator) GenerateRefreshToken(uid typedef.UserID, claims map[string]any) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRefreshToken", uid, claims)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateRefreshToken indicates an expected call of GenerateRefreshToken.
func (mr *MockTokenGeneratorMockRecorder) GenerateRefreshToken(uid, claims any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRefreshToken", reflect.TypeOf((*MockTokenGenerator)(nil).GenerateRefreshToken), uid, claims)
}

// MockTokenParser is a mock of TokenParser interface.
type MockTokenParser struct {
	ctrl     *gomock.Controller
	recorder *MockTokenParserMockRecorder
}

// MockTokenParserMockRecorder is the mock recorder for MockTokenParser.
type MockTokenParserMockRecorder struct {
	mock *MockTokenParser
}

// NewMockTokenParser creates a new mock instance.
func NewMockTokenParser(ctrl *gomock.Controller) *MockTokenParser {
	mock := &MockTokenParser{ctrl: ctrl}
	mock.recorder = &MockTokenParserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenParser) EXPECT() *MockTokenParserMockRecorder {
	return m.recorder
}

// Parse mocks base method.
func (m *MockTokenParser) Parse(token string) (jwt.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parse", token)
	ret0, _ := ret[0].(jwt.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parse indicates an expected call of Parse.
func (mr *MockTokenParserMockRecorder) Parse(token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockTokenParser)(nil).Parse), token)
}

// MockTokenProcessor is a mock of TokenProcessor interface.
type MockTokenProcessor struct {
	ctrl     *gomock.Controller
	recorder *MockTokenProcessorMockRecorder
}

// MockTokenProcessorMockRecorder is the mock recorder for MockTokenProcessor.
type MockTokenProcessorMockRecorder struct {
	mock *MockTokenProcessor
}

// NewMockTokenProcessor creates a new mock instance.
func NewMockTokenProcessor(ctrl *gomock.Controller) *MockTokenProcessor {
	mock := &MockTokenProcessor{ctrl: ctrl}
	mock.recorder = &MockTokenProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenProcessor) EXPECT() *MockTokenProcessorMockRecorder {
	return m.recorder
}

// GenerateAccessToken mocks base method.
func (m *MockTokenProcessor) GenerateAccessToken(uid typedef.UserID, claims map[string]any) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAccessToken", uid, claims)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateAccessToken indicates an expected call of GenerateAccessToken.
func (mr *MockTokenProcessorMockRecorder) GenerateAccessToken(uid, claims any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAccessToken", reflect.TypeOf((*MockTokenProcessor)(nil).GenerateAccessToken), uid, claims)
}

// GenerateIDToken mocks base method.
func (m *MockTokenProcessor) GenerateIDToken(uid typedef.UserID, audiences []string, authTime time.Time, nonce string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateIDToken", uid, audiences, authTime, nonce)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateIDToken indicates an expected call of GenerateIDToken.
func (mr *MockTokenProcessorMockRecorder) GenerateIDToken(uid, audiences, authTime, nonce any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateIDToken", reflect.TypeOf((*MockTokenProcessor)(nil).GenerateIDToken), uid, audiences, authTime, nonce)
}

// GenerateRefreshToken mocks base method.
func (m *MockTokenProcessor) GenerateRefreshToken(uid typedef.UserID, claims map[string]any) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRefreshToken", uid, claims)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateRefreshToken indicates an expected call of GenerateRefreshToken.
func (mr *MockTokenProcessorMockRecorder) GenerateRefreshToken(uid, claims any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRefreshToken", reflect.TypeOf((*MockTokenProcessor)(nil).GenerateRefreshToken), uid, claims)
}

// Parse mocks base method.
func (m *MockTokenProcessor) Parse(token string) (jwt.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parse", token)
	ret0, _ := ret[0].(jwt.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parse indicates an expected call of Parse.
func (mr *MockTokenProcessorMockRecorder) Parse(token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockTokenProcessor)(nil).Parse), token)
}

// MockAuthorizationRequestFingerprintReader is a mock of AuthorizationRequestFingerprintReader interface.
type MockAuthorizationRequestFingerprintReader struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationRequestFingerprintReaderMockRecorder
}

// MockAuthorizationRequestFingerprintReaderMockRecorder is the mock recorder for MockAuthorizationRequestFingerprintReader.
type MockAuthorizationRequestFingerprintReaderMockRecorder struct {
	mock *MockAuthorizationRequestFingerprintReader
}

// NewMockAuthorizationRequestFingerprintReader creates a new mock instance.
func NewMockAuthorizationRequestFingerprintReader(ctrl *gomock.Controller) *MockAuthorizationRequestFingerprintReader {
	mock := &MockAuthorizationRequestFingerprintReader{ctrl: ctrl}
	mock.recorder = &MockAuthorizationRequestFingerprintReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorizationRequestFingerprintReader) EXPECT() *MockAuthorizationRequestFingerprintReaderMockRecorder {
	return m.recorder
}

// ReadAuthorizationRequestFingerprint mocks base method.
func (m *MockAuthorizationRequestFingerprintReader) ReadAuthorizationRequestFingerprint(ctx context.Context, clientID typedef.ClientID, authCode string) (*typedef.AuthorizationRequestFingerprint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAuthorizationRequestFingerprint", ctx, clientID, authCode)
	ret0, _ := ret[0].(*typedef.AuthorizationRequestFingerprint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAuthorizationRequestFingerprint indicates an expected call of ReadAuthorizationRequestFingerprint.
func (mr *MockAuthorizationRequestFingerprintReaderMockRecorder) ReadAuthorizationRequestFingerprint(ctx, clientID, authCode any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAuthorizationRequestFingerprint", reflect.TypeOf((*MockAuthorizationRequestFingerprintReader)(nil).ReadAuthorizationRequestFingerprint), ctx, clientID, authCode)
}

// MockAuthorizationRequestFingerprintWriter is a mock of AuthorizationRequestFingerprintWriter interface.
type MockAuthorizationRequestFingerprintWriter struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationRequestFingerprintWriterMockRecorder
}

// MockAuthorizationRequestFingerprintWriterMockRecorder is the mock recorder for MockAuthorizationRequestFingerprintWriter.
type MockAuthorizationRequestFingerprintWriterMockRecorder struct {
	mock *MockAuthorizationRequestFingerprintWriter
}

// NewMockAuthorizationRequestFingerprintWriter creates a new mock instance.
func NewMockAuthorizationRequestFingerprintWriter(ctrl *gomock.Controller) *MockAuthorizationRequestFingerprintWriter {
	mock := &MockAuthorizationRequestFingerprintWriter{ctrl: ctrl}
	mock.recorder = &MockAuthorizationRequestFingerprintWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorizationRequestFingerprintWriter) EXPECT() *MockAuthorizationRequestFingerprintWriterMockRecorder {
	return m.recorder
}

// WriteAuthorizationRequestFingerprint mocks base method.
func (m *MockAuthorizationRequestFingerprintWriter) WriteAuthorizationRequestFingerprint(ctx context.Context, clientID typedef.ClientID, authCode string, param *typedef.AuthorizationRequestFingerprint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteAuthorizationRequestFingerprint", ctx, clientID, authCode, param)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteAuthorizationRequestFingerprint indicates an expected call of WriteAuthorizationRequestFingerprint.
func (mr *MockAuthorizationRequestFingerprintWriterMockRecorder) WriteAuthorizationRequestFingerprint(ctx, clientID, authCode, param any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteAuthorizationRequestFingerprint", reflect.TypeOf((*MockAuthorizationRequestFingerprintWriter)(nil).WriteAuthorizationRequestFingerprint), ctx, clientID, authCode, param)
}

// MockRefreshTokenReader is a mock of RefreshTokenReader interface.
type MockRefreshTokenReader struct {
	ctrl     *gomock.Controller
	recorder *MockRefreshTokenReaderMockRecorder
}

// MockRefreshTokenReaderMockRecorder is the mock recorder for MockRefreshTokenReader.
type MockRefreshTokenReaderMockRecorder struct {
	mock *MockRefreshTokenReader
}

// NewMockRefreshTokenReader creates a new mock instance.
func NewMockRefreshTokenReader(ctrl *gomock.Controller) *MockRefreshTokenReader {
	mock := &MockRefreshTokenReader{ctrl: ctrl}
	mock.recorder = &MockRefreshTokenReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRefreshTokenReader) EXPECT() *MockRefreshTokenReaderMockRecorder {
	return m.recorder
}

// ReadRefreshToken mocks base method.
func (m *MockRefreshTokenReader) ReadRefreshToken(ctx context.Context, clientID typedef.ClientID, userID typedef.UserID) (jwt.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadRefreshToken", ctx, clientID, userID)
	ret0, _ := ret[0].(jwt.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadRefreshToken indicates an expected call of ReadRefreshToken.
func (mr *MockRefreshTokenReaderMockRecorder) ReadRefreshToken(ctx, clientID, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadRefreshToken", reflect.TypeOf((*MockRefreshTokenReader)(nil).ReadRefreshToken), ctx, clientID, userID)
}

// MockRefreshTokenWriter is a mock of RefreshTokenWriter interface.
type MockRefreshTokenWriter struct {
	ctrl     *gomock.Controller
	recorder *MockRefreshTokenWriterMockRecorder
}

// MockRefreshTokenWriterMockRecorder is the mock recorder for MockRefreshTokenWriter.
type MockRefreshTokenWriterMockRecorder struct {
	mock *MockRefreshTokenWriter
}

// NewMockRefreshTokenWriter creates a new mock instance.
func NewMockRefreshTokenWriter(ctrl *gomock.Controller) *MockRefreshTokenWriter {
	mock := &MockRefreshTokenWriter{ctrl: ctrl}
	mock.recorder = &MockRefreshTokenWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRefreshTokenWriter) EXPECT() *MockRefreshTokenWriterMockRecorder {
	return m.recorder
}

// WriteRefreshToken mocks base method.
func (m *MockRefreshTokenWriter) WriteRefreshToken(ctx context.Context, token string, clientID typedef.ClientID, userID typedef.UserID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteRefreshToken", ctx, token, clientID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteRefreshToken indicates an expected call of WriteRefreshToken.
func (mr *MockRefreshTokenWriterMockRecorder) WriteRefreshToken(ctx, token, clientID, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteRefreshToken", reflect.TypeOf((*MockRefreshTokenWriter)(nil).WriteRefreshToken), ctx, token, clientID, userID)
}

// MockSessionCreator is a mock of SessionCreator interface.
type MockSessionCreator struct {
	ctrl     *gomock.Controller
	recorder *MockSessionCreatorMockRecorder
}

// MockSessionCreatorMockRecorder is the mock recorder for MockSessionCreator.
type MockSessionCreatorMockRecorder struct {
	mock *MockSessionCreator
}

// NewMockSessionCreator creates a new mock instance.
func NewMockSessionCreator(ctrl *gomock.Controller) *MockSessionCreator {
	mock := &MockSessionCreator{ctrl: ctrl}
	mock.recorder = &MockSessionCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionCreator) EXPECT() *MockSessionCreatorMockRecorder {
	return m.recorder
}

// CreateSession mocks base method.
func (m *MockSessionCreator) CreateSession(ctx context.Context, uid typedef.UserID) (typedef.SessionID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", ctx, uid)
	ret0, _ := ret[0].(typedef.SessionID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockSessionCreatorMockRecorder) CreateSession(ctx, uid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockSessionCreator)(nil).CreateSession), ctx, uid)
}

// MockStructValidator is a mock of StructValidator interface.
type MockStructValidator struct {
	ctrl     *gomock.Controller
	recorder *MockStructValidatorMockRecorder
}

// MockStructValidatorMockRecorder is the mock recorder for MockStructValidator.
type MockStructValidatorMockRecorder struct {
	mock *MockStructValidator
}

// NewMockStructValidator creates a new mock instance.
func NewMockStructValidator(ctrl *gomock.Controller) *MockStructValidator {
	mock := &MockStructValidator{ctrl: ctrl}
	mock.recorder = &MockStructValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStructValidator) EXPECT() *MockStructValidatorMockRecorder {
	return m.recorder
}

// Struct mocks base method.
func (m *MockStructValidator) Struct(s any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Struct", s)
	ret0, _ := ret[0].(error)
	return ret0
}

// Struct indicates an expected call of Struct.
func (mr *MockStructValidatorMockRecorder) Struct(s any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Struct", reflect.TypeOf((*MockStructValidator)(nil).Struct), s)
}

// MockIDGenerator is a mock of IDGenerator interface.
type MockIDGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockIDGeneratorMockRecorder
}

// MockIDGeneratorMockRecorder is the mock recorder for MockIDGenerator.
type MockIDGeneratorMockRecorder struct {
	mock *MockIDGenerator
}

// NewMockIDGenerator creates a new mock instance.
func NewMockIDGenerator(ctrl *gomock.Controller) *MockIDGenerator {
	mock := &MockIDGenerator{ctrl: ctrl}
	mock.recorder = &MockIDGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDGenerator) EXPECT() *MockIDGeneratorMockRecorder {
	return m.recorder
}

// NextID mocks base method.
func (m *MockIDGenerator) NextID() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextID")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NextID indicates an expected call of NextID.
func (mr *MockIDGeneratorMockRecorder) NextID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextID", reflect.TypeOf((*MockIDGenerator)(nil).NextID))
}
