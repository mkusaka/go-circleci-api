# InlineResponse20012Timetable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PerHour** | **int32** | Number of times a schedule triggers per hour, value must be between 1 and 60 | 
**HoursOfDay** | **[]int32** | Hours in a day in which the schedule triggers. | 
**DaysOfWeek** | **[]string** | Days in a week in which the schedule triggers. | 

## Methods

### NewInlineResponse20012Timetable

`func NewInlineResponse20012Timetable(perHour int32, hoursOfDay []int32, daysOfWeek []string, ) *InlineResponse20012Timetable`

NewInlineResponse20012Timetable instantiates a new InlineResponse20012Timetable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20012TimetableWithDefaults

`func NewInlineResponse20012TimetableWithDefaults() *InlineResponse20012Timetable`

NewInlineResponse20012TimetableWithDefaults instantiates a new InlineResponse20012Timetable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPerHour

`func (o *InlineResponse20012Timetable) GetPerHour() int32`

GetPerHour returns the PerHour field if non-nil, zero value otherwise.

### GetPerHourOk

`func (o *InlineResponse20012Timetable) GetPerHourOk() (*int32, bool)`

GetPerHourOk returns a tuple with the PerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerHour

`func (o *InlineResponse20012Timetable) SetPerHour(v int32)`

SetPerHour sets PerHour field to given value.


### GetHoursOfDay

`func (o *InlineResponse20012Timetable) GetHoursOfDay() []int32`

GetHoursOfDay returns the HoursOfDay field if non-nil, zero value otherwise.

### GetHoursOfDayOk

`func (o *InlineResponse20012Timetable) GetHoursOfDayOk() (*[]int32, bool)`

GetHoursOfDayOk returns a tuple with the HoursOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursOfDay

`func (o *InlineResponse20012Timetable) SetHoursOfDay(v []int32)`

SetHoursOfDay sets HoursOfDay field to given value.


### GetDaysOfWeek

`func (o *InlineResponse20012Timetable) GetDaysOfWeek() []string`

GetDaysOfWeek returns the DaysOfWeek field if non-nil, zero value otherwise.

### GetDaysOfWeekOk

`func (o *InlineResponse20012Timetable) GetDaysOfWeekOk() (*[]string, bool)`

GetDaysOfWeekOk returns a tuple with the DaysOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysOfWeek

`func (o *InlineResponse20012Timetable) SetDaysOfWeek(v []string)`

SetDaysOfWeek sets DaysOfWeek field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


