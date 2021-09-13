# WorkflowSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**SchemaId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**BaseType** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Invisible** | Pointer to **bool** |  | [optional] 
**Inherits** | Pointer to **string** |  | [optional] 
**AccessMeta** | Pointer to [**WorkflowAccessMeta**](WorkflowAccessMeta.md) |  | [optional] 
**VariableSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**PropertySchema** | Pointer to **map[string]interface{}** |  | [optional] 
**OutputSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**ViewConfig** | Pointer to **map[string]interface{}** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedOn** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**UpdatedOn** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**Owner** | Pointer to **NullableString** |  | [optional] 
**UniqueName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWorkflowSchema

`func NewWorkflowSchema() *WorkflowSchema`

NewWorkflowSchema instantiates a new WorkflowSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSchemaWithDefaults

`func NewWorkflowSchemaWithDefaults() *WorkflowSchema`

NewWorkflowSchemaWithDefaults instantiates a new WorkflowSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowSchema) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowSchema) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *WorkflowSchema) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *WorkflowSchema) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetSchemaId

`func (o *WorkflowSchema) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *WorkflowSchema) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *WorkflowSchema) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *WorkflowSchema) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetName

`func (o *WorkflowSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *WorkflowSchema) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WorkflowSchema) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WorkflowSchema) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WorkflowSchema) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *WorkflowSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowSchema) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowSchema) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBaseType

`func (o *WorkflowSchema) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *WorkflowSchema) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *WorkflowSchema) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *WorkflowSchema) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetVersion

`func (o *WorkflowSchema) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowSchema) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowSchema) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkflowSchema) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetInvisible

`func (o *WorkflowSchema) GetInvisible() bool`

GetInvisible returns the Invisible field if non-nil, zero value otherwise.

### GetInvisibleOk

`func (o *WorkflowSchema) GetInvisibleOk() (*bool, bool)`

GetInvisibleOk returns a tuple with the Invisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvisible

`func (o *WorkflowSchema) SetInvisible(v bool)`

SetInvisible sets Invisible field to given value.

### HasInvisible

`func (o *WorkflowSchema) HasInvisible() bool`

HasInvisible returns a boolean if a field has been set.

### GetInherits

`func (o *WorkflowSchema) GetInherits() string`

GetInherits returns the Inherits field if non-nil, zero value otherwise.

### GetInheritsOk

`func (o *WorkflowSchema) GetInheritsOk() (*string, bool)`

GetInheritsOk returns a tuple with the Inherits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInherits

`func (o *WorkflowSchema) SetInherits(v string)`

SetInherits sets Inherits field to given value.

### HasInherits

`func (o *WorkflowSchema) HasInherits() bool`

HasInherits returns a boolean if a field has been set.

### GetAccessMeta

`func (o *WorkflowSchema) GetAccessMeta() WorkflowAccessMeta`

GetAccessMeta returns the AccessMeta field if non-nil, zero value otherwise.

### GetAccessMetaOk

`func (o *WorkflowSchema) GetAccessMetaOk() (*WorkflowAccessMeta, bool)`

GetAccessMetaOk returns a tuple with the AccessMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMeta

`func (o *WorkflowSchema) SetAccessMeta(v WorkflowAccessMeta)`

SetAccessMeta sets AccessMeta field to given value.

### HasAccessMeta

`func (o *WorkflowSchema) HasAccessMeta() bool`

HasAccessMeta returns a boolean if a field has been set.

### GetVariableSchema

`func (o *WorkflowSchema) GetVariableSchema() map[string]interface{}`

GetVariableSchema returns the VariableSchema field if non-nil, zero value otherwise.

### GetVariableSchemaOk

`func (o *WorkflowSchema) GetVariableSchemaOk() (*map[string]interface{}, bool)`

GetVariableSchemaOk returns a tuple with the VariableSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableSchema

`func (o *WorkflowSchema) SetVariableSchema(v map[string]interface{})`

SetVariableSchema sets VariableSchema field to given value.

### HasVariableSchema

`func (o *WorkflowSchema) HasVariableSchema() bool`

HasVariableSchema returns a boolean if a field has been set.

### GetPropertySchema

`func (o *WorkflowSchema) GetPropertySchema() map[string]interface{}`

GetPropertySchema returns the PropertySchema field if non-nil, zero value otherwise.

### GetPropertySchemaOk

`func (o *WorkflowSchema) GetPropertySchemaOk() (*map[string]interface{}, bool)`

GetPropertySchemaOk returns a tuple with the PropertySchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertySchema

`func (o *WorkflowSchema) SetPropertySchema(v map[string]interface{})`

SetPropertySchema sets PropertySchema field to given value.

### HasPropertySchema

`func (o *WorkflowSchema) HasPropertySchema() bool`

HasPropertySchema returns a boolean if a field has been set.

### GetOutputSchema

`func (o *WorkflowSchema) GetOutputSchema() map[string]interface{}`

GetOutputSchema returns the OutputSchema field if non-nil, zero value otherwise.

### GetOutputSchemaOk

`func (o *WorkflowSchema) GetOutputSchemaOk() (*map[string]interface{}, bool)`

GetOutputSchemaOk returns a tuple with the OutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchema

`func (o *WorkflowSchema) SetOutputSchema(v map[string]interface{})`

SetOutputSchema sets OutputSchema field to given value.

### HasOutputSchema

`func (o *WorkflowSchema) HasOutputSchema() bool`

HasOutputSchema returns a boolean if a field has been set.

### GetViewConfig

`func (o *WorkflowSchema) GetViewConfig() map[string]interface{}`

GetViewConfig returns the ViewConfig field if non-nil, zero value otherwise.

### GetViewConfigOk

`func (o *WorkflowSchema) GetViewConfigOk() (*map[string]interface{}, bool)`

GetViewConfigOk returns a tuple with the ViewConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewConfig

`func (o *WorkflowSchema) SetViewConfig(v map[string]interface{})`

SetViewConfig sets ViewConfig field to given value.

### HasViewConfig

`func (o *WorkflowSchema) HasViewConfig() bool`

HasViewConfig returns a boolean if a field has been set.

### GetAttributes

`func (o *WorkflowSchema) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WorkflowSchema) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WorkflowSchema) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WorkflowSchema) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetCreatedOn

`func (o *WorkflowSchema) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *WorkflowSchema) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *WorkflowSchema) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *WorkflowSchema) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetCreatedBy

`func (o *WorkflowSchema) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WorkflowSchema) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WorkflowSchema) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *WorkflowSchema) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *WorkflowSchema) GetUpdatedOn() string`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *WorkflowSchema) GetUpdatedOnOk() (*string, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *WorkflowSchema) SetUpdatedOn(v string)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *WorkflowSchema) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *WorkflowSchema) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *WorkflowSchema) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *WorkflowSchema) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *WorkflowSchema) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *WorkflowSchema) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *WorkflowSchema) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetOwner

`func (o *WorkflowSchema) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *WorkflowSchema) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *WorkflowSchema) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *WorkflowSchema) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *WorkflowSchema) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *WorkflowSchema) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetUniqueName

`func (o *WorkflowSchema) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *WorkflowSchema) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *WorkflowSchema) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *WorkflowSchema) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### SetUniqueNameNil

`func (o *WorkflowSchema) SetUniqueNameNil(b bool)`

 SetUniqueNameNil sets the value for UniqueName to be an explicit nil

### UnsetUniqueName
`func (o *WorkflowSchema) UnsetUniqueName()`

UnsetUniqueName ensures that no value is present for UniqueName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


