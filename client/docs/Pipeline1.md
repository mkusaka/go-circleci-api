# Pipeline1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the pipeline. | 
**Errors** | [**[]PipelineListResponseErrors**](PipelineListResponseErrors.md) | A sequence of errors that have occurred within the pipeline. | 
**ProjectSlug** | **string** | The project-slug for the pipeline. | 
**UpdatedAt** | Pointer to **time.Time** | The date and time the pipeline was last updated. | [optional] 
**Number** | **int64** | The number of the pipeline. | 
**TriggerParameters** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**State** | **string** | The current state of the pipeline. | 
**CreatedAt** | **time.Time** | The date and time the pipeline was created. | 
**Trigger** | [**PipelineListResponseTrigger**](PipelineListResponseTrigger.md) |  | 
**Vcs** | Pointer to [**PipelineListResponseVcs**](PipelineListResponseVcs.md) |  | [optional] 

## Methods

### NewPipeline1

`func NewPipeline1(id string, errors []PipelineListResponseErrors, projectSlug string, number int64, state string, createdAt time.Time, trigger PipelineListResponseTrigger, ) *Pipeline1`

NewPipeline1 instantiates a new Pipeline1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipeline1WithDefaults

`func NewPipeline1WithDefaults() *Pipeline1`

NewPipeline1WithDefaults instantiates a new Pipeline1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Pipeline1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Pipeline1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Pipeline1) SetId(v string)`

SetId sets Id field to given value.


### GetErrors

`func (o *Pipeline1) GetErrors() []PipelineListResponseErrors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Pipeline1) GetErrorsOk() (*[]PipelineListResponseErrors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Pipeline1) SetErrors(v []PipelineListResponseErrors)`

SetErrors sets Errors field to given value.


### GetProjectSlug

`func (o *Pipeline1) GetProjectSlug() string`

GetProjectSlug returns the ProjectSlug field if non-nil, zero value otherwise.

### GetProjectSlugOk

`func (o *Pipeline1) GetProjectSlugOk() (*string, bool)`

GetProjectSlugOk returns a tuple with the ProjectSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectSlug

`func (o *Pipeline1) SetProjectSlug(v string)`

SetProjectSlug sets ProjectSlug field to given value.


### GetUpdatedAt

`func (o *Pipeline1) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Pipeline1) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Pipeline1) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Pipeline1) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetNumber

`func (o *Pipeline1) GetNumber() int64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Pipeline1) GetNumberOk() (*int64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Pipeline1) SetNumber(v int64)`

SetNumber sets Number field to given value.


### GetTriggerParameters

`func (o *Pipeline1) GetTriggerParameters() map[string]map[string]interface{}`

GetTriggerParameters returns the TriggerParameters field if non-nil, zero value otherwise.

### GetTriggerParametersOk

`func (o *Pipeline1) GetTriggerParametersOk() (*map[string]map[string]interface{}, bool)`

GetTriggerParametersOk returns a tuple with the TriggerParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerParameters

`func (o *Pipeline1) SetTriggerParameters(v map[string]map[string]interface{})`

SetTriggerParameters sets TriggerParameters field to given value.

### HasTriggerParameters

`func (o *Pipeline1) HasTriggerParameters() bool`

HasTriggerParameters returns a boolean if a field has been set.

### GetState

`func (o *Pipeline1) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Pipeline1) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Pipeline1) SetState(v string)`

SetState sets State field to given value.


### GetCreatedAt

`func (o *Pipeline1) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Pipeline1) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Pipeline1) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetTrigger

`func (o *Pipeline1) GetTrigger() PipelineListResponseTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *Pipeline1) GetTriggerOk() (*PipelineListResponseTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *Pipeline1) SetTrigger(v PipelineListResponseTrigger)`

SetTrigger sets Trigger field to given value.


### GetVcs

`func (o *Pipeline1) GetVcs() PipelineListResponseVcs`

GetVcs returns the Vcs field if non-nil, zero value otherwise.

### GetVcsOk

`func (o *Pipeline1) GetVcsOk() (*PipelineListResponseVcs, bool)`

GetVcsOk returns a tuple with the Vcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcs

`func (o *Pipeline1) SetVcs(v PipelineListResponseVcs)`

SetVcs sets Vcs field to given value.

### HasVcs

`func (o *Pipeline1) HasVcs() bool`

HasVcs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


