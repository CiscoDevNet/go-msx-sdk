# Product

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
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

### NewProduct

`func NewProduct(name string, label string, version int32, description string, image string, ) *Product`

NewProduct instantiates a new Product object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductWithDefaults

`func NewProductWithDefaults() *Product`

NewProductWithDefaults instantiates a new Product object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Product) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Product) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Product) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Product) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Product) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Product) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Product) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *Product) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Product) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Product) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetVersion

`func (o *Product) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Product) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Product) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetDescription

`func (o *Product) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Product) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Product) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetImage

`func (o *Product) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Product) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Product) SetImage(v string)`

SetImage sets Image field to given value.


### GetMultipleInstanceAllowed

`func (o *Product) GetMultipleInstanceAllowed() bool`

GetMultipleInstanceAllowed returns the MultipleInstanceAllowed field if non-nil, zero value otherwise.

### GetMultipleInstanceAllowedOk

`func (o *Product) GetMultipleInstanceAllowedOk() (*bool, bool)`

GetMultipleInstanceAllowedOk returns a tuple with the MultipleInstanceAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleInstanceAllowed

`func (o *Product) SetMultipleInstanceAllowed(v bool)`

SetMultipleInstanceAllowed sets MultipleInstanceAllowed field to given value.

### HasMultipleInstanceAllowed

`func (o *Product) HasMultipleInstanceAllowed() bool`

HasMultipleInstanceAllowed returns a boolean if a field has been set.

### GetPrice

`func (o *Product) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Product) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Product) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Product) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *Product) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *Product) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *Product) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *Product) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetActive

`func (o *Product) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Product) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Product) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Product) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetOrderLimit

`func (o *Product) GetOrderLimit() int32`

GetOrderLimit returns the OrderLimit field if non-nil, zero value otherwise.

### GetOrderLimitOk

`func (o *Product) GetOrderLimitOk() (*int32, bool)`

GetOrderLimitOk returns a tuple with the OrderLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderLimit

`func (o *Product) SetOrderLimit(v int32)`

SetOrderLimit sets OrderLimit field to given value.

### HasOrderLimit

`func (o *Product) HasOrderLimit() bool`

HasOrderLimit returns a boolean if a field has been set.

### GetOptions

`func (o *Product) GetOptions() []ServiceElement`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Product) GetOptionsOk() (*[]ServiceElement, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Product) SetOptions(v []ServiceElement)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Product) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *Product) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *Product) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetProperties

`func (o *Product) GetProperties() []ServiceElement`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Product) GetPropertiesOk() (*[]ServiceElement, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Product) SetProperties(v []ServiceElement)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Product) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *Product) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *Product) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetConfiguration

`func (o *Product) GetConfiguration() map[string]string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *Product) GetConfigurationOk() (*map[string]string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *Product) SetConfiguration(v map[string]string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *Product) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *Product) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *Product) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetIsResource

`func (o *Product) GetIsResource() bool`

GetIsResource returns the IsResource field if non-nil, zero value otherwise.

### GetIsResourceOk

`func (o *Product) GetIsResourceOk() (*bool, bool)`

GetIsResourceOk returns a tuple with the IsResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsResource

`func (o *Product) SetIsResource(v bool)`

SetIsResource sets IsResource field to given value.

### HasIsResource

`func (o *Product) HasIsResource() bool`

HasIsResource returns a boolean if a field has been set.

### SetIsResourceNil

`func (o *Product) SetIsResourceNil(b bool)`

 SetIsResourceNil sets the value for IsResource to be an explicit nil

### UnsetIsResource
`func (o *Product) UnsetIsResource()`

UnsetIsResource ensures that no value is present for IsResource, not even an explicit nil
### GetHasChildren

`func (o *Product) GetHasChildren() bool`

GetHasChildren returns the HasChildren field if non-nil, zero value otherwise.

### GetHasChildrenOk

`func (o *Product) GetHasChildrenOk() (*bool, bool)`

GetHasChildrenOk returns a tuple with the HasChildren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasChildren

`func (o *Product) SetHasChildren(v bool)`

SetHasChildren sets HasChildren field to given value.

### HasHasChildren

`func (o *Product) HasHasChildren() bool`

HasHasChildren returns a boolean if a field has been set.

### SetHasChildrenNil

`func (o *Product) SetHasChildrenNil(b bool)`

 SetHasChildrenNil sets the value for HasChildren to be an explicit nil

### UnsetHasChildren
`func (o *Product) UnsetHasChildren()`

UnsetHasChildren ensures that no value is present for HasChildren, not even an explicit nil
### GetParentId

`func (o *Product) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Product) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Product) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Product) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetServiceExtensions

`func (o *Product) GetServiceExtensions() []NSOConfigDataXPath`

GetServiceExtensions returns the ServiceExtensions field if non-nil, zero value otherwise.

### GetServiceExtensionsOk

`func (o *Product) GetServiceExtensionsOk() (*[]NSOConfigDataXPath, bool)`

GetServiceExtensionsOk returns a tuple with the ServiceExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceExtensions

`func (o *Product) SetServiceExtensions(v []NSOConfigDataXPath)`

SetServiceExtensions sets ServiceExtensions field to given value.

### HasServiceExtensions

`func (o *Product) HasServiceExtensions() bool`

HasServiceExtensions returns a boolean if a field has been set.

### GetServiceConfigQueryRootXPaths

`func (o *Product) GetServiceConfigQueryRootXPaths() []string`

GetServiceConfigQueryRootXPaths returns the ServiceConfigQueryRootXPaths field if non-nil, zero value otherwise.

### GetServiceConfigQueryRootXPathsOk

`func (o *Product) GetServiceConfigQueryRootXPathsOk() (*[]string, bool)`

GetServiceConfigQueryRootXPathsOk returns a tuple with the ServiceConfigQueryRootXPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceConfigQueryRootXPaths

`func (o *Product) SetServiceConfigQueryRootXPaths(v []string)`

SetServiceConfigQueryRootXPaths sets ServiceConfigQueryRootXPaths field to given value.

### HasServiceConfigQueryRootXPaths

`func (o *Product) HasServiceConfigQueryRootXPaths() bool`

HasServiceConfigQueryRootXPaths returns a boolean if a field has been set.

### GetUiConfig

`func (o *Product) GetUiConfig() ServiceUIConfig`

GetUiConfig returns the UiConfig field if non-nil, zero value otherwise.

### GetUiConfigOk

`func (o *Product) GetUiConfigOk() (*ServiceUIConfig, bool)`

GetUiConfigOk returns a tuple with the UiConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiConfig

`func (o *Product) SetUiConfig(v ServiceUIConfig)`

SetUiConfig sets UiConfig field to given value.

### HasUiConfig

`func (o *Product) HasUiConfig() bool`

HasUiConfig returns a boolean if a field has been set.

### GetSlmUiConfig

`func (o *Product) GetSlmUiConfig() ServiceSLMUIConfig`

GetSlmUiConfig returns the SlmUiConfig field if non-nil, zero value otherwise.

### GetSlmUiConfigOk

`func (o *Product) GetSlmUiConfigOk() (*ServiceSLMUIConfig, bool)`

GetSlmUiConfigOk returns a tuple with the SlmUiConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlmUiConfig

`func (o *Product) SetSlmUiConfig(v ServiceSLMUIConfig)`

SetSlmUiConfig sets SlmUiConfig field to given value.

### HasSlmUiConfig

`func (o *Product) HasSlmUiConfig() bool`

HasSlmUiConfig returns a boolean if a field has been set.

### SetSlmUiConfigNil

`func (o *Product) SetSlmUiConfigNil(b bool)`

 SetSlmUiConfigNil sets the value for SlmUiConfig to be an explicit nil

### UnsetSlmUiConfig
`func (o *Product) UnsetSlmUiConfig()`

UnsetSlmUiConfig ensures that no value is present for SlmUiConfig, not even an explicit nil
### GetExternalId

`func (o *Product) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Product) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Product) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Product) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *Product) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *Product) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetTags

`func (o *Product) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Product) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Product) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Product) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *Product) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *Product) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


