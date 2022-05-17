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

// OfferUpdate struct for OfferUpdate
type OfferUpdate struct {
	Name string `json:"name"`
	Label string `json:"label"`
	Description string `json:"description"`
	ProductId string `json:"productId"`
	Version int32 `json:"version"`
	DisplayOrder NullableInt32 `json:"displayOrder,omitempty"`
	Image *string `json:"image,omitempty"`
	Price *string `json:"price,omitempty"`
	Type *string `json:"type,omitempty"`
	SupportedProperties []string `json:"supportedProperties,omitempty"`
	SupportedOptions []NameValue `json:"supportedOptions,omitempty"`
	Approvals map[string]interface{} `json:"approvals,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OfferUpdate OfferUpdate

// NewOfferUpdate instantiates a new OfferUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferUpdate(name string, label string, description string, productId string, version int32) *OfferUpdate {
	this := OfferUpdate{}
	this.Name = name
	this.Label = label
	this.Description = description
	this.ProductId = productId
	this.Version = version
	return &this
}

// NewOfferUpdateWithDefaults instantiates a new OfferUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferUpdateWithDefaults() *OfferUpdate {
	this := OfferUpdate{}
	return &this
}

// GetName returns the Name field value
func (o *OfferUpdate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OfferUpdate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OfferUpdate) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value
func (o *OfferUpdate) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *OfferUpdate) GetLabelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *OfferUpdate) SetLabel(v string) {
	o.Label = v
}

// GetDescription returns the Description field value
func (o *OfferUpdate) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *OfferUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *OfferUpdate) SetDescription(v string) {
	o.Description = v
}

// GetProductId returns the ProductId field value
func (o *OfferUpdate) GetProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *OfferUpdate) GetProductIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *OfferUpdate) SetProductId(v string) {
	o.ProductId = v
}

// GetVersion returns the Version field value
func (o *OfferUpdate) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *OfferUpdate) GetVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *OfferUpdate) SetVersion(v int32) {
	o.Version = v
}

// GetDisplayOrder returns the DisplayOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OfferUpdate) GetDisplayOrder() int32 {
	if o == nil || o.DisplayOrder.Get() == nil {
		var ret int32
		return ret
	}
	return *o.DisplayOrder.Get()
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OfferUpdate) GetDisplayOrderOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayOrder.Get(), o.DisplayOrder.IsSet()
}

// HasDisplayOrder returns a boolean if a field has been set.
func (o *OfferUpdate) HasDisplayOrder() bool {
	if o != nil && o.DisplayOrder.IsSet() {
		return true
	}

	return false
}

// SetDisplayOrder gets a reference to the given NullableInt32 and assigns it to the DisplayOrder field.
func (o *OfferUpdate) SetDisplayOrder(v int32) {
	o.DisplayOrder.Set(&v)
}
// SetDisplayOrderNil sets the value for DisplayOrder to be an explicit nil
func (o *OfferUpdate) SetDisplayOrderNil() {
	o.DisplayOrder.Set(nil)
}

// UnsetDisplayOrder ensures that no value is present for DisplayOrder, not even an explicit nil
func (o *OfferUpdate) UnsetDisplayOrder() {
	o.DisplayOrder.Unset()
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *OfferUpdate) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferUpdate) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *OfferUpdate) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *OfferUpdate) SetImage(v string) {
	o.Image = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *OfferUpdate) GetPrice() string {
	if o == nil || o.Price == nil {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferUpdate) GetPriceOk() (*string, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *OfferUpdate) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *OfferUpdate) SetPrice(v string) {
	o.Price = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OfferUpdate) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferUpdate) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OfferUpdate) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OfferUpdate) SetType(v string) {
	o.Type = &v
}

// GetSupportedProperties returns the SupportedProperties field value if set, zero value otherwise.
func (o *OfferUpdate) GetSupportedProperties() []string {
	if o == nil || o.SupportedProperties == nil {
		var ret []string
		return ret
	}
	return o.SupportedProperties
}

// GetSupportedPropertiesOk returns a tuple with the SupportedProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferUpdate) GetSupportedPropertiesOk() ([]string, bool) {
	if o == nil || o.SupportedProperties == nil {
		return nil, false
	}
	return o.SupportedProperties, true
}

// HasSupportedProperties returns a boolean if a field has been set.
func (o *OfferUpdate) HasSupportedProperties() bool {
	if o != nil && o.SupportedProperties != nil {
		return true
	}

	return false
}

// SetSupportedProperties gets a reference to the given []string and assigns it to the SupportedProperties field.
func (o *OfferUpdate) SetSupportedProperties(v []string) {
	o.SupportedProperties = v
}

// GetSupportedOptions returns the SupportedOptions field value if set, zero value otherwise.
func (o *OfferUpdate) GetSupportedOptions() []NameValue {
	if o == nil || o.SupportedOptions == nil {
		var ret []NameValue
		return ret
	}
	return o.SupportedOptions
}

// GetSupportedOptionsOk returns a tuple with the SupportedOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferUpdate) GetSupportedOptionsOk() ([]NameValue, bool) {
	if o == nil || o.SupportedOptions == nil {
		return nil, false
	}
	return o.SupportedOptions, true
}

// HasSupportedOptions returns a boolean if a field has been set.
func (o *OfferUpdate) HasSupportedOptions() bool {
	if o != nil && o.SupportedOptions != nil {
		return true
	}

	return false
}

// SetSupportedOptions gets a reference to the given []NameValue and assigns it to the SupportedOptions field.
func (o *OfferUpdate) SetSupportedOptions(v []NameValue) {
	o.SupportedOptions = v
}

// GetApprovals returns the Approvals field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OfferUpdate) GetApprovals() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.Approvals
}

// GetApprovalsOk returns a tuple with the Approvals field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OfferUpdate) GetApprovalsOk() (map[string]interface{}, bool) {
	if o == nil || o.Approvals == nil {
		return nil, false
	}
	return o.Approvals, true
}

// HasApprovals returns a boolean if a field has been set.
func (o *OfferUpdate) HasApprovals() bool {
	if o != nil && o.Approvals != nil {
		return true
	}

	return false
}

// SetApprovals gets a reference to the given map[string]interface{} and assigns it to the Approvals field.
func (o *OfferUpdate) SetApprovals(v map[string]interface{}) {
	o.Approvals = v
}

func (o OfferUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["productId"] = o.ProductId
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if o.DisplayOrder.IsSet() {
		toSerialize["displayOrder"] = o.DisplayOrder.Get()
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.SupportedProperties != nil {
		toSerialize["supportedProperties"] = o.SupportedProperties
	}
	if o.SupportedOptions != nil {
		toSerialize["supportedOptions"] = o.SupportedOptions
	}
	if o.Approvals != nil {
		toSerialize["approvals"] = o.Approvals
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OfferUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varOfferUpdate := _OfferUpdate{}

	if err = json.Unmarshal(bytes, &varOfferUpdate); err == nil {
		*o = OfferUpdate(varOfferUpdate)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		delete(additionalProperties, "productId")
		delete(additionalProperties, "version")
		delete(additionalProperties, "displayOrder")
		delete(additionalProperties, "image")
		delete(additionalProperties, "price")
		delete(additionalProperties, "type")
		delete(additionalProperties, "supportedProperties")
		delete(additionalProperties, "supportedOptions")
		delete(additionalProperties, "approvals")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOfferUpdate struct {
	value *OfferUpdate
	isSet bool
}

func (v NullableOfferUpdate) Get() *OfferUpdate {
	return v.value
}

func (v *NullableOfferUpdate) Set(val *OfferUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferUpdate(val *OfferUpdate) *NullableOfferUpdate {
	return &NullableOfferUpdate{value: val, isSet: true}
}

func (v NullableOfferUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfferUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


