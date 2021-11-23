# \ValidationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetValidateProductVersionPage**](ValidationApi.md#GetValidateProductVersionPage) | **Get** /vulnerability/api/v8/vulnerabilities/validations | Returns a filtered page of validations.
[**ValidateProductVersion**](ValidationApi.md#ValidateProductVersion) | **Post** /vulnerability/api/v8/vulnerabilities/validations | Validate registered product / verison combinations for vulnerabilities.



## GetValidateProductVersionPage

> VulnerabilityValidationPage GetValidateProductVersionPage(ctx).Page(page).PageSize(pageSize).StartDate(startDate).EndDate(endDate).Execute()

Returns a filtered page of validations.

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
    page := int32(0) // int32 | 
    pageSize := int32(10) // int32 | 
    startDate := time.Now() // time.Time | Start date for date range filter on validation execution date. (optional)
    endDate := time.Now() // time.Time | End date for date range filter on validation execution date. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ValidationApi.GetValidateProductVersionPage(context.Background()).Page(page).PageSize(pageSize).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ValidationApi.GetValidateProductVersionPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetValidateProductVersionPage`: VulnerabilityValidationPage
    fmt.Fprintf(os.Stdout, "Response from `ValidationApi.GetValidateProductVersionPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetValidateProductVersionPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **startDate** | **time.Time** | Start date for date range filter on validation execution date. | 
 **endDate** | **time.Time** | End date for date range filter on validation execution date. | 

### Return type

[**VulnerabilityValidationPage**](VulnerabilityValidationPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateProductVersion

> VulnerabilityValidation ValidateProductVersion(ctx).Execute()

Validate registered product / verison combinations for vulnerabilities.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ValidationApi.ValidateProductVersion(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ValidationApi.ValidateProductVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateProductVersion`: VulnerabilityValidation
    fmt.Fprintf(os.Stdout, "Response from `ValidationApi.ValidateProductVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiValidateProductVersionRequest struct via the builder pattern


### Return type

[**VulnerabilityValidation**](VulnerabilityValidation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

