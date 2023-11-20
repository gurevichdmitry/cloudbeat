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

// Code generated by mockery v2.37.1. DO NOT EDIT.

package inventory

import mock "github.com/stretchr/testify/mock"

// MockTypeGenerator is an autogenerated mock type for the TypeGenerator type
type MockTypeGenerator[T interface{}] struct {
	mock.Mock
}

type MockTypeGenerator_Expecter[T interface{}] struct {
	mock *mock.Mock
}

func (_m *MockTypeGenerator[T]) EXPECT() *MockTypeGenerator_Expecter[T] {
	return &MockTypeGenerator_Expecter[T]{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: assets, projectId, projectName, orgId, orgName
func (_m *MockTypeGenerator[T]) Execute(assets []*ExtendedGcpAsset, projectId string, projectName string, orgId string, orgName string) *T {
	ret := _m.Called(assets, projectId, projectName, orgId, orgName)

	var r0 *T
	if rf, ok := ret.Get(0).(func([]*ExtendedGcpAsset, string, string, string, string) *T); ok {
		r0 = rf(assets, projectId, projectName, orgId, orgName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*T)
		}
	}

	return r0
}

// MockTypeGenerator_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockTypeGenerator_Execute_Call[T interface{}] struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - assets []*ExtendedGcpAsset
//   - projectId string
//   - projectName string
//   - orgId string
//   - orgName string
func (_e *MockTypeGenerator_Expecter[T]) Execute(assets interface{}, projectId interface{}, projectName interface{}, orgId interface{}, orgName interface{}) *MockTypeGenerator_Execute_Call[T] {
	return &MockTypeGenerator_Execute_Call[T]{Call: _e.mock.On("Execute", assets, projectId, projectName, orgId, orgName)}
}

func (_c *MockTypeGenerator_Execute_Call[T]) Run(run func(assets []*ExtendedGcpAsset, projectId string, projectName string, orgId string, orgName string)) *MockTypeGenerator_Execute_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]*ExtendedGcpAsset), args[1].(string), args[2].(string), args[3].(string), args[4].(string))
	})
	return _c
}

func (_c *MockTypeGenerator_Execute_Call[T]) Return(_a0 *T) *MockTypeGenerator_Execute_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTypeGenerator_Execute_Call[T]) RunAndReturn(run func([]*ExtendedGcpAsset, string, string, string, string) *T) *MockTypeGenerator_Execute_Call[T] {
	_c.Call.Return(run)
	return _c
}

// NewMockTypeGenerator creates a new instance of MockTypeGenerator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTypeGenerator[T interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTypeGenerator[T] {
	mock := &MockTypeGenerator[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
