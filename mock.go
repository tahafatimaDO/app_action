// Code generated by MockGen. DO NOT EDIT.
// Source: main.go

// Package main is a generated GoMock package.
package main

import (
	reflect "reflect"

	parser_struct "github.com/digitalocean/app_action/internal/parser_struct"
	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

// MockDoctlClient is a mock of DoctlClient interface.
type MockDoctlClient struct {
	ctrl     *gomock.Controller
	recorder *MockDoctlClientMockRecorder
}

// MockDoctlClientMockRecorder is the mock recorder for MockDoctlClient.
type MockDoctlClientMockRecorder struct {
	mock *MockDoctlClient
}

// NewMockDoctlClient creates a new mock instance.
func NewMockDoctlClient(ctrl *gomock.Controller) *MockDoctlClient {
	mock := &MockDoctlClient{ctrl: ctrl}
	mock.recorder = &MockDoctlClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDoctlClient) EXPECT() *MockDoctlClientMockRecorder {
	return m.recorder
}

// CreateDeployments mocks base method.
func (m *MockDoctlClient) CreateDeployments(appID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDeployments", appID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDeployments indicates an expected call of CreateDeployments.
func (mr *MockDoctlClientMockRecorder) CreateDeployments(appID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDeployments", reflect.TypeOf((*MockDoctlClient)(nil).CreateDeployments), appID)
}

// Deploy mocks base method.
func (m *MockDoctlClient) Deploy(input, appName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deploy", input, appName)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deploy indicates an expected call of Deploy.
func (mr *MockDoctlClientMockRecorder) Deploy(input, appName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deploy", reflect.TypeOf((*MockDoctlClient)(nil).Deploy), input, appName)
}

// IsDeployed mocks base method.
func (m *MockDoctlClient) IsDeployed(appID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDeployed", appID)
	ret0, _ := ret[0].(error)
	return ret0
}

// IsDeployed indicates an expected call of IsDeployed.
func (mr *MockDoctlClientMockRecorder) IsDeployed(appID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDeployed", reflect.TypeOf((*MockDoctlClient)(nil).IsDeployed), appID)
}

// ListDeployments mocks base method.
func (m *MockDoctlClient) ListDeployments(appID string) ([]godo.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDeployments", appID)
	ret0, _ := ret[0].([]godo.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDeployments indicates an expected call of ListDeployments.
func (mr *MockDoctlClientMockRecorder) ListDeployments(appID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeployments", reflect.TypeOf((*MockDoctlClient)(nil).ListDeployments), appID)
}

// RetrieveActiveDeployment mocks base method.
func (m *MockDoctlClient) RetrieveActiveDeployment(deploymentID, appID, input string) ([]parser_struct.UpdatedRepo, *godo.AppSpec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveActiveDeployment", deploymentID, appID, input)
	ret0, _ := ret[0].([]parser_struct.UpdatedRepo)
	ret1, _ := ret[1].(*godo.AppSpec)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RetrieveActiveDeployment indicates an expected call of RetrieveActiveDeployment.
func (mr *MockDoctlClientMockRecorder) RetrieveActiveDeployment(deploymentID, appID, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveActiveDeployment", reflect.TypeOf((*MockDoctlClient)(nil).RetrieveActiveDeployment), deploymentID, appID, input)
}

// RetrieveActiveDeploymentID mocks base method.
func (m *MockDoctlClient) RetrieveActiveDeploymentID(appID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveActiveDeploymentID", appID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveActiveDeploymentID indicates an expected call of RetrieveActiveDeploymentID.
func (mr *MockDoctlClientMockRecorder) RetrieveActiveDeploymentID(appID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveActiveDeploymentID", reflect.TypeOf((*MockDoctlClient)(nil).RetrieveActiveDeploymentID), appID)
}

// RetrieveAppID mocks base method.
func (m *MockDoctlClient) RetrieveAppID(appName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveAppID", appName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveAppID indicates an expected call of RetrieveAppID.
func (mr *MockDoctlClientMockRecorder) RetrieveAppID(appName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveAppID", reflect.TypeOf((*MockDoctlClient)(nil).RetrieveAppID), appName)
}

// RetrieveFromDigitalocean mocks base method.
func (m *MockDoctlClient) RetrieveFromDigitalocean() ([]godo.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveFromDigitalocean")
	ret0, _ := ret[0].([]godo.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveFromDigitalocean indicates an expected call of RetrieveFromDigitalocean.
func (mr *MockDoctlClientMockRecorder) RetrieveFromDigitalocean() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveFromDigitalocean", reflect.TypeOf((*MockDoctlClient)(nil).RetrieveFromDigitalocean))
}

// UpdateAppPlatformAppSpec mocks base method.
func (m *MockDoctlClient) UpdateAppPlatformAppSpec(tmpfile, appID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAppPlatformAppSpec", tmpfile, appID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAppPlatformAppSpec indicates an expected call of UpdateAppPlatformAppSpec.
func (mr *MockDoctlClientMockRecorder) UpdateAppPlatformAppSpec(tmpfile, appID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAppPlatformAppSpec", reflect.TypeOf((*MockDoctlClient)(nil).UpdateAppPlatformAppSpec), tmpfile, appID)
}
