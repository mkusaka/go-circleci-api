# Workflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PipelineId** | **string** | The ID of the pipeline this workflow belongs to. | 
**CanceledBy** | Pointer to **string** |  | [optional] 
**Id** | **string** | The unique ID of the workflow. | 
**Name** | **string** | The name of the workflow. | 
**ProjectSlug** | **string** | The project-slug for the pipeline this workflow belongs to. | 
**ErroredBy** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** | Tag used for the workflow | [optional] 
**Status** | **string** | The current status of the workflow. | 
**StartedBy** | **string** |  | 
**PipelineNumber** | **int64** | The number of the pipeline this workflow belongs to. | 
**CreatedAt** | **time.Time** | The date and time the workflow was created. | 
**StoppedAt** | **time.Time** | The date and time the workflow stopped. | 

## Methods

### NewWorkflow

`func NewWorkflow(pipelineId string, id string, name string, projectSlug string, status string, startedBy string, pipelineNumber int64, createdAt time.Time, stoppedAt time.Time, ) *Workflow`

NewWorkflow instantiates a new Workflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWithDefaults

`func NewWorkflowWithDefaults() *Workflow`

NewWorkflowWithDefaults instantiates a new Workflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPipelineId

`func (o *Workflow) GetPipelineId() string`

GetPipelineId returns the PipelineId field if non-nil, zero value otherwise.

### GetPipelineIdOk

`func (o *Workflow) GetPipelineIdOk() (*string, bool)`

GetPipelineIdOk returns a tuple with the PipelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineId

`func (o *Workflow) SetPipelineId(v string)`

SetPipelineId sets PipelineId field to given value.


### GetCanceledBy

`func (o *Workflow) GetCanceledBy() string`

GetCanceledBy returns the CanceledBy field if non-nil, zero value otherwise.

### GetCanceledByOk

`func (o *Workflow) GetCanceledByOk() (*string, bool)`

GetCanceledByOk returns a tuple with the CanceledBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledBy

`func (o *Workflow) SetCanceledBy(v string)`

SetCanceledBy sets CanceledBy field to given value.

### HasCanceledBy

`func (o *Workflow) HasCanceledBy() bool`

HasCanceledBy returns a boolean if a field has been set.

### GetId

`func (o *Workflow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Workflow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Workflow) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Workflow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Workflow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Workflow) SetName(v string)`

SetName sets Name field to given value.


### GetProjectSlug

`func (o *Workflow) GetProjectSlug() string`

GetProjectSlug returns the ProjectSlug field if non-nil, zero value otherwise.

### GetProjectSlugOk

`func (o *Workflow) GetProjectSlugOk() (*string, bool)`

GetProjectSlugOk returns a tuple with the ProjectSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectSlug

`func (o *Workflow) SetProjectSlug(v string)`

SetProjectSlug sets ProjectSlug field to given value.


### GetErroredBy

`func (o *Workflow) GetErroredBy() string`

GetErroredBy returns the ErroredBy field if non-nil, zero value otherwise.

### GetErroredByOk

`func (o *Workflow) GetErroredByOk() (*string, bool)`

GetErroredByOk returns a tuple with the ErroredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErroredBy

`func (o *Workflow) SetErroredBy(v string)`

SetErroredBy sets ErroredBy field to given value.

### HasErroredBy

`func (o *Workflow) HasErroredBy() bool`

HasErroredBy returns a boolean if a field has been set.

### GetTag

`func (o *Workflow) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Workflow) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Workflow) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *Workflow) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetStatus

`func (o *Workflow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Workflow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Workflow) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStartedBy

`func (o *Workflow) GetStartedBy() string`

GetStartedBy returns the StartedBy field if non-nil, zero value otherwise.

### GetStartedByOk

`func (o *Workflow) GetStartedByOk() (*string, bool)`

GetStartedByOk returns a tuple with the StartedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedBy

`func (o *Workflow) SetStartedBy(v string)`

SetStartedBy sets StartedBy field to given value.


### GetPipelineNumber

`func (o *Workflow) GetPipelineNumber() int64`

GetPipelineNumber returns the PipelineNumber field if non-nil, zero value otherwise.

### GetPipelineNumberOk

`func (o *Workflow) GetPipelineNumberOk() (*int64, bool)`

GetPipelineNumberOk returns a tuple with the PipelineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineNumber

`func (o *Workflow) SetPipelineNumber(v int64)`

SetPipelineNumber sets PipelineNumber field to given value.


### GetCreatedAt

`func (o *Workflow) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Workflow) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Workflow) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetStoppedAt

`func (o *Workflow) GetStoppedAt() time.Time`

GetStoppedAt returns the StoppedAt field if non-nil, zero value otherwise.

### GetStoppedAtOk

`func (o *Workflow) GetStoppedAtOk() (*time.Time, bool)`

GetStoppedAtOk returns a tuple with the StoppedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppedAt

`func (o *Workflow) SetStoppedAt(v time.Time)`

SetStoppedAt sets StoppedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


