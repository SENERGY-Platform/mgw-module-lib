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
	"errors"
	"fmt"
	"regexp"
)

func validateKeyNotEmptyString[T any](m map[string]T) bool {
	for ref := range m {
		if ref == "" {
			return false
		}
	}
	return true
}

func validateMapKeys[T any](m map[string]T, keys map[string]struct{}) error {
	for k := range m {
		if k == "" {
			return errors.New("empty")
		}
		if _, ok := keys[k]; ok {
			return fmt.Errorf("duplicate '%s'", k)
		}
		keys[k] = struct{}{}
	}
	return nil
}

func isValidPath(s string) bool {
	re := regexp.MustCompile(`^\/(?:[a-zA-Z0-9-_%]+)+(?:\/[a-zA-Z0-9-_%]+)*$`)
	return re.MatchString(s)
}

func isValidExtPath(s string) bool {
	re := regexp.MustCompile(`^(?:[a-zA-Z0-9-_%]+)+(?:\/[a-zA-Z0-9-_%]+)*$`)
	return re.MatchString(s)
}
