# TenantCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **NullableString** |  | [optional] 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**NumberOfChildren** | Pointer to **NullableInt64** |  | [optional] [readonly] 

## Methods

### NewTenantCreate

`func NewTenantCreate(name string, ) *TenantCreate`

NewTenantCreate instantiates a new TenantCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantCreateWithDefaults

`func NewTenantCreateWithDefaults() *TenantCreate`

NewTenantCreateWithDefaults instantiates a new TenantCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TenantCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TenantCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TenantCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TenantCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TenantCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *TenantCreate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TenantCreate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TenantCreate) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TenantCreate) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetImage

`func (o *TenantCreate) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *TenantCreate) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *TenantCreate) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *TenantCreate) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *TenantCreate) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *TenantCreate) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetParentId

`func (o *TenantCreate) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *TenantCreate) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *TenantCreate) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *TenantCreate) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *TenantCreate) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *TenantCreate) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetNumberOfChildren

`func (o *TenantCreate) GetNumberOfChildren() int64`

GetNumberOfChildren returns the NumberOfChildren field if non-nil, zero value otherwise.

### GetNumberOfChildrenOk

`func (o *TenantCreate) GetNumberOfChildrenOk() (*int64, bool)`

GetNumberOfChildrenOk returns a tuple with the NumberOfChildren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfChildren

`func (o *TenantCreate) SetNumberOfChildren(v int64)`

SetNumberOfChildren sets NumberOfChildren field to given value.

### HasNumberOfChildren

`func (o *TenantCreate) HasNumberOfChildren() bool`

HasNumberOfChildren returns a boolean if a field has been set.

### SetNumberOfChildrenNil

`func (o *TenantCreate) SetNumberOfChildrenNil(b bool)`

 SetNumberOfChildrenNil sets the value for NumberOfChildren to be an explicit nil

### UnsetNumberOfChildren
`func (o *TenantCreate) UnsetNumberOfChildren()`

UnsetNumberOfChildren ensures that no value is present for NumberOfChildren, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


