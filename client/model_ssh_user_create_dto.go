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

// checks if the SshUserCreateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SshUserCreateDto{}

// SshUserCreateDto struct for SshUserCreateDto
type SshUserCreateDto struct {
	Name string `json:"name"`
	SshPublicKey string `json:"sshPublicKey"`
}

// NewSshUserCreateDto instantiates a new SshUserCreateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSshUserCreateDto(name string, sshPublicKey string) *SshUserCreateDto {
	this := SshUserCreateDto{}
	this.Name = name
	this.SshPublicKey = sshPublicKey
	return &this
}

// NewSshUserCreateDtoWithDefaults instantiates a new SshUserCreateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshUserCreateDtoWithDefaults() *SshUserCreateDto {
	this := SshUserCreateDto{}
	return &this
}

// GetName returns the Name field value
func (o *SshUserCreateDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SshUserCreateDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SshUserCreateDto) SetName(v string) {
	o.Name = v
}

// GetSshPublicKey returns the SshPublicKey field value
func (o *SshUserCreateDto) GetSshPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SshPublicKey
}

// GetSshPublicKeyOk returns a tuple with the SshPublicKey field value
// and a boolean to check if the value has been set.
func (o *SshUserCreateDto) GetSshPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SshPublicKey, true
}

// SetSshPublicKey sets field value
func (o *SshUserCreateDto) SetSshPublicKey(v string) {
	o.SshPublicKey = v
}

func (o SshUserCreateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SshUserCreateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["sshPublicKey"] = o.SshPublicKey
	return toSerialize, nil
}

type NullableSshUserCreateDto struct {
	value *SshUserCreateDto
	isSet bool
}

func (v NullableSshUserCreateDto) Get() *SshUserCreateDto {
	return v.value
}

func (v *NullableSshUserCreateDto) Set(val *SshUserCreateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSshUserCreateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSshUserCreateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSshUserCreateDto(val *SshUserCreateDto) *NullableSshUserCreateDto {
	return &NullableSshUserCreateDto{value: val, isSet: true}
}

func (v NullableSshUserCreateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSshUserCreateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


