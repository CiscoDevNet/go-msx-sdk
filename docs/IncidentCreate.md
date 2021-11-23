# IncidentCreate

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

### NewIncidentCreate

`func NewIncidentCreate(description string, ) *IncidentCreate`

NewIncidentCreate instantiates a new IncidentCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentCreateWithDefaults

`func NewIncidentCreateWithDefaults() *IncidentCreate`

NewIncidentCreateWithDefaults instantiates a new IncidentCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *IncidentCreate) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IncidentCreate) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IncidentCreate) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IncidentCreate) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetCategory

`func (o *IncidentCreate) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *IncidentCreate) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *IncidentCreate) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *IncidentCreate) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *IncidentCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IncidentCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IncidentCreate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetImpact

`func (o *IncidentCreate) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *IncidentCreate) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *IncidentCreate) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *IncidentCreate) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetPriority

`func (o *IncidentCreate) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *IncidentCreate) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *IncidentCreate) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *IncidentCreate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSeverity

`func (o *IncidentCreate) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *IncidentCreate) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *IncidentCreate) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *IncidentCreate) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetState

`func (o *IncidentCreate) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IncidentCreate) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IncidentCreate) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IncidentCreate) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubcategory

`func (o *IncidentCreate) GetSubcategory() string`

GetSubcategory returns the Subcategory field if non-nil, zero value otherwise.

### GetSubcategoryOk

`func (o *IncidentCreate) GetSubcategoryOk() (*string, bool)`

GetSubcategoryOk returns a tuple with the Subcategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubcategory

`func (o *IncidentCreate) SetSubcategory(v string)`

SetSubcategory sets Subcategory field to given value.

### HasSubcategory

`func (o *IncidentCreate) HasSubcategory() bool`

HasSubcategory returns a boolean if a field has been set.

### GetTenant

`func (o *IncidentCreate) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *IncidentCreate) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *IncidentCreate) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *IncidentCreate) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *IncidentCreate) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *IncidentCreate) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetUrgency

`func (o *IncidentCreate) GetUrgency() string`

GetUrgency returns the Urgency field if non-nil, zero value otherwise.

### GetUrgencyOk

`func (o *IncidentCreate) GetUrgencyOk() (*string, bool)`

GetUrgencyOk returns a tuple with the Urgency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrgency

`func (o *IncidentCreate) SetUrgency(v string)`

SetUrgency sets Urgency field to given value.

### HasUrgency

`func (o *IncidentCreate) HasUrgency() bool`

HasUrgency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


