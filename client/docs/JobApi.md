# \JobApi

All URIs are relative to *https://circleci.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelJob**](JobApi.md#CancelJob) | **Post** /project/{project-slug}/job/{job-number}/cancel | Cancel job
[**GetJobArtifacts**](JobApi.md#GetJobArtifacts) | **Get** /project/{project-slug}/{job-number}/artifacts | Get a job&#39;s artifacts
[**GetJobDetails**](JobApi.md#GetJobDetails) | **Get** /project/{project-slug}/job/{job-number} | Get job details
[**GetTests**](JobApi.md#GetTests) | **Get** /project/{project-slug}/{job-number}/tests | Get test metadata



## CancelJob

> MessageResponse CancelJob(ctx, jobNumber, projectSlug).Execute()

Cancel job



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
    jobNumber := TODO // interface{} | The number of the job.
    projectSlug := "gh/CircleCI-Public/api-preview-docs" // string | Project slug in the form `vcs-slug/org-name/repo-name`. The `/` characters may be URL-escaped.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobApi.CancelJob(context.Background(), jobNumber, projectSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobApi.CancelJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelJob`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `JobApi.CancelJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobNumber** | [**interface{}**](.md) | The number of the job. | 
**projectSlug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelJobRequest struct via the builder pattern


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


## GetJobArtifacts

> ArtifactListResponse GetJobArtifacts(ctx, jobNumber, projectSlug).Execute()

Get a job's artifacts



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
    jobNumber := TODO // interface{} | The number of the job.
    projectSlug := "gh/CircleCI-Public/api-preview-docs" // string | Project slug in the form `vcs-slug/org-name/repo-name`. The `/` characters may be URL-escaped.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobApi.GetJobArtifacts(context.Background(), jobNumber, projectSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobApi.GetJobArtifacts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobArtifacts`: ArtifactListResponse
    fmt.Fprintf(os.Stdout, "Response from `JobApi.GetJobArtifacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobNumber** | [**interface{}**](.md) | The number of the job. | 
**projectSlug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ArtifactListResponse**](ArtifactListResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobDetails

> JobDetails GetJobDetails(ctx, jobNumber, projectSlug).Execute()

Get job details



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
    jobNumber := TODO // interface{} | The number of the job.
    projectSlug := "gh/CircleCI-Public/api-preview-docs" // string | Project slug in the form `vcs-slug/org-name/repo-name`. The `/` characters may be URL-escaped.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobApi.GetJobDetails(context.Background(), jobNumber, projectSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobApi.GetJobDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobDetails`: JobDetails
    fmt.Fprintf(os.Stdout, "Response from `JobApi.GetJobDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobNumber** | [**interface{}**](.md) | The number of the job. | 
**projectSlug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**JobDetails**](JobDetails.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTests

> TestsResponse GetTests(ctx, jobNumber, projectSlug).Execute()

Get test metadata



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
    jobNumber := TODO // interface{} | The number of the job.
    projectSlug := "gh/CircleCI-Public/api-preview-docs" // string | Project slug in the form `vcs-slug/org-name/repo-name`. The `/` characters may be URL-escaped.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobApi.GetTests(context.Background(), jobNumber, projectSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobApi.GetTests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTests`: TestsResponse
    fmt.Fprintf(os.Stdout, "Response from `JobApi.GetTests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobNumber** | [**interface{}**](.md) | The number of the job. | 
**projectSlug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TestsResponse**](TestsResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

