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

func TestValidate(t *testing.T) {
	m := module.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           module.AddOnModule,
		DeploymentType: module.SingleDeployment,
	}
	if err := Validate(&m); err != nil {
		t.Error("err != nil")
	}
	// ------------------------------
	m = module.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           module.AddOnModule,
		DeploymentType: module.SingleDeployment,
		Services: map[string]*module.Service{
			"a": {RequiredSrv: map[string]struct{}{"a": {}}},
		},
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	m = module.Module{
		ID: "",
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	m = module.Module{
		ID:      "test.test/test",
		Version: "",
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	m = module.Module{
		ID:      "test.test/test",
		Version: "v1.0.0",
		Type:    "",
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	m = module.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           module.AddOnModule,
		DeploymentType: "",
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	m = module.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           module.AddOnModule,
		DeploymentType: module.SingleDeployment,
		Volumes:        map[string]struct{}{"": {}},
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	m = module.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           module.AddOnModule,
		DeploymentType: module.SingleDeployment,
		Dependencies:   map[string]string{"": ""},
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	m = module.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           module.AddOnModule,
		DeploymentType: module.SingleDeployment,
		Resources:      map[string]module.Resource{"": {}},
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	m = module.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           module.AddOnModule,
		DeploymentType: module.SingleDeployment,
		Secrets:        map[string]module.Secret{"": {}},
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	m = module.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           module.AddOnModule,
		DeploymentType: module.SingleDeployment,
		Configs:        module.Configs{"": {}},
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	m = module.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           module.AddOnModule,
		DeploymentType: module.SingleDeployment,
		Inputs: module.Inputs{
			Groups: map[string]module.InputGroup{"": {}},
		},
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	g := ""
	m = module.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           module.AddOnModule,
		DeploymentType: module.SingleDeployment,
		Inputs: module.Inputs{
			Resources: map[string]module.Input{"test": {Group: &g}},
		},
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	m = module.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           module.AddOnModule,
		DeploymentType: module.SingleDeployment,
		Inputs: module.Inputs{
			Secrets: map[string]module.Input{"test": {Group: &g}},
		},
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	m = module.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           module.AddOnModule,
		DeploymentType: module.SingleDeployment,
		Inputs: module.Inputs{
			Configs: map[string]module.Input{"test": {Group: &g}},
		},
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	m = module.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           module.AddOnModule,
		DeploymentType: module.SingleDeployment,
		Inputs: module.Inputs{
			Resources: map[string]module.Input{"": {}},
		},
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	m = module.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           module.AddOnModule,
		DeploymentType: module.SingleDeployment,
		Inputs: module.Inputs{
			Secrets: map[string]module.Input{"": {}},
		},
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	m = module.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           module.AddOnModule,
		DeploymentType: module.SingleDeployment,
		Inputs: module.Inputs{
			Configs: map[string]module.Input{"": {}},
		},
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
}

func TestValidateModuleDependencies(t *testing.T) {
	var d map[string]string
	if err := validateModuleDependencies(d); err != nil {
		t.Errorf("validateModuleDependencies(%v); err != nil", d)
	}
	d = make(map[string]string)
	d["test.test/test"] = "=v1.0.0"
	if err := validateModuleDependencies(d); err != nil {
		t.Errorf("validateModuleDependencies(%v); err != nil", d)
	}
	d["test.test/test"] = "v1.0.0"
	if err := validateModuleDependencies(d); err == nil {
		t.Errorf("validateModuleDependencies(%v); err == nil", d)
	}
	delete(d, "test.test/test")
	d["test"] = "=v1.0.0"
	if err := validateModuleDependencies(d); err == nil {
		t.Errorf("validateModuleDependencies(%v); err == nil", d)
	}
}

func TestIsValidModuleID(t *testing.T) {
	ok := []string{
		"test.test/test",
		"test.test/test/test",
		"test-123_test.test/123-test_456",
	}
	notOk := []string{
		"/",
		"test123",
		"test.test",
		"test.test/test/",
		"/test.test/test",
		"test.test.test/test",
		"http://test.test/test",
		"test.!§$%&/()=?123/test",
		"test!§$%&/()=?.test/test",
	}
	for _, s := range ok {
		if isValidModuleID(s) != true {
			t.Errorf("isValidModuleID(\"%s\") != true", s)
		}
	}
	for _, s := range notOk {
		if isValidModuleID(s) != false {
			t.Errorf("isValidModuleID(\"%s\") != false", s)
		}
	}
}

func TestIsValidModuleType(t *testing.T) {
	if isValidModuleType(module.AddOnModule) != true {
		t.Errorf("isValidModuleType(\"%s\") != true", module.AddOnModule)
	}
	if isValidModuleType("test") != false {
		t.Error("isValidModuleType(\"test\") != false")
	}
}

func TestIsValidDeploymentType(t *testing.T) {
	if isValidDeploymentType(module.SingleDeployment) != true {
		t.Errorf("isValidDeploymentType(\"%s\") != true", module.SingleDeployment)
	}
	if isValidDeploymentType("test") != false {
		t.Error("isValidDeploymentType(\"test\") != false")
	}
}

func TestIsValidCPUArch(t *testing.T) {
	if isValidCPUArch(module.X86_64) != true {
		t.Errorf("isValidCPUArch(\"%s\") != true", module.X86_64)
	}
	if isValidCPUArch("test") != false {
		t.Error("isValidCPUArch(\"test\") != false")
	}
}

func TestValidateConfigs(t *testing.T) {
	var mCs module.Configs
	var inputs map[string]module.Input
	if err := validateConfigs(mCs, inputs); err != nil {
		t.Error("err != nil")
	}
	// ------------------------------
	str := "test"
	mCs = make(module.Configs)
	mCs.SetString(str, nil, nil, false, "", nil, false)
	if err := validateConfigs(mCs, inputs); err != nil {
		t.Error("err != nil")
	}
	// ------------------------------
	mCs = make(module.Configs)
	mCs.SetString("", nil, nil, false, "", nil, false)
	if err := validateConfigs(mCs, inputs); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	mCs = make(module.Configs)
	mCs.SetString(str, nil, nil, false, "", nil, true)
	if err := validateConfigs(mCs, inputs); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	mCs = make(module.Configs)
	mCs.SetString(str, &str, nil, false, "", nil, true)
	if err := validateConfigs(mCs, inputs); err != nil {
		t.Error("err != nil")
	}
	// ------------------------------
	mCs = make(module.Configs)
	inputs = make(map[string]module.Input)
	inputs[str] = module.Input{}
	mCs.SetString(str, nil, nil, false, "", nil, true)
	if err := validateConfigs(mCs, inputs); err != nil {
		t.Error("err != nil")
	}
}

func TestValidateResources(t *testing.T) {
	var mRs map[string]module.Resource
	var inputs map[string]module.Input
	if err := validateResources(mRs, inputs); err != nil {
		t.Error("err != nil")
	}
	// ------------------------------
	str := "test"
	mRs = make(map[string]module.Resource)
	mRs[str] = module.Resource{
		Tags:     nil,
		Required: false,
	}
	if err := validateResources(mRs, inputs); err != nil {
		t.Error("err != nil")
	}
	// ------------------------------
	mRs[str] = module.Resource{
		Tags:     map[string]struct{}{"": {}},
		Required: true,
	}
	if err := validateResources(mRs, inputs); err != nil {
		t.Error("err != nil")
	}
	// ------------------------------
	mRs = map[string]module.Resource{
		"": {},
	}
	if err := validateResources(mRs, inputs); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	mRs = map[string]module.Resource{
		str: {
			Tags:     nil,
			Required: true,
		},
	}
	if err := validateResources(mRs, inputs); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	inputs = make(map[string]module.Input)
	inputs[str] = module.Input{}
	mRs[str] = module.Resource{
		Tags:     nil,
		Required: true,
	}
	if err := validateResources(mRs, inputs); err != nil {
		t.Error("err != nil")
	}
}

func TestValidateSecrets(t *testing.T) {
	var mSs map[string]module.Secret
	var inputs map[string]module.Input
	if err := validateSecrets(mSs, inputs); err != nil {
		t.Error("err != nil")
	}
	// ------------------------------
	str := "test"
	mSs = make(map[string]module.Secret)
	mSs[str] = module.Secret{
		Tags:     nil,
		Required: false,
	}
	if err := validateSecrets(mSs, inputs); err != nil {
		t.Error("err != nil")
	}
	// ------------------------------
	mSs[str] = module.Secret{
		Tags:     map[string]struct{}{"": {}},
		Required: true,
	}
	if err := validateSecrets(mSs, inputs); err != nil {
		t.Error("err != nil")
	}
	// ------------------------------
	mSs = map[string]module.Secret{
		"": {},
	}
	if err := validateSecrets(mSs, inputs); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	mSs = map[string]module.Secret{
		str: {
			Tags:     nil,
			Required: true,
		},
	}
	if err := validateSecrets(mSs, inputs); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	inputs = make(map[string]module.Input)
	inputs[str] = module.Input{}
	mSs[str] = module.Secret{
		Tags:     nil,
		Required: true,
	}
	if err := validateSecrets(mSs, inputs); err != nil {
		t.Error("err != nil")
	}
}
