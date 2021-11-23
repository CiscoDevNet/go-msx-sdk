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

// DeviceUpdate struct for DeviceUpdate
type DeviceUpdate struct {
	ServiceType NullableString `json:"serviceType,omitempty"`
	Tags map[string]string `json:"tags,omitempty"`
	Managed bool `json:"managed"`
	OnboardType string `json:"onboardType"`
	OnboardInformation map[string]interface{} `json:"onboardInformation,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Name string `json:"name"`
	Model string `json:"model"`
	Type string `json:"type"`
	SubType NullableString `json:"subType,omitempty"`
	SerialKey NullableString `json:"serialKey,omitempty"`
	Version NullableString `json:"version,omitempty"`
	ComplianceState *DeviceComplianceState `json:"complianceState,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceUpdate DeviceUpdate

// NewDeviceUpdate instantiates a new DeviceUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceUpdate(managed bool, onboardType string, name string, model string, type_ string) *DeviceUpdate {
	this := DeviceUpdate{}
	this.Managed = managed
	this.OnboardType = onboardType
	this.Name = name
	this.Model = model
	this.Type = type_
	return &this
}

// NewDeviceUpdateWithDefaults instantiates a new DeviceUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceUpdateWithDefaults() *DeviceUpdate {
	this := DeviceUpdate{}
	var managed bool = false
	this.Managed = managed
	return &this
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceUpdate) GetServiceType() string {
	if o == nil || o.ServiceType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServiceType.Get()
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceUpdate) GetServiceTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServiceType.Get(), o.ServiceType.IsSet()
}

// HasServiceType returns a boolean if a field has been set.
func (o *DeviceUpdate) HasServiceType() bool {
	if o != nil && o.ServiceType.IsSet() {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given NullableString and assigns it to the ServiceType field.
func (o *DeviceUpdate) SetServiceType(v string) {
	o.ServiceType.Set(&v)
}
// SetServiceTypeNil sets the value for ServiceType to be an explicit nil
func (o *DeviceUpdate) SetServiceTypeNil() {
	o.ServiceType.Set(nil)
}

// UnsetServiceType ensures that no value is present for ServiceType, not even an explicit nil
func (o *DeviceUpdate) UnsetServiceType() {
	o.ServiceType.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceUpdate) GetTags() map[string]string {
	if o == nil  {
		var ret map[string]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceUpdate) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DeviceUpdate) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *DeviceUpdate) SetTags(v map[string]string) {
	o.Tags = v
}

// GetManaged returns the Managed field value
func (o *DeviceUpdate) GetManaged() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Managed
}

// GetManagedOk returns a tuple with the Managed field value
// and a boolean to check if the value has been set.
func (o *DeviceUpdate) GetManagedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Managed, true
}

// SetManaged sets field value
func (o *DeviceUpdate) SetManaged(v bool) {
	o.Managed = v
}

// GetOnboardType returns the OnboardType field value
func (o *DeviceUpdate) GetOnboardType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OnboardType
}

// GetOnboardTypeOk returns a tuple with the OnboardType field value
// and a boolean to check if the value has been set.
func (o *DeviceUpdate) GetOnboardTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OnboardType, true
}

// SetOnboardType sets field value
func (o *DeviceUpdate) SetOnboardType(v string) {
	o.OnboardType = v
}

// GetOnboardInformation returns the OnboardInformation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceUpdate) GetOnboardInformation() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.OnboardInformation
}

// GetOnboardInformationOk returns a tuple with the OnboardInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceUpdate) GetOnboardInformationOk() (*map[string]interface{}, bool) {
	if o == nil || o.OnboardInformation == nil {
		return nil, false
	}
	return &o.OnboardInformation, true
}

// HasOnboardInformation returns a boolean if a field has been set.
func (o *DeviceUpdate) HasOnboardInformation() bool {
	if o != nil && o.OnboardInformation != nil {
		return true
	}

	return false
}

// SetOnboardInformation gets a reference to the given map[string]interface{} and assigns it to the OnboardInformation field.
func (o *DeviceUpdate) SetOnboardInformation(v map[string]interface{}) {
	o.OnboardInformation = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceUpdate) GetAttributes() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceUpdate) GetAttributesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *DeviceUpdate) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *DeviceUpdate) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetName returns the Name field value
func (o *DeviceUpdate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeviceUpdate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeviceUpdate) SetName(v string) {
	o.Name = v
}

// GetModel returns the Model field value
func (o *DeviceUpdate) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *DeviceUpdate) GetModelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *DeviceUpdate) SetModel(v string) {
	o.Model = v
}

// GetType returns the Type field value
func (o *DeviceUpdate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DeviceUpdate) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DeviceUpdate) SetType(v string) {
	o.Type = v
}

// GetSubType returns the SubType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceUpdate) GetSubType() string {
	if o == nil || o.SubType.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubType.Get()
}

// GetSubTypeOk returns a tuple with the SubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceUpdate) GetSubTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubType.Get(), o.SubType.IsSet()
}

// HasSubType returns a boolean if a field has been set.
func (o *DeviceUpdate) HasSubType() bool {
	if o != nil && o.SubType.IsSet() {
		return true
	}

	return false
}

// SetSubType gets a reference to the given NullableString and assigns it to the SubType field.
func (o *DeviceUpdate) SetSubType(v string) {
	o.SubType.Set(&v)
}
// SetSubTypeNil sets the value for SubType to be an explicit nil
func (o *DeviceUpdate) SetSubTypeNil() {
	o.SubType.Set(nil)
}

// UnsetSubType ensures that no value is present for SubType, not even an explicit nil
func (o *DeviceUpdate) UnsetSubType() {
	o.SubType.Unset()
}

// GetSerialKey returns the SerialKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceUpdate) GetSerialKey() string {
	if o == nil || o.SerialKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.SerialKey.Get()
}

// GetSerialKeyOk returns a tuple with the SerialKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceUpdate) GetSerialKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SerialKey.Get(), o.SerialKey.IsSet()
}

// HasSerialKey returns a boolean if a field has been set.
func (o *DeviceUpdate) HasSerialKey() bool {
	if o != nil && o.SerialKey.IsSet() {
		return true
	}

	return false
}

// SetSerialKey gets a reference to the given NullableString and assigns it to the SerialKey field.
func (o *DeviceUpdate) SetSerialKey(v string) {
	o.SerialKey.Set(&v)
}
// SetSerialKeyNil sets the value for SerialKey to be an explicit nil
func (o *DeviceUpdate) SetSerialKeyNil() {
	o.SerialKey.Set(nil)
}

// UnsetSerialKey ensures that no value is present for SerialKey, not even an explicit nil
func (o *DeviceUpdate) UnsetSerialKey() {
	o.SerialKey.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceUpdate) GetVersion() string {
	if o == nil || o.Version.Get() == nil {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceUpdate) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *DeviceUpdate) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *DeviceUpdate) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *DeviceUpdate) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *DeviceUpdate) UnsetVersion() {
	o.Version.Unset()
}

// GetComplianceState returns the ComplianceState field value if set, zero value otherwise.
func (o *DeviceUpdate) GetComplianceState() DeviceComplianceState {
	if o == nil || o.ComplianceState == nil {
		var ret DeviceComplianceState
		return ret
	}
	return *o.ComplianceState
}

// GetComplianceStateOk returns a tuple with the ComplianceState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceUpdate) GetComplianceStateOk() (*DeviceComplianceState, bool) {
	if o == nil || o.ComplianceState == nil {
		return nil, false
	}
	return o.ComplianceState, true
}

// HasComplianceState returns a boolean if a field has been set.
func (o *DeviceUpdate) HasComplianceState() bool {
	if o != nil && o.ComplianceState != nil {
		return true
	}

	return false
}

// SetComplianceState gets a reference to the given DeviceComplianceState and assigns it to the ComplianceState field.
func (o *DeviceUpdate) SetComplianceState(v DeviceComplianceState) {
	o.ComplianceState = &v
}

func (o DeviceUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServiceType.IsSet() {
		toSerialize["serviceType"] = o.ServiceType.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["managed"] = o.Managed
	}
	if true {
		toSerialize["onboardType"] = o.OnboardType
	}
	if o.OnboardInformation != nil {
		toSerialize["onboardInformation"] = o.OnboardInformation
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["model"] = o.Model
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.SubType.IsSet() {
		toSerialize["subType"] = o.SubType.Get()
	}
	if o.SerialKey.IsSet() {
		toSerialize["serialKey"] = o.SerialKey.Get()
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if o.ComplianceState != nil {
		toSerialize["complianceState"] = o.ComplianceState
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeviceUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varDeviceUpdate := _DeviceUpdate{}

	if err = json.Unmarshal(bytes, &varDeviceUpdate); err == nil {
		*o = DeviceUpdate(varDeviceUpdate)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "serviceType")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "managed")
		delete(additionalProperties, "onboardType")
		delete(additionalProperties, "onboardInformation")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "name")
		delete(additionalProperties, "model")
		delete(additionalProperties, "type")
		delete(additionalProperties, "subType")
		delete(additionalProperties, "serialKey")
		delete(additionalProperties, "version")
		delete(additionalProperties, "complianceState")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceUpdate struct {
	value *DeviceUpdate
	isSet bool
}

func (v NullableDeviceUpdate) Get() *DeviceUpdate {
	return v.value
}

func (v *NullableDeviceUpdate) Set(val *DeviceUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceUpdate(val *DeviceUpdate) *NullableDeviceUpdate {
	return &NullableDeviceUpdate{value: val, isSet: true}
}

func (v NullableDeviceUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


