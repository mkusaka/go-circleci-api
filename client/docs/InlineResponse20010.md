# InlineResponse20010

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | [**InlineResponse20010Metrics**](InlineResponse20010Metrics.md) |  | 
**Trends** | [**InlineResponse20010Trends**](InlineResponse20010Trends.md) |  | 
**WorkflowNames** | **[]string** | A list of all the workflow names for a given project. | 

## Methods

### NewInlineResponse20010

`func NewInlineResponse20010(metrics InlineResponse20010Metrics, trends InlineResponse20010Trends, workflowNames []string, ) *InlineResponse20010`

NewInlineResponse20010 instantiates a new InlineResponse20010 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20010WithDefaults

`func NewInlineResponse20010WithDefaults() *InlineResponse20010`

NewInlineResponse20010WithDefaults instantiates a new InlineResponse20010 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *InlineResponse20010) GetMetrics() InlineResponse20010Metrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *InlineResponse20010) GetMetricsOk() (*InlineResponse20010Metrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *InlineResponse20010) SetMetrics(v InlineResponse20010Metrics)`

SetMetrics sets Metrics field to given value.


### GetTrends

`func (o *InlineResponse20010) GetTrends() InlineResponse20010Trends`

GetTrends returns the Trends field if non-nil, zero value otherwise.

### GetTrendsOk

`func (o *InlineResponse20010) GetTrendsOk() (*InlineResponse20010Trends, bool)`

GetTrendsOk returns a tuple with the Trends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrends

`func (o *InlineResponse20010) SetTrends(v InlineResponse20010Trends)`

SetTrends sets Trends field to given value.


### GetWorkflowNames

`func (o *InlineResponse20010) GetWorkflowNames() []string`

GetWorkflowNames returns the WorkflowNames field if non-nil, zero value otherwise.

### GetWorkflowNamesOk

`func (o *InlineResponse20010) GetWorkflowNamesOk() (*[]string, bool)`

GetWorkflowNamesOk returns a tuple with the WorkflowNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowNames

`func (o *InlineResponse20010) SetWorkflowNames(v []string)`

SetWorkflowNames sets WorkflowNames field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


