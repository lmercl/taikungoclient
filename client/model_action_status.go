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

// ActionStatus the model 'ActionStatus'
type ActionStatus int32

// List of ActionStatus
const (
	ACTIONSTATUS__0 ActionStatus = 0
	ACTIONSTATUS__1 ActionStatus = 1
	ACTIONSTATUS__2 ActionStatus = 2
)

// All allowed values of ActionStatus enum
var AllowedActionStatusEnumValues = []ActionStatus{
	0,
	1,
	2,
}

func (v *ActionStatus) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ActionStatus(value)
	for _, existing := range AllowedActionStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ActionStatus", value)
}

// NewActionStatusFromValue returns a pointer to a valid ActionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewActionStatusFromValue(v int32) (*ActionStatus, error) {
	ev := ActionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ActionStatus: valid values are %v", v, AllowedActionStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ActionStatus) IsValid() bool {
	for _, existing := range AllowedActionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ActionStatus value
func (v ActionStatus) Ptr() *ActionStatus {
	return &v
}

type NullableActionStatus struct {
	value *ActionStatus
	isSet bool
}

func (v NullableActionStatus) Get() *ActionStatus {
	return v.value
}

func (v *NullableActionStatus) Set(val *ActionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableActionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableActionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionStatus(val *ActionStatus) *NullableActionStatus {
	return &NullableActionStatus{value: val, isSet: true}
}

func (v NullableActionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

