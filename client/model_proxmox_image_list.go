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
)

// checks if the ProxmoxImageList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProxmoxImageList{}

// ProxmoxImageList struct for ProxmoxImageList
type ProxmoxImageList struct {
	Data []CommonDropdownDto `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewProxmoxImageList instantiates a new ProxmoxImageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxmoxImageList() *ProxmoxImageList {
	this := ProxmoxImageList{}
	return &this
}

// NewProxmoxImageListWithDefaults instantiates a new ProxmoxImageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxmoxImageListWithDefaults() *ProxmoxImageList {
	this := ProxmoxImageList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxmoxImageList) GetData() []CommonDropdownDto {
	if o == nil {
		var ret []CommonDropdownDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxmoxImageList) GetDataOk() ([]CommonDropdownDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ProxmoxImageList) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CommonDropdownDto and assigns it to the Data field.
func (o *ProxmoxImageList) SetData(v []CommonDropdownDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ProxmoxImageList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxmoxImageList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ProxmoxImageList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *ProxmoxImageList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o ProxmoxImageList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProxmoxImageList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableProxmoxImageList struct {
	value *ProxmoxImageList
	isSet bool
}

func (v NullableProxmoxImageList) Get() *ProxmoxImageList {
	return v.value
}

func (v *NullableProxmoxImageList) Set(val *ProxmoxImageList) {
	v.value = val
	v.isSet = true
}

func (v NullableProxmoxImageList) IsSet() bool {
	return v.isSet
}

func (v *NullableProxmoxImageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxmoxImageList(val *ProxmoxImageList) *NullableProxmoxImageList {
	return &NullableProxmoxImageList{value: val, isSet: true}
}

func (v NullableProxmoxImageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxmoxImageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


