# WorkflowEventCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**TargetId** | **string** |  | 
**SchemaId** | **string** |  | 
**Properties** | **map[string]interface{}** |  | 

## Methods

### NewWorkflowEventCreate

`func NewWorkflowEventCreate(title string, targetId string, schemaId string, properties map[string]interface{}, ) *WorkflowEventCreate`

NewWorkflowEventCreate instantiates a new WorkflowEventCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowEventCreateWithDefaults

`func NewWorkflowEventCreateWithDefaults() *WorkflowEventCreate`

NewWorkflowEventCreateWithDefaults instantiates a new WorkflowEventCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *WorkflowEventCreate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WorkflowEventCreate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WorkflowEventCreate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *WorkflowEventCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowEventCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowEventCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowEventCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTargetId

`func (o *WorkflowEventCreate) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *WorkflowEventCreate) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *WorkflowEventCreate) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.


### GetSchemaId

`func (o *WorkflowEventCreate) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *WorkflowEventCreate) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *WorkflowEventCreate) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.


### GetProperties

`func (o *WorkflowEventCreate) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowEventCreate) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowEventCreate) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


