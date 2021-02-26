# WorkflowActionBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UniqueName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**BaseType** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**Actions** | Pointer to [**[]WorkflowAction**](WorkflowAction.md) |  | [optional] 
**SubworkflowId** | Pointer to **string** |  | [optional] 
**WorkflowId** | Pointer to **string** |  | [optional] 
**SchemaId** | Pointer to **string** |  | [optional] 

## Methods

### NewWorkflowActionBlock

`func NewWorkflowActionBlock() *WorkflowActionBlock`

NewWorkflowActionBlock instantiates a new WorkflowActionBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowActionBlockWithDefaults

`func NewWorkflowActionBlockWithDefaults() *WorkflowActionBlock`

NewWorkflowActionBlockWithDefaults instantiates a new WorkflowActionBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUniqueName

`func (o *WorkflowActionBlock) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *WorkflowActionBlock) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *WorkflowActionBlock) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *WorkflowActionBlock) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### GetName

`func (o *WorkflowActionBlock) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowActionBlock) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowActionBlock) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowActionBlock) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *WorkflowActionBlock) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WorkflowActionBlock) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WorkflowActionBlock) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WorkflowActionBlock) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *WorkflowActionBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowActionBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowActionBlock) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowActionBlock) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *WorkflowActionBlock) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowActionBlock) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowActionBlock) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkflowActionBlock) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBaseType

`func (o *WorkflowActionBlock) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *WorkflowActionBlock) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *WorkflowActionBlock) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *WorkflowActionBlock) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetProperties

`func (o *WorkflowActionBlock) GetProperties() map[string]map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowActionBlock) GetPropertiesOk() (*map[string]map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowActionBlock) SetProperties(v map[string]map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowActionBlock) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetObjectType

`func (o *WorkflowActionBlock) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowActionBlock) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowActionBlock) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *WorkflowActionBlock) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetActions

`func (o *WorkflowActionBlock) GetActions() []WorkflowAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *WorkflowActionBlock) GetActionsOk() (*[]WorkflowAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *WorkflowActionBlock) SetActions(v []WorkflowAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *WorkflowActionBlock) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetSubworkflowId

`func (o *WorkflowActionBlock) GetSubworkflowId() string`

GetSubworkflowId returns the SubworkflowId field if non-nil, zero value otherwise.

### GetSubworkflowIdOk

`func (o *WorkflowActionBlock) GetSubworkflowIdOk() (*string, bool)`

GetSubworkflowIdOk returns a tuple with the SubworkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubworkflowId

`func (o *WorkflowActionBlock) SetSubworkflowId(v string)`

SetSubworkflowId sets SubworkflowId field to given value.

### HasSubworkflowId

`func (o *WorkflowActionBlock) HasSubworkflowId() bool`

HasSubworkflowId returns a boolean if a field has been set.

### GetWorkflowId

`func (o *WorkflowActionBlock) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowActionBlock) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowActionBlock) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *WorkflowActionBlock) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetSchemaId

`func (o *WorkflowActionBlock) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *WorkflowActionBlock) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *WorkflowActionBlock) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *WorkflowActionBlock) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


