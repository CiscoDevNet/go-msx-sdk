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

// UsersPageAllOf struct for UsersPageAllOf
type UsersPageAllOf struct {
	Contents []User `json:"contents,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UsersPageAllOf UsersPageAllOf

// NewUsersPageAllOf instantiates a new UsersPageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersPageAllOf() *UsersPageAllOf {
	this := UsersPageAllOf{}
	return &this
}

// NewUsersPageAllOfWithDefaults instantiates a new UsersPageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersPageAllOfWithDefaults() *UsersPageAllOf {
	this := UsersPageAllOf{}
	return &this
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *UsersPageAllOf) GetContents() []User {
	if o == nil || o.Contents == nil {
		var ret []User
		return ret
	}
	return o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersPageAllOf) GetContentsOk() ([]User, bool) {
	if o == nil || o.Contents == nil {
		return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *UsersPageAllOf) HasContents() bool {
	if o != nil && o.Contents != nil {
		return true
	}

	return false
}

// SetContents gets a reference to the given []User and assigns it to the Contents field.
func (o *UsersPageAllOf) SetContents(v []User) {
	o.Contents = v
}

func (o UsersPageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Contents != nil {
		toSerialize["contents"] = o.Contents
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UsersPageAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varUsersPageAllOf := _UsersPageAllOf{}

	if err = json.Unmarshal(bytes, &varUsersPageAllOf); err == nil {
		*o = UsersPageAllOf(varUsersPageAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "contents")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUsersPageAllOf struct {
	value *UsersPageAllOf
	isSet bool
}

func (v NullableUsersPageAllOf) Get() *UsersPageAllOf {
	return v.value
}

func (v *NullableUsersPageAllOf) Set(val *UsersPageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersPageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersPageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersPageAllOf(val *UsersPageAllOf) *NullableUsersPageAllOf {
	return &NullableUsersPageAllOf{value: val, isSet: true}
}

func (v NullableUsersPageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersPageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


