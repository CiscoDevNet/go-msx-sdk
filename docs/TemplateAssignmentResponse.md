# TemplateAssignmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**TemplateId** | Pointer to **string** |  | [optional] [readonly] 
**TemplateName** | Pointer to **string** |  | [optional] [readonly] 
**TenantId** | Pointer to **string** |  | [optional] [readonly] 
**AssignedTenantId** | Pointer to **string** |  | [optional] [readonly] 
**Inheritable** | Pointer to **NullableBool** |  | [optional] [readonly] 
**Status** | Pointer to [**TemplateStatus**](TemplateStatus.md) |  | [optional] 
**StatusDetails** | Pointer to **NullableString** |  | [optional] [readonly] 
**CreatedOn** | Pointer to **NullableTime** |  | [optional] [readonly] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] [readonly] 
**ModifiedOn** | Pointer to **NullableTime** |  | [optional] [readonly] 
**ModifiedBy** | Pointer to **NullableString** |  | [optional] [readonly] 
**ResponseStatus** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewTemplateAssignmentResponse

`func NewTemplateAssignmentResponse() *TemplateAssignmentResponse`

NewTemplateAssignmentResponse instantiates a new TemplateAssignmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateAssignmentResponseWithDefaults

`func NewTemplateAssignmentResponseWithDefaults() *TemplateAssignmentResponse`

NewTemplateAssignmentResponseWithDefaults instantiates a new TemplateAssignmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateAssignmentResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateAssignmentResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateAssignmentResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TemplateAssignmentResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTemplateId

`func (o *TemplateAssignmentResponse) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *TemplateAssignmentResponse) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *TemplateAssignmentResponse) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *TemplateAssignmentResponse) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetTemplateName

`func (o *TemplateAssignmentResponse) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *TemplateAssignmentResponse) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *TemplateAssignmentResponse) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *TemplateAssignmentResponse) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetTenantId

`func (o *TemplateAssignmentResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TemplateAssignmentResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TemplateAssignmentResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TemplateAssignmentResponse) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetAssignedTenantId

`func (o *TemplateAssignmentResponse) GetAssignedTenantId() string`

GetAssignedTenantId returns the AssignedTenantId field if non-nil, zero value otherwise.

### GetAssignedTenantIdOk

`func (o *TemplateAssignmentResponse) GetAssignedTenantIdOk() (*string, bool)`

GetAssignedTenantIdOk returns a tuple with the AssignedTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTenantId

`func (o *TemplateAssignmentResponse) SetAssignedTenantId(v string)`

SetAssignedTenantId sets AssignedTenantId field to given value.

### HasAssignedTenantId

`func (o *TemplateAssignmentResponse) HasAssignedTenantId() bool`

HasAssignedTenantId returns a boolean if a field has been set.

### GetInheritable

`func (o *TemplateAssignmentResponse) GetInheritable() bool`

GetInheritable returns the Inheritable field if non-nil, zero value otherwise.

### GetInheritableOk

`func (o *TemplateAssignmentResponse) GetInheritableOk() (*bool, bool)`

GetInheritableOk returns a tuple with the Inheritable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritable

`func (o *TemplateAssignmentResponse) SetInheritable(v bool)`

SetInheritable sets Inheritable field to given value.

### HasInheritable

`func (o *TemplateAssignmentResponse) HasInheritable() bool`

HasInheritable returns a boolean if a field has been set.

### SetInheritableNil

`func (o *TemplateAssignmentResponse) SetInheritableNil(b bool)`

 SetInheritableNil sets the value for Inheritable to be an explicit nil

### UnsetInheritable
`func (o *TemplateAssignmentResponse) UnsetInheritable()`

UnsetInheritable ensures that no value is present for Inheritable, not even an explicit nil
### GetStatus

`func (o *TemplateAssignmentResponse) GetStatus() TemplateStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TemplateAssignmentResponse) GetStatusOk() (*TemplateStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TemplateAssignmentResponse) SetStatus(v TemplateStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TemplateAssignmentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDetails

`func (o *TemplateAssignmentResponse) GetStatusDetails() string`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *TemplateAssignmentResponse) GetStatusDetailsOk() (*string, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *TemplateAssignmentResponse) SetStatusDetails(v string)`

SetStatusDetails sets StatusDetails field to given value.

### HasStatusDetails

`func (o *TemplateAssignmentResponse) HasStatusDetails() bool`

HasStatusDetails returns a boolean if a field has been set.

### SetStatusDetailsNil

`func (o *TemplateAssignmentResponse) SetStatusDetailsNil(b bool)`

 SetStatusDetailsNil sets the value for StatusDetails to be an explicit nil

### UnsetStatusDetails
`func (o *TemplateAssignmentResponse) UnsetStatusDetails()`

UnsetStatusDetails ensures that no value is present for StatusDetails, not even an explicit nil
### GetCreatedOn

`func (o *TemplateAssignmentResponse) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *TemplateAssignmentResponse) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *TemplateAssignmentResponse) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *TemplateAssignmentResponse) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### SetCreatedOnNil

`func (o *TemplateAssignmentResponse) SetCreatedOnNil(b bool)`

 SetCreatedOnNil sets the value for CreatedOn to be an explicit nil

### UnsetCreatedOn
`func (o *TemplateAssignmentResponse) UnsetCreatedOn()`

UnsetCreatedOn ensures that no value is present for CreatedOn, not even an explicit nil
### GetCreatedBy

`func (o *TemplateAssignmentResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *TemplateAssignmentResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *TemplateAssignmentResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *TemplateAssignmentResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *TemplateAssignmentResponse) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *TemplateAssignmentResponse) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetModifiedOn

`func (o *TemplateAssignmentResponse) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *TemplateAssignmentResponse) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *TemplateAssignmentResponse) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *TemplateAssignmentResponse) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### SetModifiedOnNil

`func (o *TemplateAssignmentResponse) SetModifiedOnNil(b bool)`

 SetModifiedOnNil sets the value for ModifiedOn to be an explicit nil

### UnsetModifiedOn
`func (o *TemplateAssignmentResponse) UnsetModifiedOn()`

UnsetModifiedOn ensures that no value is present for ModifiedOn, not even an explicit nil
### GetModifiedBy

`func (o *TemplateAssignmentResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *TemplateAssignmentResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *TemplateAssignmentResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *TemplateAssignmentResponse) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### SetModifiedByNil

`func (o *TemplateAssignmentResponse) SetModifiedByNil(b bool)`

 SetModifiedByNil sets the value for ModifiedBy to be an explicit nil

### UnsetModifiedBy
`func (o *TemplateAssignmentResponse) UnsetModifiedBy()`

UnsetModifiedBy ensures that no value is present for ModifiedBy, not even an explicit nil
### GetResponseStatus

`func (o *TemplateAssignmentResponse) GetResponseStatus() string`

GetResponseStatus returns the ResponseStatus field if non-nil, zero value otherwise.

### GetResponseStatusOk

`func (o *TemplateAssignmentResponse) GetResponseStatusOk() (*string, bool)`

GetResponseStatusOk returns a tuple with the ResponseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatus

`func (o *TemplateAssignmentResponse) SetResponseStatus(v string)`

SetResponseStatus sets ResponseStatus field to given value.

### HasResponseStatus

`func (o *TemplateAssignmentResponse) HasResponseStatus() bool`

HasResponseStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


