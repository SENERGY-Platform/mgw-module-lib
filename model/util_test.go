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
