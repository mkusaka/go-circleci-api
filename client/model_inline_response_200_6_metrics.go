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

// InlineResponse2006Metrics Metrics relating to a workflow's runs.
type InlineResponse2006Metrics struct {
	// The total number of runs.
	TotalRuns int64 `json:"total_runs"`
	// The number of successful runs.
	SuccessfulRuns int64 `json:"successful_runs"`
	// The mean time to recovery (mean time between failures and their next success) in seconds.
	Mttr int64 `json:"mttr"`
	// The total credits consumed by the workflow in the aggregation window. Note that Insights is not a real time financial reporting tool and should not be used for credit reporting.
	TotalCreditsUsed int64 `json:"total_credits_used"`
	// The number of failed runs.
	FailedRuns      int64                                    `json:"failed_runs"`
	SuccessRate     float32                                  `json:"success_rate"`
	DurationMetrics InlineResponse2006MetricsDurationMetrics `json:"duration_metrics"`
	// The number of recovered workflow executions per day.
	TotalRecoveries int64 `json:"total_recoveries"`
	// The average number of runs per day.
	Throughput float32 `json:"throughput"`
}

// NewInlineResponse2006Metrics instantiates a new InlineResponse2006Metrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2006Metrics(totalRuns int64, successfulRuns int64, mttr int64, totalCreditsUsed int64, failedRuns int64, successRate float32, durationMetrics InlineResponse2006MetricsDurationMetrics, totalRecoveries int64, throughput float32) *InlineResponse2006Metrics {
	this := InlineResponse2006Metrics{}
	this.TotalRuns = totalRuns
	this.SuccessfulRuns = successfulRuns
	this.Mttr = mttr
	this.TotalCreditsUsed = totalCreditsUsed
	this.FailedRuns = failedRuns
	this.SuccessRate = successRate
	this.DurationMetrics = durationMetrics
	this.TotalRecoveries = totalRecoveries
	this.Throughput = throughput
	return &this
}

// NewInlineResponse2006MetricsWithDefaults instantiates a new InlineResponse2006Metrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2006MetricsWithDefaults() *InlineResponse2006Metrics {
	this := InlineResponse2006Metrics{}
	return &this
}

// GetTotalRuns returns the TotalRuns field value
func (o *InlineResponse2006Metrics) GetTotalRuns() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalRuns
}

// GetTotalRunsOk returns a tuple with the TotalRuns field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2006Metrics) GetTotalRunsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalRuns, true
}

// SetTotalRuns sets field value
func (o *InlineResponse2006Metrics) SetTotalRuns(v int64) {
	o.TotalRuns = v
}

// GetSuccessfulRuns returns the SuccessfulRuns field value
func (o *InlineResponse2006Metrics) GetSuccessfulRuns() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SuccessfulRuns
}

// GetSuccessfulRunsOk returns a tuple with the SuccessfulRuns field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2006Metrics) GetSuccessfulRunsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuccessfulRuns, true
}

// SetSuccessfulRuns sets field value
func (o *InlineResponse2006Metrics) SetSuccessfulRuns(v int64) {
	o.SuccessfulRuns = v
}

// GetMttr returns the Mttr field value
func (o *InlineResponse2006Metrics) GetMttr() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Mttr
}

// GetMttrOk returns a tuple with the Mttr field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2006Metrics) GetMttrOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mttr, true
}

// SetMttr sets field value
func (o *InlineResponse2006Metrics) SetMttr(v int64) {
	o.Mttr = v
}

// GetTotalCreditsUsed returns the TotalCreditsUsed field value
func (o *InlineResponse2006Metrics) GetTotalCreditsUsed() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalCreditsUsed
}

// GetTotalCreditsUsedOk returns a tuple with the TotalCreditsUsed field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2006Metrics) GetTotalCreditsUsedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCreditsUsed, true
}

// SetTotalCreditsUsed sets field value
func (o *InlineResponse2006Metrics) SetTotalCreditsUsed(v int64) {
	o.TotalCreditsUsed = v
}

// GetFailedRuns returns the FailedRuns field value
func (o *InlineResponse2006Metrics) GetFailedRuns() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FailedRuns
}

// GetFailedRunsOk returns a tuple with the FailedRuns field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2006Metrics) GetFailedRunsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailedRuns, true
}

// SetFailedRuns sets field value
func (o *InlineResponse2006Metrics) SetFailedRuns(v int64) {
	o.FailedRuns = v
}

// GetSuccessRate returns the SuccessRate field value
func (o *InlineResponse2006Metrics) GetSuccessRate() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SuccessRate
}

// GetSuccessRateOk returns a tuple with the SuccessRate field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2006Metrics) GetSuccessRateOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuccessRate, true
}

// SetSuccessRate sets field value
func (o *InlineResponse2006Metrics) SetSuccessRate(v float32) {
	o.SuccessRate = v
}

// GetDurationMetrics returns the DurationMetrics field value
func (o *InlineResponse2006Metrics) GetDurationMetrics() InlineResponse2006MetricsDurationMetrics {
	if o == nil {
		var ret InlineResponse2006MetricsDurationMetrics
		return ret
	}

	return o.DurationMetrics
}

// GetDurationMetricsOk returns a tuple with the DurationMetrics field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2006Metrics) GetDurationMetricsOk() (*InlineResponse2006MetricsDurationMetrics, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationMetrics, true
}

// SetDurationMetrics sets field value
func (o *InlineResponse2006Metrics) SetDurationMetrics(v InlineResponse2006MetricsDurationMetrics) {
	o.DurationMetrics = v
}

// GetTotalRecoveries returns the TotalRecoveries field value
func (o *InlineResponse2006Metrics) GetTotalRecoveries() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalRecoveries
}

// GetTotalRecoveriesOk returns a tuple with the TotalRecoveries field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2006Metrics) GetTotalRecoveriesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalRecoveries, true
}

// SetTotalRecoveries sets field value
func (o *InlineResponse2006Metrics) SetTotalRecoveries(v int64) {
	o.TotalRecoveries = v
}

// GetThroughput returns the Throughput field value
func (o *InlineResponse2006Metrics) GetThroughput() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Throughput
}

// GetThroughputOk returns a tuple with the Throughput field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2006Metrics) GetThroughputOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Throughput, true
}

// SetThroughput sets field value
func (o *InlineResponse2006Metrics) SetThroughput(v float32) {
	o.Throughput = v
}

func (o InlineResponse2006Metrics) MarshalJSON() ([]byte, error) {
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
		toSerialize["duration_metrics"] = o.DurationMetrics
	}
	if true {
		toSerialize["total_recoveries"] = o.TotalRecoveries
	}
	if true {
		toSerialize["throughput"] = o.Throughput
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2006Metrics struct {
	value *InlineResponse2006Metrics
	isSet bool
}

func (v NullableInlineResponse2006Metrics) Get() *InlineResponse2006Metrics {
	return v.value
}

func (v *NullableInlineResponse2006Metrics) Set(val *InlineResponse2006Metrics) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2006Metrics) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2006Metrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2006Metrics(val *InlineResponse2006Metrics) *NullableInlineResponse2006Metrics {
	return &NullableInlineResponse2006Metrics{value: val, isSet: true}
}

func (v NullableInlineResponse2006Metrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2006Metrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
