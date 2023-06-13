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

// checks if the StandAloneMetaDataDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandAloneMetaDataDto{}

// StandAloneMetaDataDto struct for StandAloneMetaDataDto
type StandAloneMetaDataDto struct {
	Key NullableString `json:"key,omitempty"`
	Value NullableString `json:"value,omitempty"`
}

// NewStandAloneMetaDataDto instantiates a new StandAloneMetaDataDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandAloneMetaDataDto() *StandAloneMetaDataDto {
	this := StandAloneMetaDataDto{}
	return &this
}

// NewStandAloneMetaDataDtoWithDefaults instantiates a new StandAloneMetaDataDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandAloneMetaDataDtoWithDefaults() *StandAloneMetaDataDto {
	this := StandAloneMetaDataDto{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandAloneMetaDataDto) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandAloneMetaDataDto) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *StandAloneMetaDataDto) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *StandAloneMetaDataDto) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *StandAloneMetaDataDto) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *StandAloneMetaDataDto) UnsetKey() {
	o.Key.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandAloneMetaDataDto) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandAloneMetaDataDto) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *StandAloneMetaDataDto) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *StandAloneMetaDataDto) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *StandAloneMetaDataDto) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *StandAloneMetaDataDto) UnsetValue() {
	o.Value.Unset()
}

func (o StandAloneMetaDataDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StandAloneMetaDataDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return toSerialize, nil
}

type NullableStandAloneMetaDataDto struct {
	value *StandAloneMetaDataDto
	isSet bool
}

func (v NullableStandAloneMetaDataDto) Get() *StandAloneMetaDataDto {
	return v.value
}

func (v *NullableStandAloneMetaDataDto) Set(val *StandAloneMetaDataDto) {
	v.value = val
	v.isSet = true
}

func (v NullableStandAloneMetaDataDto) IsSet() bool {
	return v.isSet
}

func (v *NullableStandAloneMetaDataDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandAloneMetaDataDto(val *StandAloneMetaDataDto) *NullableStandAloneMetaDataDto {
	return &NullableStandAloneMetaDataDto{value: val, isSet: true}
}

func (v NullableStandAloneMetaDataDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandAloneMetaDataDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


