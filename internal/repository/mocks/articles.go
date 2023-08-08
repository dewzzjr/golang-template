// Code generated by MockGen. DO NOT EDIT.
// Source: articles.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	model "github.com/dewzzjr/ais/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockArticle is a mock of Article interface.
type MockArticle struct {
	ctrl     *gomock.Controller
	recorder *MockArticleMockRecorder
}

// MockArticleMockRecorder is the mock recorder for MockArticle.
type MockArticleMockRecorder struct {
	mock *MockArticle
}

// NewMockArticle creates a new mock instance.
func NewMockArticle(ctrl *gomock.Controller) *MockArticle {
	mock := &MockArticle{ctrl: ctrl}
	mock.recorder = &MockArticleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticle) EXPECT() *MockArticleMockRecorder {
	return m.recorder
}

// Fetch mocks base method.
func (m *MockArticle) Fetch(arg0 context.Context) ([]model.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", arg0)
	ret0, _ := ret[0].([]model.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch.
func (mr *MockArticleMockRecorder) Fetch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockArticle)(nil).Fetch), arg0)
}

// Insert mocks base method.
func (m *MockArticle) Insert(arg0 context.Context, arg1 model.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockArticleMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockArticle)(nil).Insert), arg0, arg1)
}
