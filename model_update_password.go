/*
MSX SDK

MSX SDK client.

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msxsdk

import (
	"encoding/json"
)

// UpdatePassword struct for UpdatePassword
type UpdatePassword struct {
	Username string `json:"username"`
	OldPassword NullableString `json:"oldPassword,omitempty"`
	NewPassword string `json:"newPassword"`
	AdditionalProperties map[string]interface{}
}

type _UpdatePassword UpdatePassword

// NewUpdatePassword instantiates a new UpdatePassword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePassword(username string, newPassword string) *UpdatePassword {
	this := UpdatePassword{}
	this.Username = username
	this.NewPassword = newPassword
	return &this
}

// NewUpdatePasswordWithDefaults instantiates a new UpdatePassword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePasswordWithDefaults() *UpdatePassword {
	this := UpdatePassword{}
	return &this
}

// GetUsername returns the Username field value
func (o *UpdatePassword) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *UpdatePassword) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *UpdatePassword) SetUsername(v string) {
	o.Username = v
}

// GetOldPassword returns the OldPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdatePassword) GetOldPassword() string {
	if o == nil || o.OldPassword.Get() == nil {
		var ret string
		return ret
	}
	return *o.OldPassword.Get()
}

// GetOldPasswordOk returns a tuple with the OldPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdatePassword) GetOldPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OldPassword.Get(), o.OldPassword.IsSet()
}

// HasOldPassword returns a boolean if a field has been set.
func (o *UpdatePassword) HasOldPassword() bool {
	if o != nil && o.OldPassword.IsSet() {
		return true
	}

	return false
}

// SetOldPassword gets a reference to the given NullableString and assigns it to the OldPassword field.
func (o *UpdatePassword) SetOldPassword(v string) {
	o.OldPassword.Set(&v)
}
// SetOldPasswordNil sets the value for OldPassword to be an explicit nil
func (o *UpdatePassword) SetOldPasswordNil() {
	o.OldPassword.Set(nil)
}

// UnsetOldPassword ensures that no value is present for OldPassword, not even an explicit nil
func (o *UpdatePassword) UnsetOldPassword() {
	o.OldPassword.Unset()
}

// GetNewPassword returns the NewPassword field value
func (o *UpdatePassword) GetNewPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewPassword
}

// GetNewPasswordOk returns a tuple with the NewPassword field value
// and a boolean to check if the value has been set.
func (o *UpdatePassword) GetNewPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NewPassword, true
}

// SetNewPassword sets field value
func (o *UpdatePassword) SetNewPassword(v string) {
	o.NewPassword = v
}

func (o UpdatePassword) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["username"] = o.Username
	}
	if o.OldPassword.IsSet() {
		toSerialize["oldPassword"] = o.OldPassword.Get()
	}
	if true {
		toSerialize["newPassword"] = o.NewPassword
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UpdatePassword) UnmarshalJSON(bytes []byte) (err error) {
	varUpdatePassword := _UpdatePassword{}

	if err = json.Unmarshal(bytes, &varUpdatePassword); err == nil {
		*o = UpdatePassword(varUpdatePassword)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "username")
		delete(additionalProperties, "oldPassword")
		delete(additionalProperties, "newPassword")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdatePassword struct {
	value *UpdatePassword
	isSet bool
}

func (v NullableUpdatePassword) Get() *UpdatePassword {
	return v.value
}

func (v *NullableUpdatePassword) Set(val *UpdatePassword) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePassword) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePassword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePassword(val *UpdatePassword) *NullableUpdatePassword {
	return &NullableUpdatePassword{value: val, isSet: true}
}

func (v NullableUpdatePassword) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePassword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


