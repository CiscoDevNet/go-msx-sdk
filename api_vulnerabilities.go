/*
MSX SDK

MSX SDK client.

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msxsdk

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// Linger please
var (
	_ context.Context
)

// VulnerabilitiesApiService VulnerabilitiesApi service
type VulnerabilitiesApiService service

type ApiGetIngestVulnerabilitiesTasksPageRequest struct {
	ctx context.Context
	ApiService *VulnerabilitiesApiService
	page *int32
	pageSize *int32
	startDate *time.Time
	endDate *time.Time
}

func (r ApiGetIngestVulnerabilitiesTasksPageRequest) Page(page int32) ApiGetIngestVulnerabilitiesTasksPageRequest {
	r.page = &page
	return r
}
func (r ApiGetIngestVulnerabilitiesTasksPageRequest) PageSize(pageSize int32) ApiGetIngestVulnerabilitiesTasksPageRequest {
	r.pageSize = &pageSize
	return r
}
// Start date for date range filter on validation execution date.
func (r ApiGetIngestVulnerabilitiesTasksPageRequest) StartDate(startDate time.Time) ApiGetIngestVulnerabilitiesTasksPageRequest {
	r.startDate = &startDate
	return r
}
// End date for date range filter on validation execution date.
func (r ApiGetIngestVulnerabilitiesTasksPageRequest) EndDate(endDate time.Time) ApiGetIngestVulnerabilitiesTasksPageRequest {
	r.endDate = &endDate
	return r
}

func (r ApiGetIngestVulnerabilitiesTasksPageRequest) Execute() (*VulnerabilityIngestPage, *http.Response, error) {
	return r.ApiService.GetIngestVulnerabilitiesTasksPageExecute(r)
}

/*
GetIngestVulnerabilitiesTasksPage Returns a filtered page of ingest tasks.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetIngestVulnerabilitiesTasksPageRequest
*/
func (a *VulnerabilitiesApiService) GetIngestVulnerabilitiesTasksPage(ctx context.Context) ApiGetIngestVulnerabilitiesTasksPageRequest {
	return ApiGetIngestVulnerabilitiesTasksPageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return VulnerabilityIngestPage
func (a *VulnerabilitiesApiService) GetIngestVulnerabilitiesTasksPageExecute(r ApiGetIngestVulnerabilitiesTasksPageRequest) (*VulnerabilityIngestPage, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VulnerabilityIngestPage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VulnerabilitiesApiService.GetIngestVulnerabilitiesTasksPage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/vulnerability/api/v8/vulnerabilities/ingests"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.page == nil {
		return localVarReturnValue, nil, reportError("page is required and must be specified")
	}
	if *r.page < 0 {
		return localVarReturnValue, nil, reportError("page must be greater than 0")
	}
	if r.pageSize == nil {
		return localVarReturnValue, nil, reportError("pageSize is required and must be specified")
	}
	if *r.pageSize < 1 {
		return localVarReturnValue, nil, reportError("pageSize must be greater than 1")
	}
	if *r.pageSize > 1000 {
		return localVarReturnValue, nil, reportError("pageSize must be less than 1000")
	}

	if r.startDate != nil {
		localVarQueryParams.Add("startDate", parameterToString(*r.startDate, ""))
	}
	if r.endDate != nil {
		localVarQueryParams.Add("endDate", parameterToString(*r.endDate, ""))
	}
	localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetVulnerabilitiesPageRequest struct {
	ctx context.Context
	ApiService *VulnerabilitiesApiService
	page *int32
	pageSize *int32
	cveId *string
	vendor *string
	product *string
	version *string
	severity *VulnerabilitySeverity
	startDate *time.Time
	endDate *time.Time
	year *int32
	sortBy *string
	sortOrder *string
}

func (r ApiGetVulnerabilitiesPageRequest) Page(page int32) ApiGetVulnerabilitiesPageRequest {
	r.page = &page
	return r
}
func (r ApiGetVulnerabilitiesPageRequest) PageSize(pageSize int32) ApiGetVulnerabilitiesPageRequest {
	r.pageSize = &pageSize
	return r
}
// CVE identifer (https://www.cvedetails.com/cve-help.php) to filter by.
func (r ApiGetVulnerabilitiesPageRequest) CveId(cveId string) ApiGetVulnerabilitiesPageRequest {
	r.cveId = &cveId
	return r
}
// Vendor identifier (as defined in NIST&#39;s CPE dictionary) to filter by.
func (r ApiGetVulnerabilitiesPageRequest) Vendor(vendor string) ApiGetVulnerabilitiesPageRequest {
	r.vendor = &vendor
	return r
}
// Product identifier (as defined in NIST&#39;s CPE dictionary) to filter by.
func (r ApiGetVulnerabilitiesPageRequest) Product(product string) ApiGetVulnerabilitiesPageRequest {
	r.product = &product
	return r
}
// Product version (as defined in NIST&#39;s CPE dictionary) filter to filter by.
func (r ApiGetVulnerabilitiesPageRequest) Version(version string) ApiGetVulnerabilitiesPageRequest {
	r.version = &version
	return r
}
// A CVSS severity level (https://nvd.nist.gov/vuln-metrics/cvss) to filter by.
func (r ApiGetVulnerabilitiesPageRequest) Severity(severity VulnerabilitySeverity) ApiGetVulnerabilitiesPageRequest {
	r.severity = &severity
	return r
}
// Start date for date range filter on CVE published date.
func (r ApiGetVulnerabilitiesPageRequest) StartDate(startDate time.Time) ApiGetVulnerabilitiesPageRequest {
	r.startDate = &startDate
	return r
}
// End date for date range filter on CVE published date.
func (r ApiGetVulnerabilitiesPageRequest) EndDate(endDate time.Time) ApiGetVulnerabilitiesPageRequest {
	r.endDate = &endDate
	return r
}
// Year CVE published filter.
func (r ApiGetVulnerabilitiesPageRequest) Year(year int32) ApiGetVulnerabilitiesPageRequest {
	r.year = &year
	return r
}
func (r ApiGetVulnerabilitiesPageRequest) SortBy(sortBy string) ApiGetVulnerabilitiesPageRequest {
	r.sortBy = &sortBy
	return r
}
func (r ApiGetVulnerabilitiesPageRequest) SortOrder(sortOrder string) ApiGetVulnerabilitiesPageRequest {
	r.sortOrder = &sortOrder
	return r
}

func (r ApiGetVulnerabilitiesPageRequest) Execute() (*VulnerabilitiesPage, *http.Response, error) {
	return r.ApiService.GetVulnerabilitiesPageExecute(r)
}

/*
GetVulnerabilitiesPage Returns a filtered page of vulnerabilities.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetVulnerabilitiesPageRequest
*/
func (a *VulnerabilitiesApiService) GetVulnerabilitiesPage(ctx context.Context) ApiGetVulnerabilitiesPageRequest {
	return ApiGetVulnerabilitiesPageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return VulnerabilitiesPage
func (a *VulnerabilitiesApiService) GetVulnerabilitiesPageExecute(r ApiGetVulnerabilitiesPageRequest) (*VulnerabilitiesPage, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VulnerabilitiesPage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VulnerabilitiesApiService.GetVulnerabilitiesPage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/vulnerability/api/v8/vulnerabilities"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.page == nil {
		return localVarReturnValue, nil, reportError("page is required and must be specified")
	}
	if *r.page < 0 {
		return localVarReturnValue, nil, reportError("page must be greater than 0")
	}
	if r.pageSize == nil {
		return localVarReturnValue, nil, reportError("pageSize is required and must be specified")
	}
	if *r.pageSize < 1 {
		return localVarReturnValue, nil, reportError("pageSize must be greater than 1")
	}
	if *r.pageSize > 1000 {
		return localVarReturnValue, nil, reportError("pageSize must be less than 1000")
	}

	if r.cveId != nil {
		localVarQueryParams.Add("cveId", parameterToString(*r.cveId, ""))
	}
	if r.vendor != nil {
		localVarQueryParams.Add("vendor", parameterToString(*r.vendor, ""))
	}
	if r.product != nil {
		localVarQueryParams.Add("product", parameterToString(*r.product, ""))
	}
	if r.version != nil {
		localVarQueryParams.Add("version", parameterToString(*r.version, ""))
	}
	if r.severity != nil {
		localVarQueryParams.Add("severity", parameterToString(*r.severity, ""))
	}
	if r.startDate != nil {
		localVarQueryParams.Add("startDate", parameterToString(*r.startDate, ""))
	}
	if r.endDate != nil {
		localVarQueryParams.Add("endDate", parameterToString(*r.endDate, ""))
	}
	if r.year != nil {
		localVarQueryParams.Add("year", parameterToString(*r.year, ""))
	}
	localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
	if r.sortBy != nil {
		localVarQueryParams.Add("sortBy", parameterToString(*r.sortBy, ""))
	}
	if r.sortOrder != nil {
		localVarQueryParams.Add("sortOrder", parameterToString(*r.sortOrder, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiIngestVulnerabilitiesRequest struct {
	ctx context.Context
	ApiService *VulnerabilitiesApiService
	vulnerabilityFeed *VulnerabilityFeed
}

func (r ApiIngestVulnerabilitiesRequest) VulnerabilityFeed(vulnerabilityFeed VulnerabilityFeed) ApiIngestVulnerabilitiesRequest {
	r.vulnerabilityFeed = &vulnerabilityFeed
	return r
}

func (r ApiIngestVulnerabilitiesRequest) Execute() (*VulnerabilityIngestion, *http.Response, error) {
	return r.ApiService.IngestVulnerabilitiesExecute(r)
}

/*
IngestVulnerabilities Ingests a CVE JSON feed into the Vulnerability Service datastore.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiIngestVulnerabilitiesRequest
*/
func (a *VulnerabilitiesApiService) IngestVulnerabilities(ctx context.Context) ApiIngestVulnerabilitiesRequest {
	return ApiIngestVulnerabilitiesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return VulnerabilityIngestion
func (a *VulnerabilitiesApiService) IngestVulnerabilitiesExecute(r ApiIngestVulnerabilitiesRequest) (*VulnerabilityIngestion, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VulnerabilityIngestion
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VulnerabilitiesApiService.IngestVulnerabilities")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/vulnerability/api/v8/vulnerabilities/ingests"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.vulnerabilityFeed == nil {
		return localVarReturnValue, nil, reportError("vulnerabilityFeed is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.vulnerabilityFeed
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
