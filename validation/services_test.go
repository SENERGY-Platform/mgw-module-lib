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
	"github.com/SENERGY-Platform/mgw-module-lib/model"
	"testing"
)

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
	var sResources map[string]model.ResourceTarget
	var mResources map[string]map[string]struct{}
	if err := validateServiceResources(sResources, mResources); err != nil {
		t.Errorf("validateServiceResources(%v, %v); err != nil", sResources, mResources)
	}
	sResources = make(map[string]model.ResourceTarget)
	if err := validateServiceResources(sResources, mResources); err != nil {
		t.Errorf("validateServiceResources(%v, %v); err != nil", sResources, mResources)
	}
	sResources["a"] = model.ResourceTarget{Ref: str}
	if err := validateServiceResources(sResources, mResources); err == nil {
		t.Errorf("validateServiceResources(%v, %v); err == nil", sResources, mResources)
	}
	mResources = make(map[string]map[string]struct{})
	if err := validateServiceResources(sResources, mResources); err == nil {
		t.Errorf("validateServiceResources(%v, %v); err == nil", sResources, mResources)
	}
	mResources[str] = map[string]struct{}{}
	if err := validateServiceResources(sResources, mResources); err != nil {
		t.Errorf("validateServiceResources(%v, %v); err != nil", sResources, mResources)
	}
}

func TestValidateServiceSecrets(t *testing.T) {
	str := "test"
	var sSecrets map[string]string
	var mSecrets map[string]model.Secret
	if err := validateServiceSecrets(sSecrets, mSecrets); err != nil {
		t.Errorf("validateServiceSecrets(%v, %v); err != nil", sSecrets, mSecrets)
	}
	sSecrets = make(map[string]string)
	if err := validateServiceSecrets(sSecrets, mSecrets); err != nil {
		t.Errorf("validateServiceSecrets(%v, %v); err != nil", sSecrets, mSecrets)
	}
	sSecrets["a"] = str
	if err := validateServiceSecrets(sSecrets, mSecrets); err == nil {
		t.Errorf("validateServiceSecrets(%v, %v); err == nil", sSecrets, mSecrets)
	}
	mSecrets = make(map[string]model.Secret)
	if err := validateServiceSecrets(sSecrets, mSecrets); err == nil {
		t.Errorf("validateServiceSecrets(%v, %v); err == nil", sSecrets, mSecrets)
	}
	mSecrets[str] = model.Secret{}
	if err := validateServiceSecrets(sSecrets, mSecrets); err != nil {
		t.Errorf("validateServiceSecrets(%v, %v); err != nil", sSecrets, mSecrets)
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
	mConfigs.SetString(str, nil, nil, false, "", nil)
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
	sHttpEndpoints["/test"] = model.HttpEndpoint{Path: "/test"}
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
	sHttpEndpoints["test"] = model.HttpEndpoint{Path: "/test"}
	if err := validateServiceHttpEndpoints(sHttpEndpoints, extPaths); err == nil {
		t.Errorf("validateServiceHttpEndpoints(%v); err == nil", sHttpEndpoints)
	}
	if len(extPaths) != 1 {
		t.Error("len(extPaths) != 1")
	}
	delete(sHttpEndpoints, "test")
	sHttpEndpoints["/test"] = model.HttpEndpoint{Path: "test"}
	if err := validateServiceHttpEndpoints(sHttpEndpoints, extPaths); err == nil {
		t.Errorf("validateServiceHttpEndpoints(%v); err == nil", sHttpEndpoints)
	}
	if len(extPaths) != 1 {
		t.Error("len(extPaths) != 1")
	}
}

func TestValidateServiceDependencies(t *testing.T) {
	str := "test"
	var sDependencies map[string]struct{}
	var mServices map[string]*model.Service
	if err := validateServiceDependencies(sDependencies, mServices); err != nil {
		t.Errorf("validateServiceDependencies(%v, %v); err != nil", sDependencies, mServices)
	}
	sDependencies = make(map[string]struct{})
	if err := validateServiceDependencies(sDependencies, mServices); err != nil {
		t.Errorf("validateServiceDependencies(%v, %v); err != nil", sDependencies, mServices)
	}
	sDependencies[str] = struct{}{}
	if err := validateServiceDependencies(sDependencies, mServices); err == nil {
		t.Errorf("validateServiceDependencies(%v, %v); err == nil", sDependencies, mServices)
	}
	mServices = make(map[string]*model.Service)
	if err := validateServiceDependencies(sDependencies, mServices); err == nil {
		t.Errorf("validateServiceDependencies(%v, %v); err == nil", sDependencies, mServices)
	}
	mServices[str] = &model.Service{}
	if err := validateServiceDependencies(sDependencies, mServices); err != nil {
		t.Errorf("validateServiceDependencies(%v, %v); err != nil", sDependencies, mServices)
	}
}

func TestValidateServiceExternalDependencies(t *testing.T) {
	str := "test.test/test"
	var sExtDependencies map[string]model.ExternalDependencyTarget
	var mDependencies map[string]string
	if err := validateServiceExternalDependencies(sExtDependencies, mDependencies); err != nil {
		t.Errorf("validateServiceExternalDependencies(%v, %v); err != nil", sExtDependencies, mDependencies)
	}
	sExtDependencies = make(map[string]model.ExternalDependencyTarget)
	if err := validateServiceExternalDependencies(sExtDependencies, mDependencies); err != nil {
		t.Errorf("validateServiceExternalDependencies(%v, %v); err != nil", sExtDependencies, mDependencies)
	}
	sExtDependencies["a"] = model.ExternalDependencyTarget{ID: str, Service: "test"}
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
	sExtDependencies["b"] = model.ExternalDependencyTarget{ID: str}
	if err := validateServiceExternalDependencies(sExtDependencies, mDependencies); err == nil {
		t.Errorf("validateServiceExternalDependencies(%v, %v); err != nil", sExtDependencies, mDependencies)
	}
}

func TestValidateServiceReferences(t *testing.T) {
	str := "test"
	var sReferences map[string]string
	var mServices map[string]*model.Service
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
	mServices = make(map[string]*model.Service)
	if err := validateServiceReferences(sReferences, mServices); err == nil {
		t.Errorf("validateServiceReferences(%v, %v); err == nil", sReferences, mServices)
	}
	mServices[str] = &model.Service{}
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
		Bindings: []uint{81},
	})
	if err := validateServicePorts(sPorts, hostPorts); err != nil {
		t.Errorf("validateServicePorts(%v, %v); err != nil", sPorts, hostPorts)
	}
	hostPorts = make(map[string]struct{})
	sPorts = append(sPorts, model.Port{
		Number:   82,
		Protocol: model.TcpPort,
		Bindings: []uint{82, 83},
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
		Bindings: []uint{81},
	})
	if err := validateServicePorts(sPorts, hostPorts); err != nil {
		t.Errorf("validateServicePorts(%v, %v); err != nil", sPorts, hostPorts)
	}
	hostPorts = make(map[string]struct{})
	sPorts = append(sPorts, model.Port{
		Number:   82,
		Protocol: model.UdpPort,
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
		Bindings: []uint{81},
	})
	sPorts = append(sPorts, model.Port{
		Number:   81,
		Protocol: model.TcpPort,
		Bindings: []uint{81},
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
		Bindings: []uint{81},
	})
	sPorts = append(sPorts, model.Port{
		Number:   82,
		Protocol: model.TcpPort,
		Bindings: []uint{81},
	})
	if err := validateServicePorts(sPorts, hostPorts); err == nil {
		t.Errorf("validateServicePorts(%v, %v); err == nil", sPorts, hostPorts)
	}
	if len(hostPorts) != 1 {
		t.Error("len(hostPorts) != 1")
	}
}
