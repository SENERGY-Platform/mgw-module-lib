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
	if m.Dependencies != nil {
		for mid, dependency := range m.Dependencies {
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
		if service.Volumes != nil {
			for _, volume := range service.Volumes {
				if _, ok := m.Volumes[volume]; !ok {
					return fmt.Errorf("invalid service volume: '%s' -> '%s'", ref, volume)
				}
			}
		}
		if service.Resources != nil {
			for _, target := range service.Resources {
				if _, ok := m.Resources[target.Ref]; !ok {
					return fmt.Errorf("invalid service resource: '%s' -> '%s'", ref, target.Ref)
				}
			}
		}
		if service.Secrets != nil {
			for _, secretRef := range service.Secrets {
				if _, ok := m.Secrets[secretRef]; !ok {
					return fmt.Errorf("invalid service secret: '%s' -> '%s'", ref, secretRef)
				}
			}
		}
		if service.Configs != nil {
			for _, confRef := range service.Configs {
				if _, ok := m.Configs[confRef]; !ok {
					return fmt.Errorf("invalid service secret: '%s' -> '%s'", ref, confRef)
				}
			}
		}
		if service.HttpEndpoints != nil {
			extPaths := make(map[string]string)
			for extPath, edpt := range service.HttpEndpoints {
				if !isValidPath(extPath) {
					return fmt.Errorf("invalid external path '%s'", extPath)
				}
				if v, ok := extPaths[extPath]; ok {
					return fmt.Errorf("invalid service http endpoint: '%s' -> '%s' -> '%s' & '%s'", ref, extPath, v, edpt.Path)
				}
				extPaths[extPath] = edpt.Path
			}
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
