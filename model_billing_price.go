/*
MSX SDK

MSX SDK client.

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msxsdk

import (
	"encoding/json"
	"time"
)

// BillingPrice struct for BillingPrice
type BillingPrice struct {
	Id *string `json:"id,omitempty"`
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	ModifiedOn *time.Time `json:"modifiedOn,omitempty"`
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

type _BillingPrice BillingPrice

// NewBillingPrice instantiates a new BillingPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingPrice(name string, type_ string, price float64, tenantId string) *BillingPrice {
	this := BillingPrice{}
	this.Name = name
	this.Type = type_
	this.Price = price
	this.TenantId = tenantId
	return &this
}

// NewBillingPriceWithDefaults instantiates a new BillingPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingPriceWithDefaults() *BillingPrice {
	this := BillingPrice{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BillingPrice) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPrice) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BillingPrice) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BillingPrice) SetId(v string) {
	o.Id = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *BillingPrice) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPrice) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *BillingPrice) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *BillingPrice) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetModifiedOn returns the ModifiedOn field value if set, zero value otherwise.
func (o *BillingPrice) GetModifiedOn() time.Time {
	if o == nil || o.ModifiedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPrice) GetModifiedOnOk() (*time.Time, bool) {
	if o == nil || o.ModifiedOn == nil {
		return nil, false
	}
	return o.ModifiedOn, true
}

// HasModifiedOn returns a boolean if a field has been set.
func (o *BillingPrice) HasModifiedOn() bool {
	if o != nil && o.ModifiedOn != nil {
		return true
	}

	return false
}

// SetModifiedOn gets a reference to the given time.Time and assigns it to the ModifiedOn field.
func (o *BillingPrice) SetModifiedOn(v time.Time) {
	o.ModifiedOn = &v
}

// GetName returns the Name field value
func (o *BillingPrice) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BillingPrice) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BillingPrice) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BillingPrice) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPrice) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BillingPrice) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BillingPrice) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value
func (o *BillingPrice) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BillingPrice) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BillingPrice) SetType(v string) {
	o.Type = v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *BillingPrice) GetSubtype() string {
	if o == nil || o.Subtype == nil {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPrice) GetSubtypeOk() (*string, bool) {
	if o == nil || o.Subtype == nil {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *BillingPrice) HasSubtype() bool {
	if o != nil && o.Subtype != nil {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *BillingPrice) SetSubtype(v string) {
	o.Subtype = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *BillingPrice) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPrice) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *BillingPrice) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *BillingPrice) SetSource(v string) {
	o.Source = &v
}

// GetPrice returns the Price field value
func (o *BillingPrice) GetPrice() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *BillingPrice) GetPriceOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *BillingPrice) SetPrice(v float64) {
	o.Price = v
}

// GetBillingPeriod returns the BillingPeriod field value if set, zero value otherwise.
func (o *BillingPrice) GetBillingPeriod() int64 {
	if o == nil || o.BillingPeriod == nil {
		var ret int64
		return ret
	}
	return *o.BillingPeriod
}

// GetBillingPeriodOk returns a tuple with the BillingPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPrice) GetBillingPeriodOk() (*int64, bool) {
	if o == nil || o.BillingPeriod == nil {
		return nil, false
	}
	return o.BillingPeriod, true
}

// HasBillingPeriod returns a boolean if a field has been set.
func (o *BillingPrice) HasBillingPeriod() bool {
	if o != nil && o.BillingPeriod != nil {
		return true
	}

	return false
}

// SetBillingPeriod gets a reference to the given int64 and assigns it to the BillingPeriod field.
func (o *BillingPrice) SetBillingPeriod(v int64) {
	o.BillingPeriod = &v
}

// GetService returns the Service field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingPrice) GetService() string {
	if o == nil || o.Service.Get() == nil {
		var ret string
		return ret
	}
	return *o.Service.Get()
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingPrice) GetServiceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Service.Get(), o.Service.IsSet()
}

// HasService returns a boolean if a field has been set.
func (o *BillingPrice) HasService() bool {
	if o != nil && o.Service.IsSet() {
		return true
	}

	return false
}

// SetService gets a reference to the given NullableString and assigns it to the Service field.
func (o *BillingPrice) SetService(v string) {
	o.Service.Set(&v)
}
// SetServiceNil sets the value for Service to be an explicit nil
func (o *BillingPrice) SetServiceNil() {
	o.Service.Set(nil)
}

// UnsetService ensures that no value is present for Service, not even an explicit nil
func (o *BillingPrice) UnsetService() {
	o.Service.Unset()
}

// GetTenantId returns the TenantId field value
func (o *BillingPrice) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *BillingPrice) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *BillingPrice) SetTenantId(v string) {
	o.TenantId = v
}

func (o BillingPrice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedOn != nil {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if o.ModifiedOn != nil {
		toSerialize["modifiedOn"] = o.ModifiedOn
	}
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

func (o *BillingPrice) UnmarshalJSON(bytes []byte) (err error) {
	varBillingPrice := _BillingPrice{}

	if err = json.Unmarshal(bytes, &varBillingPrice); err == nil {
		*o = BillingPrice(varBillingPrice)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "createdOn")
		delete(additionalProperties, "modifiedOn")
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

type NullableBillingPrice struct {
	value *BillingPrice
	isSet bool
}

func (v NullableBillingPrice) Get() *BillingPrice {
	return v.value
}

func (v *NullableBillingPrice) Set(val *BillingPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingPrice(val *BillingPrice) *NullableBillingPrice {
	return &NullableBillingPrice{value: val, isSet: true}
}

func (v NullableBillingPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


