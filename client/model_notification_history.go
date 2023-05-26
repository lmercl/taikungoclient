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
)

// checks if the NotificationHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationHistory{}

// NotificationHistory struct for NotificationHistory
type NotificationHistory struct {
	Data []NotificationListDto `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewNotificationHistory instantiates a new NotificationHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationHistory() *NotificationHistory {
	this := NotificationHistory{}
	return &this
}

// NewNotificationHistoryWithDefaults instantiates a new NotificationHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationHistoryWithDefaults() *NotificationHistory {
	this := NotificationHistory{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationHistory) GetData() []NotificationListDto {
	if o == nil {
		var ret []NotificationListDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationHistory) GetDataOk() ([]NotificationListDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *NotificationHistory) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []NotificationListDto and assigns it to the Data field.
func (o *NotificationHistory) SetData(v []NotificationListDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *NotificationHistory) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationHistory) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *NotificationHistory) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *NotificationHistory) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o NotificationHistory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationHistory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableNotificationHistory struct {
	value *NotificationHistory
	isSet bool
}

func (v NullableNotificationHistory) Get() *NotificationHistory {
	return v.value
}

func (v *NullableNotificationHistory) Set(val *NotificationHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationHistory(val *NotificationHistory) *NullableNotificationHistory {
	return &NullableNotificationHistory{value: val, isSet: true}
}

func (v NullableNotificationHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


