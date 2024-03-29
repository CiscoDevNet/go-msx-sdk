/*
MSX SDK

MSX SDK client.

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msxsdk

import (
	"encoding/json"
	"time"
)

// GenericEventAllOf struct for GenericEventAllOf
type GenericEventAllOf struct {
	Type *string `json:"type,omitempty"`
	Time *time.Time `json:"time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GenericEventAllOf GenericEventAllOf

// NewGenericEventAllOf instantiates a new GenericEventAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericEventAllOf() *GenericEventAllOf {
	this := GenericEventAllOf{}
	return &this
}

// NewGenericEventAllOfWithDefaults instantiates a new GenericEventAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericEventAllOfWithDefaults() *GenericEventAllOf {
	this := GenericEventAllOf{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GenericEventAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericEventAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GenericEventAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GenericEventAllOf) SetType(v string) {
	o.Type = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GenericEventAllOf) GetTime() time.Time {
	if o == nil || o.Time == nil {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericEventAllOf) GetTimeOk() (*time.Time, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GenericEventAllOf) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *GenericEventAllOf) SetTime(v time.Time) {
	o.Time = &v
}

func (o GenericEventAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GenericEventAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varGenericEventAllOf := _GenericEventAllOf{}

	if err = json.Unmarshal(bytes, &varGenericEventAllOf); err == nil {
		*o = GenericEventAllOf(varGenericEventAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGenericEventAllOf struct {
	value *GenericEventAllOf
	isSet bool
}

func (v NullableGenericEventAllOf) Get() *GenericEventAllOf {
	return v.value
}

func (v *NullableGenericEventAllOf) Set(val *GenericEventAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericEventAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericEventAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericEventAllOf(val *GenericEventAllOf) *NullableGenericEventAllOf {
	return &NullableGenericEventAllOf{value: val, isSet: true}
}

func (v NullableGenericEventAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericEventAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


