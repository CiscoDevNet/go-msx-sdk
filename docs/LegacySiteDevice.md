# LegacySiteDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Model** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**DeviceAttributes** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**DeviceOnboarding** | Pointer to [**LegacySiteDeviceOnboard**](LegacySiteDeviceOnboard.md) |  | [optional] 
**Delete** | Pointer to **bool** |  | [optional] 

## Methods

### NewLegacySiteDevice

`func NewLegacySiteDevice(name string, ) *LegacySiteDevice`

NewLegacySiteDevice instantiates a new LegacySiteDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegacySiteDeviceWithDefaults

`func NewLegacySiteDeviceWithDefaults() *LegacySiteDevice`

NewLegacySiteDeviceWithDefaults instantiates a new LegacySiteDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *LegacySiteDevice) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *LegacySiteDevice) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *LegacySiteDevice) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *LegacySiteDevice) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetName

`func (o *LegacySiteDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LegacySiteDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LegacySiteDevice) SetName(v string)`

SetName sets Name field to given value.


### GetModel

`func (o *LegacySiteDevice) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *LegacySiteDevice) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *LegacySiteDevice) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *LegacySiteDevice) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetType

`func (o *LegacySiteDevice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LegacySiteDevice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LegacySiteDevice) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LegacySiteDevice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDeviceAttributes

`func (o *LegacySiteDevice) GetDeviceAttributes() map[string]map[string]interface{}`

GetDeviceAttributes returns the DeviceAttributes field if non-nil, zero value otherwise.

### GetDeviceAttributesOk

`func (o *LegacySiteDevice) GetDeviceAttributesOk() (*map[string]map[string]interface{}, bool)`

GetDeviceAttributesOk returns a tuple with the DeviceAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAttributes

`func (o *LegacySiteDevice) SetDeviceAttributes(v map[string]map[string]interface{})`

SetDeviceAttributes sets DeviceAttributes field to given value.

### HasDeviceAttributes

`func (o *LegacySiteDevice) HasDeviceAttributes() bool`

HasDeviceAttributes returns a boolean if a field has been set.

### GetDeviceOnboarding

`func (o *LegacySiteDevice) GetDeviceOnboarding() LegacySiteDeviceOnboard`

GetDeviceOnboarding returns the DeviceOnboarding field if non-nil, zero value otherwise.

### GetDeviceOnboardingOk

`func (o *LegacySiteDevice) GetDeviceOnboardingOk() (*LegacySiteDeviceOnboard, bool)`

GetDeviceOnboardingOk returns a tuple with the DeviceOnboarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceOnboarding

`func (o *LegacySiteDevice) SetDeviceOnboarding(v LegacySiteDeviceOnboard)`

SetDeviceOnboarding sets DeviceOnboarding field to given value.

### HasDeviceOnboarding

`func (o *LegacySiteDevice) HasDeviceOnboarding() bool`

HasDeviceOnboarding returns a boolean if a field has been set.

### GetDelete

`func (o *LegacySiteDevice) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *LegacySiteDevice) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *LegacySiteDevice) SetDelete(v bool)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *LegacySiteDevice) HasDelete() bool`

HasDelete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


