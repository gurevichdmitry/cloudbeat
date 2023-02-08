// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by mockery v2.15.0. DO NOT EDIT.

package user

import mock "github.com/stretchr/testify/mock"

// MockOSUser is an autogenerated mock type for the OSUser type
type MockOSUser struct {
	mock.Mock
}

type MockOSUser_Expecter struct {
	mock *mock.Mock
}

func (_m *MockOSUser) EXPECT() *MockOSUser_Expecter {
	return &MockOSUser_Expecter{mock: &_m.Mock}
}

// GetGroupNameFromID provides a mock function with given fields: gid, groupFilePath
func (_m *MockOSUser) GetGroupNameFromID(gid string, groupFilePath string) (string, error) {
	ret := _m.Called(gid, groupFilePath)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(gid, groupFilePath)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(gid, groupFilePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockOSUser_GetGroupNameFromID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGroupNameFromID'
type MockOSUser_GetGroupNameFromID_Call struct {
	*mock.Call
}

// GetGroupNameFromID is a helper method to define mock.On call
//   - gid string
//   - groupFilePath string
func (_e *MockOSUser_Expecter) GetGroupNameFromID(gid interface{}, groupFilePath interface{}) *MockOSUser_GetGroupNameFromID_Call {
	return &MockOSUser_GetGroupNameFromID_Call{Call: _e.mock.On("GetGroupNameFromID", gid, groupFilePath)}
}

func (_c *MockOSUser_GetGroupNameFromID_Call) Run(run func(gid string, groupFilePath string)) *MockOSUser_GetGroupNameFromID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockOSUser_GetGroupNameFromID_Call) Return(_a0 string, _a1 error) *MockOSUser_GetGroupNameFromID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetUserNameFromID provides a mock function with given fields: uid, userFilePath
func (_m *MockOSUser) GetUserNameFromID(uid string, userFilePath string) (string, error) {
	ret := _m.Called(uid, userFilePath)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(uid, userFilePath)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(uid, userFilePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockOSUser_GetUserNameFromID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserNameFromID'
type MockOSUser_GetUserNameFromID_Call struct {
	*mock.Call
}

// GetUserNameFromID is a helper method to define mock.On call
//   - uid string
//   - userFilePath string
func (_e *MockOSUser_Expecter) GetUserNameFromID(uid interface{}, userFilePath interface{}) *MockOSUser_GetUserNameFromID_Call {
	return &MockOSUser_GetUserNameFromID_Call{Call: _e.mock.On("GetUserNameFromID", uid, userFilePath)}
}

func (_c *MockOSUser_GetUserNameFromID_Call) Run(run func(uid string, userFilePath string)) *MockOSUser_GetUserNameFromID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockOSUser_GetUserNameFromID_Call) Return(_a0 string, _a1 error) *MockOSUser_GetUserNameFromID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewMockOSUser interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockOSUser creates a new instance of MockOSUser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockOSUser(t mockConstructorTestingTNewMockOSUser) *MockOSUser {
	mock := &MockOSUser{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
