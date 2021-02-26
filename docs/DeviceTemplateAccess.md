# DeviceTemplateAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantIds** | Pointer to **[]string** | List of tenants to grant access. Mutually exclusive from global flag field. | [optional] 
**Global** | Pointer to **bool** | Determines if the template is globally accessible. Mutually exclusive from tenant list field. | [optional] 

## Methods

### NewDeviceTemplateAccess

`func NewDeviceTemplateAccess() *DeviceTemplateAccess`

NewDeviceTemplateAccess instantiates a new DeviceTemplateAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceTemplateAccessWithDefaults

`func NewDeviceTemplateAccessWithDefaults() *DeviceTemplateAccess`

NewDeviceTemplateAccessWithDefaults instantiates a new DeviceTemplateAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantIds

`func (o *DeviceTemplateAccess) GetTenantIds() []string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *DeviceTemplateAccess) GetTenantIdsOk() (*[]string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *DeviceTemplateAccess) SetTenantIds(v []string)`

SetTenantIds sets TenantIds field to given value.

### HasTenantIds

`func (o *DeviceTemplateAccess) HasTenantIds() bool`

HasTenantIds returns a boolean if a field has been set.

### GetGlobal

`func (o *DeviceTemplateAccess) GetGlobal() bool`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *DeviceTemplateAccess) GetGlobalOk() (*bool, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *DeviceTemplateAccess) SetGlobal(v bool)`

SetGlobal sets Global field to given value.

### HasGlobal

`func (o *DeviceTemplateAccess) HasGlobal() bool`

HasGlobal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


