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

// MockCacheReader is a mock of CacheReader interface.
type MockCacheReader struct {
	ctrl     *gomock.Controller
	recorder *MockCacheReaderMockRecorder
}

// MockCacheReaderMockRecorder is the mock recorder for MockCacheReader.
type MockCacheReaderMockRecorder struct {
	mock *MockCacheReader
}

// NewMockCacheReader creates a new mock instance.
func NewMockCacheReader(ctrl *gomock.Controller) *MockCacheReader {
	mock := &MockCacheReader{ctrl: ctrl}
	mock.recorder = &MockCacheReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacheReader) EXPECT() *MockCacheReaderMockRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockCacheReader) Read(ctx context.Context, key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockCacheReaderMockRecorder) Read(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockCacheReader)(nil).Read), ctx, key)
}

// MockCacheHashReader is a mock of CacheHashReader interface.
type MockCacheHashReader struct {
	ctrl     *gomock.Controller
	recorder *MockCacheHashReaderMockRecorder
}

// MockCacheHashReaderMockRecorder is the mock recorder for MockCacheHashReader.
type MockCacheHashReaderMockRecorder struct {
	mock *MockCacheHashReader
}

// NewMockCacheHashReader creates a new mock instance.
func NewMockCacheHashReader(ctrl *gomock.Controller) *MockCacheHashReader {
	mock := &MockCacheHashReader{ctrl: ctrl}
	mock.recorder = &MockCacheHashReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacheHashReader) EXPECT() *MockCacheHashReaderMockRecorder {
	return m.recorder
}

// ReadHash mocks base method.
func (m *MockCacheHashReader) ReadHash(ctx context.Context, key, field string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadHash", ctx, key, field)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadHash indicates an expected call of ReadHash.
func (mr *MockCacheHashReaderMockRecorder) ReadHash(ctx, key, field any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadHash", reflect.TypeOf((*MockCacheHashReader)(nil).ReadHash), ctx, key, field)
}

// MockCacheHashAllReader is a mock of CacheHashAllReader interface.
type MockCacheHashAllReader struct {
	ctrl     *gomock.Controller
	recorder *MockCacheHashAllReaderMockRecorder
}

// MockCacheHashAllReaderMockRecorder is the mock recorder for MockCacheHashAllReader.
type MockCacheHashAllReaderMockRecorder struct {
	mock *MockCacheHashAllReader
}

// NewMockCacheHashAllReader creates a new mock instance.
func NewMockCacheHashAllReader(ctrl *gomock.Controller) *MockCacheHashAllReader {
	mock := &MockCacheHashAllReader{ctrl: ctrl}
	mock.recorder = &MockCacheHashAllReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacheHashAllReader) EXPECT() *MockCacheHashAllReaderMockRecorder {
	return m.recorder
}

// ReadHashAll mocks base method.
func (m *MockCacheHashAllReader) ReadHashAll(ctx context.Context, key string) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadHashAll", ctx, key)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadHashAll indicates an expected call of ReadHashAll.
func (mr *MockCacheHashAllReaderMockRecorder) ReadHashAll(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadHashAll", reflect.TypeOf((*MockCacheHashAllReader)(nil).ReadHashAll), ctx, key)
}

// MockCacheWriter is a mock of CacheWriter interface.
type MockCacheWriter struct {
	ctrl     *gomock.Controller
	recorder *MockCacheWriterMockRecorder
}

// MockCacheWriterMockRecorder is the mock recorder for MockCacheWriter.
type MockCacheWriterMockRecorder struct {
	mock *MockCacheWriter
}

// NewMockCacheWriter creates a new mock instance.
func NewMockCacheWriter(ctrl *gomock.Controller) *MockCacheWriter {
	mock := &MockCacheWriter{ctrl: ctrl}
	mock.recorder = &MockCacheWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacheWriter) EXPECT() *MockCacheWriterMockRecorder {
	return m.recorder
}

// Write mocks base method.
func (m *MockCacheWriter) Write(ctx context.Context, key string, value any, ttl time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", ctx, key, value, ttl)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockCacheWriterMockRecorder) Write(ctx, key, value, ttl any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockCacheWriter)(nil).Write), ctx, key, value, ttl)
}

// MockCacheHashWriter is a mock of CacheHashWriter interface.
type MockCacheHashWriter struct {
	ctrl     *gomock.Controller
	recorder *MockCacheHashWriterMockRecorder
}

// MockCacheHashWriterMockRecorder is the mock recorder for MockCacheHashWriter.
type MockCacheHashWriterMockRecorder struct {
	mock *MockCacheHashWriter
}

// NewMockCacheHashWriter creates a new mock instance.
func NewMockCacheHashWriter(ctrl *gomock.Controller) *MockCacheHashWriter {
	mock := &MockCacheHashWriter{ctrl: ctrl}
	mock.recorder = &MockCacheHashWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacheHashWriter) EXPECT() *MockCacheHashWriterMockRecorder {
	return m.recorder
}

// WriteHash mocks base method.
func (m *MockCacheHashWriter) WriteHash(ctx context.Context, key string, values map[string]string, ttl time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteHash", ctx, key, values, ttl)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteHash indicates an expected call of WriteHash.
func (mr *MockCacheHashWriterMockRecorder) WriteHash(ctx, key, values, ttl any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteHash", reflect.TypeOf((*MockCacheHashWriter)(nil).WriteHash), ctx, key, values, ttl)
}

// MockCacheReadWriter is a mock of CacheReadWriter interface.
type MockCacheReadWriter struct {
	ctrl     *gomock.Controller
	recorder *MockCacheReadWriterMockRecorder
}

// MockCacheReadWriterMockRecorder is the mock recorder for MockCacheReadWriter.
type MockCacheReadWriterMockRecorder struct {
	mock *MockCacheReadWriter
}

// NewMockCacheReadWriter creates a new mock instance.
func NewMockCacheReadWriter(ctrl *gomock.Controller) *MockCacheReadWriter {
	mock := &MockCacheReadWriter{ctrl: ctrl}
	mock.recorder = &MockCacheReadWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacheReadWriter) EXPECT() *MockCacheReadWriterMockRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockCacheReadWriter) Read(ctx context.Context, key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockCacheReadWriterMockRecorder) Read(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockCacheReadWriter)(nil).Read), ctx, key)
}

// ReadHash mocks base method.
func (m *MockCacheReadWriter) ReadHash(ctx context.Context, key, field string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadHash", ctx, key, field)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadHash indicates an expected call of ReadHash.
func (mr *MockCacheReadWriterMockRecorder) ReadHash(ctx, key, field any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadHash", reflect.TypeOf((*MockCacheReadWriter)(nil).ReadHash), ctx, key, field)
}

// ReadHashAll mocks base method.
func (m *MockCacheReadWriter) ReadHashAll(ctx context.Context, key string) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadHashAll", ctx, key)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadHashAll indicates an expected call of ReadHashAll.
func (mr *MockCacheReadWriterMockRecorder) ReadHashAll(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadHashAll", reflect.TypeOf((*MockCacheReadWriter)(nil).ReadHashAll), ctx, key)
}

// Write mocks base method.
func (m *MockCacheReadWriter) Write(ctx context.Context, key string, value any, ttl time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", ctx, key, value, ttl)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockCacheReadWriterMockRecorder) Write(ctx, key, value, ttl any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockCacheReadWriter)(nil).Write), ctx, key, value, ttl)
}

// WriteHash mocks base method.
func (m *MockCacheReadWriter) WriteHash(ctx context.Context, key string, values map[string]string, ttl time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteHash", ctx, key, values, ttl)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteHash indicates an expected call of WriteHash.
func (mr *MockCacheReadWriterMockRecorder) WriteHash(ctx, key, values, ttl any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteHash", reflect.TypeOf((*MockCacheReadWriter)(nil).WriteHash), ctx, key, values, ttl)
}
