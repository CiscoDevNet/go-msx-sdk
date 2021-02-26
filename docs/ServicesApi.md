# \ServicesApi

All URIs are relative to *http://localhost:8765*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteService**](ServicesApi.md#DeleteService) | **Delete** /manage/api/v8/services/{id} | Deletes a service.
[**GetService**](ServicesApi.md#GetService) | **Get** /manage/api/v8/services/{id} | Returns a service.
[**GetServicesPage**](ServicesApi.md#GetServicesPage) | **Get** /manage/api/v8/services | Returns a page of services.
[**SubmitOrder**](ServicesApi.md#SubmitOrder) | **Post** /manage/api/v8/services | Submits an order.
[**UpdateOrder**](ServicesApi.md#UpdateOrder) | **Put** /manage/api/v8/services | Updates an order.



## DeleteService

> DeleteService(ctx, id).Execute()

Deletes a service.

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
    resp, r, err := api_client.ServicesApi.DeleteService(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.DeleteService``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteServiceRequest struct via the builder pattern


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


## GetService

> Service GetService(ctx, id).Execute()

Returns a service.

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
    resp, r, err := api_client.ServicesApi.GetService(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.GetService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetService`: Service
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.GetService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Service**](Service.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServicesPage

> ServicesPage GetServicesPage(ctx).Page(page).PageSize(pageSize).TenantIds(tenantIds).Execute()

Returns a page of services.

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
    tenantIds := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.GetServicesPage(context.Background()).Page(page).PageSize(pageSize).TenantIds(tenantIds).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.GetServicesPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServicesPage`: ServicesPage
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.GetServicesPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **tenantIds** | **[]string** |  | 

### Return type

[**ServicesPage**](ServicesPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitOrder

> LegacyServiceOrderResponse SubmitOrder(ctx).ProductId(productId).OfferId(offerId).LegacyServiceOrder(legacyServiceOrder).Execute()

Submits an order.

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
    productId := TODO // string | 
    offerId := TODO // string | 
    legacyServiceOrder := *openapiclient.NewLegacyServiceOrder() // LegacyServiceOrder | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.SubmitOrder(context.Background()).ProductId(productId).OfferId(offerId).LegacyServiceOrder(legacyServiceOrder).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.SubmitOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitOrder`: LegacyServiceOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.SubmitOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | [**string**](string.md) |  | 
 **offerId** | [**string**](string.md) |  | 
 **legacyServiceOrder** | [**LegacyServiceOrder**](LegacyServiceOrder.md) |  | 

### Return type

[**LegacyServiceOrderResponse**](LegacyServiceOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrder

> LegacyServiceOrderResponse UpdateOrder(ctx).ProductId(productId).OfferId(offerId).LegacyServiceOrder(legacyServiceOrder).Execute()

Updates an order.

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
    productId := TODO // string | 
    offerId := TODO // string | 
    legacyServiceOrder := *openapiclient.NewLegacyServiceOrder() // LegacyServiceOrder | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.UpdateOrder(context.Background()).ProductId(productId).OfferId(offerId).LegacyServiceOrder(legacyServiceOrder).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.UpdateOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrder`: LegacyServiceOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.UpdateOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | [**string**](string.md) |  | 
 **offerId** | [**string**](string.md) |  | 
 **legacyServiceOrder** | [**LegacyServiceOrder**](LegacyServiceOrder.md) |  | 

### Return type

[**LegacyServiceOrderResponse**](LegacyServiceOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

