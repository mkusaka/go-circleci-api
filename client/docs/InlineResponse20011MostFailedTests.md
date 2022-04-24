# InlineResponse20011MostFailedTests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**P95Duration** | **float64** | The 95th percentile duration, in seconds, among a group of test runs. | 
**TotalRuns** | **int64** | The total number of times the test was run. | 
**Classname** | **string** | The class the test belongs to. | 
**FailedRuns** | **int64** | The number of times the test failed | 
**Flaky** | **bool** | Whether the test is flaky. | 
**Source** | **string** | The source of the test. | 
**File** | **string** | The file the test belongs to. | 
**JobName** | **string** | The name of the job. | 
**TestName** | **string** | The name of the test. | 

## Methods

### NewInlineResponse20011MostFailedTests

`func NewInlineResponse20011MostFailedTests(p95Duration float64, totalRuns int64, classname string, failedRuns int64, flaky bool, source string, file string, jobName string, testName string, ) *InlineResponse20011MostFailedTests`

NewInlineResponse20011MostFailedTests instantiates a new InlineResponse20011MostFailedTests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20011MostFailedTestsWithDefaults

`func NewInlineResponse20011MostFailedTestsWithDefaults() *InlineResponse20011MostFailedTests`

NewInlineResponse20011MostFailedTestsWithDefaults instantiates a new InlineResponse20011MostFailedTests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetP95Duration

`func (o *InlineResponse20011MostFailedTests) GetP95Duration() float64`

GetP95Duration returns the P95Duration field if non-nil, zero value otherwise.

### GetP95DurationOk

`func (o *InlineResponse20011MostFailedTests) GetP95DurationOk() (*float64, bool)`

GetP95DurationOk returns a tuple with the P95Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP95Duration

`func (o *InlineResponse20011MostFailedTests) SetP95Duration(v float64)`

SetP95Duration sets P95Duration field to given value.


### GetTotalRuns

`func (o *InlineResponse20011MostFailedTests) GetTotalRuns() int64`

GetTotalRuns returns the TotalRuns field if non-nil, zero value otherwise.

### GetTotalRunsOk

`func (o *InlineResponse20011MostFailedTests) GetTotalRunsOk() (*int64, bool)`

GetTotalRunsOk returns a tuple with the TotalRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRuns

`func (o *InlineResponse20011MostFailedTests) SetTotalRuns(v int64)`

SetTotalRuns sets TotalRuns field to given value.


### GetClassname

`func (o *InlineResponse20011MostFailedTests) GetClassname() string`

GetClassname returns the Classname field if non-nil, zero value otherwise.

### GetClassnameOk

`func (o *InlineResponse20011MostFailedTests) GetClassnameOk() (*string, bool)`

GetClassnameOk returns a tuple with the Classname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassname

`func (o *InlineResponse20011MostFailedTests) SetClassname(v string)`

SetClassname sets Classname field to given value.


### GetFailedRuns

`func (o *InlineResponse20011MostFailedTests) GetFailedRuns() int64`

GetFailedRuns returns the FailedRuns field if non-nil, zero value otherwise.

### GetFailedRunsOk

`func (o *InlineResponse20011MostFailedTests) GetFailedRunsOk() (*int64, bool)`

GetFailedRunsOk returns a tuple with the FailedRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedRuns

`func (o *InlineResponse20011MostFailedTests) SetFailedRuns(v int64)`

SetFailedRuns sets FailedRuns field to given value.


### GetFlaky

`func (o *InlineResponse20011MostFailedTests) GetFlaky() bool`

GetFlaky returns the Flaky field if non-nil, zero value otherwise.

### GetFlakyOk

`func (o *InlineResponse20011MostFailedTests) GetFlakyOk() (*bool, bool)`

GetFlakyOk returns a tuple with the Flaky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlaky

`func (o *InlineResponse20011MostFailedTests) SetFlaky(v bool)`

SetFlaky sets Flaky field to given value.


### GetSource

`func (o *InlineResponse20011MostFailedTests) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InlineResponse20011MostFailedTests) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InlineResponse20011MostFailedTests) SetSource(v string)`

SetSource sets Source field to given value.


### GetFile

`func (o *InlineResponse20011MostFailedTests) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *InlineResponse20011MostFailedTests) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *InlineResponse20011MostFailedTests) SetFile(v string)`

SetFile sets File field to given value.


### GetJobName

`func (o *InlineResponse20011MostFailedTests) GetJobName() string`

GetJobName returns the JobName field if non-nil, zero value otherwise.

### GetJobNameOk

`func (o *InlineResponse20011MostFailedTests) GetJobNameOk() (*string, bool)`

GetJobNameOk returns a tuple with the JobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobName

`func (o *InlineResponse20011MostFailedTests) SetJobName(v string)`

SetJobName sets JobName field to given value.


### GetTestName

`func (o *InlineResponse20011MostFailedTests) GetTestName() string`

GetTestName returns the TestName field if non-nil, zero value otherwise.

### GetTestNameOk

`func (o *InlineResponse20011MostFailedTests) GetTestNameOk() (*string, bool)`

GetTestNameOk returns a tuple with the TestName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestName

`func (o *InlineResponse20011MostFailedTests) SetTestName(v string)`

SetTestName sets TestName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


