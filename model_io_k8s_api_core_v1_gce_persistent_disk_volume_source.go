/*
Argo Server API

You can get examples of requests and responses by using the CLI with `--gloglevel=9`, e.g. `argo list --gloglevel=9`

API version: VERSION
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// IoK8sApiCoreV1GCEPersistentDiskVolumeSource Represents a Persistent Disk resource in Google Compute Engine.  A GCE PD must exist before mounting to a container. The disk must also be in the same GCE project and zone as the kubelet. A GCE PD can only be mounted as read/write once or read-only many times. GCE PDs support ownership management and SELinux relabeling.
type IoK8sApiCoreV1GCEPersistentDiskVolumeSource struct {
	// Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	FsType *string `json:"fsType,omitempty"`
	// The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as \"1\". Similarly, the volume partition for /dev/sda is \"0\" (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	Partition *int32 `json:"partition,omitempty"`
	// Unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	PdName string `json:"pdName"`
	// ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	ReadOnly *bool `json:"readOnly,omitempty"`
}

// NewIoK8sApiCoreV1GCEPersistentDiskVolumeSource instantiates a new IoK8sApiCoreV1GCEPersistentDiskVolumeSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoK8sApiCoreV1GCEPersistentDiskVolumeSource(pdName string) *IoK8sApiCoreV1GCEPersistentDiskVolumeSource {
	this := IoK8sApiCoreV1GCEPersistentDiskVolumeSource{}
	this.PdName = pdName
	return &this
}

// NewIoK8sApiCoreV1GCEPersistentDiskVolumeSourceWithDefaults instantiates a new IoK8sApiCoreV1GCEPersistentDiskVolumeSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoK8sApiCoreV1GCEPersistentDiskVolumeSourceWithDefaults() *IoK8sApiCoreV1GCEPersistentDiskVolumeSource {
	this := IoK8sApiCoreV1GCEPersistentDiskVolumeSource{}
	return &this
}

// GetFsType returns the FsType field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) GetFsType() string {
	if o == nil || o.FsType == nil {
		var ret string
		return ret
	}
	return *o.FsType
}

// GetFsTypeOk returns a tuple with the FsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) GetFsTypeOk() (*string, bool) {
	if o == nil || o.FsType == nil {
		return nil, false
	}
	return o.FsType, true
}

// HasFsType returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) HasFsType() bool {
	if o != nil && o.FsType != nil {
		return true
	}

	return false
}

// SetFsType gets a reference to the given string and assigns it to the FsType field.
func (o *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) SetFsType(v string) {
	o.FsType = &v
}

// GetPartition returns the Partition field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) GetPartition() int32 {
	if o == nil || o.Partition == nil {
		var ret int32
		return ret
	}
	return *o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) GetPartitionOk() (*int32, bool) {
	if o == nil || o.Partition == nil {
		return nil, false
	}
	return o.Partition, true
}

// HasPartition returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) HasPartition() bool {
	if o != nil && o.Partition != nil {
		return true
	}

	return false
}

// SetPartition gets a reference to the given int32 and assigns it to the Partition field.
func (o *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) SetPartition(v int32) {
	o.Partition = &v
}

// GetPdName returns the PdName field value
func (o *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) GetPdName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PdName
}

// GetPdNameOk returns a tuple with the PdName field value
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) GetPdNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PdName, true
}

// SetPdName sets field value
func (o *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) SetPdName(v string) {
	o.PdName = v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) GetReadOnly() bool {
	if o == nil || o.ReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) GetReadOnlyOk() (*bool, bool) {
	if o == nil || o.ReadOnly == nil {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) HasReadOnly() bool {
	if o != nil && o.ReadOnly != nil {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

func (o IoK8sApiCoreV1GCEPersistentDiskVolumeSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FsType != nil {
		toSerialize["fsType"] = o.FsType
	}
	if o.Partition != nil {
		toSerialize["partition"] = o.Partition
	}
	if true {
		toSerialize["pdName"] = o.PdName
	}
	if o.ReadOnly != nil {
		toSerialize["readOnly"] = o.ReadOnly
	}
	return json.Marshal(toSerialize)
}

type NullableIoK8sApiCoreV1GCEPersistentDiskVolumeSource struct {
	value *IoK8sApiCoreV1GCEPersistentDiskVolumeSource
	isSet bool
}

func (v NullableIoK8sApiCoreV1GCEPersistentDiskVolumeSource) Get() *IoK8sApiCoreV1GCEPersistentDiskVolumeSource {
	return v.value
}

func (v *NullableIoK8sApiCoreV1GCEPersistentDiskVolumeSource) Set(val *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) {
	v.value = val
	v.isSet = true
}

func (v NullableIoK8sApiCoreV1GCEPersistentDiskVolumeSource) IsSet() bool {
	return v.isSet
}

func (v *NullableIoK8sApiCoreV1GCEPersistentDiskVolumeSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoK8sApiCoreV1GCEPersistentDiskVolumeSource(val *IoK8sApiCoreV1GCEPersistentDiskVolumeSource) *NullableIoK8sApiCoreV1GCEPersistentDiskVolumeSource {
	return &NullableIoK8sApiCoreV1GCEPersistentDiskVolumeSource{value: val, isSet: true}
}

func (v NullableIoK8sApiCoreV1GCEPersistentDiskVolumeSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoK8sApiCoreV1GCEPersistentDiskVolumeSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

