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
)

// Linger please
var (
	_ context.Context
)

// LicensingAccountsApiService LicensingAccountsApi service
type LicensingAccountsApiService service

type ApiGetUserAccountsListRequest struct {
	ctx context.Context
	ApiService *LicensingAccountsApiService
	userId *string
	domain *string
	roleName *string
	type_ *SmartAccountType
}

// User Id
func (r ApiGetUserAccountsListRequest) UserId(userId string) ApiGetUserAccountsListRequest {
	r.userId = &userId
	return r
}
// Smart Account domain
func (r ApiGetUserAccountsListRequest) Domain(domain string) ApiGetUserAccountsListRequest {
	r.domain = &domain
	return r
}
// Role Name
func (r ApiGetUserAccountsListRequest) RoleName(roleName string) ApiGetUserAccountsListRequest {
	r.roleName = &roleName
	return r
}
func (r ApiGetUserAccountsListRequest) Type_(type_ SmartAccountType) ApiGetUserAccountsListRequest {
	r.type_ = &type_
	return r
}

func (r ApiGetUserAccountsListRequest) Execute() (*SmartUserAccounts, *http.Response, error) {
	return r.ApiService.GetUserAccountsListExecute(r)
}

/*
GetUserAccountsList Returns a filtered page of smart accounts.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetUserAccountsListRequest
*/
func (a *LicensingAccountsApiService) GetUserAccountsList(ctx context.Context) ApiGetUserAccountsListRequest {
	return ApiGetUserAccountsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SmartUserAccounts
func (a *LicensingAccountsApiService) GetUserAccountsListExecute(r ApiGetUserAccountsListRequest) (*SmartUserAccounts, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SmartUserAccounts
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LicensingAccountsApiService.GetUserAccountsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/licensing/api/v8/accounts/smart/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.userId == nil {
		return localVarReturnValue, nil, reportError("userId is required and must be specified")
	}

	localVarQueryParams.Add("userId", parameterToString(*r.userId, ""))
	if r.domain != nil {
		localVarQueryParams.Add("domain", parameterToString(*r.domain, ""))
	}
	if r.roleName != nil {
		localVarQueryParams.Add("roleName", parameterToString(*r.roleName, ""))
	}
	if r.type_ != nil {
		localVarQueryParams.Add("type", parameterToString(*r.type_, ""))
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
