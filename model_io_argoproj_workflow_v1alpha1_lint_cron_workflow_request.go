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

// IoArgoprojWorkflowV1alpha1LintCronWorkflowRequest struct for IoArgoprojWorkflowV1alpha1LintCronWorkflowRequest
type IoArgoprojWorkflowV1alpha1LintCronWorkflowRequest struct {
	CronWorkflow *IoArgoprojWorkflowV1alpha1CronWorkflow `json:"cronWorkflow,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1LintCronWorkflowRequest instantiates a new IoArgoprojWorkflowV1alpha1LintCronWorkflowRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1LintCronWorkflowRequest() *IoArgoprojWorkflowV1alpha1LintCronWorkflowRequest {
	this := IoArgoprojWorkflowV1alpha1LintCronWorkflowRequest{}
	return &this
}

// NewIoArgoprojWorkflowV1alpha1LintCronWorkflowRequestWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1LintCronWorkflowRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1LintCronWorkflowRequestWithDefaults() *IoArgoprojWorkflowV1alpha1LintCronWorkflowRequest {
	this := IoArgoprojWorkflowV1alpha1LintCronWorkflowRequest{}
	return &this
}

// GetCronWorkflow returns the CronWorkflow field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1LintCronWorkflowRequest) GetCronWorkflow() IoArgoprojWorkflowV1alpha1CronWorkflow {
	if o == nil || o.CronWorkflow == nil {
		var ret IoArgoprojWorkflowV1alpha1CronWorkflow
		return ret
	}
	return *o.CronWorkflow
}

// GetCronWorkflowOk returns a tuple with the CronWorkflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1LintCronWorkflowRequest) GetCronWorkflowOk() (*IoArgoprojWorkflowV1alpha1CronWorkflow, bool) {
	if o == nil || o.CronWorkflow == nil {
		return nil, false
	}
	return o.CronWorkflow, true
}

// HasCronWorkflow returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1LintCronWorkflowRequest) HasCronWorkflow() bool {
	if o != nil && o.CronWorkflow != nil {
		return true
	}

	return false
}

// SetCronWorkflow gets a reference to the given IoArgoprojWorkflowV1alpha1CronWorkflow and assigns it to the CronWorkflow field.
func (o *IoArgoprojWorkflowV1alpha1LintCronWorkflowRequest) SetCronWorkflow(v IoArgoprojWorkflowV1alpha1CronWorkflow) {
	o.CronWorkflow = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1LintCronWorkflowRequest) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1LintCronWorkflowRequest) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1LintCronWorkflowRequest) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *IoArgoprojWorkflowV1alpha1LintCronWorkflowRequest) SetNamespace(v string) {
	o.Namespace = &v
}

func (o IoArgoprojWorkflowV1alpha1LintCronWorkflowRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CronWorkflow != nil {
		toSerialize["cronWorkflow"] = o.CronWorkflow
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1LintCronWorkflowRequest struct {
	value *IoArgoprojWorkflowV1alpha1LintCronWorkflowRequest
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1LintCronWorkflowRequest) Get() *IoArgoprojWorkflowV1alpha1LintCronWorkflowRequest {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1LintCronWorkflowRequest) Set(val *IoArgoprojWorkflowV1alpha1LintCronWorkflowRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1LintCronWorkflowRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1LintCronWorkflowRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1LintCronWorkflowRequest(val *IoArgoprojWorkflowV1alpha1LintCronWorkflowRequest) *NullableIoArgoprojWorkflowV1alpha1LintCronWorkflowRequest {
	return &NullableIoArgoprojWorkflowV1alpha1LintCronWorkflowRequest{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1LintCronWorkflowRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1LintCronWorkflowRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

