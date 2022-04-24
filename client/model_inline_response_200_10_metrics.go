/*
CircleCI API

This describes the resources that make up the CircleCI API v2.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse20010Metrics Metrics aggregated acrooss a workflow for a given time window.
type InlineResponse20010Metrics struct {
	// The total number of runs.
	TotalRuns int64 `json:"total_runs"`
	// The number of successful runs.
	SuccessfulRuns int64 `json:"successful_runs"`
	// The mean time to recovery (mean time between failures and their next success) in seconds.
	Mttr int64 `json:"mttr"`
	// The total credits consumed by the workflow in the aggregation window. Note that Insights is not a real time financial reporting tool and should not be used for credit reporting.
	TotalCreditsUsed int64 `json:"total_credits_used"`
	// The number of failed runs.
	FailedRuns  int64   `json:"failed_runs"`
	SuccessRate float32 `json:"success_rate"`
	// The start of the aggregation window for workflow metrics.
	WindowStart     time.Time                                `json:"window_start"`
	DurationMetrics InlineResponse2006MetricsDurationMetrics `json:"duration_metrics"`
	// The end of the aggregation window for workflow metrics.
	WindowEnd time.Time `json:"window_end"`
	// The average number of runs per day.
	Throughput float32 `json:"throughput"`
}

// NewInlineResponse20010Metrics instantiates a new InlineResponse20010Metrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20010Metrics(totalRuns int64, successfulRuns int64, mttr int64, totalCreditsUsed int64, failedRuns int64, successRate float32, windowStart time.Time, durationMetrics InlineResponse2006MetricsDurationMetrics, windowEnd time.Time, throughput float32) *InlineResponse20010Metrics {
	this := InlineResponse20010Metrics{}
	this.TotalRuns = totalRuns
	this.SuccessfulRuns = successfulRuns
	this.Mttr = mttr
	this.TotalCreditsUsed = totalCreditsUsed
	this.FailedRuns = failedRuns
	this.SuccessRate = successRate
	this.WindowStart = windowStart
	this.DurationMetrics = durationMetrics
	this.WindowEnd = windowEnd
	this.Throughput = throughput
	return &this
}

// NewInlineResponse20010MetricsWithDefaults instantiates a new InlineResponse20010Metrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20010MetricsWithDefaults() *InlineResponse20010Metrics {
	this := InlineResponse20010Metrics{}
	return &this
}

// GetTotalRuns returns the TotalRuns field value
func (o *InlineResponse20010Metrics) GetTotalRuns() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalRuns
}

// GetTotalRunsOk returns a tuple with the TotalRuns field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20010Metrics) GetTotalRunsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalRuns, true
}

// SetTotalRuns sets field value
func (o *InlineResponse20010Metrics) SetTotalRuns(v int64) {
	o.TotalRuns = v
}

// GetSuccessfulRuns returns the SuccessfulRuns field value
func (o *InlineResponse20010Metrics) GetSuccessfulRuns() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SuccessfulRuns
}

// GetSuccessfulRunsOk returns a tuple with the SuccessfulRuns field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20010Metrics) GetSuccessfulRunsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuccessfulRuns, true
}

// SetSuccessfulRuns sets field value
func (o *InlineResponse20010Metrics) SetSuccessfulRuns(v int64) {
	o.SuccessfulRuns = v
}

// GetMttr returns the Mttr field value
func (o *InlineResponse20010Metrics) GetMttr() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Mttr
}

// GetMttrOk returns a tuple with the Mttr field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20010Metrics) GetMttrOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mttr, true
}

// SetMttr sets field value
func (o *InlineResponse20010Metrics) SetMttr(v int64) {
	o.Mttr = v
}

// GetTotalCreditsUsed returns the TotalCreditsUsed field value
func (o *InlineResponse20010Metrics) GetTotalCreditsUsed() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalCreditsUsed
}

// GetTotalCreditsUsedOk returns a tuple with the TotalCreditsUsed field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20010Metrics) GetTotalCreditsUsedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCreditsUsed, true
}

// SetTotalCreditsUsed sets field value
func (o *InlineResponse20010Metrics) SetTotalCreditsUsed(v int64) {
	o.TotalCreditsUsed = v
}

// GetFailedRuns returns the FailedRuns field value
func (o *InlineResponse20010Metrics) GetFailedRuns() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FailedRuns
}

// GetFailedRunsOk returns a tuple with the FailedRuns field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20010Metrics) GetFailedRunsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailedRuns, true
}

// SetFailedRuns sets field value
func (o *InlineResponse20010Metrics) SetFailedRuns(v int64) {
	o.FailedRuns = v
}

// GetSuccessRate returns the SuccessRate field value
func (o *InlineResponse20010Metrics) GetSuccessRate() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SuccessRate
}

// GetSuccessRateOk returns a tuple with the SuccessRate field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20010Metrics) GetSuccessRateOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuccessRate, true
}

// SetSuccessRate sets field value
func (o *InlineResponse20010Metrics) SetSuccessRate(v float32) {
	o.SuccessRate = v
}

// GetWindowStart returns the WindowStart field value
func (o *InlineResponse20010Metrics) GetWindowStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.WindowStart
}

// GetWindowStartOk returns a tuple with the WindowStart field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20010Metrics) GetWindowStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WindowStart, true
}

// SetWindowStart sets field value
func (o *InlineResponse20010Metrics) SetWindowStart(v time.Time) {
	o.WindowStart = v
}

// GetDurationMetrics returns the DurationMetrics field value
func (o *InlineResponse20010Metrics) GetDurationMetrics() InlineResponse2006MetricsDurationMetrics {
	if o == nil {
		var ret InlineResponse2006MetricsDurationMetrics
		return ret
	}

	return o.DurationMetrics
}

// GetDurationMetricsOk returns a tuple with the DurationMetrics field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20010Metrics) GetDurationMetricsOk() (*InlineResponse2006MetricsDurationMetrics, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationMetrics, true
}

// SetDurationMetrics sets field value
func (o *InlineResponse20010Metrics) SetDurationMetrics(v InlineResponse2006MetricsDurationMetrics) {
	o.DurationMetrics = v
}

// GetWindowEnd returns the WindowEnd field value
func (o *InlineResponse20010Metrics) GetWindowEnd() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.WindowEnd
}

// GetWindowEndOk returns a tuple with the WindowEnd field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20010Metrics) GetWindowEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WindowEnd, true
}

// SetWindowEnd sets field value
func (o *InlineResponse20010Metrics) SetWindowEnd(v time.Time) {
	o.WindowEnd = v
}

// GetThroughput returns the Throughput field value
func (o *InlineResponse20010Metrics) GetThroughput() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Throughput
}

// GetThroughputOk returns a tuple with the Throughput field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20010Metrics) GetThroughputOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Throughput, true
}

// SetThroughput sets field value
func (o *InlineResponse20010Metrics) SetThroughput(v float32) {
	o.Throughput = v
}

func (o InlineResponse20010Metrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["total_runs"] = o.TotalRuns
	}
	if true {
		toSerialize["successful_runs"] = o.SuccessfulRuns
	}
	if true {
		toSerialize["mttr"] = o.Mttr
	}
	if true {
		toSerialize["total_credits_used"] = o.TotalCreditsUsed
	}
	if true {
		toSerialize["failed_runs"] = o.FailedRuns
	}
	if true {
		toSerialize["success_rate"] = o.SuccessRate
	}
	if true {
		toSerialize["window_start"] = o.WindowStart
	}
	if true {
		toSerialize["duration_metrics"] = o.DurationMetrics
	}
	if true {
		toSerialize["window_end"] = o.WindowEnd
	}
	if true {
		toSerialize["throughput"] = o.Throughput
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20010Metrics struct {
	value *InlineResponse20010Metrics
	isSet bool
}

func (v NullableInlineResponse20010Metrics) Get() *InlineResponse20010Metrics {
	return v.value
}

func (v *NullableInlineResponse20010Metrics) Set(val *InlineResponse20010Metrics) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20010Metrics) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20010Metrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20010Metrics(val *InlineResponse20010Metrics) *NullableInlineResponse20010Metrics {
	return &NullableInlineResponse20010Metrics{value: val, isSet: true}
}

func (v NullableInlineResponse20010Metrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20010Metrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}