# WorkflowSchemaAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
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
**VariableSchema** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**PropertySchema** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**OutputSchema** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**ViewConfig** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Attributes** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewWorkflowSchemaAllOf

`func NewWorkflowSchemaAllOf() *WorkflowSchemaAllOf`

NewWorkflowSchemaAllOf instantiates a new WorkflowSchemaAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSchemaAllOfWithDefaults

`func NewWorkflowSchemaAllOfWithDefaults() *WorkflowSchemaAllOf`

NewWorkflowSchemaAllOfWithDefaults instantiates a new WorkflowSchemaAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowSchemaAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowSchemaAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowSchemaAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowSchemaAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSchemaId

`func (o *WorkflowSchemaAllOf) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *WorkflowSchemaAllOf) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *WorkflowSchemaAllOf) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *WorkflowSchemaAllOf) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetName

`func (o *WorkflowSchemaAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowSchemaAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowSchemaAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowSchemaAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *WorkflowSchemaAllOf) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WorkflowSchemaAllOf) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WorkflowSchemaAllOf) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WorkflowSchemaAllOf) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowSchemaAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowSchemaAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowSchemaAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowSchemaAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *WorkflowSchemaAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowSchemaAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowSchemaAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowSchemaAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBaseType

`func (o *WorkflowSchemaAllOf) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *WorkflowSchemaAllOf) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *WorkflowSchemaAllOf) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *WorkflowSchemaAllOf) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetVersion

`func (o *WorkflowSchemaAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowSchemaAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowSchemaAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkflowSchemaAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetInvisible

`func (o *WorkflowSchemaAllOf) GetInvisible() bool`

GetInvisible returns the Invisible field if non-nil, zero value otherwise.

### GetInvisibleOk

`func (o *WorkflowSchemaAllOf) GetInvisibleOk() (*bool, bool)`

GetInvisibleOk returns a tuple with the Invisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvisible

`func (o *WorkflowSchemaAllOf) SetInvisible(v bool)`

SetInvisible sets Invisible field to given value.

### HasInvisible

`func (o *WorkflowSchemaAllOf) HasInvisible() bool`

HasInvisible returns a boolean if a field has been set.

### GetInherits

`func (o *WorkflowSchemaAllOf) GetInherits() string`

GetInherits returns the Inherits field if non-nil, zero value otherwise.

### GetInheritsOk

`func (o *WorkflowSchemaAllOf) GetInheritsOk() (*string, bool)`

GetInheritsOk returns a tuple with the Inherits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInherits

`func (o *WorkflowSchemaAllOf) SetInherits(v string)`

SetInherits sets Inherits field to given value.

### HasInherits

`func (o *WorkflowSchemaAllOf) HasInherits() bool`

HasInherits returns a boolean if a field has been set.

### GetAccessMeta

`func (o *WorkflowSchemaAllOf) GetAccessMeta() WorkflowAccessMeta`

GetAccessMeta returns the AccessMeta field if non-nil, zero value otherwise.

### GetAccessMetaOk

`func (o *WorkflowSchemaAllOf) GetAccessMetaOk() (*WorkflowAccessMeta, bool)`

GetAccessMetaOk returns a tuple with the AccessMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMeta

`func (o *WorkflowSchemaAllOf) SetAccessMeta(v WorkflowAccessMeta)`

SetAccessMeta sets AccessMeta field to given value.

### HasAccessMeta

`func (o *WorkflowSchemaAllOf) HasAccessMeta() bool`

HasAccessMeta returns a boolean if a field has been set.

### GetVariableSchema

`func (o *WorkflowSchemaAllOf) GetVariableSchema() map[string]map[string]interface{}`

GetVariableSchema returns the VariableSchema field if non-nil, zero value otherwise.

### GetVariableSchemaOk

`func (o *WorkflowSchemaAllOf) GetVariableSchemaOk() (*map[string]map[string]interface{}, bool)`

GetVariableSchemaOk returns a tuple with the VariableSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableSchema

`func (o *WorkflowSchemaAllOf) SetVariableSchema(v map[string]map[string]interface{})`

SetVariableSchema sets VariableSchema field to given value.

### HasVariableSchema

`func (o *WorkflowSchemaAllOf) HasVariableSchema() bool`

HasVariableSchema returns a boolean if a field has been set.

### GetPropertySchema

`func (o *WorkflowSchemaAllOf) GetPropertySchema() map[string]map[string]interface{}`

GetPropertySchema returns the PropertySchema field if non-nil, zero value otherwise.

### GetPropertySchemaOk

`func (o *WorkflowSchemaAllOf) GetPropertySchemaOk() (*map[string]map[string]interface{}, bool)`

GetPropertySchemaOk returns a tuple with the PropertySchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertySchema

`func (o *WorkflowSchemaAllOf) SetPropertySchema(v map[string]map[string]interface{})`

SetPropertySchema sets PropertySchema field to given value.

### HasPropertySchema

`func (o *WorkflowSchemaAllOf) HasPropertySchema() bool`

HasPropertySchema returns a boolean if a field has been set.

### GetOutputSchema

`func (o *WorkflowSchemaAllOf) GetOutputSchema() map[string]map[string]interface{}`

GetOutputSchema returns the OutputSchema field if non-nil, zero value otherwise.

### GetOutputSchemaOk

`func (o *WorkflowSchemaAllOf) GetOutputSchemaOk() (*map[string]map[string]interface{}, bool)`

GetOutputSchemaOk returns a tuple with the OutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchema

`func (o *WorkflowSchemaAllOf) SetOutputSchema(v map[string]map[string]interface{})`

SetOutputSchema sets OutputSchema field to given value.

### HasOutputSchema

`func (o *WorkflowSchemaAllOf) HasOutputSchema() bool`

HasOutputSchema returns a boolean if a field has been set.

### GetViewConfig

`func (o *WorkflowSchemaAllOf) GetViewConfig() map[string]map[string]interface{}`

GetViewConfig returns the ViewConfig field if non-nil, zero value otherwise.

### GetViewConfigOk

`func (o *WorkflowSchemaAllOf) GetViewConfigOk() (*map[string]map[string]interface{}, bool)`

GetViewConfigOk returns a tuple with the ViewConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewConfig

`func (o *WorkflowSchemaAllOf) SetViewConfig(v map[string]map[string]interface{})`

SetViewConfig sets ViewConfig field to given value.

### HasViewConfig

`func (o *WorkflowSchemaAllOf) HasViewConfig() bool`

HasViewConfig returns a boolean if a field has been set.

### GetAttributes

`func (o *WorkflowSchemaAllOf) GetAttributes() map[string]map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WorkflowSchemaAllOf) GetAttributesOk() (*map[string]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WorkflowSchemaAllOf) SetAttributes(v map[string]map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WorkflowSchemaAllOf) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


