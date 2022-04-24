# \InsightsApi

All URIs are relative to *https://circleci.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllInsightsBranches**](InsightsApi.md#GetAllInsightsBranches) | **Get** /insights/{project-slug}/branches | Get all branches for a project
[**GetFlakyTests**](InsightsApi.md#GetFlakyTests) | **Get** /insights/{project-slug}/flaky-tests | Get flaky tests for a project
[**GetJobTimeseries**](InsightsApi.md#GetJobTimeseries) | **Get** /insights/time-series/{project-slug}/workflows/{workflow-name}/jobs | Job timeseries data
[**GetOrgSummaryData**](InsightsApi.md#GetOrgSummaryData) | **Get** /insights/{org-slug}/summary | Get summary metrics with trends for the entire org, and for each project.
[**GetProjectJobRuns**](InsightsApi.md#GetProjectJobRuns) | **Get** /insights/{project-slug}/workflows/{workflow-name}/jobs/{job-name} | Get recent runs of a workflow job
[**GetProjectWorkflowJobMetrics**](InsightsApi.md#GetProjectWorkflowJobMetrics) | **Get** /insights/{project-slug}/workflows/{workflow-name}/jobs | Get summary metrics for a project workflow&#39;s jobs.
[**GetProjectWorkflowMetrics**](InsightsApi.md#GetProjectWorkflowMetrics) | **Get** /insights/{project-slug}/workflows | Get summary metrics for a project&#39;s workflows
[**GetProjectWorkflowRuns**](InsightsApi.md#GetProjectWorkflowRuns) | **Get** /insights/{project-slug}/workflows/{workflow-name} | Get recent runs of a workflow
[**GetProjectWorkflowTestMetrics**](InsightsApi.md#GetProjectWorkflowTestMetrics) | **Get** /insights/{project-slug}/workflows/{workflow-name}/test-metrics | Get test metrics for a project&#39;s workflows
[**GetProjectWorkflowsPageData**](InsightsApi.md#GetProjectWorkflowsPageData) | **Get** /insights/pages/{project-slug}/summary | Get summary metrics and trends for a project across it&#39;s workflows and branches
[**GetWorkflowSummary**](InsightsApi.md#GetWorkflowSummary) | **Get** /insights/{project-slug}/workflows/{workflow-name}/summary | Get metrics and trends for workflows
[**GetWorkflowTimeseries**](InsightsApi.md#GetWorkflowTimeseries) | **Get** /insights/time-series/{project-slug}/workflows | Workflow timeseries data



## GetAllInsightsBranches

> interface{} GetAllInsightsBranches(ctx, projectSlug).WorkflowName(workflowName).Execute()

Get all branches for a project



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
    workflowName := "build-and-test" // string | The name of a workflow. If not passed we will scope the API call to the project. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsApi.GetAllInsightsBranches(context.Background(), projectSlug).WorkflowName(workflowName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsApi.GetAllInsightsBranches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllInsightsBranches`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `InsightsApi.GetAllInsightsBranches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectSlug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllInsightsBranchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowName** | **string** | The name of a workflow. If not passed we will scope the API call to the project. | 

### Return type

**interface{}**

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlakyTests

> InlineResponse2005 GetFlakyTests(ctx, projectSlug).Execute()

Get flaky tests for a project



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsApi.GetFlakyTests(context.Background(), projectSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsApi.GetFlakyTests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFlakyTests`: InlineResponse2005
    fmt.Fprintf(os.Stdout, "Response from `InsightsApi.GetFlakyTests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectSlug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlakyTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2005**](InlineResponse2005.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobTimeseries

> InlineResponse2003 GetJobTimeseries(ctx, projectSlug, workflowName).Branch(branch).TimeseriesGranularity(timeseriesGranularity).StartDate(startDate).EndDate(endDate).Execute()

Job timeseries data



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
    projectSlug := "gh/CircleCI-Public/api-preview-docs" // string | Project slug in the form `vcs-slug/org-name/repo-name`. The `/` characters may be URL-escaped.
    workflowName := "build-and-test" // string | The name of the workflow.
    branch := "branch_example" // string | The name of a vcs branch. If not passed we will scope the API call to the default branch. (optional)
    timeseriesGranularity := "hourly" // string | The granularity for which to query timeseries data. (optional)
    startDate := time.Now() // time.Time | Include only executions that started at or after this date. This must be specified if an end-date is provided. (optional)
    endDate := time.Now() // time.Time | Include only executions that started before this date. This date can be at most 90 days after the start-date. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsApi.GetJobTimeseries(context.Background(), projectSlug, workflowName).Branch(branch).TimeseriesGranularity(timeseriesGranularity).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsApi.GetJobTimeseries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobTimeseries`: InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `InsightsApi.GetJobTimeseries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectSlug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 
**workflowName** | **string** | The name of the workflow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobTimeseriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **branch** | **string** | The name of a vcs branch. If not passed we will scope the API call to the default branch. | 
 **timeseriesGranularity** | **string** | The granularity for which to query timeseries data. | 
 **startDate** | **time.Time** | Include only executions that started at or after this date. This must be specified if an end-date is provided. | 
 **endDate** | **time.Time** | Include only executions that started before this date. This date can be at most 90 days after the start-date. | 

### Return type

[**InlineResponse2003**](InlineResponse2003.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgSummaryData

> InlineResponse2004 GetOrgSummaryData(ctx, orgSlug).ReportingWindow(reportingWindow).ProjectNames(projectNames).Execute()

Get summary metrics with trends for the entire org, and for each project.



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
    orgSlug := "gh/CircleCI-Public" // string | Org slug in the form `vcs-slug/org-name`. The `/` characters may be URL-escaped.
    reportingWindow := "last-90-days" // string | The time window used to calculate summary metrics. (optional)
    projectNames := map[string]interface{}{ ... } // map[string]interface{} | List of project names. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsApi.GetOrgSummaryData(context.Background(), orgSlug).ReportingWindow(reportingWindow).ProjectNames(projectNames).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsApi.GetOrgSummaryData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgSummaryData`: InlineResponse2004
    fmt.Fprintf(os.Stdout, "Response from `InsightsApi.GetOrgSummaryData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** | Org slug in the form &#x60;vcs-slug/org-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgSummaryDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportingWindow** | **string** | The time window used to calculate summary metrics. | 
 **projectNames** | [**map[string]interface{}**](map[string]interface{}.md) | List of project names. | 

### Return type

[**InlineResponse2004**](InlineResponse2004.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectJobRuns

> InlineResponse2009 GetProjectJobRuns(ctx, projectSlug, workflowName, jobName).Branch(branch).PageToken(pageToken).StartDate(startDate).EndDate(endDate).Execute()

Get recent runs of a workflow job



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
    projectSlug := "gh/CircleCI-Public/api-preview-docs" // string | Project slug in the form `vcs-slug/org-name/repo-name`. The `/` characters may be URL-escaped.
    workflowName := "build-and-test" // string | The name of the workflow.
    jobName := "lint" // string | The name of the job.
    branch := "branch_example" // string | The name of a vcs branch. If not passed we will scope the API call to the default branch.     Note - Querying all branches is not supported yet. (optional)
    pageToken := "pageToken_example" // string | A token to retrieve the next page of results. (optional)
    startDate := time.Now() // time.Time | Include only executions that started at or after this date. This must be specified if an end-date is provided. (optional)
    endDate := time.Now() // time.Time | Include only executions that started before this date. This date can be at most 90 days after the start-date. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsApi.GetProjectJobRuns(context.Background(), projectSlug, workflowName, jobName).Branch(branch).PageToken(pageToken).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsApi.GetProjectJobRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectJobRuns`: InlineResponse2009
    fmt.Fprintf(os.Stdout, "Response from `InsightsApi.GetProjectJobRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectSlug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 
**workflowName** | **string** | The name of the workflow. | 
**jobName** | **string** | The name of the job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectJobRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **branch** | **string** | The name of a vcs branch. If not passed we will scope the API call to the default branch.     Note - Querying all branches is not supported yet. | 
 **pageToken** | **string** | A token to retrieve the next page of results. | 
 **startDate** | **time.Time** | Include only executions that started at or after this date. This must be specified if an end-date is provided. | 
 **endDate** | **time.Time** | Include only executions that started before this date. This date can be at most 90 days after the start-date. | 

### Return type

[**InlineResponse2009**](InlineResponse2009.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectWorkflowJobMetrics

> InlineResponse2008 GetProjectWorkflowJobMetrics(ctx, projectSlug, workflowName).PageToken(pageToken).AllBranches(allBranches).Branch(branch).ReportingWindow(reportingWindow).Execute()

Get summary metrics for a project workflow's jobs.



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
    workflowName := "build-and-test" // string | The name of the workflow.
    pageToken := "pageToken_example" // string | A token to retrieve the next page of results. (optional)
    allBranches := true // bool | Whether to retrieve data for all branches combined. Use either this parameter OR the branch name parameter. (optional)
    branch := "branch_example" // string | The name of a vcs branch. If not passed we will scope the API call to the default branch. (optional)
    reportingWindow := "last-90-days" // string | The time window used to calculate summary metrics. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsApi.GetProjectWorkflowJobMetrics(context.Background(), projectSlug, workflowName).PageToken(pageToken).AllBranches(allBranches).Branch(branch).ReportingWindow(reportingWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsApi.GetProjectWorkflowJobMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectWorkflowJobMetrics`: InlineResponse2008
    fmt.Fprintf(os.Stdout, "Response from `InsightsApi.GetProjectWorkflowJobMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectSlug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 
**workflowName** | **string** | The name of the workflow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectWorkflowJobMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageToken** | **string** | A token to retrieve the next page of results. | 
 **allBranches** | **bool** | Whether to retrieve data for all branches combined. Use either this parameter OR the branch name parameter. | 
 **branch** | **string** | The name of a vcs branch. If not passed we will scope the API call to the default branch. | 
 **reportingWindow** | **string** | The time window used to calculate summary metrics. | 

### Return type

[**InlineResponse2008**](InlineResponse2008.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectWorkflowMetrics

> InlineResponse2006 GetProjectWorkflowMetrics(ctx, projectSlug).PageToken(pageToken).AllBranches(allBranches).Branch(branch).ReportingWindow(reportingWindow).Execute()

Get summary metrics for a project's workflows



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
    allBranches := true // bool | Whether to retrieve data for all branches combined. Use either this parameter OR the branch name parameter. (optional)
    branch := "branch_example" // string | The name of a vcs branch. If not passed we will scope the API call to the default branch. (optional)
    reportingWindow := "last-90-days" // string | The time window used to calculate summary metrics. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsApi.GetProjectWorkflowMetrics(context.Background(), projectSlug).PageToken(pageToken).AllBranches(allBranches).Branch(branch).ReportingWindow(reportingWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsApi.GetProjectWorkflowMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectWorkflowMetrics`: InlineResponse2006
    fmt.Fprintf(os.Stdout, "Response from `InsightsApi.GetProjectWorkflowMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectSlug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectWorkflowMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | A token to retrieve the next page of results. | 
 **allBranches** | **bool** | Whether to retrieve data for all branches combined. Use either this parameter OR the branch name parameter. | 
 **branch** | **string** | The name of a vcs branch. If not passed we will scope the API call to the default branch. | 
 **reportingWindow** | **string** | The time window used to calculate summary metrics. | 

### Return type

[**InlineResponse2006**](InlineResponse2006.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectWorkflowRuns

> InlineResponse2007 GetProjectWorkflowRuns(ctx, projectSlug, workflowName).AllBranches(allBranches).Branch(branch).PageToken(pageToken).StartDate(startDate).EndDate(endDate).Execute()

Get recent runs of a workflow



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
    projectSlug := "gh/CircleCI-Public/api-preview-docs" // string | Project slug in the form `vcs-slug/org-name/repo-name`. The `/` characters may be URL-escaped.
    workflowName := "build-and-test" // string | The name of the workflow.
    allBranches := true // bool | Whether to retrieve data for all branches combined. Use either this parameter OR the branch name parameter. (optional)
    branch := "branch_example" // string | The name of a vcs branch. If not passed we will scope the API call to the default branch. (optional)
    pageToken := "pageToken_example" // string | A token to retrieve the next page of results. (optional)
    startDate := time.Now() // time.Time | Include only executions that started at or after this date. This must be specified if an end-date is provided. (optional)
    endDate := time.Now() // time.Time | Include only executions that started before this date. This date can be at most 90 days after the start-date. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsApi.GetProjectWorkflowRuns(context.Background(), projectSlug, workflowName).AllBranches(allBranches).Branch(branch).PageToken(pageToken).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsApi.GetProjectWorkflowRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectWorkflowRuns`: InlineResponse2007
    fmt.Fprintf(os.Stdout, "Response from `InsightsApi.GetProjectWorkflowRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectSlug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 
**workflowName** | **string** | The name of the workflow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectWorkflowRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **allBranches** | **bool** | Whether to retrieve data for all branches combined. Use either this parameter OR the branch name parameter. | 
 **branch** | **string** | The name of a vcs branch. If not passed we will scope the API call to the default branch. | 
 **pageToken** | **string** | A token to retrieve the next page of results. | 
 **startDate** | **time.Time** | Include only executions that started at or after this date. This must be specified if an end-date is provided. | 
 **endDate** | **time.Time** | Include only executions that started before this date. This date can be at most 90 days after the start-date. | 

### Return type

[**InlineResponse2007**](InlineResponse2007.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectWorkflowTestMetrics

> InlineResponse20011 GetProjectWorkflowTestMetrics(ctx, projectSlug, workflowName).Branch(branch).AllBranches(allBranches).Execute()

Get test metrics for a project's workflows



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
    workflowName := "build-and-test" // string | The name of the workflow.
    branch := "branch_example" // string | The name of a vcs branch. If not passed we will scope the API call to the default branch. (optional)
    allBranches := true // bool | Whether to retrieve data for all branches combined. Use either this parameter OR the branch name parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsApi.GetProjectWorkflowTestMetrics(context.Background(), projectSlug, workflowName).Branch(branch).AllBranches(allBranches).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsApi.GetProjectWorkflowTestMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectWorkflowTestMetrics`: InlineResponse20011
    fmt.Fprintf(os.Stdout, "Response from `InsightsApi.GetProjectWorkflowTestMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectSlug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 
**workflowName** | **string** | The name of the workflow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectWorkflowTestMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **branch** | **string** | The name of a vcs branch. If not passed we will scope the API call to the default branch. | 
 **allBranches** | **bool** | Whether to retrieve data for all branches combined. Use either this parameter OR the branch name parameter. | 

### Return type

[**InlineResponse20011**](InlineResponse20011.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectWorkflowsPageData

> InlineResponse2002 GetProjectWorkflowsPageData(ctx, projectSlug).ReportingWindow(reportingWindow).Branches(branches).WorkflowNames(workflowNames).Execute()

Get summary metrics and trends for a project across it's workflows and branches



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
    reportingWindow := "last-90-days" // string | The time window used to calculate summary metrics. (optional)
    branches := map[string]interface{}{ ... } // map[string]interface{} | The names of VCS branches to include in branch-level workflow metrics. (optional)
    workflowNames := map[string]interface{}{ ... } // map[string]interface{} | The names of workflows to include in workflow-level metrics. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsApi.GetProjectWorkflowsPageData(context.Background(), projectSlug).ReportingWindow(reportingWindow).Branches(branches).WorkflowNames(workflowNames).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsApi.GetProjectWorkflowsPageData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectWorkflowsPageData`: InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `InsightsApi.GetProjectWorkflowsPageData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectSlug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectWorkflowsPageDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportingWindow** | **string** | The time window used to calculate summary metrics. | 
 **branches** | [**map[string]interface{}**](map[string]interface{}.md) | The names of VCS branches to include in branch-level workflow metrics. | 
 **workflowNames** | [**map[string]interface{}**](map[string]interface{}.md) | The names of workflows to include in workflow-level metrics. | 

### Return type

[**InlineResponse2002**](InlineResponse2002.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowSummary

> InlineResponse20010 GetWorkflowSummary(ctx, projectSlug, workflowName).AllBranches(allBranches).Branches(branches).Execute()

Get metrics and trends for workflows



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
    workflowName := "build-and-test" // string | The name of the workflow.
    allBranches := true // bool | Whether to retrieve data for all branches combined. Use either this parameter OR the branch name parameter. (optional)
    branches := map[string]interface{}{ ... } // map[string]interface{} | The names of VCS branches to include in branch-level workflow metrics. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsApi.GetWorkflowSummary(context.Background(), projectSlug, workflowName).AllBranches(allBranches).Branches(branches).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsApi.GetWorkflowSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowSummary`: InlineResponse20010
    fmt.Fprintf(os.Stdout, "Response from `InsightsApi.GetWorkflowSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectSlug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 
**workflowName** | **string** | The name of the workflow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **allBranches** | **bool** | Whether to retrieve data for all branches combined. Use either this parameter OR the branch name parameter. | 
 **branches** | [**map[string]interface{}**](map[string]interface{}.md) | The names of VCS branches to include in branch-level workflow metrics. | 

### Return type

[**InlineResponse20010**](InlineResponse20010.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowTimeseries

> InlineResponse2003 GetWorkflowTimeseries(ctx, projectSlug).Branch(branch).AllBranches(allBranches).TimeseriesGranularity(timeseriesGranularity).WorkflowName(workflowName).StartDate(startDate).EndDate(endDate).Execute()

Workflow timeseries data



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
    projectSlug := "gh/CircleCI-Public/api-preview-docs" // string | Project slug in the form `vcs-slug/org-name/repo-name`. The `/` characters may be URL-escaped.
    branch := "branch_example" // string | The name of a vcs branch. If not passed we will scope the API call to the default branch. (optional)
    allBranches := true // bool | Whether to retrieve data for all branches combined. Use either this parameter OR the branch name parameter. (optional)
    timeseriesGranularity := "hourly" // string | The granularity for which to query timeseries data. (optional)
    workflowName := "build-and-test" // string | The name of a workflow. If not passed we will scope the API call to the project. (optional)
    startDate := time.Now() // time.Time | Include only executions that started at or after this date. This must be specified if an end-date is provided. (optional)
    endDate := time.Now() // time.Time | Include only executions that started before this date. This date can be at most 90 days after the start-date. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsApi.GetWorkflowTimeseries(context.Background(), projectSlug).Branch(branch).AllBranches(allBranches).TimeseriesGranularity(timeseriesGranularity).WorkflowName(workflowName).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsApi.GetWorkflowTimeseries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowTimeseries`: InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `InsightsApi.GetWorkflowTimeseries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectSlug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowTimeseriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **branch** | **string** | The name of a vcs branch. If not passed we will scope the API call to the default branch. | 
 **allBranches** | **bool** | Whether to retrieve data for all branches combined. Use either this parameter OR the branch name parameter. | 
 **timeseriesGranularity** | **string** | The granularity for which to query timeseries data. | 
 **workflowName** | **string** | The name of a workflow. If not passed we will scope the API call to the project. | 
 **startDate** | **time.Time** | Include only executions that started at or after this date. This must be specified if an end-date is provided. | 
 **endDate** | **time.Time** | Include only executions that started before this date. This date can be at most 90 days after the start-date. | 

### Return type

[**InlineResponse2003**](InlineResponse2003.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

