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

// checks if the PatchNodeCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchNodeCommand{}

// PatchNodeCommand struct for PatchNodeCommand
type PatchNodeCommand struct {
	ProjectId int32 `json:"projectId"`
	Name string `json:"name"`
	Parameters []PatchNodeLabelsDto `json:"parameters,omitempty"`
}

// NewPatchNodeCommand instantiates a new PatchNodeCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchNodeCommand(projectId int32, name string) *PatchNodeCommand {
	this := PatchNodeCommand{}
	this.ProjectId = projectId
	this.Name = name
	return &this
}

// NewPatchNodeCommandWithDefaults instantiates a new PatchNodeCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchNodeCommandWithDefaults() *PatchNodeCommand {
	this := PatchNodeCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *PatchNodeCommand) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *PatchNodeCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *PatchNodeCommand) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetName returns the Name field value
func (o *PatchNodeCommand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PatchNodeCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PatchNodeCommand) SetName(v string) {
	o.Name = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchNodeCommand) GetParameters() []PatchNodeLabelsDto {
	if o == nil {
		var ret []PatchNodeLabelsDto
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchNodeCommand) GetParametersOk() ([]PatchNodeLabelsDto, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *PatchNodeCommand) HasParameters() bool {
	if o != nil && IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []PatchNodeLabelsDto and assigns it to the Parameters field.
func (o *PatchNodeCommand) SetParameters(v []PatchNodeLabelsDto) {
	o.Parameters = v
}

func (o PatchNodeCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchNodeCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["name"] = o.Name
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	return toSerialize, nil
}

type NullablePatchNodeCommand struct {
	value *PatchNodeCommand
	isSet bool
}

func (v NullablePatchNodeCommand) Get() *PatchNodeCommand {
	return v.value
}

func (v *NullablePatchNodeCommand) Set(val *PatchNodeCommand) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchNodeCommand) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchNodeCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchNodeCommand(val *PatchNodeCommand) *NullablePatchNodeCommand {
	return &NullablePatchNodeCommand{value: val, isSet: true}
}

func (v NullablePatchNodeCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchNodeCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


