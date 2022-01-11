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

// ProductsPageAllOf struct for ProductsPageAllOf
type ProductsPageAllOf struct {
	Contents *[]Product `json:"contents,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductsPageAllOf ProductsPageAllOf

// NewProductsPageAllOf instantiates a new ProductsPageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductsPageAllOf() *ProductsPageAllOf {
	this := ProductsPageAllOf{}
	return &this
}

// NewProductsPageAllOfWithDefaults instantiates a new ProductsPageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductsPageAllOfWithDefaults() *ProductsPageAllOf {
	this := ProductsPageAllOf{}
	return &this
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *ProductsPageAllOf) GetContents() []Product {
	if o == nil || o.Contents == nil {
		var ret []Product
		return ret
	}
	return *o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductsPageAllOf) GetContentsOk() (*[]Product, bool) {
	if o == nil || o.Contents == nil {
		return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *ProductsPageAllOf) HasContents() bool {
	if o != nil && o.Contents != nil {
		return true
	}

	return false
}

// SetContents gets a reference to the given []Product and assigns it to the Contents field.
func (o *ProductsPageAllOf) SetContents(v []Product) {
	o.Contents = &v
}

func (o ProductsPageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Contents != nil {
		toSerialize["contents"] = o.Contents
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProductsPageAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varProductsPageAllOf := _ProductsPageAllOf{}

	if err = json.Unmarshal(bytes, &varProductsPageAllOf); err == nil {
		*o = ProductsPageAllOf(varProductsPageAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "contents")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProductsPageAllOf struct {
	value *ProductsPageAllOf
	isSet bool
}

func (v NullableProductsPageAllOf) Get() *ProductsPageAllOf {
	return v.value
}

func (v *NullableProductsPageAllOf) Set(val *ProductsPageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableProductsPageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableProductsPageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductsPageAllOf(val *ProductsPageAllOf) *NullableProductsPageAllOf {
	return &NullableProductsPageAllOf{value: val, isSet: true}
}

func (v NullableProductsPageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductsPageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


