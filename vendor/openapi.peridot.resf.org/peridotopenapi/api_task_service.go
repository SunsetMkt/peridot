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

type TaskServiceApi interface {

	/*
	 * CancelTask CancelTask cancels a task with the given ID. Only parent tasks can be cancelled and if they're in the PENDING or RUNNING state.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param projectId
	 * @param id
	 * @return ApiCancelTaskRequest
	 */
	CancelTask(ctx _context.Context, projectId string, id string) ApiCancelTaskRequest

	/*
	 * CancelTaskExecute executes the request
	 * @return map[string]interface{}
	 */
	CancelTaskExecute(r ApiCancelTaskRequest) (map[string]interface{}, *_nethttp.Response, error)

	/*
	 * GetTask GetTask returns a specific task with the given ID
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param projectId
	 * @param id
	 * @return ApiGetTaskRequest
	 */
	GetTask(ctx _context.Context, projectId string, id string) ApiGetTaskRequest

	/*
	 * GetTaskExecute executes the request
	 * @return V1GetTaskResponse
	 */
	GetTaskExecute(r ApiGetTaskRequest) (V1GetTaskResponse, *_nethttp.Response, error)

	/*
	 * ListTasks ListTasks returns a list of tasks from all projects List mode won't return task responses. The reason being responses being able to reach huge sizes. To get the response for a specific task, you can use GetTask, either on the specific subtask or the parent task.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param projectId
	 * @return ApiListTasksRequest
	 */
	ListTasks(ctx _context.Context, projectId string) ApiListTasksRequest

	/*
	 * ListTasksExecute executes the request
	 * @return V1ListTasksResponse
	 */
	ListTasksExecute(r ApiListTasksRequest) (V1ListTasksResponse, *_nethttp.Response, error)

	/*
	 * StreamTaskLogs StreamTaskLogs streams the logs of a specific task with the given ID
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param projectId
	 * @param id
	 * @return ApiStreamTaskLogsRequest
	 */
	StreamTaskLogs(ctx _context.Context, projectId string, id string) ApiStreamTaskLogsRequest

	/*
	 * StreamTaskLogsExecute executes the request
	 * @return StreamResultOfApiHttpBody
	 */
	StreamTaskLogsExecute(r ApiStreamTaskLogsRequest) (StreamResultOfApiHttpBody, *_nethttp.Response, error)
}

// TaskServiceApiService TaskServiceApi service
type TaskServiceApiService service

type ApiCancelTaskRequest struct {
	ctx _context.Context
	ApiService TaskServiceApi
	projectId string
	id string
}


func (r ApiCancelTaskRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.CancelTaskExecute(r)
}

/*
 * CancelTask CancelTask cancels a task with the given ID. Only parent tasks can be cancelled and if they're in the PENDING or RUNNING state.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param projectId
 * @param id
 * @return ApiCancelTaskRequest
 */
func (a *TaskServiceApiService) CancelTask(ctx _context.Context, projectId string, id string) ApiCancelTaskRequest {
	return ApiCancelTaskRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *TaskServiceApiService) CancelTaskExecute(r ApiCancelTaskRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskServiceApiService.CancelTask")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/tasks/{id}/cancel"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", _neturl.PathEscape(parameterToString(r.projectId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

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

type ApiGetTaskRequest struct {
	ctx _context.Context
	ApiService TaskServiceApi
	projectId string
	id string
}


func (r ApiGetTaskRequest) Execute() (V1GetTaskResponse, *_nethttp.Response, error) {
	return r.ApiService.GetTaskExecute(r)
}

/*
 * GetTask GetTask returns a specific task with the given ID
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param projectId
 * @param id
 * @return ApiGetTaskRequest
 */
func (a *TaskServiceApiService) GetTask(ctx _context.Context, projectId string, id string) ApiGetTaskRequest {
	return ApiGetTaskRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return V1GetTaskResponse
 */
func (a *TaskServiceApiService) GetTaskExecute(r ApiGetTaskRequest) (V1GetTaskResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  V1GetTaskResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskServiceApiService.GetTask")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/tasks/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", _neturl.PathEscape(parameterToString(r.projectId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

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

type ApiListTasksRequest struct {
	ctx _context.Context
	ApiService TaskServiceApi
	projectId string
	page *int32
	limit *int32
}

func (r ApiListTasksRequest) Page(page int32) ApiListTasksRequest {
	r.page = &page
	return r
}
func (r ApiListTasksRequest) Limit(limit int32) ApiListTasksRequest {
	r.limit = &limit
	return r
}

func (r ApiListTasksRequest) Execute() (V1ListTasksResponse, *_nethttp.Response, error) {
	return r.ApiService.ListTasksExecute(r)
}

/*
 * ListTasks ListTasks returns a list of tasks from all projects List mode won't return task responses. The reason being responses being able to reach huge sizes. To get the response for a specific task, you can use GetTask, either on the specific subtask or the parent task.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param projectId
 * @return ApiListTasksRequest
 */
func (a *TaskServiceApiService) ListTasks(ctx _context.Context, projectId string) ApiListTasksRequest {
	return ApiListTasksRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
	}
}

/*
 * Execute executes the request
 * @return V1ListTasksResponse
 */
func (a *TaskServiceApiService) ListTasksExecute(r ApiListTasksRequest) (V1ListTasksResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  V1ListTasksResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskServiceApiService.ListTasks")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/tasks"
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

type ApiStreamTaskLogsRequest struct {
	ctx _context.Context
	ApiService TaskServiceApi
	projectId string
	id string
	parent *bool
}

func (r ApiStreamTaskLogsRequest) Parent(parent bool) ApiStreamTaskLogsRequest {
	r.parent = &parent
	return r
}

func (r ApiStreamTaskLogsRequest) Execute() (StreamResultOfApiHttpBody, *_nethttp.Response, error) {
	return r.ApiService.StreamTaskLogsExecute(r)
}

/*
 * StreamTaskLogs StreamTaskLogs streams the logs of a specific task with the given ID
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param projectId
 * @param id
 * @return ApiStreamTaskLogsRequest
 */
func (a *TaskServiceApiService) StreamTaskLogs(ctx _context.Context, projectId string, id string) ApiStreamTaskLogsRequest {
	return ApiStreamTaskLogsRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return StreamResultOfApiHttpBody
 */
func (a *TaskServiceApiService) StreamTaskLogsExecute(r ApiStreamTaskLogsRequest) (StreamResultOfApiHttpBody, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  StreamResultOfApiHttpBody
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskServiceApiService.StreamTaskLogs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/tasks/{id}/logs"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", _neturl.PathEscape(parameterToString(r.projectId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.parent != nil {
		localVarQueryParams.Add("parent", parameterToString(*r.parent, ""))
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
