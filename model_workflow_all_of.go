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

// WorkflowAllOf struct for WorkflowAllOf
type WorkflowAllOf struct {
	Id *string `json:"id,omitempty"`
	SchemaId *string `json:"schema_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Title *string `json:"title,omitempty"`
	Type *string `json:"type,omitempty"`
	BaseType *string `json:"base_type,omitempty"`
	ObjectType NullableString `json:"object_type,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	Valid *bool `json:"valid,omitempty"`
	WorkflowValid *bool `json:"workflow_valid,omitempty"`
	Categories []string `json:"categories,omitempty"`
	Metadata *WorkflowMetadata `json:"metadata,omitempty"`
	Status map[string]interface{} `json:"status,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	Variables []WorkflowVariable `json:"variables,omitempty"`
	Actions []WorkflowAction `json:"actions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowAllOf WorkflowAllOf

// NewWorkflowAllOf instantiates a new WorkflowAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowAllOf() *WorkflowAllOf {
	this := WorkflowAllOf{}
	return &this
}

// NewWorkflowAllOfWithDefaults instantiates a new WorkflowAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowAllOfWithDefaults() *WorkflowAllOf {
	this := WorkflowAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkflowAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkflowAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkflowAllOf) SetId(v string) {
	o.Id = &v
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *WorkflowAllOf) GetSchemaId() string {
	if o == nil || o.SchemaId == nil {
		var ret string
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAllOf) GetSchemaIdOk() (*string, bool) {
	if o == nil || o.SchemaId == nil {
		return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *WorkflowAllOf) HasSchemaId() bool {
	if o != nil && o.SchemaId != nil {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given string and assigns it to the SchemaId field.
func (o *WorkflowAllOf) SetSchemaId(v string) {
	o.SchemaId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowAllOf) SetName(v string) {
	o.Name = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *WorkflowAllOf) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAllOf) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *WorkflowAllOf) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *WorkflowAllOf) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkflowAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkflowAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkflowAllOf) SetType(v string) {
	o.Type = &v
}

// GetBaseType returns the BaseType field value if set, zero value otherwise.
func (o *WorkflowAllOf) GetBaseType() string {
	if o == nil || o.BaseType == nil {
		var ret string
		return ret
	}
	return *o.BaseType
}

// GetBaseTypeOk returns a tuple with the BaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAllOf) GetBaseTypeOk() (*string, bool) {
	if o == nil || o.BaseType == nil {
		return nil, false
	}
	return o.BaseType, true
}

// HasBaseType returns a boolean if a field has been set.
func (o *WorkflowAllOf) HasBaseType() bool {
	if o != nil && o.BaseType != nil {
		return true
	}

	return false
}

// SetBaseType gets a reference to the given string and assigns it to the BaseType field.
func (o *WorkflowAllOf) SetBaseType(v string) {
	o.BaseType = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowAllOf) GetObjectType() string {
	if o == nil || o.ObjectType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ObjectType.Get()
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ObjectType.Get(), o.ObjectType.IsSet()
}

// HasObjectType returns a boolean if a field has been set.
func (o *WorkflowAllOf) HasObjectType() bool {
	if o != nil && o.ObjectType.IsSet() {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given NullableString and assigns it to the ObjectType field.
func (o *WorkflowAllOf) SetObjectType(v string) {
	o.ObjectType.Set(&v)
}
// SetObjectTypeNil sets the value for ObjectType to be an explicit nil
func (o *WorkflowAllOf) SetObjectTypeNil() {
	o.ObjectType.Set(nil)
}

// UnsetObjectType ensures that no value is present for ObjectType, not even an explicit nil
func (o *WorkflowAllOf) UnsetObjectType() {
	o.ObjectType.Unset()
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *WorkflowAllOf) GetProperties() map[string]interface{} {
	if o == nil || o.Properties == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAllOf) GetPropertiesOk() (map[string]interface{}, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *WorkflowAllOf) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *WorkflowAllOf) SetProperties(v map[string]interface{}) {
	o.Properties = v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *WorkflowAllOf) GetValid() bool {
	if o == nil || o.Valid == nil {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAllOf) GetValidOk() (*bool, bool) {
	if o == nil || o.Valid == nil {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *WorkflowAllOf) HasValid() bool {
	if o != nil && o.Valid != nil {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *WorkflowAllOf) SetValid(v bool) {
	o.Valid = &v
}

// GetWorkflowValid returns the WorkflowValid field value if set, zero value otherwise.
func (o *WorkflowAllOf) GetWorkflowValid() bool {
	if o == nil || o.WorkflowValid == nil {
		var ret bool
		return ret
	}
	return *o.WorkflowValid
}

// GetWorkflowValidOk returns a tuple with the WorkflowValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAllOf) GetWorkflowValidOk() (*bool, bool) {
	if o == nil || o.WorkflowValid == nil {
		return nil, false
	}
	return o.WorkflowValid, true
}

// HasWorkflowValid returns a boolean if a field has been set.
func (o *WorkflowAllOf) HasWorkflowValid() bool {
	if o != nil && o.WorkflowValid != nil {
		return true
	}

	return false
}

// SetWorkflowValid gets a reference to the given bool and assigns it to the WorkflowValid field.
func (o *WorkflowAllOf) SetWorkflowValid(v bool) {
	o.WorkflowValid = &v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *WorkflowAllOf) GetCategories() []string {
	if o == nil || o.Categories == nil {
		var ret []string
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAllOf) GetCategoriesOk() ([]string, bool) {
	if o == nil || o.Categories == nil {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *WorkflowAllOf) HasCategories() bool {
	if o != nil && o.Categories != nil {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []string and assigns it to the Categories field.
func (o *WorkflowAllOf) SetCategories(v []string) {
	o.Categories = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *WorkflowAllOf) GetMetadata() WorkflowMetadata {
	if o == nil || o.Metadata == nil {
		var ret WorkflowMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAllOf) GetMetadataOk() (*WorkflowMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *WorkflowAllOf) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given WorkflowMetadata and assigns it to the Metadata field.
func (o *WorkflowAllOf) SetMetadata(v WorkflowMetadata) {
	o.Metadata = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WorkflowAllOf) GetStatus() map[string]interface{} {
	if o == nil || o.Status == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAllOf) GetStatusOk() (map[string]interface{}, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WorkflowAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given map[string]interface{} and assigns it to the Status field.
func (o *WorkflowAllOf) SetStatus(v map[string]interface{}) {
	o.Status = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *WorkflowAllOf) GetPermissions() []string {
	if o == nil || o.Permissions == nil {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAllOf) GetPermissionsOk() ([]string, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *WorkflowAllOf) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *WorkflowAllOf) SetPermissions(v []string) {
	o.Permissions = v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *WorkflowAllOf) GetVariables() []WorkflowVariable {
	if o == nil || o.Variables == nil {
		var ret []WorkflowVariable
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAllOf) GetVariablesOk() ([]WorkflowVariable, bool) {
	if o == nil || o.Variables == nil {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *WorkflowAllOf) HasVariables() bool {
	if o != nil && o.Variables != nil {
		return true
	}

	return false
}

// SetVariables gets a reference to the given []WorkflowVariable and assigns it to the Variables field.
func (o *WorkflowAllOf) SetVariables(v []WorkflowVariable) {
	o.Variables = v
}

// GetActions returns the Actions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowAllOf) GetActions() []WorkflowAction {
	if o == nil  {
		var ret []WorkflowAction
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowAllOf) GetActionsOk() ([]WorkflowAction, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *WorkflowAllOf) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given []WorkflowAction and assigns it to the Actions field.
func (o *WorkflowAllOf) SetActions(v []WorkflowAction) {
	o.Actions = v
}

func (o WorkflowAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.SchemaId != nil {
		toSerialize["schema_id"] = o.SchemaId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.BaseType != nil {
		toSerialize["base_type"] = o.BaseType
	}
	if o.ObjectType.IsSet() {
		toSerialize["object_type"] = o.ObjectType.Get()
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Valid != nil {
		toSerialize["valid"] = o.Valid
	}
	if o.WorkflowValid != nil {
		toSerialize["workflow_valid"] = o.WorkflowValid
	}
	if o.Categories != nil {
		toSerialize["categories"] = o.Categories
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	if o.Variables != nil {
		toSerialize["variables"] = o.Variables
	}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowAllOf := _WorkflowAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowAllOf); err == nil {
		*o = WorkflowAllOf(varWorkflowAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "schema_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "title")
		delete(additionalProperties, "type")
		delete(additionalProperties, "base_type")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "properties")
		delete(additionalProperties, "valid")
		delete(additionalProperties, "workflow_valid")
		delete(additionalProperties, "categories")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "status")
		delete(additionalProperties, "permissions")
		delete(additionalProperties, "variables")
		delete(additionalProperties, "actions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowAllOf struct {
	value *WorkflowAllOf
	isSet bool
}

func (v NullableWorkflowAllOf) Get() *WorkflowAllOf {
	return v.value
}

func (v *NullableWorkflowAllOf) Set(val *WorkflowAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowAllOf(val *WorkflowAllOf) *NullableWorkflowAllOf {
	return &NullableWorkflowAllOf{value: val, isSet: true}
}

func (v NullableWorkflowAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


