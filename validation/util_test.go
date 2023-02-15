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
	"testing"
)

func TestValidateKeyNotEmptyString(t *testing.T) {
	var m map[string]struct{}
	if validateKeyNotEmptyString(m) != true {
		t.Error("validateKeyNotEmptyString(m) != true")
	}
	m = make(map[string]struct{})
	if validateKeyNotEmptyString(m) != true {
		t.Error("validateKeyNotEmptyString(m) != true")
	}
	m["test"] = struct{}{}
	if validateKeyNotEmptyString(m) != true {
		t.Error("validateKeyNotEmptyString(m) != true")
	}
	m[""] = struct{}{}
	if validateKeyNotEmptyString(m) != false {
		t.Error("validateKeyNotEmptyString(m) != false")
	}
}

func TestValidateMapKeys(t *testing.T) {
	var m map[string]struct{}
	k := make(map[string]struct{})
	if err := validateMapKeys(m, k); err != nil {
		t.Error("err != nil")
	}
	if len(k) != 0 {
		t.Error("len(k) != 0")
	}
	m = make(map[string]struct{})
	m["test"] = struct{}{}
	if err := validateMapKeys(m, k); err != nil {
		t.Error("err != nil")
	}
	if len(k) != 1 {
		t.Error("len(k) != 1")
	}
	if _, ok := k["test"]; !ok {
		t.Error("_, ok := k[\"test\"]; !ok")
	}
	if err := validateMapKeys(m, k); err == nil {
		t.Error("err == nil")
	}
	if len(k) != 1 {
		t.Error("len(k) != 1")
	}
	delete(m, "test")
	delete(k, "test")
	m[""] = struct{}{}
	if err := validateMapKeys(m, k); err == nil {
		t.Error("err == nil")
	}
	if len(k) != 0 {
		t.Error("len(k) != 0")
	}
}
