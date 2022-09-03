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

type ImportServiceApi interface {

	/*
	 * GetImport GetImport gets an import by ID.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param projectId
	 * @param importId
	 * @return ApiGetImportRequest
	 */
	GetImport(ctx _context.Context, projectId string, importId string) ApiGetImportRequest

	/*
	 * GetImportExecute executes the request
	 * @return V1GetImportResponse
	 */
	GetImportExecute(r ApiGetImportRequest) (V1GetImportResponse, *_nethttp.Response, error)

	/*
	 * GetImportBatch GetImportBatch gets an import batch by ID.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param projectId
	 * @param importBatchId
	 * @return ApiGetImportBatchRequest
	 */
	GetImportBatch(ctx _context.Context, projectId string, importBatchId string) ApiGetImportBatchRequest

	/*
	 * GetImportBatchExecute executes the request
	 * @return V1GetImportBatchResponse
	 */
	GetImportBatchExecute(r ApiGetImportBatchRequest) (V1GetImportBatchResponse, *_nethttp.Response, error)

	/*
	 * ImportBatchRetryFailed ImportBatchRetryFailed retries failed imports in a batch.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param projectId
	 * @param importBatchId
	 * @return ApiImportBatchRetryFailedRequest
	 */
	ImportBatchRetryFailed(ctx _context.Context, projectId string, importBatchId string) ApiImportBatchRetryFailedRequest

	/*
	 * ImportBatchRetryFailedExecute executes the request
	 * @return V1ImportBatchRetryFailedResponse
	 */
	ImportBatchRetryFailedExecute(r ApiImportBatchRetryFailedRequest) (V1ImportBatchRetryFailedResponse, *_nethttp.Response, error)

	/*
	 * ImportPackage ImportPackage imports a package scoped to a project This method is asynchronous. Peridot uses the AsyncTask abstraction. Check out `//peridot/proto/v1:task.proto` for more information TODO low-pri: Support inter-project imports
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param projectId Project ID that we want this import to be assigned to All import requests need a project id, however after the initial import, sharing the VRE in an inter-project way is possible.
	 * @return ApiImportPackageRequest
	 */
	ImportPackage(ctx _context.Context, projectId string) ApiImportPackageRequest

	/*
	 * ImportPackageExecute executes the request
	 * @return V1AsyncTask
	 */
	ImportPackageExecute(r ApiImportPackageRequest) (V1AsyncTask, *_nethttp.Response, error)

	/*
	 * ImportPackageBatch ImportPackageBatch imports a batch of packages scoped to a project
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param projectId Only the top-most project id is used for all import requests
	 * @return ApiImportPackageBatchRequest
	 */
	ImportPackageBatch(ctx _context.Context, projectId string) ApiImportPackageBatchRequest

	/*
	 * ImportPackageBatchExecute executes the request
	 * @return V1ImportPackageBatchResponse
	 */
	ImportPackageBatchExecute(r ApiImportPackageBatchRequest) (V1ImportPackageBatchResponse, *_nethttp.Response, error)

	/*
	 * ListImportBatches ListImportBatches lists all import batches for a project.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param projectId
	 * @return ApiListImportBatchesRequest
	 */
	ListImportBatches(ctx _context.Context, projectId string) ApiListImportBatchesRequest

	/*
	 * ListImportBatchesExecute executes the request
	 * @return V1ListImportBatchesResponse
	 */
	ListImportBatchesExecute(r ApiListImportBatchesRequest) (V1ListImportBatchesResponse, *_nethttp.Response, error)

	/*
	 * ListImports ListImports lists all imports for a project.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param projectId
	 * @return ApiListImportsRequest
	 */
	ListImports(ctx _context.Context, projectId string) ApiListImportsRequest

	/*
	 * ListImportsExecute executes the request
	 * @return V1ListImportsResponse
	 */
	ListImportsExecute(r ApiListImportsRequest) (V1ListImportsResponse, *_nethttp.Response, error)
}

// ImportServiceApiService ImportServiceApi service
type ImportServiceApiService service

type ApiGetImportRequest struct {
	ctx _context.Context
	ApiService ImportServiceApi
	projectId string
	importId string
}


func (r ApiGetImportRequest) Execute() (V1GetImportResponse, *_nethttp.Response, error) {
	return r.ApiService.GetImportExecute(r)
}

/*
 * GetImport GetImport gets an import by ID.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param projectId
 * @param importId
 * @return ApiGetImportRequest
 */
func (a *ImportServiceApiService) GetImport(ctx _context.Context, projectId string, importId string) ApiGetImportRequest {
	return ApiGetImportRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
		importId: importId,
	}
}

/*
 * Execute executes the request
 * @return V1GetImportResponse
 */
func (a *ImportServiceApiService) GetImportExecute(r ApiGetImportRequest) (V1GetImportResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  V1GetImportResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportServiceApiService.GetImport")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/imports/{importId}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", _neturl.PathEscape(parameterToString(r.projectId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"importId"+"}", _neturl.PathEscape(parameterToString(r.importId, "")), -1)

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

type ApiGetImportBatchRequest struct {
	ctx _context.Context
	ApiService ImportServiceApi
	projectId string
	importBatchId string
	page *int32
	limit *int32
	filterStatus *string
}

func (r ApiGetImportBatchRequest) Page(page int32) ApiGetImportBatchRequest {
	r.page = &page
	return r
}
func (r ApiGetImportBatchRequest) Limit(limit int32) ApiGetImportBatchRequest {
	r.limit = &limit
	return r
}
func (r ApiGetImportBatchRequest) FilterStatus(filterStatus string) ApiGetImportBatchRequest {
	r.filterStatus = &filterStatus
	return r
}

func (r ApiGetImportBatchRequest) Execute() (V1GetImportBatchResponse, *_nethttp.Response, error) {
	return r.ApiService.GetImportBatchExecute(r)
}

/*
 * GetImportBatch GetImportBatch gets an import batch by ID.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param projectId
 * @param importBatchId
 * @return ApiGetImportBatchRequest
 */
func (a *ImportServiceApiService) GetImportBatch(ctx _context.Context, projectId string, importBatchId string) ApiGetImportBatchRequest {
	return ApiGetImportBatchRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
		importBatchId: importBatchId,
	}
}

/*
 * Execute executes the request
 * @return V1GetImportBatchResponse
 */
func (a *ImportServiceApiService) GetImportBatchExecute(r ApiGetImportBatchRequest) (V1GetImportBatchResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  V1GetImportBatchResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportServiceApiService.GetImportBatch")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/import_batches/{importBatchId}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", _neturl.PathEscape(parameterToString(r.projectId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"importBatchId"+"}", _neturl.PathEscape(parameterToString(r.importBatchId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.filterStatus != nil {
		localVarQueryParams.Add("filter.status", parameterToString(*r.filterStatus, ""))
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

type ApiImportBatchRetryFailedRequest struct {
	ctx _context.Context
	ApiService ImportServiceApi
	projectId string
	importBatchId string
}


func (r ApiImportBatchRetryFailedRequest) Execute() (V1ImportBatchRetryFailedResponse, *_nethttp.Response, error) {
	return r.ApiService.ImportBatchRetryFailedExecute(r)
}

/*
 * ImportBatchRetryFailed ImportBatchRetryFailed retries failed imports in a batch.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param projectId
 * @param importBatchId
 * @return ApiImportBatchRetryFailedRequest
 */
func (a *ImportServiceApiService) ImportBatchRetryFailed(ctx _context.Context, projectId string, importBatchId string) ApiImportBatchRetryFailedRequest {
	return ApiImportBatchRetryFailedRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
		importBatchId: importBatchId,
	}
}

/*
 * Execute executes the request
 * @return V1ImportBatchRetryFailedResponse
 */
func (a *ImportServiceApiService) ImportBatchRetryFailedExecute(r ApiImportBatchRetryFailedRequest) (V1ImportBatchRetryFailedResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  V1ImportBatchRetryFailedResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportServiceApiService.ImportBatchRetryFailed")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/import_batches/{importBatchId}/retry_failed"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", _neturl.PathEscape(parameterToString(r.projectId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"importBatchId"+"}", _neturl.PathEscape(parameterToString(r.importBatchId, "")), -1)

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

type ApiImportPackageRequest struct {
	ctx _context.Context
	ApiService ImportServiceApi
	projectId string
	body *ImportPackageRequestIsTheRequestMessageForImportServiceImportPackage
}

func (r ApiImportPackageRequest) Body(body ImportPackageRequestIsTheRequestMessageForImportServiceImportPackage) ApiImportPackageRequest {
	r.body = &body
	return r
}

func (r ApiImportPackageRequest) Execute() (V1AsyncTask, *_nethttp.Response, error) {
	return r.ApiService.ImportPackageExecute(r)
}

/*
 * ImportPackage ImportPackage imports a package scoped to a project This method is asynchronous. Peridot uses the AsyncTask abstraction. Check out `//peridot/proto/v1:task.proto` for more information TODO low-pri: Support inter-project imports
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param projectId Project ID that we want this import to be assigned to All import requests need a project id, however after the initial import, sharing the VRE in an inter-project way is possible.
 * @return ApiImportPackageRequest
 */
func (a *ImportServiceApiService) ImportPackage(ctx _context.Context, projectId string) ApiImportPackageRequest {
	return ApiImportPackageRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
	}
}

/*
 * Execute executes the request
 * @return V1AsyncTask
 */
func (a *ImportServiceApiService) ImportPackageExecute(r ApiImportPackageRequest) (V1AsyncTask, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  V1AsyncTask
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportServiceApiService.ImportPackage")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/imports"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", _neturl.PathEscape(parameterToString(r.projectId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	localVarPostBody = r.body
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

type ApiImportPackageBatchRequest struct {
	ctx _context.Context
	ApiService ImportServiceApi
	projectId string
	body *InlineObject7
}

func (r ApiImportPackageBatchRequest) Body(body InlineObject7) ApiImportPackageBatchRequest {
	r.body = &body
	return r
}

func (r ApiImportPackageBatchRequest) Execute() (V1ImportPackageBatchResponse, *_nethttp.Response, error) {
	return r.ApiService.ImportPackageBatchExecute(r)
}

/*
 * ImportPackageBatch ImportPackageBatch imports a batch of packages scoped to a project
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param projectId Only the top-most project id is used for all import requests
 * @return ApiImportPackageBatchRequest
 */
func (a *ImportServiceApiService) ImportPackageBatch(ctx _context.Context, projectId string) ApiImportPackageBatchRequest {
	return ApiImportPackageBatchRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
	}
}

/*
 * Execute executes the request
 * @return V1ImportPackageBatchResponse
 */
func (a *ImportServiceApiService) ImportPackageBatchExecute(r ApiImportPackageBatchRequest) (V1ImportPackageBatchResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  V1ImportPackageBatchResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportServiceApiService.ImportPackageBatch")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/import_batches"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", _neturl.PathEscape(parameterToString(r.projectId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	localVarPostBody = r.body
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

type ApiListImportBatchesRequest struct {
	ctx _context.Context
	ApiService ImportServiceApi
	projectId string
	page *int32
	limit *int32
}

func (r ApiListImportBatchesRequest) Page(page int32) ApiListImportBatchesRequest {
	r.page = &page
	return r
}
func (r ApiListImportBatchesRequest) Limit(limit int32) ApiListImportBatchesRequest {
	r.limit = &limit
	return r
}

func (r ApiListImportBatchesRequest) Execute() (V1ListImportBatchesResponse, *_nethttp.Response, error) {
	return r.ApiService.ListImportBatchesExecute(r)
}

/*
 * ListImportBatches ListImportBatches lists all import batches for a project.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param projectId
 * @return ApiListImportBatchesRequest
 */
func (a *ImportServiceApiService) ListImportBatches(ctx _context.Context, projectId string) ApiListImportBatchesRequest {
	return ApiListImportBatchesRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
	}
}

/*
 * Execute executes the request
 * @return V1ListImportBatchesResponse
 */
func (a *ImportServiceApiService) ListImportBatchesExecute(r ApiListImportBatchesRequest) (V1ListImportBatchesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  V1ListImportBatchesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportServiceApiService.ListImportBatches")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/import_batches"
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

type ApiListImportsRequest struct {
	ctx _context.Context
	ApiService ImportServiceApi
	projectId string
	page *int32
	limit *int32
}

func (r ApiListImportsRequest) Page(page int32) ApiListImportsRequest {
	r.page = &page
	return r
}
func (r ApiListImportsRequest) Limit(limit int32) ApiListImportsRequest {
	r.limit = &limit
	return r
}

func (r ApiListImportsRequest) Execute() (V1ListImportsResponse, *_nethttp.Response, error) {
	return r.ApiService.ListImportsExecute(r)
}

/*
 * ListImports ListImports lists all imports for a project.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param projectId
 * @return ApiListImportsRequest
 */
func (a *ImportServiceApiService) ListImports(ctx _context.Context, projectId string) ApiListImportsRequest {
	return ApiListImportsRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
	}
}

/*
 * Execute executes the request
 * @return V1ListImportsResponse
 */
func (a *ImportServiceApiService) ListImportsExecute(r ApiListImportsRequest) (V1ListImportsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  V1ListImportsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportServiceApiService.ListImports")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/imports"
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
