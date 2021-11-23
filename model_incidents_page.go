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

// IncidentsPage struct for IncidentsPage
type IncidentsPage struct {
	Page *int32 `json:"page,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty"`
	TotalItems NullableInt64 `json:"totalItems,omitempty"`
	HasNext NullableBool `json:"hasNext,omitempty"`
	HasPrevious NullableBool `json:"hasPrevious,omitempty"`
	SortBy NullableString `json:"sortBy,omitempty"`
	SortOrder NullableString `json:"sortOrder,omitempty"`
	Contents *[]Incident `json:"contents,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IncidentsPage IncidentsPage

// NewIncidentsPage instantiates a new IncidentsPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentsPage() *IncidentsPage {
	this := IncidentsPage{}
	return &this
}

// NewIncidentsPageWithDefaults instantiates a new IncidentsPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentsPageWithDefaults() *IncidentsPage {
	this := IncidentsPage{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *IncidentsPage) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentsPage) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *IncidentsPage) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *IncidentsPage) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *IncidentsPage) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentsPage) GetPageSizeOk() (*int32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *IncidentsPage) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *IncidentsPage) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentsPage) GetTotalItems() int64 {
	if o == nil || o.TotalItems.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TotalItems.Get()
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentsPage) GetTotalItemsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TotalItems.Get(), o.TotalItems.IsSet()
}

// HasTotalItems returns a boolean if a field has been set.
func (o *IncidentsPage) HasTotalItems() bool {
	if o != nil && o.TotalItems.IsSet() {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given NullableInt64 and assigns it to the TotalItems field.
func (o *IncidentsPage) SetTotalItems(v int64) {
	o.TotalItems.Set(&v)
}
// SetTotalItemsNil sets the value for TotalItems to be an explicit nil
func (o *IncidentsPage) SetTotalItemsNil() {
	o.TotalItems.Set(nil)
}

// UnsetTotalItems ensures that no value is present for TotalItems, not even an explicit nil
func (o *IncidentsPage) UnsetTotalItems() {
	o.TotalItems.Unset()
}

// GetHasNext returns the HasNext field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentsPage) GetHasNext() bool {
	if o == nil || o.HasNext.Get() == nil {
		var ret bool
		return ret
	}
	return *o.HasNext.Get()
}

// GetHasNextOk returns a tuple with the HasNext field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentsPage) GetHasNextOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HasNext.Get(), o.HasNext.IsSet()
}

// HasHasNext returns a boolean if a field has been set.
func (o *IncidentsPage) HasHasNext() bool {
	if o != nil && o.HasNext.IsSet() {
		return true
	}

	return false
}

// SetHasNext gets a reference to the given NullableBool and assigns it to the HasNext field.
func (o *IncidentsPage) SetHasNext(v bool) {
	o.HasNext.Set(&v)
}
// SetHasNextNil sets the value for HasNext to be an explicit nil
func (o *IncidentsPage) SetHasNextNil() {
	o.HasNext.Set(nil)
}

// UnsetHasNext ensures that no value is present for HasNext, not even an explicit nil
func (o *IncidentsPage) UnsetHasNext() {
	o.HasNext.Unset()
}

// GetHasPrevious returns the HasPrevious field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentsPage) GetHasPrevious() bool {
	if o == nil || o.HasPrevious.Get() == nil {
		var ret bool
		return ret
	}
	return *o.HasPrevious.Get()
}

// GetHasPreviousOk returns a tuple with the HasPrevious field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentsPage) GetHasPreviousOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HasPrevious.Get(), o.HasPrevious.IsSet()
}

// HasHasPrevious returns a boolean if a field has been set.
func (o *IncidentsPage) HasHasPrevious() bool {
	if o != nil && o.HasPrevious.IsSet() {
		return true
	}

	return false
}

// SetHasPrevious gets a reference to the given NullableBool and assigns it to the HasPrevious field.
func (o *IncidentsPage) SetHasPrevious(v bool) {
	o.HasPrevious.Set(&v)
}
// SetHasPreviousNil sets the value for HasPrevious to be an explicit nil
func (o *IncidentsPage) SetHasPreviousNil() {
	o.HasPrevious.Set(nil)
}

// UnsetHasPrevious ensures that no value is present for HasPrevious, not even an explicit nil
func (o *IncidentsPage) UnsetHasPrevious() {
	o.HasPrevious.Unset()
}

// GetSortBy returns the SortBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentsPage) GetSortBy() string {
	if o == nil || o.SortBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.SortBy.Get()
}

// GetSortByOk returns a tuple with the SortBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentsPage) GetSortByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SortBy.Get(), o.SortBy.IsSet()
}

// HasSortBy returns a boolean if a field has been set.
func (o *IncidentsPage) HasSortBy() bool {
	if o != nil && o.SortBy.IsSet() {
		return true
	}

	return false
}

// SetSortBy gets a reference to the given NullableString and assigns it to the SortBy field.
func (o *IncidentsPage) SetSortBy(v string) {
	o.SortBy.Set(&v)
}
// SetSortByNil sets the value for SortBy to be an explicit nil
func (o *IncidentsPage) SetSortByNil() {
	o.SortBy.Set(nil)
}

// UnsetSortBy ensures that no value is present for SortBy, not even an explicit nil
func (o *IncidentsPage) UnsetSortBy() {
	o.SortBy.Unset()
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentsPage) GetSortOrder() string {
	if o == nil || o.SortOrder.Get() == nil {
		var ret string
		return ret
	}
	return *o.SortOrder.Get()
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentsPage) GetSortOrderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SortOrder.Get(), o.SortOrder.IsSet()
}

// HasSortOrder returns a boolean if a field has been set.
func (o *IncidentsPage) HasSortOrder() bool {
	if o != nil && o.SortOrder.IsSet() {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given NullableString and assigns it to the SortOrder field.
func (o *IncidentsPage) SetSortOrder(v string) {
	o.SortOrder.Set(&v)
}
// SetSortOrderNil sets the value for SortOrder to be an explicit nil
func (o *IncidentsPage) SetSortOrderNil() {
	o.SortOrder.Set(nil)
}

// UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
func (o *IncidentsPage) UnsetSortOrder() {
	o.SortOrder.Unset()
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *IncidentsPage) GetContents() []Incident {
	if o == nil || o.Contents == nil {
		var ret []Incident
		return ret
	}
	return *o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentsPage) GetContentsOk() (*[]Incident, bool) {
	if o == nil || o.Contents == nil {
		return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *IncidentsPage) HasContents() bool {
	if o != nil && o.Contents != nil {
		return true
	}

	return false
}

// SetContents gets a reference to the given []Incident and assigns it to the Contents field.
func (o *IncidentsPage) SetContents(v []Incident) {
	o.Contents = &v
}

func (o IncidentsPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.PageSize != nil {
		toSerialize["pageSize"] = o.PageSize
	}
	if o.TotalItems.IsSet() {
		toSerialize["totalItems"] = o.TotalItems.Get()
	}
	if o.HasNext.IsSet() {
		toSerialize["hasNext"] = o.HasNext.Get()
	}
	if o.HasPrevious.IsSet() {
		toSerialize["hasPrevious"] = o.HasPrevious.Get()
	}
	if o.SortBy.IsSet() {
		toSerialize["sortBy"] = o.SortBy.Get()
	}
	if o.SortOrder.IsSet() {
		toSerialize["sortOrder"] = o.SortOrder.Get()
	}
	if o.Contents != nil {
		toSerialize["contents"] = o.Contents
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IncidentsPage) UnmarshalJSON(bytes []byte) (err error) {
	varIncidentsPage := _IncidentsPage{}

	if err = json.Unmarshal(bytes, &varIncidentsPage); err == nil {
		*o = IncidentsPage(varIncidentsPage)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "page")
		delete(additionalProperties, "pageSize")
		delete(additionalProperties, "totalItems")
		delete(additionalProperties, "hasNext")
		delete(additionalProperties, "hasPrevious")
		delete(additionalProperties, "sortBy")
		delete(additionalProperties, "sortOrder")
		delete(additionalProperties, "contents")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIncidentsPage struct {
	value *IncidentsPage
	isSet bool
}

func (v NullableIncidentsPage) Get() *IncidentsPage {
	return v.value
}

func (v *NullableIncidentsPage) Set(val *IncidentsPage) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentsPage) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentsPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentsPage(val *IncidentsPage) *NullableIncidentsPage {
	return &NullableIncidentsPage{value: val, isSet: true}
}

func (v NullableIncidentsPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentsPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

