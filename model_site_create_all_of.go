/*
 * KAKAPO - MSX SDK
 *
 * MSX Platform SDK
 *
 * API version: 1.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msxsdk

import (
	"encoding/json"
)

// SiteCreateAllOf struct for SiteCreateAllOf
type SiteCreateAllOf struct {
	TenantId string `json:"tenantId"`
	DeviceIds *[]string `json:"deviceIds,omitempty"`
	ServiceIds *[]string `json:"serviceIds,omitempty"`
}

// NewSiteCreateAllOf instantiates a new SiteCreateAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteCreateAllOf(tenantId string) *SiteCreateAllOf {
	this := SiteCreateAllOf{}
	this.TenantId = tenantId
	return &this
}

// NewSiteCreateAllOfWithDefaults instantiates a new SiteCreateAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteCreateAllOfWithDefaults() *SiteCreateAllOf {
	this := SiteCreateAllOf{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *SiteCreateAllOf) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *SiteCreateAllOf) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *SiteCreateAllOf) SetTenantId(v string) {
	o.TenantId = v
}

// GetDeviceIds returns the DeviceIds field value if set, zero value otherwise.
func (o *SiteCreateAllOf) GetDeviceIds() []string {
	if o == nil || o.DeviceIds == nil {
		var ret []string
		return ret
	}
	return *o.DeviceIds
}

// GetDeviceIdsOk returns a tuple with the DeviceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteCreateAllOf) GetDeviceIdsOk() (*[]string, bool) {
	if o == nil || o.DeviceIds == nil {
		return nil, false
	}
	return o.DeviceIds, true
}

// HasDeviceIds returns a boolean if a field has been set.
func (o *SiteCreateAllOf) HasDeviceIds() bool {
	if o != nil && o.DeviceIds != nil {
		return true
	}

	return false
}

// SetDeviceIds gets a reference to the given []string and assigns it to the DeviceIds field.
func (o *SiteCreateAllOf) SetDeviceIds(v []string) {
	o.DeviceIds = &v
}

// GetServiceIds returns the ServiceIds field value if set, zero value otherwise.
func (o *SiteCreateAllOf) GetServiceIds() []string {
	if o == nil || o.ServiceIds == nil {
		var ret []string
		return ret
	}
	return *o.ServiceIds
}

// GetServiceIdsOk returns a tuple with the ServiceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteCreateAllOf) GetServiceIdsOk() (*[]string, bool) {
	if o == nil || o.ServiceIds == nil {
		return nil, false
	}
	return o.ServiceIds, true
}

// HasServiceIds returns a boolean if a field has been set.
func (o *SiteCreateAllOf) HasServiceIds() bool {
	if o != nil && o.ServiceIds != nil {
		return true
	}

	return false
}

// SetServiceIds gets a reference to the given []string and assigns it to the ServiceIds field.
func (o *SiteCreateAllOf) SetServiceIds(v []string) {
	o.ServiceIds = &v
}

func (o SiteCreateAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tenantId"] = o.TenantId
	}
	if o.DeviceIds != nil {
		toSerialize["deviceIds"] = o.DeviceIds
	}
	if o.ServiceIds != nil {
		toSerialize["serviceIds"] = o.ServiceIds
	}
	return json.Marshal(toSerialize)
}

type NullableSiteCreateAllOf struct {
	value *SiteCreateAllOf
	isSet bool
}

func (v NullableSiteCreateAllOf) Get() *SiteCreateAllOf {
	return v.value
}

func (v *NullableSiteCreateAllOf) Set(val *SiteCreateAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteCreateAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteCreateAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteCreateAllOf(val *SiteCreateAllOf) *NullableSiteCreateAllOf {
	return &NullableSiteCreateAllOf{value: val, isSet: true}
}

func (v NullableSiteCreateAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteCreateAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


