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
	"reflect"
	"testing"
)

func testAddedNodes(t *testing.T, nAid string, nBid string, nCid string, n Nodes) {
	nA, ok := n[nAid]
	if !ok {
		t.Errorf("_, ok := n[\"%s\"]; !ok", nAid)
	}
	nB, ok := n[nBid]
	if !ok {
		t.Errorf("_, ok := n[\"%s\"]; !ok", nBid)
	}
	if _, ok = nB.InRef[nAid]; !ok {
		t.Errorf("_, ok = nB.InRef[\"%s\"]; !ok", nAid)
	}
	if _, ok = nB.OutRef[nAid]; ok {
		t.Errorf("_, ok = nB.OutRef[\"%s\"]; ok", nAid)
	}
	if _, ok = nA.OutRef[nBid]; !ok {
		t.Errorf("_, ok = nA.OutRef[\"%s\"]; !ok", nBid)
	}
	if _, ok = nA.InRef[nBid]; ok {
		t.Errorf("_, ok = nA.InRef[\"%s\"]; ok", nBid)
	}
	nC, ok := n[nCid]
	if !ok {
		t.Errorf("_, ok := n[\"%s\"]; !ok", nCid)
	}
	if _, ok = nC.InRef[nAid]; ok {
		t.Errorf("_, ok = nC.InRef[\"%s\"]; ok", nAid)
	}
	if _, ok = nC.OutRef[nAid]; !ok {
		t.Errorf("_, ok = nC.OutRef[\"%s\"]; !ok", nAid)
	}
	if _, ok = nA.OutRef[nCid]; ok {
		t.Errorf("_, ok = nA.OutRef[\"%s\"]; ok", nCid)
	}
	if _, ok = nA.InRef[nCid]; !ok {
		t.Errorf("_, ok = nA.InRef[\"%s\"]; !ok", nCid)
	}
}

func TestNodes_Add(t *testing.T) {
	nAid := "A"
	nBid := "B"
	nCid := "C"
	n := make(Nodes)
	// add node "A"
	n.Add(nAid, nil, nil)
	// add node "B" which requires node "A" (A -> B)
	n.Add(nBid, map[string]struct{}{nAid: {}}, nil)
	// add node "C" which is required by node "A" (A <- C)
	n.Add(nCid, nil, map[string]struct{}{nAid: {}})
	testAddedNodes(t, nAid, nBid, nCid, n)
	n = make(Nodes)
	// add node "B" which requires node "A" (A -> B)
	n.Add(nBid, map[string]struct{}{nAid: {}}, nil)
	// add node "C" which is required by node "A" (A <- C)
	n.Add(nCid, nil, map[string]struct{}{nAid: {}})
	testAddedNodes(t, nAid, nBid, nCid, n)
}

func TestNodes_Remove(t *testing.T) {
	nAid := "A"
	nBid := "B"
	nCid := "C"
	n := make(Nodes)
	// add node "A"
	n.Add(nAid, nil, nil)
	// add node "B" which requires node "A" (A -> B)
	n.Add(nBid, map[string]struct{}{nAid: {}}, nil)
	// add node "C" which is required by node "A" (A <- C)
	n.Add(nCid, nil, map[string]struct{}{nAid: {}})
	n.Remove(nAid)
	if _, ok := n[nAid]; ok {
		t.Errorf("_, ok := n[\"%s\"]; ok", nAid)
	}
	nB := n[nBid]
	if _, ok := nB.InRef[nAid]; ok {
		t.Errorf("_, ok = nB.InRef[\"%s\"]; ok", nAid)
	}
	nC := n[nCid]
	if _, ok := nC.OutRef[nAid]; ok {
		t.Errorf("_, ok = nC.OutRef[\"%s\"]; ok", nAid)
	}
}

func TestNodes_Copy(t *testing.T) {
	nAid := "A"
	nBid := "B"
	nCid := "C"
	n := make(Nodes)
	// add node "A"
	n.Add(nAid, nil, nil)
	// add node "B" which requires node "A" (A -> B)
	n.Add(nBid, map[string]struct{}{nAid: {}}, nil)
	// add node "C" which is required by node "A" (A <- C)
	n.Add(nCid, nil, map[string]struct{}{nAid: {}})
	nC := n.Copy()
	if !reflect.DeepEqual(n, nC) {
		t.Error("copy not equal to original")
	}
}
