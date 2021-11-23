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
	"time"
)

// DeviceTemplate struct for DeviceTemplate
type DeviceTemplate struct {
	Id *string `json:"id,omitempty"`
	UserId *string `json:"userId,omitempty"`
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Version *string `json:"version,omitempty"`
	IsLatestVersion NullableBool `json:"isLatestVersion,omitempty"`
	ServiceType *string `json:"serviceType,omitempty"`
	DeviceModels []string `json:"deviceModels,omitempty"`
	ConfigContent NullableString `json:"configContent,omitempty"`
	ResourceProvider NullableString `json:"resourceProvider,omitempty"`
	TemplateStandard *string `json:"templateStandard,omitempty"`
	TenantAccess NullableDeviceTemplateAccess `json:"tenantAccess,omitempty"`
	TemplateParameterValidators []TemplateParameterValidator `json:"templateParameterValidators,omitempty"`
	Tags []string `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceTemplate DeviceTemplate

// NewDeviceTemplate instantiates a new DeviceTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceTemplate() *DeviceTemplate {
	this := DeviceTemplate{}
	return &this
}

// NewDeviceTemplateWithDefaults instantiates a new DeviceTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceTemplateWithDefaults() *DeviceTemplate {
	this := DeviceTemplate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeviceTemplate) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTemplate) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeviceTemplate) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeviceTemplate) SetId(v string) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *DeviceTemplate) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTemplate) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *DeviceTemplate) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *DeviceTemplate) SetUserId(v string) {
	o.UserId = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *DeviceTemplate) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTemplate) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *DeviceTemplate) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *DeviceTemplate) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeviceTemplate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTemplate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeviceTemplate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeviceTemplate) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DeviceTemplate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DeviceTemplate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DeviceTemplate) SetDescription(v string) {
	o.Description = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DeviceTemplate) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTemplate) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DeviceTemplate) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *DeviceTemplate) SetVersion(v string) {
	o.Version = &v
}

// GetIsLatestVersion returns the IsLatestVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceTemplate) GetIsLatestVersion() bool {
	if o == nil || o.IsLatestVersion.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsLatestVersion.Get()
}

// GetIsLatestVersionOk returns a tuple with the IsLatestVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceTemplate) GetIsLatestVersionOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsLatestVersion.Get(), o.IsLatestVersion.IsSet()
}

// HasIsLatestVersion returns a boolean if a field has been set.
func (o *DeviceTemplate) HasIsLatestVersion() bool {
	if o != nil && o.IsLatestVersion.IsSet() {
		return true
	}

	return false
}

// SetIsLatestVersion gets a reference to the given NullableBool and assigns it to the IsLatestVersion field.
func (o *DeviceTemplate) SetIsLatestVersion(v bool) {
	o.IsLatestVersion.Set(&v)
}
// SetIsLatestVersionNil sets the value for IsLatestVersion to be an explicit nil
func (o *DeviceTemplate) SetIsLatestVersionNil() {
	o.IsLatestVersion.Set(nil)
}

// UnsetIsLatestVersion ensures that no value is present for IsLatestVersion, not even an explicit nil
func (o *DeviceTemplate) UnsetIsLatestVersion() {
	o.IsLatestVersion.Unset()
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *DeviceTemplate) GetServiceType() string {
	if o == nil || o.ServiceType == nil {
		var ret string
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTemplate) GetServiceTypeOk() (*string, bool) {
	if o == nil || o.ServiceType == nil {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *DeviceTemplate) HasServiceType() bool {
	if o != nil && o.ServiceType != nil {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given string and assigns it to the ServiceType field.
func (o *DeviceTemplate) SetServiceType(v string) {
	o.ServiceType = &v
}

// GetDeviceModels returns the DeviceModels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceTemplate) GetDeviceModels() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.DeviceModels
}

// GetDeviceModelsOk returns a tuple with the DeviceModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceTemplate) GetDeviceModelsOk() (*[]string, bool) {
	if o == nil || o.DeviceModels == nil {
		return nil, false
	}
	return &o.DeviceModels, true
}

// HasDeviceModels returns a boolean if a field has been set.
func (o *DeviceTemplate) HasDeviceModels() bool {
	if o != nil && o.DeviceModels != nil {
		return true
	}

	return false
}

// SetDeviceModels gets a reference to the given []string and assigns it to the DeviceModels field.
func (o *DeviceTemplate) SetDeviceModels(v []string) {
	o.DeviceModels = v
}

// GetConfigContent returns the ConfigContent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceTemplate) GetConfigContent() string {
	if o == nil || o.ConfigContent.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfigContent.Get()
}

// GetConfigContentOk returns a tuple with the ConfigContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceTemplate) GetConfigContentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConfigContent.Get(), o.ConfigContent.IsSet()
}

// HasConfigContent returns a boolean if a field has been set.
func (o *DeviceTemplate) HasConfigContent() bool {
	if o != nil && o.ConfigContent.IsSet() {
		return true
	}

	return false
}

// SetConfigContent gets a reference to the given NullableString and assigns it to the ConfigContent field.
func (o *DeviceTemplate) SetConfigContent(v string) {
	o.ConfigContent.Set(&v)
}
// SetConfigContentNil sets the value for ConfigContent to be an explicit nil
func (o *DeviceTemplate) SetConfigContentNil() {
	o.ConfigContent.Set(nil)
}

// UnsetConfigContent ensures that no value is present for ConfigContent, not even an explicit nil
func (o *DeviceTemplate) UnsetConfigContent() {
	o.ConfigContent.Unset()
}

// GetResourceProvider returns the ResourceProvider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceTemplate) GetResourceProvider() string {
	if o == nil || o.ResourceProvider.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResourceProvider.Get()
}

// GetResourceProviderOk returns a tuple with the ResourceProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceTemplate) GetResourceProviderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResourceProvider.Get(), o.ResourceProvider.IsSet()
}

// HasResourceProvider returns a boolean if a field has been set.
func (o *DeviceTemplate) HasResourceProvider() bool {
	if o != nil && o.ResourceProvider.IsSet() {
		return true
	}

	return false
}

// SetResourceProvider gets a reference to the given NullableString and assigns it to the ResourceProvider field.
func (o *DeviceTemplate) SetResourceProvider(v string) {
	o.ResourceProvider.Set(&v)
}
// SetResourceProviderNil sets the value for ResourceProvider to be an explicit nil
func (o *DeviceTemplate) SetResourceProviderNil() {
	o.ResourceProvider.Set(nil)
}

// UnsetResourceProvider ensures that no value is present for ResourceProvider, not even an explicit nil
func (o *DeviceTemplate) UnsetResourceProvider() {
	o.ResourceProvider.Unset()
}

// GetTemplateStandard returns the TemplateStandard field value if set, zero value otherwise.
func (o *DeviceTemplate) GetTemplateStandard() string {
	if o == nil || o.TemplateStandard == nil {
		var ret string
		return ret
	}
	return *o.TemplateStandard
}

// GetTemplateStandardOk returns a tuple with the TemplateStandard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTemplate) GetTemplateStandardOk() (*string, bool) {
	if o == nil || o.TemplateStandard == nil {
		return nil, false
	}
	return o.TemplateStandard, true
}

// HasTemplateStandard returns a boolean if a field has been set.
func (o *DeviceTemplate) HasTemplateStandard() bool {
	if o != nil && o.TemplateStandard != nil {
		return true
	}

	return false
}

// SetTemplateStandard gets a reference to the given string and assigns it to the TemplateStandard field.
func (o *DeviceTemplate) SetTemplateStandard(v string) {
	o.TemplateStandard = &v
}

// GetTenantAccess returns the TenantAccess field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceTemplate) GetTenantAccess() DeviceTemplateAccess {
	if o == nil || o.TenantAccess.Get() == nil {
		var ret DeviceTemplateAccess
		return ret
	}
	return *o.TenantAccess.Get()
}

// GetTenantAccessOk returns a tuple with the TenantAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceTemplate) GetTenantAccessOk() (*DeviceTemplateAccess, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TenantAccess.Get(), o.TenantAccess.IsSet()
}

// HasTenantAccess returns a boolean if a field has been set.
func (o *DeviceTemplate) HasTenantAccess() bool {
	if o != nil && o.TenantAccess.IsSet() {
		return true
	}

	return false
}

// SetTenantAccess gets a reference to the given NullableDeviceTemplateAccess and assigns it to the TenantAccess field.
func (o *DeviceTemplate) SetTenantAccess(v DeviceTemplateAccess) {
	o.TenantAccess.Set(&v)
}
// SetTenantAccessNil sets the value for TenantAccess to be an explicit nil
func (o *DeviceTemplate) SetTenantAccessNil() {
	o.TenantAccess.Set(nil)
}

// UnsetTenantAccess ensures that no value is present for TenantAccess, not even an explicit nil
func (o *DeviceTemplate) UnsetTenantAccess() {
	o.TenantAccess.Unset()
}

// GetTemplateParameterValidators returns the TemplateParameterValidators field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceTemplate) GetTemplateParameterValidators() []TemplateParameterValidator {
	if o == nil  {
		var ret []TemplateParameterValidator
		return ret
	}
	return o.TemplateParameterValidators
}

// GetTemplateParameterValidatorsOk returns a tuple with the TemplateParameterValidators field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceTemplate) GetTemplateParameterValidatorsOk() (*[]TemplateParameterValidator, bool) {
	if o == nil || o.TemplateParameterValidators == nil {
		return nil, false
	}
	return &o.TemplateParameterValidators, true
}

// HasTemplateParameterValidators returns a boolean if a field has been set.
func (o *DeviceTemplate) HasTemplateParameterValidators() bool {
	if o != nil && o.TemplateParameterValidators != nil {
		return true
	}

	return false
}

// SetTemplateParameterValidators gets a reference to the given []TemplateParameterValidator and assigns it to the TemplateParameterValidators field.
func (o *DeviceTemplate) SetTemplateParameterValidators(v []TemplateParameterValidator) {
	o.TemplateParameterValidators = v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceTemplate) GetTags() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceTemplate) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DeviceTemplate) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DeviceTemplate) SetTags(v []string) {
	o.Tags = v
}

func (o DeviceTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.CreatedOn != nil {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.IsLatestVersion.IsSet() {
		toSerialize["isLatestVersion"] = o.IsLatestVersion.Get()
	}
	if o.ServiceType != nil {
		toSerialize["serviceType"] = o.ServiceType
	}
	if o.DeviceModels != nil {
		toSerialize["deviceModels"] = o.DeviceModels
	}
	if o.ConfigContent.IsSet() {
		toSerialize["configContent"] = o.ConfigContent.Get()
	}
	if o.ResourceProvider.IsSet() {
		toSerialize["resourceProvider"] = o.ResourceProvider.Get()
	}
	if o.TemplateStandard != nil {
		toSerialize["templateStandard"] = o.TemplateStandard
	}
	if o.TenantAccess.IsSet() {
		toSerialize["tenantAccess"] = o.TenantAccess.Get()
	}
	if o.TemplateParameterValidators != nil {
		toSerialize["templateParameterValidators"] = o.TemplateParameterValidators
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeviceTemplate) UnmarshalJSON(bytes []byte) (err error) {
	varDeviceTemplate := _DeviceTemplate{}

	if err = json.Unmarshal(bytes, &varDeviceTemplate); err == nil {
		*o = DeviceTemplate(varDeviceTemplate)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "userId")
		delete(additionalProperties, "createdOn")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "version")
		delete(additionalProperties, "isLatestVersion")
		delete(additionalProperties, "serviceType")
		delete(additionalProperties, "deviceModels")
		delete(additionalProperties, "configContent")
		delete(additionalProperties, "resourceProvider")
		delete(additionalProperties, "templateStandard")
		delete(additionalProperties, "tenantAccess")
		delete(additionalProperties, "templateParameterValidators")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceTemplate struct {
	value *DeviceTemplate
	isSet bool
}

func (v NullableDeviceTemplate) Get() *DeviceTemplate {
	return v.value
}

func (v *NullableDeviceTemplate) Set(val *DeviceTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceTemplate(val *DeviceTemplate) *NullableDeviceTemplate {
	return &NullableDeviceTemplate{value: val, isSet: true}
}

func (v NullableDeviceTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


