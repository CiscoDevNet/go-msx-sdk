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

// BillingPriceCreate struct for BillingPriceCreate
type BillingPriceCreate struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Type string `json:"type"`
	Subtype *string `json:"subtype,omitempty"`
	Source *string `json:"source,omitempty"`
	Price float64 `json:"price"`
	BillingPeriod *int64 `json:"billingPeriod,omitempty"`
	Service NullableString `json:"service,omitempty"`
	TenantId string `json:"tenantId"`
	AdditionalProperties map[string]interface{}
}

type _BillingPriceCreate BillingPriceCreate

// NewBillingPriceCreate instantiates a new BillingPriceCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingPriceCreate(name string, type_ string, price float64, tenantId string) *BillingPriceCreate {
	this := BillingPriceCreate{}
	this.Name = name
	this.Type = type_
	this.Price = price
	this.TenantId = tenantId
	return &this
}

// NewBillingPriceCreateWithDefaults instantiates a new BillingPriceCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingPriceCreateWithDefaults() *BillingPriceCreate {
	this := BillingPriceCreate{}
	return &this
}

// GetName returns the Name field value
func (o *BillingPriceCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BillingPriceCreate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BillingPriceCreate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BillingPriceCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPriceCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BillingPriceCreate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BillingPriceCreate) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value
func (o *BillingPriceCreate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BillingPriceCreate) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BillingPriceCreate) SetType(v string) {
	o.Type = v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *BillingPriceCreate) GetSubtype() string {
	if o == nil || o.Subtype == nil {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPriceCreate) GetSubtypeOk() (*string, bool) {
	if o == nil || o.Subtype == nil {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *BillingPriceCreate) HasSubtype() bool {
	if o != nil && o.Subtype != nil {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *BillingPriceCreate) SetSubtype(v string) {
	o.Subtype = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *BillingPriceCreate) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPriceCreate) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *BillingPriceCreate) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *BillingPriceCreate) SetSource(v string) {
	o.Source = &v
}

// GetPrice returns the Price field value
func (o *BillingPriceCreate) GetPrice() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *BillingPriceCreate) GetPriceOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *BillingPriceCreate) SetPrice(v float64) {
	o.Price = v
}

// GetBillingPeriod returns the BillingPeriod field value if set, zero value otherwise.
func (o *BillingPriceCreate) GetBillingPeriod() int64 {
	if o == nil || o.BillingPeriod == nil {
		var ret int64
		return ret
	}
	return *o.BillingPeriod
}

// GetBillingPeriodOk returns a tuple with the BillingPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPriceCreate) GetBillingPeriodOk() (*int64, bool) {
	if o == nil || o.BillingPeriod == nil {
		return nil, false
	}
	return o.BillingPeriod, true
}

// HasBillingPeriod returns a boolean if a field has been set.
func (o *BillingPriceCreate) HasBillingPeriod() bool {
	if o != nil && o.BillingPeriod != nil {
		return true
	}

	return false
}

// SetBillingPeriod gets a reference to the given int64 and assigns it to the BillingPeriod field.
func (o *BillingPriceCreate) SetBillingPeriod(v int64) {
	o.BillingPeriod = &v
}

// GetService returns the Service field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingPriceCreate) GetService() string {
	if o == nil || o.Service.Get() == nil {
		var ret string
		return ret
	}
	return *o.Service.Get()
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingPriceCreate) GetServiceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Service.Get(), o.Service.IsSet()
}

// HasService returns a boolean if a field has been set.
func (o *BillingPriceCreate) HasService() bool {
	if o != nil && o.Service.IsSet() {
		return true
	}

	return false
}

// SetService gets a reference to the given NullableString and assigns it to the Service field.
func (o *BillingPriceCreate) SetService(v string) {
	o.Service.Set(&v)
}
// SetServiceNil sets the value for Service to be an explicit nil
func (o *BillingPriceCreate) SetServiceNil() {
	o.Service.Set(nil)
}

// UnsetService ensures that no value is present for Service, not even an explicit nil
func (o *BillingPriceCreate) UnsetService() {
	o.Service.Unset()
}

// GetTenantId returns the TenantId field value
func (o *BillingPriceCreate) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *BillingPriceCreate) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *BillingPriceCreate) SetTenantId(v string) {
	o.TenantId = v
}

func (o BillingPriceCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Subtype != nil {
		toSerialize["subtype"] = o.Subtype
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["price"] = o.Price
	}
	if o.BillingPeriod != nil {
		toSerialize["billingPeriod"] = o.BillingPeriod
	}
	if o.Service.IsSet() {
		toSerialize["service"] = o.Service.Get()
	}
	if true {
		toSerialize["tenantId"] = o.TenantId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BillingPriceCreate) UnmarshalJSON(bytes []byte) (err error) {
	varBillingPriceCreate := _BillingPriceCreate{}

	if err = json.Unmarshal(bytes, &varBillingPriceCreate); err == nil {
		*o = BillingPriceCreate(varBillingPriceCreate)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "type")
		delete(additionalProperties, "subtype")
		delete(additionalProperties, "source")
		delete(additionalProperties, "price")
		delete(additionalProperties, "billingPeriod")
		delete(additionalProperties, "service")
		delete(additionalProperties, "tenantId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBillingPriceCreate struct {
	value *BillingPriceCreate
	isSet bool
}

func (v NullableBillingPriceCreate) Get() *BillingPriceCreate {
	return v.value
}

func (v *NullableBillingPriceCreate) Set(val *BillingPriceCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingPriceCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingPriceCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingPriceCreate(val *BillingPriceCreate) *NullableBillingPriceCreate {
	return &NullableBillingPriceCreate{value: val, isSet: true}
}

func (v NullableBillingPriceCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingPriceCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


