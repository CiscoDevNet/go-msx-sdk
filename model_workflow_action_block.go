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

// WorkflowActionBlock struct for WorkflowActionBlock
type WorkflowActionBlock struct {
	UniqueName *string `json:"unique_name,omitempty"`
	Name *string `json:"name,omitempty"`
	Title *string `json:"title,omitempty"`
	Type *string `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
	BaseType *string `json:"base_type,omitempty"`
	Properties *map[string]map[string]interface{} `json:"properties,omitempty"`
	ObjectType *string `json:"object_type,omitempty"`
	Actions *[]WorkflowAction `json:"actions,omitempty"`
	SubworkflowId *string `json:"subworkflow_id,omitempty"`
	WorkflowId *string `json:"workflow_id,omitempty"`
	SchemaId *string `json:"schema_id,omitempty"`
}

// NewWorkflowActionBlock instantiates a new WorkflowActionBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowActionBlock() *WorkflowActionBlock {
	this := WorkflowActionBlock{}
	return &this
}

// NewWorkflowActionBlockWithDefaults instantiates a new WorkflowActionBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowActionBlockWithDefaults() *WorkflowActionBlock {
	this := WorkflowActionBlock{}
	return &this
}

// GetUniqueName returns the UniqueName field value if set, zero value otherwise.
func (o *WorkflowActionBlock) GetUniqueName() string {
	if o == nil || o.UniqueName == nil {
		var ret string
		return ret
	}
	return *o.UniqueName
}

// GetUniqueNameOk returns a tuple with the UniqueName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowActionBlock) GetUniqueNameOk() (*string, bool) {
	if o == nil || o.UniqueName == nil {
		return nil, false
	}
	return o.UniqueName, true
}

// HasUniqueName returns a boolean if a field has been set.
func (o *WorkflowActionBlock) HasUniqueName() bool {
	if o != nil && o.UniqueName != nil {
		return true
	}

	return false
}

// SetUniqueName gets a reference to the given string and assigns it to the UniqueName field.
func (o *WorkflowActionBlock) SetUniqueName(v string) {
	o.UniqueName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowActionBlock) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowActionBlock) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowActionBlock) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowActionBlock) SetName(v string) {
	o.Name = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *WorkflowActionBlock) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowActionBlock) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *WorkflowActionBlock) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *WorkflowActionBlock) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkflowActionBlock) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowActionBlock) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkflowActionBlock) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkflowActionBlock) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *WorkflowActionBlock) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowActionBlock) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *WorkflowActionBlock) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *WorkflowActionBlock) SetVersion(v string) {
	o.Version = &v
}

// GetBaseType returns the BaseType field value if set, zero value otherwise.
func (o *WorkflowActionBlock) GetBaseType() string {
	if o == nil || o.BaseType == nil {
		var ret string
		return ret
	}
	return *o.BaseType
}

// GetBaseTypeOk returns a tuple with the BaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowActionBlock) GetBaseTypeOk() (*string, bool) {
	if o == nil || o.BaseType == nil {
		return nil, false
	}
	return o.BaseType, true
}

// HasBaseType returns a boolean if a field has been set.
func (o *WorkflowActionBlock) HasBaseType() bool {
	if o != nil && o.BaseType != nil {
		return true
	}

	return false
}

// SetBaseType gets a reference to the given string and assigns it to the BaseType field.
func (o *WorkflowActionBlock) SetBaseType(v string) {
	o.BaseType = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *WorkflowActionBlock) GetProperties() map[string]map[string]interface{} {
	if o == nil || o.Properties == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowActionBlock) GetPropertiesOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *WorkflowActionBlock) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]map[string]interface{} and assigns it to the Properties field.
func (o *WorkflowActionBlock) SetProperties(v map[string]map[string]interface{}) {
	o.Properties = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *WorkflowActionBlock) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowActionBlock) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *WorkflowActionBlock) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *WorkflowActionBlock) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *WorkflowActionBlock) GetActions() []WorkflowAction {
	if o == nil || o.Actions == nil {
		var ret []WorkflowAction
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowActionBlock) GetActionsOk() (*[]WorkflowAction, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *WorkflowActionBlock) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given []WorkflowAction and assigns it to the Actions field.
func (o *WorkflowActionBlock) SetActions(v []WorkflowAction) {
	o.Actions = &v
}

// GetSubworkflowId returns the SubworkflowId field value if set, zero value otherwise.
func (o *WorkflowActionBlock) GetSubworkflowId() string {
	if o == nil || o.SubworkflowId == nil {
		var ret string
		return ret
	}
	return *o.SubworkflowId
}

// GetSubworkflowIdOk returns a tuple with the SubworkflowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowActionBlock) GetSubworkflowIdOk() (*string, bool) {
	if o == nil || o.SubworkflowId == nil {
		return nil, false
	}
	return o.SubworkflowId, true
}

// HasSubworkflowId returns a boolean if a field has been set.
func (o *WorkflowActionBlock) HasSubworkflowId() bool {
	if o != nil && o.SubworkflowId != nil {
		return true
	}

	return false
}

// SetSubworkflowId gets a reference to the given string and assigns it to the SubworkflowId field.
func (o *WorkflowActionBlock) SetSubworkflowId(v string) {
	o.SubworkflowId = &v
}

// GetWorkflowId returns the WorkflowId field value if set, zero value otherwise.
func (o *WorkflowActionBlock) GetWorkflowId() string {
	if o == nil || o.WorkflowId == nil {
		var ret string
		return ret
	}
	return *o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowActionBlock) GetWorkflowIdOk() (*string, bool) {
	if o == nil || o.WorkflowId == nil {
		return nil, false
	}
	return o.WorkflowId, true
}

// HasWorkflowId returns a boolean if a field has been set.
func (o *WorkflowActionBlock) HasWorkflowId() bool {
	if o != nil && o.WorkflowId != nil {
		return true
	}

	return false
}

// SetWorkflowId gets a reference to the given string and assigns it to the WorkflowId field.
func (o *WorkflowActionBlock) SetWorkflowId(v string) {
	o.WorkflowId = &v
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *WorkflowActionBlock) GetSchemaId() string {
	if o == nil || o.SchemaId == nil {
		var ret string
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowActionBlock) GetSchemaIdOk() (*string, bool) {
	if o == nil || o.SchemaId == nil {
		return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *WorkflowActionBlock) HasSchemaId() bool {
	if o != nil && o.SchemaId != nil {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given string and assigns it to the SchemaId field.
func (o *WorkflowActionBlock) SetSchemaId(v string) {
	o.SchemaId = &v
}

func (o WorkflowActionBlock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UniqueName != nil {
		toSerialize["unique_name"] = o.UniqueName
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
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.BaseType != nil {
		toSerialize["base_type"] = o.BaseType
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.ObjectType != nil {
		toSerialize["object_type"] = o.ObjectType
	}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.SubworkflowId != nil {
		toSerialize["subworkflow_id"] = o.SubworkflowId
	}
	if o.WorkflowId != nil {
		toSerialize["workflow_id"] = o.WorkflowId
	}
	if o.SchemaId != nil {
		toSerialize["schema_id"] = o.SchemaId
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowActionBlock struct {
	value *WorkflowActionBlock
	isSet bool
}

func (v NullableWorkflowActionBlock) Get() *WorkflowActionBlock {
	return v.value
}

func (v *NullableWorkflowActionBlock) Set(val *WorkflowActionBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowActionBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowActionBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowActionBlock(val *WorkflowActionBlock) *NullableWorkflowActionBlock {
	return &NullableWorkflowActionBlock{value: val, isSet: true}
}

func (v NullableWorkflowActionBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowActionBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


