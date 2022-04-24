# ContextOwner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique ID of the owner of the context. Specify either this or slug. | [optional] 
**Slug** | Pointer to **string** | A string that represents an organization. Specify either this or id. Cannot be used for accounts. | [optional] 
**Type** | Pointer to **string** | The type of the owner. Defaults to \&quot;organization\&quot;. Accounts are only used as context owners in server. | [optional] 

## Methods

### NewContextOwner

`func NewContextOwner() *ContextOwner`

NewContextOwner instantiates a new ContextOwner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextOwnerWithDefaults

`func NewContextOwnerWithDefaults() *ContextOwner`

NewContextOwnerWithDefaults instantiates a new ContextOwner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContextOwner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContextOwner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContextOwner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ContextOwner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSlug

`func (o *ContextOwner) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ContextOwner) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ContextOwner) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ContextOwner) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetType

`func (o *ContextOwner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContextOwner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContextOwner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ContextOwner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


