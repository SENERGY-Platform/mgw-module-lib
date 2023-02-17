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
	var mVolumes model.Set[string]
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
	mVolumes = make(model.Set[string])
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
	var mResources map[string]model.Set[string]
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
	mResources = make(map[string]model.Set[string])
	if err := validateServiceResources(sResources, mResources); err == nil {
		t.Errorf("validateServiceResources(%v, %v); err == nil", sResources, mResources)
	}
	mResources[str] = model.Set[string]{}
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
