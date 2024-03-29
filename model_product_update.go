/*
MSX SDK

MSX SDK client.

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msxsdk

import (
	"encoding/json"
)

// ProductUpdate struct for ProductUpdate
type ProductUpdate struct {
	Name string `json:"name"`
	Label string `json:"label"`
	Version int32 `json:"version"`
	Description string `json:"description"`
	Image string `json:"image"`
	MultipleInstanceAllowed *bool `json:"multipleInstanceAllowed,omitempty"`
	Price *string `json:"price,omitempty"`
	DisplayOrder *int32 `json:"displayOrder,omitempty"`
	Active *bool `json:"active,omitempty"`
	OrderLimit *int32 `json:"orderLimit,omitempty"`
	Options []ServiceElement `json:"options,omitempty"`
	Properties []ServiceElement `json:"properties,omitempty"`
	Configuration map[string]string `json:"configuration,omitempty"`
	IsResource NullableBool `json:"isResource,omitempty"`
	HasChildren NullableBool `json:"hasChildren,omitempty"`
	ParentId *string `json:"parentId,omitempty"`
	ServiceExtensions []NSOConfigDataXPath `json:"serviceExtensions,omitempty"`
	ServiceConfigQueryRootXPaths []string `json:"serviceConfigQueryRootXPaths,omitempty"`
	UiConfig *ServiceUIConfig `json:"uiConfig,omitempty"`
	SlmUiConfig NullableServiceSLMUIConfig `json:"slmUiConfig,omitempty"`
	ExternalId NullableString `json:"externalId,omitempty"`
	Tags []string `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductUpdate ProductUpdate

// NewProductUpdate instantiates a new ProductUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductUpdate(name string, label string, version int32, description string, image string) *ProductUpdate {
	this := ProductUpdate{}
	this.Name = name
	this.Label = label
	this.Version = version
	this.Description = description
	this.Image = image
	return &this
}

// NewProductUpdateWithDefaults instantiates a new ProductUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductUpdateWithDefaults() *ProductUpdate {
	this := ProductUpdate{}
	return &this
}

// GetName returns the Name field value
func (o *ProductUpdate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProductUpdate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProductUpdate) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value
func (o *ProductUpdate) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ProductUpdate) GetLabelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *ProductUpdate) SetLabel(v string) {
	o.Label = v
}

// GetVersion returns the Version field value
func (o *ProductUpdate) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ProductUpdate) GetVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ProductUpdate) SetVersion(v int32) {
	o.Version = v
}

// GetDescription returns the Description field value
func (o *ProductUpdate) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ProductUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ProductUpdate) SetDescription(v string) {
	o.Description = v
}

// GetImage returns the Image field value
func (o *ProductUpdate) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *ProductUpdate) GetImageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *ProductUpdate) SetImage(v string) {
	o.Image = v
}

// GetMultipleInstanceAllowed returns the MultipleInstanceAllowed field value if set, zero value otherwise.
func (o *ProductUpdate) GetMultipleInstanceAllowed() bool {
	if o == nil || o.MultipleInstanceAllowed == nil {
		var ret bool
		return ret
	}
	return *o.MultipleInstanceAllowed
}

// GetMultipleInstanceAllowedOk returns a tuple with the MultipleInstanceAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdate) GetMultipleInstanceAllowedOk() (*bool, bool) {
	if o == nil || o.MultipleInstanceAllowed == nil {
		return nil, false
	}
	return o.MultipleInstanceAllowed, true
}

// HasMultipleInstanceAllowed returns a boolean if a field has been set.
func (o *ProductUpdate) HasMultipleInstanceAllowed() bool {
	if o != nil && o.MultipleInstanceAllowed != nil {
		return true
	}

	return false
}

// SetMultipleInstanceAllowed gets a reference to the given bool and assigns it to the MultipleInstanceAllowed field.
func (o *ProductUpdate) SetMultipleInstanceAllowed(v bool) {
	o.MultipleInstanceAllowed = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ProductUpdate) GetPrice() string {
	if o == nil || o.Price == nil {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdate) GetPriceOk() (*string, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ProductUpdate) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *ProductUpdate) SetPrice(v string) {
	o.Price = &v
}

// GetDisplayOrder returns the DisplayOrder field value if set, zero value otherwise.
func (o *ProductUpdate) GetDisplayOrder() int32 {
	if o == nil || o.DisplayOrder == nil {
		var ret int32
		return ret
	}
	return *o.DisplayOrder
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdate) GetDisplayOrderOk() (*int32, bool) {
	if o == nil || o.DisplayOrder == nil {
		return nil, false
	}
	return o.DisplayOrder, true
}

// HasDisplayOrder returns a boolean if a field has been set.
func (o *ProductUpdate) HasDisplayOrder() bool {
	if o != nil && o.DisplayOrder != nil {
		return true
	}

	return false
}

// SetDisplayOrder gets a reference to the given int32 and assigns it to the DisplayOrder field.
func (o *ProductUpdate) SetDisplayOrder(v int32) {
	o.DisplayOrder = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ProductUpdate) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdate) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ProductUpdate) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ProductUpdate) SetActive(v bool) {
	o.Active = &v
}

// GetOrderLimit returns the OrderLimit field value if set, zero value otherwise.
func (o *ProductUpdate) GetOrderLimit() int32 {
	if o == nil || o.OrderLimit == nil {
		var ret int32
		return ret
	}
	return *o.OrderLimit
}

// GetOrderLimitOk returns a tuple with the OrderLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdate) GetOrderLimitOk() (*int32, bool) {
	if o == nil || o.OrderLimit == nil {
		return nil, false
	}
	return o.OrderLimit, true
}

// HasOrderLimit returns a boolean if a field has been set.
func (o *ProductUpdate) HasOrderLimit() bool {
	if o != nil && o.OrderLimit != nil {
		return true
	}

	return false
}

// SetOrderLimit gets a reference to the given int32 and assigns it to the OrderLimit field.
func (o *ProductUpdate) SetOrderLimit(v int32) {
	o.OrderLimit = &v
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductUpdate) GetOptions() []ServiceElement {
	if o == nil  {
		var ret []ServiceElement
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductUpdate) GetOptionsOk() ([]ServiceElement, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ProductUpdate) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []ServiceElement and assigns it to the Options field.
func (o *ProductUpdate) SetOptions(v []ServiceElement) {
	o.Options = v
}

// GetProperties returns the Properties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductUpdate) GetProperties() []ServiceElement {
	if o == nil  {
		var ret []ServiceElement
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductUpdate) GetPropertiesOk() ([]ServiceElement, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ProductUpdate) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []ServiceElement and assigns it to the Properties field.
func (o *ProductUpdate) SetProperties(v []ServiceElement) {
	o.Properties = v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductUpdate) GetConfiguration() map[string]string {
	if o == nil  {
		var ret map[string]string
		return ret
	}
	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductUpdate) GetConfigurationOk() (*map[string]string, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return &o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *ProductUpdate) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given map[string]string and assigns it to the Configuration field.
func (o *ProductUpdate) SetConfiguration(v map[string]string) {
	o.Configuration = v
}

// GetIsResource returns the IsResource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductUpdate) GetIsResource() bool {
	if o == nil || o.IsResource.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsResource.Get()
}

// GetIsResourceOk returns a tuple with the IsResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductUpdate) GetIsResourceOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsResource.Get(), o.IsResource.IsSet()
}

// HasIsResource returns a boolean if a field has been set.
func (o *ProductUpdate) HasIsResource() bool {
	if o != nil && o.IsResource.IsSet() {
		return true
	}

	return false
}

// SetIsResource gets a reference to the given NullableBool and assigns it to the IsResource field.
func (o *ProductUpdate) SetIsResource(v bool) {
	o.IsResource.Set(&v)
}
// SetIsResourceNil sets the value for IsResource to be an explicit nil
func (o *ProductUpdate) SetIsResourceNil() {
	o.IsResource.Set(nil)
}

// UnsetIsResource ensures that no value is present for IsResource, not even an explicit nil
func (o *ProductUpdate) UnsetIsResource() {
	o.IsResource.Unset()
}

// GetHasChildren returns the HasChildren field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductUpdate) GetHasChildren() bool {
	if o == nil || o.HasChildren.Get() == nil {
		var ret bool
		return ret
	}
	return *o.HasChildren.Get()
}

// GetHasChildrenOk returns a tuple with the HasChildren field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductUpdate) GetHasChildrenOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HasChildren.Get(), o.HasChildren.IsSet()
}

// HasHasChildren returns a boolean if a field has been set.
func (o *ProductUpdate) HasHasChildren() bool {
	if o != nil && o.HasChildren.IsSet() {
		return true
	}

	return false
}

// SetHasChildren gets a reference to the given NullableBool and assigns it to the HasChildren field.
func (o *ProductUpdate) SetHasChildren(v bool) {
	o.HasChildren.Set(&v)
}
// SetHasChildrenNil sets the value for HasChildren to be an explicit nil
func (o *ProductUpdate) SetHasChildrenNil() {
	o.HasChildren.Set(nil)
}

// UnsetHasChildren ensures that no value is present for HasChildren, not even an explicit nil
func (o *ProductUpdate) UnsetHasChildren() {
	o.HasChildren.Unset()
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *ProductUpdate) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdate) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *ProductUpdate) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *ProductUpdate) SetParentId(v string) {
	o.ParentId = &v
}

// GetServiceExtensions returns the ServiceExtensions field value if set, zero value otherwise.
func (o *ProductUpdate) GetServiceExtensions() []NSOConfigDataXPath {
	if o == nil || o.ServiceExtensions == nil {
		var ret []NSOConfigDataXPath
		return ret
	}
	return o.ServiceExtensions
}

// GetServiceExtensionsOk returns a tuple with the ServiceExtensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdate) GetServiceExtensionsOk() ([]NSOConfigDataXPath, bool) {
	if o == nil || o.ServiceExtensions == nil {
		return nil, false
	}
	return o.ServiceExtensions, true
}

// HasServiceExtensions returns a boolean if a field has been set.
func (o *ProductUpdate) HasServiceExtensions() bool {
	if o != nil && o.ServiceExtensions != nil {
		return true
	}

	return false
}

// SetServiceExtensions gets a reference to the given []NSOConfigDataXPath and assigns it to the ServiceExtensions field.
func (o *ProductUpdate) SetServiceExtensions(v []NSOConfigDataXPath) {
	o.ServiceExtensions = v
}

// GetServiceConfigQueryRootXPaths returns the ServiceConfigQueryRootXPaths field value if set, zero value otherwise.
func (o *ProductUpdate) GetServiceConfigQueryRootXPaths() []string {
	if o == nil || o.ServiceConfigQueryRootXPaths == nil {
		var ret []string
		return ret
	}
	return o.ServiceConfigQueryRootXPaths
}

// GetServiceConfigQueryRootXPathsOk returns a tuple with the ServiceConfigQueryRootXPaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdate) GetServiceConfigQueryRootXPathsOk() ([]string, bool) {
	if o == nil || o.ServiceConfigQueryRootXPaths == nil {
		return nil, false
	}
	return o.ServiceConfigQueryRootXPaths, true
}

// HasServiceConfigQueryRootXPaths returns a boolean if a field has been set.
func (o *ProductUpdate) HasServiceConfigQueryRootXPaths() bool {
	if o != nil && o.ServiceConfigQueryRootXPaths != nil {
		return true
	}

	return false
}

// SetServiceConfigQueryRootXPaths gets a reference to the given []string and assigns it to the ServiceConfigQueryRootXPaths field.
func (o *ProductUpdate) SetServiceConfigQueryRootXPaths(v []string) {
	o.ServiceConfigQueryRootXPaths = v
}

// GetUiConfig returns the UiConfig field value if set, zero value otherwise.
func (o *ProductUpdate) GetUiConfig() ServiceUIConfig {
	if o == nil || o.UiConfig == nil {
		var ret ServiceUIConfig
		return ret
	}
	return *o.UiConfig
}

// GetUiConfigOk returns a tuple with the UiConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdate) GetUiConfigOk() (*ServiceUIConfig, bool) {
	if o == nil || o.UiConfig == nil {
		return nil, false
	}
	return o.UiConfig, true
}

// HasUiConfig returns a boolean if a field has been set.
func (o *ProductUpdate) HasUiConfig() bool {
	if o != nil && o.UiConfig != nil {
		return true
	}

	return false
}

// SetUiConfig gets a reference to the given ServiceUIConfig and assigns it to the UiConfig field.
func (o *ProductUpdate) SetUiConfig(v ServiceUIConfig) {
	o.UiConfig = &v
}

// GetSlmUiConfig returns the SlmUiConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductUpdate) GetSlmUiConfig() ServiceSLMUIConfig {
	if o == nil || o.SlmUiConfig.Get() == nil {
		var ret ServiceSLMUIConfig
		return ret
	}
	return *o.SlmUiConfig.Get()
}

// GetSlmUiConfigOk returns a tuple with the SlmUiConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductUpdate) GetSlmUiConfigOk() (*ServiceSLMUIConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SlmUiConfig.Get(), o.SlmUiConfig.IsSet()
}

// HasSlmUiConfig returns a boolean if a field has been set.
func (o *ProductUpdate) HasSlmUiConfig() bool {
	if o != nil && o.SlmUiConfig.IsSet() {
		return true
	}

	return false
}

// SetSlmUiConfig gets a reference to the given NullableServiceSLMUIConfig and assigns it to the SlmUiConfig field.
func (o *ProductUpdate) SetSlmUiConfig(v ServiceSLMUIConfig) {
	o.SlmUiConfig.Set(&v)
}
// SetSlmUiConfigNil sets the value for SlmUiConfig to be an explicit nil
func (o *ProductUpdate) SetSlmUiConfigNil() {
	o.SlmUiConfig.Set(nil)
}

// UnsetSlmUiConfig ensures that no value is present for SlmUiConfig, not even an explicit nil
func (o *ProductUpdate) UnsetSlmUiConfig() {
	o.SlmUiConfig.Unset()
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductUpdate) GetExternalId() string {
	if o == nil || o.ExternalId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductUpdate) GetExternalIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// HasExternalId returns a boolean if a field has been set.
func (o *ProductUpdate) HasExternalId() bool {
	if o != nil && o.ExternalId.IsSet() {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given NullableString and assigns it to the ExternalId field.
func (o *ProductUpdate) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}
// SetExternalIdNil sets the value for ExternalId to be an explicit nil
func (o *ProductUpdate) SetExternalIdNil() {
	o.ExternalId.Set(nil)
}

// UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
func (o *ProductUpdate) UnsetExternalId() {
	o.ExternalId.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductUpdate) GetTags() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductUpdate) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ProductUpdate) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ProductUpdate) SetTags(v []string) {
	o.Tags = v
}

func (o ProductUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["image"] = o.Image
	}
	if o.MultipleInstanceAllowed != nil {
		toSerialize["multipleInstanceAllowed"] = o.MultipleInstanceAllowed
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.DisplayOrder != nil {
		toSerialize["displayOrder"] = o.DisplayOrder
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.OrderLimit != nil {
		toSerialize["orderLimit"] = o.OrderLimit
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	if o.IsResource.IsSet() {
		toSerialize["isResource"] = o.IsResource.Get()
	}
	if o.HasChildren.IsSet() {
		toSerialize["hasChildren"] = o.HasChildren.Get()
	}
	if o.ParentId != nil {
		toSerialize["parentId"] = o.ParentId
	}
	if o.ServiceExtensions != nil {
		toSerialize["serviceExtensions"] = o.ServiceExtensions
	}
	if o.ServiceConfigQueryRootXPaths != nil {
		toSerialize["serviceConfigQueryRootXPaths"] = o.ServiceConfigQueryRootXPaths
	}
	if o.UiConfig != nil {
		toSerialize["uiConfig"] = o.UiConfig
	}
	if o.SlmUiConfig.IsSet() {
		toSerialize["slmUiConfig"] = o.SlmUiConfig.Get()
	}
	if o.ExternalId.IsSet() {
		toSerialize["externalId"] = o.ExternalId.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProductUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varProductUpdate := _ProductUpdate{}

	if err = json.Unmarshal(bytes, &varProductUpdate); err == nil {
		*o = ProductUpdate(varProductUpdate)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "version")
		delete(additionalProperties, "description")
		delete(additionalProperties, "image")
		delete(additionalProperties, "multipleInstanceAllowed")
		delete(additionalProperties, "price")
		delete(additionalProperties, "displayOrder")
		delete(additionalProperties, "active")
		delete(additionalProperties, "orderLimit")
		delete(additionalProperties, "options")
		delete(additionalProperties, "properties")
		delete(additionalProperties, "configuration")
		delete(additionalProperties, "isResource")
		delete(additionalProperties, "hasChildren")
		delete(additionalProperties, "parentId")
		delete(additionalProperties, "serviceExtensions")
		delete(additionalProperties, "serviceConfigQueryRootXPaths")
		delete(additionalProperties, "uiConfig")
		delete(additionalProperties, "slmUiConfig")
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProductUpdate struct {
	value *ProductUpdate
	isSet bool
}

func (v NullableProductUpdate) Get() *ProductUpdate {
	return v.value
}

func (v *NullableProductUpdate) Set(val *ProductUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableProductUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableProductUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductUpdate(val *ProductUpdate) *NullableProductUpdate {
	return &NullableProductUpdate{value: val, isSet: true}
}

func (v NullableProductUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


