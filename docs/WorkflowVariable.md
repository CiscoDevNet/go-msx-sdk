# WorkflowVariable

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
**CreatedOn** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**UpdatedOn** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**Owner** | Pointer to **NullableString** |  | [optional] 
**UniqueName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWorkflowVariable

`func NewWorkflowVariable() *WorkflowVariable`

NewWorkflowVariable instantiates a new WorkflowVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowVariableWithDefaults

`func NewWorkflowVariableWithDefaults() *WorkflowVariable`

NewWorkflowVariableWithDefaults instantiates a new WorkflowVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowVariable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowVariable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowVariable) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowVariable) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *WorkflowVariable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowVariable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowVariable) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowVariable) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *WorkflowVariable) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *WorkflowVariable) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetBaseType

`func (o *WorkflowVariable) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *WorkflowVariable) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *WorkflowVariable) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *WorkflowVariable) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### SetBaseTypeNil

`func (o *WorkflowVariable) SetBaseTypeNil(b bool)`

 SetBaseTypeNil sets the value for BaseType to be an explicit nil

### UnsetBaseType
`func (o *WorkflowVariable) UnsetBaseType()`

UnsetBaseType ensures that no value is present for BaseType, not even an explicit nil
### GetSchemaId

`func (o *WorkflowVariable) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *WorkflowVariable) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *WorkflowVariable) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *WorkflowVariable) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetObjectType

`func (o *WorkflowVariable) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowVariable) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowVariable) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *WorkflowVariable) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetProperties

`func (o *WorkflowVariable) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowVariable) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowVariable) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowVariable) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetDataType

`func (o *WorkflowVariable) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *WorkflowVariable) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *WorkflowVariable) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *WorkflowVariable) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### SetDataTypeNil

`func (o *WorkflowVariable) SetDataTypeNil(b bool)`

 SetDataTypeNil sets the value for DataType to be an explicit nil

### UnsetDataType
`func (o *WorkflowVariable) UnsetDataType()`

UnsetDataType ensures that no value is present for DataType, not even an explicit nil
### GetScope

`func (o *WorkflowVariable) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *WorkflowVariable) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *WorkflowVariable) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *WorkflowVariable) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *WorkflowVariable) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *WorkflowVariable) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetCreatedOn

`func (o *WorkflowVariable) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *WorkflowVariable) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *WorkflowVariable) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *WorkflowVariable) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetCreatedBy

`func (o *WorkflowVariable) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WorkflowVariable) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WorkflowVariable) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *WorkflowVariable) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *WorkflowVariable) GetUpdatedOn() string`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *WorkflowVariable) GetUpdatedOnOk() (*string, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *WorkflowVariable) SetUpdatedOn(v string)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *WorkflowVariable) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *WorkflowVariable) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *WorkflowVariable) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *WorkflowVariable) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *WorkflowVariable) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *WorkflowVariable) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *WorkflowVariable) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetOwner

`func (o *WorkflowVariable) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *WorkflowVariable) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *WorkflowVariable) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *WorkflowVariable) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *WorkflowVariable) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *WorkflowVariable) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetUniqueName

`func (o *WorkflowVariable) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *WorkflowVariable) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *WorkflowVariable) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *WorkflowVariable) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### SetUniqueNameNil

`func (o *WorkflowVariable) SetUniqueNameNil(b bool)`

 SetUniqueNameNil sets the value for UniqueName to be an explicit nil

### UnsetUniqueName
`func (o *WorkflowVariable) UnsetUniqueName()`

UnsetUniqueName ensures that no value is present for UniqueName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


