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
