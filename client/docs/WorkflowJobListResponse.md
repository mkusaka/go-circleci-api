# WorkflowJobListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]Job**](Job.md) |  | 
**NextPageToken** | **string** | A token to pass as a &#x60;page-token&#x60; query parameter to return the next page of results. | 

## Methods

### NewWorkflowJobListResponse

`func NewWorkflowJobListResponse(items []Job, nextPageToken string, ) *WorkflowJobListResponse`

NewWorkflowJobListResponse instantiates a new WorkflowJobListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowJobListResponseWithDefaults

`func NewWorkflowJobListResponseWithDefaults() *WorkflowJobListResponse`

NewWorkflowJobListResponseWithDefaults instantiates a new WorkflowJobListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *WorkflowJobListResponse) GetItems() []Job`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *WorkflowJobListResponse) GetItemsOk() (*[]Job, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *WorkflowJobListResponse) SetItems(v []Job)`

SetItems sets Items field to given value.


### GetNextPageToken

`func (o *WorkflowJobListResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *WorkflowJobListResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *WorkflowJobListResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


