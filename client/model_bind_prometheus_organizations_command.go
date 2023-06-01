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

// checks if the BindPrometheusOrganizationsCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BindPrometheusOrganizationsCommand{}

// BindPrometheusOrganizationsCommand struct for BindPrometheusOrganizationsCommand
type BindPrometheusOrganizationsCommand struct {
	Organizations []BindOrganizationsToRuleDto `json:"organizations,omitempty"`
	PrometheusRuleId *int32 `json:"prometheusRuleId,omitempty"`
}

// NewBindPrometheusOrganizationsCommand instantiates a new BindPrometheusOrganizationsCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBindPrometheusOrganizationsCommand() *BindPrometheusOrganizationsCommand {
	this := BindPrometheusOrganizationsCommand{}
	return &this
}

// NewBindPrometheusOrganizationsCommandWithDefaults instantiates a new BindPrometheusOrganizationsCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBindPrometheusOrganizationsCommandWithDefaults() *BindPrometheusOrganizationsCommand {
	this := BindPrometheusOrganizationsCommand{}
	return &this
}

// GetOrganizations returns the Organizations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BindPrometheusOrganizationsCommand) GetOrganizations() []BindOrganizationsToRuleDto {
	if o == nil {
		var ret []BindOrganizationsToRuleDto
		return ret
	}
	return o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BindPrometheusOrganizationsCommand) GetOrganizationsOk() ([]BindOrganizationsToRuleDto, bool) {
	if o == nil || IsNil(o.Organizations) {
		return nil, false
	}
	return o.Organizations, true
}

// HasOrganizations returns a boolean if a field has been set.
func (o *BindPrometheusOrganizationsCommand) HasOrganizations() bool {
	if o != nil && IsNil(o.Organizations) {
		return true
	}

	return false
}

// SetOrganizations gets a reference to the given []BindOrganizationsToRuleDto and assigns it to the Organizations field.
func (o *BindPrometheusOrganizationsCommand) SetOrganizations(v []BindOrganizationsToRuleDto) {
	o.Organizations = v
}

// GetPrometheusRuleId returns the PrometheusRuleId field value if set, zero value otherwise.
func (o *BindPrometheusOrganizationsCommand) GetPrometheusRuleId() int32 {
	if o == nil || IsNil(o.PrometheusRuleId) {
		var ret int32
		return ret
	}
	return *o.PrometheusRuleId
}

// GetPrometheusRuleIdOk returns a tuple with the PrometheusRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BindPrometheusOrganizationsCommand) GetPrometheusRuleIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PrometheusRuleId) {
		return nil, false
	}
	return o.PrometheusRuleId, true
}

// HasPrometheusRuleId returns a boolean if a field has been set.
func (o *BindPrometheusOrganizationsCommand) HasPrometheusRuleId() bool {
	if o != nil && !IsNil(o.PrometheusRuleId) {
		return true
	}

	return false
}

// SetPrometheusRuleId gets a reference to the given int32 and assigns it to the PrometheusRuleId field.
func (o *BindPrometheusOrganizationsCommand) SetPrometheusRuleId(v int32) {
	o.PrometheusRuleId = &v
}

func (o BindPrometheusOrganizationsCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BindPrometheusOrganizationsCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Organizations != nil {
		toSerialize["organizations"] = o.Organizations
	}
	if !IsNil(o.PrometheusRuleId) {
		toSerialize["prometheusRuleId"] = o.PrometheusRuleId
	}
	return toSerialize, nil
}

type NullableBindPrometheusOrganizationsCommand struct {
	value *BindPrometheusOrganizationsCommand
	isSet bool
}

func (v NullableBindPrometheusOrganizationsCommand) Get() *BindPrometheusOrganizationsCommand {
	return v.value
}

func (v *NullableBindPrometheusOrganizationsCommand) Set(val *BindPrometheusOrganizationsCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableBindPrometheusOrganizationsCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableBindPrometheusOrganizationsCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBindPrometheusOrganizationsCommand(val *BindPrometheusOrganizationsCommand) *NullableBindPrometheusOrganizationsCommand {
	return &NullableBindPrometheusOrganizationsCommand{value: val, isSet: true}
}

func (v NullableBindPrometheusOrganizationsCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBindPrometheusOrganizationsCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


