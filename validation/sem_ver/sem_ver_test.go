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

package sem_ver

import (
	"testing"
)

func TestInSemVerRange(t *testing.T) {
	ok := [][2]string{
		{">v1.0.0;<v2.0.0", "v1.1.0"},
		{">=v1.0.0;<v2.0.0", "v1.1.0"},
		{">=v1.0.0;<v2.0.0", "v1.0.0"},
		{">v1.0.0;<=v2.0.0", "v1.1.0"},
		{">v1.0.0;<=v2.0.0", "v2.0.0"},
		{">=v1.0.0;<=v2.0.0", "v1.1.0"},
		{">=v1.0.0;<=v2.0.0", "v1.0.0"},
		{">=v1.0.0;<=v2.0.0", "v2.0.0"},
		{">v1.0;<v2.0", "v1.1"},
		{">v1;<v2", "v1.1"},
	}
	notOk := [][2]string{
		{">v1.0.0;<v2.0.0", "v0.0.1"},
		{">v1.0.0;<v2.0.0", "v2.1.0"},
		{">v1.0.0;<v2.0.0", "v2.0.0"},
		{">v1.0.0;<v2.0.0", "v1.0.0"},
		{">=v1.0.0;<v2.0.0", "v0.0.1"},
		{">=v1.0.0;<v2.0.0", "v2.1.0"},
		{">=v1.0.0;<v2.0.0", "v2.0.0"},
		{">v1.0.0;<=v2.0.0", "v0.0.1"},
		{">v1.0.0;<=v2.0.0", "v2.1.0"},
		{">v1.0.0;<=v2.0.0", "v1.0.0"},
		{">=v1.0.0;<=v2.0.0", "v0.0.1"},
		{">=v1.0.0;<=v2.0.0", "v2.1.0"},
	}
	er := [][2]string{
		{">v1.0.0;<v2.0.0", "1.1.0"},
		{">1.0.0;<v2.0.0", "v1.1.0"},
		{">v1.0.0;<2.0.0", "v1.1.0"},
		{">1.0.0;<2.0.0", "v1.1.0"},
		{">1.0.0;<2.0.0", "1.1.0"},
		{"v1.0.0;<v2.0.0", "v1.1.0"},
		{">v1.0.0;v2.0.0", "v1.1.0"},
		{"v1.0.0;v2.0.0", "v1.1.0"},
		{"=v1.0.0;<v2.0.0", "v1.1.0"},
		{">v1.0.0;=v2.0.0", "v1.1.0"},
		{"=v1.0.0;=v2.0.0", "v1.1.0"},
		{">v1.0,0;<v2.0.0", "v1.1.0"},
		{">v1.0.0;<v2.0,0", "v1.1.0"},
		{">v1.0.0;<v2.0.0", "v1.1,0"},
		{"test", "v1.1.0"},
		{">v1.0.0;<v2.0.0", "test"},
	}
	for _, v := range ok {
		k, err := InSemVerRange(v[0], v[1])
		if k != true {
			t.Errorf("InSemVerRange(\"%s\", \"%s\") != true", v[0], v[1])
		}
		if err != nil {
			t.Errorf("InSemVerRange(\"%s\", \"%s\"); err != nil", v[0], v[1])
		}
	}
	for _, v := range notOk {
		k, err := InSemVerRange(v[0], v[1])
		if k != false {
			t.Errorf("InSemVerRange(\"%s\", \"%s\") != false", v[0], v[1])
		}
		if err != nil {
			t.Errorf("InSemVerRange(\"%s\", \"%s\"); err != nil", v[0], v[1])
		}
	}
	for _, v := range er {
		k, err := InSemVerRange(v[0], v[1])
		if k != false {
			t.Errorf("InSemVerRange(\"%s\", \"%s\") != false", v[0], v[1])
		}
		if err == nil {
			t.Errorf("InSemVerRange(\"%s\", \"%s\"); err == nil", v[0], v[1])
		}
	}
}
