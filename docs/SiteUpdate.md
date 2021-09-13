# SiteUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentId** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Address** | Pointer to [**NullableSiteAddress**](SiteAddress.md) |  | [optional] 
**Contact** | Pointer to [**NullableSiteContact**](SiteContact.md) |  | [optional] 
**Location** | Pointer to [**NullableSiteLocation**](SiteLocation.md) |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewSiteUpdate

`func NewSiteUpdate(name string, ) *SiteUpdate`

NewSiteUpdate instantiates a new SiteUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteUpdateWithDefaults

`func NewSiteUpdateWithDefaults() *SiteUpdate`

NewSiteUpdateWithDefaults instantiates a new SiteUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentId

`func (o *SiteUpdate) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *SiteUpdate) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *SiteUpdate) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *SiteUpdate) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetName

`func (o *SiteUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiteUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiteUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SiteUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SiteUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SiteUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SiteUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SiteUpdate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SiteUpdate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *SiteUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SiteUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SiteUpdate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SiteUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAddress

`func (o *SiteUpdate) GetAddress() SiteAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SiteUpdate) GetAddressOk() (*SiteAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SiteUpdate) SetAddress(v SiteAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SiteUpdate) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *SiteUpdate) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *SiteUpdate) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetContact

`func (o *SiteUpdate) GetContact() SiteContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *SiteUpdate) GetContactOk() (*SiteContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *SiteUpdate) SetContact(v SiteContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *SiteUpdate) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *SiteUpdate) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *SiteUpdate) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetLocation

`func (o *SiteUpdate) GetLocation() SiteLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SiteUpdate) GetLocationOk() (*SiteLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SiteUpdate) SetLocation(v SiteLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SiteUpdate) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *SiteUpdate) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *SiteUpdate) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetImage

`func (o *SiteUpdate) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *SiteUpdate) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *SiteUpdate) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *SiteUpdate) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetAttributes

`func (o *SiteUpdate) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SiteUpdate) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SiteUpdate) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SiteUpdate) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


