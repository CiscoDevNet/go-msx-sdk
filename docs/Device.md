# Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**ProviderId** | Pointer to **string** |  | [optional] 
**VulnerabilityState** | Pointer to [**DeviceVulnerabilityState**](DeviceVulnerabilityState.md) |  | [optional] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] 
**ModifiedOn** | Pointer to **NullableTime** |  | [optional] 
**ServiceInstanceId** | Pointer to **NullableString** |  | [optional] 
**SubscriptionId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | **string** |  | 
**ServiceType** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 
**Managed** | **bool** |  | [default to false]
**OnboardType** | **string** |  | 
**OnboardInformation** | Pointer to **map[string]interface{}** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | **string** |  | 
**Model** | **string** |  | 
**Type** | **string** |  | 
**SubType** | Pointer to **NullableString** |  | [optional] 
**SerialKey** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**ComplianceState** | Pointer to [**DeviceComplianceState**](DeviceComplianceState.md) |  | [optional] 

## Methods

### NewDevice

`func NewDevice(tenantId string, managed bool, onboardType string, name string, model string, type_ string, ) *Device`

NewDevice instantiates a new Device object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithDefaults

`func NewDeviceWithDefaults() *Device`

NewDeviceWithDefaults instantiates a new Device object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Device) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Device) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Device) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Device) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *Device) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Device) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Device) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Device) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetProviderId

`func (o *Device) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *Device) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *Device) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *Device) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetVulnerabilityState

`func (o *Device) GetVulnerabilityState() DeviceVulnerabilityState`

GetVulnerabilityState returns the VulnerabilityState field if non-nil, zero value otherwise.

### GetVulnerabilityStateOk

`func (o *Device) GetVulnerabilityStateOk() (*DeviceVulnerabilityState, bool)`

GetVulnerabilityStateOk returns a tuple with the VulnerabilityState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityState

`func (o *Device) SetVulnerabilityState(v DeviceVulnerabilityState)`

SetVulnerabilityState sets VulnerabilityState field to given value.

### HasVulnerabilityState

`func (o *Device) HasVulnerabilityState() bool`

HasVulnerabilityState returns a boolean if a field has been set.

### GetCreatedOn

`func (o *Device) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Device) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Device) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *Device) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetModifiedOn

`func (o *Device) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *Device) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *Device) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *Device) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### SetModifiedOnNil

`func (o *Device) SetModifiedOnNil(b bool)`

 SetModifiedOnNil sets the value for ModifiedOn to be an explicit nil

### UnsetModifiedOn
`func (o *Device) UnsetModifiedOn()`

UnsetModifiedOn ensures that no value is present for ModifiedOn, not even an explicit nil
### GetServiceInstanceId

`func (o *Device) GetServiceInstanceId() string`

GetServiceInstanceId returns the ServiceInstanceId field if non-nil, zero value otherwise.

### GetServiceInstanceIdOk

`func (o *Device) GetServiceInstanceIdOk() (*string, bool)`

GetServiceInstanceIdOk returns a tuple with the ServiceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceInstanceId

`func (o *Device) SetServiceInstanceId(v string)`

SetServiceInstanceId sets ServiceInstanceId field to given value.

### HasServiceInstanceId

`func (o *Device) HasServiceInstanceId() bool`

HasServiceInstanceId returns a boolean if a field has been set.

### SetServiceInstanceIdNil

`func (o *Device) SetServiceInstanceIdNil(b bool)`

 SetServiceInstanceIdNil sets the value for ServiceInstanceId to be an explicit nil

### UnsetServiceInstanceId
`func (o *Device) UnsetServiceInstanceId()`

UnsetServiceInstanceId ensures that no value is present for ServiceInstanceId, not even an explicit nil
### GetSubscriptionId

`func (o *Device) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *Device) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *Device) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *Device) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *Device) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *Device) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetTenantId

`func (o *Device) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Device) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Device) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetServiceType

`func (o *Device) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *Device) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *Device) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *Device) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### SetServiceTypeNil

`func (o *Device) SetServiceTypeNil(b bool)`

 SetServiceTypeNil sets the value for ServiceType to be an explicit nil

### UnsetServiceType
`func (o *Device) UnsetServiceType()`

UnsetServiceType ensures that no value is present for ServiceType, not even an explicit nil
### GetTags

`func (o *Device) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Device) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Device) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Device) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *Device) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *Device) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetManaged

`func (o *Device) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *Device) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *Device) SetManaged(v bool)`

SetManaged sets Managed field to given value.


### GetOnboardType

`func (o *Device) GetOnboardType() string`

GetOnboardType returns the OnboardType field if non-nil, zero value otherwise.

### GetOnboardTypeOk

`func (o *Device) GetOnboardTypeOk() (*string, bool)`

GetOnboardTypeOk returns a tuple with the OnboardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardType

`func (o *Device) SetOnboardType(v string)`

SetOnboardType sets OnboardType field to given value.


### GetOnboardInformation

`func (o *Device) GetOnboardInformation() map[string]interface{}`

GetOnboardInformation returns the OnboardInformation field if non-nil, zero value otherwise.

### GetOnboardInformationOk

`func (o *Device) GetOnboardInformationOk() (*map[string]interface{}, bool)`

GetOnboardInformationOk returns a tuple with the OnboardInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardInformation

`func (o *Device) SetOnboardInformation(v map[string]interface{})`

SetOnboardInformation sets OnboardInformation field to given value.

### HasOnboardInformation

`func (o *Device) HasOnboardInformation() bool`

HasOnboardInformation returns a boolean if a field has been set.

### SetOnboardInformationNil

`func (o *Device) SetOnboardInformationNil(b bool)`

 SetOnboardInformationNil sets the value for OnboardInformation to be an explicit nil

### UnsetOnboardInformation
`func (o *Device) UnsetOnboardInformation()`

UnsetOnboardInformation ensures that no value is present for OnboardInformation, not even an explicit nil
### GetAttributes

`func (o *Device) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Device) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Device) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Device) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *Device) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *Device) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetName

`func (o *Device) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Device) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Device) SetName(v string)`

SetName sets Name field to given value.


### GetModel

`func (o *Device) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Device) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Device) SetModel(v string)`

SetModel sets Model field to given value.


### GetType

`func (o *Device) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Device) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Device) SetType(v string)`

SetType sets Type field to given value.


### GetSubType

`func (o *Device) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *Device) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *Device) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *Device) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### SetSubTypeNil

`func (o *Device) SetSubTypeNil(b bool)`

 SetSubTypeNil sets the value for SubType to be an explicit nil

### UnsetSubType
`func (o *Device) UnsetSubType()`

UnsetSubType ensures that no value is present for SubType, not even an explicit nil
### GetSerialKey

`func (o *Device) GetSerialKey() string`

GetSerialKey returns the SerialKey field if non-nil, zero value otherwise.

### GetSerialKeyOk

`func (o *Device) GetSerialKeyOk() (*string, bool)`

GetSerialKeyOk returns a tuple with the SerialKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialKey

`func (o *Device) SetSerialKey(v string)`

SetSerialKey sets SerialKey field to given value.

### HasSerialKey

`func (o *Device) HasSerialKey() bool`

HasSerialKey returns a boolean if a field has been set.

### SetSerialKeyNil

`func (o *Device) SetSerialKeyNil(b bool)`

 SetSerialKeyNil sets the value for SerialKey to be an explicit nil

### UnsetSerialKey
`func (o *Device) UnsetSerialKey()`

UnsetSerialKey ensures that no value is present for SerialKey, not even an explicit nil
### GetVersion

`func (o *Device) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Device) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Device) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Device) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *Device) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *Device) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetComplianceState

`func (o *Device) GetComplianceState() DeviceComplianceState`

GetComplianceState returns the ComplianceState field if non-nil, zero value otherwise.

### GetComplianceStateOk

`func (o *Device) GetComplianceStateOk() (*DeviceComplianceState, bool)`

GetComplianceStateOk returns a tuple with the ComplianceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceState

`func (o *Device) SetComplianceState(v DeviceComplianceState)`

SetComplianceState sets ComplianceState field to given value.

### HasComplianceState

`func (o *Device) HasComplianceState() bool`

HasComplianceState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


