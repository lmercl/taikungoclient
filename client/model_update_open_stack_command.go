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

// checks if the UpdateOpenStackCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOpenStackCommand{}

// UpdateOpenStackCommand struct for UpdateOpenStackCommand
type UpdateOpenStackCommand struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	OpenStackUser NullableString `json:"openStackUser,omitempty"`
	OpenStackPassword NullableString `json:"openStackPassword,omitempty"`
}

// NewUpdateOpenStackCommand instantiates a new UpdateOpenStackCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOpenStackCommand() *UpdateOpenStackCommand {
	this := UpdateOpenStackCommand{}
	return &this
}

// NewUpdateOpenStackCommandWithDefaults instantiates a new UpdateOpenStackCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOpenStackCommandWithDefaults() *UpdateOpenStackCommand {
	this := UpdateOpenStackCommand{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateOpenStackCommand) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOpenStackCommand) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateOpenStackCommand) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UpdateOpenStackCommand) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOpenStackCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOpenStackCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpdateOpenStackCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpdateOpenStackCommand) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *UpdateOpenStackCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpdateOpenStackCommand) UnsetName() {
	o.Name.Unset()
}

// GetOpenStackUser returns the OpenStackUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOpenStackCommand) GetOpenStackUser() string {
	if o == nil || IsNil(o.OpenStackUser.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackUser.Get()
}

// GetOpenStackUserOk returns a tuple with the OpenStackUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOpenStackCommand) GetOpenStackUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackUser.Get(), o.OpenStackUser.IsSet()
}

// HasOpenStackUser returns a boolean if a field has been set.
func (o *UpdateOpenStackCommand) HasOpenStackUser() bool {
	if o != nil && o.OpenStackUser.IsSet() {
		return true
	}

	return false
}

// SetOpenStackUser gets a reference to the given NullableString and assigns it to the OpenStackUser field.
func (o *UpdateOpenStackCommand) SetOpenStackUser(v string) {
	o.OpenStackUser.Set(&v)
}
// SetOpenStackUserNil sets the value for OpenStackUser to be an explicit nil
func (o *UpdateOpenStackCommand) SetOpenStackUserNil() {
	o.OpenStackUser.Set(nil)
}

// UnsetOpenStackUser ensures that no value is present for OpenStackUser, not even an explicit nil
func (o *UpdateOpenStackCommand) UnsetOpenStackUser() {
	o.OpenStackUser.Unset()
}

// GetOpenStackPassword returns the OpenStackPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOpenStackCommand) GetOpenStackPassword() string {
	if o == nil || IsNil(o.OpenStackPassword.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackPassword.Get()
}

// GetOpenStackPasswordOk returns a tuple with the OpenStackPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOpenStackCommand) GetOpenStackPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackPassword.Get(), o.OpenStackPassword.IsSet()
}

// HasOpenStackPassword returns a boolean if a field has been set.
func (o *UpdateOpenStackCommand) HasOpenStackPassword() bool {
	if o != nil && o.OpenStackPassword.IsSet() {
		return true
	}

	return false
}

// SetOpenStackPassword gets a reference to the given NullableString and assigns it to the OpenStackPassword field.
func (o *UpdateOpenStackCommand) SetOpenStackPassword(v string) {
	o.OpenStackPassword.Set(&v)
}
// SetOpenStackPasswordNil sets the value for OpenStackPassword to be an explicit nil
func (o *UpdateOpenStackCommand) SetOpenStackPasswordNil() {
	o.OpenStackPassword.Set(nil)
}

// UnsetOpenStackPassword ensures that no value is present for OpenStackPassword, not even an explicit nil
func (o *UpdateOpenStackCommand) UnsetOpenStackPassword() {
	o.OpenStackPassword.Unset()
}

func (o UpdateOpenStackCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOpenStackCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.OpenStackUser.IsSet() {
		toSerialize["openStackUser"] = o.OpenStackUser.Get()
	}
	if o.OpenStackPassword.IsSet() {
		toSerialize["openStackPassword"] = o.OpenStackPassword.Get()
	}
	return toSerialize, nil
}

type NullableUpdateOpenStackCommand struct {
	value *UpdateOpenStackCommand
	isSet bool
}

func (v NullableUpdateOpenStackCommand) Get() *UpdateOpenStackCommand {
	return v.value
}

func (v *NullableUpdateOpenStackCommand) Set(val *UpdateOpenStackCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOpenStackCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOpenStackCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOpenStackCommand(val *UpdateOpenStackCommand) *NullableUpdateOpenStackCommand {
	return &NullableUpdateOpenStackCommand{value: val, isSet: true}
}

func (v NullableUpdateOpenStackCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOpenStackCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


