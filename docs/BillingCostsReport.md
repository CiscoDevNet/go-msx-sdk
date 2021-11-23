# BillingCostsReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **string** |  | [optional] 
**Details** | Pointer to [**[]CostSummary**](CostSummary.md) |  | [optional] 

## Methods

### NewBillingCostsReport

`func NewBillingCostsReport() *BillingCostsReport`

NewBillingCostsReport instantiates a new BillingCostsReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingCostsReportWithDefaults

`func NewBillingCostsReportWithDefaults() *BillingCostsReport`

NewBillingCostsReportWithDefaults instantiates a new BillingCostsReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *BillingCostsReport) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BillingCostsReport) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BillingCostsReport) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *BillingCostsReport) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetDetails

`func (o *BillingCostsReport) GetDetails() []CostSummary`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *BillingCostsReport) GetDetailsOk() (*[]CostSummary, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *BillingCostsReport) SetDetails(v []CostSummary)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *BillingCostsReport) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


