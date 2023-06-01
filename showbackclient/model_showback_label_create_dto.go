/*
Taikun - Showback API

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@itera.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikunshowback

import (
	"encoding/json"
)

// checks if the ShowbackLabelCreateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShowbackLabelCreateDto{}

// ShowbackLabelCreateDto struct for ShowbackLabelCreateDto
type ShowbackLabelCreateDto struct {
	Label *string `json:"label,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewShowbackLabelCreateDto instantiates a new ShowbackLabelCreateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShowbackLabelCreateDto() *ShowbackLabelCreateDto {
	this := ShowbackLabelCreateDto{}
	return &this
}

// NewShowbackLabelCreateDtoWithDefaults instantiates a new ShowbackLabelCreateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShowbackLabelCreateDtoWithDefaults() *ShowbackLabelCreateDto {
	this := ShowbackLabelCreateDto{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ShowbackLabelCreateDto) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShowbackLabelCreateDto) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ShowbackLabelCreateDto) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ShowbackLabelCreateDto) SetLabel(v string) {
	o.Label = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ShowbackLabelCreateDto) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShowbackLabelCreateDto) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ShowbackLabelCreateDto) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ShowbackLabelCreateDto) SetValue(v string) {
	o.Value = &v
}

func (o ShowbackLabelCreateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShowbackLabelCreateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableShowbackLabelCreateDto struct {
	value *ShowbackLabelCreateDto
	isSet bool
}

func (v NullableShowbackLabelCreateDto) Get() *ShowbackLabelCreateDto {
	return v.value
}

func (v *NullableShowbackLabelCreateDto) Set(val *ShowbackLabelCreateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableShowbackLabelCreateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableShowbackLabelCreateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShowbackLabelCreateDto(val *ShowbackLabelCreateDto) *NullableShowbackLabelCreateDto {
	return &NullableShowbackLabelCreateDto{value: val, isSet: true}
}

func (v NullableShowbackLabelCreateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShowbackLabelCreateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

