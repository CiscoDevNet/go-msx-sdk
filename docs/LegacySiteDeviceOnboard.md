# LegacySiteDeviceOnboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceInstanceId** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**DeviceName** | Pointer to **string** |  | [optional] 
**Managed** | Pointer to **bool** |  | [optional] [default to true]
**DeviceModel** | Pointer to **string** |  | [optional] 
**DeviceOnboardingType** | Pointer to **string** |  | [optional] 
**DeviceOnboardInformation** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewLegacySiteDeviceOnboard

`func NewLegacySiteDeviceOnboard() *LegacySiteDeviceOnboard`

NewLegacySiteDeviceOnboard instantiates a new LegacySiteDeviceOnboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegacySiteDeviceOnboardWithDefaults

`func NewLegacySiteDeviceOnboardWithDefaults() *LegacySiteDeviceOnboard`

NewLegacySiteDeviceOnboardWithDefaults instantiates a new LegacySiteDeviceOnboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceInstanceId

`func (o *LegacySiteDeviceOnboard) GetDeviceInstanceId() string`

GetDeviceInstanceId returns the DeviceInstanceId field if non-nil, zero value otherwise.

### GetDeviceInstanceIdOk

`func (o *LegacySiteDeviceOnboard) GetDeviceInstanceIdOk() (*string, bool)`

GetDeviceInstanceIdOk returns a tuple with the DeviceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceInstanceId

`func (o *LegacySiteDeviceOnboard) SetDeviceInstanceId(v string)`

SetDeviceInstanceId sets DeviceInstanceId field to given value.

### HasDeviceInstanceId

`func (o *LegacySiteDeviceOnboard) HasDeviceInstanceId() bool`

HasDeviceInstanceId returns a boolean if a field has been set.

### GetTenantId

`func (o *LegacySiteDeviceOnboard) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *LegacySiteDeviceOnboard) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *LegacySiteDeviceOnboard) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *LegacySiteDeviceOnboard) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetDeviceName

`func (o *LegacySiteDeviceOnboard) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *LegacySiteDeviceOnboard) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *LegacySiteDeviceOnboard) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *LegacySiteDeviceOnboard) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetManaged

`func (o *LegacySiteDeviceOnboard) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *LegacySiteDeviceOnboard) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *LegacySiteDeviceOnboard) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *LegacySiteDeviceOnboard) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetDeviceModel

`func (o *LegacySiteDeviceOnboard) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *LegacySiteDeviceOnboard) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *LegacySiteDeviceOnboard) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *LegacySiteDeviceOnboard) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetDeviceOnboardingType

`func (o *LegacySiteDeviceOnboard) GetDeviceOnboardingType() string`

GetDeviceOnboardingType returns the DeviceOnboardingType field if non-nil, zero value otherwise.

### GetDeviceOnboardingTypeOk

`func (o *LegacySiteDeviceOnboard) GetDeviceOnboardingTypeOk() (*string, bool)`

GetDeviceOnboardingTypeOk returns a tuple with the DeviceOnboardingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceOnboardingType

`func (o *LegacySiteDeviceOnboard) SetDeviceOnboardingType(v string)`

SetDeviceOnboardingType sets DeviceOnboardingType field to given value.

### HasDeviceOnboardingType

`func (o *LegacySiteDeviceOnboard) HasDeviceOnboardingType() bool`

HasDeviceOnboardingType returns a boolean if a field has been set.

### GetDeviceOnboardInformation

`func (o *LegacySiteDeviceOnboard) GetDeviceOnboardInformation() map[string]interface{}`

GetDeviceOnboardInformation returns the DeviceOnboardInformation field if non-nil, zero value otherwise.

### GetDeviceOnboardInformationOk

`func (o *LegacySiteDeviceOnboard) GetDeviceOnboardInformationOk() (*map[string]interface{}, bool)`

GetDeviceOnboardInformationOk returns a tuple with the DeviceOnboardInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceOnboardInformation

`func (o *LegacySiteDeviceOnboard) SetDeviceOnboardInformation(v map[string]interface{})`

SetDeviceOnboardInformation sets DeviceOnboardInformation field to given value.

### HasDeviceOnboardInformation

`func (o *LegacySiteDeviceOnboard) HasDeviceOnboardInformation() bool`

HasDeviceOnboardInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


