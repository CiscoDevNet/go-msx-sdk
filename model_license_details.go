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

// LicenseDetails struct for LicenseDetails
type LicenseDetails struct {
	// Type of license - TERM or DEMO or PERPETUAL
	LicenseType *string `json:"licenseType,omitempty"`
	// Number of reserved licenses
	Quantity *int64 `json:"quantity,omitempty"`
	// License usage start date in yyyy-mm-dd format
	StartDate *string `json:"startDate,omitempty"`
	// License usage expiration date in yyyy-mm-dd format
	EndDate *string `json:"endDate,omitempty"`
	// Subscription refence id
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	// Licencse usage status
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LicenseDetails LicenseDetails

// NewLicenseDetails instantiates a new LicenseDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseDetails() *LicenseDetails {
	this := LicenseDetails{}
	return &this
}

// NewLicenseDetailsWithDefaults instantiates a new LicenseDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseDetailsWithDefaults() *LicenseDetails {
	this := LicenseDetails{}
	return &this
}

// GetLicenseType returns the LicenseType field value if set, zero value otherwise.
func (o *LicenseDetails) GetLicenseType() string {
	if o == nil || o.LicenseType == nil {
		var ret string
		return ret
	}
	return *o.LicenseType
}

// GetLicenseTypeOk returns a tuple with the LicenseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetails) GetLicenseTypeOk() (*string, bool) {
	if o == nil || o.LicenseType == nil {
		return nil, false
	}
	return o.LicenseType, true
}

// HasLicenseType returns a boolean if a field has been set.
func (o *LicenseDetails) HasLicenseType() bool {
	if o != nil && o.LicenseType != nil {
		return true
	}

	return false
}

// SetLicenseType gets a reference to the given string and assigns it to the LicenseType field.
func (o *LicenseDetails) SetLicenseType(v string) {
	o.LicenseType = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *LicenseDetails) GetQuantity() int64 {
	if o == nil || o.Quantity == nil {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetails) GetQuantityOk() (*int64, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *LicenseDetails) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *LicenseDetails) SetQuantity(v int64) {
	o.Quantity = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *LicenseDetails) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetails) GetStartDateOk() (*string, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *LicenseDetails) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *LicenseDetails) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *LicenseDetails) GetEndDate() string {
	if o == nil || o.EndDate == nil {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetails) GetEndDateOk() (*string, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *LicenseDetails) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *LicenseDetails) SetEndDate(v string) {
	o.EndDate = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *LicenseDetails) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetails) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || o.SubscriptionId == nil {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *LicenseDetails) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId != nil {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *LicenseDetails) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LicenseDetails) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseDetails) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LicenseDetails) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LicenseDetails) SetStatus(v string) {
	o.Status = &v
}

func (o LicenseDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LicenseType != nil {
		toSerialize["licenseType"] = o.LicenseType
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.StartDate != nil {
		toSerialize["startDate"] = o.StartDate
	}
	if o.EndDate != nil {
		toSerialize["endDate"] = o.EndDate
	}
	if o.SubscriptionId != nil {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LicenseDetails) UnmarshalJSON(bytes []byte) (err error) {
	varLicenseDetails := _LicenseDetails{}

	if err = json.Unmarshal(bytes, &varLicenseDetails); err == nil {
		*o = LicenseDetails(varLicenseDetails)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "licenseType")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "startDate")
		delete(additionalProperties, "endDate")
		delete(additionalProperties, "subscriptionId")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLicenseDetails struct {
	value *LicenseDetails
	isSet bool
}

func (v NullableLicenseDetails) Get() *LicenseDetails {
	return v.value
}

func (v *NullableLicenseDetails) Set(val *LicenseDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseDetails(val *LicenseDetails) *NullableLicenseDetails {
	return &NullableLicenseDetails{value: val, isSet: true}
}

func (v NullableLicenseDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


