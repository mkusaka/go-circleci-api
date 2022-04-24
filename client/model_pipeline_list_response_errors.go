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

// PipelineListResponseErrors An error with a type and message.
type PipelineListResponseErrors struct {
	// The type of error.
	Type string `json:"type"`
	// A human-readable error message.
	Message string `json:"message"`
}

// NewPipelineListResponseErrors instantiates a new PipelineListResponseErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineListResponseErrors(type_ string, message string) *PipelineListResponseErrors {
	this := PipelineListResponseErrors{}
	this.Type = type_
	this.Message = message
	return &this
}

// NewPipelineListResponseErrorsWithDefaults instantiates a new PipelineListResponseErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineListResponseErrorsWithDefaults() *PipelineListResponseErrors {
	this := PipelineListResponseErrors{}
	return &this
}

// GetType returns the Type field value
func (o *PipelineListResponseErrors) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PipelineListResponseErrors) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PipelineListResponseErrors) SetType(v string) {
	o.Type = v
}

// GetMessage returns the Message field value
func (o *PipelineListResponseErrors) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *PipelineListResponseErrors) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *PipelineListResponseErrors) SetMessage(v string) {
	o.Message = v
}

func (o PipelineListResponseErrors) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineListResponseErrors struct {
	value *PipelineListResponseErrors
	isSet bool
}

func (v NullablePipelineListResponseErrors) Get() *PipelineListResponseErrors {
	return v.value
}

func (v *NullablePipelineListResponseErrors) Set(val *PipelineListResponseErrors) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineListResponseErrors) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineListResponseErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineListResponseErrors(val *PipelineListResponseErrors) *NullablePipelineListResponseErrors {
	return &NullablePipelineListResponseErrors{value: val, isSet: true}
}

func (v NullablePipelineListResponseErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineListResponseErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
