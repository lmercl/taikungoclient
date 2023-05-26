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

// checks if the AdminAddBalanceCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminAddBalanceCommand{}

// AdminAddBalanceCommand struct for AdminAddBalanceCommand
type AdminAddBalanceCommand struct {
	CustomerId string `json:"customerId"`
	Balance *int64 `json:"balance,omitempty"`
}

// NewAdminAddBalanceCommand instantiates a new AdminAddBalanceCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminAddBalanceCommand(customerId string) *AdminAddBalanceCommand {
	this := AdminAddBalanceCommand{}
	this.CustomerId = customerId
	return &this
}

// NewAdminAddBalanceCommandWithDefaults instantiates a new AdminAddBalanceCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminAddBalanceCommandWithDefaults() *AdminAddBalanceCommand {
	this := AdminAddBalanceCommand{}
	return &this
}

// GetCustomerId returns the CustomerId field value
func (o *AdminAddBalanceCommand) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *AdminAddBalanceCommand) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *AdminAddBalanceCommand) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *AdminAddBalanceCommand) GetBalance() int64 {
	if o == nil || IsNil(o.Balance) {
		var ret int64
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminAddBalanceCommand) GetBalanceOk() (*int64, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *AdminAddBalanceCommand) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given int64 and assigns it to the Balance field.
func (o *AdminAddBalanceCommand) SetBalance(v int64) {
	o.Balance = &v
}

func (o AdminAddBalanceCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminAddBalanceCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customerId"] = o.CustomerId
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	return toSerialize, nil
}

type NullableAdminAddBalanceCommand struct {
	value *AdminAddBalanceCommand
	isSet bool
}

func (v NullableAdminAddBalanceCommand) Get() *AdminAddBalanceCommand {
	return v.value
}

func (v *NullableAdminAddBalanceCommand) Set(val *AdminAddBalanceCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminAddBalanceCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminAddBalanceCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminAddBalanceCommand(val *AdminAddBalanceCommand) *NullableAdminAddBalanceCommand {
	return &NullableAdminAddBalanceCommand{value: val, isSet: true}
}

func (v NullableAdminAddBalanceCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminAddBalanceCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


