/*
 * MSX SDK
 *
 * MSX SDK client.
 *
 * API version: 1.0.5
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msxsdk

import (
	"encoding/json"
)

// WorkflowVariableAllOf struct for WorkflowVariableAllOf
type WorkflowVariableAllOf struct {
	Id *string `json:"id,omitempty"`
	Type NullableString `json:"type,omitempty"`
	BaseType NullableString `json:"base_type,omitempty"`
	SchemaId *string `json:"schema_id,omitempty"`
	ObjectType *string `json:"object_type,omitempty"`
	Properties *map[string]interface{} `json:"properties,omitempty"`
	DataType NullableString `json:"data_type,omitempty"`
	Scope NullableString `json:"scope,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowVariableAllOf WorkflowVariableAllOf

// NewWorkflowVariableAllOf instantiates a new WorkflowVariableAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowVariableAllOf() *WorkflowVariableAllOf {
	this := WorkflowVariableAllOf{}
	return &this
}

// NewWorkflowVariableAllOfWithDefaults instantiates a new WorkflowVariableAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowVariableAllOfWithDefaults() *WorkflowVariableAllOf {
	this := WorkflowVariableAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkflowVariableAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowVariableAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkflowVariableAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkflowVariableAllOf) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowVariableAllOf) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowVariableAllOf) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *WorkflowVariableAllOf) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *WorkflowVariableAllOf) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *WorkflowVariableAllOf) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *WorkflowVariableAllOf) UnsetType() {
	o.Type.Unset()
}

// GetBaseType returns the BaseType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowVariableAllOf) GetBaseType() string {
	if o == nil || o.BaseType.Get() == nil {
		var ret string
		return ret
	}
	return *o.BaseType.Get()
}

// GetBaseTypeOk returns a tuple with the BaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowVariableAllOf) GetBaseTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BaseType.Get(), o.BaseType.IsSet()
}

// HasBaseType returns a boolean if a field has been set.
func (o *WorkflowVariableAllOf) HasBaseType() bool {
	if o != nil && o.BaseType.IsSet() {
		return true
	}

	return false
}

// SetBaseType gets a reference to the given NullableString and assigns it to the BaseType field.
func (o *WorkflowVariableAllOf) SetBaseType(v string) {
	o.BaseType.Set(&v)
}
// SetBaseTypeNil sets the value for BaseType to be an explicit nil
func (o *WorkflowVariableAllOf) SetBaseTypeNil() {
	o.BaseType.Set(nil)
}

// UnsetBaseType ensures that no value is present for BaseType, not even an explicit nil
func (o *WorkflowVariableAllOf) UnsetBaseType() {
	o.BaseType.Unset()
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *WorkflowVariableAllOf) GetSchemaId() string {
	if o == nil || o.SchemaId == nil {
		var ret string
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowVariableAllOf) GetSchemaIdOk() (*string, bool) {
	if o == nil || o.SchemaId == nil {
		return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *WorkflowVariableAllOf) HasSchemaId() bool {
	if o != nil && o.SchemaId != nil {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given string and assigns it to the SchemaId field.
func (o *WorkflowVariableAllOf) SetSchemaId(v string) {
	o.SchemaId = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *WorkflowVariableAllOf) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowVariableAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *WorkflowVariableAllOf) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *WorkflowVariableAllOf) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *WorkflowVariableAllOf) GetProperties() map[string]interface{} {
	if o == nil || o.Properties == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowVariableAllOf) GetPropertiesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *WorkflowVariableAllOf) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *WorkflowVariableAllOf) SetProperties(v map[string]interface{}) {
	o.Properties = &v
}

// GetDataType returns the DataType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowVariableAllOf) GetDataType() string {
	if o == nil || o.DataType.Get() == nil {
		var ret string
		return ret
	}
	return *o.DataType.Get()
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowVariableAllOf) GetDataTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DataType.Get(), o.DataType.IsSet()
}

// HasDataType returns a boolean if a field has been set.
func (o *WorkflowVariableAllOf) HasDataType() bool {
	if o != nil && o.DataType.IsSet() {
		return true
	}

	return false
}

// SetDataType gets a reference to the given NullableString and assigns it to the DataType field.
func (o *WorkflowVariableAllOf) SetDataType(v string) {
	o.DataType.Set(&v)
}
// SetDataTypeNil sets the value for DataType to be an explicit nil
func (o *WorkflowVariableAllOf) SetDataTypeNil() {
	o.DataType.Set(nil)
}

// UnsetDataType ensures that no value is present for DataType, not even an explicit nil
func (o *WorkflowVariableAllOf) UnsetDataType() {
	o.DataType.Unset()
}

// GetScope returns the Scope field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowVariableAllOf) GetScope() string {
	if o == nil || o.Scope.Get() == nil {
		var ret string
		return ret
	}
	return *o.Scope.Get()
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowVariableAllOf) GetScopeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Scope.Get(), o.Scope.IsSet()
}

// HasScope returns a boolean if a field has been set.
func (o *WorkflowVariableAllOf) HasScope() bool {
	if o != nil && o.Scope.IsSet() {
		return true
	}

	return false
}

// SetScope gets a reference to the given NullableString and assigns it to the Scope field.
func (o *WorkflowVariableAllOf) SetScope(v string) {
	o.Scope.Set(&v)
}
// SetScopeNil sets the value for Scope to be an explicit nil
func (o *WorkflowVariableAllOf) SetScopeNil() {
	o.Scope.Set(nil)
}

// UnsetScope ensures that no value is present for Scope, not even an explicit nil
func (o *WorkflowVariableAllOf) UnsetScope() {
	o.Scope.Unset()
}

func (o WorkflowVariableAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.BaseType.IsSet() {
		toSerialize["base_type"] = o.BaseType.Get()
	}
	if o.SchemaId != nil {
		toSerialize["schema_id"] = o.SchemaId
	}
	if o.ObjectType != nil {
		toSerialize["object_type"] = o.ObjectType
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.DataType.IsSet() {
		toSerialize["data_type"] = o.DataType.Get()
	}
	if o.Scope.IsSet() {
		toSerialize["scope"] = o.Scope.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowVariableAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowVariableAllOf := _WorkflowVariableAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowVariableAllOf); err == nil {
		*o = WorkflowVariableAllOf(varWorkflowVariableAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "base_type")
		delete(additionalProperties, "schema_id")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "properties")
		delete(additionalProperties, "data_type")
		delete(additionalProperties, "scope")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowVariableAllOf struct {
	value *WorkflowVariableAllOf
	isSet bool
}

func (v NullableWorkflowVariableAllOf) Get() *WorkflowVariableAllOf {
	return v.value
}

func (v *NullableWorkflowVariableAllOf) Set(val *WorkflowVariableAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowVariableAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowVariableAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowVariableAllOf(val *WorkflowVariableAllOf) *NullableWorkflowVariableAllOf {
	return &NullableWorkflowVariableAllOf{value: val, isSet: true}
}

func (v NullableWorkflowVariableAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowVariableAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


