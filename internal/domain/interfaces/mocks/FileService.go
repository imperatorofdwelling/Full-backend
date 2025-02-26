// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	file "github.com/imperatorofdwelling/Full-backend/internal/service/file"

	mock "github.com/stretchr/testify/mock"
)

// FileService is an autogenerated mock type for the FileService type
type FileService struct {
	mock.Mock
}

// GenRandomFileName provides a mock function with no fields
func (_m *FileService) GenRandomFileName() (string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GenRandomFileName")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveFile provides a mock function with given fields: fileName
func (_m *FileService) RemoveFile(fileName string) error {
	ret := _m.Called(fileName)

	if len(ret) == 0 {
		panic("no return value specified for RemoveFile")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(fileName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UploadImage provides a mock function with given fields: img, imgType, filePath
func (_m *FileService) UploadImage(img []byte, imgType file.ImageType, filePath file.PathType) (string, error) {
	ret := _m.Called(img, imgType, filePath)

	if len(ret) == 0 {
		panic("no return value specified for UploadImage")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte, file.ImageType, file.PathType) (string, error)); ok {
		return rf(img, imgType, filePath)
	}
	if rf, ok := ret.Get(0).(func([]byte, file.ImageType, file.PathType) string); ok {
		r0 = rf(img, imgType, filePath)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func([]byte, file.ImageType, file.PathType) error); ok {
		r1 = rf(img, imgType, filePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewFileService creates a new instance of FileService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFileService(t interface {
	mock.TestingT
	Cleanup(func())
}) *FileService {
	mock := &FileService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
