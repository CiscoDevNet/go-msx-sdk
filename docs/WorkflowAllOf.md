# WorkflowAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**SchemaId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**BaseType** | Pointer to **string** |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 
**WorkflowValid** | Pointer to **bool** |  | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 
**Metadata** | Pointer to [**WorkflowMetadata**](WorkflowMetadata.md) |  | [optional] 
**Status** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Permissions** | Pointer to **[]string** |  | [optional] 
**Variables** | Pointer to [**[]WorkflowVariable**](WorkflowVariable.md) |  | [optional] 
**Actions** | Pointer to [**[]WorkflowAction**](WorkflowAction.md) |  | [optional] 

## Methods

### NewWorkflowAllOf

`func NewWorkflowAllOf() *WorkflowAllOf`

NewWorkflowAllOf instantiates a new WorkflowAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowAllOfWithDefaults

`func NewWorkflowAllOfWithDefaults() *WorkflowAllOf`

NewWorkflowAllOfWithDefaults instantiates a new WorkflowAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSchemaId

`func (o *WorkflowAllOf) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *WorkflowAllOf) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *WorkflowAllOf) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *WorkflowAllOf) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetName

`func (o *WorkflowAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *WorkflowAllOf) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WorkflowAllOf) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WorkflowAllOf) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WorkflowAllOf) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *WorkflowAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBaseType

`func (o *WorkflowAllOf) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *WorkflowAllOf) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *WorkflowAllOf) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *WorkflowAllOf) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetObjectType

`func (o *WorkflowAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *WorkflowAllOf) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetProperties

`func (o *WorkflowAllOf) GetProperties() map[string]map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowAllOf) GetPropertiesOk() (*map[string]map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowAllOf) SetProperties(v map[string]map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowAllOf) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetValid

`func (o *WorkflowAllOf) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *WorkflowAllOf) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *WorkflowAllOf) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *WorkflowAllOf) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetWorkflowValid

`func (o *WorkflowAllOf) GetWorkflowValid() bool`

GetWorkflowValid returns the WorkflowValid field if non-nil, zero value otherwise.

### GetWorkflowValidOk

`func (o *WorkflowAllOf) GetWorkflowValidOk() (*bool, bool)`

GetWorkflowValidOk returns a tuple with the WorkflowValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowValid

`func (o *WorkflowAllOf) SetWorkflowValid(v bool)`

SetWorkflowValid sets WorkflowValid field to given value.

### HasWorkflowValid

`func (o *WorkflowAllOf) HasWorkflowValid() bool`

HasWorkflowValid returns a boolean if a field has been set.

### GetCategories

`func (o *WorkflowAllOf) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *WorkflowAllOf) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *WorkflowAllOf) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *WorkflowAllOf) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetMetadata

`func (o *WorkflowAllOf) GetMetadata() WorkflowMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WorkflowAllOf) GetMetadataOk() (*WorkflowMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WorkflowAllOf) SetMetadata(v WorkflowMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *WorkflowAllOf) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowAllOf) GetStatus() map[string]map[string]interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowAllOf) GetStatusOk() (*map[string]map[string]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowAllOf) SetStatus(v map[string]map[string]interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPermissions

`func (o *WorkflowAllOf) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *WorkflowAllOf) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *WorkflowAllOf) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *WorkflowAllOf) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetVariables

`func (o *WorkflowAllOf) GetVariables() []WorkflowVariable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *WorkflowAllOf) GetVariablesOk() (*[]WorkflowVariable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *WorkflowAllOf) SetVariables(v []WorkflowVariable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *WorkflowAllOf) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetActions

`func (o *WorkflowAllOf) GetActions() []WorkflowAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *WorkflowAllOf) GetActionsOk() (*[]WorkflowAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *WorkflowAllOf) SetActions(v []WorkflowAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *WorkflowAllOf) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


