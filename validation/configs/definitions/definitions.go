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

package definitions

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/SENERGY-Platform/mgw-module-lib/validation/configs/validators"
	"io/fs"
	"os"
	"path"
	"regexp"
	"strings"
)

var Definitions map[string]ConfigDefinition

//go:embed *.json
var defFiles embed.FS

func init() {
	entries, err := defFiles.ReadDir(".")
	if err != nil {
		fmt.Println("reading definitions failed: ", err)
		os.Exit(1)
	}
	Definitions = make(map[string]ConfigDefinition)
	for _, entry := range entries {
		if !entry.IsDir() {
			file, err := defFiles.Open(entry.Name())
			if err != nil {
				fmt.Println("opening definition failed: ", err)
				os.Exit(1)
			}
			definition, err := loadDefinition(file)
			if err != nil {
				fmt.Println("decoding definition failed: ", err)
				os.Exit(1)
			}
			Definitions[strings.TrimSuffix(entry.Name(), path.Ext(entry.Name()))] = definition
		}
	}
	if err = validateDefs(Definitions, validators.Validators); err != nil {
		fmt.Println("validating definitions failed: ", err)
		os.Exit(1)
	}
}

func loadDefinition(file fs.File) (ConfigDefinition, error) {
	defer file.Close()
	var definition ConfigDefinition
	if err := json.NewDecoder(file).Decode(&definition); err != nil {
		return ConfigDefinition{}, err
	}
	return definition, nil
}

func validateDefs(configDefs map[string]ConfigDefinition, validators map[string]validators.Validator) error {
	// missing tests and needs to be cleaned up
	for ref, cDef := range configDefs {
		if len(cDef.DataType) == 0 {
			return fmt.Errorf("config definition '%s' missing data type", ref)
		}
		for key, cDefOpt := range cDef.Options {
			if !cDefOpt.Inherit && len(cDefOpt.DataType) == 0 {
				return fmt.Errorf("config definition '%s' option '%s' missing data type", ref, key)
			}
		}
		if len(validators) > 0 {
			for _, validator := range cDef.Validators {
				if _, ok := validators[validator.Name]; !ok {
					return fmt.Errorf("config definition '%s' unknown validator '%s'", ref, validator.Name)
				}
				for key, param := range validator.Parameter {
					if param.Ref == nil && param.Value == nil {
						return fmt.Errorf("config definition '%s' validator '%s' parameter '%s' missing input", ref, validator.Name, key)
					}
					if param.Ref != nil {
						re := regexp.MustCompile(`^options\.[a-z0-9A-Z_]+$|^value$`)
						if !re.MatchString(*param.Ref) {
							return fmt.Errorf("config definition '%s' validator '%s' parameter '%s' invalid refrence '%s'", ref, validator.Name, key, *param.Ref)
						}
					}
				}
			}
		}
	}
	return nil
}
