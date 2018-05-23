// Copyright 2015-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package adapter

import (
	"fmt"
)

const (
	// prefix for the autogenerated volume names in the task definition
	ecsVolumeNamePrefix = "volume"
)

type Volumes struct {
	VolumeWithHost  map[string]string
	VolumeEmptyHost []string
}

// getVolumeName returns an autogenerated name for the ecs volume
func getVolumeName(suffixNum int) string {
	return fmt.Sprintf("%s-%d", ecsVolumeNamePrefix, suffixNum)
}