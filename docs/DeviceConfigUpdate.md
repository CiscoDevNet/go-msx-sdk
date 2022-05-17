# DeviceConfigUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrelationId** | **string** |  | 
**ConfigContent** | Pointer to **string** |  | [optional] 
**Format** | **string** |  | 

## Methods

### NewDeviceConfigUpdate

`func NewDeviceConfigUpdate(correlationId string, format string, ) *DeviceConfigUpdate`

NewDeviceConfigUpdate instantiates a new DeviceConfigUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceConfigUpdateWithDefaults

`func NewDeviceConfigUpdateWithDefaults() *DeviceConfigUpdate`

NewDeviceConfigUpdateWithDefaults instantiates a new DeviceConfigUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrelationId

`func (o *DeviceConfigUpdate) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *DeviceConfigUpdate) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *DeviceConfigUpdate) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.


### GetConfigContent

`func (o *DeviceConfigUpdate) GetConfigContent() string`

GetConfigContent returns the ConfigContent field if non-nil, zero value otherwise.

### GetConfigContentOk

`func (o *DeviceConfigUpdate) GetConfigContentOk() (*string, bool)`

GetConfigContentOk returns a tuple with the ConfigContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContent

`func (o *DeviceConfigUpdate) SetConfigContent(v string)`

SetConfigContent sets ConfigContent field to given value.

### HasConfigContent

`func (o *DeviceConfigUpdate) HasConfigContent() bool`

HasConfigContent returns a boolean if a field has been set.

### GetFormat

`func (o *DeviceConfigUpdate) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *DeviceConfigUpdate) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *DeviceConfigUpdate) SetFormat(v string)`

SetFormat sets Format field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


