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
)

// IncidentCreate struct for IncidentCreate
type IncidentCreate struct {
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
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

type _IncidentCreate IncidentCreate

// NewIncidentCreate instantiates a new IncidentCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentCreate(description string) *IncidentCreate {
	this := IncidentCreate{}
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

// NewIncidentCreateWithDefaults instantiates a new IncidentCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentCreateWithDefaults() *IncidentCreate {
	this := IncidentCreate{}
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
func (o *IncidentCreate) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentCreate) GetAttributesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *IncidentCreate) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *IncidentCreate) SetAttributes(v map[string]interface{}) {
	o.Attributes = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *IncidentCreate) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentCreate) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *IncidentCreate) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *IncidentCreate) SetCategory(v string) {
	o.Category = &v
}

// GetDescription returns the Description field value
func (o *IncidentCreate) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *IncidentCreate) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *IncidentCreate) SetDescription(v string) {
	o.Description = v
}

// GetImpact returns the Impact field value if set, zero value otherwise.
func (o *IncidentCreate) GetImpact() string {
	if o == nil || o.Impact == nil {
		var ret string
		return ret
	}
	return *o.Impact
}

// GetImpactOk returns a tuple with the Impact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentCreate) GetImpactOk() (*string, bool) {
	if o == nil || o.Impact == nil {
		return nil, false
	}
	return o.Impact, true
}

// HasImpact returns a boolean if a field has been set.
func (o *IncidentCreate) HasImpact() bool {
	if o != nil && o.Impact != nil {
		return true
	}

	return false
}

// SetImpact gets a reference to the given string and assigns it to the Impact field.
func (o *IncidentCreate) SetImpact(v string) {
	o.Impact = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *IncidentCreate) GetPriority() string {
	if o == nil || o.Priority == nil {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentCreate) GetPriorityOk() (*string, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *IncidentCreate) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *IncidentCreate) SetPriority(v string) {
	o.Priority = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *IncidentCreate) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentCreate) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *IncidentCreate) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *IncidentCreate) SetSeverity(v string) {
	o.Severity = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *IncidentCreate) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentCreate) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *IncidentCreate) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *IncidentCreate) SetState(v string) {
	o.State = &v
}

// GetSubcategory returns the Subcategory field value if set, zero value otherwise.
func (o *IncidentCreate) GetSubcategory() string {
	if o == nil || o.Subcategory == nil {
		var ret string
		return ret
	}
	return *o.Subcategory
}

// GetSubcategoryOk returns a tuple with the Subcategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentCreate) GetSubcategoryOk() (*string, bool) {
	if o == nil || o.Subcategory == nil {
		return nil, false
	}
	return o.Subcategory, true
}

// HasSubcategory returns a boolean if a field has been set.
func (o *IncidentCreate) HasSubcategory() bool {
	if o != nil && o.Subcategory != nil {
		return true
	}

	return false
}

// SetSubcategory gets a reference to the given string and assigns it to the Subcategory field.
func (o *IncidentCreate) SetSubcategory(v string) {
	o.Subcategory = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentCreate) GetTenant() string {
	if o == nil || o.Tenant.Get() == nil {
		var ret string
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentCreate) GetTenantOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *IncidentCreate) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableString and assigns it to the Tenant field.
func (o *IncidentCreate) SetTenant(v string) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *IncidentCreate) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *IncidentCreate) UnsetTenant() {
	o.Tenant.Unset()
}

// GetUrgency returns the Urgency field value if set, zero value otherwise.
func (o *IncidentCreate) GetUrgency() string {
	if o == nil || o.Urgency == nil {
		var ret string
		return ret
	}
	return *o.Urgency
}

// GetUrgencyOk returns a tuple with the Urgency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentCreate) GetUrgencyOk() (*string, bool) {
	if o == nil || o.Urgency == nil {
		return nil, false
	}
	return o.Urgency, true
}

// HasUrgency returns a boolean if a field has been set.
func (o *IncidentCreate) HasUrgency() bool {
	if o != nil && o.Urgency != nil {
		return true
	}

	return false
}

// SetUrgency gets a reference to the given string and assigns it to the Urgency field.
func (o *IncidentCreate) SetUrgency(v string) {
	o.Urgency = &v
}

func (o IncidentCreate) MarshalJSON() ([]byte, error) {
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

func (o *IncidentCreate) UnmarshalJSON(bytes []byte) (err error) {
	varIncidentCreate := _IncidentCreate{}

	if err = json.Unmarshal(bytes, &varIncidentCreate); err == nil {
		*o = IncidentCreate(varIncidentCreate)
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

type NullableIncidentCreate struct {
	value *IncidentCreate
	isSet bool
}

func (v NullableIncidentCreate) Get() *IncidentCreate {
	return v.value
}

func (v *NullableIncidentCreate) Set(val *IncidentCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentCreate(val *IncidentCreate) *NullableIncidentCreate {
	return &NullableIncidentCreate{value: val, isSet: true}
}

func (v NullableIncidentCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


