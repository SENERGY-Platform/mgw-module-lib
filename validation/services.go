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
)

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