// Code generated by MockGen. DO NOT EDIT.
// Source: primitive.go

// Package pbrt is a generated GoMock package.
package pbrt

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPrimitive is a mock of Primitive interface
type MockPrimitive struct {
	ctrl     *gomock.Controller
	recorder *MockPrimitiveMockRecorder
}

// MockPrimitiveMockRecorder is the mock recorder for MockPrimitive
type MockPrimitiveMockRecorder struct {
	mock *MockPrimitive
}

// NewMockPrimitive creates a new mock instance
func NewMockPrimitive(ctrl *gomock.Controller) *MockPrimitive {
	mock := &MockPrimitive{ctrl: ctrl}
	mock.recorder = &MockPrimitiveMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPrimitive) EXPECT() *MockPrimitiveMockRecorder {
	return m.recorder
}

// Intersect mocks base method
func (m *MockPrimitive) Intersect(r *Ray, si *SurfaceInteraction) bool {
	ret := m.ctrl.Call(m, "Intersect", r, si)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Intersect indicates an expected call of Intersect
func (mr *MockPrimitiveMockRecorder) Intersect(r, si interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersect", reflect.TypeOf((*MockPrimitive)(nil).Intersect), r, si)
}

// IntersectP mocks base method
func (m *MockPrimitive) IntersectP(r *Ray) bool {
	ret := m.ctrl.Call(m, "IntersectP", r)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IntersectP indicates an expected call of IntersectP
func (mr *MockPrimitiveMockRecorder) IntersectP(r interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IntersectP", reflect.TypeOf((*MockPrimitive)(nil).IntersectP), r)
}

// GetAreaLight mocks base method
func (m *MockPrimitive) GetAreaLight() AreaLighter {
	ret := m.ctrl.Call(m, "GetAreaLight")
	ret0, _ := ret[0].(AreaLighter)
	return ret0
}

// GetAreaLight indicates an expected call of GetAreaLight
func (mr *MockPrimitiveMockRecorder) GetAreaLight() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAreaLight", reflect.TypeOf((*MockPrimitive)(nil).GetAreaLight))
}

// GetMaterial mocks base method
func (m *MockPrimitive) GetMaterial() Material {
	ret := m.ctrl.Call(m, "GetMaterial")
	ret0, _ := ret[0].(Material)
	return ret0
}

// GetMaterial indicates an expected call of GetMaterial
func (mr *MockPrimitiveMockRecorder) GetMaterial() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaterial", reflect.TypeOf((*MockPrimitive)(nil).GetMaterial))
}

// ComputeScatteringFunctions mocks base method
func (m *MockPrimitive) ComputeScatteringFunctions(si *SurfaceInteraction, mode TransportMode, allowMultipleLobes bool) {
	m.ctrl.Call(m, "ComputeScatteringFunctions", si, mode, allowMultipleLobes)
}

// ComputeScatteringFunctions indicates an expected call of ComputeScatteringFunctions
func (mr *MockPrimitiveMockRecorder) ComputeScatteringFunctions(si, mode, allowMultipleLobes interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComputeScatteringFunctions", reflect.TypeOf((*MockPrimitive)(nil).ComputeScatteringFunctions), si, mode, allowMultipleLobes)
}

// WorldBound mocks base method
func (m *MockPrimitive) WorldBound() *Bounds3 {
	ret := m.ctrl.Call(m, "WorldBound")
	ret0, _ := ret[0].(*Bounds3)
	return ret0
}

// WorldBound indicates an expected call of WorldBound
func (mr *MockPrimitiveMockRecorder) WorldBound() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WorldBound", reflect.TypeOf((*MockPrimitive)(nil).WorldBound))
}

// MockAggregate is a mock of Aggregate interface
type MockAggregate struct {
	ctrl     *gomock.Controller
	recorder *MockAggregateMockRecorder
}

// MockAggregateMockRecorder is the mock recorder for MockAggregate
type MockAggregateMockRecorder struct {
	mock *MockAggregate
}

// NewMockAggregate creates a new mock instance
func NewMockAggregate(ctrl *gomock.Controller) *MockAggregate {
	mock := &MockAggregate{ctrl: ctrl}
	mock.recorder = &MockAggregateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAggregate) EXPECT() *MockAggregateMockRecorder {
	return m.recorder
}

// Intersect mocks base method
func (m *MockAggregate) Intersect(r *Ray, si *SurfaceInteraction) bool {
	ret := m.ctrl.Call(m, "Intersect", r, si)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Intersect indicates an expected call of Intersect
func (mr *MockAggregateMockRecorder) Intersect(r, si interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersect", reflect.TypeOf((*MockAggregate)(nil).Intersect), r, si)
}

// IntersectP mocks base method
func (m *MockAggregate) IntersectP(r *Ray) bool {
	ret := m.ctrl.Call(m, "IntersectP", r)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IntersectP indicates an expected call of IntersectP
func (mr *MockAggregateMockRecorder) IntersectP(r interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IntersectP", reflect.TypeOf((*MockAggregate)(nil).IntersectP), r)
}

// GetAreaLight mocks base method
func (m *MockAggregate) GetAreaLight() AreaLighter {
	ret := m.ctrl.Call(m, "GetAreaLight")
	ret0, _ := ret[0].(AreaLighter)
	return ret0
}

// GetAreaLight indicates an expected call of GetAreaLight
func (mr *MockAggregateMockRecorder) GetAreaLight() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAreaLight", reflect.TypeOf((*MockAggregate)(nil).GetAreaLight))
}

// GetMaterial mocks base method
func (m *MockAggregate) GetMaterial() Material {
	ret := m.ctrl.Call(m, "GetMaterial")
	ret0, _ := ret[0].(Material)
	return ret0
}

// GetMaterial indicates an expected call of GetMaterial
func (mr *MockAggregateMockRecorder) GetMaterial() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaterial", reflect.TypeOf((*MockAggregate)(nil).GetMaterial))
}

// ComputeScatteringFunctions mocks base method
func (m *MockAggregate) ComputeScatteringFunctions(si *SurfaceInteraction, mode TransportMode, allowMultipleLobes bool) {
	m.ctrl.Call(m, "ComputeScatteringFunctions", si, mode, allowMultipleLobes)
}

// ComputeScatteringFunctions indicates an expected call of ComputeScatteringFunctions
func (mr *MockAggregateMockRecorder) ComputeScatteringFunctions(si, mode, allowMultipleLobes interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComputeScatteringFunctions", reflect.TypeOf((*MockAggregate)(nil).ComputeScatteringFunctions), si, mode, allowMultipleLobes)
}

// WorldBound mocks base method
func (m *MockAggregate) WorldBound() *Bounds3 {
	ret := m.ctrl.Call(m, "WorldBound")
	ret0, _ := ret[0].(*Bounds3)
	return ret0
}

// WorldBound indicates an expected call of WorldBound
func (mr *MockAggregateMockRecorder) WorldBound() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WorldBound", reflect.TypeOf((*MockAggregate)(nil).WorldBound))
}
