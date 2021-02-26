# WorkflowAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**WorkflowId** | Pointer to **string** |  | [optional] 
**DefinitionId** | Pointer to **string** |  | [optional] 
**SchemaId** | Pointer to **string** |  | [optional] 
**AdapterId** | Pointer to **string** |  | [optional] 
**UniqueName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**BaseType** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**StartedBy** | Pointer to **string** |  | [optional] 
**Blocks** | Pointer to [**[]WorkflowActionBlock**](WorkflowActionBlock.md) |  | [optional] 
**Status** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Output** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**StartedOn** | Pointer to **string** |  | [optional] 
**EndedOn** | Pointer to **string** |  | [optional] 

## Methods

### NewWorkflowAction

`func NewWorkflowAction() *WorkflowAction`

NewWorkflowAction instantiates a new WorkflowAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowActionWithDefaults

`func NewWorkflowActionWithDefaults() *WorkflowAction`

NewWorkflowActionWithDefaults instantiates a new WorkflowAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowAction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowAction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowAction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowAction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWorkflowId

`func (o *WorkflowAction) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowAction) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowAction) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *WorkflowAction) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetDefinitionId

`func (o *WorkflowAction) GetDefinitionId() string`

GetDefinitionId returns the DefinitionId field if non-nil, zero value otherwise.

### GetDefinitionIdOk

`func (o *WorkflowAction) GetDefinitionIdOk() (*string, bool)`

GetDefinitionIdOk returns a tuple with the DefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionId

`func (o *WorkflowAction) SetDefinitionId(v string)`

SetDefinitionId sets DefinitionId field to given value.

### HasDefinitionId

`func (o *WorkflowAction) HasDefinitionId() bool`

HasDefinitionId returns a boolean if a field has been set.

### GetSchemaId

`func (o *WorkflowAction) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *WorkflowAction) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *WorkflowAction) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *WorkflowAction) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetAdapterId

`func (o *WorkflowAction) GetAdapterId() string`

GetAdapterId returns the AdapterId field if non-nil, zero value otherwise.

### GetAdapterIdOk

`func (o *WorkflowAction) GetAdapterIdOk() (*string, bool)`

GetAdapterIdOk returns a tuple with the AdapterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterId

`func (o *WorkflowAction) SetAdapterId(v string)`

SetAdapterId sets AdapterId field to given value.

### HasAdapterId

`func (o *WorkflowAction) HasAdapterId() bool`

HasAdapterId returns a boolean if a field has been set.

### GetUniqueName

`func (o *WorkflowAction) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *WorkflowAction) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *WorkflowAction) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *WorkflowAction) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### GetName

`func (o *WorkflowAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowAction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowAction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *WorkflowAction) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WorkflowAction) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WorkflowAction) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WorkflowAction) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *WorkflowAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowAction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowAction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBaseType

`func (o *WorkflowAction) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *WorkflowAction) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *WorkflowAction) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *WorkflowAction) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetProperties

`func (o *WorkflowAction) GetProperties() map[string]map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowAction) GetPropertiesOk() (*map[string]map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowAction) SetProperties(v map[string]map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowAction) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetObjectType

`func (o *WorkflowAction) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowAction) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowAction) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *WorkflowAction) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetStartedBy

`func (o *WorkflowAction) GetStartedBy() string`

GetStartedBy returns the StartedBy field if non-nil, zero value otherwise.

### GetStartedByOk

`func (o *WorkflowAction) GetStartedByOk() (*string, bool)`

GetStartedByOk returns a tuple with the StartedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedBy

`func (o *WorkflowAction) SetStartedBy(v string)`

SetStartedBy sets StartedBy field to given value.

### HasStartedBy

`func (o *WorkflowAction) HasStartedBy() bool`

HasStartedBy returns a boolean if a field has been set.

### GetBlocks

`func (o *WorkflowAction) GetBlocks() []WorkflowActionBlock`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *WorkflowAction) GetBlocksOk() (*[]WorkflowActionBlock, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *WorkflowAction) SetBlocks(v []WorkflowActionBlock)`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *WorkflowAction) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowAction) GetStatus() map[string]map[string]interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowAction) GetStatusOk() (*map[string]map[string]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowAction) SetStatus(v map[string]map[string]interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowAction) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowAction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowAction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOutput

`func (o *WorkflowAction) GetOutput() map[string]map[string]interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *WorkflowAction) GetOutputOk() (*map[string]map[string]interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *WorkflowAction) SetOutput(v map[string]map[string]interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *WorkflowAction) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetStartedOn

`func (o *WorkflowAction) GetStartedOn() string`

GetStartedOn returns the StartedOn field if non-nil, zero value otherwise.

### GetStartedOnOk

`func (o *WorkflowAction) GetStartedOnOk() (*string, bool)`

GetStartedOnOk returns a tuple with the StartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOn

`func (o *WorkflowAction) SetStartedOn(v string)`

SetStartedOn sets StartedOn field to given value.

### HasStartedOn

`func (o *WorkflowAction) HasStartedOn() bool`

HasStartedOn returns a boolean if a field has been set.

### GetEndedOn

`func (o *WorkflowAction) GetEndedOn() string`

GetEndedOn returns the EndedOn field if non-nil, zero value otherwise.

### GetEndedOnOk

`func (o *WorkflowAction) GetEndedOnOk() (*string, bool)`

GetEndedOnOk returns a tuple with the EndedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedOn

`func (o *WorkflowAction) SetEndedOn(v string)`

SetEndedOn sets EndedOn field to given value.

### HasEndedOn

`func (o *WorkflowAction) HasEndedOn() bool`

HasEndedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


