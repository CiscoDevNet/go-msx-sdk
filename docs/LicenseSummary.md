# LicenseSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entitled** | Pointer to **int64** | Total entitled quantity for the license | [optional] [readonly] 
**Inuse** | Pointer to **int64** | Total consumed quantity for the device | [optional] [readonly] 
**Reserved** | Pointer to **int64** | Reserved quantity (if any) | [optional] [readonly] 
**Status** | Pointer to **string** | Current compliance status for the license | [optional] [readonly] 
**Name** | Pointer to **string** | User friendly display name for the license | [optional] [readonly] 
**Enforced** | Pointer to **bool** | Identifies if the tag is for an enforced license or not | [optional] [readonly] 
**LicenseDetails** | Pointer to [**[]LicenseDetails**](LicenseDetails.md) |  | [optional] 

## Methods

### NewLicenseSummary

`func NewLicenseSummary() *LicenseSummary`

NewLicenseSummary instantiates a new LicenseSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseSummaryWithDefaults

`func NewLicenseSummaryWithDefaults() *LicenseSummary`

NewLicenseSummaryWithDefaults instantiates a new LicenseSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitled

`func (o *LicenseSummary) GetEntitled() int64`

GetEntitled returns the Entitled field if non-nil, zero value otherwise.

### GetEntitledOk

`func (o *LicenseSummary) GetEntitledOk() (*int64, bool)`

GetEntitledOk returns a tuple with the Entitled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitled

`func (o *LicenseSummary) SetEntitled(v int64)`

SetEntitled sets Entitled field to given value.

### HasEntitled

`func (o *LicenseSummary) HasEntitled() bool`

HasEntitled returns a boolean if a field has been set.

### GetInuse

`func (o *LicenseSummary) GetInuse() int64`

GetInuse returns the Inuse field if non-nil, zero value otherwise.

### GetInuseOk

`func (o *LicenseSummary) GetInuseOk() (*int64, bool)`

GetInuseOk returns a tuple with the Inuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInuse

`func (o *LicenseSummary) SetInuse(v int64)`

SetInuse sets Inuse field to given value.

### HasInuse

`func (o *LicenseSummary) HasInuse() bool`

HasInuse returns a boolean if a field has been set.

### GetReserved

`func (o *LicenseSummary) GetReserved() int64`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *LicenseSummary) GetReservedOk() (*int64, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *LicenseSummary) SetReserved(v int64)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *LicenseSummary) HasReserved() bool`

HasReserved returns a boolean if a field has been set.

### GetStatus

`func (o *LicenseSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LicenseSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LicenseSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LicenseSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetName

`func (o *LicenseSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LicenseSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LicenseSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LicenseSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnforced

`func (o *LicenseSummary) GetEnforced() bool`

GetEnforced returns the Enforced field if non-nil, zero value otherwise.

### GetEnforcedOk

`func (o *LicenseSummary) GetEnforcedOk() (*bool, bool)`

GetEnforcedOk returns a tuple with the Enforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforced

`func (o *LicenseSummary) SetEnforced(v bool)`

SetEnforced sets Enforced field to given value.

### HasEnforced

`func (o *LicenseSummary) HasEnforced() bool`

HasEnforced returns a boolean if a field has been set.

### GetLicenseDetails

`func (o *LicenseSummary) GetLicenseDetails() []LicenseDetails`

GetLicenseDetails returns the LicenseDetails field if non-nil, zero value otherwise.

### GetLicenseDetailsOk

`func (o *LicenseSummary) GetLicenseDetailsOk() (*[]LicenseDetails, bool)`

GetLicenseDetailsOk returns a tuple with the LicenseDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseDetails

`func (o *LicenseSummary) SetLicenseDetails(v []LicenseDetails)`

SetLicenseDetails sets LicenseDetails field to given value.

### HasLicenseDetails

`func (o *LicenseSummary) HasLicenseDetails() bool`

HasLicenseDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


