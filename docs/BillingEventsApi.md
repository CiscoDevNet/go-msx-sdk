# \BillingEventsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCostSummary**](BillingEventsApi.md#GetCostSummary) | **Get** /billing/api/v8/events/costs | Retrieve a summary for tenant cost.
[**GetEvent**](BillingEventsApi.md#GetEvent) | **Get** /billing/api/v8/events/{id} | Get an Event.
[**GetEventsPage**](BillingEventsApi.md#GetEventsPage) | **Get** /billing/api/v8/events | Retrieve a page of events for tenant.



## GetCostSummary

> BillingCostsReport GetCostSummary(ctx).TenantId(tenantId).FromDate(fromDate).ToDate(toDate).GroupBy(groupBy).Execute()

Retrieve a summary for tenant cost.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    tenantId := TODO // string | 
    fromDate := time.Now() // time.Time |  (optional)
    toDate := time.Now() // time.Time |  (optional)
    groupBy := "type, subtype or service" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingEventsApi.GetCostSummary(context.Background()).TenantId(tenantId).FromDate(fromDate).ToDate(toDate).GroupBy(groupBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingEventsApi.GetCostSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCostSummary`: BillingCostsReport
    fmt.Fprintf(os.Stdout, "Response from `BillingEventsApi.GetCostSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCostSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | [**string**](string.md) |  | 
 **fromDate** | **time.Time** |  | 
 **toDate** | **time.Time** |  | 
 **groupBy** | **string** |  | 

### Return type

[**BillingCostsReport**](BillingCostsReport.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvent

> BillingEvent GetEvent(ctx, id).Execute()

Get an Event.



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
    resp, r, err := api_client.BillingEventsApi.GetEvent(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingEventsApi.GetEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEvent`: BillingEvent
    fmt.Fprintf(os.Stdout, "Response from `BillingEventsApi.GetEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BillingEvent**](BillingEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventsPage

> BillingEventsPage GetEventsPage(ctx).TenantId(tenantId).Page(page).PageSize(pageSize).FromDate(fromDate).ToDate(toDate).Type_(type_).Subtype(subtype).Execute()

Retrieve a page of events for tenant.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    tenantId := TODO // string | 
    page := int32(0) // int32 | 
    pageSize := int32(10) // int32 | 
    fromDate := time.Now() // time.Time |  (optional)
    toDate := time.Now() // time.Time |  (optional)
    type_ := "type__example" // string |  (optional)
    subtype := "subtype_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillingEventsApi.GetEventsPage(context.Background()).TenantId(tenantId).Page(page).PageSize(pageSize).FromDate(fromDate).ToDate(toDate).Type_(type_).Subtype(subtype).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingEventsApi.GetEventsPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventsPage`: BillingEventsPage
    fmt.Fprintf(os.Stdout, "Response from `BillingEventsApi.GetEventsPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | [**string**](string.md) |  | 
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **fromDate** | **time.Time** |  | 
 **toDate** | **time.Time** |  | 
 **type_** | **string** |  | 
 **subtype** | **string** |  | 

### Return type

[**BillingEventsPage**](BillingEventsPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

