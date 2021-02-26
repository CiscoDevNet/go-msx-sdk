# LegacyServiceOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ServiceType** | Pointer to **string** |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**Schedule** | Pointer to [**LegacyScheduleConfig**](LegacyScheduleConfig.md) |  | [optional] 
**Payload** | Pointer to [**LegacyServiceOrderDetail**](LegacyServiceOrderDetail.md) |  | [optional] 
**TransactionUUID** | Pointer to **string** |  | [optional] 

## Methods

### NewLegacyServiceOrder

`func NewLegacyServiceOrder() *LegacyServiceOrder`

NewLegacyServiceOrder instantiates a new LegacyServiceOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegacyServiceOrderWithDefaults

`func NewLegacyServiceOrderWithDefaults() *LegacyServiceOrder`

NewLegacyServiceOrderWithDefaults instantiates a new LegacyServiceOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *LegacyServiceOrder) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *LegacyServiceOrder) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *LegacyServiceOrder) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *LegacyServiceOrder) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetType

`func (o *LegacyServiceOrder) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LegacyServiceOrder) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LegacyServiceOrder) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LegacyServiceOrder) HasType() bool`

HasType returns a boolean if a field has been set.

### GetServiceType

`func (o *LegacyServiceOrder) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *LegacyServiceOrder) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *LegacyServiceOrder) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *LegacyServiceOrder) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetAction

`func (o *LegacyServiceOrder) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *LegacyServiceOrder) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *LegacyServiceOrder) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *LegacyServiceOrder) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetSchedule

`func (o *LegacyServiceOrder) GetSchedule() LegacyScheduleConfig`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *LegacyServiceOrder) GetScheduleOk() (*LegacyScheduleConfig, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *LegacyServiceOrder) SetSchedule(v LegacyScheduleConfig)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *LegacyServiceOrder) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetPayload

`func (o *LegacyServiceOrder) GetPayload() LegacyServiceOrderDetail`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *LegacyServiceOrder) GetPayloadOk() (*LegacyServiceOrderDetail, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *LegacyServiceOrder) SetPayload(v LegacyServiceOrderDetail)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *LegacyServiceOrder) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetTransactionUUID

`func (o *LegacyServiceOrder) GetTransactionUUID() string`

GetTransactionUUID returns the TransactionUUID field if non-nil, zero value otherwise.

### GetTransactionUUIDOk

`func (o *LegacyServiceOrder) GetTransactionUUIDOk() (*string, bool)`

GetTransactionUUIDOk returns a tuple with the TransactionUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionUUID

`func (o *LegacyServiceOrder) SetTransactionUUID(v string)`

SetTransactionUUID sets TransactionUUID field to given value.

### HasTransactionUUID

`func (o *LegacyServiceOrder) HasTransactionUUID() bool`

HasTransactionUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


