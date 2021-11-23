# \VulnerabilitiesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIngestVulnerabilitiesTasksPage**](VulnerabilitiesApi.md#GetIngestVulnerabilitiesTasksPage) | **Get** /vulnerability/api/v8/vulnerabilities/ingests | Returns a filtered page of ingest tasks.
[**GetVulnerabilitiesPage**](VulnerabilitiesApi.md#GetVulnerabilitiesPage) | **Get** /vulnerability/api/v8/vulnerabilities | Returns a filtered page of vulnerabilities.
[**IngestVulnerabilities**](VulnerabilitiesApi.md#IngestVulnerabilities) | **Post** /vulnerability/api/v8/vulnerabilities/ingests | Ingests a CVE JSON feed into the Vulnerability Service datastore.



## GetIngestVulnerabilitiesTasksPage

> VulnerabilityIngestPage GetIngestVulnerabilitiesTasksPage(ctx).Page(page).PageSize(pageSize).StartDate(startDate).EndDate(endDate).Execute()

Returns a filtered page of ingest tasks.

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
    resp, r, err := api_client.VulnerabilitiesApi.GetIngestVulnerabilitiesTasksPage(context.Background()).Page(page).PageSize(pageSize).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VulnerabilitiesApi.GetIngestVulnerabilitiesTasksPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIngestVulnerabilitiesTasksPage`: VulnerabilityIngestPage
    fmt.Fprintf(os.Stdout, "Response from `VulnerabilitiesApi.GetIngestVulnerabilitiesTasksPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIngestVulnerabilitiesTasksPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **startDate** | **time.Time** | Start date for date range filter on validation execution date. | 
 **endDate** | **time.Time** | End date for date range filter on validation execution date. | 

### Return type

[**VulnerabilityIngestPage**](VulnerabilityIngestPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVulnerabilitiesPage

> VulnerabilitiesPage GetVulnerabilitiesPage(ctx).Page(page).PageSize(pageSize).CveId(cveId).Vendor(vendor).Product(product).Version(version).Severity(severity).StartDate(startDate).EndDate(endDate).Year(year).SortBy(sortBy).SortOrder(sortOrder).Execute()

Returns a filtered page of vulnerabilities.

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
    cveId := "CVE-2021-0202" // string | CVE identifer (https://www.cvedetails.com/cve-help.php) to filter by. (optional)
    vendor := "cisco" // string | Vendor identifier (as defined in NIST's CPE dictionary) to filter by. (optional)
    product := "ios_xe" // string | Product identifier (as defined in NIST's CPE dictionary) to filter by. (optional)
    version := "12.3" // string | Product version (as defined in NIST's CPE dictionary) filter to filter by. (optional)
    severity := openapiclient.VulnerabilitySeverity("NONE") // VulnerabilitySeverity | A CVSS severity level (https://nvd.nist.gov/vuln-metrics/cvss) to filter by. (optional)
    startDate := time.Now() // time.Time | Start date for date range filter on CVE published date. (optional)
    endDate := time.Now() // time.Time | End date for date range filter on CVE published date. (optional)
    year := int32(2019) // int32 | Year CVE published filter. (optional)
    sortBy := "sortBy_example" // string |  (optional) (default to "publishedOn")
    sortOrder := "sortOrder_example" // string |  (optional) (default to "asc")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VulnerabilitiesApi.GetVulnerabilitiesPage(context.Background()).Page(page).PageSize(pageSize).CveId(cveId).Vendor(vendor).Product(product).Version(version).Severity(severity).StartDate(startDate).EndDate(endDate).Year(year).SortBy(sortBy).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VulnerabilitiesApi.GetVulnerabilitiesPage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVulnerabilitiesPage`: VulnerabilitiesPage
    fmt.Fprintf(os.Stdout, "Response from `VulnerabilitiesApi.GetVulnerabilitiesPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVulnerabilitiesPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **cveId** | **string** | CVE identifer (https://www.cvedetails.com/cve-help.php) to filter by. | 
 **vendor** | **string** | Vendor identifier (as defined in NIST&#39;s CPE dictionary) to filter by. | 
 **product** | **string** | Product identifier (as defined in NIST&#39;s CPE dictionary) to filter by. | 
 **version** | **string** | Product version (as defined in NIST&#39;s CPE dictionary) filter to filter by. | 
 **severity** | [**VulnerabilitySeverity**](VulnerabilitySeverity.md) | A CVSS severity level (https://nvd.nist.gov/vuln-metrics/cvss) to filter by. | 
 **startDate** | **time.Time** | Start date for date range filter on CVE published date. | 
 **endDate** | **time.Time** | End date for date range filter on CVE published date. | 
 **year** | **int32** | Year CVE published filter. | 
 **sortBy** | **string** |  | [default to &quot;publishedOn&quot;]
 **sortOrder** | **string** |  | [default to &quot;asc&quot;]

### Return type

[**VulnerabilitiesPage**](VulnerabilitiesPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IngestVulnerabilities

> VulnerabilityIngestion IngestVulnerabilities(ctx).VulnerabilityFeed(vulnerabilityFeed).Execute()

Ingests a CVE JSON feed into the Vulnerability Service datastore.

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
    vulnerabilityFeed := *openapiclient.NewVulnerabilityFeed() // VulnerabilityFeed | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VulnerabilitiesApi.IngestVulnerabilities(context.Background()).VulnerabilityFeed(vulnerabilityFeed).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VulnerabilitiesApi.IngestVulnerabilities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IngestVulnerabilities`: VulnerabilityIngestion
    fmt.Fprintf(os.Stdout, "Response from `VulnerabilitiesApi.IngestVulnerabilities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIngestVulnerabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vulnerabilityFeed** | [**VulnerabilityFeed**](VulnerabilityFeed.md) |  | 

### Return type

[**VulnerabilityIngestion**](VulnerabilityIngestion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

