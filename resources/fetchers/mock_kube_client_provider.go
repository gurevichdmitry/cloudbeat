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

package fetchers

import (
	kubernetes "github.com/elastic/elastic-agent-autodiscover/kubernetes"
	mock "github.com/stretchr/testify/mock"
	client_gokubernetes "k8s.io/client-go/kubernetes"
)

// MockKubeClientProvider is an autogenerated mock type for the KubeClientProvider type
type MockKubeClientProvider struct {
	mock.Mock
}

type MockKubeClientProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *MockKubeClientProvider) EXPECT() *MockKubeClientProvider_Expecter {
	return &MockKubeClientProvider_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: kubeconfig, opt
func (_m *MockKubeClientProvider) Execute(kubeconfig string, opt kubernetes.KubeClientOptions) (client_gokubernetes.Interface, error) {
	ret := _m.Called(kubeconfig, opt)

	var r0 client_gokubernetes.Interface
	if rf, ok := ret.Get(0).(func(string, kubernetes.KubeClientOptions) client_gokubernetes.Interface); ok {
		r0 = rf(kubeconfig, opt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client_gokubernetes.Interface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, kubernetes.KubeClientOptions) error); ok {
		r1 = rf(kubeconfig, opt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockKubeClientProvider_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockKubeClientProvider_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - kubeconfig string
//   - opt kubernetes.KubeClientOptions
func (_e *MockKubeClientProvider_Expecter) Execute(kubeconfig interface{}, opt interface{}) *MockKubeClientProvider_Execute_Call {
	return &MockKubeClientProvider_Execute_Call{Call: _e.mock.On("Execute", kubeconfig, opt)}
}

func (_c *MockKubeClientProvider_Execute_Call) Run(run func(kubeconfig string, opt kubernetes.KubeClientOptions)) *MockKubeClientProvider_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(kubernetes.KubeClientOptions))
	})
	return _c
}

func (_c *MockKubeClientProvider_Execute_Call) Return(_a0 client_gokubernetes.Interface, _a1 error) *MockKubeClientProvider_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewMockKubeClientProvider interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockKubeClientProvider creates a new instance of MockKubeClientProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockKubeClientProvider(t mockConstructorTestingTNewMockKubeClientProvider) *MockKubeClientProvider {
	mock := &MockKubeClientProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
