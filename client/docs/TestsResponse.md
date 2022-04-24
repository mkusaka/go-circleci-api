# TestsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]TestsResponseItems**](TestsResponseItems.md) |  | 
**NextPageToken** | **string** | A token to pass as a &#x60;page-token&#x60; query parameter to return the next page of results. | 

## Methods

### NewTestsResponse

`func NewTestsResponse(items []TestsResponseItems, nextPageToken string, ) *TestsResponse`

NewTestsResponse instantiates a new TestsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestsResponseWithDefaults

`func NewTestsResponseWithDefaults() *TestsResponse`

NewTestsResponseWithDefaults instantiates a new TestsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *TestsResponse) GetItems() []TestsResponseItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TestsResponse) GetItemsOk() (*[]TestsResponseItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TestsResponse) SetItems(v []TestsResponseItems)`

SetItems sets Items field to given value.


### GetNextPageToken

`func (o *TestsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *TestsResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *TestsResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


