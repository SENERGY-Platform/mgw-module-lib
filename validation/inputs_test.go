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

func TestValidateInputsResources(t *testing.T) {
	str := "test"
	var inputs map[string]model.Input
	var mResources map[string]map[string]struct{}
	if err := validateInputsResources(inputs, mResources); err != nil {
		t.Errorf("validateInputsResources(%v, %v); err != nil", inputs, mResources)
	}
	inputs = make(map[string]model.Input)
	if err := validateInputsResources(inputs, mResources); err != nil {
		t.Errorf("validateInputsResources(%v, %v); err != nil", inputs, mResources)
	}
	inputs[str] = model.Input{}
	if err := validateInputsResources(inputs, mResources); err == nil {
		t.Errorf("validateInputsResources(%v, %v); err == nil", inputs, mResources)
	}
	mResources = make(map[string]map[string]struct{})
	if err := validateInputsResources(inputs, mResources); err == nil {
		t.Errorf("validateInputsResources(%v, %v); err == nil", inputs, mResources)
	}
	mResources[str] = map[string]struct{}{}
	if err := validateInputsResources(inputs, mResources); err != nil {
		t.Errorf("validateInputsResources(%v, %v); err != nil", inputs, mResources)
	}
	inputs[""] = model.Input{}
	if err := validateInputsResources(inputs, mResources); err == nil {
		t.Errorf("validateInputsResources(%v, %v); err == nil", inputs, mResources)
	}
}

func TestValidateInputsSecrets(t *testing.T) {
	str := "test"
	var inputs map[string]model.Input
	var mSecrets map[string]model.Secret
	if err := validateInputsSecrets(inputs, mSecrets); err != nil {
		t.Errorf("validateInputsSecrets(%v, %v); err != nil", inputs, mSecrets)
	}
	inputs = make(map[string]model.Input)
	if err := validateInputsSecrets(inputs, mSecrets); err != nil {
		t.Errorf("validateInputsSecrets(%v, %v); err != nil", inputs, mSecrets)
	}
	inputs[str] = model.Input{}
	if err := validateInputsSecrets(inputs, mSecrets); err == nil {
		t.Errorf("validateInputsSecrets(%v, %v); err == nil", inputs, mSecrets)
	}
	mSecrets = make(map[string]model.Secret)
	if err := validateInputsSecrets(inputs, mSecrets); err == nil {
		t.Errorf("validateInputsSecrets(%v, %v); err == nil", inputs, mSecrets)
	}
	mSecrets[str] = model.Secret{}
	if err := validateInputsSecrets(inputs, mSecrets); err != nil {
		t.Errorf("validateInputsSecrets(%v, %v); err != nil", inputs, mSecrets)
	}
	inputs[""] = model.Input{}
	if err := validateInputsSecrets(inputs, mSecrets); err == nil {
		t.Errorf("validateInputsSecrets(%v, %v); err == nil", inputs, mSecrets)
	}
}

func TestValidateInputsConfigs(t *testing.T) {
	str := "test"
	var inputs map[string]model.Input
	var mConfigs model.Configs
	if err := validateInputsConfigs(inputs, mConfigs); err != nil {
		t.Errorf("validateInputsConfigs(%v, %v); err != nil", inputs, mConfigs)
	}
	inputs = make(map[string]model.Input)
	if err := validateInputsConfigs(inputs, mConfigs); err != nil {
		t.Errorf("validateInputsConfigs(%v, %v); err != nil", inputs, mConfigs)
	}
	inputs[str] = model.Input{}
	if err := validateInputsConfigs(inputs, mConfigs); err == nil {
		t.Errorf("validateInputsConfigs(%v, %v); err == nil", inputs, mConfigs)
	}
	mConfigs = make(model.Configs)
	if err := validateInputsConfigs(inputs, mConfigs); err == nil {
		t.Errorf("validateInputsConfigs(%v, %v); err == nil", inputs, mConfigs)
	}
	mConfigs.SetString(str, nil, nil, false, "", nil)
	if err := validateInputsConfigs(inputs, mConfigs); err != nil {
		t.Errorf("validateInputsConfigs(%v, %v); err != nil", inputs, mConfigs)
	}
	inputs[""] = model.Input{}
	if err := validateInputsConfigs(inputs, mConfigs); err == nil {
		t.Errorf("validateInputsConfigs(%v, %v); err == nil", inputs, mConfigs)
	}
}

func TestValidateInputsAndGroups(t *testing.T) {
	str := "test"
	var inputs map[string]model.Input
	var groups map[string]model.InputGroup
	if err := validateInputsAndGroups(inputs, groups); err != nil {
		t.Errorf("validateInputsAndGroups(%v, %v); err != nil", inputs, groups)
	}
	inputs = make(map[string]model.Input)
	if err := validateInputsAndGroups(inputs, groups); err != nil {
		t.Errorf("validateInputsAndGroups(%v, %v); err != nil", inputs, groups)
	}
	inputs["a"] = model.Input{}
	if err := validateInputsAndGroups(inputs, groups); err != nil {
		t.Errorf("validateInputsAndGroups(%v, %v); err != nil", inputs, groups)
	}
	inputs["b"] = model.Input{Group: &str}
	if err := validateInputsAndGroups(inputs, groups); err == nil {
		t.Errorf("validateInputsAndGroups(%v, %v); err == nil", inputs, groups)
	}
	groups = make(map[string]model.InputGroup)
	if err := validateInputsAndGroups(inputs, groups); err == nil {
		t.Errorf("validateInputsAndGroups(%v, %v); err == nil", inputs, groups)
	}
	groups[str] = model.InputGroup{}
	if err := validateInputsAndGroups(inputs, groups); err != nil {
		t.Errorf("validateInputsAndGroups(%v, %v); err != nil", inputs, groups)
	}
}

func TestValidateInputGroups(t *testing.T) {
	str1 := "a"
	str2 := "test"
	var groups map[string]model.InputGroup
	if err := validateInputGroups(groups); err != nil {
		t.Errorf("validateInputGroups(%v); err != nil", groups)
	}
	groups = make(map[string]model.InputGroup)
	if err := validateInputGroups(groups); err != nil {
		t.Errorf("validateInputGroups(%v); err != nil", groups)
	}
	groups[str1] = model.InputGroup{}
	if err := validateInputGroups(groups); err != nil {
		t.Errorf("validateInputGroups(%v); err != nil", groups)
	}
	groups["b"] = model.InputGroup{Group: &str1}
	if err := validateInputGroups(groups); err != nil {
		t.Errorf("validateInputGroups(%v); err != nil", groups)
	}
	groups["c"] = model.InputGroup{Group: &str2}
	if err := validateInputGroups(groups); err == nil {
		t.Errorf("validateInputGroups(%v); err == nil", groups)
	}
	delete(groups, "c")
	groups[str2] = model.InputGroup{Group: &str2}
	if err := validateInputGroups(groups); err == nil {
		t.Errorf("validateInputGroups(%v); err == nil", groups)
	}
	delete(groups, str2)
	groups[""] = model.InputGroup{}
	if err := validateInputGroups(groups); err == nil {
		t.Errorf("validateInputGroups(%v); err == nil", groups)
	}
}
