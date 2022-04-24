# TriggerPipelineParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** | The branch where the pipeline ran. The HEAD commit on this branch was used for the pipeline. Note that &#x60;branch&#x60; and &#x60;tag&#x60; are mutually exclusive. To trigger a pipeline for a PR by number use &#x60;pull/&lt;number&gt;/head&#x60; for the PR ref or &#x60;pull/&lt;number&gt;/merge&#x60; for the merge ref (GitHub only). | [optional] 
**Tag** | Pointer to **string** | The tag used by the pipeline. The commit that this tag points to was used for the pipeline. Note that &#x60;branch&#x60; and &#x60;tag&#x60; are mutually exclusive. | [optional] 
**Parameters** | Pointer to **map[string]string** | An object containing pipeline parameters and their values. | [optional] 

## Methods

### NewTriggerPipelineParameters

`func NewTriggerPipelineParameters() *TriggerPipelineParameters`

NewTriggerPipelineParameters instantiates a new TriggerPipelineParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerPipelineParametersWithDefaults

`func NewTriggerPipelineParametersWithDefaults() *TriggerPipelineParameters`

NewTriggerPipelineParametersWithDefaults instantiates a new TriggerPipelineParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *TriggerPipelineParameters) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *TriggerPipelineParameters) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *TriggerPipelineParameters) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *TriggerPipelineParameters) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetTag

`func (o *TriggerPipelineParameters) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *TriggerPipelineParameters) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *TriggerPipelineParameters) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *TriggerPipelineParameters) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetParameters

`func (o *TriggerPipelineParameters) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *TriggerPipelineParameters) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *TriggerPipelineParameters) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *TriggerPipelineParameters) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


