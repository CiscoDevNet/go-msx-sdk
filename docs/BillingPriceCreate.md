# BillingPriceCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**Subtype** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Price** | **float64** |  | 
**BillingPeriod** | Pointer to **int64** |  | [optional] 
**Service** | Pointer to **NullableString** |  | [optional] 
**TenantId** | **string** |  | 

## Methods

### NewBillingPriceCreate

`func NewBillingPriceCreate(name string, type_ string, price float64, tenantId string, ) *BillingPriceCreate`

NewBillingPriceCreate instantiates a new BillingPriceCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingPriceCreateWithDefaults

`func NewBillingPriceCreateWithDefaults() *BillingPriceCreate`

NewBillingPriceCreateWithDefaults instantiates a new BillingPriceCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BillingPriceCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingPriceCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingPriceCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BillingPriceCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillingPriceCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillingPriceCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillingPriceCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *BillingPriceCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BillingPriceCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BillingPriceCreate) SetType(v string)`

SetType sets Type field to given value.


### GetSubtype

`func (o *BillingPriceCreate) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *BillingPriceCreate) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *BillingPriceCreate) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *BillingPriceCreate) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetSource

`func (o *BillingPriceCreate) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BillingPriceCreate) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BillingPriceCreate) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *BillingPriceCreate) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetPrice

`func (o *BillingPriceCreate) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillingPriceCreate) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillingPriceCreate) SetPrice(v float64)`

SetPrice sets Price field to given value.


### GetBillingPeriod

`func (o *BillingPriceCreate) GetBillingPeriod() int64`

GetBillingPeriod returns the BillingPeriod field if non-nil, zero value otherwise.

### GetBillingPeriodOk

`func (o *BillingPriceCreate) GetBillingPeriodOk() (*int64, bool)`

GetBillingPeriodOk returns a tuple with the BillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriod

`func (o *BillingPriceCreate) SetBillingPeriod(v int64)`

SetBillingPeriod sets BillingPeriod field to given value.

### HasBillingPeriod

`func (o *BillingPriceCreate) HasBillingPeriod() bool`

HasBillingPeriod returns a boolean if a field has been set.

### GetService

`func (o *BillingPriceCreate) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *BillingPriceCreate) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *BillingPriceCreate) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *BillingPriceCreate) HasService() bool`

HasService returns a boolean if a field has been set.

### SetServiceNil

`func (o *BillingPriceCreate) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *BillingPriceCreate) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil
### GetTenantId

`func (o *BillingPriceCreate) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BillingPriceCreate) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BillingPriceCreate) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


