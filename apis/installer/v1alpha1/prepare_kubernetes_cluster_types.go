/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

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
)

const (
	ResourceKindPrepareKubernetesCluster = "PrepareKubernetesCluster"
	ResourcePrepareKubernetesCluster     = "preparekubernetescluster"
	ResourcePrepareKubernetesClusters    = "preparekubernetesclusters"
)

// PrepareKubernetesCluster defines the schama for Kubernetes Cluster prepaaration process.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=preparekubernetesclusters,singular=preparekubernetescluster,categories={appscode}
type PrepareKubernetesCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrepareKubernetesClusterSpec `json:"spec,omitempty"`
}

type PKCImageRef struct {
	Repository string `json:"repository"`
	PullPolicy string `json:"pullPolicy"`
	Tag        string `json:"tag"`
	// Security options the pod should run with.
	// +optional
	SecurityContext *core.SecurityContext `json:"securityContext"`
}

type PKCCleanerRef struct {
	PKCImageRef `json:",inline"`
	Skip        bool `json:"skip"`
}

// PrepareKubernetesClusterSpec is the spec for redis version
type PrepareKubernetesClusterSpec struct {
	Preparer           PKCImageRef               `json:"preparer"`
	Cleaner            PKCCleanerRef             `json:"cleaner"`
	ImagePullSecrets   []string                  `json:"imagePullSecrets"`
	NameOverride       string                    `json:"nameOverride"`
	FullnameOverride   string                    `json:"fullnameOverride"`
	ServiceAccount     ServiceAccountSpec        `json:"serviceAccount"`
	PodAnnotations     map[string]string         `json:"podAnnotations"`
	PodSecurityContext *core.PodSecurityContext  `json:"podSecurityContext"`
	Resources          core.ResourceRequirements `json:"resources"`
	NodeSelector       map[string]string         `json:"nodeSelector"`
	Tolerations        []core.Toleration         `json:"tolerations"`
	Affinity           *core.Affinity            `json:"affinity"`
	Node               NodeConfiguration         `json:"node"`
}

type NodeConfiguration struct {
	Features []string      `json:"features"`
	Sysctls  []core.Sysctl `json:"sysctls"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PrepareKubernetesClusterList is a list of PrepareKubernetesClusters
type PrepareKubernetesClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PrepareKubernetesCluster CRD objects
	Items []PrepareKubernetesCluster `json:"items,omitempty"`
}
