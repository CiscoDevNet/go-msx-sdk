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

// LicenseSummary struct for LicenseSummary
type LicenseSummary struct {
	// Total entitled quantity for the license
	Entitled *int64 `json:"entitled,omitempty"`
	// Total consumed quantity for the device
	Inuse *int64 `json:"inuse,omitempty"`
	// Reserved quantity (if any)
	Reserved *int64 `json:"reserved,omitempty"`
	// Current compliance status for the license
	Status *string `json:"status,omitempty"`
	// User friendly display name for the license
	Name *string `json:"name,omitempty"`
	// Identifies if the tag is for an enforced license or not
	Enforced *bool `json:"enforced,omitempty"`
	LicenseDetails *[]LicenseDetails `json:"licenseDetails,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LicenseSummary LicenseSummary

// NewLicenseSummary instantiates a new LicenseSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseSummary() *LicenseSummary {
	this := LicenseSummary{}
	return &this
}

// NewLicenseSummaryWithDefaults instantiates a new LicenseSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseSummaryWithDefaults() *LicenseSummary {
	this := LicenseSummary{}
	return &this
}

// GetEntitled returns the Entitled field value if set, zero value otherwise.
func (o *LicenseSummary) GetEntitled() int64 {
	if o == nil || o.Entitled == nil {
		var ret int64
		return ret
	}
	return *o.Entitled
}

// GetEntitledOk returns a tuple with the Entitled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseSummary) GetEntitledOk() (*int64, bool) {
	if o == nil || o.Entitled == nil {
		return nil, false
	}
	return o.Entitled, true
}

// HasEntitled returns a boolean if a field has been set.
func (o *LicenseSummary) HasEntitled() bool {
	if o != nil && o.Entitled != nil {
		return true
	}

	return false
}

// SetEntitled gets a reference to the given int64 and assigns it to the Entitled field.
func (o *LicenseSummary) SetEntitled(v int64) {
	o.Entitled = &v
}

// GetInuse returns the Inuse field value if set, zero value otherwise.
func (o *LicenseSummary) GetInuse() int64 {
	if o == nil || o.Inuse == nil {
		var ret int64
		return ret
	}
	return *o.Inuse
}

// GetInuseOk returns a tuple with the Inuse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseSummary) GetInuseOk() (*int64, bool) {
	if o == nil || o.Inuse == nil {
		return nil, false
	}
	return o.Inuse, true
}

// HasInuse returns a boolean if a field has been set.
func (o *LicenseSummary) HasInuse() bool {
	if o != nil && o.Inuse != nil {
		return true
	}

	return false
}

// SetInuse gets a reference to the given int64 and assigns it to the Inuse field.
func (o *LicenseSummary) SetInuse(v int64) {
	o.Inuse = &v
}

// GetReserved returns the Reserved field value if set, zero value otherwise.
func (o *LicenseSummary) GetReserved() int64 {
	if o == nil || o.Reserved == nil {
		var ret int64
		return ret
	}
	return *o.Reserved
}

// GetReservedOk returns a tuple with the Reserved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseSummary) GetReservedOk() (*int64, bool) {
	if o == nil || o.Reserved == nil {
		return nil, false
	}
	return o.Reserved, true
}

// HasReserved returns a boolean if a field has been set.
func (o *LicenseSummary) HasReserved() bool {
	if o != nil && o.Reserved != nil {
		return true
	}

	return false
}

// SetReserved gets a reference to the given int64 and assigns it to the Reserved field.
func (o *LicenseSummary) SetReserved(v int64) {
	o.Reserved = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LicenseSummary) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseSummary) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LicenseSummary) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LicenseSummary) SetStatus(v string) {
	o.Status = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LicenseSummary) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseSummary) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LicenseSummary) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LicenseSummary) SetName(v string) {
	o.Name = &v
}

// GetEnforced returns the Enforced field value if set, zero value otherwise.
func (o *LicenseSummary) GetEnforced() bool {
	if o == nil || o.Enforced == nil {
		var ret bool
		return ret
	}
	return *o.Enforced
}

// GetEnforcedOk returns a tuple with the Enforced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseSummary) GetEnforcedOk() (*bool, bool) {
	if o == nil || o.Enforced == nil {
		return nil, false
	}
	return o.Enforced, true
}

// HasEnforced returns a boolean if a field has been set.
func (o *LicenseSummary) HasEnforced() bool {
	if o != nil && o.Enforced != nil {
		return true
	}

	return false
}

// SetEnforced gets a reference to the given bool and assigns it to the Enforced field.
func (o *LicenseSummary) SetEnforced(v bool) {
	o.Enforced = &v
}

// GetLicenseDetails returns the LicenseDetails field value if set, zero value otherwise.
func (o *LicenseSummary) GetLicenseDetails() []LicenseDetails {
	if o == nil || o.LicenseDetails == nil {
		var ret []LicenseDetails
		return ret
	}
	return *o.LicenseDetails
}

// GetLicenseDetailsOk returns a tuple with the LicenseDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseSummary) GetLicenseDetailsOk() (*[]LicenseDetails, bool) {
	if o == nil || o.LicenseDetails == nil {
		return nil, false
	}
	return o.LicenseDetails, true
}

// HasLicenseDetails returns a boolean if a field has been set.
func (o *LicenseSummary) HasLicenseDetails() bool {
	if o != nil && o.LicenseDetails != nil {
		return true
	}

	return false
}

// SetLicenseDetails gets a reference to the given []LicenseDetails and assigns it to the LicenseDetails field.
func (o *LicenseSummary) SetLicenseDetails(v []LicenseDetails) {
	o.LicenseDetails = &v
}

func (o LicenseSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Entitled != nil {
		toSerialize["entitled"] = o.Entitled
	}
	if o.Inuse != nil {
		toSerialize["inuse"] = o.Inuse
	}
	if o.Reserved != nil {
		toSerialize["reserved"] = o.Reserved
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Enforced != nil {
		toSerialize["enforced"] = o.Enforced
	}
	if o.LicenseDetails != nil {
		toSerialize["licenseDetails"] = o.LicenseDetails
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LicenseSummary) UnmarshalJSON(bytes []byte) (err error) {
	varLicenseSummary := _LicenseSummary{}

	if err = json.Unmarshal(bytes, &varLicenseSummary); err == nil {
		*o = LicenseSummary(varLicenseSummary)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "entitled")
		delete(additionalProperties, "inuse")
		delete(additionalProperties, "reserved")
		delete(additionalProperties, "status")
		delete(additionalProperties, "name")
		delete(additionalProperties, "enforced")
		delete(additionalProperties, "licenseDetails")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLicenseSummary struct {
	value *LicenseSummary
	isSet bool
}

func (v NullableLicenseSummary) Get() *LicenseSummary {
	return v.value
}

func (v *NullableLicenseSummary) Set(val *LicenseSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseSummary(val *LicenseSummary) *NullableLicenseSummary {
	return &NullableLicenseSummary{value: val, isSet: true}
}

func (v NullableLicenseSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

