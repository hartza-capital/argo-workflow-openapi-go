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

// IoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest struct for IoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest
type IoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest struct {
	CronWorkflow *IoArgoprojWorkflowV1alpha1CronWorkflow `json:"cronWorkflow,omitempty"`
	// DEPRECATED: This field is ignored.
	Name *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest instantiates a new IoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest() *IoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest {
	this := IoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest{}
	return &this
}

// NewIoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequestWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequestWithDefaults() *IoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest {
	this := IoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest{}
	return &this
}

// GetCronWorkflow returns the CronWorkflow field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest) GetCronWorkflow() IoArgoprojWorkflowV1alpha1CronWorkflow {
	if o == nil || o.CronWorkflow == nil {
		var ret IoArgoprojWorkflowV1alpha1CronWorkflow
		return ret
	}
	return *o.CronWorkflow
}

// GetCronWorkflowOk returns a tuple with the CronWorkflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest) GetCronWorkflowOk() (*IoArgoprojWorkflowV1alpha1CronWorkflow, bool) {
	if o == nil || o.CronWorkflow == nil {
		return nil, false
	}
	return o.CronWorkflow, true
}

// HasCronWorkflow returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest) HasCronWorkflow() bool {
	if o != nil && o.CronWorkflow != nil {
		return true
	}

	return false
}

// SetCronWorkflow gets a reference to the given IoArgoprojWorkflowV1alpha1CronWorkflow and assigns it to the CronWorkflow field.
func (o *IoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest) SetCronWorkflow(v IoArgoprojWorkflowV1alpha1CronWorkflow) {
	o.CronWorkflow = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *IoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest) SetNamespace(v string) {
	o.Namespace = &v
}

func (o IoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CronWorkflow != nil {
		toSerialize["cronWorkflow"] = o.CronWorkflow
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest struct {
	value *IoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest) Get() *IoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest) Set(val *IoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest(val *IoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest) *NullableIoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest {
	return &NullableIoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1UpdateCronWorkflowRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

