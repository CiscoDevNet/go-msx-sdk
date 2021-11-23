# IncidentUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**Category** | Pointer to **string** | inquiry|software|hardware|network|database | [optional] [default to "inquiry"]
**Description** | **string** |  | 
**Impact** | Pointer to **string** |  | [optional] [default to "Low"]
**Priority** | Pointer to **string** |  | [optional] [default to "Planning"]
**Severity** | Pointer to **string** |  | [optional] [default to "Low"]
**State** | Pointer to **string** |  | [optional] [default to "New"]
**Subcategory** | Pointer to **string** |  | [optional] 
**Tenant** | Pointer to **NullableString** |  | [optional] 
**Urgency** | Pointer to **string** |  | [optional] [default to "Low"]

## Methods

### NewIncidentUpdate

`func NewIncidentUpdate(description string, ) *IncidentUpdate`

NewIncidentUpdate instantiates a new IncidentUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentUpdateWithDefaults

`func NewIncidentUpdateWithDefaults() *IncidentUpdate`

NewIncidentUpdateWithDefaults instantiates a new IncidentUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *IncidentUpdate) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IncidentUpdate) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IncidentUpdate) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IncidentUpdate) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetCategory

`func (o *IncidentUpdate) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *IncidentUpdate) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *IncidentUpdate) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *IncidentUpdate) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *IncidentUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IncidentUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IncidentUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetImpact

`func (o *IncidentUpdate) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *IncidentUpdate) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *IncidentUpdate) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *IncidentUpdate) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetPriority

`func (o *IncidentUpdate) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *IncidentUpdate) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *IncidentUpdate) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *IncidentUpdate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSeverity

`func (o *IncidentUpdate) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *IncidentUpdate) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *IncidentUpdate) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *IncidentUpdate) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetState

`func (o *IncidentUpdate) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IncidentUpdate) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IncidentUpdate) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IncidentUpdate) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubcategory

`func (o *IncidentUpdate) GetSubcategory() string`

GetSubcategory returns the Subcategory field if non-nil, zero value otherwise.

### GetSubcategoryOk

`func (o *IncidentUpdate) GetSubcategoryOk() (*string, bool)`

GetSubcategoryOk returns a tuple with the Subcategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubcategory

`func (o *IncidentUpdate) SetSubcategory(v string)`

SetSubcategory sets Subcategory field to given value.

### HasSubcategory

`func (o *IncidentUpdate) HasSubcategory() bool`

HasSubcategory returns a boolean if a field has been set.

### GetTenant

`func (o *IncidentUpdate) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *IncidentUpdate) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *IncidentUpdate) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *IncidentUpdate) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *IncidentUpdate) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *IncidentUpdate) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetUrgency

`func (o *IncidentUpdate) GetUrgency() string`

GetUrgency returns the Urgency field if non-nil, zero value otherwise.

### GetUrgencyOk

`func (o *IncidentUpdate) GetUrgencyOk() (*string, bool)`

GetUrgencyOk returns a tuple with the Urgency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrgency

`func (o *IncidentUpdate) SetUrgency(v string)`

SetUrgency sets Urgency field to given value.

### HasUrgency

`func (o *IncidentUpdate) HasUrgency() bool`

HasUrgency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


