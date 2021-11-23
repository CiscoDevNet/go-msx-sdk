# Incident

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ExternalId** | **string** |  | [readonly] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**Notes** | Pointer to **[]string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] 
**ModifiedBy** | Pointer to **string** |  | [optional] 
**ModifiedOn** | Pointer to **time.Time** |  | [optional] 
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

### NewIncident

`func NewIncident(id string, externalId string, description string, ) *Incident`

NewIncident instantiates a new Incident object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentWithDefaults

`func NewIncidentWithDefaults() *Incident`

NewIncidentWithDefaults instantiates a new Incident object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Incident) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Incident) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Incident) SetId(v string)`

SetId sets Id field to given value.


### GetExternalId

`func (o *Incident) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Incident) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Incident) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetAttributes

`func (o *Incident) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Incident) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Incident) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Incident) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetNotes

`func (o *Incident) GetNotes() []string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Incident) GetNotesOk() (*[]string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Incident) SetNotes(v []string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Incident) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Incident) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Incident) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Incident) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Incident) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedOn

`func (o *Incident) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Incident) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Incident) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *Incident) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetModifiedBy

`func (o *Incident) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Incident) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Incident) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *Incident) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetModifiedOn

`func (o *Incident) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *Incident) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *Incident) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *Incident) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetCategory

`func (o *Incident) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Incident) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Incident) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Incident) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *Incident) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Incident) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Incident) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetImpact

`func (o *Incident) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *Incident) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *Incident) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *Incident) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetPriority

`func (o *Incident) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Incident) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Incident) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Incident) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSeverity

`func (o *Incident) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Incident) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Incident) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *Incident) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetState

`func (o *Incident) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Incident) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Incident) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Incident) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubcategory

`func (o *Incident) GetSubcategory() string`

GetSubcategory returns the Subcategory field if non-nil, zero value otherwise.

### GetSubcategoryOk

`func (o *Incident) GetSubcategoryOk() (*string, bool)`

GetSubcategoryOk returns a tuple with the Subcategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubcategory

`func (o *Incident) SetSubcategory(v string)`

SetSubcategory sets Subcategory field to given value.

### HasSubcategory

`func (o *Incident) HasSubcategory() bool`

HasSubcategory returns a boolean if a field has been set.

### GetTenant

`func (o *Incident) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *Incident) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *Incident) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *Incident) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *Incident) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *Incident) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetUrgency

`func (o *Incident) GetUrgency() string`

GetUrgency returns the Urgency field if non-nil, zero value otherwise.

### GetUrgencyOk

`func (o *Incident) GetUrgencyOk() (*string, bool)`

GetUrgencyOk returns a tuple with the Urgency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrgency

`func (o *Incident) SetUrgency(v string)`

SetUrgency sets Urgency field to given value.

### HasUrgency

`func (o *Incident) HasUrgency() bool`

HasUrgency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


