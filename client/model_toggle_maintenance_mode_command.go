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

// checks if the ToggleMaintenanceModeCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToggleMaintenanceModeCommand{}

// ToggleMaintenanceModeCommand struct for ToggleMaintenanceModeCommand
type ToggleMaintenanceModeCommand struct {
	Mode NullableString `json:"mode,omitempty"`
}

// NewToggleMaintenanceModeCommand instantiates a new ToggleMaintenanceModeCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToggleMaintenanceModeCommand() *ToggleMaintenanceModeCommand {
	this := ToggleMaintenanceModeCommand{}
	return &this
}

// NewToggleMaintenanceModeCommandWithDefaults instantiates a new ToggleMaintenanceModeCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToggleMaintenanceModeCommandWithDefaults() *ToggleMaintenanceModeCommand {
	this := ToggleMaintenanceModeCommand{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ToggleMaintenanceModeCommand) GetMode() string {
	if o == nil || IsNil(o.Mode.Get()) {
		var ret string
		return ret
	}
	return *o.Mode.Get()
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ToggleMaintenanceModeCommand) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mode.Get(), o.Mode.IsSet()
}

// HasMode returns a boolean if a field has been set.
func (o *ToggleMaintenanceModeCommand) HasMode() bool {
	if o != nil && o.Mode.IsSet() {
		return true
	}

	return false
}

// SetMode gets a reference to the given NullableString and assigns it to the Mode field.
func (o *ToggleMaintenanceModeCommand) SetMode(v string) {
	o.Mode.Set(&v)
}
// SetModeNil sets the value for Mode to be an explicit nil
func (o *ToggleMaintenanceModeCommand) SetModeNil() {
	o.Mode.Set(nil)
}

// UnsetMode ensures that no value is present for Mode, not even an explicit nil
func (o *ToggleMaintenanceModeCommand) UnsetMode() {
	o.Mode.Unset()
}

func (o ToggleMaintenanceModeCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToggleMaintenanceModeCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Mode.IsSet() {
		toSerialize["mode"] = o.Mode.Get()
	}
	return toSerialize, nil
}

type NullableToggleMaintenanceModeCommand struct {
	value *ToggleMaintenanceModeCommand
	isSet bool
}

func (v NullableToggleMaintenanceModeCommand) Get() *ToggleMaintenanceModeCommand {
	return v.value
}

func (v *NullableToggleMaintenanceModeCommand) Set(val *ToggleMaintenanceModeCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableToggleMaintenanceModeCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableToggleMaintenanceModeCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToggleMaintenanceModeCommand(val *ToggleMaintenanceModeCommand) *NullableToggleMaintenanceModeCommand {
	return &NullableToggleMaintenanceModeCommand{value: val, isSet: true}
}

func (v NullableToggleMaintenanceModeCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToggleMaintenanceModeCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


