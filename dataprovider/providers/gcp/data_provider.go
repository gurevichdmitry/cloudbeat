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

package gcp

import (
	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/elastic-agent-libs/logp"

	"github.com/elastic/cloudbeat/dataprovider/types"
	"github.com/elastic/cloudbeat/resources/fetching"
	"github.com/elastic/cloudbeat/version"
)

type DataProvider struct {
	log *logp.Logger
}

func New(options ...Option) DataProvider {
	adp := DataProvider{}
	for _, opt := range options {
		opt(&adp)
	}
	return adp
}

func (a DataProvider) FetchData(_ string, id string) (types.Data, error) {
	return types.Data{
		ResourceID: id,
		VersionInfo: version.CloudbeatVersionInfo{
			Version: version.CloudbeatVersion(),
			Policy:  version.PolicyVersion(),
		},
	}, nil
}

func (a DataProvider) EnrichEvent(event *beat.Event, resMetadata fetching.ResourceMetadata) error {
	return nil
}