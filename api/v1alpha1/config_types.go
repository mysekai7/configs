/*


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
	"encoding/json"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// Product
type Product struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// Global configuration
type GlobalConfig struct {
	// Platform access protocol, support http or https
	// +kubebuilder:validation:Enum:={"http","https"}
	// +kubebuilder:default:=https
	Scheme string `json:"scheme,omitempty"`

	// Platform access host address, support domain name or IP
	Host string `json:"host"`

	// The namespace deployed by the platform, default cpaas-system
	// +kubebuilder:default:=cpaas-system
	Namespace string `json:"namespace,omitempty"`

	// Platform resource instance label field, the default is cpaas.io
	// +kubebuilder:default:=cpaas.io
	LabelBaseDomain string `json:"labelBaseDomain,omitempty"`

	// The platform deploys a management account by default, and the default email admin@cpaas.io
	// +kubebuilder:default:=admin@cpaas.io
	DefaultAdmin string `json:"defaultAdmin,omitempty"`

	// The number of instances of platform common deployment components, the default is 2
	// +kubebuilder:validation:Minimum:=0
	// +kubebuilder:default:=2
	Replicas int `json:"replicas,omitempty"`

	// Platform api gateway address
	ErebusApiAddress string `json:"erebusApiAddress"`

	// Platform api address
	ApiAddress string `json:"apiAddress"`

	// Whether to deploy on openshift cluster
	// +kubebuilder:default:=false
	IsOCP bool `json:"isOCP,omitempty"`

	// Tls Secret
	TlsSecret Certificate `json:"tlsSecret"`
}

type Certificate struct {
	SecretName string `json:"secretName"`
}

type ExtValues struct {
	// Data holds the configuration keys and values.
	// Work around for https://github.com/kubernetes-sigs/kubebuilder/issues/528
	Data map[string]interface{} `json:"-"`
}

// MarshalJSON marshals the HelmValues data to a JSON blob.
func (v ExtValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Data)
}

// UnmarshalJSON sets the HelmValues to a copy of data.
func (v *ExtValues) UnmarshalJSON(data []byte) error {
	var out map[string]interface{}
	err := json.Unmarshal(data, &out)
	if err != nil {
		return err
	}
	v.Data = out
	return nil
}

// DeepCopyInto is an deepcopy function, copying the receiver, writing
// into out. In must be non-nil. Declaring this here prevents it from
// being generated in `zz_generated.deepcopy.go`.
//
// This exists here to work around https://github.com/kubernetes/code-generator/issues/50,
// and partially around https://github.com/kubernetes-sigs/controller-tools/pull/126
// and https://github.com/kubernetes-sigs/controller-tools/issues/294.
func (in *ExtValues) DeepCopyInto(out *ExtValues) {
	b, err := json.Marshal(in.Data)
	if err != nil {
		// The marshal should have been performed cleanly as otherwise
		// the resource would not have been created by the API server.
		panic(err)
	}
	var c map[string]interface{}
	err = json.Unmarshal(b, &c)
	if err != nil {
		panic(err)
	}
	out.Data = c
	return
}

type OidcConfig struct {
	// OIDC server address
	Issuer string `json:"issuer"`

	// OIDC Client ID
	// +kubebuilder:default:="alauda-auth"
	ClientID string `json:"clientID,omitempty"`

	// OIDC Client Secret
	// +kubebuilder:default:=ZXhhbXBsZS1hcHAtc2VjcmV0
	ClientSecret string `json:"clientSecret,omitempty"`

	// OIDC Response Type
	// +kubebuilder:default:=code
	ResponseType string `json:"responseType,omitempty"`

	// OIDC Scopes
	// +kubebuilder:default:="openid,profile,email,groups,ext"
	Scopes string `json:"scopes,omitempty"`
}

type EtcdConfig struct {
	// Etcd servers
	Servers []string `json:"servers"`

	// Etcd Peer secret
	EtcdPeerSecret Certificate `json:"ectdPeerSecret"`

	// Etcd Ca Secret
	EtcdCaSecret Certificate `json:"ectdCaSecret"`
}

type ElasticsearchConfig struct {
	// Server address
	// +kubebuilder:default:="http://cpaas-elasticsearch:9200"
	Host string `json:"host,omitempty"`

	// Log nodes
	Nodes []string `json:"nodes"`

	// UserName
	User v1.SecretKeySelector `json:"user"`

	// Password
	Password v1.SecretKeySelector `json:"password"`
}

type KafkaConfig struct {
	// Whether to enable authentication
	// +kubebuilder:default:=true
	Auth bool `json:"auth,omitempty"`

	// Server address
	// +kubebuilder:default:="cpaas-kafka:9092"
	Host string `json:"host,omitempty"`

	KafkaUser     v1.SecretKeySelector `json:"kafka_user"`
	KafkaPassword v1.SecretKeySelector `json:"kafka_password"`

	ZKUser        v1.SecretKeySelector `json:"zk_user"`
	ZKPassword    v1.SecretKeySelector `json:"zk_password"`
	ZKAclPassword v1.SecretKeySelector `json:"zk_acl_password"`
}

// Repository Config
type Repository struct {
	// Repository type
	// +kubebuilder:validation:Enum:={"chart","image","yum","apt"}
	Type string `json:"type"`
	// Repository url
	Url string `json:"string"`
}

// ConfigSpec defines the desired state of Config
type ConfigSpec struct {
	// Release verison
	Release string `json:"release"`

	// Deployed products
	Products []Product `json:"products"`

	// Global configuration
	Global GlobalConfig `json:"global"`

	//// Certificate
	//Certificates []Certificate `json:"certificates"`

	// OIDC Config
	Oidc OidcConfig `json:"oidc"`

	// Etcd Config
	Etcd EtcdConfig `json:"etcd"`

	// Elasticsearch Config
	Elasticsearch ElasticsearchConfig `json:"elasticsearch,omitempty"`

	// Kafka Config
	Kafka KafkaConfig `json:"kafka,omitempty"`

	// Repositories Config
	Repositories []Repository `json:"repositories"`

	// Ext Config
	Ext ExtValues `json:"ext,omitempty"`
}

// ConfigStatus defines the observed state of Config
type ConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=namespace

// Config is the Schema for the configs API
type Config struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigSpec   `json:"spec,omitempty"`
	Status ConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigList contains a list of Config
type ConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Config `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Config{}, &ConfigList{})
}