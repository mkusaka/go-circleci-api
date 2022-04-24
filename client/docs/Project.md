# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 
**Name** | **string** | The name of the project | 
**Id** | **string** |  | 
**OrganizationName** | **string** | The name of the organization the project belongs to | 
**OrganizationSlug** | **string** | The slug of the organization the project belongs to | 
**OrganizationId** | **string** | The id of the organization the project belongs to | 
**VcsInfo** | [**ProjectVcsInfo**](ProjectVcsInfo.md) |  | 

## Methods

### NewProject

`func NewProject(slug string, name string, id string, organizationName string, organizationSlug string, organizationId string, vcsInfo ProjectVcsInfo, ) *Project`

NewProject instantiates a new Project object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectWithDefaults

`func NewProjectWithDefaults() *Project`

NewProjectWithDefaults instantiates a new Project object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlug

`func (o *Project) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Project) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Project) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetName

`func (o *Project) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Project) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Project) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *Project) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Project) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Project) SetId(v string)`

SetId sets Id field to given value.


### GetOrganizationName

`func (o *Project) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *Project) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *Project) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetOrganizationSlug

`func (o *Project) GetOrganizationSlug() string`

GetOrganizationSlug returns the OrganizationSlug field if non-nil, zero value otherwise.

### GetOrganizationSlugOk

`func (o *Project) GetOrganizationSlugOk() (*string, bool)`

GetOrganizationSlugOk returns a tuple with the OrganizationSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationSlug

`func (o *Project) SetOrganizationSlug(v string)`

SetOrganizationSlug sets OrganizationSlug field to given value.


### GetOrganizationId

`func (o *Project) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Project) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Project) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetVcsInfo

`func (o *Project) GetVcsInfo() ProjectVcsInfo`

GetVcsInfo returns the VcsInfo field if non-nil, zero value otherwise.

### GetVcsInfoOk

`func (o *Project) GetVcsInfoOk() (*ProjectVcsInfo, bool)`

GetVcsInfoOk returns a tuple with the VcsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsInfo

`func (o *Project) SetVcsInfo(v ProjectVcsInfo)`

SetVcsInfo sets VcsInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


