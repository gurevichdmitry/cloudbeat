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

// Code generated by mockery v2.13.1. DO NOT EDIT.

package dataprovider

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// mockK8sDataProvider is an autogenerated mock type for the k8sDataProvider type
type mockK8sDataProvider struct {
	mock.Mock
}

type mockK8sDataProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *mockK8sDataProvider) EXPECT() *mockK8sDataProvider_Expecter {
	return &mockK8sDataProvider_Expecter{mock: &_m.Mock}
}

// CollectK8sData provides a mock function with given fields: ctx
func (_m *mockK8sDataProvider) CollectK8sData(ctx context.Context) *CommonK8sData {
	ret := _m.Called(ctx)

	var r0 *CommonK8sData
	if rf, ok := ret.Get(0).(func(context.Context) *CommonK8sData); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*CommonK8sData)
		}
	}

	return r0
}

// mockK8sDataProvider_CollectK8sData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CollectK8sData'
type mockK8sDataProvider_CollectK8sData_Call struct {
	*mock.Call
}

// CollectK8sData is a helper method to define mock.On call
//  - ctx context.Context
func (_e *mockK8sDataProvider_Expecter) CollectK8sData(ctx interface{}) *mockK8sDataProvider_CollectK8sData_Call {
	return &mockK8sDataProvider_CollectK8sData_Call{Call: _e.mock.On("CollectK8sData", ctx)}
}

func (_c *mockK8sDataProvider_CollectK8sData_Call) Run(run func(ctx context.Context)) *mockK8sDataProvider_CollectK8sData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *mockK8sDataProvider_CollectK8sData_Call) Return(_a0 *CommonK8sData) *mockK8sDataProvider_CollectK8sData_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTnewMockK8sDataProvider interface {
	mock.TestingT
	Cleanup(func())
}

// newMockK8sDataProvider creates a new instance of mockK8sDataProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockK8sDataProvider(t mockConstructorTestingTnewMockK8sDataProvider) *mockK8sDataProvider {
	mock := &mockK8sDataProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}