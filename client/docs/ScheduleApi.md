# \ScheduleApi

All URIs are relative to *https://circleci.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSchedule**](ScheduleApi.md#CreateSchedule) | **Post** /project/{project-slug}/schedule | Create a schedule
[**DeleteScheduleById**](ScheduleApi.md#DeleteScheduleById) | **Delete** /schedule/{schedule-id} | Delete a schedule
[**GetScheduleById**](ScheduleApi.md#GetScheduleById) | **Get** /schedule/{schedule-id} | Get a schedule
[**ListSchedulesForProject**](ScheduleApi.md#ListSchedulesForProject) | **Get** /project/{project-slug}/schedule | Get all schedules
[**UpdateSchedule**](ScheduleApi.md#UpdateSchedule) | **Patch** /schedule/{schedule-id} | Update a schedule



## CreateSchedule

> Schedule CreateSchedule(ctx, projectSlug).CreateScheduleParameters(createScheduleParameters).Execute()

Create a schedule



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
    createScheduleParameters := *openapiclient.NewCreateScheduleParameters("Name_example", *openapiclient.NewInlineResponse20012Timetable(int32(123), []int32{int32(123)}, []string{"DaysOfWeek_example"}), "current", map[string]string{"key": "Inner_example"}) // CreateScheduleParameters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduleApi.CreateSchedule(context.Background(), projectSlug).CreateScheduleParameters(createScheduleParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleApi.CreateSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSchedule`: Schedule
    fmt.Fprintf(os.Stdout, "Response from `ScheduleApi.CreateSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectSlug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createScheduleParameters** | [**CreateScheduleParameters**](CreateScheduleParameters.md) |  | 

### Return type

[**Schedule**](Schedule.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteScheduleById

> MessageResponse DeleteScheduleById(ctx, scheduleId).Execute()

Delete a schedule



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
    scheduleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the schedule.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduleApi.DeleteScheduleById(context.Background(), scheduleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleApi.DeleteScheduleById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteScheduleById`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `ScheduleApi.DeleteScheduleById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduleId** | **string** | The unique ID of the schedule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScheduleByIdRequest struct via the builder pattern


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


## GetScheduleById

> Schedule GetScheduleById(ctx, scheduleId).Execute()

Get a schedule



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
    scheduleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the schedule.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduleApi.GetScheduleById(context.Background(), scheduleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleApi.GetScheduleById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScheduleById`: Schedule
    fmt.Fprintf(os.Stdout, "Response from `ScheduleApi.GetScheduleById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduleId** | **string** | The unique ID of the schedule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScheduleByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Schedule**](Schedule.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSchedulesForProject

> InlineResponse20012 ListSchedulesForProject(ctx, projectSlug).PageToken(pageToken).Execute()

Get all schedules



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
    resp, r, err := apiClient.ScheduleApi.ListSchedulesForProject(context.Background(), projectSlug).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleApi.ListSchedulesForProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSchedulesForProject`: InlineResponse20012
    fmt.Fprintf(os.Stdout, "Response from `ScheduleApi.ListSchedulesForProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectSlug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSchedulesForProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | A token to retrieve the next page of results. | 

### Return type

[**InlineResponse20012**](InlineResponse20012.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSchedule

> Schedule UpdateSchedule(ctx, scheduleId).UpdateScheduleParameters(updateScheduleParameters).Execute()

Update a schedule



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
    scheduleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the schedule.
    updateScheduleParameters := *openapiclient.NewUpdateScheduleParameters() // UpdateScheduleParameters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduleApi.UpdateSchedule(context.Background(), scheduleId).UpdateScheduleParameters(updateScheduleParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleApi.UpdateSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSchedule`: Schedule
    fmt.Fprintf(os.Stdout, "Response from `ScheduleApi.UpdateSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduleId** | **string** | The unique ID of the schedule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateScheduleParameters** | [**UpdateScheduleParameters**](UpdateScheduleParameters.md) |  | 

### Return type

[**Schedule**](Schedule.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

