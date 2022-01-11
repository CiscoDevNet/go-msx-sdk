/*
 * MSX SDK
 *
 * MSX SDK client.
 *
 * API version: 1.0.9
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msxsdk

import (
	"encoding/json"
	"time"
)

// IncidentAllOf struct for IncidentAllOf
type IncidentAllOf struct {
	Id string `json:"id"`
	ExternalId string `json:"externalId"`
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	Notes *[]string `json:"notes,omitempty"`
	CreatedBy *string `json:"createdBy,omitempty"`
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	ModifiedBy *string `json:"modifiedBy,omitempty"`
	ModifiedOn *time.Time `json:"modifiedOn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IncidentAllOf IncidentAllOf

// NewIncidentAllOf instantiates a new IncidentAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentAllOf(id string, externalId string) *IncidentAllOf {
	this := IncidentAllOf{}
	this.Id = id
	this.ExternalId = externalId
	return &this
}

// NewIncidentAllOfWithDefaults instantiates a new IncidentAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentAllOfWithDefaults() *IncidentAllOf {
	this := IncidentAllOf{}
	return &this
}

// GetId returns the Id field value
func (o *IncidentAllOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IncidentAllOf) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IncidentAllOf) SetId(v string) {
	o.Id = v
}

// GetExternalId returns the ExternalId field value
func (o *IncidentAllOf) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *IncidentAllOf) GetExternalIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *IncidentAllOf) SetExternalId(v string) {
	o.ExternalId = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *IncidentAllOf) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentAllOf) GetAttributesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *IncidentAllOf) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *IncidentAllOf) SetAttributes(v map[string]interface{}) {
	o.Attributes = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *IncidentAllOf) GetNotes() []string {
	if o == nil || o.Notes == nil {
		var ret []string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentAllOf) GetNotesOk() (*[]string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *IncidentAllOf) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given []string and assigns it to the Notes field.
func (o *IncidentAllOf) SetNotes(v []string) {
	o.Notes = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *IncidentAllOf) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentAllOf) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *IncidentAllOf) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *IncidentAllOf) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *IncidentAllOf) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentAllOf) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *IncidentAllOf) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *IncidentAllOf) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *IncidentAllOf) GetModifiedBy() string {
	if o == nil || o.ModifiedBy == nil {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentAllOf) GetModifiedByOk() (*string, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *IncidentAllOf) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy != nil {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *IncidentAllOf) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetModifiedOn returns the ModifiedOn field value if set, zero value otherwise.
func (o *IncidentAllOf) GetModifiedOn() time.Time {
	if o == nil || o.ModifiedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentAllOf) GetModifiedOnOk() (*time.Time, bool) {
	if o == nil || o.ModifiedOn == nil {
		return nil, false
	}
	return o.ModifiedOn, true
}

// HasModifiedOn returns a boolean if a field has been set.
func (o *IncidentAllOf) HasModifiedOn() bool {
	if o != nil && o.ModifiedOn != nil {
		return true
	}

	return false
}

// SetModifiedOn gets a reference to the given time.Time and assigns it to the ModifiedOn field.
func (o *IncidentAllOf) SetModifiedOn(v time.Time) {
	o.ModifiedOn = &v
}

func (o IncidentAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["externalId"] = o.ExternalId
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.CreatedOn != nil {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if o.ModifiedBy != nil {
		toSerialize["modifiedBy"] = o.ModifiedBy
	}
	if o.ModifiedOn != nil {
		toSerialize["modifiedOn"] = o.ModifiedOn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IncidentAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIncidentAllOf := _IncidentAllOf{}

	if err = json.Unmarshal(bytes, &varIncidentAllOf); err == nil {
		*o = IncidentAllOf(varIncidentAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "notes")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "createdOn")
		delete(additionalProperties, "modifiedBy")
		delete(additionalProperties, "modifiedOn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIncidentAllOf struct {
	value *IncidentAllOf
	isSet bool
}

func (v NullableIncidentAllOf) Get() *IncidentAllOf {
	return v.value
}

func (v *NullableIncidentAllOf) Set(val *IncidentAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentAllOf(val *IncidentAllOf) *NullableIncidentAllOf {
	return &NullableIncidentAllOf{value: val, isSet: true}
}

func (v NullableIncidentAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


