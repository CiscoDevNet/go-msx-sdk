# \HealthApi

All URIs are relative to *http://localhost:8765*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDevicesHealthList**](HealthApi.md#GetDevicesHealthList) | **Get** /monitor/api/v8/health/devices/list | 
[**GetServicesHealthList**](HealthApi.md#GetServicesHealthList) | **Get** /monitor/api/v8/health/services/list | 



## GetDevicesHealthList

> []ResourceHealth GetDevicesHealthList(ctx).Ids(ids).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HealthApi.GetDevicesHealthList(context.Background()).Ids(ids).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthApi.GetDevicesHealthList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDevicesHealthList`: []ResourceHealth
    fmt.Fprintf(os.Stdout, "Response from `HealthApi.GetDevicesHealthList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesHealthListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** |  | 

### Return type

[**[]ResourceHealth**](ResourceHealth.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServicesHealthList

> []ResourceHealth GetServicesHealthList(ctx).Ids(ids).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HealthApi.GetServicesHealthList(context.Background()).Ids(ids).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthApi.GetServicesHealthList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServicesHealthList`: []ResourceHealth
    fmt.Fprintf(os.Stdout, "Response from `HealthApi.GetServicesHealthList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesHealthListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** |  | 

### Return type

[**[]ResourceHealth**](ResourceHealth.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

