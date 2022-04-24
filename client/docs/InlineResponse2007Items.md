# InlineResponse2007Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the workflow. | 
**Branch** | **string** | The VCS branch of a Workflow&#39;s trigger. | 
**Duration** | **int64** | The duration in seconds of a run. | 
**CreatedAt** | **time.Time** | The date and time the workflow was created. | 
**StoppedAt** | **time.Time** | The date and time the workflow stopped. | 
**CreditsUsed** | **int64** | The number of credits used during execution. Note that Insights is not a real time financial reporting tool and should not be used for credit reporting. | 
**Status** | **string** | Workflow status. | 

## Methods

### NewInlineResponse2007Items

`func NewInlineResponse2007Items(id string, branch string, duration int64, createdAt time.Time, stoppedAt time.Time, creditsUsed int64, status string, ) *InlineResponse2007Items`

NewInlineResponse2007Items instantiates a new InlineResponse2007Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2007ItemsWithDefaults

`func NewInlineResponse2007ItemsWithDefaults() *InlineResponse2007Items`

NewInlineResponse2007ItemsWithDefaults instantiates a new InlineResponse2007Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse2007Items) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse2007Items) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse2007Items) SetId(v string)`

SetId sets Id field to given value.


### GetBranch

`func (o *InlineResponse2007Items) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *InlineResponse2007Items) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *InlineResponse2007Items) SetBranch(v string)`

SetBranch sets Branch field to given value.


### GetDuration

`func (o *InlineResponse2007Items) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *InlineResponse2007Items) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *InlineResponse2007Items) SetDuration(v int64)`

SetDuration sets Duration field to given value.


### GetCreatedAt

`func (o *InlineResponse2007Items) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse2007Items) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse2007Items) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetStoppedAt

`func (o *InlineResponse2007Items) GetStoppedAt() time.Time`

GetStoppedAt returns the StoppedAt field if non-nil, zero value otherwise.

### GetStoppedAtOk

`func (o *InlineResponse2007Items) GetStoppedAtOk() (*time.Time, bool)`

GetStoppedAtOk returns a tuple with the StoppedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppedAt

`func (o *InlineResponse2007Items) SetStoppedAt(v time.Time)`

SetStoppedAt sets StoppedAt field to given value.


### GetCreditsUsed

`func (o *InlineResponse2007Items) GetCreditsUsed() int64`

GetCreditsUsed returns the CreditsUsed field if non-nil, zero value otherwise.

### GetCreditsUsedOk

`func (o *InlineResponse2007Items) GetCreditsUsedOk() (*int64, bool)`

GetCreditsUsedOk returns a tuple with the CreditsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsUsed

`func (o *InlineResponse2007Items) SetCreditsUsed(v int64)`

SetCreditsUsed sets CreditsUsed field to given value.


### GetStatus

`func (o *InlineResponse2007Items) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse2007Items) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse2007Items) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


