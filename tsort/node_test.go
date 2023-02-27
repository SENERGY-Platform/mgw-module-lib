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

import "testing"

func TestNewNode(t *testing.T) {
	n := newNode()
	if n.InRef == nil {
		t.Error("n.InRef == nil")
	}
	if len(n.InRef) != 0 {
		t.Error("len(n.InRef) != 0")
	}
	if n.OutRef == nil {
		t.Error("n.OutRef == nil")
	}
	if len(n.OutRef) != 0 {
		t.Error("len(n.OutRef) != 0")
	}
}

func TestNode_AddInRef(t *testing.T) {
	str := "test"
	n := node{InRef: make(map[string]struct{})}
	n.AddInRef(str)
	if _, ok := n.InRef[str]; !ok {
		t.Errorf("_, ok := n.InRef[\"%s\"]; !ok", str)
	}
}

func TestNode_AddOutRef(t *testing.T) {
	str := "test"
	n := node{OutRef: make(map[string]struct{})}
	n.AddOutRef(str)
	if _, ok := n.OutRef[str]; !ok {
		t.Errorf("_, ok := n.OutRef[\"%s\"]; !ok", str)
	}
}

func TestNode_RemoveInRef(t *testing.T) {
	str := "test"
	n := node{InRef: map[string]struct{}{str: {}}}
	n.RemoveInRef(str)
	if _, ok := n.InRef[str]; ok {
		t.Errorf("_, ok := n.InRef[\"%s\"]; ok", str)
	}
}

func TestNode_RemoveOutRef(t *testing.T) {
	str := "test"
	n := node{OutRef: map[string]struct{}{str: {}}}
	n.RemoveOutRef(str)
	if _, ok := n.OutRef[str]; ok {
		t.Errorf("_, ok := n.OutRef[\"%s\"]; ok", str)
	}
}
