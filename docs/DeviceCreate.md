# DeviceCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **string** |  | [optional] 
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

### NewDeviceCreate

`func NewDeviceCreate(tenantId string, name string, model string, type_ string, serialKey string, version string, managed bool, ) *DeviceCreate`

NewDeviceCreate instantiates a new DeviceCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceCreateWithDefaults

`func NewDeviceCreateWithDefaults() *DeviceCreate`

NewDeviceCreateWithDefaults instantiates a new DeviceCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *DeviceCreate) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DeviceCreate) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DeviceCreate) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *DeviceCreate) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *DeviceCreate) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DeviceCreate) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DeviceCreate) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *DeviceCreate) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTenantId

`func (o *DeviceCreate) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DeviceCreate) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DeviceCreate) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetName

`func (o *DeviceCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceCreate) SetName(v string)`

SetName sets Name field to given value.


### GetModel

`func (o *DeviceCreate) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DeviceCreate) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DeviceCreate) SetModel(v string)`

SetModel sets Model field to given value.


### GetType

`func (o *DeviceCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeviceCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeviceCreate) SetType(v string)`

SetType sets Type field to given value.


### GetSubType

`func (o *DeviceCreate) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *DeviceCreate) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *DeviceCreate) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *DeviceCreate) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetServiceType

`func (o *DeviceCreate) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *DeviceCreate) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *DeviceCreate) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *DeviceCreate) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetTags

`func (o *DeviceCreate) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeviceCreate) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DeviceCreate) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DeviceCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSerialKey

`func (o *DeviceCreate) GetSerialKey() string`

GetSerialKey returns the SerialKey field if non-nil, zero value otherwise.

### GetSerialKeyOk

`func (o *DeviceCreate) GetSerialKeyOk() (*string, bool)`

GetSerialKeyOk returns a tuple with the SerialKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialKey

`func (o *DeviceCreate) SetSerialKey(v string)`

SetSerialKey sets SerialKey field to given value.


### GetVersion

`func (o *DeviceCreate) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeviceCreate) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeviceCreate) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetManaged

`func (o *DeviceCreate) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *DeviceCreate) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *DeviceCreate) SetManaged(v bool)`

SetManaged sets Managed field to given value.


### GetOnboardType

`func (o *DeviceCreate) GetOnboardType() string`

GetOnboardType returns the OnboardType field if non-nil, zero value otherwise.

### GetOnboardTypeOk

`func (o *DeviceCreate) GetOnboardTypeOk() (*string, bool)`

GetOnboardTypeOk returns a tuple with the OnboardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardType

`func (o *DeviceCreate) SetOnboardType(v string)`

SetOnboardType sets OnboardType field to given value.

### HasOnboardType

`func (o *DeviceCreate) HasOnboardType() bool`

HasOnboardType returns a boolean if a field has been set.

### GetOnboardInformation

`func (o *DeviceCreate) GetOnboardInformation() map[string]string`

GetOnboardInformation returns the OnboardInformation field if non-nil, zero value otherwise.

### GetOnboardInformationOk

`func (o *DeviceCreate) GetOnboardInformationOk() (*map[string]string, bool)`

GetOnboardInformationOk returns a tuple with the OnboardInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardInformation

`func (o *DeviceCreate) SetOnboardInformation(v map[string]string)`

SetOnboardInformation sets OnboardInformation field to given value.

### HasOnboardInformation

`func (o *DeviceCreate) HasOnboardInformation() bool`

HasOnboardInformation returns a boolean if a field has been set.

### GetAttributes

`func (o *DeviceCreate) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *DeviceCreate) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *DeviceCreate) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *DeviceCreate) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetBilling

`func (o *DeviceCreate) GetBilling() DeviceBilling`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *DeviceCreate) GetBillingOk() (*DeviceBilling, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *DeviceCreate) SetBilling(v DeviceBilling)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *DeviceCreate) HasBilling() bool`

HasBilling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


