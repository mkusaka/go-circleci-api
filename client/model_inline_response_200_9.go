/*
CircleCI API

This describes the resources that make up the CircleCI API v2.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2009 Paginated recent job runs.
type InlineResponse2009 struct {
	// Recent job runs.
	Items []InlineResponse2009Items `json:"items"`
	// A token to pass as a `page-token` query parameter to return the next page of results.
	NextPageToken string `json:"next_page_token"`
}

// NewInlineResponse2009 instantiates a new InlineResponse2009 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2009(items []InlineResponse2009Items, nextPageToken string) *InlineResponse2009 {
	this := InlineResponse2009{}
	this.Items = items
	this.NextPageToken = nextPageToken
	return &this
}

// NewInlineResponse2009WithDefaults instantiates a new InlineResponse2009 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2009WithDefaults() *InlineResponse2009 {
	this := InlineResponse2009{}
	return &this
}

// GetItems returns the Items field value
func (o *InlineResponse2009) GetItems() []InlineResponse2009Items {
	if o == nil {
		var ret []InlineResponse2009Items
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2009) GetItemsOk() ([]InlineResponse2009Items, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *InlineResponse2009) SetItems(v []InlineResponse2009Items) {
	o.Items = v
}

// GetNextPageToken returns the NextPageToken field value
func (o *InlineResponse2009) GetNextPageToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2009) GetNextPageTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextPageToken, true
}

// SetNextPageToken sets field value
func (o *InlineResponse2009) SetNextPageToken(v string) {
	o.NextPageToken = v
}

func (o InlineResponse2009) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	if true {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2009 struct {
	value *InlineResponse2009
	isSet bool
}

func (v NullableInlineResponse2009) Get() *InlineResponse2009 {
	return v.value
}

func (v *NullableInlineResponse2009) Set(val *InlineResponse2009) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2009) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2009) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2009(val *InlineResponse2009) *NullableInlineResponse2009 {
	return &NullableInlineResponse2009{value: val, isSet: true}
}

func (v NullableInlineResponse2009) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2009) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}