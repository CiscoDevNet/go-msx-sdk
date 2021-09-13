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

// DeviceTemplateCreate struct for DeviceTemplateCreate
type DeviceTemplateCreate struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Version *string `json:"version,omitempty"`
	ServiceType string `json:"serviceType"`
	DeviceModels []string `json:"deviceModels,omitempty"`
	ConfigContent string `json:"configContent"`
	ResourceProvider string `json:"resourceProvider"`
	TemplateStandard *string `json:"templateStandard,omitempty"`
	TenantAccess NullableDeviceTemplateAccess `json:"tenantAccess,omitempty"`
	TemplateParameterValidators *[]TemplateParameterValidator `json:"templateParameterValidators,omitempty"`
	Tags []string `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceTemplateCreate DeviceTemplateCreate

// NewDeviceTemplateCreate instantiates a new DeviceTemplateCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceTemplateCreate(name string, serviceType string, configContent string, resourceProvider string) *DeviceTemplateCreate {
	this := DeviceTemplateCreate{}
	this.Name = name
	this.ServiceType = serviceType
	this.ConfigContent = configContent
	this.ResourceProvider = resourceProvider
	return &this
}

// NewDeviceTemplateCreateWithDefaults instantiates a new DeviceTemplateCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceTemplateCreateWithDefaults() *DeviceTemplateCreate {
	this := DeviceTemplateCreate{}
	return &this
}

// GetName returns the Name field value
func (o *DeviceTemplateCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeviceTemplateCreate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeviceTemplateCreate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DeviceTemplateCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTemplateCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DeviceTemplateCreate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DeviceTemplateCreate) SetDescription(v string) {
	o.Description = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DeviceTemplateCreate) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTemplateCreate) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DeviceTemplateCreate) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *DeviceTemplateCreate) SetVersion(v string) {
	o.Version = &v
}

// GetServiceType returns the ServiceType field value
func (o *DeviceTemplateCreate) GetServiceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value
// and a boolean to check if the value has been set.
func (o *DeviceTemplateCreate) GetServiceTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ServiceType, true
}

// SetServiceType sets field value
func (o *DeviceTemplateCreate) SetServiceType(v string) {
	o.ServiceType = v
}

// GetDeviceModels returns the DeviceModels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceTemplateCreate) GetDeviceModels() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.DeviceModels
}

// GetDeviceModelsOk returns a tuple with the DeviceModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceTemplateCreate) GetDeviceModelsOk() (*[]string, bool) {
	if o == nil || o.DeviceModels == nil {
		return nil, false
	}
	return &o.DeviceModels, true
}

// HasDeviceModels returns a boolean if a field has been set.
func (o *DeviceTemplateCreate) HasDeviceModels() bool {
	if o != nil && o.DeviceModels != nil {
		return true
	}

	return false
}

// SetDeviceModels gets a reference to the given []string and assigns it to the DeviceModels field.
func (o *DeviceTemplateCreate) SetDeviceModels(v []string) {
	o.DeviceModels = v
}

// GetConfigContent returns the ConfigContent field value
func (o *DeviceTemplateCreate) GetConfigContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigContent
}

// GetConfigContentOk returns a tuple with the ConfigContent field value
// and a boolean to check if the value has been set.
func (o *DeviceTemplateCreate) GetConfigContentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConfigContent, true
}

// SetConfigContent sets field value
func (o *DeviceTemplateCreate) SetConfigContent(v string) {
	o.ConfigContent = v
}

// GetResourceProvider returns the ResourceProvider field value
func (o *DeviceTemplateCreate) GetResourceProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceProvider
}

// GetResourceProviderOk returns a tuple with the ResourceProvider field value
// and a boolean to check if the value has been set.
func (o *DeviceTemplateCreate) GetResourceProviderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ResourceProvider, true
}

// SetResourceProvider sets field value
func (o *DeviceTemplateCreate) SetResourceProvider(v string) {
	o.ResourceProvider = v
}

// GetTemplateStandard returns the TemplateStandard field value if set, zero value otherwise.
func (o *DeviceTemplateCreate) GetTemplateStandard() string {
	if o == nil || o.TemplateStandard == nil {
		var ret string
		return ret
	}
	return *o.TemplateStandard
}

// GetTemplateStandardOk returns a tuple with the TemplateStandard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTemplateCreate) GetTemplateStandardOk() (*string, bool) {
	if o == nil || o.TemplateStandard == nil {
		return nil, false
	}
	return o.TemplateStandard, true
}

// HasTemplateStandard returns a boolean if a field has been set.
func (o *DeviceTemplateCreate) HasTemplateStandard() bool {
	if o != nil && o.TemplateStandard != nil {
		return true
	}

	return false
}

// SetTemplateStandard gets a reference to the given string and assigns it to the TemplateStandard field.
func (o *DeviceTemplateCreate) SetTemplateStandard(v string) {
	o.TemplateStandard = &v
}

// GetTenantAccess returns the TenantAccess field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceTemplateCreate) GetTenantAccess() DeviceTemplateAccess {
	if o == nil || o.TenantAccess.Get() == nil {
		var ret DeviceTemplateAccess
		return ret
	}
	return *o.TenantAccess.Get()
}

// GetTenantAccessOk returns a tuple with the TenantAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceTemplateCreate) GetTenantAccessOk() (*DeviceTemplateAccess, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TenantAccess.Get(), o.TenantAccess.IsSet()
}

// HasTenantAccess returns a boolean if a field has been set.
func (o *DeviceTemplateCreate) HasTenantAccess() bool {
	if o != nil && o.TenantAccess.IsSet() {
		return true
	}

	return false
}

// SetTenantAccess gets a reference to the given NullableDeviceTemplateAccess and assigns it to the TenantAccess field.
func (o *DeviceTemplateCreate) SetTenantAccess(v DeviceTemplateAccess) {
	o.TenantAccess.Set(&v)
}
// SetTenantAccessNil sets the value for TenantAccess to be an explicit nil
func (o *DeviceTemplateCreate) SetTenantAccessNil() {
	o.TenantAccess.Set(nil)
}

// UnsetTenantAccess ensures that no value is present for TenantAccess, not even an explicit nil
func (o *DeviceTemplateCreate) UnsetTenantAccess() {
	o.TenantAccess.Unset()
}

// GetTemplateParameterValidators returns the TemplateParameterValidators field value if set, zero value otherwise.
func (o *DeviceTemplateCreate) GetTemplateParameterValidators() []TemplateParameterValidator {
	if o == nil || o.TemplateParameterValidators == nil {
		var ret []TemplateParameterValidator
		return ret
	}
	return *o.TemplateParameterValidators
}

// GetTemplateParameterValidatorsOk returns a tuple with the TemplateParameterValidators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTemplateCreate) GetTemplateParameterValidatorsOk() (*[]TemplateParameterValidator, bool) {
	if o == nil || o.TemplateParameterValidators == nil {
		return nil, false
	}
	return o.TemplateParameterValidators, true
}

// HasTemplateParameterValidators returns a boolean if a field has been set.
func (o *DeviceTemplateCreate) HasTemplateParameterValidators() bool {
	if o != nil && o.TemplateParameterValidators != nil {
		return true
	}

	return false
}

// SetTemplateParameterValidators gets a reference to the given []TemplateParameterValidator and assigns it to the TemplateParameterValidators field.
func (o *DeviceTemplateCreate) SetTemplateParameterValidators(v []TemplateParameterValidator) {
	o.TemplateParameterValidators = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceTemplateCreate) GetTags() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceTemplateCreate) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DeviceTemplateCreate) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DeviceTemplateCreate) SetTags(v []string) {
	o.Tags = v
}

func (o DeviceTemplateCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["serviceType"] = o.ServiceType
	}
	if o.DeviceModels != nil {
		toSerialize["deviceModels"] = o.DeviceModels
	}
	if true {
		toSerialize["configContent"] = o.ConfigContent
	}
	if true {
		toSerialize["resourceProvider"] = o.ResourceProvider
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

func (o *DeviceTemplateCreate) UnmarshalJSON(bytes []byte) (err error) {
	varDeviceTemplateCreate := _DeviceTemplateCreate{}

	if err = json.Unmarshal(bytes, &varDeviceTemplateCreate); err == nil {
		*o = DeviceTemplateCreate(varDeviceTemplateCreate)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "version")
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

type NullableDeviceTemplateCreate struct {
	value *DeviceTemplateCreate
	isSet bool
}

func (v NullableDeviceTemplateCreate) Get() *DeviceTemplateCreate {
	return v.value
}

func (v *NullableDeviceTemplateCreate) Set(val *DeviceTemplateCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceTemplateCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceTemplateCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceTemplateCreate(val *DeviceTemplateCreate) *NullableDeviceTemplateCreate {
	return &NullableDeviceTemplateCreate{value: val, isSet: true}
}

func (v NullableDeviceTemplateCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceTemplateCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


