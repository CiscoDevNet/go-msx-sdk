# \DevicesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachDeviceTemplates**](DevicesApi.md#AttachDeviceTemplates) | **Post** /manage/api/v8/devices/{id}/templates | Attaches one or more device templates to a device instance.
[**BatchAttachDeviceTemplates**](DevicesApi.md#BatchAttachDeviceTemplates) | **Post** /manage/api/v8/devices/templates/attach | Attaches one or more device templates to a batch of device instances.
[**CreateDevice**](DevicesApi.md#CreateDevice) | **Post** /manage/api/v8/devices | Creates a device.
[**DeleteDevice**](DevicesApi.md#DeleteDevice) | **Delete** /manage/api/v8/devices/{id} | Deletes a device.
[**DetachDeviceTemplate**](DevicesApi.md#DetachDeviceTemplate) | **Delete** /manage/api/v8/devices/{id}/templates/{templateId} | Detaches a template from a device.
[**DetachDeviceTemplates**](DevicesApi.md#DetachDeviceTemplates) | **Delete** /manage/api/v8/devices/{id}/templates | Detach device templates that are already attached to a device.
[**GetDevice**](DevicesApi.md#GetDevice) | **Get** /manage/api/v8/devices/{id} | Returns a device.
[**GetDeviceConfig**](DevicesApi.md#GetDeviceConfig) | **Get** /manage/api/v8/devices/{id}/config | Returns the running configuration for a device.
[**GetDeviceTemplateHistory**](DevicesApi.md#GetDeviceTemplateHistory) | **Get** /manage/api/v8/devices/{id}/templates | Returns device template history.
[**GetDevicesPage**](DevicesApi.md#GetDevicesPage) | **Get** /manage/api/v8/devices | Returns a page of devices.
[**PatchDevice**](DevicesApi.md#PatchDevice) | **Patch** /manage/api/v8/devices/{id} | Update a device.
[**RedeployDevice**](DevicesApi.md#RedeployDevice) | **Post** /manage/api/v8/devices/{id}/redeploy | Dedeploys a device.
[**UpdateDevice**](DevicesApi.md#UpdateDevice) | **Put** /manage/api/v8/devices/{id} | Update a device.
[**UpdateDeviceTemplates**](DevicesApi.md#UpdateDeviceTemplates) | **Put** /manage/api/v8/devices/{id}/templates | Update device templates that are already attached to a device.



## AttachDeviceTemplates

> []DeviceTemplateHistory AttachDeviceTemplates(ctx, id).DeviceTemplateAttachRequest(deviceTemplateAttachRequest).Execute()

Attaches one or more device templates to a device instance.

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
    deviceTemplateAttachRequest := *openapiclient.NewDeviceTemplateAttachRequest() // DeviceTemplateAttachRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.AttachDeviceTemplates(context.Background(), id).DeviceTemplateAttachRequest(deviceTemplateAttachRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.AttachDeviceTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachDeviceTemplates`: []DeviceTemplateHistory
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.AttachDeviceTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachDeviceTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceTemplateAttachRequest** | [**DeviceTemplateAttachRequest**](DeviceTemplateAttachRequest.md) |  | 

### Return type

[**[]DeviceTemplateHistory**](DeviceTemplateHistory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchAttachDeviceTemplates

> [][]DeviceTemplateHistorySummary BatchAttachDeviceTemplates(ctx).DeviceTemplateBatchAttachRequest(deviceTemplateBatchAttachRequest).Execute()

Attaches one or more device templates to a batch of device instances.

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
    deviceTemplateBatchAttachRequest := *openapiclient.NewDeviceTemplateBatchAttachRequest() // DeviceTemplateBatchAttachRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.BatchAttachDeviceTemplates(context.Background()).DeviceTemplateBatchAttachRequest(deviceTemplateBatchAttachRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.BatchAttachDeviceTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchAttachDeviceTemplates`: [][]DeviceTemplateHistorySummary
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.BatchAttachDeviceTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchAttachDeviceTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceTemplateBatchAttachRequest** | [**DeviceTemplateBatchAttachRequest**](DeviceTemplateBatchAttachRequest.md) |  | 

### Return type

[**[][]DeviceTemplateHistorySummary**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDevice

> Device CreateDevice(ctx).DeviceCreate(deviceCreate).Execute()

Creates a device.

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
    deviceCreate := *openapiclient.NewDeviceCreate("TenantId_example", false, "OnboardType_example", "Name_example", "Model_example", "Type_example") // DeviceCreate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.CreateDevice(context.Background()).DeviceCreate(deviceCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CreateDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDevice`: Device
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.CreateDevice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceCreate** | [**DeviceCreate**](DeviceCreate.md) |  | 

### Return type

[**Device**](Device.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDevice

> DeleteDevice(ctx, id).Execute()

Deletes a device.

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
    resp, r, err := api_client.DevicesApi.DeleteDevice(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DeleteDevice``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteDeviceRequest struct via the builder pattern


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


## DetachDeviceTemplate

> []DeviceTemplateHistory DetachDeviceTemplate(ctx, id, templateId).Execute()

Detaches a template from a device.

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
    templateId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.DetachDeviceTemplate(context.Background(), id, templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DetachDeviceTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DetachDeviceTemplate`: []DeviceTemplateHistory
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.DetachDeviceTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**templateId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachDeviceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]DeviceTemplateHistory**](DeviceTemplateHistory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachDeviceTemplates

> []DeviceTemplateHistory DetachDeviceTemplates(ctx, id).Execute()

Detach device templates that are already attached to a device.

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
    resp, r, err := api_client.DevicesApi.DetachDeviceTemplates(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DetachDeviceTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DetachDeviceTemplates`: []DeviceTemplateHistory
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.DetachDeviceTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachDeviceTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DeviceTemplateHistory**](DeviceTemplateHistory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevice

> Device GetDevice(ctx, id).Execute()

Returns a device.

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
    resp, r, err := api_client.DevicesApi.GetDevice(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDevice`: Device
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Device**](Device.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceConfig

> string GetDeviceConfig(ctx, id).Execute()

Returns the running configuration for a device.

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
    resp, r, err := api_client.DevicesApi.GetDeviceConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetDeviceConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceConfig`: string
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetDeviceConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceTemplateHistory

> []DeviceTemplateHistory GetDeviceTemplateHistory(ctx, id).TemplateId(templateId).Execute()

Returns device template history.

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
    templateId := TODO // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.GetDeviceTemplateHistory(context.Background(), id).TemplateId(templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetDeviceTemplateHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceTemplateHistory`: []DeviceTemplateHistory
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetDeviceTemplateHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceTemplateHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templateId** | [**string**](string.md) |  | 

### Return type

[**[]DeviceTemplateHistory**](DeviceTemplateHistory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevicesPage

> DevicesPage GetDevicesPage(ctx).Page(page).PageSize(pageSize).DeviceIds(deviceIds).ServiceIds(serviceIds).Types(types).SerialKeys(serialKeys).ServiceTypes(serviceTypes).Models(models).Subtypes(subtypes).Names(names).Versions(versions).TenantIds(tenantIds).IncludeSubtenants(includeSubtenants).Severities(severities).ComplianceStates(complianceStates).VulnerabilityStates(vulnerabilityStates).SortBy(sortBy).SortOrder(sortOrder).Execute()

Returns a page of devices.

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
    deviceIds := []string{"Inner_example"} // []string |  (optional)
    serviceIds := []string{"Inner_example"} // []string |  (optional)
    types := []string{"Inner_example"} // []string |  (optional)
    serialKeys := []string{"Inner_example"} // []string |  (optional)
    serviceTypes := []string{"Inner_example"} // []string |  (optional)
    models := []string{"Inner_example"} // []string |  (optional)
    subtypes := []string{"Inner_example"} // []string |  (optional)
    names := []string{"Inner_example"} // []string |  (optional)
    versions := []string{"Inner_example"} // []string |  (optional)
    tenantIds := []string{"Inner_example"} // []string |  (optional)
    includeSubtenants := true // bool |  (optional) (default to false)
    severities := []string{"Inner_example"} // []string |  (optional)
    complianceStates := []openapiclient.DeviceComplianceState{openapiclient.DeviceComplianceState("COMPLIANT")} // []DeviceComplianceState |  (optional)
    vulnerabilityStates := []openapiclient.DeviceVulnerabilityState{openapiclient.DeviceVulnerabilityState("VULNERABLE")} // []DeviceVulnerabilityState |  (optional)
    sortBy := "name" // string |  (optional)
    sortOrder := "sortOrder_example" // string |  (optional) (default to "asc")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.GetDevicesPage(context.Background()).Page(page).PageSize(pageSize).DeviceIds(deviceIds).ServiceIds(serviceIds).Types(types).SerialKeys(serialKeys).ServiceTypes(serviceTypes).Models(models).Subtypes(subtypes).Names(names).Versions(versions).TenantIds(tenantIds).IncludeSubtenants(includeSubtenants).Severities(severities).ComplianceStates(complianceStates).VulnerabilityStates(vulnerabilityStates).SortBy(sortBy).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.GetDevicesPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDevicesPage`: DevicesPage
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.GetDevicesPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **deviceIds** | **[]string** |  | 
 **serviceIds** | **[]string** |  | 
 **types** | **[]string** |  | 
 **serialKeys** | **[]string** |  | 
 **serviceTypes** | **[]string** |  | 
 **models** | **[]string** |  | 
 **subtypes** | **[]string** |  | 
 **names** | **[]string** |  | 
 **versions** | **[]string** |  | 
 **tenantIds** | **[]string** |  | 
 **includeSubtenants** | **bool** |  | [default to false]
 **severities** | **[]string** |  | 
 **complianceStates** | [**[]DeviceComplianceState**](DeviceComplianceState.md) |  | 
 **vulnerabilityStates** | [**[]DeviceVulnerabilityState**](DeviceVulnerabilityState.md) |  | 
 **sortBy** | **string** |  | 
 **sortOrder** | **string** |  | [default to &quot;asc&quot;]

### Return type

[**DevicesPage**](DevicesPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchDevice

> Device PatchDevice(ctx, id).DevicePatch(devicePatch).Execute()

Update a device.

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
    devicePatch := *openapiclient.NewDevicePatch() // DevicePatch | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.PatchDevice(context.Background(), id).DevicePatch(devicePatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.PatchDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchDevice`: Device
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.PatchDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **devicePatch** | [**DevicePatch**](DevicePatch.md) |  | 

### Return type

[**Device**](Device.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RedeployDevice

> RedeployDevice(ctx, id).Execute()

Dedeploys a device.

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
    resp, r, err := api_client.DevicesApi.RedeployDevice(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.RedeployDevice``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRedeployDeviceRequest struct via the builder pattern


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


## UpdateDevice

> Device UpdateDevice(ctx, id).DeviceUpdate(deviceUpdate).Execute()

Update a device.

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
    deviceUpdate := *openapiclient.NewDeviceUpdate(false, "OnboardType_example", "Name_example", "Model_example", "Type_example") // DeviceUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.UpdateDevice(context.Background(), id).DeviceUpdate(deviceUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.UpdateDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDevice`: Device
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.UpdateDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceUpdate** | [**DeviceUpdate**](DeviceUpdate.md) |  | 

### Return type

[**Device**](Device.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceTemplates

> []DeviceTemplateHistory UpdateDeviceTemplates(ctx, id).DeviceTemplateUpdateRequest(deviceTemplateUpdateRequest).Execute()

Update device templates that are already attached to a device.

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
    deviceTemplateUpdateRequest := *openapiclient.NewDeviceTemplateUpdateRequest() // DeviceTemplateUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesApi.UpdateDeviceTemplates(context.Background(), id).DeviceTemplateUpdateRequest(deviceTemplateUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.UpdateDeviceTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceTemplates`: []DeviceTemplateHistory
    fmt.Fprintf(os.Stdout, "Response from `DevicesApi.UpdateDeviceTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceTemplateUpdateRequest** | [**DeviceTemplateUpdateRequest**](DeviceTemplateUpdateRequest.md) |  | 

### Return type

[**[]DeviceTemplateHistory**](DeviceTemplateHistory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

