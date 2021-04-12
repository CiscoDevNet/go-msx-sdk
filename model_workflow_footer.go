/*
 * KAKAPO - MSX SDK
 *
 * MSX Platform SDK
 *
 * API version: 1.0.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msxsdk

import (
	"encoding/json"
)

// WorkflowFooter struct for WorkflowFooter
type WorkflowFooter struct {
	CreatedOn *string `json:"created_on,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedOn *string `json:"updated_on,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	Owner *string `json:"owner,omitempty"`
	UniqueName *string `json:"unique_name,omitempty"`
}

// NewWorkflowFooter instantiates a new WorkflowFooter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowFooter() *WorkflowFooter {
	this := WorkflowFooter{}
	return &this
}

// NewWorkflowFooterWithDefaults instantiates a new WorkflowFooter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowFooterWithDefaults() *WorkflowFooter {
	this := WorkflowFooter{}
	return &this
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *WorkflowFooter) GetCreatedOn() string {
	if o == nil || o.CreatedOn == nil {
		var ret string
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowFooter) GetCreatedOnOk() (*string, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *WorkflowFooter) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given string and assigns it to the CreatedOn field.
func (o *WorkflowFooter) SetCreatedOn(v string) {
	o.CreatedOn = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *WorkflowFooter) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowFooter) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *WorkflowFooter) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *WorkflowFooter) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetUpdatedOn returns the UpdatedOn field value if set, zero value otherwise.
func (o *WorkflowFooter) GetUpdatedOn() string {
	if o == nil || o.UpdatedOn == nil {
		var ret string
		return ret
	}
	return *o.UpdatedOn
}

// GetUpdatedOnOk returns a tuple with the UpdatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowFooter) GetUpdatedOnOk() (*string, bool) {
	if o == nil || o.UpdatedOn == nil {
		return nil, false
	}
	return o.UpdatedOn, true
}

// HasUpdatedOn returns a boolean if a field has been set.
func (o *WorkflowFooter) HasUpdatedOn() bool {
	if o != nil && o.UpdatedOn != nil {
		return true
	}

	return false
}

// SetUpdatedOn gets a reference to the given string and assigns it to the UpdatedOn field.
func (o *WorkflowFooter) SetUpdatedOn(v string) {
	o.UpdatedOn = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *WorkflowFooter) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowFooter) GetUpdatedByOk() (*string, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *WorkflowFooter) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *WorkflowFooter) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *WorkflowFooter) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowFooter) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *WorkflowFooter) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *WorkflowFooter) SetOwner(v string) {
	o.Owner = &v
}

// GetUniqueName returns the UniqueName field value if set, zero value otherwise.
func (o *WorkflowFooter) GetUniqueName() string {
	if o == nil || o.UniqueName == nil {
		var ret string
		return ret
	}
	return *o.UniqueName
}

// GetUniqueNameOk returns a tuple with the UniqueName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowFooter) GetUniqueNameOk() (*string, bool) {
	if o == nil || o.UniqueName == nil {
		return nil, false
	}
	return o.UniqueName, true
}

// HasUniqueName returns a boolean if a field has been set.
func (o *WorkflowFooter) HasUniqueName() bool {
	if o != nil && o.UniqueName != nil {
		return true
	}

	return false
}

// SetUniqueName gets a reference to the given string and assigns it to the UniqueName field.
func (o *WorkflowFooter) SetUniqueName(v string) {
	o.UniqueName = &v
}

func (o WorkflowFooter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedOn != nil {
		toSerialize["created_on"] = o.CreatedOn
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.UpdatedOn != nil {
		toSerialize["updated_on"] = o.UpdatedOn
	}
	if o.UpdatedBy != nil {
		toSerialize["updated_by"] = o.UpdatedBy
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.UniqueName != nil {
		toSerialize["unique_name"] = o.UniqueName
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowFooter struct {
	value *WorkflowFooter
	isSet bool
}

func (v NullableWorkflowFooter) Get() *WorkflowFooter {
	return v.value
}

func (v *NullableWorkflowFooter) Set(val *WorkflowFooter) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowFooter) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowFooter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowFooter(val *WorkflowFooter) *NullableWorkflowFooter {
	return &NullableWorkflowFooter{value: val, isSet: true}
}

func (v NullableWorkflowFooter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowFooter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


