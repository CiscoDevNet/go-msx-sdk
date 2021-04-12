/*
 * KAKAPO - MSX SDK
 *
 * MSX Platform SDK
 *
 * API version: 1.0.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msxsdk

import (
	"encoding/json"
)

// SiteUpdate struct for SiteUpdate
type SiteUpdate struct {
	ParentId *string `json:"parentId,omitempty"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	Address NullableSiteAddress `json:"address,omitempty"`
	Contact NullableSiteContact `json:"contact,omitempty"`
	Location NullableSiteLocation `json:"location,omitempty"`
	Image *string `json:"image,omitempty"`
	Attributes *map[string]string `json:"attributes,omitempty"`
}

// NewSiteUpdate instantiates a new SiteUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteUpdate(name string) *SiteUpdate {
	this := SiteUpdate{}
	this.Name = name
	return &this
}

// NewSiteUpdateWithDefaults instantiates a new SiteUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteUpdateWithDefaults() *SiteUpdate {
	this := SiteUpdate{}
	return &this
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *SiteUpdate) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteUpdate) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *SiteUpdate) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *SiteUpdate) SetParentId(v string) {
	o.ParentId = &v
}

// GetName returns the Name field value
func (o *SiteUpdate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SiteUpdate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SiteUpdate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SiteUpdate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SiteUpdate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SiteUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SiteUpdate) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteUpdate) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SiteUpdate) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SiteUpdate) SetType(v string) {
	o.Type = &v
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SiteUpdate) GetAddress() SiteAddress {
	if o == nil || o.Address.Get() == nil {
		var ret SiteAddress
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SiteUpdate) GetAddressOk() (*SiteAddress, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *SiteUpdate) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableSiteAddress and assigns it to the Address field.
func (o *SiteUpdate) SetAddress(v SiteAddress) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *SiteUpdate) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *SiteUpdate) UnsetAddress() {
	o.Address.Unset()
}

// GetContact returns the Contact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SiteUpdate) GetContact() SiteContact {
	if o == nil || o.Contact.Get() == nil {
		var ret SiteContact
		return ret
	}
	return *o.Contact.Get()
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SiteUpdate) GetContactOk() (*SiteContact, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Contact.Get(), o.Contact.IsSet()
}

// HasContact returns a boolean if a field has been set.
func (o *SiteUpdate) HasContact() bool {
	if o != nil && o.Contact.IsSet() {
		return true
	}

	return false
}

// SetContact gets a reference to the given NullableSiteContact and assigns it to the Contact field.
func (o *SiteUpdate) SetContact(v SiteContact) {
	o.Contact.Set(&v)
}
// SetContactNil sets the value for Contact to be an explicit nil
func (o *SiteUpdate) SetContactNil() {
	o.Contact.Set(nil)
}

// UnsetContact ensures that no value is present for Contact, not even an explicit nil
func (o *SiteUpdate) UnsetContact() {
	o.Contact.Unset()
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SiteUpdate) GetLocation() SiteLocation {
	if o == nil || o.Location.Get() == nil {
		var ret SiteLocation
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SiteUpdate) GetLocationOk() (*SiteLocation, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *SiteUpdate) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableSiteLocation and assigns it to the Location field.
func (o *SiteUpdate) SetLocation(v SiteLocation) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *SiteUpdate) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *SiteUpdate) UnsetLocation() {
	o.Location.Unset()
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *SiteUpdate) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteUpdate) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *SiteUpdate) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *SiteUpdate) SetImage(v string) {
	o.Image = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SiteUpdate) GetAttributes() map[string]string {
	if o == nil || o.Attributes == nil {
		var ret map[string]string
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteUpdate) GetAttributesOk() (*map[string]string, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SiteUpdate) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]string and assigns it to the Attributes field.
func (o *SiteUpdate) SetAttributes(v map[string]string) {
	o.Attributes = &v
}

func (o SiteUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ParentId != nil {
		toSerialize["parentId"] = o.ParentId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if o.Contact.IsSet() {
		toSerialize["contact"] = o.Contact.Get()
	}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableSiteUpdate struct {
	value *SiteUpdate
	isSet bool
}

func (v NullableSiteUpdate) Get() *SiteUpdate {
	return v.value
}

func (v *NullableSiteUpdate) Set(val *SiteUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteUpdate(val *SiteUpdate) *NullableSiteUpdate {
	return &NullableSiteUpdate{value: val, isSet: true}
}

func (v NullableSiteUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


