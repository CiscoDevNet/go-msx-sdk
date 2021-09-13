# WorkflowTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**SchemaId** | Pointer to **string** |  | [optional] 
**AdapterId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**BaseType** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 
**DefAccessMeta** | Pointer to [**WorkflowDefAccessMeta**](WorkflowDefAccessMeta.md) |  | [optional] 
**CreatedOn** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**UpdatedOn** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**Owner** | Pointer to **NullableString** |  | [optional] 
**UniqueName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWorkflowTarget

`func NewWorkflowTarget() *WorkflowTarget`

NewWorkflowTarget instantiates a new WorkflowTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTargetWithDefaults

`func NewWorkflowTargetWithDefaults() *WorkflowTarget`

NewWorkflowTargetWithDefaults instantiates a new WorkflowTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowTarget) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowTarget) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowTarget) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowTarget) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSchemaId

`func (o *WorkflowTarget) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *WorkflowTarget) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *WorkflowTarget) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *WorkflowTarget) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetAdapterId

`func (o *WorkflowTarget) GetAdapterId() string`

GetAdapterId returns the AdapterId field if non-nil, zero value otherwise.

### GetAdapterIdOk

`func (o *WorkflowTarget) GetAdapterIdOk() (*string, bool)`

GetAdapterIdOk returns a tuple with the AdapterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterId

`func (o *WorkflowTarget) SetAdapterId(v string)`

SetAdapterId sets AdapterId field to given value.

### HasAdapterId

`func (o *WorkflowTarget) HasAdapterId() bool`

HasAdapterId returns a boolean if a field has been set.

### GetName

`func (o *WorkflowTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowTarget) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowTarget) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *WorkflowTarget) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WorkflowTarget) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WorkflowTarget) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WorkflowTarget) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *WorkflowTarget) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *WorkflowTarget) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetType

`func (o *WorkflowTarget) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowTarget) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowTarget) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowTarget) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBaseType

`func (o *WorkflowTarget) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *WorkflowTarget) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *WorkflowTarget) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *WorkflowTarget) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetProperties

`func (o *WorkflowTarget) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowTarget) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowTarget) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowTarget) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *WorkflowTarget) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *WorkflowTarget) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetValid

`func (o *WorkflowTarget) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *WorkflowTarget) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *WorkflowTarget) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *WorkflowTarget) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetDefAccessMeta

`func (o *WorkflowTarget) GetDefAccessMeta() WorkflowDefAccessMeta`

GetDefAccessMeta returns the DefAccessMeta field if non-nil, zero value otherwise.

### GetDefAccessMetaOk

`func (o *WorkflowTarget) GetDefAccessMetaOk() (*WorkflowDefAccessMeta, bool)`

GetDefAccessMetaOk returns a tuple with the DefAccessMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefAccessMeta

`func (o *WorkflowTarget) SetDefAccessMeta(v WorkflowDefAccessMeta)`

SetDefAccessMeta sets DefAccessMeta field to given value.

### HasDefAccessMeta

`func (o *WorkflowTarget) HasDefAccessMeta() bool`

HasDefAccessMeta returns a boolean if a field has been set.

### GetCreatedOn

`func (o *WorkflowTarget) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *WorkflowTarget) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *WorkflowTarget) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *WorkflowTarget) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetCreatedBy

`func (o *WorkflowTarget) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WorkflowTarget) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WorkflowTarget) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *WorkflowTarget) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *WorkflowTarget) GetUpdatedOn() string`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *WorkflowTarget) GetUpdatedOnOk() (*string, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *WorkflowTarget) SetUpdatedOn(v string)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *WorkflowTarget) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *WorkflowTarget) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *WorkflowTarget) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *WorkflowTarget) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *WorkflowTarget) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *WorkflowTarget) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *WorkflowTarget) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetOwner

`func (o *WorkflowTarget) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *WorkflowTarget) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *WorkflowTarget) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *WorkflowTarget) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *WorkflowTarget) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *WorkflowTarget) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetUniqueName

`func (o *WorkflowTarget) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *WorkflowTarget) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *WorkflowTarget) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *WorkflowTarget) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### SetUniqueNameNil

`func (o *WorkflowTarget) SetUniqueNameNil(b bool)`

 SetUniqueNameNil sets the value for UniqueName to be an explicit nil

### UnsetUniqueName
`func (o *WorkflowTarget) UnsetUniqueName()`

UnsetUniqueName ensures that no value is present for UniqueName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


