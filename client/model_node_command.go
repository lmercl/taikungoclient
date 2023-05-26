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

// checks if the NodeCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeCommand{}

// NodeCommand struct for NodeCommand
type NodeCommand struct {
	ProjectId int32 `json:"projectId"`
	Name string `json:"name"`
}

// NewNodeCommand instantiates a new NodeCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeCommand(projectId int32, name string) *NodeCommand {
	this := NodeCommand{}
	this.ProjectId = projectId
	this.Name = name
	return &this
}

// NewNodeCommandWithDefaults instantiates a new NodeCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeCommandWithDefaults() *NodeCommand {
	this := NodeCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *NodeCommand) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *NodeCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *NodeCommand) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetName returns the Name field value
func (o *NodeCommand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NodeCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NodeCommand) SetName(v string) {
	o.Name = v
}

func (o NodeCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableNodeCommand struct {
	value *NodeCommand
	isSet bool
}

func (v NullableNodeCommand) Get() *NodeCommand {
	return v.value
}

func (v *NullableNodeCommand) Set(val *NodeCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeCommand(val *NodeCommand) *NullableNodeCommand {
	return &NullableNodeCommand{value: val, isSet: true}
}

func (v NullableNodeCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


