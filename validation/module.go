/*
 * Copyright 2025 InfAI (CC SES)
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
	"regexp"

	"github.com/SENERGY-Platform/mgw-module-lib/model"
	"github.com/SENERGY-Platform/mgw-module-lib/util/sem_ver"
	"github.com/SENERGY-Platform/mgw-module-lib/validation/configs"
)

func Validate(m *model.Module) error {
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
	if err := validateResources(m.HostResources, m.Inputs.Resources); err != nil {
		return fmt.Errorf("invalid resource configuration: %s", err)
	}
	if err := validateSecrets(m.Secrets, m.Inputs.Secrets); err != nil {
		return fmt.Errorf("invalid secret configuration: %s", err)
	}
	if err := validateConfigs(m.Configs, m.Inputs.Configs); err != nil {
		return fmt.Errorf("invalid config configuration: %s", err)
	}
	if err := validateConfigTypeOptions(m.Configs, m.Inputs.Configs); err != nil {
		return fmt.Errorf("invalid config type options configuration: %s", err)
	}
	if err := validateFiles(m.Files, m.Inputs.Files); err != nil {
		return fmt.Errorf("invalid file configuration: %s", err)
	}
	if err := validateFileGroups(m.FileGroups, m.Inputs.FileGroups); err != nil {
		return fmt.Errorf("invalid file group configuration: %s", err)
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
	if err := validateInputsAndGroups(m.Inputs.Files, m.Inputs.Groups); err != nil {
		return fmt.Errorf("invalid input group configuration: %s", err)
	}
	if err := validateInputsAndGroups(m.Inputs.FileGroups, m.Inputs.Groups); err != nil {
		return fmt.Errorf("invalid input group configuration: %s", err)
	}
	if err := validateInputsResources(m.Inputs.Resources, m.HostResources); err != nil {
		return fmt.Errorf("invalid input configuration: %s", err)
	}
	if err := validateInputsSecrets(m.Inputs.Secrets, m.Secrets); err != nil {
		return fmt.Errorf("invalid input configuration: %s", err)
	}
	if err := validateInputsConfigs(m.Inputs.Configs, m.Configs); err != nil {
		return fmt.Errorf("invalid input configuration: %s", err)
	}
	if err := validateInputsFiles(m.Inputs.Files, m.Files); err != nil {
		return fmt.Errorf("invalid input configuration: %s", err)
	}
	if err := validateInputsFileGroups(m.Inputs.FileGroups, m.FileGroups); err != nil {
		return fmt.Errorf("invalid input configuration: %s", err)
	}
	if err := validateServices(m.Services, m.Volumes, m.HostResources, m.Secrets, m.Configs, m.Dependencies, m.Files, m.FileGroups); err != nil {
		return err
	}
	if err := validateAuxServices(m.AuxServices, m.Volumes, m.Configs, m.Dependencies, m.Services); err != nil {
		return err
	}
	if err := validateAuxImgSrc(m.AuxImgSrc); err != nil {
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
	_, ok := model.ModuleTypeMap[s]
	return ok
}

func isValidDeploymentType(s string) bool {
	_, ok := model.DeploymentTypeMap[s]
	return ok
}

func isValidModuleID(s string) bool {
	re := regexp.MustCompile(`^(?:[a-zA-Z0-9-_]+)\.(?:[a-zA-Z]+)(?:\/[a-zA-Z0-9-_]+)+$`)
	return re.MatchString(s)
}

func isValidCPUArch(s string) bool {
	_, ok := model.CPUArchMap[s]
	return ok
}

func validateResources(mRs map[string]model.HostResource, inputs map[string]model.Input) error {
	for ref, r := range mRs {
		if ref == "" {
			return errors.New("empty resource reference")
		}
		if r.Required && len(r.Tags) == 0 {
			if _, ok := inputs[ref]; !ok {
				return fmt.Errorf("resource '%s' is required but no tags or input defined", ref)
			}
		}
	}
	return nil
}

func validateSecrets(mSs map[string]model.Secret, inputs map[string]model.Input) error {
	for ref, s := range mSs {
		if ref == "" {
			return errors.New("empty secret reference")
		}
		if s.Required && len(s.Tags) == 0 {
			if _, ok := inputs[ref]; !ok {
				return fmt.Errorf("secret '%s' is required but no tags or input defined", ref)
			}
		}
	}
	return nil
}

func validateConfigs(mCs model.Configs, inputs map[string]model.Input) error {
	for ref, cv := range mCs {
		if ref == "" {
			return errors.New("empty config reference")
		}
		if cv.Required && cv.Default == nil {
			if _, ok := inputs[ref]; !ok {
				return fmt.Errorf("config '%s' is required but no default value or input defined", ref)
			}
		}
	}
	return nil
}

func validateFiles(mFs map[string]string, inputs map[string]model.Input) error {
	for ref := range mFs {
		if ref == "" {
			return errors.New("empty file reference")
		}
		if _, ok := inputs[ref]; !ok {
			return fmt.Errorf("file '%s' no input defined", ref)
		}
	}
	return nil
}

func validateFileGroups(mFs map[string]struct{}, inputs map[string]model.Input) error {
	for ref := range mFs {
		if ref == "" {
			return errors.New("empty file group reference")
		}
		if _, ok := inputs[ref]; !ok {
			return fmt.Errorf("file group '%s' no input defined", ref)
		}
	}
	return nil
}

func validateAuxImgSrc(sources map[string]struct{}) error {
	re := regexp.MustCompile(`^[a-zA-Z0-9\/_\.-]+(?:$|\*$)`)
	for src := range sources {
		if !re.MatchString(src) {
			return fmt.Errorf("invalid aux service image source '%s'", src)
		}
	}
	return nil
}

func validateConfigTypeOptions(mCs model.Configs, inputs map[string]model.Input) error {
	for ref, cv := range mCs {
		if _, ok := inputs[ref]; ok {
			if err := configs.ValidateBase(cv.Type, cv.TypeOpt, cv.DataType); err != nil {
				return err
			}
			if err := configs.ValidateTypeOptions(cv.Type, cv.TypeOpt); err != nil {
				return err
			}
		}
	}
	return nil
}
