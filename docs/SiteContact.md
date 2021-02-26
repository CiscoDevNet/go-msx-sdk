# SiteContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Phone** | **string** |  | 
**Email** | Pointer to **string** |  | [optional] 

## Methods

### NewSiteContact

`func NewSiteContact(name string, phone string, ) *SiteContact`

NewSiteContact instantiates a new SiteContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteContactWithDefaults

`func NewSiteContactWithDefaults() *SiteContact`

NewSiteContactWithDefaults instantiates a new SiteContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SiteContact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiteContact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiteContact) SetName(v string)`

SetName sets Name field to given value.


### GetPhone

`func (o *SiteContact) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SiteContact) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SiteContact) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetEmail

`func (o *SiteContact) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SiteContact) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SiteContact) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SiteContact) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


