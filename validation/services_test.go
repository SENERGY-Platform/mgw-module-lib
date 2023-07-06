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
	"github.com/SENERGY-Platform/mgw-module-lib/module"
	"testing"
)

func TestValidateServices(t *testing.T) {
	s := map[string]*module.Service{
		"a": {},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil); err != nil {
		t.Error("err != nil")
	}
	// ------------------------------
	s = map[string]*module.Service{
		"": {},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]*module.Service{
		"a": {
			BindMounts: map[string]module.BindMount{"": {}},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]*module.Service{
		"a": {
			Tmpfs: map[string]module.TmpfsMount{"": {}},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]*module.Service{
		"a": {
			Volumes: map[string]string{"": ""},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]*module.Service{
		"a": {
			HostResources: map[string]module.HostResTarget{"": {}},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]*module.Service{
		"a": {
			SecretMounts: map[string]module.SecretTarget{"": {}},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]*module.Service{
		"a": {
			SecretVars: map[string]module.SecretTarget{"": {}},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]*module.Service{
		"a": {
			Configs: map[string]string{"": ""},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]*module.Service{
		"a": {
			SrvReferences: map[string]string{"": ""},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]*module.Service{
		"a": {
			ExtDependencies: map[string]module.ExtDependencyTarget{"": {}},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]*module.Service{
		"a": {
			Volumes: map[string]string{"test": ""},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]*module.Service{
		"a": {
			HostResources: map[string]module.HostResTarget{"test": {}},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]*module.Service{
		"a": {
			SecretMounts: map[string]module.SecretTarget{"test": {}},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]*module.Service{
		"a": {
			SecretVars: map[string]module.SecretTarget{"test": {}},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]*module.Service{
		"a": {
			Configs: map[string]string{"test": ""},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]*module.Service{
		"a": {
			HttpEndpoints: map[string]module.HttpEndpoint{"": {}},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]*module.Service{
		"a": {
			SrvReferences: map[string]string{"test": ""},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]*module.Service{
		"a": {
			RequiredSrv: map[string]struct{}{"": {}},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]*module.Service{
		"a": {
			ExtDependencies: map[string]module.ExtDependencyTarget{"test": {}},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]*module.Service{
		"a": {
			Ports: []module.Port{
				{},
			},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]*module.Service{
		"a": {
			RequiredSrv: map[string]struct{}{
				"a": {},
			},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
}

func TestValidateServiceVolumes(t *testing.T) {
	str := "test"
	var sVolumes map[string]string
	var mVolumes map[string]struct{}
	if err := validateServiceVolumes(sVolumes, mVolumes); err != nil {
		t.Errorf("validateServiceVolumes(%v, %v); err != nil", sVolumes, mVolumes)
	}
	sVolumes = make(map[string]string)
	if err := validateServiceVolumes(sVolumes, mVolumes); err != nil {
		t.Errorf("validateServiceVolumes(%v, %v); err != nil", sVolumes, mVolumes)
	}
	sVolumes["a"] = str
	if err := validateServiceVolumes(sVolumes, mVolumes); err == nil {
		t.Errorf("validateServiceVolumes(%v, %v); err == nil", sVolumes, mVolumes)
	}
	mVolumes = make(map[string]struct{})
	if err := validateServiceVolumes(sVolumes, mVolumes); err == nil {
		t.Errorf("validateServiceVolumes(%v, %v); err == nil", sVolumes, mVolumes)
	}
	mVolumes[str] = struct{}{}
	if err := validateServiceVolumes(sVolumes, mVolumes); err != nil {
		t.Errorf("validateServiceVolumes(%v, %v); err != nil", sVolumes, mVolumes)
	}
}

func TestValidateServiceResources(t *testing.T) {
	str := "test"
	var sResources map[string]module.HostResTarget
	var mResources map[string]module.HostResource
	if err := validateServiceResources(sResources, mResources); err != nil {
		t.Errorf("validateServiceResources(%v, %v); err != nil", sResources, mResources)
	}
	sResources = make(map[string]module.HostResTarget)
	if err := validateServiceResources(sResources, mResources); err != nil {
		t.Errorf("validateServiceResources(%v, %v); err != nil", sResources, mResources)
	}
	sResources["a"] = module.HostResTarget{Ref: str}
	if err := validateServiceResources(sResources, mResources); err == nil {
		t.Errorf("validateServiceResources(%v, %v); err == nil", sResources, mResources)
	}
	mResources = make(map[string]module.HostResource)
	if err := validateServiceResources(sResources, mResources); err == nil {
		t.Errorf("validateServiceResources(%v, %v); err == nil", sResources, mResources)
	}
	mResources[str] = module.HostResource{}
	if err := validateServiceResources(sResources, mResources); err != nil {
		t.Errorf("validateServiceResources(%v, %v); err != nil", sResources, mResources)
	}
}

func TestValidateServiceSecrets(t *testing.T) {
	str := "test"
	var sSecretMounts map[string]module.SecretTarget
	var sSecretVars map[string]module.SecretTarget
	var mSecrets map[string]module.Secret
	if err := validateServiceSecrets(sSecretMounts, sSecretVars, mSecrets); err != nil {
		t.Errorf("validateServiceSecrets(%v, %v); err != nil", sSecretMounts, mSecrets)
	}
	sSecretMounts = make(map[string]module.SecretTarget)
	sSecretVars = make(map[string]module.SecretTarget)
	if err := validateServiceSecrets(sSecretMounts, sSecretVars, mSecrets); err != nil {
		t.Errorf("validateServiceSecrets(%v, %v); err != nil", sSecretMounts, mSecrets)
	}
	sSecretMounts["a"] = module.SecretTarget{Ref: str}
	if err := validateServiceSecrets(sSecretMounts, sSecretVars, mSecrets); err == nil {
		t.Errorf("validateServiceSecrets(%v, %v); err == nil", sSecretMounts, mSecrets)
	}
	mSecrets = make(map[string]module.Secret)
	if err := validateServiceSecrets(sSecretMounts, sSecretVars, mSecrets); err == nil {
		t.Errorf("validateServiceSecrets(%v, %v); err == nil", sSecretMounts, mSecrets)
	}
	mSecrets[str] = module.Secret{}
	if err := validateServiceSecrets(sSecretMounts, sSecretVars, mSecrets); err != nil {
		t.Errorf("validateServiceSecrets(%v, %v); err != nil", sSecretMounts, mSecrets)
	}
	sSecretVars["a"] = module.SecretTarget{Ref: str}
	if err := validateServiceSecrets(sSecretMounts, sSecretVars, mSecrets); err != nil {
		t.Errorf("validateServiceSecrets(%v, %v); err != nil", sSecretMounts, mSecrets)
	}
}

func TestValidateServiceConfigs(t *testing.T) {
	str := "test"
	var sConfigs map[string]string
	var mConfigs module.Configs
	if err := validateServiceConfigs(sConfigs, mConfigs); err != nil {
		t.Errorf("validateServiceConfigs(%v, %v); err != nil", sConfigs, mConfigs)
	}
	sConfigs = make(map[string]string)
	if err := validateServiceConfigs(sConfigs, mConfigs); err != nil {
		t.Errorf("validateServiceConfigs(%v, %v); err != nil", sConfigs, mConfigs)
	}
	sConfigs["a"] = str
	if err := validateServiceConfigs(sConfigs, mConfigs); err == nil {
		t.Errorf("validateServiceConfigs(%v, %v); err == nil", sConfigs, mConfigs)
	}
	mConfigs = make(module.Configs)
	if err := validateServiceConfigs(sConfigs, mConfigs); err == nil {
		t.Errorf("validateServiceConfigs(%v, %v); err == nil", sConfigs, mConfigs)
	}
	mConfigs.SetString(str, nil, nil, false, "", nil, false)
	if err := validateServiceConfigs(sConfigs, mConfigs); err != nil {
		t.Errorf("validateServiceConfigs(%v, %v); err != nil", sConfigs, mConfigs)
	}
}

func TestValidateServiceHttpEndpoints(t *testing.T) {
	var sHttpEndpoints map[string]module.HttpEndpoint
	extPaths := make(map[string]struct{})
	if err := validateServiceHttpEndpoints(sHttpEndpoints, extPaths); err != nil {
		t.Errorf("validateServiceHttpEndpoints(%v); err != nil", sHttpEndpoints)
	}
	if len(extPaths) != 0 {
		t.Error("len(extPaths) != 0")
	}
	sHttpEndpoints = make(map[string]module.HttpEndpoint)
	if err := validateServiceHttpEndpoints(sHttpEndpoints, extPaths); err != nil {
		t.Errorf("validateServiceHttpEndpoints(%v); err != nil", sHttpEndpoints)
	}
	if len(extPaths) != 0 {
		t.Error("len(extPaths) != 0")
	}
	sHttpEndpoints["/test"] = module.HttpEndpoint{Path: "/test"}
	if err := validateServiceHttpEndpoints(sHttpEndpoints, extPaths); err != nil {
		t.Errorf("validateServiceHttpEndpoints(%v); err != nil", sHttpEndpoints)
	}
	if len(extPaths) != 1 {
		t.Error("len(extPaths) != 1")
	}
	if _, ok := extPaths["/test"]; !ok {
		t.Error("_, ok := extPaths[\"/test\"]; !ok")
	}
	if err := validateServiceHttpEndpoints(sHttpEndpoints, extPaths); err == nil {
		t.Errorf("validateServiceHttpEndpoints(%v); err == nil", sHttpEndpoints)
	}
	if len(extPaths) != 1 {
		t.Error("len(extPaths) != 1")
	}
	delete(sHttpEndpoints, "/test")
	sHttpEndpoints["test"] = module.HttpEndpoint{Path: "/test"}
	if err := validateServiceHttpEndpoints(sHttpEndpoints, extPaths); err == nil {
		t.Errorf("validateServiceHttpEndpoints(%v); err == nil", sHttpEndpoints)
	}
	if len(extPaths) != 1 {
		t.Error("len(extPaths) != 1")
	}
	delete(sHttpEndpoints, "test")
	sHttpEndpoints["/test"] = module.HttpEndpoint{Path: "test"}
	if err := validateServiceHttpEndpoints(sHttpEndpoints, extPaths); err == nil {
		t.Errorf("validateServiceHttpEndpoints(%v); err == nil", sHttpEndpoints)
	}
	if len(extPaths) != 1 {
		t.Error("len(extPaths) != 1")
	}
}

func TestValidateServiceDependencies(t *testing.T) {
	str := "test"
	var requiredSrv map[string]struct{}
	var requiredBySrv map[string]struct{}
	var mServices map[string]*module.Service
	if err := validateServiceDependencies(requiredSrv, requiredBySrv, mServices); err != nil {
		t.Errorf("validateServiceDependencies(%v, %v, %v); err != nil", requiredSrv, requiredBySrv, mServices)
	}
	requiredSrv = make(map[string]struct{})
	requiredBySrv = make(map[string]struct{})
	if err := validateServiceDependencies(requiredSrv, requiredBySrv, mServices); err != nil {
		t.Errorf("validateServiceDependencies(%v, %v, %v); err != nil", requiredSrv, requiredBySrv, mServices)
	}
	requiredSrv[str] = struct{}{}
	if err := validateServiceDependencies(requiredSrv, requiredBySrv, mServices); err == nil {
		t.Errorf("validateServiceDependencies(%v, %v, %v); err == nil", requiredSrv, requiredBySrv, mServices)
	}
	requiredSrv = make(map[string]struct{})
	requiredBySrv[str] = struct{}{}
	if err := validateServiceDependencies(requiredSrv, requiredBySrv, mServices); err == nil {
		t.Errorf("validateServiceDependencies(%v, %v, %v); err == nil", requiredSrv, requiredBySrv, mServices)
	}
	requiredSrv[str] = struct{}{}
	requiredBySrv = make(map[string]struct{})
	mServices = make(map[string]*module.Service)
	if err := validateServiceDependencies(requiredSrv, requiredBySrv, mServices); err == nil {
		t.Errorf("validateServiceDependencies(%v, %v, %v); err == nil", requiredSrv, requiredBySrv, mServices)
	}
	requiredSrv = make(map[string]struct{})
	requiredBySrv[str] = struct{}{}
	if err := validateServiceDependencies(requiredSrv, requiredBySrv, mServices); err == nil {
		t.Errorf("validateServiceDependencies(%v, %v, %v); err == nil", requiredSrv, requiredBySrv, mServices)
	}
	requiredSrv[str] = struct{}{}
	mServices[str] = &module.Service{}
	if err := validateServiceDependencies(requiredSrv, requiredBySrv, mServices); err != nil {
		t.Errorf("validateServiceDependencies(%v, %v, %v); err != nil", requiredSrv, requiredBySrv, mServices)
	}
}

func TestValidateServiceExternalDependencies(t *testing.T) {
	str := "test.test/test"
	var sExtDependencies map[string]module.ExtDependencyTarget
	var mDependencies map[string]string
	if err := validateServiceExternalDependencies(sExtDependencies, mDependencies); err != nil {
		t.Errorf("validateServiceExternalDependencies(%v, %v); err != nil", sExtDependencies, mDependencies)
	}
	sExtDependencies = make(map[string]module.ExtDependencyTarget)
	if err := validateServiceExternalDependencies(sExtDependencies, mDependencies); err != nil {
		t.Errorf("validateServiceExternalDependencies(%v, %v); err != nil", sExtDependencies, mDependencies)
	}
	sExtDependencies["a"] = module.ExtDependencyTarget{ID: str, Service: "test"}
	if err := validateServiceExternalDependencies(sExtDependencies, mDependencies); err == nil {
		t.Errorf("validateServiceExternalDependencies(%v, %v); err == nil", sExtDependencies, mDependencies)
	}
	mDependencies = make(map[string]string)
	if err := validateServiceExternalDependencies(sExtDependencies, mDependencies); err == nil {
		t.Errorf("validateServiceExternalDependencies(%v, %v); err == nil", sExtDependencies, mDependencies)
	}
	mDependencies[str] = "v1.0.0"
	if err := validateServiceExternalDependencies(sExtDependencies, mDependencies); err != nil {
		t.Errorf("validateServiceExternalDependencies(%v, %v); err != nil", sExtDependencies, mDependencies)
	}
	sExtDependencies["b"] = module.ExtDependencyTarget{ID: str}
	if err := validateServiceExternalDependencies(sExtDependencies, mDependencies); err == nil {
		t.Errorf("validateServiceExternalDependencies(%v, %v); err != nil", sExtDependencies, mDependencies)
	}
}

func TestValidateServiceReferences(t *testing.T) {
	str := "test"
	var sReferences map[string]string
	var mServices map[string]*module.Service
	if err := validateServiceReferences(sReferences, mServices); err != nil {
		t.Errorf("validateServiceReferences(%v, %v); err != nil", sReferences, mServices)
	}
	sReferences = make(map[string]string)
	if err := validateServiceReferences(sReferences, mServices); err != nil {
		t.Errorf("validateServiceReferences(%v, %v); err != nil", sReferences, mServices)
	}
	sReferences["a"] = str
	if err := validateServiceReferences(sReferences, mServices); err == nil {
		t.Errorf("validateServiceReferences(%v, %v); err == nil", sReferences, mServices)
	}
	mServices = make(map[string]*module.Service)
	if err := validateServiceReferences(sReferences, mServices); err == nil {
		t.Errorf("validateServiceReferences(%v, %v); err == nil", sReferences, mServices)
	}
	mServices[str] = &module.Service{}
	if err := validateServiceReferences(sReferences, mServices); err != nil {
		t.Errorf("validateServiceReferences(%v, %v); err != nil", sReferences, mServices)
	}
}

func TestValidateServicePorts(t *testing.T) {
	var sPorts []module.Port
	hostPorts := make(map[string]struct{})
	if err := validateServicePorts(sPorts, hostPorts); err != nil {
		t.Errorf("validateServicePorts(%v, %v); err != nil", sPorts, hostPorts)
	}
	sPorts = append(sPorts, module.Port{
		Number:   80,
		Protocol: module.TcpPort,
	})
	if err := validateServicePorts(sPorts, hostPorts); err != nil {
		t.Errorf("validateServicePorts(%v, %v); err != nil", sPorts, hostPorts)
	}
	if len(hostPorts) != 0 {
		t.Error("len(hostPorts) != 0")
	}
	sPorts = append(sPorts, module.Port{
		Number:   81,
		Protocol: module.TcpPort,
		Bindings: []uint{81},
	})
	if err := validateServicePorts(sPorts, hostPorts); err != nil {
		t.Errorf("validateServicePorts(%v, %v); err != nil", sPorts, hostPorts)
	}
	hostPorts = make(map[string]struct{})
	sPorts = append(sPorts, module.Port{
		Number:   82,
		Protocol: module.TcpPort,
		Bindings: []uint{82, 83},
	})
	if err := validateServicePorts(sPorts, hostPorts); err != nil {
		t.Errorf("validateServicePorts(%v, %v); err != nil", sPorts, hostPorts)
	}
	hostPorts = make(map[string]struct{})
	sPorts = append(sPorts, module.Port{
		Number:   80,
		Protocol: module.UdpPort,
	})
	if err := validateServicePorts(sPorts, hostPorts); err != nil {
		t.Errorf("validateServicePorts(%v, %v); err != nil", sPorts, hostPorts)
	}
	hostPorts = make(map[string]struct{})
	sPorts = append(sPorts, module.Port{
		Number:   81,
		Protocol: module.UdpPort,
		Bindings: []uint{81},
	})
	if err := validateServicePorts(sPorts, hostPorts); err != nil {
		t.Errorf("validateServicePorts(%v, %v); err != nil", sPorts, hostPorts)
	}
	hostPorts = make(map[string]struct{})
	sPorts = append(sPorts, module.Port{
		Number:   82,
		Protocol: module.UdpPort,
		Bindings: []uint{82, 83},
	})
	if err := validateServicePorts(sPorts, hostPorts); err != nil {
		t.Errorf("validateServicePorts(%v, %v); err != nil", sPorts, hostPorts)
	}
	if len(hostPorts) != 6 {
		t.Error("len(hostPorts) != 6")
	}
	sPorts = nil
	hostPorts = make(map[string]struct{})
	sPorts = append(sPorts, module.Port{
		Number:   80,
		Protocol: "test",
	})
	if err := validateServicePorts(sPorts, hostPorts); err == nil {
		t.Errorf("validateServicePorts(%v, %v); err == nil", sPorts, hostPorts)
	}
	if len(hostPorts) != 0 {
		t.Error("len(hostPorts) != 0")
	}
	sPorts = nil
	sPorts = append(sPorts, module.Port{
		Number:   80,
		Protocol: module.TcpPort,
	})
	sPorts = append(sPorts, module.Port{
		Number:   80,
		Protocol: module.TcpPort,
	})
	if err := validateServicePorts(sPorts, hostPorts); err == nil {
		t.Errorf("validateServicePorts(%v, %v); err == nil", sPorts, hostPorts)
	}
	if len(hostPorts) != 0 {
		t.Error("len(hostPorts) != 0")
	}
	sPorts = nil
	hostPorts = make(map[string]struct{})
	sPorts = append(sPorts, module.Port{
		Number:   81,
		Protocol: module.TcpPort,
		Bindings: []uint{81},
	})
	sPorts = append(sPorts, module.Port{
		Number:   81,
		Protocol: module.TcpPort,
		Bindings: []uint{81},
	})
	if err := validateServicePorts(sPorts, hostPorts); err == nil {
		t.Errorf("validateServicePorts(%v, %v); err == nil", sPorts, hostPorts)
	}
	if len(hostPorts) != 1 {
		t.Error("len(hostPorts) != 1")
	}
	sPorts = nil
	sPorts = append(sPorts, module.Port{
		Number:   81,
		Protocol: module.TcpPort,
		Bindings: []uint{81},
	})
	sPorts = append(sPorts, module.Port{
		Number:   82,
		Protocol: module.TcpPort,
		Bindings: []uint{81},
	})
	if err := validateServicePorts(sPorts, hostPorts); err == nil {
		t.Errorf("validateServicePorts(%v, %v); err == nil", sPorts, hostPorts)
	}
	if len(hostPorts) != 1 {
		t.Error("len(hostPorts) != 1")
	}
}
