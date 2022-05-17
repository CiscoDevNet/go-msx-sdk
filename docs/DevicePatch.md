# DevicePatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**SubType** | Pointer to **NullableString** |  | [optional] 
**SerialKey** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**ComplianceState** | Pointer to **string** | Valid values: COMPLIANT, NOT_COMPLIANT, APPLICABLE, NOT_APPLICABLE, UNKNOWN | [optional] 
**OnboardInformation** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDevicePatch

`func NewDevicePatch() *DevicePatch`

NewDevicePatch instantiates a new DevicePatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicePatchWithDefaults

`func NewDevicePatchWithDefaults() *DevicePatch`

NewDevicePatchWithDefaults instantiates a new DevicePatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DevicePatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DevicePatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DevicePatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DevicePatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModel

`func (o *DevicePatch) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DevicePatch) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DevicePatch) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DevicePatch) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetType

`func (o *DevicePatch) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DevicePatch) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DevicePatch) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DevicePatch) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubType

`func (o *DevicePatch) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *DevicePatch) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *DevicePatch) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *DevicePatch) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### SetSubTypeNil

`func (o *DevicePatch) SetSubTypeNil(b bool)`

 SetSubTypeNil sets the value for SubType to be an explicit nil

### UnsetSubType
`func (o *DevicePatch) UnsetSubType()`

UnsetSubType ensures that no value is present for SubType, not even an explicit nil
### GetSerialKey

`func (o *DevicePatch) GetSerialKey() string`

GetSerialKey returns the SerialKey field if non-nil, zero value otherwise.

### GetSerialKeyOk

`func (o *DevicePatch) GetSerialKeyOk() (*string, bool)`

GetSerialKeyOk returns a tuple with the SerialKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialKey

`func (o *DevicePatch) SetSerialKey(v string)`

SetSerialKey sets SerialKey field to given value.

### HasSerialKey

`func (o *DevicePatch) HasSerialKey() bool`

HasSerialKey returns a boolean if a field has been set.

### SetSerialKeyNil

`func (o *DevicePatch) SetSerialKeyNil(b bool)`

 SetSerialKeyNil sets the value for SerialKey to be an explicit nil

### UnsetSerialKey
`func (o *DevicePatch) UnsetSerialKey()`

UnsetSerialKey ensures that no value is present for SerialKey, not even an explicit nil
### GetVersion

`func (o *DevicePatch) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DevicePatch) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DevicePatch) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DevicePatch) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *DevicePatch) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *DevicePatch) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetComplianceState

`func (o *DevicePatch) GetComplianceState() string`

GetComplianceState returns the ComplianceState field if non-nil, zero value otherwise.

### GetComplianceStateOk

`func (o *DevicePatch) GetComplianceStateOk() (*string, bool)`

GetComplianceStateOk returns a tuple with the ComplianceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceState

`func (o *DevicePatch) SetComplianceState(v string)`

SetComplianceState sets ComplianceState field to given value.

### HasComplianceState

`func (o *DevicePatch) HasComplianceState() bool`

HasComplianceState returns a boolean if a field has been set.

### GetOnboardInformation

`func (o *DevicePatch) GetOnboardInformation() map[string]interface{}`

GetOnboardInformation returns the OnboardInformation field if non-nil, zero value otherwise.

### GetOnboardInformationOk

`func (o *DevicePatch) GetOnboardInformationOk() (*map[string]interface{}, bool)`

GetOnboardInformationOk returns a tuple with the OnboardInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardInformation

`func (o *DevicePatch) SetOnboardInformation(v map[string]interface{})`

SetOnboardInformation sets OnboardInformation field to given value.

### HasOnboardInformation

`func (o *DevicePatch) HasOnboardInformation() bool`

HasOnboardInformation returns a boolean if a field has been set.

### SetOnboardInformationNil

`func (o *DevicePatch) SetOnboardInformationNil(b bool)`

 SetOnboardInformationNil sets the value for OnboardInformation to be an explicit nil

### UnsetOnboardInformation
`func (o *DevicePatch) UnsetOnboardInformation()`

UnsetOnboardInformation ensures that no value is present for OnboardInformation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


