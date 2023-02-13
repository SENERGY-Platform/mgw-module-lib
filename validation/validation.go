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
	if err := validateKeyNotEmptyString(m.Volumes, "invalid volume name"); err != nil {
		return err
	}
	if err := validateModuleDependencies(m.Dependencies); err != nil {
		return err
	}
	if err := validateKeyNotEmptyString(m.Resources, "invalid resource reference"); err != nil {
		return err
	}
	if err := validateKeyNotEmptyString(m.Secrets, "invalid secret reference"); err != nil {
		return err
	}
	if err := validateKeyNotEmptyString(m.Configs, "invalid config reference"); err != nil {
		return err
	}
	if err := validateKeyNotEmptyString(m.Inputs.Groups, "invalid input group reference"); err != nil {
		return err
	}
	if err := validateInputs(m.Inputs.Resources, m.Resources, "resource", m.Inputs.Groups); err != nil {
		return err
	}
	if err := validateInputs(m.Inputs.Secrets, m.Secrets, "secret", m.Inputs.Groups); err != nil {
		return err
	}
	if err := validateInputs(m.Inputs.Configs, m.Configs, "config", m.Inputs.Groups); err != nil {
		return err
	}
	if m.Services == nil || len(m.Services) == 0 {
		return errors.New("missing services")
	}
	hostPorts := make(map[uint]struct{})
	for ref, service := range m.Services {
		if ref == "" {
			return errors.New("invalid service reference")
		}
		if err := validateServiceMountPoints(service); err != nil {
			return fmt.Errorf("invalid service mount point: '%s' -> %s", ref, err)
		}
		if err := validateServiceRefVars(service); err != nil {
			return fmt.Errorf("invalid service reference variable: '%s' -> %s", ref, err)
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
		if service.SrvReferences != nil {
			for refVar, s := range service.SrvReferences {
				if _, ok := m.Services[s]; !ok {
					return fmt.Errorf("invalid service reference: '%s' -> '%s' -> '%s'", ref, refVar, s)
				}
			}
		}
		if service.Dependencies != nil {
			for s := range service.Dependencies {
				if _, ok := m.Services[s]; !ok {
					return fmt.Errorf("invalid service dependency: '%s' -> '%s'", ref, s)
				}
			}
		}
		if service.ExternalDependencies != nil {
			for _, target := range service.ExternalDependencies {
				if !isValidModuleID(target.ID) {
					return fmt.Errorf("invalid service external dependency: '%s' -> '%s'", ref, target.ID)
				}
			}
		}
		if service.PortMappings != nil {
			for _, pm := range service.PortMappings {
				if pm.HostPort != nil && len(pm.HostPort) > 0 {
					if len(pm.HostPort) > 1 {
						for i := pm.HostPort[0]; i <= pm.HostPort[1]; i++ {
							if _, ok := hostPorts[i]; ok {
								return fmt.Errorf("duplicate host port '%d'", i)
							}
							hostPorts[i] = struct{}{}
						}
					} else {
						if _, ok := hostPorts[pm.HostPort[0]]; ok {
							return fmt.Errorf("duplicate host port '%d'", pm.HostPort[0])
						}
						hostPorts[pm.HostPort[0]] = struct{}{}
					}
				}
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

func validateModuleDependencies(dependencies map[string]model.ModuleDependency) error {
	if dependencies != nil {
		for mid, dependency := range dependencies {
			if !isValidModuleID(mid) {
				return fmt.Errorf("invalid dependency module ID format '%s'", mid)
			}
			if err := sem_ver.ValidateSemVerRange(dependency.Version); err != nil {
				return fmt.Errorf("dependency '%s': %s", mid, err)
			}
			if dependency.RequiredServices == nil {
				return fmt.Errorf("missing services for dependency '%s'", mid)
			}
			if err := validateKeyNotEmptyString(dependency.RequiredServices, fmt.Sprintf("invalid service for dependency '%s'", mid)); err != nil {
				return err
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

func validateServiceExternalDependencies(sExtDependencies map[string]model.ExternalDependencyTarget, mDependencies map[string]model.ModuleDependency) error {
	if sExtDependencies != nil {
		if mDependencies == nil {
			return errors.New("no module dependencies defined")
		}
		for _, target := range sExtDependencies {
			mDep, ok := mDependencies[target.ID]
			if !ok {
				return fmt.Errorf("module dependency '%s' not defined", target.ID)
			}
			if _, k := mDep.RequiredServices[target.Service]; !k {
				return fmt.Errorf("module dependency '%s' service '%s' not defined", target.ID, target.Service)
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

func validateServiceRefVars(service *model.Service) error {
	refVars := make(map[string]string)
	if service.Configs != nil {
		for rv := range service.Configs {
			if rv == "" {
				return errors.New("invalid ref var")
			}
			refVars[rv] = "configs"
		}
	}
	if service.SrvReferences != nil {
		for rv := range service.SrvReferences {
			if rv == "" {
				return errors.New("invalid ref var")
			}
			if v, ok := refVars[rv]; ok {
				return fmt.Errorf("'%s' -> '%s' & '%s'", rv, v, "dependencies")
			}
			refVars[rv] = "dependencies"
		}
	}
	if service.ExternalDependencies != nil {
		for rv := range service.ExternalDependencies {
			if rv == "" {
				return errors.New("invalid ref var")
			}
			if v, ok := refVars[rv]; ok {
				return fmt.Errorf("'%s' -> '%s' & '%s'", rv, v, "external dependencies")
			}
			refVars[rv] = "external dependencies"
		}
	}
	return nil
}

func validateServiceMountPoints(service *model.Service) error {
	mountPoints := make(map[string]string)
	if service.Include != nil {
		for mp := range service.Include {
			if mp == "" {
				return errors.New("invalid mount point")
			}
			mountPoints[mp] = "include"
		}
	}
	if service.Tmpfs != nil {
		for mp := range service.Tmpfs {
			if mp == "" {
				return errors.New("invalid mount point")
			}
			if v, ok := mountPoints[mp]; ok {
				return fmt.Errorf("'%s' -> '%s' & '%s'", mp, v, "tmpfs")
			}
			mountPoints[mp] = "tmpfs"
		}
	}
	if service.Volumes != nil {
		for mp := range service.Volumes {
			if mp == "" {
				return errors.New("invalid mount point")
			}
			if v, ok := mountPoints[mp]; ok {
				return fmt.Errorf("'%s' -> '%s' & '%s'", mp, v, "volumes")
			}
			mountPoints[mp] = "volumes"
		}
	}
	if service.Resources != nil {
		for mp := range service.Resources {
			if mp == "" {
				return errors.New("invalid mount point")
			}
			if v, ok := mountPoints[mp]; ok {
				return fmt.Errorf("'%s' -> '%s' & '%s'", mp, v, "resources")
			}
			mountPoints[mp] = "resources"
		}
	}
	if service.Secrets != nil {
		for mp := range service.Secrets {
			if mp == "" {
				return errors.New("invalid mount point")
			}
			if v, ok := mountPoints[mp]; ok {
				return fmt.Errorf("'%s' -> '%s' & '%s'", mp, v, "secrets")
			}
			mountPoints[mp] = "secrets"
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

func validateInputs[T any](inputs map[string]model.Input, refs map[string]T, refName string, groups map[string]model.InputGroup) error {
	if inputs != nil {
		if refs == nil {
			return fmt.Errorf("missing %ss for user inputs", refName)
		}
		for ref, input := range inputs {
			if ref == "" {
				return errors.New("invalid input reference")
			}
			if _, ok := refs[ref]; !ok {
				return fmt.Errorf("missing %s for input '%s'", refName, ref)
			}
			if input.Group != nil {
				if groups == nil {
					return errors.New("missing groups for inputs")
				}
				if _, ok := groups[*input.Group]; !ok {
					return fmt.Errorf("missing group for input '%s'", ref)
				}
			}
		}
	}
	return nil
}

func validateKeyNotEmptyString[T any](m map[string]T, msg string) error {
	if m != nil {
		for ref := range m {
			if ref == "" {
				return errors.New(msg)
			}
		}
	}
	return nil
}
