/*
MSX SDK

MSX SDK client.

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msxsdk

import (
	"encoding/json"
	"time"
)

// TemplateApplication struct for TemplateApplication
type TemplateApplication struct {
	Id *string `json:"id,omitempty"`
	TemplateId *string `json:"templateId,omitempty"`
	TemplateName *string `json:"templateName,omitempty"`
	TenantId string `json:"tenantId"`
	TargetId string `json:"targetId"`
	TargetType string `json:"targetType"`
	Parameters *map[string]string `json:"parameters,omitempty"`
	Status *TemplateStatus `json:"status,omitempty"`
	StatusDetails NullableString `json:"statusDetails,omitempty"`
	CreatedOn NullableTime `json:"createdOn,omitempty"`
	CreatedBy NullableString `json:"createdBy,omitempty"`
	ModifiedOn NullableTime `json:"modifiedOn,omitempty"`
	ModifiedBy NullableString `json:"modifiedBy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TemplateApplication TemplateApplication

// NewTemplateApplication instantiates a new TemplateApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateApplication(tenantId string, targetId string, targetType string) *TemplateApplication {
	this := TemplateApplication{}
	this.TenantId = tenantId
	this.TargetId = targetId
	this.TargetType = targetType
	return &this
}

// NewTemplateApplicationWithDefaults instantiates a new TemplateApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateApplicationWithDefaults() *TemplateApplication {
	this := TemplateApplication{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TemplateApplication) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateApplication) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TemplateApplication) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TemplateApplication) SetId(v string) {
	o.Id = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *TemplateApplication) GetTemplateId() string {
	if o == nil || o.TemplateId == nil {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateApplication) GetTemplateIdOk() (*string, bool) {
	if o == nil || o.TemplateId == nil {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *TemplateApplication) HasTemplateId() bool {
	if o != nil && o.TemplateId != nil {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *TemplateApplication) SetTemplateId(v string) {
	o.TemplateId = &v
}

// GetTemplateName returns the TemplateName field value if set, zero value otherwise.
func (o *TemplateApplication) GetTemplateName() string {
	if o == nil || o.TemplateName == nil {
		var ret string
		return ret
	}
	return *o.TemplateName
}

// GetTemplateNameOk returns a tuple with the TemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateApplication) GetTemplateNameOk() (*string, bool) {
	if o == nil || o.TemplateName == nil {
		return nil, false
	}
	return o.TemplateName, true
}

// HasTemplateName returns a boolean if a field has been set.
func (o *TemplateApplication) HasTemplateName() bool {
	if o != nil && o.TemplateName != nil {
		return true
	}

	return false
}

// SetTemplateName gets a reference to the given string and assigns it to the TemplateName field.
func (o *TemplateApplication) SetTemplateName(v string) {
	o.TemplateName = &v
}

// GetTenantId returns the TenantId field value
func (o *TemplateApplication) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *TemplateApplication) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *TemplateApplication) SetTenantId(v string) {
	o.TenantId = v
}

// GetTargetId returns the TargetId field value
func (o *TemplateApplication) GetTargetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value
// and a boolean to check if the value has been set.
func (o *TemplateApplication) GetTargetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetId, true
}

// SetTargetId sets field value
func (o *TemplateApplication) SetTargetId(v string) {
	o.TargetId = v
}

// GetTargetType returns the TargetType field value
func (o *TemplateApplication) GetTargetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value
// and a boolean to check if the value has been set.
func (o *TemplateApplication) GetTargetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetType, true
}

// SetTargetType sets field value
func (o *TemplateApplication) SetTargetType(v string) {
	o.TargetType = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *TemplateApplication) GetParameters() map[string]string {
	if o == nil || o.Parameters == nil {
		var ret map[string]string
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateApplication) GetParametersOk() (*map[string]string, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *TemplateApplication) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *TemplateApplication) SetParameters(v map[string]string) {
	o.Parameters = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TemplateApplication) GetStatus() TemplateStatus {
	if o == nil || o.Status == nil {
		var ret TemplateStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateApplication) GetStatusOk() (*TemplateStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TemplateApplication) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given TemplateStatus and assigns it to the Status field.
func (o *TemplateApplication) SetStatus(v TemplateStatus) {
	o.Status = &v
}

// GetStatusDetails returns the StatusDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateApplication) GetStatusDetails() string {
	if o == nil || o.StatusDetails.Get() == nil {
		var ret string
		return ret
	}
	return *o.StatusDetails.Get()
}

// GetStatusDetailsOk returns a tuple with the StatusDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateApplication) GetStatusDetailsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StatusDetails.Get(), o.StatusDetails.IsSet()
}

// HasStatusDetails returns a boolean if a field has been set.
func (o *TemplateApplication) HasStatusDetails() bool {
	if o != nil && o.StatusDetails.IsSet() {
		return true
	}

	return false
}

// SetStatusDetails gets a reference to the given NullableString and assigns it to the StatusDetails field.
func (o *TemplateApplication) SetStatusDetails(v string) {
	o.StatusDetails.Set(&v)
}
// SetStatusDetailsNil sets the value for StatusDetails to be an explicit nil
func (o *TemplateApplication) SetStatusDetailsNil() {
	o.StatusDetails.Set(nil)
}

// UnsetStatusDetails ensures that no value is present for StatusDetails, not even an explicit nil
func (o *TemplateApplication) UnsetStatusDetails() {
	o.StatusDetails.Unset()
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateApplication) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn.Get()
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateApplication) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedOn.Get(), o.CreatedOn.IsSet()
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *TemplateApplication) HasCreatedOn() bool {
	if o != nil && o.CreatedOn.IsSet() {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given NullableTime and assigns it to the CreatedOn field.
func (o *TemplateApplication) SetCreatedOn(v time.Time) {
	o.CreatedOn.Set(&v)
}
// SetCreatedOnNil sets the value for CreatedOn to be an explicit nil
func (o *TemplateApplication) SetCreatedOnNil() {
	o.CreatedOn.Set(nil)
}

// UnsetCreatedOn ensures that no value is present for CreatedOn, not even an explicit nil
func (o *TemplateApplication) UnsetCreatedOn() {
	o.CreatedOn.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateApplication) GetCreatedBy() string {
	if o == nil || o.CreatedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateApplication) GetCreatedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *TemplateApplication) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *TemplateApplication) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}
// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *TemplateApplication) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *TemplateApplication) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetModifiedOn returns the ModifiedOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateApplication) GetModifiedOn() time.Time {
	if o == nil || o.ModifiedOn.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedOn.Get()
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateApplication) GetModifiedOnOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ModifiedOn.Get(), o.ModifiedOn.IsSet()
}

// HasModifiedOn returns a boolean if a field has been set.
func (o *TemplateApplication) HasModifiedOn() bool {
	if o != nil && o.ModifiedOn.IsSet() {
		return true
	}

	return false
}

// SetModifiedOn gets a reference to the given NullableTime and assigns it to the ModifiedOn field.
func (o *TemplateApplication) SetModifiedOn(v time.Time) {
	o.ModifiedOn.Set(&v)
}
// SetModifiedOnNil sets the value for ModifiedOn to be an explicit nil
func (o *TemplateApplication) SetModifiedOnNil() {
	o.ModifiedOn.Set(nil)
}

// UnsetModifiedOn ensures that no value is present for ModifiedOn, not even an explicit nil
func (o *TemplateApplication) UnsetModifiedOn() {
	o.ModifiedOn.Unset()
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateApplication) GetModifiedBy() string {
	if o == nil || o.ModifiedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.ModifiedBy.Get()
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateApplication) GetModifiedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ModifiedBy.Get(), o.ModifiedBy.IsSet()
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *TemplateApplication) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy.IsSet() {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given NullableString and assigns it to the ModifiedBy field.
func (o *TemplateApplication) SetModifiedBy(v string) {
	o.ModifiedBy.Set(&v)
}
// SetModifiedByNil sets the value for ModifiedBy to be an explicit nil
func (o *TemplateApplication) SetModifiedByNil() {
	o.ModifiedBy.Set(nil)
}

// UnsetModifiedBy ensures that no value is present for ModifiedBy, not even an explicit nil
func (o *TemplateApplication) UnsetModifiedBy() {
	o.ModifiedBy.Unset()
}

func (o TemplateApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.TemplateId != nil {
		toSerialize["templateId"] = o.TemplateId
	}
	if o.TemplateName != nil {
		toSerialize["templateName"] = o.TemplateName
	}
	if true {
		toSerialize["tenantId"] = o.TenantId
	}
	if true {
		toSerialize["targetId"] = o.TargetId
	}
	if true {
		toSerialize["targetType"] = o.TargetType
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StatusDetails.IsSet() {
		toSerialize["statusDetails"] = o.StatusDetails.Get()
	}
	if o.CreatedOn.IsSet() {
		toSerialize["createdOn"] = o.CreatedOn.Get()
	}
	if o.CreatedBy.IsSet() {
		toSerialize["createdBy"] = o.CreatedBy.Get()
	}
	if o.ModifiedOn.IsSet() {
		toSerialize["modifiedOn"] = o.ModifiedOn.Get()
	}
	if o.ModifiedBy.IsSet() {
		toSerialize["modifiedBy"] = o.ModifiedBy.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TemplateApplication) UnmarshalJSON(bytes []byte) (err error) {
	varTemplateApplication := _TemplateApplication{}

	if err = json.Unmarshal(bytes, &varTemplateApplication); err == nil {
		*o = TemplateApplication(varTemplateApplication)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "templateId")
		delete(additionalProperties, "templateName")
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "targetId")
		delete(additionalProperties, "targetType")
		delete(additionalProperties, "parameters")
		delete(additionalProperties, "status")
		delete(additionalProperties, "statusDetails")
		delete(additionalProperties, "createdOn")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "modifiedOn")
		delete(additionalProperties, "modifiedBy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTemplateApplication struct {
	value *TemplateApplication
	isSet bool
}

func (v NullableTemplateApplication) Get() *TemplateApplication {
	return v.value
}

func (v *NullableTemplateApplication) Set(val *TemplateApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateApplication(val *TemplateApplication) *NullableTemplateApplication {
	return &NullableTemplateApplication{value: val, isSet: true}
}

func (v NullableTemplateApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


