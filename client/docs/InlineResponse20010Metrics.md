# InlineResponse20010Metrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalRuns** | **int64** | The total number of runs. | 
**SuccessfulRuns** | **int64** | The number of successful runs. | 
**Mttr** | **int64** | The mean time to recovery (mean time between failures and their next success) in seconds. | 
**TotalCreditsUsed** | **int64** | The total credits consumed by the workflow in the aggregation window. Note that Insights is not a real time financial reporting tool and should not be used for credit reporting. | 
**FailedRuns** | **int64** | The number of failed runs. | 
**SuccessRate** | **float32** |  | 
**WindowStart** | **time.Time** | The start of the aggregation window for workflow metrics. | 
**DurationMetrics** | [**InlineResponse2006MetricsDurationMetrics**](InlineResponse2006MetricsDurationMetrics.md) |  | 
**WindowEnd** | **time.Time** | The end of the aggregation window for workflow metrics. | 
**Throughput** | **float32** | The average number of runs per day. | 

## Methods

### NewInlineResponse20010Metrics

`func NewInlineResponse20010Metrics(totalRuns int64, successfulRuns int64, mttr int64, totalCreditsUsed int64, failedRuns int64, successRate float32, windowStart time.Time, durationMetrics InlineResponse2006MetricsDurationMetrics, windowEnd time.Time, throughput float32, ) *InlineResponse20010Metrics`

NewInlineResponse20010Metrics instantiates a new InlineResponse20010Metrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20010MetricsWithDefaults

`func NewInlineResponse20010MetricsWithDefaults() *InlineResponse20010Metrics`

NewInlineResponse20010MetricsWithDefaults instantiates a new InlineResponse20010Metrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalRuns

`func (o *InlineResponse20010Metrics) GetTotalRuns() int64`

GetTotalRuns returns the TotalRuns field if non-nil, zero value otherwise.

### GetTotalRunsOk

`func (o *InlineResponse20010Metrics) GetTotalRunsOk() (*int64, bool)`

GetTotalRunsOk returns a tuple with the TotalRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRuns

`func (o *InlineResponse20010Metrics) SetTotalRuns(v int64)`

SetTotalRuns sets TotalRuns field to given value.


### GetSuccessfulRuns

`func (o *InlineResponse20010Metrics) GetSuccessfulRuns() int64`

GetSuccessfulRuns returns the SuccessfulRuns field if non-nil, zero value otherwise.

### GetSuccessfulRunsOk

`func (o *InlineResponse20010Metrics) GetSuccessfulRunsOk() (*int64, bool)`

GetSuccessfulRunsOk returns a tuple with the SuccessfulRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulRuns

`func (o *InlineResponse20010Metrics) SetSuccessfulRuns(v int64)`

SetSuccessfulRuns sets SuccessfulRuns field to given value.


### GetMttr

`func (o *InlineResponse20010Metrics) GetMttr() int64`

GetMttr returns the Mttr field if non-nil, zero value otherwise.

### GetMttrOk

`func (o *InlineResponse20010Metrics) GetMttrOk() (*int64, bool)`

GetMttrOk returns a tuple with the Mttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMttr

`func (o *InlineResponse20010Metrics) SetMttr(v int64)`

SetMttr sets Mttr field to given value.


### GetTotalCreditsUsed

`func (o *InlineResponse20010Metrics) GetTotalCreditsUsed() int64`

GetTotalCreditsUsed returns the TotalCreditsUsed field if non-nil, zero value otherwise.

### GetTotalCreditsUsedOk

`func (o *InlineResponse20010Metrics) GetTotalCreditsUsedOk() (*int64, bool)`

GetTotalCreditsUsedOk returns a tuple with the TotalCreditsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCreditsUsed

`func (o *InlineResponse20010Metrics) SetTotalCreditsUsed(v int64)`

SetTotalCreditsUsed sets TotalCreditsUsed field to given value.


### GetFailedRuns

`func (o *InlineResponse20010Metrics) GetFailedRuns() int64`

GetFailedRuns returns the FailedRuns field if non-nil, zero value otherwise.

### GetFailedRunsOk

`func (o *InlineResponse20010Metrics) GetFailedRunsOk() (*int64, bool)`

GetFailedRunsOk returns a tuple with the FailedRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedRuns

`func (o *InlineResponse20010Metrics) SetFailedRuns(v int64)`

SetFailedRuns sets FailedRuns field to given value.


### GetSuccessRate

`func (o *InlineResponse20010Metrics) GetSuccessRate() float32`

GetSuccessRate returns the SuccessRate field if non-nil, zero value otherwise.

### GetSuccessRateOk

`func (o *InlineResponse20010Metrics) GetSuccessRateOk() (*float32, bool)`

GetSuccessRateOk returns a tuple with the SuccessRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRate

`func (o *InlineResponse20010Metrics) SetSuccessRate(v float32)`

SetSuccessRate sets SuccessRate field to given value.


### GetWindowStart

`func (o *InlineResponse20010Metrics) GetWindowStart() time.Time`

GetWindowStart returns the WindowStart field if non-nil, zero value otherwise.

### GetWindowStartOk

`func (o *InlineResponse20010Metrics) GetWindowStartOk() (*time.Time, bool)`

GetWindowStartOk returns a tuple with the WindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowStart

`func (o *InlineResponse20010Metrics) SetWindowStart(v time.Time)`

SetWindowStart sets WindowStart field to given value.


### GetDurationMetrics

`func (o *InlineResponse20010Metrics) GetDurationMetrics() InlineResponse2006MetricsDurationMetrics`

GetDurationMetrics returns the DurationMetrics field if non-nil, zero value otherwise.

### GetDurationMetricsOk

`func (o *InlineResponse20010Metrics) GetDurationMetricsOk() (*InlineResponse2006MetricsDurationMetrics, bool)`

GetDurationMetricsOk returns a tuple with the DurationMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMetrics

`func (o *InlineResponse20010Metrics) SetDurationMetrics(v InlineResponse2006MetricsDurationMetrics)`

SetDurationMetrics sets DurationMetrics field to given value.


### GetWindowEnd

`func (o *InlineResponse20010Metrics) GetWindowEnd() time.Time`

GetWindowEnd returns the WindowEnd field if non-nil, zero value otherwise.

### GetWindowEndOk

`func (o *InlineResponse20010Metrics) GetWindowEndOk() (*time.Time, bool)`

GetWindowEndOk returns a tuple with the WindowEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowEnd

`func (o *InlineResponse20010Metrics) SetWindowEnd(v time.Time)`

SetWindowEnd sets WindowEnd field to given value.


### GetThroughput

`func (o *InlineResponse20010Metrics) GetThroughput() float32`

GetThroughput returns the Throughput field if non-nil, zero value otherwise.

### GetThroughputOk

`func (o *InlineResponse20010Metrics) GetThroughputOk() (*float32, bool)`

GetThroughputOk returns a tuple with the Throughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughput

`func (o *InlineResponse20010Metrics) SetThroughput(v float32)`

SetThroughput sets Throughput field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


