# DeviceUpdateAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceType** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 
**Managed** | Pointer to **bool** |  | [optional] [default to false]
**OnboardType** | **string** |  | 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDeviceUpdateAllOf

`func NewDeviceUpdateAllOf(onboardType string, ) *DeviceUpdateAllOf`

NewDeviceUpdateAllOf instantiates a new DeviceUpdateAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceUpdateAllOfWithDefaults

`func NewDeviceUpdateAllOfWithDefaults() *DeviceUpdateAllOf`

NewDeviceUpdateAllOfWithDefaults instantiates a new DeviceUpdateAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceType

`func (o *DeviceUpdateAllOf) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *DeviceUpdateAllOf) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *DeviceUpdateAllOf) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *DeviceUpdateAllOf) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### SetServiceTypeNil

`func (o *DeviceUpdateAllOf) SetServiceTypeNil(b bool)`

 SetServiceTypeNil sets the value for ServiceType to be an explicit nil

### UnsetServiceType
`func (o *DeviceUpdateAllOf) UnsetServiceType()`

UnsetServiceType ensures that no value is present for ServiceType, not even an explicit nil
### GetTags

`func (o *DeviceUpdateAllOf) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeviceUpdateAllOf) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DeviceUpdateAllOf) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DeviceUpdateAllOf) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *DeviceUpdateAllOf) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *DeviceUpdateAllOf) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetManaged

`func (o *DeviceUpdateAllOf) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *DeviceUpdateAllOf) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *DeviceUpdateAllOf) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *DeviceUpdateAllOf) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetOnboardType

`func (o *DeviceUpdateAllOf) GetOnboardType() string`

GetOnboardType returns the OnboardType field if non-nil, zero value otherwise.

### GetOnboardTypeOk

`func (o *DeviceUpdateAllOf) GetOnboardTypeOk() (*string, bool)`

GetOnboardTypeOk returns a tuple with the OnboardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardType

`func (o *DeviceUpdateAllOf) SetOnboardType(v string)`

SetOnboardType sets OnboardType field to given value.


### GetAttributes

`func (o *DeviceUpdateAllOf) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *DeviceUpdateAllOf) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *DeviceUpdateAllOf) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *DeviceUpdateAllOf) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *DeviceUpdateAllOf) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *DeviceUpdateAllOf) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


