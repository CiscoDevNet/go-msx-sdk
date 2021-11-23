# BillingEventCreate

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

### NewBillingEventCreate

`func NewBillingEventCreate() *BillingEventCreate`

NewBillingEventCreate instantiates a new BillingEventCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingEventCreateWithDefaults

`func NewBillingEventCreateWithDefaults() *BillingEventCreate`

NewBillingEventCreateWithDefaults instantiates a new BillingEventCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedOn

`func (o *BillingEventCreate) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *BillingEventCreate) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *BillingEventCreate) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *BillingEventCreate) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetName

`func (o *BillingEventCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingEventCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingEventCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillingEventCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *BillingEventCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillingEventCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillingEventCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillingEventCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *BillingEventCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BillingEventCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BillingEventCreate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BillingEventCreate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubtype

`func (o *BillingEventCreate) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *BillingEventCreate) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *BillingEventCreate) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *BillingEventCreate) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetSeverity

`func (o *BillingEventCreate) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *BillingEventCreate) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *BillingEventCreate) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *BillingEventCreate) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetAction

`func (o *BillingEventCreate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BillingEventCreate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BillingEventCreate) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *BillingEventCreate) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDetails

`func (o *BillingEventCreate) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *BillingEventCreate) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *BillingEventCreate) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *BillingEventCreate) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *BillingEventCreate) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *BillingEventCreate) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetPrice

`func (o *BillingEventCreate) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillingEventCreate) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillingEventCreate) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BillingEventCreate) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetTenantId

`func (o *BillingEventCreate) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BillingEventCreate) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BillingEventCreate) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *BillingEventCreate) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


