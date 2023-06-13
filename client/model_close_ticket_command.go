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

// checks if the CloseTicketCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloseTicketCommand{}

// CloseTicketCommand struct for CloseTicketCommand
type CloseTicketCommand struct {
	TicketId NullableString `json:"ticketId,omitempty"`
}

// NewCloseTicketCommand instantiates a new CloseTicketCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloseTicketCommand() *CloseTicketCommand {
	this := CloseTicketCommand{}
	return &this
}

// NewCloseTicketCommandWithDefaults instantiates a new CloseTicketCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloseTicketCommandWithDefaults() *CloseTicketCommand {
	this := CloseTicketCommand{}
	return &this
}

// GetTicketId returns the TicketId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloseTicketCommand) GetTicketId() string {
	if o == nil || IsNil(o.TicketId.Get()) {
		var ret string
		return ret
	}
	return *o.TicketId.Get()
}

// GetTicketIdOk returns a tuple with the TicketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloseTicketCommand) GetTicketIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TicketId.Get(), o.TicketId.IsSet()
}

// HasTicketId returns a boolean if a field has been set.
func (o *CloseTicketCommand) HasTicketId() bool {
	if o != nil && o.TicketId.IsSet() {
		return true
	}

	return false
}

// SetTicketId gets a reference to the given NullableString and assigns it to the TicketId field.
func (o *CloseTicketCommand) SetTicketId(v string) {
	o.TicketId.Set(&v)
}
// SetTicketIdNil sets the value for TicketId to be an explicit nil
func (o *CloseTicketCommand) SetTicketIdNil() {
	o.TicketId.Set(nil)
}

// UnsetTicketId ensures that no value is present for TicketId, not even an explicit nil
func (o *CloseTicketCommand) UnsetTicketId() {
	o.TicketId.Unset()
}

func (o CloseTicketCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloseTicketCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TicketId.IsSet() {
		toSerialize["ticketId"] = o.TicketId.Get()
	}
	return toSerialize, nil
}

type NullableCloseTicketCommand struct {
	value *CloseTicketCommand
	isSet bool
}

func (v NullableCloseTicketCommand) Get() *CloseTicketCommand {
	return v.value
}

func (v *NullableCloseTicketCommand) Set(val *CloseTicketCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCloseTicketCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCloseTicketCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloseTicketCommand(val *CloseTicketCommand) *NullableCloseTicketCommand {
	return &NullableCloseTicketCommand{value: val, isSet: true}
}

func (v NullableCloseTicketCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloseTicketCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


