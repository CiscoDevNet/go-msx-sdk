/*
 * MSX SDK
 *
 * MSX SDK client.
 *
 * API version: 1.0.5
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msxsdk

import (
	"encoding/json"
)

// SiteCreate struct for SiteCreate
type SiteCreate struct {
	TenantId string `json:"tenantId"`
	DeviceIds *[]string `json:"deviceIds,omitempty"`
	ServiceIds *[]string `json:"serviceIds,omitempty"`
	ParentId *string `json:"parentId,omitempty"`
	Name string `json:"name"`
	Description NullableString `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	Address NullableSiteAddress `json:"address,omitempty"`
	Contact NullableSiteContact `json:"contact,omitempty"`
	Location NullableSiteLocation `json:"location,omitempty"`
	Image *string `json:"image,omitempty"`
	Attributes *map[string]string `json:"attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SiteCreate SiteCreate

// NewSiteCreate instantiates a new SiteCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteCreate(tenantId string, name string) *SiteCreate {
	this := SiteCreate{}
	this.TenantId = tenantId
	this.Name = name
	return &this
}

// NewSiteCreateWithDefaults instantiates a new SiteCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteCreateWithDefaults() *SiteCreate {
	this := SiteCreate{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *SiteCreate) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *SiteCreate) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *SiteCreate) SetTenantId(v string) {
	o.TenantId = v
}

// GetDeviceIds returns the DeviceIds field value if set, zero value otherwise.
func (o *SiteCreate) GetDeviceIds() []string {
	if o == nil || o.DeviceIds == nil {
		var ret []string
		return ret
	}
	return *o.DeviceIds
}

// GetDeviceIdsOk returns a tuple with the DeviceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteCreate) GetDeviceIdsOk() (*[]string, bool) {
	if o == nil || o.DeviceIds == nil {
		return nil, false
	}
	return o.DeviceIds, true
}

// HasDeviceIds returns a boolean if a field has been set.
func (o *SiteCreate) HasDeviceIds() bool {
	if o != nil && o.DeviceIds != nil {
		return true
	}

	return false
}

// SetDeviceIds gets a reference to the given []string and assigns it to the DeviceIds field.
func (o *SiteCreate) SetDeviceIds(v []string) {
	o.DeviceIds = &v
}

// GetServiceIds returns the ServiceIds field value if set, zero value otherwise.
func (o *SiteCreate) GetServiceIds() []string {
	if o == nil || o.ServiceIds == nil {
		var ret []string
		return ret
	}
	return *o.ServiceIds
}

// GetServiceIdsOk returns a tuple with the ServiceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteCreate) GetServiceIdsOk() (*[]string, bool) {
	if o == nil || o.ServiceIds == nil {
		return nil, false
	}
	return o.ServiceIds, true
}

// HasServiceIds returns a boolean if a field has been set.
func (o *SiteCreate) HasServiceIds() bool {
	if o != nil && o.ServiceIds != nil {
		return true
	}

	return false
}

// SetServiceIds gets a reference to the given []string and assigns it to the ServiceIds field.
func (o *SiteCreate) SetServiceIds(v []string) {
	o.ServiceIds = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *SiteCreate) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteCreate) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *SiteCreate) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *SiteCreate) SetParentId(v string) {
	o.ParentId = &v
}

// GetName returns the Name field value
func (o *SiteCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SiteCreate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SiteCreate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SiteCreate) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SiteCreate) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *SiteCreate) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *SiteCreate) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *SiteCreate) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *SiteCreate) UnsetDescription() {
	o.Description.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SiteCreate) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteCreate) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SiteCreate) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SiteCreate) SetType(v string) {
	o.Type = &v
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SiteCreate) GetAddress() SiteAddress {
	if o == nil || o.Address.Get() == nil {
		var ret SiteAddress
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SiteCreate) GetAddressOk() (*SiteAddress, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *SiteCreate) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableSiteAddress and assigns it to the Address field.
func (o *SiteCreate) SetAddress(v SiteAddress) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *SiteCreate) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *SiteCreate) UnsetAddress() {
	o.Address.Unset()
}

// GetContact returns the Contact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SiteCreate) GetContact() SiteContact {
	if o == nil || o.Contact.Get() == nil {
		var ret SiteContact
		return ret
	}
	return *o.Contact.Get()
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SiteCreate) GetContactOk() (*SiteContact, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Contact.Get(), o.Contact.IsSet()
}

// HasContact returns a boolean if a field has been set.
func (o *SiteCreate) HasContact() bool {
	if o != nil && o.Contact.IsSet() {
		return true
	}

	return false
}

// SetContact gets a reference to the given NullableSiteContact and assigns it to the Contact field.
func (o *SiteCreate) SetContact(v SiteContact) {
	o.Contact.Set(&v)
}
// SetContactNil sets the value for Contact to be an explicit nil
func (o *SiteCreate) SetContactNil() {
	o.Contact.Set(nil)
}

// UnsetContact ensures that no value is present for Contact, not even an explicit nil
func (o *SiteCreate) UnsetContact() {
	o.Contact.Unset()
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SiteCreate) GetLocation() SiteLocation {
	if o == nil || o.Location.Get() == nil {
		var ret SiteLocation
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SiteCreate) GetLocationOk() (*SiteLocation, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *SiteCreate) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableSiteLocation and assigns it to the Location field.
func (o *SiteCreate) SetLocation(v SiteLocation) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *SiteCreate) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *SiteCreate) UnsetLocation() {
	o.Location.Unset()
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *SiteCreate) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteCreate) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *SiteCreate) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *SiteCreate) SetImage(v string) {
	o.Image = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SiteCreate) GetAttributes() map[string]string {
	if o == nil || o.Attributes == nil {
		var ret map[string]string
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteCreate) GetAttributesOk() (*map[string]string, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SiteCreate) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]string and assigns it to the Attributes field.
func (o *SiteCreate) SetAttributes(v map[string]string) {
	o.Attributes = &v
}

func (o SiteCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tenantId"] = o.TenantId
	}
	if o.DeviceIds != nil {
		toSerialize["deviceIds"] = o.DeviceIds
	}
	if o.ServiceIds != nil {
		toSerialize["serviceIds"] = o.ServiceIds
	}
	if o.ParentId != nil {
		toSerialize["parentId"] = o.ParentId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SiteCreate) UnmarshalJSON(bytes []byte) (err error) {
	varSiteCreate := _SiteCreate{}

	if err = json.Unmarshal(bytes, &varSiteCreate); err == nil {
		*o = SiteCreate(varSiteCreate)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "deviceIds")
		delete(additionalProperties, "serviceIds")
		delete(additionalProperties, "parentId")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "type")
		delete(additionalProperties, "address")
		delete(additionalProperties, "contact")
		delete(additionalProperties, "location")
		delete(additionalProperties, "image")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSiteCreate struct {
	value *SiteCreate
	isSet bool
}

func (v NullableSiteCreate) Get() *SiteCreate {
	return v.value
}

func (v *NullableSiteCreate) Set(val *SiteCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteCreate(val *SiteCreate) *NullableSiteCreate {
	return &NullableSiteCreate{value: val, isSet: true}
}

func (v NullableSiteCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


