# TemplateAssignment

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

## Methods

### NewTemplateAssignment

`func NewTemplateAssignment() *TemplateAssignment`

NewTemplateAssignment instantiates a new TemplateAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateAssignmentWithDefaults

`func NewTemplateAssignmentWithDefaults() *TemplateAssignment`

NewTemplateAssignmentWithDefaults instantiates a new TemplateAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TemplateAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTemplateId

`func (o *TemplateAssignment) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *TemplateAssignment) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *TemplateAssignment) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *TemplateAssignment) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetTemplateName

`func (o *TemplateAssignment) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *TemplateAssignment) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *TemplateAssignment) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *TemplateAssignment) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetTenantId

`func (o *TemplateAssignment) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TemplateAssignment) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TemplateAssignment) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TemplateAssignment) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetAssignedTenantId

`func (o *TemplateAssignment) GetAssignedTenantId() string`

GetAssignedTenantId returns the AssignedTenantId field if non-nil, zero value otherwise.

### GetAssignedTenantIdOk

`func (o *TemplateAssignment) GetAssignedTenantIdOk() (*string, bool)`

GetAssignedTenantIdOk returns a tuple with the AssignedTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTenantId

`func (o *TemplateAssignment) SetAssignedTenantId(v string)`

SetAssignedTenantId sets AssignedTenantId field to given value.

### HasAssignedTenantId

`func (o *TemplateAssignment) HasAssignedTenantId() bool`

HasAssignedTenantId returns a boolean if a field has been set.

### GetInheritable

`func (o *TemplateAssignment) GetInheritable() bool`

GetInheritable returns the Inheritable field if non-nil, zero value otherwise.

### GetInheritableOk

`func (o *TemplateAssignment) GetInheritableOk() (*bool, bool)`

GetInheritableOk returns a tuple with the Inheritable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritable

`func (o *TemplateAssignment) SetInheritable(v bool)`

SetInheritable sets Inheritable field to given value.

### HasInheritable

`func (o *TemplateAssignment) HasInheritable() bool`

HasInheritable returns a boolean if a field has been set.

### SetInheritableNil

`func (o *TemplateAssignment) SetInheritableNil(b bool)`

 SetInheritableNil sets the value for Inheritable to be an explicit nil

### UnsetInheritable
`func (o *TemplateAssignment) UnsetInheritable()`

UnsetInheritable ensures that no value is present for Inheritable, not even an explicit nil
### GetStatus

`func (o *TemplateAssignment) GetStatus() TemplateStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TemplateAssignment) GetStatusOk() (*TemplateStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TemplateAssignment) SetStatus(v TemplateStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TemplateAssignment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDetails

`func (o *TemplateAssignment) GetStatusDetails() string`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *TemplateAssignment) GetStatusDetailsOk() (*string, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *TemplateAssignment) SetStatusDetails(v string)`

SetStatusDetails sets StatusDetails field to given value.

### HasStatusDetails

`func (o *TemplateAssignment) HasStatusDetails() bool`

HasStatusDetails returns a boolean if a field has been set.

### SetStatusDetailsNil

`func (o *TemplateAssignment) SetStatusDetailsNil(b bool)`

 SetStatusDetailsNil sets the value for StatusDetails to be an explicit nil

### UnsetStatusDetails
`func (o *TemplateAssignment) UnsetStatusDetails()`

UnsetStatusDetails ensures that no value is present for StatusDetails, not even an explicit nil
### GetCreatedOn

`func (o *TemplateAssignment) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *TemplateAssignment) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *TemplateAssignment) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *TemplateAssignment) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### SetCreatedOnNil

`func (o *TemplateAssignment) SetCreatedOnNil(b bool)`

 SetCreatedOnNil sets the value for CreatedOn to be an explicit nil

### UnsetCreatedOn
`func (o *TemplateAssignment) UnsetCreatedOn()`

UnsetCreatedOn ensures that no value is present for CreatedOn, not even an explicit nil
### GetCreatedBy

`func (o *TemplateAssignment) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *TemplateAssignment) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *TemplateAssignment) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *TemplateAssignment) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *TemplateAssignment) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *TemplateAssignment) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetModifiedOn

`func (o *TemplateAssignment) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *TemplateAssignment) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *TemplateAssignment) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *TemplateAssignment) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### SetModifiedOnNil

`func (o *TemplateAssignment) SetModifiedOnNil(b bool)`

 SetModifiedOnNil sets the value for ModifiedOn to be an explicit nil

### UnsetModifiedOn
`func (o *TemplateAssignment) UnsetModifiedOn()`

UnsetModifiedOn ensures that no value is present for ModifiedOn, not even an explicit nil
### GetModifiedBy

`func (o *TemplateAssignment) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *TemplateAssignment) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *TemplateAssignment) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *TemplateAssignment) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### SetModifiedByNil

`func (o *TemplateAssignment) SetModifiedByNil(b bool)`

 SetModifiedByNil sets the value for ModifiedBy to be an explicit nil

### UnsetModifiedBy
`func (o *TemplateAssignment) UnsetModifiedBy()`

UnsetModifiedBy ensures that no value is present for ModifiedBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


