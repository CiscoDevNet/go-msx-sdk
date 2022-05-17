# \SitesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDevicesToSite**](SitesApi.md#AddDevicesToSite) | **Post** /manage/api/v8/sites/{id}/devices/add | Add devices to a site.
[**AddServicesToSite**](SitesApi.md#AddServicesToSite) | **Post** /manage/api/v8/sites/{id}/services/add | Add services to a site.
[**CreateSite**](SitesApi.md#CreateSite) | **Post** /manage/api/v8/sites | Creates a new site.
[**DeleteSite**](SitesApi.md#DeleteSite) | **Delete** /manage/api/v8/sites/{id} | Deletes a site.
[**GetSite**](SitesApi.md#GetSite) | **Get** /manage/api/v8/sites/{id} | Returns a site.
[**GetSitesPage**](SitesApi.md#GetSitesPage) | **Get** /manage/api/v8/sites | Returns a page of Sites. Only one filter is supported at a time.
[**RemoveDevicesFromSite**](SitesApi.md#RemoveDevicesFromSite) | **Post** /manage/api/v8/sites/{id}/devices/remove | Removes devices from a site.
[**RemoveServicesFromSite**](SitesApi.md#RemoveServicesFromSite) | **Post** /manage/api/v8/sites/{id}/services/remove | Remove services from a site.
[**UpdateSite**](SitesApi.md#UpdateSite) | **Put** /manage/api/v8/sites/{id} | Updates a site.



## AddDevicesToSite

> Site AddDevicesToSite(ctx, id).RequestBody(requestBody).Execute()

Add devices to a site.

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
    requestBody := []string{"Property_example"} // []string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.AddDevicesToSite(context.Background(), id).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.AddDevicesToSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddDevicesToSite`: Site
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.AddDevicesToSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDevicesToSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** |  | 

### Return type

[**Site**](Site.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddServicesToSite

> Site AddServicesToSite(ctx, id).RequestBody(requestBody).Execute()

Add services to a site.

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
    requestBody := []string{"Property_example"} // []string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.AddServicesToSite(context.Background(), id).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.AddServicesToSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddServicesToSite`: Site
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.AddServicesToSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddServicesToSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** |  | 

### Return type

[**Site**](Site.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSite

> Site CreateSite(ctx).SiteCreate(siteCreate).Execute()

Creates a new site.

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
    siteCreate := *openapiclient.NewSiteCreate("TenantId_example", "Name_example") // SiteCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.CreateSite(context.Background()).SiteCreate(siteCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.CreateSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSite`: Site
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.CreateSite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **siteCreate** | [**SiteCreate**](SiteCreate.md) |  | 

### Return type

[**Site**](Site.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSite

> DeleteSite(ctx, id).Execute()

Deletes a site.

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
    resp, r, err := apiClient.SitesApi.DeleteSite(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.DeleteSite``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSiteRequest struct via the builder pattern


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


## GetSite

> Site GetSite(ctx, id).ShowImage(showImage).Execute()

Returns a site.

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
    showImage := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.GetSite(context.Background(), id).ShowImage(showImage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.GetSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSite`: Site
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.GetSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **showImage** | **bool** |  | [default to false]

### Return type

[**Site**](Site.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSitesPage

> SitesPage GetSitesPage(ctx).Page(page).PageSize(pageSize).TenantId(tenantId).IncludeSubtenants(includeSubtenants).ServiceId(serviceId).ServiceType(serviceType).DeviceId(deviceId).ParentId(parentId).Type_(type_).ManagingControlPlaneId(managingControlPlaneId).ShowImage(showImage).Execute()

Returns a page of Sites. Only one filter is supported at a time.

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
    tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    includeSubtenants := true // bool |  (optional) (default to false)
    serviceId := "serviceId_example" // string |  (optional)
    serviceType := "serviceType_example" // string |  (optional)
    deviceId := "deviceId_example" // string |  (optional)
    parentId := "parentId_example" // string |  (optional)
    type_ := "type__example" // string |  (optional)
    managingControlPlaneId := "managingControlPlaneId_example" // string |  (optional)
    showImage := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.GetSitesPage(context.Background()).Page(page).PageSize(pageSize).TenantId(tenantId).IncludeSubtenants(includeSubtenants).ServiceId(serviceId).ServiceType(serviceType).DeviceId(deviceId).ParentId(parentId).Type_(type_).ManagingControlPlaneId(managingControlPlaneId).ShowImage(showImage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.GetSitesPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSitesPage`: SitesPage
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.GetSitesPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSitesPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **tenantId** | **string** |  | 
 **includeSubtenants** | **bool** |  | [default to false]
 **serviceId** | **string** |  | 
 **serviceType** | **string** |  | 
 **deviceId** | **string** |  | 
 **parentId** | **string** |  | 
 **type_** | **string** |  | 
 **managingControlPlaneId** | **string** |  | 
 **showImage** | **bool** |  | [default to false]

### Return type

[**SitesPage**](SitesPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveDevicesFromSite

> Site RemoveDevicesFromSite(ctx, id).RequestBody(requestBody).Execute()

Removes devices from a site.

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
    requestBody := []string{"Property_example"} // []string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.RemoveDevicesFromSite(context.Background(), id).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.RemoveDevicesFromSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveDevicesFromSite`: Site
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.RemoveDevicesFromSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDevicesFromSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** |  | 

### Return type

[**Site**](Site.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveServicesFromSite

> Site RemoveServicesFromSite(ctx, id).RequestBody(requestBody).Execute()

Remove services from a site.

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
    requestBody := []string{"Property_example"} // []string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.RemoveServicesFromSite(context.Background(), id).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.RemoveServicesFromSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveServicesFromSite`: Site
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.RemoveServicesFromSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveServicesFromSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** |  | 

### Return type

[**Site**](Site.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSite

> Site UpdateSite(ctx, id).SiteUpdate(siteUpdate).SendNotification(sendNotification).Execute()

Updates a site.

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
    siteUpdate := *openapiclient.NewSiteUpdate("Name_example") // SiteUpdate | 
    sendNotification := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SitesApi.UpdateSite(context.Background(), id).SiteUpdate(siteUpdate).SendNotification(sendNotification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesApi.UpdateSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSite`: Site
    fmt.Fprintf(os.Stdout, "Response from `SitesApi.UpdateSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteUpdate** | [**SiteUpdate**](SiteUpdate.md) |  | 
 **sendNotification** | **bool** |  | [default to false]

### Return type

[**Site**](Site.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

