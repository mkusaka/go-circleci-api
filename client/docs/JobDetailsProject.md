# JobDetailsProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Slug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 
**Name** | **string** | The name of the project | 
**ExternalUrl** | **string** | URL to the repository hosting the project&#39;s code | 

## Methods

### NewJobDetailsProject

`func NewJobDetailsProject(id string, slug string, name string, externalUrl string, ) *JobDetailsProject`

NewJobDetailsProject instantiates a new JobDetailsProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobDetailsProjectWithDefaults

`func NewJobDetailsProjectWithDefaults() *JobDetailsProject`

NewJobDetailsProjectWithDefaults instantiates a new JobDetailsProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobDetailsProject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobDetailsProject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobDetailsProject) SetId(v string)`

SetId sets Id field to given value.


### GetSlug

`func (o *JobDetailsProject) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *JobDetailsProject) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *JobDetailsProject) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetName

`func (o *JobDetailsProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobDetailsProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobDetailsProject) SetName(v string)`

SetName sets Name field to given value.


### GetExternalUrl

`func (o *JobDetailsProject) GetExternalUrl() string`

GetExternalUrl returns the ExternalUrl field if non-nil, zero value otherwise.

### GetExternalUrlOk

`func (o *JobDetailsProject) GetExternalUrlOk() (*string, bool)`

GetExternalUrlOk returns a tuple with the ExternalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrl

`func (o *JobDetailsProject) SetExternalUrl(v string)`

SetExternalUrl sets ExternalUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


