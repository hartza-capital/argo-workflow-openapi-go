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

// IoArgoprojWorkflowV1alpha1WorkflowStopRequest struct for IoArgoprojWorkflowV1alpha1WorkflowStopRequest
type IoArgoprojWorkflowV1alpha1WorkflowStopRequest struct {
	Message *string `json:"message,omitempty"`
	Name *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	NodeFieldSelector *string `json:"nodeFieldSelector,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1WorkflowStopRequest instantiates a new IoArgoprojWorkflowV1alpha1WorkflowStopRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1WorkflowStopRequest() *IoArgoprojWorkflowV1alpha1WorkflowStopRequest {
	this := IoArgoprojWorkflowV1alpha1WorkflowStopRequest{}
	return &this
}

// NewIoArgoprojWorkflowV1alpha1WorkflowStopRequestWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1WorkflowStopRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1WorkflowStopRequestWithDefaults() *IoArgoprojWorkflowV1alpha1WorkflowStopRequest {
	this := IoArgoprojWorkflowV1alpha1WorkflowStopRequest{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1WorkflowStopRequest) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowStopRequest) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowStopRequest) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *IoArgoprojWorkflowV1alpha1WorkflowStopRequest) SetMessage(v string) {
	o.Message = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1WorkflowStopRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowStopRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowStopRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IoArgoprojWorkflowV1alpha1WorkflowStopRequest) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1WorkflowStopRequest) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowStopRequest) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowStopRequest) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *IoArgoprojWorkflowV1alpha1WorkflowStopRequest) SetNamespace(v string) {
	o.Namespace = &v
}

// GetNodeFieldSelector returns the NodeFieldSelector field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1WorkflowStopRequest) GetNodeFieldSelector() string {
	if o == nil || o.NodeFieldSelector == nil {
		var ret string
		return ret
	}
	return *o.NodeFieldSelector
}

// GetNodeFieldSelectorOk returns a tuple with the NodeFieldSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowStopRequest) GetNodeFieldSelectorOk() (*string, bool) {
	if o == nil || o.NodeFieldSelector == nil {
		return nil, false
	}
	return o.NodeFieldSelector, true
}

// HasNodeFieldSelector returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowStopRequest) HasNodeFieldSelector() bool {
	if o != nil && o.NodeFieldSelector != nil {
		return true
	}

	return false
}

// SetNodeFieldSelector gets a reference to the given string and assigns it to the NodeFieldSelector field.
func (o *IoArgoprojWorkflowV1alpha1WorkflowStopRequest) SetNodeFieldSelector(v string) {
	o.NodeFieldSelector = &v
}

func (o IoArgoprojWorkflowV1alpha1WorkflowStopRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.NodeFieldSelector != nil {
		toSerialize["nodeFieldSelector"] = o.NodeFieldSelector
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1WorkflowStopRequest struct {
	value *IoArgoprojWorkflowV1alpha1WorkflowStopRequest
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1WorkflowStopRequest) Get() *IoArgoprojWorkflowV1alpha1WorkflowStopRequest {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1WorkflowStopRequest) Set(val *IoArgoprojWorkflowV1alpha1WorkflowStopRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1WorkflowStopRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1WorkflowStopRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1WorkflowStopRequest(val *IoArgoprojWorkflowV1alpha1WorkflowStopRequest) *NullableIoArgoprojWorkflowV1alpha1WorkflowStopRequest {
	return &NullableIoArgoprojWorkflowV1alpha1WorkflowStopRequest{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1WorkflowStopRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1WorkflowStopRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

