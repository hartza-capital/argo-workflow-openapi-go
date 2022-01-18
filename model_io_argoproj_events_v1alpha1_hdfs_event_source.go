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

// IoArgoprojEventsV1alpha1HDFSEventSource struct for IoArgoprojEventsV1alpha1HDFSEventSource
type IoArgoprojEventsV1alpha1HDFSEventSource struct {
	Addresses *[]string `json:"addresses,omitempty"`
	CheckInterval *string `json:"checkInterval,omitempty"`
	// HDFSUser is the user to access HDFS file system. It is ignored if either ccache or keytab is used.
	HdfsUser *string `json:"hdfsUser,omitempty"`
	KrbCCacheSecret *IoK8sApiCoreV1SecretKeySelector `json:"krbCCacheSecret,omitempty"`
	KrbConfigConfigMap *IoK8sApiCoreV1ConfigMapKeySelector `json:"krbConfigConfigMap,omitempty"`
	KrbKeytabSecret *IoK8sApiCoreV1SecretKeySelector `json:"krbKeytabSecret,omitempty"`
	// KrbRealm is the Kerberos realm used with Kerberos keytab It must be set if keytab is used.
	KrbRealm *string `json:"krbRealm,omitempty"`
	// KrbServicePrincipalName is the principal name of Kerberos service It must be set if either ccache or keytab is used.
	KrbServicePrincipalName *string `json:"krbServicePrincipalName,omitempty"`
	// KrbUsername is the Kerberos username used with Kerberos keytab It must be set if keytab is used.
	KrbUsername *string `json:"krbUsername,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
	Type *string `json:"type,omitempty"`
	WatchPathConfig *IoArgoprojEventsV1alpha1WatchPathConfig `json:"watchPathConfig,omitempty"`
}

// NewIoArgoprojEventsV1alpha1HDFSEventSource instantiates a new IoArgoprojEventsV1alpha1HDFSEventSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojEventsV1alpha1HDFSEventSource() *IoArgoprojEventsV1alpha1HDFSEventSource {
	this := IoArgoprojEventsV1alpha1HDFSEventSource{}
	return &this
}

// NewIoArgoprojEventsV1alpha1HDFSEventSourceWithDefaults instantiates a new IoArgoprojEventsV1alpha1HDFSEventSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojEventsV1alpha1HDFSEventSourceWithDefaults() *IoArgoprojEventsV1alpha1HDFSEventSource {
	this := IoArgoprojEventsV1alpha1HDFSEventSource{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetAddresses() []string {
	if o == nil || o.Addresses == nil {
		var ret []string
		return ret
	}
	return *o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetAddressesOk() (*[]string, bool) {
	if o == nil || o.Addresses == nil {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) HasAddresses() bool {
	if o != nil && o.Addresses != nil {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []string and assigns it to the Addresses field.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) SetAddresses(v []string) {
	o.Addresses = &v
}

// GetCheckInterval returns the CheckInterval field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetCheckInterval() string {
	if o == nil || o.CheckInterval == nil {
		var ret string
		return ret
	}
	return *o.CheckInterval
}

// GetCheckIntervalOk returns a tuple with the CheckInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetCheckIntervalOk() (*string, bool) {
	if o == nil || o.CheckInterval == nil {
		return nil, false
	}
	return o.CheckInterval, true
}

// HasCheckInterval returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) HasCheckInterval() bool {
	if o != nil && o.CheckInterval != nil {
		return true
	}

	return false
}

// SetCheckInterval gets a reference to the given string and assigns it to the CheckInterval field.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) SetCheckInterval(v string) {
	o.CheckInterval = &v
}

// GetHdfsUser returns the HdfsUser field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetHdfsUser() string {
	if o == nil || o.HdfsUser == nil {
		var ret string
		return ret
	}
	return *o.HdfsUser
}

// GetHdfsUserOk returns a tuple with the HdfsUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetHdfsUserOk() (*string, bool) {
	if o == nil || o.HdfsUser == nil {
		return nil, false
	}
	return o.HdfsUser, true
}

// HasHdfsUser returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) HasHdfsUser() bool {
	if o != nil && o.HdfsUser != nil {
		return true
	}

	return false
}

// SetHdfsUser gets a reference to the given string and assigns it to the HdfsUser field.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) SetHdfsUser(v string) {
	o.HdfsUser = &v
}

// GetKrbCCacheSecret returns the KrbCCacheSecret field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetKrbCCacheSecret() IoK8sApiCoreV1SecretKeySelector {
	if o == nil || o.KrbCCacheSecret == nil {
		var ret IoK8sApiCoreV1SecretKeySelector
		return ret
	}
	return *o.KrbCCacheSecret
}

// GetKrbCCacheSecretOk returns a tuple with the KrbCCacheSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetKrbCCacheSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool) {
	if o == nil || o.KrbCCacheSecret == nil {
		return nil, false
	}
	return o.KrbCCacheSecret, true
}

// HasKrbCCacheSecret returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) HasKrbCCacheSecret() bool {
	if o != nil && o.KrbCCacheSecret != nil {
		return true
	}

	return false
}

// SetKrbCCacheSecret gets a reference to the given IoK8sApiCoreV1SecretKeySelector and assigns it to the KrbCCacheSecret field.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) SetKrbCCacheSecret(v IoK8sApiCoreV1SecretKeySelector) {
	o.KrbCCacheSecret = &v
}

// GetKrbConfigConfigMap returns the KrbConfigConfigMap field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetKrbConfigConfigMap() IoK8sApiCoreV1ConfigMapKeySelector {
	if o == nil || o.KrbConfigConfigMap == nil {
		var ret IoK8sApiCoreV1ConfigMapKeySelector
		return ret
	}
	return *o.KrbConfigConfigMap
}

// GetKrbConfigConfigMapOk returns a tuple with the KrbConfigConfigMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetKrbConfigConfigMapOk() (*IoK8sApiCoreV1ConfigMapKeySelector, bool) {
	if o == nil || o.KrbConfigConfigMap == nil {
		return nil, false
	}
	return o.KrbConfigConfigMap, true
}

// HasKrbConfigConfigMap returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) HasKrbConfigConfigMap() bool {
	if o != nil && o.KrbConfigConfigMap != nil {
		return true
	}

	return false
}

// SetKrbConfigConfigMap gets a reference to the given IoK8sApiCoreV1ConfigMapKeySelector and assigns it to the KrbConfigConfigMap field.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) SetKrbConfigConfigMap(v IoK8sApiCoreV1ConfigMapKeySelector) {
	o.KrbConfigConfigMap = &v
}

// GetKrbKeytabSecret returns the KrbKeytabSecret field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetKrbKeytabSecret() IoK8sApiCoreV1SecretKeySelector {
	if o == nil || o.KrbKeytabSecret == nil {
		var ret IoK8sApiCoreV1SecretKeySelector
		return ret
	}
	return *o.KrbKeytabSecret
}

// GetKrbKeytabSecretOk returns a tuple with the KrbKeytabSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetKrbKeytabSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool) {
	if o == nil || o.KrbKeytabSecret == nil {
		return nil, false
	}
	return o.KrbKeytabSecret, true
}

// HasKrbKeytabSecret returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) HasKrbKeytabSecret() bool {
	if o != nil && o.KrbKeytabSecret != nil {
		return true
	}

	return false
}

// SetKrbKeytabSecret gets a reference to the given IoK8sApiCoreV1SecretKeySelector and assigns it to the KrbKeytabSecret field.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) SetKrbKeytabSecret(v IoK8sApiCoreV1SecretKeySelector) {
	o.KrbKeytabSecret = &v
}

// GetKrbRealm returns the KrbRealm field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetKrbRealm() string {
	if o == nil || o.KrbRealm == nil {
		var ret string
		return ret
	}
	return *o.KrbRealm
}

// GetKrbRealmOk returns a tuple with the KrbRealm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetKrbRealmOk() (*string, bool) {
	if o == nil || o.KrbRealm == nil {
		return nil, false
	}
	return o.KrbRealm, true
}

// HasKrbRealm returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) HasKrbRealm() bool {
	if o != nil && o.KrbRealm != nil {
		return true
	}

	return false
}

// SetKrbRealm gets a reference to the given string and assigns it to the KrbRealm field.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) SetKrbRealm(v string) {
	o.KrbRealm = &v
}

// GetKrbServicePrincipalName returns the KrbServicePrincipalName field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetKrbServicePrincipalName() string {
	if o == nil || o.KrbServicePrincipalName == nil {
		var ret string
		return ret
	}
	return *o.KrbServicePrincipalName
}

// GetKrbServicePrincipalNameOk returns a tuple with the KrbServicePrincipalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetKrbServicePrincipalNameOk() (*string, bool) {
	if o == nil || o.KrbServicePrincipalName == nil {
		return nil, false
	}
	return o.KrbServicePrincipalName, true
}

// HasKrbServicePrincipalName returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) HasKrbServicePrincipalName() bool {
	if o != nil && o.KrbServicePrincipalName != nil {
		return true
	}

	return false
}

// SetKrbServicePrincipalName gets a reference to the given string and assigns it to the KrbServicePrincipalName field.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) SetKrbServicePrincipalName(v string) {
	o.KrbServicePrincipalName = &v
}

// GetKrbUsername returns the KrbUsername field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetKrbUsername() string {
	if o == nil || o.KrbUsername == nil {
		var ret string
		return ret
	}
	return *o.KrbUsername
}

// GetKrbUsernameOk returns a tuple with the KrbUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetKrbUsernameOk() (*string, bool) {
	if o == nil || o.KrbUsername == nil {
		return nil, false
	}
	return o.KrbUsername, true
}

// HasKrbUsername returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) HasKrbUsername() bool {
	if o != nil && o.KrbUsername != nil {
		return true
	}

	return false
}

// SetKrbUsername gets a reference to the given string and assigns it to the KrbUsername field.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) SetKrbUsername(v string) {
	o.KrbUsername = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) SetType(v string) {
	o.Type = &v
}

// GetWatchPathConfig returns the WatchPathConfig field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetWatchPathConfig() IoArgoprojEventsV1alpha1WatchPathConfig {
	if o == nil || o.WatchPathConfig == nil {
		var ret IoArgoprojEventsV1alpha1WatchPathConfig
		return ret
	}
	return *o.WatchPathConfig
}

// GetWatchPathConfigOk returns a tuple with the WatchPathConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetWatchPathConfigOk() (*IoArgoprojEventsV1alpha1WatchPathConfig, bool) {
	if o == nil || o.WatchPathConfig == nil {
		return nil, false
	}
	return o.WatchPathConfig, true
}

// HasWatchPathConfig returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) HasWatchPathConfig() bool {
	if o != nil && o.WatchPathConfig != nil {
		return true
	}

	return false
}

// SetWatchPathConfig gets a reference to the given IoArgoprojEventsV1alpha1WatchPathConfig and assigns it to the WatchPathConfig field.
func (o *IoArgoprojEventsV1alpha1HDFSEventSource) SetWatchPathConfig(v IoArgoprojEventsV1alpha1WatchPathConfig) {
	o.WatchPathConfig = &v
}

func (o IoArgoprojEventsV1alpha1HDFSEventSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Addresses != nil {
		toSerialize["addresses"] = o.Addresses
	}
	if o.CheckInterval != nil {
		toSerialize["checkInterval"] = o.CheckInterval
	}
	if o.HdfsUser != nil {
		toSerialize["hdfsUser"] = o.HdfsUser
	}
	if o.KrbCCacheSecret != nil {
		toSerialize["krbCCacheSecret"] = o.KrbCCacheSecret
	}
	if o.KrbConfigConfigMap != nil {
		toSerialize["krbConfigConfigMap"] = o.KrbConfigConfigMap
	}
	if o.KrbKeytabSecret != nil {
		toSerialize["krbKeytabSecret"] = o.KrbKeytabSecret
	}
	if o.KrbRealm != nil {
		toSerialize["krbRealm"] = o.KrbRealm
	}
	if o.KrbServicePrincipalName != nil {
		toSerialize["krbServicePrincipalName"] = o.KrbServicePrincipalName
	}
	if o.KrbUsername != nil {
		toSerialize["krbUsername"] = o.KrbUsername
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.WatchPathConfig != nil {
		toSerialize["watchPathConfig"] = o.WatchPathConfig
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojEventsV1alpha1HDFSEventSource struct {
	value *IoArgoprojEventsV1alpha1HDFSEventSource
	isSet bool
}

func (v NullableIoArgoprojEventsV1alpha1HDFSEventSource) Get() *IoArgoprojEventsV1alpha1HDFSEventSource {
	return v.value
}

func (v *NullableIoArgoprojEventsV1alpha1HDFSEventSource) Set(val *IoArgoprojEventsV1alpha1HDFSEventSource) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojEventsV1alpha1HDFSEventSource) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojEventsV1alpha1HDFSEventSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojEventsV1alpha1HDFSEventSource(val *IoArgoprojEventsV1alpha1HDFSEventSource) *NullableIoArgoprojEventsV1alpha1HDFSEventSource {
	return &NullableIoArgoprojEventsV1alpha1HDFSEventSource{value: val, isSet: true}
}

func (v NullableIoArgoprojEventsV1alpha1HDFSEventSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojEventsV1alpha1HDFSEventSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

