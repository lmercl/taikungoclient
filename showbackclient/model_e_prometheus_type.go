/*
Taikun - Showback API

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikunshowback

import (
	"encoding/json"
	"fmt"
)

// EPrometheusType the model 'EPrometheusType'
type EPrometheusType int32

// List of EPrometheusType
const (
	EPROMETHEUSTYPE__100 EPrometheusType = 100
	EPROMETHEUSTYPE__200 EPrometheusType = 200
)

// All allowed values of EPrometheusType enum
var AllowedEPrometheusTypeEnumValues = []EPrometheusType{
	100,
	200,
}

func (v *EPrometheusType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EPrometheusType(value)
	for _, existing := range AllowedEPrometheusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EPrometheusType", value)
}

// NewEPrometheusTypeFromValue returns a pointer to a valid EPrometheusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEPrometheusTypeFromValue(v int32) (*EPrometheusType, error) {
	ev := EPrometheusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EPrometheusType: valid values are %v", v, AllowedEPrometheusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EPrometheusType) IsValid() bool {
	for _, existing := range AllowedEPrometheusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EPrometheusType value
func (v EPrometheusType) Ptr() *EPrometheusType {
	return &v
}

type NullableEPrometheusType struct {
	value *EPrometheusType
	isSet bool
}

func (v NullableEPrometheusType) Get() *EPrometheusType {
	return v.value
}

func (v *NullableEPrometheusType) Set(val *EPrometheusType) {
	v.value = val
	v.isSet = true
}

func (v NullableEPrometheusType) IsSet() bool {
	return v.isSet
}

func (v *NullableEPrometheusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEPrometheusType(val *EPrometheusType) *NullableEPrometheusType {
	return &NullableEPrometheusType{value: val, isSet: true}
}

func (v NullableEPrometheusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEPrometheusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

