# \WorkflowInstancesApi

All URIs are relative to *http://localhost:8765*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelWorkflowInstance**](WorkflowInstancesApi.md#CancelWorkflowInstance) | **Post** /workflow/api/v8/workflows/instances/{id}/cancel | Cancels a workflow instance.
[**DeleteWorkflowInstance**](WorkflowInstancesApi.md#DeleteWorkflowInstance) | **Delete** /workflow/api/v8/workflows/instances/{id} | Deletes a workflow instance.
[**GetWorkflowInstance**](WorkflowInstancesApi.md#GetWorkflowInstance) | **Get** /workflow/api/v8/workflows/instances/{id} | Returns a workflow instance.
[**GetWorkflowInstanceAction**](WorkflowInstancesApi.md#GetWorkflowInstanceAction) | **Get** /workflow/api/v8/workflows/instances/{id}/actions/{actionId} | Returns a workflow instance action.
[**GetWorkflowInstancesList**](WorkflowInstancesApi.md#GetWorkflowInstancesList) | **Get** /workflow/api/v8/workflows/{id}/instances/list | Returns a list of workflow instances.



## CancelWorkflowInstance

> WorkflowInstance CancelWorkflowInstance(ctx, id).Execute()

Cancels a workflow instance.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowInstancesApi.CancelWorkflowInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowInstancesApi.CancelWorkflowInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelWorkflowInstance`: WorkflowInstance
    fmt.Fprintf(os.Stdout, "Response from `WorkflowInstancesApi.CancelWorkflowInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelWorkflowInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowInstance**](WorkflowInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkflowInstance

> WorkflowInstanceDeleteResponse DeleteWorkflowInstance(ctx, id).Execute()

Deletes a workflow instance.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowInstancesApi.DeleteWorkflowInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowInstancesApi.DeleteWorkflowInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWorkflowInstance`: WorkflowInstanceDeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowInstancesApi.DeleteWorkflowInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkflowInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowInstanceDeleteResponse**](WorkflowInstanceDeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowInstance

> WorkflowInstance GetWorkflowInstance(ctx, id).Execute()

Returns a workflow instance.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowInstancesApi.GetWorkflowInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowInstancesApi.GetWorkflowInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowInstance`: WorkflowInstance
    fmt.Fprintf(os.Stdout, "Response from `WorkflowInstancesApi.GetWorkflowInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowInstance**](WorkflowInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowInstanceAction

> WorkflowAction GetWorkflowInstanceAction(ctx, id, actionId).Execute()

Returns a workflow instance action.

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
    actionId := "actionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowInstancesApi.GetWorkflowInstanceAction(context.Background(), id, actionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowInstancesApi.GetWorkflowInstanceAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowInstanceAction`: WorkflowAction
    fmt.Fprintf(os.Stdout, "Response from `WorkflowInstancesApi.GetWorkflowInstanceAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**actionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowInstanceActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WorkflowAction**](WorkflowAction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowInstancesList

> []WorkflowInstance GetWorkflowInstancesList(ctx, id).Page(page).PageSize(pageSize).DateFrom(dateFrom).DateTo(dateTo).Execute()

Returns a list of workflow instances.

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
    id := "id_example" // string | 
    page := int32(0) // int32 | 
    pageSize := int32(10) // int32 | 
    dateFrom := time.Now() // string |  (optional)
    dateTo := time.Now() // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowInstancesApi.GetWorkflowInstancesList(context.Background(), id).Page(page).PageSize(pageSize).DateFrom(dateFrom).DateTo(dateTo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowInstancesApi.GetWorkflowInstancesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowInstancesList`: []WorkflowInstance
    fmt.Fprintf(os.Stdout, "Response from `WorkflowInstancesApi.GetWorkflowInstancesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowInstancesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **dateFrom** | **string** |  | 
 **dateTo** | **string** |  | 

### Return type

[**[]WorkflowInstance**](WorkflowInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

