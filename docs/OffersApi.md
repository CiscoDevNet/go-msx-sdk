# \OffersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOffer**](OffersApi.md#CreateOffer) | **Post** /consume/api/v8/offers | Creates a product offer.
[**DeleteOffer**](OffersApi.md#DeleteOffer) | **Delete** /consume/api/v8/offers/{id} | Deletes a product offer
[**GetOffer**](OffersApi.md#GetOffer) | **Get** /consume/api/v8/offers/{id} | Returns a product offer.
[**GetOfferAssignmentsList**](OffersApi.md#GetOfferAssignmentsList) | **Get** /consume/api/v8/offers/{id}/assignments/list | Returns a list of tenant assignments for a product offer.
[**GetOffersCount**](OffersApi.md#GetOffersCount) | **Get** /consume/api/v8/offers/count | Returns the number of product offers.
[**GetOffersPage**](OffersApi.md#GetOffersPage) | **Get** /consume/api/v8/offers | Returns a page of product offers.
[**UpdateOffer**](OffersApi.md#UpdateOffer) | **Put** /consume/api/v8/offers/{id} | Updates a product offer.
[**UpdateOfferAssignments**](OffersApi.md#UpdateOfferAssignments) | **Put** /consume/api/v8/offers/{id}/assignments | Updates the tenant assignemnts for a product offer.



## CreateOffer

> Offer CreateOffer(ctx).OfferCreate(offerCreate).Execute()

Creates a product offer.

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
    offerCreate := *openapiclient.NewOfferCreate("Name_example", "Label_example", "Description_example", "ProductId_example", int32(123), int32(123), "Image_example", []string{"SupportedProperties_example"}, []openapiclient.NameValue{*openapiclient.NewNameValue()}) // OfferCreate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OffersApi.CreateOffer(context.Background()).OfferCreate(offerCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OffersApi.CreateOffer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOffer`: Offer
    fmt.Fprintf(os.Stdout, "Response from `OffersApi.CreateOffer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOfferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offerCreate** | [**OfferCreate**](OfferCreate.md) |  | 

### Return type

[**Offer**](Offer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOffer

> DeleteOffer(ctx, id).Execute()

Deletes a product offer

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
    resp, r, err := api_client.OffersApi.DeleteOffer(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OffersApi.DeleteOffer``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteOfferRequest struct via the builder pattern


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


## GetOffer

> Offer GetOffer(ctx, id).Execute()

Returns a product offer.

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
    resp, r, err := api_client.OffersApi.GetOffer(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OffersApi.GetOffer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOffer`: Offer
    fmt.Fprintf(os.Stdout, "Response from `OffersApi.GetOffer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOfferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Offer**](Offer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOfferAssignmentsList

> []CatalogAssignment GetOfferAssignmentsList(ctx, id).Execute()

Returns a list of tenant assignments for a product offer.

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
    resp, r, err := api_client.OffersApi.GetOfferAssignmentsList(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OffersApi.GetOfferAssignmentsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOfferAssignmentsList`: []CatalogAssignment
    fmt.Fprintf(os.Stdout, "Response from `OffersApi.GetOfferAssignmentsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOfferAssignmentsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CatalogAssignment**](CatalogAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOffersCount

> int64 GetOffersCount(ctx).ProductId(productId).TenantId(tenantId).Execute()

Returns the number of product offers.

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
    productId := TODO // string |  (optional)
    tenantId := TODO // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OffersApi.GetOffersCount(context.Background()).ProductId(productId).TenantId(tenantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OffersApi.GetOffersCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOffersCount`: int64
    fmt.Fprintf(os.Stdout, "Response from `OffersApi.GetOffersCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOffersCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | [**string**](string.md) |  | 
 **tenantId** | [**string**](string.md) |  | 

### Return type

**int64**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOffersPage

> OffersPage GetOffersPage(ctx).Page(page).PageSize(pageSize).ProductId(productId).TenantIds(tenantIds).Execute()

Returns a page of product offers.

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
    productId := TODO // string |  (optional)
    tenantIds := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OffersApi.GetOffersPage(context.Background()).Page(page).PageSize(pageSize).ProductId(productId).TenantIds(tenantIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OffersApi.GetOffersPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOffersPage`: OffersPage
    fmt.Fprintf(os.Stdout, "Response from `OffersApi.GetOffersPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOffersPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **productId** | [**string**](string.md) |  | 
 **tenantIds** | **[]string** |  | 

### Return type

[**OffersPage**](OffersPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOffer

> Offer UpdateOffer(ctx, id).OfferUpdate(offerUpdate).Execute()

Updates a product offer.

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
    offerUpdate := *openapiclient.NewOfferUpdate("Name_example", "Label_example", "Description_example", "ProductId_example", int32(123), int32(123), "Image_example", []string{"SupportedProperties_example"}, []openapiclient.NameValue{*openapiclient.NewNameValue()}) // OfferUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OffersApi.UpdateOffer(context.Background(), id).OfferUpdate(offerUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OffersApi.UpdateOffer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOffer`: Offer
    fmt.Fprintf(os.Stdout, "Response from `OffersApi.UpdateOffer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOfferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offerUpdate** | [**OfferUpdate**](OfferUpdate.md) |  | 

### Return type

[**Offer**](Offer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOfferAssignments

> []CatalogAssignment UpdateOfferAssignments(ctx, id).RequestBody(requestBody).Execute()

Updates the tenant assignemnts for a product offer.

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
    requestBody := []string{"Property_example"} // []string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OffersApi.UpdateOfferAssignments(context.Background(), id).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OffersApi.UpdateOfferAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOfferAssignments`: []CatalogAssignment
    fmt.Fprintf(os.Stdout, "Response from `OffersApi.UpdateOfferAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOfferAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** |  | 

### Return type

[**[]CatalogAssignment**](CatalogAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

