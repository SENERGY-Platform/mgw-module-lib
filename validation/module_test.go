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
	"fmt"
	"github.com/SENERGY-Platform/mgw-module-lib/model"
	"testing"
)

func TestValidate(t *testing.T) {
	m := model.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           model.AddOnModule,
		DeploymentType: model.SingleDeployment,
		Services: map[string]*model.Service{
			"a": {},
			"b": {RequiredSrv: map[string]struct{}{"a": {}}},
		},
	}
	if err := Validate(m); err != nil {
		fmt.Println(err)
		t.Errorf("Validate(%v); err != nil", m)
	}
	m = model.Module{
		ID:             "test.test/test",
		Version:        "v1.0.0",
		Type:           model.AddOnModule,
		DeploymentType: model.SingleDeployment,
		Services: map[string]*model.Service{
			"a": {RequiredSrv: map[string]struct{}{"a": {}}},
		},
	}
	if err := Validate(m); err == nil {
		fmt.Println(err)
		t.Errorf("Validate(%v); err == nil", m)
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
