// Code generated by MockGen. DO NOT EDIT.
// Source: core/domain/product.go

// Package mocks is a generated GoMock package.
package mocks

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/marc/go-clean-example/core/domain"
	dto "github.com/marc/go-clean-example/core/dto"
)

// MockProductService is a mock of ProductService interface.
type MockProductService struct {
	ctrl     *gomock.Controller
	recorder *MockProductServiceMockRecorder
}

// MockProductServiceMockRecorder is the mock recorder for MockProductService.
type MockProductServiceMockRecorder struct {
	mock *MockProductService
}

// NewMockProductService creates a new mock instance.
func NewMockProductService(ctrl *gomock.Controller) *MockProductService {
	mock := &MockProductService{ctrl: ctrl}
	mock.recorder = &MockProductServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductService) EXPECT() *MockProductServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProductService) Create(response http.ResponseWriter, request *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Create", response, request)
}

// Create indicates an expected call of Create.
func (mr *MockProductServiceMockRecorder) Create(response, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductService)(nil).Create), response, request)
}

// Delete mocks base method.
func (m *MockProductService) Delete(response http.ResponseWriter, request *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", response, request)
}

// Delete indicates an expected call of Delete.
func (mr *MockProductServiceMockRecorder) Delete(response, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProductService)(nil).Delete), response, request)
}

// Fetch mocks base method.
func (m *MockProductService) Fetch(response http.ResponseWriter, request *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Fetch", response, request)
}

// Fetch indicates an expected call of Fetch.
func (mr *MockProductServiceMockRecorder) Fetch(response, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockProductService)(nil).Fetch), response, request)
}

// Update mocks base method.
func (m *MockProductService) Update(response http.ResponseWriter, request *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Update", response, request)
}

// Update indicates an expected call of Update.
func (mr *MockProductServiceMockRecorder) Update(response, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProductService)(nil).Update), response, request)
}

// MockProductUseCase is a mock of ProductUseCase interface.
type MockProductUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockProductUseCaseMockRecorder
}

// MockProductUseCaseMockRecorder is the mock recorder for MockProductUseCase.
type MockProductUseCaseMockRecorder struct {
	mock *MockProductUseCase
}

// NewMockProductUseCase creates a new mock instance.
func NewMockProductUseCase(ctrl *gomock.Controller) *MockProductUseCase {
	mock := &MockProductUseCase{ctrl: ctrl}
	mock.recorder = &MockProductUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductUseCase) EXPECT() *MockProductUseCaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProductUseCase) Create(productRequest *dto.CreateProductRequest) (*domain.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", productRequest)
	ret0, _ := ret[0].(*domain.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProductUseCaseMockRecorder) Create(productRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductUseCase)(nil).Create), productRequest)
}

// Delete mocks base method.
func (m *MockProductUseCase) Delete(id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockProductUseCaseMockRecorder) Delete(id uint64) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProductUseCase)(nil).Delete), id)
}

// Fetch mocks base method.
func (m *MockProductUseCase) Fetch(paginationRequest *dto.PaginationRequestParms) (*domain.Pagination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", paginationRequest)
	ret0, _ := ret[0].(*domain.Pagination)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch.
func (mr *MockProductUseCaseMockRecorder) Fetch(paginationRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockProductUseCase)(nil).Fetch), paginationRequest)
}

// Update mocks base method.
func (m *MockProductUseCase) Update(productRequest *dto.UpdateProductRequest) (*domain.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", productRequest)
	ret0, _ := ret[0].(*domain.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockProductUseCaseMockRecorder) Update(productRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProductUseCase)(nil).Update), productRequest)
}

// MockProductRepository is a mock of ProductRepository interface.
type MockProductRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProductRepositoryMockRecorder
}

// MockProductRepositoryMockRecorder is the mock recorder for MockProductRepository.
type MockProductRepositoryMockRecorder struct {
	mock *MockProductRepository
}

// NewMockProductRepository creates a new mock instance.
func NewMockProductRepository(ctrl *gomock.Controller) *MockProductRepository {
	mock := &MockProductRepository{ctrl: ctrl}
	mock.recorder = &MockProductRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductRepository) EXPECT() *MockProductRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProductRepository) Create(productRequest *dto.CreateProductRequest) (*domain.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", productRequest)
	ret0, _ := ret[0].(*domain.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProductRepositoryMockRecorder) Create(productRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductRepository)(nil).Create), productRequest)
}

// Delete mocks base method.
func (m *MockProductRepository) Delete(id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockProductRepositoryMockRecorder) Delete(id uint64) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProductRepository)(nil).Delete), id)
}

// Fetch mocks base method.
func (m *MockProductRepository) Fetch(paginationRequest *dto.PaginationRequestParms) (*domain.Pagination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", paginationRequest)
	ret0, _ := ret[0].(*domain.Pagination)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch.
func (mr *MockProductRepositoryMockRecorder) Fetch(paginationRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockProductRepository)(nil).Fetch), paginationRequest)
}

// Update mocks base method.
func (m *MockProductRepository) Update(productRequest *dto.UpdateProductRequest) (*domain.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", productRequest)
	ret0, _ := ret[0].(*domain.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockProductRepositoryMockRecorder) Update(productRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProductRepository)(nil).Update), productRequest)
}
