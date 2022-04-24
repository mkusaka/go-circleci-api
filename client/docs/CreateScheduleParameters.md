# CreateScheduleParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the schedule. | 
**Timetable** | [**InlineResponse20012Timetable**](InlineResponse20012Timetable.md) |  | 
**AttributionActor** | **string** | The attribution-actor of the scheduled pipeline. | 
**Parameters** | **map[string]string** | Pipeline parameters represented as key-value pairs. Must contain branch or tag. | 
**Description** | Pointer to **string** | Description of the schedule. | [optional] 

## Methods

### NewCreateScheduleParameters

`func NewCreateScheduleParameters(name string, timetable InlineResponse20012Timetable, attributionActor string, parameters map[string]string, ) *CreateScheduleParameters`

NewCreateScheduleParameters instantiates a new CreateScheduleParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateScheduleParametersWithDefaults

`func NewCreateScheduleParametersWithDefaults() *CreateScheduleParameters`

NewCreateScheduleParametersWithDefaults instantiates a new CreateScheduleParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateScheduleParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateScheduleParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateScheduleParameters) SetName(v string)`

SetName sets Name field to given value.


### GetTimetable

`func (o *CreateScheduleParameters) GetTimetable() InlineResponse20012Timetable`

GetTimetable returns the Timetable field if non-nil, zero value otherwise.

### GetTimetableOk

`func (o *CreateScheduleParameters) GetTimetableOk() (*InlineResponse20012Timetable, bool)`

GetTimetableOk returns a tuple with the Timetable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimetable

`func (o *CreateScheduleParameters) SetTimetable(v InlineResponse20012Timetable)`

SetTimetable sets Timetable field to given value.


### GetAttributionActor

`func (o *CreateScheduleParameters) GetAttributionActor() string`

GetAttributionActor returns the AttributionActor field if non-nil, zero value otherwise.

### GetAttributionActorOk

`func (o *CreateScheduleParameters) GetAttributionActorOk() (*string, bool)`

GetAttributionActorOk returns a tuple with the AttributionActor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributionActor

`func (o *CreateScheduleParameters) SetAttributionActor(v string)`

SetAttributionActor sets AttributionActor field to given value.


### GetParameters

`func (o *CreateScheduleParameters) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *CreateScheduleParameters) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *CreateScheduleParameters) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.


### GetDescription

`func (o *CreateScheduleParameters) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateScheduleParameters) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateScheduleParameters) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateScheduleParameters) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


