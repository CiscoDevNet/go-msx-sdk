# Site

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**TenantId** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Address** | Pointer to [**NullableSiteAddress**](SiteAddress.md) |  | [optional] 
**Contact** | Pointer to [**NullableSiteContact**](SiteContact.md) |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to **map[string]string** |  | [optional] 
**Devices** | Pointer to [**[]DeviceSummary**](DeviceSummary.md) |  | [optional] 
**ServiceInstanceIds** | Pointer to **[]string** |  | [optional] 
**Location** | Pointer to [**SiteLocation**](SiteLocation.md) |  | [optional] 
**Status** | Pointer to [**SiteStatus**](SiteStatus.md) |  | [optional] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] [readonly] 
**CreatedBy** | Pointer to **string** |  | [optional] [readonly] 
**ModifiedOn** | Pointer to **time.Time** |  | [optional] [readonly] 
**ModifiedBy** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewSite

`func NewSite() *Site`

NewSite instantiates a new Site object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteWithDefaults

`func NewSiteWithDefaults() *Site`

NewSiteWithDefaults instantiates a new Site object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Site) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Site) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Site) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Site) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTenantId

`func (o *Site) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Site) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Site) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *Site) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetParentId

`func (o *Site) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Site) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Site) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Site) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetName

`func (o *Site) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Site) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Site) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Site) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *Site) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Site) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Site) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Site) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *Site) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Site) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Site) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Site) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAddress

`func (o *Site) GetAddress() SiteAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Site) GetAddressOk() (*SiteAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Site) SetAddress(v SiteAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Site) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *Site) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *Site) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetContact

`func (o *Site) GetContact() SiteContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Site) GetContactOk() (*SiteContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Site) SetContact(v SiteContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *Site) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *Site) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *Site) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetImage

`func (o *Site) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Site) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Site) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *Site) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetAttributes

`func (o *Site) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Site) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Site) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Site) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetDevices

`func (o *Site) GetDevices() []DeviceSummary`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *Site) GetDevicesOk() (*[]DeviceSummary, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *Site) SetDevices(v []DeviceSummary)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *Site) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetServiceInstanceIds

`func (o *Site) GetServiceInstanceIds() []string`

GetServiceInstanceIds returns the ServiceInstanceIds field if non-nil, zero value otherwise.

### GetServiceInstanceIdsOk

`func (o *Site) GetServiceInstanceIdsOk() (*[]string, bool)`

GetServiceInstanceIdsOk returns a tuple with the ServiceInstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceInstanceIds

`func (o *Site) SetServiceInstanceIds(v []string)`

SetServiceInstanceIds sets ServiceInstanceIds field to given value.

### HasServiceInstanceIds

`func (o *Site) HasServiceInstanceIds() bool`

HasServiceInstanceIds returns a boolean if a field has been set.

### GetLocation

`func (o *Site) GetLocation() SiteLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Site) GetLocationOk() (*SiteLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Site) SetLocation(v SiteLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Site) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetStatus

`func (o *Site) GetStatus() SiteStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Site) GetStatusOk() (*SiteStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Site) SetStatus(v SiteStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Site) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedOn

`func (o *Site) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Site) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Site) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *Site) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Site) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Site) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Site) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Site) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetModifiedOn

`func (o *Site) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *Site) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *Site) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *Site) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetModifiedBy

`func (o *Site) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Site) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Site) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *Site) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


