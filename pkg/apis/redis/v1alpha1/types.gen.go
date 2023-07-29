// Copyright Aeraki Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by kubetype-gen. DO NOT EDIT.

package v1alpha1

import (
	redisv1alpha1 "github.com/aeraki-mesh/api/redis/v1alpha1"
	metav1alpha1 "istio.io/api/meta/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RedisDestination defines policies that apply to redis traffic intended for a redis service
//
// <!-- crd generation tags
// +cue-gen:RedisDestination:groupName:redis.aeraki.io
// +cue-gen:RedisDestination:version:v1alpha1
// +cue-gen:RedisDestination:storageVersion
// +cue-gen:RedisDestination:subresource:status
// +cue-gen:RedisDestination:scope:Namespaced
// +cue-gen:RedisDestination:resource:categories=redis-aeraki-io,shortNames=rd
// +cue-gen:RedisDestination:printerColumn:name=Host,type=string,JSONPath=.spec.host,description="The name of a service from the service registry"
// +cue-gen:RedisDestination:printerColumn:name=Age,type=date,JSONPath=.metadata.creationTimestamp,description="CreationTimestamp is a timestamp
// representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations.
// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
// Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
// +cue-gen:RedisDestination:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=redis.aeraki.io/v1alpha1
// +genclient
// +k8s:deepcopy-gen=true
// -->
type RedisDestination struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec redisv1alpha1.RedisDestination `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	Status metav1alpha1.IstioStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RedisDestinationList is a collection of RedisDestinations.
type RedisDestinationList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []*RedisDestination `json:"items" protobuf:"bytes,2,rep,name=items"`
}

//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RedisService provide a way to config redis service in service mesh.
//
// <!-- crd generation tags
// +cue-gen:RedisService:groupName:redis.aeraki.io
// +cue-gen:RedisService:version:v1alpha1
// +cue-gen:RedisService:storageVersion
// +cue-gen:RedisService:subresource:status
// +cue-gen:RedisService:scope:Namespaced
// +cue-gen:RedisService:resource:categories=redis-aeraki-io,shortNames=rsvc
// +cue-gen:RedisService:printerColumn:name=Hosts,type=string,JSONPath=.spec.hosts,description="The destination hosts to which traffic is being sent"
// +cue-gen:RedisService:printerColumn:name=Age,type=date,JSONPath=.metadata.creationTimestamp,description="CreationTimestamp is a timestamp
// representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations.
// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
// Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
// +cue-gen:RedisService:preserveUnknownFields:true
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=redis.aeraki.io/v1alpha1
// +genclient
// +k8s:deepcopy-gen=true
// -->
type RedisService struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec redisv1alpha1.RedisService `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	Status metav1alpha1.IstioStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RedisServiceList is a collection of RedisServices.
type RedisServiceList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []*RedisService `json:"items" protobuf:"bytes,2,rep,name=items"`
}
