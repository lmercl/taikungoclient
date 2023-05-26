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
	"fmt"
)

// SecurityGroupProtocol the model 'SecurityGroupProtocol'
type SecurityGroupProtocol int32

// List of SecurityGroupProtocol
const (
	SECURITYGROUPPROTOCOL__100 SecurityGroupProtocol = 100
	SECURITYGROUPPROTOCOL__200 SecurityGroupProtocol = 200
	SECURITYGROUPPROTOCOL__300 SecurityGroupProtocol = 300
)

// All allowed values of SecurityGroupProtocol enum
var AllowedSecurityGroupProtocolEnumValues = []SecurityGroupProtocol{
	100,
	200,
	300,
}

func (v *SecurityGroupProtocol) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SecurityGroupProtocol(value)
	for _, existing := range AllowedSecurityGroupProtocolEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SecurityGroupProtocol", value)
}

// NewSecurityGroupProtocolFromValue returns a pointer to a valid SecurityGroupProtocol
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSecurityGroupProtocolFromValue(v int32) (*SecurityGroupProtocol, error) {
	ev := SecurityGroupProtocol(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SecurityGroupProtocol: valid values are %v", v, AllowedSecurityGroupProtocolEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SecurityGroupProtocol) IsValid() bool {
	for _, existing := range AllowedSecurityGroupProtocolEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityGroupProtocol value
func (v SecurityGroupProtocol) Ptr() *SecurityGroupProtocol {
	return &v
}

type NullableSecurityGroupProtocol struct {
	value *SecurityGroupProtocol
	isSet bool
}

func (v NullableSecurityGroupProtocol) Get() *SecurityGroupProtocol {
	return v.value
}

func (v *NullableSecurityGroupProtocol) Set(val *SecurityGroupProtocol) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityGroupProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityGroupProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityGroupProtocol(val *SecurityGroupProtocol) *NullableSecurityGroupProtocol {
	return &NullableSecurityGroupProtocol{value: val, isSet: true}
}

func (v NullableSecurityGroupProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityGroupProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

