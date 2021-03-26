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

// TenantsPage struct for TenantsPage
type TenantsPage struct {
	Page *int32 `json:"page,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty"`
	TotalItems *int64 `json:"totalItems,omitempty"`
	HasNext NullableBool `json:"hasNext,omitempty"`
	HasPrevious NullableBool `json:"hasPrevious,omitempty"`
	SortBy *string `json:"sortBy,omitempty"`
	SortOrder *string `json:"sortOrder,omitempty"`
	Contents *[]Tenant `json:"contents,omitempty"`
}

// NewTenantsPage instantiates a new TenantsPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantsPage() *TenantsPage {
	this := TenantsPage{}
	return &this
}

// NewTenantsPageWithDefaults instantiates a new TenantsPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantsPageWithDefaults() *TenantsPage {
	this := TenantsPage{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *TenantsPage) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantsPage) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *TenantsPage) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *TenantsPage) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *TenantsPage) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantsPage) GetPageSizeOk() (*int32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *TenantsPage) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *TenantsPage) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *TenantsPage) GetTotalItems() int64 {
	if o == nil || o.TotalItems == nil {
		var ret int64
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantsPage) GetTotalItemsOk() (*int64, bool) {
	if o == nil || o.TotalItems == nil {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *TenantsPage) HasTotalItems() bool {
	if o != nil && o.TotalItems != nil {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int64 and assigns it to the TotalItems field.
func (o *TenantsPage) SetTotalItems(v int64) {
	o.TotalItems = &v
}

// GetHasNext returns the HasNext field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantsPage) GetHasNext() bool {
	if o == nil || o.HasNext.Get() == nil {
		var ret bool
		return ret
	}
	return *o.HasNext.Get()
}

// GetHasNextOk returns a tuple with the HasNext field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantsPage) GetHasNextOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HasNext.Get(), o.HasNext.IsSet()
}

// HasHasNext returns a boolean if a field has been set.
func (o *TenantsPage) HasHasNext() bool {
	if o != nil && o.HasNext.IsSet() {
		return true
	}

	return false
}

// SetHasNext gets a reference to the given NullableBool and assigns it to the HasNext field.
func (o *TenantsPage) SetHasNext(v bool) {
	o.HasNext.Set(&v)
}
// SetHasNextNil sets the value for HasNext to be an explicit nil
func (o *TenantsPage) SetHasNextNil() {
	o.HasNext.Set(nil)
}

// UnsetHasNext ensures that no value is present for HasNext, not even an explicit nil
func (o *TenantsPage) UnsetHasNext() {
	o.HasNext.Unset()
}

// GetHasPrevious returns the HasPrevious field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantsPage) GetHasPrevious() bool {
	if o == nil || o.HasPrevious.Get() == nil {
		var ret bool
		return ret
	}
	return *o.HasPrevious.Get()
}

// GetHasPreviousOk returns a tuple with the HasPrevious field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantsPage) GetHasPreviousOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HasPrevious.Get(), o.HasPrevious.IsSet()
}

// HasHasPrevious returns a boolean if a field has been set.
func (o *TenantsPage) HasHasPrevious() bool {
	if o != nil && o.HasPrevious.IsSet() {
		return true
	}

	return false
}

// SetHasPrevious gets a reference to the given NullableBool and assigns it to the HasPrevious field.
func (o *TenantsPage) SetHasPrevious(v bool) {
	o.HasPrevious.Set(&v)
}
// SetHasPreviousNil sets the value for HasPrevious to be an explicit nil
func (o *TenantsPage) SetHasPreviousNil() {
	o.HasPrevious.Set(nil)
}

// UnsetHasPrevious ensures that no value is present for HasPrevious, not even an explicit nil
func (o *TenantsPage) UnsetHasPrevious() {
	o.HasPrevious.Unset()
}

// GetSortBy returns the SortBy field value if set, zero value otherwise.
func (o *TenantsPage) GetSortBy() string {
	if o == nil || o.SortBy == nil {
		var ret string
		return ret
	}
	return *o.SortBy
}

// GetSortByOk returns a tuple with the SortBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantsPage) GetSortByOk() (*string, bool) {
	if o == nil || o.SortBy == nil {
		return nil, false
	}
	return o.SortBy, true
}

// HasSortBy returns a boolean if a field has been set.
func (o *TenantsPage) HasSortBy() bool {
	if o != nil && o.SortBy != nil {
		return true
	}

	return false
}

// SetSortBy gets a reference to the given string and assigns it to the SortBy field.
func (o *TenantsPage) SetSortBy(v string) {
	o.SortBy = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *TenantsPage) GetSortOrder() string {
	if o == nil || o.SortOrder == nil {
		var ret string
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantsPage) GetSortOrderOk() (*string, bool) {
	if o == nil || o.SortOrder == nil {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *TenantsPage) HasSortOrder() bool {
	if o != nil && o.SortOrder != nil {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given string and assigns it to the SortOrder field.
func (o *TenantsPage) SetSortOrder(v string) {
	o.SortOrder = &v
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *TenantsPage) GetContents() []Tenant {
	if o == nil || o.Contents == nil {
		var ret []Tenant
		return ret
	}
	return *o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantsPage) GetContentsOk() (*[]Tenant, bool) {
	if o == nil || o.Contents == nil {
		return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *TenantsPage) HasContents() bool {
	if o != nil && o.Contents != nil {
		return true
	}

	return false
}

// SetContents gets a reference to the given []Tenant and assigns it to the Contents field.
func (o *TenantsPage) SetContents(v []Tenant) {
	o.Contents = &v
}

func (o TenantsPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.PageSize != nil {
		toSerialize["pageSize"] = o.PageSize
	}
	if o.TotalItems != nil {
		toSerialize["totalItems"] = o.TotalItems
	}
	if o.HasNext.IsSet() {
		toSerialize["hasNext"] = o.HasNext.Get()
	}
	if o.HasPrevious.IsSet() {
		toSerialize["hasPrevious"] = o.HasPrevious.Get()
	}
	if o.SortBy != nil {
		toSerialize["sortBy"] = o.SortBy
	}
	if o.SortOrder != nil {
		toSerialize["sortOrder"] = o.SortOrder
	}
	if o.Contents != nil {
		toSerialize["contents"] = o.Contents
	}
	return json.Marshal(toSerialize)
}

type NullableTenantsPage struct {
	value *TenantsPage
	isSet bool
}

func (v NullableTenantsPage) Get() *TenantsPage {
	return v.value
}

func (v *NullableTenantsPage) Set(val *TenantsPage) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantsPage) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantsPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantsPage(val *TenantsPage) *NullableTenantsPage {
	return &NullableTenantsPage{value: val, isSet: true}
}

func (v NullableTenantsPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantsPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


