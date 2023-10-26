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

	typedef "github.com/42milez/go-oidc-server/app/typedef"
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
func (m *MockAccessTokenGenerator) GenerateAccessToken(uid typedef.UserID) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAccessToken", uid)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateAccessToken indicates an expected call of GenerateAccessToken.
func (mr *MockAccessTokenGeneratorMockRecorder) GenerateAccessToken(uid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAccessToken", reflect.TypeOf((*MockAccessTokenGenerator)(nil).GenerateAccessToken), uid)
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
func (m *MockRefreshTokenGenerator) GenerateRefreshToken(uid typedef.UserID) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRefreshToken", uid)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateRefreshToken indicates an expected call of GenerateRefreshToken.
func (mr *MockRefreshTokenGeneratorMockRecorder) GenerateRefreshToken(uid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRefreshToken", reflect.TypeOf((*MockRefreshTokenGenerator)(nil).GenerateRefreshToken), uid)
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
func (m *MockIdTokenGenerator) GenerateIdToken(uid typedef.UserID) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateIdToken", uid)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateIdToken indicates an expected call of GenerateIdToken.
func (mr *MockIdTokenGeneratorMockRecorder) GenerateIdToken(uid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateIdToken", reflect.TypeOf((*MockIdTokenGenerator)(nil).GenerateIdToken), uid)
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
func (m *MockTokenGenerator) GenerateAccessToken(uid typedef.UserID) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAccessToken", uid)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateAccessToken indicates an expected call of GenerateAccessToken.
func (mr *MockTokenGeneratorMockRecorder) GenerateAccessToken(uid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAccessToken", reflect.TypeOf((*MockTokenGenerator)(nil).GenerateAccessToken), uid)
}

// GenerateIdToken mocks base method.
func (m *MockTokenGenerator) GenerateIdToken(uid typedef.UserID) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateIdToken", uid)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateIdToken indicates an expected call of GenerateIdToken.
func (mr *MockTokenGeneratorMockRecorder) GenerateIdToken(uid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateIdToken", reflect.TypeOf((*MockTokenGenerator)(nil).GenerateIdToken), uid)
}

// GenerateRefreshToken mocks base method.
func (m *MockTokenGenerator) GenerateRefreshToken(uid typedef.UserID) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRefreshToken", uid)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateRefreshToken indicates an expected call of GenerateRefreshToken.
func (mr *MockTokenGeneratorMockRecorder) GenerateRefreshToken(uid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRefreshToken", reflect.TypeOf((*MockTokenGenerator)(nil).GenerateRefreshToken), uid)
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
func (mr *MockTokenValidatorMockRecorder) Validate(name any) *gomock.Call {
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
func (m *MockTokenGenerateValidator) GenerateAccessToken(uid typedef.UserID) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAccessToken", uid)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateAccessToken indicates an expected call of GenerateAccessToken.
func (mr *MockTokenGenerateValidatorMockRecorder) GenerateAccessToken(uid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAccessToken", reflect.TypeOf((*MockTokenGenerateValidator)(nil).GenerateAccessToken), uid)
}

// GenerateIdToken mocks base method.
func (m *MockTokenGenerateValidator) GenerateIdToken(uid typedef.UserID) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateIdToken", uid)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateIdToken indicates an expected call of GenerateIdToken.
func (mr *MockTokenGenerateValidatorMockRecorder) GenerateIdToken(uid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateIdToken", reflect.TypeOf((*MockTokenGenerateValidator)(nil).GenerateIdToken), uid)
}

// GenerateRefreshToken mocks base method.
func (m *MockTokenGenerateValidator) GenerateRefreshToken(uid typedef.UserID) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRefreshToken", uid)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateRefreshToken indicates an expected call of GenerateRefreshToken.
func (mr *MockTokenGenerateValidatorMockRecorder) GenerateRefreshToken(uid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRefreshToken", reflect.TypeOf((*MockTokenGenerateValidator)(nil).GenerateRefreshToken), uid)
}

// Validate mocks base method.
func (m *MockTokenGenerateValidator) Validate(name *string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockTokenGenerateValidatorMockRecorder) Validate(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockTokenGenerateValidator)(nil).Validate), name)
}

// MockRedirectUriSessionReader is a mock of RedirectUriSessionReader interface.
type MockRedirectUriSessionReader struct {
	ctrl     *gomock.Controller
	recorder *MockRedirectUriSessionReaderMockRecorder
}

// MockRedirectUriSessionReaderMockRecorder is the mock recorder for MockRedirectUriSessionReader.
type MockRedirectUriSessionReaderMockRecorder struct {
	mock *MockRedirectUriSessionReader
}

// NewMockRedirectUriSessionReader creates a new mock instance.
func NewMockRedirectUriSessionReader(ctrl *gomock.Controller) *MockRedirectUriSessionReader {
	mock := &MockRedirectUriSessionReader{ctrl: ctrl}
	mock.recorder = &MockRedirectUriSessionReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedirectUriSessionReader) EXPECT() *MockRedirectUriSessionReaderMockRecorder {
	return m.recorder
}

// ReadRedirectUri mocks base method.
func (m *MockRedirectUriSessionReader) ReadRedirectUri(ctx context.Context, clientId, authCode string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadRedirectUri", ctx, clientId, authCode)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadRedirectUri indicates an expected call of ReadRedirectUri.
func (mr *MockRedirectUriSessionReaderMockRecorder) ReadRedirectUri(ctx, clientId, authCode any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadRedirectUri", reflect.TypeOf((*MockRedirectUriSessionReader)(nil).ReadRedirectUri), ctx, clientId, authCode)
}

// MockRedirectUriSessionWriter is a mock of RedirectUriSessionWriter interface.
type MockRedirectUriSessionWriter struct {
	ctrl     *gomock.Controller
	recorder *MockRedirectUriSessionWriterMockRecorder
}

// MockRedirectUriSessionWriterMockRecorder is the mock recorder for MockRedirectUriSessionWriter.
type MockRedirectUriSessionWriterMockRecorder struct {
	mock *MockRedirectUriSessionWriter
}

// NewMockRedirectUriSessionWriter creates a new mock instance.
func NewMockRedirectUriSessionWriter(ctrl *gomock.Controller) *MockRedirectUriSessionWriter {
	mock := &MockRedirectUriSessionWriter{ctrl: ctrl}
	mock.recorder = &MockRedirectUriSessionWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedirectUriSessionWriter) EXPECT() *MockRedirectUriSessionWriterMockRecorder {
	return m.recorder
}

// WriteRedirectUriAssociation mocks base method.
func (m *MockRedirectUriSessionWriter) WriteRedirectUriAssociation(ctx context.Context, uri, clientId, authCode string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteRedirectUriAssociation", ctx, uri, clientId, authCode)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteRedirectUriAssociation indicates an expected call of WriteRedirectUriAssociation.
func (mr *MockRedirectUriSessionWriterMockRecorder) WriteRedirectUriAssociation(ctx, uri, clientId, authCode any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteRedirectUriAssociation", reflect.TypeOf((*MockRedirectUriSessionWriter)(nil).WriteRedirectUriAssociation), ctx, uri, clientId, authCode)
}

// MockRefreshTokenOwnerSessionWriter is a mock of RefreshTokenOwnerSessionWriter interface.
type MockRefreshTokenOwnerSessionWriter struct {
	ctrl     *gomock.Controller
	recorder *MockRefreshTokenOwnerSessionWriterMockRecorder
}

// MockRefreshTokenOwnerSessionWriterMockRecorder is the mock recorder for MockRefreshTokenOwnerSessionWriter.
type MockRefreshTokenOwnerSessionWriterMockRecorder struct {
	mock *MockRefreshTokenOwnerSessionWriter
}

// NewMockRefreshTokenOwnerSessionWriter creates a new mock instance.
func NewMockRefreshTokenOwnerSessionWriter(ctrl *gomock.Controller) *MockRefreshTokenOwnerSessionWriter {
	mock := &MockRefreshTokenOwnerSessionWriter{ctrl: ctrl}
	mock.recorder = &MockRefreshTokenOwnerSessionWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRefreshTokenOwnerSessionWriter) EXPECT() *MockRefreshTokenOwnerSessionWriterMockRecorder {
	return m.recorder
}

// WriteRefreshTokenOwner mocks base method.
func (m *MockRefreshTokenOwnerSessionWriter) WriteRefreshTokenOwner(ctx context.Context, token, clientId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteRefreshTokenOwner", ctx, token, clientId)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteRefreshTokenOwner indicates an expected call of WriteRefreshTokenOwner.
func (mr *MockRefreshTokenOwnerSessionWriterMockRecorder) WriteRefreshTokenOwner(ctx, token, clientId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteRefreshTokenOwner", reflect.TypeOf((*MockRefreshTokenOwnerSessionWriter)(nil).WriteRefreshTokenOwner), ctx, token, clientId)
}

// MockUserIdSessionWriter is a mock of UserIdSessionWriter interface.
type MockUserIdSessionWriter struct {
	ctrl     *gomock.Controller
	recorder *MockUserIdSessionWriterMockRecorder
}

// MockUserIdSessionWriterMockRecorder is the mock recorder for MockUserIdSessionWriter.
type MockUserIdSessionWriterMockRecorder struct {
	mock *MockUserIdSessionWriter
}

// NewMockUserIdSessionWriter creates a new mock instance.
func NewMockUserIdSessionWriter(ctrl *gomock.Controller) *MockUserIdSessionWriter {
	mock := &MockUserIdSessionWriter{ctrl: ctrl}
	mock.recorder = &MockUserIdSessionWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserIdSessionWriter) EXPECT() *MockUserIdSessionWriterMockRecorder {
	return m.recorder
}

// WriteUserId mocks base method.
func (m *MockUserIdSessionWriter) WriteUserId(ctx context.Context, userId typedef.UserID) (typedef.SessionID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteUserId", ctx, userId)
	ret0, _ := ret[0].(typedef.SessionID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteUserId indicates an expected call of WriteUserId.
func (mr *MockUserIdSessionWriterMockRecorder) WriteUserId(ctx, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteUserId", reflect.TypeOf((*MockUserIdSessionWriter)(nil).WriteUserId), ctx, userId)
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
