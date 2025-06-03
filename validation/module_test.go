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
	"github.com/SENERGY-Platform/mgw-module-lib/model"
	"testing"
)

func TestValidate(t *testing.T) {
	m := model.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           model.AddOnModule,
		DeploymentType: model.SingleDeployment,
	}
	if err := Validate(&m); err != nil {
		t.Error("err != nil")
	}
	// ------------------------------
	m = model.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           model.AddOnModule,
		DeploymentType: model.SingleDeployment,
		Services: map[string]*model.Service{
			"a": {RequiredSrv: map[string]struct{}{"a": {}}},
		},
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	m = model.Module{
		ID: "",
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	m = model.Module{
		ID:      "test.test/test",
		Version: "",
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	m = model.Module{
		ID:      "test.test/test",
		Version: "v1.0.0",
		Type:    "",
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	m = model.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           model.AddOnModule,
		DeploymentType: "",
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	m = model.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           model.AddOnModule,
		DeploymentType: model.SingleDeployment,
		Volumes:        map[string]struct{}{"": {}},
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	m = model.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           model.AddOnModule,
		DeploymentType: model.SingleDeployment,
		Dependencies:   map[string]string{"": ""},
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	m = model.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           model.AddOnModule,
		DeploymentType: model.SingleDeployment,
		HostResources:  map[string]model.HostResource{"": {}},
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	m = model.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           model.AddOnModule,
		DeploymentType: model.SingleDeployment,
		Secrets:        map[string]model.Secret{"": {}},
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	m = model.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           model.AddOnModule,
		DeploymentType: model.SingleDeployment,
		Configs:        model.Configs{"": {}},
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	m = model.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           model.AddOnModule,
		DeploymentType: model.SingleDeployment,
		Inputs: model.Inputs{
			Groups: map[string]model.InputGroup{"": {}},
		},
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	g := ""
	m = model.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           model.AddOnModule,
		DeploymentType: model.SingleDeployment,
		Inputs: model.Inputs{
			Resources: map[string]model.Input{"test": {Group: &g}},
		},
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	m = model.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           model.AddOnModule,
		DeploymentType: model.SingleDeployment,
		Inputs: model.Inputs{
			Secrets: map[string]model.Input{"test": {Group: &g}},
		},
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	m = model.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           model.AddOnModule,
		DeploymentType: model.SingleDeployment,
		Inputs: model.Inputs{
			Configs: map[string]model.Input{"test": {Group: &g}},
		},
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	m = model.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           model.AddOnModule,
		DeploymentType: model.SingleDeployment,
		Inputs: model.Inputs{
			Resources: map[string]model.Input{"": {}},
		},
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	m = model.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           model.AddOnModule,
		DeploymentType: model.SingleDeployment,
		Inputs: model.Inputs{
			Secrets: map[string]model.Input{"": {}},
		},
	}
	if err := Validate(&m); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	m = model.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           model.AddOnModule,
		DeploymentType: model.SingleDeployment,
		Inputs: model.Inputs{
			Configs: map[string]model.Input{"": {}},
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
	if isValidModuleType(model.AddOnModule) != true {
		t.Errorf("isValidModuleType(\"%s\") != true", model.AddOnModule)
	}
	if isValidModuleType("test") != false {
		t.Error("isValidModuleType(\"test\") != false")
	}
}

func TestIsValidDeploymentType(t *testing.T) {
	if isValidDeploymentType(model.SingleDeployment) != true {
		t.Errorf("isValidDeploymentType(\"%s\") != true", model.SingleDeployment)
	}
	if isValidDeploymentType("test") != false {
		t.Error("isValidDeploymentType(\"test\") != false")
	}
}

func TestIsValidCPUArch(t *testing.T) {
	if isValidCPUArch(model.X86_64) != true {
		t.Errorf("isValidCPUArch(\"%s\") != true", model.X86_64)
	}
	if isValidCPUArch("test") != false {
		t.Error("isValidCPUArch(\"test\") != false")
	}
}

func TestValidateConfigs(t *testing.T) {
	var mCs model.Configs
	var inputs map[string]model.Input
	if err := validateConfigs(mCs, inputs); err != nil {
		t.Error("err != nil")
	}
	// ------------------------------
	str := "test"
	mCs = make(model.Configs)
	mCs.SetString(str, nil, nil, false, "", nil, false)
	if err := validateConfigs(mCs, inputs); err != nil {
		t.Error("err != nil")
	}
	// ------------------------------
	mCs = make(model.Configs)
	mCs.SetString("", nil, nil, false, "", nil, false)
	if err := validateConfigs(mCs, inputs); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	mCs = make(model.Configs)
	mCs.SetString(str, nil, nil, false, "", nil, true)
	if err := validateConfigs(mCs, inputs); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	mCs = make(model.Configs)
	mCs.SetString(str, &str, nil, false, "", nil, true)
	if err := validateConfigs(mCs, inputs); err != nil {
		t.Error("err != nil")
	}
	// ------------------------------
	mCs = make(model.Configs)
	inputs = make(map[string]model.Input)
	inputs[str] = model.Input{}
	mCs.SetString(str, nil, nil, false, "", nil, true)
	if err := validateConfigs(mCs, inputs); err != nil {
		t.Error("err != nil")
	}
}

func TestValidateResources(t *testing.T) {
	var mRs map[string]model.HostResource
	var inputs map[string]model.Input
	if err := validateResources(mRs, inputs); err != nil {
		t.Error("err != nil")
	}
	// ------------------------------
	str := "test"
	mRs = make(map[string]model.HostResource)
	mRs[str] = model.HostResource{
		Resource: model.Resource{
			Tags:     nil,
			Required: false,
		},
	}
	if err := validateResources(mRs, inputs); err != nil {
		t.Error("err != nil")
	}
	// ------------------------------
	mRs[str] = model.HostResource{
		Resource: model.Resource{
			Tags:     map[string]struct{}{"": {}},
			Required: true,
		},
	}
	if err := validateResources(mRs, inputs); err != nil {
		t.Error("err != nil")
	}
	// ------------------------------
	mRs = map[string]model.HostResource{
		"": {},
	}
	if err := validateResources(mRs, inputs); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	mRs = map[string]model.HostResource{
		str: {
			Resource: model.Resource{
				Tags:     nil,
				Required: true,
			},
		},
	}
	if err := validateResources(mRs, inputs); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	inputs = make(map[string]model.Input)
	inputs[str] = model.Input{}
	mRs[str] = model.HostResource{
		Resource: model.Resource{
			Tags:     nil,
			Required: true,
		},
	}
	if err := validateResources(mRs, inputs); err != nil {
		t.Error("err != nil")
	}
}

func TestValidateSecrets(t *testing.T) {
	var mSs map[string]model.Secret
	var inputs map[string]model.Input
	if err := validateSecrets(mSs, inputs); err != nil {
		t.Error("err != nil")
	}
	// ------------------------------
	str := "test"
	mSs = make(map[string]model.Secret)
	mSs[str] = model.Secret{
		Resource: model.Resource{
			Tags:     nil,
			Required: false,
		},
	}
	if err := validateSecrets(mSs, inputs); err != nil {
		t.Error("err != nil")
	}
	// ------------------------------
	mSs[str] = model.Secret{
		Resource: model.Resource{
			Tags:     map[string]struct{}{"": {}},
			Required: true,
		},
	}
	if err := validateSecrets(mSs, inputs); err != nil {
		t.Error("err != nil")
	}
	// ------------------------------
	mSs = map[string]model.Secret{
		"": {},
	}
	if err := validateSecrets(mSs, inputs); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	mSs = map[string]model.Secret{
		str: {
			Resource: model.Resource{
				Tags:     nil,
				Required: true,
			},
		},
	}
	if err := validateSecrets(mSs, inputs); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	inputs = make(map[string]model.Input)
	inputs[str] = model.Input{}
	mSs[str] = model.Secret{
		Resource: model.Resource{
			Tags:     nil,
			Required: true,
		},
	}
	if err := validateSecrets(mSs, inputs); err != nil {
		t.Error("err != nil")
	}
}
