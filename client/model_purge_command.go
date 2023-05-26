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

// checks if the PurgeCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PurgeCommand{}

// PurgeCommand struct for PurgeCommand
type PurgeCommand struct {
	ProjectId *int32 `json:"projectId,omitempty"`
	ServerIds []int32 `json:"serverIds,omitempty"`
}

// NewPurgeCommand instantiates a new PurgeCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurgeCommand() *PurgeCommand {
	this := PurgeCommand{}
	return &this
}

// NewPurgeCommandWithDefaults instantiates a new PurgeCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurgeCommandWithDefaults() *PurgeCommand {
	this := PurgeCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *PurgeCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurgeCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *PurgeCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *PurgeCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetServerIds returns the ServerIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurgeCommand) GetServerIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.ServerIds
}

// GetServerIdsOk returns a tuple with the ServerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurgeCommand) GetServerIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.ServerIds) {
		return nil, false
	}
	return o.ServerIds, true
}

// HasServerIds returns a boolean if a field has been set.
func (o *PurgeCommand) HasServerIds() bool {
	if o != nil && IsNil(o.ServerIds) {
		return true
	}

	return false
}

// SetServerIds gets a reference to the given []int32 and assigns it to the ServerIds field.
func (o *PurgeCommand) SetServerIds(v []int32) {
	o.ServerIds = v
}

func (o PurgeCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PurgeCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.ServerIds != nil {
		toSerialize["serverIds"] = o.ServerIds
	}
	return toSerialize, nil
}

type NullablePurgeCommand struct {
	value *PurgeCommand
	isSet bool
}

func (v NullablePurgeCommand) Get() *PurgeCommand {
	return v.value
}

func (v *NullablePurgeCommand) Set(val *PurgeCommand) {
	v.value = val
	v.isSet = true
}

func (v NullablePurgeCommand) IsSet() bool {
	return v.isSet
}

func (v *NullablePurgeCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurgeCommand(val *PurgeCommand) *NullablePurgeCommand {
	return &NullablePurgeCommand{value: val, isSet: true}
}

func (v NullablePurgeCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePurgeCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


