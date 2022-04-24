# ProjectVcsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VcsUrl** | **string** | URL to the repository hosting the project&#39;s code | 
**Provider** | **string** | The VCS provider | 
**DefaultBranch** | **string** |  | 

## Methods

### NewProjectVcsInfo

`func NewProjectVcsInfo(vcsUrl string, provider string, defaultBranch string, ) *ProjectVcsInfo`

NewProjectVcsInfo instantiates a new ProjectVcsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectVcsInfoWithDefaults

`func NewProjectVcsInfoWithDefaults() *ProjectVcsInfo`

NewProjectVcsInfoWithDefaults instantiates a new ProjectVcsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVcsUrl

`func (o *ProjectVcsInfo) GetVcsUrl() string`

GetVcsUrl returns the VcsUrl field if non-nil, zero value otherwise.

### GetVcsUrlOk

`func (o *ProjectVcsInfo) GetVcsUrlOk() (*string, bool)`

GetVcsUrlOk returns a tuple with the VcsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsUrl

`func (o *ProjectVcsInfo) SetVcsUrl(v string)`

SetVcsUrl sets VcsUrl field to given value.


### GetProvider

`func (o *ProjectVcsInfo) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ProjectVcsInfo) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ProjectVcsInfo) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetDefaultBranch

`func (o *ProjectVcsInfo) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *ProjectVcsInfo) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *ProjectVcsInfo) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


