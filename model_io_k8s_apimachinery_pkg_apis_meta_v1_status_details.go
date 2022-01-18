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

// IoK8sApimachineryPkgApisMetaV1StatusDetails StatusDetails is a set of additional properties that MAY be set by the server to provide additional information about a response. The Reason field of a Status object defines what attributes will be set. Clients must ignore fields that do not match the defined type of each attribute, and should assume that any attribute may be empty, invalid, or under defined.
type IoK8sApimachineryPkgApisMetaV1StatusDetails struct {
	// The Causes array includes more details associated with the StatusReason failure. Not all StatusReasons may provide detailed causes.
	Causes *[]IoK8sApimachineryPkgApisMetaV1StatusCause `json:"causes,omitempty"`
	// The group attribute of the resource associated with the status StatusReason.
	Group *string `json:"group,omitempty"`
	// The kind attribute of the resource associated with the status StatusReason. On some operations may differ from the requested resource Kind. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	// The name attribute of the resource associated with the status StatusReason (when there is a single name which can be described).
	Name *string `json:"name,omitempty"`
	// If specified, the time in seconds before the operation should be retried. Some errors may indicate the client must take an alternate action - for those errors this field may indicate how long to wait before taking the alternate action.
	RetryAfterSeconds *int32 `json:"retryAfterSeconds,omitempty"`
	// UID of the resource. (when there is a single resource which can be described). More info: http://kubernetes.io/docs/user-guide/identifiers#uids
	Uid *string `json:"uid,omitempty"`
}

// NewIoK8sApimachineryPkgApisMetaV1StatusDetails instantiates a new IoK8sApimachineryPkgApisMetaV1StatusDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoK8sApimachineryPkgApisMetaV1StatusDetails() *IoK8sApimachineryPkgApisMetaV1StatusDetails {
	this := IoK8sApimachineryPkgApisMetaV1StatusDetails{}
	return &this
}

// NewIoK8sApimachineryPkgApisMetaV1StatusDetailsWithDefaults instantiates a new IoK8sApimachineryPkgApisMetaV1StatusDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoK8sApimachineryPkgApisMetaV1StatusDetailsWithDefaults() *IoK8sApimachineryPkgApisMetaV1StatusDetails {
	this := IoK8sApimachineryPkgApisMetaV1StatusDetails{}
	return &this
}

// GetCauses returns the Causes field value if set, zero value otherwise.
func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) GetCauses() []IoK8sApimachineryPkgApisMetaV1StatusCause {
	if o == nil || o.Causes == nil {
		var ret []IoK8sApimachineryPkgApisMetaV1StatusCause
		return ret
	}
	return *o.Causes
}

// GetCausesOk returns a tuple with the Causes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) GetCausesOk() (*[]IoK8sApimachineryPkgApisMetaV1StatusCause, bool) {
	if o == nil || o.Causes == nil {
		return nil, false
	}
	return o.Causes, true
}

// HasCauses returns a boolean if a field has been set.
func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) HasCauses() bool {
	if o != nil && o.Causes != nil {
		return true
	}

	return false
}

// SetCauses gets a reference to the given []IoK8sApimachineryPkgApisMetaV1StatusCause and assigns it to the Causes field.
func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) SetCauses(v []IoK8sApimachineryPkgApisMetaV1StatusCause) {
	o.Causes = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) SetGroup(v string) {
	o.Group = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) SetKind(v string) {
	o.Kind = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) SetName(v string) {
	o.Name = &v
}

// GetRetryAfterSeconds returns the RetryAfterSeconds field value if set, zero value otherwise.
func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) GetRetryAfterSeconds() int32 {
	if o == nil || o.RetryAfterSeconds == nil {
		var ret int32
		return ret
	}
	return *o.RetryAfterSeconds
}

// GetRetryAfterSecondsOk returns a tuple with the RetryAfterSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) GetRetryAfterSecondsOk() (*int32, bool) {
	if o == nil || o.RetryAfterSeconds == nil {
		return nil, false
	}
	return o.RetryAfterSeconds, true
}

// HasRetryAfterSeconds returns a boolean if a field has been set.
func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) HasRetryAfterSeconds() bool {
	if o != nil && o.RetryAfterSeconds != nil {
		return true
	}

	return false
}

// SetRetryAfterSeconds gets a reference to the given int32 and assigns it to the RetryAfterSeconds field.
func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) SetRetryAfterSeconds(v int32) {
	o.RetryAfterSeconds = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) GetUid() string {
	if o == nil || o.Uid == nil {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) GetUidOk() (*string, bool) {
	if o == nil || o.Uid == nil {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) HasUid() bool {
	if o != nil && o.Uid != nil {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *IoK8sApimachineryPkgApisMetaV1StatusDetails) SetUid(v string) {
	o.Uid = &v
}

func (o IoK8sApimachineryPkgApisMetaV1StatusDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Causes != nil {
		toSerialize["causes"] = o.Causes
	}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RetryAfterSeconds != nil {
		toSerialize["retryAfterSeconds"] = o.RetryAfterSeconds
	}
	if o.Uid != nil {
		toSerialize["uid"] = o.Uid
	}
	return json.Marshal(toSerialize)
}

type NullableIoK8sApimachineryPkgApisMetaV1StatusDetails struct {
	value *IoK8sApimachineryPkgApisMetaV1StatusDetails
	isSet bool
}

func (v NullableIoK8sApimachineryPkgApisMetaV1StatusDetails) Get() *IoK8sApimachineryPkgApisMetaV1StatusDetails {
	return v.value
}

func (v *NullableIoK8sApimachineryPkgApisMetaV1StatusDetails) Set(val *IoK8sApimachineryPkgApisMetaV1StatusDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableIoK8sApimachineryPkgApisMetaV1StatusDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableIoK8sApimachineryPkgApisMetaV1StatusDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoK8sApimachineryPkgApisMetaV1StatusDetails(val *IoK8sApimachineryPkgApisMetaV1StatusDetails) *NullableIoK8sApimachineryPkgApisMetaV1StatusDetails {
	return &NullableIoK8sApimachineryPkgApisMetaV1StatusDetails{value: val, isSet: true}
}

func (v NullableIoK8sApimachineryPkgApisMetaV1StatusDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoK8sApimachineryPkgApisMetaV1StatusDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

