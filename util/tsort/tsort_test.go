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

func checkOrder(o []string, refO []string) bool {
	lPos := -1
	for i := len(refO) - 1; i >= 0; i-- {
		for i2, s := range o {
			if s == refO[i] {
				if lPos == -1 {
					lPos = i2
				} else {
					if i2 > lPos {
						return false
					}
				}
			}
		}
	}
	return true
}

func TestGetTopOrder(t *testing.T) {
	nAid := "A"
	nBid := "B"
	nCid := "C"
	nDid := "D"
	order := []string{nCid, nAid, nBid} // (C -> A -> B)
	n := make(Nodes)
	// add node "A"
	n.Add(nAid, nil, nil)
	if o, err := GetTopOrder(n); err != nil {
		t.Error("err != nil")
	} else if o[0] != nAid {
		t.Errorf("o[0] != \"%s\"", nAid)
	}
	n = make(Nodes)
	// add node "A" which requires node "C" (C -> A)
	n.Add(nAid, map[string]struct{}{nCid: {}}, nil)
	// add node "B" which requires node "A" (A -> B)
	n.Add(nBid, map[string]struct{}{nAid: {}}, nil)
	if o, err := GetTopOrder(n); err != nil {
		t.Error("err != nil")
	} else if !reflect.DeepEqual(order, o) {
		t.Errorf("%v != %v", order, o)
	}
	n = make(Nodes)
	// add node "A" which is required by node "B" (B <- A)
	n.Add(nAid, nil, map[string]struct{}{nBid: {}})
	// add node "C" which is required by node "A" (A <- C)
	n.Add(nCid, nil, map[string]struct{}{nAid: {}})
	if o, err := GetTopOrder(n); err != nil {
		t.Error("err != nil")
	} else if !reflect.DeepEqual(order, o) {
		t.Errorf("%v != %v", order, o)
	}
	n = make(Nodes)
	// add node "B" which requires node "A" (A -> B)
	n.Add(nBid, map[string]struct{}{nAid: {}}, nil)
	// add node "C" which is required by node "A" (A <- C)
	n.Add(nCid, nil, map[string]struct{}{nAid: {}})
	if o, err := GetTopOrder(n); err != nil {
		t.Error("err != nil")
	} else if !reflect.DeepEqual(order, o) {
		t.Errorf("%v != %v", order, o)
	}
	n = make(Nodes)
	// add node "A" which is required by node "B" (B <- A)
	n.Add(nAid, nil, map[string]struct{}{nBid: {}})
	// add node "A" which requires node "C" (C -> A)
	n.Add(nAid, map[string]struct{}{nCid: {}}, nil)
	if o, err := GetTopOrder(n); err != nil {
		t.Error("err != nil")
	} else if !reflect.DeepEqual(order, o) {
		t.Errorf("%v != %v", order, o)
	}
	n = make(Nodes)
	// add node "A" which requires node "C" (C -> A) and is required by node "B" (B <- A)
	n.Add(nAid, map[string]struct{}{nCid: {}}, map[string]struct{}{nBid: {}})
	if o, err := GetTopOrder(n); err != nil {
		t.Error("err != nil")
	} else if !reflect.DeepEqual(order, o) {
		t.Errorf("%v != %v", order, o)
	}
	n = make(Nodes)
	// add node "A" which requires node "C" (C -> A)
	n.Add(nAid, map[string]struct{}{nCid: {}}, nil)
	// add node "B" which requires node "A" (A -> B)
	n.Add(nBid, map[string]struct{}{nAid: {}}, nil)
	// add node "D"
	n.Add(nDid, nil, nil)
	if o, err := GetTopOrder(n); err != nil {
		t.Error("err != nil")
	} else {
		var ok bool
		for _, s := range o {
			ok = s == nDid
			if ok {
				break
			}
		}
		if !ok {
			t.Errorf("\"%s\" not in %v", nDid, o)
		}
		if !checkOrder(o, order) {
			t.Errorf("%v != %v", order, o)
		}
	}
	n = make(Nodes)
	// add node "A" which requires node "A" (A -> A)
	n.Add(nAid, map[string]struct{}{nAid: {}}, nil)
	if _, err := GetTopOrder(n); err == nil {
		t.Error("err == nil")
	}
	n = make(Nodes)
	// add node "A" which is required by node "A" (A <- A)
	n.Add(nAid, nil, map[string]struct{}{nAid: {}})
	if _, err := GetTopOrder(n); err == nil {
		t.Error("err == nil")
	}
	n = make(Nodes)
	// add node "A" which requires node "B" (B -> A)
	n.Add(nAid, map[string]struct{}{nBid: {}}, nil)
	// add node "B" which requires node "A" (A -> B)
	n.Add(nBid, map[string]struct{}{nAid: {}}, nil)
	if _, err := GetTopOrder(n); err == nil {
		t.Error("err == nil")
	}
	n = make(Nodes)
	// add node "A" which requires node "C" and "D" (C -> A, D -> A)
	n.Add(nAid, map[string]struct{}{nCid: {}, nDid: {}}, nil)
	// add node "B" which requires node "A" (A -> B)
	n.Add(nBid, map[string]struct{}{nAid: {}}, nil)
	// add node "D" which requires node "B" (B -> D)
	n.Add(nDid, map[string]struct{}{nBid: {}}, nil)
	if _, err := GetTopOrder(n); err == nil {
		t.Error("err == nil")
	}
}
