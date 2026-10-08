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
	ResourceKindDrControlplaneAgent = "DrControlplaneAgent"
	ResourceDrControlplaneAgent     = "drcontrolplaneagent"
	ResourceDrControlplaneAgents    = "drcontrolplaneagents"
)

// DrControlplaneAgent defines the schema for the dr-controlplane-agent OCM addon chart.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=drcontrolplaneagents,singular=drcontrolplaneagent,categories={kubeops,appscode}
type DrControlplaneAgent struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DrControlplaneAgentSpec `json:"spec,omitempty"`
}

// DrControlplaneAgentSpec is the schema for the dr-controlplane-agent chart values file
type DrControlplaneAgentSpec struct {
	//+optional
	ClusterName string `json:"clusterName"`
	Namespace   string `json:"namespace"`
	//+optional
	AddonInstallNamespace string              `json:"addonInstallNamespace"`
	CreateNamespace       bool                `json:"createNamespace"`
	Image                 DrControlplaneImage `json:"image"`
	//+optional
	ImagePullSecrets []core.LocalObjectReference `json:"imagePullSecrets"`
	Replicas         int                         `json:"replicas"`
	Agent            DrControlplaneAgentAgent    `json:"agent"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
	//+optional
	Tolerations []core.Toleration `json:"tolerations"`
}

type DrControlplaneAgentAgent struct {
	//+optional
	CoordKubeconfigData string `json:"coordKubeconfigData"`
	//+optional
	CoordKubeconfigSecret string                 `json:"coordKubeconfigSecret"`
	MetricsAddr           string                 `json:"metricsAddr"`
	Health                DrControlplaneHealth   `json:"health"`
	Election              DrControlplaneElection `json:"election"`
	//+optional
	Resources core.ResourceRequirements `json:"resources"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DrControlplaneAgentList is a list of DrControlplaneAgents
type DrControlplaneAgentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DrControlplaneAgent CRD objects
	Items []DrControlplaneAgent `json:"items,omitempty"`
}
