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

package module

import (
	"reflect"
	"testing"
)

func TestNewConfigValue(t *testing.T) {
	str := "test"
	cv1 := newConfigValue[string](nil, nil, StringType, false, str, nil, true)
	if cv1.IsSlice != false {
		t.Error("cv1.IsSlice != false")
	}
	if cv1.Default != nil {
		t.Error("cv1.Default != nil")
	}
	if cv1.Options != nil {
		t.Error("cv1.Options != nil")
	}
	if cv1.OptionsLen() != 0 {
		t.Error("cv1.OptionsLen() != 0")
	}
	if cv1.DataType != StringType {
		t.Error("cv1.DataType !=", StringType)
	}
	if cv1.OptExt != false {
		t.Error("cv1.OptExt != false")
	}
	if cv1.Type != str {
		t.Error("cv1.Type !=", str)
	}
	if cv1.TypeOpt != nil {
		t.Error("cv1.TypeOpt != nil")
	}
	if cv1.Required == false {
		t.Error("cv1.Required == false")
	}
	var opt []string
	cto := make(ConfigTypeOptions)
	cv2 := newConfigValue(nil, opt, StringType, false, str, cto, false)
	if cv2.Options != nil {
		t.Error("cv2.Options != nil")
	}
	if cv2.OptionsLen() != 0 {
		t.Error("cv2.OptionsLen() != 0")
	}
	if cv2.TypeOpt != nil {
		t.Error("cv2.TypeOpt != nil")
	}
	opt = append(opt, str)
	cto.SetString(str, str)
	cv3 := newConfigValue(&str, opt, StringType, true, str, cto, false)
	if cv3.Default == nil {
		t.Error("cv3.Default == nil")
	}
	if cv3.Default.(string) != str {
		t.Error("cv3.Default.(string) !=", str)
	}
	if cv3.Options == nil {
		t.Error("cv3.Options == nil")
	}
	if cv3.OptionsLen() != 1 {
		t.Error("cv3.OptionsLen() != 1")
	}
	if cv3.Options.([]string)[0] != str {
		t.Error("cv3.Options.([]string)[0] !=", str)
	}
	if cv3.OptExt != true {
		t.Error("cv3.OptExt != true")
	}
	if cv3.TypeOpt == nil {
		t.Error("cv3.TypeOpt == nil")
	}
}

func TestNewConfigValueSlice(t *testing.T) {
	str := "test"
	cvs1 := newConfigValueSlice[string](nil, nil, StringType, false, str, nil, str, true)
	if cvs1.IsSlice != true {
		t.Error("cvs1.IsSlice != true")
	}
	if cvs1.Default != nil {
		t.Error("cvs1.Default != nil")
	}
	if cvs1.Options != nil {
		t.Error("cvs1.Options != nil")
	}
	if cvs1.OptionsLen() != 0 {
		t.Error("cvs1.OptionsLen() != 0")
	}
	if cvs1.DataType != StringType {
		t.Error("cvs1.DataType !=", StringType)
	}
	if cvs1.OptExt != false {
		t.Error("cvs1.OptExt != false")
	}
	if cvs1.Type != str {
		t.Error("cvs1.Type !=", str)
	}
	if cvs1.TypeOpt != nil {
		t.Error("cvs1.TypeOpt != nil")
	}
	if cvs1.Delimiter != str {
		t.Error("cvs1.Delimiter !=", str)
	}
	if cvs1.Required == false {
		t.Error("cvs1.Required == false")
	}
	var def []string
	var opt []string
	cto := make(ConfigTypeOptions)
	cvs2 := newConfigValueSlice(def, opt, StringType, false, str, cto, "", false)
	if cvs2.Default != nil {
		t.Error("cvs2.Default != nil")
	}
	if cvs2.Options != nil {
		t.Error("cvs2.Options != nil")
	}
	if cvs2.OptionsLen() != 0 {
		t.Error("cvs2.OptionsLen() != 0")
	}
	if cvs2.TypeOpt != nil {
		t.Error("cvs2.TypeOpt != nil")
	}
	def = append(def, str)
	opt = append(opt, str)
	cto.SetString(str, str)
	cvs3 := newConfigValueSlice(def, opt, StringType, true, str, cto, "", false)
	if cvs3.Default == nil {
		t.Error("cvs3.Default == nil")
	}
	if cvs3.Default.([]string)[0] != str {
		t.Error("cvs3.Default.([]string)[0] !=", str)
	}
	if cvs3.Options == nil {
		t.Error("cvs3.Options == nil")
	}
	if cvs3.OptionsLen() != 1 {
		t.Error("cvs3.OptionsLen() != 1")
	}
	if cvs3.Options.([]string)[0] != str {
		t.Error("cvs3.Options.([]string)[0] !=", str)
	}
	if cvs3.OptExt != true {
		t.Error("cvs3.OptExt != true")
	}
	if cvs3.TypeOpt == nil {
		t.Error("cvs3.TypeOpt == nil")
	}
}

func TestConfigs_SetString(t *testing.T) {
	configs := make(Configs)
	configs.SetString("", nil, nil, false, "", nil, false)
	for _, config := range configs {
		if config.DataType != StringType {
			t.Error("config.DataType != StringType")
		}
		if config.IsSlice != false {
			t.Error("config.IsSlice != false")
		}
	}
}

func TestConfigs_SetStringSlice(t *testing.T) {
	configs := make(Configs)
	configs.SetStringSlice("", nil, nil, false, "", nil, "", false)
	for _, config := range configs {
		if config.DataType != StringType {
			t.Error("config.DataType != StringType")
		}
		if config.IsSlice != true {
			t.Error("config.IsSlice != true")
		}
	}
}

func TestConfigs_SetBool(t *testing.T) {
	configs := make(Configs)
	configs.SetBool("", nil, nil, false, "", nil, false)
	for _, config := range configs {
		if config.DataType != BoolType {
			t.Error("config.DataType != BoolType")
		}
		if config.IsSlice != false {
			t.Error("config.IsSlice != false")
		}
	}
}

func TestConfigs_SetBoolSlice(t *testing.T) {
	configs := make(Configs)
	configs.SetBoolSlice("", nil, nil, false, "", nil, "", false)
	for _, config := range configs {
		if config.DataType != BoolType {
			t.Error("config.DataType != BoolType")
		}
		if config.IsSlice != true {
			t.Error("config.IsSlice != true")
		}
	}
}

func TestConfigs_SetFloat64(t *testing.T) {
	configs := make(Configs)
	configs.SetFloat64("", nil, nil, false, "", nil, false)
	for _, config := range configs {
		if config.DataType != Float64Type {
			t.Error("config.DataType != Float64Type")
		}
		if config.IsSlice != false {
			t.Error("config.IsSlice != false")
		}
	}
}

func TestConfigs_SetFloat64Slice(t *testing.T) {
	configs := make(Configs)
	configs.SetFloat64Slice("", nil, nil, false, "", nil, "", false)
	for _, config := range configs {
		if config.DataType != Float64Type {
			t.Error("config.DataType != Float64Type")
		}
		if config.IsSlice != true {
			t.Error("config.IsSlice != true")
		}
	}
}

func TestConfigs_SetInt64(t *testing.T) {
	configs := make(Configs)
	configs.SetInt64("", nil, nil, false, "", nil, false)
	for _, config := range configs {
		if config.DataType != Int64Type {
			t.Error("config.DataType != Int64Type")
		}
		if config.IsSlice != false {
			t.Error("config.IsSlice != false")
		}
	}
}

func TestConfigs_SetInt64Slice(t *testing.T) {
	configs := make(Configs)
	configs.SetInt64Slice("", nil, nil, false, "", nil, "", false)
	for _, config := range configs {
		if config.DataType != Int64Type {
			t.Error("config.DataType != Int64Type")
		}
		if config.IsSlice != true {
			t.Error("config.IsSlice != true")
		}
	}
}

func TestConfigTypeOptions_SetString(t *testing.T) {
	cto := make(ConfigTypeOptions)
	val := "test"
	cto.SetString("", val)
	for _, opt := range cto {
		if opt.DataType != StringType {
			t.Error("opt.DataType != StringType")
		}
		if opt.Value != val {
			t.Errorf("opt.Value != \"%s\"", val)
		}
	}
}

func TestConfigTypeOptions_SetBool(t *testing.T) {
	cto := make(ConfigTypeOptions)
	val := true
	cto.SetBool("", val)
	for _, opt := range cto {
		if opt.DataType != BoolType {
			t.Error("opt.DataType != BoolType")
		}
		if opt.Value != val {
			t.Errorf("opt.Value != %v", val)
		}
	}
}

func TestConfigTypeOptions_SetFloat64(t *testing.T) {
	cto := make(ConfigTypeOptions)
	val := float64(1.0)
	cto.SetFloat64("", val)
	for _, opt := range cto {
		if opt.DataType != Float64Type {
			t.Error("opt.DataType != Float64Type")
		}
		if opt.Value != val {
			t.Errorf("opt.Value != %f", val)
		}
	}
}

func TestConfigTypeOptions_SetInt64(t *testing.T) {
	cto := make(ConfigTypeOptions)
	val := int64(1)
	cto.SetInt64("", val)
	for _, opt := range cto {
		if opt.DataType != Int64Type {
			t.Error("opt.DataType != Int64Type")
		}
		if opt.Value != val {
			t.Errorf("opt.Value != %d", val)
		}
	}
}

func TestConfigValue_OptionsLen(t *testing.T) {
	cv1 := newConfigValue(nil, []string{"test"}, StringType, false, "", nil, false)
	if cv1.OptionsLen() != 1 {
		t.Error("cv1.OptionsLen() != 1")
	}
	cv2 := newConfigValue(nil, []bool{true}, StringType, false, "", nil, false)
	if cv2.OptionsLen() != 1 {
		t.Error("cv2.OptionsLen() != 1")
	}
	cv3 := newConfigValue(nil, []int64{1}, StringType, false, "", nil, false)
	if cv3.OptionsLen() != 1 {
		t.Error("cv3.OptionsLen() != 1")
	}
	cv4 := newConfigValue(nil, []float64{1.0}, StringType, false, "", nil, false)
	if cv4.OptionsLen() != 1 {
		t.Error("cv4.OptionsLen() != 1")
	}
}

func TestGetServiceStartOrder(t *testing.T) {
	sAref := "A"
	sBref := "B"
	sCref := "C"
	order := []string{sCref, sAref, sBref} // (C -> A -> B)
	var services map[string]*Service
	if o, err := GetServiceStartOrder(services); err != nil {
		t.Errorf("GetServiceStartOrder(%v); err != nil", services)
	} else if len(o) > 0 {
		t.Error("len(o) > 0")
	}
	services = make(map[string]*Service)
	if o, err := GetServiceStartOrder(services); err != nil {
		t.Errorf("GetServiceStartOrder(%v); err != nil", services)
	} else if len(o) > 0 {
		t.Error("len(o) > 0")
	}
	// add service "A" which requires service "C" (C -> A)
	services[sAref] = &Service{RequiredSrv: map[string]struct{}{sCref: {}}}
	// add service "B" which requires service "A" (A -> B)
	services[sBref] = &Service{RequiredSrv: map[string]struct{}{sAref: {}}}
	// add service "C"
	services[sCref] = &Service{}
	if o, err := GetServiceStartOrder(services); err != nil {
		t.Errorf("GetServiceStartOrder(%v); err != nil", services)
	} else if !reflect.DeepEqual(order, o) {
		t.Errorf("%v != %v", order, o)
	}
}

func TestSrvRefTarget_FillTemplate(t *testing.T) {
	str := "test"
	target := SrvRefTarget{}
	a := str
	b := target.FillTemplate(str)
	if a != b {
		t.Errorf("%s != %s", a, b)
	}
	tmp := "http://{" + RefPlaceholder + "}/api"
	target = SrvRefTarget{Template: &tmp}
	a = "http://" + str + "/api"
	b = target.FillTemplate(str)
	if a != b {
		t.Errorf("%s != %s", a, b)
	}
	tmp2 := "http://api"
	target = SrvRefTarget{Template: &tmp2}
	a = tmp2
	b = target.FillTemplate(str)
	if a != b {
		t.Errorf("%s != %s", a, b)
	}
}
