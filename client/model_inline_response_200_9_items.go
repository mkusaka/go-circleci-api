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

// InlineResponse2009Items struct for InlineResponse2009Items
type InlineResponse2009Items struct {
	// The unique ID of the job.
	Id string `json:"id"`
	// The date and time the job started.
	StartedAt time.Time `json:"started_at"`
	// The time when the job stopped.
	StoppedAt time.Time `json:"stopped_at"`
	// Job status.
	Status string `json:"status"`
	// The duration in seconds of a run.
	Duration int64 `json:"duration"`
	// The number of credits used during execution. Note that Insights is not a real time financial reporting tool and should not be used for credit reporting.
	CreditsUsed int64 `json:"credits_used"`
}

// NewInlineResponse2009Items instantiates a new InlineResponse2009Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2009Items(id string, startedAt time.Time, stoppedAt time.Time, status string, duration int64, creditsUsed int64) *InlineResponse2009Items {
	this := InlineResponse2009Items{}
	this.Id = id
	this.StartedAt = startedAt
	this.StoppedAt = stoppedAt
	this.Status = status
	this.Duration = duration
	this.CreditsUsed = creditsUsed
	return &this
}

// NewInlineResponse2009ItemsWithDefaults instantiates a new InlineResponse2009Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2009ItemsWithDefaults() *InlineResponse2009Items {
	this := InlineResponse2009Items{}
	return &this
}

// GetId returns the Id field value
func (o *InlineResponse2009Items) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2009Items) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InlineResponse2009Items) SetId(v string) {
	o.Id = v
}

// GetStartedAt returns the StartedAt field value
func (o *InlineResponse2009Items) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2009Items) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *InlineResponse2009Items) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetStoppedAt returns the StoppedAt field value
func (o *InlineResponse2009Items) GetStoppedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StoppedAt
}

// GetStoppedAtOk returns a tuple with the StoppedAt field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2009Items) GetStoppedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StoppedAt, true
}

// SetStoppedAt sets field value
func (o *InlineResponse2009Items) SetStoppedAt(v time.Time) {
	o.StoppedAt = v
}

// GetStatus returns the Status field value
func (o *InlineResponse2009Items) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2009Items) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InlineResponse2009Items) SetStatus(v string) {
	o.Status = v
}

// GetDuration returns the Duration field value
func (o *InlineResponse2009Items) GetDuration() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2009Items) GetDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *InlineResponse2009Items) SetDuration(v int64) {
	o.Duration = v
}

// GetCreditsUsed returns the CreditsUsed field value
func (o *InlineResponse2009Items) GetCreditsUsed() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreditsUsed
}

// GetCreditsUsedOk returns a tuple with the CreditsUsed field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2009Items) GetCreditsUsedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreditsUsed, true
}

// SetCreditsUsed sets field value
func (o *InlineResponse2009Items) SetCreditsUsed(v int64) {
	o.CreditsUsed = v
}

func (o InlineResponse2009Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["started_at"] = o.StartedAt
	}
	if true {
		toSerialize["stopped_at"] = o.StoppedAt
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["duration"] = o.Duration
	}
	if true {
		toSerialize["credits_used"] = o.CreditsUsed
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2009Items struct {
	value *InlineResponse2009Items
	isSet bool
}

func (v NullableInlineResponse2009Items) Get() *InlineResponse2009Items {
	return v.value
}

func (v *NullableInlineResponse2009Items) Set(val *InlineResponse2009Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2009Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2009Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2009Items(val *InlineResponse2009Items) *NullableInlineResponse2009Items {
	return &NullableInlineResponse2009Items{value: val, isSet: true}
}

func (v NullableInlineResponse2009Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2009Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
