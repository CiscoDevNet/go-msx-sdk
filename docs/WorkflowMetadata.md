# WorkflowMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitInfo** | Pointer to [**WorkflowMetadataGitInfo**](WorkflowMetadataGitInfo.md) |  | [optional] 

## Methods

### NewWorkflowMetadata

`func NewWorkflowMetadata() *WorkflowMetadata`

NewWorkflowMetadata instantiates a new WorkflowMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowMetadataWithDefaults

`func NewWorkflowMetadataWithDefaults() *WorkflowMetadata`

NewWorkflowMetadataWithDefaults instantiates a new WorkflowMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitInfo

`func (o *WorkflowMetadata) GetGitInfo() WorkflowMetadataGitInfo`

GetGitInfo returns the GitInfo field if non-nil, zero value otherwise.

### GetGitInfoOk

`func (o *WorkflowMetadata) GetGitInfoOk() (*WorkflowMetadataGitInfo, bool)`

GetGitInfoOk returns a tuple with the GitInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitInfo

`func (o *WorkflowMetadata) SetGitInfo(v WorkflowMetadataGitInfo)`

SetGitInfo sets GitInfo field to given value.

### HasGitInfo

`func (o *WorkflowMetadata) HasGitInfo() bool`

HasGitInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


