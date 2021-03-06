// Copyright 2014 Google Inc. All Rights Reserved.
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

package interference

import "github.com/google/cadvisor/info"

// InterferenceDectector detects if there's a container which
// interferences with a set of containers. The detector tracks
// a set of containers and find the victims and antagonist.
type InterferenceDetector interface {
	// Tracks the behavior of the container.
	AddContainer(ref info.ContainerReference)

	// Returns a list of possible interferences. The upper layer may take action
	// based on the interference.
	Detect() ([]*info.Interference, error)

	// The name of the detector.
	Name() string
}
