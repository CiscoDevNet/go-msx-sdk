/*
 * KAKAPO - MSX SDK
 *
 * MSX Platform SDK
 *
 * API version: 1.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msxsdk

import (
	"encoding/json"
)

// DeviceCreateAllOf struct for DeviceCreateAllOf
type DeviceCreateAllOf struct {
	ServiceInstanceId *string `json:"serviceInstanceId,omitempty"`
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	TenantId string `json:"tenantId"`
}

// NewDeviceCreateAllOf instantiates a new DeviceCreateAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceCreateAllOf(tenantId string) *DeviceCreateAllOf {
	this := DeviceCreateAllOf{}
	this.TenantId = tenantId
	return &this
}

// NewDeviceCreateAllOfWithDefaults instantiates a new DeviceCreateAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceCreateAllOfWithDefaults() *DeviceCreateAllOf {
	this := DeviceCreateAllOf{}
	return &this
}

// GetServiceInstanceId returns the ServiceInstanceId field value if set, zero value otherwise.
func (o *DeviceCreateAllOf) GetServiceInstanceId() string {
	if o == nil || o.ServiceInstanceId == nil {
		var ret string
		return ret
	}
	return *o.ServiceInstanceId
}

// GetServiceInstanceIdOk returns a tuple with the ServiceInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateAllOf) GetServiceInstanceIdOk() (*string, bool) {
	if o == nil || o.ServiceInstanceId == nil {
		return nil, false
	}
	return o.ServiceInstanceId, true
}

// HasServiceInstanceId returns a boolean if a field has been set.
func (o *DeviceCreateAllOf) HasServiceInstanceId() bool {
	if o != nil && o.ServiceInstanceId != nil {
		return true
	}

	return false
}

// SetServiceInstanceId gets a reference to the given string and assigns it to the ServiceInstanceId field.
func (o *DeviceCreateAllOf) SetServiceInstanceId(v string) {
	o.ServiceInstanceId = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *DeviceCreateAllOf) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreateAllOf) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || o.SubscriptionId == nil {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *DeviceCreateAllOf) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId != nil {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *DeviceCreateAllOf) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetTenantId returns the TenantId field value
func (o *DeviceCreateAllOf) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *DeviceCreateAllOf) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *DeviceCreateAllOf) SetTenantId(v string) {
	o.TenantId = v
}

func (o DeviceCreateAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServiceInstanceId != nil {
		toSerialize["serviceInstanceId"] = o.ServiceInstanceId
	}
	if o.SubscriptionId != nil {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if true {
		toSerialize["tenantId"] = o.TenantId
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceCreateAllOf struct {
	value *DeviceCreateAllOf
	isSet bool
}

func (v NullableDeviceCreateAllOf) Get() *DeviceCreateAllOf {
	return v.value
}

func (v *NullableDeviceCreateAllOf) Set(val *DeviceCreateAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceCreateAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceCreateAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceCreateAllOf(val *DeviceCreateAllOf) *NullableDeviceCreateAllOf {
	return &NullableDeviceCreateAllOf{value: val, isSet: true}
}

func (v NullableDeviceCreateAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceCreateAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


