# InlineResponse2003Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the workflow. | 
**MinStartedAt** | **time.Time** | The start time for the earliest execution included in the metrics. | 
**MaxEndedAt** | **time.Time** | The end time of the last execution included in the metrics. | 
**Timestamp** | **time.Time** | The start of the interval for timeseries metrics. | 
**Metrics** | [**InlineResponse2003Metrics**](InlineResponse2003Metrics.md) |  | 

## Methods

### NewInlineResponse2003Items

`func NewInlineResponse2003Items(name string, minStartedAt time.Time, maxEndedAt time.Time, timestamp time.Time, metrics InlineResponse2003Metrics, ) *InlineResponse2003Items`

NewInlineResponse2003Items instantiates a new InlineResponse2003Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2003ItemsWithDefaults

`func NewInlineResponse2003ItemsWithDefaults() *InlineResponse2003Items`

NewInlineResponse2003ItemsWithDefaults instantiates a new InlineResponse2003Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse2003Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse2003Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse2003Items) SetName(v string)`

SetName sets Name field to given value.


### GetMinStartedAt

`func (o *InlineResponse2003Items) GetMinStartedAt() time.Time`

GetMinStartedAt returns the MinStartedAt field if non-nil, zero value otherwise.

### GetMinStartedAtOk

`func (o *InlineResponse2003Items) GetMinStartedAtOk() (*time.Time, bool)`

GetMinStartedAtOk returns a tuple with the MinStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinStartedAt

`func (o *InlineResponse2003Items) SetMinStartedAt(v time.Time)`

SetMinStartedAt sets MinStartedAt field to given value.


### GetMaxEndedAt

`func (o *InlineResponse2003Items) GetMaxEndedAt() time.Time`

GetMaxEndedAt returns the MaxEndedAt field if non-nil, zero value otherwise.

### GetMaxEndedAtOk

`func (o *InlineResponse2003Items) GetMaxEndedAtOk() (*time.Time, bool)`

GetMaxEndedAtOk returns a tuple with the MaxEndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEndedAt

`func (o *InlineResponse2003Items) SetMaxEndedAt(v time.Time)`

SetMaxEndedAt sets MaxEndedAt field to given value.


### GetTimestamp

`func (o *InlineResponse2003Items) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InlineResponse2003Items) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InlineResponse2003Items) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetMetrics

`func (o *InlineResponse2003Items) GetMetrics() InlineResponse2003Metrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *InlineResponse2003Items) GetMetricsOk() (*InlineResponse2003Metrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *InlineResponse2003Items) SetMetrics(v InlineResponse2003Metrics)`

SetMetrics sets Metrics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


