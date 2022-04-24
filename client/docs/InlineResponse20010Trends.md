# InlineResponse20010Trends

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalRuns** | **float32** | The trend value for total number of runs. | 
**FailedRuns** | **float32** | The trend value for number of failed runs. | 
**SuccessRate** | **float32** | The trend value for the success rate. | 
**P95DurationSecs** | **float32** | Trend value for the 95th percentile duration for a workflow for a given time window. | 
**MedianDurationSecs** | **float32** | Trend value for the 50th percentile duration for a workflow for a given time window. | 
**TotalCreditsUsed** | **float32** | The trend value for total credits consumed. | 
**Mttr** | **float32** | trend for mean time to recovery (mean time between failures and their next success). | 
**Throughput** | **float32** | Trend value for the average number of runs per day. | 

## Methods

### NewInlineResponse20010Trends

`func NewInlineResponse20010Trends(totalRuns float32, failedRuns float32, successRate float32, p95DurationSecs float32, medianDurationSecs float32, totalCreditsUsed float32, mttr float32, throughput float32, ) *InlineResponse20010Trends`

NewInlineResponse20010Trends instantiates a new InlineResponse20010Trends object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20010TrendsWithDefaults

`func NewInlineResponse20010TrendsWithDefaults() *InlineResponse20010Trends`

NewInlineResponse20010TrendsWithDefaults instantiates a new InlineResponse20010Trends object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalRuns

`func (o *InlineResponse20010Trends) GetTotalRuns() float32`

GetTotalRuns returns the TotalRuns field if non-nil, zero value otherwise.

### GetTotalRunsOk

`func (o *InlineResponse20010Trends) GetTotalRunsOk() (*float32, bool)`

GetTotalRunsOk returns a tuple with the TotalRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRuns

`func (o *InlineResponse20010Trends) SetTotalRuns(v float32)`

SetTotalRuns sets TotalRuns field to given value.


### GetFailedRuns

`func (o *InlineResponse20010Trends) GetFailedRuns() float32`

GetFailedRuns returns the FailedRuns field if non-nil, zero value otherwise.

### GetFailedRunsOk

`func (o *InlineResponse20010Trends) GetFailedRunsOk() (*float32, bool)`

GetFailedRunsOk returns a tuple with the FailedRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedRuns

`func (o *InlineResponse20010Trends) SetFailedRuns(v float32)`

SetFailedRuns sets FailedRuns field to given value.


### GetSuccessRate

`func (o *InlineResponse20010Trends) GetSuccessRate() float32`

GetSuccessRate returns the SuccessRate field if non-nil, zero value otherwise.

### GetSuccessRateOk

`func (o *InlineResponse20010Trends) GetSuccessRateOk() (*float32, bool)`

GetSuccessRateOk returns a tuple with the SuccessRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRate

`func (o *InlineResponse20010Trends) SetSuccessRate(v float32)`

SetSuccessRate sets SuccessRate field to given value.


### GetP95DurationSecs

`func (o *InlineResponse20010Trends) GetP95DurationSecs() float32`

GetP95DurationSecs returns the P95DurationSecs field if non-nil, zero value otherwise.

### GetP95DurationSecsOk

`func (o *InlineResponse20010Trends) GetP95DurationSecsOk() (*float32, bool)`

GetP95DurationSecsOk returns a tuple with the P95DurationSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP95DurationSecs

`func (o *InlineResponse20010Trends) SetP95DurationSecs(v float32)`

SetP95DurationSecs sets P95DurationSecs field to given value.


### GetMedianDurationSecs

`func (o *InlineResponse20010Trends) GetMedianDurationSecs() float32`

GetMedianDurationSecs returns the MedianDurationSecs field if non-nil, zero value otherwise.

### GetMedianDurationSecsOk

`func (o *InlineResponse20010Trends) GetMedianDurationSecsOk() (*float32, bool)`

GetMedianDurationSecsOk returns a tuple with the MedianDurationSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianDurationSecs

`func (o *InlineResponse20010Trends) SetMedianDurationSecs(v float32)`

SetMedianDurationSecs sets MedianDurationSecs field to given value.


### GetTotalCreditsUsed

`func (o *InlineResponse20010Trends) GetTotalCreditsUsed() float32`

GetTotalCreditsUsed returns the TotalCreditsUsed field if non-nil, zero value otherwise.

### GetTotalCreditsUsedOk

`func (o *InlineResponse20010Trends) GetTotalCreditsUsedOk() (*float32, bool)`

GetTotalCreditsUsedOk returns a tuple with the TotalCreditsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCreditsUsed

`func (o *InlineResponse20010Trends) SetTotalCreditsUsed(v float32)`

SetTotalCreditsUsed sets TotalCreditsUsed field to given value.


### GetMttr

`func (o *InlineResponse20010Trends) GetMttr() float32`

GetMttr returns the Mttr field if non-nil, zero value otherwise.

### GetMttrOk

`func (o *InlineResponse20010Trends) GetMttrOk() (*float32, bool)`

GetMttrOk returns a tuple with the Mttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMttr

`func (o *InlineResponse20010Trends) SetMttr(v float32)`

SetMttr sets Mttr field to given value.


### GetThroughput

`func (o *InlineResponse20010Trends) GetThroughput() float32`

GetThroughput returns the Throughput field if non-nil, zero value otherwise.

### GetThroughputOk

`func (o *InlineResponse20010Trends) GetThroughputOk() (*float32, bool)`

GetThroughputOk returns a tuple with the Throughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughput

`func (o *InlineResponse20010Trends) SetThroughput(v float32)`

SetThroughput sets Throughput field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


