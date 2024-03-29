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

// ChangeRequestUpdate struct for ChangeRequestUpdate
type ChangeRequestUpdate struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Attributes map[string]map[string]interface{} `json:"attributes,omitempty"`
	TenantId NullableString `json:"tenantId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChangeRequestUpdate ChangeRequestUpdate

// NewChangeRequestUpdate instantiates a new ChangeRequestUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeRequestUpdate() *ChangeRequestUpdate {
	this := ChangeRequestUpdate{}
	return &this
}

// NewChangeRequestUpdateWithDefaults instantiates a new ChangeRequestUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeRequestUpdateWithDefaults() *ChangeRequestUpdate {
	this := ChangeRequestUpdate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ChangeRequestUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ChangeRequestUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ChangeRequestUpdate) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ChangeRequestUpdate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ChangeRequestUpdate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ChangeRequestUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ChangeRequestUpdate) GetAttributes() map[string]map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestUpdate) GetAttributesOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ChangeRequestUpdate) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]map[string]interface{} and assigns it to the Attributes field.
func (o *ChangeRequestUpdate) SetAttributes(v map[string]map[string]interface{}) {
	o.Attributes = v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChangeRequestUpdate) GetTenantId() string {
	if o == nil || o.TenantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChangeRequestUpdate) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *ChangeRequestUpdate) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *ChangeRequestUpdate) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *ChangeRequestUpdate) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *ChangeRequestUpdate) UnsetTenantId() {
	o.TenantId.Unset()
}

func (o ChangeRequestUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ChangeRequestUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varChangeRequestUpdate := _ChangeRequestUpdate{}

	if err = json.Unmarshal(bytes, &varChangeRequestUpdate); err == nil {
		*o = ChangeRequestUpdate(varChangeRequestUpdate)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "tenantId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChangeRequestUpdate struct {
	value *ChangeRequestUpdate
	isSet bool
}

func (v NullableChangeRequestUpdate) Get() *ChangeRequestUpdate {
	return v.value
}

func (v *NullableChangeRequestUpdate) Set(val *ChangeRequestUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeRequestUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeRequestUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeRequestUpdate(val *ChangeRequestUpdate) *NullableChangeRequestUpdate {
	return &NullableChangeRequestUpdate{value: val, isSet: true}
}

func (v NullableChangeRequestUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeRequestUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


