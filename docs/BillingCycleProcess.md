# BillingCycleProcess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextBilledOn** | **time.Time** |  | 

## Methods

### NewBillingCycleProcess

`func NewBillingCycleProcess(nextBilledOn time.Time, ) *BillingCycleProcess`

NewBillingCycleProcess instantiates a new BillingCycleProcess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingCycleProcessWithDefaults

`func NewBillingCycleProcessWithDefaults() *BillingCycleProcess`

NewBillingCycleProcessWithDefaults instantiates a new BillingCycleProcess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextBilledOn

`func (o *BillingCycleProcess) GetNextBilledOn() time.Time`

GetNextBilledOn returns the NextBilledOn field if non-nil, zero value otherwise.

### GetNextBilledOnOk

`func (o *BillingCycleProcess) GetNextBilledOnOk() (*time.Time, bool)`

GetNextBilledOnOk returns a tuple with the NextBilledOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextBilledOn

`func (o *BillingCycleProcess) SetNextBilledOn(v time.Time)`

SetNextBilledOn sets NextBilledOn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


