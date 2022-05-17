/*
MSX SDK

MSX SDK client.

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msxsdk

import (
	"encoding/json"
)

// DeviceTemplateAttachRequest struct for DeviceTemplateAttachRequest
type DeviceTemplateAttachRequest struct {
	TemplateDetails []DeviceTemplateDetails `json:"templateDetails,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceTemplateAttachRequest DeviceTemplateAttachRequest

// NewDeviceTemplateAttachRequest instantiates a new DeviceTemplateAttachRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceTemplateAttachRequest() *DeviceTemplateAttachRequest {
	this := DeviceTemplateAttachRequest{}
	return &this
}

// NewDeviceTemplateAttachRequestWithDefaults instantiates a new DeviceTemplateAttachRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceTemplateAttachRequestWithDefaults() *DeviceTemplateAttachRequest {
	this := DeviceTemplateAttachRequest{}
	return &this
}

// GetTemplateDetails returns the TemplateDetails field value if set, zero value otherwise.
func (o *DeviceTemplateAttachRequest) GetTemplateDetails() []DeviceTemplateDetails {
	if o == nil || o.TemplateDetails == nil {
		var ret []DeviceTemplateDetails
		return ret
	}
	return o.TemplateDetails
}

// GetTemplateDetailsOk returns a tuple with the TemplateDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTemplateAttachRequest) GetTemplateDetailsOk() ([]DeviceTemplateDetails, bool) {
	if o == nil || o.TemplateDetails == nil {
		return nil, false
	}
	return o.TemplateDetails, true
}

// HasTemplateDetails returns a boolean if a field has been set.
func (o *DeviceTemplateAttachRequest) HasTemplateDetails() bool {
	if o != nil && o.TemplateDetails != nil {
		return true
	}

	return false
}

// SetTemplateDetails gets a reference to the given []DeviceTemplateDetails and assigns it to the TemplateDetails field.
func (o *DeviceTemplateAttachRequest) SetTemplateDetails(v []DeviceTemplateDetails) {
	o.TemplateDetails = v
}

func (o DeviceTemplateAttachRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TemplateDetails != nil {
		toSerialize["templateDetails"] = o.TemplateDetails
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeviceTemplateAttachRequest) UnmarshalJSON(bytes []byte) (err error) {
	varDeviceTemplateAttachRequest := _DeviceTemplateAttachRequest{}

	if err = json.Unmarshal(bytes, &varDeviceTemplateAttachRequest); err == nil {
		*o = DeviceTemplateAttachRequest(varDeviceTemplateAttachRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "templateDetails")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceTemplateAttachRequest struct {
	value *DeviceTemplateAttachRequest
	isSet bool
}

func (v NullableDeviceTemplateAttachRequest) Get() *DeviceTemplateAttachRequest {
	return v.value
}

func (v *NullableDeviceTemplateAttachRequest) Set(val *DeviceTemplateAttachRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceTemplateAttachRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceTemplateAttachRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceTemplateAttachRequest(val *DeviceTemplateAttachRequest) *NullableDeviceTemplateAttachRequest {
	return &NullableDeviceTemplateAttachRequest{value: val, isSet: true}
}

func (v NullableDeviceTemplateAttachRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceTemplateAttachRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


