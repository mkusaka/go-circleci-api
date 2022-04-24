# Schedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the schedule. | 
**Timetable** | [**InlineResponse20012Timetable**](InlineResponse20012Timetable.md) |  | 
**UpdatedAt** | **time.Time** | The date and time the pipeline was last updated. | 
**Name** | **string** | Name of the schedule. | 
**CreatedAt** | **time.Time** | The date and time the pipeline was created. | 
**ProjectSlug** | **string** | The project-slug for the schedule | 
**Parameters** | **map[string]string** | Pipeline parameters represented as key-value pairs. Must contain branch or tag. | 
**Actor** | [**User1**](User1.md) |  | 
**Description** | **string** | Description of the schedule. | 

## Methods

### NewSchedule

`func NewSchedule(id string, timetable InlineResponse20012Timetable, updatedAt time.Time, name string, createdAt time.Time, projectSlug string, parameters map[string]string, actor User1, description string, ) *Schedule`

NewSchedule instantiates a new Schedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleWithDefaults

`func NewScheduleWithDefaults() *Schedule`

NewScheduleWithDefaults instantiates a new Schedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Schedule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Schedule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Schedule) SetId(v string)`

SetId sets Id field to given value.


### GetTimetable

`func (o *Schedule) GetTimetable() InlineResponse20012Timetable`

GetTimetable returns the Timetable field if non-nil, zero value otherwise.

### GetTimetableOk

`func (o *Schedule) GetTimetableOk() (*InlineResponse20012Timetable, bool)`

GetTimetableOk returns a tuple with the Timetable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimetable

`func (o *Schedule) SetTimetable(v InlineResponse20012Timetable)`

SetTimetable sets Timetable field to given value.


### GetUpdatedAt

`func (o *Schedule) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Schedule) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Schedule) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetName

`func (o *Schedule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Schedule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Schedule) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *Schedule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Schedule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Schedule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetProjectSlug

`func (o *Schedule) GetProjectSlug() string`

GetProjectSlug returns the ProjectSlug field if non-nil, zero value otherwise.

### GetProjectSlugOk

`func (o *Schedule) GetProjectSlugOk() (*string, bool)`

GetProjectSlugOk returns a tuple with the ProjectSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectSlug

`func (o *Schedule) SetProjectSlug(v string)`

SetProjectSlug sets ProjectSlug field to given value.


### GetParameters

`func (o *Schedule) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *Schedule) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *Schedule) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.


### GetActor

`func (o *Schedule) GetActor() User1`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *Schedule) GetActorOk() (*User1, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *Schedule) SetActor(v User1)`

SetActor sets Actor field to given value.


### GetDescription

`func (o *Schedule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Schedule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Schedule) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


