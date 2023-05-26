/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@itera.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikungoclient

import (
	"encoding/json"
)

// checks if the RefreshTokenCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefreshTokenCommand{}

// RefreshTokenCommand struct for RefreshTokenCommand
type RefreshTokenCommand struct {
	Token NullableString `json:"token,omitempty"`
	RefreshToken NullableString `json:"refreshToken,omitempty"`
}

// NewRefreshTokenCommand instantiates a new RefreshTokenCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshTokenCommand() *RefreshTokenCommand {
	this := RefreshTokenCommand{}
	return &this
}

// NewRefreshTokenCommandWithDefaults instantiates a new RefreshTokenCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshTokenCommandWithDefaults() *RefreshTokenCommand {
	this := RefreshTokenCommand{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RefreshTokenCommand) GetToken() string {
	if o == nil || IsNil(o.Token.Get()) {
		var ret string
		return ret
	}
	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RefreshTokenCommand) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// HasToken returns a boolean if a field has been set.
func (o *RefreshTokenCommand) HasToken() bool {
	if o != nil && o.Token.IsSet() {
		return true
	}

	return false
}

// SetToken gets a reference to the given NullableString and assigns it to the Token field.
func (o *RefreshTokenCommand) SetToken(v string) {
	o.Token.Set(&v)
}
// SetTokenNil sets the value for Token to be an explicit nil
func (o *RefreshTokenCommand) SetTokenNil() {
	o.Token.Set(nil)
}

// UnsetToken ensures that no value is present for Token, not even an explicit nil
func (o *RefreshTokenCommand) UnsetToken() {
	o.Token.Unset()
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RefreshTokenCommand) GetRefreshToken() string {
	if o == nil || IsNil(o.RefreshToken.Get()) {
		var ret string
		return ret
	}
	return *o.RefreshToken.Get()
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RefreshTokenCommand) GetRefreshTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefreshToken.Get(), o.RefreshToken.IsSet()
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *RefreshTokenCommand) HasRefreshToken() bool {
	if o != nil && o.RefreshToken.IsSet() {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given NullableString and assigns it to the RefreshToken field.
func (o *RefreshTokenCommand) SetRefreshToken(v string) {
	o.RefreshToken.Set(&v)
}
// SetRefreshTokenNil sets the value for RefreshToken to be an explicit nil
func (o *RefreshTokenCommand) SetRefreshTokenNil() {
	o.RefreshToken.Set(nil)
}

// UnsetRefreshToken ensures that no value is present for RefreshToken, not even an explicit nil
func (o *RefreshTokenCommand) UnsetRefreshToken() {
	o.RefreshToken.Unset()
}

func (o RefreshTokenCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefreshTokenCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Token.IsSet() {
		toSerialize["token"] = o.Token.Get()
	}
	if o.RefreshToken.IsSet() {
		toSerialize["refreshToken"] = o.RefreshToken.Get()
	}
	return toSerialize, nil
}

type NullableRefreshTokenCommand struct {
	value *RefreshTokenCommand
	isSet bool
}

func (v NullableRefreshTokenCommand) Get() *RefreshTokenCommand {
	return v.value
}

func (v *NullableRefreshTokenCommand) Set(val *RefreshTokenCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshTokenCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshTokenCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshTokenCommand(val *RefreshTokenCommand) *NullableRefreshTokenCommand {
	return &NullableRefreshTokenCommand{value: val, isSet: true}
}

func (v NullableRefreshTokenCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshTokenCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


