# TenantsPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** |  | [optional] [readonly] 
**PageSize** | Pointer to **int32** |  | [optional] [readonly] 
**TotalItems** | Pointer to **NullableInt64** |  | [optional] [readonly] 
**HasNext** | Pointer to **NullableBool** |  | [optional] [readonly] 
**HasPrevious** | Pointer to **NullableBool** |  | [optional] [readonly] 
**SortBy** | Pointer to **NullableString** |  | [optional] [readonly] 
**SortOrder** | Pointer to **NullableString** |  | [optional] [readonly] 
**Contents** | Pointer to [**[]Tenant**](Tenant.md) |  | [optional] 

## Methods

### NewTenantsPage

`func NewTenantsPage() *TenantsPage`

NewTenantsPage instantiates a new TenantsPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantsPageWithDefaults

`func NewTenantsPageWithDefaults() *TenantsPage`

NewTenantsPageWithDefaults instantiates a new TenantsPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *TenantsPage) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *TenantsPage) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *TenantsPage) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *TenantsPage) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *TenantsPage) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *TenantsPage) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *TenantsPage) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *TenantsPage) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalItems

`func (o *TenantsPage) GetTotalItems() int64`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *TenantsPage) GetTotalItemsOk() (*int64, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *TenantsPage) SetTotalItems(v int64)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *TenantsPage) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### SetTotalItemsNil

`func (o *TenantsPage) SetTotalItemsNil(b bool)`

 SetTotalItemsNil sets the value for TotalItems to be an explicit nil

### UnsetTotalItems
`func (o *TenantsPage) UnsetTotalItems()`

UnsetTotalItems ensures that no value is present for TotalItems, not even an explicit nil
### GetHasNext

`func (o *TenantsPage) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *TenantsPage) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *TenantsPage) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *TenantsPage) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.

### SetHasNextNil

`func (o *TenantsPage) SetHasNextNil(b bool)`

 SetHasNextNil sets the value for HasNext to be an explicit nil

### UnsetHasNext
`func (o *TenantsPage) UnsetHasNext()`

UnsetHasNext ensures that no value is present for HasNext, not even an explicit nil
### GetHasPrevious

`func (o *TenantsPage) GetHasPrevious() bool`

GetHasPrevious returns the HasPrevious field if non-nil, zero value otherwise.

### GetHasPreviousOk

`func (o *TenantsPage) GetHasPreviousOk() (*bool, bool)`

GetHasPreviousOk returns a tuple with the HasPrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrevious

`func (o *TenantsPage) SetHasPrevious(v bool)`

SetHasPrevious sets HasPrevious field to given value.

### HasHasPrevious

`func (o *TenantsPage) HasHasPrevious() bool`

HasHasPrevious returns a boolean if a field has been set.

### SetHasPreviousNil

`func (o *TenantsPage) SetHasPreviousNil(b bool)`

 SetHasPreviousNil sets the value for HasPrevious to be an explicit nil

### UnsetHasPrevious
`func (o *TenantsPage) UnsetHasPrevious()`

UnsetHasPrevious ensures that no value is present for HasPrevious, not even an explicit nil
### GetSortBy

`func (o *TenantsPage) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *TenantsPage) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *TenantsPage) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *TenantsPage) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### SetSortByNil

`func (o *TenantsPage) SetSortByNil(b bool)`

 SetSortByNil sets the value for SortBy to be an explicit nil

### UnsetSortBy
`func (o *TenantsPage) UnsetSortBy()`

UnsetSortBy ensures that no value is present for SortBy, not even an explicit nil
### GetSortOrder

`func (o *TenantsPage) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *TenantsPage) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *TenantsPage) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *TenantsPage) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### SetSortOrderNil

`func (o *TenantsPage) SetSortOrderNil(b bool)`

 SetSortOrderNil sets the value for SortOrder to be an explicit nil

### UnsetSortOrder
`func (o *TenantsPage) UnsetSortOrder()`

UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
### GetContents

`func (o *TenantsPage) GetContents() []Tenant`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *TenantsPage) GetContentsOk() (*[]Tenant, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *TenantsPage) SetContents(v []Tenant)`

SetContents sets Contents field to given value.

### HasContents

`func (o *TenantsPage) HasContents() bool`

HasContents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


