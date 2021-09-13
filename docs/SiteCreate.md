# SiteCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**DeviceIds** | Pointer to **[]string** |  | [optional] 
**ServiceIds** | Pointer to **[]string** |  | [optional] 
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

### NewSiteCreate

`func NewSiteCreate(tenantId string, name string, ) *SiteCreate`

NewSiteCreate instantiates a new SiteCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteCreateWithDefaults

`func NewSiteCreateWithDefaults() *SiteCreate`

NewSiteCreateWithDefaults instantiates a new SiteCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *SiteCreate) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SiteCreate) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SiteCreate) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetDeviceIds

`func (o *SiteCreate) GetDeviceIds() []string`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *SiteCreate) GetDeviceIdsOk() (*[]string, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIds

`func (o *SiteCreate) SetDeviceIds(v []string)`

SetDeviceIds sets DeviceIds field to given value.

### HasDeviceIds

`func (o *SiteCreate) HasDeviceIds() bool`

HasDeviceIds returns a boolean if a field has been set.

### GetServiceIds

`func (o *SiteCreate) GetServiceIds() []string`

GetServiceIds returns the ServiceIds field if non-nil, zero value otherwise.

### GetServiceIdsOk

`func (o *SiteCreate) GetServiceIdsOk() (*[]string, bool)`

GetServiceIdsOk returns a tuple with the ServiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceIds

`func (o *SiteCreate) SetServiceIds(v []string)`

SetServiceIds sets ServiceIds field to given value.

### HasServiceIds

`func (o *SiteCreate) HasServiceIds() bool`

HasServiceIds returns a boolean if a field has been set.

### GetParentId

`func (o *SiteCreate) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *SiteCreate) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *SiteCreate) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *SiteCreate) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetName

`func (o *SiteCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiteCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiteCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SiteCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SiteCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SiteCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SiteCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SiteCreate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SiteCreate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *SiteCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SiteCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SiteCreate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SiteCreate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAddress

`func (o *SiteCreate) GetAddress() SiteAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SiteCreate) GetAddressOk() (*SiteAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SiteCreate) SetAddress(v SiteAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SiteCreate) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *SiteCreate) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *SiteCreate) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetContact

`func (o *SiteCreate) GetContact() SiteContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *SiteCreate) GetContactOk() (*SiteContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *SiteCreate) SetContact(v SiteContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *SiteCreate) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *SiteCreate) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *SiteCreate) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetLocation

`func (o *SiteCreate) GetLocation() SiteLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SiteCreate) GetLocationOk() (*SiteLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SiteCreate) SetLocation(v SiteLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SiteCreate) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *SiteCreate) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *SiteCreate) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetImage

`func (o *SiteCreate) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *SiteCreate) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *SiteCreate) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *SiteCreate) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetAttributes

`func (o *SiteCreate) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SiteCreate) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SiteCreate) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SiteCreate) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


