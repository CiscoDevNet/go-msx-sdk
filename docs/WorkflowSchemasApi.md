# \WorkflowSchemasApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWorkflowSchema**](WorkflowSchemasApi.md#GetWorkflowSchema) | **Get** /workflow/api/v8/schemas/{id} | Returns a workflow schema.
[**GetWorkflowSchemasList**](WorkflowSchemasApi.md#GetWorkflowSchemasList) | **Get** /workflow/api/v8/schemas/list | Returns a list of workflow schemas.



## GetWorkflowSchema

> WorkflowSchemaByTypeResponse GetWorkflowSchema(ctx, id).SchemaType(schemaType).Execute()

Returns a workflow schema.

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
    schemaType := "schemaType_example" // string |  (optional) (default to "view")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowSchemasApi.GetWorkflowSchema(context.Background(), id).SchemaType(schemaType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemasApi.GetWorkflowSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowSchema`: WorkflowSchemaByTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemasApi.GetWorkflowSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **schemaType** | **string** |  | [default to &quot;view&quot;]

### Return type

[**WorkflowSchemaByTypeResponse**](WorkflowSchemaByTypeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowSchemasList

> []WorkflowSchema GetWorkflowSchemasList(ctx).BaseType(baseType).SchemaType(schemaType).VariableType(variableType).Execute()

Returns a list of workflow schemas.

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
    baseType := "baseType_example" // string |  (default to "category")
    schemaType := "schemaType_example" // string |  (optional) (default to "view")
    variableType := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowSchemasApi.GetWorkflowSchemasList(context.Background()).BaseType(baseType).SchemaType(schemaType).VariableType(variableType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemasApi.GetWorkflowSchemasList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowSchemasList`: []WorkflowSchema
    fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemasApi.GetWorkflowSchemasList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowSchemasListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseType** | **string** |  | [default to &quot;category&quot;]
 **schemaType** | **string** |  | [default to &quot;view&quot;]
 **variableType** | **bool** |  | [default to false]

### Return type

[**[]WorkflowSchema**](WorkflowSchema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

