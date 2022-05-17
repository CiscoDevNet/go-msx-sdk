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

// TenantsPageAllOf struct for TenantsPageAllOf
type TenantsPageAllOf struct {
	Contents []Tenant `json:"contents,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TenantsPageAllOf TenantsPageAllOf

// NewTenantsPageAllOf instantiates a new TenantsPageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantsPageAllOf() *TenantsPageAllOf {
	this := TenantsPageAllOf{}
	return &this
}

// NewTenantsPageAllOfWithDefaults instantiates a new TenantsPageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantsPageAllOfWithDefaults() *TenantsPageAllOf {
	this := TenantsPageAllOf{}
	return &this
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *TenantsPageAllOf) GetContents() []Tenant {
	if o == nil || o.Contents == nil {
		var ret []Tenant
		return ret
	}
	return o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantsPageAllOf) GetContentsOk() ([]Tenant, bool) {
	if o == nil || o.Contents == nil {
		return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *TenantsPageAllOf) HasContents() bool {
	if o != nil && o.Contents != nil {
		return true
	}

	return false
}

// SetContents gets a reference to the given []Tenant and assigns it to the Contents field.
func (o *TenantsPageAllOf) SetContents(v []Tenant) {
	o.Contents = v
}

func (o TenantsPageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Contents != nil {
		toSerialize["contents"] = o.Contents
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TenantsPageAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTenantsPageAllOf := _TenantsPageAllOf{}

	if err = json.Unmarshal(bytes, &varTenantsPageAllOf); err == nil {
		*o = TenantsPageAllOf(varTenantsPageAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "contents")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTenantsPageAllOf struct {
	value *TenantsPageAllOf
	isSet bool
}

func (v NullableTenantsPageAllOf) Get() *TenantsPageAllOf {
	return v.value
}

func (v *NullableTenantsPageAllOf) Set(val *TenantsPageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantsPageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantsPageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantsPageAllOf(val *TenantsPageAllOf) *NullableTenantsPageAllOf {
	return &NullableTenantsPageAllOf{value: val, isSet: true}
}

func (v NullableTenantsPageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantsPageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


