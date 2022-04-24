# InlineResponse2008Metrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuccessRate** | **float32** |  | 
**TotalRuns** | **int64** | The total number of runs. | 
**FailedRuns** | **int64** | The number of failed runs. | 
**SuccessfulRuns** | **int64** | The number of successful runs. | 
**Throughput** | **float32** | The average number of runs per day. | 
**TotalCreditsUsed** | **int64** | The total credits consumed by the job in the aggregation window. Note that Insights is not a real time financial reporting tool and should not be used for credit reporting. | 
**DurationMetrics** | [**InlineResponse2008MetricsDurationMetrics**](InlineResponse2008MetricsDurationMetrics.md) |  | 

## Methods

### NewInlineResponse2008Metrics

`func NewInlineResponse2008Metrics(successRate float32, totalRuns int64, failedRuns int64, successfulRuns int64, throughput float32, totalCreditsUsed int64, durationMetrics InlineResponse2008MetricsDurationMetrics, ) *InlineResponse2008Metrics`

NewInlineResponse2008Metrics instantiates a new InlineResponse2008Metrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2008MetricsWithDefaults

`func NewInlineResponse2008MetricsWithDefaults() *InlineResponse2008Metrics`

NewInlineResponse2008MetricsWithDefaults instantiates a new InlineResponse2008Metrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccessRate

`func (o *InlineResponse2008Metrics) GetSuccessRate() float32`

GetSuccessRate returns the SuccessRate field if non-nil, zero value otherwise.

### GetSuccessRateOk

`func (o *InlineResponse2008Metrics) GetSuccessRateOk() (*float32, bool)`

GetSuccessRateOk returns a tuple with the SuccessRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRate

`func (o *InlineResponse2008Metrics) SetSuccessRate(v float32)`

SetSuccessRate sets SuccessRate field to given value.


### GetTotalRuns

`func (o *InlineResponse2008Metrics) GetTotalRuns() int64`

GetTotalRuns returns the TotalRuns field if non-nil, zero value otherwise.

### GetTotalRunsOk

`func (o *InlineResponse2008Metrics) GetTotalRunsOk() (*int64, bool)`

GetTotalRunsOk returns a tuple with the TotalRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRuns

`func (o *InlineResponse2008Metrics) SetTotalRuns(v int64)`

SetTotalRuns sets TotalRuns field to given value.


### GetFailedRuns

`func (o *InlineResponse2008Metrics) GetFailedRuns() int64`

GetFailedRuns returns the FailedRuns field if non-nil, zero value otherwise.

### GetFailedRunsOk

`func (o *InlineResponse2008Metrics) GetFailedRunsOk() (*int64, bool)`

GetFailedRunsOk returns a tuple with the FailedRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedRuns

`func (o *InlineResponse2008Metrics) SetFailedRuns(v int64)`

SetFailedRuns sets FailedRuns field to given value.


### GetSuccessfulRuns

`func (o *InlineResponse2008Metrics) GetSuccessfulRuns() int64`

GetSuccessfulRuns returns the SuccessfulRuns field if non-nil, zero value otherwise.

### GetSuccessfulRunsOk

`func (o *InlineResponse2008Metrics) GetSuccessfulRunsOk() (*int64, bool)`

GetSuccessfulRunsOk returns a tuple with the SuccessfulRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulRuns

`func (o *InlineResponse2008Metrics) SetSuccessfulRuns(v int64)`

SetSuccessfulRuns sets SuccessfulRuns field to given value.


### GetThroughput

`func (o *InlineResponse2008Metrics) GetThroughput() float32`

GetThroughput returns the Throughput field if non-nil, zero value otherwise.

### GetThroughputOk

`func (o *InlineResponse2008Metrics) GetThroughputOk() (*float32, bool)`

GetThroughputOk returns a tuple with the Throughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughput

`func (o *InlineResponse2008Metrics) SetThroughput(v float32)`

SetThroughput sets Throughput field to given value.


### GetTotalCreditsUsed

`func (o *InlineResponse2008Metrics) GetTotalCreditsUsed() int64`

GetTotalCreditsUsed returns the TotalCreditsUsed field if non-nil, zero value otherwise.

### GetTotalCreditsUsedOk

`func (o *InlineResponse2008Metrics) GetTotalCreditsUsedOk() (*int64, bool)`

GetTotalCreditsUsedOk returns a tuple with the TotalCreditsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCreditsUsed

`func (o *InlineResponse2008Metrics) SetTotalCreditsUsed(v int64)`

SetTotalCreditsUsed sets TotalCreditsUsed field to given value.


### GetDurationMetrics

`func (o *InlineResponse2008Metrics) GetDurationMetrics() InlineResponse2008MetricsDurationMetrics`

GetDurationMetrics returns the DurationMetrics field if non-nil, zero value otherwise.

### GetDurationMetricsOk

`func (o *InlineResponse2008Metrics) GetDurationMetricsOk() (*InlineResponse2008MetricsDurationMetrics, bool)`

GetDurationMetricsOk returns a tuple with the DurationMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMetrics

`func (o *InlineResponse2008Metrics) SetDurationMetrics(v InlineResponse2008MetricsDurationMetrics)`

SetDurationMetrics sets DurationMetrics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


