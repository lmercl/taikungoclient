/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikuncore

import (
	"encoding/json"
)

// checks if the BackupCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupCredentials{}

// BackupCredentials struct for BackupCredentials
type BackupCredentials struct {
	Data []BackupCredentialsListDto `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewBackupCredentials instantiates a new BackupCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupCredentials() *BackupCredentials {
	this := BackupCredentials{}
	return &this
}

// NewBackupCredentialsWithDefaults instantiates a new BackupCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupCredentialsWithDefaults() *BackupCredentials {
	this := BackupCredentials{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupCredentials) GetData() []BackupCredentialsListDto {
	if o == nil {
		var ret []BackupCredentialsListDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupCredentials) GetDataOk() ([]BackupCredentialsListDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BackupCredentials) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []BackupCredentialsListDto and assigns it to the Data field.
func (o *BackupCredentials) SetData(v []BackupCredentialsListDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *BackupCredentials) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupCredentials) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *BackupCredentials) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *BackupCredentials) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o BackupCredentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableBackupCredentials struct {
	value *BackupCredentials
	isSet bool
}

func (v NullableBackupCredentials) Get() *BackupCredentials {
	return v.value
}

func (v *NullableBackupCredentials) Set(val *BackupCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupCredentials(val *BackupCredentials) *NullableBackupCredentials {
	return &NullableBackupCredentials{value: val, isSet: true}
}

func (v NullableBackupCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


