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

// TicketPriority the model 'TicketPriority'
type TicketPriority string

// List of TicketPriority
const (
	TICKETPRIORITY_LOW TicketPriority = "Low"
	TICKETPRIORITY_MEDIUM TicketPriority = "Medium"
	TICKETPRIORITY_HIGH TicketPriority = "High"
)

// All allowed values of TicketPriority enum
var AllowedTicketPriorityEnumValues = []TicketPriority{
	"Low",
	"Medium",
	"High",
}

func (v *TicketPriority) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TicketPriority(value)
	for _, existing := range AllowedTicketPriorityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TicketPriority", value)
}

// NewTicketPriorityFromValue returns a pointer to a valid TicketPriority
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTicketPriorityFromValue(v string) (*TicketPriority, error) {
	ev := TicketPriority(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TicketPriority: valid values are %v", v, AllowedTicketPriorityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TicketPriority) IsValid() bool {
	for _, existing := range AllowedTicketPriorityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TicketPriority value
func (v TicketPriority) Ptr() *TicketPriority {
	return &v
}

type NullableTicketPriority struct {
	value *TicketPriority
	isSet bool
}

func (v NullableTicketPriority) Get() *TicketPriority {
	return v.value
}

func (v *NullableTicketPriority) Set(val *TicketPriority) {
	v.value = val
	v.isSet = true
}

func (v NullableTicketPriority) IsSet() bool {
	return v.isSet
}

func (v *NullableTicketPriority) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTicketPriority(val *TicketPriority) *NullableTicketPriority {
	return &NullableTicketPriority{value: val, isSet: true}
}

func (v NullableTicketPriority) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTicketPriority) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

