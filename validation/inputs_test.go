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

func TestValidateInputsResources(t *testing.T) {
	str := "test"
	var inputs map[string]module.Input
	var mResources map[string]module.Resource
	if err := validateInputsResources(inputs, mResources); err != nil {
		t.Errorf("validateInputsResources(%v, %v); err != nil", inputs, mResources)
	}
	inputs = make(map[string]module.Input)
	if err := validateInputsResources(inputs, mResources); err != nil {
		t.Errorf("validateInputsResources(%v, %v); err != nil", inputs, mResources)
	}
	inputs[str] = module.Input{}
	if err := validateInputsResources(inputs, mResources); err == nil {
		t.Errorf("validateInputsResources(%v, %v); err == nil", inputs, mResources)
	}
	mResources = make(map[string]module.Resource)
	if err := validateInputsResources(inputs, mResources); err == nil {
		t.Errorf("validateInputsResources(%v, %v); err == nil", inputs, mResources)
	}
	mResources[str] = module.Resource{}
	if err := validateInputsResources(inputs, mResources); err != nil {
		t.Errorf("validateInputsResources(%v, %v); err != nil", inputs, mResources)
	}
	inputs[""] = module.Input{}
	if err := validateInputsResources(inputs, mResources); err == nil {
		t.Errorf("validateInputsResources(%v, %v); err == nil", inputs, mResources)
	}
}

func TestValidateInputsSecrets(t *testing.T) {
	str := "test"
	var inputs map[string]module.Input
	var mSecrets map[string]module.Secret
	if err := validateInputsSecrets(inputs, mSecrets); err != nil {
		t.Errorf("validateInputsSecrets(%v, %v); err != nil", inputs, mSecrets)
	}
	inputs = make(map[string]module.Input)
	if err := validateInputsSecrets(inputs, mSecrets); err != nil {
		t.Errorf("validateInputsSecrets(%v, %v); err != nil", inputs, mSecrets)
	}
	inputs[str] = module.Input{}
	if err := validateInputsSecrets(inputs, mSecrets); err == nil {
		t.Errorf("validateInputsSecrets(%v, %v); err == nil", inputs, mSecrets)
	}
	mSecrets = make(map[string]module.Secret)
	if err := validateInputsSecrets(inputs, mSecrets); err == nil {
		t.Errorf("validateInputsSecrets(%v, %v); err == nil", inputs, mSecrets)
	}
	mSecrets[str] = module.Secret{}
	if err := validateInputsSecrets(inputs, mSecrets); err != nil {
		t.Errorf("validateInputsSecrets(%v, %v); err != nil", inputs, mSecrets)
	}
	inputs[""] = module.Input{}
	if err := validateInputsSecrets(inputs, mSecrets); err == nil {
		t.Errorf("validateInputsSecrets(%v, %v); err == nil", inputs, mSecrets)
	}
}

func TestValidateInputsConfigs(t *testing.T) {
	str := "test"
	var inputs map[string]module.Input
	var mConfigs module.Configs
	if err := validateInputsConfigs(inputs, mConfigs); err != nil {
		t.Errorf("validateInputsConfigs(%v, %v); err != nil", inputs, mConfigs)
	}
	inputs = make(map[string]module.Input)
	if err := validateInputsConfigs(inputs, mConfigs); err != nil {
		t.Errorf("validateInputsConfigs(%v, %v); err != nil", inputs, mConfigs)
	}
	inputs[str] = module.Input{}
	if err := validateInputsConfigs(inputs, mConfigs); err == nil {
		t.Errorf("validateInputsConfigs(%v, %v); err == nil", inputs, mConfigs)
	}
	mConfigs = make(module.Configs)
	if err := validateInputsConfigs(inputs, mConfigs); err == nil {
		t.Errorf("validateInputsConfigs(%v, %v); err == nil", inputs, mConfigs)
	}
	mConfigs.SetString(str, nil, nil, false, "", nil, false)
	if err := validateInputsConfigs(inputs, mConfigs); err != nil {
		t.Errorf("validateInputsConfigs(%v, %v); err != nil", inputs, mConfigs)
	}
	inputs[""] = module.Input{}
	if err := validateInputsConfigs(inputs, mConfigs); err == nil {
		t.Errorf("validateInputsConfigs(%v, %v); err == nil", inputs, mConfigs)
	}
}

func TestValidateInputsAndGroups(t *testing.T) {
	str := "test"
	var inputs map[string]module.Input
	var groups map[string]module.InputGroup
	if err := validateInputsAndGroups(inputs, groups); err != nil {
		t.Errorf("validateInputsAndGroups(%v, %v); err != nil", inputs, groups)
	}
	inputs = make(map[string]module.Input)
	if err := validateInputsAndGroups(inputs, groups); err != nil {
		t.Errorf("validateInputsAndGroups(%v, %v); err != nil", inputs, groups)
	}
	inputs["a"] = module.Input{}
	if err := validateInputsAndGroups(inputs, groups); err != nil {
		t.Errorf("validateInputsAndGroups(%v, %v); err != nil", inputs, groups)
	}
	inputs["b"] = module.Input{Group: &str}
	if err := validateInputsAndGroups(inputs, groups); err == nil {
		t.Errorf("validateInputsAndGroups(%v, %v); err == nil", inputs, groups)
	}
	groups = make(map[string]module.InputGroup)
	if err := validateInputsAndGroups(inputs, groups); err == nil {
		t.Errorf("validateInputsAndGroups(%v, %v); err == nil", inputs, groups)
	}
	groups[str] = module.InputGroup{}
	if err := validateInputsAndGroups(inputs, groups); err != nil {
		t.Errorf("validateInputsAndGroups(%v, %v); err != nil", inputs, groups)
	}
}

func TestValidateInputGroups(t *testing.T) {
	str1 := "a"
	str2 := "test"
	var groups map[string]module.InputGroup
	if err := validateInputGroups(groups); err != nil {
		t.Errorf("validateInputGroups(%v); err != nil", groups)
	}
	groups = make(map[string]module.InputGroup)
	if err := validateInputGroups(groups); err != nil {
		t.Errorf("validateInputGroups(%v); err != nil", groups)
	}
	groups[str1] = module.InputGroup{}
	if err := validateInputGroups(groups); err != nil {
		t.Errorf("validateInputGroups(%v); err != nil", groups)
	}
	groups["b"] = module.InputGroup{Group: &str1}
	if err := validateInputGroups(groups); err != nil {
		t.Errorf("validateInputGroups(%v); err != nil", groups)
	}
	groups["c"] = module.InputGroup{Group: &str2}
	if err := validateInputGroups(groups); err == nil {
		t.Errorf("validateInputGroups(%v); err == nil", groups)
	}
	delete(groups, "c")
	groups[str2] = module.InputGroup{Group: &str2}
	if err := validateInputGroups(groups); err == nil {
		t.Errorf("validateInputGroups(%v); err == nil", groups)
	}
	delete(groups, str2)
	groups[""] = module.InputGroup{}
	if err := validateInputGroups(groups); err == nil {
		t.Errorf("validateInputGroups(%v); err == nil", groups)
	}
}
