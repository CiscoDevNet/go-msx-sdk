# \WorkflowEventsApi

All URIs are relative to *http://localhost:8765*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkflowEvent**](WorkflowEventsApi.md#CreateWorkflowEvent) | **Post** /workflow/api/v8/events | Creates a new workflow event.
[**DeleteWorkflowEvent**](WorkflowEventsApi.md#DeleteWorkflowEvent) | **Delete** /workflow/api/v8/events/{id} | Deletes a workflow event.
[**GetWorkflowEvent**](WorkflowEventsApi.md#GetWorkflowEvent) | **Get** /workflow/api/v8/events/{id} | Returns a workflow event.
[**GetWorkflowEventsList**](WorkflowEventsApi.md#GetWorkflowEventsList) | **Get** /workflow/api/v8/events/list | Returns a list of workflow events.
[**UpdateWorkflowEvent**](WorkflowEventsApi.md#UpdateWorkflowEvent) | **Put** /workflow/api/v8/events/{id} | Updates a workflow event.



## CreateWorkflowEvent

> WorkflowEvent CreateWorkflowEvent(ctx).WorkflowEventCreate(workflowEventCreate).Execute()

Creates a new workflow event.

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
    workflowEventCreate := *openapiclient.NewWorkflowEventCreate("Title_example", "TargetId_example", "SchemaId_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}) // WorkflowEventCreate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowEventsApi.CreateWorkflowEvent(context.Background()).WorkflowEventCreate(workflowEventCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowEventsApi.CreateWorkflowEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowEvent`: WorkflowEvent
    fmt.Fprintf(os.Stdout, "Response from `WorkflowEventsApi.CreateWorkflowEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowEventCreate** | [**WorkflowEventCreate**](WorkflowEventCreate.md) |  | 

### Return type

[**WorkflowEvent**](WorkflowEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkflowEvent

> DeleteWorkflowEvent(ctx, id).Execute()

Deletes a workflow event.

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
    resp, r, err := api_client.WorkflowEventsApi.DeleteWorkflowEvent(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowEventsApi.DeleteWorkflowEvent``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteWorkflowEventRequest struct via the builder pattern


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


## GetWorkflowEvent

> WorkflowEvent GetWorkflowEvent(ctx, id).Execute()

Returns a workflow event.

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
    resp, r, err := api_client.WorkflowEventsApi.GetWorkflowEvent(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowEventsApi.GetWorkflowEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowEvent`: WorkflowEvent
    fmt.Fprintf(os.Stdout, "Response from `WorkflowEventsApi.GetWorkflowEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowEvent**](WorkflowEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowEventsList

> []WorkflowEvent GetWorkflowEventsList(ctx).Execute()

Returns a list of workflow events.

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowEventsApi.GetWorkflowEventsList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowEventsApi.GetWorkflowEventsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowEventsList`: []WorkflowEvent
    fmt.Fprintf(os.Stdout, "Response from `WorkflowEventsApi.GetWorkflowEventsList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowEventsListRequest struct via the builder pattern


### Return type

[**[]WorkflowEvent**](WorkflowEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkflowEvent

> WorkflowEvent UpdateWorkflowEvent(ctx, id).WorkflowEventUpdate(workflowEventUpdate).Execute()

Updates a workflow event.

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
    workflowEventUpdate := *openapiclient.NewWorkflowEventUpdate("Title_example", "TargetId_example", "SchemaId_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}) // WorkflowEventUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowEventsApi.UpdateWorkflowEvent(context.Background(), id).WorkflowEventUpdate(workflowEventUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowEventsApi.UpdateWorkflowEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkflowEvent`: WorkflowEvent
    fmt.Fprintf(os.Stdout, "Response from `WorkflowEventsApi.UpdateWorkflowEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkflowEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowEventUpdate** | [**WorkflowEventUpdate**](WorkflowEventUpdate.md) |  | 

### Return type

[**WorkflowEvent**](WorkflowEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

