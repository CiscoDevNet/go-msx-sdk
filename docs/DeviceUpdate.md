# DeviceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceType** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 
**Managed** | Pointer to **bool** |  | [optional] [default to false]
**OnboardType** | **string** |  | 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | **string** |  | 
**Model** | **string** |  | 
**Type** | **string** |  | 
**SubType** | Pointer to **NullableString** |  | [optional] 
**SerialKey** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**ComplianceState** | Pointer to **string** | Valid values: COMPLIANT, NOT_COMPLIANT, APPLICABLE, NOT_APPLICABLE, UNKNOWN | [optional] 
**OnboardInformation** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDeviceUpdate

`func NewDeviceUpdate(onboardType string, name string, model string, type_ string, ) *DeviceUpdate`

NewDeviceUpdate instantiates a new DeviceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceUpdateWithDefaults

`func NewDeviceUpdateWithDefaults() *DeviceUpdate`

NewDeviceUpdateWithDefaults instantiates a new DeviceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceType

`func (o *DeviceUpdate) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *DeviceUpdate) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *DeviceUpdate) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *DeviceUpdate) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### SetServiceTypeNil

`func (o *DeviceUpdate) SetServiceTypeNil(b bool)`

 SetServiceTypeNil sets the value for ServiceType to be an explicit nil

### UnsetServiceType
`func (o *DeviceUpdate) UnsetServiceType()`

UnsetServiceType ensures that no value is present for ServiceType, not even an explicit nil
### GetTags

`func (o *DeviceUpdate) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeviceUpdate) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DeviceUpdate) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DeviceUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *DeviceUpdate) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *DeviceUpdate) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetManaged

`func (o *DeviceUpdate) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *DeviceUpdate) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *DeviceUpdate) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *DeviceUpdate) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetOnboardType

`func (o *DeviceUpdate) GetOnboardType() string`

GetOnboardType returns the OnboardType field if non-nil, zero value otherwise.

### GetOnboardTypeOk

`func (o *DeviceUpdate) GetOnboardTypeOk() (*string, bool)`

GetOnboardTypeOk returns a tuple with the OnboardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardType

`func (o *DeviceUpdate) SetOnboardType(v string)`

SetOnboardType sets OnboardType field to given value.


### GetAttributes

`func (o *DeviceUpdate) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *DeviceUpdate) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *DeviceUpdate) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *DeviceUpdate) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *DeviceUpdate) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *DeviceUpdate) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetName

`func (o *DeviceUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetModel

`func (o *DeviceUpdate) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DeviceUpdate) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DeviceUpdate) SetModel(v string)`

SetModel sets Model field to given value.


### GetType

`func (o *DeviceUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeviceUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeviceUpdate) SetType(v string)`

SetType sets Type field to given value.


### GetSubType

`func (o *DeviceUpdate) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *DeviceUpdate) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *DeviceUpdate) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *DeviceUpdate) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### SetSubTypeNil

`func (o *DeviceUpdate) SetSubTypeNil(b bool)`

 SetSubTypeNil sets the value for SubType to be an explicit nil

### UnsetSubType
`func (o *DeviceUpdate) UnsetSubType()`

UnsetSubType ensures that no value is present for SubType, not even an explicit nil
### GetSerialKey

`func (o *DeviceUpdate) GetSerialKey() string`

GetSerialKey returns the SerialKey field if non-nil, zero value otherwise.

### GetSerialKeyOk

`func (o *DeviceUpdate) GetSerialKeyOk() (*string, bool)`

GetSerialKeyOk returns a tuple with the SerialKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialKey

`func (o *DeviceUpdate) SetSerialKey(v string)`

SetSerialKey sets SerialKey field to given value.

### HasSerialKey

`func (o *DeviceUpdate) HasSerialKey() bool`

HasSerialKey returns a boolean if a field has been set.

### SetSerialKeyNil

`func (o *DeviceUpdate) SetSerialKeyNil(b bool)`

 SetSerialKeyNil sets the value for SerialKey to be an explicit nil

### UnsetSerialKey
`func (o *DeviceUpdate) UnsetSerialKey()`

UnsetSerialKey ensures that no value is present for SerialKey, not even an explicit nil
### GetVersion

`func (o *DeviceUpdate) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeviceUpdate) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeviceUpdate) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DeviceUpdate) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *DeviceUpdate) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *DeviceUpdate) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetComplianceState

`func (o *DeviceUpdate) GetComplianceState() string`

GetComplianceState returns the ComplianceState field if non-nil, zero value otherwise.

### GetComplianceStateOk

`func (o *DeviceUpdate) GetComplianceStateOk() (*string, bool)`

GetComplianceStateOk returns a tuple with the ComplianceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceState

`func (o *DeviceUpdate) SetComplianceState(v string)`

SetComplianceState sets ComplianceState field to given value.

### HasComplianceState

`func (o *DeviceUpdate) HasComplianceState() bool`

HasComplianceState returns a boolean if a field has been set.

### GetOnboardInformation

`func (o *DeviceUpdate) GetOnboardInformation() map[string]interface{}`

GetOnboardInformation returns the OnboardInformation field if non-nil, zero value otherwise.

### GetOnboardInformationOk

`func (o *DeviceUpdate) GetOnboardInformationOk() (*map[string]interface{}, bool)`

GetOnboardInformationOk returns a tuple with the OnboardInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardInformation

`func (o *DeviceUpdate) SetOnboardInformation(v map[string]interface{})`

SetOnboardInformation sets OnboardInformation field to given value.

### HasOnboardInformation

`func (o *DeviceUpdate) HasOnboardInformation() bool`

HasOnboardInformation returns a boolean if a field has been set.

### SetOnboardInformationNil

`func (o *DeviceUpdate) SetOnboardInformationNil(b bool)`

 SetOnboardInformationNil sets the value for OnboardInformation to be an explicit nil

### UnsetOnboardInformation
`func (o *DeviceUpdate) UnsetOnboardInformation()`

UnsetOnboardInformation ensures that no value is present for OnboardInformation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


