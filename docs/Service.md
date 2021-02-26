# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ProviderId** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] [readonly] 
**ModifiedOn** | Pointer to **time.Time** |  | [optional] [readonly] 
**ProvisionedOn** | Pointer to **time.Time** |  | [optional] [readonly] 
**Status** | Pointer to **map[string]string** |  | [optional] 
**DefinitionAttributes** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Attributes** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewService

`func NewService() *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Service) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Service) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Service) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Service) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProviderId

`func (o *Service) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *Service) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *Service) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *Service) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetTenantId

`func (o *Service) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Service) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Service) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *Service) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUserId

`func (o *Service) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Service) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Service) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Service) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *Service) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *Service) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *Service) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *Service) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetCreatedOn

`func (o *Service) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Service) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Service) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *Service) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetModifiedOn

`func (o *Service) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *Service) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *Service) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *Service) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetProvisionedOn

`func (o *Service) GetProvisionedOn() time.Time`

GetProvisionedOn returns the ProvisionedOn field if non-nil, zero value otherwise.

### GetProvisionedOnOk

`func (o *Service) GetProvisionedOnOk() (*time.Time, bool)`

GetProvisionedOnOk returns a tuple with the ProvisionedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedOn

`func (o *Service) SetProvisionedOn(v time.Time)`

SetProvisionedOn sets ProvisionedOn field to given value.

### HasProvisionedOn

`func (o *Service) HasProvisionedOn() bool`

HasProvisionedOn returns a boolean if a field has been set.

### GetStatus

`func (o *Service) GetStatus() map[string]string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Service) GetStatusOk() (*map[string]string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Service) SetStatus(v map[string]string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Service) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDefinitionAttributes

`func (o *Service) GetDefinitionAttributes() map[string]map[string]interface{}`

GetDefinitionAttributes returns the DefinitionAttributes field if non-nil, zero value otherwise.

### GetDefinitionAttributesOk

`func (o *Service) GetDefinitionAttributesOk() (*map[string]map[string]interface{}, bool)`

GetDefinitionAttributesOk returns a tuple with the DefinitionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionAttributes

`func (o *Service) SetDefinitionAttributes(v map[string]map[string]interface{})`

SetDefinitionAttributes sets DefinitionAttributes field to given value.

### HasDefinitionAttributes

`func (o *Service) HasDefinitionAttributes() bool`

HasDefinitionAttributes returns a boolean if a field has been set.

### GetAttributes

`func (o *Service) GetAttributes() map[string]map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Service) GetAttributesOk() (*map[string]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Service) SetAttributes(v map[string]map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Service) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


