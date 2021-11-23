# BillingPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] [readonly] 
**ModifiedOn** | Pointer to **time.Time** |  | [optional] [readonly] 
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

### NewBillingPrice

`func NewBillingPrice(name string, type_ string, price float64, tenantId string, ) *BillingPrice`

NewBillingPrice instantiates a new BillingPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingPriceWithDefaults

`func NewBillingPriceWithDefaults() *BillingPrice`

NewBillingPriceWithDefaults instantiates a new BillingPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillingPrice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingPrice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingPrice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingPrice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedOn

`func (o *BillingPrice) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *BillingPrice) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *BillingPrice) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *BillingPrice) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetModifiedOn

`func (o *BillingPrice) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *BillingPrice) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *BillingPrice) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *BillingPrice) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetName

`func (o *BillingPrice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingPrice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingPrice) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BillingPrice) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillingPrice) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillingPrice) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillingPrice) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *BillingPrice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BillingPrice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BillingPrice) SetType(v string)`

SetType sets Type field to given value.


### GetSubtype

`func (o *BillingPrice) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *BillingPrice) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *BillingPrice) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *BillingPrice) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetSource

`func (o *BillingPrice) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BillingPrice) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BillingPrice) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *BillingPrice) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetPrice

`func (o *BillingPrice) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillingPrice) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillingPrice) SetPrice(v float64)`

SetPrice sets Price field to given value.


### GetBillingPeriod

`func (o *BillingPrice) GetBillingPeriod() int64`

GetBillingPeriod returns the BillingPeriod field if non-nil, zero value otherwise.

### GetBillingPeriodOk

`func (o *BillingPrice) GetBillingPeriodOk() (*int64, bool)`

GetBillingPeriodOk returns a tuple with the BillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriod

`func (o *BillingPrice) SetBillingPeriod(v int64)`

SetBillingPeriod sets BillingPeriod field to given value.

### HasBillingPeriod

`func (o *BillingPrice) HasBillingPeriod() bool`

HasBillingPeriod returns a boolean if a field has been set.

### GetService

`func (o *BillingPrice) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *BillingPrice) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *BillingPrice) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *BillingPrice) HasService() bool`

HasService returns a boolean if a field has been set.

### SetServiceNil

`func (o *BillingPrice) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *BillingPrice) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil
### GetTenantId

`func (o *BillingPrice) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BillingPrice) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BillingPrice) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


