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

// InlineResponse2002ProjectDataTrends Metric trends aggregated across all workflows and branches for a project.
type InlineResponse2002ProjectDataTrends struct {
	// The trend value for total number of runs.
	TotalRuns float32 `json:"total_runs"`
	// Trend value for total duration.
	TotalDurationSecs float32 `json:"total_duration_secs"`
	// The trend value for total credits consumed.
	TotalCreditsUsed float32 `json:"total_credits_used"`
	// The trend value for the success rate.
	SuccessRate float32 `json:"success_rate"`
	// Trend value for the average number of runs per day.
	Throughput float32 `json:"throughput"`
}

// NewInlineResponse2002ProjectDataTrends instantiates a new InlineResponse2002ProjectDataTrends object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2002ProjectDataTrends(totalRuns float32, totalDurationSecs float32, totalCreditsUsed float32, successRate float32, throughput float32) *InlineResponse2002ProjectDataTrends {
	this := InlineResponse2002ProjectDataTrends{}
	this.TotalRuns = totalRuns
	this.TotalDurationSecs = totalDurationSecs
	this.TotalCreditsUsed = totalCreditsUsed
	this.SuccessRate = successRate
	this.Throughput = throughput
	return &this
}

// NewInlineResponse2002ProjectDataTrendsWithDefaults instantiates a new InlineResponse2002ProjectDataTrends object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2002ProjectDataTrendsWithDefaults() *InlineResponse2002ProjectDataTrends {
	this := InlineResponse2002ProjectDataTrends{}
	return &this
}

// GetTotalRuns returns the TotalRuns field value
func (o *InlineResponse2002ProjectDataTrends) GetTotalRuns() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalRuns
}

// GetTotalRunsOk returns a tuple with the TotalRuns field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2002ProjectDataTrends) GetTotalRunsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalRuns, true
}

// SetTotalRuns sets field value
func (o *InlineResponse2002ProjectDataTrends) SetTotalRuns(v float32) {
	o.TotalRuns = v
}

// GetTotalDurationSecs returns the TotalDurationSecs field value
func (o *InlineResponse2002ProjectDataTrends) GetTotalDurationSecs() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalDurationSecs
}

// GetTotalDurationSecsOk returns a tuple with the TotalDurationSecs field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2002ProjectDataTrends) GetTotalDurationSecsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalDurationSecs, true
}

// SetTotalDurationSecs sets field value
func (o *InlineResponse2002ProjectDataTrends) SetTotalDurationSecs(v float32) {
	o.TotalDurationSecs = v
}

// GetTotalCreditsUsed returns the TotalCreditsUsed field value
func (o *InlineResponse2002ProjectDataTrends) GetTotalCreditsUsed() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalCreditsUsed
}

// GetTotalCreditsUsedOk returns a tuple with the TotalCreditsUsed field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2002ProjectDataTrends) GetTotalCreditsUsedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCreditsUsed, true
}

// SetTotalCreditsUsed sets field value
func (o *InlineResponse2002ProjectDataTrends) SetTotalCreditsUsed(v float32) {
	o.TotalCreditsUsed = v
}

// GetSuccessRate returns the SuccessRate field value
func (o *InlineResponse2002ProjectDataTrends) GetSuccessRate() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SuccessRate
}

// GetSuccessRateOk returns a tuple with the SuccessRate field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2002ProjectDataTrends) GetSuccessRateOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuccessRate, true
}

// SetSuccessRate sets field value
func (o *InlineResponse2002ProjectDataTrends) SetSuccessRate(v float32) {
	o.SuccessRate = v
}

// GetThroughput returns the Throughput field value
func (o *InlineResponse2002ProjectDataTrends) GetThroughput() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Throughput
}

// GetThroughputOk returns a tuple with the Throughput field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2002ProjectDataTrends) GetThroughputOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Throughput, true
}

// SetThroughput sets field value
func (o *InlineResponse2002ProjectDataTrends) SetThroughput(v float32) {
	o.Throughput = v
}

func (o InlineResponse2002ProjectDataTrends) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["total_runs"] = o.TotalRuns
	}
	if true {
		toSerialize["total_duration_secs"] = o.TotalDurationSecs
	}
	if true {
		toSerialize["total_credits_used"] = o.TotalCreditsUsed
	}
	if true {
		toSerialize["success_rate"] = o.SuccessRate
	}
	if true {
		toSerialize["throughput"] = o.Throughput
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2002ProjectDataTrends struct {
	value *InlineResponse2002ProjectDataTrends
	isSet bool
}

func (v NullableInlineResponse2002ProjectDataTrends) Get() *InlineResponse2002ProjectDataTrends {
	return v.value
}

func (v *NullableInlineResponse2002ProjectDataTrends) Set(val *InlineResponse2002ProjectDataTrends) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2002ProjectDataTrends) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2002ProjectDataTrends) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2002ProjectDataTrends(val *InlineResponse2002ProjectDataTrends) *NullableInlineResponse2002ProjectDataTrends {
	return &NullableInlineResponse2002ProjectDataTrends{value: val, isSet: true}
}

func (v NullableInlineResponse2002ProjectDataTrends) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2002ProjectDataTrends) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
