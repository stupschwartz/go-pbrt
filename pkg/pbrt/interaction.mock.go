// Code generated by MockGen. DO NOT EDIT.
// Source: interaction.go

// Package pbrt is a generated GoMock package.
package pbrt

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockInteraction is a mock of Interaction interface
type MockInteraction struct {
	ctrl     *gomock.Controller
	recorder *MockInteractionMockRecorder
}

// MockInteractionMockRecorder is the mock recorder for MockInteraction
type MockInteractionMockRecorder struct {
	mock *MockInteraction
}

// NewMockInteraction creates a new mock instance
func NewMockInteraction(ctrl *gomock.Controller) *MockInteraction {
	mock := &MockInteraction{ctrl: ctrl}
	mock.recorder = &MockInteractionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInteraction) EXPECT() *MockInteractionMockRecorder {
	return m.recorder
}

// SpawnRay mocks base method
func (m *MockInteraction) SpawnRay(direction *Vector3f) *Ray {
	ret := m.ctrl.Call(m, "SpawnRay", direction)
	ret0, _ := ret[0].(*Ray)
	return ret0
}

// SpawnRay indicates an expected call of SpawnRay
func (mr *MockInteractionMockRecorder) SpawnRay(direction interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpawnRay", reflect.TypeOf((*MockInteraction)(nil).SpawnRay), direction)
}

// SpawnRayToPoint mocks base method
func (m *MockInteraction) SpawnRayToPoint(p *Point3f) *Ray {
	ret := m.ctrl.Call(m, "SpawnRayToPoint", p)
	ret0, _ := ret[0].(*Ray)
	return ret0
}

// SpawnRayToPoint indicates an expected call of SpawnRayToPoint
func (mr *MockInteractionMockRecorder) SpawnRayToPoint(p interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpawnRayToPoint", reflect.TypeOf((*MockInteraction)(nil).SpawnRayToPoint), p)
}

// SpawnRayToInteraction mocks base method
func (m *MockInteraction) SpawnRayToInteraction(to Interaction) *Ray {
	ret := m.ctrl.Call(m, "SpawnRayToInteraction", to)
	ret0, _ := ret[0].(*Ray)
	return ret0
}

// SpawnRayToInteraction indicates an expected call of SpawnRayToInteraction
func (mr *MockInteractionMockRecorder) SpawnRayToInteraction(to interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpawnRayToInteraction", reflect.TypeOf((*MockInteraction)(nil).SpawnRayToInteraction), to)
}

// GetPoint mocks base method
func (m *MockInteraction) GetPoint() *Point3f {
	ret := m.ctrl.Call(m, "GetPoint")
	ret0, _ := ret[0].(*Point3f)
	return ret0
}

// GetPoint indicates an expected call of GetPoint
func (mr *MockInteractionMockRecorder) GetPoint() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPoint", reflect.TypeOf((*MockInteraction)(nil).GetPoint))
}

// GetPointError mocks base method
func (m *MockInteraction) GetPointError() *Vector3f {
	ret := m.ctrl.Call(m, "GetPointError")
	ret0, _ := ret[0].(*Vector3f)
	return ret0
}

// GetPointError indicates an expected call of GetPointError
func (mr *MockInteractionMockRecorder) GetPointError() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPointError", reflect.TypeOf((*MockInteraction)(nil).GetPointError))
}

// GetTime mocks base method
func (m *MockInteraction) GetTime() float64 {
	ret := m.ctrl.Call(m, "GetTime")
	ret0, _ := ret[0].(float64)
	return ret0
}

// GetTime indicates an expected call of GetTime
func (mr *MockInteractionMockRecorder) GetTime() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTime", reflect.TypeOf((*MockInteraction)(nil).GetTime))
}

// GetNormal mocks base method
func (m *MockInteraction) GetNormal() *Normal3f {
	ret := m.ctrl.Call(m, "GetNormal")
	ret0, _ := ret[0].(*Normal3f)
	return ret0
}

// GetNormal indicates an expected call of GetNormal
func (mr *MockInteractionMockRecorder) GetNormal() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNormal", reflect.TypeOf((*MockInteraction)(nil).GetNormal))
}

// GetMediumAccessor mocks base method
func (m *MockInteraction) GetMediumAccessor() *MediumAccessor {
	ret := m.ctrl.Call(m, "GetMediumAccessor")
	ret0, _ := ret[0].(*MediumAccessor)
	return ret0
}

// GetMediumAccessor indicates an expected call of GetMediumAccessor
func (mr *MockInteractionMockRecorder) GetMediumAccessor() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMediumAccessor", reflect.TypeOf((*MockInteraction)(nil).GetMediumAccessor))
}

// SetMediumAccessor mocks base method
func (m *MockInteraction) SetMediumAccessor(accessor *MediumAccessor) {
	m.ctrl.Call(m, "SetMediumAccessor", accessor)
}

// SetMediumAccessor indicates an expected call of SetMediumAccessor
func (mr *MockInteractionMockRecorder) SetMediumAccessor(accessor interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMediumAccessor", reflect.TypeOf((*MockInteraction)(nil).SetMediumAccessor), accessor)
}
