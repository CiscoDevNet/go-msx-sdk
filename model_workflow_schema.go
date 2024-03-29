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

// WorkflowSchema struct for WorkflowSchema
type WorkflowSchema struct {
	Id NullableString `json:"id,omitempty"`
	SchemaId *string `json:"schema_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Title *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	BaseType *string `json:"base_type,omitempty"`
	Version *string `json:"version,omitempty"`
	Invisible *bool `json:"invisible,omitempty"`
	Inherits *string `json:"inherits,omitempty"`
	AccessMeta *WorkflowAccessMeta `json:"access_meta,omitempty"`
	VariableSchema map[string]interface{} `json:"variable_schema,omitempty"`
	PropertySchema map[string]interface{} `json:"property_schema,omitempty"`
	OutputSchema map[string]interface{} `json:"output_schema,omitempty"`
	ViewConfig map[string]interface{} `json:"view_config,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	CreatedOn *string `json:"created_on,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedOn *string `json:"updated_on,omitempty"`
	UpdatedBy NullableString `json:"updated_by,omitempty"`
	Owner NullableString `json:"owner,omitempty"`
	UniqueName NullableString `json:"unique_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowSchema WorkflowSchema

// NewWorkflowSchema instantiates a new WorkflowSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowSchema() *WorkflowSchema {
	this := WorkflowSchema{}
	return &this
}

// NewWorkflowSchemaWithDefaults instantiates a new WorkflowSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowSchemaWithDefaults() *WorkflowSchema {
	this := WorkflowSchema{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowSchema) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowSchema) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *WorkflowSchema) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *WorkflowSchema) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *WorkflowSchema) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *WorkflowSchema) UnsetId() {
	o.Id.Unset()
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *WorkflowSchema) GetSchemaId() string {
	if o == nil || o.SchemaId == nil {
		var ret string
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchema) GetSchemaIdOk() (*string, bool) {
	if o == nil || o.SchemaId == nil {
		return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *WorkflowSchema) HasSchemaId() bool {
	if o != nil && o.SchemaId != nil {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given string and assigns it to the SchemaId field.
func (o *WorkflowSchema) SetSchemaId(v string) {
	o.SchemaId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowSchema) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchema) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowSchema) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowSchema) SetName(v string) {
	o.Name = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *WorkflowSchema) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchema) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *WorkflowSchema) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *WorkflowSchema) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowSchema) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchema) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowSchema) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowSchema) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkflowSchema) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchema) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkflowSchema) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkflowSchema) SetType(v string) {
	o.Type = &v
}

// GetBaseType returns the BaseType field value if set, zero value otherwise.
func (o *WorkflowSchema) GetBaseType() string {
	if o == nil || o.BaseType == nil {
		var ret string
		return ret
	}
	return *o.BaseType
}

// GetBaseTypeOk returns a tuple with the BaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchema) GetBaseTypeOk() (*string, bool) {
	if o == nil || o.BaseType == nil {
		return nil, false
	}
	return o.BaseType, true
}

// HasBaseType returns a boolean if a field has been set.
func (o *WorkflowSchema) HasBaseType() bool {
	if o != nil && o.BaseType != nil {
		return true
	}

	return false
}

// SetBaseType gets a reference to the given string and assigns it to the BaseType field.
func (o *WorkflowSchema) SetBaseType(v string) {
	o.BaseType = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *WorkflowSchema) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchema) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *WorkflowSchema) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *WorkflowSchema) SetVersion(v string) {
	o.Version = &v
}

// GetInvisible returns the Invisible field value if set, zero value otherwise.
func (o *WorkflowSchema) GetInvisible() bool {
	if o == nil || o.Invisible == nil {
		var ret bool
		return ret
	}
	return *o.Invisible
}

// GetInvisibleOk returns a tuple with the Invisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchema) GetInvisibleOk() (*bool, bool) {
	if o == nil || o.Invisible == nil {
		return nil, false
	}
	return o.Invisible, true
}

// HasInvisible returns a boolean if a field has been set.
func (o *WorkflowSchema) HasInvisible() bool {
	if o != nil && o.Invisible != nil {
		return true
	}

	return false
}

// SetInvisible gets a reference to the given bool and assigns it to the Invisible field.
func (o *WorkflowSchema) SetInvisible(v bool) {
	o.Invisible = &v
}

// GetInherits returns the Inherits field value if set, zero value otherwise.
func (o *WorkflowSchema) GetInherits() string {
	if o == nil || o.Inherits == nil {
		var ret string
		return ret
	}
	return *o.Inherits
}

// GetInheritsOk returns a tuple with the Inherits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchema) GetInheritsOk() (*string, bool) {
	if o == nil || o.Inherits == nil {
		return nil, false
	}
	return o.Inherits, true
}

// HasInherits returns a boolean if a field has been set.
func (o *WorkflowSchema) HasInherits() bool {
	if o != nil && o.Inherits != nil {
		return true
	}

	return false
}

// SetInherits gets a reference to the given string and assigns it to the Inherits field.
func (o *WorkflowSchema) SetInherits(v string) {
	o.Inherits = &v
}

// GetAccessMeta returns the AccessMeta field value if set, zero value otherwise.
func (o *WorkflowSchema) GetAccessMeta() WorkflowAccessMeta {
	if o == nil || o.AccessMeta == nil {
		var ret WorkflowAccessMeta
		return ret
	}
	return *o.AccessMeta
}

// GetAccessMetaOk returns a tuple with the AccessMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchema) GetAccessMetaOk() (*WorkflowAccessMeta, bool) {
	if o == nil || o.AccessMeta == nil {
		return nil, false
	}
	return o.AccessMeta, true
}

// HasAccessMeta returns a boolean if a field has been set.
func (o *WorkflowSchema) HasAccessMeta() bool {
	if o != nil && o.AccessMeta != nil {
		return true
	}

	return false
}

// SetAccessMeta gets a reference to the given WorkflowAccessMeta and assigns it to the AccessMeta field.
func (o *WorkflowSchema) SetAccessMeta(v WorkflowAccessMeta) {
	o.AccessMeta = &v
}

// GetVariableSchema returns the VariableSchema field value if set, zero value otherwise.
func (o *WorkflowSchema) GetVariableSchema() map[string]interface{} {
	if o == nil || o.VariableSchema == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.VariableSchema
}

// GetVariableSchemaOk returns a tuple with the VariableSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchema) GetVariableSchemaOk() (map[string]interface{}, bool) {
	if o == nil || o.VariableSchema == nil {
		return nil, false
	}
	return o.VariableSchema, true
}

// HasVariableSchema returns a boolean if a field has been set.
func (o *WorkflowSchema) HasVariableSchema() bool {
	if o != nil && o.VariableSchema != nil {
		return true
	}

	return false
}

// SetVariableSchema gets a reference to the given map[string]interface{} and assigns it to the VariableSchema field.
func (o *WorkflowSchema) SetVariableSchema(v map[string]interface{}) {
	o.VariableSchema = v
}

// GetPropertySchema returns the PropertySchema field value if set, zero value otherwise.
func (o *WorkflowSchema) GetPropertySchema() map[string]interface{} {
	if o == nil || o.PropertySchema == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.PropertySchema
}

// GetPropertySchemaOk returns a tuple with the PropertySchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchema) GetPropertySchemaOk() (map[string]interface{}, bool) {
	if o == nil || o.PropertySchema == nil {
		return nil, false
	}
	return o.PropertySchema, true
}

// HasPropertySchema returns a boolean if a field has been set.
func (o *WorkflowSchema) HasPropertySchema() bool {
	if o != nil && o.PropertySchema != nil {
		return true
	}

	return false
}

// SetPropertySchema gets a reference to the given map[string]interface{} and assigns it to the PropertySchema field.
func (o *WorkflowSchema) SetPropertySchema(v map[string]interface{}) {
	o.PropertySchema = v
}

// GetOutputSchema returns the OutputSchema field value if set, zero value otherwise.
func (o *WorkflowSchema) GetOutputSchema() map[string]interface{} {
	if o == nil || o.OutputSchema == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.OutputSchema
}

// GetOutputSchemaOk returns a tuple with the OutputSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchema) GetOutputSchemaOk() (map[string]interface{}, bool) {
	if o == nil || o.OutputSchema == nil {
		return nil, false
	}
	return o.OutputSchema, true
}

// HasOutputSchema returns a boolean if a field has been set.
func (o *WorkflowSchema) HasOutputSchema() bool {
	if o != nil && o.OutputSchema != nil {
		return true
	}

	return false
}

// SetOutputSchema gets a reference to the given map[string]interface{} and assigns it to the OutputSchema field.
func (o *WorkflowSchema) SetOutputSchema(v map[string]interface{}) {
	o.OutputSchema = v
}

// GetViewConfig returns the ViewConfig field value if set, zero value otherwise.
func (o *WorkflowSchema) GetViewConfig() map[string]interface{} {
	if o == nil || o.ViewConfig == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ViewConfig
}

// GetViewConfigOk returns a tuple with the ViewConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchema) GetViewConfigOk() (map[string]interface{}, bool) {
	if o == nil || o.ViewConfig == nil {
		return nil, false
	}
	return o.ViewConfig, true
}

// HasViewConfig returns a boolean if a field has been set.
func (o *WorkflowSchema) HasViewConfig() bool {
	if o != nil && o.ViewConfig != nil {
		return true
	}

	return false
}

// SetViewConfig gets a reference to the given map[string]interface{} and assigns it to the ViewConfig field.
func (o *WorkflowSchema) SetViewConfig(v map[string]interface{}) {
	o.ViewConfig = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *WorkflowSchema) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchema) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *WorkflowSchema) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *WorkflowSchema) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *WorkflowSchema) GetCreatedOn() string {
	if o == nil || o.CreatedOn == nil {
		var ret string
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchema) GetCreatedOnOk() (*string, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *WorkflowSchema) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given string and assigns it to the CreatedOn field.
func (o *WorkflowSchema) SetCreatedOn(v string) {
	o.CreatedOn = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *WorkflowSchema) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchema) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *WorkflowSchema) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *WorkflowSchema) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetUpdatedOn returns the UpdatedOn field value if set, zero value otherwise.
func (o *WorkflowSchema) GetUpdatedOn() string {
	if o == nil || o.UpdatedOn == nil {
		var ret string
		return ret
	}
	return *o.UpdatedOn
}

// GetUpdatedOnOk returns a tuple with the UpdatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchema) GetUpdatedOnOk() (*string, bool) {
	if o == nil || o.UpdatedOn == nil {
		return nil, false
	}
	return o.UpdatedOn, true
}

// HasUpdatedOn returns a boolean if a field has been set.
func (o *WorkflowSchema) HasUpdatedOn() bool {
	if o != nil && o.UpdatedOn != nil {
		return true
	}

	return false
}

// SetUpdatedOn gets a reference to the given string and assigns it to the UpdatedOn field.
func (o *WorkflowSchema) SetUpdatedOn(v string) {
	o.UpdatedOn = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowSchema) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy.Get()
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowSchema) GetUpdatedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdatedBy.Get(), o.UpdatedBy.IsSet()
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *WorkflowSchema) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy.IsSet() {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given NullableString and assigns it to the UpdatedBy field.
func (o *WorkflowSchema) SetUpdatedBy(v string) {
	o.UpdatedBy.Set(&v)
}
// SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil
func (o *WorkflowSchema) SetUpdatedByNil() {
	o.UpdatedBy.Set(nil)
}

// UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
func (o *WorkflowSchema) UnsetUpdatedBy() {
	o.UpdatedBy.Unset()
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowSchema) GetOwner() string {
	if o == nil || o.Owner.Get() == nil {
		var ret string
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowSchema) GetOwnerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *WorkflowSchema) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableString and assigns it to the Owner field.
func (o *WorkflowSchema) SetOwner(v string) {
	o.Owner.Set(&v)
}
// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *WorkflowSchema) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *WorkflowSchema) UnsetOwner() {
	o.Owner.Unset()
}

// GetUniqueName returns the UniqueName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowSchema) GetUniqueName() string {
	if o == nil || o.UniqueName.Get() == nil {
		var ret string
		return ret
	}
	return *o.UniqueName.Get()
}

// GetUniqueNameOk returns a tuple with the UniqueName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowSchema) GetUniqueNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UniqueName.Get(), o.UniqueName.IsSet()
}

// HasUniqueName returns a boolean if a field has been set.
func (o *WorkflowSchema) HasUniqueName() bool {
	if o != nil && o.UniqueName.IsSet() {
		return true
	}

	return false
}

// SetUniqueName gets a reference to the given NullableString and assigns it to the UniqueName field.
func (o *WorkflowSchema) SetUniqueName(v string) {
	o.UniqueName.Set(&v)
}
// SetUniqueNameNil sets the value for UniqueName to be an explicit nil
func (o *WorkflowSchema) SetUniqueNameNil() {
	o.UniqueName.Set(nil)
}

// UnsetUniqueName ensures that no value is present for UniqueName, not even an explicit nil
func (o *WorkflowSchema) UnsetUniqueName() {
	o.UniqueName.Unset()
}

func (o WorkflowSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
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
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.BaseType != nil {
		toSerialize["base_type"] = o.BaseType
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Invisible != nil {
		toSerialize["invisible"] = o.Invisible
	}
	if o.Inherits != nil {
		toSerialize["inherits"] = o.Inherits
	}
	if o.AccessMeta != nil {
		toSerialize["access_meta"] = o.AccessMeta
	}
	if o.VariableSchema != nil {
		toSerialize["variable_schema"] = o.VariableSchema
	}
	if o.PropertySchema != nil {
		toSerialize["property_schema"] = o.PropertySchema
	}
	if o.OutputSchema != nil {
		toSerialize["output_schema"] = o.OutputSchema
	}
	if o.ViewConfig != nil {
		toSerialize["view_config"] = o.ViewConfig
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
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
	if o.UpdatedBy.IsSet() {
		toSerialize["updated_by"] = o.UpdatedBy.Get()
	}
	if o.Owner.IsSet() {
		toSerialize["owner"] = o.Owner.Get()
	}
	if o.UniqueName.IsSet() {
		toSerialize["unique_name"] = o.UniqueName.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowSchema) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowSchema := _WorkflowSchema{}

	if err = json.Unmarshal(bytes, &varWorkflowSchema); err == nil {
		*o = WorkflowSchema(varWorkflowSchema)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "schema_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "title")
		delete(additionalProperties, "description")
		delete(additionalProperties, "type")
		delete(additionalProperties, "base_type")
		delete(additionalProperties, "version")
		delete(additionalProperties, "invisible")
		delete(additionalProperties, "inherits")
		delete(additionalProperties, "access_meta")
		delete(additionalProperties, "variable_schema")
		delete(additionalProperties, "property_schema")
		delete(additionalProperties, "output_schema")
		delete(additionalProperties, "view_config")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "created_on")
		delete(additionalProperties, "created_by")
		delete(additionalProperties, "updated_on")
		delete(additionalProperties, "updated_by")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "unique_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowSchema struct {
	value *WorkflowSchema
	isSet bool
}

func (v NullableWorkflowSchema) Get() *WorkflowSchema {
	return v.value
}

func (v *NullableWorkflowSchema) Set(val *WorkflowSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowSchema(val *WorkflowSchema) *NullableWorkflowSchema {
	return &NullableWorkflowSchema{value: val, isSet: true}
}

func (v NullableWorkflowSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


