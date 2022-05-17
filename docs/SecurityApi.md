# \SecurityApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccessToken**](SecurityApi.md#GetAccessToken) | **Post** /idm/v2/token | Returns an access token.



## GetAccessToken

> AccessToken GetAccessToken(ctx).Authorization(authorization).GrantType(grantType).Username(username).Password(password).AccessToken(accessToken).SwitchUsername(switchUsername).TenantId(tenantId).Scope(scope).Nonce(nonce).TenantName(tenantName).Execute()

Returns an access token.

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
    authorization := "authorization_example" // string | 
    grantType := "grantType_example" // string | 
    username := "username_example" // string |  (optional)
    password := "password_example" // string |  (optional)
    accessToken := "accessToken_example" // string |  (optional)
    switchUsername := "switchUsername_example" // string |  (optional)
    tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    scope := "scope_example" // string |  (optional)
    nonce := "nonce_example" // string |  (optional)
    tenantName := "tenantName_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.GetAccessToken(context.Background()).Authorization(authorization).GrantType(grantType).Username(username).Password(password).AccessToken(accessToken).SwitchUsername(switchUsername).TenantId(tenantId).Scope(scope).Nonce(nonce).TenantName(tenantName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.GetAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessToken`: AccessToken
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.GetAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **grantType** | **string** |  | 
 **username** | **string** |  | 
 **password** | **string** |  | 
 **accessToken** | **string** |  | 
 **switchUsername** | **string** |  | 
 **tenantId** | **string** |  | 
 **scope** | **string** |  | 
 **nonce** | **string** |  | 
 **tenantName** | **string** |  | 

### Return type

[**AccessToken**](AccessToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

