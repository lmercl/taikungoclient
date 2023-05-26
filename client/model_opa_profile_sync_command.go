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

// checks if the OpaProfileSyncCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpaProfileSyncCommand{}

// OpaProfileSyncCommand struct for OpaProfileSyncCommand
type OpaProfileSyncCommand struct {
	ProjectId int32 `json:"projectId"`
}

// NewOpaProfileSyncCommand instantiates a new OpaProfileSyncCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpaProfileSyncCommand(projectId int32) *OpaProfileSyncCommand {
	this := OpaProfileSyncCommand{}
	this.ProjectId = projectId
	return &this
}

// NewOpaProfileSyncCommandWithDefaults instantiates a new OpaProfileSyncCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpaProfileSyncCommandWithDefaults() *OpaProfileSyncCommand {
	this := OpaProfileSyncCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *OpaProfileSyncCommand) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *OpaProfileSyncCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *OpaProfileSyncCommand) SetProjectId(v int32) {
	o.ProjectId = v
}

func (o OpaProfileSyncCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpaProfileSyncCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableOpaProfileSyncCommand struct {
	value *OpaProfileSyncCommand
	isSet bool
}

func (v NullableOpaProfileSyncCommand) Get() *OpaProfileSyncCommand {
	return v.value
}

func (v *NullableOpaProfileSyncCommand) Set(val *OpaProfileSyncCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableOpaProfileSyncCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableOpaProfileSyncCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpaProfileSyncCommand(val *OpaProfileSyncCommand) *NullableOpaProfileSyncCommand {
	return &NullableOpaProfileSyncCommand{value: val, isSet: true}
}

func (v NullableOpaProfileSyncCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpaProfileSyncCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


