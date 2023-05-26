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

// checks if the ProxmoxNetworkListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProxmoxNetworkListDto{}

// ProxmoxNetworkListDto struct for ProxmoxNetworkListDto
type ProxmoxNetworkListDto struct {
	Bridge NullableString `json:"bridge,omitempty"`
	Gateway NullableString `json:"gateway,omitempty"`
	IpAddress NullableString `json:"ipAddress,omitempty"`
	NetMask *int32 `json:"netMask,omitempty"`
	IsPrivate *bool `json:"isPrivate,omitempty"`
}

// NewProxmoxNetworkListDto instantiates a new ProxmoxNetworkListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxmoxNetworkListDto() *ProxmoxNetworkListDto {
	this := ProxmoxNetworkListDto{}
	return &this
}

// NewProxmoxNetworkListDtoWithDefaults instantiates a new ProxmoxNetworkListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxmoxNetworkListDtoWithDefaults() *ProxmoxNetworkListDto {
	this := ProxmoxNetworkListDto{}
	return &this
}

// GetBridge returns the Bridge field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxmoxNetworkListDto) GetBridge() string {
	if o == nil || IsNil(o.Bridge.Get()) {
		var ret string
		return ret
	}
	return *o.Bridge.Get()
}

// GetBridgeOk returns a tuple with the Bridge field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxmoxNetworkListDto) GetBridgeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bridge.Get(), o.Bridge.IsSet()
}

// HasBridge returns a boolean if a field has been set.
func (o *ProxmoxNetworkListDto) HasBridge() bool {
	if o != nil && o.Bridge.IsSet() {
		return true
	}

	return false
}

// SetBridge gets a reference to the given NullableString and assigns it to the Bridge field.
func (o *ProxmoxNetworkListDto) SetBridge(v string) {
	o.Bridge.Set(&v)
}
// SetBridgeNil sets the value for Bridge to be an explicit nil
func (o *ProxmoxNetworkListDto) SetBridgeNil() {
	o.Bridge.Set(nil)
}

// UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
func (o *ProxmoxNetworkListDto) UnsetBridge() {
	o.Bridge.Unset()
}

// GetGateway returns the Gateway field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxmoxNetworkListDto) GetGateway() string {
	if o == nil || IsNil(o.Gateway.Get()) {
		var ret string
		return ret
	}
	return *o.Gateway.Get()
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxmoxNetworkListDto) GetGatewayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gateway.Get(), o.Gateway.IsSet()
}

// HasGateway returns a boolean if a field has been set.
func (o *ProxmoxNetworkListDto) HasGateway() bool {
	if o != nil && o.Gateway.IsSet() {
		return true
	}

	return false
}

// SetGateway gets a reference to the given NullableString and assigns it to the Gateway field.
func (o *ProxmoxNetworkListDto) SetGateway(v string) {
	o.Gateway.Set(&v)
}
// SetGatewayNil sets the value for Gateway to be an explicit nil
func (o *ProxmoxNetworkListDto) SetGatewayNil() {
	o.Gateway.Set(nil)
}

// UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
func (o *ProxmoxNetworkListDto) UnsetGateway() {
	o.Gateway.Unset()
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxmoxNetworkListDto) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress.Get()) {
		var ret string
		return ret
	}
	return *o.IpAddress.Get()
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxmoxNetworkListDto) GetIpAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpAddress.Get(), o.IpAddress.IsSet()
}

// HasIpAddress returns a boolean if a field has been set.
func (o *ProxmoxNetworkListDto) HasIpAddress() bool {
	if o != nil && o.IpAddress.IsSet() {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given NullableString and assigns it to the IpAddress field.
func (o *ProxmoxNetworkListDto) SetIpAddress(v string) {
	o.IpAddress.Set(&v)
}
// SetIpAddressNil sets the value for IpAddress to be an explicit nil
func (o *ProxmoxNetworkListDto) SetIpAddressNil() {
	o.IpAddress.Set(nil)
}

// UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
func (o *ProxmoxNetworkListDto) UnsetIpAddress() {
	o.IpAddress.Unset()
}

// GetNetMask returns the NetMask field value if set, zero value otherwise.
func (o *ProxmoxNetworkListDto) GetNetMask() int32 {
	if o == nil || IsNil(o.NetMask) {
		var ret int32
		return ret
	}
	return *o.NetMask
}

// GetNetMaskOk returns a tuple with the NetMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxmoxNetworkListDto) GetNetMaskOk() (*int32, bool) {
	if o == nil || IsNil(o.NetMask) {
		return nil, false
	}
	return o.NetMask, true
}

// HasNetMask returns a boolean if a field has been set.
func (o *ProxmoxNetworkListDto) HasNetMask() bool {
	if o != nil && !IsNil(o.NetMask) {
		return true
	}

	return false
}

// SetNetMask gets a reference to the given int32 and assigns it to the NetMask field.
func (o *ProxmoxNetworkListDto) SetNetMask(v int32) {
	o.NetMask = &v
}

// GetIsPrivate returns the IsPrivate field value if set, zero value otherwise.
func (o *ProxmoxNetworkListDto) GetIsPrivate() bool {
	if o == nil || IsNil(o.IsPrivate) {
		var ret bool
		return ret
	}
	return *o.IsPrivate
}

// GetIsPrivateOk returns a tuple with the IsPrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxmoxNetworkListDto) GetIsPrivateOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrivate) {
		return nil, false
	}
	return o.IsPrivate, true
}

// HasIsPrivate returns a boolean if a field has been set.
func (o *ProxmoxNetworkListDto) HasIsPrivate() bool {
	if o != nil && !IsNil(o.IsPrivate) {
		return true
	}

	return false
}

// SetIsPrivate gets a reference to the given bool and assigns it to the IsPrivate field.
func (o *ProxmoxNetworkListDto) SetIsPrivate(v bool) {
	o.IsPrivate = &v
}

func (o ProxmoxNetworkListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProxmoxNetworkListDto) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.IsPrivate) {
		toSerialize["isPrivate"] = o.IsPrivate
	}
	return toSerialize, nil
}

type NullableProxmoxNetworkListDto struct {
	value *ProxmoxNetworkListDto
	isSet bool
}

func (v NullableProxmoxNetworkListDto) Get() *ProxmoxNetworkListDto {
	return v.value
}

func (v *NullableProxmoxNetworkListDto) Set(val *ProxmoxNetworkListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProxmoxNetworkListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProxmoxNetworkListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxmoxNetworkListDto(val *ProxmoxNetworkListDto) *NullableProxmoxNetworkListDto {
	return &NullableProxmoxNetworkListDto{value: val, isSet: true}
}

func (v NullableProxmoxNetworkListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxmoxNetworkListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


