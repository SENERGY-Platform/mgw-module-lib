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
	var mResources map[string]model.Set[string]
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
	mResources = make(map[string]model.Set[string])
	if err := validateInputsResources(inputs, mResources); err == nil {
		t.Errorf("validateInputsResources(%v, %v); err == nil", inputs, mResources)
	}
	mResources[str] = model.Set[string]{}
	if err := validateInputsResources(inputs, mResources); err != nil {
		t.Errorf("validateInputsResources(%v, %v); err != nil", inputs, mResources)
	}
}

func TestValidateInputsSecrets(t *testing.T) {
	str := "test"
	var inputs map[string]model.Input
	var mSecrets map[string]model.Secret
	if err := validateInputsSecrets(inputs, mSecrets); err != nil {
		t.Errorf("validateInputsResources(%v, %v); err != nil", inputs, mSecrets)
	}
	inputs = make(map[string]model.Input)
	if err := validateInputsSecrets(inputs, mSecrets); err != nil {
		t.Errorf("validateInputsResources(%v, %v); err != nil", inputs, mSecrets)
	}
	inputs[str] = model.Input{}
	if err := validateInputsSecrets(inputs, mSecrets); err == nil {
		t.Errorf("validateInputsResources(%v, %v); err == nil", inputs, mSecrets)
	}
	mSecrets = make(map[string]model.Secret)
	if err := validateInputsSecrets(inputs, mSecrets); err == nil {
		t.Errorf("validateInputsResources(%v, %v); err == nil", inputs, mSecrets)
	}
	mSecrets[str] = model.Secret{}
	if err := validateInputsSecrets(inputs, mSecrets); err != nil {
		t.Errorf("validateInputsResources(%v, %v); err != nil", inputs, mSecrets)
	}
}
