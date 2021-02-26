# LegacyScheduleConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to **string** |  | [optional] 
**Relative** | Pointer to [**LegacyRelativeConfig**](LegacyRelativeConfig.md) |  | [optional] 
**Absolute** | Pointer to [**LegacyAbsoluteConfig**](LegacyAbsoluteConfig.md) |  | [optional] 

## Methods

### NewLegacyScheduleConfig

`func NewLegacyScheduleConfig() *LegacyScheduleConfig`

NewLegacyScheduleConfig instantiates a new LegacyScheduleConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegacyScheduleConfigWithDefaults

`func NewLegacyScheduleConfigWithDefaults() *LegacyScheduleConfig`

NewLegacyScheduleConfigWithDefaults instantiates a new LegacyScheduleConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *LegacyScheduleConfig) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *LegacyScheduleConfig) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *LegacyScheduleConfig) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *LegacyScheduleConfig) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetRelative

`func (o *LegacyScheduleConfig) GetRelative() LegacyRelativeConfig`

GetRelative returns the Relative field if non-nil, zero value otherwise.

### GetRelativeOk

`func (o *LegacyScheduleConfig) GetRelativeOk() (*LegacyRelativeConfig, bool)`

GetRelativeOk returns a tuple with the Relative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelative

`func (o *LegacyScheduleConfig) SetRelative(v LegacyRelativeConfig)`

SetRelative sets Relative field to given value.

### HasRelative

`func (o *LegacyScheduleConfig) HasRelative() bool`

HasRelative returns a boolean if a field has been set.

### GetAbsolute

`func (o *LegacyScheduleConfig) GetAbsolute() LegacyAbsoluteConfig`

GetAbsolute returns the Absolute field if non-nil, zero value otherwise.

### GetAbsoluteOk

`func (o *LegacyScheduleConfig) GetAbsoluteOk() (*LegacyAbsoluteConfig, bool)`

GetAbsoluteOk returns a tuple with the Absolute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsolute

`func (o *LegacyScheduleConfig) SetAbsolute(v LegacyAbsoluteConfig)`

SetAbsolute sets Absolute field to given value.

### HasAbsolute

`func (o *LegacyScheduleConfig) HasAbsolute() bool`

HasAbsolute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


