/*
 * KAKAPO - MSX SDK
 *
 * MSX Platform SDK
 *
 * API version: 1.0.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msxsdk

import (
	"encoding/json"
)

// WorkflowInstanceAllOf struct for WorkflowInstanceAllOf
type WorkflowInstanceAllOf struct {
	Id *string `json:"id,omitempty"`
	DefinitionId *string `json:"definition_id,omitempty"`
	Name *string `json:"name,omitempty"`
	SchemaId *string `json:"schema_id,omitempty"`
	Version *string `json:"version,omitempty"`
	Type *string `json:"type,omitempty"`
	BaseType *string `json:"base_type,omitempty"`
	Properties *map[string]map[string]interface{} `json:"properties,omitempty"`
	Actions *[]WorkflowAction `json:"actions,omitempty"`
	Variables *[]WorkflowVariable `json:"variables,omitempty"`
	Status *map[string]map[string]interface{} `json:"status,omitempty"`
	StartedOn *string `json:"started_on,omitempty"`
	EndedOn *string `json:"ended_on,omitempty"`
}

// NewWorkflowInstanceAllOf instantiates a new WorkflowInstanceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowInstanceAllOf() *WorkflowInstanceAllOf {
	this := WorkflowInstanceAllOf{}
	return &this
}

// NewWorkflowInstanceAllOfWithDefaults instantiates a new WorkflowInstanceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowInstanceAllOfWithDefaults() *WorkflowInstanceAllOf {
	this := WorkflowInstanceAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkflowInstanceAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkflowInstanceAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkflowInstanceAllOf) SetId(v string) {
	o.Id = &v
}

// GetDefinitionId returns the DefinitionId field value if set, zero value otherwise.
func (o *WorkflowInstanceAllOf) GetDefinitionId() string {
	if o == nil || o.DefinitionId == nil {
		var ret string
		return ret
	}
	return *o.DefinitionId
}

// GetDefinitionIdOk returns a tuple with the DefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceAllOf) GetDefinitionIdOk() (*string, bool) {
	if o == nil || o.DefinitionId == nil {
		return nil, false
	}
	return o.DefinitionId, true
}

// HasDefinitionId returns a boolean if a field has been set.
func (o *WorkflowInstanceAllOf) HasDefinitionId() bool {
	if o != nil && o.DefinitionId != nil {
		return true
	}

	return false
}

// SetDefinitionId gets a reference to the given string and assigns it to the DefinitionId field.
func (o *WorkflowInstanceAllOf) SetDefinitionId(v string) {
	o.DefinitionId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowInstanceAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowInstanceAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowInstanceAllOf) SetName(v string) {
	o.Name = &v
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *WorkflowInstanceAllOf) GetSchemaId() string {
	if o == nil || o.SchemaId == nil {
		var ret string
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceAllOf) GetSchemaIdOk() (*string, bool) {
	if o == nil || o.SchemaId == nil {
		return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *WorkflowInstanceAllOf) HasSchemaId() bool {
	if o != nil && o.SchemaId != nil {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given string and assigns it to the SchemaId field.
func (o *WorkflowInstanceAllOf) SetSchemaId(v string) {
	o.SchemaId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *WorkflowInstanceAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *WorkflowInstanceAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *WorkflowInstanceAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkflowInstanceAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkflowInstanceAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkflowInstanceAllOf) SetType(v string) {
	o.Type = &v
}

// GetBaseType returns the BaseType field value if set, zero value otherwise.
func (o *WorkflowInstanceAllOf) GetBaseType() string {
	if o == nil || o.BaseType == nil {
		var ret string
		return ret
	}
	return *o.BaseType
}

// GetBaseTypeOk returns a tuple with the BaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceAllOf) GetBaseTypeOk() (*string, bool) {
	if o == nil || o.BaseType == nil {
		return nil, false
	}
	return o.BaseType, true
}

// HasBaseType returns a boolean if a field has been set.
func (o *WorkflowInstanceAllOf) HasBaseType() bool {
	if o != nil && o.BaseType != nil {
		return true
	}

	return false
}

// SetBaseType gets a reference to the given string and assigns it to the BaseType field.
func (o *WorkflowInstanceAllOf) SetBaseType(v string) {
	o.BaseType = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *WorkflowInstanceAllOf) GetProperties() map[string]map[string]interface{} {
	if o == nil || o.Properties == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceAllOf) GetPropertiesOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *WorkflowInstanceAllOf) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]map[string]interface{} and assigns it to the Properties field.
func (o *WorkflowInstanceAllOf) SetProperties(v map[string]map[string]interface{}) {
	o.Properties = &v
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *WorkflowInstanceAllOf) GetActions() []WorkflowAction {
	if o == nil || o.Actions == nil {
		var ret []WorkflowAction
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceAllOf) GetActionsOk() (*[]WorkflowAction, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *WorkflowInstanceAllOf) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given []WorkflowAction and assigns it to the Actions field.
func (o *WorkflowInstanceAllOf) SetActions(v []WorkflowAction) {
	o.Actions = &v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *WorkflowInstanceAllOf) GetVariables() []WorkflowVariable {
	if o == nil || o.Variables == nil {
		var ret []WorkflowVariable
		return ret
	}
	return *o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceAllOf) GetVariablesOk() (*[]WorkflowVariable, bool) {
	if o == nil || o.Variables == nil {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *WorkflowInstanceAllOf) HasVariables() bool {
	if o != nil && o.Variables != nil {
		return true
	}

	return false
}

// SetVariables gets a reference to the given []WorkflowVariable and assigns it to the Variables field.
func (o *WorkflowInstanceAllOf) SetVariables(v []WorkflowVariable) {
	o.Variables = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WorkflowInstanceAllOf) GetStatus() map[string]map[string]interface{} {
	if o == nil || o.Status == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceAllOf) GetStatusOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WorkflowInstanceAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given map[string]map[string]interface{} and assigns it to the Status field.
func (o *WorkflowInstanceAllOf) SetStatus(v map[string]map[string]interface{}) {
	o.Status = &v
}

// GetStartedOn returns the StartedOn field value if set, zero value otherwise.
func (o *WorkflowInstanceAllOf) GetStartedOn() string {
	if o == nil || o.StartedOn == nil {
		var ret string
		return ret
	}
	return *o.StartedOn
}

// GetStartedOnOk returns a tuple with the StartedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceAllOf) GetStartedOnOk() (*string, bool) {
	if o == nil || o.StartedOn == nil {
		return nil, false
	}
	return o.StartedOn, true
}

// HasStartedOn returns a boolean if a field has been set.
func (o *WorkflowInstanceAllOf) HasStartedOn() bool {
	if o != nil && o.StartedOn != nil {
		return true
	}

	return false
}

// SetStartedOn gets a reference to the given string and assigns it to the StartedOn field.
func (o *WorkflowInstanceAllOf) SetStartedOn(v string) {
	o.StartedOn = &v
}

// GetEndedOn returns the EndedOn field value if set, zero value otherwise.
func (o *WorkflowInstanceAllOf) GetEndedOn() string {
	if o == nil || o.EndedOn == nil {
		var ret string
		return ret
	}
	return *o.EndedOn
}

// GetEndedOnOk returns a tuple with the EndedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInstanceAllOf) GetEndedOnOk() (*string, bool) {
	if o == nil || o.EndedOn == nil {
		return nil, false
	}
	return o.EndedOn, true
}

// HasEndedOn returns a boolean if a field has been set.
func (o *WorkflowInstanceAllOf) HasEndedOn() bool {
	if o != nil && o.EndedOn != nil {
		return true
	}

	return false
}

// SetEndedOn gets a reference to the given string and assigns it to the EndedOn field.
func (o *WorkflowInstanceAllOf) SetEndedOn(v string) {
	o.EndedOn = &v
}

func (o WorkflowInstanceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DefinitionId != nil {
		toSerialize["definition_id"] = o.DefinitionId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SchemaId != nil {
		toSerialize["schema_id"] = o.SchemaId
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.BaseType != nil {
		toSerialize["base_type"] = o.BaseType
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.Variables != nil {
		toSerialize["variables"] = o.Variables
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StartedOn != nil {
		toSerialize["started_on"] = o.StartedOn
	}
	if o.EndedOn != nil {
		toSerialize["ended_on"] = o.EndedOn
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowInstanceAllOf struct {
	value *WorkflowInstanceAllOf
	isSet bool
}

func (v NullableWorkflowInstanceAllOf) Get() *WorkflowInstanceAllOf {
	return v.value
}

func (v *NullableWorkflowInstanceAllOf) Set(val *WorkflowInstanceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowInstanceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowInstanceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowInstanceAllOf(val *WorkflowInstanceAllOf) *NullableWorkflowInstanceAllOf {
	return &NullableWorkflowInstanceAllOf{value: val, isSet: true}
}

func (v NullableWorkflowInstanceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowInstanceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


