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

package tsort

import (
	"fmt"
	"sort"
	"strings"
)

func GetTopOrder(nodes Nodes) ([]string, error) {
	var stack []string
	var err error
	stack, err = topSort(nodes.Copy(), stack)
	return stack, err
}

func topSort(nodes Nodes, stack []string) ([]string, error) {
	del := false
	for ndeId, nde := range nodes {
		if len(nde.InRef) == 0 && len(nde.OutRef) > 0 {
			stack = append(stack, ndeId)
			del = true
			break
		}
	}
	if del {
		nodes.Remove(stack[len(stack)-1])
		return topSort(nodes, stack)
	} else {
		var errStr []string
		for ndeId, nde := range nodes {
			if len(nde.InRef) != 0 {
				errStr = append(errStr, fmt.Sprintf("[%s->%s->%s]", keysToStr(nde.InRef), ndeId, keysToStr(nde.OutRef)))
			} else {
				stack = append(stack, ndeId)
			}
		}
		if len(errStr) > 0 {
			return nil, fmt.Errorf("non acyclic graph: %v", strings.Join(errStr, " "))
		}
	}
	return stack, nil
}

func keysToStr(m map[string]struct{}) string {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return strings.Join(keys, ",")
}
