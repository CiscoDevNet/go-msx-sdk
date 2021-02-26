# DeviceTemplateCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**ServiceType** | **string** |  | 
**DeviceModels** | Pointer to **[]string** |  | [optional] 
**ConfigContent** | **string** |  | 
**ResourceProvider** | **string** |  | 
**TemplateStandard** | Pointer to **string** |  | [optional] 
**TenantAccess** | Pointer to [**NullableDeviceTemplateAccess**](DeviceTemplateAccess.md) |  | [optional] 
**TemplateParameterValidators** | Pointer to [**[]TemplateParameterValidator**](TemplateParameterValidator.md) |  | [optional] 

## Methods

### NewDeviceTemplateCreate

`func NewDeviceTemplateCreate(name string, serviceType string, configContent string, resourceProvider string, ) *DeviceTemplateCreate`

NewDeviceTemplateCreate instantiates a new DeviceTemplateCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceTemplateCreateWithDefaults

`func NewDeviceTemplateCreateWithDefaults() *DeviceTemplateCreate`

NewDeviceTemplateCreateWithDefaults instantiates a new DeviceTemplateCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeviceTemplateCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceTemplateCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceTemplateCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DeviceTemplateCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeviceTemplateCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeviceTemplateCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeviceTemplateCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *DeviceTemplateCreate) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeviceTemplateCreate) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeviceTemplateCreate) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DeviceTemplateCreate) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetServiceType

`func (o *DeviceTemplateCreate) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *DeviceTemplateCreate) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *DeviceTemplateCreate) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.


### GetDeviceModels

`func (o *DeviceTemplateCreate) GetDeviceModels() []string`

GetDeviceModels returns the DeviceModels field if non-nil, zero value otherwise.

### GetDeviceModelsOk

`func (o *DeviceTemplateCreate) GetDeviceModelsOk() (*[]string, bool)`

GetDeviceModelsOk returns a tuple with the DeviceModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModels

`func (o *DeviceTemplateCreate) SetDeviceModels(v []string)`

SetDeviceModels sets DeviceModels field to given value.

### HasDeviceModels

`func (o *DeviceTemplateCreate) HasDeviceModels() bool`

HasDeviceModels returns a boolean if a field has been set.

### GetConfigContent

`func (o *DeviceTemplateCreate) GetConfigContent() string`

GetConfigContent returns the ConfigContent field if non-nil, zero value otherwise.

### GetConfigContentOk

`func (o *DeviceTemplateCreate) GetConfigContentOk() (*string, bool)`

GetConfigContentOk returns a tuple with the ConfigContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContent

`func (o *DeviceTemplateCreate) SetConfigContent(v string)`

SetConfigContent sets ConfigContent field to given value.


### GetResourceProvider

`func (o *DeviceTemplateCreate) GetResourceProvider() string`

GetResourceProvider returns the ResourceProvider field if non-nil, zero value otherwise.

### GetResourceProviderOk

`func (o *DeviceTemplateCreate) GetResourceProviderOk() (*string, bool)`

GetResourceProviderOk returns a tuple with the ResourceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceProvider

`func (o *DeviceTemplateCreate) SetResourceProvider(v string)`

SetResourceProvider sets ResourceProvider field to given value.


### GetTemplateStandard

`func (o *DeviceTemplateCreate) GetTemplateStandard() string`

GetTemplateStandard returns the TemplateStandard field if non-nil, zero value otherwise.

### GetTemplateStandardOk

`func (o *DeviceTemplateCreate) GetTemplateStandardOk() (*string, bool)`

GetTemplateStandardOk returns a tuple with the TemplateStandard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateStandard

`func (o *DeviceTemplateCreate) SetTemplateStandard(v string)`

SetTemplateStandard sets TemplateStandard field to given value.

### HasTemplateStandard

`func (o *DeviceTemplateCreate) HasTemplateStandard() bool`

HasTemplateStandard returns a boolean if a field has been set.

### GetTenantAccess

`func (o *DeviceTemplateCreate) GetTenantAccess() DeviceTemplateAccess`

GetTenantAccess returns the TenantAccess field if non-nil, zero value otherwise.

### GetTenantAccessOk

`func (o *DeviceTemplateCreate) GetTenantAccessOk() (*DeviceTemplateAccess, bool)`

GetTenantAccessOk returns a tuple with the TenantAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantAccess

`func (o *DeviceTemplateCreate) SetTenantAccess(v DeviceTemplateAccess)`

SetTenantAccess sets TenantAccess field to given value.

### HasTenantAccess

`func (o *DeviceTemplateCreate) HasTenantAccess() bool`

HasTenantAccess returns a boolean if a field has been set.

### SetTenantAccessNil

`func (o *DeviceTemplateCreate) SetTenantAccessNil(b bool)`

 SetTenantAccessNil sets the value for TenantAccess to be an explicit nil

### UnsetTenantAccess
`func (o *DeviceTemplateCreate) UnsetTenantAccess()`

UnsetTenantAccess ensures that no value is present for TenantAccess, not even an explicit nil
### GetTemplateParameterValidators

`func (o *DeviceTemplateCreate) GetTemplateParameterValidators() []TemplateParameterValidator`

GetTemplateParameterValidators returns the TemplateParameterValidators field if non-nil, zero value otherwise.

### GetTemplateParameterValidatorsOk

`func (o *DeviceTemplateCreate) GetTemplateParameterValidatorsOk() (*[]TemplateParameterValidator, bool)`

GetTemplateParameterValidatorsOk returns a tuple with the TemplateParameterValidators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateParameterValidators

`func (o *DeviceTemplateCreate) SetTemplateParameterValidators(v []TemplateParameterValidator)`

SetTemplateParameterValidators sets TemplateParameterValidators field to given value.

### HasTemplateParameterValidators

`func (o *DeviceTemplateCreate) HasTemplateParameterValidators() bool`

HasTemplateParameterValidators returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


