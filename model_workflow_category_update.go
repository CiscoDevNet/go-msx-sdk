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

// WorkflowCategoryUpdate struct for WorkflowCategoryUpdate
type WorkflowCategoryUpdate struct {
	Name string `json:"name"`
	Title string `json:"title"`
	Description string `json:"description"`
	SchemaId string `json:"schema_id"`
	UniqueName *string `json:"unique_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowCategoryUpdate WorkflowCategoryUpdate

// NewWorkflowCategoryUpdate instantiates a new WorkflowCategoryUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowCategoryUpdate(name string, title string, description string, schemaId string) *WorkflowCategoryUpdate {
	this := WorkflowCategoryUpdate{}
	this.Name = name
	this.Title = title
	this.Description = description
	this.SchemaId = schemaId
	return &this
}

// NewWorkflowCategoryUpdateWithDefaults instantiates a new WorkflowCategoryUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowCategoryUpdateWithDefaults() *WorkflowCategoryUpdate {
	this := WorkflowCategoryUpdate{}
	return &this
}

// GetName returns the Name field value
func (o *WorkflowCategoryUpdate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WorkflowCategoryUpdate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WorkflowCategoryUpdate) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value
func (o *WorkflowCategoryUpdate) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *WorkflowCategoryUpdate) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *WorkflowCategoryUpdate) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value
func (o *WorkflowCategoryUpdate) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *WorkflowCategoryUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *WorkflowCategoryUpdate) SetDescription(v string) {
	o.Description = v
}

// GetSchemaId returns the SchemaId field value
func (o *WorkflowCategoryUpdate) GetSchemaId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value
// and a boolean to check if the value has been set.
func (o *WorkflowCategoryUpdate) GetSchemaIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SchemaId, true
}

// SetSchemaId sets field value
func (o *WorkflowCategoryUpdate) SetSchemaId(v string) {
	o.SchemaId = v
}

// GetUniqueName returns the UniqueName field value if set, zero value otherwise.
func (o *WorkflowCategoryUpdate) GetUniqueName() string {
	if o == nil || o.UniqueName == nil {
		var ret string
		return ret
	}
	return *o.UniqueName
}

// GetUniqueNameOk returns a tuple with the UniqueName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCategoryUpdate) GetUniqueNameOk() (*string, bool) {
	if o == nil || o.UniqueName == nil {
		return nil, false
	}
	return o.UniqueName, true
}

// HasUniqueName returns a boolean if a field has been set.
func (o *WorkflowCategoryUpdate) HasUniqueName() bool {
	if o != nil && o.UniqueName != nil {
		return true
	}

	return false
}

// SetUniqueName gets a reference to the given string and assigns it to the UniqueName field.
func (o *WorkflowCategoryUpdate) SetUniqueName(v string) {
	o.UniqueName = &v
}

func (o WorkflowCategoryUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["schema_id"] = o.SchemaId
	}
	if o.UniqueName != nil {
		toSerialize["unique_name"] = o.UniqueName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowCategoryUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowCategoryUpdate := _WorkflowCategoryUpdate{}

	if err = json.Unmarshal(bytes, &varWorkflowCategoryUpdate); err == nil {
		*o = WorkflowCategoryUpdate(varWorkflowCategoryUpdate)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "title")
		delete(additionalProperties, "description")
		delete(additionalProperties, "schema_id")
		delete(additionalProperties, "unique_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowCategoryUpdate struct {
	value *WorkflowCategoryUpdate
	isSet bool
}

func (v NullableWorkflowCategoryUpdate) Get() *WorkflowCategoryUpdate {
	return v.value
}

func (v *NullableWorkflowCategoryUpdate) Set(val *WorkflowCategoryUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowCategoryUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowCategoryUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowCategoryUpdate(val *WorkflowCategoryUpdate) *NullableWorkflowCategoryUpdate {
	return &NullableWorkflowCategoryUpdate{value: val, isSet: true}
}

func (v NullableWorkflowCategoryUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowCategoryUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


