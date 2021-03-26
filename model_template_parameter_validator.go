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

// TemplateParameterValidator It's metadata about a parameter for use in the UI.  A name to put on the field, a type so it can be validated (e.g. an IP address or an integer value, or list of allowed values for an enumeration/list).
type TemplateParameterValidator struct {
	Name *string `json:"name,omitempty"`
	HintText *string `json:"hintText,omitempty"`
	Label *string `json:"label,omitempty"`
	Type *string `json:"type,omitempty"`
	DisplayType *string `json:"displayType,omitempty"`
	AllowedValues *[]string `json:"allowedValues,omitempty"`
	ToolTipText *string `json:"toolTipText,omitempty"`
}

// NewTemplateParameterValidator instantiates a new TemplateParameterValidator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateParameterValidator() *TemplateParameterValidator {
	this := TemplateParameterValidator{}
	return &this
}

// NewTemplateParameterValidatorWithDefaults instantiates a new TemplateParameterValidator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateParameterValidatorWithDefaults() *TemplateParameterValidator {
	this := TemplateParameterValidator{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TemplateParameterValidator) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateParameterValidator) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TemplateParameterValidator) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TemplateParameterValidator) SetName(v string) {
	o.Name = &v
}

// GetHintText returns the HintText field value if set, zero value otherwise.
func (o *TemplateParameterValidator) GetHintText() string {
	if o == nil || o.HintText == nil {
		var ret string
		return ret
	}
	return *o.HintText
}

// GetHintTextOk returns a tuple with the HintText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateParameterValidator) GetHintTextOk() (*string, bool) {
	if o == nil || o.HintText == nil {
		return nil, false
	}
	return o.HintText, true
}

// HasHintText returns a boolean if a field has been set.
func (o *TemplateParameterValidator) HasHintText() bool {
	if o != nil && o.HintText != nil {
		return true
	}

	return false
}

// SetHintText gets a reference to the given string and assigns it to the HintText field.
func (o *TemplateParameterValidator) SetHintText(v string) {
	o.HintText = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *TemplateParameterValidator) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateParameterValidator) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *TemplateParameterValidator) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *TemplateParameterValidator) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TemplateParameterValidator) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateParameterValidator) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TemplateParameterValidator) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TemplateParameterValidator) SetType(v string) {
	o.Type = &v
}

// GetDisplayType returns the DisplayType field value if set, zero value otherwise.
func (o *TemplateParameterValidator) GetDisplayType() string {
	if o == nil || o.DisplayType == nil {
		var ret string
		return ret
	}
	return *o.DisplayType
}

// GetDisplayTypeOk returns a tuple with the DisplayType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateParameterValidator) GetDisplayTypeOk() (*string, bool) {
	if o == nil || o.DisplayType == nil {
		return nil, false
	}
	return o.DisplayType, true
}

// HasDisplayType returns a boolean if a field has been set.
func (o *TemplateParameterValidator) HasDisplayType() bool {
	if o != nil && o.DisplayType != nil {
		return true
	}

	return false
}

// SetDisplayType gets a reference to the given string and assigns it to the DisplayType field.
func (o *TemplateParameterValidator) SetDisplayType(v string) {
	o.DisplayType = &v
}

// GetAllowedValues returns the AllowedValues field value if set, zero value otherwise.
func (o *TemplateParameterValidator) GetAllowedValues() []string {
	if o == nil || o.AllowedValues == nil {
		var ret []string
		return ret
	}
	return *o.AllowedValues
}

// GetAllowedValuesOk returns a tuple with the AllowedValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateParameterValidator) GetAllowedValuesOk() (*[]string, bool) {
	if o == nil || o.AllowedValues == nil {
		return nil, false
	}
	return o.AllowedValues, true
}

// HasAllowedValues returns a boolean if a field has been set.
func (o *TemplateParameterValidator) HasAllowedValues() bool {
	if o != nil && o.AllowedValues != nil {
		return true
	}

	return false
}

// SetAllowedValues gets a reference to the given []string and assigns it to the AllowedValues field.
func (o *TemplateParameterValidator) SetAllowedValues(v []string) {
	o.AllowedValues = &v
}

// GetToolTipText returns the ToolTipText field value if set, zero value otherwise.
func (o *TemplateParameterValidator) GetToolTipText() string {
	if o == nil || o.ToolTipText == nil {
		var ret string
		return ret
	}
	return *o.ToolTipText
}

// GetToolTipTextOk returns a tuple with the ToolTipText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateParameterValidator) GetToolTipTextOk() (*string, bool) {
	if o == nil || o.ToolTipText == nil {
		return nil, false
	}
	return o.ToolTipText, true
}

// HasToolTipText returns a boolean if a field has been set.
func (o *TemplateParameterValidator) HasToolTipText() bool {
	if o != nil && o.ToolTipText != nil {
		return true
	}

	return false
}

// SetToolTipText gets a reference to the given string and assigns it to the ToolTipText field.
func (o *TemplateParameterValidator) SetToolTipText(v string) {
	o.ToolTipText = &v
}

func (o TemplateParameterValidator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.HintText != nil {
		toSerialize["hintText"] = o.HintText
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.DisplayType != nil {
		toSerialize["displayType"] = o.DisplayType
	}
	if o.AllowedValues != nil {
		toSerialize["allowedValues"] = o.AllowedValues
	}
	if o.ToolTipText != nil {
		toSerialize["toolTipText"] = o.ToolTipText
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateParameterValidator struct {
	value *TemplateParameterValidator
	isSet bool
}

func (v NullableTemplateParameterValidator) Get() *TemplateParameterValidator {
	return v.value
}

func (v *NullableTemplateParameterValidator) Set(val *TemplateParameterValidator) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateParameterValidator) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateParameterValidator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateParameterValidator(val *TemplateParameterValidator) *NullableTemplateParameterValidator {
	return &NullableTemplateParameterValidator{value: val, isSet: true}
}

func (v NullableTemplateParameterValidator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateParameterValidator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


