# Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanceledBy** | Pointer to **string** | The unique ID of the user. | [optional] 
**Dependencies** | **[]string** | A sequence of the unique job IDs for the jobs that this job depends upon in the workflow. | 
**JobNumber** | Pointer to **int64** | The number of the job. | [optional] 
**Id** | **string** | The unique ID of the job. | 
**StartedAt** | **time.Time** | The date and time the job started. | 
**Name** | **string** | The name of the job. | 
**ApprovedBy** | Pointer to **string** | The unique ID of the user. | [optional] 
**ProjectSlug** | **string** | The project-slug for the job. | 
**Status** | **interface{}** | The current status of the job. | 
**Type** | **string** | The type of job. | 
**StoppedAt** | Pointer to **time.Time** | The time when the job stopped. | [optional] 
**ApprovalRequestId** | Pointer to **string** | The unique ID of the job. | [optional] 

## Methods

### NewJob

`func NewJob(dependencies []string, id string, startedAt time.Time, name string, projectSlug string, status interface{}, type_ string, ) *Job`

NewJob instantiates a new Job object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobWithDefaults

`func NewJobWithDefaults() *Job`

NewJobWithDefaults instantiates a new Job object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanceledBy

`func (o *Job) GetCanceledBy() string`

GetCanceledBy returns the CanceledBy field if non-nil, zero value otherwise.

### GetCanceledByOk

`func (o *Job) GetCanceledByOk() (*string, bool)`

GetCanceledByOk returns a tuple with the CanceledBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledBy

`func (o *Job) SetCanceledBy(v string)`

SetCanceledBy sets CanceledBy field to given value.

### HasCanceledBy

`func (o *Job) HasCanceledBy() bool`

HasCanceledBy returns a boolean if a field has been set.

### GetDependencies

`func (o *Job) GetDependencies() []string`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *Job) GetDependenciesOk() (*[]string, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *Job) SetDependencies(v []string)`

SetDependencies sets Dependencies field to given value.


### GetJobNumber

`func (o *Job) GetJobNumber() int64`

GetJobNumber returns the JobNumber field if non-nil, zero value otherwise.

### GetJobNumberOk

`func (o *Job) GetJobNumberOk() (*int64, bool)`

GetJobNumberOk returns a tuple with the JobNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobNumber

`func (o *Job) SetJobNumber(v int64)`

SetJobNumber sets JobNumber field to given value.

### HasJobNumber

`func (o *Job) HasJobNumber() bool`

HasJobNumber returns a boolean if a field has been set.

### GetId

`func (o *Job) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Job) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Job) SetId(v string)`

SetId sets Id field to given value.


### GetStartedAt

`func (o *Job) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Job) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Job) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetName

`func (o *Job) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Job) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Job) SetName(v string)`

SetName sets Name field to given value.


### GetApprovedBy

`func (o *Job) GetApprovedBy() string`

GetApprovedBy returns the ApprovedBy field if non-nil, zero value otherwise.

### GetApprovedByOk

`func (o *Job) GetApprovedByOk() (*string, bool)`

GetApprovedByOk returns a tuple with the ApprovedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedBy

`func (o *Job) SetApprovedBy(v string)`

SetApprovedBy sets ApprovedBy field to given value.

### HasApprovedBy

`func (o *Job) HasApprovedBy() bool`

HasApprovedBy returns a boolean if a field has been set.

### GetProjectSlug

`func (o *Job) GetProjectSlug() string`

GetProjectSlug returns the ProjectSlug field if non-nil, zero value otherwise.

### GetProjectSlugOk

`func (o *Job) GetProjectSlugOk() (*string, bool)`

GetProjectSlugOk returns a tuple with the ProjectSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectSlug

`func (o *Job) SetProjectSlug(v string)`

SetProjectSlug sets ProjectSlug field to given value.


### GetStatus

`func (o *Job) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Job) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Job) SetStatus(v interface{})`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *Job) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Job) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetType

`func (o *Job) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Job) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Job) SetType(v string)`

SetType sets Type field to given value.


### GetStoppedAt

`func (o *Job) GetStoppedAt() time.Time`

GetStoppedAt returns the StoppedAt field if non-nil, zero value otherwise.

### GetStoppedAtOk

`func (o *Job) GetStoppedAtOk() (*time.Time, bool)`

GetStoppedAtOk returns a tuple with the StoppedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppedAt

`func (o *Job) SetStoppedAt(v time.Time)`

SetStoppedAt sets StoppedAt field to given value.

### HasStoppedAt

`func (o *Job) HasStoppedAt() bool`

HasStoppedAt returns a boolean if a field has been set.

### GetApprovalRequestId

`func (o *Job) GetApprovalRequestId() string`

GetApprovalRequestId returns the ApprovalRequestId field if non-nil, zero value otherwise.

### GetApprovalRequestIdOk

`func (o *Job) GetApprovalRequestIdOk() (*string, bool)`

GetApprovalRequestIdOk returns a tuple with the ApprovalRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalRequestId

`func (o *Job) SetApprovalRequestId(v string)`

SetApprovalRequestId sets ApprovalRequestId field to given value.

### HasApprovalRequestId

`func (o *Job) HasApprovalRequestId() bool`

HasApprovalRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


