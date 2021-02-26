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

// WorkflowCategory struct for WorkflowCategory
type WorkflowCategory struct {
	Id *string `json:"id,omitempty"`
	SchemaId *string `json:"schema_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Title *string `json:"title,omitempty"`
	Type *string `json:"type,omitempty"`
	BaseType *string `json:"base_type,omitempty"`
	Description *string `json:"description,omitempty"`
	CategoryType *string `json:"category_type,omitempty"`
	ObjectType *string `json:"object_type,omitempty"`
	CreatedOn *string `json:"created_on,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedOn *string `json:"updated_on,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	Owner *string `json:"owner,omitempty"`
	UniqueName *string `json:"unique_name,omitempty"`
}

// NewWorkflowCategory instantiates a new WorkflowCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowCategory() *WorkflowCategory {
	this := WorkflowCategory{}
	return &this
}

// NewWorkflowCategoryWithDefaults instantiates a new WorkflowCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowCategoryWithDefaults() *WorkflowCategory {
	this := WorkflowCategory{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkflowCategory) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCategory) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkflowCategory) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkflowCategory) SetId(v string) {
	o.Id = &v
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *WorkflowCategory) GetSchemaId() string {
	if o == nil || o.SchemaId == nil {
		var ret string
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCategory) GetSchemaIdOk() (*string, bool) {
	if o == nil || o.SchemaId == nil {
		return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *WorkflowCategory) HasSchemaId() bool {
	if o != nil && o.SchemaId != nil {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given string and assigns it to the SchemaId field.
func (o *WorkflowCategory) SetSchemaId(v string) {
	o.SchemaId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowCategory) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCategory) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowCategory) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowCategory) SetName(v string) {
	o.Name = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *WorkflowCategory) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCategory) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *WorkflowCategory) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *WorkflowCategory) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkflowCategory) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCategory) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkflowCategory) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkflowCategory) SetType(v string) {
	o.Type = &v
}

// GetBaseType returns the BaseType field value if set, zero value otherwise.
func (o *WorkflowCategory) GetBaseType() string {
	if o == nil || o.BaseType == nil {
		var ret string
		return ret
	}
	return *o.BaseType
}

// GetBaseTypeOk returns a tuple with the BaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCategory) GetBaseTypeOk() (*string, bool) {
	if o == nil || o.BaseType == nil {
		return nil, false
	}
	return o.BaseType, true
}

// HasBaseType returns a boolean if a field has been set.
func (o *WorkflowCategory) HasBaseType() bool {
	if o != nil && o.BaseType != nil {
		return true
	}

	return false
}

// SetBaseType gets a reference to the given string and assigns it to the BaseType field.
func (o *WorkflowCategory) SetBaseType(v string) {
	o.BaseType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowCategory) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCategory) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowCategory) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowCategory) SetDescription(v string) {
	o.Description = &v
}

// GetCategoryType returns the CategoryType field value if set, zero value otherwise.
func (o *WorkflowCategory) GetCategoryType() string {
	if o == nil || o.CategoryType == nil {
		var ret string
		return ret
	}
	return *o.CategoryType
}

// GetCategoryTypeOk returns a tuple with the CategoryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCategory) GetCategoryTypeOk() (*string, bool) {
	if o == nil || o.CategoryType == nil {
		return nil, false
	}
	return o.CategoryType, true
}

// HasCategoryType returns a boolean if a field has been set.
func (o *WorkflowCategory) HasCategoryType() bool {
	if o != nil && o.CategoryType != nil {
		return true
	}

	return false
}

// SetCategoryType gets a reference to the given string and assigns it to the CategoryType field.
func (o *WorkflowCategory) SetCategoryType(v string) {
	o.CategoryType = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *WorkflowCategory) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCategory) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *WorkflowCategory) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *WorkflowCategory) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *WorkflowCategory) GetCreatedOn() string {
	if o == nil || o.CreatedOn == nil {
		var ret string
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCategory) GetCreatedOnOk() (*string, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *WorkflowCategory) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given string and assigns it to the CreatedOn field.
func (o *WorkflowCategory) SetCreatedOn(v string) {
	o.CreatedOn = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *WorkflowCategory) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCategory) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *WorkflowCategory) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *WorkflowCategory) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetUpdatedOn returns the UpdatedOn field value if set, zero value otherwise.
func (o *WorkflowCategory) GetUpdatedOn() string {
	if o == nil || o.UpdatedOn == nil {
		var ret string
		return ret
	}
	return *o.UpdatedOn
}

// GetUpdatedOnOk returns a tuple with the UpdatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCategory) GetUpdatedOnOk() (*string, bool) {
	if o == nil || o.UpdatedOn == nil {
		return nil, false
	}
	return o.UpdatedOn, true
}

// HasUpdatedOn returns a boolean if a field has been set.
func (o *WorkflowCategory) HasUpdatedOn() bool {
	if o != nil && o.UpdatedOn != nil {
		return true
	}

	return false
}

// SetUpdatedOn gets a reference to the given string and assigns it to the UpdatedOn field.
func (o *WorkflowCategory) SetUpdatedOn(v string) {
	o.UpdatedOn = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *WorkflowCategory) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCategory) GetUpdatedByOk() (*string, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *WorkflowCategory) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *WorkflowCategory) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *WorkflowCategory) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCategory) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *WorkflowCategory) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *WorkflowCategory) SetOwner(v string) {
	o.Owner = &v
}

// GetUniqueName returns the UniqueName field value if set, zero value otherwise.
func (o *WorkflowCategory) GetUniqueName() string {
	if o == nil || o.UniqueName == nil {
		var ret string
		return ret
	}
	return *o.UniqueName
}

// GetUniqueNameOk returns a tuple with the UniqueName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCategory) GetUniqueNameOk() (*string, bool) {
	if o == nil || o.UniqueName == nil {
		return nil, false
	}
	return o.UniqueName, true
}

// HasUniqueName returns a boolean if a field has been set.
func (o *WorkflowCategory) HasUniqueName() bool {
	if o != nil && o.UniqueName != nil {
		return true
	}

	return false
}

// SetUniqueName gets a reference to the given string and assigns it to the UniqueName field.
func (o *WorkflowCategory) SetUniqueName(v string) {
	o.UniqueName = &v
}

func (o WorkflowCategory) MarshalJSON() ([]byte, error) {
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
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.CategoryType != nil {
		toSerialize["category_type"] = o.CategoryType
	}
	if o.ObjectType != nil {
		toSerialize["object_type"] = o.ObjectType
	}
	if o.CreatedOn != nil {
		toSerialize["created_on"] = o.CreatedOn
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.UpdatedOn != nil {
		toSerialize["updated_on"] = o.UpdatedOn
	}
	if o.UpdatedBy != nil {
		toSerialize["updated_by"] = o.UpdatedBy
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.UniqueName != nil {
		toSerialize["unique_name"] = o.UniqueName
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowCategory struct {
	value *WorkflowCategory
	isSet bool
}

func (v NullableWorkflowCategory) Get() *WorkflowCategory {
	return v.value
}

func (v *NullableWorkflowCategory) Set(val *WorkflowCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowCategory(val *WorkflowCategory) *NullableWorkflowCategory {
	return &NullableWorkflowCategory{value: val, isSet: true}
}

func (v NullableWorkflowCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


