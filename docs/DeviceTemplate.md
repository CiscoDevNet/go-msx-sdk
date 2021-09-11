# DeviceTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**UserId** | Pointer to **string** |  | [optional] [readonly] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**IsLatestVersion** | Pointer to **NullableBool** |  | [optional] 
**ServiceType** | Pointer to **string** |  | [optional] 
**DeviceModels** | Pointer to **[]string** |  | [optional] 
**ConfigContent** | Pointer to **NullableString** |  | [optional] 
**ResourceProvider** | Pointer to **NullableString** |  | [optional] 
**TemplateStandard** | Pointer to **string** |  | [optional] 
**TenantAccess** | Pointer to [**NullableDeviceTemplateAccess**](DeviceTemplateAccess.md) |  | [optional] 
**TemplateParameterValidators** | Pointer to [**[]TemplateParameterValidator**](TemplateParameterValidator.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDeviceTemplate

`func NewDeviceTemplate() *DeviceTemplate`

NewDeviceTemplate instantiates a new DeviceTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceTemplateWithDefaults

`func NewDeviceTemplateWithDefaults() *DeviceTemplate`

NewDeviceTemplateWithDefaults instantiates a new DeviceTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceTemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *DeviceTemplate) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DeviceTemplate) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DeviceTemplate) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *DeviceTemplate) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCreatedOn

`func (o *DeviceTemplate) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *DeviceTemplate) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *DeviceTemplate) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *DeviceTemplate) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetName

`func (o *DeviceTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *DeviceTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeviceTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeviceTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeviceTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *DeviceTemplate) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeviceTemplate) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeviceTemplate) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DeviceTemplate) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetIsLatestVersion

`func (o *DeviceTemplate) GetIsLatestVersion() bool`

GetIsLatestVersion returns the IsLatestVersion field if non-nil, zero value otherwise.

### GetIsLatestVersionOk

`func (o *DeviceTemplate) GetIsLatestVersionOk() (*bool, bool)`

GetIsLatestVersionOk returns a tuple with the IsLatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLatestVersion

`func (o *DeviceTemplate) SetIsLatestVersion(v bool)`

SetIsLatestVersion sets IsLatestVersion field to given value.

### HasIsLatestVersion

`func (o *DeviceTemplate) HasIsLatestVersion() bool`

HasIsLatestVersion returns a boolean if a field has been set.

### SetIsLatestVersionNil

`func (o *DeviceTemplate) SetIsLatestVersionNil(b bool)`

 SetIsLatestVersionNil sets the value for IsLatestVersion to be an explicit nil

### UnsetIsLatestVersion
`func (o *DeviceTemplate) UnsetIsLatestVersion()`

UnsetIsLatestVersion ensures that no value is present for IsLatestVersion, not even an explicit nil
### GetServiceType

`func (o *DeviceTemplate) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *DeviceTemplate) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *DeviceTemplate) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *DeviceTemplate) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetDeviceModels

`func (o *DeviceTemplate) GetDeviceModels() []string`

GetDeviceModels returns the DeviceModels field if non-nil, zero value otherwise.

### GetDeviceModelsOk

`func (o *DeviceTemplate) GetDeviceModelsOk() (*[]string, bool)`

GetDeviceModelsOk returns a tuple with the DeviceModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModels

`func (o *DeviceTemplate) SetDeviceModels(v []string)`

SetDeviceModels sets DeviceModels field to given value.

### HasDeviceModels

`func (o *DeviceTemplate) HasDeviceModels() bool`

HasDeviceModels returns a boolean if a field has been set.

### SetDeviceModelsNil

`func (o *DeviceTemplate) SetDeviceModelsNil(b bool)`

 SetDeviceModelsNil sets the value for DeviceModels to be an explicit nil

### UnsetDeviceModels
`func (o *DeviceTemplate) UnsetDeviceModels()`

UnsetDeviceModels ensures that no value is present for DeviceModels, not even an explicit nil
### GetConfigContent

`func (o *DeviceTemplate) GetConfigContent() string`

GetConfigContent returns the ConfigContent field if non-nil, zero value otherwise.

### GetConfigContentOk

`func (o *DeviceTemplate) GetConfigContentOk() (*string, bool)`

GetConfigContentOk returns a tuple with the ConfigContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContent

`func (o *DeviceTemplate) SetConfigContent(v string)`

SetConfigContent sets ConfigContent field to given value.

### HasConfigContent

`func (o *DeviceTemplate) HasConfigContent() bool`

HasConfigContent returns a boolean if a field has been set.

### SetConfigContentNil

`func (o *DeviceTemplate) SetConfigContentNil(b bool)`

 SetConfigContentNil sets the value for ConfigContent to be an explicit nil

### UnsetConfigContent
`func (o *DeviceTemplate) UnsetConfigContent()`

UnsetConfigContent ensures that no value is present for ConfigContent, not even an explicit nil
### GetResourceProvider

`func (o *DeviceTemplate) GetResourceProvider() string`

GetResourceProvider returns the ResourceProvider field if non-nil, zero value otherwise.

### GetResourceProviderOk

`func (o *DeviceTemplate) GetResourceProviderOk() (*string, bool)`

GetResourceProviderOk returns a tuple with the ResourceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceProvider

`func (o *DeviceTemplate) SetResourceProvider(v string)`

SetResourceProvider sets ResourceProvider field to given value.

### HasResourceProvider

`func (o *DeviceTemplate) HasResourceProvider() bool`

HasResourceProvider returns a boolean if a field has been set.

### SetResourceProviderNil

`func (o *DeviceTemplate) SetResourceProviderNil(b bool)`

 SetResourceProviderNil sets the value for ResourceProvider to be an explicit nil

### UnsetResourceProvider
`func (o *DeviceTemplate) UnsetResourceProvider()`

UnsetResourceProvider ensures that no value is present for ResourceProvider, not even an explicit nil
### GetTemplateStandard

`func (o *DeviceTemplate) GetTemplateStandard() string`

GetTemplateStandard returns the TemplateStandard field if non-nil, zero value otherwise.

### GetTemplateStandardOk

`func (o *DeviceTemplate) GetTemplateStandardOk() (*string, bool)`

GetTemplateStandardOk returns a tuple with the TemplateStandard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateStandard

`func (o *DeviceTemplate) SetTemplateStandard(v string)`

SetTemplateStandard sets TemplateStandard field to given value.

### HasTemplateStandard

`func (o *DeviceTemplate) HasTemplateStandard() bool`

HasTemplateStandard returns a boolean if a field has been set.

### GetTenantAccess

`func (o *DeviceTemplate) GetTenantAccess() DeviceTemplateAccess`

GetTenantAccess returns the TenantAccess field if non-nil, zero value otherwise.

### GetTenantAccessOk

`func (o *DeviceTemplate) GetTenantAccessOk() (*DeviceTemplateAccess, bool)`

GetTenantAccessOk returns a tuple with the TenantAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantAccess

`func (o *DeviceTemplate) SetTenantAccess(v DeviceTemplateAccess)`

SetTenantAccess sets TenantAccess field to given value.

### HasTenantAccess

`func (o *DeviceTemplate) HasTenantAccess() bool`

HasTenantAccess returns a boolean if a field has been set.

### SetTenantAccessNil

`func (o *DeviceTemplate) SetTenantAccessNil(b bool)`

 SetTenantAccessNil sets the value for TenantAccess to be an explicit nil

### UnsetTenantAccess
`func (o *DeviceTemplate) UnsetTenantAccess()`

UnsetTenantAccess ensures that no value is present for TenantAccess, not even an explicit nil
### GetTemplateParameterValidators

`func (o *DeviceTemplate) GetTemplateParameterValidators() []TemplateParameterValidator`

GetTemplateParameterValidators returns the TemplateParameterValidators field if non-nil, zero value otherwise.

### GetTemplateParameterValidatorsOk

`func (o *DeviceTemplate) GetTemplateParameterValidatorsOk() (*[]TemplateParameterValidator, bool)`

GetTemplateParameterValidatorsOk returns a tuple with the TemplateParameterValidators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateParameterValidators

`func (o *DeviceTemplate) SetTemplateParameterValidators(v []TemplateParameterValidator)`

SetTemplateParameterValidators sets TemplateParameterValidators field to given value.

### HasTemplateParameterValidators

`func (o *DeviceTemplate) HasTemplateParameterValidators() bool`

HasTemplateParameterValidators returns a boolean if a field has been set.

### SetTemplateParameterValidatorsNil

`func (o *DeviceTemplate) SetTemplateParameterValidatorsNil(b bool)`

 SetTemplateParameterValidatorsNil sets the value for TemplateParameterValidators to be an explicit nil

### UnsetTemplateParameterValidators
`func (o *DeviceTemplate) UnsetTemplateParameterValidators()`

UnsetTemplateParameterValidators ensures that no value is present for TemplateParameterValidators, not even an explicit nil
### GetTags

`func (o *DeviceTemplate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeviceTemplate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DeviceTemplate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DeviceTemplate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *DeviceTemplate) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *DeviceTemplate) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


