// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kastenhq/kubestr/pkg/csi (interfaces: ApplicationCreator)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/kastenhq/kubestr/pkg/csi/types"
	v1 "k8s.io/api/core/v1"
)

// MockApplicationCreator is a mock of ApplicationCreator interface.
type MockApplicationCreator struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationCreatorMockRecorder
}

// MockApplicationCreatorMockRecorder is the mock recorder for MockApplicationCreator.
type MockApplicationCreatorMockRecorder struct {
	mock *MockApplicationCreator
}

// NewMockApplicationCreator creates a new mock instance.
func NewMockApplicationCreator(ctrl *gomock.Controller) *MockApplicationCreator {
	mock := &MockApplicationCreator{ctrl: ctrl}
	mock.recorder = &MockApplicationCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationCreator) EXPECT() *MockApplicationCreatorMockRecorder {
	return m.recorder
}

// CreatePVC mocks base method.
func (m *MockApplicationCreator) CreatePVC(arg0 context.Context, arg1 *types.CreatePVCArgs) (*v1.PersistentVolumeClaim, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePVC", arg0, arg1)
	ret0, _ := ret[0].(*v1.PersistentVolumeClaim)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePVC indicates an expected call of CreatePVC.
func (mr *MockApplicationCreatorMockRecorder) CreatePVC(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePVC", reflect.TypeOf((*MockApplicationCreator)(nil).CreatePVC), arg0, arg1)
}

// CreatePod mocks base method.
func (m *MockApplicationCreator) CreatePod(arg0 context.Context, arg1 *types.CreatePodArgs) (*v1.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePod", arg0, arg1)
	ret0, _ := ret[0].(*v1.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePod indicates an expected call of CreatePod.
func (mr *MockApplicationCreatorMockRecorder) CreatePod(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePod", reflect.TypeOf((*MockApplicationCreator)(nil).CreatePod), arg0, arg1)
}

// WaitForPodReady mocks base method.
func (m *MockApplicationCreator) WaitForPodReady(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForPodReady", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForPodReady indicates an expected call of WaitForPodReady.
func (mr *MockApplicationCreatorMockRecorder) WaitForPodReady(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForPodReady", reflect.TypeOf((*MockApplicationCreator)(nil).WaitForPodReady), arg0, arg1, arg2)
}
