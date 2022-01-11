/*
 * MSX SDK
 *
 * MSX SDK client.
 *
 * API version: 1.0.9
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msxsdk

import (
	"encoding/json"
)

// BillingCycleProcess struct for BillingCycleProcess
type BillingCycleProcess struct {
	NextBilledOn string `json:"nextBilledOn"`
	AdditionalProperties map[string]interface{}
}

type _BillingCycleProcess BillingCycleProcess

// NewBillingCycleProcess instantiates a new BillingCycleProcess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingCycleProcess(nextBilledOn string) *BillingCycleProcess {
	this := BillingCycleProcess{}
	this.NextBilledOn = nextBilledOn
	return &this
}

// NewBillingCycleProcessWithDefaults instantiates a new BillingCycleProcess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingCycleProcessWithDefaults() *BillingCycleProcess {
	this := BillingCycleProcess{}
	return &this
}

// GetNextBilledOn returns the NextBilledOn field value
func (o *BillingCycleProcess) GetNextBilledOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextBilledOn
}

// GetNextBilledOnOk returns a tuple with the NextBilledOn field value
// and a boolean to check if the value has been set.
func (o *BillingCycleProcess) GetNextBilledOnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NextBilledOn, true
}

// SetNextBilledOn sets field value
func (o *BillingCycleProcess) SetNextBilledOn(v string) {
	o.NextBilledOn = v
}

func (o BillingCycleProcess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["nextBilledOn"] = o.NextBilledOn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BillingCycleProcess) UnmarshalJSON(bytes []byte) (err error) {
	varBillingCycleProcess := _BillingCycleProcess{}

	if err = json.Unmarshal(bytes, &varBillingCycleProcess); err == nil {
		*o = BillingCycleProcess(varBillingCycleProcess)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "nextBilledOn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBillingCycleProcess struct {
	value *BillingCycleProcess
	isSet bool
}

func (v NullableBillingCycleProcess) Get() *BillingCycleProcess {
	return v.value
}

func (v *NullableBillingCycleProcess) Set(val *BillingCycleProcess) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingCycleProcess) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingCycleProcess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingCycleProcess(val *BillingCycleProcess) *NullableBillingCycleProcess {
	return &NullableBillingCycleProcess{value: val, isSet: true}
}

func (v NullableBillingCycleProcess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingCycleProcess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


