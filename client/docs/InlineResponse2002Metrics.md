# InlineResponse2002Metrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCreditsUsed** | **int64** | The total credits consumed over the current timeseries interval. | 
**P95DurationSecs** | **float32** | The 95th percentile duration among a group of workflow runs. | 
**TotalRuns** | **int64** | The total number of runs. | 
**SuccessRate** | **float32** |  | 

## Methods

### NewInlineResponse2002Metrics

`func NewInlineResponse2002Metrics(totalCreditsUsed int64, p95DurationSecs float32, totalRuns int64, successRate float32, ) *InlineResponse2002Metrics`

NewInlineResponse2002Metrics instantiates a new InlineResponse2002Metrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2002MetricsWithDefaults

`func NewInlineResponse2002MetricsWithDefaults() *InlineResponse2002Metrics`

NewInlineResponse2002MetricsWithDefaults instantiates a new InlineResponse2002Metrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCreditsUsed

`func (o *InlineResponse2002Metrics) GetTotalCreditsUsed() int64`

GetTotalCreditsUsed returns the TotalCreditsUsed field if non-nil, zero value otherwise.

### GetTotalCreditsUsedOk

`func (o *InlineResponse2002Metrics) GetTotalCreditsUsedOk() (*int64, bool)`

GetTotalCreditsUsedOk returns a tuple with the TotalCreditsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCreditsUsed

`func (o *InlineResponse2002Metrics) SetTotalCreditsUsed(v int64)`

SetTotalCreditsUsed sets TotalCreditsUsed field to given value.


### GetP95DurationSecs

`func (o *InlineResponse2002Metrics) GetP95DurationSecs() float32`

GetP95DurationSecs returns the P95DurationSecs field if non-nil, zero value otherwise.

### GetP95DurationSecsOk

`func (o *InlineResponse2002Metrics) GetP95DurationSecsOk() (*float32, bool)`

GetP95DurationSecsOk returns a tuple with the P95DurationSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP95DurationSecs

`func (o *InlineResponse2002Metrics) SetP95DurationSecs(v float32)`

SetP95DurationSecs sets P95DurationSecs field to given value.


### GetTotalRuns

`func (o *InlineResponse2002Metrics) GetTotalRuns() int64`

GetTotalRuns returns the TotalRuns field if non-nil, zero value otherwise.

### GetTotalRunsOk

`func (o *InlineResponse2002Metrics) GetTotalRunsOk() (*int64, bool)`

GetTotalRunsOk returns a tuple with the TotalRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRuns

`func (o *InlineResponse2002Metrics) SetTotalRuns(v int64)`

SetTotalRuns sets TotalRuns field to given value.


### GetSuccessRate

`func (o *InlineResponse2002Metrics) GetSuccessRate() float32`

GetSuccessRate returns the SuccessRate field if non-nil, zero value otherwise.

### GetSuccessRateOk

`func (o *InlineResponse2002Metrics) GetSuccessRateOk() (*float32, bool)`

GetSuccessRateOk returns a tuple with the SuccessRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRate

`func (o *InlineResponse2002Metrics) SetSuccessRate(v float32)`

SetSuccessRate sets SuccessRate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


