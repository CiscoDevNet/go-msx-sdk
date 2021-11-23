/*
 * MSX SDK
 *
 * MSX SDK client.
 *
 * API version: 1.0.8
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msxsdk

import (
	"encoding/json"
	"time"
)

// BillingEventCreate struct for BillingEventCreate
type BillingEventCreate struct {
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	Subtype *string `json:"subtype,omitempty"`
	Severity *string `json:"severity,omitempty"`
	Action *string `json:"action,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
	Price *float64 `json:"price,omitempty"`
	TenantId *string `json:"tenantId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BillingEventCreate BillingEventCreate

// NewBillingEventCreate instantiates a new BillingEventCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingEventCreate() *BillingEventCreate {
	this := BillingEventCreate{}
	return &this
}

// NewBillingEventCreateWithDefaults instantiates a new BillingEventCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingEventCreateWithDefaults() *BillingEventCreate {
	this := BillingEventCreate{}
	return &this
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *BillingEventCreate) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEventCreate) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *BillingEventCreate) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *BillingEventCreate) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BillingEventCreate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEventCreate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BillingEventCreate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BillingEventCreate) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BillingEventCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEventCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BillingEventCreate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BillingEventCreate) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BillingEventCreate) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEventCreate) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BillingEventCreate) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BillingEventCreate) SetType(v string) {
	o.Type = &v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *BillingEventCreate) GetSubtype() string {
	if o == nil || o.Subtype == nil {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEventCreate) GetSubtypeOk() (*string, bool) {
	if o == nil || o.Subtype == nil {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *BillingEventCreate) HasSubtype() bool {
	if o != nil && o.Subtype != nil {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *BillingEventCreate) SetSubtype(v string) {
	o.Subtype = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *BillingEventCreate) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEventCreate) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *BillingEventCreate) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *BillingEventCreate) SetSeverity(v string) {
	o.Severity = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *BillingEventCreate) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEventCreate) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *BillingEventCreate) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *BillingEventCreate) SetAction(v string) {
	o.Action = &v
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingEventCreate) GetDetails() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingEventCreate) GetDetailsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return &o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *BillingEventCreate) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given map[string]interface{} and assigns it to the Details field.
func (o *BillingEventCreate) SetDetails(v map[string]interface{}) {
	o.Details = v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *BillingEventCreate) GetPrice() float64 {
	if o == nil || o.Price == nil {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEventCreate) GetPriceOk() (*float64, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *BillingEventCreate) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *BillingEventCreate) SetPrice(v float64) {
	o.Price = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *BillingEventCreate) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEventCreate) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *BillingEventCreate) HasTenantId() bool {
	if o != nil && o.TenantId != nil {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *BillingEventCreate) SetTenantId(v string) {
	o.TenantId = &v
}

func (o BillingEventCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedOn != nil {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Subtype != nil {
		toSerialize["subtype"] = o.Subtype
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.TenantId != nil {
		toSerialize["tenantId"] = o.TenantId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BillingEventCreate) UnmarshalJSON(bytes []byte) (err error) {
	varBillingEventCreate := _BillingEventCreate{}

	if err = json.Unmarshal(bytes, &varBillingEventCreate); err == nil {
		*o = BillingEventCreate(varBillingEventCreate)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "createdOn")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "type")
		delete(additionalProperties, "subtype")
		delete(additionalProperties, "severity")
		delete(additionalProperties, "action")
		delete(additionalProperties, "details")
		delete(additionalProperties, "price")
		delete(additionalProperties, "tenantId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBillingEventCreate struct {
	value *BillingEventCreate
	isSet bool
}

func (v NullableBillingEventCreate) Get() *BillingEventCreate {
	return v.value
}

func (v *NullableBillingEventCreate) Set(val *BillingEventCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingEventCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingEventCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingEventCreate(val *BillingEventCreate) *NullableBillingEventCreate {
	return &NullableBillingEventCreate{value: val, isSet: true}
}

func (v NullableBillingEventCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingEventCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


