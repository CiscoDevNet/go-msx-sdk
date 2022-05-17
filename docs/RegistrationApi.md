# \RegistrationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRegisteredProductVersion**](RegistrationApi.md#DeleteRegisteredProductVersion) | **Delete** /vulnerability/api/v8/vulnerabilities/registrations/{id} | Delete a registration.
[**GetRegisteredProductVersionPage**](RegistrationApi.md#GetRegisteredProductVersionPage) | **Get** /vulnerability/api/v8/vulnerabilities/registrations | Returns a filtered page of registered products / versions.
[**RegisterProductVersion**](RegistrationApi.md#RegisterProductVersion) | **Post** /vulnerability/api/v8/vulnerabilities/registrations | Register a product / verison combination for vulnerability inspection.



## DeleteRegisteredProductVersion

> DeleteRegisteredProductVersion(ctx, id).Execute()

Delete a registration.

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
    resp, r, err := apiClient.RegistrationApi.DeleteRegisteredProductVersion(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistrationApi.DeleteRegisteredProductVersion``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteRegisteredProductVersionRequest struct via the builder pattern


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


## GetRegisteredProductVersionPage

> VulnerabilitiesRegistrationPage GetRegisteredProductVersionPage(ctx).Page(page).PageSize(pageSize).Product(product).Version(version).Execute()

Returns a filtered page of registered products / versions.

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
    product := "ios_xe" // string | Product identifier (as defined in NIST's CPE dictionary) to filter by. (optional)
    version := "12.3" // string | Product version (as defined in NIST's CPE dictionary) filter to filter by. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistrationApi.GetRegisteredProductVersionPage(context.Background()).Page(page).PageSize(pageSize).Product(product).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistrationApi.GetRegisteredProductVersionPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegisteredProductVersionPage`: VulnerabilitiesRegistrationPage
    fmt.Fprintf(os.Stdout, "Response from `RegistrationApi.GetRegisteredProductVersionPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRegisteredProductVersionPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **product** | **string** | Product identifier (as defined in NIST&#39;s CPE dictionary) to filter by. | 
 **version** | **string** | Product version (as defined in NIST&#39;s CPE dictionary) filter to filter by. | 

### Return type

[**VulnerabilitiesRegistrationPage**](VulnerabilitiesRegistrationPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterProductVersion

> VulnerabilityRegistration RegisterProductVersion(ctx).VulnerabilityRegistrationCreate(vulnerabilityRegistrationCreate).Execute()

Register a product / verison combination for vulnerability inspection.

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
    vulnerabilityRegistrationCreate := *openapiclient.NewVulnerabilityRegistrationCreate() // VulnerabilityRegistrationCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistrationApi.RegisterProductVersion(context.Background()).VulnerabilityRegistrationCreate(vulnerabilityRegistrationCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistrationApi.RegisterProductVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterProductVersion`: VulnerabilityRegistration
    fmt.Fprintf(os.Stdout, "Response from `RegistrationApi.RegisterProductVersion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterProductVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vulnerabilityRegistrationCreate** | [**VulnerabilityRegistrationCreate**](VulnerabilityRegistrationCreate.md) |  | 

### Return type

[**VulnerabilityRegistration**](VulnerabilityRegistration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

