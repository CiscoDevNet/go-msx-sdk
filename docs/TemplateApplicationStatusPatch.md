# TemplateApplicationStatusPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**TemplateStatus**](TemplateStatus.md) |  | 
**StatusDetails** | Pointer to **string** |  | [optional] 

## Methods

### NewTemplateApplicationStatusPatch

`func NewTemplateApplicationStatusPatch(status TemplateStatus, ) *TemplateApplicationStatusPatch`

NewTemplateApplicationStatusPatch instantiates a new TemplateApplicationStatusPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateApplicationStatusPatchWithDefaults

`func NewTemplateApplicationStatusPatchWithDefaults() *TemplateApplicationStatusPatch`

NewTemplateApplicationStatusPatchWithDefaults instantiates a new TemplateApplicationStatusPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *TemplateApplicationStatusPatch) GetStatus() TemplateStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TemplateApplicationStatusPatch) GetStatusOk() (*TemplateStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TemplateApplicationStatusPatch) SetStatus(v TemplateStatus)`

SetStatus sets Status field to given value.


### GetStatusDetails

`func (o *TemplateApplicationStatusPatch) GetStatusDetails() string`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *TemplateApplicationStatusPatch) GetStatusDetailsOk() (*string, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *TemplateApplicationStatusPatch) SetStatusDetails(v string)`

SetStatusDetails sets StatusDetails field to given value.

### HasStatusDetails

`func (o *TemplateApplicationStatusPatch) HasStatusDetails() bool`

HasStatusDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


