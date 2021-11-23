# TemplatePatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ServiceType** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Subtype** | Pointer to **NullableString** |  | [optional] 
**Configuration** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to **map[string]string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**TemplateStatus**](TemplateStatus.md) |  | [optional] 
**StatusDetails** | Pointer to **string** |  | [optional] 

## Methods

### NewTemplatePatch

`func NewTemplatePatch() *TemplatePatch`

NewTemplatePatch instantiates a new TemplatePatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatePatchWithDefaults

`func NewTemplatePatchWithDefaults() *TemplatePatch`

NewTemplatePatchWithDefaults instantiates a new TemplatePatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TemplatePatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplatePatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplatePatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TemplatePatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *TemplatePatch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplatePatch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplatePatch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TemplatePatch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetServiceType

`func (o *TemplatePatch) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *TemplatePatch) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *TemplatePatch) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *TemplatePatch) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetType

`func (o *TemplatePatch) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TemplatePatch) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TemplatePatch) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TemplatePatch) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubtype

`func (o *TemplatePatch) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *TemplatePatch) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *TemplatePatch) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *TemplatePatch) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### SetSubtypeNil

`func (o *TemplatePatch) SetSubtypeNil(b bool)`

 SetSubtypeNil sets the value for Subtype to be an explicit nil

### UnsetSubtype
`func (o *TemplatePatch) UnsetSubtype()`

UnsetSubtype ensures that no value is present for Subtype, not even an explicit nil
### GetConfiguration

`func (o *TemplatePatch) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *TemplatePatch) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *TemplatePatch) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *TemplatePatch) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetAttributes

`func (o *TemplatePatch) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TemplatePatch) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TemplatePatch) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TemplatePatch) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetTags

`func (o *TemplatePatch) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TemplatePatch) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TemplatePatch) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TemplatePatch) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNotes

`func (o *TemplatePatch) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *TemplatePatch) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *TemplatePatch) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *TemplatePatch) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetStatus

`func (o *TemplatePatch) GetStatus() TemplateStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TemplatePatch) GetStatusOk() (*TemplateStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TemplatePatch) SetStatus(v TemplateStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TemplatePatch) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDetails

`func (o *TemplatePatch) GetStatusDetails() string`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *TemplatePatch) GetStatusDetailsOk() (*string, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *TemplatePatch) SetStatusDetails(v string)`

SetStatusDetails sets StatusDetails field to given value.

### HasStatusDetails

`func (o *TemplatePatch) HasStatusDetails() bool`

HasStatusDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


