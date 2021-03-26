/*
 * KAKAPO - MSX SDK
 *
 * MSX Platform SDK
 *
 * API version: 1.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msxsdk

import (
	"encoding/json"
)

// WorkflowEventUpdate struct for WorkflowEventUpdate
type WorkflowEventUpdate struct {
	Title string `json:"title"`
	Description *string `json:"description,omitempty"`
	TargetId string `json:"target_id"`
	SchemaId string `json:"schema_id"`
	Properties map[string]map[string]interface{} `json:"properties"`
}

// NewWorkflowEventUpdate instantiates a new WorkflowEventUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowEventUpdate(title string, targetId string, schemaId string, properties map[string]map[string]interface{}) *WorkflowEventUpdate {
	this := WorkflowEventUpdate{}
	this.Title = title
	this.TargetId = targetId
	this.SchemaId = schemaId
	this.Properties = properties
	return &this
}

// NewWorkflowEventUpdateWithDefaults instantiates a new WorkflowEventUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowEventUpdateWithDefaults() *WorkflowEventUpdate {
	this := WorkflowEventUpdate{}
	return &this
}

// GetTitle returns the Title field value
func (o *WorkflowEventUpdate) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *WorkflowEventUpdate) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *WorkflowEventUpdate) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowEventUpdate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowEventUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowEventUpdate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowEventUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetTargetId returns the TargetId field value
func (o *WorkflowEventUpdate) GetTargetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value
// and a boolean to check if the value has been set.
func (o *WorkflowEventUpdate) GetTargetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetId, true
}

// SetTargetId sets field value
func (o *WorkflowEventUpdate) SetTargetId(v string) {
	o.TargetId = v
}

// GetSchemaId returns the SchemaId field value
func (o *WorkflowEventUpdate) GetSchemaId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value
// and a boolean to check if the value has been set.
func (o *WorkflowEventUpdate) GetSchemaIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SchemaId, true
}

// SetSchemaId sets field value
func (o *WorkflowEventUpdate) SetSchemaId(v string) {
	o.SchemaId = v
}

// GetProperties returns the Properties field value
func (o *WorkflowEventUpdate) GetProperties() map[string]map[string]interface{} {
	if o == nil {
		var ret map[string]map[string]interface{}
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *WorkflowEventUpdate) GetPropertiesOk() (*map[string]map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *WorkflowEventUpdate) SetProperties(v map[string]map[string]interface{}) {
	o.Properties = v
}

func (o WorkflowEventUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["title"] = o.Title
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["target_id"] = o.TargetId
	}
	if true {
		toSerialize["schema_id"] = o.SchemaId
	}
	if true {
		toSerialize["properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowEventUpdate struct {
	value *WorkflowEventUpdate
	isSet bool
}

func (v NullableWorkflowEventUpdate) Get() *WorkflowEventUpdate {
	return v.value
}

func (v *NullableWorkflowEventUpdate) Set(val *WorkflowEventUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowEventUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowEventUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowEventUpdate(val *WorkflowEventUpdate) *NullableWorkflowEventUpdate {
	return &NullableWorkflowEventUpdate{value: val, isSet: true}
}

func (v NullableWorkflowEventUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowEventUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


