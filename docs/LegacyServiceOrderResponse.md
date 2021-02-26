# LegacyServiceOrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**TransactionUUID** | Pointer to **string** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ServiceInstanceIds** | Pointer to **[]string** |  | [optional] 
**Sites** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Devices** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewLegacyServiceOrderResponse

`func NewLegacyServiceOrderResponse() *LegacyServiceOrderResponse`

NewLegacyServiceOrderResponse instantiates a new LegacyServiceOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegacyServiceOrderResponseWithDefaults

`func NewLegacyServiceOrderResponseWithDefaults() *LegacyServiceOrderResponse`

NewLegacyServiceOrderResponseWithDefaults instantiates a new LegacyServiceOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *LegacyServiceOrderResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LegacyServiceOrderResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LegacyServiceOrderResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *LegacyServiceOrderResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTransactionUUID

`func (o *LegacyServiceOrderResponse) GetTransactionUUID() string`

GetTransactionUUID returns the TransactionUUID field if non-nil, zero value otherwise.

### GetTransactionUUIDOk

`func (o *LegacyServiceOrderResponse) GetTransactionUUIDOk() (*string, bool)`

GetTransactionUUIDOk returns a tuple with the TransactionUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionUUID

`func (o *LegacyServiceOrderResponse) SetTransactionUUID(v string)`

SetTransactionUUID sets TransactionUUID field to given value.

### HasTransactionUUID

`func (o *LegacyServiceOrderResponse) HasTransactionUUID() bool`

HasTransactionUUID returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *LegacyServiceOrderResponse) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *LegacyServiceOrderResponse) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *LegacyServiceOrderResponse) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *LegacyServiceOrderResponse) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetUrl

`func (o *LegacyServiceOrderResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LegacyServiceOrderResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LegacyServiceOrderResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *LegacyServiceOrderResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStatus

`func (o *LegacyServiceOrderResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LegacyServiceOrderResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LegacyServiceOrderResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LegacyServiceOrderResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetServiceInstanceIds

`func (o *LegacyServiceOrderResponse) GetServiceInstanceIds() []string`

GetServiceInstanceIds returns the ServiceInstanceIds field if non-nil, zero value otherwise.

### GetServiceInstanceIdsOk

`func (o *LegacyServiceOrderResponse) GetServiceInstanceIdsOk() (*[]string, bool)`

GetServiceInstanceIdsOk returns a tuple with the ServiceInstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceInstanceIds

`func (o *LegacyServiceOrderResponse) SetServiceInstanceIds(v []string)`

SetServiceInstanceIds sets ServiceInstanceIds field to given value.

### HasServiceInstanceIds

`func (o *LegacyServiceOrderResponse) HasServiceInstanceIds() bool`

HasServiceInstanceIds returns a boolean if a field has been set.

### GetSites

`func (o *LegacyServiceOrderResponse) GetSites() []map[string]interface{}`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *LegacyServiceOrderResponse) GetSitesOk() (*[]map[string]interface{}, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *LegacyServiceOrderResponse) SetSites(v []map[string]interface{})`

SetSites sets Sites field to given value.

### HasSites

`func (o *LegacyServiceOrderResponse) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetDevices

`func (o *LegacyServiceOrderResponse) GetDevices() []map[string]interface{}`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *LegacyServiceOrderResponse) GetDevicesOk() (*[]map[string]interface{}, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *LegacyServiceOrderResponse) SetDevices(v []map[string]interface{})`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *LegacyServiceOrderResponse) HasDevices() bool`

HasDevices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


