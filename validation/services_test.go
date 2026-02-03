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
	"testing"

	"github.com/SENERGY-Platform/mgw-module-lib/model"
)

func TestValidateServices(t *testing.T) {
	s := map[string]model.Service{
		"a": {},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil, nil, nil); err != nil {
		t.Error("err != nil")
	}
	// ------------------------------
	s = map[string]model.Service{
		"": {},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.Service{
		"a": {
			BindMounts: map[string]model.BindMount{"": {}},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.Service{
		"a": {
			Tmpfs: map[string]model.TmpfsMount{"": {}},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.Service{
		"a": {
			Volumes: map[string]string{"": ""},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.Service{
		"a": {
			HostResources: map[string]model.HostResTarget{"": {}},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.Service{
		"a": {
			SecretMounts: map[string]model.SecretTarget{"": {}},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.Service{
		"a": {
			SecretVars: map[string]model.SecretTarget{"": {}},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.Service{
		"a": {
			Configs: map[string]string{"": ""},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.Service{
		"a": {
			SrvReferences: map[string]model.SrvRefTarget{"": {}},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.Service{
		"a": {
			ExtDependencies: map[string]model.ExtDependencyTarget{"": {}},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.Service{
		"a": {
			Volumes: map[string]string{"test": ""},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.Service{
		"a": {
			HostResources: map[string]model.HostResTarget{"test": {}},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.Service{
		"a": {
			SecretMounts: map[string]model.SecretTarget{"test": {}},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.Service{
		"a": {
			SecretVars: map[string]model.SecretTarget{"test": {}},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.Service{
		"a": {
			Configs: map[string]string{"test": ""},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.Service{
		"a": {
			HttpEndpoints: map[string]model.HttpEndpoint{"///": {}},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.Service{
		"a": {
			SrvReferences: map[string]model.SrvRefTarget{"test": {}},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.Service{
		"a": {
			RequiredSrv: map[string]struct{}{"": {}},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.Service{
		"a": {
			ExtDependencies: map[string]model.ExtDependencyTarget{"test": {}},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.Service{
		"a": {
			Ports: []model.Port{
				{},
			},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.Service{
		"a": {
			RequiredSrv: map[string]struct{}{
				"a": {},
			},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.Service{
		"a": {
			Files: map[string]model.FileTarget{"test": {}},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.Service{
		"a": {
			FileGroups: map[string]string{"test": ""},
		},
	}
	if err := validateServices(s, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
}

func TestValidateAuxServices(t *testing.T) {
	s := map[string]model.AuxService{
		"a": {},
	}
	if err := validateAuxServices(s, nil, nil, nil, nil); err != nil {
		t.Error("err != nil")
	}
	// ------------------------------
	s = map[string]model.AuxService{
		"": {},
	}
	if err := validateAuxServices(s, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.AuxService{
		"a": {
			BindMounts: map[string]model.BindMount{"": {}},
		},
	}
	if err := validateAuxServices(s, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.AuxService{
		"a": {
			Tmpfs: map[string]model.TmpfsMount{"": {}},
		},
	}
	if err := validateAuxServices(s, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.AuxService{
		"a": {
			Volumes: map[string]string{"": ""},
		},
	}
	if err := validateAuxServices(s, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.AuxService{
		"a": {
			Configs: map[string]string{"": ""},
		},
	}
	if err := validateAuxServices(s, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.AuxService{
		"a": {
			SrvReferences: map[string]model.SrvRefTarget{"": {}},
		},
	}
	if err := validateAuxServices(s, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.AuxService{
		"a": {
			ExtDependencies: map[string]model.ExtDependencyTarget{"": {}},
		},
	}
	if err := validateAuxServices(s, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.AuxService{
		"a": {
			Volumes: map[string]string{"test": ""},
		},
	}
	if err := validateAuxServices(s, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.AuxService{
		"a": {
			Configs: map[string]string{"test": ""},
		},
	}
	if err := validateAuxServices(s, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.AuxService{
		"a": {
			SrvReferences: map[string]model.SrvRefTarget{"test": {}},
		},
	}
	if err := validateAuxServices(s, nil, nil, nil, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	s = map[string]model.AuxService{
		"a": {
			ExtDependencies: map[string]model.ExtDependencyTarget{"test": {}},
		},
	}
	if err := validateAuxServices(s, nil, nil, nil, nil); err == nil {
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
	var sResources map[string]model.HostResTarget
	var mResources map[string]model.HostResource
	if err := validateServiceResources(sResources, mResources); err != nil {
		t.Errorf("validateServiceResources(%v, %v); err != nil", sResources, mResources)
	}
	sResources = make(map[string]model.HostResTarget)
	if err := validateServiceResources(sResources, mResources); err != nil {
		t.Errorf("validateServiceResources(%v, %v); err != nil", sResources, mResources)
	}
	sResources["a"] = model.HostResTarget{Ref: str}
	if err := validateServiceResources(sResources, mResources); err == nil {
		t.Errorf("validateServiceResources(%v, %v); err == nil", sResources, mResources)
	}
	mResources = make(map[string]model.HostResource)
	if err := validateServiceResources(sResources, mResources); err == nil {
		t.Errorf("validateServiceResources(%v, %v); err == nil", sResources, mResources)
	}
	mResources[str] = model.HostResource{}
	if err := validateServiceResources(sResources, mResources); err != nil {
		t.Errorf("validateServiceResources(%v, %v); err != nil", sResources, mResources)
	}
}

func TestValidateServiceSecrets(t *testing.T) {
	str := "test"
	var sSecretMounts map[string]model.SecretTarget
	var sSecretVars map[string]model.SecretTarget
	var mSecrets map[string]model.Secret
	if err := validateServiceSecrets(sSecretMounts, sSecretVars, mSecrets); err != nil {
		t.Errorf("validateServiceSecrets(%v, %v); err != nil", sSecretMounts, mSecrets)
	}
	sSecretMounts = make(map[string]model.SecretTarget)
	sSecretVars = make(map[string]model.SecretTarget)
	if err := validateServiceSecrets(sSecretMounts, sSecretVars, mSecrets); err != nil {
		t.Errorf("validateServiceSecrets(%v, %v); err != nil", sSecretMounts, mSecrets)
	}
	sSecretMounts["a"] = model.SecretTarget{Ref: str}
	if err := validateServiceSecrets(sSecretMounts, sSecretVars, mSecrets); err == nil {
		t.Errorf("validateServiceSecrets(%v, %v); err == nil", sSecretMounts, mSecrets)
	}
	mSecrets = make(map[string]model.Secret)
	if err := validateServiceSecrets(sSecretMounts, sSecretVars, mSecrets); err == nil {
		t.Errorf("validateServiceSecrets(%v, %v); err == nil", sSecretMounts, mSecrets)
	}
	mSecrets[str] = model.Secret{}
	if err := validateServiceSecrets(sSecretMounts, sSecretVars, mSecrets); err != nil {
		t.Errorf("validateServiceSecrets(%v, %v); err != nil", sSecretMounts, mSecrets)
	}
	sSecretVars["a"] = model.SecretTarget{Ref: str}
	if err := validateServiceSecrets(sSecretMounts, sSecretVars, mSecrets); err != nil {
		t.Errorf("validateServiceSecrets(%v, %v); err != nil", sSecretMounts, mSecrets)
	}
}

func TestValidateServiceConfigs(t *testing.T) {
	str := "test"
	var sConfigs map[string]string
	var mConfigs model.Configs
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
	mConfigs = make(model.Configs)
	if err := validateServiceConfigs(sConfigs, mConfigs); err == nil {
		t.Errorf("validateServiceConfigs(%v, %v); err == nil", sConfigs, mConfigs)
	}
	mConfigs.SetString(str, nil, nil, false, "", nil, false)
	if err := validateServiceConfigs(sConfigs, mConfigs); err != nil {
		t.Errorf("validateServiceConfigs(%v, %v); err != nil", sConfigs, mConfigs)
	}
}

func TestValidateServiceHttpEndpoints(t *testing.T) {
	var sHttpEndpoints map[string]model.HttpEndpoint
	extPaths := make(map[string]struct{})
	if err := validateServiceHttpEndpoints(sHttpEndpoints, extPaths); err != nil {
		t.Errorf("validateServiceHttpEndpoints(%v); err != nil", sHttpEndpoints)
	}
	if len(extPaths) != 0 {
		t.Error("len(extPaths) != 0")
	}
	sHttpEndpoints = make(map[string]model.HttpEndpoint)
	if err := validateServiceHttpEndpoints(sHttpEndpoints, extPaths); err != nil {
		t.Errorf("validateServiceHttpEndpoints(%v); err != nil", sHttpEndpoints)
	}
	if len(extPaths) != 0 {
		t.Error("len(extPaths) != 0")
	}
	p1 := "/test"
	sHttpEndpoints["test"] = model.HttpEndpoint{Path: p1}
	if err := validateServiceHttpEndpoints(sHttpEndpoints, extPaths); err != nil {
		t.Errorf("validateServiceHttpEndpoints(%v); err != nil", sHttpEndpoints)
	}
	if len(extPaths) != 1 {
		t.Error("len(extPaths) != 1")
	}
	if _, ok := extPaths["test"]; !ok {
		t.Error("_, ok := extPaths[\"test\"]; !ok")
	}
	if err := validateServiceHttpEndpoints(sHttpEndpoints, extPaths); err == nil {
		t.Errorf("validateServiceHttpEndpoints(%v); err == nil", sHttpEndpoints)
	}
	if len(extPaths) != 1 {
		t.Error("len(extPaths) != 1")
	}
	delete(sHttpEndpoints, "test")
	sHttpEndpoints["/test"] = model.HttpEndpoint{Path: p1}
	if err := validateServiceHttpEndpoints(sHttpEndpoints, extPaths); err == nil {
		t.Errorf("validateServiceHttpEndpoints(%v); err == nil", sHttpEndpoints)
	}
	if len(extPaths) != 1 {
		t.Error("len(extPaths) != 1")
	}
	delete(sHttpEndpoints, "test")
	p2 := "test"
	sHttpEndpoints["/test"] = model.HttpEndpoint{Path: p2}
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
	var mServices map[string]model.Service
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
	mServices = make(map[string]model.Service)
	if err := validateServiceDependencies(requiredSrv, requiredBySrv, mServices); err == nil {
		t.Errorf("validateServiceDependencies(%v, %v, %v); err == nil", requiredSrv, requiredBySrv, mServices)
	}
	requiredSrv = make(map[string]struct{})
	requiredBySrv[str] = struct{}{}
	if err := validateServiceDependencies(requiredSrv, requiredBySrv, mServices); err == nil {
		t.Errorf("validateServiceDependencies(%v, %v, %v); err == nil", requiredSrv, requiredBySrv, mServices)
	}
	requiredSrv[str] = struct{}{}
	mServices[str] = model.Service{}
	if err := validateServiceDependencies(requiredSrv, requiredBySrv, mServices); err != nil {
		t.Errorf("validateServiceDependencies(%v, %v, %v); err != nil", requiredSrv, requiredBySrv, mServices)
	}
}

func TestValidateServiceExternalDependencies(t *testing.T) {
	str := "test.test/test"
	var sExtDependencies map[string]model.ExtDependencyTarget
	var mDependencies map[string]string
	if err := validateServiceExternalDependencies(sExtDependencies, mDependencies); err != nil {
		t.Errorf("validateServiceExternalDependencies(%v, %v); err != nil", sExtDependencies, mDependencies)
	}
	sExtDependencies = make(map[string]model.ExtDependencyTarget)
	if err := validateServiceExternalDependencies(sExtDependencies, mDependencies); err != nil {
		t.Errorf("validateServiceExternalDependencies(%v, %v); err != nil", sExtDependencies, mDependencies)
	}
	sExtDependencies["a"] = model.ExtDependencyTarget{ID: str, Service: "test"}
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
	sExtDependencies["b"] = model.ExtDependencyTarget{ID: str}
	if err := validateServiceExternalDependencies(sExtDependencies, mDependencies); err == nil {
		t.Errorf("validateServiceExternalDependencies(%v, %v); err != nil", sExtDependencies, mDependencies)
	}
}

func TestValidateServiceReferences(t *testing.T) {
	str := "test"
	var sReferences map[string]model.SrvRefTarget
	var mServices map[string]model.Service
	if err := validateServiceReferences(sReferences, mServices); err != nil {
		t.Errorf("validateServiceReferences(%v, %v); err != nil", sReferences, mServices)
	}
	sReferences = make(map[string]model.SrvRefTarget)
	if err := validateServiceReferences(sReferences, mServices); err != nil {
		t.Errorf("validateServiceReferences(%v, %v); err != nil", sReferences, mServices)
	}
	sReferences["a"] = model.SrvRefTarget{Ref: str}
	if err := validateServiceReferences(sReferences, mServices); err == nil {
		t.Errorf("validateServiceReferences(%v, %v); err == nil", sReferences, mServices)
	}
	mServices = make(map[string]model.Service)
	if err := validateServiceReferences(sReferences, mServices); err == nil {
		t.Errorf("validateServiceReferences(%v, %v); err == nil", sReferences, mServices)
	}
	mServices[str] = model.Service{}
	if err := validateServiceReferences(sReferences, mServices); err != nil {
		t.Errorf("validateServiceReferences(%v, %v); err != nil", sReferences, mServices)
	}
}

func TestValidateServicePorts(t *testing.T) {
	var sPorts []model.Port
	hostPorts := make(map[string]struct{})
	if err := validateServicePorts(sPorts, hostPorts); err != nil {
		t.Errorf("validateServicePorts(%v, %v); err != nil", sPorts, hostPorts)
	}
	sPorts = append(sPorts, model.Port{
		Number:   80,
		Protocol: model.TcpPort,
	})
	if err := validateServicePorts(sPorts, hostPorts); err != nil {
		t.Errorf("validateServicePorts(%v, %v); err != nil", sPorts, hostPorts)
	}
	if len(hostPorts) != 0 {
		t.Error("len(hostPorts) != 0")
	}
	sPorts = append(sPorts, model.Port{
		Number:   81,
		Protocol: model.TcpPort,
		Bindings: []int{81},
	})
	if err := validateServicePorts(sPorts, hostPorts); err != nil {
		t.Errorf("validateServicePorts(%v, %v); err != nil", sPorts, hostPorts)
	}
	hostPorts = make(map[string]struct{})
	sPorts = append(sPorts, model.Port{
		Number:   82,
		Protocol: model.TcpPort,
		Bindings: []int{82, 83},
	})
	if err := validateServicePorts(sPorts, hostPorts); err != nil {
		t.Errorf("validateServicePorts(%v, %v); err != nil", sPorts, hostPorts)
	}
	hostPorts = make(map[string]struct{})
	sPorts = append(sPorts, model.Port{
		Number:   80,
		Protocol: model.UdpPort,
	})
	if err := validateServicePorts(sPorts, hostPorts); err != nil {
		t.Errorf("validateServicePorts(%v, %v); err != nil", sPorts, hostPorts)
	}
	hostPorts = make(map[string]struct{})
	sPorts = append(sPorts, model.Port{
		Number:   81,
		Protocol: model.UdpPort,
		Bindings: []int{81},
	})
	if err := validateServicePorts(sPorts, hostPorts); err != nil {
		t.Errorf("validateServicePorts(%v, %v); err != nil", sPorts, hostPorts)
	}
	hostPorts = make(map[string]struct{})
	sPorts = append(sPorts, model.Port{
		Number:   82,
		Protocol: model.UdpPort,
		Bindings: []int{82, 83},
	})
	if err := validateServicePorts(sPorts, hostPorts); err != nil {
		t.Errorf("validateServicePorts(%v, %v); err != nil", sPorts, hostPorts)
	}
	if len(hostPorts) != 6 {
		t.Error("len(hostPorts) != 6")
	}
	sPorts = nil
	hostPorts = make(map[string]struct{})
	sPorts = append(sPorts, model.Port{
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
	sPorts = append(sPorts, model.Port{
		Number:   80,
		Protocol: model.TcpPort,
	})
	sPorts = append(sPorts, model.Port{
		Number:   80,
		Protocol: model.TcpPort,
	})
	if err := validateServicePorts(sPorts, hostPorts); err == nil {
		t.Errorf("validateServicePorts(%v, %v); err == nil", sPorts, hostPorts)
	}
	if len(hostPorts) != 0 {
		t.Error("len(hostPorts) != 0")
	}
	sPorts = nil
	hostPorts = make(map[string]struct{})
	sPorts = append(sPorts, model.Port{
		Number:   81,
		Protocol: model.TcpPort,
		Bindings: []int{81},
	})
	sPorts = append(sPorts, model.Port{
		Number:   81,
		Protocol: model.TcpPort,
		Bindings: []int{81},
	})
	if err := validateServicePorts(sPorts, hostPorts); err == nil {
		t.Errorf("validateServicePorts(%v, %v); err == nil", sPorts, hostPorts)
	}
	if len(hostPorts) != 1 {
		t.Error("len(hostPorts) != 1")
	}
	sPorts = nil
	sPorts = append(sPorts, model.Port{
		Number:   81,
		Protocol: model.TcpPort,
		Bindings: []int{81},
	})
	sPorts = append(sPorts, model.Port{
		Number:   82,
		Protocol: model.TcpPort,
		Bindings: []int{81},
	})
	if err := validateServicePorts(sPorts, hostPorts); err == nil {
		t.Errorf("validateServicePorts(%v, %v); err == nil", sPorts, hostPorts)
	}
	if len(hostPorts) != 1 {
		t.Error("len(hostPorts) != 1")
	}
}
