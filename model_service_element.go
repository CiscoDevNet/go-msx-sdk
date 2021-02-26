/*
 * KAKAPO - MSX SDK
 *
 * MSX Platform SDK
 *
 * API version: 3.10.0
 * Contact: somecontact@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msxsdk

import (
	"encoding/json"
)

// ServiceElement struct for ServiceElement
type ServiceElement struct {
	Name *string `json:"name,omitempty"`
	Label *string `json:"label,omitempty"`
	Header *string `json:"header,omitempty"`
	Description *string `json:"description,omitempty"`
	Hint *string `json:"hint,omitempty"`
	InputType *string `json:"inputType,omitempty"`
	Type *string `json:"type,omitempty"`
	Component *string `json:"component,omitempty"`
	MaxLimit *string `json:"maxLimit,omitempty"`
	MinLimit *string `json:"minLimit,omitempty"`
	Value *string `json:"value,omitempty"`
	ValueList *[]map[string]interface{} `json:"valueList,omitempty"`
	AllowedOptionValues *[]string `json:"allowedOptionValues,omitempty"`
	AllowedValues *[]map[string]map[string]interface{} `json:"allowedValues,omitempty"`
	Mandatory *bool `json:"mandatory,omitempty"`
	Section *string `json:"section,omitempty"`
	Billable *bool `json:"billable,omitempty"`
	Hidden *bool `json:"hidden,omitempty"`
	ParentName *string `json:"parentName,omitempty"`
	Supported *bool `json:"supported,omitempty"`
	DynamicData *bool `json:"dynamicData,omitempty"`
	MinValue *int32 `json:"minValue,omitempty"`
	MaxValue *int32 `json:"maxValue,omitempty"`
	StepSize *int32 `json:"stepSize,omitempty"`
	PricingAtttributes *ServiceElementPrice `json:"pricingAtttributes,omitempty"`
	ChildElements *[]ServiceElement `json:"childElements,omitempty"`
}

// NewServiceElement instantiates a new ServiceElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceElement() *ServiceElement {
	this := ServiceElement{}
	return &this
}

// NewServiceElementWithDefaults instantiates a new ServiceElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceElementWithDefaults() *ServiceElement {
	this := ServiceElement{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceElement) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceElement) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceElement) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceElement) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ServiceElement) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceElement) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ServiceElement) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ServiceElement) SetLabel(v string) {
	o.Label = &v
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *ServiceElement) GetHeader() string {
	if o == nil || o.Header == nil {
		var ret string
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceElement) GetHeaderOk() (*string, bool) {
	if o == nil || o.Header == nil {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *ServiceElement) HasHeader() bool {
	if o != nil && o.Header != nil {
		return true
	}

	return false
}

// SetHeader gets a reference to the given string and assigns it to the Header field.
func (o *ServiceElement) SetHeader(v string) {
	o.Header = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceElement) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceElement) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceElement) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceElement) SetDescription(v string) {
	o.Description = &v
}

// GetHint returns the Hint field value if set, zero value otherwise.
func (o *ServiceElement) GetHint() string {
	if o == nil || o.Hint == nil {
		var ret string
		return ret
	}
	return *o.Hint
}

// GetHintOk returns a tuple with the Hint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceElement) GetHintOk() (*string, bool) {
	if o == nil || o.Hint == nil {
		return nil, false
	}
	return o.Hint, true
}

// HasHint returns a boolean if a field has been set.
func (o *ServiceElement) HasHint() bool {
	if o != nil && o.Hint != nil {
		return true
	}

	return false
}

// SetHint gets a reference to the given string and assigns it to the Hint field.
func (o *ServiceElement) SetHint(v string) {
	o.Hint = &v
}

// GetInputType returns the InputType field value if set, zero value otherwise.
func (o *ServiceElement) GetInputType() string {
	if o == nil || o.InputType == nil {
		var ret string
		return ret
	}
	return *o.InputType
}

// GetInputTypeOk returns a tuple with the InputType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceElement) GetInputTypeOk() (*string, bool) {
	if o == nil || o.InputType == nil {
		return nil, false
	}
	return o.InputType, true
}

// HasInputType returns a boolean if a field has been set.
func (o *ServiceElement) HasInputType() bool {
	if o != nil && o.InputType != nil {
		return true
	}

	return false
}

// SetInputType gets a reference to the given string and assigns it to the InputType field.
func (o *ServiceElement) SetInputType(v string) {
	o.InputType = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ServiceElement) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceElement) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ServiceElement) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ServiceElement) SetType(v string) {
	o.Type = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *ServiceElement) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceElement) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *ServiceElement) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *ServiceElement) SetComponent(v string) {
	o.Component = &v
}

// GetMaxLimit returns the MaxLimit field value if set, zero value otherwise.
func (o *ServiceElement) GetMaxLimit() string {
	if o == nil || o.MaxLimit == nil {
		var ret string
		return ret
	}
	return *o.MaxLimit
}

// GetMaxLimitOk returns a tuple with the MaxLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceElement) GetMaxLimitOk() (*string, bool) {
	if o == nil || o.MaxLimit == nil {
		return nil, false
	}
	return o.MaxLimit, true
}

// HasMaxLimit returns a boolean if a field has been set.
func (o *ServiceElement) HasMaxLimit() bool {
	if o != nil && o.MaxLimit != nil {
		return true
	}

	return false
}

// SetMaxLimit gets a reference to the given string and assigns it to the MaxLimit field.
func (o *ServiceElement) SetMaxLimit(v string) {
	o.MaxLimit = &v
}

// GetMinLimit returns the MinLimit field value if set, zero value otherwise.
func (o *ServiceElement) GetMinLimit() string {
	if o == nil || o.MinLimit == nil {
		var ret string
		return ret
	}
	return *o.MinLimit
}

// GetMinLimitOk returns a tuple with the MinLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceElement) GetMinLimitOk() (*string, bool) {
	if o == nil || o.MinLimit == nil {
		return nil, false
	}
	return o.MinLimit, true
}

// HasMinLimit returns a boolean if a field has been set.
func (o *ServiceElement) HasMinLimit() bool {
	if o != nil && o.MinLimit != nil {
		return true
	}

	return false
}

// SetMinLimit gets a reference to the given string and assigns it to the MinLimit field.
func (o *ServiceElement) SetMinLimit(v string) {
	o.MinLimit = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ServiceElement) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceElement) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ServiceElement) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ServiceElement) SetValue(v string) {
	o.Value = &v
}

// GetValueList returns the ValueList field value if set, zero value otherwise.
func (o *ServiceElement) GetValueList() []map[string]interface{} {
	if o == nil || o.ValueList == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.ValueList
}

// GetValueListOk returns a tuple with the ValueList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceElement) GetValueListOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.ValueList == nil {
		return nil, false
	}
	return o.ValueList, true
}

// HasValueList returns a boolean if a field has been set.
func (o *ServiceElement) HasValueList() bool {
	if o != nil && o.ValueList != nil {
		return true
	}

	return false
}

// SetValueList gets a reference to the given []map[string]interface{} and assigns it to the ValueList field.
func (o *ServiceElement) SetValueList(v []map[string]interface{}) {
	o.ValueList = &v
}

// GetAllowedOptionValues returns the AllowedOptionValues field value if set, zero value otherwise.
func (o *ServiceElement) GetAllowedOptionValues() []string {
	if o == nil || o.AllowedOptionValues == nil {
		var ret []string
		return ret
	}
	return *o.AllowedOptionValues
}

// GetAllowedOptionValuesOk returns a tuple with the AllowedOptionValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceElement) GetAllowedOptionValuesOk() (*[]string, bool) {
	if o == nil || o.AllowedOptionValues == nil {
		return nil, false
	}
	return o.AllowedOptionValues, true
}

// HasAllowedOptionValues returns a boolean if a field has been set.
func (o *ServiceElement) HasAllowedOptionValues() bool {
	if o != nil && o.AllowedOptionValues != nil {
		return true
	}

	return false
}

// SetAllowedOptionValues gets a reference to the given []string and assigns it to the AllowedOptionValues field.
func (o *ServiceElement) SetAllowedOptionValues(v []string) {
	o.AllowedOptionValues = &v
}

// GetAllowedValues returns the AllowedValues field value if set, zero value otherwise.
func (o *ServiceElement) GetAllowedValues() []map[string]map[string]interface{} {
	if o == nil || o.AllowedValues == nil {
		var ret []map[string]map[string]interface{}
		return ret
	}
	return *o.AllowedValues
}

// GetAllowedValuesOk returns a tuple with the AllowedValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceElement) GetAllowedValuesOk() (*[]map[string]map[string]interface{}, bool) {
	if o == nil || o.AllowedValues == nil {
		return nil, false
	}
	return o.AllowedValues, true
}

// HasAllowedValues returns a boolean if a field has been set.
func (o *ServiceElement) HasAllowedValues() bool {
	if o != nil && o.AllowedValues != nil {
		return true
	}

	return false
}

// SetAllowedValues gets a reference to the given []map[string]map[string]interface{} and assigns it to the AllowedValues field.
func (o *ServiceElement) SetAllowedValues(v []map[string]map[string]interface{}) {
	o.AllowedValues = &v
}

// GetMandatory returns the Mandatory field value if set, zero value otherwise.
func (o *ServiceElement) GetMandatory() bool {
	if o == nil || o.Mandatory == nil {
		var ret bool
		return ret
	}
	return *o.Mandatory
}

// GetMandatoryOk returns a tuple with the Mandatory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceElement) GetMandatoryOk() (*bool, bool) {
	if o == nil || o.Mandatory == nil {
		return nil, false
	}
	return o.Mandatory, true
}

// HasMandatory returns a boolean if a field has been set.
func (o *ServiceElement) HasMandatory() bool {
	if o != nil && o.Mandatory != nil {
		return true
	}

	return false
}

// SetMandatory gets a reference to the given bool and assigns it to the Mandatory field.
func (o *ServiceElement) SetMandatory(v bool) {
	o.Mandatory = &v
}

// GetSection returns the Section field value if set, zero value otherwise.
func (o *ServiceElement) GetSection() string {
	if o == nil || o.Section == nil {
		var ret string
		return ret
	}
	return *o.Section
}

// GetSectionOk returns a tuple with the Section field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceElement) GetSectionOk() (*string, bool) {
	if o == nil || o.Section == nil {
		return nil, false
	}
	return o.Section, true
}

// HasSection returns a boolean if a field has been set.
func (o *ServiceElement) HasSection() bool {
	if o != nil && o.Section != nil {
		return true
	}

	return false
}

// SetSection gets a reference to the given string and assigns it to the Section field.
func (o *ServiceElement) SetSection(v string) {
	o.Section = &v
}

// GetBillable returns the Billable field value if set, zero value otherwise.
func (o *ServiceElement) GetBillable() bool {
	if o == nil || o.Billable == nil {
		var ret bool
		return ret
	}
	return *o.Billable
}

// GetBillableOk returns a tuple with the Billable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceElement) GetBillableOk() (*bool, bool) {
	if o == nil || o.Billable == nil {
		return nil, false
	}
	return o.Billable, true
}

// HasBillable returns a boolean if a field has been set.
func (o *ServiceElement) HasBillable() bool {
	if o != nil && o.Billable != nil {
		return true
	}

	return false
}

// SetBillable gets a reference to the given bool and assigns it to the Billable field.
func (o *ServiceElement) SetBillable(v bool) {
	o.Billable = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *ServiceElement) GetHidden() bool {
	if o == nil || o.Hidden == nil {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceElement) GetHiddenOk() (*bool, bool) {
	if o == nil || o.Hidden == nil {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *ServiceElement) HasHidden() bool {
	if o != nil && o.Hidden != nil {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *ServiceElement) SetHidden(v bool) {
	o.Hidden = &v
}

// GetParentName returns the ParentName field value if set, zero value otherwise.
func (o *ServiceElement) GetParentName() string {
	if o == nil || o.ParentName == nil {
		var ret string
		return ret
	}
	return *o.ParentName
}

// GetParentNameOk returns a tuple with the ParentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceElement) GetParentNameOk() (*string, bool) {
	if o == nil || o.ParentName == nil {
		return nil, false
	}
	return o.ParentName, true
}

// HasParentName returns a boolean if a field has been set.
func (o *ServiceElement) HasParentName() bool {
	if o != nil && o.ParentName != nil {
		return true
	}

	return false
}

// SetParentName gets a reference to the given string and assigns it to the ParentName field.
func (o *ServiceElement) SetParentName(v string) {
	o.ParentName = &v
}

// GetSupported returns the Supported field value if set, zero value otherwise.
func (o *ServiceElement) GetSupported() bool {
	if o == nil || o.Supported == nil {
		var ret bool
		return ret
	}
	return *o.Supported
}

// GetSupportedOk returns a tuple with the Supported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceElement) GetSupportedOk() (*bool, bool) {
	if o == nil || o.Supported == nil {
		return nil, false
	}
	return o.Supported, true
}

// HasSupported returns a boolean if a field has been set.
func (o *ServiceElement) HasSupported() bool {
	if o != nil && o.Supported != nil {
		return true
	}

	return false
}

// SetSupported gets a reference to the given bool and assigns it to the Supported field.
func (o *ServiceElement) SetSupported(v bool) {
	o.Supported = &v
}

// GetDynamicData returns the DynamicData field value if set, zero value otherwise.
func (o *ServiceElement) GetDynamicData() bool {
	if o == nil || o.DynamicData == nil {
		var ret bool
		return ret
	}
	return *o.DynamicData
}

// GetDynamicDataOk returns a tuple with the DynamicData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceElement) GetDynamicDataOk() (*bool, bool) {
	if o == nil || o.DynamicData == nil {
		return nil, false
	}
	return o.DynamicData, true
}

// HasDynamicData returns a boolean if a field has been set.
func (o *ServiceElement) HasDynamicData() bool {
	if o != nil && o.DynamicData != nil {
		return true
	}

	return false
}

// SetDynamicData gets a reference to the given bool and assigns it to the DynamicData field.
func (o *ServiceElement) SetDynamicData(v bool) {
	o.DynamicData = &v
}

// GetMinValue returns the MinValue field value if set, zero value otherwise.
func (o *ServiceElement) GetMinValue() int32 {
	if o == nil || o.MinValue == nil {
		var ret int32
		return ret
	}
	return *o.MinValue
}

// GetMinValueOk returns a tuple with the MinValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceElement) GetMinValueOk() (*int32, bool) {
	if o == nil || o.MinValue == nil {
		return nil, false
	}
	return o.MinValue, true
}

// HasMinValue returns a boolean if a field has been set.
func (o *ServiceElement) HasMinValue() bool {
	if o != nil && o.MinValue != nil {
		return true
	}

	return false
}

// SetMinValue gets a reference to the given int32 and assigns it to the MinValue field.
func (o *ServiceElement) SetMinValue(v int32) {
	o.MinValue = &v
}

// GetMaxValue returns the MaxValue field value if set, zero value otherwise.
func (o *ServiceElement) GetMaxValue() int32 {
	if o == nil || o.MaxValue == nil {
		var ret int32
		return ret
	}
	return *o.MaxValue
}

// GetMaxValueOk returns a tuple with the MaxValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceElement) GetMaxValueOk() (*int32, bool) {
	if o == nil || o.MaxValue == nil {
		return nil, false
	}
	return o.MaxValue, true
}

// HasMaxValue returns a boolean if a field has been set.
func (o *ServiceElement) HasMaxValue() bool {
	if o != nil && o.MaxValue != nil {
		return true
	}

	return false
}

// SetMaxValue gets a reference to the given int32 and assigns it to the MaxValue field.
func (o *ServiceElement) SetMaxValue(v int32) {
	o.MaxValue = &v
}

// GetStepSize returns the StepSize field value if set, zero value otherwise.
func (o *ServiceElement) GetStepSize() int32 {
	if o == nil || o.StepSize == nil {
		var ret int32
		return ret
	}
	return *o.StepSize
}

// GetStepSizeOk returns a tuple with the StepSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceElement) GetStepSizeOk() (*int32, bool) {
	if o == nil || o.StepSize == nil {
		return nil, false
	}
	return o.StepSize, true
}

// HasStepSize returns a boolean if a field has been set.
func (o *ServiceElement) HasStepSize() bool {
	if o != nil && o.StepSize != nil {
		return true
	}

	return false
}

// SetStepSize gets a reference to the given int32 and assigns it to the StepSize field.
func (o *ServiceElement) SetStepSize(v int32) {
	o.StepSize = &v
}

// GetPricingAtttributes returns the PricingAtttributes field value if set, zero value otherwise.
func (o *ServiceElement) GetPricingAtttributes() ServiceElementPrice {
	if o == nil || o.PricingAtttributes == nil {
		var ret ServiceElementPrice
		return ret
	}
	return *o.PricingAtttributes
}

// GetPricingAtttributesOk returns a tuple with the PricingAtttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceElement) GetPricingAtttributesOk() (*ServiceElementPrice, bool) {
	if o == nil || o.PricingAtttributes == nil {
		return nil, false
	}
	return o.PricingAtttributes, true
}

// HasPricingAtttributes returns a boolean if a field has been set.
func (o *ServiceElement) HasPricingAtttributes() bool {
	if o != nil && o.PricingAtttributes != nil {
		return true
	}

	return false
}

// SetPricingAtttributes gets a reference to the given ServiceElementPrice and assigns it to the PricingAtttributes field.
func (o *ServiceElement) SetPricingAtttributes(v ServiceElementPrice) {
	o.PricingAtttributes = &v
}

// GetChildElements returns the ChildElements field value if set, zero value otherwise.
func (o *ServiceElement) GetChildElements() []ServiceElement {
	if o == nil || o.ChildElements == nil {
		var ret []ServiceElement
		return ret
	}
	return *o.ChildElements
}

// GetChildElementsOk returns a tuple with the ChildElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceElement) GetChildElementsOk() (*[]ServiceElement, bool) {
	if o == nil || o.ChildElements == nil {
		return nil, false
	}
	return o.ChildElements, true
}

// HasChildElements returns a boolean if a field has been set.
func (o *ServiceElement) HasChildElements() bool {
	if o != nil && o.ChildElements != nil {
		return true
	}

	return false
}

// SetChildElements gets a reference to the given []ServiceElement and assigns it to the ChildElements field.
func (o *ServiceElement) SetChildElements(v []ServiceElement) {
	o.ChildElements = &v
}

func (o ServiceElement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Header != nil {
		toSerialize["header"] = o.Header
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Hint != nil {
		toSerialize["hint"] = o.Hint
	}
	if o.InputType != nil {
		toSerialize["inputType"] = o.InputType
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.MaxLimit != nil {
		toSerialize["maxLimit"] = o.MaxLimit
	}
	if o.MinLimit != nil {
		toSerialize["minLimit"] = o.MinLimit
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.ValueList != nil {
		toSerialize["valueList"] = o.ValueList
	}
	if o.AllowedOptionValues != nil {
		toSerialize["allowedOptionValues"] = o.AllowedOptionValues
	}
	if o.AllowedValues != nil {
		toSerialize["allowedValues"] = o.AllowedValues
	}
	if o.Mandatory != nil {
		toSerialize["mandatory"] = o.Mandatory
	}
	if o.Section != nil {
		toSerialize["section"] = o.Section
	}
	if o.Billable != nil {
		toSerialize["billable"] = o.Billable
	}
	if o.Hidden != nil {
		toSerialize["hidden"] = o.Hidden
	}
	if o.ParentName != nil {
		toSerialize["parentName"] = o.ParentName
	}
	if o.Supported != nil {
		toSerialize["supported"] = o.Supported
	}
	if o.DynamicData != nil {
		toSerialize["dynamicData"] = o.DynamicData
	}
	if o.MinValue != nil {
		toSerialize["minValue"] = o.MinValue
	}
	if o.MaxValue != nil {
		toSerialize["maxValue"] = o.MaxValue
	}
	if o.StepSize != nil {
		toSerialize["stepSize"] = o.StepSize
	}
	if o.PricingAtttributes != nil {
		toSerialize["pricingAtttributes"] = o.PricingAtttributes
	}
	if o.ChildElements != nil {
		toSerialize["childElements"] = o.ChildElements
	}
	return json.Marshal(toSerialize)
}

type NullableServiceElement struct {
	value *ServiceElement
	isSet bool
}

func (v NullableServiceElement) Get() *ServiceElement {
	return v.value
}

func (v *NullableServiceElement) Set(val *ServiceElement) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceElement) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceElement(val *ServiceElement) *NullableServiceElement {
	return &NullableServiceElement{value: val, isSet: true}
}

func (v NullableServiceElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


