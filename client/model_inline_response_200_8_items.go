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

// InlineResponse2008Items struct for InlineResponse2008Items
type InlineResponse2008Items struct {
	// The name of the job.
	Name string `json:"name"`
	// The start of the aggregation window for job metrics.
	WindowStart time.Time `json:"window_start"`
	// The end of the aggregation window for job metrics.
	WindowEnd time.Time                 `json:"window_end"`
	Metrics   InlineResponse2008Metrics `json:"metrics"`
}

// NewInlineResponse2008Items instantiates a new InlineResponse2008Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2008Items(name string, windowStart time.Time, windowEnd time.Time, metrics InlineResponse2008Metrics) *InlineResponse2008Items {
	this := InlineResponse2008Items{}
	this.Name = name
	this.WindowStart = windowStart
	this.WindowEnd = windowEnd
	this.Metrics = metrics
	return &this
}

// NewInlineResponse2008ItemsWithDefaults instantiates a new InlineResponse2008Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2008ItemsWithDefaults() *InlineResponse2008Items {
	this := InlineResponse2008Items{}
	return &this
}

// GetName returns the Name field value
func (o *InlineResponse2008Items) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2008Items) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineResponse2008Items) SetName(v string) {
	o.Name = v
}

// GetWindowStart returns the WindowStart field value
func (o *InlineResponse2008Items) GetWindowStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.WindowStart
}

// GetWindowStartOk returns a tuple with the WindowStart field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2008Items) GetWindowStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WindowStart, true
}

// SetWindowStart sets field value
func (o *InlineResponse2008Items) SetWindowStart(v time.Time) {
	o.WindowStart = v
}

// GetWindowEnd returns the WindowEnd field value
func (o *InlineResponse2008Items) GetWindowEnd() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.WindowEnd
}

// GetWindowEndOk returns a tuple with the WindowEnd field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2008Items) GetWindowEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WindowEnd, true
}

// SetWindowEnd sets field value
func (o *InlineResponse2008Items) SetWindowEnd(v time.Time) {
	o.WindowEnd = v
}

// GetMetrics returns the Metrics field value
func (o *InlineResponse2008Items) GetMetrics() InlineResponse2008Metrics {
	if o == nil {
		var ret InlineResponse2008Metrics
		return ret
	}

	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2008Items) GetMetricsOk() (*InlineResponse2008Metrics, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// SetMetrics sets field value
func (o *InlineResponse2008Items) SetMetrics(v InlineResponse2008Metrics) {
	o.Metrics = v
}

func (o InlineResponse2008Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["window_start"] = o.WindowStart
	}
	if true {
		toSerialize["window_end"] = o.WindowEnd
	}
	if true {
		toSerialize["metrics"] = o.Metrics
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2008Items struct {
	value *InlineResponse2008Items
	isSet bool
}

func (v NullableInlineResponse2008Items) Get() *InlineResponse2008Items {
	return v.value
}

func (v *NullableInlineResponse2008Items) Set(val *InlineResponse2008Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2008Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2008Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2008Items(val *InlineResponse2008Items) *NullableInlineResponse2008Items {
	return &NullableInlineResponse2008Items{value: val, isSet: true}
}

func (v NullableInlineResponse2008Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2008Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
