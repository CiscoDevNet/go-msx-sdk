# DeviceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**ProviderId** | Pointer to **string** |  | [optional] 
**VulnerabilityState** | Pointer to **string** | Valid values: VULNERABLE, NOT_VULNERABLE, NOT_APPLICABLE, UNKNOWN | [optional] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] 
**ModifiedOn** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewDeviceAllOf

`func NewDeviceAllOf() *DeviceAllOf`

NewDeviceAllOf instantiates a new DeviceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAllOfWithDefaults

`func NewDeviceAllOfWithDefaults() *DeviceAllOf`

NewDeviceAllOfWithDefaults instantiates a new DeviceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *DeviceAllOf) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DeviceAllOf) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DeviceAllOf) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *DeviceAllOf) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetProviderId

`func (o *DeviceAllOf) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *DeviceAllOf) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *DeviceAllOf) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *DeviceAllOf) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetVulnerabilityState

`func (o *DeviceAllOf) GetVulnerabilityState() string`

GetVulnerabilityState returns the VulnerabilityState field if non-nil, zero value otherwise.

### GetVulnerabilityStateOk

`func (o *DeviceAllOf) GetVulnerabilityStateOk() (*string, bool)`

GetVulnerabilityStateOk returns a tuple with the VulnerabilityState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityState

`func (o *DeviceAllOf) SetVulnerabilityState(v string)`

SetVulnerabilityState sets VulnerabilityState field to given value.

### HasVulnerabilityState

`func (o *DeviceAllOf) HasVulnerabilityState() bool`

HasVulnerabilityState returns a boolean if a field has been set.

### GetCreatedOn

`func (o *DeviceAllOf) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *DeviceAllOf) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *DeviceAllOf) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *DeviceAllOf) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetModifiedOn

`func (o *DeviceAllOf) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *DeviceAllOf) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *DeviceAllOf) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *DeviceAllOf) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### SetModifiedOnNil

`func (o *DeviceAllOf) SetModifiedOnNil(b bool)`

 SetModifiedOnNil sets the value for ModifiedOn to be an explicit nil

### UnsetModifiedOn
`func (o *DeviceAllOf) UnsetModifiedOn()`

UnsetModifiedOn ensures that no value is present for ModifiedOn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


