/*
 * KAKAPO - MSX SDK
 *
 * MSX Platform SDK
 *
 * API version: 1.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msxsdk

import (
	"encoding/json"
)

// ValidateWorkflowResponse struct for ValidateWorkflowResponse
type ValidateWorkflowResponse struct {
	TotalActions *int32 `json:"total_actions,omitempty"`
	TotalValid *int32 `json:"total_valid,omitempty"`
	WorkflowValid *bool `json:"workflow_valid,omitempty"`
	InvalidActionIds *[]string `json:"invalid_action_ids,omitempty"`
}

// NewValidateWorkflowResponse instantiates a new ValidateWorkflowResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateWorkflowResponse() *ValidateWorkflowResponse {
	this := ValidateWorkflowResponse{}
	return &this
}

// NewValidateWorkflowResponseWithDefaults instantiates a new ValidateWorkflowResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateWorkflowResponseWithDefaults() *ValidateWorkflowResponse {
	this := ValidateWorkflowResponse{}
	return &this
}

// GetTotalActions returns the TotalActions field value if set, zero value otherwise.
func (o *ValidateWorkflowResponse) GetTotalActions() int32 {
	if o == nil || o.TotalActions == nil {
		var ret int32
		return ret
	}
	return *o.TotalActions
}

// GetTotalActionsOk returns a tuple with the TotalActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateWorkflowResponse) GetTotalActionsOk() (*int32, bool) {
	if o == nil || o.TotalActions == nil {
		return nil, false
	}
	return o.TotalActions, true
}

// HasTotalActions returns a boolean if a field has been set.
func (o *ValidateWorkflowResponse) HasTotalActions() bool {
	if o != nil && o.TotalActions != nil {
		return true
	}

	return false
}

// SetTotalActions gets a reference to the given int32 and assigns it to the TotalActions field.
func (o *ValidateWorkflowResponse) SetTotalActions(v int32) {
	o.TotalActions = &v
}

// GetTotalValid returns the TotalValid field value if set, zero value otherwise.
func (o *ValidateWorkflowResponse) GetTotalValid() int32 {
	if o == nil || o.TotalValid == nil {
		var ret int32
		return ret
	}
	return *o.TotalValid
}

// GetTotalValidOk returns a tuple with the TotalValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateWorkflowResponse) GetTotalValidOk() (*int32, bool) {
	if o == nil || o.TotalValid == nil {
		return nil, false
	}
	return o.TotalValid, true
}

// HasTotalValid returns a boolean if a field has been set.
func (o *ValidateWorkflowResponse) HasTotalValid() bool {
	if o != nil && o.TotalValid != nil {
		return true
	}

	return false
}

// SetTotalValid gets a reference to the given int32 and assigns it to the TotalValid field.
func (o *ValidateWorkflowResponse) SetTotalValid(v int32) {
	o.TotalValid = &v
}

// GetWorkflowValid returns the WorkflowValid field value if set, zero value otherwise.
func (o *ValidateWorkflowResponse) GetWorkflowValid() bool {
	if o == nil || o.WorkflowValid == nil {
		var ret bool
		return ret
	}
	return *o.WorkflowValid
}

// GetWorkflowValidOk returns a tuple with the WorkflowValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateWorkflowResponse) GetWorkflowValidOk() (*bool, bool) {
	if o == nil || o.WorkflowValid == nil {
		return nil, false
	}
	return o.WorkflowValid, true
}

// HasWorkflowValid returns a boolean if a field has been set.
func (o *ValidateWorkflowResponse) HasWorkflowValid() bool {
	if o != nil && o.WorkflowValid != nil {
		return true
	}

	return false
}

// SetWorkflowValid gets a reference to the given bool and assigns it to the WorkflowValid field.
func (o *ValidateWorkflowResponse) SetWorkflowValid(v bool) {
	o.WorkflowValid = &v
}

// GetInvalidActionIds returns the InvalidActionIds field value if set, zero value otherwise.
func (o *ValidateWorkflowResponse) GetInvalidActionIds() []string {
	if o == nil || o.InvalidActionIds == nil {
		var ret []string
		return ret
	}
	return *o.InvalidActionIds
}

// GetInvalidActionIdsOk returns a tuple with the InvalidActionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateWorkflowResponse) GetInvalidActionIdsOk() (*[]string, bool) {
	if o == nil || o.InvalidActionIds == nil {
		return nil, false
	}
	return o.InvalidActionIds, true
}

// HasInvalidActionIds returns a boolean if a field has been set.
func (o *ValidateWorkflowResponse) HasInvalidActionIds() bool {
	if o != nil && o.InvalidActionIds != nil {
		return true
	}

	return false
}

// SetInvalidActionIds gets a reference to the given []string and assigns it to the InvalidActionIds field.
func (o *ValidateWorkflowResponse) SetInvalidActionIds(v []string) {
	o.InvalidActionIds = &v
}

func (o ValidateWorkflowResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TotalActions != nil {
		toSerialize["total_actions"] = o.TotalActions
	}
	if o.TotalValid != nil {
		toSerialize["total_valid"] = o.TotalValid
	}
	if o.WorkflowValid != nil {
		toSerialize["workflow_valid"] = o.WorkflowValid
	}
	if o.InvalidActionIds != nil {
		toSerialize["invalid_action_ids"] = o.InvalidActionIds
	}
	return json.Marshal(toSerialize)
}

type NullableValidateWorkflowResponse struct {
	value *ValidateWorkflowResponse
	isSet bool
}

func (v NullableValidateWorkflowResponse) Get() *ValidateWorkflowResponse {
	return v.value
}

func (v *NullableValidateWorkflowResponse) Set(val *ValidateWorkflowResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateWorkflowResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateWorkflowResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateWorkflowResponse(val *ValidateWorkflowResponse) *NullableValidateWorkflowResponse {
	return &NullableValidateWorkflowResponse{value: val, isSet: true}
}

func (v NullableValidateWorkflowResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateWorkflowResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


