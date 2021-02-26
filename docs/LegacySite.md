# LegacySite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | **string** |  | 
**SiteName** | **string** |  | 
**DisplayName** | Pointer to **string** |  | [optional] 
**Address** | [**LegacyAddress**](LegacyAddress.md) |  | 
**Devices** | Pointer to [**[]LegacySiteDevice**](LegacySiteDevice.md) |  | [optional] 
**SiteAttributes** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Delete** | Pointer to **bool** |  | [optional] 
**OperationalStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewLegacySite

`func NewLegacySite(siteId string, siteName string, address LegacyAddress, ) *LegacySite`

NewLegacySite instantiates a new LegacySite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegacySiteWithDefaults

`func NewLegacySiteWithDefaults() *LegacySite`

NewLegacySiteWithDefaults instantiates a new LegacySite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *LegacySite) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *LegacySite) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *LegacySite) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.


### GetSiteName

`func (o *LegacySite) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *LegacySite) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *LegacySite) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.


### GetDisplayName

`func (o *LegacySite) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *LegacySite) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *LegacySite) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *LegacySite) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetAddress

`func (o *LegacySite) GetAddress() LegacyAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *LegacySite) GetAddressOk() (*LegacyAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *LegacySite) SetAddress(v LegacyAddress)`

SetAddress sets Address field to given value.


### GetDevices

`func (o *LegacySite) GetDevices() []LegacySiteDevice`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *LegacySite) GetDevicesOk() (*[]LegacySiteDevice, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *LegacySite) SetDevices(v []LegacySiteDevice)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *LegacySite) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetSiteAttributes

`func (o *LegacySite) GetSiteAttributes() map[string]map[string]interface{}`

GetSiteAttributes returns the SiteAttributes field if non-nil, zero value otherwise.

### GetSiteAttributesOk

`func (o *LegacySite) GetSiteAttributesOk() (*map[string]map[string]interface{}, bool)`

GetSiteAttributesOk returns a tuple with the SiteAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteAttributes

`func (o *LegacySite) SetSiteAttributes(v map[string]map[string]interface{})`

SetSiteAttributes sets SiteAttributes field to given value.

### HasSiteAttributes

`func (o *LegacySite) HasSiteAttributes() bool`

HasSiteAttributes returns a boolean if a field has been set.

### GetDelete

`func (o *LegacySite) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *LegacySite) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *LegacySite) SetDelete(v bool)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *LegacySite) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetOperationalStatus

`func (o *LegacySite) GetOperationalStatus() string`

GetOperationalStatus returns the OperationalStatus field if non-nil, zero value otherwise.

### GetOperationalStatusOk

`func (o *LegacySite) GetOperationalStatusOk() (*string, bool)`

GetOperationalStatusOk returns a tuple with the OperationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStatus

`func (o *LegacySite) SetOperationalStatus(v string)`

SetOperationalStatus sets OperationalStatus field to given value.

### HasOperationalStatus

`func (o *LegacySite) HasOperationalStatus() bool`

HasOperationalStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


