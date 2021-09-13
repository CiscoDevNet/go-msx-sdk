# Tenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] [readonly] 
**ModifiedOn** | Pointer to **time.Time** |  | [optional] [readonly] 
**Suspended** | Pointer to **NullableBool** |  | [optional] [readonly] 
**NumberOfChildren** | Pointer to **NullableInt64** |  | [optional] [readonly] 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Image** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTenant

`func NewTenant(name string, ) *Tenant`

NewTenant instantiates a new Tenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantWithDefaults

`func NewTenantWithDefaults() *Tenant`

NewTenantWithDefaults instantiates a new Tenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Tenant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Tenant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Tenant) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Tenant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedOn

`func (o *Tenant) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Tenant) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Tenant) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *Tenant) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetModifiedOn

`func (o *Tenant) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *Tenant) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *Tenant) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *Tenant) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetSuspended

`func (o *Tenant) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *Tenant) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *Tenant) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *Tenant) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### SetSuspendedNil

`func (o *Tenant) SetSuspendedNil(b bool)`

 SetSuspendedNil sets the value for Suspended to be an explicit nil

### UnsetSuspended
`func (o *Tenant) UnsetSuspended()`

UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
### GetNumberOfChildren

`func (o *Tenant) GetNumberOfChildren() int64`

GetNumberOfChildren returns the NumberOfChildren field if non-nil, zero value otherwise.

### GetNumberOfChildrenOk

`func (o *Tenant) GetNumberOfChildrenOk() (*int64, bool)`

GetNumberOfChildrenOk returns a tuple with the NumberOfChildren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfChildren

`func (o *Tenant) SetNumberOfChildren(v int64)`

SetNumberOfChildren sets NumberOfChildren field to given value.

### HasNumberOfChildren

`func (o *Tenant) HasNumberOfChildren() bool`

HasNumberOfChildren returns a boolean if a field has been set.

### SetNumberOfChildrenNil

`func (o *Tenant) SetNumberOfChildrenNil(b bool)`

 SetNumberOfChildrenNil sets the value for NumberOfChildren to be an explicit nil

### UnsetNumberOfChildren
`func (o *Tenant) UnsetNumberOfChildren()`

UnsetNumberOfChildren ensures that no value is present for NumberOfChildren, not even an explicit nil
### GetParentId

`func (o *Tenant) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Tenant) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Tenant) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Tenant) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *Tenant) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *Tenant) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetExternalId

`func (o *Tenant) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Tenant) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Tenant) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Tenant) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *Tenant) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *Tenant) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetName

`func (o *Tenant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tenant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tenant) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Tenant) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Tenant) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Tenant) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Tenant) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Tenant) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Tenant) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUrl

`func (o *Tenant) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Tenant) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Tenant) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Tenant) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *Tenant) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *Tenant) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetImage

`func (o *Tenant) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Tenant) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Tenant) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *Tenant) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *Tenant) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *Tenant) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetEmail

`func (o *Tenant) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Tenant) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Tenant) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Tenant) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *Tenant) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *Tenant) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


