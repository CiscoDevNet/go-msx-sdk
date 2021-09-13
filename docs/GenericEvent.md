# GenericEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**Severity** | Pointer to [**GenericEventSeverity**](GenericEventSeverity.md) |  | [optional] 
**Subtype** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** | Service that generated the event | [optional] 
**Keywords** | Pointer to **string** | Space separated list of keywords to be used for search | [optional] 
**Details** | Pointer to **map[string]string** |  | [optional] 
**Trace** | Pointer to [**GenericEventTrace**](GenericEventTrace.md) |  | [optional] 
**Security** | Pointer to [**GenericEventSecurity**](GenericEventSecurity.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewGenericEvent

`func NewGenericEvent() *GenericEvent`

NewGenericEvent instantiates a new GenericEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericEventWithDefaults

`func NewGenericEventWithDefaults() *GenericEvent`

NewGenericEventWithDefaults instantiates a new GenericEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GenericEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GenericEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GenericEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GenericEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTime

`func (o *GenericEvent) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *GenericEvent) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *GenericEvent) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *GenericEvent) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetSeverity

`func (o *GenericEvent) GetSeverity() GenericEventSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GenericEvent) GetSeverityOk() (*GenericEventSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GenericEvent) SetSeverity(v GenericEventSeverity)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *GenericEvent) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSubtype

`func (o *GenericEvent) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *GenericEvent) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *GenericEvent) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *GenericEvent) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetService

`func (o *GenericEvent) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *GenericEvent) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *GenericEvent) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *GenericEvent) HasService() bool`

HasService returns a boolean if a field has been set.

### GetKeywords

`func (o *GenericEvent) GetKeywords() string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *GenericEvent) GetKeywordsOk() (*string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *GenericEvent) SetKeywords(v string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *GenericEvent) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetDetails

`func (o *GenericEvent) GetDetails() map[string]string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GenericEvent) GetDetailsOk() (*map[string]string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GenericEvent) SetDetails(v map[string]string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GenericEvent) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetTrace

`func (o *GenericEvent) GetTrace() GenericEventTrace`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *GenericEvent) GetTraceOk() (*GenericEventTrace, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *GenericEvent) SetTrace(v GenericEventTrace)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *GenericEvent) HasTrace() bool`

HasTrace returns a boolean if a field has been set.

### GetSecurity

`func (o *GenericEvent) GetSecurity() GenericEventSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *GenericEvent) GetSecurityOk() (*GenericEventSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *GenericEvent) SetSecurity(v GenericEventSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *GenericEvent) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetDescription

`func (o *GenericEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GenericEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GenericEvent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GenericEvent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


