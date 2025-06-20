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

package configs

import (
	"errors"
	"github.com/SENERGY-Platform/mgw-module-lib/model"
	"github.com/SENERGY-Platform/mgw-module-lib/validation/configs/definitions"
	"github.com/SENERGY-Platform/mgw-module-lib/validation/configs/validators"
	"reflect"
	"testing"
)

func TestGenVltOptParams(t *testing.T) {
	cDefVP := make(map[string]definitions.ConfigDefinitionValidatorParam)
	var cTypeO model.ConfigTypeOptions
	if b := genVltOptParams(cDefVP, cTypeO); len(b) != 0 {
		t.Errorf("len(%v) != 0", b)
	}
	// ------------------------------
	str := "test"
	vRef := "value"
	oRef := ".opt"
	cDefVP[""] = definitions.ConfigDefinitionValidatorParam{
		Value: str,
		Ref:   nil,
	}
	a := map[string]any{
		"": str,
	}
	b := genVltOptParams(cDefVP, cTypeO)
	if reflect.DeepEqual(a, b) == false {
		t.Errorf("%v != %v", a, b)
	}
	// ------------------------------
	cDefVP[""] = definitions.ConfigDefinitionValidatorParam{
		Value: str,
		Ref:   &vRef,
	}
	b = genVltOptParams(cDefVP, cTypeO)
	if reflect.DeepEqual(a, b) == false {
		t.Errorf("%v != %v", a, b)
	}
	// ------------------------------
	cTypeO = make(model.ConfigTypeOptions)
	cDefVP[""] = definitions.ConfigDefinitionValidatorParam{
		Value: str,
		Ref:   &oRef,
	}
	b = genVltOptParams(cDefVP, cTypeO)
	if reflect.DeepEqual(a, b) == false {
		t.Errorf("%v != %v", a, b)
	}
	// ------------------------------
	cTypeO.SetString("opt", str)
	cDefVP[""] = definitions.ConfigDefinitionValidatorParam{
		Value: nil,
		Ref:   &oRef,
	}
	b = genVltOptParams(cDefVP, cTypeO)
	if reflect.DeepEqual(a, b) == false {
		t.Errorf("%v != %v", a, b)
	}
	// ------------------------------
	cDefVP[""] = definitions.ConfigDefinitionValidatorParam{
		Value: "val",
		Ref:   &oRef,
	}
	b = genVltOptParams(cDefVP, cTypeO)
	if reflect.DeepEqual(a, b) == false {
		t.Errorf("%v != %v", a, b)
	}
	// ------------------------------
	cDefVP[""] = definitions.ConfigDefinitionValidatorParam{
		Value: nil,
		Ref:   &vRef,
	}
	if b = genVltOptParams(cDefVP, cTypeO); len(b) != 0 {
		t.Errorf("len(%v) != 0", b)
	}
	// ------------------------------
	cDefVP[""] = definitions.ConfigDefinitionValidatorParam{
		Value: nil,
		Ref:   &vRef,
	}
	if b = genVltOptParams(cDefVP, cTypeO); len(b) != 0 {
		t.Errorf("len(%v) != 0", b)
	}
	// ------------------------------
	oRef2 := ".opt2"
	cDefVP[""] = definitions.ConfigDefinitionValidatorParam{
		Value: nil,
		Ref:   &oRef2,
	}
	if b = genVltOptParams(cDefVP, cTypeO); len(b) != 0 {
		t.Errorf("len(%v) != 0", b)
	}
}

func TestGenVltValParams(t *testing.T) {
	cDefVP := make(map[string]definitions.ConfigDefinitionValidatorParam)
	var cTypeO model.ConfigTypeOptions
	if b := genVltValParams(cDefVP, cTypeO, nil); len(b) != 0 {
		t.Errorf("len(%v) != 0", b)
	}
	// ------------------------------
	str := "test"
	vRef := "value"
	oRef := ".opt"
	cDefVP[""] = definitions.ConfigDefinitionValidatorParam{
		Value: str,
		Ref:   nil,
	}
	a := map[string]any{
		"": str,
	}
	b := genVltValParams(cDefVP, cTypeO, nil)
	if reflect.DeepEqual(a, b) == false {
		t.Errorf("%v != %v", a, b)
	}
	// ------------------------------
	cDefVP[""] = definitions.ConfigDefinitionValidatorParam{
		Value: nil,
		Ref:   &vRef,
	}
	b = genVltValParams(cDefVP, cTypeO, str)
	if reflect.DeepEqual(a, b) == false {
		t.Errorf("%v != %v", a, b)
	}
	// ------------------------------
	cDefVP[""] = definitions.ConfigDefinitionValidatorParam{
		Value: "test2",
		Ref:   &vRef,
	}
	b = genVltValParams(cDefVP, cTypeO, str)
	if reflect.DeepEqual(a, b) == false {
		t.Errorf("%v != %v", a, b)
	}
	// ------------------------------
	cTypeO = make(model.ConfigTypeOptions)
	cDefVP[""] = definitions.ConfigDefinitionValidatorParam{
		Value: str,
		Ref:   &oRef,
	}
	b = genVltValParams(cDefVP, cTypeO, nil)
	if reflect.DeepEqual(a, b) == false {
		t.Errorf("%v != %v", a, b)
	}
	// ------------------------------
	cDefVP[""] = definitions.ConfigDefinitionValidatorParam{
		Value: nil,
		Ref:   &oRef,
	}
	if b = genVltValParams(cDefVP, cTypeO, nil); len(b) != 0 {
		t.Errorf("len(%v) != 0", b)
	}
	// ------------------------------
	cTypeO.SetString("opt", str)
	cDefVP[""] = definitions.ConfigDefinitionValidatorParam{
		Value: nil,
		Ref:   &oRef,
	}
	b = genVltValParams(cDefVP, cTypeO, nil)
	if reflect.DeepEqual(a, b) == false {
		t.Errorf("%v != %v", a, b)
	}
	// ------------------------------
	cDefVP[""] = definitions.ConfigDefinitionValidatorParam{
		Value: "test2",
		Ref:   &oRef,
	}
	b = genVltValParams(cDefVP, cTypeO, nil)
	if reflect.DeepEqual(a, b) == false {
		t.Errorf("%v != %v", a, b)
	}
	// ------------------------------
	cDefVP[""] = definitions.ConfigDefinitionValidatorParam{
		Value: nil,
		Ref:   &vRef,
	}
	if b = genVltValParams(cDefVP, cTypeO, nil); len(b) != 0 {
		t.Errorf("len(%v) != 0", b)
	}
}

func TestVltOptions(t *testing.T) {
	var cDefVlts []definitions.ConfigDefinitionValidator
	vlts := make(map[string]validators.Validator)
	if err := vltTypeOpts(cDefVlts, nil, vlts); err != nil {
		t.Error("err != nil")
	}
	// ------------------------------
	cDefVlts = []definitions.ConfigDefinitionValidator{
		{
			Name: "vlt",
			Parameter: map[string]definitions.ConfigDefinitionValidatorParam{
				"": {
					Value: "val",
					Ref:   nil,
				},
			},
		},
	}
	if err := vltTypeOpts(cDefVlts, nil, vlts); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	vlts["vlt"] = func(params map[string]any) error {
		return nil
	}
	if err := vltTypeOpts(cDefVlts, nil, vlts); err != nil {
		t.Error("err != nil")
	}
	// ------------------------------
	vlts["vlt"] = func(params map[string]any) error {
		return errors.New("test")
	}
	if err := vltTypeOpts(cDefVlts, nil, vlts); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	vRef := "value"
	cDefVlts = []definitions.ConfigDefinitionValidator{
		{
			Name: "vlt",
			Parameter: map[string]definitions.ConfigDefinitionValidatorParam{
				"a": {
					Value: "val",
					Ref:   nil,
				},
				"b": {
					Value: nil,
					Ref:   &vRef,
				},
			},
		},
	}
	if err := vltTypeOpts(cDefVlts, nil, vlts); err != nil {
		t.Error("err != nil")
	}
}

func TestVltValue(t *testing.T) {
	var cDefVlts []definitions.ConfigDefinitionValidator
	vlts := make(map[string]validators.Validator)
	if err := vltValue(cDefVlts, nil, vlts, nil); err != nil {
		t.Error("err != nil")
	}
	// ------------------------------
	cDefVlts = []definitions.ConfigDefinitionValidator{
		{
			Name: "vlt",
			Parameter: map[string]definitions.ConfigDefinitionValidatorParam{
				"": {
					Value: "val",
					Ref:   nil,
				},
			},
		},
	}
	if err := vltValue(cDefVlts, nil, vlts, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	vlts["vlt"] = func(params map[string]any) error {
		return nil
	}
	if err := vltValue(cDefVlts, nil, vlts, nil); err != nil {
		t.Error("err != nil")
	}
	// ------------------------------
	vlts["vlt"] = func(params map[string]any) error {
		return errors.New("test")
	}
	if err := vltValue(cDefVlts, nil, vlts, nil); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	vRef := "value"
	cDefVlts = []definitions.ConfigDefinitionValidator{
		{
			Name: "vlt",
			Parameter: map[string]definitions.ConfigDefinitionValidatorParam{
				"a": {
					Value: "val",
					Ref:   nil,
				},
				"b": {
					Value: nil,
					Ref:   &vRef,
				},
			},
		},
	}
	if err := vltValue(cDefVlts, nil, vlts, nil); err != nil {
		t.Error("err != nil")
	}
}

func TestVltBase(t *testing.T) {
	var cDef definitions.ConfigDefinition
	var cTypeOpts model.ConfigTypeOptions
	if err := vltBase(cDef, cTypeOpts, ""); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	dt := "dt"
	opt := "opt"
	cDef = definitions.ConfigDefinition{
		DataType:   model.Set[string]{dt: {}},
		Options:    nil,
		Validators: nil,
	}
	if err := vltBase(cDef, cTypeOpts, dt); err != nil {
		t.Error("err != nil")
	}
	// ------------------------------
	cDef = definitions.ConfigDefinition{
		DataType: model.Set[string]{dt: {}},
		Options: map[string]definitions.ConfigDefinitionOption{
			opt: {},
		},
		Validators: nil,
	}
	if err := vltBase(cDef, cTypeOpts, dt); err != nil {
		t.Error("err != nil")
	}
	// ------------------------------
	cDef = definitions.ConfigDefinition{
		DataType: model.Set[string]{dt: {}},
		Options: map[string]definitions.ConfigDefinitionOption{
			opt: {
				Required: true,
			},
		},
		Validators: nil,
	}
	if err := vltBase(cDef, cTypeOpts, dt); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	cTypeOpts = make(model.ConfigTypeOptions)
	cTypeOpts.SetString(opt, dt)
	cDef = definitions.ConfigDefinition{
		DataType:   model.Set[string]{dt: {}},
		Options:    nil,
		Validators: nil,
	}
	if err := vltBase(cDef, cTypeOpts, dt); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	cDef = definitions.ConfigDefinition{
		DataType: model.Set[string]{dt: {}},
		Options: map[string]definitions.ConfigDefinitionOption{
			opt: {},
		},
		Validators: nil,
	}
	if err := vltBase(cDef, cTypeOpts, dt); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	cDef = definitions.ConfigDefinition{
		DataType: model.Set[string]{dt: {}},
		Options: map[string]definitions.ConfigDefinitionOption{
			opt: {
				DataType: model.Set[string]{dt: {}},
			},
		},
		Validators: nil,
	}
	if err := vltBase(cDef, cTypeOpts, dt); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	cDef = definitions.ConfigDefinition{
		DataType: model.Set[string]{dt: {}},
		Options: map[string]definitions.ConfigDefinitionOption{
			opt: {
				Inherit: true,
			},
		},
		Validators: nil,
	}
	if err := vltBase(cDef, cTypeOpts, dt); err == nil {
		t.Error("err == nil")
	}
	// ------------------------------
	cDef = definitions.ConfigDefinition{
		DataType: model.Set[string]{dt: {}},
		Options: map[string]definitions.ConfigDefinitionOption{
			opt: {
				DataType: model.Set[string]{model.StringType: {}},
			},
		},
		Validators: nil,
	}
	if err := vltBase(cDef, cTypeOpts, dt); err != nil {
		t.Error("err != nil")
	}
	// ------------------------------
	cDef = definitions.ConfigDefinition{
		DataType: model.Set[string]{model.StringType: {}},
		Options: map[string]definitions.ConfigDefinitionOption{
			opt: {
				Inherit: true,
			},
		},
		Validators: nil,
	}
	if err := vltBase(cDef, cTypeOpts, model.StringType); err != nil {
		t.Error("err != nil")
	}
	// ------------------------------
	cTypeOpts.SetString("test", dt)
	if err := vltBase(cDef, cTypeOpts, model.StringType); err == nil {
		t.Error("err == nil")
	}
}

func TestVltValInOpt(t *testing.T) {
	if _, err := vltValInOpt[int](nil, nil); err == nil {
		t.Error("err == nil")
	}
	if _, err := vltValInOpt[int](1, nil); err == nil {
		t.Error("err == nil")
	}
	if _, err := vltValInOpt[int](nil, 1); err == nil {
		t.Error("err == nil")
	}
	if ok, err := vltValInOpt[int](1, []int{}); err != nil {
		t.Error("err == nil")
	} else if ok == true {
		t.Error("ok == true")
	}
	if ok, err := vltValInOpt[int](1, []int{1}); err != nil {
		t.Error("err == nil")
	} else if ok == false {
		t.Error("ok == false")
	}
}

func TestVltValSlInOpt(t *testing.T) {
	if _, err := vltValSlInOpt[int](nil, nil); err == nil {
		t.Error("err == nil")
	}
	if _, err := vltValSlInOpt[int]([]int{}, nil); err == nil {
		t.Error("err == nil")
	}
	if ok, err := vltValSlInOpt[int]([]int{1}, []int{}); err != nil {
		t.Error("err != nil")
	} else if ok == true {
		t.Error("ok == true")
	}
	if ok, err := vltValSlInOpt[int]([]int{1}, []int{1}); err != nil {
		t.Error("err != nil")
	} else if ok == false {
		t.Error("ok == false")
	}
}
