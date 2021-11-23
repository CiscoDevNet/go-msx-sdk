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

// VulnerabilitiesRegistrationPageAllOf struct for VulnerabilitiesRegistrationPageAllOf
type VulnerabilitiesRegistrationPageAllOf struct {
	Contents *[]VulnerabilityRegistration `json:"contents,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VulnerabilitiesRegistrationPageAllOf VulnerabilitiesRegistrationPageAllOf

// NewVulnerabilitiesRegistrationPageAllOf instantiates a new VulnerabilitiesRegistrationPageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVulnerabilitiesRegistrationPageAllOf() *VulnerabilitiesRegistrationPageAllOf {
	this := VulnerabilitiesRegistrationPageAllOf{}
	return &this
}

// NewVulnerabilitiesRegistrationPageAllOfWithDefaults instantiates a new VulnerabilitiesRegistrationPageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVulnerabilitiesRegistrationPageAllOfWithDefaults() *VulnerabilitiesRegistrationPageAllOf {
	this := VulnerabilitiesRegistrationPageAllOf{}
	return &this
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *VulnerabilitiesRegistrationPageAllOf) GetContents() []VulnerabilityRegistration {
	if o == nil || o.Contents == nil {
		var ret []VulnerabilityRegistration
		return ret
	}
	return *o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnerabilitiesRegistrationPageAllOf) GetContentsOk() (*[]VulnerabilityRegistration, bool) {
	if o == nil || o.Contents == nil {
		return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *VulnerabilitiesRegistrationPageAllOf) HasContents() bool {
	if o != nil && o.Contents != nil {
		return true
	}

	return false
}

// SetContents gets a reference to the given []VulnerabilityRegistration and assigns it to the Contents field.
func (o *VulnerabilitiesRegistrationPageAllOf) SetContents(v []VulnerabilityRegistration) {
	o.Contents = &v
}

func (o VulnerabilitiesRegistrationPageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Contents != nil {
		toSerialize["contents"] = o.Contents
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VulnerabilitiesRegistrationPageAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVulnerabilitiesRegistrationPageAllOf := _VulnerabilitiesRegistrationPageAllOf{}

	if err = json.Unmarshal(bytes, &varVulnerabilitiesRegistrationPageAllOf); err == nil {
		*o = VulnerabilitiesRegistrationPageAllOf(varVulnerabilitiesRegistrationPageAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "contents")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVulnerabilitiesRegistrationPageAllOf struct {
	value *VulnerabilitiesRegistrationPageAllOf
	isSet bool
}

func (v NullableVulnerabilitiesRegistrationPageAllOf) Get() *VulnerabilitiesRegistrationPageAllOf {
	return v.value
}

func (v *NullableVulnerabilitiesRegistrationPageAllOf) Set(val *VulnerabilitiesRegistrationPageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVulnerabilitiesRegistrationPageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVulnerabilitiesRegistrationPageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVulnerabilitiesRegistrationPageAllOf(val *VulnerabilitiesRegistrationPageAllOf) *NullableVulnerabilitiesRegistrationPageAllOf {
	return &NullableVulnerabilitiesRegistrationPageAllOf{value: val, isSet: true}
}

func (v NullableVulnerabilitiesRegistrationPageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVulnerabilitiesRegistrationPageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


