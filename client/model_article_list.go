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

// checks if the ArticleList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArticleList{}

// ArticleList struct for ArticleList
type ArticleList struct {
	Data []ArticlesListDto `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewArticleList instantiates a new ArticleList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArticleList() *ArticleList {
	this := ArticleList{}
	return &this
}

// NewArticleListWithDefaults instantiates a new ArticleList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArticleListWithDefaults() *ArticleList {
	this := ArticleList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArticleList) GetData() []ArticlesListDto {
	if o == nil {
		var ret []ArticlesListDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArticleList) GetDataOk() ([]ArticlesListDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ArticleList) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ArticlesListDto and assigns it to the Data field.
func (o *ArticleList) SetData(v []ArticlesListDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ArticleList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArticleList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ArticleList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *ArticleList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o ArticleList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArticleList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableArticleList struct {
	value *ArticleList
	isSet bool
}

func (v NullableArticleList) Get() *ArticleList {
	return v.value
}

func (v *NullableArticleList) Set(val *ArticleList) {
	v.value = val
	v.isSet = true
}

func (v NullableArticleList) IsSet() bool {
	return v.isSet
}

func (v *NullableArticleList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArticleList(val *ArticleList) *NullableArticleList {
	return &NullableArticleList{value: val, isSet: true}
}

func (v NullableArticleList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArticleList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


