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

package awsfetcher

import (
	context "context"

	awslib "github.com/elastic/cloudbeat/internal/resources/providers/awslib"

	mock "github.com/stretchr/testify/mock"
)

// mockIamPolicyProvider is an autogenerated mock type for the iamPolicyProvider type
type mockIamPolicyProvider struct {
	mock.Mock
}

type mockIamPolicyProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *mockIamPolicyProvider) EXPECT() *mockIamPolicyProvider_Expecter {
	return &mockIamPolicyProvider_Expecter{mock: &_m.Mock}
}

// GetPolicies provides a mock function with given fields: ctx
func (_m *mockIamPolicyProvider) GetPolicies(ctx context.Context) ([]awslib.AwsResource, error) {
	ret := _m.Called(ctx)

	var r0 []awslib.AwsResource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]awslib.AwsResource, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []awslib.AwsResource); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]awslib.AwsResource)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockIamPolicyProvider_GetPolicies_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPolicies'
type mockIamPolicyProvider_GetPolicies_Call struct {
	*mock.Call
}

// GetPolicies is a helper method to define mock.On call
//   - ctx context.Context
func (_e *mockIamPolicyProvider_Expecter) GetPolicies(ctx interface{}) *mockIamPolicyProvider_GetPolicies_Call {
	return &mockIamPolicyProvider_GetPolicies_Call{Call: _e.mock.On("GetPolicies", ctx)}
}

func (_c *mockIamPolicyProvider_GetPolicies_Call) Run(run func(ctx context.Context)) *mockIamPolicyProvider_GetPolicies_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *mockIamPolicyProvider_GetPolicies_Call) Return(_a0 []awslib.AwsResource, _a1 error) *mockIamPolicyProvider_GetPolicies_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockIamPolicyProvider_GetPolicies_Call) RunAndReturn(run func(context.Context) ([]awslib.AwsResource, error)) *mockIamPolicyProvider_GetPolicies_Call {
	_c.Call.Return(run)
	return _c
}

// newMockIamPolicyProvider creates a new instance of mockIamPolicyProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockIamPolicyProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockIamPolicyProvider {
	mock := &mockIamPolicyProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
