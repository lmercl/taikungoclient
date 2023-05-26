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

// checks if the DescribeSecretCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeSecretCommand{}

// DescribeSecretCommand struct for DescribeSecretCommand
type DescribeSecretCommand struct {
	ProjectId int32 `json:"projectId"`
	Name string `json:"name"`
	Namespace string `json:"namespace"`
}

// NewDescribeSecretCommand instantiates a new DescribeSecretCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeSecretCommand(projectId int32, name string, namespace string) *DescribeSecretCommand {
	this := DescribeSecretCommand{}
	this.ProjectId = projectId
	this.Name = name
	this.Namespace = namespace
	return &this
}

// NewDescribeSecretCommandWithDefaults instantiates a new DescribeSecretCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeSecretCommandWithDefaults() *DescribeSecretCommand {
	this := DescribeSecretCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *DescribeSecretCommand) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *DescribeSecretCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *DescribeSecretCommand) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetName returns the Name field value
func (o *DescribeSecretCommand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DescribeSecretCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DescribeSecretCommand) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value
func (o *DescribeSecretCommand) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *DescribeSecretCommand) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *DescribeSecretCommand) SetNamespace(v string) {
	o.Namespace = v
}

func (o DescribeSecretCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeSecretCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["name"] = o.Name
	toSerialize["namespace"] = o.Namespace
	return toSerialize, nil
}

type NullableDescribeSecretCommand struct {
	value *DescribeSecretCommand
	isSet bool
}

func (v NullableDescribeSecretCommand) Get() *DescribeSecretCommand {
	return v.value
}

func (v *NullableDescribeSecretCommand) Set(val *DescribeSecretCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeSecretCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeSecretCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeSecretCommand(val *DescribeSecretCommand) *NullableDescribeSecretCommand {
	return &NullableDescribeSecretCommand{value: val, isSet: true}
}

func (v NullableDescribeSecretCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeSecretCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


