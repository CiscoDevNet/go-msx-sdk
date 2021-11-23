# \LicensingAccountsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserAccountsList**](LicensingAccountsApi.md#GetUserAccountsList) | **Get** /licensing/api/v8/accounts/smart/list | Returns a filtered page of smart accounts.



## GetUserAccountsList

> SmartUserAccounts GetUserAccountsList(ctx).UserId(userId).Domain(domain).RoleName(roleName).Type_(type_).Execute()

Returns a filtered page of smart accounts.

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
    userId := "bdavis" // string | User Id
    domain := "abhtest001.cisco.com" // string | Smart Account domain (optional)
    roleName := "AccountAdmin" // string | Role Name (optional)
    type_ := openapiclient.SmartAccountType("CUSTOMER") // SmartAccountType |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LicensingAccountsApi.GetUserAccountsList(context.Background()).UserId(userId).Domain(domain).RoleName(roleName).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensingAccountsApi.GetUserAccountsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserAccountsList`: SmartUserAccounts
    fmt.Fprintf(os.Stdout, "Response from `LicensingAccountsApi.GetUserAccountsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAccountsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | User Id | 
 **domain** | **string** | Smart Account domain | 
 **roleName** | **string** | Role Name | 
 **type_** | [**SmartAccountType**](SmartAccountType.md) |  | 

### Return type

[**SmartUserAccounts**](SmartUserAccounts.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

