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

// checks if the PrometheusDashboardUpdateCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrometheusDashboardUpdateCommand{}

// PrometheusDashboardUpdateCommand struct for PrometheusDashboardUpdateCommand
type PrometheusDashboardUpdateCommand struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Expression NullableString `json:"expression,omitempty"`
	Description NullableString `json:"description,omitempty"`
	CategoryName NullableString `json:"categoryName,omitempty"`
}

// NewPrometheusDashboardUpdateCommand instantiates a new PrometheusDashboardUpdateCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrometheusDashboardUpdateCommand() *PrometheusDashboardUpdateCommand {
	this := PrometheusDashboardUpdateCommand{}
	return &this
}

// NewPrometheusDashboardUpdateCommandWithDefaults instantiates a new PrometheusDashboardUpdateCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrometheusDashboardUpdateCommandWithDefaults() *PrometheusDashboardUpdateCommand {
	this := PrometheusDashboardUpdateCommand{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PrometheusDashboardUpdateCommand) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusDashboardUpdateCommand) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PrometheusDashboardUpdateCommand) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PrometheusDashboardUpdateCommand) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusDashboardUpdateCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusDashboardUpdateCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PrometheusDashboardUpdateCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PrometheusDashboardUpdateCommand) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *PrometheusDashboardUpdateCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PrometheusDashboardUpdateCommand) UnsetName() {
	o.Name.Unset()
}

// GetExpression returns the Expression field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusDashboardUpdateCommand) GetExpression() string {
	if o == nil || IsNil(o.Expression.Get()) {
		var ret string
		return ret
	}
	return *o.Expression.Get()
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusDashboardUpdateCommand) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expression.Get(), o.Expression.IsSet()
}

// HasExpression returns a boolean if a field has been set.
func (o *PrometheusDashboardUpdateCommand) HasExpression() bool {
	if o != nil && o.Expression.IsSet() {
		return true
	}

	return false
}

// SetExpression gets a reference to the given NullableString and assigns it to the Expression field.
func (o *PrometheusDashboardUpdateCommand) SetExpression(v string) {
	o.Expression.Set(&v)
}
// SetExpressionNil sets the value for Expression to be an explicit nil
func (o *PrometheusDashboardUpdateCommand) SetExpressionNil() {
	o.Expression.Set(nil)
}

// UnsetExpression ensures that no value is present for Expression, not even an explicit nil
func (o *PrometheusDashboardUpdateCommand) UnsetExpression() {
	o.Expression.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusDashboardUpdateCommand) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusDashboardUpdateCommand) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PrometheusDashboardUpdateCommand) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PrometheusDashboardUpdateCommand) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PrometheusDashboardUpdateCommand) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PrometheusDashboardUpdateCommand) UnsetDescription() {
	o.Description.Unset()
}

// GetCategoryName returns the CategoryName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusDashboardUpdateCommand) GetCategoryName() string {
	if o == nil || IsNil(o.CategoryName.Get()) {
		var ret string
		return ret
	}
	return *o.CategoryName.Get()
}

// GetCategoryNameOk returns a tuple with the CategoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusDashboardUpdateCommand) GetCategoryNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CategoryName.Get(), o.CategoryName.IsSet()
}

// HasCategoryName returns a boolean if a field has been set.
func (o *PrometheusDashboardUpdateCommand) HasCategoryName() bool {
	if o != nil && o.CategoryName.IsSet() {
		return true
	}

	return false
}

// SetCategoryName gets a reference to the given NullableString and assigns it to the CategoryName field.
func (o *PrometheusDashboardUpdateCommand) SetCategoryName(v string) {
	o.CategoryName.Set(&v)
}
// SetCategoryNameNil sets the value for CategoryName to be an explicit nil
func (o *PrometheusDashboardUpdateCommand) SetCategoryNameNil() {
	o.CategoryName.Set(nil)
}

// UnsetCategoryName ensures that no value is present for CategoryName, not even an explicit nil
func (o *PrometheusDashboardUpdateCommand) UnsetCategoryName() {
	o.CategoryName.Unset()
}

func (o PrometheusDashboardUpdateCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrometheusDashboardUpdateCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Expression.IsSet() {
		toSerialize["expression"] = o.Expression.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.CategoryName.IsSet() {
		toSerialize["categoryName"] = o.CategoryName.Get()
	}
	return toSerialize, nil
}

type NullablePrometheusDashboardUpdateCommand struct {
	value *PrometheusDashboardUpdateCommand
	isSet bool
}

func (v NullablePrometheusDashboardUpdateCommand) Get() *PrometheusDashboardUpdateCommand {
	return v.value
}

func (v *NullablePrometheusDashboardUpdateCommand) Set(val *PrometheusDashboardUpdateCommand) {
	v.value = val
	v.isSet = true
}

func (v NullablePrometheusDashboardUpdateCommand) IsSet() bool {
	return v.isSet
}

func (v *NullablePrometheusDashboardUpdateCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrometheusDashboardUpdateCommand(val *PrometheusDashboardUpdateCommand) *NullablePrometheusDashboardUpdateCommand {
	return &NullablePrometheusDashboardUpdateCommand{value: val, isSet: true}
}

func (v NullablePrometheusDashboardUpdateCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrometheusDashboardUpdateCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


