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

// checks if the DeleteProjectCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteProjectCommand{}

// DeleteProjectCommand struct for DeleteProjectCommand
type DeleteProjectCommand struct {
	ProjectId *int32 `json:"projectId,omitempty"`
	IsForceDelete *bool `json:"isForceDelete,omitempty"`
}

// NewDeleteProjectCommand instantiates a new DeleteProjectCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteProjectCommand() *DeleteProjectCommand {
	this := DeleteProjectCommand{}
	return &this
}

// NewDeleteProjectCommandWithDefaults instantiates a new DeleteProjectCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteProjectCommandWithDefaults() *DeleteProjectCommand {
	this := DeleteProjectCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *DeleteProjectCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteProjectCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *DeleteProjectCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *DeleteProjectCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetIsForceDelete returns the IsForceDelete field value if set, zero value otherwise.
func (o *DeleteProjectCommand) GetIsForceDelete() bool {
	if o == nil || IsNil(o.IsForceDelete) {
		var ret bool
		return ret
	}
	return *o.IsForceDelete
}

// GetIsForceDeleteOk returns a tuple with the IsForceDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteProjectCommand) GetIsForceDeleteOk() (*bool, bool) {
	if o == nil || IsNil(o.IsForceDelete) {
		return nil, false
	}
	return o.IsForceDelete, true
}

// HasIsForceDelete returns a boolean if a field has been set.
func (o *DeleteProjectCommand) HasIsForceDelete() bool {
	if o != nil && !IsNil(o.IsForceDelete) {
		return true
	}

	return false
}

// SetIsForceDelete gets a reference to the given bool and assigns it to the IsForceDelete field.
func (o *DeleteProjectCommand) SetIsForceDelete(v bool) {
	o.IsForceDelete = &v
}

func (o DeleteProjectCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteProjectCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.IsForceDelete) {
		toSerialize["isForceDelete"] = o.IsForceDelete
	}
	return toSerialize, nil
}

type NullableDeleteProjectCommand struct {
	value *DeleteProjectCommand
	isSet bool
}

func (v NullableDeleteProjectCommand) Get() *DeleteProjectCommand {
	return v.value
}

func (v *NullableDeleteProjectCommand) Set(val *DeleteProjectCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteProjectCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteProjectCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteProjectCommand(val *DeleteProjectCommand) *NullableDeleteProjectCommand {
	return &NullableDeleteProjectCommand{value: val, isSet: true}
}

func (v NullableDeleteProjectCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteProjectCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


