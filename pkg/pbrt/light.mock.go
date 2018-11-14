// Code generated by MockGen. DO NOT EDIT.
// Source: light.go

// Package pbrt is a generated GoMock package.
package pbrt

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockLight is a mock of Light interface
type MockLight struct {
	ctrl     *gomock.Controller
	recorder *MockLightMockRecorder
}

// MockLightMockRecorder is the mock recorder for MockLight
type MockLightMockRecorder struct {
	mock *MockLight
}

// NewMockLight creates a new mock instance
func NewMockLight(ctrl *gomock.Controller) *MockLight {
	mock := &MockLight{ctrl: ctrl}
	mock.recorder = &MockLightMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLight) EXPECT() *MockLightMockRecorder {
	return m.recorder
}

// GetFlags mocks base method
func (m *MockLight) GetFlags() LightFlag {
	ret := m.ctrl.Call(m, "GetFlags")
	ret0, _ := ret[0].(LightFlag)
	return ret0
}

// GetFlags indicates an expected call of GetFlags
func (mr *MockLightMockRecorder) GetFlags() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFlags", reflect.TypeOf((*MockLight)(nil).GetFlags))
}

// GetSamples mocks base method
func (m *MockLight) GetSamples() int {
	ret := m.ctrl.Call(m, "GetSamples")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetSamples indicates an expected call of GetSamples
func (mr *MockLightMockRecorder) GetSamples() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSamples", reflect.TypeOf((*MockLight)(nil).GetSamples))
}

// Preprocess mocks base method
func (m *MockLight) Preprocess() {
	m.ctrl.Call(m, "Preprocess")
}

// Preprocess indicates an expected call of Preprocess
func (mr *MockLightMockRecorder) Preprocess() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Preprocess", reflect.TypeOf((*MockLight)(nil).Preprocess))
}

// SampleLi mocks base method
func (m *MockLight) SampleLi(ref Interaction, u *Point2f) (Spectrum, *Vector3f, float64, *VisibilityTester) {
	ret := m.ctrl.Call(m, "SampleLi", ref, u)
	ret0, _ := ret[0].(Spectrum)
	ret1, _ := ret[1].(*Vector3f)
	ret2, _ := ret[2].(float64)
	ret3, _ := ret[3].(*VisibilityTester)
	return ret0, ret1, ret2, ret3
}

// SampleLi indicates an expected call of SampleLi
func (mr *MockLightMockRecorder) SampleLi(ref, u interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SampleLi", reflect.TypeOf((*MockLight)(nil).SampleLi), ref, u)
}

// Power mocks base method
func (m *MockLight) Power() Spectrum {
	ret := m.ctrl.Call(m, "Power")
	ret0, _ := ret[0].(Spectrum)
	return ret0
}

// Power indicates an expected call of Power
func (mr *MockLightMockRecorder) Power() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Power", reflect.TypeOf((*MockLight)(nil).Power))
}

// Le mocks base method
func (m *MockLight) Le(r *RayDifferential) Spectrum {
	ret := m.ctrl.Call(m, "Le", r)
	ret0, _ := ret[0].(Spectrum)
	return ret0
}

// Le indicates an expected call of Le
func (mr *MockLightMockRecorder) Le(r interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Le", reflect.TypeOf((*MockLight)(nil).Le), r)
}

// PdfLi mocks base method
func (m *MockLight) PdfLi(ref Interaction) (float64, *Vector3f) {
	ret := m.ctrl.Call(m, "PdfLi", ref)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(*Vector3f)
	return ret0, ret1
}

// PdfLi indicates an expected call of PdfLi
func (mr *MockLightMockRecorder) PdfLi(ref interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PdfLi", reflect.TypeOf((*MockLight)(nil).PdfLi), ref)
}

// SampleLe mocks base method
func (m *MockLight) SampleLe(u1, u2 *Point2f, time float64) (Spectrum, *Ray, *Normal3f, float64, float64) {
	ret := m.ctrl.Call(m, "SampleLe", u1, u2, time)
	ret0, _ := ret[0].(Spectrum)
	ret1, _ := ret[1].(*Ray)
	ret2, _ := ret[2].(*Normal3f)
	ret3, _ := ret[3].(float64)
	ret4, _ := ret[4].(float64)
	return ret0, ret1, ret2, ret3, ret4
}

// SampleLe indicates an expected call of SampleLe
func (mr *MockLightMockRecorder) SampleLe(u1, u2, time interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SampleLe", reflect.TypeOf((*MockLight)(nil).SampleLe), u1, u2, time)
}

// PdfLe mocks base method
func (m *MockLight) PdfLe(r *Ray, nLight *Normal3f) (float64, float64) {
	ret := m.ctrl.Call(m, "PdfLe", r, nLight)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(float64)
	return ret0, ret1
}

// PdfLe indicates an expected call of PdfLe
func (mr *MockLightMockRecorder) PdfLe(r, nLight interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PdfLe", reflect.TypeOf((*MockLight)(nil).PdfLe), r, nLight)
}

// MockAreaLighter is a mock of AreaLighter interface
type MockAreaLighter struct {
	ctrl     *gomock.Controller
	recorder *MockAreaLighterMockRecorder
}

// MockAreaLighterMockRecorder is the mock recorder for MockAreaLighter
type MockAreaLighterMockRecorder struct {
	mock *MockAreaLighter
}

// NewMockAreaLighter creates a new mock instance
func NewMockAreaLighter(ctrl *gomock.Controller) *MockAreaLighter {
	mock := &MockAreaLighter{ctrl: ctrl}
	mock.recorder = &MockAreaLighterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAreaLighter) EXPECT() *MockAreaLighterMockRecorder {
	return m.recorder
}

// GetFlags mocks base method
func (m *MockAreaLighter) GetFlags() LightFlag {
	ret := m.ctrl.Call(m, "GetFlags")
	ret0, _ := ret[0].(LightFlag)
	return ret0
}

// GetFlags indicates an expected call of GetFlags
func (mr *MockAreaLighterMockRecorder) GetFlags() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFlags", reflect.TypeOf((*MockAreaLighter)(nil).GetFlags))
}

// GetSamples mocks base method
func (m *MockAreaLighter) GetSamples() int {
	ret := m.ctrl.Call(m, "GetSamples")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetSamples indicates an expected call of GetSamples
func (mr *MockAreaLighterMockRecorder) GetSamples() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSamples", reflect.TypeOf((*MockAreaLighter)(nil).GetSamples))
}

// Preprocess mocks base method
func (m *MockAreaLighter) Preprocess() {
	m.ctrl.Call(m, "Preprocess")
}

// Preprocess indicates an expected call of Preprocess
func (mr *MockAreaLighterMockRecorder) Preprocess() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Preprocess", reflect.TypeOf((*MockAreaLighter)(nil).Preprocess))
}

// SampleLi mocks base method
func (m *MockAreaLighter) SampleLi(ref Interaction, u *Point2f) (Spectrum, *Vector3f, float64, *VisibilityTester) {
	ret := m.ctrl.Call(m, "SampleLi", ref, u)
	ret0, _ := ret[0].(Spectrum)
	ret1, _ := ret[1].(*Vector3f)
	ret2, _ := ret[2].(float64)
	ret3, _ := ret[3].(*VisibilityTester)
	return ret0, ret1, ret2, ret3
}

// SampleLi indicates an expected call of SampleLi
func (mr *MockAreaLighterMockRecorder) SampleLi(ref, u interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SampleLi", reflect.TypeOf((*MockAreaLighter)(nil).SampleLi), ref, u)
}

// Power mocks base method
func (m *MockAreaLighter) Power() Spectrum {
	ret := m.ctrl.Call(m, "Power")
	ret0, _ := ret[0].(Spectrum)
	return ret0
}

// Power indicates an expected call of Power
func (mr *MockAreaLighterMockRecorder) Power() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Power", reflect.TypeOf((*MockAreaLighter)(nil).Power))
}

// Le mocks base method
func (m *MockAreaLighter) Le(r *RayDifferential) Spectrum {
	ret := m.ctrl.Call(m, "Le", r)
	ret0, _ := ret[0].(Spectrum)
	return ret0
}

// Le indicates an expected call of Le
func (mr *MockAreaLighterMockRecorder) Le(r interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Le", reflect.TypeOf((*MockAreaLighter)(nil).Le), r)
}

// PdfLi mocks base method
func (m *MockAreaLighter) PdfLi(ref Interaction) (float64, *Vector3f) {
	ret := m.ctrl.Call(m, "PdfLi", ref)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(*Vector3f)
	return ret0, ret1
}

// PdfLi indicates an expected call of PdfLi
func (mr *MockAreaLighterMockRecorder) PdfLi(ref interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PdfLi", reflect.TypeOf((*MockAreaLighter)(nil).PdfLi), ref)
}

// SampleLe mocks base method
func (m *MockAreaLighter) SampleLe(u1, u2 *Point2f, time float64) (Spectrum, *Ray, *Normal3f, float64, float64) {
	ret := m.ctrl.Call(m, "SampleLe", u1, u2, time)
	ret0, _ := ret[0].(Spectrum)
	ret1, _ := ret[1].(*Ray)
	ret2, _ := ret[2].(*Normal3f)
	ret3, _ := ret[3].(float64)
	ret4, _ := ret[4].(float64)
	return ret0, ret1, ret2, ret3, ret4
}

// SampleLe indicates an expected call of SampleLe
func (mr *MockAreaLighterMockRecorder) SampleLe(u1, u2, time interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SampleLe", reflect.TypeOf((*MockAreaLighter)(nil).SampleLe), u1, u2, time)
}

// PdfLe mocks base method
func (m *MockAreaLighter) PdfLe(r *Ray, nLight *Normal3f) (float64, float64) {
	ret := m.ctrl.Call(m, "PdfLe", r, nLight)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(float64)
	return ret0, ret1
}

// PdfLe indicates an expected call of PdfLe
func (mr *MockAreaLighterMockRecorder) PdfLe(r, nLight interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PdfLe", reflect.TypeOf((*MockAreaLighter)(nil).PdfLe), r, nLight)
}

// L mocks base method
func (m *MockAreaLighter) L(i Interaction, w *Vector3f) Spectrum {
	ret := m.ctrl.Call(m, "L", i, w)
	ret0, _ := ret[0].(Spectrum)
	return ret0
}

// L indicates an expected call of L
func (mr *MockAreaLighterMockRecorder) L(i, w interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "L", reflect.TypeOf((*MockAreaLighter)(nil).L), i, w)
}