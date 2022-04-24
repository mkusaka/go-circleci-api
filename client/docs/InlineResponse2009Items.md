# InlineResponse2009Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the job. | 
**StartedAt** | **time.Time** | The date and time the job started. | 
**StoppedAt** | **time.Time** | The time when the job stopped. | 
**Status** | **string** | Job status. | 
**Duration** | **int64** | The duration in seconds of a run. | 
**CreditsUsed** | **int64** | The number of credits used during execution. Note that Insights is not a real time financial reporting tool and should not be used for credit reporting. | 

## Methods

### NewInlineResponse2009Items

`func NewInlineResponse2009Items(id string, startedAt time.Time, stoppedAt time.Time, status string, duration int64, creditsUsed int64, ) *InlineResponse2009Items`

NewInlineResponse2009Items instantiates a new InlineResponse2009Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2009ItemsWithDefaults

`func NewInlineResponse2009ItemsWithDefaults() *InlineResponse2009Items`

NewInlineResponse2009ItemsWithDefaults instantiates a new InlineResponse2009Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse2009Items) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse2009Items) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse2009Items) SetId(v string)`

SetId sets Id field to given value.


### GetStartedAt

`func (o *InlineResponse2009Items) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *InlineResponse2009Items) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *InlineResponse2009Items) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetStoppedAt

`func (o *InlineResponse2009Items) GetStoppedAt() time.Time`

GetStoppedAt returns the StoppedAt field if non-nil, zero value otherwise.

### GetStoppedAtOk

`func (o *InlineResponse2009Items) GetStoppedAtOk() (*time.Time, bool)`

GetStoppedAtOk returns a tuple with the StoppedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppedAt

`func (o *InlineResponse2009Items) SetStoppedAt(v time.Time)`

SetStoppedAt sets StoppedAt field to given value.


### GetStatus

`func (o *InlineResponse2009Items) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse2009Items) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse2009Items) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDuration

`func (o *InlineResponse2009Items) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *InlineResponse2009Items) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *InlineResponse2009Items) SetDuration(v int64)`

SetDuration sets Duration field to given value.


### GetCreditsUsed

`func (o *InlineResponse2009Items) GetCreditsUsed() int64`

GetCreditsUsed returns the CreditsUsed field if non-nil, zero value otherwise.

### GetCreditsUsedOk

`func (o *InlineResponse2009Items) GetCreditsUsedOk() (*int64, bool)`

GetCreditsUsedOk returns a tuple with the CreditsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsUsed

`func (o *InlineResponse2009Items) SetCreditsUsed(v int64)`

SetCreditsUsed sets CreditsUsed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


