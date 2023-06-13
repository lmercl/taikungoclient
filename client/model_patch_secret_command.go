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

// checks if the PatchSecretCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchSecretCommand{}

// PatchSecretCommand struct for PatchSecretCommand
type PatchSecretCommand struct {
	ProjectId *int32 `json:"projectId,omitempty"`
	Yaml NullableString `json:"yaml,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Namespace NullableString `json:"namespace,omitempty"`
}

// NewPatchSecretCommand instantiates a new PatchSecretCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchSecretCommand() *PatchSecretCommand {
	this := PatchSecretCommand{}
	return &this
}

// NewPatchSecretCommandWithDefaults instantiates a new PatchSecretCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchSecretCommandWithDefaults() *PatchSecretCommand {
	this := PatchSecretCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *PatchSecretCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSecretCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *PatchSecretCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *PatchSecretCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetYaml returns the Yaml field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchSecretCommand) GetYaml() string {
	if o == nil || IsNil(o.Yaml.Get()) {
		var ret string
		return ret
	}
	return *o.Yaml.Get()
}

// GetYamlOk returns a tuple with the Yaml field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchSecretCommand) GetYamlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Yaml.Get(), o.Yaml.IsSet()
}

// HasYaml returns a boolean if a field has been set.
func (o *PatchSecretCommand) HasYaml() bool {
	if o != nil && o.Yaml.IsSet() {
		return true
	}

	return false
}

// SetYaml gets a reference to the given NullableString and assigns it to the Yaml field.
func (o *PatchSecretCommand) SetYaml(v string) {
	o.Yaml.Set(&v)
}
// SetYamlNil sets the value for Yaml to be an explicit nil
func (o *PatchSecretCommand) SetYamlNil() {
	o.Yaml.Set(nil)
}

// UnsetYaml ensures that no value is present for Yaml, not even an explicit nil
func (o *PatchSecretCommand) UnsetYaml() {
	o.Yaml.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchSecretCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchSecretCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PatchSecretCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PatchSecretCommand) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *PatchSecretCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PatchSecretCommand) UnsetName() {
	o.Name.Unset()
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchSecretCommand) GetNamespace() string {
	if o == nil || IsNil(o.Namespace.Get()) {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchSecretCommand) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *PatchSecretCommand) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *PatchSecretCommand) SetNamespace(v string) {
	o.Namespace.Set(&v)
}
// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *PatchSecretCommand) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *PatchSecretCommand) UnsetNamespace() {
	o.Namespace.Unset()
}

func (o PatchSecretCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchSecretCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.Yaml.IsSet() {
		toSerialize["yaml"] = o.Yaml.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Namespace.IsSet() {
		toSerialize["namespace"] = o.Namespace.Get()
	}
	return toSerialize, nil
}

type NullablePatchSecretCommand struct {
	value *PatchSecretCommand
	isSet bool
}

func (v NullablePatchSecretCommand) Get() *PatchSecretCommand {
	return v.value
}

func (v *NullablePatchSecretCommand) Set(val *PatchSecretCommand) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchSecretCommand) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchSecretCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchSecretCommand(val *PatchSecretCommand) *NullablePatchSecretCommand {
	return &NullablePatchSecretCommand{value: val, isSet: true}
}

func (v NullablePatchSecretCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchSecretCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


