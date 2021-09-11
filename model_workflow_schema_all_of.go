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

// WorkflowSchemaAllOf struct for WorkflowSchemaAllOf
type WorkflowSchemaAllOf struct {
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
	VariableSchema *map[string]interface{} `json:"variable_schema,omitempty"`
	PropertySchema *map[string]interface{} `json:"property_schema,omitempty"`
	OutputSchema *map[string]interface{} `json:"output_schema,omitempty"`
	ViewConfig *map[string]interface{} `json:"view_config,omitempty"`
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowSchemaAllOf WorkflowSchemaAllOf

// NewWorkflowSchemaAllOf instantiates a new WorkflowSchemaAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowSchemaAllOf() *WorkflowSchemaAllOf {
	this := WorkflowSchemaAllOf{}
	return &this
}

// NewWorkflowSchemaAllOfWithDefaults instantiates a new WorkflowSchemaAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowSchemaAllOfWithDefaults() *WorkflowSchemaAllOf {
	this := WorkflowSchemaAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowSchemaAllOf) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowSchemaAllOf) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *WorkflowSchemaAllOf) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *WorkflowSchemaAllOf) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *WorkflowSchemaAllOf) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *WorkflowSchemaAllOf) UnsetId() {
	o.Id.Unset()
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *WorkflowSchemaAllOf) GetSchemaId() string {
	if o == nil || o.SchemaId == nil {
		var ret string
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchemaAllOf) GetSchemaIdOk() (*string, bool) {
	if o == nil || o.SchemaId == nil {
		return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *WorkflowSchemaAllOf) HasSchemaId() bool {
	if o != nil && o.SchemaId != nil {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given string and assigns it to the SchemaId field.
func (o *WorkflowSchemaAllOf) SetSchemaId(v string) {
	o.SchemaId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowSchemaAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchemaAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowSchemaAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowSchemaAllOf) SetName(v string) {
	o.Name = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *WorkflowSchemaAllOf) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchemaAllOf) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *WorkflowSchemaAllOf) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *WorkflowSchemaAllOf) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowSchemaAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchemaAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowSchemaAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowSchemaAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkflowSchemaAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchemaAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkflowSchemaAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkflowSchemaAllOf) SetType(v string) {
	o.Type = &v
}

// GetBaseType returns the BaseType field value if set, zero value otherwise.
func (o *WorkflowSchemaAllOf) GetBaseType() string {
	if o == nil || o.BaseType == nil {
		var ret string
		return ret
	}
	return *o.BaseType
}

// GetBaseTypeOk returns a tuple with the BaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchemaAllOf) GetBaseTypeOk() (*string, bool) {
	if o == nil || o.BaseType == nil {
		return nil, false
	}
	return o.BaseType, true
}

// HasBaseType returns a boolean if a field has been set.
func (o *WorkflowSchemaAllOf) HasBaseType() bool {
	if o != nil && o.BaseType != nil {
		return true
	}

	return false
}

// SetBaseType gets a reference to the given string and assigns it to the BaseType field.
func (o *WorkflowSchemaAllOf) SetBaseType(v string) {
	o.BaseType = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *WorkflowSchemaAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchemaAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *WorkflowSchemaAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *WorkflowSchemaAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetInvisible returns the Invisible field value if set, zero value otherwise.
func (o *WorkflowSchemaAllOf) GetInvisible() bool {
	if o == nil || o.Invisible == nil {
		var ret bool
		return ret
	}
	return *o.Invisible
}

// GetInvisibleOk returns a tuple with the Invisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchemaAllOf) GetInvisibleOk() (*bool, bool) {
	if o == nil || o.Invisible == nil {
		return nil, false
	}
	return o.Invisible, true
}

// HasInvisible returns a boolean if a field has been set.
func (o *WorkflowSchemaAllOf) HasInvisible() bool {
	if o != nil && o.Invisible != nil {
		return true
	}

	return false
}

// SetInvisible gets a reference to the given bool and assigns it to the Invisible field.
func (o *WorkflowSchemaAllOf) SetInvisible(v bool) {
	o.Invisible = &v
}

// GetInherits returns the Inherits field value if set, zero value otherwise.
func (o *WorkflowSchemaAllOf) GetInherits() string {
	if o == nil || o.Inherits == nil {
		var ret string
		return ret
	}
	return *o.Inherits
}

// GetInheritsOk returns a tuple with the Inherits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchemaAllOf) GetInheritsOk() (*string, bool) {
	if o == nil || o.Inherits == nil {
		return nil, false
	}
	return o.Inherits, true
}

// HasInherits returns a boolean if a field has been set.
func (o *WorkflowSchemaAllOf) HasInherits() bool {
	if o != nil && o.Inherits != nil {
		return true
	}

	return false
}

// SetInherits gets a reference to the given string and assigns it to the Inherits field.
func (o *WorkflowSchemaAllOf) SetInherits(v string) {
	o.Inherits = &v
}

// GetAccessMeta returns the AccessMeta field value if set, zero value otherwise.
func (o *WorkflowSchemaAllOf) GetAccessMeta() WorkflowAccessMeta {
	if o == nil || o.AccessMeta == nil {
		var ret WorkflowAccessMeta
		return ret
	}
	return *o.AccessMeta
}

// GetAccessMetaOk returns a tuple with the AccessMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchemaAllOf) GetAccessMetaOk() (*WorkflowAccessMeta, bool) {
	if o == nil || o.AccessMeta == nil {
		return nil, false
	}
	return o.AccessMeta, true
}

// HasAccessMeta returns a boolean if a field has been set.
func (o *WorkflowSchemaAllOf) HasAccessMeta() bool {
	if o != nil && o.AccessMeta != nil {
		return true
	}

	return false
}

// SetAccessMeta gets a reference to the given WorkflowAccessMeta and assigns it to the AccessMeta field.
func (o *WorkflowSchemaAllOf) SetAccessMeta(v WorkflowAccessMeta) {
	o.AccessMeta = &v
}

// GetVariableSchema returns the VariableSchema field value if set, zero value otherwise.
func (o *WorkflowSchemaAllOf) GetVariableSchema() map[string]interface{} {
	if o == nil || o.VariableSchema == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.VariableSchema
}

// GetVariableSchemaOk returns a tuple with the VariableSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchemaAllOf) GetVariableSchemaOk() (*map[string]interface{}, bool) {
	if o == nil || o.VariableSchema == nil {
		return nil, false
	}
	return o.VariableSchema, true
}

// HasVariableSchema returns a boolean if a field has been set.
func (o *WorkflowSchemaAllOf) HasVariableSchema() bool {
	if o != nil && o.VariableSchema != nil {
		return true
	}

	return false
}

// SetVariableSchema gets a reference to the given map[string]interface{} and assigns it to the VariableSchema field.
func (o *WorkflowSchemaAllOf) SetVariableSchema(v map[string]interface{}) {
	o.VariableSchema = &v
}

// GetPropertySchema returns the PropertySchema field value if set, zero value otherwise.
func (o *WorkflowSchemaAllOf) GetPropertySchema() map[string]interface{} {
	if o == nil || o.PropertySchema == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.PropertySchema
}

// GetPropertySchemaOk returns a tuple with the PropertySchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchemaAllOf) GetPropertySchemaOk() (*map[string]interface{}, bool) {
	if o == nil || o.PropertySchema == nil {
		return nil, false
	}
	return o.PropertySchema, true
}

// HasPropertySchema returns a boolean if a field has been set.
func (o *WorkflowSchemaAllOf) HasPropertySchema() bool {
	if o != nil && o.PropertySchema != nil {
		return true
	}

	return false
}

// SetPropertySchema gets a reference to the given map[string]interface{} and assigns it to the PropertySchema field.
func (o *WorkflowSchemaAllOf) SetPropertySchema(v map[string]interface{}) {
	o.PropertySchema = &v
}

// GetOutputSchema returns the OutputSchema field value if set, zero value otherwise.
func (o *WorkflowSchemaAllOf) GetOutputSchema() map[string]interface{} {
	if o == nil || o.OutputSchema == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.OutputSchema
}

// GetOutputSchemaOk returns a tuple with the OutputSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchemaAllOf) GetOutputSchemaOk() (*map[string]interface{}, bool) {
	if o == nil || o.OutputSchema == nil {
		return nil, false
	}
	return o.OutputSchema, true
}

// HasOutputSchema returns a boolean if a field has been set.
func (o *WorkflowSchemaAllOf) HasOutputSchema() bool {
	if o != nil && o.OutputSchema != nil {
		return true
	}

	return false
}

// SetOutputSchema gets a reference to the given map[string]interface{} and assigns it to the OutputSchema field.
func (o *WorkflowSchemaAllOf) SetOutputSchema(v map[string]interface{}) {
	o.OutputSchema = &v
}

// GetViewConfig returns the ViewConfig field value if set, zero value otherwise.
func (o *WorkflowSchemaAllOf) GetViewConfig() map[string]interface{} {
	if o == nil || o.ViewConfig == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.ViewConfig
}

// GetViewConfigOk returns a tuple with the ViewConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchemaAllOf) GetViewConfigOk() (*map[string]interface{}, bool) {
	if o == nil || o.ViewConfig == nil {
		return nil, false
	}
	return o.ViewConfig, true
}

// HasViewConfig returns a boolean if a field has been set.
func (o *WorkflowSchemaAllOf) HasViewConfig() bool {
	if o != nil && o.ViewConfig != nil {
		return true
	}

	return false
}

// SetViewConfig gets a reference to the given map[string]interface{} and assigns it to the ViewConfig field.
func (o *WorkflowSchemaAllOf) SetViewConfig(v map[string]interface{}) {
	o.ViewConfig = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *WorkflowSchemaAllOf) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSchemaAllOf) GetAttributesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *WorkflowSchemaAllOf) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *WorkflowSchemaAllOf) SetAttributes(v map[string]interface{}) {
	o.Attributes = &v
}

func (o WorkflowSchemaAllOf) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowSchemaAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowSchemaAllOf := _WorkflowSchemaAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowSchemaAllOf); err == nil {
		*o = WorkflowSchemaAllOf(varWorkflowSchemaAllOf)
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
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowSchemaAllOf struct {
	value *WorkflowSchemaAllOf
	isSet bool
}

func (v NullableWorkflowSchemaAllOf) Get() *WorkflowSchemaAllOf {
	return v.value
}

func (v *NullableWorkflowSchemaAllOf) Set(val *WorkflowSchemaAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowSchemaAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowSchemaAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowSchemaAllOf(val *WorkflowSchemaAllOf) *NullableWorkflowSchemaAllOf {
	return &NullableWorkflowSchemaAllOf{value: val, isSet: true}
}

func (v NullableWorkflowSchemaAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowSchemaAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


