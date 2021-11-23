# \TemplateAssignmentsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchAssignTemplate**](TemplateAssignmentsApi.md#BatchAssignTemplate) | **Post** /template/api/v8/templates/{id}/assignments/add | Assigns a template to one or more tenants.
[**BatchUnassignTemplate**](TemplateAssignmentsApi.md#BatchUnassignTemplate) | **Post** /template/api/v8/templates/{id}/assignments/remove | Unassigns a template from one or more tenants.
[**GetAssignment**](TemplateAssignmentsApi.md#GetAssignment) | **Get** /template/api/v8/templates/assignments/{id} | Gets a template assignment.
[**GetAssignmentHistory**](TemplateAssignmentsApi.md#GetAssignmentHistory) | **Get** /template/api/v8/templates/assignments/{id}/history | Gets a template assignment history.
[**GetTemplateAssignmentsPage**](TemplateAssignmentsApi.md#GetTemplateAssignmentsPage) | **Get** /template/api/v8/templates/assignments | Returns a page of template assignments.
[**UpdateAssignmentStatus**](TemplateAssignmentsApi.md#UpdateAssignmentStatus) | **Patch** /template/api/v8/templates/assignments/{id} | Updates a template assignment status.



## BatchAssignTemplate

> []TemplateAssignmentResponse BatchAssignTemplate(ctx, id).RequestBody(requestBody).Inheritable(inheritable).Execute()

Assigns a template to one or more tenants.

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
    id := TODO // string | 
    requestBody := []string{"Property_example"} // []string | 
    inheritable := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplateAssignmentsApi.BatchAssignTemplate(context.Background(), id).RequestBody(requestBody).Inheritable(inheritable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateAssignmentsApi.BatchAssignTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchAssignTemplate`: []TemplateAssignmentResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplateAssignmentsApi.BatchAssignTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchAssignTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** |  | 
 **inheritable** | **bool** |  | 

### Return type

[**[]TemplateAssignmentResponse**](TemplateAssignmentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchUnassignTemplate

> []TemplateAssignmentResponse BatchUnassignTemplate(ctx, id).RequestBody(requestBody).Execute()

Unassigns a template from one or more tenants.

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
    id := TODO // string | 
    requestBody := []string{"Property_example"} // []string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplateAssignmentsApi.BatchUnassignTemplate(context.Background(), id).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateAssignmentsApi.BatchUnassignTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchUnassignTemplate`: []TemplateAssignmentResponse
    fmt.Fprintf(os.Stdout, "Response from `TemplateAssignmentsApi.BatchUnassignTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchUnassignTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** |  | 

### Return type

[**[]TemplateAssignmentResponse**](TemplateAssignmentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssignment

> TemplateAssignment GetAssignment(ctx, id).Execute()

Gets a template assignment.

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
    id := TODO // string | ID of template assignment record.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplateAssignmentsApi.GetAssignment(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateAssignmentsApi.GetAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssignment`: TemplateAssignment
    fmt.Fprintf(os.Stdout, "Response from `TemplateAssignmentsApi.GetAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | ID of template assignment record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TemplateAssignment**](TemplateAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssignmentHistory

> []TemplateAssignment GetAssignmentHistory(ctx, id).Execute()

Gets a template assignment history.

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
    id := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplateAssignmentsApi.GetAssignmentHistory(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateAssignmentsApi.GetAssignmentHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssignmentHistory`: []TemplateAssignment
    fmt.Fprintf(os.Stdout, "Response from `TemplateAssignmentsApi.GetAssignmentHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssignmentHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TemplateAssignment**](TemplateAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplateAssignmentsPage

> TemplateAssignmentsPage GetTemplateAssignmentsPage(ctx).Page(page).PageSize(pageSize).TemplateId(templateId).TenantId(tenantId).CalculateTotalItems(calculateTotalItems).Execute()

Returns a page of template assignments.

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
    templateId := TODO // string |  (optional)
    tenantId := TODO // string |  (optional)
    calculateTotalItems := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplateAssignmentsApi.GetTemplateAssignmentsPage(context.Background()).Page(page).PageSize(pageSize).TemplateId(templateId).TenantId(tenantId).CalculateTotalItems(calculateTotalItems).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateAssignmentsApi.GetTemplateAssignmentsPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTemplateAssignmentsPage`: TemplateAssignmentsPage
    fmt.Fprintf(os.Stdout, "Response from `TemplateAssignmentsApi.GetTemplateAssignmentsPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateAssignmentsPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **templateId** | [**string**](string.md) |  | 
 **tenantId** | [**string**](string.md) |  | 
 **calculateTotalItems** | **bool** |  | 

### Return type

[**TemplateAssignmentsPage**](TemplateAssignmentsPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssignmentStatus

> TemplateAssignment UpdateAssignmentStatus(ctx, id).TemplateAssignmentStatusPatch(templateAssignmentStatusPatch).Execute()

Updates a template assignment status.

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
    id := TODO // string | ID of template assignment record.
    templateAssignmentStatusPatch := *openapiclient.NewTemplateAssignmentStatusPatch(openapiclient.TemplateStatus("NEW")) // TemplateAssignmentStatusPatch | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplateAssignmentsApi.UpdateAssignmentStatus(context.Background(), id).TemplateAssignmentStatusPatch(templateAssignmentStatusPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateAssignmentsApi.UpdateAssignmentStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAssignmentStatus`: TemplateAssignment
    fmt.Fprintf(os.Stdout, "Response from `TemplateAssignmentsApi.UpdateAssignmentStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | ID of template assignment record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAssignmentStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templateAssignmentStatusPatch** | [**TemplateAssignmentStatusPatch**](TemplateAssignmentStatusPatch.md) |  | 

### Return type

[**TemplateAssignment**](TemplateAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

