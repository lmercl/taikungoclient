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

// checks if the PrometheusBillingSummaryDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrometheusBillingSummaryDto{}

// PrometheusBillingSummaryDto struct for PrometheusBillingSummaryDto
type PrometheusBillingSummaryDto struct {
	Price *float64 `json:"price,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty"`
	EndDate NullableTime `json:"endDate,omitempty"`
	PrometheusRuleId *int32 `json:"prometheusRuleId,omitempty"`
	RuleName NullableString `json:"ruleName,omitempty"`
	CreatedBy NullableString `json:"createdBy,omitempty"`
	LastModified NullableString `json:"lastModified,omitempty"`
	LastModifiedBy NullableString `json:"lastModifiedBy,omitempty"`
}

// NewPrometheusBillingSummaryDto instantiates a new PrometheusBillingSummaryDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrometheusBillingSummaryDto() *PrometheusBillingSummaryDto {
	this := PrometheusBillingSummaryDto{}
	return &this
}

// NewPrometheusBillingSummaryDtoWithDefaults instantiates a new PrometheusBillingSummaryDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrometheusBillingSummaryDtoWithDefaults() *PrometheusBillingSummaryDto {
	this := PrometheusBillingSummaryDto{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *PrometheusBillingSummaryDto) GetPrice() float64 {
	if o == nil || IsNil(o.Price) {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusBillingSummaryDto) GetPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *PrometheusBillingSummaryDto) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *PrometheusBillingSummaryDto) SetPrice(v float64) {
	o.Price = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *PrometheusBillingSummaryDto) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusBillingSummaryDto) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *PrometheusBillingSummaryDto) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *PrometheusBillingSummaryDto) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusBillingSummaryDto) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusBillingSummaryDto) GetEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *PrometheusBillingSummaryDto) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableTime and assigns it to the EndDate field.
func (o *PrometheusBillingSummaryDto) SetEndDate(v time.Time) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *PrometheusBillingSummaryDto) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *PrometheusBillingSummaryDto) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetPrometheusRuleId returns the PrometheusRuleId field value if set, zero value otherwise.
func (o *PrometheusBillingSummaryDto) GetPrometheusRuleId() int32 {
	if o == nil || IsNil(o.PrometheusRuleId) {
		var ret int32
		return ret
	}
	return *o.PrometheusRuleId
}

// GetPrometheusRuleIdOk returns a tuple with the PrometheusRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusBillingSummaryDto) GetPrometheusRuleIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PrometheusRuleId) {
		return nil, false
	}
	return o.PrometheusRuleId, true
}

// HasPrometheusRuleId returns a boolean if a field has been set.
func (o *PrometheusBillingSummaryDto) HasPrometheusRuleId() bool {
	if o != nil && !IsNil(o.PrometheusRuleId) {
		return true
	}

	return false
}

// SetPrometheusRuleId gets a reference to the given int32 and assigns it to the PrometheusRuleId field.
func (o *PrometheusBillingSummaryDto) SetPrometheusRuleId(v int32) {
	o.PrometheusRuleId = &v
}

// GetRuleName returns the RuleName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusBillingSummaryDto) GetRuleName() string {
	if o == nil || IsNil(o.RuleName.Get()) {
		var ret string
		return ret
	}
	return *o.RuleName.Get()
}

// GetRuleNameOk returns a tuple with the RuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusBillingSummaryDto) GetRuleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RuleName.Get(), o.RuleName.IsSet()
}

// HasRuleName returns a boolean if a field has been set.
func (o *PrometheusBillingSummaryDto) HasRuleName() bool {
	if o != nil && o.RuleName.IsSet() {
		return true
	}

	return false
}

// SetRuleName gets a reference to the given NullableString and assigns it to the RuleName field.
func (o *PrometheusBillingSummaryDto) SetRuleName(v string) {
	o.RuleName.Set(&v)
}
// SetRuleNameNil sets the value for RuleName to be an explicit nil
func (o *PrometheusBillingSummaryDto) SetRuleNameNil() {
	o.RuleName.Set(nil)
}

// UnsetRuleName ensures that no value is present for RuleName, not even an explicit nil
func (o *PrometheusBillingSummaryDto) UnsetRuleName() {
	o.RuleName.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusBillingSummaryDto) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusBillingSummaryDto) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *PrometheusBillingSummaryDto) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *PrometheusBillingSummaryDto) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}
// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *PrometheusBillingSummaryDto) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *PrometheusBillingSummaryDto) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetLastModified returns the LastModified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusBillingSummaryDto) GetLastModified() string {
	if o == nil || IsNil(o.LastModified.Get()) {
		var ret string
		return ret
	}
	return *o.LastModified.Get()
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusBillingSummaryDto) GetLastModifiedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModified.Get(), o.LastModified.IsSet()
}

// HasLastModified returns a boolean if a field has been set.
func (o *PrometheusBillingSummaryDto) HasLastModified() bool {
	if o != nil && o.LastModified.IsSet() {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given NullableString and assigns it to the LastModified field.
func (o *PrometheusBillingSummaryDto) SetLastModified(v string) {
	o.LastModified.Set(&v)
}
// SetLastModifiedNil sets the value for LastModified to be an explicit nil
func (o *PrometheusBillingSummaryDto) SetLastModifiedNil() {
	o.LastModified.Set(nil)
}

// UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
func (o *PrometheusBillingSummaryDto) UnsetLastModified() {
	o.LastModified.Unset()
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusBillingSummaryDto) GetLastModifiedBy() string {
	if o == nil || IsNil(o.LastModifiedBy.Get()) {
		var ret string
		return ret
	}
	return *o.LastModifiedBy.Get()
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusBillingSummaryDto) GetLastModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModifiedBy.Get(), o.LastModifiedBy.IsSet()
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *PrometheusBillingSummaryDto) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given NullableString and assigns it to the LastModifiedBy field.
func (o *PrometheusBillingSummaryDto) SetLastModifiedBy(v string) {
	o.LastModifiedBy.Set(&v)
}
// SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil
func (o *PrometheusBillingSummaryDto) SetLastModifiedByNil() {
	o.LastModifiedBy.Set(nil)
}

// UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
func (o *PrometheusBillingSummaryDto) UnsetLastModifiedBy() {
	o.LastModifiedBy.Unset()
}

func (o PrometheusBillingSummaryDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrometheusBillingSummaryDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if o.EndDate.IsSet() {
		toSerialize["endDate"] = o.EndDate.Get()
	}
	if !IsNil(o.PrometheusRuleId) {
		toSerialize["prometheusRuleId"] = o.PrometheusRuleId
	}
	if o.RuleName.IsSet() {
		toSerialize["ruleName"] = o.RuleName.Get()
	}
	if o.CreatedBy.IsSet() {
		toSerialize["createdBy"] = o.CreatedBy.Get()
	}
	if o.LastModified.IsSet() {
		toSerialize["lastModified"] = o.LastModified.Get()
	}
	if o.LastModifiedBy.IsSet() {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy.Get()
	}
	return toSerialize, nil
}

type NullablePrometheusBillingSummaryDto struct {
	value *PrometheusBillingSummaryDto
	isSet bool
}

func (v NullablePrometheusBillingSummaryDto) Get() *PrometheusBillingSummaryDto {
	return v.value
}

func (v *NullablePrometheusBillingSummaryDto) Set(val *PrometheusBillingSummaryDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePrometheusBillingSummaryDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePrometheusBillingSummaryDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrometheusBillingSummaryDto(val *PrometheusBillingSummaryDto) *NullablePrometheusBillingSummaryDto {
	return &NullablePrometheusBillingSummaryDto{value: val, isSet: true}
}

func (v NullablePrometheusBillingSummaryDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrometheusBillingSummaryDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


