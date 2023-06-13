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

// StandAloneVmStatus the model 'StandAloneVmStatus'
type StandAloneVmStatus string

// List of StandAloneVmStatus
const (
	STANDALONEVMSTATUS_DELETING StandAloneVmStatus = "Deleting"
	STANDALONEVMSTATUS_FAILURE StandAloneVmStatus = "Failure"
	STANDALONEVMSTATUS_WAITING StandAloneVmStatus = "Waiting"
	STANDALONEVMSTATUS_PENDING_DELETE StandAloneVmStatus = "PendingDelete"
	STANDALONEVMSTATUS_READY StandAloneVmStatus = "Ready"
	STANDALONEVMSTATUS_UPDATING StandAloneVmStatus = "Updating"
)

// All allowed values of StandAloneVmStatus enum
var AllowedStandAloneVmStatusEnumValues = []StandAloneVmStatus{
	"Deleting",
	"Failure",
	"Waiting",
	"PendingDelete",
	"Ready",
	"Updating",
}

func (v *StandAloneVmStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StandAloneVmStatus(value)
	for _, existing := range AllowedStandAloneVmStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StandAloneVmStatus", value)
}

// NewStandAloneVmStatusFromValue returns a pointer to a valid StandAloneVmStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStandAloneVmStatusFromValue(v string) (*StandAloneVmStatus, error) {
	ev := StandAloneVmStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StandAloneVmStatus: valid values are %v", v, AllowedStandAloneVmStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StandAloneVmStatus) IsValid() bool {
	for _, existing := range AllowedStandAloneVmStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StandAloneVmStatus value
func (v StandAloneVmStatus) Ptr() *StandAloneVmStatus {
	return &v
}

type NullableStandAloneVmStatus struct {
	value *StandAloneVmStatus
	isSet bool
}

func (v NullableStandAloneVmStatus) Get() *StandAloneVmStatus {
	return v.value
}

func (v *NullableStandAloneVmStatus) Set(val *StandAloneVmStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableStandAloneVmStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableStandAloneVmStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandAloneVmStatus(val *StandAloneVmStatus) *NullableStandAloneVmStatus {
	return &NullableStandAloneVmStatus{value: val, isSet: true}
}

func (v NullableStandAloneVmStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandAloneVmStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

