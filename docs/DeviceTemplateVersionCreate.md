# DeviceTemplateVersionCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ConfigContent** | **string** |  | 
**TemplateParameterValidators** | Pointer to [**[]TemplateParameterValidator**](TemplateParameterValidator.md) |  | [optional] 

## Methods

### NewDeviceTemplateVersionCreate

`func NewDeviceTemplateVersionCreate(name string, configContent string, ) *DeviceTemplateVersionCreate`

NewDeviceTemplateVersionCreate instantiates a new DeviceTemplateVersionCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceTemplateVersionCreateWithDefaults

`func NewDeviceTemplateVersionCreateWithDefaults() *DeviceTemplateVersionCreate`

NewDeviceTemplateVersionCreateWithDefaults instantiates a new DeviceTemplateVersionCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeviceTemplateVersionCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceTemplateVersionCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceTemplateVersionCreate) SetName(v string)`

SetName sets Name field to given value.


### GetConfigContent

`func (o *DeviceTemplateVersionCreate) GetConfigContent() string`

GetConfigContent returns the ConfigContent field if non-nil, zero value otherwise.

### GetConfigContentOk

`func (o *DeviceTemplateVersionCreate) GetConfigContentOk() (*string, bool)`

GetConfigContentOk returns a tuple with the ConfigContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContent

`func (o *DeviceTemplateVersionCreate) SetConfigContent(v string)`

SetConfigContent sets ConfigContent field to given value.


### GetTemplateParameterValidators

`func (o *DeviceTemplateVersionCreate) GetTemplateParameterValidators() []TemplateParameterValidator`

GetTemplateParameterValidators returns the TemplateParameterValidators field if non-nil, zero value otherwise.

### GetTemplateParameterValidatorsOk

`func (o *DeviceTemplateVersionCreate) GetTemplateParameterValidatorsOk() (*[]TemplateParameterValidator, bool)`

GetTemplateParameterValidatorsOk returns a tuple with the TemplateParameterValidators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateParameterValidators

`func (o *DeviceTemplateVersionCreate) SetTemplateParameterValidators(v []TemplateParameterValidator)`

SetTemplateParameterValidators sets TemplateParameterValidators field to given value.

### HasTemplateParameterValidators

`func (o *DeviceTemplateVersionCreate) HasTemplateParameterValidators() bool`

HasTemplateParameterValidators returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


