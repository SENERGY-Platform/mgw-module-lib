/*
 * Copyright 2023 InfAI (CC SES)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package validation

import (
	"errors"
	"fmt"
	"github.com/SENERGY-Platform/mgw-module-lib/module"
	"github.com/SENERGY-Platform/mgw-module-lib/validation/sem_ver"
	"regexp"
)

func Validate(m *module.Module) error {
	if !isValidModuleID(m.ID) {
		return fmt.Errorf("invalid module ID format '%s'", m.ID)
	}
	if !sem_ver.IsValidSemVer(m.Version) {
		return fmt.Errorf("invalid version format '%s'", m.Version)
	}
	if !isValidModuleType(m.Type) {
		return fmt.Errorf("invalid module type '%s'", m.Type)
	}
	if !isValidDeploymentType(m.DeploymentType) {
		return fmt.Errorf("invlaid deployment type '%s'", m.DeploymentType)
	}
	if !validateKeyNotEmptyString(m.Volumes) {
		return errors.New("empty volume name")
	}
	if err := validateModuleDependencies(m.Dependencies); err != nil {
		return fmt.Errorf("invalid dependency configuration: %s", err)
	}
	if !validateKeyNotEmptyString(m.Resources) {
		return errors.New("empty resource reference")
	}
	if !validateKeyNotEmptyString(m.Secrets) {
		return errors.New("empty secret reference")
	}
	if !validateKeyNotEmptyString(m.Configs) {
		return errors.New("empty config reference")
	}
	if err := validateInputGroups(m.Inputs.Groups); err != nil {
		return fmt.Errorf("invalid input group configuration: %s", err)
	}
	if err := validateInputsAndGroups(m.Inputs.Resources, m.Inputs.Groups); err != nil {
		return fmt.Errorf("invalid input group configuration: %s", err)
	}
	if err := validateInputsAndGroups(m.Inputs.Secrets, m.Inputs.Groups); err != nil {
		return fmt.Errorf("invalid input group configuration: %s", err)
	}
	if err := validateInputsAndGroups(m.Inputs.Configs, m.Inputs.Groups); err != nil {
		return fmt.Errorf("invalid input group configuration: %s", err)
	}
	if err := validateInputsResources(m.Inputs.Resources, m.Resources); err != nil {
		return fmt.Errorf("invalid input configuration: %s", err)
	}
	if err := validateInputsSecrets(m.Inputs.Secrets, m.Secrets); err != nil {
		return fmt.Errorf("invalid input configuration: %s", err)
	}
	if err := validateInputsConfigs(m.Inputs.Configs, m.Configs); err != nil {
		return fmt.Errorf("invalid input configuration: %s", err)
	}
	if err := validateServices(m.Services, m.Volumes, m.Resources, m.Secrets, m.Configs, m.Dependencies); err != nil {
		return err
	}
	return nil
}

func validateModuleDependencies(dependencies map[string]string) error {
	for mid, ver := range dependencies {
		if !isValidModuleID(mid) {
			return fmt.Errorf("invalid module ID format '%s'", mid)
		}
		if err := sem_ver.ValidateSemVerRange(ver); err != nil {
			return fmt.Errorf("version %s", err)
		}
	}
	return nil
}

func isValidModuleType(s string) bool {
	_, ok := module.ModuleTypeMap[s]
	return ok
}

func isValidDeploymentType(s string) bool {
	_, ok := module.DeploymentTypeMap[s]
	return ok
}

func isValidModuleID(s string) bool {
	re := regexp.MustCompile(`^(?:[a-zA-Z0-9-_]+)\.(?:[a-zA-Z]+)(?:\/[a-zA-Z0-9-_]+)+$`)
	return re.MatchString(s)
}

func isValidCPUArch(s string) bool {
	_, ok := module.CPUArchMap[s]
	return ok
}
