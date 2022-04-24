# TestsResponseItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | The failure message associated with the test. | 
**Source** | **string** | The program that generated the test results | 
**RunTime** | **float64** | The time it took to run the test in seconds | 
**File** | **string** | The file in which the test is defined. | 
**Result** | **string** | Indication of whether the test succeeded. | 
**Name** | **string** | The name of the test. | 
**Classname** | **string** | The programmatic location of the test. | 

## Methods

### NewTestsResponseItems

`func NewTestsResponseItems(message string, source string, runTime float64, file string, result string, name string, classname string, ) *TestsResponseItems`

NewTestsResponseItems instantiates a new TestsResponseItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestsResponseItemsWithDefaults

`func NewTestsResponseItemsWithDefaults() *TestsResponseItems`

NewTestsResponseItemsWithDefaults instantiates a new TestsResponseItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *TestsResponseItems) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TestsResponseItems) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TestsResponseItems) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSource

`func (o *TestsResponseItems) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TestsResponseItems) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TestsResponseItems) SetSource(v string)`

SetSource sets Source field to given value.


### GetRunTime

`func (o *TestsResponseItems) GetRunTime() float64`

GetRunTime returns the RunTime field if non-nil, zero value otherwise.

### GetRunTimeOk

`func (o *TestsResponseItems) GetRunTimeOk() (*float64, bool)`

GetRunTimeOk returns a tuple with the RunTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunTime

`func (o *TestsResponseItems) SetRunTime(v float64)`

SetRunTime sets RunTime field to given value.


### GetFile

`func (o *TestsResponseItems) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *TestsResponseItems) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *TestsResponseItems) SetFile(v string)`

SetFile sets File field to given value.


### GetResult

`func (o *TestsResponseItems) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TestsResponseItems) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TestsResponseItems) SetResult(v string)`

SetResult sets Result field to given value.


### GetName

`func (o *TestsResponseItems) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestsResponseItems) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestsResponseItems) SetName(v string)`

SetName sets Name field to given value.


### GetClassname

`func (o *TestsResponseItems) GetClassname() string`

GetClassname returns the Classname field if non-nil, zero value otherwise.

### GetClassnameOk

`func (o *TestsResponseItems) GetClassnameOk() (*string, bool)`

GetClassnameOk returns a tuple with the Classname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassname

`func (o *TestsResponseItems) SetClassname(v string)`

SetClassname sets Classname field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


