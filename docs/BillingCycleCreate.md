# BillingCycleCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** |  | 
**LastBilledOn** | **time.Time** |  | 
**NextBilledOn** | **time.Time** |  | 
**TenantId** | **string** |  | 

## Methods

### NewBillingCycleCreate

`func NewBillingCycleCreate(eventId string, lastBilledOn time.Time, nextBilledOn time.Time, tenantId string, ) *BillingCycleCreate`

NewBillingCycleCreate instantiates a new BillingCycleCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingCycleCreateWithDefaults

`func NewBillingCycleCreateWithDefaults() *BillingCycleCreate`

NewBillingCycleCreateWithDefaults instantiates a new BillingCycleCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *BillingCycleCreate) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *BillingCycleCreate) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *BillingCycleCreate) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetLastBilledOn

`func (o *BillingCycleCreate) GetLastBilledOn() time.Time`

GetLastBilledOn returns the LastBilledOn field if non-nil, zero value otherwise.

### GetLastBilledOnOk

`func (o *BillingCycleCreate) GetLastBilledOnOk() (*time.Time, bool)`

GetLastBilledOnOk returns a tuple with the LastBilledOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBilledOn

`func (o *BillingCycleCreate) SetLastBilledOn(v time.Time)`

SetLastBilledOn sets LastBilledOn field to given value.


### GetNextBilledOn

`func (o *BillingCycleCreate) GetNextBilledOn() time.Time`

GetNextBilledOn returns the NextBilledOn field if non-nil, zero value otherwise.

### GetNextBilledOnOk

`func (o *BillingCycleCreate) GetNextBilledOnOk() (*time.Time, bool)`

GetNextBilledOnOk returns a tuple with the NextBilledOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextBilledOn

`func (o *BillingCycleCreate) SetNextBilledOn(v time.Time)`

SetNextBilledOn sets NextBilledOn field to given value.


### GetTenantId

`func (o *BillingCycleCreate) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BillingCycleCreate) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BillingCycleCreate) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


