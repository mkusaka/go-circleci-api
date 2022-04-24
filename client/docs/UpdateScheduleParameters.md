# UpdateScheduleParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the schedule. | [optional] 
**Name** | Pointer to **string** | Name of the schedule. | [optional] 
**Timetable** | Pointer to [**ScheduleScheduleIdTimetable**](ScheduleScheduleIdTimetable.md) |  | [optional] 
**AttributionActor** | Pointer to **string** | The attribution-actor of the scheduled pipeline. | [optional] 
**Parameters** | Pointer to **map[string]string** | Pipeline parameters represented as key-value pairs. Must contain branch or tag. | [optional] 

## Methods

### NewUpdateScheduleParameters

`func NewUpdateScheduleParameters() *UpdateScheduleParameters`

NewUpdateScheduleParameters instantiates a new UpdateScheduleParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateScheduleParametersWithDefaults

`func NewUpdateScheduleParametersWithDefaults() *UpdateScheduleParameters`

NewUpdateScheduleParametersWithDefaults instantiates a new UpdateScheduleParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateScheduleParameters) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateScheduleParameters) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateScheduleParameters) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateScheduleParameters) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *UpdateScheduleParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateScheduleParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateScheduleParameters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateScheduleParameters) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTimetable

`func (o *UpdateScheduleParameters) GetTimetable() ScheduleScheduleIdTimetable`

GetTimetable returns the Timetable field if non-nil, zero value otherwise.

### GetTimetableOk

`func (o *UpdateScheduleParameters) GetTimetableOk() (*ScheduleScheduleIdTimetable, bool)`

GetTimetableOk returns a tuple with the Timetable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimetable

`func (o *UpdateScheduleParameters) SetTimetable(v ScheduleScheduleIdTimetable)`

SetTimetable sets Timetable field to given value.

### HasTimetable

`func (o *UpdateScheduleParameters) HasTimetable() bool`

HasTimetable returns a boolean if a field has been set.

### GetAttributionActor

`func (o *UpdateScheduleParameters) GetAttributionActor() string`

GetAttributionActor returns the AttributionActor field if non-nil, zero value otherwise.

### GetAttributionActorOk

`func (o *UpdateScheduleParameters) GetAttributionActorOk() (*string, bool)`

GetAttributionActorOk returns a tuple with the AttributionActor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributionActor

`func (o *UpdateScheduleParameters) SetAttributionActor(v string)`

SetAttributionActor sets AttributionActor field to given value.

### HasAttributionActor

`func (o *UpdateScheduleParameters) HasAttributionActor() bool`

HasAttributionActor returns a boolean if a field has been set.

### GetParameters

`func (o *UpdateScheduleParameters) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *UpdateScheduleParameters) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *UpdateScheduleParameters) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *UpdateScheduleParameters) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


