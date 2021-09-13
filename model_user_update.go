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

// UserUpdate struct for UserUpdate
type UserUpdate struct {
	FirstName *string `json:"firstName,omitempty"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	RoleIds *[]string `json:"roleIds,omitempty"`
	TenantIds *[]string `json:"tenantIds,omitempty"`
	PasswordPolicyName string `json:"passwordPolicyName"`
	Locale NullableString `json:"locale,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserUpdate UserUpdate

// NewUserUpdate instantiates a new UserUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserUpdate(lastName string, email string, passwordPolicyName string) *UserUpdate {
	this := UserUpdate{}
	this.LastName = lastName
	this.Email = email
	this.PasswordPolicyName = passwordPolicyName
	return &this
}

// NewUserUpdateWithDefaults instantiates a new UserUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserUpdateWithDefaults() *UserUpdate {
	this := UserUpdate{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UserUpdate) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdate) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UserUpdate) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UserUpdate) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value
func (o *UserUpdate) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *UserUpdate) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *UserUpdate) SetLastName(v string) {
	o.LastName = v
}

// GetEmail returns the Email field value
func (o *UserUpdate) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserUpdate) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserUpdate) SetEmail(v string) {
	o.Email = v
}

// GetRoleIds returns the RoleIds field value if set, zero value otherwise.
func (o *UserUpdate) GetRoleIds() []string {
	if o == nil || o.RoleIds == nil {
		var ret []string
		return ret
	}
	return *o.RoleIds
}

// GetRoleIdsOk returns a tuple with the RoleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdate) GetRoleIdsOk() (*[]string, bool) {
	if o == nil || o.RoleIds == nil {
		return nil, false
	}
	return o.RoleIds, true
}

// HasRoleIds returns a boolean if a field has been set.
func (o *UserUpdate) HasRoleIds() bool {
	if o != nil && o.RoleIds != nil {
		return true
	}

	return false
}

// SetRoleIds gets a reference to the given []string and assigns it to the RoleIds field.
func (o *UserUpdate) SetRoleIds(v []string) {
	o.RoleIds = &v
}

// GetTenantIds returns the TenantIds field value if set, zero value otherwise.
func (o *UserUpdate) GetTenantIds() []string {
	if o == nil || o.TenantIds == nil {
		var ret []string
		return ret
	}
	return *o.TenantIds
}

// GetTenantIdsOk returns a tuple with the TenantIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdate) GetTenantIdsOk() (*[]string, bool) {
	if o == nil || o.TenantIds == nil {
		return nil, false
	}
	return o.TenantIds, true
}

// HasTenantIds returns a boolean if a field has been set.
func (o *UserUpdate) HasTenantIds() bool {
	if o != nil && o.TenantIds != nil {
		return true
	}

	return false
}

// SetTenantIds gets a reference to the given []string and assigns it to the TenantIds field.
func (o *UserUpdate) SetTenantIds(v []string) {
	o.TenantIds = &v
}

// GetPasswordPolicyName returns the PasswordPolicyName field value
func (o *UserUpdate) GetPasswordPolicyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PasswordPolicyName
}

// GetPasswordPolicyNameOk returns a tuple with the PasswordPolicyName field value
// and a boolean to check if the value has been set.
func (o *UserUpdate) GetPasswordPolicyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PasswordPolicyName, true
}

// SetPasswordPolicyName sets field value
func (o *UserUpdate) SetPasswordPolicyName(v string) {
	o.PasswordPolicyName = v
}

// GetLocale returns the Locale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUpdate) GetLocale() string {
	if o == nil || o.Locale.Get() == nil {
		var ret string
		return ret
	}
	return *o.Locale.Get()
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUpdate) GetLocaleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Locale.Get(), o.Locale.IsSet()
}

// HasLocale returns a boolean if a field has been set.
func (o *UserUpdate) HasLocale() bool {
	if o != nil && o.Locale.IsSet() {
		return true
	}

	return false
}

// SetLocale gets a reference to the given NullableString and assigns it to the Locale field.
func (o *UserUpdate) SetLocale(v string) {
	o.Locale.Set(&v)
}
// SetLocaleNil sets the value for Locale to be an explicit nil
func (o *UserUpdate) SetLocaleNil() {
	o.Locale.Set(nil)
}

// UnsetLocale ensures that no value is present for Locale, not even an explicit nil
func (o *UserUpdate) UnsetLocale() {
	o.Locale.Unset()
}

func (o UserUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FirstName != nil {
		toSerialize["firstName"] = o.FirstName
	}
	if true {
		toSerialize["lastName"] = o.LastName
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if o.RoleIds != nil {
		toSerialize["roleIds"] = o.RoleIds
	}
	if o.TenantIds != nil {
		toSerialize["tenantIds"] = o.TenantIds
	}
	if true {
		toSerialize["passwordPolicyName"] = o.PasswordPolicyName
	}
	if o.Locale.IsSet() {
		toSerialize["locale"] = o.Locale.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varUserUpdate := _UserUpdate{}

	if err = json.Unmarshal(bytes, &varUserUpdate); err == nil {
		*o = UserUpdate(varUserUpdate)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "firstName")
		delete(additionalProperties, "lastName")
		delete(additionalProperties, "email")
		delete(additionalProperties, "roleIds")
		delete(additionalProperties, "tenantIds")
		delete(additionalProperties, "passwordPolicyName")
		delete(additionalProperties, "locale")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserUpdate struct {
	value *UserUpdate
	isSet bool
}

func (v NullableUserUpdate) Get() *UserUpdate {
	return v.value
}

func (v *NullableUserUpdate) Set(val *UserUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableUserUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableUserUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserUpdate(val *UserUpdate) *NullableUserUpdate {
	return &NullableUserUpdate{value: val, isSet: true}
}

func (v NullableUserUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


