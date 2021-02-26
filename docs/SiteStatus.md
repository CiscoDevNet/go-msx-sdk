# SiteStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | [default to "unknown"]
**Name** | **string** |  | [default to "unknown"]
**Value** | **string** |  | [default to "unknown"]
**Severity** | **string** |  | 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**LastUpdatedMessage** | **string** |  | 

## Methods

### NewSiteStatus

`func NewSiteStatus(type_ string, name string, value string, severity string, lastUpdatedMessage string, ) *SiteStatus`

NewSiteStatus instantiates a new SiteStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteStatusWithDefaults

`func NewSiteStatusWithDefaults() *SiteStatus`

NewSiteStatusWithDefaults instantiates a new SiteStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SiteStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SiteStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SiteStatus) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *SiteStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiteStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiteStatus) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *SiteStatus) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SiteStatus) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SiteStatus) SetValue(v string)`

SetValue sets Value field to given value.


### GetSeverity

`func (o *SiteStatus) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *SiteStatus) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *SiteStatus) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetLastUpdated

`func (o *SiteStatus) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SiteStatus) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SiteStatus) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *SiteStatus) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastUpdatedMessage

`func (o *SiteStatus) GetLastUpdatedMessage() string`

GetLastUpdatedMessage returns the LastUpdatedMessage field if non-nil, zero value otherwise.

### GetLastUpdatedMessageOk

`func (o *SiteStatus) GetLastUpdatedMessageOk() (*string, bool)`

GetLastUpdatedMessageOk returns a tuple with the LastUpdatedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedMessage

`func (o *SiteStatus) SetLastUpdatedMessage(v string)`

SetLastUpdatedMessage sets LastUpdatedMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


