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
)

// UserCreate struct for UserCreate
type UserCreate struct {
	FirstName *string `json:"firstName,omitempty"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	RoleIds *[]string `json:"roleIds,omitempty"`
	TenantIds *[]string `json:"tenantIds,omitempty"`
	PasswordPolicyName string `json:"passwordPolicyName"`
	Locale NullableString `json:"locale,omitempty"`
	Username *string `json:"username,omitempty"`
	Password NullableString `json:"password,omitempty"`
}

// NewUserCreate instantiates a new UserCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCreate(lastName string, email string, passwordPolicyName string) *UserCreate {
	this := UserCreate{}
	this.LastName = lastName
	this.Email = email
	this.PasswordPolicyName = passwordPolicyName
	return &this
}

// NewUserCreateWithDefaults instantiates a new UserCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCreateWithDefaults() *UserCreate {
	this := UserCreate{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UserCreate) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreate) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UserCreate) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UserCreate) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value
func (o *UserCreate) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *UserCreate) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *UserCreate) SetLastName(v string) {
	o.LastName = v
}

// GetEmail returns the Email field value
func (o *UserCreate) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserCreate) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserCreate) SetEmail(v string) {
	o.Email = v
}

// GetRoleIds returns the RoleIds field value if set, zero value otherwise.
func (o *UserCreate) GetRoleIds() []string {
	if o == nil || o.RoleIds == nil {
		var ret []string
		return ret
	}
	return *o.RoleIds
}

// GetRoleIdsOk returns a tuple with the RoleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreate) GetRoleIdsOk() (*[]string, bool) {
	if o == nil || o.RoleIds == nil {
		return nil, false
	}
	return o.RoleIds, true
}

// HasRoleIds returns a boolean if a field has been set.
func (o *UserCreate) HasRoleIds() bool {
	if o != nil && o.RoleIds != nil {
		return true
	}

	return false
}

// SetRoleIds gets a reference to the given []string and assigns it to the RoleIds field.
func (o *UserCreate) SetRoleIds(v []string) {
	o.RoleIds = &v
}

// GetTenantIds returns the TenantIds field value if set, zero value otherwise.
func (o *UserCreate) GetTenantIds() []string {
	if o == nil || o.TenantIds == nil {
		var ret []string
		return ret
	}
	return *o.TenantIds
}

// GetTenantIdsOk returns a tuple with the TenantIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreate) GetTenantIdsOk() (*[]string, bool) {
	if o == nil || o.TenantIds == nil {
		return nil, false
	}
	return o.TenantIds, true
}

// HasTenantIds returns a boolean if a field has been set.
func (o *UserCreate) HasTenantIds() bool {
	if o != nil && o.TenantIds != nil {
		return true
	}

	return false
}

// SetTenantIds gets a reference to the given []string and assigns it to the TenantIds field.
func (o *UserCreate) SetTenantIds(v []string) {
	o.TenantIds = &v
}

// GetPasswordPolicyName returns the PasswordPolicyName field value
func (o *UserCreate) GetPasswordPolicyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PasswordPolicyName
}

// GetPasswordPolicyNameOk returns a tuple with the PasswordPolicyName field value
// and a boolean to check if the value has been set.
func (o *UserCreate) GetPasswordPolicyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PasswordPolicyName, true
}

// SetPasswordPolicyName sets field value
func (o *UserCreate) SetPasswordPolicyName(v string) {
	o.PasswordPolicyName = v
}

// GetLocale returns the Locale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserCreate) GetLocale() string {
	if o == nil || o.Locale.Get() == nil {
		var ret string
		return ret
	}
	return *o.Locale.Get()
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserCreate) GetLocaleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Locale.Get(), o.Locale.IsSet()
}

// HasLocale returns a boolean if a field has been set.
func (o *UserCreate) HasLocale() bool {
	if o != nil && o.Locale.IsSet() {
		return true
	}

	return false
}

// SetLocale gets a reference to the given NullableString and assigns it to the Locale field.
func (o *UserCreate) SetLocale(v string) {
	o.Locale.Set(&v)
}
// SetLocaleNil sets the value for Locale to be an explicit nil
func (o *UserCreate) SetLocaleNil() {
	o.Locale.Set(nil)
}

// UnsetLocale ensures that no value is present for Locale, not even an explicit nil
func (o *UserCreate) UnsetLocale() {
	o.Locale.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UserCreate) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreate) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UserCreate) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UserCreate) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserCreate) GetPassword() string {
	if o == nil || o.Password.Get() == nil {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserCreate) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *UserCreate) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *UserCreate) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *UserCreate) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *UserCreate) UnsetPassword() {
	o.Password.Unset()
}

func (o UserCreate) MarshalJSON() ([]byte, error) {
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
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUserCreate struct {
	value *UserCreate
	isSet bool
}

func (v NullableUserCreate) Get() *UserCreate {
	return v.value
}

func (v *NullableUserCreate) Set(val *UserCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCreate(val *UserCreate) *NullableUserCreate {
	return &NullableUserCreate{value: val, isSet: true}
}

func (v NullableUserCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


