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

// InlineResponse2008MetricsDurationMetrics Metrics relating to the duration of runs for a workflow job.
type InlineResponse2008MetricsDurationMetrics struct {
	// The minimum duration, in seconds, among a group of runs.
	Min int64 `json:"min"`
	// The mean duration, in seconds, among a group of runs.
	Mean int64 `json:"mean"`
	// The median duration, in seconds, among a group of runs.
	Median int64 `json:"median"`
	// The 95th percentile duration, in seconds, among a group of runs.
	P95 int64 `json:"p95"`
	// The max duration, in seconds, among a group of runs.
	Max int64 `json:"max"`
	// The standard deviation, in seconds, among a group of runs.
	StandardDeviation float32 `json:"standard_deviation"`
}

// NewInlineResponse2008MetricsDurationMetrics instantiates a new InlineResponse2008MetricsDurationMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2008MetricsDurationMetrics(min int64, mean int64, median int64, p95 int64, max int64, standardDeviation float32) *InlineResponse2008MetricsDurationMetrics {
	this := InlineResponse2008MetricsDurationMetrics{}
	this.Min = min
	this.Mean = mean
	this.Median = median
	this.P95 = p95
	this.Max = max
	this.StandardDeviation = standardDeviation
	return &this
}

// NewInlineResponse2008MetricsDurationMetricsWithDefaults instantiates a new InlineResponse2008MetricsDurationMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2008MetricsDurationMetricsWithDefaults() *InlineResponse2008MetricsDurationMetrics {
	this := InlineResponse2008MetricsDurationMetrics{}
	return &this
}

// GetMin returns the Min field value
func (o *InlineResponse2008MetricsDurationMetrics) GetMin() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Min
}

// GetMinOk returns a tuple with the Min field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2008MetricsDurationMetrics) GetMinOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Min, true
}

// SetMin sets field value
func (o *InlineResponse2008MetricsDurationMetrics) SetMin(v int64) {
	o.Min = v
}

// GetMean returns the Mean field value
func (o *InlineResponse2008MetricsDurationMetrics) GetMean() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Mean
}

// GetMeanOk returns a tuple with the Mean field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2008MetricsDurationMetrics) GetMeanOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mean, true
}

// SetMean sets field value
func (o *InlineResponse2008MetricsDurationMetrics) SetMean(v int64) {
	o.Mean = v
}

// GetMedian returns the Median field value
func (o *InlineResponse2008MetricsDurationMetrics) GetMedian() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Median
}

// GetMedianOk returns a tuple with the Median field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2008MetricsDurationMetrics) GetMedianOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Median, true
}

// SetMedian sets field value
func (o *InlineResponse2008MetricsDurationMetrics) SetMedian(v int64) {
	o.Median = v
}

// GetP95 returns the P95 field value
func (o *InlineResponse2008MetricsDurationMetrics) GetP95() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.P95
}

// GetP95Ok returns a tuple with the P95 field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2008MetricsDurationMetrics) GetP95Ok() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.P95, true
}

// SetP95 sets field value
func (o *InlineResponse2008MetricsDurationMetrics) SetP95(v int64) {
	o.P95 = v
}

// GetMax returns the Max field value
func (o *InlineResponse2008MetricsDurationMetrics) GetMax() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Max
}

// GetMaxOk returns a tuple with the Max field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2008MetricsDurationMetrics) GetMaxOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Max, true
}

// SetMax sets field value
func (o *InlineResponse2008MetricsDurationMetrics) SetMax(v int64) {
	o.Max = v
}

// GetStandardDeviation returns the StandardDeviation field value
func (o *InlineResponse2008MetricsDurationMetrics) GetStandardDeviation() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.StandardDeviation
}

// GetStandardDeviationOk returns a tuple with the StandardDeviation field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2008MetricsDurationMetrics) GetStandardDeviationOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StandardDeviation, true
}

// SetStandardDeviation sets field value
func (o *InlineResponse2008MetricsDurationMetrics) SetStandardDeviation(v float32) {
	o.StandardDeviation = v
}

func (o InlineResponse2008MetricsDurationMetrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["min"] = o.Min
	}
	if true {
		toSerialize["mean"] = o.Mean
	}
	if true {
		toSerialize["median"] = o.Median
	}
	if true {
		toSerialize["p95"] = o.P95
	}
	if true {
		toSerialize["max"] = o.Max
	}
	if true {
		toSerialize["standard_deviation"] = o.StandardDeviation
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2008MetricsDurationMetrics struct {
	value *InlineResponse2008MetricsDurationMetrics
	isSet bool
}

func (v NullableInlineResponse2008MetricsDurationMetrics) Get() *InlineResponse2008MetricsDurationMetrics {
	return v.value
}

func (v *NullableInlineResponse2008MetricsDurationMetrics) Set(val *InlineResponse2008MetricsDurationMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2008MetricsDurationMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2008MetricsDurationMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2008MetricsDurationMetrics(val *InlineResponse2008MetricsDurationMetrics) *NullableInlineResponse2008MetricsDurationMetrics {
	return &NullableInlineResponse2008MetricsDurationMetrics{value: val, isSet: true}
}

func (v NullableInlineResponse2008MetricsDurationMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2008MetricsDurationMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
