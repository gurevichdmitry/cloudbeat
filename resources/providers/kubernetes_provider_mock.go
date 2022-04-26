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
// Code generated by mockery v2.10.4. DO NOT EDIT.

package providers

import (
	kubernetes "github.com/elastic/beats/v7/libbeat/common/kubernetes"
	client_gokubernetes "k8s.io/client-go/kubernetes"

	mock "github.com/stretchr/testify/mock"
)

// KubernetesClientGetter is an autogenerated mock type for the KubernetesClientGetter type
type MockedKubernetesClientGetter struct {
	mock.Mock
}

type KubernetesClientGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockedKubernetesClientGetter) EXPECT() *KubernetesClientGetter_Expecter {
	return &KubernetesClientGetter_Expecter{mock: &_m.Mock}
}

// GetClient provides a mock function with given fields: kubeConfig, options
func (_m *MockedKubernetesClientGetter) GetClient(kubeConfig string, options kubernetes.KubeClientOptions) (client_gokubernetes.Interface, error) {
	ret := _m.Called(kubeConfig, options)

	var r0 client_gokubernetes.Interface
	if rf, ok := ret.Get(0).(func(string, kubernetes.KubeClientOptions) client_gokubernetes.Interface); ok {
		r0 = rf(kubeConfig, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client_gokubernetes.Interface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, kubernetes.KubeClientOptions) error); ok {
		r1 = rf(kubeConfig, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KubernetesClientGetter_GetClient_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetClient'
type KubernetesClientGetter_GetClient_Call struct {
	*mock.Call
}

// GetClient is a helper method to define mock.On call
//  - kubeConfig string
//  - options kubernetes.KubeClientOptions
func (_e *KubernetesClientGetter_Expecter) GetClient(kubeConfig interface{}, options interface{}) *KubernetesClientGetter_GetClient_Call {
	return &KubernetesClientGetter_GetClient_Call{Call: _e.mock.On("GetClient", kubeConfig, options)}
}

func (_c *KubernetesClientGetter_GetClient_Call) Run(run func(kubeConfig string, options kubernetes.KubeClientOptions)) *KubernetesClientGetter_GetClient_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(kubernetes.KubeClientOptions))
	})
	return _c
}

func (_c *KubernetesClientGetter_GetClient_Call) Return(_a0 client_gokubernetes.Interface, _a1 error) *KubernetesClientGetter_GetClient_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}
