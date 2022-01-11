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

// ServiceUIResource struct for ServiceUIResource
type ServiceUIResource struct {
	Type *string `json:"type,omitempty"`
	Href *string `json:"href,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceUIResource ServiceUIResource

// NewServiceUIResource instantiates a new ServiceUIResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceUIResource() *ServiceUIResource {
	this := ServiceUIResource{}
	return &this
}

// NewServiceUIResourceWithDefaults instantiates a new ServiceUIResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceUIResourceWithDefaults() *ServiceUIResource {
	this := ServiceUIResource{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ServiceUIResource) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceUIResource) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ServiceUIResource) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ServiceUIResource) SetType(v string) {
	o.Type = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ServiceUIResource) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceUIResource) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ServiceUIResource) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ServiceUIResource) SetHref(v string) {
	o.Href = &v
}

func (o ServiceUIResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServiceUIResource) UnmarshalJSON(bytes []byte) (err error) {
	varServiceUIResource := _ServiceUIResource{}

	if err = json.Unmarshal(bytes, &varServiceUIResource); err == nil {
		*o = ServiceUIResource(varServiceUIResource)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "href")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceUIResource struct {
	value *ServiceUIResource
	isSet bool
}

func (v NullableServiceUIResource) Get() *ServiceUIResource {
	return v.value
}

func (v *NullableServiceUIResource) Set(val *ServiceUIResource) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceUIResource) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceUIResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceUIResource(val *ServiceUIResource) *NullableServiceUIResource {
	return &NullableServiceUIResource{value: val, isSet: true}
}

func (v NullableServiceUIResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceUIResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


