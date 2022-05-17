# OfferCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Label** | **string** |  | 
**Description** | **string** |  | 
**ProductId** | **string** |  | 
**Version** | **int32** |  | 
**DisplayOrder** | Pointer to **NullableInt32** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**SupportedProperties** | Pointer to **[]string** |  | [optional] 
**SupportedOptions** | Pointer to [**[]NameValue**](NameValue.md) |  | [optional] 
**Approvals** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewOfferCreate

`func NewOfferCreate(name string, label string, description string, productId string, version int32, ) *OfferCreate`

NewOfferCreate instantiates a new OfferCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferCreateWithDefaults

`func NewOfferCreateWithDefaults() *OfferCreate`

NewOfferCreateWithDefaults instantiates a new OfferCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OfferCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OfferCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OfferCreate) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *OfferCreate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *OfferCreate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *OfferCreate) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *OfferCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OfferCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OfferCreate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetProductId

`func (o *OfferCreate) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *OfferCreate) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *OfferCreate) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetVersion

`func (o *OfferCreate) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OfferCreate) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OfferCreate) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetDisplayOrder

`func (o *OfferCreate) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *OfferCreate) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *OfferCreate) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *OfferCreate) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### SetDisplayOrderNil

`func (o *OfferCreate) SetDisplayOrderNil(b bool)`

 SetDisplayOrderNil sets the value for DisplayOrder to be an explicit nil

### UnsetDisplayOrder
`func (o *OfferCreate) UnsetDisplayOrder()`

UnsetDisplayOrder ensures that no value is present for DisplayOrder, not even an explicit nil
### GetImage

`func (o *OfferCreate) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *OfferCreate) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *OfferCreate) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *OfferCreate) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetPrice

`func (o *OfferCreate) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OfferCreate) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OfferCreate) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OfferCreate) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetType

`func (o *OfferCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OfferCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OfferCreate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OfferCreate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSupportedProperties

`func (o *OfferCreate) GetSupportedProperties() []string`

GetSupportedProperties returns the SupportedProperties field if non-nil, zero value otherwise.

### GetSupportedPropertiesOk

`func (o *OfferCreate) GetSupportedPropertiesOk() (*[]string, bool)`

GetSupportedPropertiesOk returns a tuple with the SupportedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedProperties

`func (o *OfferCreate) SetSupportedProperties(v []string)`

SetSupportedProperties sets SupportedProperties field to given value.

### HasSupportedProperties

`func (o *OfferCreate) HasSupportedProperties() bool`

HasSupportedProperties returns a boolean if a field has been set.

### GetSupportedOptions

`func (o *OfferCreate) GetSupportedOptions() []NameValue`

GetSupportedOptions returns the SupportedOptions field if non-nil, zero value otherwise.

### GetSupportedOptionsOk

`func (o *OfferCreate) GetSupportedOptionsOk() (*[]NameValue, bool)`

GetSupportedOptionsOk returns a tuple with the SupportedOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedOptions

`func (o *OfferCreate) SetSupportedOptions(v []NameValue)`

SetSupportedOptions sets SupportedOptions field to given value.

### HasSupportedOptions

`func (o *OfferCreate) HasSupportedOptions() bool`

HasSupportedOptions returns a boolean if a field has been set.

### GetApprovals

`func (o *OfferCreate) GetApprovals() map[string]interface{}`

GetApprovals returns the Approvals field if non-nil, zero value otherwise.

### GetApprovalsOk

`func (o *OfferCreate) GetApprovalsOk() (*map[string]interface{}, bool)`

GetApprovalsOk returns a tuple with the Approvals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovals

`func (o *OfferCreate) SetApprovals(v map[string]interface{})`

SetApprovals sets Approvals field to given value.

### HasApprovals

`func (o *OfferCreate) HasApprovals() bool`

HasApprovals returns a boolean if a field has been set.

### SetApprovalsNil

`func (o *OfferCreate) SetApprovalsNil(b bool)`

 SetApprovalsNil sets the value for Approvals to be an explicit nil

### UnsetApprovals
`func (o *OfferCreate) UnsetApprovals()`

UnsetApprovals ensures that no value is present for Approvals, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


