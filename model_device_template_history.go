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
	"time"
)

// DeviceTemplateHistory struct for DeviceTemplateHistory
type DeviceTemplateHistory struct {
	Id *string `json:"id,omitempty"`
	DeviceId *string `json:"deviceId,omitempty"`
	InstanceId *string `json:"instanceId,omitempty"`
	TemplateId *string `json:"templateId,omitempty"`
	UserId *string `json:"userId,omitempty"`
	Status *string `json:"status,omitempty"`
	StatusError *string `json:"statusError,omitempty"`
	TemplateParams *[]NameValue `json:"templateParams,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceTemplateHistory DeviceTemplateHistory

// NewDeviceTemplateHistory instantiates a new DeviceTemplateHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceTemplateHistory() *DeviceTemplateHistory {
	this := DeviceTemplateHistory{}
	return &this
}

// NewDeviceTemplateHistoryWithDefaults instantiates a new DeviceTemplateHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceTemplateHistoryWithDefaults() *DeviceTemplateHistory {
	this := DeviceTemplateHistory{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeviceTemplateHistory) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTemplateHistory) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeviceTemplateHistory) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeviceTemplateHistory) SetId(v string) {
	o.Id = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *DeviceTemplateHistory) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTemplateHistory) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *DeviceTemplateHistory) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *DeviceTemplateHistory) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *DeviceTemplateHistory) GetInstanceId() string {
	if o == nil || o.InstanceId == nil {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTemplateHistory) GetInstanceIdOk() (*string, bool) {
	if o == nil || o.InstanceId == nil {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *DeviceTemplateHistory) HasInstanceId() bool {
	if o != nil && o.InstanceId != nil {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *DeviceTemplateHistory) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *DeviceTemplateHistory) GetTemplateId() string {
	if o == nil || o.TemplateId == nil {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTemplateHistory) GetTemplateIdOk() (*string, bool) {
	if o == nil || o.TemplateId == nil {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *DeviceTemplateHistory) HasTemplateId() bool {
	if o != nil && o.TemplateId != nil {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *DeviceTemplateHistory) SetTemplateId(v string) {
	o.TemplateId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *DeviceTemplateHistory) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTemplateHistory) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *DeviceTemplateHistory) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *DeviceTemplateHistory) SetUserId(v string) {
	o.UserId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeviceTemplateHistory) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTemplateHistory) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeviceTemplateHistory) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DeviceTemplateHistory) SetStatus(v string) {
	o.Status = &v
}

// GetStatusError returns the StatusError field value if set, zero value otherwise.
func (o *DeviceTemplateHistory) GetStatusError() string {
	if o == nil || o.StatusError == nil {
		var ret string
		return ret
	}
	return *o.StatusError
}

// GetStatusErrorOk returns a tuple with the StatusError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTemplateHistory) GetStatusErrorOk() (*string, bool) {
	if o == nil || o.StatusError == nil {
		return nil, false
	}
	return o.StatusError, true
}

// HasStatusError returns a boolean if a field has been set.
func (o *DeviceTemplateHistory) HasStatusError() bool {
	if o != nil && o.StatusError != nil {
		return true
	}

	return false
}

// SetStatusError gets a reference to the given string and assigns it to the StatusError field.
func (o *DeviceTemplateHistory) SetStatusError(v string) {
	o.StatusError = &v
}

// GetTemplateParams returns the TemplateParams field value if set, zero value otherwise.
func (o *DeviceTemplateHistory) GetTemplateParams() []NameValue {
	if o == nil || o.TemplateParams == nil {
		var ret []NameValue
		return ret
	}
	return *o.TemplateParams
}

// GetTemplateParamsOk returns a tuple with the TemplateParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTemplateHistory) GetTemplateParamsOk() (*[]NameValue, bool) {
	if o == nil || o.TemplateParams == nil {
		return nil, false
	}
	return o.TemplateParams, true
}

// HasTemplateParams returns a boolean if a field has been set.
func (o *DeviceTemplateHistory) HasTemplateParams() bool {
	if o != nil && o.TemplateParams != nil {
		return true
	}

	return false
}

// SetTemplateParams gets a reference to the given []NameValue and assigns it to the TemplateParams field.
func (o *DeviceTemplateHistory) SetTemplateParams(v []NameValue) {
	o.TemplateParams = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *DeviceTemplateHistory) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTemplateHistory) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *DeviceTemplateHistory) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *DeviceTemplateHistory) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o DeviceTemplateHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DeviceId != nil {
		toSerialize["deviceId"] = o.DeviceId
	}
	if o.InstanceId != nil {
		toSerialize["instanceId"] = o.InstanceId
	}
	if o.TemplateId != nil {
		toSerialize["templateId"] = o.TemplateId
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StatusError != nil {
		toSerialize["statusError"] = o.StatusError
	}
	if o.TemplateParams != nil {
		toSerialize["templateParams"] = o.TemplateParams
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeviceTemplateHistory) UnmarshalJSON(bytes []byte) (err error) {
	varDeviceTemplateHistory := _DeviceTemplateHistory{}

	if err = json.Unmarshal(bytes, &varDeviceTemplateHistory); err == nil {
		*o = DeviceTemplateHistory(varDeviceTemplateHistory)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "deviceId")
		delete(additionalProperties, "instanceId")
		delete(additionalProperties, "templateId")
		delete(additionalProperties, "userId")
		delete(additionalProperties, "status")
		delete(additionalProperties, "statusError")
		delete(additionalProperties, "templateParams")
		delete(additionalProperties, "lastUpdated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceTemplateHistory struct {
	value *DeviceTemplateHistory
	isSet bool
}

func (v NullableDeviceTemplateHistory) Get() *DeviceTemplateHistory {
	return v.value
}

func (v *NullableDeviceTemplateHistory) Set(val *DeviceTemplateHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceTemplateHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceTemplateHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceTemplateHistory(val *DeviceTemplateHistory) *NullableDeviceTemplateHistory {
	return &NullableDeviceTemplateHistory{value: val, isSet: true}
}

func (v NullableDeviceTemplateHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceTemplateHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


