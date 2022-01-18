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

// IoArgoprojEventsV1alpha1StorageGridEventSource struct for IoArgoprojEventsV1alpha1StorageGridEventSource
type IoArgoprojEventsV1alpha1StorageGridEventSource struct {
	// APIURL is the url of the storagegrid api.
	ApiURL *string `json:"apiURL,omitempty"`
	AuthToken *IoK8sApiCoreV1SecretKeySelector `json:"authToken,omitempty"`
	// Name of the bucket to register notifications for.
	Bucket *string `json:"bucket,omitempty"`
	Events *[]string `json:"events,omitempty"`
	Filter *IoArgoprojEventsV1alpha1StorageGridFilter `json:"filter,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
	Region *string `json:"region,omitempty"`
	TopicArn *string `json:"topicArn,omitempty"`
	Webhook *IoArgoprojEventsV1alpha1WebhookContext `json:"webhook,omitempty"`
}

// NewIoArgoprojEventsV1alpha1StorageGridEventSource instantiates a new IoArgoprojEventsV1alpha1StorageGridEventSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojEventsV1alpha1StorageGridEventSource() *IoArgoprojEventsV1alpha1StorageGridEventSource {
	this := IoArgoprojEventsV1alpha1StorageGridEventSource{}
	return &this
}

// NewIoArgoprojEventsV1alpha1StorageGridEventSourceWithDefaults instantiates a new IoArgoprojEventsV1alpha1StorageGridEventSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojEventsV1alpha1StorageGridEventSourceWithDefaults() *IoArgoprojEventsV1alpha1StorageGridEventSource {
	this := IoArgoprojEventsV1alpha1StorageGridEventSource{}
	return &this
}

// GetApiURL returns the ApiURL field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetApiURL() string {
	if o == nil || o.ApiURL == nil {
		var ret string
		return ret
	}
	return *o.ApiURL
}

// GetApiURLOk returns a tuple with the ApiURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetApiURLOk() (*string, bool) {
	if o == nil || o.ApiURL == nil {
		return nil, false
	}
	return o.ApiURL, true
}

// HasApiURL returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) HasApiURL() bool {
	if o != nil && o.ApiURL != nil {
		return true
	}

	return false
}

// SetApiURL gets a reference to the given string and assigns it to the ApiURL field.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) SetApiURL(v string) {
	o.ApiURL = &v
}

// GetAuthToken returns the AuthToken field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetAuthToken() IoK8sApiCoreV1SecretKeySelector {
	if o == nil || o.AuthToken == nil {
		var ret IoK8sApiCoreV1SecretKeySelector
		return ret
	}
	return *o.AuthToken
}

// GetAuthTokenOk returns a tuple with the AuthToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetAuthTokenOk() (*IoK8sApiCoreV1SecretKeySelector, bool) {
	if o == nil || o.AuthToken == nil {
		return nil, false
	}
	return o.AuthToken, true
}

// HasAuthToken returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) HasAuthToken() bool {
	if o != nil && o.AuthToken != nil {
		return true
	}

	return false
}

// SetAuthToken gets a reference to the given IoK8sApiCoreV1SecretKeySelector and assigns it to the AuthToken field.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) SetAuthToken(v IoK8sApiCoreV1SecretKeySelector) {
	o.AuthToken = &v
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetBucket() string {
	if o == nil || o.Bucket == nil {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetBucketOk() (*string, bool) {
	if o == nil || o.Bucket == nil {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) HasBucket() bool {
	if o != nil && o.Bucket != nil {
		return true
	}

	return false
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) SetBucket(v string) {
	o.Bucket = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetEvents() []string {
	if o == nil || o.Events == nil {
		var ret []string
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetEventsOk() (*[]string, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []string and assigns it to the Events field.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) SetEvents(v []string) {
	o.Events = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetFilter() IoArgoprojEventsV1alpha1StorageGridFilter {
	if o == nil || o.Filter == nil {
		var ret IoArgoprojEventsV1alpha1StorageGridFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetFilterOk() (*IoArgoprojEventsV1alpha1StorageGridFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given IoArgoprojEventsV1alpha1StorageGridFilter and assigns it to the Filter field.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) SetFilter(v IoArgoprojEventsV1alpha1StorageGridFilter) {
	o.Filter = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) SetRegion(v string) {
	o.Region = &v
}

// GetTopicArn returns the TopicArn field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetTopicArn() string {
	if o == nil || o.TopicArn == nil {
		var ret string
		return ret
	}
	return *o.TopicArn
}

// GetTopicArnOk returns a tuple with the TopicArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetTopicArnOk() (*string, bool) {
	if o == nil || o.TopicArn == nil {
		return nil, false
	}
	return o.TopicArn, true
}

// HasTopicArn returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) HasTopicArn() bool {
	if o != nil && o.TopicArn != nil {
		return true
	}

	return false
}

// SetTopicArn gets a reference to the given string and assigns it to the TopicArn field.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) SetTopicArn(v string) {
	o.TopicArn = &v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetWebhook() IoArgoprojEventsV1alpha1WebhookContext {
	if o == nil || o.Webhook == nil {
		var ret IoArgoprojEventsV1alpha1WebhookContext
		return ret
	}
	return *o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetWebhookOk() (*IoArgoprojEventsV1alpha1WebhookContext, bool) {
	if o == nil || o.Webhook == nil {
		return nil, false
	}
	return o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) HasWebhook() bool {
	if o != nil && o.Webhook != nil {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given IoArgoprojEventsV1alpha1WebhookContext and assigns it to the Webhook field.
func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) SetWebhook(v IoArgoprojEventsV1alpha1WebhookContext) {
	o.Webhook = &v
}

func (o IoArgoprojEventsV1alpha1StorageGridEventSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiURL != nil {
		toSerialize["apiURL"] = o.ApiURL
	}
	if o.AuthToken != nil {
		toSerialize["authToken"] = o.AuthToken
	}
	if o.Bucket != nil {
		toSerialize["bucket"] = o.Bucket
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.TopicArn != nil {
		toSerialize["topicArn"] = o.TopicArn
	}
	if o.Webhook != nil {
		toSerialize["webhook"] = o.Webhook
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojEventsV1alpha1StorageGridEventSource struct {
	value *IoArgoprojEventsV1alpha1StorageGridEventSource
	isSet bool
}

func (v NullableIoArgoprojEventsV1alpha1StorageGridEventSource) Get() *IoArgoprojEventsV1alpha1StorageGridEventSource {
	return v.value
}

func (v *NullableIoArgoprojEventsV1alpha1StorageGridEventSource) Set(val *IoArgoprojEventsV1alpha1StorageGridEventSource) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojEventsV1alpha1StorageGridEventSource) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojEventsV1alpha1StorageGridEventSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojEventsV1alpha1StorageGridEventSource(val *IoArgoprojEventsV1alpha1StorageGridEventSource) *NullableIoArgoprojEventsV1alpha1StorageGridEventSource {
	return &NullableIoArgoprojEventsV1alpha1StorageGridEventSource{value: val, isSet: true}
}

func (v NullableIoArgoprojEventsV1alpha1StorageGridEventSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojEventsV1alpha1StorageGridEventSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

