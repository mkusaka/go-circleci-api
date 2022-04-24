# EnvironmentVariableListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]EnvironmentVariablePair**](EnvironmentVariablePair.md) |  | 
**NextPageToken** | **string** | A token to pass as a &#x60;page-token&#x60; query parameter to return the next page of results. | 

## Methods

### NewEnvironmentVariableListResponse

`func NewEnvironmentVariableListResponse(items []EnvironmentVariablePair, nextPageToken string, ) *EnvironmentVariableListResponse`

NewEnvironmentVariableListResponse instantiates a new EnvironmentVariableListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentVariableListResponseWithDefaults

`func NewEnvironmentVariableListResponseWithDefaults() *EnvironmentVariableListResponse`

NewEnvironmentVariableListResponseWithDefaults instantiates a new EnvironmentVariableListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *EnvironmentVariableListResponse) GetItems() []EnvironmentVariablePair`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EnvironmentVariableListResponse) GetItemsOk() (*[]EnvironmentVariablePair, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EnvironmentVariableListResponse) SetItems(v []EnvironmentVariablePair)`

SetItems sets Items field to given value.


### GetNextPageToken

`func (o *EnvironmentVariableListResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *EnvironmentVariableListResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *EnvironmentVariableListResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


