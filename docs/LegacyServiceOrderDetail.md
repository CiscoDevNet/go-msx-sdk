# LegacyServiceOrderDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to **map[string]string** |  | [optional] 
**Tenant** | Pointer to **map[string]string** |  | [optional] 
**User** | Pointer to **map[string]string** |  | [optional] 
**Provider** | Pointer to **map[string]string** |  | [optional] 
**Offer** | Pointer to **map[string]string** |  | [optional] 
**Cost** | Pointer to **map[string]string** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**SubscriptionName** | Pointer to **string** |  | [optional] 
**SubscriptionDetail** | Pointer to [**LegacySubscriptionDetail**](LegacySubscriptionDetail.md) |  | [optional] 
**ServiceDowngrade** | Pointer to **map[string]interface{}** |  | [optional] 
**NsoResponseTypes** | Pointer to [**LegacyNsoResponseTypes**](LegacyNsoResponseTypes.md) |  | [optional] 

## Methods

### NewLegacyServiceOrderDetail

`func NewLegacyServiceOrderDetail() *LegacyServiceOrderDetail`

NewLegacyServiceOrderDetail instantiates a new LegacyServiceOrderDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegacyServiceOrderDetailWithDefaults

`func NewLegacyServiceOrderDetailWithDefaults() *LegacyServiceOrderDetail`

NewLegacyServiceOrderDetailWithDefaults instantiates a new LegacyServiceOrderDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *LegacyServiceOrderDetail) GetService() map[string]string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *LegacyServiceOrderDetail) GetServiceOk() (*map[string]string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *LegacyServiceOrderDetail) SetService(v map[string]string)`

SetService sets Service field to given value.

### HasService

`func (o *LegacyServiceOrderDetail) HasService() bool`

HasService returns a boolean if a field has been set.

### GetTenant

`func (o *LegacyServiceOrderDetail) GetTenant() map[string]string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *LegacyServiceOrderDetail) GetTenantOk() (*map[string]string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *LegacyServiceOrderDetail) SetTenant(v map[string]string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *LegacyServiceOrderDetail) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetUser

`func (o *LegacyServiceOrderDetail) GetUser() map[string]string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LegacyServiceOrderDetail) GetUserOk() (*map[string]string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LegacyServiceOrderDetail) SetUser(v map[string]string)`

SetUser sets User field to given value.

### HasUser

`func (o *LegacyServiceOrderDetail) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetProvider

`func (o *LegacyServiceOrderDetail) GetProvider() map[string]string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *LegacyServiceOrderDetail) GetProviderOk() (*map[string]string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *LegacyServiceOrderDetail) SetProvider(v map[string]string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *LegacyServiceOrderDetail) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetOffer

`func (o *LegacyServiceOrderDetail) GetOffer() map[string]string`

GetOffer returns the Offer field if non-nil, zero value otherwise.

### GetOfferOk

`func (o *LegacyServiceOrderDetail) GetOfferOk() (*map[string]string, bool)`

GetOfferOk returns a tuple with the Offer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffer

`func (o *LegacyServiceOrderDetail) SetOffer(v map[string]string)`

SetOffer sets Offer field to given value.

### HasOffer

`func (o *LegacyServiceOrderDetail) HasOffer() bool`

HasOffer returns a boolean if a field has been set.

### GetCost

`func (o *LegacyServiceOrderDetail) GetCost() map[string]string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *LegacyServiceOrderDetail) GetCostOk() (*map[string]string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *LegacyServiceOrderDetail) SetCost(v map[string]string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *LegacyServiceOrderDetail) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *LegacyServiceOrderDetail) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *LegacyServiceOrderDetail) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *LegacyServiceOrderDetail) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *LegacyServiceOrderDetail) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSubscriptionName

`func (o *LegacyServiceOrderDetail) GetSubscriptionName() string`

GetSubscriptionName returns the SubscriptionName field if non-nil, zero value otherwise.

### GetSubscriptionNameOk

`func (o *LegacyServiceOrderDetail) GetSubscriptionNameOk() (*string, bool)`

GetSubscriptionNameOk returns a tuple with the SubscriptionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionName

`func (o *LegacyServiceOrderDetail) SetSubscriptionName(v string)`

SetSubscriptionName sets SubscriptionName field to given value.

### HasSubscriptionName

`func (o *LegacyServiceOrderDetail) HasSubscriptionName() bool`

HasSubscriptionName returns a boolean if a field has been set.

### GetSubscriptionDetail

`func (o *LegacyServiceOrderDetail) GetSubscriptionDetail() LegacySubscriptionDetail`

GetSubscriptionDetail returns the SubscriptionDetail field if non-nil, zero value otherwise.

### GetSubscriptionDetailOk

`func (o *LegacyServiceOrderDetail) GetSubscriptionDetailOk() (*LegacySubscriptionDetail, bool)`

GetSubscriptionDetailOk returns a tuple with the SubscriptionDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionDetail

`func (o *LegacyServiceOrderDetail) SetSubscriptionDetail(v LegacySubscriptionDetail)`

SetSubscriptionDetail sets SubscriptionDetail field to given value.

### HasSubscriptionDetail

`func (o *LegacyServiceOrderDetail) HasSubscriptionDetail() bool`

HasSubscriptionDetail returns a boolean if a field has been set.

### GetServiceDowngrade

`func (o *LegacyServiceOrderDetail) GetServiceDowngrade() map[string]interface{}`

GetServiceDowngrade returns the ServiceDowngrade field if non-nil, zero value otherwise.

### GetServiceDowngradeOk

`func (o *LegacyServiceOrderDetail) GetServiceDowngradeOk() (*map[string]interface{}, bool)`

GetServiceDowngradeOk returns a tuple with the ServiceDowngrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDowngrade

`func (o *LegacyServiceOrderDetail) SetServiceDowngrade(v map[string]interface{})`

SetServiceDowngrade sets ServiceDowngrade field to given value.

### HasServiceDowngrade

`func (o *LegacyServiceOrderDetail) HasServiceDowngrade() bool`

HasServiceDowngrade returns a boolean if a field has been set.

### GetNsoResponseTypes

`func (o *LegacyServiceOrderDetail) GetNsoResponseTypes() LegacyNsoResponseTypes`

GetNsoResponseTypes returns the NsoResponseTypes field if non-nil, zero value otherwise.

### GetNsoResponseTypesOk

`func (o *LegacyServiceOrderDetail) GetNsoResponseTypesOk() (*LegacyNsoResponseTypes, bool)`

GetNsoResponseTypesOk returns a tuple with the NsoResponseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsoResponseTypes

`func (o *LegacyServiceOrderDetail) SetNsoResponseTypes(v LegacyNsoResponseTypes)`

SetNsoResponseTypes sets NsoResponseTypes field to given value.

### HasNsoResponseTypes

`func (o *LegacyServiceOrderDetail) HasNsoResponseTypes() bool`

HasNsoResponseTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


