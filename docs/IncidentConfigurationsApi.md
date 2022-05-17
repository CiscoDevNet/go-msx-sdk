# \IncidentConfigurationsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceNowConfiguration**](IncidentConfigurationsApi.md#CreateServiceNowConfiguration) | **Post** /incident/api/v8/configurations | Creates a ServiceNow configuration.
[**DeleteServiceNowConfiguration**](IncidentConfigurationsApi.md#DeleteServiceNowConfiguration) | **Delete** /incident/api/v8/configurations/{id} | Deletes a ServiceNow configuration.
[**GetConfiguration**](IncidentConfigurationsApi.md#GetConfiguration) | **Get** /incident/api/v1/config | API to get configuration of encloud service.
[**GetServiceNowConfiguration**](IncidentConfigurationsApi.md#GetServiceNowConfiguration) | **Get** /incident/api/v8/configurations/{id} | Returns a ServiceNow configuration.
[**GetServiceNowConfigurationsPage**](IncidentConfigurationsApi.md#GetServiceNowConfigurationsPage) | **Get** /incident/api/v8/configurations | Returns a ServiceNow configurations.
[**PatchConfiguration**](IncidentConfigurationsApi.md#PatchConfiguration) | **Patch** /incident/api/v1/config | API to patch configure encloud service
[**UpdateConfiguration**](IncidentConfigurationsApi.md#UpdateConfiguration) | **Put** /incident/api/v1/config | API to configure encloud service.



## CreateServiceNowConfiguration

> ServiceNowConfiguration CreateServiceNowConfiguration(ctx).ServiceNowConfigurationCreate(serviceNowConfigurationCreate).Execute()

Creates a ServiceNow configuration.

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
    serviceNowConfigurationCreate := *openapiclient.NewServiceNowConfigurationCreate() // ServiceNowConfigurationCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IncidentConfigurationsApi.CreateServiceNowConfiguration(context.Background()).ServiceNowConfigurationCreate(serviceNowConfigurationCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentConfigurationsApi.CreateServiceNowConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServiceNowConfiguration`: ServiceNowConfiguration
    fmt.Fprintf(os.Stdout, "Response from `IncidentConfigurationsApi.CreateServiceNowConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceNowConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceNowConfigurationCreate** | [**ServiceNowConfigurationCreate**](ServiceNowConfigurationCreate.md) |  | 

### Return type

[**ServiceNowConfiguration**](ServiceNowConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServiceNowConfiguration

> DeleteServiceNowConfiguration(ctx, id).Execute()

Deletes a ServiceNow configuration.



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IncidentConfigurationsApi.DeleteServiceNowConfiguration(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentConfigurationsApi.DeleteServiceNowConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceNowConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfiguration

> IncidentConfig GetConfiguration(ctx).Execute()

API to get configuration of encloud service.

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
    resp, r, err := apiClient.IncidentConfigurationsApi.GetConfiguration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentConfigurationsApi.GetConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfiguration`: IncidentConfig
    fmt.Fprintf(os.Stdout, "Response from `IncidentConfigurationsApi.GetConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationRequest struct via the builder pattern


### Return type

[**IncidentConfig**](IncidentConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceNowConfiguration

> ServiceNowConfiguration GetServiceNowConfiguration(ctx, id).Execute()

Returns a ServiceNow configuration.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IncidentConfigurationsApi.GetServiceNowConfiguration(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentConfigurationsApi.GetServiceNowConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceNowConfiguration`: ServiceNowConfiguration
    fmt.Fprintf(os.Stdout, "Response from `IncidentConfigurationsApi.GetServiceNowConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceNowConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceNowConfiguration**](ServiceNowConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceNowConfigurationsPage

> ServiceNowConfigurationsPage GetServiceNowConfigurationsPage(ctx).Page(page).PageSize(pageSize).TenantId(tenantId).Domain(domain).Execute()

Returns a ServiceNow configurations.

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
    page := int32(0) // int32 | 
    pageSize := int32(10) // int32 | 
    tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    domain := "domain_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IncidentConfigurationsApi.GetServiceNowConfigurationsPage(context.Background()).Page(page).PageSize(pageSize).TenantId(tenantId).Domain(domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentConfigurationsApi.GetServiceNowConfigurationsPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceNowConfigurationsPage`: ServiceNowConfigurationsPage
    fmt.Fprintf(os.Stdout, "Response from `IncidentConfigurationsApi.GetServiceNowConfigurationsPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceNowConfigurationsPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **tenantId** | **string** |  | 
 **domain** | **string** |  | 

### Return type

[**ServiceNowConfigurationsPage**](ServiceNowConfigurationsPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchConfiguration

> IncidentConfig PatchConfiguration(ctx).IncidentConfigPatch(incidentConfigPatch).Execute()

API to patch configure encloud service

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
    incidentConfigPatch := *openapiclient.NewIncidentConfigPatch() // IncidentConfigPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IncidentConfigurationsApi.PatchConfiguration(context.Background()).IncidentConfigPatch(incidentConfigPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentConfigurationsApi.PatchConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchConfiguration`: IncidentConfig
    fmt.Fprintf(os.Stdout, "Response from `IncidentConfigurationsApi.PatchConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **incidentConfigPatch** | [**IncidentConfigPatch**](IncidentConfigPatch.md) |  | 

### Return type

[**IncidentConfig**](IncidentConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfiguration

> IncidentConfig UpdateConfiguration(ctx).IncidentConfigUpdate(incidentConfigUpdate).Execute()

API to configure encloud service.

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
    incidentConfigUpdate := *openapiclient.NewIncidentConfigUpdate("ClientId_example", "ClientSecret_example", "Domain_example", "Password_example", "UserName_example") // IncidentConfigUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IncidentConfigurationsApi.UpdateConfiguration(context.Background()).IncidentConfigUpdate(incidentConfigUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentConfigurationsApi.UpdateConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConfiguration`: IncidentConfig
    fmt.Fprintf(os.Stdout, "Response from `IncidentConfigurationsApi.UpdateConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **incidentConfigUpdate** | [**IncidentConfigUpdate**](IncidentConfigUpdate.md) |  | 

### Return type

[**IncidentConfig**](IncidentConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

