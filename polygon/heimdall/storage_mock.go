// Code generated by MockGen. DO NOT EDIT.
// Source: ./storage.go
//
// Generated by this command:
//
//	mockgen -destination=./storage_mock.go -package=heimdall -source=./storage.go
//

// Package heimdall is a generated GoMock package.
package heimdall

import (
	context "context"
	reflect "reflect"

	common "github.com/ledgerwatch/erigon-lib/common"
	kv "github.com/ledgerwatch/erigon-lib/kv"
	rlp "github.com/optimism-java/erigon/rlp"
	gomock "go.uber.org/mock/gomock"
)

// MockSpanReader is a mock of SpanReader interface.
type MockSpanReader struct {
	ctrl     *gomock.Controller
	recorder *MockSpanReaderMockRecorder
}

// MockSpanReaderMockRecorder is the mock recorder for MockSpanReader.
type MockSpanReaderMockRecorder struct {
	mock *MockSpanReader
}

// NewMockSpanReader creates a new mock instance.
func NewMockSpanReader(ctrl *gomock.Controller) *MockSpanReader {
	mock := &MockSpanReader{ctrl: ctrl}
	mock.recorder = &MockSpanReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpanReader) EXPECT() *MockSpanReaderMockRecorder {
	return m.recorder
}

// GetSpan mocks base method.
func (m *MockSpanReader) GetSpan(ctx context.Context, spanId SpanId) (*Span, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSpan", ctx, spanId)
	ret0, _ := ret[0].(*Span)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSpan indicates an expected call of GetSpan.
func (mr *MockSpanReaderMockRecorder) GetSpan(ctx, spanId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSpan", reflect.TypeOf((*MockSpanReader)(nil).GetSpan), ctx, spanId)
}

// LastSpanId mocks base method.
func (m *MockSpanReader) LastSpanId(ctx context.Context) (SpanId, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastSpanId", ctx)
	ret0, _ := ret[0].(SpanId)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LastSpanId indicates an expected call of LastSpanId.
func (mr *MockSpanReaderMockRecorder) LastSpanId(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastSpanId", reflect.TypeOf((*MockSpanReader)(nil).LastSpanId), ctx)
}

// MockSpanWriter is a mock of SpanWriter interface.
type MockSpanWriter struct {
	ctrl     *gomock.Controller
	recorder *MockSpanWriterMockRecorder
}

// MockSpanWriterMockRecorder is the mock recorder for MockSpanWriter.
type MockSpanWriterMockRecorder struct {
	mock *MockSpanWriter
}

// NewMockSpanWriter creates a new mock instance.
func NewMockSpanWriter(ctrl *gomock.Controller) *MockSpanWriter {
	mock := &MockSpanWriter{ctrl: ctrl}
	mock.recorder = &MockSpanWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpanWriter) EXPECT() *MockSpanWriterMockRecorder {
	return m.recorder
}

// PutSpan mocks base method.
func (m *MockSpanWriter) PutSpan(ctx context.Context, span *Span) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutSpan", ctx, span)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutSpan indicates an expected call of PutSpan.
func (mr *MockSpanWriterMockRecorder) PutSpan(ctx, span any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutSpan", reflect.TypeOf((*MockSpanWriter)(nil).PutSpan), ctx, span)
}

// MockSpanStore is a mock of SpanStore interface.
type MockSpanStore struct {
	ctrl     *gomock.Controller
	recorder *MockSpanStoreMockRecorder
}

// MockSpanStoreMockRecorder is the mock recorder for MockSpanStore.
type MockSpanStoreMockRecorder struct {
	mock *MockSpanStore
}

// NewMockSpanStore creates a new mock instance.
func NewMockSpanStore(ctrl *gomock.Controller) *MockSpanStore {
	mock := &MockSpanStore{ctrl: ctrl}
	mock.recorder = &MockSpanStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpanStore) EXPECT() *MockSpanStoreMockRecorder {
	return m.recorder
}

// GetSpan mocks base method.
func (m *MockSpanStore) GetSpan(ctx context.Context, spanId SpanId) (*Span, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSpan", ctx, spanId)
	ret0, _ := ret[0].(*Span)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSpan indicates an expected call of GetSpan.
func (mr *MockSpanStoreMockRecorder) GetSpan(ctx, spanId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSpan", reflect.TypeOf((*MockSpanStore)(nil).GetSpan), ctx, spanId)
}

// LastSpanId mocks base method.
func (m *MockSpanStore) LastSpanId(ctx context.Context) (SpanId, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastSpanId", ctx)
	ret0, _ := ret[0].(SpanId)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LastSpanId indicates an expected call of LastSpanId.
func (mr *MockSpanStoreMockRecorder) LastSpanId(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastSpanId", reflect.TypeOf((*MockSpanStore)(nil).LastSpanId), ctx)
}

// PutSpan mocks base method.
func (m *MockSpanStore) PutSpan(ctx context.Context, span *Span) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutSpan", ctx, span)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutSpan indicates an expected call of PutSpan.
func (mr *MockSpanStoreMockRecorder) PutSpan(ctx, span any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutSpan", reflect.TypeOf((*MockSpanStore)(nil).PutSpan), ctx, span)
}

// MockMilestoneReader is a mock of MilestoneReader interface.
type MockMilestoneReader struct {
	ctrl     *gomock.Controller
	recorder *MockMilestoneReaderMockRecorder
}

// MockMilestoneReaderMockRecorder is the mock recorder for MockMilestoneReader.
type MockMilestoneReaderMockRecorder struct {
	mock *MockMilestoneReader
}

// NewMockMilestoneReader creates a new mock instance.
func NewMockMilestoneReader(ctrl *gomock.Controller) *MockMilestoneReader {
	mock := &MockMilestoneReader{ctrl: ctrl}
	mock.recorder = &MockMilestoneReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMilestoneReader) EXPECT() *MockMilestoneReaderMockRecorder {
	return m.recorder
}

// GetMilestone mocks base method.
func (m *MockMilestoneReader) GetMilestone(ctx context.Context, milestoneId MilestoneId) (*Milestone, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMilestone", ctx, milestoneId)
	ret0, _ := ret[0].(*Milestone)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMilestone indicates an expected call of GetMilestone.
func (mr *MockMilestoneReaderMockRecorder) GetMilestone(ctx, milestoneId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMilestone", reflect.TypeOf((*MockMilestoneReader)(nil).GetMilestone), ctx, milestoneId)
}

// LastMilestoneId mocks base method.
func (m *MockMilestoneReader) LastMilestoneId(ctx context.Context) (MilestoneId, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastMilestoneId", ctx)
	ret0, _ := ret[0].(MilestoneId)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LastMilestoneId indicates an expected call of LastMilestoneId.
func (mr *MockMilestoneReaderMockRecorder) LastMilestoneId(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastMilestoneId", reflect.TypeOf((*MockMilestoneReader)(nil).LastMilestoneId), ctx)
}

// MockMilestoneWriter is a mock of MilestoneWriter interface.
type MockMilestoneWriter struct {
	ctrl     *gomock.Controller
	recorder *MockMilestoneWriterMockRecorder
}

// MockMilestoneWriterMockRecorder is the mock recorder for MockMilestoneWriter.
type MockMilestoneWriterMockRecorder struct {
	mock *MockMilestoneWriter
}

// NewMockMilestoneWriter creates a new mock instance.
func NewMockMilestoneWriter(ctrl *gomock.Controller) *MockMilestoneWriter {
	mock := &MockMilestoneWriter{ctrl: ctrl}
	mock.recorder = &MockMilestoneWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMilestoneWriter) EXPECT() *MockMilestoneWriterMockRecorder {
	return m.recorder
}

// PutMilestone mocks base method.
func (m *MockMilestoneWriter) PutMilestone(ctx context.Context, milestoneId MilestoneId, milestone *Milestone) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutMilestone", ctx, milestoneId, milestone)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutMilestone indicates an expected call of PutMilestone.
func (mr *MockMilestoneWriterMockRecorder) PutMilestone(ctx, milestoneId, milestone any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutMilestone", reflect.TypeOf((*MockMilestoneWriter)(nil).PutMilestone), ctx, milestoneId, milestone)
}

// MockMilestoneStore is a mock of MilestoneStore interface.
type MockMilestoneStore struct {
	ctrl     *gomock.Controller
	recorder *MockMilestoneStoreMockRecorder
}

// MockMilestoneStoreMockRecorder is the mock recorder for MockMilestoneStore.
type MockMilestoneStoreMockRecorder struct {
	mock *MockMilestoneStore
}

// NewMockMilestoneStore creates a new mock instance.
func NewMockMilestoneStore(ctrl *gomock.Controller) *MockMilestoneStore {
	mock := &MockMilestoneStore{ctrl: ctrl}
	mock.recorder = &MockMilestoneStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMilestoneStore) EXPECT() *MockMilestoneStoreMockRecorder {
	return m.recorder
}

// GetMilestone mocks base method.
func (m *MockMilestoneStore) GetMilestone(ctx context.Context, milestoneId MilestoneId) (*Milestone, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMilestone", ctx, milestoneId)
	ret0, _ := ret[0].(*Milestone)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMilestone indicates an expected call of GetMilestone.
func (mr *MockMilestoneStoreMockRecorder) GetMilestone(ctx, milestoneId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMilestone", reflect.TypeOf((*MockMilestoneStore)(nil).GetMilestone), ctx, milestoneId)
}

// LastMilestoneId mocks base method.
func (m *MockMilestoneStore) LastMilestoneId(ctx context.Context) (MilestoneId, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastMilestoneId", ctx)
	ret0, _ := ret[0].(MilestoneId)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LastMilestoneId indicates an expected call of LastMilestoneId.
func (mr *MockMilestoneStoreMockRecorder) LastMilestoneId(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastMilestoneId", reflect.TypeOf((*MockMilestoneStore)(nil).LastMilestoneId), ctx)
}

// PutMilestone mocks base method.
func (m *MockMilestoneStore) PutMilestone(ctx context.Context, milestoneId MilestoneId, milestone *Milestone) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutMilestone", ctx, milestoneId, milestone)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutMilestone indicates an expected call of PutMilestone.
func (mr *MockMilestoneStoreMockRecorder) PutMilestone(ctx, milestoneId, milestone any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutMilestone", reflect.TypeOf((*MockMilestoneStore)(nil).PutMilestone), ctx, milestoneId, milestone)
}

// MockCheckpointReader is a mock of CheckpointReader interface.
type MockCheckpointReader struct {
	ctrl     *gomock.Controller
	recorder *MockCheckpointReaderMockRecorder
}

// MockCheckpointReaderMockRecorder is the mock recorder for MockCheckpointReader.
type MockCheckpointReaderMockRecorder struct {
	mock *MockCheckpointReader
}

// NewMockCheckpointReader creates a new mock instance.
func NewMockCheckpointReader(ctrl *gomock.Controller) *MockCheckpointReader {
	mock := &MockCheckpointReader{ctrl: ctrl}
	mock.recorder = &MockCheckpointReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCheckpointReader) EXPECT() *MockCheckpointReaderMockRecorder {
	return m.recorder
}

// GetCheckpoint mocks base method.
func (m *MockCheckpointReader) GetCheckpoint(ctx context.Context, checkpointId CheckpointId) (*Checkpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCheckpoint", ctx, checkpointId)
	ret0, _ := ret[0].(*Checkpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCheckpoint indicates an expected call of GetCheckpoint.
func (mr *MockCheckpointReaderMockRecorder) GetCheckpoint(ctx, checkpointId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCheckpoint", reflect.TypeOf((*MockCheckpointReader)(nil).GetCheckpoint), ctx, checkpointId)
}

// LastCheckpointId mocks base method.
func (m *MockCheckpointReader) LastCheckpointId(ctx context.Context) (CheckpointId, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastCheckpointId", ctx)
	ret0, _ := ret[0].(CheckpointId)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LastCheckpointId indicates an expected call of LastCheckpointId.
func (mr *MockCheckpointReaderMockRecorder) LastCheckpointId(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastCheckpointId", reflect.TypeOf((*MockCheckpointReader)(nil).LastCheckpointId), ctx)
}

// MockCheckpointWriter is a mock of CheckpointWriter interface.
type MockCheckpointWriter struct {
	ctrl     *gomock.Controller
	recorder *MockCheckpointWriterMockRecorder
}

// MockCheckpointWriterMockRecorder is the mock recorder for MockCheckpointWriter.
type MockCheckpointWriterMockRecorder struct {
	mock *MockCheckpointWriter
}

// NewMockCheckpointWriter creates a new mock instance.
func NewMockCheckpointWriter(ctrl *gomock.Controller) *MockCheckpointWriter {
	mock := &MockCheckpointWriter{ctrl: ctrl}
	mock.recorder = &MockCheckpointWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCheckpointWriter) EXPECT() *MockCheckpointWriterMockRecorder {
	return m.recorder
}

// PutCheckpoint mocks base method.
func (m *MockCheckpointWriter) PutCheckpoint(ctx context.Context, checkpointId CheckpointId, checkpoint *Checkpoint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutCheckpoint", ctx, checkpointId, checkpoint)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutCheckpoint indicates an expected call of PutCheckpoint.
func (mr *MockCheckpointWriterMockRecorder) PutCheckpoint(ctx, checkpointId, checkpoint any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutCheckpoint", reflect.TypeOf((*MockCheckpointWriter)(nil).PutCheckpoint), ctx, checkpointId, checkpoint)
}

// MockCheckpointStore is a mock of CheckpointStore interface.
type MockCheckpointStore struct {
	ctrl     *gomock.Controller
	recorder *MockCheckpointStoreMockRecorder
}

// MockCheckpointStoreMockRecorder is the mock recorder for MockCheckpointStore.
type MockCheckpointStoreMockRecorder struct {
	mock *MockCheckpointStore
}

// NewMockCheckpointStore creates a new mock instance.
func NewMockCheckpointStore(ctrl *gomock.Controller) *MockCheckpointStore {
	mock := &MockCheckpointStore{ctrl: ctrl}
	mock.recorder = &MockCheckpointStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCheckpointStore) EXPECT() *MockCheckpointStoreMockRecorder {
	return m.recorder
}

// GetCheckpoint mocks base method.
func (m *MockCheckpointStore) GetCheckpoint(ctx context.Context, checkpointId CheckpointId) (*Checkpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCheckpoint", ctx, checkpointId)
	ret0, _ := ret[0].(*Checkpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCheckpoint indicates an expected call of GetCheckpoint.
func (mr *MockCheckpointStoreMockRecorder) GetCheckpoint(ctx, checkpointId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCheckpoint", reflect.TypeOf((*MockCheckpointStore)(nil).GetCheckpoint), ctx, checkpointId)
}

// LastCheckpointId mocks base method.
func (m *MockCheckpointStore) LastCheckpointId(ctx context.Context) (CheckpointId, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastCheckpointId", ctx)
	ret0, _ := ret[0].(CheckpointId)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LastCheckpointId indicates an expected call of LastCheckpointId.
func (mr *MockCheckpointStoreMockRecorder) LastCheckpointId(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastCheckpointId", reflect.TypeOf((*MockCheckpointStore)(nil).LastCheckpointId), ctx)
}

// PutCheckpoint mocks base method.
func (m *MockCheckpointStore) PutCheckpoint(ctx context.Context, checkpointId CheckpointId, checkpoint *Checkpoint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutCheckpoint", ctx, checkpointId, checkpoint)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutCheckpoint indicates an expected call of PutCheckpoint.
func (mr *MockCheckpointStoreMockRecorder) PutCheckpoint(ctx, checkpointId, checkpoint any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutCheckpoint", reflect.TypeOf((*MockCheckpointStore)(nil).PutCheckpoint), ctx, checkpointId, checkpoint)
}

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// GetCheckpoint mocks base method.
func (m *MockStore) GetCheckpoint(ctx context.Context, checkpointId CheckpointId) (*Checkpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCheckpoint", ctx, checkpointId)
	ret0, _ := ret[0].(*Checkpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCheckpoint indicates an expected call of GetCheckpoint.
func (mr *MockStoreMockRecorder) GetCheckpoint(ctx, checkpointId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCheckpoint", reflect.TypeOf((*MockStore)(nil).GetCheckpoint), ctx, checkpointId)
}

// GetMilestone mocks base method.
func (m *MockStore) GetMilestone(ctx context.Context, milestoneId MilestoneId) (*Milestone, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMilestone", ctx, milestoneId)
	ret0, _ := ret[0].(*Milestone)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMilestone indicates an expected call of GetMilestone.
func (mr *MockStoreMockRecorder) GetMilestone(ctx, milestoneId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMilestone", reflect.TypeOf((*MockStore)(nil).GetMilestone), ctx, milestoneId)
}

// GetSpan mocks base method.
func (m *MockStore) GetSpan(ctx context.Context, spanId SpanId) (*Span, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSpan", ctx, spanId)
	ret0, _ := ret[0].(*Span)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSpan indicates an expected call of GetSpan.
func (mr *MockStoreMockRecorder) GetSpan(ctx, spanId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSpan", reflect.TypeOf((*MockStore)(nil).GetSpan), ctx, spanId)
}

// LastCheckpointId mocks base method.
func (m *MockStore) LastCheckpointId(ctx context.Context) (CheckpointId, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastCheckpointId", ctx)
	ret0, _ := ret[0].(CheckpointId)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LastCheckpointId indicates an expected call of LastCheckpointId.
func (mr *MockStoreMockRecorder) LastCheckpointId(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastCheckpointId", reflect.TypeOf((*MockStore)(nil).LastCheckpointId), ctx)
}

// LastMilestoneId mocks base method.
func (m *MockStore) LastMilestoneId(ctx context.Context) (MilestoneId, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastMilestoneId", ctx)
	ret0, _ := ret[0].(MilestoneId)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LastMilestoneId indicates an expected call of LastMilestoneId.
func (mr *MockStoreMockRecorder) LastMilestoneId(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastMilestoneId", reflect.TypeOf((*MockStore)(nil).LastMilestoneId), ctx)
}

// LastSpanId mocks base method.
func (m *MockStore) LastSpanId(ctx context.Context) (SpanId, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastSpanId", ctx)
	ret0, _ := ret[0].(SpanId)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LastSpanId indicates an expected call of LastSpanId.
func (mr *MockStoreMockRecorder) LastSpanId(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastSpanId", reflect.TypeOf((*MockStore)(nil).LastSpanId), ctx)
}

// PutCheckpoint mocks base method.
func (m *MockStore) PutCheckpoint(ctx context.Context, checkpointId CheckpointId, checkpoint *Checkpoint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutCheckpoint", ctx, checkpointId, checkpoint)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutCheckpoint indicates an expected call of PutCheckpoint.
func (mr *MockStoreMockRecorder) PutCheckpoint(ctx, checkpointId, checkpoint any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutCheckpoint", reflect.TypeOf((*MockStore)(nil).PutCheckpoint), ctx, checkpointId, checkpoint)
}

// PutMilestone mocks base method.
func (m *MockStore) PutMilestone(ctx context.Context, milestoneId MilestoneId, milestone *Milestone) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutMilestone", ctx, milestoneId, milestone)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutMilestone indicates an expected call of PutMilestone.
func (mr *MockStoreMockRecorder) PutMilestone(ctx, milestoneId, milestone any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutMilestone", reflect.TypeOf((*MockStore)(nil).PutMilestone), ctx, milestoneId, milestone)
}

// PutSpan mocks base method.
func (m *MockStore) PutSpan(ctx context.Context, span *Span) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutSpan", ctx, span)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutSpan indicates an expected call of PutSpan.
func (mr *MockStoreMockRecorder) PutSpan(ctx, span any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutSpan", reflect.TypeOf((*MockStore)(nil).PutSpan), ctx, span)
}

// Mockreader is a mock of reader interface.
type Mockreader struct {
	ctrl     *gomock.Controller
	recorder *MockreaderMockRecorder
}

// MockreaderMockRecorder is the mock recorder for Mockreader.
type MockreaderMockRecorder struct {
	mock *Mockreader
}

// NewMockreader creates a new mock instance.
func NewMockreader(ctrl *gomock.Controller) *Mockreader {
	mock := &Mockreader{ctrl: ctrl}
	mock.recorder = &MockreaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockreader) EXPECT() *MockreaderMockRecorder {
	return m.recorder
}

// BorStartEventID mocks base method.
func (m *Mockreader) BorStartEventID(ctx context.Context, tx kv.Tx, hash common.Hash, blockNum uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BorStartEventID", ctx, tx, hash, blockNum)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BorStartEventID indicates an expected call of BorStartEventID.
func (mr *MockreaderMockRecorder) BorStartEventID(ctx, tx, hash, blockNum any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BorStartEventID", reflect.TypeOf((*Mockreader)(nil).BorStartEventID), ctx, tx, hash, blockNum)
}

// Checkpoint mocks base method.
func (m *Mockreader) Checkpoint(ctx context.Context, tx kv.Getter, checkpointId uint64) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Checkpoint", ctx, tx, checkpointId)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Checkpoint indicates an expected call of Checkpoint.
func (mr *MockreaderMockRecorder) Checkpoint(ctx, tx, checkpointId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Checkpoint", reflect.TypeOf((*Mockreader)(nil).Checkpoint), ctx, tx, checkpointId)
}

// EventLookup mocks base method.
func (m *Mockreader) EventLookup(ctx context.Context, tx kv.Getter, txnHash common.Hash) (uint64, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EventLookup", ctx, tx, txnHash)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EventLookup indicates an expected call of EventLookup.
func (mr *MockreaderMockRecorder) EventLookup(ctx, tx, txnHash any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventLookup", reflect.TypeOf((*Mockreader)(nil).EventLookup), ctx, tx, txnHash)
}

// EventsByBlock mocks base method.
func (m *Mockreader) EventsByBlock(ctx context.Context, tx kv.Tx, hash common.Hash, blockNum uint64) ([]rlp.RawValue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EventsByBlock", ctx, tx, hash, blockNum)
	ret0, _ := ret[0].([]rlp.RawValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EventsByBlock indicates an expected call of EventsByBlock.
func (mr *MockreaderMockRecorder) EventsByBlock(ctx, tx, hash, blockNum any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventsByBlock", reflect.TypeOf((*Mockreader)(nil).EventsByBlock), ctx, tx, hash, blockNum)
}

// LastCheckpointId mocks base method.
func (m *Mockreader) LastCheckpointId(ctx context.Context, tx kv.Tx) (uint64, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastCheckpointId", ctx, tx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LastCheckpointId indicates an expected call of LastCheckpointId.
func (mr *MockreaderMockRecorder) LastCheckpointId(ctx, tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastCheckpointId", reflect.TypeOf((*Mockreader)(nil).LastCheckpointId), ctx, tx)
}

// LastEventId mocks base method.
func (m *Mockreader) LastEventId(ctx context.Context, tx kv.Tx) (uint64, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastEventId", ctx, tx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LastEventId indicates an expected call of LastEventId.
func (mr *MockreaderMockRecorder) LastEventId(ctx, tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastEventId", reflect.TypeOf((*Mockreader)(nil).LastEventId), ctx, tx)
}

// LastFrozenEventId mocks base method.
func (m *Mockreader) LastFrozenEventId() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastFrozenEventId")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// LastFrozenEventId indicates an expected call of LastFrozenEventId.
func (mr *MockreaderMockRecorder) LastFrozenEventId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastFrozenEventId", reflect.TypeOf((*Mockreader)(nil).LastFrozenEventId))
}

// LastFrozenSpanId mocks base method.
func (m *Mockreader) LastFrozenSpanId() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastFrozenSpanId")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// LastFrozenSpanId indicates an expected call of LastFrozenSpanId.
func (mr *MockreaderMockRecorder) LastFrozenSpanId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastFrozenSpanId", reflect.TypeOf((*Mockreader)(nil).LastFrozenSpanId))
}

// LastMilestoneId mocks base method.
func (m *Mockreader) LastMilestoneId(ctx context.Context, tx kv.Tx) (uint64, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastMilestoneId", ctx, tx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LastMilestoneId indicates an expected call of LastMilestoneId.
func (mr *MockreaderMockRecorder) LastMilestoneId(ctx, tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastMilestoneId", reflect.TypeOf((*Mockreader)(nil).LastMilestoneId), ctx, tx)
}

// LastSpanId mocks base method.
func (m *Mockreader) LastSpanId(ctx context.Context, tx kv.Tx) (uint64, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastSpanId", ctx, tx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LastSpanId indicates an expected call of LastSpanId.
func (mr *MockreaderMockRecorder) LastSpanId(ctx, tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastSpanId", reflect.TypeOf((*Mockreader)(nil).LastSpanId), ctx, tx)
}

// Milestone mocks base method.
func (m *Mockreader) Milestone(ctx context.Context, tx kv.Getter, milestoneId uint64) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Milestone", ctx, tx, milestoneId)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Milestone indicates an expected call of Milestone.
func (mr *MockreaderMockRecorder) Milestone(ctx, tx, milestoneId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Milestone", reflect.TypeOf((*Mockreader)(nil).Milestone), ctx, tx, milestoneId)
}

// Span mocks base method.
func (m *Mockreader) Span(ctx context.Context, tx kv.Getter, spanId uint64) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Span", ctx, tx, spanId)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Span indicates an expected call of Span.
func (mr *MockreaderMockRecorder) Span(ctx, tx, spanId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Span", reflect.TypeOf((*Mockreader)(nil).Span), ctx, tx, spanId)
}
