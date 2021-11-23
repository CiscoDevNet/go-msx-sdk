# BillingEventUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedOn** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Subtype** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **map[string]interface{}** |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 

## Methods

### NewBillingEventUpdate

`func NewBillingEventUpdate() *BillingEventUpdate`

NewBillingEventUpdate instantiates a new BillingEventUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingEventUpdateWithDefaults

`func NewBillingEventUpdateWithDefaults() *BillingEventUpdate`

NewBillingEventUpdateWithDefaults instantiates a new BillingEventUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedOn

`func (o *BillingEventUpdate) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *BillingEventUpdate) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *BillingEventUpdate) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *BillingEventUpdate) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetName

`func (o *BillingEventUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingEventUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingEventUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillingEventUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *BillingEventUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillingEventUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillingEventUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillingEventUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *BillingEventUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BillingEventUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BillingEventUpdate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BillingEventUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubtype

`func (o *BillingEventUpdate) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *BillingEventUpdate) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *BillingEventUpdate) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *BillingEventUpdate) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetSeverity

`func (o *BillingEventUpdate) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *BillingEventUpdate) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *BillingEventUpdate) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *BillingEventUpdate) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetAction

`func (o *BillingEventUpdate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BillingEventUpdate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BillingEventUpdate) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *BillingEventUpdate) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDetails

`func (o *BillingEventUpdate) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *BillingEventUpdate) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *BillingEventUpdate) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *BillingEventUpdate) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *BillingEventUpdate) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *BillingEventUpdate) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetPrice

`func (o *BillingEventUpdate) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillingEventUpdate) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillingEventUpdate) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BillingEventUpdate) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetTenantId

`func (o *BillingEventUpdate) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BillingEventUpdate) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BillingEventUpdate) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *BillingEventUpdate) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


