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

// IoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest struct for IoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest
type IoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest struct {
	Name *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest instantiates a new IoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest() *IoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest {
	this := IoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest{}
	return &this
}

// NewIoArgoprojWorkflowV1alpha1CronWorkflowResumeRequestWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1CronWorkflowResumeRequestWithDefaults() *IoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest {
	this := IoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *IoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest) SetNamespace(v string) {
	o.Namespace = &v
}

func (o IoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest struct {
	value *IoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest) Get() *IoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest) Set(val *IoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest(val *IoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest) *NullableIoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest {
	return &NullableIoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1CronWorkflowResumeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

