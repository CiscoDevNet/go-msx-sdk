# WorkflowTargetAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**SchemaId** | Pointer to **string** |  | [optional] 
**AdapterId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**BaseType** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 
**DefAccessMeta** | Pointer to [**WorkflowDefAccessMeta**](WorkflowDefAccessMeta.md) |  | [optional] 

## Methods

### NewWorkflowTargetAllOf

`func NewWorkflowTargetAllOf() *WorkflowTargetAllOf`

NewWorkflowTargetAllOf instantiates a new WorkflowTargetAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTargetAllOfWithDefaults

`func NewWorkflowTargetAllOfWithDefaults() *WorkflowTargetAllOf`

NewWorkflowTargetAllOfWithDefaults instantiates a new WorkflowTargetAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowTargetAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowTargetAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowTargetAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowTargetAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSchemaId

`func (o *WorkflowTargetAllOf) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *WorkflowTargetAllOf) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *WorkflowTargetAllOf) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *WorkflowTargetAllOf) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetAdapterId

`func (o *WorkflowTargetAllOf) GetAdapterId() string`

GetAdapterId returns the AdapterId field if non-nil, zero value otherwise.

### GetAdapterIdOk

`func (o *WorkflowTargetAllOf) GetAdapterIdOk() (*string, bool)`

GetAdapterIdOk returns a tuple with the AdapterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterId

`func (o *WorkflowTargetAllOf) SetAdapterId(v string)`

SetAdapterId sets AdapterId field to given value.

### HasAdapterId

`func (o *WorkflowTargetAllOf) HasAdapterId() bool`

HasAdapterId returns a boolean if a field has been set.

### GetName

`func (o *WorkflowTargetAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowTargetAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowTargetAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowTargetAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *WorkflowTargetAllOf) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WorkflowTargetAllOf) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WorkflowTargetAllOf) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WorkflowTargetAllOf) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *WorkflowTargetAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowTargetAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowTargetAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowTargetAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBaseType

`func (o *WorkflowTargetAllOf) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *WorkflowTargetAllOf) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *WorkflowTargetAllOf) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *WorkflowTargetAllOf) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetProperties

`func (o *WorkflowTargetAllOf) GetProperties() map[string]map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowTargetAllOf) GetPropertiesOk() (*map[string]map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowTargetAllOf) SetProperties(v map[string]map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowTargetAllOf) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetValid

`func (o *WorkflowTargetAllOf) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *WorkflowTargetAllOf) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *WorkflowTargetAllOf) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *WorkflowTargetAllOf) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetDefAccessMeta

`func (o *WorkflowTargetAllOf) GetDefAccessMeta() WorkflowDefAccessMeta`

GetDefAccessMeta returns the DefAccessMeta field if non-nil, zero value otherwise.

### GetDefAccessMetaOk

`func (o *WorkflowTargetAllOf) GetDefAccessMetaOk() (*WorkflowDefAccessMeta, bool)`

GetDefAccessMetaOk returns a tuple with the DefAccessMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefAccessMeta

`func (o *WorkflowTargetAllOf) SetDefAccessMeta(v WorkflowDefAccessMeta)`

SetDefAccessMeta sets DefAccessMeta field to given value.

### HasDefAccessMeta

`func (o *WorkflowTargetAllOf) HasDefAccessMeta() bool`

HasDefAccessMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


