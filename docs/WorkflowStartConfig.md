# WorkflowStartConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputVariables** | Pointer to [**[]WorkflowVariable**](WorkflowVariable.md) |  | [optional] 
**TypeOfTargetNeeded** | Pointer to **NullableString** |  | [optional] 
**TargetId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWorkflowStartConfig

`func NewWorkflowStartConfig() *WorkflowStartConfig`

NewWorkflowStartConfig instantiates a new WorkflowStartConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowStartConfigWithDefaults

`func NewWorkflowStartConfigWithDefaults() *WorkflowStartConfig`

NewWorkflowStartConfigWithDefaults instantiates a new WorkflowStartConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputVariables

`func (o *WorkflowStartConfig) GetInputVariables() []WorkflowVariable`

GetInputVariables returns the InputVariables field if non-nil, zero value otherwise.

### GetInputVariablesOk

`func (o *WorkflowStartConfig) GetInputVariablesOk() (*[]WorkflowVariable, bool)`

GetInputVariablesOk returns a tuple with the InputVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputVariables

`func (o *WorkflowStartConfig) SetInputVariables(v []WorkflowVariable)`

SetInputVariables sets InputVariables field to given value.

### HasInputVariables

`func (o *WorkflowStartConfig) HasInputVariables() bool`

HasInputVariables returns a boolean if a field has been set.

### GetTypeOfTargetNeeded

`func (o *WorkflowStartConfig) GetTypeOfTargetNeeded() string`

GetTypeOfTargetNeeded returns the TypeOfTargetNeeded field if non-nil, zero value otherwise.

### GetTypeOfTargetNeededOk

`func (o *WorkflowStartConfig) GetTypeOfTargetNeededOk() (*string, bool)`

GetTypeOfTargetNeededOk returns a tuple with the TypeOfTargetNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOfTargetNeeded

`func (o *WorkflowStartConfig) SetTypeOfTargetNeeded(v string)`

SetTypeOfTargetNeeded sets TypeOfTargetNeeded field to given value.

### HasTypeOfTargetNeeded

`func (o *WorkflowStartConfig) HasTypeOfTargetNeeded() bool`

HasTypeOfTargetNeeded returns a boolean if a field has been set.

### SetTypeOfTargetNeededNil

`func (o *WorkflowStartConfig) SetTypeOfTargetNeededNil(b bool)`

 SetTypeOfTargetNeededNil sets the value for TypeOfTargetNeeded to be an explicit nil

### UnsetTypeOfTargetNeeded
`func (o *WorkflowStartConfig) UnsetTypeOfTargetNeeded()`

UnsetTypeOfTargetNeeded ensures that no value is present for TypeOfTargetNeeded, not even an explicit nil
### GetTargetId

`func (o *WorkflowStartConfig) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *WorkflowStartConfig) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *WorkflowStartConfig) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *WorkflowStartConfig) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *WorkflowStartConfig) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *WorkflowStartConfig) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


