# \LicensingLicensesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLicensesList**](LicensingLicensesApi.md#GetLicensesList) | **Get** /licensing/api/v8/licensing/api/v8/licenses/list | Returns a filtered list of licenses.



## GetLicensesList

> []LicenseSummary GetLicensesList(ctx).SmartAccountId(smartAccountId).VirtualAccountId(virtualAccountId).Execute()

Returns a filtered list of licenses.

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
    smartAccountId := "295183" // string | Smart Account Identifier
    virtualAccountId := "123123" // string | Virtual Account Identifier

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicensingLicensesApi.GetLicensesList(context.Background()).SmartAccountId(smartAccountId).VirtualAccountId(virtualAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensingLicensesApi.GetLicensesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicensesList`: []LicenseSummary
    fmt.Fprintf(os.Stdout, "Response from `LicensingLicensesApi.GetLicensesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLicensesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smartAccountId** | **string** | Smart Account Identifier | 
 **virtualAccountId** | **string** | Virtual Account Identifier | 

### Return type

[**[]LicenseSummary**](LicenseSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

