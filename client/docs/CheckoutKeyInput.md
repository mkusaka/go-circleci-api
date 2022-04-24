# CheckoutKeyInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of checkout key to create. This may be either &#x60;deploy-key&#x60; or &#x60;user-key&#x60;. | 

## Methods

### NewCheckoutKeyInput

`func NewCheckoutKeyInput(type_ string, ) *CheckoutKeyInput`

NewCheckoutKeyInput instantiates a new CheckoutKeyInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutKeyInputWithDefaults

`func NewCheckoutKeyInputWithDefaults() *CheckoutKeyInput`

NewCheckoutKeyInputWithDefaults instantiates a new CheckoutKeyInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CheckoutKeyInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckoutKeyInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckoutKeyInput) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


