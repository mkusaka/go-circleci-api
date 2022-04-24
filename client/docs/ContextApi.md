# \ContextApi

All URIs are relative to *https://circleci.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddEnvironmentVariableToContext**](ContextApi.md#AddEnvironmentVariableToContext) | **Put** /context/{context-id}/environment-variable/{env-var-name} | Add or update an environment variable
[**CreateContext**](ContextApi.md#CreateContext) | **Post** /context | Create a new context
[**DeleteContext**](ContextApi.md#DeleteContext) | **Delete** /context/{context-id} | Delete a context
[**DeleteEnvironmentVariableFromContext**](ContextApi.md#DeleteEnvironmentVariableFromContext) | **Delete** /context/{context-id}/environment-variable/{env-var-name} | Remove an environment variable
[**GetContext**](ContextApi.md#GetContext) | **Get** /context/{context-id} | Get a context
[**ListContexts**](ContextApi.md#ListContexts) | **Get** /context | List contexts
[**ListEnvironmentVariablesFromContext**](ContextApi.md#ListEnvironmentVariablesFromContext) | **Get** /context/{context-id}/environment-variable | List environment variables



## AddEnvironmentVariableToContext

> InlineResponse2001Items AddEnvironmentVariableToContext(ctx, contextId, envVarName).InlineObject1(inlineObject1).Execute()

Add or update an environment variable



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
    contextId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the context (UUID)
    envVarName := "POSTGRES_USER" // string | The name of the environment variable
    inlineObject1 := *openapiclient.NewInlineObject1("some-secret-value") // InlineObject1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextApi.AddEnvironmentVariableToContext(context.Background(), contextId, envVarName).InlineObject1(inlineObject1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextApi.AddEnvironmentVariableToContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddEnvironmentVariableToContext`: InlineResponse2001Items
    fmt.Fprintf(os.Stdout, "Response from `ContextApi.AddEnvironmentVariableToContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextId** | **string** | ID of the context (UUID) | 
**envVarName** | **string** | The name of the environment variable | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddEnvironmentVariableToContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject1** | [**InlineObject1**](InlineObject1.md) |  | 

### Return type

[**InlineResponse2001Items**](InlineResponse2001Items.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContext

> Context CreateContext(ctx).InlineObject(inlineObject).Execute()

Create a new context

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
    inlineObject := *openapiclient.NewInlineObject("Name_example", *openapiclient.NewContextOwner()) // InlineObject |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextApi.CreateContext(context.Background()).InlineObject(inlineObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextApi.CreateContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContext`: Context
    fmt.Fprintf(os.Stdout, "Response from `ContextApi.CreateContext`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject** | [**InlineObject**](InlineObject.md) |  | 

### Return type

[**Context**](Context.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContext

> MessageResponse DeleteContext(ctx, contextId).Execute()

Delete a context

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
    contextId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the context (UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextApi.DeleteContext(context.Background(), contextId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextApi.DeleteContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteContext`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `ContextApi.DeleteContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextId** | **string** | ID of the context (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContextRequest struct via the builder pattern


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


## DeleteEnvironmentVariableFromContext

> MessageResponse DeleteEnvironmentVariableFromContext(ctx, envVarName, contextId).Execute()

Remove an environment variable



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
    envVarName := "POSTGRES_USER" // string | The name of the environment variable
    contextId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the context (UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextApi.DeleteEnvironmentVariableFromContext(context.Background(), envVarName, contextId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextApi.DeleteEnvironmentVariableFromContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteEnvironmentVariableFromContext`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `ContextApi.DeleteEnvironmentVariableFromContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envVarName** | **string** | The name of the environment variable | 
**contextId** | **string** | ID of the context (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentVariableFromContextRequest struct via the builder pattern


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


## GetContext

> Context GetContext(ctx, contextId).Execute()

Get a context



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
    contextId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the context (UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextApi.GetContext(context.Background(), contextId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextApi.GetContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContext`: Context
    fmt.Fprintf(os.Stdout, "Response from `ContextApi.GetContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextId** | **string** | ID of the context (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Context**](Context.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContexts

> InlineResponse200 ListContexts(ctx).OwnerId(ownerId).OwnerSlug(ownerSlug).OwnerType(ownerType).PageToken(pageToken).Execute()

List contexts



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
    ownerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the owner of the context. Specify either this or owner-slug. (optional)
    ownerSlug := "ownerSlug_example" // string | A string that represents an organization. Specify either this or owner-id. Cannot be used for accounts. (optional)
    ownerType := "ownerType_example" // string | The type of the owner. Defaults to \"organization\". Accounts are only used as context owners in server. (optional)
    pageToken := "pageToken_example" // string | A token to retrieve the next page of results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextApi.ListContexts(context.Background()).OwnerId(ownerId).OwnerSlug(ownerSlug).OwnerType(ownerType).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextApi.ListContexts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListContexts`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `ContextApi.ListContexts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListContextsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ownerId** | **string** | The unique ID of the owner of the context. Specify either this or owner-slug. | 
 **ownerSlug** | **string** | A string that represents an organization. Specify either this or owner-id. Cannot be used for accounts. | 
 **ownerType** | **string** | The type of the owner. Defaults to \&quot;organization\&quot;. Accounts are only used as context owners in server. | 
 **pageToken** | **string** | A token to retrieve the next page of results. | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEnvironmentVariablesFromContext

> InlineResponse2001 ListEnvironmentVariablesFromContext(ctx, contextId).Execute()

List environment variables



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
    contextId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the context (UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextApi.ListEnvironmentVariablesFromContext(context.Background(), contextId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextApi.ListEnvironmentVariablesFromContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEnvironmentVariablesFromContext`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `ContextApi.ListEnvironmentVariablesFromContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextId** | **string** | ID of the context (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEnvironmentVariablesFromContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

