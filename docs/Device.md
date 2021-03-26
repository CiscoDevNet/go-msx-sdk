# Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**ProviderId** | Pointer to **string** |  | [optional] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] 
**ModifiedOn** | Pointer to **time.Time** |  | [optional] 
**ServiceInstanceId** | Pointer to **string** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**TenantId** | **string** |  | 
**Name** | **string** |  | 
**Model** | **string** |  | 
**Type** | **string** |  | 
**SubType** | Pointer to **string** |  | [optional] 
**ServiceType** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 
**SerialKey** | **string** |  | 
**Version** | **string** |  | 
**Managed** | **bool** |  | [default to false]
**OnboardType** | Pointer to **string** |  | [optional] 
**OnboardInformation** | Pointer to **map[string]string** |  | [optional] 
**Attributes** | Pointer to **map[string]string** |  | [optional] 
**Billing** | Pointer to [**DeviceBilling**](DeviceBilling.md) |  | [optional] 

## Methods

### NewDevice

`func NewDevice(tenantId string, name string, model string, type_ string, serialKey string, version string, managed bool, ) *Device`

NewDevice instantiates a new Device object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithDefaults

`func NewDeviceWithDefaults() *Device`

NewDeviceWithDefaults instantiates a new Device object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Device) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Device) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Device) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Device) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *Device) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Device) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Device) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Device) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetProviderId

`func (o *Device) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *Device) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *Device) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *Device) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetCreatedOn

`func (o *Device) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Device) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Device) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *Device) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetModifiedOn

`func (o *Device) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *Device) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *Device) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *Device) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetServiceInstanceId

`func (o *Device) GetServiceInstanceId() string`

GetServiceInstanceId returns the ServiceInstanceId field if non-nil, zero value otherwise.

### GetServiceInstanceIdOk

`func (o *Device) GetServiceInstanceIdOk() (*string, bool)`

GetServiceInstanceIdOk returns a tuple with the ServiceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceInstanceId

`func (o *Device) SetServiceInstanceId(v string)`

SetServiceInstanceId sets ServiceInstanceId field to given value.

### HasServiceInstanceId

`func (o *Device) HasServiceInstanceId() bool`

HasServiceInstanceId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *Device) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *Device) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *Device) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *Device) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTenantId

`func (o *Device) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Device) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Device) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetName

`func (o *Device) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Device) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Device) SetName(v string)`

SetName sets Name field to given value.


### GetModel

`func (o *Device) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Device) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Device) SetModel(v string)`

SetModel sets Model field to given value.


### GetType

`func (o *Device) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Device) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Device) SetType(v string)`

SetType sets Type field to given value.


### GetSubType

`func (o *Device) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *Device) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *Device) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *Device) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetServiceType

`func (o *Device) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *Device) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *Device) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *Device) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetTags

`func (o *Device) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Device) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Device) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Device) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSerialKey

`func (o *Device) GetSerialKey() string`

GetSerialKey returns the SerialKey field if non-nil, zero value otherwise.

### GetSerialKeyOk

`func (o *Device) GetSerialKeyOk() (*string, bool)`

GetSerialKeyOk returns a tuple with the SerialKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialKey

`func (o *Device) SetSerialKey(v string)`

SetSerialKey sets SerialKey field to given value.


### GetVersion

`func (o *Device) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Device) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Device) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetManaged

`func (o *Device) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *Device) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *Device) SetManaged(v bool)`

SetManaged sets Managed field to given value.


### GetOnboardType

`func (o *Device) GetOnboardType() string`

GetOnboardType returns the OnboardType field if non-nil, zero value otherwise.

### GetOnboardTypeOk

`func (o *Device) GetOnboardTypeOk() (*string, bool)`

GetOnboardTypeOk returns a tuple with the OnboardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardType

`func (o *Device) SetOnboardType(v string)`

SetOnboardType sets OnboardType field to given value.

### HasOnboardType

`func (o *Device) HasOnboardType() bool`

HasOnboardType returns a boolean if a field has been set.

### GetOnboardInformation

`func (o *Device) GetOnboardInformation() map[string]string`

GetOnboardInformation returns the OnboardInformation field if non-nil, zero value otherwise.

### GetOnboardInformationOk

`func (o *Device) GetOnboardInformationOk() (*map[string]string, bool)`

GetOnboardInformationOk returns a tuple with the OnboardInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardInformation

`func (o *Device) SetOnboardInformation(v map[string]string)`

SetOnboardInformation sets OnboardInformation field to given value.

### HasOnboardInformation

`func (o *Device) HasOnboardInformation() bool`

HasOnboardInformation returns a boolean if a field has been set.

### GetAttributes

`func (o *Device) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Device) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Device) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Device) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetBilling

`func (o *Device) GetBilling() DeviceBilling`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *Device) GetBillingOk() (*DeviceBilling, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *Device) SetBilling(v DeviceBilling)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *Device) HasBilling() bool`

HasBilling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


