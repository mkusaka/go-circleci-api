# InlineResponse20011

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageTestCount** | **int64** | The average number of tests executed per run | 
**MostFailedTests** | [**[]InlineResponse20011MostFailedTests**](InlineResponse20011MostFailedTests.md) | Metrics for the most frequently failing tests | 
**MostFailedTestsExtra** | **int64** | The number of tests with the same success rate being omitted from most_failed_tests | 
**SlowestTests** | [**[]InlineResponse20011MostFailedTests**](InlineResponse20011MostFailedTests.md) | Metrics for the slowest running tests | 
**SlowestTestsExtra** | **int64** | The number of tests with the same duration rate being omitted from slowest_tests | 
**TotalTestRuns** | **int64** | The total number of test runs | 
**TestRuns** | [**[]InlineResponse20011TestRuns**](InlineResponse20011TestRuns.md) | Test counts grouped by pipeline number and workflow id | 

## Methods

### NewInlineResponse20011

`func NewInlineResponse20011(averageTestCount int64, mostFailedTests []InlineResponse20011MostFailedTests, mostFailedTestsExtra int64, slowestTests []InlineResponse20011MostFailedTests, slowestTestsExtra int64, totalTestRuns int64, testRuns []InlineResponse20011TestRuns, ) *InlineResponse20011`

NewInlineResponse20011 instantiates a new InlineResponse20011 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20011WithDefaults

`func NewInlineResponse20011WithDefaults() *InlineResponse20011`

NewInlineResponse20011WithDefaults instantiates a new InlineResponse20011 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageTestCount

`func (o *InlineResponse20011) GetAverageTestCount() int64`

GetAverageTestCount returns the AverageTestCount field if non-nil, zero value otherwise.

### GetAverageTestCountOk

`func (o *InlineResponse20011) GetAverageTestCountOk() (*int64, bool)`

GetAverageTestCountOk returns a tuple with the AverageTestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageTestCount

`func (o *InlineResponse20011) SetAverageTestCount(v int64)`

SetAverageTestCount sets AverageTestCount field to given value.


### GetMostFailedTests

`func (o *InlineResponse20011) GetMostFailedTests() []InlineResponse20011MostFailedTests`

GetMostFailedTests returns the MostFailedTests field if non-nil, zero value otherwise.

### GetMostFailedTestsOk

`func (o *InlineResponse20011) GetMostFailedTestsOk() (*[]InlineResponse20011MostFailedTests, bool)`

GetMostFailedTestsOk returns a tuple with the MostFailedTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostFailedTests

`func (o *InlineResponse20011) SetMostFailedTests(v []InlineResponse20011MostFailedTests)`

SetMostFailedTests sets MostFailedTests field to given value.


### GetMostFailedTestsExtra

`func (o *InlineResponse20011) GetMostFailedTestsExtra() int64`

GetMostFailedTestsExtra returns the MostFailedTestsExtra field if non-nil, zero value otherwise.

### GetMostFailedTestsExtraOk

`func (o *InlineResponse20011) GetMostFailedTestsExtraOk() (*int64, bool)`

GetMostFailedTestsExtraOk returns a tuple with the MostFailedTestsExtra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostFailedTestsExtra

`func (o *InlineResponse20011) SetMostFailedTestsExtra(v int64)`

SetMostFailedTestsExtra sets MostFailedTestsExtra field to given value.


### GetSlowestTests

`func (o *InlineResponse20011) GetSlowestTests() []InlineResponse20011MostFailedTests`

GetSlowestTests returns the SlowestTests field if non-nil, zero value otherwise.

### GetSlowestTestsOk

`func (o *InlineResponse20011) GetSlowestTestsOk() (*[]InlineResponse20011MostFailedTests, bool)`

GetSlowestTestsOk returns a tuple with the SlowestTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowestTests

`func (o *InlineResponse20011) SetSlowestTests(v []InlineResponse20011MostFailedTests)`

SetSlowestTests sets SlowestTests field to given value.


### GetSlowestTestsExtra

`func (o *InlineResponse20011) GetSlowestTestsExtra() int64`

GetSlowestTestsExtra returns the SlowestTestsExtra field if non-nil, zero value otherwise.

### GetSlowestTestsExtraOk

`func (o *InlineResponse20011) GetSlowestTestsExtraOk() (*int64, bool)`

GetSlowestTestsExtraOk returns a tuple with the SlowestTestsExtra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowestTestsExtra

`func (o *InlineResponse20011) SetSlowestTestsExtra(v int64)`

SetSlowestTestsExtra sets SlowestTestsExtra field to given value.


### GetTotalTestRuns

`func (o *InlineResponse20011) GetTotalTestRuns() int64`

GetTotalTestRuns returns the TotalTestRuns field if non-nil, zero value otherwise.

### GetTotalTestRunsOk

`func (o *InlineResponse20011) GetTotalTestRunsOk() (*int64, bool)`

GetTotalTestRunsOk returns a tuple with the TotalTestRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTestRuns

`func (o *InlineResponse20011) SetTotalTestRuns(v int64)`

SetTotalTestRuns sets TotalTestRuns field to given value.


### GetTestRuns

`func (o *InlineResponse20011) GetTestRuns() []InlineResponse20011TestRuns`

GetTestRuns returns the TestRuns field if non-nil, zero value otherwise.

### GetTestRunsOk

`func (o *InlineResponse20011) GetTestRunsOk() (*[]InlineResponse20011TestRuns, bool)`

GetTestRunsOk returns a tuple with the TestRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRuns

`func (o *InlineResponse20011) SetTestRuns(v []InlineResponse20011TestRuns)`

SetTestRuns sets TestRuns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


