// Code generated by MockGen. DO NOT EDIT.
// Source: client/produtos.go
//
// Generated by this command:
//
//	mockgen -source=client/produtos.go -package=mock_client -destination=test/mock/client/produtos.go
//

// Package mock_client is a generated GoMock package.
package mock_client

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockProduto is a mock of Produto interface.
type MockProduto struct {
	ctrl     *gomock.Controller
	recorder *MockProdutoMockRecorder
}

// MockProdutoMockRecorder is the mock recorder for MockProduto.
type MockProdutoMockRecorder struct {
	mock *MockProduto
}

// NewMockProduto creates a new mock instance.
func NewMockProduto(ctrl *gomock.Controller) *MockProduto {
	mock := &MockProduto{ctrl: ctrl}
	mock.recorder = &MockProdutoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProduto) EXPECT() *MockProdutoMockRecorder {
	return m.recorder
}

// PesquisaPorIDS mocks base method.
func (m *MockProduto) PesquisaPorIDS(ctx context.Context, ids []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PesquisaPorIDS", ctx, ids)
	ret0, _ := ret[0].(error)
	return ret0
}

// PesquisaPorIDS indicates an expected call of PesquisaPorIDS.
func (mr *MockProdutoMockRecorder) PesquisaPorIDS(ctx, ids any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PesquisaPorIDS", reflect.TypeOf((*MockProduto)(nil).PesquisaPorIDS), ctx, ids)
}
