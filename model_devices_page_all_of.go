/*
 * MSX SDK
 *
 * MSX SDK client.
 *
 * API version: 1.0.8
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msxsdk

import (
	"encoding/json"
)

// DevicesPageAllOf struct for DevicesPageAllOf
type DevicesPageAllOf struct {
	Contents *[]Device `json:"contents,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DevicesPageAllOf DevicesPageAllOf

// NewDevicesPageAllOf instantiates a new DevicesPageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesPageAllOf() *DevicesPageAllOf {
	this := DevicesPageAllOf{}
	return &this
}

// NewDevicesPageAllOfWithDefaults instantiates a new DevicesPageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesPageAllOfWithDefaults() *DevicesPageAllOf {
	this := DevicesPageAllOf{}
	return &this
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *DevicesPageAllOf) GetContents() []Device {
	if o == nil || o.Contents == nil {
		var ret []Device
		return ret
	}
	return *o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesPageAllOf) GetContentsOk() (*[]Device, bool) {
	if o == nil || o.Contents == nil {
		return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *DevicesPageAllOf) HasContents() bool {
	if o != nil && o.Contents != nil {
		return true
	}

	return false
}

// SetContents gets a reference to the given []Device and assigns it to the Contents field.
func (o *DevicesPageAllOf) SetContents(v []Device) {
	o.Contents = &v
}

func (o DevicesPageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Contents != nil {
		toSerialize["contents"] = o.Contents
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DevicesPageAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varDevicesPageAllOf := _DevicesPageAllOf{}

	if err = json.Unmarshal(bytes, &varDevicesPageAllOf); err == nil {
		*o = DevicesPageAllOf(varDevicesPageAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "contents")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDevicesPageAllOf struct {
	value *DevicesPageAllOf
	isSet bool
}

func (v NullableDevicesPageAllOf) Get() *DevicesPageAllOf {
	return v.value
}

func (v *NullableDevicesPageAllOf) Set(val *DevicesPageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesPageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesPageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesPageAllOf(val *DevicesPageAllOf) *NullableDevicesPageAllOf {
	return &NullableDevicesPageAllOf{value: val, isSet: true}
}

func (v NullableDevicesPageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesPageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


