//  Copyright (c) 2020 Cisco and/or its affiliates.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package version

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

// Following variables should normally be updated via `-ldflags "-X ..."`.
// However, the version string is hard-coded to ensure it is always included
// even with bare go build/install.
var (
	name       = "vpp-probe"
	version    = "v0.1.0-dev"
	commit     = "unknown"
	branch     = "HEAD"
	buildStamp = ""
	buildUser  = ""
	buildHost  = ""

	buildDate time.Time
)

func init() {
	buildstampInt64, _ := strconv.ParseInt(buildStamp, 10, 64)
	if buildstampInt64 == 0 {
		buildstampInt64 = time.Now().Unix()
	}
	buildDate = time.Unix(buildstampInt64, 0)
}

func String() string {
	return version
}

func Verbose() string {
	return fmt.Sprintf(`%s
  Version:      %s
  Branch:   	%s
  Revision: 	%s
  Built by:  	%s@%s 
  Build date:	%s
  Go runtime:	%s (%s/%s)`,
		name,
		version, branch, commit,
		buildUser, buildHost, buildDate.Format(time.UnixDate),
		runtime.Version(), runtime.GOOS, runtime.GOARCH,
	)
}
