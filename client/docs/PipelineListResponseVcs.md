# PipelineListResponseVcs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderName** | **string** | Name of the VCS provider (e.g. GitHub, Bitbucket). | 
**TargetRepositoryUrl** | **string** | URL for the repository the trigger targets (i.e. the repository where the PR will be merged). For fork-PR pipelines, this is the URL to the parent repo. For other pipelines, the &#x60;origin_&#x60; and &#x60;target_repository_url&#x60;s will be the same. | 
**Branch** | Pointer to **string** | The branch where the pipeline ran. The HEAD commit on this branch was used for the pipeline. Note that &#x60;branch&#x60; and &#x60;tag&#x60; are mutually exclusive. To trigger a pipeline for a PR by number use &#x60;pull/&lt;number&gt;/head&#x60; for the PR ref or &#x60;pull/&lt;number&gt;/merge&#x60; for the merge ref (GitHub only). | [optional] 
**ReviewId** | Pointer to **string** | The code review id. | [optional] 
**ReviewUrl** | Pointer to **string** | The code review URL. | [optional] 
**Revision** | **string** | The code revision the pipeline ran. | 
**Tag** | Pointer to **string** | The tag used by the pipeline. The commit that this tag points to was used for the pipeline. Note that &#x60;branch&#x60; and &#x60;tag&#x60; are mutually exclusive. | [optional] 
**Commit** | Pointer to [**PipelineListResponseVcsCommit**](PipelineListResponseVcsCommit.md) |  | [optional] 
**OriginRepositoryUrl** | **string** | URL for the repository where the trigger originated. For fork-PR pipelines, this is the URL to the fork. For other pipelines the &#x60;origin_&#x60; and &#x60;target_repository_url&#x60;s will be the same. | 

## Methods

### NewPipelineListResponseVcs

`func NewPipelineListResponseVcs(providerName string, targetRepositoryUrl string, revision string, originRepositoryUrl string, ) *PipelineListResponseVcs`

NewPipelineListResponseVcs instantiates a new PipelineListResponseVcs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineListResponseVcsWithDefaults

`func NewPipelineListResponseVcsWithDefaults() *PipelineListResponseVcs`

NewPipelineListResponseVcsWithDefaults instantiates a new PipelineListResponseVcs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderName

`func (o *PipelineListResponseVcs) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *PipelineListResponseVcs) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *PipelineListResponseVcs) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetTargetRepositoryUrl

`func (o *PipelineListResponseVcs) GetTargetRepositoryUrl() string`

GetTargetRepositoryUrl returns the TargetRepositoryUrl field if non-nil, zero value otherwise.

### GetTargetRepositoryUrlOk

`func (o *PipelineListResponseVcs) GetTargetRepositoryUrlOk() (*string, bool)`

GetTargetRepositoryUrlOk returns a tuple with the TargetRepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRepositoryUrl

`func (o *PipelineListResponseVcs) SetTargetRepositoryUrl(v string)`

SetTargetRepositoryUrl sets TargetRepositoryUrl field to given value.


### GetBranch

`func (o *PipelineListResponseVcs) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *PipelineListResponseVcs) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *PipelineListResponseVcs) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *PipelineListResponseVcs) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetReviewId

`func (o *PipelineListResponseVcs) GetReviewId() string`

GetReviewId returns the ReviewId field if non-nil, zero value otherwise.

### GetReviewIdOk

`func (o *PipelineListResponseVcs) GetReviewIdOk() (*string, bool)`

GetReviewIdOk returns a tuple with the ReviewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewId

`func (o *PipelineListResponseVcs) SetReviewId(v string)`

SetReviewId sets ReviewId field to given value.

### HasReviewId

`func (o *PipelineListResponseVcs) HasReviewId() bool`

HasReviewId returns a boolean if a field has been set.

### GetReviewUrl

`func (o *PipelineListResponseVcs) GetReviewUrl() string`

GetReviewUrl returns the ReviewUrl field if non-nil, zero value otherwise.

### GetReviewUrlOk

`func (o *PipelineListResponseVcs) GetReviewUrlOk() (*string, bool)`

GetReviewUrlOk returns a tuple with the ReviewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewUrl

`func (o *PipelineListResponseVcs) SetReviewUrl(v string)`

SetReviewUrl sets ReviewUrl field to given value.

### HasReviewUrl

`func (o *PipelineListResponseVcs) HasReviewUrl() bool`

HasReviewUrl returns a boolean if a field has been set.

### GetRevision

`func (o *PipelineListResponseVcs) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *PipelineListResponseVcs) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *PipelineListResponseVcs) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetTag

`func (o *PipelineListResponseVcs) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *PipelineListResponseVcs) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *PipelineListResponseVcs) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *PipelineListResponseVcs) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetCommit

`func (o *PipelineListResponseVcs) GetCommit() PipelineListResponseVcsCommit`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *PipelineListResponseVcs) GetCommitOk() (*PipelineListResponseVcsCommit, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *PipelineListResponseVcs) SetCommit(v PipelineListResponseVcsCommit)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *PipelineListResponseVcs) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetOriginRepositoryUrl

`func (o *PipelineListResponseVcs) GetOriginRepositoryUrl() string`

GetOriginRepositoryUrl returns the OriginRepositoryUrl field if non-nil, zero value otherwise.

### GetOriginRepositoryUrlOk

`func (o *PipelineListResponseVcs) GetOriginRepositoryUrlOk() (*string, bool)`

GetOriginRepositoryUrlOk returns a tuple with the OriginRepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRepositoryUrl

`func (o *PipelineListResponseVcs) SetOriginRepositoryUrl(v string)`

SetOriginRepositoryUrl sets OriginRepositoryUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


