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

// LegacyNsoResponseTypes struct for LegacyNsoResponseTypes
type LegacyNsoResponseTypes struct {
	CreateOperation *string `json:"createOperation,omitempty"`
	UpdateOperation *string `json:"updateOperation,omitempty"`
	DeleteOperation *string `json:"deleteOperation,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LegacyNsoResponseTypes LegacyNsoResponseTypes

// NewLegacyNsoResponseTypes instantiates a new LegacyNsoResponseTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLegacyNsoResponseTypes() *LegacyNsoResponseTypes {
	this := LegacyNsoResponseTypes{}
	return &this
}

// NewLegacyNsoResponseTypesWithDefaults instantiates a new LegacyNsoResponseTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLegacyNsoResponseTypesWithDefaults() *LegacyNsoResponseTypes {
	this := LegacyNsoResponseTypes{}
	return &this
}

// GetCreateOperation returns the CreateOperation field value if set, zero value otherwise.
func (o *LegacyNsoResponseTypes) GetCreateOperation() string {
	if o == nil || o.CreateOperation == nil {
		var ret string
		return ret
	}
	return *o.CreateOperation
}

// GetCreateOperationOk returns a tuple with the CreateOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyNsoResponseTypes) GetCreateOperationOk() (*string, bool) {
	if o == nil || o.CreateOperation == nil {
		return nil, false
	}
	return o.CreateOperation, true
}

// HasCreateOperation returns a boolean if a field has been set.
func (o *LegacyNsoResponseTypes) HasCreateOperation() bool {
	if o != nil && o.CreateOperation != nil {
		return true
	}

	return false
}

// SetCreateOperation gets a reference to the given string and assigns it to the CreateOperation field.
func (o *LegacyNsoResponseTypes) SetCreateOperation(v string) {
	o.CreateOperation = &v
}

// GetUpdateOperation returns the UpdateOperation field value if set, zero value otherwise.
func (o *LegacyNsoResponseTypes) GetUpdateOperation() string {
	if o == nil || o.UpdateOperation == nil {
		var ret string
		return ret
	}
	return *o.UpdateOperation
}

// GetUpdateOperationOk returns a tuple with the UpdateOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyNsoResponseTypes) GetUpdateOperationOk() (*string, bool) {
	if o == nil || o.UpdateOperation == nil {
		return nil, false
	}
	return o.UpdateOperation, true
}

// HasUpdateOperation returns a boolean if a field has been set.
func (o *LegacyNsoResponseTypes) HasUpdateOperation() bool {
	if o != nil && o.UpdateOperation != nil {
		return true
	}

	return false
}

// SetUpdateOperation gets a reference to the given string and assigns it to the UpdateOperation field.
func (o *LegacyNsoResponseTypes) SetUpdateOperation(v string) {
	o.UpdateOperation = &v
}

// GetDeleteOperation returns the DeleteOperation field value if set, zero value otherwise.
func (o *LegacyNsoResponseTypes) GetDeleteOperation() string {
	if o == nil || o.DeleteOperation == nil {
		var ret string
		return ret
	}
	return *o.DeleteOperation
}

// GetDeleteOperationOk returns a tuple with the DeleteOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyNsoResponseTypes) GetDeleteOperationOk() (*string, bool) {
	if o == nil || o.DeleteOperation == nil {
		return nil, false
	}
	return o.DeleteOperation, true
}

// HasDeleteOperation returns a boolean if a field has been set.
func (o *LegacyNsoResponseTypes) HasDeleteOperation() bool {
	if o != nil && o.DeleteOperation != nil {
		return true
	}

	return false
}

// SetDeleteOperation gets a reference to the given string and assigns it to the DeleteOperation field.
func (o *LegacyNsoResponseTypes) SetDeleteOperation(v string) {
	o.DeleteOperation = &v
}

func (o LegacyNsoResponseTypes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreateOperation != nil {
		toSerialize["createOperation"] = o.CreateOperation
	}
	if o.UpdateOperation != nil {
		toSerialize["updateOperation"] = o.UpdateOperation
	}
	if o.DeleteOperation != nil {
		toSerialize["deleteOperation"] = o.DeleteOperation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LegacyNsoResponseTypes) UnmarshalJSON(bytes []byte) (err error) {
	varLegacyNsoResponseTypes := _LegacyNsoResponseTypes{}

	if err = json.Unmarshal(bytes, &varLegacyNsoResponseTypes); err == nil {
		*o = LegacyNsoResponseTypes(varLegacyNsoResponseTypes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "createOperation")
		delete(additionalProperties, "updateOperation")
		delete(additionalProperties, "deleteOperation")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLegacyNsoResponseTypes struct {
	value *LegacyNsoResponseTypes
	isSet bool
}

func (v NullableLegacyNsoResponseTypes) Get() *LegacyNsoResponseTypes {
	return v.value
}

func (v *NullableLegacyNsoResponseTypes) Set(val *LegacyNsoResponseTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableLegacyNsoResponseTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableLegacyNsoResponseTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLegacyNsoResponseTypes(val *LegacyNsoResponseTypes) *NullableLegacyNsoResponseTypes {
	return &NullableLegacyNsoResponseTypes{value: val, isSet: true}
}

func (v NullableLegacyNsoResponseTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLegacyNsoResponseTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


