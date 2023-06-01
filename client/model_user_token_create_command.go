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
	"time"
)

// checks if the UserTokenCreateCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserTokenCreateCommand{}

// UserTokenCreateCommand struct for UserTokenCreateCommand
type UserTokenCreateCommand struct {
	ExpireDate NullableTime `json:"expireDate,omitempty"`
	IsReadonly *bool `json:"isReadonly,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Endpoints []AvailableEndpointData `json:"endpoints,omitempty"`
	BindALL *bool `json:"bindALL,omitempty"`
}

// NewUserTokenCreateCommand instantiates a new UserTokenCreateCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserTokenCreateCommand() *UserTokenCreateCommand {
	this := UserTokenCreateCommand{}
	return &this
}

// NewUserTokenCreateCommandWithDefaults instantiates a new UserTokenCreateCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserTokenCreateCommandWithDefaults() *UserTokenCreateCommand {
	this := UserTokenCreateCommand{}
	return &this
}

// GetExpireDate returns the ExpireDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserTokenCreateCommand) GetExpireDate() time.Time {
	if o == nil || IsNil(o.ExpireDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpireDate.Get()
}

// GetExpireDateOk returns a tuple with the ExpireDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserTokenCreateCommand) GetExpireDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpireDate.Get(), o.ExpireDate.IsSet()
}

// HasExpireDate returns a boolean if a field has been set.
func (o *UserTokenCreateCommand) HasExpireDate() bool {
	if o != nil && o.ExpireDate.IsSet() {
		return true
	}

	return false
}

// SetExpireDate gets a reference to the given NullableTime and assigns it to the ExpireDate field.
func (o *UserTokenCreateCommand) SetExpireDate(v time.Time) {
	o.ExpireDate.Set(&v)
}
// SetExpireDateNil sets the value for ExpireDate to be an explicit nil
func (o *UserTokenCreateCommand) SetExpireDateNil() {
	o.ExpireDate.Set(nil)
}

// UnsetExpireDate ensures that no value is present for ExpireDate, not even an explicit nil
func (o *UserTokenCreateCommand) UnsetExpireDate() {
	o.ExpireDate.Unset()
}

// GetIsReadonly returns the IsReadonly field value if set, zero value otherwise.
func (o *UserTokenCreateCommand) GetIsReadonly() bool {
	if o == nil || IsNil(o.IsReadonly) {
		var ret bool
		return ret
	}
	return *o.IsReadonly
}

// GetIsReadonlyOk returns a tuple with the IsReadonly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTokenCreateCommand) GetIsReadonlyOk() (*bool, bool) {
	if o == nil || IsNil(o.IsReadonly) {
		return nil, false
	}
	return o.IsReadonly, true
}

// HasIsReadonly returns a boolean if a field has been set.
func (o *UserTokenCreateCommand) HasIsReadonly() bool {
	if o != nil && !IsNil(o.IsReadonly) {
		return true
	}

	return false
}

// SetIsReadonly gets a reference to the given bool and assigns it to the IsReadonly field.
func (o *UserTokenCreateCommand) SetIsReadonly(v bool) {
	o.IsReadonly = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserTokenCreateCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserTokenCreateCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UserTokenCreateCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UserTokenCreateCommand) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *UserTokenCreateCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UserTokenCreateCommand) UnsetName() {
	o.Name.Unset()
}

// GetEndpoints returns the Endpoints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserTokenCreateCommand) GetEndpoints() []AvailableEndpointData {
	if o == nil {
		var ret []AvailableEndpointData
		return ret
	}
	return o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserTokenCreateCommand) GetEndpointsOk() ([]AvailableEndpointData, bool) {
	if o == nil || IsNil(o.Endpoints) {
		return nil, false
	}
	return o.Endpoints, true
}

// HasEndpoints returns a boolean if a field has been set.
func (o *UserTokenCreateCommand) HasEndpoints() bool {
	if o != nil && IsNil(o.Endpoints) {
		return true
	}

	return false
}

// SetEndpoints gets a reference to the given []AvailableEndpointData and assigns it to the Endpoints field.
func (o *UserTokenCreateCommand) SetEndpoints(v []AvailableEndpointData) {
	o.Endpoints = v
}

// GetBindALL returns the BindALL field value if set, zero value otherwise.
func (o *UserTokenCreateCommand) GetBindALL() bool {
	if o == nil || IsNil(o.BindALL) {
		var ret bool
		return ret
	}
	return *o.BindALL
}

// GetBindALLOk returns a tuple with the BindALL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTokenCreateCommand) GetBindALLOk() (*bool, bool) {
	if o == nil || IsNil(o.BindALL) {
		return nil, false
	}
	return o.BindALL, true
}

// HasBindALL returns a boolean if a field has been set.
func (o *UserTokenCreateCommand) HasBindALL() bool {
	if o != nil && !IsNil(o.BindALL) {
		return true
	}

	return false
}

// SetBindALL gets a reference to the given bool and assigns it to the BindALL field.
func (o *UserTokenCreateCommand) SetBindALL(v bool) {
	o.BindALL = &v
}

func (o UserTokenCreateCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserTokenCreateCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpireDate.IsSet() {
		toSerialize["expireDate"] = o.ExpireDate.Get()
	}
	if !IsNil(o.IsReadonly) {
		toSerialize["isReadonly"] = o.IsReadonly
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Endpoints != nil {
		toSerialize["endpoints"] = o.Endpoints
	}
	if !IsNil(o.BindALL) {
		toSerialize["bindALL"] = o.BindALL
	}
	return toSerialize, nil
}

type NullableUserTokenCreateCommand struct {
	value *UserTokenCreateCommand
	isSet bool
}

func (v NullableUserTokenCreateCommand) Get() *UserTokenCreateCommand {
	return v.value
}

func (v *NullableUserTokenCreateCommand) Set(val *UserTokenCreateCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableUserTokenCreateCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableUserTokenCreateCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserTokenCreateCommand(val *UserTokenCreateCommand) *NullableUserTokenCreateCommand {
	return &NullableUserTokenCreateCommand{value: val, isSet: true}
}

func (v NullableUserTokenCreateCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserTokenCreateCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


