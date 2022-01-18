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

// GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource struct for GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource
type GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource struct {
	Kafka *GithubComArgoprojLabsArgoDataflowApiV1alpha1Kafka `json:"kafka,omitempty"`
	StartOffset *string `json:"startOffset,omitempty"`
}

// NewGithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource() *GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource {
	this := GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource{}
	return &this
}

// NewGithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSourceWithDefaults instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSourceWithDefaults() *GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource {
	this := GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource{}
	return &this
}

// GetKafka returns the Kafka field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource) GetKafka() GithubComArgoprojLabsArgoDataflowApiV1alpha1Kafka {
	if o == nil || o.Kafka == nil {
		var ret GithubComArgoprojLabsArgoDataflowApiV1alpha1Kafka
		return ret
	}
	return *o.Kafka
}

// GetKafkaOk returns a tuple with the Kafka field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource) GetKafkaOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1Kafka, bool) {
	if o == nil || o.Kafka == nil {
		return nil, false
	}
	return o.Kafka, true
}

// HasKafka returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource) HasKafka() bool {
	if o != nil && o.Kafka != nil {
		return true
	}

	return false
}

// SetKafka gets a reference to the given GithubComArgoprojLabsArgoDataflowApiV1alpha1Kafka and assigns it to the Kafka field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource) SetKafka(v GithubComArgoprojLabsArgoDataflowApiV1alpha1Kafka) {
	o.Kafka = &v
}

// GetStartOffset returns the StartOffset field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource) GetStartOffset() string {
	if o == nil || o.StartOffset == nil {
		var ret string
		return ret
	}
	return *o.StartOffset
}

// GetStartOffsetOk returns a tuple with the StartOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource) GetStartOffsetOk() (*string, bool) {
	if o == nil || o.StartOffset == nil {
		return nil, false
	}
	return o.StartOffset, true
}

// HasStartOffset returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource) HasStartOffset() bool {
	if o != nil && o.StartOffset != nil {
		return true
	}

	return false
}

// SetStartOffset gets a reference to the given string and assigns it to the StartOffset field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource) SetStartOffset(v string) {
	o.StartOffset = &v
}

func (o GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Kafka != nil {
		toSerialize["kafka"] = o.Kafka
	}
	if o.StartOffset != nil {
		toSerialize["startOffset"] = o.StartOffset
	}
	return json.Marshal(toSerialize)
}

type NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource struct {
	value *GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource
	isSet bool
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource) Get() *GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource {
	return v.value
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource) Set(val *GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource(val *GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource) *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource {
	return &NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource{value: val, isSet: true}
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

