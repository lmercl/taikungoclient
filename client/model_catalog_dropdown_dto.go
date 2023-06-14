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

// checks if the CatalogDropdownDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogDropdownDto{}

// CatalogDropdownDto struct for CatalogDropdownDto
type CatalogDropdownDto struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	PackageIds []string `json:"packageIds,omitempty"`
}

// NewCatalogDropdownDto instantiates a new CatalogDropdownDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogDropdownDto() *CatalogDropdownDto {
	this := CatalogDropdownDto{}
	return &this
}

// NewCatalogDropdownDtoWithDefaults instantiates a new CatalogDropdownDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogDropdownDtoWithDefaults() *CatalogDropdownDto {
	this := CatalogDropdownDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CatalogDropdownDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogDropdownDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CatalogDropdownDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CatalogDropdownDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogDropdownDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogDropdownDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CatalogDropdownDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CatalogDropdownDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CatalogDropdownDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CatalogDropdownDto) UnsetName() {
	o.Name.Unset()
}

// GetPackageIds returns the PackageIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogDropdownDto) GetPackageIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.PackageIds
}

// GetPackageIdsOk returns a tuple with the PackageIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogDropdownDto) GetPackageIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.PackageIds) {
		return nil, false
	}
	return o.PackageIds, true
}

// HasPackageIds returns a boolean if a field has been set.
func (o *CatalogDropdownDto) HasPackageIds() bool {
	if o != nil && IsNil(o.PackageIds) {
		return true
	}

	return false
}

// SetPackageIds gets a reference to the given []string and assigns it to the PackageIds field.
func (o *CatalogDropdownDto) SetPackageIds(v []string) {
	o.PackageIds = v
}

func (o CatalogDropdownDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogDropdownDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.PackageIds != nil {
		toSerialize["packageIds"] = o.PackageIds
	}
	return toSerialize, nil
}

type NullableCatalogDropdownDto struct {
	value *CatalogDropdownDto
	isSet bool
}

func (v NullableCatalogDropdownDto) Get() *CatalogDropdownDto {
	return v.value
}

func (v *NullableCatalogDropdownDto) Set(val *CatalogDropdownDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogDropdownDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogDropdownDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogDropdownDto(val *CatalogDropdownDto) *NullableCatalogDropdownDto {
	return &NullableCatalogDropdownDto{value: val, isSet: true}
}

func (v NullableCatalogDropdownDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogDropdownDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


