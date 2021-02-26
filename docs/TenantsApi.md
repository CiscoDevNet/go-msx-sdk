# \TenantsApi

All URIs are relative to *http://localhost:8765*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTenant**](TenantsApi.md#CreateTenant) | **Post** /idm/api/v8/tenants | Creates a new tenant.
[**DeleteTenant**](TenantsApi.md#DeleteTenant) | **Delete** /idm/api/v8/tenants/{id} | Deletes a tenant by id.
[**GetTenant**](TenantsApi.md#GetTenant) | **Get** /idm/api/v8/tenants/{id} | Returns a tenant by id.
[**GetTenantsList**](TenantsApi.md#GetTenantsList) | **Get** /idm/api/v8/tenants/list | Returns a list of tenants.
[**GetTenantsPage**](TenantsApi.md#GetTenantsPage) | **Get** /idm/api/v8/tenants | Returns a page of tenants.
[**UpdateTenant**](TenantsApi.md#UpdateTenant) | **Put** /idm/api/v8/tenants/{id} | Updates a tenant by id.



## CreateTenant

> Tenant CreateTenant(ctx).TenantCreate(tenantCreate).Execute()

Creates a new tenant.

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
    tenantCreate := *openapiclient.NewTenantCreate("Name_example") // TenantCreate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantsApi.CreateTenant(context.Background()).TenantCreate(tenantCreate).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.CreateTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTenant`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.CreateTenant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantCreate** | [**TenantCreate**](TenantCreate.md) |  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTenant

> DeleteTenant(ctx, id).Execute()

Deletes a tenant by id.

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
    resp, r, err := api_client.TenantsApi.DeleteTenant(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.DeleteTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantRequest struct via the builder pattern


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


## GetTenant

> Tenant GetTenant(ctx, id).Execute()

Returns a tenant by id.

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
    resp, r, err := api_client.TenantsApi.GetTenant(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.GetTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenant`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.GetTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Tenant**](Tenant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantsList

> []Tenant GetTenantsList(ctx).Ids(ids).ShowImage(showImage).Execute()

Returns a list of tenants.

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
    ids := []string{"Inner_example"} // []string | 
    showImage := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantsApi.GetTenantsList(context.Background()).Ids(ids).ShowImage(showImage).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.GetTenantsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenantsList`: []Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.GetTenantsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** |  | 
 **showImage** | **bool** |  | [default to false]

### Return type

[**[]Tenant**](Tenant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantsPage

> TenantsPage GetTenantsPage(ctx).Page(page).PageSize(pageSize).ParentId(parentId).ShowImage(showImage).SortBy(sortBy).SortOrder(sortOrder).Execute()

Returns a page of tenants.

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
    parentId := TODO // string |  (optional)
    showImage := true // bool |  (optional) (default to false)
    sortBy := "name" // string |  (optional)
    sortOrder := "sortOrder_example" // string |  (optional) (default to "asc")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantsApi.GetTenantsPage(context.Background()).Page(page).PageSize(pageSize).ParentId(parentId).ShowImage(showImage).SortBy(sortBy).SortOrder(sortOrder).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.GetTenantsPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenantsPage`: TenantsPage
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.GetTenantsPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantsPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **parentId** | [**string**](string.md) |  | 
 **showImage** | **bool** |  | [default to false]
 **sortBy** | **string** |  | 
 **sortOrder** | **string** |  | [default to &quot;asc&quot;]

### Return type

[**TenantsPage**](TenantsPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTenant

> Tenant UpdateTenant(ctx, id).TenantUpdate(tenantUpdate).Execute()

Updates a tenant by id.

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
    tenantUpdate := *openapiclient.NewTenantUpdate("Name_example") // TenantUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantsApi.UpdateTenant(context.Background(), id).TenantUpdate(tenantUpdate).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantsApi.UpdateTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTenant`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenantsApi.UpdateTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantUpdate** | [**TenantUpdate**](TenantUpdate.md) |  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

