# DeviceTemplateBatchAttachRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceIds** | Pointer to **[]string** |  | [optional] 
**TemplateDetails** | Pointer to [**[]DeviceTemplateDetails**](DeviceTemplateDetails.md) |  | [optional] 

## Methods

### NewDeviceTemplateBatchAttachRequest

`func NewDeviceTemplateBatchAttachRequest() *DeviceTemplateBatchAttachRequest`

NewDeviceTemplateBatchAttachRequest instantiates a new DeviceTemplateBatchAttachRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceTemplateBatchAttachRequestWithDefaults

`func NewDeviceTemplateBatchAttachRequestWithDefaults() *DeviceTemplateBatchAttachRequest`

NewDeviceTemplateBatchAttachRequestWithDefaults instantiates a new DeviceTemplateBatchAttachRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceIds

`func (o *DeviceTemplateBatchAttachRequest) GetDeviceIds() []string`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *DeviceTemplateBatchAttachRequest) GetDeviceIdsOk() (*[]string, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIds

`func (o *DeviceTemplateBatchAttachRequest) SetDeviceIds(v []string)`

SetDeviceIds sets DeviceIds field to given value.

### HasDeviceIds

`func (o *DeviceTemplateBatchAttachRequest) HasDeviceIds() bool`

HasDeviceIds returns a boolean if a field has been set.

### GetTemplateDetails

`func (o *DeviceTemplateBatchAttachRequest) GetTemplateDetails() []DeviceTemplateDetails`

GetTemplateDetails returns the TemplateDetails field if non-nil, zero value otherwise.

### GetTemplateDetailsOk

`func (o *DeviceTemplateBatchAttachRequest) GetTemplateDetailsOk() (*[]DeviceTemplateDetails, bool)`

GetTemplateDetailsOk returns a tuple with the TemplateDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateDetails

`func (o *DeviceTemplateBatchAttachRequest) SetTemplateDetails(v []DeviceTemplateDetails)`

SetTemplateDetails sets TemplateDetails field to given value.

### HasTemplateDetails

`func (o *DeviceTemplateBatchAttachRequest) HasTemplateDetails() bool`

HasTemplateDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


