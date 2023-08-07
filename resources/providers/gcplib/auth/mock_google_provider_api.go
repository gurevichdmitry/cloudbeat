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

// Code generated by mockery v2.24.0. DO NOT EDIT.

package auth

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	google "golang.org/x/oauth2/google"
)

// MockGoogleProviderAPI is an autogenerated mock type for the GoogleAuthProviderAPI type
type MockGoogleProviderAPI struct {
	mock.Mock
}

type MockGoogleProviderAPI_Expecter struct {
	mock *mock.Mock
}

func (_m *MockGoogleProviderAPI) EXPECT() *MockGoogleProviderAPI_Expecter {
	return &MockGoogleProviderAPI_Expecter{mock: &_m.Mock}
}

// FindDefaultCredentials provides a mock function with given fields: ctx
func (_m *MockGoogleProviderAPI) FindDefaultCredentials(ctx context.Context) (*google.Credentials, error) {
	ret := _m.Called(ctx)

	var r0 *google.Credentials
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*google.Credentials, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *google.Credentials); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*google.Credentials)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGoogleProviderAPI_FindDefaultCredentials_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindDefaultCredentials'
type MockGoogleProviderAPI_FindDefaultCredentials_Call struct {
	*mock.Call
}

// FindDefaultCredentials is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockGoogleProviderAPI_Expecter) FindDefaultCredentials(ctx interface{}) *MockGoogleProviderAPI_FindDefaultCredentials_Call {
	return &MockGoogleProviderAPI_FindDefaultCredentials_Call{Call: _e.mock.On("FindDefaultCredentials", ctx)}
}

func (_c *MockGoogleProviderAPI_FindDefaultCredentials_Call) Run(run func(ctx context.Context)) *MockGoogleProviderAPI_FindDefaultCredentials_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockGoogleProviderAPI_FindDefaultCredentials_Call) Return(_a0 *google.Credentials, _a1 error) *MockGoogleProviderAPI_FindDefaultCredentials_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGoogleProviderAPI_FindDefaultCredentials_Call) RunAndReturn(run func(context.Context) (*google.Credentials, error)) *MockGoogleProviderAPI_FindDefaultCredentials_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockGoogleProviderAPI interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockGoogleProviderAPI creates a new instance of MockGoogleProviderAPI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockGoogleProviderAPI(t mockConstructorTestingTNewMockGoogleProviderAPI) *MockGoogleProviderAPI {
	mock := &MockGoogleProviderAPI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}