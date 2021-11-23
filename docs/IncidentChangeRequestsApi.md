# \IncidentChangeRequestsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChangeRequest**](IncidentChangeRequestsApi.md#CreateChangeRequest) | **Post** /incident/api/v8/changerequests | Create a change request.
[**DeleteChangeRequest**](IncidentChangeRequestsApi.md#DeleteChangeRequest) | **Delete** /incident/api/v8/changerequests/{id} | Deletes a change request.
[**GetChangeRequest**](IncidentChangeRequestsApi.md#GetChangeRequest) | **Get** /incident/api/v8/changerequests/{id} | Returns a change request.
[**GetChangeRequestsPage**](IncidentChangeRequestsApi.md#GetChangeRequestsPage) | **Get** /incident/api/v8/changerequests | Returns a filtered list of change requests.
[**UpdateChangeRequest**](IncidentChangeRequestsApi.md#UpdateChangeRequest) | **Put** /incident/api/v8/changerequests/{id} | Updates a change request.



## CreateChangeRequest

> ChangeRequest CreateChangeRequest(ctx).ChangeRequestCreate(changeRequestCreate).Execute()

Create a change request.

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
    changeRequestCreate := *openapiclient.NewChangeRequestCreate() // ChangeRequestCreate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentChangeRequestsApi.CreateChangeRequest(context.Background()).ChangeRequestCreate(changeRequestCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentChangeRequestsApi.CreateChangeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateChangeRequest`: ChangeRequest
    fmt.Fprintf(os.Stdout, "Response from `IncidentChangeRequestsApi.CreateChangeRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **changeRequestCreate** | [**ChangeRequestCreate**](ChangeRequestCreate.md) |  | 

### Return type

[**ChangeRequest**](ChangeRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteChangeRequest

> DeleteChangeRequest(ctx, id).TenantId(tenantId).Execute()

Deletes a change request.

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
    id := "id_example" // string | 
    tenantId := TODO // string | Required for bi-directional scenario (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentChangeRequestsApi.DeleteChangeRequest(context.Background(), id).TenantId(tenantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentChangeRequestsApi.DeleteChangeRequest``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | [**string**](string.md) | Required for bi-directional scenario | 

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


## GetChangeRequest

> ChangeRequest GetChangeRequest(ctx, id).TenantId(tenantId).Execute()

Returns a change request.

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
    id := "id_example" // string | Change Request Number  CHG0030022
    tenantId := TODO // string | Required for bi-directional scenario (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentChangeRequestsApi.GetChangeRequest(context.Background(), id).TenantId(tenantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentChangeRequestsApi.GetChangeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChangeRequest`: ChangeRequest
    fmt.Fprintf(os.Stdout, "Response from `IncidentChangeRequestsApi.GetChangeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Change Request Number  CHG0030022 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | [**string**](string.md) | Required for bi-directional scenario | 

### Return type

[**ChangeRequest**](ChangeRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChangeRequestsPage

> ChangeRequestsPage GetChangeRequestsPage(ctx).Page(page).PageSize(pageSize).TenantId(tenantId).Execute()

Returns a filtered list of change requests.

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
    tenantId := TODO // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentChangeRequestsApi.GetChangeRequestsPage(context.Background()).Page(page).PageSize(pageSize).TenantId(tenantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentChangeRequestsApi.GetChangeRequestsPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChangeRequestsPage`: ChangeRequestsPage
    fmt.Fprintf(os.Stdout, "Response from `IncidentChangeRequestsApi.GetChangeRequestsPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChangeRequestsPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **tenantId** | [**string**](string.md) |  | 

### Return type

[**ChangeRequestsPage**](ChangeRequestsPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChangeRequest

> ChangeRequest UpdateChangeRequest(ctx, id).ChangeRequestUpdate(changeRequestUpdate).Execute()

Updates a change request.



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
    id := "id_example" // string | 
    changeRequestUpdate := *openapiclient.NewChangeRequestUpdate() // ChangeRequestUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentChangeRequestsApi.UpdateChangeRequest(context.Background(), id).ChangeRequestUpdate(changeRequestUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentChangeRequestsApi.UpdateChangeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateChangeRequest`: ChangeRequest
    fmt.Fprintf(os.Stdout, "Response from `IncidentChangeRequestsApi.UpdateChangeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changeRequestUpdate** | [**ChangeRequestUpdate**](ChangeRequestUpdate.md) |  | 

### Return type

[**ChangeRequest**](ChangeRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

