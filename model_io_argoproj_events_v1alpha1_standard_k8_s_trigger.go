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

// IoArgoprojEventsV1alpha1StandardK8STrigger struct for IoArgoprojEventsV1alpha1StandardK8STrigger
type IoArgoprojEventsV1alpha1StandardK8STrigger struct {
	GroupVersionResource *IoK8sApimachineryPkgApisMetaV1GroupVersionResource `json:"groupVersionResource,omitempty"`
	LiveObject *bool `json:"liveObject,omitempty"`
	Operation *string `json:"operation,omitempty"`
	// Parameters is the list of parameters that is applied to resolved K8s trigger object.
	Parameters *[]IoArgoprojEventsV1alpha1TriggerParameter `json:"parameters,omitempty"`
	PatchStrategy *string `json:"patchStrategy,omitempty"`
	Source *IoArgoprojEventsV1alpha1ArtifactLocation `json:"source,omitempty"`
}

// NewIoArgoprojEventsV1alpha1StandardK8STrigger instantiates a new IoArgoprojEventsV1alpha1StandardK8STrigger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojEventsV1alpha1StandardK8STrigger() *IoArgoprojEventsV1alpha1StandardK8STrigger {
	this := IoArgoprojEventsV1alpha1StandardK8STrigger{}
	return &this
}

// NewIoArgoprojEventsV1alpha1StandardK8STriggerWithDefaults instantiates a new IoArgoprojEventsV1alpha1StandardK8STrigger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojEventsV1alpha1StandardK8STriggerWithDefaults() *IoArgoprojEventsV1alpha1StandardK8STrigger {
	this := IoArgoprojEventsV1alpha1StandardK8STrigger{}
	return &this
}

// GetGroupVersionResource returns the GroupVersionResource field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) GetGroupVersionResource() IoK8sApimachineryPkgApisMetaV1GroupVersionResource {
	if o == nil || o.GroupVersionResource == nil {
		var ret IoK8sApimachineryPkgApisMetaV1GroupVersionResource
		return ret
	}
	return *o.GroupVersionResource
}

// GetGroupVersionResourceOk returns a tuple with the GroupVersionResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) GetGroupVersionResourceOk() (*IoK8sApimachineryPkgApisMetaV1GroupVersionResource, bool) {
	if o == nil || o.GroupVersionResource == nil {
		return nil, false
	}
	return o.GroupVersionResource, true
}

// HasGroupVersionResource returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) HasGroupVersionResource() bool {
	if o != nil && o.GroupVersionResource != nil {
		return true
	}

	return false
}

// SetGroupVersionResource gets a reference to the given IoK8sApimachineryPkgApisMetaV1GroupVersionResource and assigns it to the GroupVersionResource field.
func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) SetGroupVersionResource(v IoK8sApimachineryPkgApisMetaV1GroupVersionResource) {
	o.GroupVersionResource = &v
}

// GetLiveObject returns the LiveObject field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) GetLiveObject() bool {
	if o == nil || o.LiveObject == nil {
		var ret bool
		return ret
	}
	return *o.LiveObject
}

// GetLiveObjectOk returns a tuple with the LiveObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) GetLiveObjectOk() (*bool, bool) {
	if o == nil || o.LiveObject == nil {
		return nil, false
	}
	return o.LiveObject, true
}

// HasLiveObject returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) HasLiveObject() bool {
	if o != nil && o.LiveObject != nil {
		return true
	}

	return false
}

// SetLiveObject gets a reference to the given bool and assigns it to the LiveObject field.
func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) SetLiveObject(v bool) {
	o.LiveObject = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) GetOperation() string {
	if o == nil || o.Operation == nil {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) GetOperationOk() (*string, bool) {
	if o == nil || o.Operation == nil {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) HasOperation() bool {
	if o != nil && o.Operation != nil {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) SetOperation(v string) {
	o.Operation = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) GetParameters() []IoArgoprojEventsV1alpha1TriggerParameter {
	if o == nil || o.Parameters == nil {
		var ret []IoArgoprojEventsV1alpha1TriggerParameter
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) GetParametersOk() (*[]IoArgoprojEventsV1alpha1TriggerParameter, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []IoArgoprojEventsV1alpha1TriggerParameter and assigns it to the Parameters field.
func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) SetParameters(v []IoArgoprojEventsV1alpha1TriggerParameter) {
	o.Parameters = &v
}

// GetPatchStrategy returns the PatchStrategy field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) GetPatchStrategy() string {
	if o == nil || o.PatchStrategy == nil {
		var ret string
		return ret
	}
	return *o.PatchStrategy
}

// GetPatchStrategyOk returns a tuple with the PatchStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) GetPatchStrategyOk() (*string, bool) {
	if o == nil || o.PatchStrategy == nil {
		return nil, false
	}
	return o.PatchStrategy, true
}

// HasPatchStrategy returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) HasPatchStrategy() bool {
	if o != nil && o.PatchStrategy != nil {
		return true
	}

	return false
}

// SetPatchStrategy gets a reference to the given string and assigns it to the PatchStrategy field.
func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) SetPatchStrategy(v string) {
	o.PatchStrategy = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) GetSource() IoArgoprojEventsV1alpha1ArtifactLocation {
	if o == nil || o.Source == nil {
		var ret IoArgoprojEventsV1alpha1ArtifactLocation
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) GetSourceOk() (*IoArgoprojEventsV1alpha1ArtifactLocation, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given IoArgoprojEventsV1alpha1ArtifactLocation and assigns it to the Source field.
func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) SetSource(v IoArgoprojEventsV1alpha1ArtifactLocation) {
	o.Source = &v
}

func (o IoArgoprojEventsV1alpha1StandardK8STrigger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupVersionResource != nil {
		toSerialize["groupVersionResource"] = o.GroupVersionResource
	}
	if o.LiveObject != nil {
		toSerialize["liveObject"] = o.LiveObject
	}
	if o.Operation != nil {
		toSerialize["operation"] = o.Operation
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.PatchStrategy != nil {
		toSerialize["patchStrategy"] = o.PatchStrategy
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojEventsV1alpha1StandardK8STrigger struct {
	value *IoArgoprojEventsV1alpha1StandardK8STrigger
	isSet bool
}

func (v NullableIoArgoprojEventsV1alpha1StandardK8STrigger) Get() *IoArgoprojEventsV1alpha1StandardK8STrigger {
	return v.value
}

func (v *NullableIoArgoprojEventsV1alpha1StandardK8STrigger) Set(val *IoArgoprojEventsV1alpha1StandardK8STrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojEventsV1alpha1StandardK8STrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojEventsV1alpha1StandardK8STrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojEventsV1alpha1StandardK8STrigger(val *IoArgoprojEventsV1alpha1StandardK8STrigger) *NullableIoArgoprojEventsV1alpha1StandardK8STrigger {
	return &NullableIoArgoprojEventsV1alpha1StandardK8STrigger{value: val, isSet: true}
}

func (v NullableIoArgoprojEventsV1alpha1StandardK8STrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojEventsV1alpha1StandardK8STrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

