# \WorkflowTargetsApi

All URIs are relative to *http://localhost:8765*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkflowTarget**](WorkflowTargetsApi.md#CreateWorkflowTarget) | **Post** /workflow/api/v8/targets | Creates a new workflow target.
[**DeleteWorkflowTarget**](WorkflowTargetsApi.md#DeleteWorkflowTarget) | **Delete** /workflow/api/v8/targets/{id} | Deletes a workflow target.
[**GetWorkflowTarget**](WorkflowTargetsApi.md#GetWorkflowTarget) | **Get** /workflow/api/v8/targets/{id} | Returns a workflow target.
[**GetWorkflowTargetsList**](WorkflowTargetsApi.md#GetWorkflowTargetsList) | **Get** /workflow/api/v8/targets/list | Returns a list of workflow targets.
[**UpdateWorkflowTarget**](WorkflowTargetsApi.md#UpdateWorkflowTarget) | **Put** /workflow/api/v8/targets/{id} | Updates a workflow target.



## CreateWorkflowTarget

> WorkflowTarget CreateWorkflowTarget(ctx).WorkflowTargetCreate(workflowTargetCreate).Execute()

Creates a new workflow target.

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
    workflowTargetCreate := *openapiclient.NewWorkflowTargetCreate("Name_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}) // WorkflowTargetCreate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowTargetsApi.CreateWorkflowTarget(context.Background()).WorkflowTargetCreate(workflowTargetCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTargetsApi.CreateWorkflowTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowTarget`: WorkflowTarget
    fmt.Fprintf(os.Stdout, "Response from `WorkflowTargetsApi.CreateWorkflowTarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowTargetCreate** | [**WorkflowTargetCreate**](WorkflowTargetCreate.md) |  | 

### Return type

[**WorkflowTarget**](WorkflowTarget.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkflowTarget

> DeleteWorkflowTarget(ctx, id).Execute()

Deletes a workflow target.

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
    resp, r, err := api_client.WorkflowTargetsApi.DeleteWorkflowTarget(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTargetsApi.DeleteWorkflowTarget``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteWorkflowTargetRequest struct via the builder pattern


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


## GetWorkflowTarget

> WorkflowTarget GetWorkflowTarget(ctx, id).Execute()

Returns a workflow target.

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
    resp, r, err := api_client.WorkflowTargetsApi.GetWorkflowTarget(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTargetsApi.GetWorkflowTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowTarget`: WorkflowTarget
    fmt.Fprintf(os.Stdout, "Response from `WorkflowTargetsApi.GetWorkflowTarget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowTarget**](WorkflowTarget.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowTargetsList

> []WorkflowTarget GetWorkflowTargetsList(ctx).Internal(internal).Type_(type_).Execute()

Returns a list of workflow targets.

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
    internal := true // bool |  (optional)
    type_ := "type__example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowTargetsApi.GetWorkflowTargetsList(context.Background()).Internal(internal).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTargetsApi.GetWorkflowTargetsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowTargetsList`: []WorkflowTarget
    fmt.Fprintf(os.Stdout, "Response from `WorkflowTargetsApi.GetWorkflowTargetsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowTargetsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **internal** | **bool** |  | 
 **type_** | **string** |  | 

### Return type

[**[]WorkflowTarget**](WorkflowTarget.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkflowTarget

> WorkflowTarget UpdateWorkflowTarget(ctx, id).WorkflowTargetUpdate(workflowTargetUpdate).Execute()

Updates a workflow target.

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
    workflowTargetUpdate := *openapiclient.NewWorkflowTargetUpdate("Name_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}) // WorkflowTargetUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowTargetsApi.UpdateWorkflowTarget(context.Background(), id).WorkflowTargetUpdate(workflowTargetUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTargetsApi.UpdateWorkflowTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkflowTarget`: WorkflowTarget
    fmt.Fprintf(os.Stdout, "Response from `WorkflowTargetsApi.UpdateWorkflowTarget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkflowTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowTargetUpdate** | [**WorkflowTargetUpdate**](WorkflowTargetUpdate.md) |  | 

### Return type

[**WorkflowTarget**](WorkflowTarget.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

