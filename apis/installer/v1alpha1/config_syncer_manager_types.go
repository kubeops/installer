/*
Copyright AppsCode Inc. and Contributors

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
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kmodules.xyz/resource-metadata/apis/shared"
)

const (
	ResourceKindConfigSyncerManager = "ConfigSyncerManager"
	ResourceConfigSyncerManager     = "configsyncermanager"
	ResourceConfigSyncerManagers    = "configsyncermanagers"
)

// ConfigSyncerManager defines the schema for the config-syncer OCM AddOn manager installer chart.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=configsyncermanagers,singular=configsyncermanager,categories={kubeops,appscode}
type ConfigSyncerManager struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigSyncerManagerSpec `json:"spec,omitempty"`
}

// ConfigSyncerManagerSpec is the schema for the config-syncer-manager values file.
type ConfigSyncerManagerSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string `json:"fullnameOverride"`

	RegistryFQDN string `json:"registryFQDN"`
	Image        string `json:"image"`
	//+optional
	Tag             string `json:"tag"`
	ImagePullPolicy string `json:"imagePullPolicy"`

	// ClusterName is recorded as origin cluster on objects synced from the hub.
	ClusterName string `json:"clusterName"`

	Placement ConfigSyncerManagerPlacement `json:"placement"`

	//+optional
	KubeconfigSecretName string `json:"kubeconfigSecretName"`

	SecurityContext *core.SecurityContext `json:"securityContext"`
	//+optional
	EnvFrom []core.EnvFromSource `json:"envFrom"`
	//+optional
	Env []core.EnvVar `json:"env"`
	//+optional
	Monitoring Monitoring `json:"monitoring"`
	//+optional
	Distro shared.DistroSpec `json:"distro"`
}

type ConfigSyncerManagerPlacement struct {
	Create bool   `json:"create"`
	Name   string `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConfigSyncerManagerList is a list of ConfigSyncerManagers
type ConfigSyncerManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ConfigSyncerManager CRD objects
	Items []ConfigSyncerManager `json:"items,omitempty"`
}
