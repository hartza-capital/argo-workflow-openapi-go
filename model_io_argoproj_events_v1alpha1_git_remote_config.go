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

// IoArgoprojEventsV1alpha1GitRemoteConfig struct for IoArgoprojEventsV1alpha1GitRemoteConfig
type IoArgoprojEventsV1alpha1GitRemoteConfig struct {
	// Name of the remote to fetch from.
	Name *string `json:"name,omitempty"`
	// URLs the URLs of a remote repository. It must be non-empty. Fetch will always use the first URL, while push will use all of them.
	Urls *[]string `json:"urls,omitempty"`
}

// NewIoArgoprojEventsV1alpha1GitRemoteConfig instantiates a new IoArgoprojEventsV1alpha1GitRemoteConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojEventsV1alpha1GitRemoteConfig() *IoArgoprojEventsV1alpha1GitRemoteConfig {
	this := IoArgoprojEventsV1alpha1GitRemoteConfig{}
	return &this
}

// NewIoArgoprojEventsV1alpha1GitRemoteConfigWithDefaults instantiates a new IoArgoprojEventsV1alpha1GitRemoteConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojEventsV1alpha1GitRemoteConfigWithDefaults() *IoArgoprojEventsV1alpha1GitRemoteConfig {
	this := IoArgoprojEventsV1alpha1GitRemoteConfig{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1GitRemoteConfig) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1GitRemoteConfig) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1GitRemoteConfig) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IoArgoprojEventsV1alpha1GitRemoteConfig) SetName(v string) {
	o.Name = &v
}

// GetUrls returns the Urls field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1GitRemoteConfig) GetUrls() []string {
	if o == nil || o.Urls == nil {
		var ret []string
		return ret
	}
	return *o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1GitRemoteConfig) GetUrlsOk() (*[]string, bool) {
	if o == nil || o.Urls == nil {
		return nil, false
	}
	return o.Urls, true
}

// HasUrls returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1GitRemoteConfig) HasUrls() bool {
	if o != nil && o.Urls != nil {
		return true
	}

	return false
}

// SetUrls gets a reference to the given []string and assigns it to the Urls field.
func (o *IoArgoprojEventsV1alpha1GitRemoteConfig) SetUrls(v []string) {
	o.Urls = &v
}

func (o IoArgoprojEventsV1alpha1GitRemoteConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Urls != nil {
		toSerialize["urls"] = o.Urls
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojEventsV1alpha1GitRemoteConfig struct {
	value *IoArgoprojEventsV1alpha1GitRemoteConfig
	isSet bool
}

func (v NullableIoArgoprojEventsV1alpha1GitRemoteConfig) Get() *IoArgoprojEventsV1alpha1GitRemoteConfig {
	return v.value
}

func (v *NullableIoArgoprojEventsV1alpha1GitRemoteConfig) Set(val *IoArgoprojEventsV1alpha1GitRemoteConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojEventsV1alpha1GitRemoteConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojEventsV1alpha1GitRemoteConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojEventsV1alpha1GitRemoteConfig(val *IoArgoprojEventsV1alpha1GitRemoteConfig) *NullableIoArgoprojEventsV1alpha1GitRemoteConfig {
	return &NullableIoArgoprojEventsV1alpha1GitRemoteConfig{value: val, isSet: true}
}

func (v NullableIoArgoprojEventsV1alpha1GitRemoteConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojEventsV1alpha1GitRemoteConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

