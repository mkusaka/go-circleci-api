# InlineResponse2003Metrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalRuns** | **int64** | The total number of runs. | 
**FailedRuns** | **int64** | The number of failed runs. | 
**SuccessfulRuns** | **int64** | The number of successful runs. | 
**Throughput** | **float32** | The average number of runs per day. | 
**MedianCreditsUsed** | **int64** | The median credits consumed over the current timeseries interval. | 
**TotalCreditsUsed** | **int64** | The total credits consumed over the current timeseries interval. | 
**DurationMetrics** | [**InlineResponse2003MetricsDurationMetrics**](InlineResponse2003MetricsDurationMetrics.md) |  | 

## Methods

### NewInlineResponse2003Metrics

`func NewInlineResponse2003Metrics(totalRuns int64, failedRuns int64, successfulRuns int64, throughput float32, medianCreditsUsed int64, totalCreditsUsed int64, durationMetrics InlineResponse2003MetricsDurationMetrics, ) *InlineResponse2003Metrics`

NewInlineResponse2003Metrics instantiates a new InlineResponse2003Metrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2003MetricsWithDefaults

`func NewInlineResponse2003MetricsWithDefaults() *InlineResponse2003Metrics`

NewInlineResponse2003MetricsWithDefaults instantiates a new InlineResponse2003Metrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalRuns

`func (o *InlineResponse2003Metrics) GetTotalRuns() int64`

GetTotalRuns returns the TotalRuns field if non-nil, zero value otherwise.

### GetTotalRunsOk

`func (o *InlineResponse2003Metrics) GetTotalRunsOk() (*int64, bool)`

GetTotalRunsOk returns a tuple with the TotalRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRuns

`func (o *InlineResponse2003Metrics) SetTotalRuns(v int64)`

SetTotalRuns sets TotalRuns field to given value.


### GetFailedRuns

`func (o *InlineResponse2003Metrics) GetFailedRuns() int64`

GetFailedRuns returns the FailedRuns field if non-nil, zero value otherwise.

### GetFailedRunsOk

`func (o *InlineResponse2003Metrics) GetFailedRunsOk() (*int64, bool)`

GetFailedRunsOk returns a tuple with the FailedRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedRuns

`func (o *InlineResponse2003Metrics) SetFailedRuns(v int64)`

SetFailedRuns sets FailedRuns field to given value.


### GetSuccessfulRuns

`func (o *InlineResponse2003Metrics) GetSuccessfulRuns() int64`

GetSuccessfulRuns returns the SuccessfulRuns field if non-nil, zero value otherwise.

### GetSuccessfulRunsOk

`func (o *InlineResponse2003Metrics) GetSuccessfulRunsOk() (*int64, bool)`

GetSuccessfulRunsOk returns a tuple with the SuccessfulRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulRuns

`func (o *InlineResponse2003Metrics) SetSuccessfulRuns(v int64)`

SetSuccessfulRuns sets SuccessfulRuns field to given value.


### GetThroughput

`func (o *InlineResponse2003Metrics) GetThroughput() float32`

GetThroughput returns the Throughput field if non-nil, zero value otherwise.

### GetThroughputOk

`func (o *InlineResponse2003Metrics) GetThroughputOk() (*float32, bool)`

GetThroughputOk returns a tuple with the Throughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughput

`func (o *InlineResponse2003Metrics) SetThroughput(v float32)`

SetThroughput sets Throughput field to given value.


### GetMedianCreditsUsed

`func (o *InlineResponse2003Metrics) GetMedianCreditsUsed() int64`

GetMedianCreditsUsed returns the MedianCreditsUsed field if non-nil, zero value otherwise.

### GetMedianCreditsUsedOk

`func (o *InlineResponse2003Metrics) GetMedianCreditsUsedOk() (*int64, bool)`

GetMedianCreditsUsedOk returns a tuple with the MedianCreditsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianCreditsUsed

`func (o *InlineResponse2003Metrics) SetMedianCreditsUsed(v int64)`

SetMedianCreditsUsed sets MedianCreditsUsed field to given value.


### GetTotalCreditsUsed

`func (o *InlineResponse2003Metrics) GetTotalCreditsUsed() int64`

GetTotalCreditsUsed returns the TotalCreditsUsed field if non-nil, zero value otherwise.

### GetTotalCreditsUsedOk

`func (o *InlineResponse2003Metrics) GetTotalCreditsUsedOk() (*int64, bool)`

GetTotalCreditsUsedOk returns a tuple with the TotalCreditsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCreditsUsed

`func (o *InlineResponse2003Metrics) SetTotalCreditsUsed(v int64)`

SetTotalCreditsUsed sets TotalCreditsUsed field to given value.


### GetDurationMetrics

`func (o *InlineResponse2003Metrics) GetDurationMetrics() InlineResponse2003MetricsDurationMetrics`

GetDurationMetrics returns the DurationMetrics field if non-nil, zero value otherwise.

### GetDurationMetricsOk

`func (o *InlineResponse2003Metrics) GetDurationMetricsOk() (*InlineResponse2003MetricsDurationMetrics, bool)`

GetDurationMetricsOk returns a tuple with the DurationMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMetrics

`func (o *InlineResponse2003Metrics) SetDurationMetrics(v InlineResponse2003MetricsDurationMetrics)`

SetDurationMetrics sets DurationMetrics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


