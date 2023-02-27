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

type node struct {
	InRef  map[string]struct{} // inbound references: this node requires node n (n -> this), e.g. B requires A (A -> B)
	OutRef map[string]struct{} // outbound references: this node is required by node n (n <- this), e.g. B is required by A (A <- B)
}

func newNode() *node {
	return &node{
		InRef:  make(map[string]struct{}),
		OutRef: make(map[string]struct{}),
	}
}

func (n *node) AddInRef(id string) {
	n.InRef[id] = struct{}{}
}

func (n *node) AddOutRef(id string) {
	n.OutRef[id] = struct{}{}
}

func (n *node) RemoveInRef(id string) {
	delete(n.InRef, id)
}

func (n *node) RemoveOutRef(id string) {
	delete(n.OutRef, id)
}
