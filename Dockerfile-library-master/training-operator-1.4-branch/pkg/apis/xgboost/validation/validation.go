// Copyright 2021 The Kubeflow Authors
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

package validation

import (
	"fmt"

	xgboostv1 "github.com/kubeflow/training-operator/pkg/apis/xgboost/v1"

	commonv1 "github.com/kubeflow/common/pkg/apis/common/v1"
)

func ValidateV1XGBoostJobSpec(c *xgboostv1.XGBoostJobSpec) error {
	if c.XGBReplicaSpecs == nil {
		return fmt.Errorf("XGBoostJobSpec is not valid")
	}
	masterExists := false
	for rType, value := range c.XGBReplicaSpecs {
		if value == nil || len(value.Template.Spec.Containers) == 0 {
			return fmt.Errorf("XGBoostJobSpec is not valid: containers definition expected in %v", rType)
		}
		// Make sure the replica type is valid.
		validReplicaTypes := []commonv1.ReplicaType{xgboostv1.XGBoostReplicaTypeMaster, xgboostv1.XGBoostReplicaTypeWorker}

		isValidReplicaType := false
		for _, t := range validReplicaTypes {
			if t == rType {
				isValidReplicaType = true
				break
			}
		}

		if !isValidReplicaType {
			return fmt.Errorf("XGBoostReplicaType is %v but must be one of %v", rType, validReplicaTypes)
		}

		//Make sure the image is defined in the container
		defaultContainerPresent := false
		for _, container := range value.Template.Spec.Containers {
			if container.Image == "" {
				msg := fmt.Sprintf("XGBoostReplicaType is not valid: Image is undefined in the container of %v", rType)
				return fmt.Errorf(msg)
			}
			if container.Name == xgboostv1.DefaultContainerName {
				defaultContainerPresent = true
			}
		}
		//Make sure there has at least one container named "xgboost"
		if !defaultContainerPresent {
			msg := fmt.Sprintf("XGBoostReplicaType is not valid: There is no container named %s in %v", xgboostv1.DefaultContainerName, rType)
			return fmt.Errorf(msg)
		}
		if rType == xgboostv1.XGBoostReplicaTypeMaster {
			masterExists = true
			if value.Replicas != nil && int(*value.Replicas) != 1 {
				return fmt.Errorf("XGBoostReplicaType is not valid: There must be only 1 master replica")
			}
		}

	}

	if !masterExists {
		return fmt.Errorf("XGBoostReplicaType is not valid: Master ReplicaSpec must be present")
	}
	return nil

}
