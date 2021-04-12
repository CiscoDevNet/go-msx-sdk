# PageHeader

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

## Methods

### NewPageHeader

`func NewPageHeader() *PageHeader`

NewPageHeader instantiates a new PageHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageHeaderWithDefaults

`func NewPageHeaderWithDefaults() *PageHeader`

NewPageHeaderWithDefaults instantiates a new PageHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *PageHeader) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PageHeader) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PageHeader) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *PageHeader) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *PageHeader) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PageHeader) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PageHeader) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *PageHeader) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalItems

`func (o *PageHeader) GetTotalItems() int64`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *PageHeader) GetTotalItemsOk() (*int64, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *PageHeader) SetTotalItems(v int64)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *PageHeader) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### SetTotalItemsNil

`func (o *PageHeader) SetTotalItemsNil(b bool)`

 SetTotalItemsNil sets the value for TotalItems to be an explicit nil

### UnsetTotalItems
`func (o *PageHeader) UnsetTotalItems()`

UnsetTotalItems ensures that no value is present for TotalItems, not even an explicit nil
### GetHasNext

`func (o *PageHeader) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *PageHeader) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *PageHeader) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *PageHeader) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.

### SetHasNextNil

`func (o *PageHeader) SetHasNextNil(b bool)`

 SetHasNextNil sets the value for HasNext to be an explicit nil

### UnsetHasNext
`func (o *PageHeader) UnsetHasNext()`

UnsetHasNext ensures that no value is present for HasNext, not even an explicit nil
### GetHasPrevious

`func (o *PageHeader) GetHasPrevious() bool`

GetHasPrevious returns the HasPrevious field if non-nil, zero value otherwise.

### GetHasPreviousOk

`func (o *PageHeader) GetHasPreviousOk() (*bool, bool)`

GetHasPreviousOk returns a tuple with the HasPrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrevious

`func (o *PageHeader) SetHasPrevious(v bool)`

SetHasPrevious sets HasPrevious field to given value.

### HasHasPrevious

`func (o *PageHeader) HasHasPrevious() bool`

HasHasPrevious returns a boolean if a field has been set.

### SetHasPreviousNil

`func (o *PageHeader) SetHasPreviousNil(b bool)`

 SetHasPreviousNil sets the value for HasPrevious to be an explicit nil

### UnsetHasPrevious
`func (o *PageHeader) UnsetHasPrevious()`

UnsetHasPrevious ensures that no value is present for HasPrevious, not even an explicit nil
### GetSortBy

`func (o *PageHeader) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *PageHeader) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *PageHeader) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *PageHeader) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### SetSortByNil

`func (o *PageHeader) SetSortByNil(b bool)`

 SetSortByNil sets the value for SortBy to be an explicit nil

### UnsetSortBy
`func (o *PageHeader) UnsetSortBy()`

UnsetSortBy ensures that no value is present for SortBy, not even an explicit nil
### GetSortOrder

`func (o *PageHeader) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *PageHeader) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *PageHeader) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *PageHeader) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### SetSortOrderNil

`func (o *PageHeader) SetSortOrderNil(b bool)`

 SetSortOrderNil sets the value for SortOrder to be an explicit nil

### UnsetSortOrder
`func (o *PageHeader) UnsetSortOrder()`

UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


