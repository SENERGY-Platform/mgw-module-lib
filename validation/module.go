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
	"github.com/SENERGY-Platform/mgw-module-lib/model"
	"github.com/SENERGY-Platform/mgw-module-lib/tsort"
	"github.com/SENERGY-Platform/mgw-module-lib/validation/sem_ver"
	"regexp"
)

func Validate(m model.Module) error {
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
	refVars := make(map[string]struct{})
	mntPts := make(map[string]struct{})
	extPaths := make(map[string]struct{})
	hostPorts := make(map[string]struct{})
	nodes := make(tsort.Nodes)
	for ref, service := range m.Services {
		if ref == "" {
			return errors.New("empty service reference")
		}
		if err := validateMapKeys(service.Include, mntPts); err != nil {
			return fmt.Errorf("service '%s' invalid include mount point configuration: %s", ref, err)
		}
		if err := validateMapKeys(service.Tmpfs, mntPts); err != nil {
			return fmt.Errorf("service '%s' invalid tmpfs mount point configuration: %s", ref, err)
		}
		if err := validateMapKeys(service.Volumes, mntPts); err != nil {
			return fmt.Errorf("service '%s' invalid volume mount point configuration: %s", ref, err)
		}
		if err := validateMapKeys(service.Resources, mntPts); err != nil {
			return fmt.Errorf("service '%s' invalid resource mount point configuration: %s", ref, err)
		}
		if err := validateMapKeys(service.Secrets, mntPts); err != nil {
			return fmt.Errorf("service '%s' invalid secret mount point configuration: %s", ref, err)
		}
		if err := validateMapKeys(service.Configs, refVars); err != nil {
			return fmt.Errorf("service '%s' invalid config reference variable configuration: %s", ref, err)
		}
		if err := validateMapKeys(service.SrvReferences, refVars); err != nil {
			return fmt.Errorf("service '%s' invalid service reference variable configuration: %s", ref, err)
		}
		if err := validateMapKeys(service.ExternalDependencies, refVars); err != nil {
			return fmt.Errorf("service '%s' invalid external dependency reference variable configuration: %s", ref, err)
		}
		if err := validateServiceVolumes(service.Volumes, m.Volumes); err != nil {
			return fmt.Errorf("service '%s' invalid volume configuration: %s", ref, err)
		}
		if err := validateServiceResources(service.Resources, m.Resources); err != nil {
			return fmt.Errorf("service '%s' invalid resource configuration: %s", ref, err)
		}
		if err := validateServiceSecrets(service.Secrets, m.Secrets); err != nil {
			return fmt.Errorf("service '%s' invalid secret configuration: %s", ref, err)
		}
		if err := validateServiceConfigs(service.Configs, m.Configs); err != nil {
			return fmt.Errorf("service '%s' invalid config configuration: %s", ref, err)
		}
		if err := validateServiceHttpEndpoints(service.HttpEndpoints, extPaths); err != nil {
			return fmt.Errorf("service '%s' invalid http endpoint configuration: %s", ref, err)
		}
		if err := validateServiceReferences(service.SrvReferences, m.Services); err != nil {
			return fmt.Errorf("service '%s' invalid reference configuration: %s", ref, err)
		}
		if err := validateServiceDependencies(service.RequiredSrv, service.RequiredBySrv, m.Services); err != nil {
			return fmt.Errorf("service '%s' invalid dependency configuration: %s", ref, err)
		}
		if err := validateServiceExternalDependencies(service.ExternalDependencies, m.Dependencies); err != nil {
			return fmt.Errorf("service '%s' invalid external dependency configuration: %s", ref, err)
		}
		if err := validateServicePorts(service.Ports, hostPorts); err != nil {
			return fmt.Errorf("service '%s' invalid port mapping configuration: %s", ref, err)
		}
		nodes.Add(ref, service.RequiredSrv, service.RequiredBySrv)
	}
	_, err := tsort.GetTopOrder(nodes)
	if err != nil {
		return fmt.Errorf("invalid service startup configuration: %s", err)
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
