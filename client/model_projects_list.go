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

// checks if the ProjectsList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectsList{}

// ProjectsList struct for ProjectsList
type ProjectsList struct {
	Data []ProjectListDetailDto `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewProjectsList instantiates a new ProjectsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsList() *ProjectsList {
	this := ProjectsList{}
	return &this
}

// NewProjectsListWithDefaults instantiates a new ProjectsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsListWithDefaults() *ProjectsList {
	this := ProjectsList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectsList) GetData() []ProjectListDetailDto {
	if o == nil {
		var ret []ProjectListDetailDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectsList) GetDataOk() ([]ProjectListDetailDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ProjectsList) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ProjectListDetailDto and assigns it to the Data field.
func (o *ProjectsList) SetData(v []ProjectListDetailDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ProjectsList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ProjectsList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *ProjectsList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o ProjectsList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectsList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableProjectsList struct {
	value *ProjectsList
	isSet bool
}

func (v NullableProjectsList) Get() *ProjectsList {
	return v.value
}

func (v *NullableProjectsList) Set(val *ProjectsList) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsList) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsList(val *ProjectsList) *NullableProjectsList {
	return &NullableProjectsList{value: val, isSet: true}
}

func (v NullableProjectsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


