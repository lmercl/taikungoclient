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

// checks if the CreateKubernetesProfileCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateKubernetesProfileCommand{}

// CreateKubernetesProfileCommand struct for CreateKubernetesProfileCommand
type CreateKubernetesProfileCommand struct {
	Name string `json:"name"`
	OctaviaEnabled *bool `json:"octaviaEnabled,omitempty"`
	ExposeNodePortOnBastion *bool `json:"exposeNodePortOnBastion,omitempty"`
	OrganizationId NullableInt32 `json:"organizationId,omitempty"`
	TaikunLBEnabled *bool `json:"taikunLBEnabled,omitempty"`
	AllowSchedulingOnMaster *bool `json:"allowSchedulingOnMaster,omitempty"`
	UniqueClusterName *bool `json:"uniqueClusterName,omitempty"`
}

// NewCreateKubernetesProfileCommand instantiates a new CreateKubernetesProfileCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateKubernetesProfileCommand(name string) *CreateKubernetesProfileCommand {
	this := CreateKubernetesProfileCommand{}
	this.Name = name
	return &this
}

// NewCreateKubernetesProfileCommandWithDefaults instantiates a new CreateKubernetesProfileCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateKubernetesProfileCommandWithDefaults() *CreateKubernetesProfileCommand {
	this := CreateKubernetesProfileCommand{}
	return &this
}

// GetName returns the Name field value
func (o *CreateKubernetesProfileCommand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateKubernetesProfileCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateKubernetesProfileCommand) SetName(v string) {
	o.Name = v
}

// GetOctaviaEnabled returns the OctaviaEnabled field value if set, zero value otherwise.
func (o *CreateKubernetesProfileCommand) GetOctaviaEnabled() bool {
	if o == nil || IsNil(o.OctaviaEnabled) {
		var ret bool
		return ret
	}
	return *o.OctaviaEnabled
}

// GetOctaviaEnabledOk returns a tuple with the OctaviaEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKubernetesProfileCommand) GetOctaviaEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.OctaviaEnabled) {
		return nil, false
	}
	return o.OctaviaEnabled, true
}

// HasOctaviaEnabled returns a boolean if a field has been set.
func (o *CreateKubernetesProfileCommand) HasOctaviaEnabled() bool {
	if o != nil && !IsNil(o.OctaviaEnabled) {
		return true
	}

	return false
}

// SetOctaviaEnabled gets a reference to the given bool and assigns it to the OctaviaEnabled field.
func (o *CreateKubernetesProfileCommand) SetOctaviaEnabled(v bool) {
	o.OctaviaEnabled = &v
}

// GetExposeNodePortOnBastion returns the ExposeNodePortOnBastion field value if set, zero value otherwise.
func (o *CreateKubernetesProfileCommand) GetExposeNodePortOnBastion() bool {
	if o == nil || IsNil(o.ExposeNodePortOnBastion) {
		var ret bool
		return ret
	}
	return *o.ExposeNodePortOnBastion
}

// GetExposeNodePortOnBastionOk returns a tuple with the ExposeNodePortOnBastion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKubernetesProfileCommand) GetExposeNodePortOnBastionOk() (*bool, bool) {
	if o == nil || IsNil(o.ExposeNodePortOnBastion) {
		return nil, false
	}
	return o.ExposeNodePortOnBastion, true
}

// HasExposeNodePortOnBastion returns a boolean if a field has been set.
func (o *CreateKubernetesProfileCommand) HasExposeNodePortOnBastion() bool {
	if o != nil && !IsNil(o.ExposeNodePortOnBastion) {
		return true
	}

	return false
}

// SetExposeNodePortOnBastion gets a reference to the given bool and assigns it to the ExposeNodePortOnBastion field.
func (o *CreateKubernetesProfileCommand) SetExposeNodePortOnBastion(v bool) {
	o.ExposeNodePortOnBastion = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateKubernetesProfileCommand) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret int32
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateKubernetesProfileCommand) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *CreateKubernetesProfileCommand) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableInt32 and assigns it to the OrganizationId field.
func (o *CreateKubernetesProfileCommand) SetOrganizationId(v int32) {
	o.OrganizationId.Set(&v)
}
// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *CreateKubernetesProfileCommand) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *CreateKubernetesProfileCommand) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

// GetTaikunLBEnabled returns the TaikunLBEnabled field value if set, zero value otherwise.
func (o *CreateKubernetesProfileCommand) GetTaikunLBEnabled() bool {
	if o == nil || IsNil(o.TaikunLBEnabled) {
		var ret bool
		return ret
	}
	return *o.TaikunLBEnabled
}

// GetTaikunLBEnabledOk returns a tuple with the TaikunLBEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKubernetesProfileCommand) GetTaikunLBEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.TaikunLBEnabled) {
		return nil, false
	}
	return o.TaikunLBEnabled, true
}

// HasTaikunLBEnabled returns a boolean if a field has been set.
func (o *CreateKubernetesProfileCommand) HasTaikunLBEnabled() bool {
	if o != nil && !IsNil(o.TaikunLBEnabled) {
		return true
	}

	return false
}

// SetTaikunLBEnabled gets a reference to the given bool and assigns it to the TaikunLBEnabled field.
func (o *CreateKubernetesProfileCommand) SetTaikunLBEnabled(v bool) {
	o.TaikunLBEnabled = &v
}

// GetAllowSchedulingOnMaster returns the AllowSchedulingOnMaster field value if set, zero value otherwise.
func (o *CreateKubernetesProfileCommand) GetAllowSchedulingOnMaster() bool {
	if o == nil || IsNil(o.AllowSchedulingOnMaster) {
		var ret bool
		return ret
	}
	return *o.AllowSchedulingOnMaster
}

// GetAllowSchedulingOnMasterOk returns a tuple with the AllowSchedulingOnMaster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKubernetesProfileCommand) GetAllowSchedulingOnMasterOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowSchedulingOnMaster) {
		return nil, false
	}
	return o.AllowSchedulingOnMaster, true
}

// HasAllowSchedulingOnMaster returns a boolean if a field has been set.
func (o *CreateKubernetesProfileCommand) HasAllowSchedulingOnMaster() bool {
	if o != nil && !IsNil(o.AllowSchedulingOnMaster) {
		return true
	}

	return false
}

// SetAllowSchedulingOnMaster gets a reference to the given bool and assigns it to the AllowSchedulingOnMaster field.
func (o *CreateKubernetesProfileCommand) SetAllowSchedulingOnMaster(v bool) {
	o.AllowSchedulingOnMaster = &v
}

// GetUniqueClusterName returns the UniqueClusterName field value if set, zero value otherwise.
func (o *CreateKubernetesProfileCommand) GetUniqueClusterName() bool {
	if o == nil || IsNil(o.UniqueClusterName) {
		var ret bool
		return ret
	}
	return *o.UniqueClusterName
}

// GetUniqueClusterNameOk returns a tuple with the UniqueClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateKubernetesProfileCommand) GetUniqueClusterNameOk() (*bool, bool) {
	if o == nil || IsNil(o.UniqueClusterName) {
		return nil, false
	}
	return o.UniqueClusterName, true
}

// HasUniqueClusterName returns a boolean if a field has been set.
func (o *CreateKubernetesProfileCommand) HasUniqueClusterName() bool {
	if o != nil && !IsNil(o.UniqueClusterName) {
		return true
	}

	return false
}

// SetUniqueClusterName gets a reference to the given bool and assigns it to the UniqueClusterName field.
func (o *CreateKubernetesProfileCommand) SetUniqueClusterName(v bool) {
	o.UniqueClusterName = &v
}

func (o CreateKubernetesProfileCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateKubernetesProfileCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.OctaviaEnabled) {
		toSerialize["octaviaEnabled"] = o.OctaviaEnabled
	}
	if !IsNil(o.ExposeNodePortOnBastion) {
		toSerialize["exposeNodePortOnBastion"] = o.ExposeNodePortOnBastion
	}
	if o.OrganizationId.IsSet() {
		toSerialize["organizationId"] = o.OrganizationId.Get()
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
	return toSerialize, nil
}

type NullableCreateKubernetesProfileCommand struct {
	value *CreateKubernetesProfileCommand
	isSet bool
}

func (v NullableCreateKubernetesProfileCommand) Get() *CreateKubernetesProfileCommand {
	return v.value
}

func (v *NullableCreateKubernetesProfileCommand) Set(val *CreateKubernetesProfileCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateKubernetesProfileCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateKubernetesProfileCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateKubernetesProfileCommand(val *CreateKubernetesProfileCommand) *NullableCreateKubernetesProfileCommand {
	return &NullableCreateKubernetesProfileCommand{value: val, isSet: true}
}

func (v NullableCreateKubernetesProfileCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateKubernetesProfileCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


