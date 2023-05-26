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

// checks if the AzResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzResult{}

// AzResult struct for AzResult
type AzResult struct {
	List []string `json:"list,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewAzResult instantiates a new AzResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzResult() *AzResult {
	this := AzResult{}
	return &this
}

// NewAzResultWithDefaults instantiates a new AzResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzResultWithDefaults() *AzResult {
	this := AzResult{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzResult) GetList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzResult) GetListOk() ([]string, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *AzResult) HasList() bool {
	if o != nil && IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []string and assigns it to the List field.
func (o *AzResult) SetList(v []string) {
	o.List = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *AzResult) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzResult) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *AzResult) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *AzResult) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o AzResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.List != nil {
		toSerialize["list"] = o.List
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableAzResult struct {
	value *AzResult
	isSet bool
}

func (v NullableAzResult) Get() *AzResult {
	return v.value
}

func (v *NullableAzResult) Set(val *AzResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAzResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAzResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzResult(val *AzResult) *NullableAzResult {
	return &NullableAzResult{value: val, isSet: true}
}

func (v NullableAzResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


