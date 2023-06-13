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
	"fmt"
)

// CloudType the model 'CloudType'
type CloudType string

// List of CloudType
const (
	CLOUDTYPE_AWS CloudType = "AWS"
	CLOUDTYPE_AZURE CloudType = "AZURE"
	CLOUDTYPE_GOOGLE CloudType = "GOOGLE"
	CLOUDTYPE_OPENSTACK CloudType = "OPENSTACK"
	CLOUDTYPE_PROXMOX CloudType = "PROXMOX"
	CLOUDTYPE_TANZU CloudType = "TANZU"
)

// All allowed values of CloudType enum
var AllowedCloudTypeEnumValues = []CloudType{
	"AWS",
	"AZURE",
	"GOOGLE",
	"OPENSTACK",
	"PROXMOX",
	"TANZU",
}

func (v *CloudType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CloudType(value)
	for _, existing := range AllowedCloudTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CloudType", value)
}

// NewCloudTypeFromValue returns a pointer to a valid CloudType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCloudTypeFromValue(v string) (*CloudType, error) {
	ev := CloudType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CloudType: valid values are %v", v, AllowedCloudTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CloudType) IsValid() bool {
	for _, existing := range AllowedCloudTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloudType value
func (v CloudType) Ptr() *CloudType {
	return &v
}

type NullableCloudType struct {
	value *CloudType
	isSet bool
}

func (v NullableCloudType) Get() *CloudType {
	return v.value
}

func (v *NullableCloudType) Set(val *CloudType) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudType) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudType(val *CloudType) *NullableCloudType {
	return &NullableCloudType{value: val, isSet: true}
}

func (v NullableCloudType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

