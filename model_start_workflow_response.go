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

// StartWorkflowResponse struct for StartWorkflowResponse
type StartWorkflowResponse struct {
	Id *string `json:"id,omitempty"`
	DefinitionId *string `json:"definition_id,omitempty"`
	SchemaId *string `json:"schema_id,omitempty"`
	Type *string `json:"type,omitempty"`
	BaseType *string `json:"base_type,omitempty"`
	Status *map[string]map[string]interface{} `json:"status,omitempty"`
	Version *string `json:"version,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	CreatedOn *string `json:"created_on,omitempty"`
}

// NewStartWorkflowResponse instantiates a new StartWorkflowResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartWorkflowResponse() *StartWorkflowResponse {
	this := StartWorkflowResponse{}
	return &this
}

// NewStartWorkflowResponseWithDefaults instantiates a new StartWorkflowResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartWorkflowResponseWithDefaults() *StartWorkflowResponse {
	this := StartWorkflowResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StartWorkflowResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartWorkflowResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StartWorkflowResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StartWorkflowResponse) SetId(v string) {
	o.Id = &v
}

// GetDefinitionId returns the DefinitionId field value if set, zero value otherwise.
func (o *StartWorkflowResponse) GetDefinitionId() string {
	if o == nil || o.DefinitionId == nil {
		var ret string
		return ret
	}
	return *o.DefinitionId
}

// GetDefinitionIdOk returns a tuple with the DefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartWorkflowResponse) GetDefinitionIdOk() (*string, bool) {
	if o == nil || o.DefinitionId == nil {
		return nil, false
	}
	return o.DefinitionId, true
}

// HasDefinitionId returns a boolean if a field has been set.
func (o *StartWorkflowResponse) HasDefinitionId() bool {
	if o != nil && o.DefinitionId != nil {
		return true
	}

	return false
}

// SetDefinitionId gets a reference to the given string and assigns it to the DefinitionId field.
func (o *StartWorkflowResponse) SetDefinitionId(v string) {
	o.DefinitionId = &v
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *StartWorkflowResponse) GetSchemaId() string {
	if o == nil || o.SchemaId == nil {
		var ret string
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartWorkflowResponse) GetSchemaIdOk() (*string, bool) {
	if o == nil || o.SchemaId == nil {
		return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *StartWorkflowResponse) HasSchemaId() bool {
	if o != nil && o.SchemaId != nil {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given string and assigns it to the SchemaId field.
func (o *StartWorkflowResponse) SetSchemaId(v string) {
	o.SchemaId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StartWorkflowResponse) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartWorkflowResponse) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StartWorkflowResponse) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StartWorkflowResponse) SetType(v string) {
	o.Type = &v
}

// GetBaseType returns the BaseType field value if set, zero value otherwise.
func (o *StartWorkflowResponse) GetBaseType() string {
	if o == nil || o.BaseType == nil {
		var ret string
		return ret
	}
	return *o.BaseType
}

// GetBaseTypeOk returns a tuple with the BaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartWorkflowResponse) GetBaseTypeOk() (*string, bool) {
	if o == nil || o.BaseType == nil {
		return nil, false
	}
	return o.BaseType, true
}

// HasBaseType returns a boolean if a field has been set.
func (o *StartWorkflowResponse) HasBaseType() bool {
	if o != nil && o.BaseType != nil {
		return true
	}

	return false
}

// SetBaseType gets a reference to the given string and assigns it to the BaseType field.
func (o *StartWorkflowResponse) SetBaseType(v string) {
	o.BaseType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StartWorkflowResponse) GetStatus() map[string]map[string]interface{} {
	if o == nil || o.Status == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartWorkflowResponse) GetStatusOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StartWorkflowResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given map[string]map[string]interface{} and assigns it to the Status field.
func (o *StartWorkflowResponse) SetStatus(v map[string]map[string]interface{}) {
	o.Status = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *StartWorkflowResponse) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartWorkflowResponse) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *StartWorkflowResponse) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *StartWorkflowResponse) SetVersion(v string) {
	o.Version = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *StartWorkflowResponse) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartWorkflowResponse) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *StartWorkflowResponse) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *StartWorkflowResponse) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *StartWorkflowResponse) GetCreatedOn() string {
	if o == nil || o.CreatedOn == nil {
		var ret string
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartWorkflowResponse) GetCreatedOnOk() (*string, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *StartWorkflowResponse) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given string and assigns it to the CreatedOn field.
func (o *StartWorkflowResponse) SetCreatedOn(v string) {
	o.CreatedOn = &v
}

func (o StartWorkflowResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DefinitionId != nil {
		toSerialize["definition_id"] = o.DefinitionId
	}
	if o.SchemaId != nil {
		toSerialize["schema_id"] = o.SchemaId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.BaseType != nil {
		toSerialize["base_type"] = o.BaseType
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.CreatedOn != nil {
		toSerialize["created_on"] = o.CreatedOn
	}
	return json.Marshal(toSerialize)
}

type NullableStartWorkflowResponse struct {
	value *StartWorkflowResponse
	isSet bool
}

func (v NullableStartWorkflowResponse) Get() *StartWorkflowResponse {
	return v.value
}

func (v *NullableStartWorkflowResponse) Set(val *StartWorkflowResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStartWorkflowResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStartWorkflowResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartWorkflowResponse(val *StartWorkflowResponse) *NullableStartWorkflowResponse {
	return &NullableStartWorkflowResponse{value: val, isSet: true}
}

func (v NullableStartWorkflowResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartWorkflowResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


