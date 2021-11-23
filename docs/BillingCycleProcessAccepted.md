# BillingCycleProcessAccepted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**NextBilledOn** | **time.Time** |  | 
**CreatedOn** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewBillingCycleProcessAccepted

`func NewBillingCycleProcessAccepted(nextBilledOn time.Time, ) *BillingCycleProcessAccepted`

NewBillingCycleProcessAccepted instantiates a new BillingCycleProcessAccepted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingCycleProcessAcceptedWithDefaults

`func NewBillingCycleProcessAcceptedWithDefaults() *BillingCycleProcessAccepted`

NewBillingCycleProcessAcceptedWithDefaults instantiates a new BillingCycleProcessAccepted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillingCycleProcessAccepted) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingCycleProcessAccepted) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingCycleProcessAccepted) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingCycleProcessAccepted) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BillingCycleProcessAccepted) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BillingCycleProcessAccepted) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetNextBilledOn

`func (o *BillingCycleProcessAccepted) GetNextBilledOn() time.Time`

GetNextBilledOn returns the NextBilledOn field if non-nil, zero value otherwise.

### GetNextBilledOnOk

`func (o *BillingCycleProcessAccepted) GetNextBilledOnOk() (*time.Time, bool)`

GetNextBilledOnOk returns a tuple with the NextBilledOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextBilledOn

`func (o *BillingCycleProcessAccepted) SetNextBilledOn(v time.Time)`

SetNextBilledOn sets NextBilledOn field to given value.


### GetCreatedOn

`func (o *BillingCycleProcessAccepted) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *BillingCycleProcessAccepted) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *BillingCycleProcessAccepted) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *BillingCycleProcessAccepted) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


