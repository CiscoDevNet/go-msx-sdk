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

// Device struct for Device
type Device struct {
	Id *string `json:"id,omitempty"`
	UserId *string `json:"userId,omitempty"`
	ProviderId *string `json:"providerId,omitempty"`
	// Valid values: VULNERABLE, NOT_VULNERABLE, NOT_APPLICABLE, UNKNOWN
	VulnerabilityState *string `json:"vulnerabilityState,omitempty"`
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	ModifiedOn NullableTime `json:"modifiedOn,omitempty"`
	ServiceInstanceId NullableString `json:"serviceInstanceId,omitempty"`
	SubscriptionId NullableString `json:"subscriptionId,omitempty"`
	TenantId string `json:"tenantId"`
	ServiceType NullableString `json:"serviceType,omitempty"`
	Tags map[string]string `json:"tags,omitempty"`
	Managed *bool `json:"managed,omitempty"`
	OnboardType string `json:"onboardType"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Name string `json:"name"`
	Model string `json:"model"`
	Type string `json:"type"`
	SubType NullableString `json:"subType,omitempty"`
	SerialKey NullableString `json:"serialKey,omitempty"`
	Version NullableString `json:"version,omitempty"`
	// Valid values: COMPLIANT, NOT_COMPLIANT, APPLICABLE, NOT_APPLICABLE, UNKNOWN
	ComplianceState *string `json:"complianceState,omitempty"`
	OnboardInformation map[string]interface{} `json:"onboardInformation,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Device Device

// NewDevice instantiates a new Device object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevice(tenantId string, onboardType string, name string, model string, type_ string) *Device {
	this := Device{}
	this.TenantId = tenantId
	var managed bool = false
	this.Managed = &managed
	this.OnboardType = onboardType
	this.Name = name
	this.Model = model
	this.Type = type_
	return &this
}

// NewDeviceWithDefaults instantiates a new Device object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceWithDefaults() *Device {
	this := Device{}
	var managed bool = false
	this.Managed = &managed
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Device) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Device) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Device) SetId(v string) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *Device) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *Device) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *Device) SetUserId(v string) {
	o.UserId = &v
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *Device) GetProviderId() string {
	if o == nil || o.ProviderId == nil {
		var ret string
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetProviderIdOk() (*string, bool) {
	if o == nil || o.ProviderId == nil {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *Device) HasProviderId() bool {
	if o != nil && o.ProviderId != nil {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given string and assigns it to the ProviderId field.
func (o *Device) SetProviderId(v string) {
	o.ProviderId = &v
}

// GetVulnerabilityState returns the VulnerabilityState field value if set, zero value otherwise.
func (o *Device) GetVulnerabilityState() string {
	if o == nil || o.VulnerabilityState == nil {
		var ret string
		return ret
	}
	return *o.VulnerabilityState
}

// GetVulnerabilityStateOk returns a tuple with the VulnerabilityState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetVulnerabilityStateOk() (*string, bool) {
	if o == nil || o.VulnerabilityState == nil {
		return nil, false
	}
	return o.VulnerabilityState, true
}

// HasVulnerabilityState returns a boolean if a field has been set.
func (o *Device) HasVulnerabilityState() bool {
	if o != nil && o.VulnerabilityState != nil {
		return true
	}

	return false
}

// SetVulnerabilityState gets a reference to the given string and assigns it to the VulnerabilityState field.
func (o *Device) SetVulnerabilityState(v string) {
	o.VulnerabilityState = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *Device) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *Device) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *Device) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetModifiedOn returns the ModifiedOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetModifiedOn() time.Time {
	if o == nil || o.ModifiedOn.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedOn.Get()
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetModifiedOnOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ModifiedOn.Get(), o.ModifiedOn.IsSet()
}

// HasModifiedOn returns a boolean if a field has been set.
func (o *Device) HasModifiedOn() bool {
	if o != nil && o.ModifiedOn.IsSet() {
		return true
	}

	return false
}

// SetModifiedOn gets a reference to the given NullableTime and assigns it to the ModifiedOn field.
func (o *Device) SetModifiedOn(v time.Time) {
	o.ModifiedOn.Set(&v)
}
// SetModifiedOnNil sets the value for ModifiedOn to be an explicit nil
func (o *Device) SetModifiedOnNil() {
	o.ModifiedOn.Set(nil)
}

// UnsetModifiedOn ensures that no value is present for ModifiedOn, not even an explicit nil
func (o *Device) UnsetModifiedOn() {
	o.ModifiedOn.Unset()
}

// GetServiceInstanceId returns the ServiceInstanceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetServiceInstanceId() string {
	if o == nil || o.ServiceInstanceId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServiceInstanceId.Get()
}

// GetServiceInstanceIdOk returns a tuple with the ServiceInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetServiceInstanceIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServiceInstanceId.Get(), o.ServiceInstanceId.IsSet()
}

// HasServiceInstanceId returns a boolean if a field has been set.
func (o *Device) HasServiceInstanceId() bool {
	if o != nil && o.ServiceInstanceId.IsSet() {
		return true
	}

	return false
}

// SetServiceInstanceId gets a reference to the given NullableString and assigns it to the ServiceInstanceId field.
func (o *Device) SetServiceInstanceId(v string) {
	o.ServiceInstanceId.Set(&v)
}
// SetServiceInstanceIdNil sets the value for ServiceInstanceId to be an explicit nil
func (o *Device) SetServiceInstanceIdNil() {
	o.ServiceInstanceId.Set(nil)
}

// UnsetServiceInstanceId ensures that no value is present for ServiceInstanceId, not even an explicit nil
func (o *Device) UnsetServiceInstanceId() {
	o.ServiceInstanceId.Unset()
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId.Get()
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetSubscriptionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubscriptionId.Get(), o.SubscriptionId.IsSet()
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *Device) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given NullableString and assigns it to the SubscriptionId field.
func (o *Device) SetSubscriptionId(v string) {
	o.SubscriptionId.Set(&v)
}
// SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil
func (o *Device) SetSubscriptionIdNil() {
	o.SubscriptionId.Set(nil)
}

// UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
func (o *Device) UnsetSubscriptionId() {
	o.SubscriptionId.Unset()
}

// GetTenantId returns the TenantId field value
func (o *Device) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *Device) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *Device) SetTenantId(v string) {
	o.TenantId = v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetServiceType() string {
	if o == nil || o.ServiceType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServiceType.Get()
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetServiceTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServiceType.Get(), o.ServiceType.IsSet()
}

// HasServiceType returns a boolean if a field has been set.
func (o *Device) HasServiceType() bool {
	if o != nil && o.ServiceType.IsSet() {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given NullableString and assigns it to the ServiceType field.
func (o *Device) SetServiceType(v string) {
	o.ServiceType.Set(&v)
}
// SetServiceTypeNil sets the value for ServiceType to be an explicit nil
func (o *Device) SetServiceTypeNil() {
	o.ServiceType.Set(nil)
}

// UnsetServiceType ensures that no value is present for ServiceType, not even an explicit nil
func (o *Device) UnsetServiceType() {
	o.ServiceType.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetTags() map[string]string {
	if o == nil  {
		var ret map[string]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Device) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *Device) SetTags(v map[string]string) {
	o.Tags = v
}

// GetManaged returns the Managed field value if set, zero value otherwise.
func (o *Device) GetManaged() bool {
	if o == nil || o.Managed == nil {
		var ret bool
		return ret
	}
	return *o.Managed
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetManagedOk() (*bool, bool) {
	if o == nil || o.Managed == nil {
		return nil, false
	}
	return o.Managed, true
}

// HasManaged returns a boolean if a field has been set.
func (o *Device) HasManaged() bool {
	if o != nil && o.Managed != nil {
		return true
	}

	return false
}

// SetManaged gets a reference to the given bool and assigns it to the Managed field.
func (o *Device) SetManaged(v bool) {
	o.Managed = &v
}

// GetOnboardType returns the OnboardType field value
func (o *Device) GetOnboardType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OnboardType
}

// GetOnboardTypeOk returns a tuple with the OnboardType field value
// and a boolean to check if the value has been set.
func (o *Device) GetOnboardTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OnboardType, true
}

// SetOnboardType sets field value
func (o *Device) SetOnboardType(v string) {
	o.OnboardType = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetAttributes() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Device) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *Device) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetName returns the Name field value
func (o *Device) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Device) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Device) SetName(v string) {
	o.Name = v
}

// GetModel returns the Model field value
func (o *Device) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *Device) GetModelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *Device) SetModel(v string) {
	o.Model = v
}

// GetType returns the Type field value
func (o *Device) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Device) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Device) SetType(v string) {
	o.Type = v
}

// GetSubType returns the SubType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetSubType() string {
	if o == nil || o.SubType.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubType.Get()
}

// GetSubTypeOk returns a tuple with the SubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetSubTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubType.Get(), o.SubType.IsSet()
}

// HasSubType returns a boolean if a field has been set.
func (o *Device) HasSubType() bool {
	if o != nil && o.SubType.IsSet() {
		return true
	}

	return false
}

// SetSubType gets a reference to the given NullableString and assigns it to the SubType field.
func (o *Device) SetSubType(v string) {
	o.SubType.Set(&v)
}
// SetSubTypeNil sets the value for SubType to be an explicit nil
func (o *Device) SetSubTypeNil() {
	o.SubType.Set(nil)
}

// UnsetSubType ensures that no value is present for SubType, not even an explicit nil
func (o *Device) UnsetSubType() {
	o.SubType.Unset()
}

// GetSerialKey returns the SerialKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetSerialKey() string {
	if o == nil || o.SerialKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.SerialKey.Get()
}

// GetSerialKeyOk returns a tuple with the SerialKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetSerialKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SerialKey.Get(), o.SerialKey.IsSet()
}

// HasSerialKey returns a boolean if a field has been set.
func (o *Device) HasSerialKey() bool {
	if o != nil && o.SerialKey.IsSet() {
		return true
	}

	return false
}

// SetSerialKey gets a reference to the given NullableString and assigns it to the SerialKey field.
func (o *Device) SetSerialKey(v string) {
	o.SerialKey.Set(&v)
}
// SetSerialKeyNil sets the value for SerialKey to be an explicit nil
func (o *Device) SetSerialKeyNil() {
	o.SerialKey.Set(nil)
}

// UnsetSerialKey ensures that no value is present for SerialKey, not even an explicit nil
func (o *Device) UnsetSerialKey() {
	o.SerialKey.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetVersion() string {
	if o == nil || o.Version.Get() == nil {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *Device) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *Device) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *Device) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *Device) UnsetVersion() {
	o.Version.Unset()
}

// GetComplianceState returns the ComplianceState field value if set, zero value otherwise.
func (o *Device) GetComplianceState() string {
	if o == nil || o.ComplianceState == nil {
		var ret string
		return ret
	}
	return *o.ComplianceState
}

// GetComplianceStateOk returns a tuple with the ComplianceState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetComplianceStateOk() (*string, bool) {
	if o == nil || o.ComplianceState == nil {
		return nil, false
	}
	return o.ComplianceState, true
}

// HasComplianceState returns a boolean if a field has been set.
func (o *Device) HasComplianceState() bool {
	if o != nil && o.ComplianceState != nil {
		return true
	}

	return false
}

// SetComplianceState gets a reference to the given string and assigns it to the ComplianceState field.
func (o *Device) SetComplianceState(v string) {
	o.ComplianceState = &v
}

// GetOnboardInformation returns the OnboardInformation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetOnboardInformation() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.OnboardInformation
}

// GetOnboardInformationOk returns a tuple with the OnboardInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetOnboardInformationOk() (map[string]interface{}, bool) {
	if o == nil || o.OnboardInformation == nil {
		return nil, false
	}
	return o.OnboardInformation, true
}

// HasOnboardInformation returns a boolean if a field has been set.
func (o *Device) HasOnboardInformation() bool {
	if o != nil && o.OnboardInformation != nil {
		return true
	}

	return false
}

// SetOnboardInformation gets a reference to the given map[string]interface{} and assigns it to the OnboardInformation field.
func (o *Device) SetOnboardInformation(v map[string]interface{}) {
	o.OnboardInformation = v
}

func (o Device) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.ProviderId != nil {
		toSerialize["providerId"] = o.ProviderId
	}
	if o.VulnerabilityState != nil {
		toSerialize["vulnerabilityState"] = o.VulnerabilityState
	}
	if o.CreatedOn != nil {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if o.ModifiedOn.IsSet() {
		toSerialize["modifiedOn"] = o.ModifiedOn.Get()
	}
	if o.ServiceInstanceId.IsSet() {
		toSerialize["serviceInstanceId"] = o.ServiceInstanceId.Get()
	}
	if o.SubscriptionId.IsSet() {
		toSerialize["subscriptionId"] = o.SubscriptionId.Get()
	}
	if true {
		toSerialize["tenantId"] = o.TenantId
	}
	if o.ServiceType.IsSet() {
		toSerialize["serviceType"] = o.ServiceType.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Managed != nil {
		toSerialize["managed"] = o.Managed
	}
	if true {
		toSerialize["onboardType"] = o.OnboardType
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
	if o.OnboardInformation != nil {
		toSerialize["onboardInformation"] = o.OnboardInformation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Device) UnmarshalJSON(bytes []byte) (err error) {
	varDevice := _Device{}

	if err = json.Unmarshal(bytes, &varDevice); err == nil {
		*o = Device(varDevice)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "userId")
		delete(additionalProperties, "providerId")
		delete(additionalProperties, "vulnerabilityState")
		delete(additionalProperties, "createdOn")
		delete(additionalProperties, "modifiedOn")
		delete(additionalProperties, "serviceInstanceId")
		delete(additionalProperties, "subscriptionId")
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "serviceType")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "managed")
		delete(additionalProperties, "onboardType")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "name")
		delete(additionalProperties, "model")
		delete(additionalProperties, "type")
		delete(additionalProperties, "subType")
		delete(additionalProperties, "serialKey")
		delete(additionalProperties, "version")
		delete(additionalProperties, "complianceState")
		delete(additionalProperties, "onboardInformation")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDevice struct {
	value *Device
	isSet bool
}

func (v NullableDevice) Get() *Device {
	return v.value
}

func (v *NullableDevice) Set(val *Device) {
	v.value = val
	v.isSet = true
}

func (v NullableDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevice(val *Device) *NullableDevice {
	return &NullableDevice{value: val, isSet: true}
}

func (v NullableDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


