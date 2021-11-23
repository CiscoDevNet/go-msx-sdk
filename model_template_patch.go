/*
 * MSX SDK
 *
 * MSX SDK client.
 *
 * API version: 1.0.8
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msxsdk

import (
	"encoding/json"
)

// TemplatePatch struct for TemplatePatch
type TemplatePatch struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ServiceType *string `json:"serviceType,omitempty"`
	Type *string `json:"type,omitempty"`
	Subtype NullableString `json:"subtype,omitempty"`
	Configuration *string `json:"configuration,omitempty"`
	Attributes *map[string]string `json:"attributes,omitempty"`
	Tags *[]string `json:"tags,omitempty"`
	Notes *string `json:"notes,omitempty"`
	Status *TemplateStatus `json:"status,omitempty"`
	StatusDetails *string `json:"statusDetails,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TemplatePatch TemplatePatch

// NewTemplatePatch instantiates a new TemplatePatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplatePatch() *TemplatePatch {
	this := TemplatePatch{}
	return &this
}

// NewTemplatePatchWithDefaults instantiates a new TemplatePatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplatePatchWithDefaults() *TemplatePatch {
	this := TemplatePatch{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TemplatePatch) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatePatch) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TemplatePatch) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TemplatePatch) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TemplatePatch) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatePatch) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TemplatePatch) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TemplatePatch) SetDescription(v string) {
	o.Description = &v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *TemplatePatch) GetServiceType() string {
	if o == nil || o.ServiceType == nil {
		var ret string
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatePatch) GetServiceTypeOk() (*string, bool) {
	if o == nil || o.ServiceType == nil {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *TemplatePatch) HasServiceType() bool {
	if o != nil && o.ServiceType != nil {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given string and assigns it to the ServiceType field.
func (o *TemplatePatch) SetServiceType(v string) {
	o.ServiceType = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TemplatePatch) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatePatch) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TemplatePatch) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TemplatePatch) SetType(v string) {
	o.Type = &v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplatePatch) GetSubtype() string {
	if o == nil || o.Subtype.Get() == nil {
		var ret string
		return ret
	}
	return *o.Subtype.Get()
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplatePatch) GetSubtypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Subtype.Get(), o.Subtype.IsSet()
}

// HasSubtype returns a boolean if a field has been set.
func (o *TemplatePatch) HasSubtype() bool {
	if o != nil && o.Subtype.IsSet() {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given NullableString and assigns it to the Subtype field.
func (o *TemplatePatch) SetSubtype(v string) {
	o.Subtype.Set(&v)
}
// SetSubtypeNil sets the value for Subtype to be an explicit nil
func (o *TemplatePatch) SetSubtypeNil() {
	o.Subtype.Set(nil)
}

// UnsetSubtype ensures that no value is present for Subtype, not even an explicit nil
func (o *TemplatePatch) UnsetSubtype() {
	o.Subtype.Unset()
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *TemplatePatch) GetConfiguration() string {
	if o == nil || o.Configuration == nil {
		var ret string
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatePatch) GetConfigurationOk() (*string, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *TemplatePatch) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given string and assigns it to the Configuration field.
func (o *TemplatePatch) SetConfiguration(v string) {
	o.Configuration = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TemplatePatch) GetAttributes() map[string]string {
	if o == nil || o.Attributes == nil {
		var ret map[string]string
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatePatch) GetAttributesOk() (*map[string]string, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TemplatePatch) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]string and assigns it to the Attributes field.
func (o *TemplatePatch) SetAttributes(v map[string]string) {
	o.Attributes = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *TemplatePatch) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatePatch) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TemplatePatch) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *TemplatePatch) SetTags(v []string) {
	o.Tags = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *TemplatePatch) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatePatch) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *TemplatePatch) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *TemplatePatch) SetNotes(v string) {
	o.Notes = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TemplatePatch) GetStatus() TemplateStatus {
	if o == nil || o.Status == nil {
		var ret TemplateStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatePatch) GetStatusOk() (*TemplateStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TemplatePatch) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given TemplateStatus and assigns it to the Status field.
func (o *TemplatePatch) SetStatus(v TemplateStatus) {
	o.Status = &v
}

// GetStatusDetails returns the StatusDetails field value if set, zero value otherwise.
func (o *TemplatePatch) GetStatusDetails() string {
	if o == nil || o.StatusDetails == nil {
		var ret string
		return ret
	}
	return *o.StatusDetails
}

// GetStatusDetailsOk returns a tuple with the StatusDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatePatch) GetStatusDetailsOk() (*string, bool) {
	if o == nil || o.StatusDetails == nil {
		return nil, false
	}
	return o.StatusDetails, true
}

// HasStatusDetails returns a boolean if a field has been set.
func (o *TemplatePatch) HasStatusDetails() bool {
	if o != nil && o.StatusDetails != nil {
		return true
	}

	return false
}

// SetStatusDetails gets a reference to the given string and assigns it to the StatusDetails field.
func (o *TemplatePatch) SetStatusDetails(v string) {
	o.StatusDetails = &v
}

func (o TemplatePatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ServiceType != nil {
		toSerialize["serviceType"] = o.ServiceType
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Subtype.IsSet() {
		toSerialize["subtype"] = o.Subtype.Get()
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StatusDetails != nil {
		toSerialize["statusDetails"] = o.StatusDetails
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TemplatePatch) UnmarshalJSON(bytes []byte) (err error) {
	varTemplatePatch := _TemplatePatch{}

	if err = json.Unmarshal(bytes, &varTemplatePatch); err == nil {
		*o = TemplatePatch(varTemplatePatch)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "serviceType")
		delete(additionalProperties, "type")
		delete(additionalProperties, "subtype")
		delete(additionalProperties, "configuration")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "notes")
		delete(additionalProperties, "status")
		delete(additionalProperties, "statusDetails")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTemplatePatch struct {
	value *TemplatePatch
	isSet bool
}

func (v NullableTemplatePatch) Get() *TemplatePatch {
	return v.value
}

func (v *NullableTemplatePatch) Set(val *TemplatePatch) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplatePatch) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplatePatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplatePatch(val *TemplatePatch) *NullableTemplatePatch {
	return &NullableTemplatePatch{value: val, isSet: true}
}

func (v NullableTemplatePatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplatePatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


