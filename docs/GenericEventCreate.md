# GenericEventCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Severity** | Pointer to [**GenericEventSeverity**](GenericEventSeverity.md) |  | [optional] 
**Subtype** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** | Service that generated the event | [optional] 
**Keywords** | Pointer to **string** | Space separated list of keywords to be used for search | [optional] 
**Details** | Pointer to **map[string]string** |  | [optional] 
**Trace** | Pointer to [**GenericEventTrace**](GenericEventTrace.md) |  | [optional] 
**Security** | Pointer to [**GenericEventSecurity**](GenericEventSecurity.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewGenericEventCreate

`func NewGenericEventCreate() *GenericEventCreate`

NewGenericEventCreate instantiates a new GenericEventCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericEventCreateWithDefaults

`func NewGenericEventCreateWithDefaults() *GenericEventCreate`

NewGenericEventCreateWithDefaults instantiates a new GenericEventCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeverity

`func (o *GenericEventCreate) GetSeverity() GenericEventSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GenericEventCreate) GetSeverityOk() (*GenericEventSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GenericEventCreate) SetSeverity(v GenericEventSeverity)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *GenericEventCreate) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSubtype

`func (o *GenericEventCreate) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *GenericEventCreate) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *GenericEventCreate) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *GenericEventCreate) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetService

`func (o *GenericEventCreate) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *GenericEventCreate) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *GenericEventCreate) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *GenericEventCreate) HasService() bool`

HasService returns a boolean if a field has been set.

### GetKeywords

`func (o *GenericEventCreate) GetKeywords() string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *GenericEventCreate) GetKeywordsOk() (*string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *GenericEventCreate) SetKeywords(v string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *GenericEventCreate) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetDetails

`func (o *GenericEventCreate) GetDetails() map[string]string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GenericEventCreate) GetDetailsOk() (*map[string]string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GenericEventCreate) SetDetails(v map[string]string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GenericEventCreate) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetTrace

`func (o *GenericEventCreate) GetTrace() GenericEventTrace`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *GenericEventCreate) GetTraceOk() (*GenericEventTrace, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *GenericEventCreate) SetTrace(v GenericEventTrace)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *GenericEventCreate) HasTrace() bool`

HasTrace returns a boolean if a field has been set.

### GetSecurity

`func (o *GenericEventCreate) GetSecurity() GenericEventSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *GenericEventCreate) GetSecurityOk() (*GenericEventSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *GenericEventCreate) SetSecurity(v GenericEventSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *GenericEventCreate) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetDescription

`func (o *GenericEventCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GenericEventCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GenericEventCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GenericEventCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


