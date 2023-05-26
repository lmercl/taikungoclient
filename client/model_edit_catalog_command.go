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

// checks if the EditCatalogCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditCatalogCommand{}

// EditCatalogCommand struct for EditCatalogCommand
type EditCatalogCommand struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	Description NullableString `json:"description,omitempty"`
}

// NewEditCatalogCommand instantiates a new EditCatalogCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditCatalogCommand(id int32, name string) *EditCatalogCommand {
	this := EditCatalogCommand{}
	this.Id = id
	this.Name = name
	return &this
}

// NewEditCatalogCommandWithDefaults instantiates a new EditCatalogCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditCatalogCommandWithDefaults() *EditCatalogCommand {
	this := EditCatalogCommand{}
	return &this
}

// GetId returns the Id field value
func (o *EditCatalogCommand) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EditCatalogCommand) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EditCatalogCommand) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *EditCatalogCommand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EditCatalogCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EditCatalogCommand) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditCatalogCommand) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditCatalogCommand) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *EditCatalogCommand) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *EditCatalogCommand) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *EditCatalogCommand) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *EditCatalogCommand) UnsetDescription() {
	o.Description.Unset()
}

func (o EditCatalogCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditCatalogCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	return toSerialize, nil
}

type NullableEditCatalogCommand struct {
	value *EditCatalogCommand
	isSet bool
}

func (v NullableEditCatalogCommand) Get() *EditCatalogCommand {
	return v.value
}

func (v *NullableEditCatalogCommand) Set(val *EditCatalogCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableEditCatalogCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableEditCatalogCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditCatalogCommand(val *EditCatalogCommand) *NullableEditCatalogCommand {
	return &NullableEditCatalogCommand{value: val, isSet: true}
}

func (v NullableEditCatalogCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditCatalogCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


