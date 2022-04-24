# InlineResponse2006Metrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalRuns** | **int64** | The total number of runs. | 
**SuccessfulRuns** | **int64** | The number of successful runs. | 
**Mttr** | **int64** | The mean time to recovery (mean time between failures and their next success) in seconds. | 
**TotalCreditsUsed** | **int64** | The total credits consumed by the workflow in the aggregation window. Note that Insights is not a real time financial reporting tool and should not be used for credit reporting. | 
**FailedRuns** | **int64** | The number of failed runs. | 
**SuccessRate** | **float32** |  | 
**DurationMetrics** | [**InlineResponse2006MetricsDurationMetrics**](InlineResponse2006MetricsDurationMetrics.md) |  | 
**TotalRecoveries** | **int64** | The number of recovered workflow executions per day. | 
**Throughput** | **float32** | The average number of runs per day. | 

## Methods

### NewInlineResponse2006Metrics

`func NewInlineResponse2006Metrics(totalRuns int64, successfulRuns int64, mttr int64, totalCreditsUsed int64, failedRuns int64, successRate float32, durationMetrics InlineResponse2006MetricsDurationMetrics, totalRecoveries int64, throughput float32, ) *InlineResponse2006Metrics`

NewInlineResponse2006Metrics instantiates a new InlineResponse2006Metrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2006MetricsWithDefaults

`func NewInlineResponse2006MetricsWithDefaults() *InlineResponse2006Metrics`

NewInlineResponse2006MetricsWithDefaults instantiates a new InlineResponse2006Metrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalRuns

`func (o *InlineResponse2006Metrics) GetTotalRuns() int64`

GetTotalRuns returns the TotalRuns field if non-nil, zero value otherwise.

### GetTotalRunsOk

`func (o *InlineResponse2006Metrics) GetTotalRunsOk() (*int64, bool)`

GetTotalRunsOk returns a tuple with the TotalRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRuns

`func (o *InlineResponse2006Metrics) SetTotalRuns(v int64)`

SetTotalRuns sets TotalRuns field to given value.


### GetSuccessfulRuns

`func (o *InlineResponse2006Metrics) GetSuccessfulRuns() int64`

GetSuccessfulRuns returns the SuccessfulRuns field if non-nil, zero value otherwise.

### GetSuccessfulRunsOk

`func (o *InlineResponse2006Metrics) GetSuccessfulRunsOk() (*int64, bool)`

GetSuccessfulRunsOk returns a tuple with the SuccessfulRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulRuns

`func (o *InlineResponse2006Metrics) SetSuccessfulRuns(v int64)`

SetSuccessfulRuns sets SuccessfulRuns field to given value.


### GetMttr

`func (o *InlineResponse2006Metrics) GetMttr() int64`

GetMttr returns the Mttr field if non-nil, zero value otherwise.

### GetMttrOk

`func (o *InlineResponse2006Metrics) GetMttrOk() (*int64, bool)`

GetMttrOk returns a tuple with the Mttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMttr

`func (o *InlineResponse2006Metrics) SetMttr(v int64)`

SetMttr sets Mttr field to given value.


### GetTotalCreditsUsed

`func (o *InlineResponse2006Metrics) GetTotalCreditsUsed() int64`

GetTotalCreditsUsed returns the TotalCreditsUsed field if non-nil, zero value otherwise.

### GetTotalCreditsUsedOk

`func (o *InlineResponse2006Metrics) GetTotalCreditsUsedOk() (*int64, bool)`

GetTotalCreditsUsedOk returns a tuple with the TotalCreditsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCreditsUsed

`func (o *InlineResponse2006Metrics) SetTotalCreditsUsed(v int64)`

SetTotalCreditsUsed sets TotalCreditsUsed field to given value.


### GetFailedRuns

`func (o *InlineResponse2006Metrics) GetFailedRuns() int64`

GetFailedRuns returns the FailedRuns field if non-nil, zero value otherwise.

### GetFailedRunsOk

`func (o *InlineResponse2006Metrics) GetFailedRunsOk() (*int64, bool)`

GetFailedRunsOk returns a tuple with the FailedRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedRuns

`func (o *InlineResponse2006Metrics) SetFailedRuns(v int64)`

SetFailedRuns sets FailedRuns field to given value.


### GetSuccessRate

`func (o *InlineResponse2006Metrics) GetSuccessRate() float32`

GetSuccessRate returns the SuccessRate field if non-nil, zero value otherwise.

### GetSuccessRateOk

`func (o *InlineResponse2006Metrics) GetSuccessRateOk() (*float32, bool)`

GetSuccessRateOk returns a tuple with the SuccessRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRate

`func (o *InlineResponse2006Metrics) SetSuccessRate(v float32)`

SetSuccessRate sets SuccessRate field to given value.


### GetDurationMetrics

`func (o *InlineResponse2006Metrics) GetDurationMetrics() InlineResponse2006MetricsDurationMetrics`

GetDurationMetrics returns the DurationMetrics field if non-nil, zero value otherwise.

### GetDurationMetricsOk

`func (o *InlineResponse2006Metrics) GetDurationMetricsOk() (*InlineResponse2006MetricsDurationMetrics, bool)`

GetDurationMetricsOk returns a tuple with the DurationMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMetrics

`func (o *InlineResponse2006Metrics) SetDurationMetrics(v InlineResponse2006MetricsDurationMetrics)`

SetDurationMetrics sets DurationMetrics field to given value.


### GetTotalRecoveries

`func (o *InlineResponse2006Metrics) GetTotalRecoveries() int64`

GetTotalRecoveries returns the TotalRecoveries field if non-nil, zero value otherwise.

### GetTotalRecoveriesOk

`func (o *InlineResponse2006Metrics) GetTotalRecoveriesOk() (*int64, bool)`

GetTotalRecoveriesOk returns a tuple with the TotalRecoveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecoveries

`func (o *InlineResponse2006Metrics) SetTotalRecoveries(v int64)`

SetTotalRecoveries sets TotalRecoveries field to given value.


### GetThroughput

`func (o *InlineResponse2006Metrics) GetThroughput() float32`

GetThroughput returns the Throughput field if non-nil, zero value otherwise.

### GetThroughputOk

`func (o *InlineResponse2006Metrics) GetThroughputOk() (*float32, bool)`

GetThroughputOk returns a tuple with the Throughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughput

`func (o *InlineResponse2006Metrics) SetThroughput(v float32)`

SetThroughput sets Throughput field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


