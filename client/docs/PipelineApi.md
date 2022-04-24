# \PipelineApi

All URIs are relative to *https://circleci.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContinuePipeline**](PipelineApi.md#ContinuePipeline) | **Post** /pipeline/continue | Continue a pipeline
[**GetPipelineById**](PipelineApi.md#GetPipelineById) | **Get** /pipeline/{pipeline-id} | Get a pipeline
[**GetPipelineByNumber**](PipelineApi.md#GetPipelineByNumber) | **Get** /project/{project-slug}/pipeline/{pipeline-number} | Get a pipeline
[**GetPipelineConfigById**](PipelineApi.md#GetPipelineConfigById) | **Get** /pipeline/{pipeline-id}/config | Get a pipeline&#39;s configuration
[**ListMyPipelines**](PipelineApi.md#ListMyPipelines) | **Get** /project/{project-slug}/pipeline/mine | Get your pipelines
[**ListPipelines**](PipelineApi.md#ListPipelines) | **Get** /pipeline | Get a list of pipelines
[**ListPipelinesForProject**](PipelineApi.md#ListPipelinesForProject) | **Get** /project/{project-slug}/pipeline | Get all pipelines
[**ListWorkflowsByPipelineId**](PipelineApi.md#ListWorkflowsByPipelineId) | **Get** /pipeline/{pipeline-id}/workflow | Get a pipeline&#39;s workflows
[**TriggerPipeline**](PipelineApi.md#TriggerPipeline) | **Post** /project/{project-slug}/pipeline | Trigger a new pipeline



## ContinuePipeline

> MessageResponse ContinuePipeline(ctx).InlineObject2(inlineObject2).Execute()

Continue a pipeline



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
    inlineObject2 := *openapiclient.NewInlineObject2("ContinuationKey_example", "Configuration_example") // InlineObject2 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelineApi.ContinuePipeline(context.Background()).InlineObject2(inlineObject2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineApi.ContinuePipeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContinuePipeline`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `PipelineApi.ContinuePipeline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContinuePipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject2** | [**InlineObject2**](InlineObject2.md) |  | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelineById

> Pipeline GetPipelineById(ctx, pipelineId).Execute()

Get a pipeline



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
    pipelineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the pipeline.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelineApi.GetPipelineById(context.Background(), pipelineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineApi.GetPipelineById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPipelineById`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `PipelineApi.GetPipelineById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **string** | The unique ID of the pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelineByNumber

> Pipeline1 GetPipelineByNumber(ctx, projectSlug, pipelineNumber).Execute()

Get a pipeline



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
    projectSlug := "gh/CircleCI-Public/api-preview-docs" // string | Project slug in the form `vcs-slug/org-name/repo-name`. The `/` characters may be URL-escaped.
    pipelineNumber := TODO // interface{} | The number of the pipeline.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelineApi.GetPipelineByNumber(context.Background(), projectSlug, pipelineNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineApi.GetPipelineByNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPipelineByNumber`: Pipeline1
    fmt.Fprintf(os.Stdout, "Response from `PipelineApi.GetPipelineByNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectSlug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 
**pipelineNumber** | [**interface{}**](.md) | The number of the pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineByNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Pipeline1**](Pipeline1.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelineConfigById

> PipelineConfig GetPipelineConfigById(ctx, pipelineId).Execute()

Get a pipeline's configuration



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
    pipelineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the pipeline.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelineApi.GetPipelineConfigById(context.Background(), pipelineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineApi.GetPipelineConfigById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPipelineConfigById`: PipelineConfig
    fmt.Fprintf(os.Stdout, "Response from `PipelineApi.GetPipelineConfigById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **string** | The unique ID of the pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineConfigByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PipelineConfig**](PipelineConfig.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMyPipelines

> PipelineListResponse ListMyPipelines(ctx, projectSlug).PageToken(pageToken).Execute()

Get your pipelines



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
    projectSlug := "gh/CircleCI-Public/api-preview-docs" // string | Project slug in the form `vcs-slug/org-name/repo-name`. The `/` characters may be URL-escaped.
    pageToken := "pageToken_example" // string | A token to retrieve the next page of results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelineApi.ListMyPipelines(context.Background(), projectSlug).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineApi.ListMyPipelines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMyPipelines`: PipelineListResponse
    fmt.Fprintf(os.Stdout, "Response from `PipelineApi.ListMyPipelines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectSlug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMyPipelinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | A token to retrieve the next page of results. | 

### Return type

[**PipelineListResponse**](PipelineListResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPipelines

> PipelineListResponse ListPipelines(ctx).OrgSlug(orgSlug).PageToken(pageToken).Mine(mine).Execute()

Get a list of pipelines



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
    orgSlug := "gh/CircleCI-Public" // string | Org slug in the form `vcs-slug/org-name` (optional)
    pageToken := "pageToken_example" // string | A token to retrieve the next page of results. (optional)
    mine := true // bool | Only include entries created by your user. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelineApi.ListPipelines(context.Background()).OrgSlug(orgSlug).PageToken(pageToken).Mine(mine).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineApi.ListPipelines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPipelines`: PipelineListResponse
    fmt.Fprintf(os.Stdout, "Response from `PipelineApi.ListPipelines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPipelinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgSlug** | **string** | Org slug in the form &#x60;vcs-slug/org-name&#x60; | 
 **pageToken** | **string** | A token to retrieve the next page of results. | 
 **mine** | **bool** | Only include entries created by your user. | 

### Return type

[**PipelineListResponse**](PipelineListResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPipelinesForProject

> PipelineListResponse1 ListPipelinesForProject(ctx, projectSlug).Branch(branch).PageToken(pageToken).Execute()

Get all pipelines



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
    projectSlug := "gh/CircleCI-Public/api-preview-docs" // string | Project slug in the form `vcs-slug/org-name/repo-name`. The `/` characters may be URL-escaped.
    branch := "branch_example" // string | The name of a vcs branch. (optional)
    pageToken := "pageToken_example" // string | A token to retrieve the next page of results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelineApi.ListPipelinesForProject(context.Background(), projectSlug).Branch(branch).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineApi.ListPipelinesForProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPipelinesForProject`: PipelineListResponse1
    fmt.Fprintf(os.Stdout, "Response from `PipelineApi.ListPipelinesForProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectSlug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPipelinesForProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **branch** | **string** | The name of a vcs branch. | 
 **pageToken** | **string** | A token to retrieve the next page of results. | 

### Return type

[**PipelineListResponse1**](PipelineListResponse1.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkflowsByPipelineId

> WorkflowListResponse ListWorkflowsByPipelineId(ctx, pipelineId).PageToken(pageToken).Execute()

Get a pipeline's workflows



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
    pipelineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the pipeline.
    pageToken := "pageToken_example" // string | A token to retrieve the next page of results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelineApi.ListWorkflowsByPipelineId(context.Background(), pipelineId).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineApi.ListWorkflowsByPipelineId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkflowsByPipelineId`: WorkflowListResponse
    fmt.Fprintf(os.Stdout, "Response from `PipelineApi.ListWorkflowsByPipelineId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **string** | The unique ID of the pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowsByPipelineIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | A token to retrieve the next page of results. | 

### Return type

[**WorkflowListResponse**](WorkflowListResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerPipeline

> PipelineCreation TriggerPipeline(ctx, projectSlug).TriggerPipelineParameters(triggerPipelineParameters).Execute()

Trigger a new pipeline



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
    projectSlug := "gh/CircleCI-Public/api-preview-docs" // string | Project slug in the form `vcs-slug/org-name/repo-name`. The `/` characters may be URL-escaped.
    triggerPipelineParameters := *openapiclient.NewTriggerPipelineParameters() // TriggerPipelineParameters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PipelineApi.TriggerPipeline(context.Background(), projectSlug).TriggerPipelineParameters(triggerPipelineParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelineApi.TriggerPipeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TriggerPipeline`: PipelineCreation
    fmt.Fprintf(os.Stdout, "Response from `PipelineApi.TriggerPipeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectSlug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTriggerPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **triggerPipelineParameters** | [**TriggerPipelineParameters**](TriggerPipelineParameters.md) |  | 

### Return type

[**PipelineCreation**](PipelineCreation.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

