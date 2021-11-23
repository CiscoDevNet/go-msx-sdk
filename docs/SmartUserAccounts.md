# SmartUserAccounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]SmartAccountUser**](SmartAccountUser.md) |  | [optional] 

## Methods

### NewSmartUserAccounts

`func NewSmartUserAccounts() *SmartUserAccounts`

NewSmartUserAccounts instantiates a new SmartUserAccounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartUserAccountsWithDefaults

`func NewSmartUserAccountsWithDefaults() *SmartUserAccounts`

NewSmartUserAccountsWithDefaults instantiates a new SmartUserAccounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *SmartUserAccounts) GetAccounts() []SmartAccountUser`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *SmartUserAccounts) GetAccountsOk() (*[]SmartAccountUser, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *SmartUserAccounts) SetAccounts(v []SmartAccountUser)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *SmartUserAccounts) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


