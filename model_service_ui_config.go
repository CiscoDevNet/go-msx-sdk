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

// ServiceUIConfig struct for ServiceUIConfig
type ServiceUIConfig struct {
	BannerImage *string `json:"bannerImage,omitempty"`
	Links *[]ServiceUILink `json:"links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceUIConfig ServiceUIConfig

// NewServiceUIConfig instantiates a new ServiceUIConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceUIConfig() *ServiceUIConfig {
	this := ServiceUIConfig{}
	return &this
}

// NewServiceUIConfigWithDefaults instantiates a new ServiceUIConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceUIConfigWithDefaults() *ServiceUIConfig {
	this := ServiceUIConfig{}
	return &this
}

// GetBannerImage returns the BannerImage field value if set, zero value otherwise.
func (o *ServiceUIConfig) GetBannerImage() string {
	if o == nil || o.BannerImage == nil {
		var ret string
		return ret
	}
	return *o.BannerImage
}

// GetBannerImageOk returns a tuple with the BannerImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceUIConfig) GetBannerImageOk() (*string, bool) {
	if o == nil || o.BannerImage == nil {
		return nil, false
	}
	return o.BannerImage, true
}

// HasBannerImage returns a boolean if a field has been set.
func (o *ServiceUIConfig) HasBannerImage() bool {
	if o != nil && o.BannerImage != nil {
		return true
	}

	return false
}

// SetBannerImage gets a reference to the given string and assigns it to the BannerImage field.
func (o *ServiceUIConfig) SetBannerImage(v string) {
	o.BannerImage = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ServiceUIConfig) GetLinks() []ServiceUILink {
	if o == nil || o.Links == nil {
		var ret []ServiceUILink
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceUIConfig) GetLinksOk() (*[]ServiceUILink, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ServiceUIConfig) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ServiceUILink and assigns it to the Links field.
func (o *ServiceUIConfig) SetLinks(v []ServiceUILink) {
	o.Links = &v
}

func (o ServiceUIConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BannerImage != nil {
		toSerialize["bannerImage"] = o.BannerImage
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServiceUIConfig) UnmarshalJSON(bytes []byte) (err error) {
	varServiceUIConfig := _ServiceUIConfig{}

	if err = json.Unmarshal(bytes, &varServiceUIConfig); err == nil {
		*o = ServiceUIConfig(varServiceUIConfig)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "bannerImage")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceUIConfig struct {
	value *ServiceUIConfig
	isSet bool
}

func (v NullableServiceUIConfig) Get() *ServiceUIConfig {
	return v.value
}

func (v *NullableServiceUIConfig) Set(val *ServiceUIConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceUIConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceUIConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceUIConfig(val *ServiceUIConfig) *NullableServiceUIConfig {
	return &NullableServiceUIConfig{value: val, isSet: true}
}

func (v NullableServiceUIConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceUIConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


