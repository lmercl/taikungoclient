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

// CloudStatus the model 'CloudStatus'
type CloudStatus string

// List of CloudStatus
const (
	CLOUDSTATUS_DELETING CloudStatus = "Deleting"
	CLOUDSTATUS_FAILURE CloudStatus = "Failure"
	CLOUDSTATUS_FAILED_UPGRADE CloudStatus = "FailedUpgrade"
	CLOUDSTATUS_WAITING CloudStatus = "Waiting"
	CLOUDSTATUS_PENDING_DELETE CloudStatus = "PendingDelete"
	CLOUDSTATUS_PENDING_UPGRADE CloudStatus = "PendingUpgrade"
	CLOUDSTATUS_READY CloudStatus = "Ready"
	CLOUDSTATUS_UPDATING CloudStatus = "Updating"
	CLOUDSTATUS_UPGRADING CloudStatus = "Upgrading"
)

// All allowed values of CloudStatus enum
var AllowedCloudStatusEnumValues = []CloudStatus{
	"Deleting",
	"Failure",
	"FailedUpgrade",
	"Waiting",
	"PendingDelete",
	"PendingUpgrade",
	"Ready",
	"Updating",
	"Upgrading",
}

func (v *CloudStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CloudStatus(value)
	for _, existing := range AllowedCloudStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CloudStatus", value)
}

// NewCloudStatusFromValue returns a pointer to a valid CloudStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCloudStatusFromValue(v string) (*CloudStatus, error) {
	ev := CloudStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CloudStatus: valid values are %v", v, AllowedCloudStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CloudStatus) IsValid() bool {
	for _, existing := range AllowedCloudStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloudStatus value
func (v CloudStatus) Ptr() *CloudStatus {
	return &v
}

type NullableCloudStatus struct {
	value *CloudStatus
	isSet bool
}

func (v NullableCloudStatus) Get() *CloudStatus {
	return v.value
}

func (v *NullableCloudStatus) Set(val *CloudStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudStatus(val *CloudStatus) *NullableCloudStatus {
	return &NullableCloudStatus{value: val, isSet: true}
}

func (v NullableCloudStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

