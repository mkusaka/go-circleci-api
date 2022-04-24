# PipelineListResponse1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]Pipeline1**](Pipeline1.md) |  | 
**NextPageToken** | **string** | A token to pass as a &#x60;page-token&#x60; query parameter to return the next page of results. | 

## Methods

### NewPipelineListResponse1

`func NewPipelineListResponse1(items []Pipeline1, nextPageToken string, ) *PipelineListResponse1`

NewPipelineListResponse1 instantiates a new PipelineListResponse1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineListResponse1WithDefaults

`func NewPipelineListResponse1WithDefaults() *PipelineListResponse1`

NewPipelineListResponse1WithDefaults instantiates a new PipelineListResponse1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *PipelineListResponse1) GetItems() []Pipeline1`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PipelineListResponse1) GetItemsOk() (*[]Pipeline1, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PipelineListResponse1) SetItems(v []Pipeline1)`

SetItems sets Items field to given value.


### GetNextPageToken

`func (o *PipelineListResponse1) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *PipelineListResponse1) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *PipelineListResponse1) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


