# TemplateStatusMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**TemplateStatus**](TemplateStatus.md) |  | [optional] 
**StatusDetails** | Pointer to **NullableString** |  | [optional] [readonly] 
**CreatedOn** | Pointer to **NullableTime** |  | [optional] [readonly] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] [readonly] 
**ModifiedOn** | Pointer to **NullableTime** |  | [optional] [readonly] 
**ModifiedBy** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewTemplateStatusMeta

`func NewTemplateStatusMeta() *TemplateStatusMeta`

NewTemplateStatusMeta instantiates a new TemplateStatusMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateStatusMetaWithDefaults

`func NewTemplateStatusMetaWithDefaults() *TemplateStatusMeta`

NewTemplateStatusMetaWithDefaults instantiates a new TemplateStatusMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *TemplateStatusMeta) GetStatus() TemplateStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TemplateStatusMeta) GetStatusOk() (*TemplateStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TemplateStatusMeta) SetStatus(v TemplateStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TemplateStatusMeta) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDetails

`func (o *TemplateStatusMeta) GetStatusDetails() string`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *TemplateStatusMeta) GetStatusDetailsOk() (*string, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *TemplateStatusMeta) SetStatusDetails(v string)`

SetStatusDetails sets StatusDetails field to given value.

### HasStatusDetails

`func (o *TemplateStatusMeta) HasStatusDetails() bool`

HasStatusDetails returns a boolean if a field has been set.

### SetStatusDetailsNil

`func (o *TemplateStatusMeta) SetStatusDetailsNil(b bool)`

 SetStatusDetailsNil sets the value for StatusDetails to be an explicit nil

### UnsetStatusDetails
`func (o *TemplateStatusMeta) UnsetStatusDetails()`

UnsetStatusDetails ensures that no value is present for StatusDetails, not even an explicit nil
### GetCreatedOn

`func (o *TemplateStatusMeta) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *TemplateStatusMeta) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *TemplateStatusMeta) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *TemplateStatusMeta) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### SetCreatedOnNil

`func (o *TemplateStatusMeta) SetCreatedOnNil(b bool)`

 SetCreatedOnNil sets the value for CreatedOn to be an explicit nil

### UnsetCreatedOn
`func (o *TemplateStatusMeta) UnsetCreatedOn()`

UnsetCreatedOn ensures that no value is present for CreatedOn, not even an explicit nil
### GetCreatedBy

`func (o *TemplateStatusMeta) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *TemplateStatusMeta) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *TemplateStatusMeta) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *TemplateStatusMeta) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *TemplateStatusMeta) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *TemplateStatusMeta) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetModifiedOn

`func (o *TemplateStatusMeta) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *TemplateStatusMeta) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *TemplateStatusMeta) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *TemplateStatusMeta) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### SetModifiedOnNil

`func (o *TemplateStatusMeta) SetModifiedOnNil(b bool)`

 SetModifiedOnNil sets the value for ModifiedOn to be an explicit nil

### UnsetModifiedOn
`func (o *TemplateStatusMeta) UnsetModifiedOn()`

UnsetModifiedOn ensures that no value is present for ModifiedOn, not even an explicit nil
### GetModifiedBy

`func (o *TemplateStatusMeta) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *TemplateStatusMeta) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *TemplateStatusMeta) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *TemplateStatusMeta) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### SetModifiedByNil

`func (o *TemplateStatusMeta) SetModifiedByNil(b bool)`

 SetModifiedByNil sets the value for ModifiedBy to be an explicit nil

### UnsetModifiedBy
`func (o *TemplateStatusMeta) UnsetModifiedBy()`

UnsetModifiedBy ensures that no value is present for ModifiedBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


