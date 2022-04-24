# InlineResponse2004

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgData** | [**InlineResponse2004OrgData**](InlineResponse2004OrgData.md) |  | 
**OrgProjectData** | [**[]InlineResponse2004OrgProjectData**](InlineResponse2004OrgProjectData.md) | Metrics for a single project, across all branches | 
**AllProjects** | **[]string** | A list of all the project names in the organization. | 

## Methods

### NewInlineResponse2004

`func NewInlineResponse2004(orgData InlineResponse2004OrgData, orgProjectData []InlineResponse2004OrgProjectData, allProjects []string, ) *InlineResponse2004`

NewInlineResponse2004 instantiates a new InlineResponse2004 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2004WithDefaults

`func NewInlineResponse2004WithDefaults() *InlineResponse2004`

NewInlineResponse2004WithDefaults instantiates a new InlineResponse2004 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgData

`func (o *InlineResponse2004) GetOrgData() InlineResponse2004OrgData`

GetOrgData returns the OrgData field if non-nil, zero value otherwise.

### GetOrgDataOk

`func (o *InlineResponse2004) GetOrgDataOk() (*InlineResponse2004OrgData, bool)`

GetOrgDataOk returns a tuple with the OrgData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgData

`func (o *InlineResponse2004) SetOrgData(v InlineResponse2004OrgData)`

SetOrgData sets OrgData field to given value.


### GetOrgProjectData

`func (o *InlineResponse2004) GetOrgProjectData() []InlineResponse2004OrgProjectData`

GetOrgProjectData returns the OrgProjectData field if non-nil, zero value otherwise.

### GetOrgProjectDataOk

`func (o *InlineResponse2004) GetOrgProjectDataOk() (*[]InlineResponse2004OrgProjectData, bool)`

GetOrgProjectDataOk returns a tuple with the OrgProjectData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgProjectData

`func (o *InlineResponse2004) SetOrgProjectData(v []InlineResponse2004OrgProjectData)`

SetOrgProjectData sets OrgProjectData field to given value.


### GetAllProjects

`func (o *InlineResponse2004) GetAllProjects() []string`

GetAllProjects returns the AllProjects field if non-nil, zero value otherwise.

### GetAllProjectsOk

`func (o *InlineResponse2004) GetAllProjectsOk() (*[]string, bool)`

GetAllProjectsOk returns a tuple with the AllProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllProjects

`func (o *InlineResponse2004) SetAllProjects(v []string)`

SetAllProjects sets AllProjects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


