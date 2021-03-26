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

// SiteContact struct for SiteContact
type SiteContact struct {
	Name string `json:"name"`
	Phone string `json:"phone"`
	Email *string `json:"email,omitempty"`
}

// NewSiteContact instantiates a new SiteContact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteContact(name string, phone string) *SiteContact {
	this := SiteContact{}
	this.Name = name
	this.Phone = phone
	return &this
}

// NewSiteContactWithDefaults instantiates a new SiteContact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteContactWithDefaults() *SiteContact {
	this := SiteContact{}
	return &this
}

// GetName returns the Name field value
func (o *SiteContact) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SiteContact) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SiteContact) SetName(v string) {
	o.Name = v
}

// GetPhone returns the Phone field value
func (o *SiteContact) GetPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
func (o *SiteContact) GetPhoneOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Phone, true
}

// SetPhone sets field value
func (o *SiteContact) SetPhone(v string) {
	o.Phone = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *SiteContact) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteContact) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *SiteContact) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *SiteContact) SetEmail(v string) {
	o.Email = &v
}

func (o SiteContact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["phone"] = o.Phone
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableSiteContact struct {
	value *SiteContact
	isSet bool
}

func (v NullableSiteContact) Get() *SiteContact {
	return v.value
}

func (v *NullableSiteContact) Set(val *SiteContact) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteContact) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteContact(val *SiteContact) *NullableSiteContact {
	return &NullableSiteContact{value: val, isSet: true}
}

func (v NullableSiteContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


