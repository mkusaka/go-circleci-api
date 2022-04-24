# ScheduleScheduleIdTimetable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PerHour** | Pointer to **int32** | Number of times a schedule triggers per hour, value must be between 1 and 60 | [optional] 
**HoursOfDay** | Pointer to **[]int32** | Hours in a day in which the schedule triggers. | [optional] 
**DaysOfWeek** | Pointer to **[]string** | Days in a week in which the schedule triggers. | [optional] 

## Methods

### NewScheduleScheduleIdTimetable

`func NewScheduleScheduleIdTimetable() *ScheduleScheduleIdTimetable`

NewScheduleScheduleIdTimetable instantiates a new ScheduleScheduleIdTimetable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleScheduleIdTimetableWithDefaults

`func NewScheduleScheduleIdTimetableWithDefaults() *ScheduleScheduleIdTimetable`

NewScheduleScheduleIdTimetableWithDefaults instantiates a new ScheduleScheduleIdTimetable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPerHour

`func (o *ScheduleScheduleIdTimetable) GetPerHour() int32`

GetPerHour returns the PerHour field if non-nil, zero value otherwise.

### GetPerHourOk

`func (o *ScheduleScheduleIdTimetable) GetPerHourOk() (*int32, bool)`

GetPerHourOk returns a tuple with the PerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerHour

`func (o *ScheduleScheduleIdTimetable) SetPerHour(v int32)`

SetPerHour sets PerHour field to given value.

### HasPerHour

`func (o *ScheduleScheduleIdTimetable) HasPerHour() bool`

HasPerHour returns a boolean if a field has been set.

### GetHoursOfDay

`func (o *ScheduleScheduleIdTimetable) GetHoursOfDay() []int32`

GetHoursOfDay returns the HoursOfDay field if non-nil, zero value otherwise.

### GetHoursOfDayOk

`func (o *ScheduleScheduleIdTimetable) GetHoursOfDayOk() (*[]int32, bool)`

GetHoursOfDayOk returns a tuple with the HoursOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursOfDay

`func (o *ScheduleScheduleIdTimetable) SetHoursOfDay(v []int32)`

SetHoursOfDay sets HoursOfDay field to given value.

### HasHoursOfDay

`func (o *ScheduleScheduleIdTimetable) HasHoursOfDay() bool`

HasHoursOfDay returns a boolean if a field has been set.

### GetDaysOfWeek

`func (o *ScheduleScheduleIdTimetable) GetDaysOfWeek() []string`

GetDaysOfWeek returns the DaysOfWeek field if non-nil, zero value otherwise.

### GetDaysOfWeekOk

`func (o *ScheduleScheduleIdTimetable) GetDaysOfWeekOk() (*[]string, bool)`

GetDaysOfWeekOk returns a tuple with the DaysOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysOfWeek

`func (o *ScheduleScheduleIdTimetable) SetDaysOfWeek(v []string)`

SetDaysOfWeek sets DaysOfWeek field to given value.

### HasDaysOfWeek

`func (o *ScheduleScheduleIdTimetable) HasDaysOfWeek() bool`

HasDaysOfWeek returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


