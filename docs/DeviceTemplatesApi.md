# \DeviceTemplatesApi

All URIs are relative to *http://localhost:8765*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceTemplate**](DeviceTemplatesApi.md#CreateDeviceTemplate) | **Post** /manage/api/v8/devicetemplates | Creates a device template.
[**DeleteDeviceTemplate**](DeviceTemplatesApi.md#DeleteDeviceTemplate) | **Delete** /manage/api/v8/devicetemplates/{id} | Deletes a device template.
[**GetDeviceTemplate**](DeviceTemplatesApi.md#GetDeviceTemplate) | **Get** /manage/api/v8/devicetemplates/{id} | Returns a device template.
[**GetDeviceTemplatesList**](DeviceTemplatesApi.md#GetDeviceTemplatesList) | **Get** /manage/api/v8/devicetemplates/list | Returns a list of device templates.
[**ScanDeviceTemplateParameters**](DeviceTemplatesApi.md#ScanDeviceTemplateParameters) | **Post** /manage/api/v8/devicetemplates/parameters/scan | API to scan parameters from the device template XML.
[**UpdateDeviceTemplateAccess**](DeviceTemplatesApi.md#UpdateDeviceTemplateAccess) | **Put** /manage/api/v8/devicetemplates/{id}/access | Updates device template access.



## CreateDeviceTemplate

> DeviceTemplate CreateDeviceTemplate(ctx).DeviceTemplateCreate(deviceTemplateCreate).Execute()

Creates a device template.

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
    deviceTemplateCreate := *openapiclient.NewDeviceTemplateCreate("Name_example", "ServiceType_example", "ConfigContent_example", "ResourceProvider_example") // DeviceTemplateCreate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceTemplatesApi.CreateDeviceTemplate(context.Background()).DeviceTemplateCreate(deviceTemplateCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceTemplatesApi.CreateDeviceTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceTemplate`: DeviceTemplate
    fmt.Fprintf(os.Stdout, "Response from `DeviceTemplatesApi.CreateDeviceTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceTemplateCreate** | [**DeviceTemplateCreate**](DeviceTemplateCreate.md) |  | 

### Return type

[**DeviceTemplate**](DeviceTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeviceTemplate

> DeleteDeviceTemplate(ctx, id).Execute()

Deletes a device template.

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
    resp, r, err := api_client.DeviceTemplatesApi.DeleteDeviceTemplate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceTemplatesApi.DeleteDeviceTemplate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteDeviceTemplateRequest struct via the builder pattern


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


## GetDeviceTemplate

> DeviceTemplate GetDeviceTemplate(ctx, id).Execute()

Returns a device template.

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
    resp, r, err := api_client.DeviceTemplatesApi.GetDeviceTemplate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceTemplatesApi.GetDeviceTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceTemplate`: DeviceTemplate
    fmt.Fprintf(os.Stdout, "Response from `DeviceTemplatesApi.GetDeviceTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeviceTemplate**](DeviceTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceTemplatesList

> []DeviceTemplate GetDeviceTemplatesList(ctx).ServiceType(serviceType).TenantId(tenantId).Execute()

Returns a list of device templates.

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
    serviceType := "manageddevice" // string |  (optional)
    tenantId := TODO // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceTemplatesApi.GetDeviceTemplatesList(context.Background()).ServiceType(serviceType).TenantId(tenantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceTemplatesApi.GetDeviceTemplatesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceTemplatesList`: []DeviceTemplate
    fmt.Fprintf(os.Stdout, "Response from `DeviceTemplatesApi.GetDeviceTemplatesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceTemplatesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceType** | **string** |  | 
 **tenantId** | [**string**](string.md) |  | 

### Return type

[**[]DeviceTemplate**](DeviceTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScanDeviceTemplateParameters

> []string ScanDeviceTemplateParameters(ctx).File(file).Execute()

API to scan parameters from the device template XML.

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
    file := os.NewFile(1234, "some_file") // *os.File | The XML template file of a device template (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceTemplatesApi.ScanDeviceTemplateParameters(context.Background()).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceTemplatesApi.ScanDeviceTemplateParameters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScanDeviceTemplateParameters`: []string
    fmt.Fprintf(os.Stdout, "Response from `DeviceTemplatesApi.ScanDeviceTemplateParameters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScanDeviceTemplateParametersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The XML template file of a device template | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceTemplateAccess

> DeviceTemplateAccessResponse UpdateDeviceTemplateAccess(ctx, id).DeviceTemplateAccess(deviceTemplateAccess).Execute()

Updates device template access.

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
    deviceTemplateAccess := *openapiclient.NewDeviceTemplateAccess() // DeviceTemplateAccess | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceTemplatesApi.UpdateDeviceTemplateAccess(context.Background(), id).DeviceTemplateAccess(deviceTemplateAccess).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceTemplatesApi.UpdateDeviceTemplateAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceTemplateAccess`: DeviceTemplateAccessResponse
    fmt.Fprintf(os.Stdout, "Response from `DeviceTemplatesApi.UpdateDeviceTemplateAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceTemplateAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceTemplateAccess** | [**DeviceTemplateAccess**](DeviceTemplateAccess.md) |  | 

### Return type

[**DeviceTemplateAccessResponse**](DeviceTemplateAccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

