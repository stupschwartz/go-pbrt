// Code generated by MockGen. DO NOT EDIT.
// Source: reflection.go

// Package pbrt is a generated GoMock package.
package pbrt

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBxDFer is a mock of BxDF interface
type MockBxDFer struct {
	ctrl     *gomock.Controller
	recorder *MockBxDFerMockRecorder
}

// MockBxDFerMockRecorder is the mock recorder for MockBxDFer
type MockBxDFerMockRecorder struct {
	mock *MockBxDFer
}

// NewMockBxDFer creates a new mock instance
func NewMockBxDFer(ctrl *gomock.Controller) *MockBxDFer {
	mock := &MockBxDFer{ctrl: ctrl}
	mock.recorder = &MockBxDFerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBxDFer) EXPECT() *MockBxDFerMockRecorder {
	return m.recorder
}

// GetType mocks base method
func (m *MockBxDFer) GetType() BxDFType {
	ret := m.ctrl.Call(m, "GetType")
	ret0, _ := ret[0].(BxDFType)
	return ret0
}

// GetType indicates an expected call of GetType
func (mr *MockBxDFerMockRecorder) GetType() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetType", reflect.TypeOf((*MockBxDFer)(nil).GetType))
}

// MatchesFlags mocks base method
func (m *MockBxDFer) MatchesFlags(t BxDFType) bool {
	ret := m.ctrl.Call(m, "MatchesFlags", t)
	ret0, _ := ret[0].(bool)
	return ret0
}

// MatchesFlags indicates an expected call of MatchesFlags
func (mr *MockBxDFerMockRecorder) MatchesFlags(t interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MatchesFlags", reflect.TypeOf((*MockBxDFer)(nil).MatchesFlags), t)
}

// F mocks base method
func (m *MockBxDFer) F(woW, wiW *Vector3f) Spectrum {
	ret := m.ctrl.Call(m, "F", woW, wiW)
	ret0, _ := ret[0].(Spectrum)
	return ret0
}

// F indicates an expected call of F
func (mr *MockBxDFerMockRecorder) F(woW, wiW interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "F", reflect.TypeOf((*MockBxDFer)(nil).F), woW, wiW)
}

// SampleF mocks base method
func (m *MockBxDFer) SampleF(wo *Vector3f, sample *Point2f) (Spectrum, *Vector3f, float64, BxDFType) {
	ret := m.ctrl.Call(m, "SampleF", wo, sample)
	ret0, _ := ret[0].(Spectrum)
	ret1, _ := ret[1].(*Vector3f)
	ret2, _ := ret[2].(float64)
	ret3, _ := ret[3].(BxDFType)
	return ret0, ret1, ret2, ret3
}

// SampleF indicates an expected call of SampleF
func (mr *MockBxDFerMockRecorder) SampleF(wo, sample interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SampleF", reflect.TypeOf((*MockBxDFer)(nil).SampleF), wo, sample)
}

// Rho mocks base method
func (m *MockBxDFer) Rho(wo *Vector3f, samples []*Point2f) Spectrum {
	ret := m.ctrl.Call(m, "Rho", wo, samples)
	ret0, _ := ret[0].(Spectrum)
	return ret0
}

// Rho indicates an expected call of Rho
func (mr *MockBxDFerMockRecorder) Rho(wo, samples interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rho", reflect.TypeOf((*MockBxDFer)(nil).Rho), wo, samples)
}

// RhoSamples mocks base method
func (m *MockBxDFer) RhoSamples(samples1, samples2 []*Point2f) Spectrum {
	ret := m.ctrl.Call(m, "RhoSamples", samples1, samples2)
	ret0, _ := ret[0].(Spectrum)
	return ret0
}

// RhoSamples indicates an expected call of RhoSamples
func (mr *MockBxDFerMockRecorder) RhoSamples(samples1, samples2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RhoSamples", reflect.TypeOf((*MockBxDFer)(nil).RhoSamples), samples1, samples2)
}

// Pdf mocks base method
func (m *MockBxDFer) Pdf(wo, wi *Vector3f) float64 {
	ret := m.ctrl.Call(m, "Pdf", wo, wi)
	ret0, _ := ret[0].(float64)
	return ret0
}

// Pdf indicates an expected call of Pdf
func (mr *MockBxDFerMockRecorder) Pdf(wo, wi interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pdf", reflect.TypeOf((*MockBxDFer)(nil).Pdf), wo, wi)
}