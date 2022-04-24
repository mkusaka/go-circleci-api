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

// Pipeline1 A pipeline response.
type Pipeline1 struct {
	// The unique ID of the pipeline.
	Id string `json:"id"`
	// A sequence of errors that have occurred within the pipeline.
	Errors []PipelineListResponseErrors `json:"errors"`
	// The project-slug for the pipeline.
	ProjectSlug string `json:"project_slug"`
	// The date and time the pipeline was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The number of the pipeline.
	Number            int64                             `json:"number"`
	TriggerParameters map[string]map[string]interface{} `json:"trigger_parameters,omitempty"`
	// The current state of the pipeline.
	State string `json:"state"`
	// The date and time the pipeline was created.
	CreatedAt time.Time                   `json:"created_at"`
	Trigger   PipelineListResponseTrigger `json:"trigger"`
	Vcs       *PipelineListResponseVcs    `json:"vcs,omitempty"`
}

// NewPipeline1 instantiates a new Pipeline1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipeline1(id string, errors []PipelineListResponseErrors, projectSlug string, number int64, state string, createdAt time.Time, trigger PipelineListResponseTrigger) *Pipeline1 {
	this := Pipeline1{}
	this.Id = id
	this.Errors = errors
	this.ProjectSlug = projectSlug
	this.Number = number
	this.State = state
	this.CreatedAt = createdAt
	this.Trigger = trigger
	return &this
}

// NewPipeline1WithDefaults instantiates a new Pipeline1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipeline1WithDefaults() *Pipeline1 {
	this := Pipeline1{}
	return &this
}

// GetId returns the Id field value
func (o *Pipeline1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Pipeline1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Pipeline1) SetId(v string) {
	o.Id = v
}

// GetErrors returns the Errors field value
func (o *Pipeline1) GetErrors() []PipelineListResponseErrors {
	if o == nil {
		var ret []PipelineListResponseErrors
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *Pipeline1) GetErrorsOk() ([]PipelineListResponseErrors, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *Pipeline1) SetErrors(v []PipelineListResponseErrors) {
	o.Errors = v
}

// GetProjectSlug returns the ProjectSlug field value
func (o *Pipeline1) GetProjectSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectSlug
}

// GetProjectSlugOk returns a tuple with the ProjectSlug field value
// and a boolean to check if the value has been set.
func (o *Pipeline1) GetProjectSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectSlug, true
}

// SetProjectSlug sets field value
func (o *Pipeline1) SetProjectSlug(v string) {
	o.ProjectSlug = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Pipeline1) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pipeline1) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Pipeline1) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Pipeline1) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetNumber returns the Number field value
func (o *Pipeline1) GetNumber() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *Pipeline1) GetNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *Pipeline1) SetNumber(v int64) {
	o.Number = v
}

// GetTriggerParameters returns the TriggerParameters field value if set, zero value otherwise.
func (o *Pipeline1) GetTriggerParameters() map[string]map[string]interface{} {
	if o == nil || o.TriggerParameters == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.TriggerParameters
}

// GetTriggerParametersOk returns a tuple with the TriggerParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pipeline1) GetTriggerParametersOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.TriggerParameters == nil {
		return nil, false
	}
	return o.TriggerParameters, true
}

// HasTriggerParameters returns a boolean if a field has been set.
func (o *Pipeline1) HasTriggerParameters() bool {
	if o != nil && o.TriggerParameters != nil {
		return true
	}

	return false
}

// SetTriggerParameters gets a reference to the given map[string]map[string]interface{} and assigns it to the TriggerParameters field.
func (o *Pipeline1) SetTriggerParameters(v map[string]map[string]interface{}) {
	o.TriggerParameters = v
}

// GetState returns the State field value
func (o *Pipeline1) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *Pipeline1) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *Pipeline1) SetState(v string) {
	o.State = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Pipeline1) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Pipeline1) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Pipeline1) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetTrigger returns the Trigger field value
func (o *Pipeline1) GetTrigger() PipelineListResponseTrigger {
	if o == nil {
		var ret PipelineListResponseTrigger
		return ret
	}

	return o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value
// and a boolean to check if the value has been set.
func (o *Pipeline1) GetTriggerOk() (*PipelineListResponseTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Trigger, true
}

// SetTrigger sets field value
func (o *Pipeline1) SetTrigger(v PipelineListResponseTrigger) {
	o.Trigger = v
}

// GetVcs returns the Vcs field value if set, zero value otherwise.
func (o *Pipeline1) GetVcs() PipelineListResponseVcs {
	if o == nil || o.Vcs == nil {
		var ret PipelineListResponseVcs
		return ret
	}
	return *o.Vcs
}

// GetVcsOk returns a tuple with the Vcs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pipeline1) GetVcsOk() (*PipelineListResponseVcs, bool) {
	if o == nil || o.Vcs == nil {
		return nil, false
	}
	return o.Vcs, true
}

// HasVcs returns a boolean if a field has been set.
func (o *Pipeline1) HasVcs() bool {
	if o != nil && o.Vcs != nil {
		return true
	}

	return false
}

// SetVcs gets a reference to the given PipelineListResponseVcs and assigns it to the Vcs field.
func (o *Pipeline1) SetVcs(v PipelineListResponseVcs) {
	o.Vcs = &v
}

func (o Pipeline1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["errors"] = o.Errors
	}
	if true {
		toSerialize["project_slug"] = o.ProjectSlug
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["number"] = o.Number
	}
	if o.TriggerParameters != nil {
		toSerialize["trigger_parameters"] = o.TriggerParameters
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["trigger"] = o.Trigger
	}
	if o.Vcs != nil {
		toSerialize["vcs"] = o.Vcs
	}
	return json.Marshal(toSerialize)
}

type NullablePipeline1 struct {
	value *Pipeline1
	isSet bool
}

func (v NullablePipeline1) Get() *Pipeline1 {
	return v.value
}

func (v *NullablePipeline1) Set(val *Pipeline1) {
	v.value = val
	v.isSet = true
}

func (v NullablePipeline1) IsSet() bool {
	return v.isSet
}

func (v *NullablePipeline1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipeline1(val *Pipeline1) *NullablePipeline1 {
	return &NullablePipeline1{value: val, isSet: true}
}

func (v NullablePipeline1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipeline1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}