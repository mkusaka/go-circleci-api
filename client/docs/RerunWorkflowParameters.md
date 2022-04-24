# RerunWorkflowParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableSsh** | Pointer to **bool** | Whether to enable SSH access for the triggering user on the newly-rerun job. Requires the jobs parameter to be used and so is mutually exclusive with the from_failed parameter. | [optional] 
**FromFailed** | Pointer to **bool** | Whether to rerun the workflow from the failed job. Mutually exclusive with the jobs parameter. | [optional] 
**Jobs** | Pointer to **[]string** | A list of job IDs to rerun. | [optional] 
**SparseTree** | Pointer to **bool** | Completes rerun using sparse trees logic, an optimization for workflows that have disconnected subgraphs. Requires jobs parameter and so is mutually exclusive with the from_failed parameter. | [optional] 

## Methods

### NewRerunWorkflowParameters

`func NewRerunWorkflowParameters() *RerunWorkflowParameters`

NewRerunWorkflowParameters instantiates a new RerunWorkflowParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRerunWorkflowParametersWithDefaults

`func NewRerunWorkflowParametersWithDefaults() *RerunWorkflowParameters`

NewRerunWorkflowParametersWithDefaults instantiates a new RerunWorkflowParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableSsh

`func (o *RerunWorkflowParameters) GetEnableSsh() bool`

GetEnableSsh returns the EnableSsh field if non-nil, zero value otherwise.

### GetEnableSshOk

`func (o *RerunWorkflowParameters) GetEnableSshOk() (*bool, bool)`

GetEnableSshOk returns a tuple with the EnableSsh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSsh

`func (o *RerunWorkflowParameters) SetEnableSsh(v bool)`

SetEnableSsh sets EnableSsh field to given value.

### HasEnableSsh

`func (o *RerunWorkflowParameters) HasEnableSsh() bool`

HasEnableSsh returns a boolean if a field has been set.

### GetFromFailed

`func (o *RerunWorkflowParameters) GetFromFailed() bool`

GetFromFailed returns the FromFailed field if non-nil, zero value otherwise.

### GetFromFailedOk

`func (o *RerunWorkflowParameters) GetFromFailedOk() (*bool, bool)`

GetFromFailedOk returns a tuple with the FromFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromFailed

`func (o *RerunWorkflowParameters) SetFromFailed(v bool)`

SetFromFailed sets FromFailed field to given value.

### HasFromFailed

`func (o *RerunWorkflowParameters) HasFromFailed() bool`

HasFromFailed returns a boolean if a field has been set.

### GetJobs

`func (o *RerunWorkflowParameters) GetJobs() []string`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *RerunWorkflowParameters) GetJobsOk() (*[]string, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *RerunWorkflowParameters) SetJobs(v []string)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *RerunWorkflowParameters) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetSparseTree

`func (o *RerunWorkflowParameters) GetSparseTree() bool`

GetSparseTree returns the SparseTree field if non-nil, zero value otherwise.

### GetSparseTreeOk

`func (o *RerunWorkflowParameters) GetSparseTreeOk() (*bool, bool)`

GetSparseTreeOk returns a tuple with the SparseTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSparseTree

`func (o *RerunWorkflowParameters) SetSparseTree(v bool)`

SetSparseTree sets SparseTree field to given value.

### HasSparseTree

`func (o *RerunWorkflowParameters) HasSparseTree() bool`

HasSparseTree returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


