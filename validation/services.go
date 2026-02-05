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

	"github.com/SENERGY-Platform/mgw-module-lib/model"
	tsort2 "github.com/SENERGY-Platform/mgw-module-lib/util/tsort"
)

func validateServices(
	mServices map[string]model.Service,
	mVolumes map[string]struct{},
	mResources map[string]model.HostResource,
	mSecrets map[string]model.Secret,
	mConfigs model.Configs,
	mDependencies map[string]string,
	mFiles map[string]string,
	mFileGroups map[string]struct{},
) error {
	extPaths := make(map[string]struct{})
	hostPorts := make(map[string]struct{})
	nodes := make(tsort2.Nodes)
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
		if err := validateMapKeys(service.Files, mntPts); err != nil {
			return fmt.Errorf("service '%s' invalid file mount point configuration: %s", ref, err)
		}
		if err := validateMapKeys(service.FileGroups, mntPts); err != nil {
			return fmt.Errorf("service '%s' invalid file group mount point configuration: %s", ref, err)
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
		if err := validateServiceFiles(service.Files, mFiles); err != nil {
			return fmt.Errorf("service '%s' invalid file configuration: %s", ref, err)
		}
		if err := validateServiceFileGroups(service.FileGroups, mFileGroups); err != nil {
			return fmt.Errorf("service '%s' invalid file configuration: %s", ref, err)
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
	_, err := tsort2.GetTopOrder(nodes)
	if err != nil {
		return fmt.Errorf("invalid service startup configuration: %s", err)
	}
	return nil
}

func validateAuxServices(auxServices map[string]model.AuxService, mVolumes map[string]struct{}, mConfigs model.Configs, mDependencies map[string]string, mServices map[string]model.Service) error {
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
	if len(sVolumes) > 0 && len(mVolumes) == 0 {
		return errors.New("no volumes defined")
	}
	for _, volume := range sVolumes {
		if _, ok := mVolumes[volume]; !ok {
			return fmt.Errorf("volume '%s' not defined", volume)
		}
	}
	return nil
}

func validateServiceResources(sResources map[string]model.HostResTarget, mResources map[string]model.HostResource) error {
	if len(sResources) > 0 && len(mResources) == 0 {
		return errors.New("no resources defined")
	}
	for _, target := range sResources {
		if _, ok := mResources[target.Ref]; !ok {
			return fmt.Errorf("resource '%s' not defined", target.Ref)
		}
	}
	return nil
}

func validateServiceSecrets(sSecretMounts, sSecretVars map[string]model.SecretTarget, mSecrets map[string]model.Secret) error {
	if len(sSecretMounts)+len(sSecretVars) > 0 && len(mSecrets) == 0 {
		return errors.New("no secrets defined")
	}
	for _, target := range sSecretMounts {
		if _, ok := mSecrets[target.Ref]; !ok {
			return fmt.Errorf("secret '%s' not defined", target.Ref)
		}
	}
	for _, target := range sSecretVars {
		if _, ok := mSecrets[target.Ref]; !ok {
			return fmt.Errorf("secret '%s' not defined", target.Ref)
		}
	}
	return nil
}

func validateServiceConfigs(sConfigs map[string]string, mConfigs model.Configs) error {
	if len(sConfigs) > 0 && len(mConfigs) == 0 {
		return errors.New("no configs defined")
	}
	for _, confRef := range sConfigs {
		if _, ok := mConfigs[confRef]; !ok {
			return fmt.Errorf("config '%s' not defined", confRef)
		}
	}
	return nil
}

func validateServiceFiles(sFiles map[string]string, mFiles map[string]string) error {
	if len(sFiles) > 0 && len(mFiles) == 0 {
		return errors.New("no files defined")
	}
	for _, ref := range sFiles {
		if _, ok := mFiles[ref]; !ok {
			return fmt.Errorf("file '%s' not defined", ref)
		}
	}
	return nil
}

func validateServiceFileGroups(sFileGroups map[string]string, mFileGroups map[string]struct{}) error {
	if len(sFileGroups) > 0 && len(mFileGroups) == 0 {
		return errors.New("no file groups defined")
	}
	for _, ref := range sFileGroups {
		if _, ok := mFileGroups[ref]; !ok {
			return fmt.Errorf("file group '%s' not defined", ref)
		}
	}
	return nil
}

func validateServiceHttpEndpoints(sHttpEndpoints map[string]model.HttpEndpoint, extPaths map[string]struct{}) error {
	for extPath, ept := range sHttpEndpoints {
		if !isValidExtPath(extPath) {
			return fmt.Errorf("invalid external path '%s'", extPath)
		}
		if ept.Path != "" && !isValidPath(ept.Path) {
			return fmt.Errorf("invalid internal path '%s'", ept.Path)
		}
		if _, ok := extPaths[extPath]; ok {
			return fmt.Errorf("duplicate path '%s'", extPath)
		}
		mt := make(map[string]struct{})
		for _, t := range ept.StringSub.MimeTypes {
			if _, ok := mt[t]; ok {
				return fmt.Errorf("duplicate mime type '%s'", t)
			}
			mt[t] = struct{}{}
		}
		extPaths[extPath] = struct{}{}
	}
	return nil
}

func validateServiceDependencies(requiredSrv map[string]struct{}, requiredBySrv map[string]struct{}, mServices map[string]model.Service) error {
	if len(requiredSrv)+len(requiredBySrv) > 0 && len(mServices) == 0 {
		return errors.New("no services defined")
	}
	for srvRef := range requiredSrv {
		if _, ok := mServices[srvRef]; !ok {
			return fmt.Errorf("service '%s' not defined", srvRef)
		}
	}
	for srvRef := range requiredBySrv {
		if _, ok := mServices[srvRef]; !ok {
			return fmt.Errorf("service '%s' not defined", srvRef)
		}
	}
	return nil
}

func validateServiceExternalDependencies(sExtDependencies map[string]model.ExtDependencyTarget, mDependencies map[string]string) error {
	if len(sExtDependencies) > 0 && len(mDependencies) == 0 {
		return errors.New("no module dependencies defined")
	}
	for _, target := range sExtDependencies {
		if target.Service == "" {
			return errors.New("empty service reference")
		}
		if _, ok := mDependencies[target.ID]; !ok {
			return fmt.Errorf("module dependency '%s' not defined", target.ID)
		}
	}
	return nil
}

func validateServiceReferences(sReferences map[string]model.SrvRefTarget, mServices map[string]model.Service) error {
	if len(sReferences) > 0 && len(mServices) == 0 {
		return errors.New("no services defined")
	}
	for _, target := range sReferences {
		if _, ok := mServices[target.Ref]; !ok {
			return fmt.Errorf("service '%s' not defined", target.Ref)
		}
	}
	return nil
}

func genPortKey(n int, p model.PortProtocol) string {
	return fmt.Sprintf("%d%s", n, p)
}

func validateServicePorts(sPorts []model.Port, hostPorts map[string]struct{}) error {
	expPorts := make(map[string]struct{})
	for _, port := range sPorts {
		if _, ok := model.PortProtocolMap[port.Protocol]; !ok {
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
