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

// Schedule A schedule response
type Schedule struct {
	// The unique ID of the schedule.
	Id        string                       `json:"id"`
	Timetable InlineResponse20012Timetable `json:"timetable"`
	// The date and time the pipeline was last updated.
	UpdatedAt time.Time `json:"updated-at"`
	// Name of the schedule.
	Name string `json:"name"`
	// The date and time the pipeline was created.
	CreatedAt time.Time `json:"created-at"`
	// The project-slug for the schedule
	ProjectSlug string `json:"project-slug"`
	// Pipeline parameters represented as key-value pairs. Must contain branch or tag.
	Parameters map[string]string `json:"parameters"`
	Actor      User1             `json:"actor"`
	// Description of the schedule.
	Description string `json:"description"`
}

// NewSchedule instantiates a new Schedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedule(id string, timetable InlineResponse20012Timetable, updatedAt time.Time, name string, createdAt time.Time, projectSlug string, parameters map[string]string, actor User1, description string) *Schedule {
	this := Schedule{}
	this.Id = id
	this.Timetable = timetable
	this.UpdatedAt = updatedAt
	this.Name = name
	this.CreatedAt = createdAt
	this.ProjectSlug = projectSlug
	this.Parameters = parameters
	this.Actor = actor
	this.Description = description
	return &this
}

// NewScheduleWithDefaults instantiates a new Schedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleWithDefaults() *Schedule {
	this := Schedule{}
	return &this
}

// GetId returns the Id field value
func (o *Schedule) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Schedule) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Schedule) SetId(v string) {
	o.Id = v
}

// GetTimetable returns the Timetable field value
func (o *Schedule) GetTimetable() InlineResponse20012Timetable {
	if o == nil {
		var ret InlineResponse20012Timetable
		return ret
	}

	return o.Timetable
}

// GetTimetableOk returns a tuple with the Timetable field value
// and a boolean to check if the value has been set.
func (o *Schedule) GetTimetableOk() (*InlineResponse20012Timetable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timetable, true
}

// SetTimetable sets field value
func (o *Schedule) SetTimetable(v InlineResponse20012Timetable) {
	o.Timetable = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Schedule) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Schedule) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Schedule) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetName returns the Name field value
func (o *Schedule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Schedule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Schedule) SetName(v string) {
	o.Name = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Schedule) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Schedule) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Schedule) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetProjectSlug returns the ProjectSlug field value
func (o *Schedule) GetProjectSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectSlug
}

// GetProjectSlugOk returns a tuple with the ProjectSlug field value
// and a boolean to check if the value has been set.
func (o *Schedule) GetProjectSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectSlug, true
}

// SetProjectSlug sets field value
func (o *Schedule) SetProjectSlug(v string) {
	o.ProjectSlug = v
}

// GetParameters returns the Parameters field value
func (o *Schedule) GetParameters() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *Schedule) GetParametersOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// SetParameters sets field value
func (o *Schedule) SetParameters(v map[string]string) {
	o.Parameters = v
}

// GetActor returns the Actor field value
func (o *Schedule) GetActor() User1 {
	if o == nil {
		var ret User1
		return ret
	}

	return o.Actor
}

// GetActorOk returns a tuple with the Actor field value
// and a boolean to check if the value has been set.
func (o *Schedule) GetActorOk() (*User1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actor, true
}

// SetActor sets field value
func (o *Schedule) SetActor(v User1) {
	o.Actor = v
}

// GetDescription returns the Description field value
func (o *Schedule) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Schedule) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Schedule) SetDescription(v string) {
	o.Description = v
}

func (o Schedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["timetable"] = o.Timetable
	}
	if true {
		toSerialize["updated-at"] = o.UpdatedAt
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["created-at"] = o.CreatedAt
	}
	if true {
		toSerialize["project-slug"] = o.ProjectSlug
	}
	if true {
		toSerialize["parameters"] = o.Parameters
	}
	if true {
		toSerialize["actor"] = o.Actor
	}
	if true {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableSchedule struct {
	value *Schedule
	isSet bool
}

func (v NullableSchedule) Get() *Schedule {
	return v.value
}

func (v *NullableSchedule) Set(val *Schedule) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedule(val *Schedule) *NullableSchedule {
	return &NullableSchedule{value: val, isSet: true}
}

func (v NullableSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
