# LicenseDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseType** | Pointer to **string** | Type of license - TERM or DEMO or PERPETUAL | [optional] [readonly] 
**Quantity** | Pointer to **int64** | Number of reserved licenses | [optional] [readonly] 
**StartDate** | Pointer to **string** | License usage start date in yyyy-mm-dd format | [optional] [readonly] 
**EndDate** | Pointer to **string** | License usage expiration date in yyyy-mm-dd format | [optional] [readonly] 
**SubscriptionId** | Pointer to **string** | Subscription refence id | [optional] [readonly] 
**Status** | Pointer to **string** | Licencse usage status | [optional] [readonly] 

## Methods

### NewLicenseDetails

`func NewLicenseDetails() *LicenseDetails`

NewLicenseDetails instantiates a new LicenseDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseDetailsWithDefaults

`func NewLicenseDetailsWithDefaults() *LicenseDetails`

NewLicenseDetailsWithDefaults instantiates a new LicenseDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenseType

`func (o *LicenseDetails) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *LicenseDetails) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *LicenseDetails) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *LicenseDetails) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetQuantity

`func (o *LicenseDetails) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *LicenseDetails) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *LicenseDetails) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *LicenseDetails) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStartDate

`func (o *LicenseDetails) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *LicenseDetails) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *LicenseDetails) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *LicenseDetails) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *LicenseDetails) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *LicenseDetails) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *LicenseDetails) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *LicenseDetails) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *LicenseDetails) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *LicenseDetails) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *LicenseDetails) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *LicenseDetails) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetStatus

`func (o *LicenseDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LicenseDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LicenseDetails) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LicenseDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


