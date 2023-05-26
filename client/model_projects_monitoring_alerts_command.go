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

// checks if the ProjectsMonitoringAlertsCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectsMonitoringAlertsCommand{}

// ProjectsMonitoringAlertsCommand struct for ProjectsMonitoringAlertsCommand
type ProjectsMonitoringAlertsCommand struct {
	ProjectId *int32 `json:"projectId,omitempty"`
}

// NewProjectsMonitoringAlertsCommand instantiates a new ProjectsMonitoringAlertsCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsMonitoringAlertsCommand() *ProjectsMonitoringAlertsCommand {
	this := ProjectsMonitoringAlertsCommand{}
	return &this
}

// NewProjectsMonitoringAlertsCommandWithDefaults instantiates a new ProjectsMonitoringAlertsCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsMonitoringAlertsCommandWithDefaults() *ProjectsMonitoringAlertsCommand {
	this := ProjectsMonitoringAlertsCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ProjectsMonitoringAlertsCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsMonitoringAlertsCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ProjectsMonitoringAlertsCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *ProjectsMonitoringAlertsCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

func (o ProjectsMonitoringAlertsCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectsMonitoringAlertsCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	return toSerialize, nil
}

type NullableProjectsMonitoringAlertsCommand struct {
	value *ProjectsMonitoringAlertsCommand
	isSet bool
}

func (v NullableProjectsMonitoringAlertsCommand) Get() *ProjectsMonitoringAlertsCommand {
	return v.value
}

func (v *NullableProjectsMonitoringAlertsCommand) Set(val *ProjectsMonitoringAlertsCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsMonitoringAlertsCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsMonitoringAlertsCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsMonitoringAlertsCommand(val *ProjectsMonitoringAlertsCommand) *NullableProjectsMonitoringAlertsCommand {
	return &NullableProjectsMonitoringAlertsCommand{value: val, isSet: true}
}

func (v NullableProjectsMonitoringAlertsCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsMonitoringAlertsCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


