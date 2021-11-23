# OfferUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Label** | **string** |  | 
**Description** | **string** |  | 
**ProductId** | **string** |  | 
**Version** | **int32** |  | 
**DisplayOrder** | **int32** |  | 
**Image** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**SupportedProperties** | Pointer to **[]string** |  | [optional] 
**SupportedOptions** | Pointer to [**[]NameValue**](NameValue.md) |  | [optional] 
**Approvals** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewOfferUpdate

`func NewOfferUpdate(name string, label string, description string, productId string, version int32, displayOrder int32, ) *OfferUpdate`

NewOfferUpdate instantiates a new OfferUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferUpdateWithDefaults

`func NewOfferUpdateWithDefaults() *OfferUpdate`

NewOfferUpdateWithDefaults instantiates a new OfferUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OfferUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OfferUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OfferUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *OfferUpdate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *OfferUpdate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *OfferUpdate) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *OfferUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OfferUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OfferUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetProductId

`func (o *OfferUpdate) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *OfferUpdate) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *OfferUpdate) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetVersion

`func (o *OfferUpdate) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OfferUpdate) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OfferUpdate) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetDisplayOrder

`func (o *OfferUpdate) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *OfferUpdate) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *OfferUpdate) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.


### GetImage

`func (o *OfferUpdate) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *OfferUpdate) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *OfferUpdate) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *OfferUpdate) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetPrice

`func (o *OfferUpdate) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OfferUpdate) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OfferUpdate) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OfferUpdate) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetType

`func (o *OfferUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OfferUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OfferUpdate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OfferUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSupportedProperties

`func (o *OfferUpdate) GetSupportedProperties() []string`

GetSupportedProperties returns the SupportedProperties field if non-nil, zero value otherwise.

### GetSupportedPropertiesOk

`func (o *OfferUpdate) GetSupportedPropertiesOk() (*[]string, bool)`

GetSupportedPropertiesOk returns a tuple with the SupportedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedProperties

`func (o *OfferUpdate) SetSupportedProperties(v []string)`

SetSupportedProperties sets SupportedProperties field to given value.

### HasSupportedProperties

`func (o *OfferUpdate) HasSupportedProperties() bool`

HasSupportedProperties returns a boolean if a field has been set.

### GetSupportedOptions

`func (o *OfferUpdate) GetSupportedOptions() []NameValue`

GetSupportedOptions returns the SupportedOptions field if non-nil, zero value otherwise.

### GetSupportedOptionsOk

`func (o *OfferUpdate) GetSupportedOptionsOk() (*[]NameValue, bool)`

GetSupportedOptionsOk returns a tuple with the SupportedOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedOptions

`func (o *OfferUpdate) SetSupportedOptions(v []NameValue)`

SetSupportedOptions sets SupportedOptions field to given value.

### HasSupportedOptions

`func (o *OfferUpdate) HasSupportedOptions() bool`

HasSupportedOptions returns a boolean if a field has been set.

### GetApprovals

`func (o *OfferUpdate) GetApprovals() map[string]interface{}`

GetApprovals returns the Approvals field if non-nil, zero value otherwise.

### GetApprovalsOk

`func (o *OfferUpdate) GetApprovalsOk() (*map[string]interface{}, bool)`

GetApprovalsOk returns a tuple with the Approvals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovals

`func (o *OfferUpdate) SetApprovals(v map[string]interface{})`

SetApprovals sets Approvals field to given value.

### HasApprovals

`func (o *OfferUpdate) HasApprovals() bool`

HasApprovals returns a boolean if a field has been set.

### SetApprovalsNil

`func (o *OfferUpdate) SetApprovalsNil(b bool)`

 SetApprovalsNil sets the value for Approvals to be an explicit nil

### UnsetApprovals
`func (o *OfferUpdate) UnsetApprovals()`

UnsetApprovals ensures that no value is present for Approvals, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


