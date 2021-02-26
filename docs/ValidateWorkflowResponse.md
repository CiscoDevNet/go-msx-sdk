# ValidateWorkflowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalActions** | Pointer to **int32** |  | [optional] 
**TotalValid** | Pointer to **int32** |  | [optional] 
**WorkflowValid** | Pointer to **bool** |  | [optional] 
**InvalidActionIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewValidateWorkflowResponse

`func NewValidateWorkflowResponse() *ValidateWorkflowResponse`

NewValidateWorkflowResponse instantiates a new ValidateWorkflowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateWorkflowResponseWithDefaults

`func NewValidateWorkflowResponseWithDefaults() *ValidateWorkflowResponse`

NewValidateWorkflowResponseWithDefaults instantiates a new ValidateWorkflowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalActions

`func (o *ValidateWorkflowResponse) GetTotalActions() int32`

GetTotalActions returns the TotalActions field if non-nil, zero value otherwise.

### GetTotalActionsOk

`func (o *ValidateWorkflowResponse) GetTotalActionsOk() (*int32, bool)`

GetTotalActionsOk returns a tuple with the TotalActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalActions

`func (o *ValidateWorkflowResponse) SetTotalActions(v int32)`

SetTotalActions sets TotalActions field to given value.

### HasTotalActions

`func (o *ValidateWorkflowResponse) HasTotalActions() bool`

HasTotalActions returns a boolean if a field has been set.

### GetTotalValid

`func (o *ValidateWorkflowResponse) GetTotalValid() int32`

GetTotalValid returns the TotalValid field if non-nil, zero value otherwise.

### GetTotalValidOk

`func (o *ValidateWorkflowResponse) GetTotalValidOk() (*int32, bool)`

GetTotalValidOk returns a tuple with the TotalValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalValid

`func (o *ValidateWorkflowResponse) SetTotalValid(v int32)`

SetTotalValid sets TotalValid field to given value.

### HasTotalValid

`func (o *ValidateWorkflowResponse) HasTotalValid() bool`

HasTotalValid returns a boolean if a field has been set.

### GetWorkflowValid

`func (o *ValidateWorkflowResponse) GetWorkflowValid() bool`

GetWorkflowValid returns the WorkflowValid field if non-nil, zero value otherwise.

### GetWorkflowValidOk

`func (o *ValidateWorkflowResponse) GetWorkflowValidOk() (*bool, bool)`

GetWorkflowValidOk returns a tuple with the WorkflowValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowValid

`func (o *ValidateWorkflowResponse) SetWorkflowValid(v bool)`

SetWorkflowValid sets WorkflowValid field to given value.

### HasWorkflowValid

`func (o *ValidateWorkflowResponse) HasWorkflowValid() bool`

HasWorkflowValid returns a boolean if a field has been set.

### GetInvalidActionIds

`func (o *ValidateWorkflowResponse) GetInvalidActionIds() []string`

GetInvalidActionIds returns the InvalidActionIds field if non-nil, zero value otherwise.

### GetInvalidActionIdsOk

`func (o *ValidateWorkflowResponse) GetInvalidActionIdsOk() (*[]string, bool)`

GetInvalidActionIdsOk returns a tuple with the InvalidActionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidActionIds

`func (o *ValidateWorkflowResponse) SetInvalidActionIds(v []string)`

SetInvalidActionIds sets InvalidActionIds field to given value.

### HasInvalidActionIds

`func (o *ValidateWorkflowResponse) HasInvalidActionIds() bool`

HasInvalidActionIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


