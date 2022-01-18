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

// IoArgoprojWorkflowV1alpha1DataSource DataSource sources external data into a data template
type IoArgoprojWorkflowV1alpha1DataSource struct {
	ArtifactPaths *IoArgoprojWorkflowV1alpha1ArtifactPaths `json:"artifactPaths,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1DataSource instantiates a new IoArgoprojWorkflowV1alpha1DataSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1DataSource() *IoArgoprojWorkflowV1alpha1DataSource {
	this := IoArgoprojWorkflowV1alpha1DataSource{}
	return &this
}

// NewIoArgoprojWorkflowV1alpha1DataSourceWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1DataSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1DataSourceWithDefaults() *IoArgoprojWorkflowV1alpha1DataSource {
	this := IoArgoprojWorkflowV1alpha1DataSource{}
	return &this
}

// GetArtifactPaths returns the ArtifactPaths field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1DataSource) GetArtifactPaths() IoArgoprojWorkflowV1alpha1ArtifactPaths {
	if o == nil || o.ArtifactPaths == nil {
		var ret IoArgoprojWorkflowV1alpha1ArtifactPaths
		return ret
	}
	return *o.ArtifactPaths
}

// GetArtifactPathsOk returns a tuple with the ArtifactPaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1DataSource) GetArtifactPathsOk() (*IoArgoprojWorkflowV1alpha1ArtifactPaths, bool) {
	if o == nil || o.ArtifactPaths == nil {
		return nil, false
	}
	return o.ArtifactPaths, true
}

// HasArtifactPaths returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1DataSource) HasArtifactPaths() bool {
	if o != nil && o.ArtifactPaths != nil {
		return true
	}

	return false
}

// SetArtifactPaths gets a reference to the given IoArgoprojWorkflowV1alpha1ArtifactPaths and assigns it to the ArtifactPaths field.
func (o *IoArgoprojWorkflowV1alpha1DataSource) SetArtifactPaths(v IoArgoprojWorkflowV1alpha1ArtifactPaths) {
	o.ArtifactPaths = &v
}

func (o IoArgoprojWorkflowV1alpha1DataSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArtifactPaths != nil {
		toSerialize["artifactPaths"] = o.ArtifactPaths
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1DataSource struct {
	value *IoArgoprojWorkflowV1alpha1DataSource
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1DataSource) Get() *IoArgoprojWorkflowV1alpha1DataSource {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1DataSource) Set(val *IoArgoprojWorkflowV1alpha1DataSource) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1DataSource) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1DataSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1DataSource(val *IoArgoprojWorkflowV1alpha1DataSource) *NullableIoArgoprojWorkflowV1alpha1DataSource {
	return &NullableIoArgoprojWorkflowV1alpha1DataSource{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1DataSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1DataSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

