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

// checks if the CreateProxmoxNetworkDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProxmoxNetworkDto{}

// CreateProxmoxNetworkDto struct for CreateProxmoxNetworkDto
type CreateProxmoxNetworkDto struct {
	Bridge NullableString `json:"bridge,omitempty"`
	Gateway NullableString `json:"gateway,omitempty"`
	IpAddress NullableString `json:"ipAddress,omitempty"`
	NetMask *int32 `json:"netMask,omitempty"`
	BeginAllocationRange NullableString `json:"beginAllocationRange,omitempty"`
	EndAllocationRange NullableString `json:"endAllocationRange,omitempty"`
}

// NewCreateProxmoxNetworkDto instantiates a new CreateProxmoxNetworkDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProxmoxNetworkDto() *CreateProxmoxNetworkDto {
	this := CreateProxmoxNetworkDto{}
	return &this
}

// NewCreateProxmoxNetworkDtoWithDefaults instantiates a new CreateProxmoxNetworkDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProxmoxNetworkDtoWithDefaults() *CreateProxmoxNetworkDto {
	this := CreateProxmoxNetworkDto{}
	return &this
}

// GetBridge returns the Bridge field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateProxmoxNetworkDto) GetBridge() string {
	if o == nil || IsNil(o.Bridge.Get()) {
		var ret string
		return ret
	}
	return *o.Bridge.Get()
}

// GetBridgeOk returns a tuple with the Bridge field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateProxmoxNetworkDto) GetBridgeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bridge.Get(), o.Bridge.IsSet()
}

// HasBridge returns a boolean if a field has been set.
func (o *CreateProxmoxNetworkDto) HasBridge() bool {
	if o != nil && o.Bridge.IsSet() {
		return true
	}

	return false
}

// SetBridge gets a reference to the given NullableString and assigns it to the Bridge field.
func (o *CreateProxmoxNetworkDto) SetBridge(v string) {
	o.Bridge.Set(&v)
}
// SetBridgeNil sets the value for Bridge to be an explicit nil
func (o *CreateProxmoxNetworkDto) SetBridgeNil() {
	o.Bridge.Set(nil)
}

// UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
func (o *CreateProxmoxNetworkDto) UnsetBridge() {
	o.Bridge.Unset()
}

// GetGateway returns the Gateway field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateProxmoxNetworkDto) GetGateway() string {
	if o == nil || IsNil(o.Gateway.Get()) {
		var ret string
		return ret
	}
	return *o.Gateway.Get()
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateProxmoxNetworkDto) GetGatewayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gateway.Get(), o.Gateway.IsSet()
}

// HasGateway returns a boolean if a field has been set.
func (o *CreateProxmoxNetworkDto) HasGateway() bool {
	if o != nil && o.Gateway.IsSet() {
		return true
	}

	return false
}

// SetGateway gets a reference to the given NullableString and assigns it to the Gateway field.
func (o *CreateProxmoxNetworkDto) SetGateway(v string) {
	o.Gateway.Set(&v)
}
// SetGatewayNil sets the value for Gateway to be an explicit nil
func (o *CreateProxmoxNetworkDto) SetGatewayNil() {
	o.Gateway.Set(nil)
}

// UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
func (o *CreateProxmoxNetworkDto) UnsetGateway() {
	o.Gateway.Unset()
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateProxmoxNetworkDto) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress.Get()) {
		var ret string
		return ret
	}
	return *o.IpAddress.Get()
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateProxmoxNetworkDto) GetIpAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpAddress.Get(), o.IpAddress.IsSet()
}

// HasIpAddress returns a boolean if a field has been set.
func (o *CreateProxmoxNetworkDto) HasIpAddress() bool {
	if o != nil && o.IpAddress.IsSet() {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given NullableString and assigns it to the IpAddress field.
func (o *CreateProxmoxNetworkDto) SetIpAddress(v string) {
	o.IpAddress.Set(&v)
}
// SetIpAddressNil sets the value for IpAddress to be an explicit nil
func (o *CreateProxmoxNetworkDto) SetIpAddressNil() {
	o.IpAddress.Set(nil)
}

// UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
func (o *CreateProxmoxNetworkDto) UnsetIpAddress() {
	o.IpAddress.Unset()
}

// GetNetMask returns the NetMask field value if set, zero value otherwise.
func (o *CreateProxmoxNetworkDto) GetNetMask() int32 {
	if o == nil || IsNil(o.NetMask) {
		var ret int32
		return ret
	}
	return *o.NetMask
}

// GetNetMaskOk returns a tuple with the NetMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProxmoxNetworkDto) GetNetMaskOk() (*int32, bool) {
	if o == nil || IsNil(o.NetMask) {
		return nil, false
	}
	return o.NetMask, true
}

// HasNetMask returns a boolean if a field has been set.
func (o *CreateProxmoxNetworkDto) HasNetMask() bool {
	if o != nil && !IsNil(o.NetMask) {
		return true
	}

	return false
}

// SetNetMask gets a reference to the given int32 and assigns it to the NetMask field.
func (o *CreateProxmoxNetworkDto) SetNetMask(v int32) {
	o.NetMask = &v
}

// GetBeginAllocationRange returns the BeginAllocationRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateProxmoxNetworkDto) GetBeginAllocationRange() string {
	if o == nil || IsNil(o.BeginAllocationRange.Get()) {
		var ret string
		return ret
	}
	return *o.BeginAllocationRange.Get()
}

// GetBeginAllocationRangeOk returns a tuple with the BeginAllocationRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateProxmoxNetworkDto) GetBeginAllocationRangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BeginAllocationRange.Get(), o.BeginAllocationRange.IsSet()
}

// HasBeginAllocationRange returns a boolean if a field has been set.
func (o *CreateProxmoxNetworkDto) HasBeginAllocationRange() bool {
	if o != nil && o.BeginAllocationRange.IsSet() {
		return true
	}

	return false
}

// SetBeginAllocationRange gets a reference to the given NullableString and assigns it to the BeginAllocationRange field.
func (o *CreateProxmoxNetworkDto) SetBeginAllocationRange(v string) {
	o.BeginAllocationRange.Set(&v)
}
// SetBeginAllocationRangeNil sets the value for BeginAllocationRange to be an explicit nil
func (o *CreateProxmoxNetworkDto) SetBeginAllocationRangeNil() {
	o.BeginAllocationRange.Set(nil)
}

// UnsetBeginAllocationRange ensures that no value is present for BeginAllocationRange, not even an explicit nil
func (o *CreateProxmoxNetworkDto) UnsetBeginAllocationRange() {
	o.BeginAllocationRange.Unset()
}

// GetEndAllocationRange returns the EndAllocationRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateProxmoxNetworkDto) GetEndAllocationRange() string {
	if o == nil || IsNil(o.EndAllocationRange.Get()) {
		var ret string
		return ret
	}
	return *o.EndAllocationRange.Get()
}

// GetEndAllocationRangeOk returns a tuple with the EndAllocationRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateProxmoxNetworkDto) GetEndAllocationRangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndAllocationRange.Get(), o.EndAllocationRange.IsSet()
}

// HasEndAllocationRange returns a boolean if a field has been set.
func (o *CreateProxmoxNetworkDto) HasEndAllocationRange() bool {
	if o != nil && o.EndAllocationRange.IsSet() {
		return true
	}

	return false
}

// SetEndAllocationRange gets a reference to the given NullableString and assigns it to the EndAllocationRange field.
func (o *CreateProxmoxNetworkDto) SetEndAllocationRange(v string) {
	o.EndAllocationRange.Set(&v)
}
// SetEndAllocationRangeNil sets the value for EndAllocationRange to be an explicit nil
func (o *CreateProxmoxNetworkDto) SetEndAllocationRangeNil() {
	o.EndAllocationRange.Set(nil)
}

// UnsetEndAllocationRange ensures that no value is present for EndAllocationRange, not even an explicit nil
func (o *CreateProxmoxNetworkDto) UnsetEndAllocationRange() {
	o.EndAllocationRange.Unset()
}

func (o CreateProxmoxNetworkDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateProxmoxNetworkDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Bridge.IsSet() {
		toSerialize["bridge"] = o.Bridge.Get()
	}
	if o.Gateway.IsSet() {
		toSerialize["gateway"] = o.Gateway.Get()
	}
	if o.IpAddress.IsSet() {
		toSerialize["ipAddress"] = o.IpAddress.Get()
	}
	if !IsNil(o.NetMask) {
		toSerialize["netMask"] = o.NetMask
	}
	if o.BeginAllocationRange.IsSet() {
		toSerialize["beginAllocationRange"] = o.BeginAllocationRange.Get()
	}
	if o.EndAllocationRange.IsSet() {
		toSerialize["endAllocationRange"] = o.EndAllocationRange.Get()
	}
	return toSerialize, nil
}

type NullableCreateProxmoxNetworkDto struct {
	value *CreateProxmoxNetworkDto
	isSet bool
}

func (v NullableCreateProxmoxNetworkDto) Get() *CreateProxmoxNetworkDto {
	return v.value
}

func (v *NullableCreateProxmoxNetworkDto) Set(val *CreateProxmoxNetworkDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProxmoxNetworkDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProxmoxNetworkDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProxmoxNetworkDto(val *CreateProxmoxNetworkDto) *NullableCreateProxmoxNetworkDto {
	return &NullableCreateProxmoxNetworkDto{value: val, isSet: true}
}

func (v NullableCreateProxmoxNetworkDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProxmoxNetworkDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


