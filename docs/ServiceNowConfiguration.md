# ServiceNowConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**CriticalEvent** | Pointer to **bool** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**Proxy** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewServiceNowConfiguration

`func NewServiceNowConfiguration() *ServiceNowConfiguration`

NewServiceNowConfiguration instantiates a new ServiceNowConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceNowConfigurationWithDefaults

`func NewServiceNowConfigurationWithDefaults() *ServiceNowConfiguration`

NewServiceNowConfigurationWithDefaults instantiates a new ServiceNowConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceNowConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceNowConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceNowConfiguration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceNowConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClientId

`func (o *ServiceNowConfiguration) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ServiceNowConfiguration) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ServiceNowConfiguration) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ServiceNowConfiguration) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *ServiceNowConfiguration) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ServiceNowConfiguration) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ServiceNowConfiguration) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *ServiceNowConfiguration) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetCriticalEvent

`func (o *ServiceNowConfiguration) GetCriticalEvent() bool`

GetCriticalEvent returns the CriticalEvent field if non-nil, zero value otherwise.

### GetCriticalEventOk

`func (o *ServiceNowConfiguration) GetCriticalEventOk() (*bool, bool)`

GetCriticalEventOk returns a tuple with the CriticalEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalEvent

`func (o *ServiceNowConfiguration) SetCriticalEvent(v bool)`

SetCriticalEvent sets CriticalEvent field to given value.

### HasCriticalEvent

`func (o *ServiceNowConfiguration) HasCriticalEvent() bool`

HasCriticalEvent returns a boolean if a field has been set.

### GetDomain

`func (o *ServiceNowConfiguration) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ServiceNowConfiguration) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ServiceNowConfiguration) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ServiceNowConfiguration) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetPassword

`func (o *ServiceNowConfiguration) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ServiceNowConfiguration) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ServiceNowConfiguration) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ServiceNowConfiguration) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUserName

`func (o *ServiceNowConfiguration) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ServiceNowConfiguration) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ServiceNowConfiguration) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ServiceNowConfiguration) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetTenantId

`func (o *ServiceNowConfiguration) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ServiceNowConfiguration) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ServiceNowConfiguration) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ServiceNowConfiguration) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ServiceNowConfiguration) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ServiceNowConfiguration) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetProxy

`func (o *ServiceNowConfiguration) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *ServiceNowConfiguration) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *ServiceNowConfiguration) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *ServiceNowConfiguration) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *ServiceNowConfiguration) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *ServiceNowConfiguration) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


