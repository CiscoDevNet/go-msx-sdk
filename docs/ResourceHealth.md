# ResourceHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | [**ResourceType**](ResourceType.md) |  | 
**Status** | [**ResourceStatus**](ResourceStatus.md) |  | 
**Description** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewResourceHealth

`func NewResourceHealth(id string, type_ ResourceType, status ResourceStatus, ) *ResourceHealth`

NewResourceHealth instantiates a new ResourceHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceHealthWithDefaults

`func NewResourceHealthWithDefaults() *ResourceHealth`

NewResourceHealthWithDefaults instantiates a new ResourceHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceHealth) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceHealth) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceHealth) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ResourceHealth) GetType() ResourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceHealth) GetTypeOk() (*ResourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceHealth) SetType(v ResourceType)`

SetType sets Type field to given value.


### GetStatus

`func (o *ResourceHealth) GetStatus() ResourceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResourceHealth) GetStatusOk() (*ResourceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResourceHealth) SetStatus(v ResourceStatus)`

SetStatus sets Status field to given value.


### GetDescription

`func (o *ResourceHealth) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceHealth) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceHealth) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResourceHealth) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ResourceHealth) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ResourceHealth) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


