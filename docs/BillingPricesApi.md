# \BillingPricesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPrice**](BillingPricesApi.md#AddPrice) | **Post** /billing/api/v8/prices | Add price for tenant and event type.
[**DeletePrice**](BillingPricesApi.md#DeletePrice) | **Delete** /billing/api/v8/prices/{id} | Delete a price.
[**GetPrice**](BillingPricesApi.md#GetPrice) | **Get** /billing/api/v8/prices/{id} | Get a price.
[**GetPricesPage**](BillingPricesApi.md#GetPricesPage) | **Get** /billing/api/v8/prices | Retrieve a page of prices.
[**UpdatePrice**](BillingPricesApi.md#UpdatePrice) | **Put** /billing/api/v8/prices/{id} | Update price for an event type and tenant.



## AddPrice

> BillingPrice AddPrice(ctx).BillingPriceCreate(billingPriceCreate).Execute()

Add price for tenant and event type.



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
    billingPriceCreate := *openapiclient.NewBillingPriceCreate("Name_example", "Type_example", float64(123), "TenantId_example") // BillingPriceCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingPricesApi.AddPrice(context.Background()).BillingPriceCreate(billingPriceCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingPricesApi.AddPrice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPrice`: BillingPrice
    fmt.Fprintf(os.Stdout, "Response from `BillingPricesApi.AddPrice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **billingPriceCreate** | [**BillingPriceCreate**](BillingPriceCreate.md) |  | 

### Return type

[**BillingPrice**](BillingPrice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePrice

> DeletePrice(ctx, id).Execute()

Delete a price.



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
    resp, r, err := apiClient.BillingPricesApi.DeletePrice(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingPricesApi.DeletePrice``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeletePriceRequest struct via the builder pattern


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


## GetPrice

> BillingPrice GetPrice(ctx, id).Execute()

Get a price.



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
    resp, r, err := apiClient.BillingPricesApi.GetPrice(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingPricesApi.GetPrice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrice`: BillingPrice
    fmt.Fprintf(os.Stdout, "Response from `BillingPricesApi.GetPrice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BillingPrice**](BillingPrice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricesPage

> BillingPricesPage GetPricesPage(ctx).TenantId(tenantId).Page(page).PageSize(pageSize).Type_(type_).Subtype(subtype).Execute()

Retrieve a page of prices.



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
    type_ := "type__example" // string |  (optional)
    subtype := "subtype_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingPricesApi.GetPricesPage(context.Background()).TenantId(tenantId).Page(page).PageSize(pageSize).Type_(type_).Subtype(subtype).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingPricesApi.GetPricesPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPricesPage`: BillingPricesPage
    fmt.Fprintf(os.Stdout, "Response from `BillingPricesApi.GetPricesPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPricesPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **type_** | **string** |  | 
 **subtype** | **string** |  | 

### Return type

[**BillingPricesPage**](BillingPricesPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePrice

> BillingPrice UpdatePrice(ctx, id).BillingPriceUpdate(billingPriceUpdate).Execute()

Update price for an event type and tenant.



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
    billingPriceUpdate := *openapiclient.NewBillingPriceUpdate("Name_example", "Type_example", float64(123), "TenantId_example") // BillingPriceUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingPricesApi.UpdatePrice(context.Background(), id).BillingPriceUpdate(billingPriceUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingPricesApi.UpdatePrice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePrice`: BillingPrice
    fmt.Fprintf(os.Stdout, "Response from `BillingPricesApi.UpdatePrice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **billingPriceUpdate** | [**BillingPriceUpdate**](BillingPriceUpdate.md) |  | 

### Return type

[**BillingPrice**](BillingPrice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

