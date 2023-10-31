// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go
//
// Generated by this command:
//
//	mockgen -source=interface.go -destination=interface_mock.go -package=httpstore
//
// Package httpstore is a generated GoMock package.
package httpstore

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockSessionBasicReader is a mock of SessionBasicReader interface.
type MockSessionBasicReader struct {
	ctrl     *gomock.Controller
	recorder *MockSessionBasicReaderMockRecorder
}

// MockSessionBasicReaderMockRecorder is the mock recorder for MockSessionBasicReader.
type MockSessionBasicReaderMockRecorder struct {
	mock *MockSessionBasicReader
}

// NewMockSessionBasicReader creates a new mock instance.
func NewMockSessionBasicReader(ctrl *gomock.Controller) *MockSessionBasicReader {
	mock := &MockSessionBasicReader{ctrl: ctrl}
	mock.recorder = &MockSessionBasicReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionBasicReader) EXPECT() *MockSessionBasicReaderMockRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockSessionBasicReader) Read(ctx context.Context, key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockSessionBasicReaderMockRecorder) Read(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockSessionBasicReader)(nil).Read), ctx, key)
}

// MockSessionHashReader is a mock of SessionHashReader interface.
type MockSessionHashReader struct {
	ctrl     *gomock.Controller
	recorder *MockSessionHashReaderMockRecorder
}

// MockSessionHashReaderMockRecorder is the mock recorder for MockSessionHashReader.
type MockSessionHashReaderMockRecorder struct {
	mock *MockSessionHashReader
}

// NewMockSessionHashReader creates a new mock instance.
func NewMockSessionHashReader(ctrl *gomock.Controller) *MockSessionHashReader {
	mock := &MockSessionHashReader{ctrl: ctrl}
	mock.recorder = &MockSessionHashReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionHashReader) EXPECT() *MockSessionHashReaderMockRecorder {
	return m.recorder
}

// ReadHash mocks base method.
func (m *MockSessionHashReader) ReadHash(ctx context.Context, key, field string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadHash", ctx, key, field)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadHash indicates an expected call of ReadHash.
func (mr *MockSessionHashReaderMockRecorder) ReadHash(ctx, key, field any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadHash", reflect.TypeOf((*MockSessionHashReader)(nil).ReadHash), ctx, key, field)
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
func (m *MockSessionReader) Read(ctx context.Context, key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockSessionReaderMockRecorder) Read(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockSessionReader)(nil).Read), ctx, key)
}

// ReadHash mocks base method.
func (m *MockSessionReader) ReadHash(ctx context.Context, key, field string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadHash", ctx, key, field)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadHash indicates an expected call of ReadHash.
func (mr *MockSessionReaderMockRecorder) ReadHash(ctx, key, field any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadHash", reflect.TypeOf((*MockSessionReader)(nil).ReadHash), ctx, key, field)
}

// MockSessionBasicWriter is a mock of SessionBasicWriter interface.
type MockSessionBasicWriter struct {
	ctrl     *gomock.Controller
	recorder *MockSessionBasicWriterMockRecorder
}

// MockSessionBasicWriterMockRecorder is the mock recorder for MockSessionBasicWriter.
type MockSessionBasicWriterMockRecorder struct {
	mock *MockSessionBasicWriter
}

// NewMockSessionBasicWriter creates a new mock instance.
func NewMockSessionBasicWriter(ctrl *gomock.Controller) *MockSessionBasicWriter {
	mock := &MockSessionBasicWriter{ctrl: ctrl}
	mock.recorder = &MockSessionBasicWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionBasicWriter) EXPECT() *MockSessionBasicWriterMockRecorder {
	return m.recorder
}

// Write mocks base method.
func (m *MockSessionBasicWriter) Write(ctx context.Context, key string, value any, ttl time.Duration) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", ctx, key, value, ttl)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockSessionBasicWriterMockRecorder) Write(ctx, key, value, ttl any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockSessionBasicWriter)(nil).Write), ctx, key, value, ttl)
}

// MockSessionHashWriter is a mock of SessionHashWriter interface.
type MockSessionHashWriter struct {
	ctrl     *gomock.Controller
	recorder *MockSessionHashWriterMockRecorder
}

// MockSessionHashWriterMockRecorder is the mock recorder for MockSessionHashWriter.
type MockSessionHashWriterMockRecorder struct {
	mock *MockSessionHashWriter
}

// NewMockSessionHashWriter creates a new mock instance.
func NewMockSessionHashWriter(ctrl *gomock.Controller) *MockSessionHashWriter {
	mock := &MockSessionHashWriter{ctrl: ctrl}
	mock.recorder = &MockSessionHashWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionHashWriter) EXPECT() *MockSessionHashWriterMockRecorder {
	return m.recorder
}

// WriteHash mocks base method.
func (m *MockSessionHashWriter) WriteHash(ctx context.Context, key string, values map[string]string, ttl time.Duration) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteHash", ctx, key, values, ttl)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteHash indicates an expected call of WriteHash.
func (mr *MockSessionHashWriterMockRecorder) WriteHash(ctx, key, values, ttl any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteHash", reflect.TypeOf((*MockSessionHashWriter)(nil).WriteHash), ctx, key, values, ttl)
}

// MockSessionWriter is a mock of SessionWriter interface.
type MockSessionWriter struct {
	ctrl     *gomock.Controller
	recorder *MockSessionWriterMockRecorder
}

// MockSessionWriterMockRecorder is the mock recorder for MockSessionWriter.
type MockSessionWriterMockRecorder struct {
	mock *MockSessionWriter
}

// NewMockSessionWriter creates a new mock instance.
func NewMockSessionWriter(ctrl *gomock.Controller) *MockSessionWriter {
	mock := &MockSessionWriter{ctrl: ctrl}
	mock.recorder = &MockSessionWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionWriter) EXPECT() *MockSessionWriterMockRecorder {
	return m.recorder
}

// Write mocks base method.
func (m *MockSessionWriter) Write(ctx context.Context, key string, value any, ttl time.Duration) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", ctx, key, value, ttl)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockSessionWriterMockRecorder) Write(ctx, key, value, ttl any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockSessionWriter)(nil).Write), ctx, key, value, ttl)
}

// WriteHash mocks base method.
func (m *MockSessionWriter) WriteHash(ctx context.Context, key string, values map[string]string, ttl time.Duration) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteHash", ctx, key, values, ttl)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteHash indicates an expected call of WriteHash.
func (mr *MockSessionWriterMockRecorder) WriteHash(ctx, key, values, ttl any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteHash", reflect.TypeOf((*MockSessionWriter)(nil).WriteHash), ctx, key, values, ttl)
}

// MockIdGenerator is a mock of IdGenerator interface.
type MockIdGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockIdGeneratorMockRecorder
}

// MockIdGeneratorMockRecorder is the mock recorder for MockIdGenerator.
type MockIdGeneratorMockRecorder struct {
	mock *MockIdGenerator
}

// NewMockIdGenerator creates a new mock instance.
func NewMockIdGenerator(ctrl *gomock.Controller) *MockIdGenerator {
	mock := &MockIdGenerator{ctrl: ctrl}
	mock.recorder = &MockIdGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIdGenerator) EXPECT() *MockIdGeneratorMockRecorder {
	return m.recorder
}

// NextID mocks base method.
func (m *MockIdGenerator) NextID() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextID")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NextID indicates an expected call of NextID.
func (mr *MockIdGeneratorMockRecorder) NextID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextID", reflect.TypeOf((*MockIdGenerator)(nil).NextID))
}
