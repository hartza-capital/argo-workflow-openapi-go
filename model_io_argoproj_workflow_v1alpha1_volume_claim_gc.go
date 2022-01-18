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

// IoArgoprojWorkflowV1alpha1VolumeClaimGC VolumeClaimGC describes how to delete volumes from completed Workflows
type IoArgoprojWorkflowV1alpha1VolumeClaimGC struct {
	// Strategy is the strategy to use. One of \"OnWorkflowCompletion\", \"OnWorkflowSuccess\"
	Strategy *string `json:"strategy,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1VolumeClaimGC instantiates a new IoArgoprojWorkflowV1alpha1VolumeClaimGC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1VolumeClaimGC() *IoArgoprojWorkflowV1alpha1VolumeClaimGC {
	this := IoArgoprojWorkflowV1alpha1VolumeClaimGC{}
	return &this
}

// NewIoArgoprojWorkflowV1alpha1VolumeClaimGCWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1VolumeClaimGC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1VolumeClaimGCWithDefaults() *IoArgoprojWorkflowV1alpha1VolumeClaimGC {
	this := IoArgoprojWorkflowV1alpha1VolumeClaimGC{}
	return &this
}

// GetStrategy returns the Strategy field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1VolumeClaimGC) GetStrategy() string {
	if o == nil || o.Strategy == nil {
		var ret string
		return ret
	}
	return *o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1VolumeClaimGC) GetStrategyOk() (*string, bool) {
	if o == nil || o.Strategy == nil {
		return nil, false
	}
	return o.Strategy, true
}

// HasStrategy returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1VolumeClaimGC) HasStrategy() bool {
	if o != nil && o.Strategy != nil {
		return true
	}

	return false
}

// SetStrategy gets a reference to the given string and assigns it to the Strategy field.
func (o *IoArgoprojWorkflowV1alpha1VolumeClaimGC) SetStrategy(v string) {
	o.Strategy = &v
}

func (o IoArgoprojWorkflowV1alpha1VolumeClaimGC) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Strategy != nil {
		toSerialize["strategy"] = o.Strategy
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1VolumeClaimGC struct {
	value *IoArgoprojWorkflowV1alpha1VolumeClaimGC
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1VolumeClaimGC) Get() *IoArgoprojWorkflowV1alpha1VolumeClaimGC {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1VolumeClaimGC) Set(val *IoArgoprojWorkflowV1alpha1VolumeClaimGC) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1VolumeClaimGC) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1VolumeClaimGC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1VolumeClaimGC(val *IoArgoprojWorkflowV1alpha1VolumeClaimGC) *NullableIoArgoprojWorkflowV1alpha1VolumeClaimGC {
	return &NullableIoArgoprojWorkflowV1alpha1VolumeClaimGC{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1VolumeClaimGC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1VolumeClaimGC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

