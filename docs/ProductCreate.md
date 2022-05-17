# ProductCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Label** | **string** |  | 
**Version** | **int32** |  | 
**Description** | **string** |  | 
**Image** | **string** |  | 
**MultipleInstanceAllowed** | Pointer to **bool** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**DisplayOrder** | Pointer to **int32** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**OrderLimit** | Pointer to **int32** |  | [optional] 
**Options** | Pointer to [**[]ServiceElement**](ServiceElement.md) |  | [optional] 
**Properties** | Pointer to [**[]ServiceElement**](ServiceElement.md) |  | [optional] 
**Configuration** | Pointer to **map[string]string** |  | [optional] 
**IsResource** | Pointer to **NullableBool** |  | [optional] 
**HasChildren** | Pointer to **NullableBool** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**ServiceExtensions** | Pointer to [**[]NSOConfigDataXPath**](NSOConfigDataXPath.md) |  | [optional] 
**ServiceConfigQueryRootXPaths** | Pointer to **[]string** |  | [optional] 
**UiConfig** | Pointer to [**ServiceUIConfig**](ServiceUIConfig.md) |  | [optional] 
**SlmUiConfig** | Pointer to [**NullableServiceSLMUIConfig**](ServiceSLMUIConfig.md) |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewProductCreate

`func NewProductCreate(name string, label string, version int32, description string, image string, ) *ProductCreate`

NewProductCreate instantiates a new ProductCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductCreateWithDefaults

`func NewProductCreateWithDefaults() *ProductCreate`

NewProductCreateWithDefaults instantiates a new ProductCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProductCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductCreate) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *ProductCreate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ProductCreate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ProductCreate) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetVersion

`func (o *ProductCreate) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ProductCreate) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ProductCreate) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetDescription

`func (o *ProductCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductCreate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetImage

`func (o *ProductCreate) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ProductCreate) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ProductCreate) SetImage(v string)`

SetImage sets Image field to given value.


### GetMultipleInstanceAllowed

`func (o *ProductCreate) GetMultipleInstanceAllowed() bool`

GetMultipleInstanceAllowed returns the MultipleInstanceAllowed field if non-nil, zero value otherwise.

### GetMultipleInstanceAllowedOk

`func (o *ProductCreate) GetMultipleInstanceAllowedOk() (*bool, bool)`

GetMultipleInstanceAllowedOk returns a tuple with the MultipleInstanceAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleInstanceAllowed

`func (o *ProductCreate) SetMultipleInstanceAllowed(v bool)`

SetMultipleInstanceAllowed sets MultipleInstanceAllowed field to given value.

### HasMultipleInstanceAllowed

`func (o *ProductCreate) HasMultipleInstanceAllowed() bool`

HasMultipleInstanceAllowed returns a boolean if a field has been set.

### GetPrice

`func (o *ProductCreate) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProductCreate) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProductCreate) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ProductCreate) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *ProductCreate) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *ProductCreate) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *ProductCreate) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *ProductCreate) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetActive

`func (o *ProductCreate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ProductCreate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ProductCreate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ProductCreate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetOrderLimit

`func (o *ProductCreate) GetOrderLimit() int32`

GetOrderLimit returns the OrderLimit field if non-nil, zero value otherwise.

### GetOrderLimitOk

`func (o *ProductCreate) GetOrderLimitOk() (*int32, bool)`

GetOrderLimitOk returns a tuple with the OrderLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderLimit

`func (o *ProductCreate) SetOrderLimit(v int32)`

SetOrderLimit sets OrderLimit field to given value.

### HasOrderLimit

`func (o *ProductCreate) HasOrderLimit() bool`

HasOrderLimit returns a boolean if a field has been set.

### GetOptions

`func (o *ProductCreate) GetOptions() []ServiceElement`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ProductCreate) GetOptionsOk() (*[]ServiceElement, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ProductCreate) SetOptions(v []ServiceElement)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ProductCreate) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *ProductCreate) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *ProductCreate) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetProperties

`func (o *ProductCreate) GetProperties() []ServiceElement`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ProductCreate) GetPropertiesOk() (*[]ServiceElement, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ProductCreate) SetProperties(v []ServiceElement)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ProductCreate) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *ProductCreate) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *ProductCreate) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetConfiguration

`func (o *ProductCreate) GetConfiguration() map[string]string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ProductCreate) GetConfigurationOk() (*map[string]string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ProductCreate) SetConfiguration(v map[string]string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ProductCreate) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *ProductCreate) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *ProductCreate) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetIsResource

`func (o *ProductCreate) GetIsResource() bool`

GetIsResource returns the IsResource field if non-nil, zero value otherwise.

### GetIsResourceOk

`func (o *ProductCreate) GetIsResourceOk() (*bool, bool)`

GetIsResourceOk returns a tuple with the IsResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsResource

`func (o *ProductCreate) SetIsResource(v bool)`

SetIsResource sets IsResource field to given value.

### HasIsResource

`func (o *ProductCreate) HasIsResource() bool`

HasIsResource returns a boolean if a field has been set.

### SetIsResourceNil

`func (o *ProductCreate) SetIsResourceNil(b bool)`

 SetIsResourceNil sets the value for IsResource to be an explicit nil

### UnsetIsResource
`func (o *ProductCreate) UnsetIsResource()`

UnsetIsResource ensures that no value is present for IsResource, not even an explicit nil
### GetHasChildren

`func (o *ProductCreate) GetHasChildren() bool`

GetHasChildren returns the HasChildren field if non-nil, zero value otherwise.

### GetHasChildrenOk

`func (o *ProductCreate) GetHasChildrenOk() (*bool, bool)`

GetHasChildrenOk returns a tuple with the HasChildren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasChildren

`func (o *ProductCreate) SetHasChildren(v bool)`

SetHasChildren sets HasChildren field to given value.

### HasHasChildren

`func (o *ProductCreate) HasHasChildren() bool`

HasHasChildren returns a boolean if a field has been set.

### SetHasChildrenNil

`func (o *ProductCreate) SetHasChildrenNil(b bool)`

 SetHasChildrenNil sets the value for HasChildren to be an explicit nil

### UnsetHasChildren
`func (o *ProductCreate) UnsetHasChildren()`

UnsetHasChildren ensures that no value is present for HasChildren, not even an explicit nil
### GetParentId

`func (o *ProductCreate) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ProductCreate) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ProductCreate) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ProductCreate) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetServiceExtensions

`func (o *ProductCreate) GetServiceExtensions() []NSOConfigDataXPath`

GetServiceExtensions returns the ServiceExtensions field if non-nil, zero value otherwise.

### GetServiceExtensionsOk

`func (o *ProductCreate) GetServiceExtensionsOk() (*[]NSOConfigDataXPath, bool)`

GetServiceExtensionsOk returns a tuple with the ServiceExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceExtensions

`func (o *ProductCreate) SetServiceExtensions(v []NSOConfigDataXPath)`

SetServiceExtensions sets ServiceExtensions field to given value.

### HasServiceExtensions

`func (o *ProductCreate) HasServiceExtensions() bool`

HasServiceExtensions returns a boolean if a field has been set.

### GetServiceConfigQueryRootXPaths

`func (o *ProductCreate) GetServiceConfigQueryRootXPaths() []string`

GetServiceConfigQueryRootXPaths returns the ServiceConfigQueryRootXPaths field if non-nil, zero value otherwise.

### GetServiceConfigQueryRootXPathsOk

`func (o *ProductCreate) GetServiceConfigQueryRootXPathsOk() (*[]string, bool)`

GetServiceConfigQueryRootXPathsOk returns a tuple with the ServiceConfigQueryRootXPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceConfigQueryRootXPaths

`func (o *ProductCreate) SetServiceConfigQueryRootXPaths(v []string)`

SetServiceConfigQueryRootXPaths sets ServiceConfigQueryRootXPaths field to given value.

### HasServiceConfigQueryRootXPaths

`func (o *ProductCreate) HasServiceConfigQueryRootXPaths() bool`

HasServiceConfigQueryRootXPaths returns a boolean if a field has been set.

### GetUiConfig

`func (o *ProductCreate) GetUiConfig() ServiceUIConfig`

GetUiConfig returns the UiConfig field if non-nil, zero value otherwise.

### GetUiConfigOk

`func (o *ProductCreate) GetUiConfigOk() (*ServiceUIConfig, bool)`

GetUiConfigOk returns a tuple with the UiConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiConfig

`func (o *ProductCreate) SetUiConfig(v ServiceUIConfig)`

SetUiConfig sets UiConfig field to given value.

### HasUiConfig

`func (o *ProductCreate) HasUiConfig() bool`

HasUiConfig returns a boolean if a field has been set.

### GetSlmUiConfig

`func (o *ProductCreate) GetSlmUiConfig() ServiceSLMUIConfig`

GetSlmUiConfig returns the SlmUiConfig field if non-nil, zero value otherwise.

### GetSlmUiConfigOk

`func (o *ProductCreate) GetSlmUiConfigOk() (*ServiceSLMUIConfig, bool)`

GetSlmUiConfigOk returns a tuple with the SlmUiConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlmUiConfig

`func (o *ProductCreate) SetSlmUiConfig(v ServiceSLMUIConfig)`

SetSlmUiConfig sets SlmUiConfig field to given value.

### HasSlmUiConfig

`func (o *ProductCreate) HasSlmUiConfig() bool`

HasSlmUiConfig returns a boolean if a field has been set.

### SetSlmUiConfigNil

`func (o *ProductCreate) SetSlmUiConfigNil(b bool)`

 SetSlmUiConfigNil sets the value for SlmUiConfig to be an explicit nil

### UnsetSlmUiConfig
`func (o *ProductCreate) UnsetSlmUiConfig()`

UnsetSlmUiConfig ensures that no value is present for SlmUiConfig, not even an explicit nil
### GetExternalId

`func (o *ProductCreate) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ProductCreate) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ProductCreate) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ProductCreate) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *ProductCreate) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *ProductCreate) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetTags

`func (o *ProductCreate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProductCreate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProductCreate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProductCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ProductCreate) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ProductCreate) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


