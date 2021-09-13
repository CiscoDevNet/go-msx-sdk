# TenantCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentId** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Image** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 

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
### GetExternalId

`func (o *TenantCreate) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *TenantCreate) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *TenantCreate) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *TenantCreate) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *TenantCreate) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *TenantCreate) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
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

### SetDescriptionNil

`func (o *TenantCreate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TenantCreate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
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

### SetUrlNil

`func (o *TenantCreate) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TenantCreate) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
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
### GetEmail

`func (o *TenantCreate) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TenantCreate) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TenantCreate) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *TenantCreate) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *TenantCreate) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *TenantCreate) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


