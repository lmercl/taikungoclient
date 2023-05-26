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

// SlackType the model 'SlackType'
type SlackType int32

// List of SlackType
const (
	SLACKTYPE__100 SlackType = 100
	SLACKTYPE__200 SlackType = 200
)

// All allowed values of SlackType enum
var AllowedSlackTypeEnumValues = []SlackType{
	100,
	200,
}

func (v *SlackType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SlackType(value)
	for _, existing := range AllowedSlackTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SlackType", value)
}

// NewSlackTypeFromValue returns a pointer to a valid SlackType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSlackTypeFromValue(v int32) (*SlackType, error) {
	ev := SlackType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SlackType: valid values are %v", v, AllowedSlackTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SlackType) IsValid() bool {
	for _, existing := range AllowedSlackTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SlackType value
func (v SlackType) Ptr() *SlackType {
	return &v
}

type NullableSlackType struct {
	value *SlackType
	isSet bool
}

func (v NullableSlackType) Get() *SlackType {
	return v.value
}

func (v *NullableSlackType) Set(val *SlackType) {
	v.value = val
	v.isSet = true
}

func (v NullableSlackType) IsSet() bool {
	return v.isSet
}

func (v *NullableSlackType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlackType(val *SlackType) *NullableSlackType {
	return &NullableSlackType{value: val, isSet: true}
}

func (v NullableSlackType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlackType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

