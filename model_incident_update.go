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

// IncidentUpdate struct for IncidentUpdate
type IncidentUpdate struct {
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// inquiry|software|hardware|network|database
	Category *string `json:"category,omitempty"`
	Description string `json:"description"`
	Impact *string `json:"impact,omitempty"`
	Priority *string `json:"priority,omitempty"`
	Severity *string `json:"severity,omitempty"`
	State *string `json:"state,omitempty"`
	Subcategory *string `json:"subcategory,omitempty"`
	Tenant NullableString `json:"tenant,omitempty"`
	Urgency *string `json:"urgency,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IncidentUpdate IncidentUpdate

// NewIncidentUpdate instantiates a new IncidentUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentUpdate(description string) *IncidentUpdate {
	this := IncidentUpdate{}
	var category string = "inquiry"
	this.Category = &category
	this.Description = description
	var impact string = "Low"
	this.Impact = &impact
	var priority string = "Planning"
	this.Priority = &priority
	var severity string = "Low"
	this.Severity = &severity
	var state string = "New"
	this.State = &state
	var urgency string = "Low"
	this.Urgency = &urgency
	return &this
}

// NewIncidentUpdateWithDefaults instantiates a new IncidentUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentUpdateWithDefaults() *IncidentUpdate {
	this := IncidentUpdate{}
	var category string = "inquiry"
	this.Category = &category
	var impact string = "Low"
	this.Impact = &impact
	var priority string = "Planning"
	this.Priority = &priority
	var severity string = "Low"
	this.Severity = &severity
	var state string = "New"
	this.State = &state
	var urgency string = "Low"
	this.Urgency = &urgency
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *IncidentUpdate) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUpdate) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *IncidentUpdate) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *IncidentUpdate) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *IncidentUpdate) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUpdate) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *IncidentUpdate) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *IncidentUpdate) SetCategory(v string) {
	o.Category = &v
}

// GetDescription returns the Description field value
func (o *IncidentUpdate) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *IncidentUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *IncidentUpdate) SetDescription(v string) {
	o.Description = v
}

// GetImpact returns the Impact field value if set, zero value otherwise.
func (o *IncidentUpdate) GetImpact() string {
	if o == nil || o.Impact == nil {
		var ret string
		return ret
	}
	return *o.Impact
}

// GetImpactOk returns a tuple with the Impact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUpdate) GetImpactOk() (*string, bool) {
	if o == nil || o.Impact == nil {
		return nil, false
	}
	return o.Impact, true
}

// HasImpact returns a boolean if a field has been set.
func (o *IncidentUpdate) HasImpact() bool {
	if o != nil && o.Impact != nil {
		return true
	}

	return false
}

// SetImpact gets a reference to the given string and assigns it to the Impact field.
func (o *IncidentUpdate) SetImpact(v string) {
	o.Impact = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *IncidentUpdate) GetPriority() string {
	if o == nil || o.Priority == nil {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUpdate) GetPriorityOk() (*string, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *IncidentUpdate) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *IncidentUpdate) SetPriority(v string) {
	o.Priority = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *IncidentUpdate) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUpdate) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *IncidentUpdate) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *IncidentUpdate) SetSeverity(v string) {
	o.Severity = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *IncidentUpdate) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUpdate) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *IncidentUpdate) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *IncidentUpdate) SetState(v string) {
	o.State = &v
}

// GetSubcategory returns the Subcategory field value if set, zero value otherwise.
func (o *IncidentUpdate) GetSubcategory() string {
	if o == nil || o.Subcategory == nil {
		var ret string
		return ret
	}
	return *o.Subcategory
}

// GetSubcategoryOk returns a tuple with the Subcategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUpdate) GetSubcategoryOk() (*string, bool) {
	if o == nil || o.Subcategory == nil {
		return nil, false
	}
	return o.Subcategory, true
}

// HasSubcategory returns a boolean if a field has been set.
func (o *IncidentUpdate) HasSubcategory() bool {
	if o != nil && o.Subcategory != nil {
		return true
	}

	return false
}

// SetSubcategory gets a reference to the given string and assigns it to the Subcategory field.
func (o *IncidentUpdate) SetSubcategory(v string) {
	o.Subcategory = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentUpdate) GetTenant() string {
	if o == nil || o.Tenant.Get() == nil {
		var ret string
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentUpdate) GetTenantOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *IncidentUpdate) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableString and assigns it to the Tenant field.
func (o *IncidentUpdate) SetTenant(v string) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *IncidentUpdate) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *IncidentUpdate) UnsetTenant() {
	o.Tenant.Unset()
}

// GetUrgency returns the Urgency field value if set, zero value otherwise.
func (o *IncidentUpdate) GetUrgency() string {
	if o == nil || o.Urgency == nil {
		var ret string
		return ret
	}
	return *o.Urgency
}

// GetUrgencyOk returns a tuple with the Urgency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUpdate) GetUrgencyOk() (*string, bool) {
	if o == nil || o.Urgency == nil {
		return nil, false
	}
	return o.Urgency, true
}

// HasUrgency returns a boolean if a field has been set.
func (o *IncidentUpdate) HasUrgency() bool {
	if o != nil && o.Urgency != nil {
		return true
	}

	return false
}

// SetUrgency gets a reference to the given string and assigns it to the Urgency field.
func (o *IncidentUpdate) SetUrgency(v string) {
	o.Urgency = &v
}

func (o IncidentUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.Impact != nil {
		toSerialize["impact"] = o.Impact
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Subcategory != nil {
		toSerialize["subcategory"] = o.Subcategory
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.Urgency != nil {
		toSerialize["urgency"] = o.Urgency
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IncidentUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varIncidentUpdate := _IncidentUpdate{}

	if err = json.Unmarshal(bytes, &varIncidentUpdate); err == nil {
		*o = IncidentUpdate(varIncidentUpdate)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "category")
		delete(additionalProperties, "description")
		delete(additionalProperties, "impact")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "severity")
		delete(additionalProperties, "state")
		delete(additionalProperties, "subcategory")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "urgency")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIncidentUpdate struct {
	value *IncidentUpdate
	isSet bool
}

func (v NullableIncidentUpdate) Get() *IncidentUpdate {
	return v.value
}

func (v *NullableIncidentUpdate) Set(val *IncidentUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentUpdate(val *IncidentUpdate) *NullableIncidentUpdate {
	return &NullableIncidentUpdate{value: val, isSet: true}
}

func (v NullableIncidentUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


