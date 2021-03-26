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

// DeviceTemplateUpdateRequest struct for DeviceTemplateUpdateRequest
type DeviceTemplateUpdateRequest struct {
	TemplateDetails *[]DeviceTemplateUpdateDetails `json:"templateDetails,omitempty"`
}

// NewDeviceTemplateUpdateRequest instantiates a new DeviceTemplateUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceTemplateUpdateRequest() *DeviceTemplateUpdateRequest {
	this := DeviceTemplateUpdateRequest{}
	return &this
}

// NewDeviceTemplateUpdateRequestWithDefaults instantiates a new DeviceTemplateUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceTemplateUpdateRequestWithDefaults() *DeviceTemplateUpdateRequest {
	this := DeviceTemplateUpdateRequest{}
	return &this
}

// GetTemplateDetails returns the TemplateDetails field value if set, zero value otherwise.
func (o *DeviceTemplateUpdateRequest) GetTemplateDetails() []DeviceTemplateUpdateDetails {
	if o == nil || o.TemplateDetails == nil {
		var ret []DeviceTemplateUpdateDetails
		return ret
	}
	return *o.TemplateDetails
}

// GetTemplateDetailsOk returns a tuple with the TemplateDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTemplateUpdateRequest) GetTemplateDetailsOk() (*[]DeviceTemplateUpdateDetails, bool) {
	if o == nil || o.TemplateDetails == nil {
		return nil, false
	}
	return o.TemplateDetails, true
}

// HasTemplateDetails returns a boolean if a field has been set.
func (o *DeviceTemplateUpdateRequest) HasTemplateDetails() bool {
	if o != nil && o.TemplateDetails != nil {
		return true
	}

	return false
}

// SetTemplateDetails gets a reference to the given []DeviceTemplateUpdateDetails and assigns it to the TemplateDetails field.
func (o *DeviceTemplateUpdateRequest) SetTemplateDetails(v []DeviceTemplateUpdateDetails) {
	o.TemplateDetails = &v
}

func (o DeviceTemplateUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TemplateDetails != nil {
		toSerialize["templateDetails"] = o.TemplateDetails
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceTemplateUpdateRequest struct {
	value *DeviceTemplateUpdateRequest
	isSet bool
}

func (v NullableDeviceTemplateUpdateRequest) Get() *DeviceTemplateUpdateRequest {
	return v.value
}

func (v *NullableDeviceTemplateUpdateRequest) Set(val *DeviceTemplateUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceTemplateUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceTemplateUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceTemplateUpdateRequest(val *DeviceTemplateUpdateRequest) *NullableDeviceTemplateUpdateRequest {
	return &NullableDeviceTemplateUpdateRequest{value: val, isSet: true}
}

func (v NullableDeviceTemplateUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceTemplateUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


