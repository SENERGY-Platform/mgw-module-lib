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
	"testing"
)

func TestValidateKeyNotEmptyString(t *testing.T) {
	var m map[string]struct{}
	if validateKeyNotEmptyString(m) != true {
		t.Errorf("validateKeyNotEmptyString(%v) != true", m)
	}
	m = make(map[string]struct{})
	if validateKeyNotEmptyString(m) != true {
		t.Errorf("validateKeyNotEmptyString(%v) != true", m)
	}
	m["test"] = struct{}{}
	if validateKeyNotEmptyString(m) != true {
		t.Errorf("validateKeyNotEmptyString(%v) != true", m)
	}
	m[""] = struct{}{}
	if validateKeyNotEmptyString(m) != false {
		t.Errorf("validateKeyNotEmptyString(%v) != false", m)
	}
}

func TestValidateMapKeys(t *testing.T) {
	str := "test"
	var m map[string]struct{}
	k := make(map[string]struct{})
	if err := validateMapKeys(m, k); err != nil {
		t.Errorf("validateMapKeys(%v, %v); err != nil", m, k)
	}
	if len(k) != 0 {
		t.Error("len(k) != 0")
	}
	m = make(map[string]struct{})
	m[str] = struct{}{}
	if err := validateMapKeys(m, k); err != nil {
		t.Errorf("validateMapKeys(%v, %v); err != nil", m, k)
	}
	if len(k) != 1 {
		t.Error("len(k) != 1")
	}
	if _, ok := k[str]; !ok {
		t.Errorf("_, ok := k[\"%s\"]; !ok", str)
	}
	if err := validateMapKeys(m, k); err == nil {
		t.Errorf("validateMapKeys(%v, %v); err == nil", m, k)
	}
	if len(k) != 1 {
		t.Error("len(k) != 1")
	}
	delete(m, str)
	delete(k, str)
	m[""] = struct{}{}
	if err := validateMapKeys(m, k); err == nil {
		t.Errorf("validateMapKeys(%v, %v); err == nil", m, k)
	}
	if len(k) != 0 {
		t.Error("len(k) != 0")
	}
}

func TestIsValidPath(t *testing.T) {
	ok := []string{
		"",
		"/",
		"/test",
		"/test/test",
		"/test/",
		"/test/test/",
		"/123",
		"/123/123",
		"/Test",
		"/Test/Test",
		"/test/123",
		"/Test/123",
		"/t1st",
		"/t1st/t2st",
		"/-_%",
	}
	notOk := []string{
		"test",
		"test/",
		"test//test",
		"//",
		"//test",
		"///test",
		"123",
		"/test!§$%&/()=?",
	}
	for _, s := range ok {
		if isValidPath(s) != true {
			t.Errorf("isValidPath(\"%s\") != true", s)
		}
	}
	for _, s := range notOk {
		if isValidPath(s) != false {
			t.Errorf("isValidPath(\"%s\") != false", s)
		}
	}
}

func TestIsValidExtPath(t *testing.T) {
	ok := []string{
		"test",
		"test/test",
		"test/",
		"test/test/",
		"123",
		"123/123",
		"Test",
		"Test/Test",
		"test/123",
		"Test/123",
		"t1st",
		"t1st/t2st",
		"-_%",
	}
	notOk := []string{
		"/",
		"test//test",
		"/test/",
		"/test/test/",
		"//",
		"//test",
		"///test",
		"test!§$%&/()=?",
	}
	for _, s := range ok {
		if isValidExtPath(s) != true {
			t.Errorf("isValidPath(\"%s\") != true", s)
		}
	}
	for _, s := range notOk {
		if isValidExtPath(s) != false {
			t.Errorf("isValidPath(\"%s\") != false", s)
		}
	}
}
