# \WorkflowCategoriesApi

All URIs are relative to *http://localhost:8765*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkflowCategory**](WorkflowCategoriesApi.md#CreateWorkflowCategory) | **Post** /workflow/api/v8/categories | Creates a new workflow category.
[**DeleteWorkflowCategory**](WorkflowCategoriesApi.md#DeleteWorkflowCategory) | **Delete** /workflow/api/v8/categories/{id} | Deletes a workflow category.
[**GetWorkflowCategoriesList**](WorkflowCategoriesApi.md#GetWorkflowCategoriesList) | **Get** /workflow/api/v8/categories/list | Returns a list of workflow categories.
[**GetWorkflowCategory**](WorkflowCategoriesApi.md#GetWorkflowCategory) | **Get** /workflow/api/v8/categories/{id} | Returns a workflow category.
[**UpdateWorkflowCategory**](WorkflowCategoriesApi.md#UpdateWorkflowCategory) | **Put** /workflow/api/v8/categories/{id} | Updates a workflow category.



## CreateWorkflowCategory

> WorkflowCategory CreateWorkflowCategory(ctx).TenantId(tenantId).WorkflowCategoryCreate(workflowCategoryCreate).Execute()

Creates a new workflow category.

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
    tenantId := TODO // string | 
    workflowCategoryCreate := *openapiclient.NewWorkflowCategoryCreate("Name_example", "Title_example", "Description_example", "SchemaId_example") // WorkflowCategoryCreate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowCategoriesApi.CreateWorkflowCategory(context.Background()).TenantId(tenantId).WorkflowCategoryCreate(workflowCategoryCreate).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowCategoriesApi.CreateWorkflowCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowCategory`: WorkflowCategory
    fmt.Fprintf(os.Stdout, "Response from `WorkflowCategoriesApi.CreateWorkflowCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | [**string**](string.md) |  | 
 **workflowCategoryCreate** | [**WorkflowCategoryCreate**](WorkflowCategoryCreate.md) |  | 

### Return type

[**WorkflowCategory**](WorkflowCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkflowCategory

> DeleteWorkflowCategory(ctx, id).Execute()

Deletes a workflow category.

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
    resp, r, err := api_client.WorkflowCategoriesApi.DeleteWorkflowCategory(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowCategoriesApi.DeleteWorkflowCategory``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteWorkflowCategoryRequest struct via the builder pattern


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


## GetWorkflowCategoriesList

> []WorkflowCategory GetWorkflowCategoriesList(ctx).TenantId(tenantId).Execute()

Returns a list of workflow categories.

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
    tenantId := TODO // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowCategoriesApi.GetWorkflowCategoriesList(context.Background()).TenantId(tenantId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowCategoriesApi.GetWorkflowCategoriesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowCategoriesList`: []WorkflowCategory
    fmt.Fprintf(os.Stdout, "Response from `WorkflowCategoriesApi.GetWorkflowCategoriesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowCategoriesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | [**string**](string.md) |  | 

### Return type

[**[]WorkflowCategory**](WorkflowCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowCategory

> WorkflowCategory GetWorkflowCategory(ctx, id).Execute()

Returns a workflow category.

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
    resp, r, err := api_client.WorkflowCategoriesApi.GetWorkflowCategory(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowCategoriesApi.GetWorkflowCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowCategory`: WorkflowCategory
    fmt.Fprintf(os.Stdout, "Response from `WorkflowCategoriesApi.GetWorkflowCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowCategory**](WorkflowCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkflowCategory

> WorkflowCategory UpdateWorkflowCategory(ctx, id).WorkflowCategoryUpdate(workflowCategoryUpdate).Execute()

Updates a workflow category.

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
    workflowCategoryUpdate := *openapiclient.NewWorkflowCategoryUpdate("Name_example", "Title_example", "Description_example", "SchemaId_example") // WorkflowCategoryUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowCategoriesApi.UpdateWorkflowCategory(context.Background(), id).WorkflowCategoryUpdate(workflowCategoryUpdate).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowCategoriesApi.UpdateWorkflowCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkflowCategory`: WorkflowCategory
    fmt.Fprintf(os.Stdout, "Response from `WorkflowCategoriesApi.UpdateWorkflowCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkflowCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowCategoryUpdate** | [**WorkflowCategoryUpdate**](WorkflowCategoryUpdate.md) |  | 

### Return type

[**WorkflowCategory**](WorkflowCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

