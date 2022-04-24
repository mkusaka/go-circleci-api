# JobDetailsMessages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Message type. | 
**Message** | **string** | Information describing message. | 
**Reason** | Pointer to **string** | Value describing the reason for message to be added to the job. | [optional] 

## Methods

### NewJobDetailsMessages

`func NewJobDetailsMessages(type_ string, message string, ) *JobDetailsMessages`

NewJobDetailsMessages instantiates a new JobDetailsMessages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobDetailsMessagesWithDefaults

`func NewJobDetailsMessagesWithDefaults() *JobDetailsMessages`

NewJobDetailsMessagesWithDefaults instantiates a new JobDetailsMessages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *JobDetailsMessages) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JobDetailsMessages) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JobDetailsMessages) SetType(v string)`

SetType sets Type field to given value.


### GetMessage

`func (o *JobDetailsMessages) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *JobDetailsMessages) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *JobDetailsMessages) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetReason

`func (o *JobDetailsMessages) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *JobDetailsMessages) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *JobDetailsMessages) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *JobDetailsMessages) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


