/*
 * Copyright 2022 InfAI (CC SES)
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
	"fmt"
	"golang.org/x/mod/semver"
	"strings"
)

func InSemVerRange(r string, v string) (bool, error) {
	opr, ver, err := semVerRangeParse(r)
	if err != nil {
		return false, err
	}
	if !semver.IsValid(v) {
		return false, fmt.Errorf("format '%s' invalid", v)
	}
	ok := semVerRangeCheck(opr[0], ver[0], v)
	if ok && len(opr) > 1 && len(ver) > 1 {
		ok = semVerRangeCheck(opr[1], ver[1], v)
	}
	return ok, nil
}

func IsValidSemVer(s string) bool {
	return semver.IsValid(s)
}

func ValidateSemVerRange(s string) error {
	_, _, err := semVerRangeParse(s)
	return err
}

func CompareSemVer(v, w string) (int, error) {
	for _, s := range [2]string{w, v} {
		if !semver.IsValid(s) {
			return 0, fmt.Errorf("format '%s' invalid", s)
		}
	}
	return semver.Compare(v, w), nil
}

func semVerRangeCheck(o string, w, v string) bool {
	res := semver.Compare(v, w)
	switch res {
	case -1:
		if o == LessEqual || o == Less {
			return true
		}
	case 0:
		if o == Equal || o == LessEqual || o == GreaterEqual {
			return true
		}
	case 1:
		if o == GreaterEqual || o == Greater {
			return true
		}
	}
	return false
}

func semVerRangeParsePart(s string) (string, string, error) {
	pos := strings.Index(s, "v")
	if pos < 1 || pos > 2 {
		return "", "", fmt.Errorf("format '%s' invalid", s)
	}
	if _, ok := OperatorMap[(s[:pos])]; ok {
		if !semver.IsValid(s[pos:]) {
			return "", "", fmt.Errorf("format '%s' invalid", s[pos:])
		}
		return s[:pos], s[pos:], nil
	}
	return "", "", fmt.Errorf("operator format '%s' invalid", s[:pos])
}

func semVerRangeParse(s string) (opr []string, ver []string, err error) {
	partsStr := strings.Split(s, ";")
	if len(partsStr) > 2 {
		err = fmt.Errorf("format '%s' invalid", s)
		return
	}
	for _, p := range partsStr {
		o, v, e := semVerRangeParsePart(p)
		if e != nil {
			err = e
			return
		}
		opr = append(opr, o)
		ver = append(ver, v)
	}
	if len(opr) > 1 && len(ver) > 1 {
		if opr[0] == Less || opr[0] == LessEqual || opr[0] == Equal {
			err = fmt.Errorf("operator order '%s' + '%s' invalid", opr[0], opr[1])
			return
		}
		if opr[1] == Greater || opr[1] == GreaterEqual || opr[1] == Equal {
			err = fmt.Errorf("operator order '%s' - '%s' invalid", opr[0], opr[1])
			return
		}
		if semver.Compare(ver[0], ver[1]) > -1 {
			err = fmt.Errorf("version order '%s' - '%s' invalid", ver[0], ver[1])
			return
		}
	}
	return
}
