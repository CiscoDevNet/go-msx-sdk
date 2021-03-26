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

// DeviceTemplateAccess struct for DeviceTemplateAccess
type DeviceTemplateAccess struct {
	// List of tenants to grant access. Mutually exclusive from global flag field.
	TenantIds *[]string `json:"tenantIds,omitempty"`
	// Determines if the template is globally accessible. Mutually exclusive from tenant list field.
	Global *bool `json:"global,omitempty"`
}

// NewDeviceTemplateAccess instantiates a new DeviceTemplateAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceTemplateAccess() *DeviceTemplateAccess {
	this := DeviceTemplateAccess{}
	return &this
}

// NewDeviceTemplateAccessWithDefaults instantiates a new DeviceTemplateAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceTemplateAccessWithDefaults() *DeviceTemplateAccess {
	this := DeviceTemplateAccess{}
	return &this
}

// GetTenantIds returns the TenantIds field value if set, zero value otherwise.
func (o *DeviceTemplateAccess) GetTenantIds() []string {
	if o == nil || o.TenantIds == nil {
		var ret []string
		return ret
	}
	return *o.TenantIds
}

// GetTenantIdsOk returns a tuple with the TenantIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTemplateAccess) GetTenantIdsOk() (*[]string, bool) {
	if o == nil || o.TenantIds == nil {
		return nil, false
	}
	return o.TenantIds, true
}

// HasTenantIds returns a boolean if a field has been set.
func (o *DeviceTemplateAccess) HasTenantIds() bool {
	if o != nil && o.TenantIds != nil {
		return true
	}

	return false
}

// SetTenantIds gets a reference to the given []string and assigns it to the TenantIds field.
func (o *DeviceTemplateAccess) SetTenantIds(v []string) {
	o.TenantIds = &v
}

// GetGlobal returns the Global field value if set, zero value otherwise.
func (o *DeviceTemplateAccess) GetGlobal() bool {
	if o == nil || o.Global == nil {
		var ret bool
		return ret
	}
	return *o.Global
}

// GetGlobalOk returns a tuple with the Global field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTemplateAccess) GetGlobalOk() (*bool, bool) {
	if o == nil || o.Global == nil {
		return nil, false
	}
	return o.Global, true
}

// HasGlobal returns a boolean if a field has been set.
func (o *DeviceTemplateAccess) HasGlobal() bool {
	if o != nil && o.Global != nil {
		return true
	}

	return false
}

// SetGlobal gets a reference to the given bool and assigns it to the Global field.
func (o *DeviceTemplateAccess) SetGlobal(v bool) {
	o.Global = &v
}

func (o DeviceTemplateAccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TenantIds != nil {
		toSerialize["tenantIds"] = o.TenantIds
	}
	if o.Global != nil {
		toSerialize["global"] = o.Global
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceTemplateAccess struct {
	value *DeviceTemplateAccess
	isSet bool
}

func (v NullableDeviceTemplateAccess) Get() *DeviceTemplateAccess {
	return v.value
}

func (v *NullableDeviceTemplateAccess) Set(val *DeviceTemplateAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceTemplateAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceTemplateAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceTemplateAccess(val *DeviceTemplateAccess) *NullableDeviceTemplateAccess {
	return &NullableDeviceTemplateAccess{value: val, isSet: true}
}

func (v NullableDeviceTemplateAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceTemplateAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


