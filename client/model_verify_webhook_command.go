/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@itera.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikuncore

import (
	"encoding/json"
)

// checks if the VerifyWebhookCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifyWebhookCommand{}

// VerifyWebhookCommand struct for VerifyWebhookCommand
type VerifyWebhookCommand struct {
	Url NullableString `json:"url,omitempty"`
}

// NewVerifyWebhookCommand instantiates a new VerifyWebhookCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyWebhookCommand() *VerifyWebhookCommand {
	this := VerifyWebhookCommand{}
	return &this
}

// NewVerifyWebhookCommandWithDefaults instantiates a new VerifyWebhookCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyWebhookCommandWithDefaults() *VerifyWebhookCommand {
	this := VerifyWebhookCommand{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VerifyWebhookCommand) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VerifyWebhookCommand) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *VerifyWebhookCommand) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *VerifyWebhookCommand) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *VerifyWebhookCommand) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *VerifyWebhookCommand) UnsetUrl() {
	o.Url.Unset()
}

func (o VerifyWebhookCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifyWebhookCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	return toSerialize, nil
}

type NullableVerifyWebhookCommand struct {
	value *VerifyWebhookCommand
	isSet bool
}

func (v NullableVerifyWebhookCommand) Get() *VerifyWebhookCommand {
	return v.value
}

func (v *NullableVerifyWebhookCommand) Set(val *VerifyWebhookCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyWebhookCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyWebhookCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyWebhookCommand(val *VerifyWebhookCommand) *NullableVerifyWebhookCommand {
	return &NullableVerifyWebhookCommand{value: val, isSet: true}
}

func (v NullableVerifyWebhookCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyWebhookCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


