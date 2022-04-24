# InlineResponse2002

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **interface{}** | The unique ID of the organization | [optional] 
**ProjectId** | Pointer to **interface{}** | The unique ID of the project | [optional] 
**ProjectData** | Pointer to [**InlineResponse2002ProjectData**](InlineResponse2002ProjectData.md) |  | [optional] 
**ProjectWorkflowData** | Pointer to [**[]InlineResponse2002ProjectWorkflowData**](InlineResponse2002ProjectWorkflowData.md) | A list of metrics and trends data for workflows for a given project. | [optional] 
**ProjectWorkflowBranchData** | Pointer to [**[]InlineResponse2002ProjectWorkflowBranchData**](InlineResponse2002ProjectWorkflowBranchData.md) | A list of metrics and trends data for branches for a given project. | [optional] 
**AllBranches** | Pointer to **[]string** | A list of all the branches for a given project. | [optional] 
**AllWorkflows** | Pointer to **[]string** | A list of all the workflows for a given project. | [optional] 

## Methods

### NewInlineResponse2002

`func NewInlineResponse2002() *InlineResponse2002`

NewInlineResponse2002 instantiates a new InlineResponse2002 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2002WithDefaults

`func NewInlineResponse2002WithDefaults() *InlineResponse2002`

NewInlineResponse2002WithDefaults instantiates a new InlineResponse2002 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *InlineResponse2002) GetOrgId() interface{}`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *InlineResponse2002) GetOrgIdOk() (*interface{}, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *InlineResponse2002) SetOrgId(v interface{})`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *InlineResponse2002) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *InlineResponse2002) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *InlineResponse2002) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetProjectId

`func (o *InlineResponse2002) GetProjectId() interface{}`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *InlineResponse2002) GetProjectIdOk() (*interface{}, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *InlineResponse2002) SetProjectId(v interface{})`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *InlineResponse2002) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *InlineResponse2002) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *InlineResponse2002) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetProjectData

`func (o *InlineResponse2002) GetProjectData() InlineResponse2002ProjectData`

GetProjectData returns the ProjectData field if non-nil, zero value otherwise.

### GetProjectDataOk

`func (o *InlineResponse2002) GetProjectDataOk() (*InlineResponse2002ProjectData, bool)`

GetProjectDataOk returns a tuple with the ProjectData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectData

`func (o *InlineResponse2002) SetProjectData(v InlineResponse2002ProjectData)`

SetProjectData sets ProjectData field to given value.

### HasProjectData

`func (o *InlineResponse2002) HasProjectData() bool`

HasProjectData returns a boolean if a field has been set.

### GetProjectWorkflowData

`func (o *InlineResponse2002) GetProjectWorkflowData() []InlineResponse2002ProjectWorkflowData`

GetProjectWorkflowData returns the ProjectWorkflowData field if non-nil, zero value otherwise.

### GetProjectWorkflowDataOk

`func (o *InlineResponse2002) GetProjectWorkflowDataOk() (*[]InlineResponse2002ProjectWorkflowData, bool)`

GetProjectWorkflowDataOk returns a tuple with the ProjectWorkflowData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectWorkflowData

`func (o *InlineResponse2002) SetProjectWorkflowData(v []InlineResponse2002ProjectWorkflowData)`

SetProjectWorkflowData sets ProjectWorkflowData field to given value.

### HasProjectWorkflowData

`func (o *InlineResponse2002) HasProjectWorkflowData() bool`

HasProjectWorkflowData returns a boolean if a field has been set.

### GetProjectWorkflowBranchData

`func (o *InlineResponse2002) GetProjectWorkflowBranchData() []InlineResponse2002ProjectWorkflowBranchData`

GetProjectWorkflowBranchData returns the ProjectWorkflowBranchData field if non-nil, zero value otherwise.

### GetProjectWorkflowBranchDataOk

`func (o *InlineResponse2002) GetProjectWorkflowBranchDataOk() (*[]InlineResponse2002ProjectWorkflowBranchData, bool)`

GetProjectWorkflowBranchDataOk returns a tuple with the ProjectWorkflowBranchData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectWorkflowBranchData

`func (o *InlineResponse2002) SetProjectWorkflowBranchData(v []InlineResponse2002ProjectWorkflowBranchData)`

SetProjectWorkflowBranchData sets ProjectWorkflowBranchData field to given value.

### HasProjectWorkflowBranchData

`func (o *InlineResponse2002) HasProjectWorkflowBranchData() bool`

HasProjectWorkflowBranchData returns a boolean if a field has been set.

### GetAllBranches

`func (o *InlineResponse2002) GetAllBranches() []string`

GetAllBranches returns the AllBranches field if non-nil, zero value otherwise.

### GetAllBranchesOk

`func (o *InlineResponse2002) GetAllBranchesOk() (*[]string, bool)`

GetAllBranchesOk returns a tuple with the AllBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllBranches

`func (o *InlineResponse2002) SetAllBranches(v []string)`

SetAllBranches sets AllBranches field to given value.

### HasAllBranches

`func (o *InlineResponse2002) HasAllBranches() bool`

HasAllBranches returns a boolean if a field has been set.

### GetAllWorkflows

`func (o *InlineResponse2002) GetAllWorkflows() []string`

GetAllWorkflows returns the AllWorkflows field if non-nil, zero value otherwise.

### GetAllWorkflowsOk

`func (o *InlineResponse2002) GetAllWorkflowsOk() (*[]string, bool)`

GetAllWorkflowsOk returns a tuple with the AllWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllWorkflows

`func (o *InlineResponse2002) SetAllWorkflows(v []string)`

SetAllWorkflows sets AllWorkflows field to given value.

### HasAllWorkflows

`func (o *InlineResponse2002) HasAllWorkflows() bool`

HasAllWorkflows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


