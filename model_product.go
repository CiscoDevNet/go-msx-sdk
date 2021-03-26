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

// Product struct for Product
type Product struct {
	Id *string `json:"id,omitempty"`
	Name string `json:"name"`
	Label string `json:"label"`
	Version int32 `json:"version"`
	Description string `json:"description"`
	Image string `json:"image"`
	MultipleInstanceAllowed *bool `json:"multipleInstanceAllowed,omitempty"`
	Price string `json:"price"`
	DisplayOrder *int32 `json:"displayOrder,omitempty"`
	Active *bool `json:"active,omitempty"`
	OrderLimit *int32 `json:"orderLimit,omitempty"`
	Options []ServiceElement `json:"options"`
	Properties []ServiceElement `json:"properties"`
	Configuration map[string]string `json:"configuration"`
	IsResource bool `json:"isResource"`
	HasChildren bool `json:"hasChildren"`
	ParentId *string `json:"parentId,omitempty"`
	ServiceExtensions *[]NSOConfigDataXPath `json:"serviceExtensions,omitempty"`
	ServiceConfigQueryRootXPaths *[]string `json:"serviceConfigQueryRootXPaths,omitempty"`
	UiConfig *ServiceUIConfig `json:"uiConfig,omitempty"`
}

// NewProduct instantiates a new Product object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct(name string, label string, version int32, description string, image string, price string, options []ServiceElement, properties []ServiceElement, configuration map[string]string, isResource bool, hasChildren bool) *Product {
	this := Product{}
	this.Name = name
	this.Label = label
	this.Version = version
	this.Description = description
	this.Image = image
	this.Price = price
	this.Options = options
	this.Properties = properties
	this.Configuration = configuration
	this.IsResource = isResource
	this.HasChildren = hasChildren
	return &this
}

// NewProductWithDefaults instantiates a new Product object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductWithDefaults() *Product {
	this := Product{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Product) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Product) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Product) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *Product) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Product) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Product) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value
func (o *Product) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *Product) GetLabelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *Product) SetLabel(v string) {
	o.Label = v
}

// GetVersion returns the Version field value
func (o *Product) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *Product) GetVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *Product) SetVersion(v int32) {
	o.Version = v
}

// GetDescription returns the Description field value
func (o *Product) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Product) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Product) SetDescription(v string) {
	o.Description = v
}

// GetImage returns the Image field value
func (o *Product) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *Product) GetImageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *Product) SetImage(v string) {
	o.Image = v
}

// GetMultipleInstanceAllowed returns the MultipleInstanceAllowed field value if set, zero value otherwise.
func (o *Product) GetMultipleInstanceAllowed() bool {
	if o == nil || o.MultipleInstanceAllowed == nil {
		var ret bool
		return ret
	}
	return *o.MultipleInstanceAllowed
}

// GetMultipleInstanceAllowedOk returns a tuple with the MultipleInstanceAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetMultipleInstanceAllowedOk() (*bool, bool) {
	if o == nil || o.MultipleInstanceAllowed == nil {
		return nil, false
	}
	return o.MultipleInstanceAllowed, true
}

// HasMultipleInstanceAllowed returns a boolean if a field has been set.
func (o *Product) HasMultipleInstanceAllowed() bool {
	if o != nil && o.MultipleInstanceAllowed != nil {
		return true
	}

	return false
}

// SetMultipleInstanceAllowed gets a reference to the given bool and assigns it to the MultipleInstanceAllowed field.
func (o *Product) SetMultipleInstanceAllowed(v bool) {
	o.MultipleInstanceAllowed = &v
}

// GetPrice returns the Price field value
func (o *Product) GetPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *Product) GetPriceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *Product) SetPrice(v string) {
	o.Price = v
}

// GetDisplayOrder returns the DisplayOrder field value if set, zero value otherwise.
func (o *Product) GetDisplayOrder() int32 {
	if o == nil || o.DisplayOrder == nil {
		var ret int32
		return ret
	}
	return *o.DisplayOrder
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetDisplayOrderOk() (*int32, bool) {
	if o == nil || o.DisplayOrder == nil {
		return nil, false
	}
	return o.DisplayOrder, true
}

// HasDisplayOrder returns a boolean if a field has been set.
func (o *Product) HasDisplayOrder() bool {
	if o != nil && o.DisplayOrder != nil {
		return true
	}

	return false
}

// SetDisplayOrder gets a reference to the given int32 and assigns it to the DisplayOrder field.
func (o *Product) SetDisplayOrder(v int32) {
	o.DisplayOrder = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *Product) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *Product) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *Product) SetActive(v bool) {
	o.Active = &v
}

// GetOrderLimit returns the OrderLimit field value if set, zero value otherwise.
func (o *Product) GetOrderLimit() int32 {
	if o == nil || o.OrderLimit == nil {
		var ret int32
		return ret
	}
	return *o.OrderLimit
}

// GetOrderLimitOk returns a tuple with the OrderLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetOrderLimitOk() (*int32, bool) {
	if o == nil || o.OrderLimit == nil {
		return nil, false
	}
	return o.OrderLimit, true
}

// HasOrderLimit returns a boolean if a field has been set.
func (o *Product) HasOrderLimit() bool {
	if o != nil && o.OrderLimit != nil {
		return true
	}

	return false
}

// SetOrderLimit gets a reference to the given int32 and assigns it to the OrderLimit field.
func (o *Product) SetOrderLimit(v int32) {
	o.OrderLimit = &v
}

// GetOptions returns the Options field value
func (o *Product) GetOptions() []ServiceElement {
	if o == nil {
		var ret []ServiceElement
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *Product) GetOptionsOk() (*[]ServiceElement, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *Product) SetOptions(v []ServiceElement) {
	o.Options = v
}

// GetProperties returns the Properties field value
func (o *Product) GetProperties() []ServiceElement {
	if o == nil {
		var ret []ServiceElement
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *Product) GetPropertiesOk() (*[]ServiceElement, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *Product) SetProperties(v []ServiceElement) {
	o.Properties = v
}

// GetConfiguration returns the Configuration field value
func (o *Product) GetConfiguration() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value
// and a boolean to check if the value has been set.
func (o *Product) GetConfigurationOk() (*map[string]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Configuration, true
}

// SetConfiguration sets field value
func (o *Product) SetConfiguration(v map[string]string) {
	o.Configuration = v
}

// GetIsResource returns the IsResource field value
func (o *Product) GetIsResource() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsResource
}

// GetIsResourceOk returns a tuple with the IsResource field value
// and a boolean to check if the value has been set.
func (o *Product) GetIsResourceOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsResource, true
}

// SetIsResource sets field value
func (o *Product) SetIsResource(v bool) {
	o.IsResource = v
}

// GetHasChildren returns the HasChildren field value
func (o *Product) GetHasChildren() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasChildren
}

// GetHasChildrenOk returns a tuple with the HasChildren field value
// and a boolean to check if the value has been set.
func (o *Product) GetHasChildrenOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HasChildren, true
}

// SetHasChildren sets field value
func (o *Product) SetHasChildren(v bool) {
	o.HasChildren = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *Product) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *Product) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *Product) SetParentId(v string) {
	o.ParentId = &v
}

// GetServiceExtensions returns the ServiceExtensions field value if set, zero value otherwise.
func (o *Product) GetServiceExtensions() []NSOConfigDataXPath {
	if o == nil || o.ServiceExtensions == nil {
		var ret []NSOConfigDataXPath
		return ret
	}
	return *o.ServiceExtensions
}

// GetServiceExtensionsOk returns a tuple with the ServiceExtensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetServiceExtensionsOk() (*[]NSOConfigDataXPath, bool) {
	if o == nil || o.ServiceExtensions == nil {
		return nil, false
	}
	return o.ServiceExtensions, true
}

// HasServiceExtensions returns a boolean if a field has been set.
func (o *Product) HasServiceExtensions() bool {
	if o != nil && o.ServiceExtensions != nil {
		return true
	}

	return false
}

// SetServiceExtensions gets a reference to the given []NSOConfigDataXPath and assigns it to the ServiceExtensions field.
func (o *Product) SetServiceExtensions(v []NSOConfigDataXPath) {
	o.ServiceExtensions = &v
}

// GetServiceConfigQueryRootXPaths returns the ServiceConfigQueryRootXPaths field value if set, zero value otherwise.
func (o *Product) GetServiceConfigQueryRootXPaths() []string {
	if o == nil || o.ServiceConfigQueryRootXPaths == nil {
		var ret []string
		return ret
	}
	return *o.ServiceConfigQueryRootXPaths
}

// GetServiceConfigQueryRootXPathsOk returns a tuple with the ServiceConfigQueryRootXPaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetServiceConfigQueryRootXPathsOk() (*[]string, bool) {
	if o == nil || o.ServiceConfigQueryRootXPaths == nil {
		return nil, false
	}
	return o.ServiceConfigQueryRootXPaths, true
}

// HasServiceConfigQueryRootXPaths returns a boolean if a field has been set.
func (o *Product) HasServiceConfigQueryRootXPaths() bool {
	if o != nil && o.ServiceConfigQueryRootXPaths != nil {
		return true
	}

	return false
}

// SetServiceConfigQueryRootXPaths gets a reference to the given []string and assigns it to the ServiceConfigQueryRootXPaths field.
func (o *Product) SetServiceConfigQueryRootXPaths(v []string) {
	o.ServiceConfigQueryRootXPaths = &v
}

// GetUiConfig returns the UiConfig field value if set, zero value otherwise.
func (o *Product) GetUiConfig() ServiceUIConfig {
	if o == nil || o.UiConfig == nil {
		var ret ServiceUIConfig
		return ret
	}
	return *o.UiConfig
}

// GetUiConfigOk returns a tuple with the UiConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetUiConfigOk() (*ServiceUIConfig, bool) {
	if o == nil || o.UiConfig == nil {
		return nil, false
	}
	return o.UiConfig, true
}

// HasUiConfig returns a boolean if a field has been set.
func (o *Product) HasUiConfig() bool {
	if o != nil && o.UiConfig != nil {
		return true
	}

	return false
}

// SetUiConfig gets a reference to the given ServiceUIConfig and assigns it to the UiConfig field.
func (o *Product) SetUiConfig(v ServiceUIConfig) {
	o.UiConfig = &v
}

func (o Product) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
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
	if true {
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
	if true {
		toSerialize["options"] = o.Options
	}
	if true {
		toSerialize["properties"] = o.Properties
	}
	if true {
		toSerialize["configuration"] = o.Configuration
	}
	if true {
		toSerialize["isResource"] = o.IsResource
	}
	if true {
		toSerialize["hasChildren"] = o.HasChildren
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
	return json.Marshal(toSerialize)
}

type NullableProduct struct {
	value *Product
	isSet bool
}

func (v NullableProduct) Get() *Product {
	return v.value
}

func (v *NullableProduct) Set(val *Product) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct(val *Product) *NullableProduct {
	return &NullableProduct{value: val, isSet: true}
}

func (v NullableProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


