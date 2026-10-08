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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

const (
	ResourceKindDrControlplane = "DrControlplane"
	ResourceDrControlplane     = "drcontrolplane"
	ResourceDrControlplanes    = "drcontrolplanes"
)

// DrControlplane defines the schema for the dr-controlplane installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=drcontrolplanes,singular=drcontrolplane,categories={kubeops,appscode}
type DrControlplane struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DrControlplaneSpec `json:"spec,omitempty"`
}

// DrControlplaneSpec is the schema for the dr-controlplane chart values file
type DrControlplaneSpec struct {
	Namespace       string              `json:"namespace"`
	CreateNamespace bool                `json:"createNamespace"`
	Image           DrControlplaneImage `json:"image"`
	//+optional
	ImagePullSecrets []core.LocalObjectReference `json:"imagePullSecrets"`
	Etcd             DrControlplaneEtcd          `json:"etcd"`
	Controlplane     DrControlplaneControlplane  `json:"controlplane"`
	Agent            DrControlplaneAgentValues   `json:"agent"`
	Topology         DrControlplaneTopology      `json:"topology"`
	Addon            DrControlplaneAddon         `json:"addon"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
	//+optional
	Tolerations []core.Toleration `json:"tolerations"`
	//+optional
	Affinity *core.Affinity `json:"affinity"`
}

type DrControlplaneImage struct {
	Repository string `json:"repository"`
	//+optional
	Tag        string `json:"tag"`
	PullPolicy string `json:"pullPolicy"`
}

type DrControlplaneEtcd struct {
	Deploy   bool   `json:"deploy"`
	Replicas int    `json:"replicas"`
	Image    string `json:"image"`
	//+optional
	StorageClassName    string                   `json:"storageClassName"`
	StorageSize         string                   `json:"storageSize"`
	HeartbeatIntervalMs int                      `json:"heartbeatIntervalMs"`
	ElectionTimeoutMs   int                      `json:"electionTimeoutMs"`
	Member              DrControlplaneEtcdMember `json:"member"`
	//+optional
	Peers               []string `json:"peers"`
	InitialClusterState string   `json:"initialClusterState"`
	//+optional
	ExternalEndpoints []string `json:"externalEndpoints"`
}

type DrControlplaneEtcdMember struct {
	//+optional
	Name string `json:"name"`
	//+optional
	AdvertiseHost string                    `json:"advertiseHost"`
	ClientPort    int                       `json:"clientPort"`
	PeerPort      int                       `json:"peerPort"`
	HostNetwork   bool                      `json:"hostNetwork"`
	HostPath      string                    `json:"hostPath"`
	Persistence   DrControlplanePersistence `json:"persistence"`
	Service       DrControlplaneEtcdService `json:"service"`
	TLS           DrControlplaneEtcdTLS     `json:"tls"`
}

type DrControlplanePersistence struct {
	Enabled bool `json:"enabled"`
	//+optional
	ExistingClaim string `json:"existingClaim"`
	Size          string `json:"size"`
	//+optional
	StorageClassName string                            `json:"storageClassName"`
	AccessModes      []core.PersistentVolumeAccessMode `json:"accessModes"`
}

type DrControlplaneEtcdService struct {
	Enabled bool   `json:"enabled"`
	Type    string `json:"type"`
	//+optional
	Annotations map[string]string `json:"annotations"`
	//+optional
	LoadBalancerIP string `json:"loadBalancerIP"`
}

type DrControlplaneEtcdTLS struct {
	Enabled bool `json:"enabled"`
	//+optional
	SecretName         string `json:"secretName"`
	ClientCertAuth     bool   `json:"clientCertAuth"`
	PeerClientCertAuth bool   `json:"peerClientCertAuth"`
}

type DrControlplaneControlplane struct {
	Enabled bool `json:"enabled"`
	//+optional
	Image string `json:"image"`
	//+optional
	Command          []string `json:"command"`
	Replicas         int      `json:"replicas"`
	ExternalHostname string   `json:"externalHostname"`
	Port             int      `json:"port"`
	EtcdPrefix       string   `json:"etcdPrefix"`
	//+optional
	EtcdTLSSecret string `json:"etcdTLSSecret"`
	//+optional
	ApiserverCASecret   string                            `json:"apiserverCASecret"`
	GenerateApiserverCA bool                              `json:"generateApiserverCA"`
	Persistence         DrControlplanePersistence         `json:"persistence"`
	Service             DrControlplaneControlplaneService `json:"service"`
}

type DrControlplaneControlplaneService struct {
	Type string `json:"type"`
}

type DrControlplaneAgentValues struct {
	Enabled  bool `json:"enabled"`
	Replicas int  `json:"replicas"`
	//+optional
	DcName string `json:"dcName"`
	//+optional
	CoordKubeconfigSecret string                 `json:"coordKubeconfigSecret"`
	MetricsAddr           string                 `json:"metricsAddr"`
	Health                DrControlplaneHealth   `json:"health"`
	Election              DrControlplaneElection `json:"election"`
	//+optional
	Resources core.ResourceRequirements `json:"resources"`
}

type DrControlplaneHealth struct {
	LeaseDurationSeconds int `json:"leaseDurationSeconds"`
	RenewIntervalSeconds int `json:"renewIntervalSeconds"`
}

type DrControlplaneElection struct {
	LeaseDurationSeconds int `json:"leaseDurationSeconds"`
	RenewDeadlineSeconds int `json:"renewDeadlineSeconds"`
	RetryPeriodSeconds   int `json:"retryPeriodSeconds"`
}

type DrControlplaneTopology struct {
	Enabled       bool `json:"enabled"`
	RequireSpread bool `json:"requireSpread"`
	//+optional
	HubKubeconfigSecret string `json:"hubKubeconfigSecret"`
	//+optional
	CoordKubeconfigSecret string `json:"coordKubeconfigSecret"`
	//+optional
	Regions map[string]string `json:"regions"`
	//+optional
	Resources core.ResourceRequirements `json:"resources"`
}

type DrControlplaneAddon struct {
	AutoInstall DrControlplaneAutoInstall `json:"autoInstall"`
	ClusterSet  string                    `json:"clusterSet"`
	//+optional
	CoordKubeconfigMirrorNamespaces []string `json:"coordKubeconfigMirrorNamespaces"`
	//+optional
	PlacementNamespace string                     `json:"placementNamespace"`
	Placement          DrControlplanePlacement    `json:"placement"`
	Agent              DrControlplaneAddonAgent   `json:"agent"`
	Manager            DrControlplaneAddonManager `json:"manager"`
}

type DrControlplaneAutoInstall struct {
	Enabled bool `json:"enabled"`
}

type DrControlplanePlacement struct {
	// Raw passthrough into Placement .spec.predicates.
	//+optional
	Predicates []runtime.RawExtension `json:"predicates"`
	// Shortcut used only when predicates is empty.
	//+optional
	LabelSelector metav1.LabelSelector `json:"labelSelector"`
}

type DrControlplaneAddonAgent struct {
	InstallNamespace            string `json:"installNamespace"`
	CreateNamespace             bool   `json:"createNamespace"`
	CoordKubeconfigSecret       string `json:"coordKubeconfigSecret"`
	Replicas                    int    `json:"replicas"`
	CoordKubeconfigSourceSecret string `json:"coordKubeconfigSourceSecret"`
	//+optional
	CoordExternalEndpoint string                   `json:"coordExternalEndpoint"`
	Image                 DrControlplaneAddonImage `json:"image"`
	//+optional
	ImagePullSecrets []string               `json:"imagePullSecrets"`
	Health           DrControlplaneHealth   `json:"health"`
	Election         DrControlplaneElection `json:"election"`
}

type DrControlplaneAddonImage struct {
	//+optional
	Repository string `json:"repository"`
	//+optional
	Tag string `json:"tag"`
}

type DrControlplaneAddonManager struct {
	Replicas int                             `json:"replicas"`
	Image    DrControlplaneAddonManagerImage `json:"image"`
	//+optional
	Args []string `json:"args"`
	//+optional
	Resources core.ResourceRequirements `json:"resources"`
}

type DrControlplaneAddonManagerImage struct {
	//+optional
	Repository string `json:"repository"`
	//+optional
	Tag string `json:"tag"`
	//+optional
	PullPolicy string `json:"pullPolicy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DrControlplaneList is a list of DrControlplanes
type DrControlplaneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DrControlplane CRD objects
	Items []DrControlplane `json:"items,omitempty"`
}
