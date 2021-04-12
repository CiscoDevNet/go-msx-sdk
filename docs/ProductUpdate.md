# ProductUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Label** | **string** |  | 
**Version** | **int32** |  | 
**Description** | **string** |  | 
**Image** | **string** |  | 
**MultipleInstanceAllowed** | Pointer to **bool** |  | [optional] 
**Price** | **string** |  | 
**DisplayOrder** | Pointer to **int32** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**OrderLimit** | Pointer to **int32** |  | [optional] 
**Options** | [**[]ServiceElement**](ServiceElement.md) |  | 
**Properties** | [**[]ServiceElement**](ServiceElement.md) |  | 
**Configuration** | **map[string]string** |  | 
**IsResource** | **bool** |  | 
**HasChildren** | **bool** |  | 
**ParentId** | Pointer to **string** |  | [optional] 
**ServiceExtensions** | Pointer to [**[]NSOConfigDataXPath**](NSOConfigDataXPath.md) |  | [optional] 
**ServiceConfigQueryRootXPaths** | Pointer to **[]string** |  | [optional] 
**UiConfig** | Pointer to [**ServiceUIConfig**](ServiceUIConfig.md) |  | [optional] 

## Methods

### NewProductUpdate

`func NewProductUpdate(name string, label string, version int32, description string, image string, price string, options []ServiceElement, properties []ServiceElement, configuration map[string]string, isResource bool, hasChildren bool, ) *ProductUpdate`

NewProductUpdate instantiates a new ProductUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductUpdateWithDefaults

`func NewProductUpdateWithDefaults() *ProductUpdate`

NewProductUpdateWithDefaults instantiates a new ProductUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProductUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *ProductUpdate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ProductUpdate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ProductUpdate) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetVersion

`func (o *ProductUpdate) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ProductUpdate) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ProductUpdate) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetDescription

`func (o *ProductUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetImage

`func (o *ProductUpdate) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ProductUpdate) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ProductUpdate) SetImage(v string)`

SetImage sets Image field to given value.


### GetMultipleInstanceAllowed

`func (o *ProductUpdate) GetMultipleInstanceAllowed() bool`

GetMultipleInstanceAllowed returns the MultipleInstanceAllowed field if non-nil, zero value otherwise.

### GetMultipleInstanceAllowedOk

`func (o *ProductUpdate) GetMultipleInstanceAllowedOk() (*bool, bool)`

GetMultipleInstanceAllowedOk returns a tuple with the MultipleInstanceAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleInstanceAllowed

`func (o *ProductUpdate) SetMultipleInstanceAllowed(v bool)`

SetMultipleInstanceAllowed sets MultipleInstanceAllowed field to given value.

### HasMultipleInstanceAllowed

`func (o *ProductUpdate) HasMultipleInstanceAllowed() bool`

HasMultipleInstanceAllowed returns a boolean if a field has been set.

### GetPrice

`func (o *ProductUpdate) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProductUpdate) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProductUpdate) SetPrice(v string)`

SetPrice sets Price field to given value.


### GetDisplayOrder

`func (o *ProductUpdate) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *ProductUpdate) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *ProductUpdate) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *ProductUpdate) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetActive

`func (o *ProductUpdate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ProductUpdate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ProductUpdate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ProductUpdate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetOrderLimit

`func (o *ProductUpdate) GetOrderLimit() int32`

GetOrderLimit returns the OrderLimit field if non-nil, zero value otherwise.

### GetOrderLimitOk

`func (o *ProductUpdate) GetOrderLimitOk() (*int32, bool)`

GetOrderLimitOk returns a tuple with the OrderLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderLimit

`func (o *ProductUpdate) SetOrderLimit(v int32)`

SetOrderLimit sets OrderLimit field to given value.

### HasOrderLimit

`func (o *ProductUpdate) HasOrderLimit() bool`

HasOrderLimit returns a boolean if a field has been set.

### GetOptions

`func (o *ProductUpdate) GetOptions() []ServiceElement`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ProductUpdate) GetOptionsOk() (*[]ServiceElement, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ProductUpdate) SetOptions(v []ServiceElement)`

SetOptions sets Options field to given value.


### GetProperties

`func (o *ProductUpdate) GetProperties() []ServiceElement`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ProductUpdate) GetPropertiesOk() (*[]ServiceElement, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ProductUpdate) SetProperties(v []ServiceElement)`

SetProperties sets Properties field to given value.


### SetPropertiesNil

`func (o *ProductUpdate) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *ProductUpdate) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetConfiguration

`func (o *ProductUpdate) GetConfiguration() map[string]string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ProductUpdate) GetConfigurationOk() (*map[string]string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ProductUpdate) SetConfiguration(v map[string]string)`

SetConfiguration sets Configuration field to given value.


### GetIsResource

`func (o *ProductUpdate) GetIsResource() bool`

GetIsResource returns the IsResource field if non-nil, zero value otherwise.

### GetIsResourceOk

`func (o *ProductUpdate) GetIsResourceOk() (*bool, bool)`

GetIsResourceOk returns a tuple with the IsResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsResource

`func (o *ProductUpdate) SetIsResource(v bool)`

SetIsResource sets IsResource field to given value.


### GetHasChildren

`func (o *ProductUpdate) GetHasChildren() bool`

GetHasChildren returns the HasChildren field if non-nil, zero value otherwise.

### GetHasChildrenOk

`func (o *ProductUpdate) GetHasChildrenOk() (*bool, bool)`

GetHasChildrenOk returns a tuple with the HasChildren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasChildren

`func (o *ProductUpdate) SetHasChildren(v bool)`

SetHasChildren sets HasChildren field to given value.


### GetParentId

`func (o *ProductUpdate) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ProductUpdate) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ProductUpdate) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ProductUpdate) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetServiceExtensions

`func (o *ProductUpdate) GetServiceExtensions() []NSOConfigDataXPath`

GetServiceExtensions returns the ServiceExtensions field if non-nil, zero value otherwise.

### GetServiceExtensionsOk

`func (o *ProductUpdate) GetServiceExtensionsOk() (*[]NSOConfigDataXPath, bool)`

GetServiceExtensionsOk returns a tuple with the ServiceExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceExtensions

`func (o *ProductUpdate) SetServiceExtensions(v []NSOConfigDataXPath)`

SetServiceExtensions sets ServiceExtensions field to given value.

### HasServiceExtensions

`func (o *ProductUpdate) HasServiceExtensions() bool`

HasServiceExtensions returns a boolean if a field has been set.

### GetServiceConfigQueryRootXPaths

`func (o *ProductUpdate) GetServiceConfigQueryRootXPaths() []string`

GetServiceConfigQueryRootXPaths returns the ServiceConfigQueryRootXPaths field if non-nil, zero value otherwise.

### GetServiceConfigQueryRootXPathsOk

`func (o *ProductUpdate) GetServiceConfigQueryRootXPathsOk() (*[]string, bool)`

GetServiceConfigQueryRootXPathsOk returns a tuple with the ServiceConfigQueryRootXPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceConfigQueryRootXPaths

`func (o *ProductUpdate) SetServiceConfigQueryRootXPaths(v []string)`

SetServiceConfigQueryRootXPaths sets ServiceConfigQueryRootXPaths field to given value.

### HasServiceConfigQueryRootXPaths

`func (o *ProductUpdate) HasServiceConfigQueryRootXPaths() bool`

HasServiceConfigQueryRootXPaths returns a boolean if a field has been set.

### GetUiConfig

`func (o *ProductUpdate) GetUiConfig() ServiceUIConfig`

GetUiConfig returns the UiConfig field if non-nil, zero value otherwise.

### GetUiConfigOk

`func (o *ProductUpdate) GetUiConfigOk() (*ServiceUIConfig, bool)`

GetUiConfigOk returns a tuple with the UiConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiConfig

`func (o *ProductUpdate) SetUiConfig(v ServiceUIConfig)`

SetUiConfig sets UiConfig field to given value.

### HasUiConfig

`func (o *ProductUpdate) HasUiConfig() bool`

HasUiConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


