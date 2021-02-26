# ServiceElementPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OneTimePrice** | Pointer to **string** |  | [optional] 
**PeriodicPrice** | Pointer to **string** |  | [optional] 
**TimePeriod** | Pointer to **string** |  | [optional] 
**UnitOfMeasure** | Pointer to **string** |  | [optional] 
**IncludedQuantity** | Pointer to **int32** |  | [optional] 
**AdditionalOneTimePrice** | Pointer to **string** |  | [optional] 
**AdditionalPeriodicPrice** | Pointer to **string** |  | [optional] 
**AdditionalQuantity** | Pointer to **int32** |  | [optional] 

## Methods

### NewServiceElementPrice

`func NewServiceElementPrice() *ServiceElementPrice`

NewServiceElementPrice instantiates a new ServiceElementPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceElementPriceWithDefaults

`func NewServiceElementPriceWithDefaults() *ServiceElementPrice`

NewServiceElementPriceWithDefaults instantiates a new ServiceElementPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOneTimePrice

`func (o *ServiceElementPrice) GetOneTimePrice() string`

GetOneTimePrice returns the OneTimePrice field if non-nil, zero value otherwise.

### GetOneTimePriceOk

`func (o *ServiceElementPrice) GetOneTimePriceOk() (*string, bool)`

GetOneTimePriceOk returns a tuple with the OneTimePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneTimePrice

`func (o *ServiceElementPrice) SetOneTimePrice(v string)`

SetOneTimePrice sets OneTimePrice field to given value.

### HasOneTimePrice

`func (o *ServiceElementPrice) HasOneTimePrice() bool`

HasOneTimePrice returns a boolean if a field has been set.

### GetPeriodicPrice

`func (o *ServiceElementPrice) GetPeriodicPrice() string`

GetPeriodicPrice returns the PeriodicPrice field if non-nil, zero value otherwise.

### GetPeriodicPriceOk

`func (o *ServiceElementPrice) GetPeriodicPriceOk() (*string, bool)`

GetPeriodicPriceOk returns a tuple with the PeriodicPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicPrice

`func (o *ServiceElementPrice) SetPeriodicPrice(v string)`

SetPeriodicPrice sets PeriodicPrice field to given value.

### HasPeriodicPrice

`func (o *ServiceElementPrice) HasPeriodicPrice() bool`

HasPeriodicPrice returns a boolean if a field has been set.

### GetTimePeriod

`func (o *ServiceElementPrice) GetTimePeriod() string`

GetTimePeriod returns the TimePeriod field if non-nil, zero value otherwise.

### GetTimePeriodOk

`func (o *ServiceElementPrice) GetTimePeriodOk() (*string, bool)`

GetTimePeriodOk returns a tuple with the TimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriod

`func (o *ServiceElementPrice) SetTimePeriod(v string)`

SetTimePeriod sets TimePeriod field to given value.

### HasTimePeriod

`func (o *ServiceElementPrice) HasTimePeriod() bool`

HasTimePeriod returns a boolean if a field has been set.

### GetUnitOfMeasure

`func (o *ServiceElementPrice) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *ServiceElementPrice) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *ServiceElementPrice) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *ServiceElementPrice) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.

### GetIncludedQuantity

`func (o *ServiceElementPrice) GetIncludedQuantity() int32`

GetIncludedQuantity returns the IncludedQuantity field if non-nil, zero value otherwise.

### GetIncludedQuantityOk

`func (o *ServiceElementPrice) GetIncludedQuantityOk() (*int32, bool)`

GetIncludedQuantityOk returns a tuple with the IncludedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedQuantity

`func (o *ServiceElementPrice) SetIncludedQuantity(v int32)`

SetIncludedQuantity sets IncludedQuantity field to given value.

### HasIncludedQuantity

`func (o *ServiceElementPrice) HasIncludedQuantity() bool`

HasIncludedQuantity returns a boolean if a field has been set.

### GetAdditionalOneTimePrice

`func (o *ServiceElementPrice) GetAdditionalOneTimePrice() string`

GetAdditionalOneTimePrice returns the AdditionalOneTimePrice field if non-nil, zero value otherwise.

### GetAdditionalOneTimePriceOk

`func (o *ServiceElementPrice) GetAdditionalOneTimePriceOk() (*string, bool)`

GetAdditionalOneTimePriceOk returns a tuple with the AdditionalOneTimePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalOneTimePrice

`func (o *ServiceElementPrice) SetAdditionalOneTimePrice(v string)`

SetAdditionalOneTimePrice sets AdditionalOneTimePrice field to given value.

### HasAdditionalOneTimePrice

`func (o *ServiceElementPrice) HasAdditionalOneTimePrice() bool`

HasAdditionalOneTimePrice returns a boolean if a field has been set.

### GetAdditionalPeriodicPrice

`func (o *ServiceElementPrice) GetAdditionalPeriodicPrice() string`

GetAdditionalPeriodicPrice returns the AdditionalPeriodicPrice field if non-nil, zero value otherwise.

### GetAdditionalPeriodicPriceOk

`func (o *ServiceElementPrice) GetAdditionalPeriodicPriceOk() (*string, bool)`

GetAdditionalPeriodicPriceOk returns a tuple with the AdditionalPeriodicPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPeriodicPrice

`func (o *ServiceElementPrice) SetAdditionalPeriodicPrice(v string)`

SetAdditionalPeriodicPrice sets AdditionalPeriodicPrice field to given value.

### HasAdditionalPeriodicPrice

`func (o *ServiceElementPrice) HasAdditionalPeriodicPrice() bool`

HasAdditionalPeriodicPrice returns a boolean if a field has been set.

### GetAdditionalQuantity

`func (o *ServiceElementPrice) GetAdditionalQuantity() int32`

GetAdditionalQuantity returns the AdditionalQuantity field if non-nil, zero value otherwise.

### GetAdditionalQuantityOk

`func (o *ServiceElementPrice) GetAdditionalQuantityOk() (*int32, bool)`

GetAdditionalQuantityOk returns a tuple with the AdditionalQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalQuantity

`func (o *ServiceElementPrice) SetAdditionalQuantity(v int32)`

SetAdditionalQuantity sets AdditionalQuantity field to given value.

### HasAdditionalQuantity

`func (o *ServiceElementPrice) HasAdditionalQuantity() bool`

HasAdditionalQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


