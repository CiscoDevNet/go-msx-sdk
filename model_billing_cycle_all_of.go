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
)

// BillingCycleAllOf struct for BillingCycleAllOf
type BillingCycleAllOf struct {
	Id *string `json:"id,omitempty"`
	EventId *string `json:"eventId,omitempty"`
	LastBilledOn *string `json:"lastBilledOn,omitempty"`
	NextBilledOn *string `json:"nextBilledOn,omitempty"`
	TenantId *string `json:"tenantId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BillingCycleAllOf BillingCycleAllOf

// NewBillingCycleAllOf instantiates a new BillingCycleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingCycleAllOf() *BillingCycleAllOf {
	this := BillingCycleAllOf{}
	return &this
}

// NewBillingCycleAllOfWithDefaults instantiates a new BillingCycleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingCycleAllOfWithDefaults() *BillingCycleAllOf {
	this := BillingCycleAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BillingCycleAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingCycleAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BillingCycleAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BillingCycleAllOf) SetId(v string) {
	o.Id = &v
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *BillingCycleAllOf) GetEventId() string {
	if o == nil || o.EventId == nil {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingCycleAllOf) GetEventIdOk() (*string, bool) {
	if o == nil || o.EventId == nil {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *BillingCycleAllOf) HasEventId() bool {
	if o != nil && o.EventId != nil {
		return true
	}

	return false
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *BillingCycleAllOf) SetEventId(v string) {
	o.EventId = &v
}

// GetLastBilledOn returns the LastBilledOn field value if set, zero value otherwise.
func (o *BillingCycleAllOf) GetLastBilledOn() string {
	if o == nil || o.LastBilledOn == nil {
		var ret string
		return ret
	}
	return *o.LastBilledOn
}

// GetLastBilledOnOk returns a tuple with the LastBilledOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingCycleAllOf) GetLastBilledOnOk() (*string, bool) {
	if o == nil || o.LastBilledOn == nil {
		return nil, false
	}
	return o.LastBilledOn, true
}

// HasLastBilledOn returns a boolean if a field has been set.
func (o *BillingCycleAllOf) HasLastBilledOn() bool {
	if o != nil && o.LastBilledOn != nil {
		return true
	}

	return false
}

// SetLastBilledOn gets a reference to the given string and assigns it to the LastBilledOn field.
func (o *BillingCycleAllOf) SetLastBilledOn(v string) {
	o.LastBilledOn = &v
}

// GetNextBilledOn returns the NextBilledOn field value if set, zero value otherwise.
func (o *BillingCycleAllOf) GetNextBilledOn() string {
	if o == nil || o.NextBilledOn == nil {
		var ret string
		return ret
	}
	return *o.NextBilledOn
}

// GetNextBilledOnOk returns a tuple with the NextBilledOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingCycleAllOf) GetNextBilledOnOk() (*string, bool) {
	if o == nil || o.NextBilledOn == nil {
		return nil, false
	}
	return o.NextBilledOn, true
}

// HasNextBilledOn returns a boolean if a field has been set.
func (o *BillingCycleAllOf) HasNextBilledOn() bool {
	if o != nil && o.NextBilledOn != nil {
		return true
	}

	return false
}

// SetNextBilledOn gets a reference to the given string and assigns it to the NextBilledOn field.
func (o *BillingCycleAllOf) SetNextBilledOn(v string) {
	o.NextBilledOn = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *BillingCycleAllOf) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingCycleAllOf) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *BillingCycleAllOf) HasTenantId() bool {
	if o != nil && o.TenantId != nil {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *BillingCycleAllOf) SetTenantId(v string) {
	o.TenantId = &v
}

func (o BillingCycleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.EventId != nil {
		toSerialize["eventId"] = o.EventId
	}
	if o.LastBilledOn != nil {
		toSerialize["lastBilledOn"] = o.LastBilledOn
	}
	if o.NextBilledOn != nil {
		toSerialize["nextBilledOn"] = o.NextBilledOn
	}
	if o.TenantId != nil {
		toSerialize["tenantId"] = o.TenantId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BillingCycleAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBillingCycleAllOf := _BillingCycleAllOf{}

	if err = json.Unmarshal(bytes, &varBillingCycleAllOf); err == nil {
		*o = BillingCycleAllOf(varBillingCycleAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "eventId")
		delete(additionalProperties, "lastBilledOn")
		delete(additionalProperties, "nextBilledOn")
		delete(additionalProperties, "tenantId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBillingCycleAllOf struct {
	value *BillingCycleAllOf
	isSet bool
}

func (v NullableBillingCycleAllOf) Get() *BillingCycleAllOf {
	return v.value
}

func (v *NullableBillingCycleAllOf) Set(val *BillingCycleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingCycleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingCycleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingCycleAllOf(val *BillingCycleAllOf) *NullableBillingCycleAllOf {
	return &NullableBillingCycleAllOf{value: val, isSet: true}
}

func (v NullableBillingCycleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingCycleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


