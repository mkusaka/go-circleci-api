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

// InlineResponse20011TestCounts Test counts for a given pipeline number
type InlineResponse20011TestCounts struct {
	// The number of tests with the error status
	Error int64 `json:"error"`
	// The number of tests with the failure status
	Failure int64 `json:"failure"`
	// The number of tests with the skipped status
	Skipped int64 `json:"skipped"`
	// The number of tests with the success status
	Success int64 `json:"success"`
	// The total number of tests
	Total int64 `json:"total"`
}

// NewInlineResponse20011TestCounts instantiates a new InlineResponse20011TestCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20011TestCounts(error_ int64, failure int64, skipped int64, success int64, total int64) *InlineResponse20011TestCounts {
	this := InlineResponse20011TestCounts{}
	this.Error = error_
	this.Failure = failure
	this.Skipped = skipped
	this.Success = success
	this.Total = total
	return &this
}

// NewInlineResponse20011TestCountsWithDefaults instantiates a new InlineResponse20011TestCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20011TestCountsWithDefaults() *InlineResponse20011TestCounts {
	this := InlineResponse20011TestCounts{}
	return &this
}

// GetError returns the Error field value
func (o *InlineResponse20011TestCounts) GetError() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20011TestCounts) GetErrorOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *InlineResponse20011TestCounts) SetError(v int64) {
	o.Error = v
}

// GetFailure returns the Failure field value
func (o *InlineResponse20011TestCounts) GetFailure() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Failure
}

// GetFailureOk returns a tuple with the Failure field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20011TestCounts) GetFailureOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Failure, true
}

// SetFailure sets field value
func (o *InlineResponse20011TestCounts) SetFailure(v int64) {
	o.Failure = v
}

// GetSkipped returns the Skipped field value
func (o *InlineResponse20011TestCounts) GetSkipped() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Skipped
}

// GetSkippedOk returns a tuple with the Skipped field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20011TestCounts) GetSkippedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Skipped, true
}

// SetSkipped sets field value
func (o *InlineResponse20011TestCounts) SetSkipped(v int64) {
	o.Skipped = v
}

// GetSuccess returns the Success field value
func (o *InlineResponse20011TestCounts) GetSuccess() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20011TestCounts) GetSuccessOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *InlineResponse20011TestCounts) SetSuccess(v int64) {
	o.Success = v
}

// GetTotal returns the Total field value
func (o *InlineResponse20011TestCounts) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20011TestCounts) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *InlineResponse20011TestCounts) SetTotal(v int64) {
	o.Total = v
}

func (o InlineResponse20011TestCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["error"] = o.Error
	}
	if true {
		toSerialize["failure"] = o.Failure
	}
	if true {
		toSerialize["skipped"] = o.Skipped
	}
	if true {
		toSerialize["success"] = o.Success
	}
	if true {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20011TestCounts struct {
	value *InlineResponse20011TestCounts
	isSet bool
}

func (v NullableInlineResponse20011TestCounts) Get() *InlineResponse20011TestCounts {
	return v.value
}

func (v *NullableInlineResponse20011TestCounts) Set(val *InlineResponse20011TestCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20011TestCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20011TestCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20011TestCounts(val *InlineResponse20011TestCounts) *NullableInlineResponse20011TestCounts {
	return &NullableInlineResponse20011TestCounts{value: val, isSet: true}
}

func (v NullableInlineResponse20011TestCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20011TestCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
