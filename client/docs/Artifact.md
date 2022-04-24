# Artifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | The artifact path. | 
**NodeIndex** | **int64** | The index of the node that stored the artifact. | 
**Url** | **string** | The URL to download the artifact contents. | 

## Methods

### NewArtifact

`func NewArtifact(path string, nodeIndex int64, url string, ) *Artifact`

NewArtifact instantiates a new Artifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactWithDefaults

`func NewArtifactWithDefaults() *Artifact`

NewArtifactWithDefaults instantiates a new Artifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *Artifact) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Artifact) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Artifact) SetPath(v string)`

SetPath sets Path field to given value.


### GetNodeIndex

`func (o *Artifact) GetNodeIndex() int64`

GetNodeIndex returns the NodeIndex field if non-nil, zero value otherwise.

### GetNodeIndexOk

`func (o *Artifact) GetNodeIndexOk() (*int64, bool)`

GetNodeIndexOk returns a tuple with the NodeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIndex

`func (o *Artifact) SetNodeIndex(v int64)`

SetNodeIndex sets NodeIndex field to given value.


### GetUrl

`func (o *Artifact) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Artifact) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Artifact) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


