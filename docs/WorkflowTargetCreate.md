# WorkflowTargetCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**SchemaId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Properties** | **map[string]map[string]interface{}** |  | 
**UniqueName** | Pointer to **string** |  | [optional] 

## Methods

### NewWorkflowTargetCreate

`func NewWorkflowTargetCreate(name string, properties map[string]map[string]interface{}, ) *WorkflowTargetCreate`

NewWorkflowTargetCreate instantiates a new WorkflowTargetCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTargetCreateWithDefaults

`func NewWorkflowTargetCreateWithDefaults() *WorkflowTargetCreate`

NewWorkflowTargetCreateWithDefaults instantiates a new WorkflowTargetCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WorkflowTargetCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowTargetCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowTargetCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WorkflowTargetCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowTargetCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowTargetCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowTargetCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTitle

`func (o *WorkflowTargetCreate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WorkflowTargetCreate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WorkflowTargetCreate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WorkflowTargetCreate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSchemaId

`func (o *WorkflowTargetCreate) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *WorkflowTargetCreate) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *WorkflowTargetCreate) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *WorkflowTargetCreate) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetType

`func (o *WorkflowTargetCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowTargetCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowTargetCreate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowTargetCreate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProperties

`func (o *WorkflowTargetCreate) GetProperties() map[string]map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowTargetCreate) GetPropertiesOk() (*map[string]map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowTargetCreate) SetProperties(v map[string]map[string]interface{})`

SetProperties sets Properties field to given value.


### GetUniqueName

`func (o *WorkflowTargetCreate) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *WorkflowTargetCreate) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *WorkflowTargetCreate) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *WorkflowTargetCreate) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


