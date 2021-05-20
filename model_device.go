/*
 * KAKAPO - MSX SDK
 *
 * MSX Platform SDK
 *
 * API version: 1.0.2
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
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	ModifiedOn *time.Time `json:"modifiedOn,omitempty"`
	ServiceInstanceId *string `json:"serviceInstanceId,omitempty"`
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	TenantId string `json:"tenantId"`
	Name string `json:"name"`
	Model string `json:"model"`
	Type string `json:"type"`
	SubType NullableString `json:"subType,omitempty"`
	ServiceType *string `json:"serviceType,omitempty"`
	Tags map[string]string `json:"tags,omitempty"`
	SerialKey string `json:"serialKey"`
	Version string `json:"version"`
	Managed bool `json:"managed"`
	OnboardType *string `json:"onboardType,omitempty"`
	OnboardInformation *map[string]interface{} `json:"onboardInformation,omitempty"`
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	Billing *DeviceBilling `json:"billing,omitempty"`
}

// NewDevice instantiates a new Device object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevice(tenantId string, name string, model string, type_ string, serialKey string, version string, managed bool) *Device {
	this := Device{}
	this.TenantId = tenantId
	this.Name = name
	this.Model = model
	this.Type = type_
	this.SerialKey = serialKey
	this.Version = version
	this.Managed = managed
	return &this
}

// NewDeviceWithDefaults instantiates a new Device object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceWithDefaults() *Device {
	this := Device{}
	var managed bool = false
	this.Managed = managed
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

// GetModifiedOn returns the ModifiedOn field value if set, zero value otherwise.
func (o *Device) GetModifiedOn() time.Time {
	if o == nil || o.ModifiedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetModifiedOnOk() (*time.Time, bool) {
	if o == nil || o.ModifiedOn == nil {
		return nil, false
	}
	return o.ModifiedOn, true
}

// HasModifiedOn returns a boolean if a field has been set.
func (o *Device) HasModifiedOn() bool {
	if o != nil && o.ModifiedOn != nil {
		return true
	}

	return false
}

// SetModifiedOn gets a reference to the given time.Time and assigns it to the ModifiedOn field.
func (o *Device) SetModifiedOn(v time.Time) {
	o.ModifiedOn = &v
}

// GetServiceInstanceId returns the ServiceInstanceId field value if set, zero value otherwise.
func (o *Device) GetServiceInstanceId() string {
	if o == nil || o.ServiceInstanceId == nil {
		var ret string
		return ret
	}
	return *o.ServiceInstanceId
}

// GetServiceInstanceIdOk returns a tuple with the ServiceInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetServiceInstanceIdOk() (*string, bool) {
	if o == nil || o.ServiceInstanceId == nil {
		return nil, false
	}
	return o.ServiceInstanceId, true
}

// HasServiceInstanceId returns a boolean if a field has been set.
func (o *Device) HasServiceInstanceId() bool {
	if o != nil && o.ServiceInstanceId != nil {
		return true
	}

	return false
}

// SetServiceInstanceId gets a reference to the given string and assigns it to the ServiceInstanceId field.
func (o *Device) SetServiceInstanceId(v string) {
	o.ServiceInstanceId = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *Device) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || o.SubscriptionId == nil {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *Device) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId != nil {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *Device) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
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

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *Device) GetServiceType() string {
	if o == nil || o.ServiceType == nil {
		var ret string
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetServiceTypeOk() (*string, bool) {
	if o == nil || o.ServiceType == nil {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *Device) HasServiceType() bool {
	if o != nil && o.ServiceType != nil {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given string and assigns it to the ServiceType field.
func (o *Device) SetServiceType(v string) {
	o.ServiceType = &v
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

// GetSerialKey returns the SerialKey field value
func (o *Device) GetSerialKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SerialKey
}

// GetSerialKeyOk returns a tuple with the SerialKey field value
// and a boolean to check if the value has been set.
func (o *Device) GetSerialKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SerialKey, true
}

// SetSerialKey sets field value
func (o *Device) SetSerialKey(v string) {
	o.SerialKey = v
}

// GetVersion returns the Version field value
func (o *Device) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *Device) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *Device) SetVersion(v string) {
	o.Version = v
}

// GetManaged returns the Managed field value
func (o *Device) GetManaged() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Managed
}

// GetManagedOk returns a tuple with the Managed field value
// and a boolean to check if the value has been set.
func (o *Device) GetManagedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Managed, true
}

// SetManaged sets field value
func (o *Device) SetManaged(v bool) {
	o.Managed = v
}

// GetOnboardType returns the OnboardType field value if set, zero value otherwise.
func (o *Device) GetOnboardType() string {
	if o == nil || o.OnboardType == nil {
		var ret string
		return ret
	}
	return *o.OnboardType
}

// GetOnboardTypeOk returns a tuple with the OnboardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetOnboardTypeOk() (*string, bool) {
	if o == nil || o.OnboardType == nil {
		return nil, false
	}
	return o.OnboardType, true
}

// HasOnboardType returns a boolean if a field has been set.
func (o *Device) HasOnboardType() bool {
	if o != nil && o.OnboardType != nil {
		return true
	}

	return false
}

// SetOnboardType gets a reference to the given string and assigns it to the OnboardType field.
func (o *Device) SetOnboardType(v string) {
	o.OnboardType = &v
}

// GetOnboardInformation returns the OnboardInformation field value if set, zero value otherwise.
func (o *Device) GetOnboardInformation() map[string]string {
	if o == nil || o.OnboardInformation == nil {
		var ret map[string]string
		return ret
	}
	return *o.OnboardInformation
}

// GetOnboardInformationOk returns a tuple with the OnboardInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetOnboardInformationOk() (*map[string]string, bool) {
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

// SetOnboardInformation gets a reference to the given map[string]string and assigns it to the OnboardInformation field.
func (o *Device) SetOnboardInformation(v map[string]string) {
	o.OnboardInformation = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Device) GetAttributes() map[string]string {
	if o == nil || o.Attributes == nil {
		var ret map[string]string
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetAttributesOk() (*map[string]string, bool) {
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

// SetAttributes gets a reference to the given map[string]string and assigns it to the Attributes field.
func (o *Device) SetAttributes(v map[string]string) {
	o.Attributes = &v
}

// GetBilling returns the Billing field value if set, zero value otherwise.
func (o *Device) GetBilling() DeviceBilling {
	if o == nil || o.Billing == nil {
		var ret DeviceBilling
		return ret
	}
	return *o.Billing
}

// GetBillingOk returns a tuple with the Billing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetBillingOk() (*DeviceBilling, bool) {
	if o == nil || o.Billing == nil {
		return nil, false
	}
	return o.Billing, true
}

// HasBilling returns a boolean if a field has been set.
func (o *Device) HasBilling() bool {
	if o != nil && o.Billing != nil {
		return true
	}

	return false
}

// SetBilling gets a reference to the given DeviceBilling and assigns it to the Billing field.
func (o *Device) SetBilling(v DeviceBilling) {
	o.Billing = &v
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
	if o.CreatedOn != nil {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if o.ModifiedOn != nil {
		toSerialize["modifiedOn"] = o.ModifiedOn
	}
	if o.ServiceInstanceId != nil {
		toSerialize["serviceInstanceId"] = o.ServiceInstanceId
	}
	if o.SubscriptionId != nil {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if true {
		toSerialize["tenantId"] = o.TenantId
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
	if o.ServiceType != nil {
		toSerialize["serviceType"] = o.ServiceType
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["serialKey"] = o.SerialKey
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["managed"] = o.Managed
	}
	if o.OnboardType != nil {
		toSerialize["onboardType"] = o.OnboardType
	}
	if o.OnboardInformation != nil {
		toSerialize["onboardInformation"] = o.OnboardInformation
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Billing != nil {
		toSerialize["billing"] = o.Billing
	}
	return json.Marshal(toSerialize)
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


