# \ProjectApi

All URIs are relative to *https://circleci.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCheckoutKey**](ProjectApi.md#CreateCheckoutKey) | **Post** /project/{project-slug}/checkout-key | Create a new checkout key
[**CreateEnvVar**](ProjectApi.md#CreateEnvVar) | **Post** /project/{project-slug}/envvar | Create an environment variable
[**DeleteCheckoutKey**](ProjectApi.md#DeleteCheckoutKey) | **Delete** /project/{project-slug}/checkout-key/{fingerprint} | Delete a checkout key
[**DeleteEnvVar**](ProjectApi.md#DeleteEnvVar) | **Delete** /project/{project-slug}/envvar/{name} | Delete an environment variable
[**GetCheckoutKey**](ProjectApi.md#GetCheckoutKey) | **Get** /project/{project-slug}/checkout-key/{fingerprint} | Get a checkout key
[**GetEnvVar**](ProjectApi.md#GetEnvVar) | **Get** /project/{project-slug}/envvar/{name} | Get a masked environment variable
[**GetProjectBySlug**](ProjectApi.md#GetProjectBySlug) | **Get** /project/{project-slug} | Get a project
[**ListCheckoutKeys**](ProjectApi.md#ListCheckoutKeys) | **Get** /project/{project-slug}/checkout-key | Get all checkout keys
[**ListEnvVars**](ProjectApi.md#ListEnvVars) | **Get** /project/{project-slug}/envvar | List all environment variables



## CreateCheckoutKey

> CheckoutKey CreateCheckoutKey(ctx, projectSlug).CheckoutKeyInput(checkoutKeyInput).Execute()

Create a new checkout key



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
    checkoutKeyInput := *openapiclient.NewCheckoutKeyInput("deploy-key") // CheckoutKeyInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.CreateCheckoutKey(context.Background(), projectSlug).CheckoutKeyInput(checkoutKeyInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.CreateCheckoutKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCheckoutKey`: CheckoutKey
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.CreateCheckoutKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectSlug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCheckoutKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checkoutKeyInput** | [**CheckoutKeyInput**](CheckoutKeyInput.md) |  | 

### Return type

[**CheckoutKey**](CheckoutKey.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEnvVar

> EnvironmentVariablePair1 CreateEnvVar(ctx, projectSlug).EnvironmentVariablePair1(environmentVariablePair1).Execute()

Create an environment variable



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
    environmentVariablePair1 := *openapiclient.NewEnvironmentVariablePair1("foo", "xxxx1234") // EnvironmentVariablePair1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.CreateEnvVar(context.Background(), projectSlug).EnvironmentVariablePair1(environmentVariablePair1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.CreateEnvVar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnvVar`: EnvironmentVariablePair1
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.CreateEnvVar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectSlug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvVarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentVariablePair1** | [**EnvironmentVariablePair1**](EnvironmentVariablePair1.md) |  | 

### Return type

[**EnvironmentVariablePair1**](EnvironmentVariablePair1.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCheckoutKey

> MessageResponse DeleteCheckoutKey(ctx, projectSlug, fingerprint).Execute()

Delete a checkout key



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
    fingerprint := "c9:0b:1c:4f:d5:65:56:b9:ad:88:f9:81:2b:37:74:2f" // string | An SSH key fingerprint.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.DeleteCheckoutKey(context.Background(), projectSlug, fingerprint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.DeleteCheckoutKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCheckoutKey`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.DeleteCheckoutKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectSlug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 
**fingerprint** | **string** | An SSH key fingerprint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCheckoutKeyRequest struct via the builder pattern


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


## DeleteEnvVar

> MessageResponse DeleteEnvVar(ctx, projectSlug, name).Execute()

Delete an environment variable



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
    name := "foo" // string | The name of the environment variable.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.DeleteEnvVar(context.Background(), projectSlug, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.DeleteEnvVar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteEnvVar`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.DeleteEnvVar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectSlug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 
**name** | **string** | The name of the environment variable. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvVarRequest struct via the builder pattern


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


## GetCheckoutKey

> CheckoutKey GetCheckoutKey(ctx, projectSlug, fingerprint).Execute()

Get a checkout key



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
    fingerprint := "c9:0b:1c:4f:d5:65:56:b9:ad:88:f9:81:2b:37:74:2f" // string | An SSH key fingerprint.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.GetCheckoutKey(context.Background(), projectSlug, fingerprint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.GetCheckoutKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCheckoutKey`: CheckoutKey
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.GetCheckoutKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectSlug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 
**fingerprint** | **string** | An SSH key fingerprint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCheckoutKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CheckoutKey**](CheckoutKey.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvVar

> EnvironmentVariablePair1 GetEnvVar(ctx, projectSlug, name).Execute()

Get a masked environment variable



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
    name := "foo" // string | The name of the environment variable.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.GetEnvVar(context.Background(), projectSlug, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.GetEnvVar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvVar`: EnvironmentVariablePair1
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.GetEnvVar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectSlug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 
**name** | **string** | The name of the environment variable. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvVarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EnvironmentVariablePair1**](EnvironmentVariablePair1.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectBySlug

> Project GetProjectBySlug(ctx, projectSlug).Execute()

Get a project



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
    resp, r, err := apiClient.ProjectApi.GetProjectBySlug(context.Background(), projectSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.GetProjectBySlug``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectBySlug`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.GetProjectBySlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectSlug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectBySlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Project**](Project.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCheckoutKeys

> CheckoutKeyListResponse ListCheckoutKeys(ctx, projectSlug).Execute()

Get all checkout keys



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
    resp, r, err := apiClient.ProjectApi.ListCheckoutKeys(context.Background(), projectSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ListCheckoutKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCheckoutKeys`: CheckoutKeyListResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ListCheckoutKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectSlug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCheckoutKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CheckoutKeyListResponse**](CheckoutKeyListResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEnvVars

> EnvironmentVariableListResponse ListEnvVars(ctx, projectSlug).Execute()

List all environment variables



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
    resp, r, err := apiClient.ProjectApi.ListEnvVars(context.Background(), projectSlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ListEnvVars``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEnvVars`: EnvironmentVariableListResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ListEnvVars`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectSlug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEnvVarsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentVariableListResponse**](EnvironmentVariableListResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

