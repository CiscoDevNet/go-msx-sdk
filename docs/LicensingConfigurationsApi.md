# \LicensingConfigurationsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSmartAccountConfiguration**](LicensingConfigurationsApi.md#CreateSmartAccountConfiguration) | **Post** /licensing/api/v8/configuration | Creates a smart account configuration.
[**GetSmartAccountConfiguration**](LicensingConfigurationsApi.md#GetSmartAccountConfiguration) | **Get** /licensing/api/v8/configuration | Returns a smart account configuration.
[**UpdateSmartAccountConfiguration**](LicensingConfigurationsApi.md#UpdateSmartAccountConfiguration) | **Put** /licensing/api/v8/configuration | Updates a smart account configuration.



## CreateSmartAccountConfiguration

> SmartAccountConfiguration CreateSmartAccountConfiguration(ctx).SmartAccountConfigurationCreate(smartAccountConfigurationCreate).Execute()

Creates a smart account configuration.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    smartAccountConfigurationCreate := *openapiclient.NewSmartAccountConfigurationCreate() // SmartAccountConfigurationCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensingConfigurationsApi.CreateSmartAccountConfiguration(context.Background()).SmartAccountConfigurationCreate(smartAccountConfigurationCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensingConfigurationsApi.CreateSmartAccountConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSmartAccountConfiguration`: SmartAccountConfiguration
    fmt.Fprintf(os.Stdout, "Response from `LicensingConfigurationsApi.CreateSmartAccountConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSmartAccountConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smartAccountConfigurationCreate** | [**SmartAccountConfigurationCreate**](SmartAccountConfigurationCreate.md) |  | 

### Return type

[**SmartAccountConfiguration**](SmartAccountConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSmartAccountConfiguration

> SmartAccountConfiguration GetSmartAccountConfiguration(ctx).Execute()

Returns a smart account configuration.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensingConfigurationsApi.GetSmartAccountConfiguration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensingConfigurationsApi.GetSmartAccountConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSmartAccountConfiguration`: SmartAccountConfiguration
    fmt.Fprintf(os.Stdout, "Response from `LicensingConfigurationsApi.GetSmartAccountConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSmartAccountConfigurationRequest struct via the builder pattern


### Return type

[**SmartAccountConfiguration**](SmartAccountConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSmartAccountConfiguration

> SmartAccountConfiguration UpdateSmartAccountConfiguration(ctx).SmartAccountConfigurationUpdate(smartAccountConfigurationUpdate).Execute()

Updates a smart account configuration.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    smartAccountConfigurationUpdate := *openapiclient.NewSmartAccountConfigurationUpdate() // SmartAccountConfigurationUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensingConfigurationsApi.UpdateSmartAccountConfiguration(context.Background()).SmartAccountConfigurationUpdate(smartAccountConfigurationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensingConfigurationsApi.UpdateSmartAccountConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSmartAccountConfiguration`: SmartAccountConfiguration
    fmt.Fprintf(os.Stdout, "Response from `LicensingConfigurationsApi.UpdateSmartAccountConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSmartAccountConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smartAccountConfigurationUpdate** | [**SmartAccountConfigurationUpdate**](SmartAccountConfigurationUpdate.md) |  | 

### Return type

[**SmartAccountConfiguration**](SmartAccountConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

