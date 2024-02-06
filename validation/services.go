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
	"github.com/SENERGY-Platform/mgw-module-lib/tsort"
)

func validateServices(mServices map[string]*module.Service, mVolumes map[string]struct{}, mResources map[string]module.HostResource, mSecrets map[string]module.Secret, mConfigs module.Configs, mDependencies map[string]string) error {
	extPaths := make(map[string]struct{})
	hostPorts := make(map[string]struct{})
	nodes := make(tsort.Nodes)
	for ref, service := range mServices {
		refVars := make(map[string]struct{})
		mntPts := make(map[string]struct{})
		if ref == "" {
			return errors.New("empty service reference")
		}
		if err := validateMapKeys(service.BindMounts, mntPts); err != nil {
			return fmt.Errorf("service '%s' invalid include mount point configuration: %s", ref, err)
		}
		if err := validateMapKeys(service.Tmpfs, mntPts); err != nil {
			return fmt.Errorf("service '%s' invalid tmpfs mount point configuration: %s", ref, err)
		}
		if err := validateMapKeys(service.Volumes, mntPts); err != nil {
			return fmt.Errorf("service '%s' invalid volume mount point configuration: %s", ref, err)
		}
		if err := validateMapKeys(service.HostResources, mntPts); err != nil {
			return fmt.Errorf("service '%s' invalid resource mount point configuration: %s", ref, err)
		}
		if err := validateMapKeys(service.SecretMounts, mntPts); err != nil {
			return fmt.Errorf("service '%s' invalid secret mount point configuration: %s", ref, err)
		}
		if err := validateMapKeys(service.SecretVars, refVars); err != nil {
			return fmt.Errorf("service '%s' invalid secret reference variable configuration: %s", ref, err)
		}
		if err := validateMapKeys(service.Configs, refVars); err != nil {
			return fmt.Errorf("service '%s' invalid config reference variable configuration: %s", ref, err)
		}
		if err := validateMapKeys(service.SrvReferences, refVars); err != nil {
			return fmt.Errorf("service '%s' invalid service reference variable configuration: %s", ref, err)
		}
		if err := validateMapKeys(service.ExtDependencies, refVars); err != nil {
			return fmt.Errorf("service '%s' invalid external dependency reference variable configuration: %s", ref, err)
		}
		if err := validateServiceVolumes(service.Volumes, mVolumes); err != nil {
			return fmt.Errorf("service '%s' invalid volume configuration: %s", ref, err)
		}
		if err := validateServiceResources(service.HostResources, mResources); err != nil {
			return fmt.Errorf("service '%s' invalid resource configuration: %s", ref, err)
		}
		if err := validateServiceSecrets(service.SecretMounts, service.SecretVars, mSecrets); err != nil {
			return fmt.Errorf("service '%s' invalid secret configuration: %s", ref, err)
		}
		if err := validateServiceConfigs(service.Configs, mConfigs); err != nil {
			return fmt.Errorf("service '%s' invalid config configuration: %s", ref, err)
		}
		if err := validateServiceHttpEndpoints(service.HttpEndpoints, extPaths); err != nil {
			return fmt.Errorf("service '%s' invalid http endpoint configuration: %s", ref, err)
		}
		if err := validateServiceReferences(service.SrvReferences, mServices); err != nil {
			return fmt.Errorf("service '%s' invalid reference configuration: %s", ref, err)
		}
		if err := validateServiceDependencies(service.RequiredSrv, service.RequiredBySrv, mServices); err != nil {
			return fmt.Errorf("service '%s' invalid dependency configuration: %s", ref, err)
		}
		if err := validateServiceExternalDependencies(service.ExtDependencies, mDependencies); err != nil {
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

func validateAuxServices(auxServices map[string]*module.AuxService, mVolumes map[string]struct{}, mConfigs module.Configs, mDependencies map[string]string, mServices map[string]*module.Service) error {
	for ref, service := range auxServices {
		refVars := make(map[string]struct{})
		mntPts := make(map[string]struct{})
		if ref == "" {
			return errors.New("empty service reference")
		}
		if err := validateMapKeys(service.BindMounts, mntPts); err != nil {
			return fmt.Errorf("aux service '%s' invalid include mount point configuration: %s", ref, err)
		}
		if err := validateMapKeys(service.Tmpfs, mntPts); err != nil {
			return fmt.Errorf("aux service '%s' invalid tmpfs mount point configuration: %s", ref, err)
		}
		if err := validateMapKeys(service.Volumes, mntPts); err != nil {
			return fmt.Errorf("aux service '%s' invalid volume mount point configuration: %s", ref, err)
		}
		if err := validateMapKeys(service.Configs, refVars); err != nil {
			return fmt.Errorf("aux service '%s' invalid config reference variable configuration: %s", ref, err)
		}
		if err := validateMapKeys(service.SrvReferences, refVars); err != nil {
			return fmt.Errorf("aux service '%s' invalid service reference variable configuration: %s", ref, err)
		}
		if err := validateMapKeys(service.ExtDependencies, refVars); err != nil {
			return fmt.Errorf("aux service '%s' invalid external dependency reference variable configuration: %s", ref, err)
		}
		if err := validateServiceVolumes(service.Volumes, mVolumes); err != nil {
			return fmt.Errorf("aux service '%s' invalid volume configuration: %s", ref, err)
		}
		if err := validateServiceConfigs(service.Configs, mConfigs); err != nil {
			return fmt.Errorf("aux service '%s' invalid config configuration: %s", ref, err)
		}
		if err := validateServiceReferences(service.SrvReferences, mServices); err != nil {
			return fmt.Errorf("aux service '%s' invalid reference configuration: %s", ref, err)
		}
		if err := validateServiceExternalDependencies(service.ExtDependencies, mDependencies); err != nil {
			return fmt.Errorf("aux service '%s' invalid external dependency configuration: %s", ref, err)
		}
	}
	return nil
}

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

func validateServiceResources(sResources map[string]module.HostResTarget, mResources map[string]module.HostResource) error {
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

func validateServiceSecrets(sSecretMounts, sSecretVars map[string]module.SecretTarget, mSecrets map[string]module.Secret) error {
	for _, target := range sSecretMounts {
		if mSecrets != nil {
			if _, ok := mSecrets[target.Ref]; !ok {
				return fmt.Errorf("secret '%s' not defined", target.Ref)
			}
		} else {
			return errors.New("no secrets defined")
		}
	}
	for _, target := range sSecretVars {
		if mSecrets != nil {
			if _, ok := mSecrets[target.Ref]; !ok {
				return fmt.Errorf("secret '%s' not defined", target.Ref)
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
		if !isValidExtPath(extPath) {
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

func validateServiceReferences(sReferences map[string]module.SrvRefTarget, mServices map[string]*module.Service) error {
	for _, target := range sReferences {
		if mServices != nil {
			if _, ok := mServices[target.Ref]; !ok {
				return fmt.Errorf("service '%s' not defined", target.Ref)
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
