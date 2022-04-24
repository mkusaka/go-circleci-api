# JobDetailsExecutor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceClass** | **string** | Resource class name. | 
**Type** | Pointer to **string** | Executor type. | [optional] 

## Methods

### NewJobDetailsExecutor

`func NewJobDetailsExecutor(resourceClass string, ) *JobDetailsExecutor`

NewJobDetailsExecutor instantiates a new JobDetailsExecutor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobDetailsExecutorWithDefaults

`func NewJobDetailsExecutorWithDefaults() *JobDetailsExecutor`

NewJobDetailsExecutorWithDefaults instantiates a new JobDetailsExecutor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceClass

`func (o *JobDetailsExecutor) GetResourceClass() string`

GetResourceClass returns the ResourceClass field if non-nil, zero value otherwise.

### GetResourceClassOk

`func (o *JobDetailsExecutor) GetResourceClassOk() (*string, bool)`

GetResourceClassOk returns a tuple with the ResourceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceClass

`func (o *JobDetailsExecutor) SetResourceClass(v string)`

SetResourceClass sets ResourceClass field to given value.


### GetType

`func (o *JobDetailsExecutor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JobDetailsExecutor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JobDetailsExecutor) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *JobDetailsExecutor) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


