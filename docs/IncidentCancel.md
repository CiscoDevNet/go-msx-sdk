# IncidentCancel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notes** | Pointer to **string** |  | [optional] 
**Tenant** | **string** |  | 

## Methods

### NewIncidentCancel

`func NewIncidentCancel(tenant string, ) *IncidentCancel`

NewIncidentCancel instantiates a new IncidentCancel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentCancelWithDefaults

`func NewIncidentCancelWithDefaults() *IncidentCancel`

NewIncidentCancelWithDefaults instantiates a new IncidentCancel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotes

`func (o *IncidentCancel) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *IncidentCancel) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *IncidentCancel) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *IncidentCancel) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetTenant

`func (o *IncidentCancel) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *IncidentCancel) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *IncidentCancel) SetTenant(v string)`

SetTenant sets Tenant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


