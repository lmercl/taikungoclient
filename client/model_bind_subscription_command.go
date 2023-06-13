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

// checks if the BindSubscriptionCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BindSubscriptionCommand{}

// BindSubscriptionCommand struct for BindSubscriptionCommand
type BindSubscriptionCommand struct {
	Id *int32 `json:"id,omitempty"`
	Yearly *bool `json:"yearly,omitempty"`
}

// NewBindSubscriptionCommand instantiates a new BindSubscriptionCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBindSubscriptionCommand() *BindSubscriptionCommand {
	this := BindSubscriptionCommand{}
	return &this
}

// NewBindSubscriptionCommandWithDefaults instantiates a new BindSubscriptionCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBindSubscriptionCommandWithDefaults() *BindSubscriptionCommand {
	this := BindSubscriptionCommand{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BindSubscriptionCommand) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BindSubscriptionCommand) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BindSubscriptionCommand) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *BindSubscriptionCommand) SetId(v int32) {
	o.Id = &v
}

// GetYearly returns the Yearly field value if set, zero value otherwise.
func (o *BindSubscriptionCommand) GetYearly() bool {
	if o == nil || IsNil(o.Yearly) {
		var ret bool
		return ret
	}
	return *o.Yearly
}

// GetYearlyOk returns a tuple with the Yearly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BindSubscriptionCommand) GetYearlyOk() (*bool, bool) {
	if o == nil || IsNil(o.Yearly) {
		return nil, false
	}
	return o.Yearly, true
}

// HasYearly returns a boolean if a field has been set.
func (o *BindSubscriptionCommand) HasYearly() bool {
	if o != nil && !IsNil(o.Yearly) {
		return true
	}

	return false
}

// SetYearly gets a reference to the given bool and assigns it to the Yearly field.
func (o *BindSubscriptionCommand) SetYearly(v bool) {
	o.Yearly = &v
}

func (o BindSubscriptionCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BindSubscriptionCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Yearly) {
		toSerialize["yearly"] = o.Yearly
	}
	return toSerialize, nil
}

type NullableBindSubscriptionCommand struct {
	value *BindSubscriptionCommand
	isSet bool
}

func (v NullableBindSubscriptionCommand) Get() *BindSubscriptionCommand {
	return v.value
}

func (v *NullableBindSubscriptionCommand) Set(val *BindSubscriptionCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableBindSubscriptionCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableBindSubscriptionCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBindSubscriptionCommand(val *BindSubscriptionCommand) *NullableBindSubscriptionCommand {
	return &NullableBindSubscriptionCommand{value: val, isSet: true}
}

func (v NullableBindSubscriptionCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBindSubscriptionCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


