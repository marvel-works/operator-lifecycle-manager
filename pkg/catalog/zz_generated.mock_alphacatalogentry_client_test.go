// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/client/alphacatalogentry_client.go

// Package catalog is a generated GoMock package.
package catalog

import (
	v1alpha1 "github.com/coreos-inc/alm/pkg/apis/alphacatalogentry/v1alpha1"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAlphaCatalogEntryInterface is a mock of AlphaCatalogEntryInterface interface
type MockAlphaCatalogEntryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAlphaCatalogEntryInterfaceMockRecorder
}

// MockAlphaCatalogEntryInterfaceMockRecorder is the mock recorder for MockAlphaCatalogEntryInterface
type MockAlphaCatalogEntryInterfaceMockRecorder struct {
	mock *MockAlphaCatalogEntryInterface
}

// NewMockAlphaCatalogEntryInterface creates a new mock instance
func NewMockAlphaCatalogEntryInterface(ctrl *gomock.Controller) *MockAlphaCatalogEntryInterface {
	mock := &MockAlphaCatalogEntryInterface{ctrl: ctrl}
	mock.recorder = &MockAlphaCatalogEntryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAlphaCatalogEntryInterface) EXPECT() *MockAlphaCatalogEntryInterfaceMockRecorder {
	return m.recorder
}

// UpdateEntry mocks base method
func (m *MockAlphaCatalogEntryInterface) UpdateEntry(csv *v1alpha1.AlphaCatalogEntry) (*v1alpha1.AlphaCatalogEntry, error) {
	ret := m.ctrl.Call(m, "UpdateEntry", csv)
	ret0, _ := ret[0].(*v1alpha1.AlphaCatalogEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEntry indicates an expected call of UpdateEntry
func (mr *MockAlphaCatalogEntryInterfaceMockRecorder) UpdateEntry(csv interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEntry", reflect.TypeOf((*MockAlphaCatalogEntryInterface)(nil).UpdateEntry), csv)
}