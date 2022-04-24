# PipelineListResponseTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of trigger. | 
**ReceivedAt** | **time.Time** | The date and time the trigger was received. | 
**Actor** | [**PipelineListResponseTriggerActor**](PipelineListResponseTriggerActor.md) |  | 

## Methods

### NewPipelineListResponseTrigger

`func NewPipelineListResponseTrigger(type_ string, receivedAt time.Time, actor PipelineListResponseTriggerActor, ) *PipelineListResponseTrigger`

NewPipelineListResponseTrigger instantiates a new PipelineListResponseTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineListResponseTriggerWithDefaults

`func NewPipelineListResponseTriggerWithDefaults() *PipelineListResponseTrigger`

NewPipelineListResponseTriggerWithDefaults instantiates a new PipelineListResponseTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PipelineListResponseTrigger) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PipelineListResponseTrigger) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PipelineListResponseTrigger) SetType(v string)`

SetType sets Type field to given value.


### GetReceivedAt

`func (o *PipelineListResponseTrigger) GetReceivedAt() time.Time`

GetReceivedAt returns the ReceivedAt field if non-nil, zero value otherwise.

### GetReceivedAtOk

`func (o *PipelineListResponseTrigger) GetReceivedAtOk() (*time.Time, bool)`

GetReceivedAtOk returns a tuple with the ReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAt

`func (o *PipelineListResponseTrigger) SetReceivedAt(v time.Time)`

SetReceivedAt sets ReceivedAt field to given value.


### GetActor

`func (o *PipelineListResponseTrigger) GetActor() PipelineListResponseTriggerActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *PipelineListResponseTrigger) GetActorOk() (*PipelineListResponseTriggerActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *PipelineListResponseTrigger) SetActor(v PipelineListResponseTriggerActor)`

SetActor sets Actor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


