# \WorkflowApi

All URIs are relative to *https://circleci.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApprovePendingApprovalJobById**](WorkflowApi.md#ApprovePendingApprovalJobById) | **Post** /workflow/{id}/approve/{approval_request_id} | Approve a job
[**CancelWorkflow**](WorkflowApi.md#CancelWorkflow) | **Post** /workflow/{id}/cancel | Cancel a workflow
[**GetWorkflowById**](WorkflowApi.md#GetWorkflowById) | **Get** /workflow/{id} | Get a workflow
[**ListWorkflowJobs**](WorkflowApi.md#ListWorkflowJobs) | **Get** /workflow/{id}/job | Get a workflow&#39;s jobs
[**RerunWorkflow**](WorkflowApi.md#RerunWorkflow) | **Post** /workflow/{id}/rerun | Rerun a workflow



## ApprovePendingApprovalJobById

> MessageResponse ApprovePendingApprovalJobById(ctx, approvalRequestId, id).Execute()

Approve a job



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
    approvalRequestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the job being approved.
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the workflow.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowApi.ApprovePendingApprovalJobById(context.Background(), approvalRequestId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowApi.ApprovePendingApprovalJobById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApprovePendingApprovalJobById`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowApi.ApprovePendingApprovalJobById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**approvalRequestId** | **string** | The ID of the job being approved. | 
**id** | **string** | The unique ID of the workflow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApprovePendingApprovalJobByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelWorkflow

> MessageResponse CancelWorkflow(ctx, id).Execute()

Cancel a workflow



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the workflow.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowApi.CancelWorkflow(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowApi.CancelWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelWorkflow`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowApi.CancelWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique ID of the workflow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowById

> Workflow GetWorkflowById(ctx, id).Execute()

Get a workflow



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the workflow.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowApi.GetWorkflowById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowApi.GetWorkflowById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowById`: Workflow
    fmt.Fprintf(os.Stdout, "Response from `WorkflowApi.GetWorkflowById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique ID of the workflow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Workflow**](Workflow.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkflowJobs

> WorkflowJobListResponse ListWorkflowJobs(ctx, id).Execute()

Get a workflow's jobs



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the workflow.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowApi.ListWorkflowJobs(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowApi.ListWorkflowJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkflowJobs`: WorkflowJobListResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowApi.ListWorkflowJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique ID of the workflow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowJobListResponse**](WorkflowJobListResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RerunWorkflow

> InlineResponse202 RerunWorkflow(ctx, id).RerunWorkflowParameters(rerunWorkflowParameters).Execute()

Rerun a workflow



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the workflow.
    rerunWorkflowParameters := *openapiclient.NewRerunWorkflowParameters() // RerunWorkflowParameters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowApi.RerunWorkflow(context.Background(), id).RerunWorkflowParameters(rerunWorkflowParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowApi.RerunWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RerunWorkflow`: InlineResponse202
    fmt.Fprintf(os.Stdout, "Response from `WorkflowApi.RerunWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique ID of the workflow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRerunWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rerunWorkflowParameters** | [**RerunWorkflowParameters**](RerunWorkflowParameters.md) |  | 

### Return type

[**InlineResponse202**](InlineResponse202.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

