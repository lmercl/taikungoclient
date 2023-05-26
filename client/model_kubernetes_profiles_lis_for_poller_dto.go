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

// checks if the KubernetesProfilesLisForPollerDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesProfilesLisForPollerDto{}

// KubernetesProfilesLisForPollerDto struct for KubernetesProfilesLisForPollerDto
type KubernetesProfilesLisForPollerDto struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	OrganizationId NullableInt32 `json:"organizationId,omitempty"`
	OrganizationName NullableString `json:"organizationName,omitempty"`
	Cni NullableString `json:"cni,omitempty"`
	OctaviaEnabled *bool `json:"octaviaEnabled,omitempty"`
	ExposeNodePortOnBastion *bool `json:"exposeNodePortOnBastion,omitempty"`
	IsLocked *bool `json:"isLocked,omitempty"`
	TaikunLBEnabled *bool `json:"taikunLBEnabled,omitempty"`
	AllowSchedulingOnMaster *bool `json:"allowSchedulingOnMaster,omitempty"`
	UniqueClusterName *bool `json:"uniqueClusterName,omitempty"`
	ProxmoxStorage NullableString `json:"proxmoxStorage,omitempty"`
}

// NewKubernetesProfilesLisForPollerDto instantiates a new KubernetesProfilesLisForPollerDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesProfilesLisForPollerDto() *KubernetesProfilesLisForPollerDto {
	this := KubernetesProfilesLisForPollerDto{}
	return &this
}

// NewKubernetesProfilesLisForPollerDtoWithDefaults instantiates a new KubernetesProfilesLisForPollerDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesProfilesLisForPollerDtoWithDefaults() *KubernetesProfilesLisForPollerDto {
	this := KubernetesProfilesLisForPollerDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KubernetesProfilesLisForPollerDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesProfilesLisForPollerDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KubernetesProfilesLisForPollerDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KubernetesProfilesLisForPollerDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesProfilesLisForPollerDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesProfilesLisForPollerDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *KubernetesProfilesLisForPollerDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *KubernetesProfilesLisForPollerDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *KubernetesProfilesLisForPollerDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *KubernetesProfilesLisForPollerDto) UnsetName() {
	o.Name.Unset()
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesProfilesLisForPollerDto) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret int32
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesProfilesLisForPollerDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *KubernetesProfilesLisForPollerDto) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableInt32 and assigns it to the OrganizationId field.
func (o *KubernetesProfilesLisForPollerDto) SetOrganizationId(v int32) {
	o.OrganizationId.Set(&v)
}
// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *KubernetesProfilesLisForPollerDto) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *KubernetesProfilesLisForPollerDto) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesProfilesLisForPollerDto) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationName.Get()
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesProfilesLisForPollerDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationName.Get(), o.OrganizationName.IsSet()
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *KubernetesProfilesLisForPollerDto) HasOrganizationName() bool {
	if o != nil && o.OrganizationName.IsSet() {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given NullableString and assigns it to the OrganizationName field.
func (o *KubernetesProfilesLisForPollerDto) SetOrganizationName(v string) {
	o.OrganizationName.Set(&v)
}
// SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil
func (o *KubernetesProfilesLisForPollerDto) SetOrganizationNameNil() {
	o.OrganizationName.Set(nil)
}

// UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
func (o *KubernetesProfilesLisForPollerDto) UnsetOrganizationName() {
	o.OrganizationName.Unset()
}

// GetCni returns the Cni field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesProfilesLisForPollerDto) GetCni() string {
	if o == nil || IsNil(o.Cni.Get()) {
		var ret string
		return ret
	}
	return *o.Cni.Get()
}

// GetCniOk returns a tuple with the Cni field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesProfilesLisForPollerDto) GetCniOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cni.Get(), o.Cni.IsSet()
}

// HasCni returns a boolean if a field has been set.
func (o *KubernetesProfilesLisForPollerDto) HasCni() bool {
	if o != nil && o.Cni.IsSet() {
		return true
	}

	return false
}

// SetCni gets a reference to the given NullableString and assigns it to the Cni field.
func (o *KubernetesProfilesLisForPollerDto) SetCni(v string) {
	o.Cni.Set(&v)
}
// SetCniNil sets the value for Cni to be an explicit nil
func (o *KubernetesProfilesLisForPollerDto) SetCniNil() {
	o.Cni.Set(nil)
}

// UnsetCni ensures that no value is present for Cni, not even an explicit nil
func (o *KubernetesProfilesLisForPollerDto) UnsetCni() {
	o.Cni.Unset()
}

// GetOctaviaEnabled returns the OctaviaEnabled field value if set, zero value otherwise.
func (o *KubernetesProfilesLisForPollerDto) GetOctaviaEnabled() bool {
	if o == nil || IsNil(o.OctaviaEnabled) {
		var ret bool
		return ret
	}
	return *o.OctaviaEnabled
}

// GetOctaviaEnabledOk returns a tuple with the OctaviaEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesProfilesLisForPollerDto) GetOctaviaEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.OctaviaEnabled) {
		return nil, false
	}
	return o.OctaviaEnabled, true
}

// HasOctaviaEnabled returns a boolean if a field has been set.
func (o *KubernetesProfilesLisForPollerDto) HasOctaviaEnabled() bool {
	if o != nil && !IsNil(o.OctaviaEnabled) {
		return true
	}

	return false
}

// SetOctaviaEnabled gets a reference to the given bool and assigns it to the OctaviaEnabled field.
func (o *KubernetesProfilesLisForPollerDto) SetOctaviaEnabled(v bool) {
	o.OctaviaEnabled = &v
}

// GetExposeNodePortOnBastion returns the ExposeNodePortOnBastion field value if set, zero value otherwise.
func (o *KubernetesProfilesLisForPollerDto) GetExposeNodePortOnBastion() bool {
	if o == nil || IsNil(o.ExposeNodePortOnBastion) {
		var ret bool
		return ret
	}
	return *o.ExposeNodePortOnBastion
}

// GetExposeNodePortOnBastionOk returns a tuple with the ExposeNodePortOnBastion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesProfilesLisForPollerDto) GetExposeNodePortOnBastionOk() (*bool, bool) {
	if o == nil || IsNil(o.ExposeNodePortOnBastion) {
		return nil, false
	}
	return o.ExposeNodePortOnBastion, true
}

// HasExposeNodePortOnBastion returns a boolean if a field has been set.
func (o *KubernetesProfilesLisForPollerDto) HasExposeNodePortOnBastion() bool {
	if o != nil && !IsNil(o.ExposeNodePortOnBastion) {
		return true
	}

	return false
}

// SetExposeNodePortOnBastion gets a reference to the given bool and assigns it to the ExposeNodePortOnBastion field.
func (o *KubernetesProfilesLisForPollerDto) SetExposeNodePortOnBastion(v bool) {
	o.ExposeNodePortOnBastion = &v
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *KubernetesProfilesLisForPollerDto) GetIsLocked() bool {
	if o == nil || IsNil(o.IsLocked) {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesProfilesLisForPollerDto) GetIsLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocked) {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *KubernetesProfilesLisForPollerDto) HasIsLocked() bool {
	if o != nil && !IsNil(o.IsLocked) {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *KubernetesProfilesLisForPollerDto) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetTaikunLBEnabled returns the TaikunLBEnabled field value if set, zero value otherwise.
func (o *KubernetesProfilesLisForPollerDto) GetTaikunLBEnabled() bool {
	if o == nil || IsNil(o.TaikunLBEnabled) {
		var ret bool
		return ret
	}
	return *o.TaikunLBEnabled
}

// GetTaikunLBEnabledOk returns a tuple with the TaikunLBEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesProfilesLisForPollerDto) GetTaikunLBEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.TaikunLBEnabled) {
		return nil, false
	}
	return o.TaikunLBEnabled, true
}

// HasTaikunLBEnabled returns a boolean if a field has been set.
func (o *KubernetesProfilesLisForPollerDto) HasTaikunLBEnabled() bool {
	if o != nil && !IsNil(o.TaikunLBEnabled) {
		return true
	}

	return false
}

// SetTaikunLBEnabled gets a reference to the given bool and assigns it to the TaikunLBEnabled field.
func (o *KubernetesProfilesLisForPollerDto) SetTaikunLBEnabled(v bool) {
	o.TaikunLBEnabled = &v
}

// GetAllowSchedulingOnMaster returns the AllowSchedulingOnMaster field value if set, zero value otherwise.
func (o *KubernetesProfilesLisForPollerDto) GetAllowSchedulingOnMaster() bool {
	if o == nil || IsNil(o.AllowSchedulingOnMaster) {
		var ret bool
		return ret
	}
	return *o.AllowSchedulingOnMaster
}

// GetAllowSchedulingOnMasterOk returns a tuple with the AllowSchedulingOnMaster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesProfilesLisForPollerDto) GetAllowSchedulingOnMasterOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowSchedulingOnMaster) {
		return nil, false
	}
	return o.AllowSchedulingOnMaster, true
}

// HasAllowSchedulingOnMaster returns a boolean if a field has been set.
func (o *KubernetesProfilesLisForPollerDto) HasAllowSchedulingOnMaster() bool {
	if o != nil && !IsNil(o.AllowSchedulingOnMaster) {
		return true
	}

	return false
}

// SetAllowSchedulingOnMaster gets a reference to the given bool and assigns it to the AllowSchedulingOnMaster field.
func (o *KubernetesProfilesLisForPollerDto) SetAllowSchedulingOnMaster(v bool) {
	o.AllowSchedulingOnMaster = &v
}

// GetUniqueClusterName returns the UniqueClusterName field value if set, zero value otherwise.
func (o *KubernetesProfilesLisForPollerDto) GetUniqueClusterName() bool {
	if o == nil || IsNil(o.UniqueClusterName) {
		var ret bool
		return ret
	}
	return *o.UniqueClusterName
}

// GetUniqueClusterNameOk returns a tuple with the UniqueClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesProfilesLisForPollerDto) GetUniqueClusterNameOk() (*bool, bool) {
	if o == nil || IsNil(o.UniqueClusterName) {
		return nil, false
	}
	return o.UniqueClusterName, true
}

// HasUniqueClusterName returns a boolean if a field has been set.
func (o *KubernetesProfilesLisForPollerDto) HasUniqueClusterName() bool {
	if o != nil && !IsNil(o.UniqueClusterName) {
		return true
	}

	return false
}

// SetUniqueClusterName gets a reference to the given bool and assigns it to the UniqueClusterName field.
func (o *KubernetesProfilesLisForPollerDto) SetUniqueClusterName(v bool) {
	o.UniqueClusterName = &v
}

// GetProxmoxStorage returns the ProxmoxStorage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesProfilesLisForPollerDto) GetProxmoxStorage() string {
	if o == nil || IsNil(o.ProxmoxStorage.Get()) {
		var ret string
		return ret
	}
	return *o.ProxmoxStorage.Get()
}

// GetProxmoxStorageOk returns a tuple with the ProxmoxStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesProfilesLisForPollerDto) GetProxmoxStorageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProxmoxStorage.Get(), o.ProxmoxStorage.IsSet()
}

// HasProxmoxStorage returns a boolean if a field has been set.
func (o *KubernetesProfilesLisForPollerDto) HasProxmoxStorage() bool {
	if o != nil && o.ProxmoxStorage.IsSet() {
		return true
	}

	return false
}

// SetProxmoxStorage gets a reference to the given NullableString and assigns it to the ProxmoxStorage field.
func (o *KubernetesProfilesLisForPollerDto) SetProxmoxStorage(v string) {
	o.ProxmoxStorage.Set(&v)
}
// SetProxmoxStorageNil sets the value for ProxmoxStorage to be an explicit nil
func (o *KubernetesProfilesLisForPollerDto) SetProxmoxStorageNil() {
	o.ProxmoxStorage.Set(nil)
}

// UnsetProxmoxStorage ensures that no value is present for ProxmoxStorage, not even an explicit nil
func (o *KubernetesProfilesLisForPollerDto) UnsetProxmoxStorage() {
	o.ProxmoxStorage.Unset()
}

func (o KubernetesProfilesLisForPollerDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesProfilesLisForPollerDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.OrganizationId.IsSet() {
		toSerialize["organizationId"] = o.OrganizationId.Get()
	}
	if o.OrganizationName.IsSet() {
		toSerialize["organizationName"] = o.OrganizationName.Get()
	}
	if o.Cni.IsSet() {
		toSerialize["cni"] = o.Cni.Get()
	}
	if !IsNil(o.OctaviaEnabled) {
		toSerialize["octaviaEnabled"] = o.OctaviaEnabled
	}
	if !IsNil(o.ExposeNodePortOnBastion) {
		toSerialize["exposeNodePortOnBastion"] = o.ExposeNodePortOnBastion
	}
	if !IsNil(o.IsLocked) {
		toSerialize["isLocked"] = o.IsLocked
	}
	if !IsNil(o.TaikunLBEnabled) {
		toSerialize["taikunLBEnabled"] = o.TaikunLBEnabled
	}
	if !IsNil(o.AllowSchedulingOnMaster) {
		toSerialize["allowSchedulingOnMaster"] = o.AllowSchedulingOnMaster
	}
	if !IsNil(o.UniqueClusterName) {
		toSerialize["uniqueClusterName"] = o.UniqueClusterName
	}
	if o.ProxmoxStorage.IsSet() {
		toSerialize["proxmoxStorage"] = o.ProxmoxStorage.Get()
	}
	return toSerialize, nil
}

type NullableKubernetesProfilesLisForPollerDto struct {
	value *KubernetesProfilesLisForPollerDto
	isSet bool
}

func (v NullableKubernetesProfilesLisForPollerDto) Get() *KubernetesProfilesLisForPollerDto {
	return v.value
}

func (v *NullableKubernetesProfilesLisForPollerDto) Set(val *KubernetesProfilesLisForPollerDto) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesProfilesLisForPollerDto) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesProfilesLisForPollerDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesProfilesLisForPollerDto(val *KubernetesProfilesLisForPollerDto) *NullableKubernetesProfilesLisForPollerDto {
	return &NullableKubernetesProfilesLisForPollerDto{value: val, isSet: true}
}

func (v NullableKubernetesProfilesLisForPollerDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesProfilesLisForPollerDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


