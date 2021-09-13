# GenericEventTrace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TraceId** | Pointer to **string** |  | [optional] 
**SpanId** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 

## Methods

### NewGenericEventTrace

`func NewGenericEventTrace() *GenericEventTrace`

NewGenericEventTrace instantiates a new GenericEventTrace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericEventTraceWithDefaults

`func NewGenericEventTraceWithDefaults() *GenericEventTrace`

NewGenericEventTraceWithDefaults instantiates a new GenericEventTrace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTraceId

`func (o *GenericEventTrace) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *GenericEventTrace) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *GenericEventTrace) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *GenericEventTrace) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetSpanId

`func (o *GenericEventTrace) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *GenericEventTrace) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *GenericEventTrace) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.

### HasSpanId

`func (o *GenericEventTrace) HasSpanId() bool`

HasSpanId returns a boolean if a field has been set.

### GetParentId

`func (o *GenericEventTrace) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *GenericEventTrace) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *GenericEventTrace) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *GenericEventTrace) HasParentId() bool`

HasParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


