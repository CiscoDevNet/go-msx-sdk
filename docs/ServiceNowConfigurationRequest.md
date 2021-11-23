# ServiceNowConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**CriticalEvent** | Pointer to **bool** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**Proxy** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewServiceNowConfigurationRequest

`func NewServiceNowConfigurationRequest() *ServiceNowConfigurationRequest`

NewServiceNowConfigurationRequest instantiates a new ServiceNowConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceNowConfigurationRequestWithDefaults

`func NewServiceNowConfigurationRequestWithDefaults() *ServiceNowConfigurationRequest`

NewServiceNowConfigurationRequestWithDefaults instantiates a new ServiceNowConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *ServiceNowConfigurationRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ServiceNowConfigurationRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ServiceNowConfigurationRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ServiceNowConfigurationRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *ServiceNowConfigurationRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ServiceNowConfigurationRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ServiceNowConfigurationRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *ServiceNowConfigurationRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetCriticalEvent

`func (o *ServiceNowConfigurationRequest) GetCriticalEvent() bool`

GetCriticalEvent returns the CriticalEvent field if non-nil, zero value otherwise.

### GetCriticalEventOk

`func (o *ServiceNowConfigurationRequest) GetCriticalEventOk() (*bool, bool)`

GetCriticalEventOk returns a tuple with the CriticalEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalEvent

`func (o *ServiceNowConfigurationRequest) SetCriticalEvent(v bool)`

SetCriticalEvent sets CriticalEvent field to given value.

### HasCriticalEvent

`func (o *ServiceNowConfigurationRequest) HasCriticalEvent() bool`

HasCriticalEvent returns a boolean if a field has been set.

### GetDomain

`func (o *ServiceNowConfigurationRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ServiceNowConfigurationRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ServiceNowConfigurationRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ServiceNowConfigurationRequest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetPassword

`func (o *ServiceNowConfigurationRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ServiceNowConfigurationRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ServiceNowConfigurationRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ServiceNowConfigurationRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUserName

`func (o *ServiceNowConfigurationRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ServiceNowConfigurationRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ServiceNowConfigurationRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ServiceNowConfigurationRequest) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetTenantId

`func (o *ServiceNowConfigurationRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ServiceNowConfigurationRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ServiceNowConfigurationRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ServiceNowConfigurationRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ServiceNowConfigurationRequest) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ServiceNowConfigurationRequest) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetProxy

`func (o *ServiceNowConfigurationRequest) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *ServiceNowConfigurationRequest) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *ServiceNowConfigurationRequest) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *ServiceNowConfigurationRequest) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *ServiceNowConfigurationRequest) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *ServiceNowConfigurationRequest) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


