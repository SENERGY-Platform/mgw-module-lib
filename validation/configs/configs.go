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

package configs

import (
	"errors"
	"fmt"
	"github.com/SENERGY-Platform/mgw-module-lib/model"
	"github.com/SENERGY-Platform/mgw-module-lib/validation/configs/definitions"
	"github.com/SENERGY-Platform/mgw-module-lib/validation/configs/validators"
	"strings"
)

func ValidateBase(cType string, cTypeOpts model.ConfigTypeOptions, dataType model.DataType) error {
	cDef, ok := definitions.Definitions[cType]
	if !ok {
		return fmt.Errorf("config type '%s' not defined", cType)
	}
	return vltBase(cDef, cTypeOpts, dataType)
}

func ValidateTypeOptions(cType string, cTypeOpts model.ConfigTypeOptions) error {
	cDef, ok := definitions.Definitions[cType]
	if !ok {
		return fmt.Errorf("config type '%s' not defined", cType)
	}
	return vltTypeOpts(cDef.Validators, cTypeOpts, validators.Validators)
}

func ValidateValue(cType string, cTypeOpts model.ConfigTypeOptions, value any, isSlice bool, dataType model.DataType) error {
	cDef, ok := definitions.Definitions[cType]
	if !ok {
		return fmt.Errorf("config type '%s' not defined", cType)
	}
	if isSlice {
		switch dataType {
		case model.StringType:
			return vltValSlice[string](cDef.Validators, cTypeOpts, validators.Validators, value)
		case model.BoolType:
			return vltValSlice[bool](cDef.Validators, cTypeOpts, validators.Validators, value)
		case model.Int64Type:
			return vltValSlice[int64](cDef.Validators, cTypeOpts, validators.Validators, value)
		case model.Float64Type:
			return vltValSlice[float64](cDef.Validators, cTypeOpts, validators.Validators, value)
		default:
			return fmt.Errorf("unknown data type '%s'", dataType)
		}
	} else {
		return vltValue(cDef.Validators, cTypeOpts, validators.Validators, value)
	}
}

func ValidateValInOpt(cOpt any, value any, isSlice bool, dataType model.DataType) (err error) {
	var ok bool
	switch dataType {
	case model.StringType:
		ok, err = vltValInOptF[string](isSlice)(value, cOpt)
	case model.BoolType:
		ok, err = vltValInOptF[bool](isSlice)(value, cOpt)
	case model.Int64Type:
		ok, err = vltValInOptF[int64](isSlice)(value, cOpt)
	case model.Float64Type:
		ok, err = vltValInOptF[float64](isSlice)(value, cOpt)
	default:
		err = fmt.Errorf("unknown data type '%s'", dataType)
	}
	if !ok {
		err = errors.New("value not allowed")
	}
	return
}

func vltValInOptF[T comparable](isSl bool) func(any, any) (bool, error) {
	if isSl {
		return vltValSlInOpt[T]
	} else {
		return vltValInOpt[T]
	}
}

func vltValInOpt[T comparable](val any, opt any) (bool, error) {
	v, ok := val.(T)
	if !ok {
		return false, fmt.Errorf("invalid data type '%T'", val)
	}
	o, ok := opt.([]T)
	if !ok {
		return false, fmt.Errorf("invalid data type '%T'", opt)
	}
	for _, e := range o {
		if v == e {
			return true, nil
		}
	}
	return false, nil
}

func vltValSlInOpt[T comparable](val any, opt any) (bool, error) {
	vSl, ok := val.([]T)
	if !ok {
		return false, fmt.Errorf("invalid data type '%T'", val)
	}
	o, ok := opt.([]T)
	if !ok {
		return false, fmt.Errorf("invalid data type '%T'", opt)
	}
	var k bool
	for _, v := range vSl {
		k = false
		for _, e := range o {
			if v == e {
				k = true
				break
			}
		}
		if !k {
			break
		}
	}
	return k, nil
}

func vltValSlice[T any](cDefVlts []definitions.ConfigDefinitionValidator, cTypeOpts model.ConfigTypeOptions, validators map[string]validators.Validator, value any) error {
	valSl, ok := value.([]T)
	if !ok {
		return fmt.Errorf("invlaid data type: %T != %T", value, *new(T))
	}
	for _, val := range valSl {
		if err := vltValue(cDefVlts, cTypeOpts, validators, val); err != nil {
			return err
		}
	}
	return nil
}

func vltBase(cDef definitions.ConfigDefinition, cTypeOpts model.ConfigTypeOptions, dataType model.DataType) error {
	if _, ok := cDef.DataType[dataType]; !ok {
		return fmt.Errorf("data type '%s' not supported", dataType)
	}
	if len(cTypeOpts) > 0 && len(cDef.Options) == 0 {
		return fmt.Errorf("options not supported")
	}
	for name := range cTypeOpts {
		if _, ok := cDef.Options[name]; !ok {
			return fmt.Errorf("option '%s' not supported", name)
		}
	}
	for name, cDefO := range cDef.Options {
		if cTypeO, ok := cTypeOpts[name]; ok {
			if cDefO.Inherit {
				if cTypeO.DataType != dataType {
					return fmt.Errorf("data type '%s' not supported by option '%s'", cTypeO.DataType, name)
				}
			} else {
				if _, ok := cDefO.DataType[cTypeO.DataType]; !ok {
					return fmt.Errorf("data type '%s' not supported by option '%s'", cTypeO.DataType, name)
				}
			}
		} else if cDefO.Required {
			return fmt.Errorf("option '%s' required", name)
		}
	}
	return nil
}

func genVltOptParams(cDefVltParams map[string]definitions.ConfigDefinitionValidatorParam, cTypeOpts model.ConfigTypeOptions) map[string]any {
	vp := make(map[string]any)
	for name, cDefVP := range cDefVltParams {
		if cDefVP.Ref != nil {
			if *cDefVP.Ref == "value" {
				if cDefVP.Value != nil {
					vp[name] = cDefVP.Value
				} else {
					vp = nil
					break
				}
			} else {
				cTypeOName := strings.Split(*cDefVP.Ref, ".")[1]
				if cTypeO, ok := cTypeOpts[cTypeOName]; ok {
					vp[name] = cTypeO.Value
				} else {
					if cDefVP.Value != nil {
						vp[name] = cDefVP.Value
					} else {
						vp = nil
						break
					}
				}
			}
		} else {
			vp[name] = cDefVP.Value
		}
	}
	return vp
}

func vltTypeOpts(cDefVlts []definitions.ConfigDefinitionValidator, cTypeOpts model.ConfigTypeOptions, validators map[string]validators.Validator) error {
	for _, cDefVlt := range cDefVlts {
		p := genVltOptParams(cDefVlt.Parameter, cTypeOpts)
		if len(p) > 0 {
			vFunc, ok := validators[cDefVlt.Name]
			if !ok {
				return fmt.Errorf("validator '%s' not defined", cDefVlt.Name)
			}
			err := vFunc(p)
			if err != nil {
				return fmt.Errorf("validator '%s' returned with: %s", cDefVlt.Name, err)
			}
		}
	}
	return nil
}

func genVltValParams(cDefVltParams map[string]definitions.ConfigDefinitionValidatorParam, cTypeOpts model.ConfigTypeOptions, value any) map[string]any {
	vp := make(map[string]any)
	for name, cDefVP := range cDefVltParams {
		if cDefVP.Ref != nil {
			if *cDefVP.Ref == "value" {
				if value != nil {
					vp[name] = value
				} else {
					if cDefVP.Value != nil {
						vp[name] = cDefVP.Value
					} else {
						vp = nil
						break
					}
				}
			} else {
				cTypeOName := strings.Split(*cDefVP.Ref, ".")[1]
				if cTypeO, ok := cTypeOpts[cTypeOName]; ok {
					vp[name] = cTypeO.Value
				} else {
					if cDefVP.Value != nil {
						vp[name] = cDefVP.Value
					} else {
						vp = nil
						break
					}
				}
			}
		} else {
			vp[name] = cDefVP.Value
		}
	}
	return vp
}

func vltValue(cDefVlts []definitions.ConfigDefinitionValidator, cTypeOpts model.ConfigTypeOptions, validators map[string]validators.Validator, value any) error {
	for _, cDefVlt := range cDefVlts {
		p := genVltValParams(cDefVlt.Parameter, cTypeOpts, value)
		if len(p) > 0 {
			vFunc, ok := validators[cDefVlt.Name]
			if !ok {
				return fmt.Errorf("validator '%s' not defined", cDefVlt.Name)
			}
			err := vFunc(p)
			if err != nil {
				return fmt.Errorf("validator '%s' returned with: %s", cDefVlt.Name, err)
			}
		}
	}
	return nil
}
