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

// checks if the UpdateQuotaCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateQuotaCommand{}

// UpdateQuotaCommand struct for UpdateQuotaCommand
type UpdateQuotaCommand struct {
	QuotaId *int32 `json:"quotaId,omitempty"`
	ServerCpu *int64 `json:"serverCpu,omitempty"`
	ServerRam *int64 `json:"serverRam,omitempty"`
	ServerDiskSize *int64 `json:"serverDiskSize,omitempty"`
	VmCpu *int64 `json:"vmCpu,omitempty"`
	VmRam *int64 `json:"vmRam,omitempty"`
	VmVolumeSize *int64 `json:"vmVolumeSize,omitempty"`
}

// NewUpdateQuotaCommand instantiates a new UpdateQuotaCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateQuotaCommand() *UpdateQuotaCommand {
	this := UpdateQuotaCommand{}
	return &this
}

// NewUpdateQuotaCommandWithDefaults instantiates a new UpdateQuotaCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateQuotaCommandWithDefaults() *UpdateQuotaCommand {
	this := UpdateQuotaCommand{}
	return &this
}

// GetQuotaId returns the QuotaId field value if set, zero value otherwise.
func (o *UpdateQuotaCommand) GetQuotaId() int32 {
	if o == nil || IsNil(o.QuotaId) {
		var ret int32
		return ret
	}
	return *o.QuotaId
}

// GetQuotaIdOk returns a tuple with the QuotaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateQuotaCommand) GetQuotaIdOk() (*int32, bool) {
	if o == nil || IsNil(o.QuotaId) {
		return nil, false
	}
	return o.QuotaId, true
}

// HasQuotaId returns a boolean if a field has been set.
func (o *UpdateQuotaCommand) HasQuotaId() bool {
	if o != nil && !IsNil(o.QuotaId) {
		return true
	}

	return false
}

// SetQuotaId gets a reference to the given int32 and assigns it to the QuotaId field.
func (o *UpdateQuotaCommand) SetQuotaId(v int32) {
	o.QuotaId = &v
}

// GetServerCpu returns the ServerCpu field value if set, zero value otherwise.
func (o *UpdateQuotaCommand) GetServerCpu() int64 {
	if o == nil || IsNil(o.ServerCpu) {
		var ret int64
		return ret
	}
	return *o.ServerCpu
}

// GetServerCpuOk returns a tuple with the ServerCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateQuotaCommand) GetServerCpuOk() (*int64, bool) {
	if o == nil || IsNil(o.ServerCpu) {
		return nil, false
	}
	return o.ServerCpu, true
}

// HasServerCpu returns a boolean if a field has been set.
func (o *UpdateQuotaCommand) HasServerCpu() bool {
	if o != nil && !IsNil(o.ServerCpu) {
		return true
	}

	return false
}

// SetServerCpu gets a reference to the given int64 and assigns it to the ServerCpu field.
func (o *UpdateQuotaCommand) SetServerCpu(v int64) {
	o.ServerCpu = &v
}

// GetServerRam returns the ServerRam field value if set, zero value otherwise.
func (o *UpdateQuotaCommand) GetServerRam() int64 {
	if o == nil || IsNil(o.ServerRam) {
		var ret int64
		return ret
	}
	return *o.ServerRam
}

// GetServerRamOk returns a tuple with the ServerRam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateQuotaCommand) GetServerRamOk() (*int64, bool) {
	if o == nil || IsNil(o.ServerRam) {
		return nil, false
	}
	return o.ServerRam, true
}

// HasServerRam returns a boolean if a field has been set.
func (o *UpdateQuotaCommand) HasServerRam() bool {
	if o != nil && !IsNil(o.ServerRam) {
		return true
	}

	return false
}

// SetServerRam gets a reference to the given int64 and assigns it to the ServerRam field.
func (o *UpdateQuotaCommand) SetServerRam(v int64) {
	o.ServerRam = &v
}

// GetServerDiskSize returns the ServerDiskSize field value if set, zero value otherwise.
func (o *UpdateQuotaCommand) GetServerDiskSize() int64 {
	if o == nil || IsNil(o.ServerDiskSize) {
		var ret int64
		return ret
	}
	return *o.ServerDiskSize
}

// GetServerDiskSizeOk returns a tuple with the ServerDiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateQuotaCommand) GetServerDiskSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.ServerDiskSize) {
		return nil, false
	}
	return o.ServerDiskSize, true
}

// HasServerDiskSize returns a boolean if a field has been set.
func (o *UpdateQuotaCommand) HasServerDiskSize() bool {
	if o != nil && !IsNil(o.ServerDiskSize) {
		return true
	}

	return false
}

// SetServerDiskSize gets a reference to the given int64 and assigns it to the ServerDiskSize field.
func (o *UpdateQuotaCommand) SetServerDiskSize(v int64) {
	o.ServerDiskSize = &v
}

// GetVmCpu returns the VmCpu field value if set, zero value otherwise.
func (o *UpdateQuotaCommand) GetVmCpu() int64 {
	if o == nil || IsNil(o.VmCpu) {
		var ret int64
		return ret
	}
	return *o.VmCpu
}

// GetVmCpuOk returns a tuple with the VmCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateQuotaCommand) GetVmCpuOk() (*int64, bool) {
	if o == nil || IsNil(o.VmCpu) {
		return nil, false
	}
	return o.VmCpu, true
}

// HasVmCpu returns a boolean if a field has been set.
func (o *UpdateQuotaCommand) HasVmCpu() bool {
	if o != nil && !IsNil(o.VmCpu) {
		return true
	}

	return false
}

// SetVmCpu gets a reference to the given int64 and assigns it to the VmCpu field.
func (o *UpdateQuotaCommand) SetVmCpu(v int64) {
	o.VmCpu = &v
}

// GetVmRam returns the VmRam field value if set, zero value otherwise.
func (o *UpdateQuotaCommand) GetVmRam() int64 {
	if o == nil || IsNil(o.VmRam) {
		var ret int64
		return ret
	}
	return *o.VmRam
}

// GetVmRamOk returns a tuple with the VmRam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateQuotaCommand) GetVmRamOk() (*int64, bool) {
	if o == nil || IsNil(o.VmRam) {
		return nil, false
	}
	return o.VmRam, true
}

// HasVmRam returns a boolean if a field has been set.
func (o *UpdateQuotaCommand) HasVmRam() bool {
	if o != nil && !IsNil(o.VmRam) {
		return true
	}

	return false
}

// SetVmRam gets a reference to the given int64 and assigns it to the VmRam field.
func (o *UpdateQuotaCommand) SetVmRam(v int64) {
	o.VmRam = &v
}

// GetVmVolumeSize returns the VmVolumeSize field value if set, zero value otherwise.
func (o *UpdateQuotaCommand) GetVmVolumeSize() int64 {
	if o == nil || IsNil(o.VmVolumeSize) {
		var ret int64
		return ret
	}
	return *o.VmVolumeSize
}

// GetVmVolumeSizeOk returns a tuple with the VmVolumeSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateQuotaCommand) GetVmVolumeSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.VmVolumeSize) {
		return nil, false
	}
	return o.VmVolumeSize, true
}

// HasVmVolumeSize returns a boolean if a field has been set.
func (o *UpdateQuotaCommand) HasVmVolumeSize() bool {
	if o != nil && !IsNil(o.VmVolumeSize) {
		return true
	}

	return false
}

// SetVmVolumeSize gets a reference to the given int64 and assigns it to the VmVolumeSize field.
func (o *UpdateQuotaCommand) SetVmVolumeSize(v int64) {
	o.VmVolumeSize = &v
}

func (o UpdateQuotaCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateQuotaCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QuotaId) {
		toSerialize["quotaId"] = o.QuotaId
	}
	if !IsNil(o.ServerCpu) {
		toSerialize["serverCpu"] = o.ServerCpu
	}
	if !IsNil(o.ServerRam) {
		toSerialize["serverRam"] = o.ServerRam
	}
	if !IsNil(o.ServerDiskSize) {
		toSerialize["serverDiskSize"] = o.ServerDiskSize
	}
	if !IsNil(o.VmCpu) {
		toSerialize["vmCpu"] = o.VmCpu
	}
	if !IsNil(o.VmRam) {
		toSerialize["vmRam"] = o.VmRam
	}
	if !IsNil(o.VmVolumeSize) {
		toSerialize["vmVolumeSize"] = o.VmVolumeSize
	}
	return toSerialize, nil
}

type NullableUpdateQuotaCommand struct {
	value *UpdateQuotaCommand
	isSet bool
}

func (v NullableUpdateQuotaCommand) Get() *UpdateQuotaCommand {
	return v.value
}

func (v *NullableUpdateQuotaCommand) Set(val *UpdateQuotaCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateQuotaCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateQuotaCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateQuotaCommand(val *UpdateQuotaCommand) *NullableUpdateQuotaCommand {
	return &NullableUpdateQuotaCommand{value: val, isSet: true}
}

func (v NullableUpdateQuotaCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateQuotaCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


