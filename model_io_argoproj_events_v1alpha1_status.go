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

// IoArgoprojEventsV1alpha1Status Status is a common structure which can be used for Status field.
type IoArgoprojEventsV1alpha1Status struct {
	Conditions *[]IoArgoprojEventsV1alpha1Condition `json:"conditions,omitempty"`
}

// NewIoArgoprojEventsV1alpha1Status instantiates a new IoArgoprojEventsV1alpha1Status object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojEventsV1alpha1Status() *IoArgoprojEventsV1alpha1Status {
	this := IoArgoprojEventsV1alpha1Status{}
	return &this
}

// NewIoArgoprojEventsV1alpha1StatusWithDefaults instantiates a new IoArgoprojEventsV1alpha1Status object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojEventsV1alpha1StatusWithDefaults() *IoArgoprojEventsV1alpha1Status {
	this := IoArgoprojEventsV1alpha1Status{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1Status) GetConditions() []IoArgoprojEventsV1alpha1Condition {
	if o == nil || o.Conditions == nil {
		var ret []IoArgoprojEventsV1alpha1Condition
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1Status) GetConditionsOk() (*[]IoArgoprojEventsV1alpha1Condition, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1Status) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []IoArgoprojEventsV1alpha1Condition and assigns it to the Conditions field.
func (o *IoArgoprojEventsV1alpha1Status) SetConditions(v []IoArgoprojEventsV1alpha1Condition) {
	o.Conditions = &v
}

func (o IoArgoprojEventsV1alpha1Status) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojEventsV1alpha1Status struct {
	value *IoArgoprojEventsV1alpha1Status
	isSet bool
}

func (v NullableIoArgoprojEventsV1alpha1Status) Get() *IoArgoprojEventsV1alpha1Status {
	return v.value
}

func (v *NullableIoArgoprojEventsV1alpha1Status) Set(val *IoArgoprojEventsV1alpha1Status) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojEventsV1alpha1Status) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojEventsV1alpha1Status) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojEventsV1alpha1Status(val *IoArgoprojEventsV1alpha1Status) *NullableIoArgoprojEventsV1alpha1Status {
	return &NullableIoArgoprojEventsV1alpha1Status{value: val, isSet: true}
}

func (v NullableIoArgoprojEventsV1alpha1Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojEventsV1alpha1Status) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

