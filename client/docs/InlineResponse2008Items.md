# InlineResponse2008Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the job. | 
**WindowStart** | **time.Time** | The start of the aggregation window for job metrics. | 
**WindowEnd** | **time.Time** | The end of the aggregation window for job metrics. | 
**Metrics** | [**InlineResponse2008Metrics**](InlineResponse2008Metrics.md) |  | 

## Methods

### NewInlineResponse2008Items

`func NewInlineResponse2008Items(name string, windowStart time.Time, windowEnd time.Time, metrics InlineResponse2008Metrics, ) *InlineResponse2008Items`

NewInlineResponse2008Items instantiates a new InlineResponse2008Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2008ItemsWithDefaults

`func NewInlineResponse2008ItemsWithDefaults() *InlineResponse2008Items`

NewInlineResponse2008ItemsWithDefaults instantiates a new InlineResponse2008Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse2008Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse2008Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse2008Items) SetName(v string)`

SetName sets Name field to given value.


### GetWindowStart

`func (o *InlineResponse2008Items) GetWindowStart() time.Time`

GetWindowStart returns the WindowStart field if non-nil, zero value otherwise.

### GetWindowStartOk

`func (o *InlineResponse2008Items) GetWindowStartOk() (*time.Time, bool)`

GetWindowStartOk returns a tuple with the WindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowStart

`func (o *InlineResponse2008Items) SetWindowStart(v time.Time)`

SetWindowStart sets WindowStart field to given value.


### GetWindowEnd

`func (o *InlineResponse2008Items) GetWindowEnd() time.Time`

GetWindowEnd returns the WindowEnd field if non-nil, zero value otherwise.

### GetWindowEndOk

`func (o *InlineResponse2008Items) GetWindowEndOk() (*time.Time, bool)`

GetWindowEndOk returns a tuple with the WindowEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowEnd

`func (o *InlineResponse2008Items) SetWindowEnd(v time.Time)`

SetWindowEnd sets WindowEnd field to given value.


### GetMetrics

`func (o *InlineResponse2008Items) GetMetrics() InlineResponse2008Metrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *InlineResponse2008Items) GetMetricsOk() (*InlineResponse2008Metrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *InlineResponse2008Items) SetMetrics(v InlineResponse2008Metrics)`

SetMetrics sets Metrics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


