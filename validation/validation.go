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
	return validateServices(m)
}

func validateServices(m model.Module) error {
	if m.Services != nil {
		hostPorts := make(map[uint]struct{})
		refVars := make(map[string]struct{})
		mntPts := make(map[string]struct{})
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
			if err := validateServiceHttpEndpoints(service.HttpEndpoints); err != nil {
				return fmt.Errorf("service '%s' invalid http endpoint configuration: %s", ref, err)
			}
			if err := validateServiceReferences(service.SrvReferences, m.Services); err != nil {
				return fmt.Errorf("service '%s' invalid reference configuration: %s", ref, err)
			}
			if err := validateServiceDependencies(service.Dependencies, m.Services); err != nil {
				return fmt.Errorf("service '%s' invalid dependency configuration: %s", ref, err)
			}
			if err := validateServiceExternalDependencies(service.ExternalDependencies, m.Dependencies); err != nil {
				return fmt.Errorf("service '%s' invalid external dependency configuration: %s", ref, err)
			}
			if err := validateServicePortMappings(service.PortMappings, hostPorts); err != nil {
				return fmt.Errorf("service '%s' invalid port mapping configuration: %s", ref, err)
			}
		}
	}
	return nil
}

func validateServiceVolumes(sVolumes map[string]string, mVolumes model.Set[string]) error {
	if sVolumes != nil {
		if mVolumes == nil {
			return errors.New("no volumes defined")
		}
		for _, volume := range sVolumes {
			if _, ok := mVolumes[volume]; !ok {
				return fmt.Errorf("volume '%s' not defined", volume)
			}
		}
	}
	return nil
}

func validateServiceResources(sResources map[string]model.ResourceTarget, mResources map[string]model.Set[string]) error {
	if sResources != nil {
		if mResources == nil {
			return errors.New("no resources defined")
		}
		for _, target := range sResources {
			if _, ok := mResources[target.Ref]; !ok {
				return fmt.Errorf("resource '%s' not defined", target.Ref)
			}
		}
	}
	return nil
}

func validateModuleDependencies(dependencies map[string]string) error {
	if dependencies != nil {
		for mid, ver := range dependencies {
			if !isValidModuleID(mid) {
				return fmt.Errorf("invalid module ID format '%s'", mid)
			}
			if err := sem_ver.ValidateSemVerRange(ver); err != nil {
				return fmt.Errorf("version %s", err)
			}
		}
	}
	return nil
}

func validateInputsResources(inputs map[string]model.Input, mResources map[string]model.Set[string]) error {
	if inputs != nil {
		if mResources == nil {
			return errors.New("no resources defined")
		}
		for ref := range inputs {
			if ref == "" {
				return errors.New("invalid input reference")
			}
			if _, ok := mResources[ref]; !ok {
				return fmt.Errorf("resource '%s' not defined", ref)
			}
		}
	}
	return nil
}

func validateInputsSecrets(inputs map[string]model.Input, mSecrets map[string]model.Secret) error {
	if inputs != nil {
		if mSecrets == nil {
			return errors.New("no secrets defined")
		}
		for ref := range inputs {
			if ref == "" {
				return errors.New("invalid input reference")
			}
			if _, ok := mSecrets[ref]; !ok {
				return fmt.Errorf("secret '%s' not defined", ref)
			}
		}
	}
	return nil
}

func validateInputsConfigs(inputs map[string]model.Input, mConfigs model.Configs) error {
	if inputs != nil {
		if mConfigs == nil {
			return errors.New("no configs defined")
		}
		for ref := range inputs {
			if ref == "" {
				return errors.New("invalid input reference")
			}
			if _, ok := mConfigs[ref]; !ok {
				return fmt.Errorf("config '%s' not defined", ref)
			}
		}
	}
	return nil
}

func validateInputsAndGroups(inputs map[string]model.Input, groups map[string]model.InputGroup) error {
	for _, input := range inputs {
		if input.Group != nil {
			if groups == nil {
				return errors.New("no input groups defined")
			}
			if _, ok := groups[*input.Group]; !ok {
				return fmt.Errorf("input group '%s' not defined", *input.Group)
			}
		}
	}
	return nil
}

func validateInputGroups(groups map[string]model.InputGroup) error {
	if groups != nil {
		for ref, group := range groups {
			if ref == "" {
				return errors.New("invalid input group reference")
			}
			if group.Group != nil {
				if _, ok := groups[*group.Group]; !ok {
					return fmt.Errorf("input group '%s' not defined", *group.Group)
				}
			}
		}
	}
	return nil
}

func validateServiceSecrets(sSecrets map[string]string, mSecrets map[string]model.Secret) error {
	if sSecrets != nil {
		if mSecrets == nil {
			return errors.New("no secrets defined")
		}
		for _, secretRef := range sSecrets {
			if _, ok := mSecrets[secretRef]; !ok {
				return fmt.Errorf("secret '%s' not defined", secretRef)
			}
		}
	}
	return nil
}

func validateServiceConfigs(sConfigs map[string]string, mConfigs model.Configs) error {
	if sConfigs != nil {
		if mConfigs == nil {
			return errors.New("no configs defined")
		}
		for _, confRef := range sConfigs {
			if _, ok := mConfigs[confRef]; !ok {
				return fmt.Errorf("config '%s' not defined", confRef)
			}
		}
	}
	return nil
}

func validateServiceHttpEndpoints(sHttpEndpoints map[string]model.HttpEndpoint) error {
	if sHttpEndpoints != nil {
		extPaths := make(map[string]struct{})
		for extPath := range sHttpEndpoints {
			if !isValidPath(extPath) {
				return fmt.Errorf("invalid path '%s'", extPath)
			}
			if _, ok := extPaths[extPath]; ok {
				return fmt.Errorf("duplicate path '%s'", extPath)
			}
			extPaths[extPath] = struct{}{}
		}
	}
	return nil
}

func validateServiceDependencies(sDependencies model.Set[string], mServices map[string]*model.Service) error {
	if sDependencies != nil {
		if mServices == nil {
			return errors.New("no services defined")
		}
		for srvRef := range sDependencies {
			if _, ok := mServices[srvRef]; !ok {
				return fmt.Errorf("service '%s' not defined", srvRef)
			}
		}
	}
	return nil
}

func validateServiceExternalDependencies(sExtDependencies map[string]model.ExternalDependencyTarget, mDependencies map[string]string) error {
	if sExtDependencies != nil {
		if mDependencies == nil {
			return errors.New("no module dependencies defined")
		}
		for _, target := range sExtDependencies {
			if _, ok := mDependencies[target.ID]; !ok {
				return fmt.Errorf("module dependency '%s' not defined", target.ID)
			}
		}
	}
	return nil
}

func validateServiceReferences(sReferences map[string]string, mServices map[string]*model.Service) error {
	if sReferences != nil {
		if mServices == nil {
			return errors.New("no services defined")
		}
		for _, srvRef := range sReferences {
			if _, ok := mServices[srvRef]; !ok {
				return fmt.Errorf("service '%s' not defined", srvRef)
			}
		}
	}
	return nil
}

func validateServicePortMappings(sPortMappings model.PortMappings, hostPorts map[uint]struct{}) error {
	if sPortMappings != nil {
		for _, pm := range sPortMappings {
			if pm.HostPort != nil && len(pm.HostPort) > 0 {
				if len(pm.HostPort) > 1 {
					for i := pm.HostPort[0]; i <= pm.HostPort[1]; i++ {
						if _, ok := hostPorts[i]; ok {
							return fmt.Errorf("duplicate port mapping '%d'", i)
						}
						hostPorts[i] = struct{}{}
					}
				} else {
					if _, ok := hostPorts[pm.HostPort[0]]; ok {
						return fmt.Errorf("duplicate port mapping '%d'", pm.HostPort[0])
					}
					hostPorts[pm.HostPort[0]] = struct{}{}
				}
			}
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
	re := regexp.MustCompile(`^([a-z0-9A-Z-_.]+)(:\d+)?([\/a-zA-Z0-9-\.]+)?$`)
	return re.MatchString(s)
}

func isValidPath(s string) bool {
	re := regexp.MustCompile(`^\/(?:[a-zA-Z0-9-_%]+)*(?:\/[a-zA-Z0-9-_%]+)*$`)
	return re.MatchString(s)
}

func validateKeyNotEmptyString[T any](m map[string]T) bool {
	if m != nil {
		for ref := range m {
			if ref == "" {
				return false
			}
		}
	}
	return true
}

func validateMapKeys[T any](m map[string]T, keys map[string]struct{}) error {
	if m != nil {
		for k := range m {
			if k == "" {
				return errors.New("empty")
			}
			if _, ok := keys[k]; ok {
				return fmt.Errorf("duplicate '%s'", k)
			}
			keys[k] = struct{}{}
		}
	}
	return nil
}
