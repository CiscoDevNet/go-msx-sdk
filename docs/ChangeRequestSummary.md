# ChangeRequestSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**Attributes** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewChangeRequestSummary

`func NewChangeRequestSummary() *ChangeRequestSummary`

NewChangeRequestSummary instantiates a new ChangeRequestSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeRequestSummaryWithDefaults

`func NewChangeRequestSummaryWithDefaults() *ChangeRequestSummary`

NewChangeRequestSummaryWithDefaults instantiates a new ChangeRequestSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChangeRequestSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChangeRequestSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChangeRequestSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChangeRequestSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExternalId

`func (o *ChangeRequestSummary) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ChangeRequestSummary) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ChangeRequestSummary) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ChangeRequestSummary) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetName

`func (o *ChangeRequestSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChangeRequestSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChangeRequestSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChangeRequestSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ChangeRequestSummary) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChangeRequestSummary) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChangeRequestSummary) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChangeRequestSummary) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTenantId

`func (o *ChangeRequestSummary) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ChangeRequestSummary) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ChangeRequestSummary) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ChangeRequestSummary) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ChangeRequestSummary) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ChangeRequestSummary) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetAttributes

`func (o *ChangeRequestSummary) GetAttributes() map[string]map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ChangeRequestSummary) GetAttributesOk() (*map[string]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ChangeRequestSummary) SetAttributes(v map[string]map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ChangeRequestSummary) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


