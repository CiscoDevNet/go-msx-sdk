# ServiceElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Header** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Hint** | Pointer to **string** |  | [optional] 
**InputType** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Component** | Pointer to **string** |  | [optional] 
**MaxLimit** | Pointer to **string** |  | [optional] 
**MinLimit** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**ValueList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**AllowedOptionValues** | Pointer to **[]string** |  | [optional] 
**AllowedValues** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Mandatory** | Pointer to **bool** |  | [optional] 
**Section** | Pointer to **string** |  | [optional] 
**Billable** | Pointer to **bool** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**ParentName** | Pointer to **string** |  | [optional] 
**Supported** | Pointer to **bool** |  | [optional] 
**DynamicData** | Pointer to **bool** |  | [optional] 
**MinValue** | Pointer to **int32** |  | [optional] 
**MaxValue** | Pointer to **int32** |  | [optional] 
**StepSize** | Pointer to **int32** |  | [optional] 
**PricingAtttributes** | Pointer to [**NullableServiceElementPrice**](ServiceElementPrice.md) |  | [optional] 
**ChildElements** | Pointer to [**[]ServiceElement**](ServiceElement.md) |  | [optional] 

## Methods

### NewServiceElement

`func NewServiceElement() *ServiceElement`

NewServiceElement instantiates a new ServiceElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceElementWithDefaults

`func NewServiceElementWithDefaults() *ServiceElement`

NewServiceElementWithDefaults instantiates a new ServiceElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServiceElement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceElement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceElement) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceElement) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *ServiceElement) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServiceElement) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServiceElement) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ServiceElement) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetHeader

`func (o *ServiceElement) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *ServiceElement) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *ServiceElement) SetHeader(v string)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *ServiceElement) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetDescription

`func (o *ServiceElement) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceElement) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceElement) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceElement) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHint

`func (o *ServiceElement) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *ServiceElement) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *ServiceElement) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *ServiceElement) HasHint() bool`

HasHint returns a boolean if a field has been set.

### GetInputType

`func (o *ServiceElement) GetInputType() string`

GetInputType returns the InputType field if non-nil, zero value otherwise.

### GetInputTypeOk

`func (o *ServiceElement) GetInputTypeOk() (*string, bool)`

GetInputTypeOk returns a tuple with the InputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputType

`func (o *ServiceElement) SetInputType(v string)`

SetInputType sets InputType field to given value.

### HasInputType

`func (o *ServiceElement) HasInputType() bool`

HasInputType returns a boolean if a field has been set.

### GetType

`func (o *ServiceElement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceElement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceElement) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ServiceElement) HasType() bool`

HasType returns a boolean if a field has been set.

### GetComponent

`func (o *ServiceElement) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ServiceElement) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ServiceElement) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *ServiceElement) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetMaxLimit

`func (o *ServiceElement) GetMaxLimit() string`

GetMaxLimit returns the MaxLimit field if non-nil, zero value otherwise.

### GetMaxLimitOk

`func (o *ServiceElement) GetMaxLimitOk() (*string, bool)`

GetMaxLimitOk returns a tuple with the MaxLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLimit

`func (o *ServiceElement) SetMaxLimit(v string)`

SetMaxLimit sets MaxLimit field to given value.

### HasMaxLimit

`func (o *ServiceElement) HasMaxLimit() bool`

HasMaxLimit returns a boolean if a field has been set.

### GetMinLimit

`func (o *ServiceElement) GetMinLimit() string`

GetMinLimit returns the MinLimit field if non-nil, zero value otherwise.

### GetMinLimitOk

`func (o *ServiceElement) GetMinLimitOk() (*string, bool)`

GetMinLimitOk returns a tuple with the MinLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLimit

`func (o *ServiceElement) SetMinLimit(v string)`

SetMinLimit sets MinLimit field to given value.

### HasMinLimit

`func (o *ServiceElement) HasMinLimit() bool`

HasMinLimit returns a boolean if a field has been set.

### GetValue

`func (o *ServiceElement) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ServiceElement) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ServiceElement) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ServiceElement) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueList

`func (o *ServiceElement) GetValueList() []map[string]interface{}`

GetValueList returns the ValueList field if non-nil, zero value otherwise.

### GetValueListOk

`func (o *ServiceElement) GetValueListOk() (*[]map[string]interface{}, bool)`

GetValueListOk returns a tuple with the ValueList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueList

`func (o *ServiceElement) SetValueList(v []map[string]interface{})`

SetValueList sets ValueList field to given value.

### HasValueList

`func (o *ServiceElement) HasValueList() bool`

HasValueList returns a boolean if a field has been set.

### SetValueListNil

`func (o *ServiceElement) SetValueListNil(b bool)`

 SetValueListNil sets the value for ValueList to be an explicit nil

### UnsetValueList
`func (o *ServiceElement) UnsetValueList()`

UnsetValueList ensures that no value is present for ValueList, not even an explicit nil
### GetAllowedOptionValues

`func (o *ServiceElement) GetAllowedOptionValues() []string`

GetAllowedOptionValues returns the AllowedOptionValues field if non-nil, zero value otherwise.

### GetAllowedOptionValuesOk

`func (o *ServiceElement) GetAllowedOptionValuesOk() (*[]string, bool)`

GetAllowedOptionValuesOk returns a tuple with the AllowedOptionValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOptionValues

`func (o *ServiceElement) SetAllowedOptionValues(v []string)`

SetAllowedOptionValues sets AllowedOptionValues field to given value.

### HasAllowedOptionValues

`func (o *ServiceElement) HasAllowedOptionValues() bool`

HasAllowedOptionValues returns a boolean if a field has been set.

### GetAllowedValues

`func (o *ServiceElement) GetAllowedValues() []map[string]interface{}`

GetAllowedValues returns the AllowedValues field if non-nil, zero value otherwise.

### GetAllowedValuesOk

`func (o *ServiceElement) GetAllowedValuesOk() (*[]map[string]interface{}, bool)`

GetAllowedValuesOk returns a tuple with the AllowedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValues

`func (o *ServiceElement) SetAllowedValues(v []map[string]interface{})`

SetAllowedValues sets AllowedValues field to given value.

### HasAllowedValues

`func (o *ServiceElement) HasAllowedValues() bool`

HasAllowedValues returns a boolean if a field has been set.

### SetAllowedValuesNil

`func (o *ServiceElement) SetAllowedValuesNil(b bool)`

 SetAllowedValuesNil sets the value for AllowedValues to be an explicit nil

### UnsetAllowedValues
`func (o *ServiceElement) UnsetAllowedValues()`

UnsetAllowedValues ensures that no value is present for AllowedValues, not even an explicit nil
### GetMandatory

`func (o *ServiceElement) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *ServiceElement) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *ServiceElement) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.

### HasMandatory

`func (o *ServiceElement) HasMandatory() bool`

HasMandatory returns a boolean if a field has been set.

### GetSection

`func (o *ServiceElement) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *ServiceElement) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *ServiceElement) SetSection(v string)`

SetSection sets Section field to given value.

### HasSection

`func (o *ServiceElement) HasSection() bool`

HasSection returns a boolean if a field has been set.

### GetBillable

`func (o *ServiceElement) GetBillable() bool`

GetBillable returns the Billable field if non-nil, zero value otherwise.

### GetBillableOk

`func (o *ServiceElement) GetBillableOk() (*bool, bool)`

GetBillableOk returns a tuple with the Billable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillable

`func (o *ServiceElement) SetBillable(v bool)`

SetBillable sets Billable field to given value.

### HasBillable

`func (o *ServiceElement) HasBillable() bool`

HasBillable returns a boolean if a field has been set.

### GetHidden

`func (o *ServiceElement) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *ServiceElement) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *ServiceElement) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *ServiceElement) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetParentName

`func (o *ServiceElement) GetParentName() string`

GetParentName returns the ParentName field if non-nil, zero value otherwise.

### GetParentNameOk

`func (o *ServiceElement) GetParentNameOk() (*string, bool)`

GetParentNameOk returns a tuple with the ParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentName

`func (o *ServiceElement) SetParentName(v string)`

SetParentName sets ParentName field to given value.

### HasParentName

`func (o *ServiceElement) HasParentName() bool`

HasParentName returns a boolean if a field has been set.

### GetSupported

`func (o *ServiceElement) GetSupported() bool`

GetSupported returns the Supported field if non-nil, zero value otherwise.

### GetSupportedOk

`func (o *ServiceElement) GetSupportedOk() (*bool, bool)`

GetSupportedOk returns a tuple with the Supported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupported

`func (o *ServiceElement) SetSupported(v bool)`

SetSupported sets Supported field to given value.

### HasSupported

`func (o *ServiceElement) HasSupported() bool`

HasSupported returns a boolean if a field has been set.

### GetDynamicData

`func (o *ServiceElement) GetDynamicData() bool`

GetDynamicData returns the DynamicData field if non-nil, zero value otherwise.

### GetDynamicDataOk

`func (o *ServiceElement) GetDynamicDataOk() (*bool, bool)`

GetDynamicDataOk returns a tuple with the DynamicData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicData

`func (o *ServiceElement) SetDynamicData(v bool)`

SetDynamicData sets DynamicData field to given value.

### HasDynamicData

`func (o *ServiceElement) HasDynamicData() bool`

HasDynamicData returns a boolean if a field has been set.

### GetMinValue

`func (o *ServiceElement) GetMinValue() int32`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *ServiceElement) GetMinValueOk() (*int32, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *ServiceElement) SetMinValue(v int32)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *ServiceElement) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.

### GetMaxValue

`func (o *ServiceElement) GetMaxValue() int32`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *ServiceElement) GetMaxValueOk() (*int32, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *ServiceElement) SetMaxValue(v int32)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *ServiceElement) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### GetStepSize

`func (o *ServiceElement) GetStepSize() int32`

GetStepSize returns the StepSize field if non-nil, zero value otherwise.

### GetStepSizeOk

`func (o *ServiceElement) GetStepSizeOk() (*int32, bool)`

GetStepSizeOk returns a tuple with the StepSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepSize

`func (o *ServiceElement) SetStepSize(v int32)`

SetStepSize sets StepSize field to given value.

### HasStepSize

`func (o *ServiceElement) HasStepSize() bool`

HasStepSize returns a boolean if a field has been set.

### GetPricingAtttributes

`func (o *ServiceElement) GetPricingAtttributes() ServiceElementPrice`

GetPricingAtttributes returns the PricingAtttributes field if non-nil, zero value otherwise.

### GetPricingAtttributesOk

`func (o *ServiceElement) GetPricingAtttributesOk() (*ServiceElementPrice, bool)`

GetPricingAtttributesOk returns a tuple with the PricingAtttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingAtttributes

`func (o *ServiceElement) SetPricingAtttributes(v ServiceElementPrice)`

SetPricingAtttributes sets PricingAtttributes field to given value.

### HasPricingAtttributes

`func (o *ServiceElement) HasPricingAtttributes() bool`

HasPricingAtttributes returns a boolean if a field has been set.

### SetPricingAtttributesNil

`func (o *ServiceElement) SetPricingAtttributesNil(b bool)`

 SetPricingAtttributesNil sets the value for PricingAtttributes to be an explicit nil

### UnsetPricingAtttributes
`func (o *ServiceElement) UnsetPricingAtttributes()`

UnsetPricingAtttributes ensures that no value is present for PricingAtttributes, not even an explicit nil
### GetChildElements

`func (o *ServiceElement) GetChildElements() []ServiceElement`

GetChildElements returns the ChildElements field if non-nil, zero value otherwise.

### GetChildElementsOk

`func (o *ServiceElement) GetChildElementsOk() (*[]ServiceElement, bool)`

GetChildElementsOk returns a tuple with the ChildElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildElements

`func (o *ServiceElement) SetChildElements(v []ServiceElement)`

SetChildElements sets ChildElements field to given value.

### HasChildElements

`func (o *ServiceElement) HasChildElements() bool`

HasChildElements returns a boolean if a field has been set.

### SetChildElementsNil

`func (o *ServiceElement) SetChildElementsNil(b bool)`

 SetChildElementsNil sets the value for ChildElements to be an explicit nil

### UnsetChildElements
`func (o *ServiceElement) UnsetChildElements()`

UnsetChildElements ensures that no value is present for ChildElements, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


