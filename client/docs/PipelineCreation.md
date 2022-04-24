# PipelineCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the pipeline. | 
**State** | **string** | The current state of the pipeline. | 
**Number** | **int64** | The number of the pipeline. | 
**CreatedAt** | **time.Time** | The date and time the pipeline was created. | 

## Methods

### NewPipelineCreation

`func NewPipelineCreation(id string, state string, number int64, createdAt time.Time, ) *PipelineCreation`

NewPipelineCreation instantiates a new PipelineCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineCreationWithDefaults

`func NewPipelineCreationWithDefaults() *PipelineCreation`

NewPipelineCreationWithDefaults instantiates a new PipelineCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PipelineCreation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PipelineCreation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PipelineCreation) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *PipelineCreation) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PipelineCreation) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PipelineCreation) SetState(v string)`

SetState sets State field to given value.


### GetNumber

`func (o *PipelineCreation) GetNumber() int64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PipelineCreation) GetNumberOk() (*int64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PipelineCreation) SetNumber(v int64)`

SetNumber sets Number field to given value.


### GetCreatedAt

`func (o *PipelineCreation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PipelineCreation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PipelineCreation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


