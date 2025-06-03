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

type Nodes map[string]*node

func (n Nodes) Add(id string, inRef map[string]struct{}, outRef map[string]struct{}) {
	nde, ok := n[id]
	if !ok {
		nde = newNode()
		n[id] = nde
	}
	for ref := range inRef {
		nde.AddInRef(ref)
		reqNde, k := n[ref]
		if !k {
			reqNde = newNode()
			n[ref] = reqNde
		}
		reqNde.AddOutRef(id)
	}
	for ref := range outRef {
		nde.AddOutRef(ref)
		reqByNde, k := n[ref]
		if !k {
			reqByNde = newNode()
			n[ref] = reqByNde
		}
		reqByNde.AddInRef(id)
	}
}

func (n Nodes) Remove(id string) {
	if nde, ok := n[id]; ok {
		for ref := range nde.InRef {
			if iNde, k := n[ref]; k {
				iNde.RemoveOutRef(id)
			}
		}
		for ref := range nde.OutRef {
			if oNde, k := n[ref]; k {
				oNde.RemoveInRef(id)
			}
		}
		delete(n, id)
	}
}

func (n Nodes) Copy() Nodes {
	cNodes := make(Nodes)
	for id, nde := range n {
		cNde := newNode()
		for ref := range nde.InRef {
			cNde.AddInRef(ref)
		}
		for ref := range nde.OutRef {
			cNde.AddOutRef(ref)
		}
		cNodes[id] = cNde
	}
	return cNodes
}
