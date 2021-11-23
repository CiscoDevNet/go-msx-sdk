# SmartAccountUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **int32** | Smart Account identifier | [optional] [readonly] 
**AccountName** | Pointer to **string** | Smart Account Name | [optional] [readonly] 
**AccountDomain** | Pointer to **string** | Smart Accont Domain | [optional] [readonly] 
**AccountType** | Pointer to [**SmartAccountType**](SmartAccountType.md) |  | [optional] 
**Roles** | Pointer to [**[]SmartAccountUserRole**](SmartAccountUserRole.md) |  | [optional] 

## Methods

### NewSmartAccountUser

`func NewSmartAccountUser() *SmartAccountUser`

NewSmartAccountUser instantiates a new SmartAccountUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartAccountUserWithDefaults

`func NewSmartAccountUserWithDefaults() *SmartAccountUser`

NewSmartAccountUserWithDefaults instantiates a new SmartAccountUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *SmartAccountUser) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SmartAccountUser) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SmartAccountUser) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *SmartAccountUser) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountName

`func (o *SmartAccountUser) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *SmartAccountUser) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *SmartAccountUser) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *SmartAccountUser) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAccountDomain

`func (o *SmartAccountUser) GetAccountDomain() string`

GetAccountDomain returns the AccountDomain field if non-nil, zero value otherwise.

### GetAccountDomainOk

`func (o *SmartAccountUser) GetAccountDomainOk() (*string, bool)`

GetAccountDomainOk returns a tuple with the AccountDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDomain

`func (o *SmartAccountUser) SetAccountDomain(v string)`

SetAccountDomain sets AccountDomain field to given value.

### HasAccountDomain

`func (o *SmartAccountUser) HasAccountDomain() bool`

HasAccountDomain returns a boolean if a field has been set.

### GetAccountType

`func (o *SmartAccountUser) GetAccountType() SmartAccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *SmartAccountUser) GetAccountTypeOk() (*SmartAccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *SmartAccountUser) SetAccountType(v SmartAccountType)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *SmartAccountUser) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetRoles

`func (o *SmartAccountUser) GetRoles() []SmartAccountUserRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *SmartAccountUser) GetRolesOk() (*[]SmartAccountUserRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *SmartAccountUser) SetRoles(v []SmartAccountUserRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *SmartAccountUser) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


