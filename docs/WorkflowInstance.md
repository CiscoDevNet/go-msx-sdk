# WorkflowInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DefinitionId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SchemaId** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**BaseType** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Actions** | Pointer to [**[]WorkflowAction**](WorkflowAction.md) |  | [optional] 
**Variables** | Pointer to [**[]WorkflowVariable**](WorkflowVariable.md) |  | [optional] 
**Status** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**StartedOn** | Pointer to **string** |  | [optional] 
**EndedOn** | Pointer to **string** |  | [optional] 
**CreatedOn** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**UpdatedOn** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**UniqueName** | Pointer to **string** |  | [optional] 

## Methods

### NewWorkflowInstance

`func NewWorkflowInstance() *WorkflowInstance`

NewWorkflowInstance instantiates a new WorkflowInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowInstanceWithDefaults

`func NewWorkflowInstanceWithDefaults() *WorkflowInstance`

NewWorkflowInstanceWithDefaults instantiates a new WorkflowInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowInstance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDefinitionId

`func (o *WorkflowInstance) GetDefinitionId() string`

GetDefinitionId returns the DefinitionId field if non-nil, zero value otherwise.

### GetDefinitionIdOk

`func (o *WorkflowInstance) GetDefinitionIdOk() (*string, bool)`

GetDefinitionIdOk returns a tuple with the DefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionId

`func (o *WorkflowInstance) SetDefinitionId(v string)`

SetDefinitionId sets DefinitionId field to given value.

### HasDefinitionId

`func (o *WorkflowInstance) HasDefinitionId() bool`

HasDefinitionId returns a boolean if a field has been set.

### GetName

`func (o *WorkflowInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchemaId

`func (o *WorkflowInstance) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *WorkflowInstance) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *WorkflowInstance) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *WorkflowInstance) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetVersion

`func (o *WorkflowInstance) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowInstance) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowInstance) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkflowInstance) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetType

`func (o *WorkflowInstance) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowInstance) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowInstance) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowInstance) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBaseType

`func (o *WorkflowInstance) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *WorkflowInstance) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *WorkflowInstance) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *WorkflowInstance) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetProperties

`func (o *WorkflowInstance) GetProperties() map[string]map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowInstance) GetPropertiesOk() (*map[string]map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowInstance) SetProperties(v map[string]map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowInstance) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetActions

`func (o *WorkflowInstance) GetActions() []WorkflowAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *WorkflowInstance) GetActionsOk() (*[]WorkflowAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *WorkflowInstance) SetActions(v []WorkflowAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *WorkflowInstance) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetVariables

`func (o *WorkflowInstance) GetVariables() []WorkflowVariable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *WorkflowInstance) GetVariablesOk() (*[]WorkflowVariable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *WorkflowInstance) SetVariables(v []WorkflowVariable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *WorkflowInstance) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowInstance) GetStatus() map[string]map[string]interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowInstance) GetStatusOk() (*map[string]map[string]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowInstance) SetStatus(v map[string]map[string]interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowInstance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartedOn

`func (o *WorkflowInstance) GetStartedOn() string`

GetStartedOn returns the StartedOn field if non-nil, zero value otherwise.

### GetStartedOnOk

`func (o *WorkflowInstance) GetStartedOnOk() (*string, bool)`

GetStartedOnOk returns a tuple with the StartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOn

`func (o *WorkflowInstance) SetStartedOn(v string)`

SetStartedOn sets StartedOn field to given value.

### HasStartedOn

`func (o *WorkflowInstance) HasStartedOn() bool`

HasStartedOn returns a boolean if a field has been set.

### GetEndedOn

`func (o *WorkflowInstance) GetEndedOn() string`

GetEndedOn returns the EndedOn field if non-nil, zero value otherwise.

### GetEndedOnOk

`func (o *WorkflowInstance) GetEndedOnOk() (*string, bool)`

GetEndedOnOk returns a tuple with the EndedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedOn

`func (o *WorkflowInstance) SetEndedOn(v string)`

SetEndedOn sets EndedOn field to given value.

### HasEndedOn

`func (o *WorkflowInstance) HasEndedOn() bool`

HasEndedOn returns a boolean if a field has been set.

### GetCreatedOn

`func (o *WorkflowInstance) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *WorkflowInstance) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *WorkflowInstance) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *WorkflowInstance) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetCreatedBy

`func (o *WorkflowInstance) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WorkflowInstance) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WorkflowInstance) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *WorkflowInstance) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *WorkflowInstance) GetUpdatedOn() string`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *WorkflowInstance) GetUpdatedOnOk() (*string, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *WorkflowInstance) SetUpdatedOn(v string)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *WorkflowInstance) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *WorkflowInstance) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *WorkflowInstance) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *WorkflowInstance) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *WorkflowInstance) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetOwner

`func (o *WorkflowInstance) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *WorkflowInstance) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *WorkflowInstance) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *WorkflowInstance) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetUniqueName

`func (o *WorkflowInstance) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *WorkflowInstance) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *WorkflowInstance) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *WorkflowInstance) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


