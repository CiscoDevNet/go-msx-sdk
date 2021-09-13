# Offer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Label** | **string** |  | 
**Description** | **string** |  | 
**ProductId** | **string** |  | 
**Version** | **int32** |  | 
**DisplayOrder** | **int32** |  | 
**Image** | **string** |  | 
**Price** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**SupportedProperties** | **[]string** |  | 
**SupportedOptions** | [**[]NameValue**](NameValue.md) |  | 
**Approvals** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewOffer

`func NewOffer(name string, label string, description string, productId string, version int32, displayOrder int32, image string, supportedProperties []string, supportedOptions []NameValue, ) *Offer`

NewOffer instantiates a new Offer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferWithDefaults

`func NewOfferWithDefaults() *Offer`

NewOfferWithDefaults instantiates a new Offer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Offer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Offer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Offer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Offer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Offer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Offer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Offer) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *Offer) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Offer) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Offer) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *Offer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Offer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Offer) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetProductId

`func (o *Offer) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *Offer) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *Offer) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetVersion

`func (o *Offer) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Offer) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Offer) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetDisplayOrder

`func (o *Offer) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *Offer) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *Offer) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.


### GetImage

`func (o *Offer) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Offer) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Offer) SetImage(v string)`

SetImage sets Image field to given value.


### GetPrice

`func (o *Offer) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Offer) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Offer) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Offer) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetType

`func (o *Offer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Offer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Offer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Offer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSupportedProperties

`func (o *Offer) GetSupportedProperties() []string`

GetSupportedProperties returns the SupportedProperties field if non-nil, zero value otherwise.

### GetSupportedPropertiesOk

`func (o *Offer) GetSupportedPropertiesOk() (*[]string, bool)`

GetSupportedPropertiesOk returns a tuple with the SupportedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedProperties

`func (o *Offer) SetSupportedProperties(v []string)`

SetSupportedProperties sets SupportedProperties field to given value.


### GetSupportedOptions

`func (o *Offer) GetSupportedOptions() []NameValue`

GetSupportedOptions returns the SupportedOptions field if non-nil, zero value otherwise.

### GetSupportedOptionsOk

`func (o *Offer) GetSupportedOptionsOk() (*[]NameValue, bool)`

GetSupportedOptionsOk returns a tuple with the SupportedOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedOptions

`func (o *Offer) SetSupportedOptions(v []NameValue)`

SetSupportedOptions sets SupportedOptions field to given value.


### GetApprovals

`func (o *Offer) GetApprovals() map[string]interface{}`

GetApprovals returns the Approvals field if non-nil, zero value otherwise.

### GetApprovalsOk

`func (o *Offer) GetApprovalsOk() (*map[string]interface{}, bool)`

GetApprovalsOk returns a tuple with the Approvals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovals

`func (o *Offer) SetApprovals(v map[string]interface{})`

SetApprovals sets Approvals field to given value.

### HasApprovals

`func (o *Offer) HasApprovals() bool`

HasApprovals returns a boolean if a field has been set.

### SetApprovalsNil

`func (o *Offer) SetApprovalsNil(b bool)`

 SetApprovalsNil sets the value for Approvals to be an explicit nil

### UnsetApprovals
`func (o *Offer) UnsetApprovals()`

UnsetApprovals ensures that no value is present for Approvals, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


