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

// checks if the DownloadKubeConfigCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DownloadKubeConfigCommand{}

// DownloadKubeConfigCommand struct for DownloadKubeConfigCommand
type DownloadKubeConfigCommand struct {
	Id *int32 `json:"id,omitempty"`
	ProjectId *int32 `json:"projectId,omitempty"`
}

// NewDownloadKubeConfigCommand instantiates a new DownloadKubeConfigCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDownloadKubeConfigCommand() *DownloadKubeConfigCommand {
	this := DownloadKubeConfigCommand{}
	return &this
}

// NewDownloadKubeConfigCommandWithDefaults instantiates a new DownloadKubeConfigCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDownloadKubeConfigCommandWithDefaults() *DownloadKubeConfigCommand {
	this := DownloadKubeConfigCommand{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DownloadKubeConfigCommand) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadKubeConfigCommand) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DownloadKubeConfigCommand) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *DownloadKubeConfigCommand) SetId(v int32) {
	o.Id = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *DownloadKubeConfigCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadKubeConfigCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *DownloadKubeConfigCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *DownloadKubeConfigCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

func (o DownloadKubeConfigCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DownloadKubeConfigCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	return toSerialize, nil
}

type NullableDownloadKubeConfigCommand struct {
	value *DownloadKubeConfigCommand
	isSet bool
}

func (v NullableDownloadKubeConfigCommand) Get() *DownloadKubeConfigCommand {
	return v.value
}

func (v *NullableDownloadKubeConfigCommand) Set(val *DownloadKubeConfigCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDownloadKubeConfigCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDownloadKubeConfigCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDownloadKubeConfigCommand(val *DownloadKubeConfigCommand) *NullableDownloadKubeConfigCommand {
	return &NullableDownloadKubeConfigCommand{value: val, isSet: true}
}

func (v NullableDownloadKubeConfigCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDownloadKubeConfigCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


