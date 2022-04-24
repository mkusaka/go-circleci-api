# InlineResponse2005FlakyTests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeWasted** | Pointer to **int64** |  | [optional] 
**WorkflowCreatedAt** | **interface{}** | The date and time when workflow was created. | 
**WorkflowId** | **interface{}** | The ID of the workflow associated with the provided test counts | 
**Classname** | **string** | The class the test belongs to. | 
**PipelineNumber** | **int64** | The number of the pipeline. | 
**WorkflowName** | **string** | The name of the workflow. | 
**TestName** | **string** | The name of the test. | 
**JobName** | **string** | The name of the job. | 
**JobNumber** | **int64** | The number of the job. | 
**TimesFlaked** | **int64** | The number of times the test flaked. | 
**Source** | **string** | The source of the test. | 
**File** | **string** | The file the test belongs to. | 

## Methods

### NewInlineResponse2005FlakyTests

`func NewInlineResponse2005FlakyTests(workflowCreatedAt interface{}, workflowId interface{}, classname string, pipelineNumber int64, workflowName string, testName string, jobName string, jobNumber int64, timesFlaked int64, source string, file string, ) *InlineResponse2005FlakyTests`

NewInlineResponse2005FlakyTests instantiates a new InlineResponse2005FlakyTests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2005FlakyTestsWithDefaults

`func NewInlineResponse2005FlakyTestsWithDefaults() *InlineResponse2005FlakyTests`

NewInlineResponse2005FlakyTestsWithDefaults instantiates a new InlineResponse2005FlakyTests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeWasted

`func (o *InlineResponse2005FlakyTests) GetTimeWasted() int64`

GetTimeWasted returns the TimeWasted field if non-nil, zero value otherwise.

### GetTimeWastedOk

`func (o *InlineResponse2005FlakyTests) GetTimeWastedOk() (*int64, bool)`

GetTimeWastedOk returns a tuple with the TimeWasted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWasted

`func (o *InlineResponse2005FlakyTests) SetTimeWasted(v int64)`

SetTimeWasted sets TimeWasted field to given value.

### HasTimeWasted

`func (o *InlineResponse2005FlakyTests) HasTimeWasted() bool`

HasTimeWasted returns a boolean if a field has been set.

### GetWorkflowCreatedAt

`func (o *InlineResponse2005FlakyTests) GetWorkflowCreatedAt() interface{}`

GetWorkflowCreatedAt returns the WorkflowCreatedAt field if non-nil, zero value otherwise.

### GetWorkflowCreatedAtOk

`func (o *InlineResponse2005FlakyTests) GetWorkflowCreatedAtOk() (*interface{}, bool)`

GetWorkflowCreatedAtOk returns a tuple with the WorkflowCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowCreatedAt

`func (o *InlineResponse2005FlakyTests) SetWorkflowCreatedAt(v interface{})`

SetWorkflowCreatedAt sets WorkflowCreatedAt field to given value.


### SetWorkflowCreatedAtNil

`func (o *InlineResponse2005FlakyTests) SetWorkflowCreatedAtNil(b bool)`

 SetWorkflowCreatedAtNil sets the value for WorkflowCreatedAt to be an explicit nil

### UnsetWorkflowCreatedAt
`func (o *InlineResponse2005FlakyTests) UnsetWorkflowCreatedAt()`

UnsetWorkflowCreatedAt ensures that no value is present for WorkflowCreatedAt, not even an explicit nil
### GetWorkflowId

`func (o *InlineResponse2005FlakyTests) GetWorkflowId() interface{}`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *InlineResponse2005FlakyTests) GetWorkflowIdOk() (*interface{}, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *InlineResponse2005FlakyTests) SetWorkflowId(v interface{})`

SetWorkflowId sets WorkflowId field to given value.


### SetWorkflowIdNil

`func (o *InlineResponse2005FlakyTests) SetWorkflowIdNil(b bool)`

 SetWorkflowIdNil sets the value for WorkflowId to be an explicit nil

### UnsetWorkflowId
`func (o *InlineResponse2005FlakyTests) UnsetWorkflowId()`

UnsetWorkflowId ensures that no value is present for WorkflowId, not even an explicit nil
### GetClassname

`func (o *InlineResponse2005FlakyTests) GetClassname() string`

GetClassname returns the Classname field if non-nil, zero value otherwise.

### GetClassnameOk

`func (o *InlineResponse2005FlakyTests) GetClassnameOk() (*string, bool)`

GetClassnameOk returns a tuple with the Classname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassname

`func (o *InlineResponse2005FlakyTests) SetClassname(v string)`

SetClassname sets Classname field to given value.


### GetPipelineNumber

`func (o *InlineResponse2005FlakyTests) GetPipelineNumber() int64`

GetPipelineNumber returns the PipelineNumber field if non-nil, zero value otherwise.

### GetPipelineNumberOk

`func (o *InlineResponse2005FlakyTests) GetPipelineNumberOk() (*int64, bool)`

GetPipelineNumberOk returns a tuple with the PipelineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineNumber

`func (o *InlineResponse2005FlakyTests) SetPipelineNumber(v int64)`

SetPipelineNumber sets PipelineNumber field to given value.


### GetWorkflowName

`func (o *InlineResponse2005FlakyTests) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *InlineResponse2005FlakyTests) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *InlineResponse2005FlakyTests) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.


### GetTestName

`func (o *InlineResponse2005FlakyTests) GetTestName() string`

GetTestName returns the TestName field if non-nil, zero value otherwise.

### GetTestNameOk

`func (o *InlineResponse2005FlakyTests) GetTestNameOk() (*string, bool)`

GetTestNameOk returns a tuple with the TestName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestName

`func (o *InlineResponse2005FlakyTests) SetTestName(v string)`

SetTestName sets TestName field to given value.


### GetJobName

`func (o *InlineResponse2005FlakyTests) GetJobName() string`

GetJobName returns the JobName field if non-nil, zero value otherwise.

### GetJobNameOk

`func (o *InlineResponse2005FlakyTests) GetJobNameOk() (*string, bool)`

GetJobNameOk returns a tuple with the JobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobName

`func (o *InlineResponse2005FlakyTests) SetJobName(v string)`

SetJobName sets JobName field to given value.


### GetJobNumber

`func (o *InlineResponse2005FlakyTests) GetJobNumber() int64`

GetJobNumber returns the JobNumber field if non-nil, zero value otherwise.

### GetJobNumberOk

`func (o *InlineResponse2005FlakyTests) GetJobNumberOk() (*int64, bool)`

GetJobNumberOk returns a tuple with the JobNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobNumber

`func (o *InlineResponse2005FlakyTests) SetJobNumber(v int64)`

SetJobNumber sets JobNumber field to given value.


### GetTimesFlaked

`func (o *InlineResponse2005FlakyTests) GetTimesFlaked() int64`

GetTimesFlaked returns the TimesFlaked field if non-nil, zero value otherwise.

### GetTimesFlakedOk

`func (o *InlineResponse2005FlakyTests) GetTimesFlakedOk() (*int64, bool)`

GetTimesFlakedOk returns a tuple with the TimesFlaked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimesFlaked

`func (o *InlineResponse2005FlakyTests) SetTimesFlaked(v int64)`

SetTimesFlaked sets TimesFlaked field to given value.


### GetSource

`func (o *InlineResponse2005FlakyTests) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InlineResponse2005FlakyTests) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InlineResponse2005FlakyTests) SetSource(v string)`

SetSource sets Source field to given value.


### GetFile

`func (o *InlineResponse2005FlakyTests) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *InlineResponse2005FlakyTests) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *InlineResponse2005FlakyTests) SetFile(v string)`

SetFile sets File field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


