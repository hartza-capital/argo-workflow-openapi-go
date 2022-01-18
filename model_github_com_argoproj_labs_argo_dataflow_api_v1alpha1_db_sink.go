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

// GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink struct for GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink
type GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink struct {
	Actions *[]GithubComArgoprojLabsArgoDataflowApiV1alpha1SQLAction `json:"actions,omitempty"`
	Database *GithubComArgoprojLabsArgoDataflowApiV1alpha1Database `json:"database,omitempty"`
}

// NewGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink() *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink {
	this := GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink{}
	return &this
}

// NewGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSinkWithDefaults instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSinkWithDefaults() *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink {
	this := GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink) GetActions() []GithubComArgoprojLabsArgoDataflowApiV1alpha1SQLAction {
	if o == nil || o.Actions == nil {
		var ret []GithubComArgoprojLabsArgoDataflowApiV1alpha1SQLAction
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink) GetActionsOk() (*[]GithubComArgoprojLabsArgoDataflowApiV1alpha1SQLAction, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given []GithubComArgoprojLabsArgoDataflowApiV1alpha1SQLAction and assigns it to the Actions field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink) SetActions(v []GithubComArgoprojLabsArgoDataflowApiV1alpha1SQLAction) {
	o.Actions = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink) GetDatabase() GithubComArgoprojLabsArgoDataflowApiV1alpha1Database {
	if o == nil || o.Database == nil {
		var ret GithubComArgoprojLabsArgoDataflowApiV1alpha1Database
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink) GetDatabaseOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1Database, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink) HasDatabase() bool {
	if o != nil && o.Database != nil {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given GithubComArgoprojLabsArgoDataflowApiV1alpha1Database and assigns it to the Database field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink) SetDatabase(v GithubComArgoprojLabsArgoDataflowApiV1alpha1Database) {
	o.Database = &v
}

func (o GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	return json.Marshal(toSerialize)
}

type NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink struct {
	value *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink
	isSet bool
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink) Get() *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink {
	return v.value
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink) Set(val *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink(val *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink) *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink {
	return &NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink{value: val, isSet: true}
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

