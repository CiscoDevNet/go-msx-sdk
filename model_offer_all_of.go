/*
 * KAKAPO - MSX SDK
 *
 * MSX Platform SDK
 *
 * API version: 3.10.0
 * Contact: somecontact@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msxsdk

import (
	"encoding/json"
)

// OfferAllOf struct for OfferAllOf
type OfferAllOf struct {
	Id *string `json:"id,omitempty"`
}

// NewOfferAllOf instantiates a new OfferAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferAllOf() *OfferAllOf {
	this := OfferAllOf{}
	return &this
}

// NewOfferAllOfWithDefaults instantiates a new OfferAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferAllOfWithDefaults() *OfferAllOf {
	this := OfferAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OfferAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OfferAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OfferAllOf) SetId(v string) {
	o.Id = &v
}

func (o OfferAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableOfferAllOf struct {
	value *OfferAllOf
	isSet bool
}

func (v NullableOfferAllOf) Get() *OfferAllOf {
	return v.value
}

func (v *NullableOfferAllOf) Set(val *OfferAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferAllOf(val *OfferAllOf) *NullableOfferAllOf {
	return &NullableOfferAllOf{value: val, isSet: true}
}

func (v NullableOfferAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfferAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


