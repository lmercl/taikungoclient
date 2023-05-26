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

// checks if the DescribeStsCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeStsCommand{}

// DescribeStsCommand struct for DescribeStsCommand
type DescribeStsCommand struct {
	ProjectId int32 `json:"projectId"`
	Name string `json:"name"`
	Namespace string `json:"namespace"`
}

// NewDescribeStsCommand instantiates a new DescribeStsCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeStsCommand(projectId int32, name string, namespace string) *DescribeStsCommand {
	this := DescribeStsCommand{}
	this.ProjectId = projectId
	this.Name = name
	this.Namespace = namespace
	return &this
}

// NewDescribeStsCommandWithDefaults instantiates a new DescribeStsCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeStsCommandWithDefaults() *DescribeStsCommand {
	this := DescribeStsCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *DescribeStsCommand) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *DescribeStsCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *DescribeStsCommand) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetName returns the Name field value
func (o *DescribeStsCommand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DescribeStsCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DescribeStsCommand) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value
func (o *DescribeStsCommand) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *DescribeStsCommand) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *DescribeStsCommand) SetNamespace(v string) {
	o.Namespace = v
}

func (o DescribeStsCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeStsCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["name"] = o.Name
	toSerialize["namespace"] = o.Namespace
	return toSerialize, nil
}

type NullableDescribeStsCommand struct {
	value *DescribeStsCommand
	isSet bool
}

func (v NullableDescribeStsCommand) Get() *DescribeStsCommand {
	return v.value
}

func (v *NullableDescribeStsCommand) Set(val *DescribeStsCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeStsCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeStsCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeStsCommand(val *DescribeStsCommand) *NullableDescribeStsCommand {
	return &NullableDescribeStsCommand{value: val, isSet: true}
}

func (v NullableDescribeStsCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeStsCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


