# InlineResponse20011TestRuns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PipelineNumber** | **int64** | The number of the pipeline associated with the provided test counts | 
**WorkflowId** | **interface{}** | The ID of the workflow associated with the provided test counts | 
**SuccessRate** | **float32** | The success rate calculated from test counts | 
**TestCounts** | [**InlineResponse20011TestCounts**](InlineResponse20011TestCounts.md) |  | 

## Methods

### NewInlineResponse20011TestRuns

`func NewInlineResponse20011TestRuns(pipelineNumber int64, workflowId interface{}, successRate float32, testCounts InlineResponse20011TestCounts, ) *InlineResponse20011TestRuns`

NewInlineResponse20011TestRuns instantiates a new InlineResponse20011TestRuns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20011TestRunsWithDefaults

`func NewInlineResponse20011TestRunsWithDefaults() *InlineResponse20011TestRuns`

NewInlineResponse20011TestRunsWithDefaults instantiates a new InlineResponse20011TestRuns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPipelineNumber

`func (o *InlineResponse20011TestRuns) GetPipelineNumber() int64`

GetPipelineNumber returns the PipelineNumber field if non-nil, zero value otherwise.

### GetPipelineNumberOk

`func (o *InlineResponse20011TestRuns) GetPipelineNumberOk() (*int64, bool)`

GetPipelineNumberOk returns a tuple with the PipelineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineNumber

`func (o *InlineResponse20011TestRuns) SetPipelineNumber(v int64)`

SetPipelineNumber sets PipelineNumber field to given value.


### GetWorkflowId

`func (o *InlineResponse20011TestRuns) GetWorkflowId() interface{}`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *InlineResponse20011TestRuns) GetWorkflowIdOk() (*interface{}, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *InlineResponse20011TestRuns) SetWorkflowId(v interface{})`

SetWorkflowId sets WorkflowId field to given value.


### SetWorkflowIdNil

`func (o *InlineResponse20011TestRuns) SetWorkflowIdNil(b bool)`

 SetWorkflowIdNil sets the value for WorkflowId to be an explicit nil

### UnsetWorkflowId
`func (o *InlineResponse20011TestRuns) UnsetWorkflowId()`

UnsetWorkflowId ensures that no value is present for WorkflowId, not even an explicit nil
### GetSuccessRate

`func (o *InlineResponse20011TestRuns) GetSuccessRate() float32`

GetSuccessRate returns the SuccessRate field if non-nil, zero value otherwise.

### GetSuccessRateOk

`func (o *InlineResponse20011TestRuns) GetSuccessRateOk() (*float32, bool)`

GetSuccessRateOk returns a tuple with the SuccessRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRate

`func (o *InlineResponse20011TestRuns) SetSuccessRate(v float32)`

SetSuccessRate sets SuccessRate field to given value.


### GetTestCounts

`func (o *InlineResponse20011TestRuns) GetTestCounts() InlineResponse20011TestCounts`

GetTestCounts returns the TestCounts field if non-nil, zero value otherwise.

### GetTestCountsOk

`func (o *InlineResponse20011TestRuns) GetTestCountsOk() (*InlineResponse20011TestCounts, bool)`

GetTestCountsOk returns a tuple with the TestCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCounts

`func (o *InlineResponse20011TestRuns) SetTestCounts(v InlineResponse20011TestCounts)`

SetTestCounts sets TestCounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


