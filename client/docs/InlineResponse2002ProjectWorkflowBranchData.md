# InlineResponse2002ProjectWorkflowBranchData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowName** | **string** | The name of the workflow. | 
**Branch** | **string** | The VCS branch of a workflow&#39;s trigger. | 
**Metrics** | [**InlineResponse2002Metrics**](InlineResponse2002Metrics.md) |  | 
**Trends** | [**InlineResponse2002Trends**](InlineResponse2002Trends.md) |  | 

## Methods

### NewInlineResponse2002ProjectWorkflowBranchData

`func NewInlineResponse2002ProjectWorkflowBranchData(workflowName string, branch string, metrics InlineResponse2002Metrics, trends InlineResponse2002Trends, ) *InlineResponse2002ProjectWorkflowBranchData`

NewInlineResponse2002ProjectWorkflowBranchData instantiates a new InlineResponse2002ProjectWorkflowBranchData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2002ProjectWorkflowBranchDataWithDefaults

`func NewInlineResponse2002ProjectWorkflowBranchDataWithDefaults() *InlineResponse2002ProjectWorkflowBranchData`

NewInlineResponse2002ProjectWorkflowBranchDataWithDefaults instantiates a new InlineResponse2002ProjectWorkflowBranchData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowName

`func (o *InlineResponse2002ProjectWorkflowBranchData) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *InlineResponse2002ProjectWorkflowBranchData) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *InlineResponse2002ProjectWorkflowBranchData) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.


### GetBranch

`func (o *InlineResponse2002ProjectWorkflowBranchData) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *InlineResponse2002ProjectWorkflowBranchData) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *InlineResponse2002ProjectWorkflowBranchData) SetBranch(v string)`

SetBranch sets Branch field to given value.


### GetMetrics

`func (o *InlineResponse2002ProjectWorkflowBranchData) GetMetrics() InlineResponse2002Metrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *InlineResponse2002ProjectWorkflowBranchData) GetMetricsOk() (*InlineResponse2002Metrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *InlineResponse2002ProjectWorkflowBranchData) SetMetrics(v InlineResponse2002Metrics)`

SetMetrics sets Metrics field to given value.


### GetTrends

`func (o *InlineResponse2002ProjectWorkflowBranchData) GetTrends() InlineResponse2002Trends`

GetTrends returns the Trends field if non-nil, zero value otherwise.

### GetTrendsOk

`func (o *InlineResponse2002ProjectWorkflowBranchData) GetTrendsOk() (*InlineResponse2002Trends, bool)`

GetTrendsOk returns a tuple with the Trends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrends

`func (o *InlineResponse2002ProjectWorkflowBranchData) SetTrends(v InlineResponse2002Trends)`

SetTrends sets Trends field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


