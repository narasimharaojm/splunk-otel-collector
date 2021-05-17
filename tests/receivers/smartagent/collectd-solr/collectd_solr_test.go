// Copyright Splunk, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tests

import (
	"path"
	"testing"

	"github.com/signalfx/splunk-otel-collector/tests/testutils"
)

func TestCollectdSolrReceiverProvidesAllMetrics(t *testing.T) {
	containers := []testutils.Container{
		testutils.NewContainer().WithContext(
			path.Join(".", "testdata", "server"),
		).WithExposedPorts("8983:8983").WithName(
			"solr",
		).WillWaitForPorts("8983").WillWaitForLogs("example launched successfully"),
	}

	testutils.AssertAllMetricsReceived(
		t, "all.yaml", "all_metrics_config.yaml", containers,
	)
}