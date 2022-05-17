# UserCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | **string** |  | 
**Email** | **string** |  | 
**RoleIds** | Pointer to **[]string** |  | [optional] 
**TenantIds** | Pointer to **[]string** |  | [optional] 
**PasswordPolicyName** | Pointer to **string** |  | [optional] 
**Locale** | Pointer to **NullableString** |  | [optional] 
**Username** | **string** |  | 
**Password** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUserCreate

`func NewUserCreate(lastName string, email string, username string, ) *UserCreate`

NewUserCreate instantiates a new UserCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCreateWithDefaults

`func NewUserCreateWithDefaults() *UserCreate`

NewUserCreateWithDefaults instantiates a new UserCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *UserCreate) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserCreate) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserCreate) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserCreate) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UserCreate) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserCreate) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserCreate) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetEmail

`func (o *UserCreate) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserCreate) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserCreate) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRoleIds

`func (o *UserCreate) GetRoleIds() []string`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *UserCreate) GetRoleIdsOk() (*[]string, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *UserCreate) SetRoleIds(v []string)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *UserCreate) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.

### GetTenantIds

`func (o *UserCreate) GetTenantIds() []string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *UserCreate) GetTenantIdsOk() (*[]string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *UserCreate) SetTenantIds(v []string)`

SetTenantIds sets TenantIds field to given value.

### HasTenantIds

`func (o *UserCreate) HasTenantIds() bool`

HasTenantIds returns a boolean if a field has been set.

### GetPasswordPolicyName

`func (o *UserCreate) GetPasswordPolicyName() string`

GetPasswordPolicyName returns the PasswordPolicyName field if non-nil, zero value otherwise.

### GetPasswordPolicyNameOk

`func (o *UserCreate) GetPasswordPolicyNameOk() (*string, bool)`

GetPasswordPolicyNameOk returns a tuple with the PasswordPolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicyName

`func (o *UserCreate) SetPasswordPolicyName(v string)`

SetPasswordPolicyName sets PasswordPolicyName field to given value.

### HasPasswordPolicyName

`func (o *UserCreate) HasPasswordPolicyName() bool`

HasPasswordPolicyName returns a boolean if a field has been set.

### GetLocale

`func (o *UserCreate) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *UserCreate) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *UserCreate) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *UserCreate) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *UserCreate) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *UserCreate) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetUsername

`func (o *UserCreate) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserCreate) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserCreate) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *UserCreate) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserCreate) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserCreate) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserCreate) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *UserCreate) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *UserCreate) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


