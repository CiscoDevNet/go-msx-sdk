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

// User struct for User
type User struct {
	Id *string `json:"id,omitempty"`
	Status *string `json:"status,omitempty"`
	Deleted *string `json:"deleted,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	RoleIds *[]string `json:"roleIds,omitempty"`
	TenantIds *[]string `json:"tenantIds,omitempty"`
	PasswordPolicyName string `json:"passwordPolicyName"`
	Username *string `json:"username,omitempty"`
	Password NullableString `json:"password,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser(lastName string, email string, passwordPolicyName string) *User {
	this := User{}
	this.LastName = lastName
	this.Email = email
	this.PasswordPolicyName = passwordPolicyName
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *User) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *User) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *User) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *User) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *User) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *User) SetStatus(v string) {
	o.Status = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *User) GetDeleted() string {
	if o == nil || o.Deleted == nil {
		var ret string
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetDeletedOk() (*string, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *User) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given string and assigns it to the Deleted field.
func (o *User) SetDeleted(v string) {
	o.Deleted = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *User) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *User) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *User) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value
func (o *User) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *User) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *User) SetLastName(v string) {
	o.LastName = v
}

// GetEmail returns the Email field value
func (o *User) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *User) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *User) SetEmail(v string) {
	o.Email = v
}

// GetRoleIds returns the RoleIds field value if set, zero value otherwise.
func (o *User) GetRoleIds() []string {
	if o == nil || o.RoleIds == nil {
		var ret []string
		return ret
	}
	return *o.RoleIds
}

// GetRoleIdsOk returns a tuple with the RoleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetRoleIdsOk() (*[]string, bool) {
	if o == nil || o.RoleIds == nil {
		return nil, false
	}
	return o.RoleIds, true
}

// HasRoleIds returns a boolean if a field has been set.
func (o *User) HasRoleIds() bool {
	if o != nil && o.RoleIds != nil {
		return true
	}

	return false
}

// SetRoleIds gets a reference to the given []string and assigns it to the RoleIds field.
func (o *User) SetRoleIds(v []string) {
	o.RoleIds = &v
}

// GetTenantIds returns the TenantIds field value if set, zero value otherwise.
func (o *User) GetTenantIds() []string {
	if o == nil || o.TenantIds == nil {
		var ret []string
		return ret
	}
	return *o.TenantIds
}

// GetTenantIdsOk returns a tuple with the TenantIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetTenantIdsOk() (*[]string, bool) {
	if o == nil || o.TenantIds == nil {
		return nil, false
	}
	return o.TenantIds, true
}

// HasTenantIds returns a boolean if a field has been set.
func (o *User) HasTenantIds() bool {
	if o != nil && o.TenantIds != nil {
		return true
	}

	return false
}

// SetTenantIds gets a reference to the given []string and assigns it to the TenantIds field.
func (o *User) SetTenantIds(v []string) {
	o.TenantIds = &v
}

// GetPasswordPolicyName returns the PasswordPolicyName field value
func (o *User) GetPasswordPolicyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PasswordPolicyName
}

// GetPasswordPolicyNameOk returns a tuple with the PasswordPolicyName field value
// and a boolean to check if the value has been set.
func (o *User) GetPasswordPolicyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PasswordPolicyName, true
}

// SetPasswordPolicyName sets field value
func (o *User) SetPasswordPolicyName(v string) {
	o.PasswordPolicyName = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *User) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *User) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *User) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetPassword() string {
	if o == nil || o.Password.Get() == nil {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *User) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *User) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *User) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *User) UnsetPassword() {
	o.Password.Unset()
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
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
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


