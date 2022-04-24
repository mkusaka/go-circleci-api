# Collaboration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The UUID of the organization | 
**VcsType** | **string** | The VCS provider | 
**Name** | **string** | The name of the organization | 
**AvatarUrl** | **string** | URL to the user&#39;s avatar on the VCS | 
**Slug** | **string** | The slug of the organization | 

## Methods

### NewCollaboration

`func NewCollaboration(id string, vcsType string, name string, avatarUrl string, slug string, ) *Collaboration`

NewCollaboration instantiates a new Collaboration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollaborationWithDefaults

`func NewCollaborationWithDefaults() *Collaboration`

NewCollaborationWithDefaults instantiates a new Collaboration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Collaboration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Collaboration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Collaboration) SetId(v string)`

SetId sets Id field to given value.


### GetVcsType

`func (o *Collaboration) GetVcsType() string`

GetVcsType returns the VcsType field if non-nil, zero value otherwise.

### GetVcsTypeOk

`func (o *Collaboration) GetVcsTypeOk() (*string, bool)`

GetVcsTypeOk returns a tuple with the VcsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsType

`func (o *Collaboration) SetVcsType(v string)`

SetVcsType sets VcsType field to given value.


### GetName

`func (o *Collaboration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Collaboration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Collaboration) SetName(v string)`

SetName sets Name field to given value.


### GetAvatarUrl

`func (o *Collaboration) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *Collaboration) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *Collaboration) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.


### GetSlug

`func (o *Collaboration) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Collaboration) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Collaboration) SetSlug(v string)`

SetSlug sets Slug field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


