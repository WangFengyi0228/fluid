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
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/fluid-cloudnative/fluid/pkg/common"
)

type AlluxioRuntimeRole common.RuntimeRole

const (
	// Master is the type for master of Alluxio cluster.
	Master AlluxioRuntimeRole = "master"

	// Worker is the type for workers of Alluxio cluster.
	Worker AlluxioRuntimeRole = "worker"

	// Fuse is the type for chief worker of Alluxio cluster.
	Fuse AlluxioRuntimeRole = "fuse"

	// API Gateway is the API Gateway of Alluxio cluster.
	APIGateway AlluxioRuntimeRole = "apiGateway"
)

// AlluxioCompTemplateSpec is a description of the Alluxio commponents
type AlluxioCompTemplateSpec struct {
	// Replicas is the desired number of replicas of the given template.
	// If unspecified, defaults to 1.
	// +kubebuilder:validation:Minimum=1
	// replicas is the min replicas of dataset in the cluster
	// +optional
	Replicas int32 `json:"replicas,omitempty"`

	// Options for JVM
	JvmOptions []string `json:"jvmOptions,omitempty"`

	// Configurable properties for the Alluxio component. <br>
	// Refer to <a href="https://docs.alluxio.io/os/user/stable/en/reference/Properties-List.html">Alluxio Configuration Properties</a> for more info
	// +optional
	Properties map[string]string `json:"properties,omitempty"`

	// Ports used by Alluxio(e.g. rpc: 19998 for master)
	// +optional
	Ports map[string]int `json:"ports,omitempty"`

	// Resources that will be requested by the Alluxio component. <br>
	// <br>
	// Resources are not allowed for ephemeral containers. Ephemeral containers use spare resources
	// already allocated to the pod.
	// +optional
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`

	// Environment variables that will be used by Alluxio component. <br>
	Env map[string]string `json:"env,omitempty"`

	// Enabled or Disabled for the components. For now, only  API Gateway is enabled or disabled.
	// +optional
	Enabled bool `json:"enabled,omitempty"`

	// NodeSelector is a selector which must be true for the master to fit on a node
	// +optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// Whether to use hostnetwork or not
	// +kubebuilder:validation:Enum=HostNetwork;"";ContainerNetwork
	// +optional
	NetworkMode NetworkMode `json:"networkMode,omitempty"`
	// VolumeMounts specifies the volumes listed in ".spec.volumes" to mount into the alluxio runtime component's filesystem.
	// +optional
	VolumeMounts []corev1.VolumeMount `json:"volumeMounts,omitempty"`

	// PodMetadata defines labels and annotations that will be propagated to Alluxio's pods
	// +optional
	PodMetadata PodMetadata `json:"podMetadata,omitempty"`
}

// AlluxioFuseSpec is a description of the Alluxio Fuse
type AlluxioFuseSpec struct {

	// Image for Alluxio Fuse(e.g. alluxio/alluxio-fuse)
	Image string `json:"image,omitempty"`

	// Image Tag for Alluxio Fuse(e.g. 2.3.0-SNAPSHOT)
	ImageTag string `json:"imageTag,omitempty"`

	// One of the three policies: `Always`, `IfNotPresent`, `Never`
	ImagePullPolicy string `json:"imagePullPolicy,omitempty"`

	// Options for JVM
	JvmOptions []string `json:"jvmOptions,omitempty"`

	// Configurable properties for Alluxio System. <br>
	// Refer to <a href="https://docs.alluxio.io/os/user/stable/en/reference/Properties-List.html">Alluxio Configuration Properties</a> for more info
	Properties map[string]string `json:"properties,omitempty"`

	// Environment variables that will be used by Alluxio Fuse
	Env map[string]string `json:"env,omitempty"`

	// ShortCircuitPolicy string            `json:"shortCircuitPolicy,omitempty"`

	// Resources that will be requested by Alluxio Fuse. <br>
	// <br>
	// Resources are not allowed for ephemeral containers. Ephemeral containers use spare resources
	// already allocated to the pod.
	// +optional
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`

	// Arguments that will be passed to Alluxio Fuse
	Args []string `json:"args,omitempty"`

	// If the fuse client should be deployed in global mode,
	// otherwise the affinity should be considered
	// +optional
	Global bool `json:"global,omitempty"`

	// NodeSelector is a selector which must be true for the fuse client to fit on a node,
	// this option only effect when global is enabled
	// +optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// CleanPolicy decides when to clean Alluxio Fuse pods.
	// Currently Fluid supports two policies: OnDemand and OnRuntimeDeleted
	// OnDemand cleans fuse pod once th fuse pod on some node is not needed
	// OnRuntimeDeleted cleans fuse pod only when the cache runtime is deleted
	// Defaults to OnRuntimeDeleted
	// +optional
	CleanPolicy FuseCleanPolicy `json:"cleanPolicy,omitempty"`

	// Whether to use hostnetwork or not
	// +kubebuilder:validation:Enum=HostNetwork;"";ContainerNetwork
	// +optional
	NetworkMode NetworkMode `json:"networkMode,omitempty"`
	// VolumeMounts specifies the volumes listed in ".spec.volumes" to mount into the alluxio runtime component's filesystem.
	// +optional
	VolumeMounts []corev1.VolumeMount `json:"volumeMounts,omitempty"`

	// PodMetadata defines labels and annotations that will be propagated to Alluxio's fuse pods
	// +optional
	PodMetadata PodMetadata `json:"podMetadata,omitempty"`
}

// Level describes configurations a tier needs. <br>
// Refer to <a href="https://docs.alluxio.io/os/user/stable/en/core-services/Caching.html#configuring-tiered-storage">Configuring Tiered Storage</a> for more info
type Level struct {
	// Alias string `json:"alias,omitempty"`

	// Medium Type of the tier. One of the three types: `MEM`, `SSD`, `HDD`
	// +kubebuilder:validation:Enum=MEM;SSD;HDD
	// +required
	MediumType common.MediumType `json:"mediumtype"`

	// VolumeType is the volume type of the tier. Should be one of the three types: `hostPath`, `emptyDir` and `volumeTemplate`.
	// If not set, defaults to hostPath.
	// +kubebuilder:default=hostPath
	// +kubebuilder:validation:Enum=hostPath;emptyDir
	// +optional
	VolumeType common.VolumeType `json:"volumeType"`

	// VolumeSource is the volume source of the tier. It follows the form of corev1.VolumeSource.
	// For now, users should only specify VolumeSource when VolumeType is set to emptyDir.
	VolumeSource VolumeSource `json:"volumeSource,omitempty"`

	// File paths to be used for the tier. Multiple paths are supported.
	// Multiple paths should be separated with comma. For example: "/mnt/cache1,/mnt/cache2".
	// +kubebuilder:validation:MinLength=1
	// +required
	Path string `json:"path,omitempty"`

	// Quota for the whole tier. (e.g. 100Gi)
	// Please note that if there're multiple paths used for this tierstore,
	// the quota will be equally divided into these paths. If you'd like to
	// set quota for each, path, see QuotaList for more information.
	// +optional
	Quota *resource.Quantity `json:"quota,omitempty"`

	// QuotaList are quotas used to set quota on multiple paths. Quotas should be separated with comma.
	// Quotas in this list will be set to paths with the same order in Path.
	// For example, with Path defined with "/mnt/cache1,/mnt/cache2" and QuotaList set to "100Gi, 50Gi",
	// then we get 100GiB cache storage under "/mnt/cache1" and 50GiB under "/mnt/cache2".
	// Also note that num of quotas must be consistent with the num of paths defined in Path.
	// +optional
	// +kubebuilder:validation:Pattern:="^((\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+)))),)+((\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))))?)$"
	QuotaList string `json:"quotaList,omitempty"`

	// StorageType common.CacheStoreType `json:"storageType,omitempty"`
	// float64 is not supported, https://github.com/kubernetes-sigs/controller-tools/issues/245

	// Ratio of high watermark of the tier (e.g. 0.9)
	High string `json:"high,omitempty"`

	// Ratio of low watermark of the tier (e.g. 0.7)
	Low string `json:"low,omitempty"`
}

// TieredStore is a description of the tiered store
type TieredStore struct {
	// configurations for multiple tiers
	Levels []Level `json:"levels,omitempty"`
}

// Data management strategies
type Data struct {
	// The copies of the dataset
	// +optional
	Replicas int32 `json:"replicas"`

	// Pin the dataset or not. Refer to <a href="https://docs.alluxio.io/os/user/stable/en/operation/User-CLI.html#pin">Alluxio User-CLI pin</a>
	// +optional
	Pin bool `json:"pin"`
}

// AlluxioRuntimeSpec defines the desired state of AlluxioRuntime
type AlluxioRuntimeSpec struct {
	// The version information that instructs fluid to orchestrate a particular version of Alluxio.
	AlluxioVersion VersionSpec `json:"alluxioVersion,omitempty"`

	// The component spec of Alluxio master
	Master AlluxioCompTemplateSpec `json:"master,omitempty"`

	// The component spec of Alluxio job master
	JobMaster AlluxioCompTemplateSpec `json:"jobMaster,omitempty"`

	// The component spec of Alluxio worker
	Worker AlluxioCompTemplateSpec `json:"worker,omitempty"`

	// The component spec of Alluxio job Worker
	JobWorker AlluxioCompTemplateSpec `json:"jobWorker,omitempty"`

	// The component spec of Alluxio API Gateway
	APIGateway AlluxioCompTemplateSpec `json:"apiGateway,omitempty"`

	// The spec of init users
	InitUsers InitUsersSpec `json:"initUsers,omitempty"`

	// The component spec of Alluxio Fuse
	Fuse AlluxioFuseSpec `json:"fuse,omitempty"`

	// Configurable properties for Alluxio system. <br>
	// Refer to <a href="https://docs.alluxio.io/os/user/stable/en/reference/Properties-List.html">Alluxio Configuration Properties</a> for more info
	Properties map[string]string `json:"properties,omitempty"`

	// Options for JVM
	JvmOptions []string `json:"jvmOptions,omitempty"`

	// Tiered storage used by Alluxio
	TieredStore TieredStore `json:"tieredstore,omitempty"`

	// Management strategies for the dataset to which the runtime is bound
	Data Data `json:"data,omitempty"`

	// The replicas of the worker, need to be specified
	Replicas int32 `json:"replicas,omitempty"`

	// Manage the user to run Alluxio Runtime
	RunAs *User `json:"runAs,omitempty"`

	// Disable monitoring for Alluxio Runtime
	// Prometheus is enabled by default
	// +optional
	DisablePrometheus bool `json:"disablePrometheus,omitempty"`

	// Name of the configMap used to support HDFS configurations when using HDFS as Alluxio's UFS. The configMap
	// must be in the same namespace with the AlluxioRuntime. The configMap should contain user-specific HDFS conf files in it.
	// For now, only "hdfs-site.xml" and "core-site.xml" are supported. It must take the filename of the conf file as the key and content
	// of the file as the value.
	// +optional
	HadoopConfig string `json:"hadoopConfig,omitempty"`

	// Volumes is the list of Kubernetes volumes that can be mounted by the alluxio runtime components and/or fuses.
	// +optional
	Volumes []corev1.Volume `json:"volumes,omitempty"`

	// PodMetadata defines labels and annotations that will be propagated to Alluxio's pods
	// +optional
	PodMetadata PodMetadata `json:"podMetadata,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:subresource:scale:specpath=.spec.replicas,statuspath=.status.currentWorkerNumberScheduled,selectorpath=.status.selector
// +kubebuilder:printcolumn:name="Ready Masters",type="integer",JSONPath=`.status.masterNumberReady`,priority=10
// +kubebuilder:printcolumn:name="Desired Masters",type="integer",JSONPath=`.status.desiredMasterNumberScheduled`,priority=10
// +kubebuilder:printcolumn:name="Master Phase",type="string",JSONPath=`.status.masterPhase`,priority=0
// +kubebuilder:printcolumn:name="Ready Workers",type="integer",JSONPath=`.status.workerNumberReady`,priority=10
// +kubebuilder:printcolumn:name="Desired Workers",type="integer",JSONPath=`.status.desiredWorkerNumberScheduled`,priority=10
// +kubebuilder:printcolumn:name="Worker Phase",type="string",JSONPath=`.status.workerPhase`,priority=0
// +kubebuilder:printcolumn:name="Ready Fuses",type="integer",JSONPath=`.status.fuseNumberReady`,priority=10
// +kubebuilder:printcolumn:name="Desired Fuses",type="integer",JSONPath=`.status.desiredFuseNumberScheduled`,priority=10
// +kubebuilder:printcolumn:name="Fuse Phase",type="string",JSONPath=`.status.fusePhase`,priority=0
// +kubebuilder:printcolumn:name="API Gateway",type="string",JSONPath=`.status.apiGateway.endpoint`,priority=10
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=`.metadata.creationTimestamp`,priority=0
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:resource:categories={fluid},shortName=alluxio
// +genclient

// AlluxioRuntime is the Schema for the alluxioruntimes API
type AlluxioRuntime struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlluxioRuntimeSpec `json:"spec,omitempty"`
	Status RuntimeStatus      `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced

// AlluxioRuntimeList contains a list of AlluxioRuntime
type AlluxioRuntimeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlluxioRuntime `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AlluxioRuntime{}, &AlluxioRuntimeList{})
}

// Replicas gets the replicas of runtime worker
func (runtime *AlluxioRuntime) Replicas() int32 {
	return runtime.Spec.Replicas
}

func (runtime *AlluxioRuntime) GetStatus() *RuntimeStatus {
	return &runtime.Status
}
