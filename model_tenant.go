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
	"time"
)

// Tenant struct for Tenant
type Tenant struct {
	Id *string `json:"id,omitempty"`
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	ModifiedOn *time.Time `json:"modifiedOn,omitempty"`
	Suspended NullableBool `json:"suspended,omitempty"`
	NumberOfChildren NullableInt64 `json:"numberOfChildren,omitempty"`
	ParentId NullableString `json:"parentId,omitempty"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Url *string `json:"url,omitempty"`
	Image NullableString `json:"image,omitempty"`
}

// NewTenant instantiates a new Tenant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenant(name string) *Tenant {
	this := Tenant{}
	this.Name = name
	return &this
}

// NewTenantWithDefaults instantiates a new Tenant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantWithDefaults() *Tenant {
	this := Tenant{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Tenant) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Tenant) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Tenant) SetId(v string) {
	o.Id = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *Tenant) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *Tenant) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *Tenant) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetModifiedOn returns the ModifiedOn field value if set, zero value otherwise.
func (o *Tenant) GetModifiedOn() time.Time {
	if o == nil || o.ModifiedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetModifiedOnOk() (*time.Time, bool) {
	if o == nil || o.ModifiedOn == nil {
		return nil, false
	}
	return o.ModifiedOn, true
}

// HasModifiedOn returns a boolean if a field has been set.
func (o *Tenant) HasModifiedOn() bool {
	if o != nil && o.ModifiedOn != nil {
		return true
	}

	return false
}

// SetModifiedOn gets a reference to the given time.Time and assigns it to the ModifiedOn field.
func (o *Tenant) SetModifiedOn(v time.Time) {
	o.ModifiedOn = &v
}

// GetSuspended returns the Suspended field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tenant) GetSuspended() bool {
	if o == nil || o.Suspended.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Suspended.Get()
}

// GetSuspendedOk returns a tuple with the Suspended field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tenant) GetSuspendedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Suspended.Get(), o.Suspended.IsSet()
}

// HasSuspended returns a boolean if a field has been set.
func (o *Tenant) HasSuspended() bool {
	if o != nil && o.Suspended.IsSet() {
		return true
	}

	return false
}

// SetSuspended gets a reference to the given NullableBool and assigns it to the Suspended field.
func (o *Tenant) SetSuspended(v bool) {
	o.Suspended.Set(&v)
}
// SetSuspendedNil sets the value for Suspended to be an explicit nil
func (o *Tenant) SetSuspendedNil() {
	o.Suspended.Set(nil)
}

// UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
func (o *Tenant) UnsetSuspended() {
	o.Suspended.Unset()
}

// GetNumberOfChildren returns the NumberOfChildren field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tenant) GetNumberOfChildren() int64 {
	if o == nil || o.NumberOfChildren.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfChildren.Get()
}

// GetNumberOfChildrenOk returns a tuple with the NumberOfChildren field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tenant) GetNumberOfChildrenOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NumberOfChildren.Get(), o.NumberOfChildren.IsSet()
}

// HasNumberOfChildren returns a boolean if a field has been set.
func (o *Tenant) HasNumberOfChildren() bool {
	if o != nil && o.NumberOfChildren.IsSet() {
		return true
	}

	return false
}

// SetNumberOfChildren gets a reference to the given NullableInt64 and assigns it to the NumberOfChildren field.
func (o *Tenant) SetNumberOfChildren(v int64) {
	o.NumberOfChildren.Set(&v)
}
// SetNumberOfChildrenNil sets the value for NumberOfChildren to be an explicit nil
func (o *Tenant) SetNumberOfChildrenNil() {
	o.NumberOfChildren.Set(nil)
}

// UnsetNumberOfChildren ensures that no value is present for NumberOfChildren, not even an explicit nil
func (o *Tenant) UnsetNumberOfChildren() {
	o.NumberOfChildren.Unset()
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tenant) GetParentId() string {
	if o == nil || o.ParentId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tenant) GetParentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *Tenant) HasParentId() bool {
	if o != nil && o.ParentId.IsSet() {
		return true
	}

	return false
}

// SetParentId gets a reference to the given NullableString and assigns it to the ParentId field.
func (o *Tenant) SetParentId(v string) {
	o.ParentId.Set(&v)
}
// SetParentIdNil sets the value for ParentId to be an explicit nil
func (o *Tenant) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
func (o *Tenant) UnsetParentId() {
	o.ParentId.Unset()
}

// GetName returns the Name field value
func (o *Tenant) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Tenant) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Tenant) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Tenant) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Tenant) SetDescription(v string) {
	o.Description = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Tenant) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Tenant) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Tenant) SetUrl(v string) {
	o.Url = &v
}

// GetImage returns the Image field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tenant) GetImage() string {
	if o == nil || o.Image.Get() == nil {
		var ret string
		return ret
	}
	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tenant) GetImageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// HasImage returns a boolean if a field has been set.
func (o *Tenant) HasImage() bool {
	if o != nil && o.Image.IsSet() {
		return true
	}

	return false
}

// SetImage gets a reference to the given NullableString and assigns it to the Image field.
func (o *Tenant) SetImage(v string) {
	o.Image.Set(&v)
}
// SetImageNil sets the value for Image to be an explicit nil
func (o *Tenant) SetImageNil() {
	o.Image.Set(nil)
}

// UnsetImage ensures that no value is present for Image, not even an explicit nil
func (o *Tenant) UnsetImage() {
	o.Image.Unset()
}

func (o Tenant) MarshalJSON() ([]byte, error) {
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
	if o.Suspended.IsSet() {
		toSerialize["suspended"] = o.Suspended.Get()
	}
	if o.NumberOfChildren.IsSet() {
		toSerialize["numberOfChildren"] = o.NumberOfChildren.Get()
	}
	if o.ParentId.IsSet() {
		toSerialize["parentId"] = o.ParentId.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Image.IsSet() {
		toSerialize["image"] = o.Image.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTenant struct {
	value *Tenant
	isSet bool
}

func (v NullableTenant) Get() *Tenant {
	return v.value
}

func (v *NullableTenant) Set(val *Tenant) {
	v.value = val
	v.isSet = true
}

func (v NullableTenant) IsSet() bool {
	return v.isSet
}

func (v *NullableTenant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenant(val *Tenant) *NullableTenant {
	return &NullableTenant{value: val, isSet: true}
}

func (v NullableTenant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


