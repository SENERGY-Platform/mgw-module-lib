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
	"github.com/SENERGY-Platform/mgw-module-lib/module"
	"github.com/SENERGY-Platform/mgw-module-lib/util"
)

func validateInputsResources(inputs map[string]module.Input, mResources map[string]util.Set[string]) error {
	for ref := range inputs {
		if ref == "" {
			return errors.New("invalid input reference")
		}
		if mResources != nil {
			if _, ok := mResources[ref]; !ok {
				return fmt.Errorf("resource '%s' not defined", ref)
			}
		} else {
			return errors.New("no resources defined")
		}
	}
	return nil
}

func validateInputsSecrets(inputs map[string]module.Input, mSecrets map[string]module.Secret) error {
	for ref := range inputs {
		if ref == "" {
			return errors.New("invalid input reference")
		}
		if mSecrets != nil {
			if _, ok := mSecrets[ref]; !ok {
				return fmt.Errorf("secret '%s' not defined", ref)
			}
		} else {
			return errors.New("no secrets defined")
		}
	}
	return nil
}

func validateInputsConfigs(inputs map[string]module.Input, mConfigs module.Configs) error {
	for ref := range inputs {
		if ref == "" {
			return errors.New("invalid input reference")
		}
		if mConfigs != nil {
			if _, ok := mConfigs[ref]; !ok {
				return fmt.Errorf("config '%s' not defined", ref)
			}
		} else {
			return errors.New("no configs defined")
		}
	}
	return nil
}

func validateInputsAndGroups(inputs map[string]module.Input, groups map[string]module.InputGroup) error {
	for _, input := range inputs {
		if input.Group != nil {
			if groups == nil {
				return errors.New("no input groups defined")
			}
			if _, ok := groups[*input.Group]; !ok {
				return fmt.Errorf("input group '%s' not defined", *input.Group)
			}
		}
	}
	return nil
}

func validateInputGroups(groups map[string]module.InputGroup) error {
	for ref, group := range groups {
		if ref == "" {
			return errors.New("invalid input group reference")
		}
		if group.Group != nil {
			if g, ok := groups[*group.Group]; !ok {
				return fmt.Errorf("input group '%s' not defined", *group.Group)
			} else {
				if g.Group != nil && *g.Group == *group.Group {
					return fmt.Errorf("input group '%s' reference cycle", *group.Group)
				}
			}
		}
	}
	return nil
}
