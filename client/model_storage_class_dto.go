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

// checks if the StorageClassDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageClassDto{}

// StorageClassDto struct for StorageClassDto
type StorageClassDto struct {
	MetadataName NullableString `json:"metadataName,omitempty"`
	Age NullableString `json:"age,omitempty"`
	Provisioner NullableString `json:"provisioner,omitempty"`
	ReclaimPolicy NullableString `json:"reclaimPolicy,omitempty"`
	VolumeBindingMode NullableString `json:"volumeBindingMode,omitempty"`
	AllowVolumeExpansion NullableBool `json:"allowVolumeExpansion,omitempty"`
}

// NewStorageClassDto instantiates a new StorageClassDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageClassDto() *StorageClassDto {
	this := StorageClassDto{}
	return &this
}

// NewStorageClassDtoWithDefaults instantiates a new StorageClassDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageClassDtoWithDefaults() *StorageClassDto {
	this := StorageClassDto{}
	return &this
}

// GetMetadataName returns the MetadataName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageClassDto) GetMetadataName() string {
	if o == nil || IsNil(o.MetadataName.Get()) {
		var ret string
		return ret
	}
	return *o.MetadataName.Get()
}

// GetMetadataNameOk returns a tuple with the MetadataName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageClassDto) GetMetadataNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetadataName.Get(), o.MetadataName.IsSet()
}

// HasMetadataName returns a boolean if a field has been set.
func (o *StorageClassDto) HasMetadataName() bool {
	if o != nil && o.MetadataName.IsSet() {
		return true
	}

	return false
}

// SetMetadataName gets a reference to the given NullableString and assigns it to the MetadataName field.
func (o *StorageClassDto) SetMetadataName(v string) {
	o.MetadataName.Set(&v)
}
// SetMetadataNameNil sets the value for MetadataName to be an explicit nil
func (o *StorageClassDto) SetMetadataNameNil() {
	o.MetadataName.Set(nil)
}

// UnsetMetadataName ensures that no value is present for MetadataName, not even an explicit nil
func (o *StorageClassDto) UnsetMetadataName() {
	o.MetadataName.Unset()
}

// GetAge returns the Age field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageClassDto) GetAge() string {
	if o == nil || IsNil(o.Age.Get()) {
		var ret string
		return ret
	}
	return *o.Age.Get()
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageClassDto) GetAgeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Age.Get(), o.Age.IsSet()
}

// HasAge returns a boolean if a field has been set.
func (o *StorageClassDto) HasAge() bool {
	if o != nil && o.Age.IsSet() {
		return true
	}

	return false
}

// SetAge gets a reference to the given NullableString and assigns it to the Age field.
func (o *StorageClassDto) SetAge(v string) {
	o.Age.Set(&v)
}
// SetAgeNil sets the value for Age to be an explicit nil
func (o *StorageClassDto) SetAgeNil() {
	o.Age.Set(nil)
}

// UnsetAge ensures that no value is present for Age, not even an explicit nil
func (o *StorageClassDto) UnsetAge() {
	o.Age.Unset()
}

// GetProvisioner returns the Provisioner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageClassDto) GetProvisioner() string {
	if o == nil || IsNil(o.Provisioner.Get()) {
		var ret string
		return ret
	}
	return *o.Provisioner.Get()
}

// GetProvisionerOk returns a tuple with the Provisioner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageClassDto) GetProvisionerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Provisioner.Get(), o.Provisioner.IsSet()
}

// HasProvisioner returns a boolean if a field has been set.
func (o *StorageClassDto) HasProvisioner() bool {
	if o != nil && o.Provisioner.IsSet() {
		return true
	}

	return false
}

// SetProvisioner gets a reference to the given NullableString and assigns it to the Provisioner field.
func (o *StorageClassDto) SetProvisioner(v string) {
	o.Provisioner.Set(&v)
}
// SetProvisionerNil sets the value for Provisioner to be an explicit nil
func (o *StorageClassDto) SetProvisionerNil() {
	o.Provisioner.Set(nil)
}

// UnsetProvisioner ensures that no value is present for Provisioner, not even an explicit nil
func (o *StorageClassDto) UnsetProvisioner() {
	o.Provisioner.Unset()
}

// GetReclaimPolicy returns the ReclaimPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageClassDto) GetReclaimPolicy() string {
	if o == nil || IsNil(o.ReclaimPolicy.Get()) {
		var ret string
		return ret
	}
	return *o.ReclaimPolicy.Get()
}

// GetReclaimPolicyOk returns a tuple with the ReclaimPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageClassDto) GetReclaimPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReclaimPolicy.Get(), o.ReclaimPolicy.IsSet()
}

// HasReclaimPolicy returns a boolean if a field has been set.
func (o *StorageClassDto) HasReclaimPolicy() bool {
	if o != nil && o.ReclaimPolicy.IsSet() {
		return true
	}

	return false
}

// SetReclaimPolicy gets a reference to the given NullableString and assigns it to the ReclaimPolicy field.
func (o *StorageClassDto) SetReclaimPolicy(v string) {
	o.ReclaimPolicy.Set(&v)
}
// SetReclaimPolicyNil sets the value for ReclaimPolicy to be an explicit nil
func (o *StorageClassDto) SetReclaimPolicyNil() {
	o.ReclaimPolicy.Set(nil)
}

// UnsetReclaimPolicy ensures that no value is present for ReclaimPolicy, not even an explicit nil
func (o *StorageClassDto) UnsetReclaimPolicy() {
	o.ReclaimPolicy.Unset()
}

// GetVolumeBindingMode returns the VolumeBindingMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageClassDto) GetVolumeBindingMode() string {
	if o == nil || IsNil(o.VolumeBindingMode.Get()) {
		var ret string
		return ret
	}
	return *o.VolumeBindingMode.Get()
}

// GetVolumeBindingModeOk returns a tuple with the VolumeBindingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageClassDto) GetVolumeBindingModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VolumeBindingMode.Get(), o.VolumeBindingMode.IsSet()
}

// HasVolumeBindingMode returns a boolean if a field has been set.
func (o *StorageClassDto) HasVolumeBindingMode() bool {
	if o != nil && o.VolumeBindingMode.IsSet() {
		return true
	}

	return false
}

// SetVolumeBindingMode gets a reference to the given NullableString and assigns it to the VolumeBindingMode field.
func (o *StorageClassDto) SetVolumeBindingMode(v string) {
	o.VolumeBindingMode.Set(&v)
}
// SetVolumeBindingModeNil sets the value for VolumeBindingMode to be an explicit nil
func (o *StorageClassDto) SetVolumeBindingModeNil() {
	o.VolumeBindingMode.Set(nil)
}

// UnsetVolumeBindingMode ensures that no value is present for VolumeBindingMode, not even an explicit nil
func (o *StorageClassDto) UnsetVolumeBindingMode() {
	o.VolumeBindingMode.Unset()
}

// GetAllowVolumeExpansion returns the AllowVolumeExpansion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageClassDto) GetAllowVolumeExpansion() bool {
	if o == nil || IsNil(o.AllowVolumeExpansion.Get()) {
		var ret bool
		return ret
	}
	return *o.AllowVolumeExpansion.Get()
}

// GetAllowVolumeExpansionOk returns a tuple with the AllowVolumeExpansion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageClassDto) GetAllowVolumeExpansionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowVolumeExpansion.Get(), o.AllowVolumeExpansion.IsSet()
}

// HasAllowVolumeExpansion returns a boolean if a field has been set.
func (o *StorageClassDto) HasAllowVolumeExpansion() bool {
	if o != nil && o.AllowVolumeExpansion.IsSet() {
		return true
	}

	return false
}

// SetAllowVolumeExpansion gets a reference to the given NullableBool and assigns it to the AllowVolumeExpansion field.
func (o *StorageClassDto) SetAllowVolumeExpansion(v bool) {
	o.AllowVolumeExpansion.Set(&v)
}
// SetAllowVolumeExpansionNil sets the value for AllowVolumeExpansion to be an explicit nil
func (o *StorageClassDto) SetAllowVolumeExpansionNil() {
	o.AllowVolumeExpansion.Set(nil)
}

// UnsetAllowVolumeExpansion ensures that no value is present for AllowVolumeExpansion, not even an explicit nil
func (o *StorageClassDto) UnsetAllowVolumeExpansion() {
	o.AllowVolumeExpansion.Unset()
}

func (o StorageClassDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageClassDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MetadataName.IsSet() {
		toSerialize["metadataName"] = o.MetadataName.Get()
	}
	if o.Age.IsSet() {
		toSerialize["age"] = o.Age.Get()
	}
	if o.Provisioner.IsSet() {
		toSerialize["provisioner"] = o.Provisioner.Get()
	}
	if o.ReclaimPolicy.IsSet() {
		toSerialize["reclaimPolicy"] = o.ReclaimPolicy.Get()
	}
	if o.VolumeBindingMode.IsSet() {
		toSerialize["volumeBindingMode"] = o.VolumeBindingMode.Get()
	}
	if o.AllowVolumeExpansion.IsSet() {
		toSerialize["allowVolumeExpansion"] = o.AllowVolumeExpansion.Get()
	}
	return toSerialize, nil
}

type NullableStorageClassDto struct {
	value *StorageClassDto
	isSet bool
}

func (v NullableStorageClassDto) Get() *StorageClassDto {
	return v.value
}

func (v *NullableStorageClassDto) Set(val *StorageClassDto) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageClassDto) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageClassDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageClassDto(val *StorageClassDto) *NullableStorageClassDto {
	return &NullableStorageClassDto{value: val, isSet: true}
}

func (v NullableStorageClassDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageClassDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


