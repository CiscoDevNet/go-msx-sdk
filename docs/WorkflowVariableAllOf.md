# WorkflowVariableAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**BaseType** | Pointer to **NullableString** |  | [optional] 
**SchemaId** | Pointer to **string** |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 
**DataType** | Pointer to **NullableString** |  | [optional] 
**Scope** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWorkflowVariableAllOf

`func NewWorkflowVariableAllOf() *WorkflowVariableAllOf`

NewWorkflowVariableAllOf instantiates a new WorkflowVariableAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowVariableAllOfWithDefaults

`func NewWorkflowVariableAllOfWithDefaults() *WorkflowVariableAllOf`

NewWorkflowVariableAllOfWithDefaults instantiates a new WorkflowVariableAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowVariableAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowVariableAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowVariableAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowVariableAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *WorkflowVariableAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowVariableAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowVariableAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowVariableAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *WorkflowVariableAllOf) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *WorkflowVariableAllOf) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetBaseType

`func (o *WorkflowVariableAllOf) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *WorkflowVariableAllOf) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *WorkflowVariableAllOf) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *WorkflowVariableAllOf) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### SetBaseTypeNil

`func (o *WorkflowVariableAllOf) SetBaseTypeNil(b bool)`

 SetBaseTypeNil sets the value for BaseType to be an explicit nil

### UnsetBaseType
`func (o *WorkflowVariableAllOf) UnsetBaseType()`

UnsetBaseType ensures that no value is present for BaseType, not even an explicit nil
### GetSchemaId

`func (o *WorkflowVariableAllOf) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *WorkflowVariableAllOf) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *WorkflowVariableAllOf) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *WorkflowVariableAllOf) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetObjectType

`func (o *WorkflowVariableAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowVariableAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowVariableAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *WorkflowVariableAllOf) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetProperties

`func (o *WorkflowVariableAllOf) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowVariableAllOf) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowVariableAllOf) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowVariableAllOf) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetDataType

`func (o *WorkflowVariableAllOf) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *WorkflowVariableAllOf) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *WorkflowVariableAllOf) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *WorkflowVariableAllOf) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### SetDataTypeNil

`func (o *WorkflowVariableAllOf) SetDataTypeNil(b bool)`

 SetDataTypeNil sets the value for DataType to be an explicit nil

### UnsetDataType
`func (o *WorkflowVariableAllOf) UnsetDataType()`

UnsetDataType ensures that no value is present for DataType, not even an explicit nil
### GetScope

`func (o *WorkflowVariableAllOf) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *WorkflowVariableAllOf) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *WorkflowVariableAllOf) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *WorkflowVariableAllOf) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *WorkflowVariableAllOf) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *WorkflowVariableAllOf) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


