/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikuncore

import (
	"encoding/json"
)

// checks if the CountryListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountryListDto{}

// CountryListDto struct for CountryListDto
type CountryListDto struct {
	Name NullableString `json:"name,omitempty"`
	IsEu *bool `json:"isEu,omitempty"`
}

// NewCountryListDto instantiates a new CountryListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountryListDto() *CountryListDto {
	this := CountryListDto{}
	return &this
}

// NewCountryListDtoWithDefaults instantiates a new CountryListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountryListDtoWithDefaults() *CountryListDto {
	this := CountryListDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CountryListDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CountryListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CountryListDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CountryListDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CountryListDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CountryListDto) UnsetName() {
	o.Name.Unset()
}

// GetIsEu returns the IsEu field value if set, zero value otherwise.
func (o *CountryListDto) GetIsEu() bool {
	if o == nil || IsNil(o.IsEu) {
		var ret bool
		return ret
	}
	return *o.IsEu
}

// GetIsEuOk returns a tuple with the IsEu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryListDto) GetIsEuOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEu) {
		return nil, false
	}
	return o.IsEu, true
}

// HasIsEu returns a boolean if a field has been set.
func (o *CountryListDto) HasIsEu() bool {
	if o != nil && !IsNil(o.IsEu) {
		return true
	}

	return false
}

// SetIsEu gets a reference to the given bool and assigns it to the IsEu field.
func (o *CountryListDto) SetIsEu(v bool) {
	o.IsEu = &v
}

func (o CountryListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountryListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.IsEu) {
		toSerialize["isEu"] = o.IsEu
	}
	return toSerialize, nil
}

type NullableCountryListDto struct {
	value *CountryListDto
	isSet bool
}

func (v NullableCountryListDto) Get() *CountryListDto {
	return v.value
}

func (v *NullableCountryListDto) Set(val *CountryListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCountryListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCountryListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountryListDto(val *CountryListDto) *NullableCountryListDto {
	return &NullableCountryListDto{value: val, isSet: true}
}

func (v NullableCountryListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountryListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


