// Code generated by MockGen. DO NOT EDIT.
// Source: kubectl.go

// Package kubernetes is a generated GoMock package.
package kubernetes

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
)

// MockKubernetesClient is a mock of KubernetesClient interface.
type MockKubernetesClient struct {
	ctrl     *gomock.Controller
	recorder *MockKubernetesClientMockRecorder
}

// MockKubernetesClientMockRecorder is the mock recorder for MockKubernetesClient.
type MockKubernetesClientMockRecorder struct {
	mock *MockKubernetesClient
}

// NewMockKubernetesClient creates a new mock instance.
func NewMockKubernetesClient(ctrl *gomock.Controller) *MockKubernetesClient {
	mock := &MockKubernetesClient{ctrl: ctrl}
	mock.recorder = &MockKubernetesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKubernetesClient) EXPECT() *MockKubernetesClientMockRecorder {
	return m.recorder
}

// CreateConfigMap mocks base method.
func (m *MockKubernetesClient) CreateConfigMap(namespace string, configMap *v1.ConfigMap, dryrun bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateConfigMap", namespace, configMap, dryrun)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateConfigMap indicates an expected call of CreateConfigMap.
func (mr *MockKubernetesClientMockRecorder) CreateConfigMap(namespace, configMap, dryrun interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateConfigMap", reflect.TypeOf((*MockKubernetesClient)(nil).CreateConfigMap), namespace, configMap, dryrun)
}

// GetConfigMap mocks base method.
func (m *MockKubernetesClient) GetConfigMap(namespace, configMap string) (*v1.ConfigMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfigMap", namespace, configMap)
	ret0, _ := ret[0].(*v1.ConfigMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfigMap indicates an expected call of GetConfigMap.
func (mr *MockKubernetesClientMockRecorder) GetConfigMap(namespace, configMap interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfigMap", reflect.TypeOf((*MockKubernetesClient)(nil).GetConfigMap), namespace, configMap)
}

// GetPod mocks base method.
func (m *MockKubernetesClient) GetPod(namespace string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPod", namespace)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPod indicates an expected call of GetPod.
func (mr *MockKubernetesClientMockRecorder) GetPod(namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPod", reflect.TypeOf((*MockKubernetesClient)(nil).GetPod), namespace)
}

// UpdateAPIServerURL mocks base method.
func (m *MockKubernetesClient) UpdateAPIServerURL(apiServerURL string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAPIServerURL", apiServerURL)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAPIServerURL indicates an expected call of UpdateAPIServerURL.
func (mr *MockKubernetesClientMockRecorder) UpdateAPIServerURL(apiServerURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAPIServerURL", reflect.TypeOf((*MockKubernetesClient)(nil).UpdateAPIServerURL), apiServerURL)
}

// UpdateConfigMap mocks base method.
func (m *MockKubernetesClient) UpdateConfigMap(namespace string, configMap *v1.ConfigMap, dryrun bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateConfigMap", namespace, configMap, dryrun)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateConfigMap indicates an expected call of UpdateConfigMap.
func (mr *MockKubernetesClientMockRecorder) UpdateConfigMap(namespace, configMap, dryrun interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConfigMap", reflect.TypeOf((*MockKubernetesClient)(nil).UpdateConfigMap), namespace, configMap, dryrun)
}
