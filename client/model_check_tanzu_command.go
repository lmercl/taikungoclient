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

// checks if the CheckTanzuCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckTanzuCommand{}

// CheckTanzuCommand struct for CheckTanzuCommand
type CheckTanzuCommand struct {
	Username string `json:"username"`
	Url string `json:"url"`
	Password string `json:"password"`
	Namespace string `json:"namespace"`
	Port NullableInt32 `json:"port,omitempty"`
	VolumeType string `json:"volumeType"`
}

// NewCheckTanzuCommand instantiates a new CheckTanzuCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckTanzuCommand(username string, url string, password string, namespace string, volumeType string) *CheckTanzuCommand {
	this := CheckTanzuCommand{}
	this.Username = username
	this.Url = url
	this.Password = password
	this.Namespace = namespace
	this.VolumeType = volumeType
	return &this
}

// NewCheckTanzuCommandWithDefaults instantiates a new CheckTanzuCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckTanzuCommandWithDefaults() *CheckTanzuCommand {
	this := CheckTanzuCommand{}
	return &this
}

// GetUsername returns the Username field value
func (o *CheckTanzuCommand) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *CheckTanzuCommand) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *CheckTanzuCommand) SetUsername(v string) {
	o.Username = v
}

// GetUrl returns the Url field value
func (o *CheckTanzuCommand) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CheckTanzuCommand) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CheckTanzuCommand) SetUrl(v string) {
	o.Url = v
}

// GetPassword returns the Password field value
func (o *CheckTanzuCommand) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CheckTanzuCommand) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *CheckTanzuCommand) SetPassword(v string) {
	o.Password = v
}

// GetNamespace returns the Namespace field value
func (o *CheckTanzuCommand) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *CheckTanzuCommand) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *CheckTanzuCommand) SetNamespace(v string) {
	o.Namespace = v
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckTanzuCommand) GetPort() int32 {
	if o == nil || IsNil(o.Port.Get()) {
		var ret int32
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckTanzuCommand) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *CheckTanzuCommand) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableInt32 and assigns it to the Port field.
func (o *CheckTanzuCommand) SetPort(v int32) {
	o.Port.Set(&v)
}
// SetPortNil sets the value for Port to be an explicit nil
func (o *CheckTanzuCommand) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *CheckTanzuCommand) UnsetPort() {
	o.Port.Unset()
}

// GetVolumeType returns the VolumeType field value
func (o *CheckTanzuCommand) GetVolumeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VolumeType
}

// GetVolumeTypeOk returns a tuple with the VolumeType field value
// and a boolean to check if the value has been set.
func (o *CheckTanzuCommand) GetVolumeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VolumeType, true
}

// SetVolumeType sets field value
func (o *CheckTanzuCommand) SetVolumeType(v string) {
	o.VolumeType = v
}

func (o CheckTanzuCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckTanzuCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["username"] = o.Username
	toSerialize["url"] = o.Url
	toSerialize["password"] = o.Password
	toSerialize["namespace"] = o.Namespace
	if o.Port.IsSet() {
		toSerialize["port"] = o.Port.Get()
	}
	toSerialize["volumeType"] = o.VolumeType
	return toSerialize, nil
}

type NullableCheckTanzuCommand struct {
	value *CheckTanzuCommand
	isSet bool
}

func (v NullableCheckTanzuCommand) Get() *CheckTanzuCommand {
	return v.value
}

func (v *NullableCheckTanzuCommand) Set(val *CheckTanzuCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckTanzuCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckTanzuCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckTanzuCommand(val *CheckTanzuCommand) *NullableCheckTanzuCommand {
	return &NullableCheckTanzuCommand{value: val, isSet: true}
}

func (v NullableCheckTanzuCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckTanzuCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


