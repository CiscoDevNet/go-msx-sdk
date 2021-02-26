# WorkflowAccessMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adapter** | Pointer to [**WorkflowAccessMetaType**](WorkflowAccessMetaType.md) |  | [optional] 
**RuntimeUsers** | Pointer to [**[]WorkflowAccessMetaType**](WorkflowAccessMetaType.md) |  | [optional] 
**Targets** | Pointer to [**[]WorkflowAccessMetaType**](WorkflowAccessMetaType.md) |  | [optional] 
**IsIntegration** | Pointer to **bool** |  | [optional] 
**IsInternal** | Pointer to **bool** |  | [optional] 

## Methods

### NewWorkflowAccessMeta

`func NewWorkflowAccessMeta() *WorkflowAccessMeta`

NewWorkflowAccessMeta instantiates a new WorkflowAccessMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowAccessMetaWithDefaults

`func NewWorkflowAccessMetaWithDefaults() *WorkflowAccessMeta`

NewWorkflowAccessMetaWithDefaults instantiates a new WorkflowAccessMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdapter

`func (o *WorkflowAccessMeta) GetAdapter() WorkflowAccessMetaType`

GetAdapter returns the Adapter field if non-nil, zero value otherwise.

### GetAdapterOk

`func (o *WorkflowAccessMeta) GetAdapterOk() (*WorkflowAccessMetaType, bool)`

GetAdapterOk returns a tuple with the Adapter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapter

`func (o *WorkflowAccessMeta) SetAdapter(v WorkflowAccessMetaType)`

SetAdapter sets Adapter field to given value.

### HasAdapter

`func (o *WorkflowAccessMeta) HasAdapter() bool`

HasAdapter returns a boolean if a field has been set.

### GetRuntimeUsers

`func (o *WorkflowAccessMeta) GetRuntimeUsers() []WorkflowAccessMetaType`

GetRuntimeUsers returns the RuntimeUsers field if non-nil, zero value otherwise.

### GetRuntimeUsersOk

`func (o *WorkflowAccessMeta) GetRuntimeUsersOk() (*[]WorkflowAccessMetaType, bool)`

GetRuntimeUsersOk returns a tuple with the RuntimeUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeUsers

`func (o *WorkflowAccessMeta) SetRuntimeUsers(v []WorkflowAccessMetaType)`

SetRuntimeUsers sets RuntimeUsers field to given value.

### HasRuntimeUsers

`func (o *WorkflowAccessMeta) HasRuntimeUsers() bool`

HasRuntimeUsers returns a boolean if a field has been set.

### GetTargets

`func (o *WorkflowAccessMeta) GetTargets() []WorkflowAccessMetaType`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *WorkflowAccessMeta) GetTargetsOk() (*[]WorkflowAccessMetaType, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *WorkflowAccessMeta) SetTargets(v []WorkflowAccessMetaType)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *WorkflowAccessMeta) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetIsIntegration

`func (o *WorkflowAccessMeta) GetIsIntegration() bool`

GetIsIntegration returns the IsIntegration field if non-nil, zero value otherwise.

### GetIsIntegrationOk

`func (o *WorkflowAccessMeta) GetIsIntegrationOk() (*bool, bool)`

GetIsIntegrationOk returns a tuple with the IsIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIntegration

`func (o *WorkflowAccessMeta) SetIsIntegration(v bool)`

SetIsIntegration sets IsIntegration field to given value.

### HasIsIntegration

`func (o *WorkflowAccessMeta) HasIsIntegration() bool`

HasIsIntegration returns a boolean if a field has been set.

### GetIsInternal

`func (o *WorkflowAccessMeta) GetIsInternal() bool`

GetIsInternal returns the IsInternal field if non-nil, zero value otherwise.

### GetIsInternalOk

`func (o *WorkflowAccessMeta) GetIsInternalOk() (*bool, bool)`

GetIsInternalOk returns a tuple with the IsInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternal

`func (o *WorkflowAccessMeta) SetIsInternal(v bool)`

SetIsInternal sets IsInternal field to given value.

### HasIsInternal

`func (o *WorkflowAccessMeta) HasIsInternal() bool`

HasIsInternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


