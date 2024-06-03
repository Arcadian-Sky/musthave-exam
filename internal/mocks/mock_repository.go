// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/repository.go

// Package mock_repository is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	model "github.com/Arcadian-Sky/musthave-exam/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// AddOrder mocks base method.
func (m *MockRepository) AddOrder(ctx context.Context, userID int64, number string) (*model.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOrder", ctx, userID, number)
	ret0, _ := ret[0].(*model.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddOrder indicates an expected call of AddOrder.
func (mr *MockRepositoryMockRecorder) AddOrder(ctx, userID, number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrder", reflect.TypeOf((*MockRepository)(nil).AddOrder), ctx, userID, number)
}

// CheckUserExisis mocks base method.
func (m *MockRepository) CheckUserExisis(ctx context.Context, user model.User) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserExisis", ctx, user)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUserExisis indicates an expected call of CheckUserExisis.
func (mr *MockRepositoryMockRecorder) CheckUserExisis(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserExisis", reflect.TypeOf((*MockRepository)(nil).CheckUserExisis), ctx, user)
}

// GetBalance mocks base method.
func (m *MockRepository) GetBalance(ctx context.Context, userID int64) (model.Balance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", ctx, userID)
	ret0, _ := ret[0].(model.Balance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockRepositoryMockRecorder) GetBalance(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockRepository)(nil).GetBalance), ctx, userID)
}

// GetNewTransactions mocks base method.
func (m *MockRepository) GetNewTransactions(ctx context.Context) ([]model.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNewTransactions", ctx)
	ret0, _ := ret[0].([]model.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNewTransactions indicates an expected call of GetNewTransactions.
func (mr *MockRepositoryMockRecorder) GetNewTransactions(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNewTransactions", reflect.TypeOf((*MockRepository)(nil).GetNewTransactions), ctx)
}

// GetOrders mocks base method.
func (m *MockRepository) GetOrders(ctx context.Context, userID int64, act string) ([]model.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrders", ctx, userID, act)
	ret0, _ := ret[0].([]model.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrders indicates an expected call of GetOrders.
func (mr *MockRepositoryMockRecorder) GetOrders(ctx, userID, act interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrders", reflect.TypeOf((*MockRepository)(nil).GetOrders), ctx, userID, act)
}

// GetUser mocks base method.
func (m *MockRepository) GetUser(ctx context.Context, id int64) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, id)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockRepositoryMockRecorder) GetUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockRepository)(nil).GetUser), ctx, id)
}

// LoginUser mocks base method.
func (m *MockRepository) LoginUser(ctx context.Context, user model.User) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginUser", ctx, user)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginUser indicates an expected call of LoginUser.
func (mr *MockRepositoryMockRecorder) LoginUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginUser", reflect.TypeOf((*MockRepository)(nil).LoginUser), ctx, user)
}

// RegisterUser mocks base method.
func (m *MockRepository) RegisterUser(ctx context.Context, user model.User) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", ctx, user)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockRepositoryMockRecorder) RegisterUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockRepository)(nil).RegisterUser), ctx, user)
}

// UpdateTransactionStatus mocks base method.
func (m *MockRepository) UpdateTransactionStatus(ctx context.Context, transactionID, status string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTransactionStatus", ctx, transactionID, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTransactionStatus indicates an expected call of UpdateTransactionStatus.
func (mr *MockRepositoryMockRecorder) UpdateTransactionStatus(ctx, transactionID, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTransactionStatus", reflect.TypeOf((*MockRepository)(nil).UpdateTransactionStatus), ctx, transactionID, status)
}

// UpdateTransactionStatusAndAccrual mocks base method.
func (m *MockRepository) UpdateTransactionStatusAndAccrual(ctx context.Context, transactionID, status string, accrual float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTransactionStatusAndAccrual", ctx, transactionID, status, accrual)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTransactionStatusAndAccrual indicates an expected call of UpdateTransactionStatusAndAccrual.
func (mr *MockRepositoryMockRecorder) UpdateTransactionStatusAndAccrual(ctx, transactionID, status, accrual interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTransactionStatusAndAccrual", reflect.TypeOf((*MockRepository)(nil).UpdateTransactionStatusAndAccrual), ctx, transactionID, status, accrual)
}

// Withdraw mocks base method.
func (m *MockRepository) Withdraw(ctx context.Context, userID int64, order string, sum float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Withdraw", ctx, userID, order, sum)
	ret0, _ := ret[0].(error)
	return ret0
}

// Withdraw indicates an expected call of Withdraw.
func (mr *MockRepositoryMockRecorder) Withdraw(ctx, userID, order, sum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Withdraw", reflect.TypeOf((*MockRepository)(nil).Withdraw), ctx, userID, order, sum)
}
