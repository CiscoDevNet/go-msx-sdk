/*
 * KAKAPO - MSX SDK
 *
 * MSX Platform SDK
 *
 * API version: 3.10.0
 * Contact: somecontact@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msxsdk

import (
	"encoding/json"
)

// LegacySiteDeviceOnboard struct for LegacySiteDeviceOnboard
type LegacySiteDeviceOnboard struct {
	DeviceInstanceId *string `json:"deviceInstanceId,omitempty"`
	TenantId *string `json:"tenantId,omitempty"`
	DeviceName *string `json:"deviceName,omitempty"`
	Managed *bool `json:"managed,omitempty"`
	DeviceModel *string `json:"deviceModel,omitempty"`
	DeviceOnboardingType *string `json:"deviceOnboardingType,omitempty"`
	DeviceOnboardInformation *map[string]map[string]interface{} `json:"deviceOnboardInformation,omitempty"`
}

// NewLegacySiteDeviceOnboard instantiates a new LegacySiteDeviceOnboard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLegacySiteDeviceOnboard() *LegacySiteDeviceOnboard {
	this := LegacySiteDeviceOnboard{}
	var managed bool = true
	this.Managed = &managed
	return &this
}

// NewLegacySiteDeviceOnboardWithDefaults instantiates a new LegacySiteDeviceOnboard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLegacySiteDeviceOnboardWithDefaults() *LegacySiteDeviceOnboard {
	this := LegacySiteDeviceOnboard{}
	var managed bool = true
	this.Managed = &managed
	return &this
}

// GetDeviceInstanceId returns the DeviceInstanceId field value if set, zero value otherwise.
func (o *LegacySiteDeviceOnboard) GetDeviceInstanceId() string {
	if o == nil || o.DeviceInstanceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceInstanceId
}

// GetDeviceInstanceIdOk returns a tuple with the DeviceInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacySiteDeviceOnboard) GetDeviceInstanceIdOk() (*string, bool) {
	if o == nil || o.DeviceInstanceId == nil {
		return nil, false
	}
	return o.DeviceInstanceId, true
}

// HasDeviceInstanceId returns a boolean if a field has been set.
func (o *LegacySiteDeviceOnboard) HasDeviceInstanceId() bool {
	if o != nil && o.DeviceInstanceId != nil {
		return true
	}

	return false
}

// SetDeviceInstanceId gets a reference to the given string and assigns it to the DeviceInstanceId field.
func (o *LegacySiteDeviceOnboard) SetDeviceInstanceId(v string) {
	o.DeviceInstanceId = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *LegacySiteDeviceOnboard) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacySiteDeviceOnboard) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *LegacySiteDeviceOnboard) HasTenantId() bool {
	if o != nil && o.TenantId != nil {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *LegacySiteDeviceOnboard) SetTenantId(v string) {
	o.TenantId = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *LegacySiteDeviceOnboard) GetDeviceName() string {
	if o == nil || o.DeviceName == nil {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacySiteDeviceOnboard) GetDeviceNameOk() (*string, bool) {
	if o == nil || o.DeviceName == nil {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *LegacySiteDeviceOnboard) HasDeviceName() bool {
	if o != nil && o.DeviceName != nil {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *LegacySiteDeviceOnboard) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetManaged returns the Managed field value if set, zero value otherwise.
func (o *LegacySiteDeviceOnboard) GetManaged() bool {
	if o == nil || o.Managed == nil {
		var ret bool
		return ret
	}
	return *o.Managed
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacySiteDeviceOnboard) GetManagedOk() (*bool, bool) {
	if o == nil || o.Managed == nil {
		return nil, false
	}
	return o.Managed, true
}

// HasManaged returns a boolean if a field has been set.
func (o *LegacySiteDeviceOnboard) HasManaged() bool {
	if o != nil && o.Managed != nil {
		return true
	}

	return false
}

// SetManaged gets a reference to the given bool and assigns it to the Managed field.
func (o *LegacySiteDeviceOnboard) SetManaged(v bool) {
	o.Managed = &v
}

// GetDeviceModel returns the DeviceModel field value if set, zero value otherwise.
func (o *LegacySiteDeviceOnboard) GetDeviceModel() string {
	if o == nil || o.DeviceModel == nil {
		var ret string
		return ret
	}
	return *o.DeviceModel
}

// GetDeviceModelOk returns a tuple with the DeviceModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacySiteDeviceOnboard) GetDeviceModelOk() (*string, bool) {
	if o == nil || o.DeviceModel == nil {
		return nil, false
	}
	return o.DeviceModel, true
}

// HasDeviceModel returns a boolean if a field has been set.
func (o *LegacySiteDeviceOnboard) HasDeviceModel() bool {
	if o != nil && o.DeviceModel != nil {
		return true
	}

	return false
}

// SetDeviceModel gets a reference to the given string and assigns it to the DeviceModel field.
func (o *LegacySiteDeviceOnboard) SetDeviceModel(v string) {
	o.DeviceModel = &v
}

// GetDeviceOnboardingType returns the DeviceOnboardingType field value if set, zero value otherwise.
func (o *LegacySiteDeviceOnboard) GetDeviceOnboardingType() string {
	if o == nil || o.DeviceOnboardingType == nil {
		var ret string
		return ret
	}
	return *o.DeviceOnboardingType
}

// GetDeviceOnboardingTypeOk returns a tuple with the DeviceOnboardingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacySiteDeviceOnboard) GetDeviceOnboardingTypeOk() (*string, bool) {
	if o == nil || o.DeviceOnboardingType == nil {
		return nil, false
	}
	return o.DeviceOnboardingType, true
}

// HasDeviceOnboardingType returns a boolean if a field has been set.
func (o *LegacySiteDeviceOnboard) HasDeviceOnboardingType() bool {
	if o != nil && o.DeviceOnboardingType != nil {
		return true
	}

	return false
}

// SetDeviceOnboardingType gets a reference to the given string and assigns it to the DeviceOnboardingType field.
func (o *LegacySiteDeviceOnboard) SetDeviceOnboardingType(v string) {
	o.DeviceOnboardingType = &v
}

// GetDeviceOnboardInformation returns the DeviceOnboardInformation field value if set, zero value otherwise.
func (o *LegacySiteDeviceOnboard) GetDeviceOnboardInformation() map[string]map[string]interface{} {
	if o == nil || o.DeviceOnboardInformation == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.DeviceOnboardInformation
}

// GetDeviceOnboardInformationOk returns a tuple with the DeviceOnboardInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacySiteDeviceOnboard) GetDeviceOnboardInformationOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.DeviceOnboardInformation == nil {
		return nil, false
	}
	return o.DeviceOnboardInformation, true
}

// HasDeviceOnboardInformation returns a boolean if a field has been set.
func (o *LegacySiteDeviceOnboard) HasDeviceOnboardInformation() bool {
	if o != nil && o.DeviceOnboardInformation != nil {
		return true
	}

	return false
}

// SetDeviceOnboardInformation gets a reference to the given map[string]map[string]interface{} and assigns it to the DeviceOnboardInformation field.
func (o *LegacySiteDeviceOnboard) SetDeviceOnboardInformation(v map[string]map[string]interface{}) {
	o.DeviceOnboardInformation = &v
}

func (o LegacySiteDeviceOnboard) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceInstanceId != nil {
		toSerialize["deviceInstanceId"] = o.DeviceInstanceId
	}
	if o.TenantId != nil {
		toSerialize["tenantId"] = o.TenantId
	}
	if o.DeviceName != nil {
		toSerialize["deviceName"] = o.DeviceName
	}
	if o.Managed != nil {
		toSerialize["managed"] = o.Managed
	}
	if o.DeviceModel != nil {
		toSerialize["deviceModel"] = o.DeviceModel
	}
	if o.DeviceOnboardingType != nil {
		toSerialize["deviceOnboardingType"] = o.DeviceOnboardingType
	}
	if o.DeviceOnboardInformation != nil {
		toSerialize["deviceOnboardInformation"] = o.DeviceOnboardInformation
	}
	return json.Marshal(toSerialize)
}

type NullableLegacySiteDeviceOnboard struct {
	value *LegacySiteDeviceOnboard
	isSet bool
}

func (v NullableLegacySiteDeviceOnboard) Get() *LegacySiteDeviceOnboard {
	return v.value
}

func (v *NullableLegacySiteDeviceOnboard) Set(val *LegacySiteDeviceOnboard) {
	v.value = val
	v.isSet = true
}

func (v NullableLegacySiteDeviceOnboard) IsSet() bool {
	return v.isSet
}

func (v *NullableLegacySiteDeviceOnboard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLegacySiteDeviceOnboard(val *LegacySiteDeviceOnboard) *NullableLegacySiteDeviceOnboard {
	return &NullableLegacySiteDeviceOnboard{value: val, isSet: true}
}

func (v NullableLegacySiteDeviceOnboard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLegacySiteDeviceOnboard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


