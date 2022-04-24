# JobDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebUrl** | **string** | URL of the job in CircleCI Web UI. | 
**Project** | [**JobDetailsProject**](JobDetailsProject.md) |  | 
**ParallelRuns** | [**[]JobDetailsParallelRuns**](JobDetailsParallelRuns.md) | Info about parallels runs and their status. | 
**StartedAt** | **time.Time** | The date and time the job started. | 
**LatestWorkflow** | [**JobDetailsLatestWorkflow**](JobDetailsLatestWorkflow.md) |  | 
**Name** | **string** | The name of the job. | 
**Executor** | [**JobDetailsExecutor**](JobDetailsExecutor.md) |  | 
**Parallelism** | **int64** | A number of parallel runs the job has. | 
**Status** | **interface{}** | The current status of the job. | 
**Number** | **int64** | The number of the job. | 
**Pipeline** | [**JobDetailsPipeline**](JobDetailsPipeline.md) |  | 
**Duration** | **int64** | Duration of a job in milliseconds. | 
**CreatedAt** | **time.Time** | The time when the job was created. | 
**Messages** | [**[]JobDetailsMessages**](JobDetailsMessages.md) | Messages from CircleCI execution platform. | 
**Contexts** | [**[]JobDetailsContexts**](JobDetailsContexts.md) | List of contexts used by the job. | 
**Organization** | [**JobDetailsOrganization**](JobDetailsOrganization.md) |  | 
**QueuedAt** | **time.Time** | The time when the job was placed in a queue. | 
**StoppedAt** | Pointer to **time.Time** | The time when the job stopped. | [optional] 

## Methods

### NewJobDetails

`func NewJobDetails(webUrl string, project JobDetailsProject, parallelRuns []JobDetailsParallelRuns, startedAt time.Time, latestWorkflow JobDetailsLatestWorkflow, name string, executor JobDetailsExecutor, parallelism int64, status interface{}, number int64, pipeline JobDetailsPipeline, duration int64, createdAt time.Time, messages []JobDetailsMessages, contexts []JobDetailsContexts, organization JobDetailsOrganization, queuedAt time.Time, ) *JobDetails`

NewJobDetails instantiates a new JobDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobDetailsWithDefaults

`func NewJobDetailsWithDefaults() *JobDetails`

NewJobDetailsWithDefaults instantiates a new JobDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebUrl

`func (o *JobDetails) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *JobDetails) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *JobDetails) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.


### GetProject

`func (o *JobDetails) GetProject() JobDetailsProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *JobDetails) GetProjectOk() (*JobDetailsProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *JobDetails) SetProject(v JobDetailsProject)`

SetProject sets Project field to given value.


### GetParallelRuns

`func (o *JobDetails) GetParallelRuns() []JobDetailsParallelRuns`

GetParallelRuns returns the ParallelRuns field if non-nil, zero value otherwise.

### GetParallelRunsOk

`func (o *JobDetails) GetParallelRunsOk() (*[]JobDetailsParallelRuns, bool)`

GetParallelRunsOk returns a tuple with the ParallelRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelRuns

`func (o *JobDetails) SetParallelRuns(v []JobDetailsParallelRuns)`

SetParallelRuns sets ParallelRuns field to given value.


### GetStartedAt

`func (o *JobDetails) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *JobDetails) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *JobDetails) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetLatestWorkflow

`func (o *JobDetails) GetLatestWorkflow() JobDetailsLatestWorkflow`

GetLatestWorkflow returns the LatestWorkflow field if non-nil, zero value otherwise.

### GetLatestWorkflowOk

`func (o *JobDetails) GetLatestWorkflowOk() (*JobDetailsLatestWorkflow, bool)`

GetLatestWorkflowOk returns a tuple with the LatestWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestWorkflow

`func (o *JobDetails) SetLatestWorkflow(v JobDetailsLatestWorkflow)`

SetLatestWorkflow sets LatestWorkflow field to given value.


### GetName

`func (o *JobDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobDetails) SetName(v string)`

SetName sets Name field to given value.


### GetExecutor

`func (o *JobDetails) GetExecutor() JobDetailsExecutor`

GetExecutor returns the Executor field if non-nil, zero value otherwise.

### GetExecutorOk

`func (o *JobDetails) GetExecutorOk() (*JobDetailsExecutor, bool)`

GetExecutorOk returns a tuple with the Executor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutor

`func (o *JobDetails) SetExecutor(v JobDetailsExecutor)`

SetExecutor sets Executor field to given value.


### GetParallelism

`func (o *JobDetails) GetParallelism() int64`

GetParallelism returns the Parallelism field if non-nil, zero value otherwise.

### GetParallelismOk

`func (o *JobDetails) GetParallelismOk() (*int64, bool)`

GetParallelismOk returns a tuple with the Parallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelism

`func (o *JobDetails) SetParallelism(v int64)`

SetParallelism sets Parallelism field to given value.


### GetStatus

`func (o *JobDetails) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JobDetails) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JobDetails) SetStatus(v interface{})`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *JobDetails) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *JobDetails) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetNumber

`func (o *JobDetails) GetNumber() int64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *JobDetails) GetNumberOk() (*int64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *JobDetails) SetNumber(v int64)`

SetNumber sets Number field to given value.


### GetPipeline

`func (o *JobDetails) GetPipeline() JobDetailsPipeline`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *JobDetails) GetPipelineOk() (*JobDetailsPipeline, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *JobDetails) SetPipeline(v JobDetailsPipeline)`

SetPipeline sets Pipeline field to given value.


### GetDuration

`func (o *JobDetails) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *JobDetails) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *JobDetails) SetDuration(v int64)`

SetDuration sets Duration field to given value.


### GetCreatedAt

`func (o *JobDetails) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *JobDetails) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *JobDetails) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetMessages

`func (o *JobDetails) GetMessages() []JobDetailsMessages`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *JobDetails) GetMessagesOk() (*[]JobDetailsMessages, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *JobDetails) SetMessages(v []JobDetailsMessages)`

SetMessages sets Messages field to given value.


### GetContexts

`func (o *JobDetails) GetContexts() []JobDetailsContexts`

GetContexts returns the Contexts field if non-nil, zero value otherwise.

### GetContextsOk

`func (o *JobDetails) GetContextsOk() (*[]JobDetailsContexts, bool)`

GetContextsOk returns a tuple with the Contexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContexts

`func (o *JobDetails) SetContexts(v []JobDetailsContexts)`

SetContexts sets Contexts field to given value.


### GetOrganization

`func (o *JobDetails) GetOrganization() JobDetailsOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *JobDetails) GetOrganizationOk() (*JobDetailsOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *JobDetails) SetOrganization(v JobDetailsOrganization)`

SetOrganization sets Organization field to given value.


### GetQueuedAt

`func (o *JobDetails) GetQueuedAt() time.Time`

GetQueuedAt returns the QueuedAt field if non-nil, zero value otherwise.

### GetQueuedAtOk

`func (o *JobDetails) GetQueuedAtOk() (*time.Time, bool)`

GetQueuedAtOk returns a tuple with the QueuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedAt

`func (o *JobDetails) SetQueuedAt(v time.Time)`

SetQueuedAt sets QueuedAt field to given value.


### GetStoppedAt

`func (o *JobDetails) GetStoppedAt() time.Time`

GetStoppedAt returns the StoppedAt field if non-nil, zero value otherwise.

### GetStoppedAtOk

`func (o *JobDetails) GetStoppedAtOk() (*time.Time, bool)`

GetStoppedAtOk returns a tuple with the StoppedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppedAt

`func (o *JobDetails) SetStoppedAt(v time.Time)`

SetStoppedAt sets StoppedAt field to given value.

### HasStoppedAt

`func (o *JobDetails) HasStoppedAt() bool`

HasStoppedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


