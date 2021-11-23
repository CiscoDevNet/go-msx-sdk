# VulnerabilitiesPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **int32** |  | [readonly] 
**PageSize** | **int32** |  | [readonly] 
**TotalItems** | Pointer to **NullableInt64** |  | [optional] [readonly] 
**HasNext** | **NullableBool** |  | [readonly] 
**HasPrevious** | **NullableBool** |  | [readonly] 
**SortBy** | Pointer to **NullableString** |  | [optional] [readonly] 
**SortOrder** | Pointer to **NullableString** |  | [optional] [readonly] 
**Contents** | [**[]Vulnerability**](Vulnerability.md) |  | 

## Methods

### NewVulnerabilitiesPage

`func NewVulnerabilitiesPage(page int32, pageSize int32, hasNext NullableBool, hasPrevious NullableBool, contents []Vulnerability, ) *VulnerabilitiesPage`

NewVulnerabilitiesPage instantiates a new VulnerabilitiesPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVulnerabilitiesPageWithDefaults

`func NewVulnerabilitiesPageWithDefaults() *VulnerabilitiesPage`

NewVulnerabilitiesPageWithDefaults instantiates a new VulnerabilitiesPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *VulnerabilitiesPage) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *VulnerabilitiesPage) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *VulnerabilitiesPage) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *VulnerabilitiesPage) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *VulnerabilitiesPage) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *VulnerabilitiesPage) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetTotalItems

`func (o *VulnerabilitiesPage) GetTotalItems() int64`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *VulnerabilitiesPage) GetTotalItemsOk() (*int64, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *VulnerabilitiesPage) SetTotalItems(v int64)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *VulnerabilitiesPage) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### SetTotalItemsNil

`func (o *VulnerabilitiesPage) SetTotalItemsNil(b bool)`

 SetTotalItemsNil sets the value for TotalItems to be an explicit nil

### UnsetTotalItems
`func (o *VulnerabilitiesPage) UnsetTotalItems()`

UnsetTotalItems ensures that no value is present for TotalItems, not even an explicit nil
### GetHasNext

`func (o *VulnerabilitiesPage) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *VulnerabilitiesPage) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *VulnerabilitiesPage) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.


### SetHasNextNil

`func (o *VulnerabilitiesPage) SetHasNextNil(b bool)`

 SetHasNextNil sets the value for HasNext to be an explicit nil

### UnsetHasNext
`func (o *VulnerabilitiesPage) UnsetHasNext()`

UnsetHasNext ensures that no value is present for HasNext, not even an explicit nil
### GetHasPrevious

`func (o *VulnerabilitiesPage) GetHasPrevious() bool`

GetHasPrevious returns the HasPrevious field if non-nil, zero value otherwise.

### GetHasPreviousOk

`func (o *VulnerabilitiesPage) GetHasPreviousOk() (*bool, bool)`

GetHasPreviousOk returns a tuple with the HasPrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrevious

`func (o *VulnerabilitiesPage) SetHasPrevious(v bool)`

SetHasPrevious sets HasPrevious field to given value.


### SetHasPreviousNil

`func (o *VulnerabilitiesPage) SetHasPreviousNil(b bool)`

 SetHasPreviousNil sets the value for HasPrevious to be an explicit nil

### UnsetHasPrevious
`func (o *VulnerabilitiesPage) UnsetHasPrevious()`

UnsetHasPrevious ensures that no value is present for HasPrevious, not even an explicit nil
### GetSortBy

`func (o *VulnerabilitiesPage) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *VulnerabilitiesPage) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *VulnerabilitiesPage) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *VulnerabilitiesPage) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### SetSortByNil

`func (o *VulnerabilitiesPage) SetSortByNil(b bool)`

 SetSortByNil sets the value for SortBy to be an explicit nil

### UnsetSortBy
`func (o *VulnerabilitiesPage) UnsetSortBy()`

UnsetSortBy ensures that no value is present for SortBy, not even an explicit nil
### GetSortOrder

`func (o *VulnerabilitiesPage) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *VulnerabilitiesPage) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *VulnerabilitiesPage) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *VulnerabilitiesPage) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### SetSortOrderNil

`func (o *VulnerabilitiesPage) SetSortOrderNil(b bool)`

 SetSortOrderNil sets the value for SortOrder to be an explicit nil

### UnsetSortOrder
`func (o *VulnerabilitiesPage) UnsetSortOrder()`

UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
### GetContents

`func (o *VulnerabilitiesPage) GetContents() []Vulnerability`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *VulnerabilitiesPage) GetContentsOk() (*[]Vulnerability, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *VulnerabilitiesPage) SetContents(v []Vulnerability)`

SetContents sets Contents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


