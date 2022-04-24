# PipelineConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | The source configuration for the pipeline, before any config compilation has been performed. If there is no config, then this field will be empty. | 
**Compiled** | **string** | The compiled configuration for the pipeline, after all orb expansion has been performed. If there were errors processing the pipeline&#39;s configuration, then this field may be empty. | 
**SetupConfig** | Pointer to **string** | The setup configuration for the pipeline used for Setup Workflows. If there were errors processing the pipeline&#39;s configuration or if setup workflows are not enabled, then this field should not exist | [optional] 
**CompiledSetupConfig** | Pointer to **string** | The compiled setup configuration for the pipeline, after all orb expansion has been performed. If there were errors processing the pipeline&#39;s setup workflows, then this field may be empty. | [optional] 

## Methods

### NewPipelineConfig

`func NewPipelineConfig(source string, compiled string, ) *PipelineConfig`

NewPipelineConfig instantiates a new PipelineConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineConfigWithDefaults

`func NewPipelineConfigWithDefaults() *PipelineConfig`

NewPipelineConfigWithDefaults instantiates a new PipelineConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *PipelineConfig) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PipelineConfig) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PipelineConfig) SetSource(v string)`

SetSource sets Source field to given value.


### GetCompiled

`func (o *PipelineConfig) GetCompiled() string`

GetCompiled returns the Compiled field if non-nil, zero value otherwise.

### GetCompiledOk

`func (o *PipelineConfig) GetCompiledOk() (*string, bool)`

GetCompiledOk returns a tuple with the Compiled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompiled

`func (o *PipelineConfig) SetCompiled(v string)`

SetCompiled sets Compiled field to given value.


### GetSetupConfig

`func (o *PipelineConfig) GetSetupConfig() string`

GetSetupConfig returns the SetupConfig field if non-nil, zero value otherwise.

### GetSetupConfigOk

`func (o *PipelineConfig) GetSetupConfigOk() (*string, bool)`

GetSetupConfigOk returns a tuple with the SetupConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupConfig

`func (o *PipelineConfig) SetSetupConfig(v string)`

SetSetupConfig sets SetupConfig field to given value.

### HasSetupConfig

`func (o *PipelineConfig) HasSetupConfig() bool`

HasSetupConfig returns a boolean if a field has been set.

### GetCompiledSetupConfig

`func (o *PipelineConfig) GetCompiledSetupConfig() string`

GetCompiledSetupConfig returns the CompiledSetupConfig field if non-nil, zero value otherwise.

### GetCompiledSetupConfigOk

`func (o *PipelineConfig) GetCompiledSetupConfigOk() (*string, bool)`

GetCompiledSetupConfigOk returns a tuple with the CompiledSetupConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompiledSetupConfig

`func (o *PipelineConfig) SetCompiledSetupConfig(v string)`

SetCompiledSetupConfig sets CompiledSetupConfig field to given value.

### HasCompiledSetupConfig

`func (o *PipelineConfig) HasCompiledSetupConfig() bool`

HasCompiledSetupConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


