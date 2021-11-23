# TemplateApplicationCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | 
**TargetId** | **string** |  | 
**TargetType** | **string** |  | 
**Parameters** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTemplateApplicationCreate

`func NewTemplateApplicationCreate(tenantId string, targetId string, targetType string, ) *TemplateApplicationCreate`

NewTemplateApplicationCreate instantiates a new TemplateApplicationCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateApplicationCreateWithDefaults

`func NewTemplateApplicationCreateWithDefaults() *TemplateApplicationCreate`

NewTemplateApplicationCreateWithDefaults instantiates a new TemplateApplicationCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *TemplateApplicationCreate) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TemplateApplicationCreate) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TemplateApplicationCreate) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetTargetId

`func (o *TemplateApplicationCreate) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *TemplateApplicationCreate) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *TemplateApplicationCreate) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.


### GetTargetType

`func (o *TemplateApplicationCreate) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *TemplateApplicationCreate) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *TemplateApplicationCreate) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.


### GetParameters

`func (o *TemplateApplicationCreate) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *TemplateApplicationCreate) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *TemplateApplicationCreate) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *TemplateApplicationCreate) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


