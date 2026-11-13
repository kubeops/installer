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
	"kmodules.xyz/resource-metadata/apis/shared"
)

const (
	ResourceKindCattleSet = "CattleSet"
	ResourceCattleSet     = "cattleset"
	ResourceCattleSets    = "cattlesets"
)

// CattleSet defines the schama for CattleSet operator installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
type CattleSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CattleSetSpec `json:"spec,omitempty"`
}

// CattleSetSpec is the schema for Operator Operator values file
type CattleSetSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string    `json:"fullnameOverride"`
	RegistryFQDN     string    `json:"registryFQDN"`
	ReplicaCount     int32     `json:"replicaCount"`
	Image            Container `json:"image"`
	ImagePullPolicy  string    `json:"imagePullPolicy"`
	//+optional
	ImagePullSecrets []string `json:"imagePullSecrets"`
	//+optional
	CriticalAddon bool `json:"criticalAddon"`
	//+optional
	LogLevel int32 `json:"logLevel"`
	//+optional
	Annotations map[string]string `json:"annotations"`
	//+optional
	PodAnnotations map[string]string `json:"podAnnotations"`
	//+optional
	PodLabels map[string]string `json:"podLabels"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity *core.Affinity `json:"affinity"`
	// PodSecurityContext holds pod-level security attributes and common container settings.
	// Optional: Defaults to empty.  See type description for default values of each field.
	// +optional
	PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext"`
	ServiceAccount     ServiceAccountSpec       `json:"serviceAccount"`
	// +optional
	Apiserver  CattleSetApiserver `json:"apiserver"`
	Monitoring Monitoring         `json:"monitoring"`
	// +optional
	NetworkPolicy NetworkPolicySpec `json:"networkPolicy"`
	// +optional
	Features CattleSetFeatures `json:"features"`
	// +optional
	Distro shared.DistroSpec `json:"distro"`
	// +optional
	MaxConcurrentReconciles int `json:"maxConcurrentReconciles"`
	// The hub controller: reconciler and webhooks. Install it once, on the OCM hub.
	// +optional
	HubController CattleSetHubController `json:"hubController"`
	// The coordination control plane that serves the dc-failover Leases.
	// +optional
	Coordination CattleSetCoordination `json:"coordination"`
	// The gate agent. It runs on every Member spoke with the hub controller off.
	// +optional
	GateAgent CattleSetGateAgent `json:"gateAgent"`
	// +optional
	RequiresPlacementPolicyCRD bool `json:"requiresPlacementPolicyCRD"`
}

type CattleSetHubController struct {
	Enabled bool `json:"enabled"`
	// The OCM managed cluster name of the cluster the hub controller runs in. When it
	// is not a Member of a herd, the herd's hub-local workload is retired, not adopted.
	// +optional
	ClusterName string `json:"clusterName"`
	// This hub is an Arbiter DC and runs no workloads.
	// +optional
	ArbiterHub bool `json:"arbiterHub"`
}

type CattleSetCoordination struct {
	// Name of a Secret with a `kubeconfig` key for the coordination control plane.
	// Empty means the hub's own API server. The hub controller only reads from it.
	// +optional
	KubeconfigSecret string `json:"kubeconfigSecret"`
}

type CattleSetGateAgent struct {
	Enabled bool `json:"enabled"`
	// This spoke's OCM managed cluster name, the DC name the scope markers carry.
	// Required when the agent is enabled.
	// +optional
	DcName string `json:"dcName"`
	// How often a gated Service is re-evaluated, for example 5s.
	// +optional
	Resync string `json:"resync"`
}

type CattleSetFeatures struct {
	// +optional
	Ocm OCMSpec `json:"ocm"`
}

type CattleSetApiserver struct {
	EnableMutatingWebhook   bool                  `json:"enableMutatingWebhook"`
	EnableValidatingWebhook bool                  `json:"enableValidatingWebhook"`
	Healthcheck             HealthcheckSpec       `json:"healthcheck"`
	ServingCerts            CattleSetServingCerts `json:"servingCerts"`
}

type CattleSetServingCerts struct {
	Generate bool `json:"generate"`
	//+optional
	CertManager CattleSetCertManagerCerts `json:"certManager"`
	//+optional
	CaCrt string `json:"caCrt"`
	//+optional
	ServerCrt string `json:"serverCrt"`
	//+optional
	ServerKey string `json:"serverKey"`
}

type CattleSetCertManagerCerts struct {
	Enabled bool `json:"enabled"`
	//+optional
	IssuerRef CattleSetCertManagerIssuerRef `json:"issuerRef"`
}

type CattleSetCertManagerIssuerRef struct {
	//+optional
	Name string `json:"name"`
	//+optional
	Kind string `json:"kind"`
	//+optional
	Group string `json:"group"`
}

// OCMSpec and OCMPlacement are shared with the petset installer types in this
// package (apis/installer/v1alpha1/petset_types.go); they are not redeclared here.

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CattleSetList is a list of CattleSets
type CattleSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CattleSet CRD objects
	Items []CattleSet `json:"items,omitempty"`
}
