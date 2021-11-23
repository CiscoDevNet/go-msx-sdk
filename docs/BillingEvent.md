# BillingEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ModifiedOn** | Pointer to **time.Time** |  | [optional] [readonly] 
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

### NewBillingEvent

`func NewBillingEvent() *BillingEvent`

NewBillingEvent instantiates a new BillingEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingEventWithDefaults

`func NewBillingEventWithDefaults() *BillingEvent`

NewBillingEventWithDefaults instantiates a new BillingEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillingEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedOn

`func (o *BillingEvent) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *BillingEvent) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *BillingEvent) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *BillingEvent) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetCreatedOn

`func (o *BillingEvent) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *BillingEvent) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *BillingEvent) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *BillingEvent) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetName

`func (o *BillingEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillingEvent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *BillingEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillingEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillingEvent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillingEvent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *BillingEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BillingEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BillingEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BillingEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubtype

`func (o *BillingEvent) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *BillingEvent) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *BillingEvent) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *BillingEvent) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetSeverity

`func (o *BillingEvent) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *BillingEvent) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *BillingEvent) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *BillingEvent) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetAction

`func (o *BillingEvent) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BillingEvent) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BillingEvent) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *BillingEvent) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDetails

`func (o *BillingEvent) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *BillingEvent) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *BillingEvent) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *BillingEvent) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *BillingEvent) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *BillingEvent) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetPrice

`func (o *BillingEvent) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillingEvent) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillingEvent) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BillingEvent) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetTenantId

`func (o *BillingEvent) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BillingEvent) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BillingEvent) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *BillingEvent) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


