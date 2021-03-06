// Copyright 2019 The Secreter Authors
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

package version

import (
	"runtime"
)

// Variables reflect version information for the secreter CLI.
// All but GoVersion are supposed to be passed via go linker, .e.g. "-ldflags -X"
var (
	Version   = "v0.0.0"
	GitCommit = "unknown"
	BuildDate = "1970-01-01T00:00:00Z"
	GoVersion = runtime.Version()
)
