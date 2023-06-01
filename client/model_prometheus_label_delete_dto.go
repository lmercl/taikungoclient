/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@itera.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikuncore

import (
	"encoding/json"
)

// checks if the PrometheusLabelDeleteDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrometheusLabelDeleteDto{}

// PrometheusLabelDeleteDto struct for PrometheusLabelDeleteDto
type PrometheusLabelDeleteDto struct {
	Id *int32 `json:"id,omitempty"`
}

// NewPrometheusLabelDeleteDto instantiates a new PrometheusLabelDeleteDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrometheusLabelDeleteDto() *PrometheusLabelDeleteDto {
	this := PrometheusLabelDeleteDto{}
	return &this
}

// NewPrometheusLabelDeleteDtoWithDefaults instantiates a new PrometheusLabelDeleteDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrometheusLabelDeleteDtoWithDefaults() *PrometheusLabelDeleteDto {
	this := PrometheusLabelDeleteDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PrometheusLabelDeleteDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusLabelDeleteDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PrometheusLabelDeleteDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PrometheusLabelDeleteDto) SetId(v int32) {
	o.Id = &v
}

func (o PrometheusLabelDeleteDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrometheusLabelDeleteDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePrometheusLabelDeleteDto struct {
	value *PrometheusLabelDeleteDto
	isSet bool
}

func (v NullablePrometheusLabelDeleteDto) Get() *PrometheusLabelDeleteDto {
	return v.value
}

func (v *NullablePrometheusLabelDeleteDto) Set(val *PrometheusLabelDeleteDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePrometheusLabelDeleteDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePrometheusLabelDeleteDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrometheusLabelDeleteDto(val *PrometheusLabelDeleteDto) *NullablePrometheusLabelDeleteDto {
	return &NullablePrometheusLabelDeleteDto{value: val, isSet: true}
}

func (v NullablePrometheusLabelDeleteDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrometheusLabelDeleteDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

