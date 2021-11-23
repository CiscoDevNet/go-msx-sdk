# IncidentConfigUpdate

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

### NewIncidentConfigUpdate

`func NewIncidentConfigUpdate(clientId string, clientSecret string, domain string, password string, userName string, ) *IncidentConfigUpdate`

NewIncidentConfigUpdate instantiates a new IncidentConfigUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentConfigUpdateWithDefaults

`func NewIncidentConfigUpdateWithDefaults() *IncidentConfigUpdate`

NewIncidentConfigUpdateWithDefaults instantiates a new IncidentConfigUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *IncidentConfigUpdate) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IncidentConfigUpdate) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IncidentConfigUpdate) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *IncidentConfigUpdate) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *IncidentConfigUpdate) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *IncidentConfigUpdate) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetCriticalEvent

`func (o *IncidentConfigUpdate) GetCriticalEvent() bool`

GetCriticalEvent returns the CriticalEvent field if non-nil, zero value otherwise.

### GetCriticalEventOk

`func (o *IncidentConfigUpdate) GetCriticalEventOk() (*bool, bool)`

GetCriticalEventOk returns a tuple with the CriticalEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalEvent

`func (o *IncidentConfigUpdate) SetCriticalEvent(v bool)`

SetCriticalEvent sets CriticalEvent field to given value.

### HasCriticalEvent

`func (o *IncidentConfigUpdate) HasCriticalEvent() bool`

HasCriticalEvent returns a boolean if a field has been set.

### GetDomain

`func (o *IncidentConfigUpdate) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *IncidentConfigUpdate) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *IncidentConfigUpdate) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetPassword

`func (o *IncidentConfigUpdate) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *IncidentConfigUpdate) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *IncidentConfigUpdate) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUserName

`func (o *IncidentConfigUpdate) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *IncidentConfigUpdate) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *IncidentConfigUpdate) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetProxy

`func (o *IncidentConfigUpdate) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *IncidentConfigUpdate) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *IncidentConfigUpdate) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *IncidentConfigUpdate) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *IncidentConfigUpdate) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *IncidentConfigUpdate) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


