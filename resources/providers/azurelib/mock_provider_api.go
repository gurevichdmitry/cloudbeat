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

package azurelib

import (
	context "context"

	cycle "github.com/elastic/cloudbeat/resources/fetching/cycle"
	governance "github.com/elastic/cloudbeat/resources/providers/azurelib/governance"

	inventory "github.com/elastic/cloudbeat/resources/providers/azurelib/inventory"

	mock "github.com/stretchr/testify/mock"
)

// MockProviderAPI is an autogenerated mock type for the ProviderAPI type
type MockProviderAPI struct {
	mock.Mock
}

type MockProviderAPI_Expecter struct {
	mock *mock.Mock
}

func (_m *MockProviderAPI) EXPECT() *MockProviderAPI_Expecter {
	return &MockProviderAPI_Expecter{mock: &_m.Mock}
}

// GetSubscriptions provides a mock function with given fields: ctx, cycleMetadata
func (_m *MockProviderAPI) GetSubscriptions(ctx context.Context, cycleMetadata cycle.Metadata) (map[string]governance.Subscription, error) {
	ret := _m.Called(ctx, cycleMetadata)

	var r0 map[string]governance.Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, cycle.Metadata) (map[string]governance.Subscription, error)); ok {
		return rf(ctx, cycleMetadata)
	}
	if rf, ok := ret.Get(0).(func(context.Context, cycle.Metadata) map[string]governance.Subscription); ok {
		r0 = rf(ctx, cycleMetadata)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]governance.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, cycle.Metadata) error); ok {
		r1 = rf(ctx, cycleMetadata)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProviderAPI_GetSubscriptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSubscriptions'
type MockProviderAPI_GetSubscriptions_Call struct {
	*mock.Call
}

// GetSubscriptions is a helper method to define mock.On call
//   - ctx context.Context
//   - cycleMetadata cycle.Metadata
func (_e *MockProviderAPI_Expecter) GetSubscriptions(ctx interface{}, cycleMetadata interface{}) *MockProviderAPI_GetSubscriptions_Call {
	return &MockProviderAPI_GetSubscriptions_Call{Call: _e.mock.On("GetSubscriptions", ctx, cycleMetadata)}
}

func (_c *MockProviderAPI_GetSubscriptions_Call) Run(run func(ctx context.Context, cycleMetadata cycle.Metadata)) *MockProviderAPI_GetSubscriptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(cycle.Metadata))
	})
	return _c
}

func (_c *MockProviderAPI_GetSubscriptions_Call) Return(_a0 map[string]governance.Subscription, _a1 error) *MockProviderAPI_GetSubscriptions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProviderAPI_GetSubscriptions_Call) RunAndReturn(run func(context.Context, cycle.Metadata) (map[string]governance.Subscription, error)) *MockProviderAPI_GetSubscriptions_Call {
	_c.Call.Return(run)
	return _c
}

// ListAllAssetTypesByName provides a mock function with given fields: ctx, assetsGroup, assets
func (_m *MockProviderAPI) ListAllAssetTypesByName(ctx context.Context, assetsGroup string, assets []string) ([]inventory.AzureAsset, error) {
	ret := _m.Called(ctx, assetsGroup, assets)

	var r0 []inventory.AzureAsset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) ([]inventory.AzureAsset, error)); ok {
		return rf(ctx, assetsGroup, assets)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) []inventory.AzureAsset); ok {
		r0 = rf(ctx, assetsGroup, assets)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]inventory.AzureAsset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, []string) error); ok {
		r1 = rf(ctx, assetsGroup, assets)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProviderAPI_ListAllAssetTypesByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAllAssetTypesByName'
type MockProviderAPI_ListAllAssetTypesByName_Call struct {
	*mock.Call
}

// ListAllAssetTypesByName is a helper method to define mock.On call
//   - ctx context.Context
//   - assetsGroup string
//   - assets []string
func (_e *MockProviderAPI_Expecter) ListAllAssetTypesByName(ctx interface{}, assetsGroup interface{}, assets interface{}) *MockProviderAPI_ListAllAssetTypesByName_Call {
	return &MockProviderAPI_ListAllAssetTypesByName_Call{Call: _e.mock.On("ListAllAssetTypesByName", ctx, assetsGroup, assets)}
}

func (_c *MockProviderAPI_ListAllAssetTypesByName_Call) Run(run func(ctx context.Context, assetsGroup string, assets []string)) *MockProviderAPI_ListAllAssetTypesByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].([]string))
	})
	return _c
}

func (_c *MockProviderAPI_ListAllAssetTypesByName_Call) Return(_a0 []inventory.AzureAsset, _a1 error) *MockProviderAPI_ListAllAssetTypesByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProviderAPI_ListAllAssetTypesByName_Call) RunAndReturn(run func(context.Context, string, []string) ([]inventory.AzureAsset, error)) *MockProviderAPI_ListAllAssetTypesByName_Call {
	_c.Call.Return(run)
	return _c
}

// ListDiagnosticSettingsAssetTypes provides a mock function with given fields: ctx, cycleMetadata, subscriptionIDs
func (_m *MockProviderAPI) ListDiagnosticSettingsAssetTypes(ctx context.Context, cycleMetadata cycle.Metadata, subscriptionIDs []string) ([]inventory.AzureAsset, error) {
	ret := _m.Called(ctx, cycleMetadata, subscriptionIDs)

	var r0 []inventory.AzureAsset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, cycle.Metadata, []string) ([]inventory.AzureAsset, error)); ok {
		return rf(ctx, cycleMetadata, subscriptionIDs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, cycle.Metadata, []string) []inventory.AzureAsset); ok {
		r0 = rf(ctx, cycleMetadata, subscriptionIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]inventory.AzureAsset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, cycle.Metadata, []string) error); ok {
		r1 = rf(ctx, cycleMetadata, subscriptionIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProviderAPI_ListDiagnosticSettingsAssetTypes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListDiagnosticSettingsAssetTypes'
type MockProviderAPI_ListDiagnosticSettingsAssetTypes_Call struct {
	*mock.Call
}

// ListDiagnosticSettingsAssetTypes is a helper method to define mock.On call
//   - ctx context.Context
//   - cycleMetadata cycle.Metadata
//   - subscriptionIDs []string
func (_e *MockProviderAPI_Expecter) ListDiagnosticSettingsAssetTypes(ctx interface{}, cycleMetadata interface{}, subscriptionIDs interface{}) *MockProviderAPI_ListDiagnosticSettingsAssetTypes_Call {
	return &MockProviderAPI_ListDiagnosticSettingsAssetTypes_Call{Call: _e.mock.On("ListDiagnosticSettingsAssetTypes", ctx, cycleMetadata, subscriptionIDs)}
}

func (_c *MockProviderAPI_ListDiagnosticSettingsAssetTypes_Call) Run(run func(ctx context.Context, cycleMetadata cycle.Metadata, subscriptionIDs []string)) *MockProviderAPI_ListDiagnosticSettingsAssetTypes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(cycle.Metadata), args[2].([]string))
	})
	return _c
}

func (_c *MockProviderAPI_ListDiagnosticSettingsAssetTypes_Call) Return(_a0 []inventory.AzureAsset, _a1 error) *MockProviderAPI_ListDiagnosticSettingsAssetTypes_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProviderAPI_ListDiagnosticSettingsAssetTypes_Call) RunAndReturn(run func(context.Context, cycle.Metadata, []string) ([]inventory.AzureAsset, error)) *MockProviderAPI_ListDiagnosticSettingsAssetTypes_Call {
	_c.Call.Return(run)
	return _c
}

// ListStorageAccountBlobServices provides a mock function with given fields: ctx, storageAccounts
func (_m *MockProviderAPI) ListStorageAccountBlobServices(ctx context.Context, storageAccounts []inventory.AzureAsset) ([]inventory.AzureAsset, error) {
	ret := _m.Called(ctx, storageAccounts)

	var r0 []inventory.AzureAsset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []inventory.AzureAsset) ([]inventory.AzureAsset, error)); ok {
		return rf(ctx, storageAccounts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []inventory.AzureAsset) []inventory.AzureAsset); ok {
		r0 = rf(ctx, storageAccounts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]inventory.AzureAsset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []inventory.AzureAsset) error); ok {
		r1 = rf(ctx, storageAccounts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProviderAPI_ListStorageAccountBlobServices_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListStorageAccountBlobServices'
type MockProviderAPI_ListStorageAccountBlobServices_Call struct {
	*mock.Call
}

// ListStorageAccountBlobServices is a helper method to define mock.On call
//   - ctx context.Context
//   - storageAccounts []inventory.AzureAsset
func (_e *MockProviderAPI_Expecter) ListStorageAccountBlobServices(ctx interface{}, storageAccounts interface{}) *MockProviderAPI_ListStorageAccountBlobServices_Call {
	return &MockProviderAPI_ListStorageAccountBlobServices_Call{Call: _e.mock.On("ListStorageAccountBlobServices", ctx, storageAccounts)}
}

func (_c *MockProviderAPI_ListStorageAccountBlobServices_Call) Run(run func(ctx context.Context, storageAccounts []inventory.AzureAsset)) *MockProviderAPI_ListStorageAccountBlobServices_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]inventory.AzureAsset))
	})
	return _c
}

func (_c *MockProviderAPI_ListStorageAccountBlobServices_Call) Return(_a0 []inventory.AzureAsset, _a1 error) *MockProviderAPI_ListStorageAccountBlobServices_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProviderAPI_ListStorageAccountBlobServices_Call) RunAndReturn(run func(context.Context, []inventory.AzureAsset) ([]inventory.AzureAsset, error)) *MockProviderAPI_ListStorageAccountBlobServices_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockProviderAPI creates a new instance of MockProviderAPI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockProviderAPI(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockProviderAPI {
	mock := &MockProviderAPI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
