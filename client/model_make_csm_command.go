/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikuncore

import (
	"encoding/json"
)

// checks if the MakeCsmCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MakeCsmCommand{}

// MakeCsmCommand struct for MakeCsmCommand
type MakeCsmCommand struct {
	UserId NullableString `json:"userId,omitempty"`
	Mode NullableString `json:"mode,omitempty"`
}

// NewMakeCsmCommand instantiates a new MakeCsmCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMakeCsmCommand() *MakeCsmCommand {
	this := MakeCsmCommand{}
	return &this
}

// NewMakeCsmCommandWithDefaults instantiates a new MakeCsmCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMakeCsmCommandWithDefaults() *MakeCsmCommand {
	this := MakeCsmCommand{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MakeCsmCommand) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MakeCsmCommand) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *MakeCsmCommand) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *MakeCsmCommand) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *MakeCsmCommand) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *MakeCsmCommand) UnsetUserId() {
	o.UserId.Unset()
}

// GetMode returns the Mode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MakeCsmCommand) GetMode() string {
	if o == nil || IsNil(o.Mode.Get()) {
		var ret string
		return ret
	}
	return *o.Mode.Get()
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MakeCsmCommand) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mode.Get(), o.Mode.IsSet()
}

// HasMode returns a boolean if a field has been set.
func (o *MakeCsmCommand) HasMode() bool {
	if o != nil && o.Mode.IsSet() {
		return true
	}

	return false
}

// SetMode gets a reference to the given NullableString and assigns it to the Mode field.
func (o *MakeCsmCommand) SetMode(v string) {
	o.Mode.Set(&v)
}
// SetModeNil sets the value for Mode to be an explicit nil
func (o *MakeCsmCommand) SetModeNil() {
	o.Mode.Set(nil)
}

// UnsetMode ensures that no value is present for Mode, not even an explicit nil
func (o *MakeCsmCommand) UnsetMode() {
	o.Mode.Unset()
}

func (o MakeCsmCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MakeCsmCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.UserId.IsSet() {
		toSerialize["userId"] = o.UserId.Get()
	}
	if o.Mode.IsSet() {
		toSerialize["mode"] = o.Mode.Get()
	}
	return toSerialize, nil
}

type NullableMakeCsmCommand struct {
	value *MakeCsmCommand
	isSet bool
}

func (v NullableMakeCsmCommand) Get() *MakeCsmCommand {
	return v.value
}

func (v *NullableMakeCsmCommand) Set(val *MakeCsmCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableMakeCsmCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableMakeCsmCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMakeCsmCommand(val *MakeCsmCommand) *NullableMakeCsmCommand {
	return &NullableMakeCsmCommand{value: val, isSet: true}
}

func (v NullableMakeCsmCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMakeCsmCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


