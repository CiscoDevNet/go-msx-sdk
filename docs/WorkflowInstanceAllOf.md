# WorkflowInstanceAllOf

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

## Methods

### NewWorkflowInstanceAllOf

`func NewWorkflowInstanceAllOf() *WorkflowInstanceAllOf`

NewWorkflowInstanceAllOf instantiates a new WorkflowInstanceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowInstanceAllOfWithDefaults

`func NewWorkflowInstanceAllOfWithDefaults() *WorkflowInstanceAllOf`

NewWorkflowInstanceAllOfWithDefaults instantiates a new WorkflowInstanceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowInstanceAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowInstanceAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowInstanceAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowInstanceAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDefinitionId

`func (o *WorkflowInstanceAllOf) GetDefinitionId() string`

GetDefinitionId returns the DefinitionId field if non-nil, zero value otherwise.

### GetDefinitionIdOk

`func (o *WorkflowInstanceAllOf) GetDefinitionIdOk() (*string, bool)`

GetDefinitionIdOk returns a tuple with the DefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionId

`func (o *WorkflowInstanceAllOf) SetDefinitionId(v string)`

SetDefinitionId sets DefinitionId field to given value.

### HasDefinitionId

`func (o *WorkflowInstanceAllOf) HasDefinitionId() bool`

HasDefinitionId returns a boolean if a field has been set.

### GetName

`func (o *WorkflowInstanceAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowInstanceAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowInstanceAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowInstanceAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchemaId

`func (o *WorkflowInstanceAllOf) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *WorkflowInstanceAllOf) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *WorkflowInstanceAllOf) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *WorkflowInstanceAllOf) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetVersion

`func (o *WorkflowInstanceAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowInstanceAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowInstanceAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkflowInstanceAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetType

`func (o *WorkflowInstanceAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowInstanceAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowInstanceAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowInstanceAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBaseType

`func (o *WorkflowInstanceAllOf) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *WorkflowInstanceAllOf) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *WorkflowInstanceAllOf) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *WorkflowInstanceAllOf) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetProperties

`func (o *WorkflowInstanceAllOf) GetProperties() map[string]map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowInstanceAllOf) GetPropertiesOk() (*map[string]map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowInstanceAllOf) SetProperties(v map[string]map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowInstanceAllOf) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetActions

`func (o *WorkflowInstanceAllOf) GetActions() []WorkflowAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *WorkflowInstanceAllOf) GetActionsOk() (*[]WorkflowAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *WorkflowInstanceAllOf) SetActions(v []WorkflowAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *WorkflowInstanceAllOf) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetVariables

`func (o *WorkflowInstanceAllOf) GetVariables() []WorkflowVariable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *WorkflowInstanceAllOf) GetVariablesOk() (*[]WorkflowVariable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *WorkflowInstanceAllOf) SetVariables(v []WorkflowVariable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *WorkflowInstanceAllOf) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowInstanceAllOf) GetStatus() map[string]map[string]interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowInstanceAllOf) GetStatusOk() (*map[string]map[string]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowInstanceAllOf) SetStatus(v map[string]map[string]interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowInstanceAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartedOn

`func (o *WorkflowInstanceAllOf) GetStartedOn() string`

GetStartedOn returns the StartedOn field if non-nil, zero value otherwise.

### GetStartedOnOk

`func (o *WorkflowInstanceAllOf) GetStartedOnOk() (*string, bool)`

GetStartedOnOk returns a tuple with the StartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOn

`func (o *WorkflowInstanceAllOf) SetStartedOn(v string)`

SetStartedOn sets StartedOn field to given value.

### HasStartedOn

`func (o *WorkflowInstanceAllOf) HasStartedOn() bool`

HasStartedOn returns a boolean if a field has been set.

### GetEndedOn

`func (o *WorkflowInstanceAllOf) GetEndedOn() string`

GetEndedOn returns the EndedOn field if non-nil, zero value otherwise.

### GetEndedOnOk

`func (o *WorkflowInstanceAllOf) GetEndedOnOk() (*string, bool)`

GetEndedOnOk returns a tuple with the EndedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedOn

`func (o *WorkflowInstanceAllOf) SetEndedOn(v string)`

SetEndedOn sets EndedOn field to given value.

### HasEndedOn

`func (o *WorkflowInstanceAllOf) HasEndedOn() bool`

HasEndedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


