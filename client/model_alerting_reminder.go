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

// AlertingReminder the model 'AlertingReminder'
type AlertingReminder int32

// List of AlertingReminder
const (
	ALERTINGREMINDER__100 AlertingReminder = 100
	ALERTINGREMINDER__200 AlertingReminder = 200
	ALERTINGREMINDER__300 AlertingReminder = 300
	ALERTINGREMINDER__MINUS_1 AlertingReminder = -1
)

// All allowed values of AlertingReminder enum
var AllowedAlertingReminderEnumValues = []AlertingReminder{
	100,
	200,
	300,
	-1,
}

func (v *AlertingReminder) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AlertingReminder(value)
	for _, existing := range AllowedAlertingReminderEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AlertingReminder", value)
}

// NewAlertingReminderFromValue returns a pointer to a valid AlertingReminder
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAlertingReminderFromValue(v int32) (*AlertingReminder, error) {
	ev := AlertingReminder(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AlertingReminder: valid values are %v", v, AllowedAlertingReminderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AlertingReminder) IsValid() bool {
	for _, existing := range AllowedAlertingReminderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlertingReminder value
func (v AlertingReminder) Ptr() *AlertingReminder {
	return &v
}

type NullableAlertingReminder struct {
	value *AlertingReminder
	isSet bool
}

func (v NullableAlertingReminder) Get() *AlertingReminder {
	return v.value
}

func (v *NullableAlertingReminder) Set(val *AlertingReminder) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertingReminder) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertingReminder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertingReminder(val *AlertingReminder) *NullableAlertingReminder {
	return &NullableAlertingReminder{value: val, isSet: true}
}

func (v NullableAlertingReminder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertingReminder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

