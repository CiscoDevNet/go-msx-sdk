# \TemplateApplicationsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyTemplate**](TemplateApplicationsApi.md#ApplyTemplate) | **Post** /template/api/v8/templates/{id}/applications | Applies a template to a target.
[**DeleteTemplateApplication**](TemplateApplicationsApi.md#DeleteTemplateApplication) | **Delete** /template/api/v8/templates/applications/{id} | Deletes a template application.
[**GetTemplateApplication**](TemplateApplicationsApi.md#GetTemplateApplication) | **Get** /template/api/v8/templates/applications/{id} | Gets a template application.
[**GetTemplateApplicationHistory**](TemplateApplicationsApi.md#GetTemplateApplicationHistory) | **Get** /template/api/v8/templates/applications/{id}/history | Gets a template application history.
[**GetTemplateApplicationsPage**](TemplateApplicationsApi.md#GetTemplateApplicationsPage) | **Get** /template/api/v8/templates/applications | Get a page of template applications.
[**UpdateApplicationStatus**](TemplateApplicationsApi.md#UpdateApplicationStatus) | **Patch** /template/api/v8/templates/applications/{id} | Updates an application status.



## ApplyTemplate

> TemplateApplication ApplyTemplate(ctx, id).TemplateApplicationCreate(templateApplicationCreate).Execute()

Applies a template to a target.

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
    templateApplicationCreate := *openapiclient.NewTemplateApplicationCreate("TenantId_example", "TargetId_example", "TargetType_example") // TemplateApplicationCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateApplicationsApi.ApplyTemplate(context.Background(), id).TemplateApplicationCreate(templateApplicationCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApplicationsApi.ApplyTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplyTemplate`: TemplateApplication
    fmt.Fprintf(os.Stdout, "Response from `TemplateApplicationsApi.ApplyTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templateApplicationCreate** | [**TemplateApplicationCreate**](TemplateApplicationCreate.md) |  | 

### Return type

[**TemplateApplication**](TemplateApplication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTemplateApplication

> DeleteTemplateApplication(ctx, id).Execute()

Deletes a template application.

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
    resp, r, err := apiClient.TemplateApplicationsApi.DeleteTemplateApplication(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApplicationsApi.DeleteTemplateApplication``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteTemplateApplicationRequest struct via the builder pattern


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


## GetTemplateApplication

> TemplateApplication GetTemplateApplication(ctx, id).Execute()

Gets a template application.

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
    resp, r, err := apiClient.TemplateApplicationsApi.GetTemplateApplication(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApplicationsApi.GetTemplateApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTemplateApplication`: TemplateApplication
    fmt.Fprintf(os.Stdout, "Response from `TemplateApplicationsApi.GetTemplateApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TemplateApplication**](TemplateApplication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplateApplicationHistory

> []TemplateApplication GetTemplateApplicationHistory(ctx, id).Execute()

Gets a template application history.

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
    resp, r, err := apiClient.TemplateApplicationsApi.GetTemplateApplicationHistory(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApplicationsApi.GetTemplateApplicationHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTemplateApplicationHistory`: []TemplateApplication
    fmt.Fprintf(os.Stdout, "Response from `TemplateApplicationsApi.GetTemplateApplicationHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateApplicationHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TemplateApplication**](TemplateApplication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplateApplicationsPage

> TemplateApplicationsPage GetTemplateApplicationsPage(ctx).TenantId(tenantId).Page(page).PageSize(pageSize).TemplateId(templateId).TargetId(targetId).TargetType(targetType).CalculateTotalItems(calculateTotalItems).SortBy(sortBy).SortOrder(sortOrder).Execute()

Get a page of template applications.

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
    tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    page := int32(0) // int32 | 
    pageSize := int32(10) // int32 | 
    templateId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    targetId := "targetId_example" // string |  (optional)
    targetType := "targetType_example" // string |  (optional)
    calculateTotalItems := true // bool |  (optional)
    sortBy := "sortBy_example" // string |  (optional) (default to "createdOn")
    sortOrder := "sortOrder_example" // string |  (optional) (default to "asc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateApplicationsApi.GetTemplateApplicationsPage(context.Background()).TenantId(tenantId).Page(page).PageSize(pageSize).TemplateId(templateId).TargetId(targetId).TargetType(targetType).CalculateTotalItems(calculateTotalItems).SortBy(sortBy).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApplicationsApi.GetTemplateApplicationsPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTemplateApplicationsPage`: TemplateApplicationsPage
    fmt.Fprintf(os.Stdout, "Response from `TemplateApplicationsApi.GetTemplateApplicationsPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateApplicationsPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **templateId** | **string** |  | 
 **targetId** | **string** |  | 
 **targetType** | **string** |  | 
 **calculateTotalItems** | **bool** |  | 
 **sortBy** | **string** |  | [default to &quot;createdOn&quot;]
 **sortOrder** | **string** |  | [default to &quot;asc&quot;]

### Return type

[**TemplateApplicationsPage**](TemplateApplicationsPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationStatus

> TemplateApplication UpdateApplicationStatus(ctx, id).TemplateApplicationStatusPatch(templateApplicationStatusPatch).Execute()

Updates an application status.

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
    templateApplicationStatusPatch := *openapiclient.NewTemplateApplicationStatusPatch(openapiclient.TemplateStatus("NEW")) // TemplateApplicationStatusPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateApplicationsApi.UpdateApplicationStatus(context.Background(), id).TemplateApplicationStatusPatch(templateApplicationStatusPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateApplicationsApi.UpdateApplicationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplicationStatus`: TemplateApplication
    fmt.Fprintf(os.Stdout, "Response from `TemplateApplicationsApi.UpdateApplicationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templateApplicationStatusPatch** | [**TemplateApplicationStatusPatch**](TemplateApplicationStatusPatch.md) |  | 

### Return type

[**TemplateApplication**](TemplateApplication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

