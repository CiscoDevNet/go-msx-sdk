# BillingCycle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**EventId** | Pointer to **string** |  | [optional] [readonly] 
**LastBilledOn** | Pointer to **time.Time** |  | [optional] 
**NextBilledOn** | Pointer to **time.Time** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 

## Methods

### NewBillingCycle

`func NewBillingCycle() *BillingCycle`

NewBillingCycle instantiates a new BillingCycle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingCycleWithDefaults

`func NewBillingCycleWithDefaults() *BillingCycle`

NewBillingCycleWithDefaults instantiates a new BillingCycle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillingCycle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingCycle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingCycle) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingCycle) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEventId

`func (o *BillingCycle) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *BillingCycle) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *BillingCycle) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *BillingCycle) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetLastBilledOn

`func (o *BillingCycle) GetLastBilledOn() time.Time`

GetLastBilledOn returns the LastBilledOn field if non-nil, zero value otherwise.

### GetLastBilledOnOk

`func (o *BillingCycle) GetLastBilledOnOk() (*time.Time, bool)`

GetLastBilledOnOk returns a tuple with the LastBilledOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBilledOn

`func (o *BillingCycle) SetLastBilledOn(v time.Time)`

SetLastBilledOn sets LastBilledOn field to given value.

### HasLastBilledOn

`func (o *BillingCycle) HasLastBilledOn() bool`

HasLastBilledOn returns a boolean if a field has been set.

### GetNextBilledOn

`func (o *BillingCycle) GetNextBilledOn() time.Time`

GetNextBilledOn returns the NextBilledOn field if non-nil, zero value otherwise.

### GetNextBilledOnOk

`func (o *BillingCycle) GetNextBilledOnOk() (*time.Time, bool)`

GetNextBilledOnOk returns a tuple with the NextBilledOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextBilledOn

`func (o *BillingCycle) SetNextBilledOn(v time.Time)`

SetNextBilledOn sets NextBilledOn field to given value.

### HasNextBilledOn

`func (o *BillingCycle) HasNextBilledOn() bool`

HasNextBilledOn returns a boolean if a field has been set.

### GetTenantId

`func (o *BillingCycle) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BillingCycle) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BillingCycle) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *BillingCycle) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


