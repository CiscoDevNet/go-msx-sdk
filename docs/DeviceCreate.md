# DeviceCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceInstanceId** | Pointer to **NullableString** |  | [optional] 
**SubscriptionId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | **string** |  | 
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

### NewDeviceCreate

`func NewDeviceCreate(tenantId string, onboardType string, name string, model string, type_ string, ) *DeviceCreate`

NewDeviceCreate instantiates a new DeviceCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceCreateWithDefaults

`func NewDeviceCreateWithDefaults() *DeviceCreate`

NewDeviceCreateWithDefaults instantiates a new DeviceCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceInstanceId

`func (o *DeviceCreate) GetServiceInstanceId() string`

GetServiceInstanceId returns the ServiceInstanceId field if non-nil, zero value otherwise.

### GetServiceInstanceIdOk

`func (o *DeviceCreate) GetServiceInstanceIdOk() (*string, bool)`

GetServiceInstanceIdOk returns a tuple with the ServiceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceInstanceId

`func (o *DeviceCreate) SetServiceInstanceId(v string)`

SetServiceInstanceId sets ServiceInstanceId field to given value.

### HasServiceInstanceId

`func (o *DeviceCreate) HasServiceInstanceId() bool`

HasServiceInstanceId returns a boolean if a field has been set.

### SetServiceInstanceIdNil

`func (o *DeviceCreate) SetServiceInstanceIdNil(b bool)`

 SetServiceInstanceIdNil sets the value for ServiceInstanceId to be an explicit nil

### UnsetServiceInstanceId
`func (o *DeviceCreate) UnsetServiceInstanceId()`

UnsetServiceInstanceId ensures that no value is present for ServiceInstanceId, not even an explicit nil
### GetSubscriptionId

`func (o *DeviceCreate) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DeviceCreate) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DeviceCreate) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *DeviceCreate) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *DeviceCreate) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *DeviceCreate) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetTenantId

`func (o *DeviceCreate) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DeviceCreate) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DeviceCreate) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetServiceType

`func (o *DeviceCreate) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *DeviceCreate) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *DeviceCreate) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *DeviceCreate) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### SetServiceTypeNil

`func (o *DeviceCreate) SetServiceTypeNil(b bool)`

 SetServiceTypeNil sets the value for ServiceType to be an explicit nil

### UnsetServiceType
`func (o *DeviceCreate) UnsetServiceType()`

UnsetServiceType ensures that no value is present for ServiceType, not even an explicit nil
### GetTags

`func (o *DeviceCreate) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeviceCreate) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DeviceCreate) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DeviceCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *DeviceCreate) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *DeviceCreate) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetManaged

`func (o *DeviceCreate) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *DeviceCreate) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *DeviceCreate) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *DeviceCreate) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetOnboardType

`func (o *DeviceCreate) GetOnboardType() string`

GetOnboardType returns the OnboardType field if non-nil, zero value otherwise.

### GetOnboardTypeOk

`func (o *DeviceCreate) GetOnboardTypeOk() (*string, bool)`

GetOnboardTypeOk returns a tuple with the OnboardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardType

`func (o *DeviceCreate) SetOnboardType(v string)`

SetOnboardType sets OnboardType field to given value.


### GetAttributes

`func (o *DeviceCreate) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *DeviceCreate) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *DeviceCreate) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *DeviceCreate) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *DeviceCreate) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *DeviceCreate) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetName

`func (o *DeviceCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceCreate) SetName(v string)`

SetName sets Name field to given value.


### GetModel

`func (o *DeviceCreate) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DeviceCreate) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DeviceCreate) SetModel(v string)`

SetModel sets Model field to given value.


### GetType

`func (o *DeviceCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeviceCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeviceCreate) SetType(v string)`

SetType sets Type field to given value.


### GetSubType

`func (o *DeviceCreate) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *DeviceCreate) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *DeviceCreate) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *DeviceCreate) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### SetSubTypeNil

`func (o *DeviceCreate) SetSubTypeNil(b bool)`

 SetSubTypeNil sets the value for SubType to be an explicit nil

### UnsetSubType
`func (o *DeviceCreate) UnsetSubType()`

UnsetSubType ensures that no value is present for SubType, not even an explicit nil
### GetSerialKey

`func (o *DeviceCreate) GetSerialKey() string`

GetSerialKey returns the SerialKey field if non-nil, zero value otherwise.

### GetSerialKeyOk

`func (o *DeviceCreate) GetSerialKeyOk() (*string, bool)`

GetSerialKeyOk returns a tuple with the SerialKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialKey

`func (o *DeviceCreate) SetSerialKey(v string)`

SetSerialKey sets SerialKey field to given value.

### HasSerialKey

`func (o *DeviceCreate) HasSerialKey() bool`

HasSerialKey returns a boolean if a field has been set.

### SetSerialKeyNil

`func (o *DeviceCreate) SetSerialKeyNil(b bool)`

 SetSerialKeyNil sets the value for SerialKey to be an explicit nil

### UnsetSerialKey
`func (o *DeviceCreate) UnsetSerialKey()`

UnsetSerialKey ensures that no value is present for SerialKey, not even an explicit nil
### GetVersion

`func (o *DeviceCreate) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeviceCreate) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeviceCreate) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DeviceCreate) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *DeviceCreate) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *DeviceCreate) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetComplianceState

`func (o *DeviceCreate) GetComplianceState() string`

GetComplianceState returns the ComplianceState field if non-nil, zero value otherwise.

### GetComplianceStateOk

`func (o *DeviceCreate) GetComplianceStateOk() (*string, bool)`

GetComplianceStateOk returns a tuple with the ComplianceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceState

`func (o *DeviceCreate) SetComplianceState(v string)`

SetComplianceState sets ComplianceState field to given value.

### HasComplianceState

`func (o *DeviceCreate) HasComplianceState() bool`

HasComplianceState returns a boolean if a field has been set.

### GetOnboardInformation

`func (o *DeviceCreate) GetOnboardInformation() map[string]interface{}`

GetOnboardInformation returns the OnboardInformation field if non-nil, zero value otherwise.

### GetOnboardInformationOk

`func (o *DeviceCreate) GetOnboardInformationOk() (*map[string]interface{}, bool)`

GetOnboardInformationOk returns a tuple with the OnboardInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardInformation

`func (o *DeviceCreate) SetOnboardInformation(v map[string]interface{})`

SetOnboardInformation sets OnboardInformation field to given value.

### HasOnboardInformation

`func (o *DeviceCreate) HasOnboardInformation() bool`

HasOnboardInformation returns a boolean if a field has been set.

### SetOnboardInformationNil

`func (o *DeviceCreate) SetOnboardInformationNil(b bool)`

 SetOnboardInformationNil sets the value for OnboardInformation to be an explicit nil

### UnsetOnboardInformation
`func (o *DeviceCreate) UnsetOnboardInformation()`

UnsetOnboardInformation ensures that no value is present for OnboardInformation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


