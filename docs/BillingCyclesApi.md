# \BillingCyclesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBillingCycle**](BillingCyclesApi.md#AddBillingCycle) | **Post** /billing/api/v8/cycles | Add a billing cycle.
[**DeleteBillingCycle**](BillingCyclesApi.md#DeleteBillingCycle) | **Delete** /billing/api/v8/cycles/{id} | Delete a billing cycle.
[**GetBillingCycle**](BillingCyclesApi.md#GetBillingCycle) | **Get** /billing/api/v8/cycles/{id} | Get a billing cycle.
[**GetBillingCyclesPage**](BillingCyclesApi.md#GetBillingCyclesPage) | **Get** /billing/api/v8/cycles | Retrieve a page of billing cycles.
[**ProcessBillingCycle**](BillingCyclesApi.md#ProcessBillingCycle) | **Post** /billing/api/v8/cycles/process | Process a billing cycle.
[**UpdateBillingCycle**](BillingCyclesApi.md#UpdateBillingCycle) | **Put** /billing/api/v8/cycles/{id} | Update billing cycle for an event type and tenant.



## AddBillingCycle

> BillingCycle AddBillingCycle(ctx).BillingCycleCreate(billingCycleCreate).Execute()

Add a billing cycle.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    billingCycleCreate := *openapiclient.NewBillingCycleCreate("EventId_example", time.Now(), time.Now(), "TenantId_example") // BillingCycleCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingCyclesApi.AddBillingCycle(context.Background()).BillingCycleCreate(billingCycleCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingCyclesApi.AddBillingCycle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddBillingCycle`: BillingCycle
    fmt.Fprintf(os.Stdout, "Response from `BillingCyclesApi.AddBillingCycle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddBillingCycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **billingCycleCreate** | [**BillingCycleCreate**](BillingCycleCreate.md) |  | 

### Return type

[**BillingCycle**](BillingCycle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBillingCycle

> DeleteBillingCycle(ctx, id).Execute()

Delete a billing cycle.



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
    resp, r, err := apiClient.BillingCyclesApi.DeleteBillingCycle(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingCyclesApi.DeleteBillingCycle``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteBillingCycleRequest struct via the builder pattern


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


## GetBillingCycle

> BillingCycle GetBillingCycle(ctx, id).Execute()

Get a billing cycle.



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
    resp, r, err := apiClient.BillingCyclesApi.GetBillingCycle(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingCyclesApi.GetBillingCycle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBillingCycle`: BillingCycle
    fmt.Fprintf(os.Stdout, "Response from `BillingCyclesApi.GetBillingCycle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingCycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BillingCycle**](BillingCycle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingCyclesPage

> BillingCyclesPage GetBillingCyclesPage(ctx).TenantId(tenantId).Page(page).PageSize(pageSize).NextBilledOn(nextBilledOn).Execute()

Retrieve a page of billing cycles.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    page := int32(0) // int32 | 
    pageSize := int32(10) // int32 | 
    nextBilledOn := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingCyclesApi.GetBillingCyclesPage(context.Background()).TenantId(tenantId).Page(page).PageSize(pageSize).NextBilledOn(nextBilledOn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingCyclesApi.GetBillingCyclesPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBillingCyclesPage`: BillingCyclesPage
    fmt.Fprintf(os.Stdout, "Response from `BillingCyclesApi.GetBillingCyclesPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingCyclesPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **nextBilledOn** | **time.Time** |  | 

### Return type

[**BillingCyclesPage**](BillingCyclesPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessBillingCycle

> BillingCycleProcessAccepted ProcessBillingCycle(ctx).BillingCycleProcess(billingCycleProcess).Execute()

Process a billing cycle.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    billingCycleProcess := *openapiclient.NewBillingCycleProcess(time.Now()) // BillingCycleProcess | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingCyclesApi.ProcessBillingCycle(context.Background()).BillingCycleProcess(billingCycleProcess).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingCyclesApi.ProcessBillingCycle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessBillingCycle`: BillingCycleProcessAccepted
    fmt.Fprintf(os.Stdout, "Response from `BillingCyclesApi.ProcessBillingCycle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessBillingCycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **billingCycleProcess** | [**BillingCycleProcess**](BillingCycleProcess.md) |  | 

### Return type

[**BillingCycleProcessAccepted**](BillingCycleProcessAccepted.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBillingCycle

> BillingCycle UpdateBillingCycle(ctx, id).BillingCycleUpdate(billingCycleUpdate).Execute()

Update billing cycle for an event type and tenant.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    billingCycleUpdate := *openapiclient.NewBillingCycleUpdate("EventId_example", time.Now(), time.Now(), "TenantId_example") // BillingCycleUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingCyclesApi.UpdateBillingCycle(context.Background(), id).BillingCycleUpdate(billingCycleUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingCyclesApi.UpdateBillingCycle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBillingCycle`: BillingCycle
    fmt.Fprintf(os.Stdout, "Response from `BillingCyclesApi.UpdateBillingCycle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBillingCycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **billingCycleUpdate** | [**BillingCycleUpdate**](BillingCycleUpdate.md) |  | 

### Return type

[**BillingCycle**](BillingCycle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

