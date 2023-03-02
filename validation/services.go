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
)

func validateServiceVolumes(sVolumes map[string]string, mVolumes map[string]struct{}) error {
	for _, volume := range sVolumes {
		if mVolumes != nil {
			if _, ok := mVolumes[volume]; !ok {
				return fmt.Errorf("volume '%s' not defined", volume)
			}
		} else {
			return errors.New("no volumes defined")
		}
	}
	return nil
}

func validateServiceResources(sResources map[string]module.ResourceTarget, mResources map[string]map[string]struct{}) error {
	for _, target := range sResources {
		if mResources != nil {
			if _, ok := mResources[target.Ref]; !ok {
				return fmt.Errorf("resource '%s' not defined", target.Ref)
			}
		} else {
			return errors.New("no resources defined")
		}
	}
	return nil
}

func validateServiceSecrets(sSecrets map[string]string, mSecrets map[string]module.Secret) error {
	for _, secretRef := range sSecrets {
		if mSecrets != nil {
			if _, ok := mSecrets[secretRef]; !ok {
				return fmt.Errorf("secret '%s' not defined", secretRef)
			}
		} else {
			return errors.New("no secrets defined")
		}
	}
	return nil
}

func validateServiceConfigs(sConfigs map[string]string, mConfigs module.Configs) error {
	for _, confRef := range sConfigs {
		if mConfigs != nil {
			if _, ok := mConfigs[confRef]; !ok {
				return fmt.Errorf("config '%s' not defined", confRef)
			}
		} else {
			return errors.New("no configs defined")
		}
	}
	return nil
}

func validateServiceHttpEndpoints(sHttpEndpoints map[string]module.HttpEndpoint, extPaths map[string]struct{}) error {
	for extPath, ept := range sHttpEndpoints {
		if !isValidPath(extPath) {
			return fmt.Errorf("invalid external path '%s'", extPath)
		}
		if !isValidPath(ept.Path) {
			return fmt.Errorf("invalid internal path '%s'", ept.Path)
		}
		if _, ok := extPaths[extPath]; ok {
			return fmt.Errorf("duplicate path '%s'", extPath)
		}
		extPaths[extPath] = struct{}{}
	}
	return nil
}

func validateServiceDependencies(requiredSrv map[string]struct{}, requiredBySrv map[string]struct{}, mServices map[string]*module.Service) error {
	for srvRef := range requiredSrv {
		if mServices != nil {
			if _, ok := mServices[srvRef]; !ok {
				return fmt.Errorf("service '%s' not defined", srvRef)
			}
		} else {
			return errors.New("no services defined")
		}
	}
	for srvRef := range requiredBySrv {
		if mServices != nil {
			if _, ok := mServices[srvRef]; !ok {
				return fmt.Errorf("service '%s' not defined", srvRef)
			}
		} else {
			return errors.New("no services defined")
		}
	}
	return nil
}

func validateServiceExternalDependencies(sExtDependencies map[string]module.ExtDependencyTarget, mDependencies map[string]string) error {
	for _, target := range sExtDependencies {
		if target.Service == "" {
			return errors.New("empty service reference")
		}
		if mDependencies != nil {
			if _, ok := mDependencies[target.ID]; !ok {
				return fmt.Errorf("module dependency '%s' not defined", target.ID)
			}
		} else {
			return errors.New("no module dependencies defined")
		}
	}
	return nil
}

func validateServiceReferences(sReferences map[string]string, mServices map[string]*module.Service) error {
	for _, srvRef := range sReferences {
		if mServices != nil {
			if _, ok := mServices[srvRef]; !ok {
				return fmt.Errorf("service '%s' not defined", srvRef)
			}
		} else {
			return errors.New("no services defined")
		}
	}
	return nil
}

func genPortKey(n uint, p module.PortProtocol) string {
	return fmt.Sprintf("%d%s", n, p)
}

func validateServicePorts(sPorts []module.Port, hostPorts map[string]struct{}) error {
	expPorts := make(map[string]struct{})
	for _, port := range sPorts {
		if _, ok := module.PortProtocolMap[port.Protocol]; !ok {
			return fmt.Errorf("invalid protocol '%s'", port.Protocol)
		}
		pKey := genPortKey(port.Number, port.Protocol)
		if _, ok := expPorts[pKey]; ok {
			return fmt.Errorf("duplicate port '%d/%s'", port.Number, port.Protocol)
		}
		expPorts[pKey] = struct{}{}
		for _, binding := range port.Bindings {
			bpKey := genPortKey(binding, port.Protocol)
			if _, ok := hostPorts[bpKey]; ok {
				return fmt.Errorf("duplicate port binding '%d/%s'", binding, port.Protocol)
			}
			hostPorts[bpKey] = struct{}{}
		}
	}
	return nil
}
