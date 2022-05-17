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

// DeviceTemplateDetails struct for DeviceTemplateDetails
type DeviceTemplateDetails struct {
	TemplateId *string `json:"templateId,omitempty"`
	TemplateParams []NameValue `json:"templateParams,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceTemplateDetails DeviceTemplateDetails

// NewDeviceTemplateDetails instantiates a new DeviceTemplateDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceTemplateDetails() *DeviceTemplateDetails {
	this := DeviceTemplateDetails{}
	return &this
}

// NewDeviceTemplateDetailsWithDefaults instantiates a new DeviceTemplateDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceTemplateDetailsWithDefaults() *DeviceTemplateDetails {
	this := DeviceTemplateDetails{}
	return &this
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *DeviceTemplateDetails) GetTemplateId() string {
	if o == nil || o.TemplateId == nil {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTemplateDetails) GetTemplateIdOk() (*string, bool) {
	if o == nil || o.TemplateId == nil {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *DeviceTemplateDetails) HasTemplateId() bool {
	if o != nil && o.TemplateId != nil {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *DeviceTemplateDetails) SetTemplateId(v string) {
	o.TemplateId = &v
}

// GetTemplateParams returns the TemplateParams field value if set, zero value otherwise.
func (o *DeviceTemplateDetails) GetTemplateParams() []NameValue {
	if o == nil || o.TemplateParams == nil {
		var ret []NameValue
		return ret
	}
	return o.TemplateParams
}

// GetTemplateParamsOk returns a tuple with the TemplateParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTemplateDetails) GetTemplateParamsOk() ([]NameValue, bool) {
	if o == nil || o.TemplateParams == nil {
		return nil, false
	}
	return o.TemplateParams, true
}

// HasTemplateParams returns a boolean if a field has been set.
func (o *DeviceTemplateDetails) HasTemplateParams() bool {
	if o != nil && o.TemplateParams != nil {
		return true
	}

	return false
}

// SetTemplateParams gets a reference to the given []NameValue and assigns it to the TemplateParams field.
func (o *DeviceTemplateDetails) SetTemplateParams(v []NameValue) {
	o.TemplateParams = v
}

func (o DeviceTemplateDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TemplateId != nil {
		toSerialize["templateId"] = o.TemplateId
	}
	if o.TemplateParams != nil {
		toSerialize["templateParams"] = o.TemplateParams
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeviceTemplateDetails) UnmarshalJSON(bytes []byte) (err error) {
	varDeviceTemplateDetails := _DeviceTemplateDetails{}

	if err = json.Unmarshal(bytes, &varDeviceTemplateDetails); err == nil {
		*o = DeviceTemplateDetails(varDeviceTemplateDetails)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "templateId")
		delete(additionalProperties, "templateParams")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceTemplateDetails struct {
	value *DeviceTemplateDetails
	isSet bool
}

func (v NullableDeviceTemplateDetails) Get() *DeviceTemplateDetails {
	return v.value
}

func (v *NullableDeviceTemplateDetails) Set(val *DeviceTemplateDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceTemplateDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceTemplateDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceTemplateDetails(val *DeviceTemplateDetails) *NullableDeviceTemplateDetails {
	return &NullableDeviceTemplateDetails{value: val, isSet: true}
}

func (v NullableDeviceTemplateDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceTemplateDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


