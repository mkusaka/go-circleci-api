# InlineObject2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinuationKey** | **string** | A pipeline continuation key. | 
**Configuration** | **string** | A configuration string for the pipeline. | 
**Parameters** | Pointer to **map[string]string** | An object containing pipeline parameters and their values. | [optional] 

## Methods

### NewInlineObject2

`func NewInlineObject2(continuationKey string, configuration string, ) *InlineObject2`

NewInlineObject2 instantiates a new InlineObject2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject2WithDefaults

`func NewInlineObject2WithDefaults() *InlineObject2`

NewInlineObject2WithDefaults instantiates a new InlineObject2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinuationKey

`func (o *InlineObject2) GetContinuationKey() string`

GetContinuationKey returns the ContinuationKey field if non-nil, zero value otherwise.

### GetContinuationKeyOk

`func (o *InlineObject2) GetContinuationKeyOk() (*string, bool)`

GetContinuationKeyOk returns a tuple with the ContinuationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationKey

`func (o *InlineObject2) SetContinuationKey(v string)`

SetContinuationKey sets ContinuationKey field to given value.


### GetConfiguration

`func (o *InlineObject2) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *InlineObject2) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *InlineObject2) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.


### GetParameters

`func (o *InlineObject2) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *InlineObject2) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *InlineObject2) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *InlineObject2) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


