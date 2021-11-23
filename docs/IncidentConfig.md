# IncidentConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** |  | 
**ClientSecret** | **string** |  | 
**CriticalEvent** | Pointer to **bool** |  | [optional] 
**Domain** | **string** |  | 
**Password** | **string** |  | 
**UserName** | **string** |  | 
**Proxy** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIncidentConfig

`func NewIncidentConfig(clientId string, clientSecret string, domain string, password string, userName string, ) *IncidentConfig`

NewIncidentConfig instantiates a new IncidentConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentConfigWithDefaults

`func NewIncidentConfigWithDefaults() *IncidentConfig`

NewIncidentConfigWithDefaults instantiates a new IncidentConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *IncidentConfig) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IncidentConfig) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IncidentConfig) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *IncidentConfig) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *IncidentConfig) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *IncidentConfig) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetCriticalEvent

`func (o *IncidentConfig) GetCriticalEvent() bool`

GetCriticalEvent returns the CriticalEvent field if non-nil, zero value otherwise.

### GetCriticalEventOk

`func (o *IncidentConfig) GetCriticalEventOk() (*bool, bool)`

GetCriticalEventOk returns a tuple with the CriticalEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalEvent

`func (o *IncidentConfig) SetCriticalEvent(v bool)`

SetCriticalEvent sets CriticalEvent field to given value.

### HasCriticalEvent

`func (o *IncidentConfig) HasCriticalEvent() bool`

HasCriticalEvent returns a boolean if a field has been set.

### GetDomain

`func (o *IncidentConfig) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *IncidentConfig) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *IncidentConfig) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetPassword

`func (o *IncidentConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *IncidentConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *IncidentConfig) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUserName

`func (o *IncidentConfig) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *IncidentConfig) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *IncidentConfig) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetProxy

`func (o *IncidentConfig) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *IncidentConfig) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *IncidentConfig) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *IncidentConfig) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *IncidentConfig) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *IncidentConfig) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


