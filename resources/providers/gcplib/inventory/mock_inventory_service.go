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

package inventory

import (
	assetpb "cloud.google.com/go/asset/apiv1/assetpb"
	mock "github.com/stretchr/testify/mock"
)

// MockInventoryService is an autogenerated mock type for the InventoryService type
type MockInventoryService struct {
	mock.Mock
}

type MockInventoryService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockInventoryService) EXPECT() *MockInventoryService_Expecter {
	return &MockInventoryService_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *MockInventoryService) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockInventoryService_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockInventoryService_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockInventoryService_Expecter) Close() *MockInventoryService_Close_Call {
	return &MockInventoryService_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockInventoryService_Close_Call) Run(run func()) *MockInventoryService_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockInventoryService_Close_Call) Return(_a0 error) *MockInventoryService_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInventoryService_Close_Call) RunAndReturn(run func() error) *MockInventoryService_Close_Call {
	_c.Call.Return(run)
	return _c
}

// ListAllAssetTypesByName provides a mock function with given fields: assets
func (_m *MockInventoryService) ListAllAssetTypesByName(assets []string) ([]*assetpb.Asset, error) {
	ret := _m.Called(assets)

	var r0 []*assetpb.Asset
	var r1 error
	if rf, ok := ret.Get(0).(func([]string) ([]*assetpb.Asset, error)); ok {
		return rf(assets)
	}
	if rf, ok := ret.Get(0).(func([]string) []*assetpb.Asset); ok {
		r0 = rf(assets)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*assetpb.Asset)
		}
	}

	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(assets)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInventoryService_ListAllAssetTypesByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAllAssetTypesByName'
type MockInventoryService_ListAllAssetTypesByName_Call struct {
	*mock.Call
}

// ListAllAssetTypesByName is a helper method to define mock.On call
//   - assets []string
func (_e *MockInventoryService_Expecter) ListAllAssetTypesByName(assets interface{}) *MockInventoryService_ListAllAssetTypesByName_Call {
	return &MockInventoryService_ListAllAssetTypesByName_Call{Call: _e.mock.On("ListAllAssetTypesByName", assets)}
}

func (_c *MockInventoryService_ListAllAssetTypesByName_Call) Run(run func(assets []string)) *MockInventoryService_ListAllAssetTypesByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *MockInventoryService_ListAllAssetTypesByName_Call) Return(_a0 []*assetpb.Asset, _a1 error) *MockInventoryService_ListAllAssetTypesByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInventoryService_ListAllAssetTypesByName_Call) RunAndReturn(run func([]string) ([]*assetpb.Asset, error)) *MockInventoryService_ListAllAssetTypesByName_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockInventoryService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockInventoryService creates a new instance of MockInventoryService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockInventoryService(t mockConstructorTestingTNewMockInventoryService) *MockInventoryService {
	mock := &MockInventoryService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
