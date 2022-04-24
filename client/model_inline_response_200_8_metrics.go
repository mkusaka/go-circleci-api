/*
CircleCI API

This describes the resources that make up the CircleCI API v2.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2008Metrics Metrics relating to a workflow job's runs.
type InlineResponse2008Metrics struct {
	SuccessRate float32 `json:"success_rate"`
	// The total number of runs.
	TotalRuns int64 `json:"total_runs"`
	// The number of failed runs.
	FailedRuns int64 `json:"failed_runs"`
	// The number of successful runs.
	SuccessfulRuns int64 `json:"successful_runs"`
	// The average number of runs per day.
	Throughput float32 `json:"throughput"`
	// The total credits consumed by the job in the aggregation window. Note that Insights is not a real time financial reporting tool and should not be used for credit reporting.
	TotalCreditsUsed int64                                    `json:"total_credits_used"`
	DurationMetrics  InlineResponse2008MetricsDurationMetrics `json:"duration_metrics"`
}

// NewInlineResponse2008Metrics instantiates a new InlineResponse2008Metrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2008Metrics(successRate float32, totalRuns int64, failedRuns int64, successfulRuns int64, throughput float32, totalCreditsUsed int64, durationMetrics InlineResponse2008MetricsDurationMetrics) *InlineResponse2008Metrics {
	this := InlineResponse2008Metrics{}
	this.SuccessRate = successRate
	this.TotalRuns = totalRuns
	this.FailedRuns = failedRuns
	this.SuccessfulRuns = successfulRuns
	this.Throughput = throughput
	this.TotalCreditsUsed = totalCreditsUsed
	this.DurationMetrics = durationMetrics
	return &this
}

// NewInlineResponse2008MetricsWithDefaults instantiates a new InlineResponse2008Metrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2008MetricsWithDefaults() *InlineResponse2008Metrics {
	this := InlineResponse2008Metrics{}
	return &this
}

// GetSuccessRate returns the SuccessRate field value
func (o *InlineResponse2008Metrics) GetSuccessRate() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SuccessRate
}

// GetSuccessRateOk returns a tuple with the SuccessRate field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2008Metrics) GetSuccessRateOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuccessRate, true
}

// SetSuccessRate sets field value
func (o *InlineResponse2008Metrics) SetSuccessRate(v float32) {
	o.SuccessRate = v
}

// GetTotalRuns returns the TotalRuns field value
func (o *InlineResponse2008Metrics) GetTotalRuns() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalRuns
}

// GetTotalRunsOk returns a tuple with the TotalRuns field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2008Metrics) GetTotalRunsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalRuns, true
}

// SetTotalRuns sets field value
func (o *InlineResponse2008Metrics) SetTotalRuns(v int64) {
	o.TotalRuns = v
}

// GetFailedRuns returns the FailedRuns field value
func (o *InlineResponse2008Metrics) GetFailedRuns() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FailedRuns
}

// GetFailedRunsOk returns a tuple with the FailedRuns field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2008Metrics) GetFailedRunsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailedRuns, true
}

// SetFailedRuns sets field value
func (o *InlineResponse2008Metrics) SetFailedRuns(v int64) {
	o.FailedRuns = v
}

// GetSuccessfulRuns returns the SuccessfulRuns field value
func (o *InlineResponse2008Metrics) GetSuccessfulRuns() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SuccessfulRuns
}

// GetSuccessfulRunsOk returns a tuple with the SuccessfulRuns field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2008Metrics) GetSuccessfulRunsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuccessfulRuns, true
}

// SetSuccessfulRuns sets field value
func (o *InlineResponse2008Metrics) SetSuccessfulRuns(v int64) {
	o.SuccessfulRuns = v
}

// GetThroughput returns the Throughput field value
func (o *InlineResponse2008Metrics) GetThroughput() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Throughput
}

// GetThroughputOk returns a tuple with the Throughput field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2008Metrics) GetThroughputOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Throughput, true
}

// SetThroughput sets field value
func (o *InlineResponse2008Metrics) SetThroughput(v float32) {
	o.Throughput = v
}

// GetTotalCreditsUsed returns the TotalCreditsUsed field value
func (o *InlineResponse2008Metrics) GetTotalCreditsUsed() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalCreditsUsed
}

// GetTotalCreditsUsedOk returns a tuple with the TotalCreditsUsed field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2008Metrics) GetTotalCreditsUsedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCreditsUsed, true
}

// SetTotalCreditsUsed sets field value
func (o *InlineResponse2008Metrics) SetTotalCreditsUsed(v int64) {
	o.TotalCreditsUsed = v
}

// GetDurationMetrics returns the DurationMetrics field value
func (o *InlineResponse2008Metrics) GetDurationMetrics() InlineResponse2008MetricsDurationMetrics {
	if o == nil {
		var ret InlineResponse2008MetricsDurationMetrics
		return ret
	}

	return o.DurationMetrics
}

// GetDurationMetricsOk returns a tuple with the DurationMetrics field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2008Metrics) GetDurationMetricsOk() (*InlineResponse2008MetricsDurationMetrics, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationMetrics, true
}

// SetDurationMetrics sets field value
func (o *InlineResponse2008Metrics) SetDurationMetrics(v InlineResponse2008MetricsDurationMetrics) {
	o.DurationMetrics = v
}

func (o InlineResponse2008Metrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["success_rate"] = o.SuccessRate
	}
	if true {
		toSerialize["total_runs"] = o.TotalRuns
	}
	if true {
		toSerialize["failed_runs"] = o.FailedRuns
	}
	if true {
		toSerialize["successful_runs"] = o.SuccessfulRuns
	}
	if true {
		toSerialize["throughput"] = o.Throughput
	}
	if true {
		toSerialize["total_credits_used"] = o.TotalCreditsUsed
	}
	if true {
		toSerialize["duration_metrics"] = o.DurationMetrics
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2008Metrics struct {
	value *InlineResponse2008Metrics
	isSet bool
}

func (v NullableInlineResponse2008Metrics) Get() *InlineResponse2008Metrics {
	return v.value
}

func (v *NullableInlineResponse2008Metrics) Set(val *InlineResponse2008Metrics) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2008Metrics) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2008Metrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2008Metrics(val *InlineResponse2008Metrics) *NullableInlineResponse2008Metrics {
	return &NullableInlineResponse2008Metrics{value: val, isSet: true}
}

func (v NullableInlineResponse2008Metrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2008Metrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
