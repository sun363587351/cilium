// Copyright 2018 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package workloads

import (
	"time"
)

const (
	// EndpointCorrelationMaxRetries is the number of retries to correlate
	// a workload start event with an existing endpoint before giving up.
	EndpointCorrelationMaxRetries = 20
)

// EndpointCorrelationSleepTime returns the sleep time between correlation
// attempts
func EndpointCorrelationSleepTime(try int) time.Duration {
	return time.Duration(try) * time.Second
}
