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

// checks if the CatalogAppParamsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogAppParamsDto{}

// CatalogAppParamsDto struct for CatalogAppParamsDto
type CatalogAppParamsDto struct {
	Key string `json:"key"`
	Value NullableString `json:"value,omitempty"`
	IsEditableWhenInstalling *bool `json:"isEditableWhenInstalling,omitempty"`
	IsEditableAfterInstallation *bool `json:"isEditableAfterInstallation,omitempty"`
	IsMandatory *bool `json:"isMandatory,omitempty"`
}

// NewCatalogAppParamsDto instantiates a new CatalogAppParamsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogAppParamsDto(key string) *CatalogAppParamsDto {
	this := CatalogAppParamsDto{}
	this.Key = key
	return &this
}

// NewCatalogAppParamsDtoWithDefaults instantiates a new CatalogAppParamsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogAppParamsDtoWithDefaults() *CatalogAppParamsDto {
	this := CatalogAppParamsDto{}
	return &this
}

// GetKey returns the Key field value
func (o *CatalogAppParamsDto) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *CatalogAppParamsDto) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *CatalogAppParamsDto) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogAppParamsDto) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogAppParamsDto) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *CatalogAppParamsDto) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *CatalogAppParamsDto) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *CatalogAppParamsDto) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *CatalogAppParamsDto) UnsetValue() {
	o.Value.Unset()
}

// GetIsEditableWhenInstalling returns the IsEditableWhenInstalling field value if set, zero value otherwise.
func (o *CatalogAppParamsDto) GetIsEditableWhenInstalling() bool {
	if o == nil || IsNil(o.IsEditableWhenInstalling) {
		var ret bool
		return ret
	}
	return *o.IsEditableWhenInstalling
}

// GetIsEditableWhenInstallingOk returns a tuple with the IsEditableWhenInstalling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogAppParamsDto) GetIsEditableWhenInstallingOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEditableWhenInstalling) {
		return nil, false
	}
	return o.IsEditableWhenInstalling, true
}

// HasIsEditableWhenInstalling returns a boolean if a field has been set.
func (o *CatalogAppParamsDto) HasIsEditableWhenInstalling() bool {
	if o != nil && !IsNil(o.IsEditableWhenInstalling) {
		return true
	}

	return false
}

// SetIsEditableWhenInstalling gets a reference to the given bool and assigns it to the IsEditableWhenInstalling field.
func (o *CatalogAppParamsDto) SetIsEditableWhenInstalling(v bool) {
	o.IsEditableWhenInstalling = &v
}

// GetIsEditableAfterInstallation returns the IsEditableAfterInstallation field value if set, zero value otherwise.
func (o *CatalogAppParamsDto) GetIsEditableAfterInstallation() bool {
	if o == nil || IsNil(o.IsEditableAfterInstallation) {
		var ret bool
		return ret
	}
	return *o.IsEditableAfterInstallation
}

// GetIsEditableAfterInstallationOk returns a tuple with the IsEditableAfterInstallation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogAppParamsDto) GetIsEditableAfterInstallationOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEditableAfterInstallation) {
		return nil, false
	}
	return o.IsEditableAfterInstallation, true
}

// HasIsEditableAfterInstallation returns a boolean if a field has been set.
func (o *CatalogAppParamsDto) HasIsEditableAfterInstallation() bool {
	if o != nil && !IsNil(o.IsEditableAfterInstallation) {
		return true
	}

	return false
}

// SetIsEditableAfterInstallation gets a reference to the given bool and assigns it to the IsEditableAfterInstallation field.
func (o *CatalogAppParamsDto) SetIsEditableAfterInstallation(v bool) {
	o.IsEditableAfterInstallation = &v
}

// GetIsMandatory returns the IsMandatory field value if set, zero value otherwise.
func (o *CatalogAppParamsDto) GetIsMandatory() bool {
	if o == nil || IsNil(o.IsMandatory) {
		var ret bool
		return ret
	}
	return *o.IsMandatory
}

// GetIsMandatoryOk returns a tuple with the IsMandatory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogAppParamsDto) GetIsMandatoryOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMandatory) {
		return nil, false
	}
	return o.IsMandatory, true
}

// HasIsMandatory returns a boolean if a field has been set.
func (o *CatalogAppParamsDto) HasIsMandatory() bool {
	if o != nil && !IsNil(o.IsMandatory) {
		return true
	}

	return false
}

// SetIsMandatory gets a reference to the given bool and assigns it to the IsMandatory field.
func (o *CatalogAppParamsDto) SetIsMandatory(v bool) {
	o.IsMandatory = &v
}

func (o CatalogAppParamsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogAppParamsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if !IsNil(o.IsEditableWhenInstalling) {
		toSerialize["isEditableWhenInstalling"] = o.IsEditableWhenInstalling
	}
	if !IsNil(o.IsEditableAfterInstallation) {
		toSerialize["isEditableAfterInstallation"] = o.IsEditableAfterInstallation
	}
	if !IsNil(o.IsMandatory) {
		toSerialize["isMandatory"] = o.IsMandatory
	}
	return toSerialize, nil
}

type NullableCatalogAppParamsDto struct {
	value *CatalogAppParamsDto
	isSet bool
}

func (v NullableCatalogAppParamsDto) Get() *CatalogAppParamsDto {
	return v.value
}

func (v *NullableCatalogAppParamsDto) Set(val *CatalogAppParamsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogAppParamsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogAppParamsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogAppParamsDto(val *CatalogAppParamsDto) *NullableCatalogAppParamsDto {
	return &NullableCatalogAppParamsDto{value: val, isSet: true}
}

func (v NullableCatalogAppParamsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogAppParamsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


