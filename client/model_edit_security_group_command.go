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

// checks if the EditSecurityGroupCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditSecurityGroupCommand{}

// EditSecurityGroupCommand struct for EditSecurityGroupCommand
type EditSecurityGroupCommand struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	Protocol *SecurityGroupProtocol `json:"protocol,omitempty"`
	PortMinRange *int32 `json:"portMinRange,omitempty"`
	PortMaxRange *int32 `json:"portMaxRange,omitempty"`
	RemoteIpPrefix string `json:"remoteIpPrefix"`
}

// NewEditSecurityGroupCommand instantiates a new EditSecurityGroupCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditSecurityGroupCommand(id int32, name string, remoteIpPrefix string) *EditSecurityGroupCommand {
	this := EditSecurityGroupCommand{}
	this.Id = id
	this.Name = name
	this.RemoteIpPrefix = remoteIpPrefix
	return &this
}

// NewEditSecurityGroupCommandWithDefaults instantiates a new EditSecurityGroupCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditSecurityGroupCommandWithDefaults() *EditSecurityGroupCommand {
	this := EditSecurityGroupCommand{}
	return &this
}

// GetId returns the Id field value
func (o *EditSecurityGroupCommand) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EditSecurityGroupCommand) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EditSecurityGroupCommand) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *EditSecurityGroupCommand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EditSecurityGroupCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EditSecurityGroupCommand) SetName(v string) {
	o.Name = v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *EditSecurityGroupCommand) GetProtocol() SecurityGroupProtocol {
	if o == nil || IsNil(o.Protocol) {
		var ret SecurityGroupProtocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditSecurityGroupCommand) GetProtocolOk() (*SecurityGroupProtocol, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *EditSecurityGroupCommand) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given SecurityGroupProtocol and assigns it to the Protocol field.
func (o *EditSecurityGroupCommand) SetProtocol(v SecurityGroupProtocol) {
	o.Protocol = &v
}

// GetPortMinRange returns the PortMinRange field value if set, zero value otherwise.
func (o *EditSecurityGroupCommand) GetPortMinRange() int32 {
	if o == nil || IsNil(o.PortMinRange) {
		var ret int32
		return ret
	}
	return *o.PortMinRange
}

// GetPortMinRangeOk returns a tuple with the PortMinRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditSecurityGroupCommand) GetPortMinRangeOk() (*int32, bool) {
	if o == nil || IsNil(o.PortMinRange) {
		return nil, false
	}
	return o.PortMinRange, true
}

// HasPortMinRange returns a boolean if a field has been set.
func (o *EditSecurityGroupCommand) HasPortMinRange() bool {
	if o != nil && !IsNil(o.PortMinRange) {
		return true
	}

	return false
}

// SetPortMinRange gets a reference to the given int32 and assigns it to the PortMinRange field.
func (o *EditSecurityGroupCommand) SetPortMinRange(v int32) {
	o.PortMinRange = &v
}

// GetPortMaxRange returns the PortMaxRange field value if set, zero value otherwise.
func (o *EditSecurityGroupCommand) GetPortMaxRange() int32 {
	if o == nil || IsNil(o.PortMaxRange) {
		var ret int32
		return ret
	}
	return *o.PortMaxRange
}

// GetPortMaxRangeOk returns a tuple with the PortMaxRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditSecurityGroupCommand) GetPortMaxRangeOk() (*int32, bool) {
	if o == nil || IsNil(o.PortMaxRange) {
		return nil, false
	}
	return o.PortMaxRange, true
}

// HasPortMaxRange returns a boolean if a field has been set.
func (o *EditSecurityGroupCommand) HasPortMaxRange() bool {
	if o != nil && !IsNil(o.PortMaxRange) {
		return true
	}

	return false
}

// SetPortMaxRange gets a reference to the given int32 and assigns it to the PortMaxRange field.
func (o *EditSecurityGroupCommand) SetPortMaxRange(v int32) {
	o.PortMaxRange = &v
}

// GetRemoteIpPrefix returns the RemoteIpPrefix field value
func (o *EditSecurityGroupCommand) GetRemoteIpPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RemoteIpPrefix
}

// GetRemoteIpPrefixOk returns a tuple with the RemoteIpPrefix field value
// and a boolean to check if the value has been set.
func (o *EditSecurityGroupCommand) GetRemoteIpPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoteIpPrefix, true
}

// SetRemoteIpPrefix sets field value
func (o *EditSecurityGroupCommand) SetRemoteIpPrefix(v string) {
	o.RemoteIpPrefix = v
}

func (o EditSecurityGroupCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditSecurityGroupCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.PortMinRange) {
		toSerialize["portMinRange"] = o.PortMinRange
	}
	if !IsNil(o.PortMaxRange) {
		toSerialize["portMaxRange"] = o.PortMaxRange
	}
	toSerialize["remoteIpPrefix"] = o.RemoteIpPrefix
	return toSerialize, nil
}

type NullableEditSecurityGroupCommand struct {
	value *EditSecurityGroupCommand
	isSet bool
}

func (v NullableEditSecurityGroupCommand) Get() *EditSecurityGroupCommand {
	return v.value
}

func (v *NullableEditSecurityGroupCommand) Set(val *EditSecurityGroupCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableEditSecurityGroupCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableEditSecurityGroupCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditSecurityGroupCommand(val *EditSecurityGroupCommand) *NullableEditSecurityGroupCommand {
	return &NullableEditSecurityGroupCommand{value: val, isSet: true}
}

func (v NullableEditSecurityGroupCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditSecurityGroupCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


