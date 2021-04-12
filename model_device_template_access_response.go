/*
 * KAKAPO - MSX SDK
 *
 * MSX Platform SDK
 *
 * API version: 1.0.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msxsdk

import (
	"encoding/json"
)

// DeviceTemplateAccessResponse struct for DeviceTemplateAccessResponse
type DeviceTemplateAccessResponse struct {
	FailureListOfTenants []string `json:"failureListOfTenants,omitempty"`
	Global NullableBool `json:"global,omitempty"`
	SuccessListOfTenants []string `json:"successListOfTenants,omitempty"`
}

// NewDeviceTemplateAccessResponse instantiates a new DeviceTemplateAccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceTemplateAccessResponse() *DeviceTemplateAccessResponse {
	this := DeviceTemplateAccessResponse{}
	return &this
}

// NewDeviceTemplateAccessResponseWithDefaults instantiates a new DeviceTemplateAccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceTemplateAccessResponseWithDefaults() *DeviceTemplateAccessResponse {
	this := DeviceTemplateAccessResponse{}
	return &this
}

// GetFailureListOfTenants returns the FailureListOfTenants field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceTemplateAccessResponse) GetFailureListOfTenants() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.FailureListOfTenants
}

// GetFailureListOfTenantsOk returns a tuple with the FailureListOfTenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceTemplateAccessResponse) GetFailureListOfTenantsOk() (*[]string, bool) {
	if o == nil || o.FailureListOfTenants == nil {
		return nil, false
	}
	return &o.FailureListOfTenants, true
}

// HasFailureListOfTenants returns a boolean if a field has been set.
func (o *DeviceTemplateAccessResponse) HasFailureListOfTenants() bool {
	if o != nil && o.FailureListOfTenants != nil {
		return true
	}

	return false
}

// SetFailureListOfTenants gets a reference to the given []string and assigns it to the FailureListOfTenants field.
func (o *DeviceTemplateAccessResponse) SetFailureListOfTenants(v []string) {
	o.FailureListOfTenants = v
}

// GetGlobal returns the Global field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceTemplateAccessResponse) GetGlobal() bool {
	if o == nil || o.Global.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Global.Get()
}

// GetGlobalOk returns a tuple with the Global field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceTemplateAccessResponse) GetGlobalOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Global.Get(), o.Global.IsSet()
}

// HasGlobal returns a boolean if a field has been set.
func (o *DeviceTemplateAccessResponse) HasGlobal() bool {
	if o != nil && o.Global.IsSet() {
		return true
	}

	return false
}

// SetGlobal gets a reference to the given NullableBool and assigns it to the Global field.
func (o *DeviceTemplateAccessResponse) SetGlobal(v bool) {
	o.Global.Set(&v)
}
// SetGlobalNil sets the value for Global to be an explicit nil
func (o *DeviceTemplateAccessResponse) SetGlobalNil() {
	o.Global.Set(nil)
}

// UnsetGlobal ensures that no value is present for Global, not even an explicit nil
func (o *DeviceTemplateAccessResponse) UnsetGlobal() {
	o.Global.Unset()
}

// GetSuccessListOfTenants returns the SuccessListOfTenants field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceTemplateAccessResponse) GetSuccessListOfTenants() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.SuccessListOfTenants
}

// GetSuccessListOfTenantsOk returns a tuple with the SuccessListOfTenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceTemplateAccessResponse) GetSuccessListOfTenantsOk() (*[]string, bool) {
	if o == nil || o.SuccessListOfTenants == nil {
		return nil, false
	}
	return &o.SuccessListOfTenants, true
}

// HasSuccessListOfTenants returns a boolean if a field has been set.
func (o *DeviceTemplateAccessResponse) HasSuccessListOfTenants() bool {
	if o != nil && o.SuccessListOfTenants != nil {
		return true
	}

	return false
}

// SetSuccessListOfTenants gets a reference to the given []string and assigns it to the SuccessListOfTenants field.
func (o *DeviceTemplateAccessResponse) SetSuccessListOfTenants(v []string) {
	o.SuccessListOfTenants = v
}

func (o DeviceTemplateAccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FailureListOfTenants != nil {
		toSerialize["failureListOfTenants"] = o.FailureListOfTenants
	}
	if o.Global.IsSet() {
		toSerialize["global"] = o.Global.Get()
	}
	if o.SuccessListOfTenants != nil {
		toSerialize["successListOfTenants"] = o.SuccessListOfTenants
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceTemplateAccessResponse struct {
	value *DeviceTemplateAccessResponse
	isSet bool
}

func (v NullableDeviceTemplateAccessResponse) Get() *DeviceTemplateAccessResponse {
	return v.value
}

func (v *NullableDeviceTemplateAccessResponse) Set(val *DeviceTemplateAccessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceTemplateAccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceTemplateAccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceTemplateAccessResponse(val *DeviceTemplateAccessResponse) *NullableDeviceTemplateAccessResponse {
	return &NullableDeviceTemplateAccessResponse{value: val, isSet: true}
}

func (v NullableDeviceTemplateAccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceTemplateAccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


