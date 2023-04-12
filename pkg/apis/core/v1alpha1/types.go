/*
Copyright 2023 The Kuberiter Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/apimachinery/pkg/types"
)

// kubebuilder:object:root=true
// kubebuilder:subresource:status
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Article struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ArticleSpec   `json:"spec,omitempty"`
	Status ArticleStatus `json:"status,omitempty"`
}

type ArticleSpec struct {
	Title       string      `json:"title,omitempty"`
	Description string      `json:"description,omitempty"`
	Content     string      `json:"content,omitempty"`
	Authors     []string    `json:"authors,omitempty"`
	CreatedAt   metav1.Time `json:"createdAt,omitempty"`
	UpdatedAt   metav1.Time `json:"updatedAt,omitempty"`
	Likes       int         `json:"likes,omitempty"`
}

type ArticleStatus struct {
	// Likes int `json:"likes,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ArticleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Article `json:"items"`
}

// kubebuilder:object:root=true
// kubebuilder:subresource:status
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UserSpec   `json:"spec,omitempty"`
	Status UserStatus `json:"status,omitempty"`
}

type UserSpec struct {
	Gender uint8 `json:"gender,omitempty"`
}

type UserStatus struct {
	Followers []string `json:"followers,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []User `json:"items"`
}

// kubebuilder:object:root=true
// kubebuilder:subresource:status
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Tag struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec  TagSpec   `json:"spec,omitempty"`
	State TagStatus `json:"status,omitempty"`
}

type TagSpec struct {
	Owner string `json:"owner,omitempty"`
}

type TagStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Tag `json:"items"`
}
