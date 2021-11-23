# CostSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Subtype** | Pointer to **string** |  | [optional] 
**Cost** | Pointer to **NullableFloat64** |  | [optional] 
**Price** | Pointer to **NullableFloat64** |  | [optional] 

## Methods

### NewCostSummary

`func NewCostSummary() *CostSummary`

NewCostSummary instantiates a new CostSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostSummaryWithDefaults

`func NewCostSummaryWithDefaults() *CostSummary`

NewCostSummaryWithDefaults instantiates a new CostSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CostSummary) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CostSummary) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CostSummary) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CostSummary) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubtype

`func (o *CostSummary) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *CostSummary) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *CostSummary) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *CostSummary) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetCost

`func (o *CostSummary) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *CostSummary) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *CostSummary) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *CostSummary) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *CostSummary) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *CostSummary) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetPrice

`func (o *CostSummary) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CostSummary) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CostSummary) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CostSummary) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *CostSummary) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *CostSummary) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


