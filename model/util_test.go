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

package model

import (
	"testing"
)

func TestNewConfigValue(t *testing.T) {
	str := "test"
	cv1 := newConfigValue[string](nil, nil, StringType, false, str, nil)
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
	var opt []string
	cto := make(ConfigTypeOptions)
	cv2 := newConfigValue(nil, opt, StringType, false, str, cto)
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
	cv3 := newConfigValue(&str, opt, StringType, true, str, cto)
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
	cvs1 := newConfigValueSlice[string](nil, nil, StringType, false, str, nil, nil)
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
	if cvs1.Delimiter != nil {
		t.Error("cvs1.Delimiter != nil")
	}
	var def []string
	var opt []string
	cto := make(ConfigTypeOptions)
	cvs2 := newConfigValueSlice(def, opt, StringType, false, str, cto, nil)
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
	cvs3 := newConfigValueSlice(def, opt, StringType, true, str, cto, &str)
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
	if cvs3.Delimiter == nil {
		t.Error("cvs3.Delimiter == nil")
	}
	if *cvs3.Delimiter != str {
		t.Error("*cvs3.Delimiter != str")
	}
}

func TestIsValidPort(t *testing.T) {
	if isValidPort(nil) == true {
		t.Error("isValidPort(nil) == true")
	}
	if isValidPort([]uint{}) == true {
		t.Error("isValidPort([]uint{}) == true")
	}
	if isValidPort([]uint{1, 1}) == true {
		t.Error("isValidPort([]uint{1, 1}) == true")
	}
	if isValidPort([]uint{2, 1}) == true {
		t.Error("isValidPort([]uint{2, 1}) == true")
	}
	if isValidPort([]uint{1, 2, 3}) == true {
		t.Error("isValidPort([]uint{1, 2, 3}) == true")
	}
	if isValidPort([]uint{1}) == false {
		t.Error("isValidPort([]uint{1}) == false")
	}
	if isValidPort([]uint{1, 2}) == false {
		t.Error("isValidPort([]uint{1, 2}) == false")
	}
	if isValidPort([]uint{1, 10}) == false {
		t.Error("isValidPort([]uint{1, 10}) == false")
	}
}

func TestPortMappings_Add(t *testing.T) {
	pm := make(PortMappings)
	if err := pm.Add(nil, nil, nil, TcpPort); err == nil {
		t.Errorf("pm.Add(nil, nil, nil, \"%s\"); err == nil", TcpPort)
	}
	if len(pm) != 0 {
		t.Error("len(pm) != 0")
	}
	if err := pm.Add(nil, []uint{0}, []uint{1, 0}, TcpPort); err == nil {
		t.Errorf("pm.Add(nil, []uint{0}, []uint{1, 0}, \"%s\"); err == nil", TcpPort)
	}
	if len(pm) != 0 {
		t.Error("len(pm) != 0")
	}
	if err := pm.Add(nil, []uint{0}, nil, ""); err == nil {
		t.Error("Add(nil, []uint{0}, nil, \"\") -> err == nil")
	}
	if len(pm) != 0 {
		t.Error("len(pm) != 0")
	}
	if err := pm.Add(nil, []uint{0}, nil, TcpPort); err != nil {
		t.Errorf("pm.Add(nil, []uint{0}, nil, \"%s\"); err != nil", TcpPort)
	}
	if len(pm) != 1 {
		t.Error("len(pm) != 1")
	}
	if err := pm.Add(nil, []uint{0}, []uint{0}, TcpPort); err != nil {
		t.Errorf("pm.Add(nil, []uint{0}, []uint{0}, \"%s\"); err != nil", TcpPort)
	}
	if len(pm) != 1 {
		t.Error("len(pm) != 1")
	}
	str := "test"
	if err := pm.Add(&str, []uint{0}, nil, TcpPort); err != nil {
		t.Errorf("pm.Add(\"%s\", []uint{0}, nil, \"%s\"); err != nil", str, TcpPort)
	}
	if len(pm) != 1 {
		t.Error("len(pm) != 1")
	}
}
