# CheckoutKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicKey** | **string** | A public SSH key. | 
**Type** | **string** | The type of checkout key. This may be either &#x60;deploy-key&#x60; or &#x60;github-user-key&#x60;. | 
**Fingerprint** | **string** | An SSH key fingerprint. | 
**Preferred** | **bool** | A boolean value that indicates if this key is preferred. | 
**CreatedAt** | **time.Time** | The date and time the checkout key was created. | 

## Methods

### NewCheckoutKey

`func NewCheckoutKey(publicKey string, type_ string, fingerprint string, preferred bool, createdAt time.Time, ) *CheckoutKey`

NewCheckoutKey instantiates a new CheckoutKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutKeyWithDefaults

`func NewCheckoutKeyWithDefaults() *CheckoutKey`

NewCheckoutKeyWithDefaults instantiates a new CheckoutKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicKey

`func (o *CheckoutKey) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *CheckoutKey) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *CheckoutKey) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetType

`func (o *CheckoutKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckoutKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckoutKey) SetType(v string)`

SetType sets Type field to given value.


### GetFingerprint

`func (o *CheckoutKey) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *CheckoutKey) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *CheckoutKey) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.


### GetPreferred

`func (o *CheckoutKey) GetPreferred() bool`

GetPreferred returns the Preferred field if non-nil, zero value otherwise.

### GetPreferredOk

`func (o *CheckoutKey) GetPreferredOk() (*bool, bool)`

GetPreferredOk returns a tuple with the Preferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferred

`func (o *CheckoutKey) SetPreferred(v bool)`

SetPreferred sets Preferred field to given value.


### GetCreatedAt

`func (o *CheckoutKey) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CheckoutKey) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CheckoutKey) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


