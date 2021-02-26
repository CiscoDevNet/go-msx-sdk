# DeviceTemplateUpdateDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateHistoryId** | Pointer to **string** |  | [optional] 
**TemplateParams** | Pointer to [**[]NameValue**](NameValue.md) |  | [optional] 

## Methods

### NewDeviceTemplateUpdateDetails

`func NewDeviceTemplateUpdateDetails() *DeviceTemplateUpdateDetails`

NewDeviceTemplateUpdateDetails instantiates a new DeviceTemplateUpdateDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceTemplateUpdateDetailsWithDefaults

`func NewDeviceTemplateUpdateDetailsWithDefaults() *DeviceTemplateUpdateDetails`

NewDeviceTemplateUpdateDetailsWithDefaults instantiates a new DeviceTemplateUpdateDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateHistoryId

`func (o *DeviceTemplateUpdateDetails) GetTemplateHistoryId() string`

GetTemplateHistoryId returns the TemplateHistoryId field if non-nil, zero value otherwise.

### GetTemplateHistoryIdOk

`func (o *DeviceTemplateUpdateDetails) GetTemplateHistoryIdOk() (*string, bool)`

GetTemplateHistoryIdOk returns a tuple with the TemplateHistoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateHistoryId

`func (o *DeviceTemplateUpdateDetails) SetTemplateHistoryId(v string)`

SetTemplateHistoryId sets TemplateHistoryId field to given value.

### HasTemplateHistoryId

`func (o *DeviceTemplateUpdateDetails) HasTemplateHistoryId() bool`

HasTemplateHistoryId returns a boolean if a field has been set.

### GetTemplateParams

`func (o *DeviceTemplateUpdateDetails) GetTemplateParams() []NameValue`

GetTemplateParams returns the TemplateParams field if non-nil, zero value otherwise.

### GetTemplateParamsOk

`func (o *DeviceTemplateUpdateDetails) GetTemplateParamsOk() (*[]NameValue, bool)`

GetTemplateParamsOk returns a tuple with the TemplateParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateParams

`func (o *DeviceTemplateUpdateDetails) SetTemplateParams(v []NameValue)`

SetTemplateParams sets TemplateParams field to given value.

### HasTemplateParams

`func (o *DeviceTemplateUpdateDetails) HasTemplateParams() bool`

HasTemplateParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


