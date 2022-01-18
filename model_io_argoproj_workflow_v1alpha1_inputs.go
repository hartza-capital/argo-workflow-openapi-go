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

// IoArgoprojWorkflowV1alpha1Inputs Inputs are the mechanism for passing parameters, artifacts, volumes from one template to another
type IoArgoprojWorkflowV1alpha1Inputs struct {
	// Artifact are a list of artifacts passed as inputs
	Artifacts *[]IoArgoprojWorkflowV1alpha1Artifact `json:"artifacts,omitempty"`
	// Parameters are a list of parameters passed as inputs
	Parameters *[]IoArgoprojWorkflowV1alpha1Parameter `json:"parameters,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1Inputs instantiates a new IoArgoprojWorkflowV1alpha1Inputs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1Inputs() *IoArgoprojWorkflowV1alpha1Inputs {
	this := IoArgoprojWorkflowV1alpha1Inputs{}
	return &this
}

// NewIoArgoprojWorkflowV1alpha1InputsWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1Inputs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1InputsWithDefaults() *IoArgoprojWorkflowV1alpha1Inputs {
	this := IoArgoprojWorkflowV1alpha1Inputs{}
	return &this
}

// GetArtifacts returns the Artifacts field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1Inputs) GetArtifacts() []IoArgoprojWorkflowV1alpha1Artifact {
	if o == nil || o.Artifacts == nil {
		var ret []IoArgoprojWorkflowV1alpha1Artifact
		return ret
	}
	return *o.Artifacts
}

// GetArtifactsOk returns a tuple with the Artifacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1Inputs) GetArtifactsOk() (*[]IoArgoprojWorkflowV1alpha1Artifact, bool) {
	if o == nil || o.Artifacts == nil {
		return nil, false
	}
	return o.Artifacts, true
}

// HasArtifacts returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1Inputs) HasArtifacts() bool {
	if o != nil && o.Artifacts != nil {
		return true
	}

	return false
}

// SetArtifacts gets a reference to the given []IoArgoprojWorkflowV1alpha1Artifact and assigns it to the Artifacts field.
func (o *IoArgoprojWorkflowV1alpha1Inputs) SetArtifacts(v []IoArgoprojWorkflowV1alpha1Artifact) {
	o.Artifacts = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1Inputs) GetParameters() []IoArgoprojWorkflowV1alpha1Parameter {
	if o == nil || o.Parameters == nil {
		var ret []IoArgoprojWorkflowV1alpha1Parameter
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1Inputs) GetParametersOk() (*[]IoArgoprojWorkflowV1alpha1Parameter, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1Inputs) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []IoArgoprojWorkflowV1alpha1Parameter and assigns it to the Parameters field.
func (o *IoArgoprojWorkflowV1alpha1Inputs) SetParameters(v []IoArgoprojWorkflowV1alpha1Parameter) {
	o.Parameters = &v
}

func (o IoArgoprojWorkflowV1alpha1Inputs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Artifacts != nil {
		toSerialize["artifacts"] = o.Artifacts
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1Inputs struct {
	value *IoArgoprojWorkflowV1alpha1Inputs
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1Inputs) Get() *IoArgoprojWorkflowV1alpha1Inputs {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1Inputs) Set(val *IoArgoprojWorkflowV1alpha1Inputs) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1Inputs) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1Inputs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1Inputs(val *IoArgoprojWorkflowV1alpha1Inputs) *NullableIoArgoprojWorkflowV1alpha1Inputs {
	return &NullableIoArgoprojWorkflowV1alpha1Inputs{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1Inputs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1Inputs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

