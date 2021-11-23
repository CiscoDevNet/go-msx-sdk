# \IncidentsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelIncident**](IncidentsApi.md#CancelIncident) | **Patch** /incident/api/v1/incidents/{id}/cancel | Cancels an incident.
[**CreateIncident**](IncidentsApi.md#CreateIncident) | **Post** /incident/api/v1/incidents | Create a new Incident.
[**DeleteIncident**](IncidentsApi.md#DeleteIncident) | **Delete** /incident/api/v1/incidents/{id} | Deletes an incident.
[**GetIncident**](IncidentsApi.md#GetIncident) | **Get** /incident/api/v1/incidents/{id} | Get incident details.
[**GetIncidents**](IncidentsApi.md#GetIncidents) | **Get** /incident/api/v1/incidents | Get Incidents by filter.
[**UpdateIncident**](IncidentsApi.md#UpdateIncident) | **Put** /incident/api/v1/incidents/{id} | Updates an incident.



## CancelIncident

> Incident CancelIncident(ctx, id).IncidentCancel(incidentCancel).Execute()

Cancels an incident.

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
    incidentCancel := *openapiclient.NewIncidentCancel("Tenant_example") // IncidentCancel | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentsApi.CancelIncident(context.Background(), id).IncidentCancel(incidentCancel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CancelIncident``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelIncident`: Incident
    fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.CancelIncident`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelIncidentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **incidentCancel** | [**IncidentCancel**](IncidentCancel.md) |  | 

### Return type

[**Incident**](Incident.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIncident

> Incident CreateIncident(ctx).IncidentCreate(incidentCreate).Execute()

Create a new Incident.

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
    incidentCreate := *openapiclient.NewIncidentCreate("Description_example") // IncidentCreate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentsApi.CreateIncident(context.Background()).IncidentCreate(incidentCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CreateIncident``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIncident`: Incident
    fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.CreateIncident`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIncidentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **incidentCreate** | [**IncidentCreate**](IncidentCreate.md) |  | 

### Return type

[**Incident**](Incident.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIncident

> DeleteIncident(ctx, id).Execute()

Deletes an incident.

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
    resp, r, err := api_client.IncidentsApi.DeleteIncident(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.DeleteIncident``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteIncidentRequest struct via the builder pattern


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


## GetIncident

> Incident GetIncident(ctx, id).Execute()

Get incident details.

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
    resp, r, err := api_client.IncidentsApi.GetIncident(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.GetIncident``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIncident`: Incident
    fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.GetIncident`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIncidentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Incident**](Incident.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIncidents

> IncidentsPage GetIncidents(ctx).Page(page).PageSize(pageSize).Id(id).ExternalId(externalId).TenantId(tenantId).Category(category).Subcategory(subcategory).State(state).Priority(priority).Severity(severity).SortBy(sortBy).SortOrder(sortOrder).Execute()

Get Incidents by filter.

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
    id := "id_example" // string |  (optional)
    externalId := "externalId_example" // string |  (optional)
    tenantId := "tenantId_example" // string |  (optional)
    category := "category_example" // string |  (optional)
    subcategory := "subcategory_example" // string |  (optional)
    state := "state_example" // string | New|InProgress|OnHold|Resolved|Canceled (optional)
    priority := "priority_example" // string | Critical|Low|High|Moderate|Planning (optional)
    severity := "severity_example" // string | High|Medium|Low (optional)
    sortBy := "sortBy_example" // string | category|subcategory|priority|severity|state (optional)
    sortOrder := "sortOrder_example" // string | ASC/DESC (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentsApi.GetIncidents(context.Background()).Page(page).PageSize(pageSize).Id(id).ExternalId(externalId).TenantId(tenantId).Category(category).Subcategory(subcategory).State(state).Priority(priority).Severity(severity).SortBy(sortBy).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.GetIncidents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIncidents`: IncidentsPage
    fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.GetIncidents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIncidentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **id** | **string** |  | 
 **externalId** | **string** |  | 
 **tenantId** | **string** |  | 
 **category** | **string** |  | 
 **subcategory** | **string** |  | 
 **state** | **string** | New|InProgress|OnHold|Resolved|Canceled | 
 **priority** | **string** | Critical|Low|High|Moderate|Planning | 
 **severity** | **string** | High|Medium|Low | 
 **sortBy** | **string** | category|subcategory|priority|severity|state | 
 **sortOrder** | **string** | ASC/DESC | 

### Return type

[**IncidentsPage**](IncidentsPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIncident

> Incident UpdateIncident(ctx, id).IncidentUpdate(incidentUpdate).Execute()

Updates an incident.

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
    incidentUpdate := *openapiclient.NewIncidentUpdate("Description_example") // IncidentUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentsApi.UpdateIncident(context.Background(), id).IncidentUpdate(incidentUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.UpdateIncident``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIncident`: Incident
    fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.UpdateIncident`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIncidentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **incidentUpdate** | [**IncidentUpdate**](IncidentUpdate.md) |  | 

### Return type

[**Incident**](Incident.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

