/*
 * peridot/proto/v1/batch.proto
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: version not set
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package peridotopenapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

type PackageServiceApi interface {

	/*
	 * GetPackage GetPackage returns a package by its id or name
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param projectId
	 * @param field
	 * @param value
	 * @return ApiGetPackageRequest
	 */
	GetPackage(ctx _context.Context, projectId string, field string, value string) ApiGetPackageRequest

	/*
	 * GetPackageExecute executes the request
	 * @return V1GetPackageResponse
	 */
	GetPackageExecute(r ApiGetPackageRequest) (V1GetPackageResponse, *_nethttp.Response, error)

	/*
	 * ListPackages ListPackages returns all packages with filters applied
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param projectId Project ID that should be queried
	 * @return ApiListPackagesRequest
	 */
	ListPackages(ctx _context.Context, projectId string) ApiListPackagesRequest

	/*
	 * ListPackagesExecute executes the request
	 * @return V1ListPackagesResponse
	 */
	ListPackagesExecute(r ApiListPackagesRequest) (V1ListPackagesResponse, *_nethttp.Response, error)
}

// PackageServiceApiService PackageServiceApi service
type PackageServiceApiService service

type ApiGetPackageRequest struct {
	ctx _context.Context
	ApiService PackageServiceApi
	projectId string
	field string
	value string
}


func (r ApiGetPackageRequest) Execute() (V1GetPackageResponse, *_nethttp.Response, error) {
	return r.ApiService.GetPackageExecute(r)
}

/*
 * GetPackage GetPackage returns a package by its id or name
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param projectId
 * @param field
 * @param value
 * @return ApiGetPackageRequest
 */
func (a *PackageServiceApiService) GetPackage(ctx _context.Context, projectId string, field string, value string) ApiGetPackageRequest {
	return ApiGetPackageRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
		field: field,
		value: value,
	}
}

/*
 * Execute executes the request
 * @return V1GetPackageResponse
 */
func (a *PackageServiceApiService) GetPackageExecute(r ApiGetPackageRequest) (V1GetPackageResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  V1GetPackageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PackageServiceApiService.GetPackage")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/packages/{field}/{value}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", _neturl.PathEscape(parameterToString(r.projectId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"field"+"}", _neturl.PathEscape(parameterToString(r.field, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"value"+"}", _neturl.PathEscape(parameterToString(r.value, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v RpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListPackagesRequest struct {
	ctx _context.Context
	ApiService PackageServiceApi
	projectId string
	page *int32
	limit *int32
	filtersId *string
	filtersName *string
	filtersModular *bool
	filtersNameExact *string
	filtersNoImports *bool
	filtersNoBuilds *bool
}

func (r ApiListPackagesRequest) Page(page int32) ApiListPackagesRequest {
	r.page = &page
	return r
}
func (r ApiListPackagesRequest) Limit(limit int32) ApiListPackagesRequest {
	r.limit = &limit
	return r
}
func (r ApiListPackagesRequest) FiltersId(filtersId string) ApiListPackagesRequest {
	r.filtersId = &filtersId
	return r
}
func (r ApiListPackagesRequest) FiltersName(filtersName string) ApiListPackagesRequest {
	r.filtersName = &filtersName
	return r
}
func (r ApiListPackagesRequest) FiltersModular(filtersModular bool) ApiListPackagesRequest {
	r.filtersModular = &filtersModular
	return r
}
func (r ApiListPackagesRequest) FiltersNameExact(filtersNameExact string) ApiListPackagesRequest {
	r.filtersNameExact = &filtersNameExact
	return r
}
func (r ApiListPackagesRequest) FiltersNoImports(filtersNoImports bool) ApiListPackagesRequest {
	r.filtersNoImports = &filtersNoImports
	return r
}
func (r ApiListPackagesRequest) FiltersNoBuilds(filtersNoBuilds bool) ApiListPackagesRequest {
	r.filtersNoBuilds = &filtersNoBuilds
	return r
}

func (r ApiListPackagesRequest) Execute() (V1ListPackagesResponse, *_nethttp.Response, error) {
	return r.ApiService.ListPackagesExecute(r)
}

/*
 * ListPackages ListPackages returns all packages with filters applied
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param projectId Project ID that should be queried
 * @return ApiListPackagesRequest
 */
func (a *PackageServiceApiService) ListPackages(ctx _context.Context, projectId string) ApiListPackagesRequest {
	return ApiListPackagesRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
	}
}

/*
 * Execute executes the request
 * @return V1ListPackagesResponse
 */
func (a *PackageServiceApiService) ListPackagesExecute(r ApiListPackagesRequest) (V1ListPackagesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  V1ListPackagesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PackageServiceApiService.ListPackages")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/packages"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", _neturl.PathEscape(parameterToString(r.projectId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.filtersId != nil {
		localVarQueryParams.Add("filters.id", parameterToString(*r.filtersId, ""))
	}
	if r.filtersName != nil {
		localVarQueryParams.Add("filters.name", parameterToString(*r.filtersName, ""))
	}
	if r.filtersModular != nil {
		localVarQueryParams.Add("filters.modular", parameterToString(*r.filtersModular, ""))
	}
	if r.filtersNameExact != nil {
		localVarQueryParams.Add("filters.nameExact", parameterToString(*r.filtersNameExact, ""))
	}
	if r.filtersNoImports != nil {
		localVarQueryParams.Add("filters.noImports", parameterToString(*r.filtersNoImports, ""))
	}
	if r.filtersNoBuilds != nil {
		localVarQueryParams.Add("filters.noBuilds", parameterToString(*r.filtersNoBuilds, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v RpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
