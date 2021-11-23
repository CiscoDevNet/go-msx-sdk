# TemplateApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**TemplateId** | Pointer to **string** |  | [optional] [readonly] 
**TemplateName** | Pointer to **string** |  | [optional] [readonly] 
**TenantId** | **string** |  | 
**TargetId** | **string** |  | 
**TargetType** | **string** |  | 
**Parameters** | Pointer to **map[string]string** |  | [optional] 
**Status** | Pointer to [**TemplateStatus**](TemplateStatus.md) |  | [optional] 
**StatusDetails** | Pointer to **NullableString** |  | [optional] [readonly] 
**CreatedOn** | Pointer to **NullableTime** |  | [optional] [readonly] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] [readonly] 
**ModifiedOn** | Pointer to **NullableTime** |  | [optional] [readonly] 
**ModifiedBy** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewTemplateApplication

`func NewTemplateApplication(tenantId string, targetId string, targetType string, ) *TemplateApplication`

NewTemplateApplication instantiates a new TemplateApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateApplicationWithDefaults

`func NewTemplateApplicationWithDefaults() *TemplateApplication`

NewTemplateApplicationWithDefaults instantiates a new TemplateApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateApplication) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateApplication) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateApplication) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TemplateApplication) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTemplateId

`func (o *TemplateApplication) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *TemplateApplication) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *TemplateApplication) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *TemplateApplication) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetTemplateName

`func (o *TemplateApplication) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *TemplateApplication) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *TemplateApplication) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *TemplateApplication) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetTenantId

`func (o *TemplateApplication) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TemplateApplication) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TemplateApplication) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetTargetId

`func (o *TemplateApplication) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *TemplateApplication) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *TemplateApplication) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.


### GetTargetType

`func (o *TemplateApplication) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *TemplateApplication) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *TemplateApplication) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.


### GetParameters

`func (o *TemplateApplication) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *TemplateApplication) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *TemplateApplication) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *TemplateApplication) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetStatus

`func (o *TemplateApplication) GetStatus() TemplateStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TemplateApplication) GetStatusOk() (*TemplateStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TemplateApplication) SetStatus(v TemplateStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TemplateApplication) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDetails

`func (o *TemplateApplication) GetStatusDetails() string`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *TemplateApplication) GetStatusDetailsOk() (*string, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *TemplateApplication) SetStatusDetails(v string)`

SetStatusDetails sets StatusDetails field to given value.

### HasStatusDetails

`func (o *TemplateApplication) HasStatusDetails() bool`

HasStatusDetails returns a boolean if a field has been set.

### SetStatusDetailsNil

`func (o *TemplateApplication) SetStatusDetailsNil(b bool)`

 SetStatusDetailsNil sets the value for StatusDetails to be an explicit nil

### UnsetStatusDetails
`func (o *TemplateApplication) UnsetStatusDetails()`

UnsetStatusDetails ensures that no value is present for StatusDetails, not even an explicit nil
### GetCreatedOn

`func (o *TemplateApplication) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *TemplateApplication) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *TemplateApplication) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *TemplateApplication) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### SetCreatedOnNil

`func (o *TemplateApplication) SetCreatedOnNil(b bool)`

 SetCreatedOnNil sets the value for CreatedOn to be an explicit nil

### UnsetCreatedOn
`func (o *TemplateApplication) UnsetCreatedOn()`

UnsetCreatedOn ensures that no value is present for CreatedOn, not even an explicit nil
### GetCreatedBy

`func (o *TemplateApplication) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *TemplateApplication) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *TemplateApplication) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *TemplateApplication) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *TemplateApplication) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *TemplateApplication) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetModifiedOn

`func (o *TemplateApplication) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *TemplateApplication) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *TemplateApplication) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *TemplateApplication) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### SetModifiedOnNil

`func (o *TemplateApplication) SetModifiedOnNil(b bool)`

 SetModifiedOnNil sets the value for ModifiedOn to be an explicit nil

### UnsetModifiedOn
`func (o *TemplateApplication) UnsetModifiedOn()`

UnsetModifiedOn ensures that no value is present for ModifiedOn, not even an explicit nil
### GetModifiedBy

`func (o *TemplateApplication) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *TemplateApplication) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *TemplateApplication) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *TemplateApplication) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### SetModifiedByNil

`func (o *TemplateApplication) SetModifiedByNil(b bool)`

 SetModifiedByNil sets the value for ModifiedBy to be an explicit nil

### UnsetModifiedBy
`func (o *TemplateApplication) UnsetModifiedBy()`

UnsetModifiedBy ensures that no value is present for ModifiedBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


