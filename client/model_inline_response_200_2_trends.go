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

// InlineResponse2002Trends Trends aggregated across a workflow or branch for a project.
type InlineResponse2002Trends struct {
	// The trend value for total credits consumed.
	TotalCreditsUsed float32 `json:"total_credits_used"`
	// The 95th percentile duration among a group of workflow runs.
	P95DurationSecs float32 `json:"p95_duration_secs"`
	// The trend value for total number of runs.
	TotalRuns float32 `json:"total_runs"`
	// The trend value for the success rate.
	SuccessRate float32 `json:"success_rate"`
}

// NewInlineResponse2002Trends instantiates a new InlineResponse2002Trends object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2002Trends(totalCreditsUsed float32, p95DurationSecs float32, totalRuns float32, successRate float32) *InlineResponse2002Trends {
	this := InlineResponse2002Trends{}
	this.TotalCreditsUsed = totalCreditsUsed
	this.P95DurationSecs = p95DurationSecs
	this.TotalRuns = totalRuns
	this.SuccessRate = successRate
	return &this
}

// NewInlineResponse2002TrendsWithDefaults instantiates a new InlineResponse2002Trends object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2002TrendsWithDefaults() *InlineResponse2002Trends {
	this := InlineResponse2002Trends{}
	return &this
}

// GetTotalCreditsUsed returns the TotalCreditsUsed field value
func (o *InlineResponse2002Trends) GetTotalCreditsUsed() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalCreditsUsed
}

// GetTotalCreditsUsedOk returns a tuple with the TotalCreditsUsed field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2002Trends) GetTotalCreditsUsedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCreditsUsed, true
}

// SetTotalCreditsUsed sets field value
func (o *InlineResponse2002Trends) SetTotalCreditsUsed(v float32) {
	o.TotalCreditsUsed = v
}

// GetP95DurationSecs returns the P95DurationSecs field value
func (o *InlineResponse2002Trends) GetP95DurationSecs() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.P95DurationSecs
}

// GetP95DurationSecsOk returns a tuple with the P95DurationSecs field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2002Trends) GetP95DurationSecsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.P95DurationSecs, true
}

// SetP95DurationSecs sets field value
func (o *InlineResponse2002Trends) SetP95DurationSecs(v float32) {
	o.P95DurationSecs = v
}

// GetTotalRuns returns the TotalRuns field value
func (o *InlineResponse2002Trends) GetTotalRuns() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalRuns
}

// GetTotalRunsOk returns a tuple with the TotalRuns field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2002Trends) GetTotalRunsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalRuns, true
}

// SetTotalRuns sets field value
func (o *InlineResponse2002Trends) SetTotalRuns(v float32) {
	o.TotalRuns = v
}

// GetSuccessRate returns the SuccessRate field value
func (o *InlineResponse2002Trends) GetSuccessRate() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SuccessRate
}

// GetSuccessRateOk returns a tuple with the SuccessRate field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2002Trends) GetSuccessRateOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuccessRate, true
}

// SetSuccessRate sets field value
func (o *InlineResponse2002Trends) SetSuccessRate(v float32) {
	o.SuccessRate = v
}

func (o InlineResponse2002Trends) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["total_credits_used"] = o.TotalCreditsUsed
	}
	if true {
		toSerialize["p95_duration_secs"] = o.P95DurationSecs
	}
	if true {
		toSerialize["total_runs"] = o.TotalRuns
	}
	if true {
		toSerialize["success_rate"] = o.SuccessRate
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2002Trends struct {
	value *InlineResponse2002Trends
	isSet bool
}

func (v NullableInlineResponse2002Trends) Get() *InlineResponse2002Trends {
	return v.value
}

func (v *NullableInlineResponse2002Trends) Set(val *InlineResponse2002Trends) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2002Trends) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2002Trends) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2002Trends(val *InlineResponse2002Trends) *NullableInlineResponse2002Trends {
	return &NullableInlineResponse2002Trends{value: val, isSet: true}
}

func (v NullableInlineResponse2002Trends) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2002Trends) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
