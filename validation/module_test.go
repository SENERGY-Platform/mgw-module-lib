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

import "testing"

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
