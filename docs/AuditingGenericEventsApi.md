# \AuditingGenericEventsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGenericEvent**](AuditingGenericEventsApi.md#CreateGenericEvent) | **Post** /auditing/api/v8/genericevents | Create Generic Event



## CreateGenericEvent

> GenericEvent CreateGenericEvent(ctx).GenericEventCreate(genericEventCreate).Execute()

Create Generic Event



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
    genericEventCreate := *openapiclient.NewGenericEventCreate() // GenericEventCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuditingGenericEventsApi.CreateGenericEvent(context.Background()).GenericEventCreate(genericEventCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditingGenericEventsApi.CreateGenericEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGenericEvent`: GenericEvent
    fmt.Fprintf(os.Stdout, "Response from `AuditingGenericEventsApi.CreateGenericEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGenericEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **genericEventCreate** | [**GenericEventCreate**](GenericEventCreate.md) |  | 

### Return type

[**GenericEvent**](GenericEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

