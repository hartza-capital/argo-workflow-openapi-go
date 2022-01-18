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

// IoArgoprojEventsV1alpha1CatchupConfiguration struct for IoArgoprojEventsV1alpha1CatchupConfiguration
type IoArgoprojEventsV1alpha1CatchupConfiguration struct {
	Enabled *bool `json:"enabled,omitempty"`
	MaxDuration *string `json:"maxDuration,omitempty"`
}

// NewIoArgoprojEventsV1alpha1CatchupConfiguration instantiates a new IoArgoprojEventsV1alpha1CatchupConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojEventsV1alpha1CatchupConfiguration() *IoArgoprojEventsV1alpha1CatchupConfiguration {
	this := IoArgoprojEventsV1alpha1CatchupConfiguration{}
	return &this
}

// NewIoArgoprojEventsV1alpha1CatchupConfigurationWithDefaults instantiates a new IoArgoprojEventsV1alpha1CatchupConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojEventsV1alpha1CatchupConfigurationWithDefaults() *IoArgoprojEventsV1alpha1CatchupConfiguration {
	this := IoArgoprojEventsV1alpha1CatchupConfiguration{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1CatchupConfiguration) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1CatchupConfiguration) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1CatchupConfiguration) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *IoArgoprojEventsV1alpha1CatchupConfiguration) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMaxDuration returns the MaxDuration field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1CatchupConfiguration) GetMaxDuration() string {
	if o == nil || o.MaxDuration == nil {
		var ret string
		return ret
	}
	return *o.MaxDuration
}

// GetMaxDurationOk returns a tuple with the MaxDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1CatchupConfiguration) GetMaxDurationOk() (*string, bool) {
	if o == nil || o.MaxDuration == nil {
		return nil, false
	}
	return o.MaxDuration, true
}

// HasMaxDuration returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1CatchupConfiguration) HasMaxDuration() bool {
	if o != nil && o.MaxDuration != nil {
		return true
	}

	return false
}

// SetMaxDuration gets a reference to the given string and assigns it to the MaxDuration field.
func (o *IoArgoprojEventsV1alpha1CatchupConfiguration) SetMaxDuration(v string) {
	o.MaxDuration = &v
}

func (o IoArgoprojEventsV1alpha1CatchupConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.MaxDuration != nil {
		toSerialize["maxDuration"] = o.MaxDuration
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojEventsV1alpha1CatchupConfiguration struct {
	value *IoArgoprojEventsV1alpha1CatchupConfiguration
	isSet bool
}

func (v NullableIoArgoprojEventsV1alpha1CatchupConfiguration) Get() *IoArgoprojEventsV1alpha1CatchupConfiguration {
	return v.value
}

func (v *NullableIoArgoprojEventsV1alpha1CatchupConfiguration) Set(val *IoArgoprojEventsV1alpha1CatchupConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojEventsV1alpha1CatchupConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojEventsV1alpha1CatchupConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojEventsV1alpha1CatchupConfiguration(val *IoArgoprojEventsV1alpha1CatchupConfiguration) *NullableIoArgoprojEventsV1alpha1CatchupConfiguration {
	return &NullableIoArgoprojEventsV1alpha1CatchupConfiguration{value: val, isSet: true}
}

func (v NullableIoArgoprojEventsV1alpha1CatchupConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojEventsV1alpha1CatchupConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

