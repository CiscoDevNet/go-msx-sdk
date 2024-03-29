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

// OffersPageAllOf struct for OffersPageAllOf
type OffersPageAllOf struct {
	Contents []Offer `json:"contents,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OffersPageAllOf OffersPageAllOf

// NewOffersPageAllOf instantiates a new OffersPageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOffersPageAllOf() *OffersPageAllOf {
	this := OffersPageAllOf{}
	return &this
}

// NewOffersPageAllOfWithDefaults instantiates a new OffersPageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOffersPageAllOfWithDefaults() *OffersPageAllOf {
	this := OffersPageAllOf{}
	return &this
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *OffersPageAllOf) GetContents() []Offer {
	if o == nil || o.Contents == nil {
		var ret []Offer
		return ret
	}
	return o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OffersPageAllOf) GetContentsOk() ([]Offer, bool) {
	if o == nil || o.Contents == nil {
		return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *OffersPageAllOf) HasContents() bool {
	if o != nil && o.Contents != nil {
		return true
	}

	return false
}

// SetContents gets a reference to the given []Offer and assigns it to the Contents field.
func (o *OffersPageAllOf) SetContents(v []Offer) {
	o.Contents = v
}

func (o OffersPageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Contents != nil {
		toSerialize["contents"] = o.Contents
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OffersPageAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varOffersPageAllOf := _OffersPageAllOf{}

	if err = json.Unmarshal(bytes, &varOffersPageAllOf); err == nil {
		*o = OffersPageAllOf(varOffersPageAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "contents")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOffersPageAllOf struct {
	value *OffersPageAllOf
	isSet bool
}

func (v NullableOffersPageAllOf) Get() *OffersPageAllOf {
	return v.value
}

func (v *NullableOffersPageAllOf) Set(val *OffersPageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOffersPageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOffersPageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOffersPageAllOf(val *OffersPageAllOf) *NullableOffersPageAllOf {
	return &NullableOffersPageAllOf{value: val, isSet: true}
}

func (v NullableOffersPageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOffersPageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


