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
	"time"
)

// checks if the StripeInvoiceListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StripeInvoiceListDto{}

// StripeInvoiceListDto struct for StripeInvoiceListDto
type StripeInvoiceListDto struct {
	Id NullableString `json:"id,omitempty"`
	InvoiceStatus NullableString `json:"invoiceStatus,omitempty"`
	ChargeStatus NullableString `json:"chargeStatus,omitempty"`
	ChargeReason NullableString `json:"chargeReason,omitempty"`
	Price *float64 `json:"price,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty"`
	EndDate *time.Time `json:"endDate,omitempty"`
}

// NewStripeInvoiceListDto instantiates a new StripeInvoiceListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStripeInvoiceListDto() *StripeInvoiceListDto {
	this := StripeInvoiceListDto{}
	return &this
}

// NewStripeInvoiceListDtoWithDefaults instantiates a new StripeInvoiceListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStripeInvoiceListDtoWithDefaults() *StripeInvoiceListDto {
	this := StripeInvoiceListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StripeInvoiceListDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StripeInvoiceListDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *StripeInvoiceListDto) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *StripeInvoiceListDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *StripeInvoiceListDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *StripeInvoiceListDto) UnsetId() {
	o.Id.Unset()
}

// GetInvoiceStatus returns the InvoiceStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StripeInvoiceListDto) GetInvoiceStatus() string {
	if o == nil || IsNil(o.InvoiceStatus.Get()) {
		var ret string
		return ret
	}
	return *o.InvoiceStatus.Get()
}

// GetInvoiceStatusOk returns a tuple with the InvoiceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StripeInvoiceListDto) GetInvoiceStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceStatus.Get(), o.InvoiceStatus.IsSet()
}

// HasInvoiceStatus returns a boolean if a field has been set.
func (o *StripeInvoiceListDto) HasInvoiceStatus() bool {
	if o != nil && o.InvoiceStatus.IsSet() {
		return true
	}

	return false
}

// SetInvoiceStatus gets a reference to the given NullableString and assigns it to the InvoiceStatus field.
func (o *StripeInvoiceListDto) SetInvoiceStatus(v string) {
	o.InvoiceStatus.Set(&v)
}
// SetInvoiceStatusNil sets the value for InvoiceStatus to be an explicit nil
func (o *StripeInvoiceListDto) SetInvoiceStatusNil() {
	o.InvoiceStatus.Set(nil)
}

// UnsetInvoiceStatus ensures that no value is present for InvoiceStatus, not even an explicit nil
func (o *StripeInvoiceListDto) UnsetInvoiceStatus() {
	o.InvoiceStatus.Unset()
}

// GetChargeStatus returns the ChargeStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StripeInvoiceListDto) GetChargeStatus() string {
	if o == nil || IsNil(o.ChargeStatus.Get()) {
		var ret string
		return ret
	}
	return *o.ChargeStatus.Get()
}

// GetChargeStatusOk returns a tuple with the ChargeStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StripeInvoiceListDto) GetChargeStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChargeStatus.Get(), o.ChargeStatus.IsSet()
}

// HasChargeStatus returns a boolean if a field has been set.
func (o *StripeInvoiceListDto) HasChargeStatus() bool {
	if o != nil && o.ChargeStatus.IsSet() {
		return true
	}

	return false
}

// SetChargeStatus gets a reference to the given NullableString and assigns it to the ChargeStatus field.
func (o *StripeInvoiceListDto) SetChargeStatus(v string) {
	o.ChargeStatus.Set(&v)
}
// SetChargeStatusNil sets the value for ChargeStatus to be an explicit nil
func (o *StripeInvoiceListDto) SetChargeStatusNil() {
	o.ChargeStatus.Set(nil)
}

// UnsetChargeStatus ensures that no value is present for ChargeStatus, not even an explicit nil
func (o *StripeInvoiceListDto) UnsetChargeStatus() {
	o.ChargeStatus.Unset()
}

// GetChargeReason returns the ChargeReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StripeInvoiceListDto) GetChargeReason() string {
	if o == nil || IsNil(o.ChargeReason.Get()) {
		var ret string
		return ret
	}
	return *o.ChargeReason.Get()
}

// GetChargeReasonOk returns a tuple with the ChargeReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StripeInvoiceListDto) GetChargeReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChargeReason.Get(), o.ChargeReason.IsSet()
}

// HasChargeReason returns a boolean if a field has been set.
func (o *StripeInvoiceListDto) HasChargeReason() bool {
	if o != nil && o.ChargeReason.IsSet() {
		return true
	}

	return false
}

// SetChargeReason gets a reference to the given NullableString and assigns it to the ChargeReason field.
func (o *StripeInvoiceListDto) SetChargeReason(v string) {
	o.ChargeReason.Set(&v)
}
// SetChargeReasonNil sets the value for ChargeReason to be an explicit nil
func (o *StripeInvoiceListDto) SetChargeReasonNil() {
	o.ChargeReason.Set(nil)
}

// UnsetChargeReason ensures that no value is present for ChargeReason, not even an explicit nil
func (o *StripeInvoiceListDto) UnsetChargeReason() {
	o.ChargeReason.Unset()
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *StripeInvoiceListDto) GetPrice() float64 {
	if o == nil || IsNil(o.Price) {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeInvoiceListDto) GetPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *StripeInvoiceListDto) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *StripeInvoiceListDto) SetPrice(v float64) {
	o.Price = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *StripeInvoiceListDto) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeInvoiceListDto) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *StripeInvoiceListDto) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *StripeInvoiceListDto) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *StripeInvoiceListDto) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeInvoiceListDto) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *StripeInvoiceListDto) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *StripeInvoiceListDto) SetEndDate(v time.Time) {
	o.EndDate = &v
}

func (o StripeInvoiceListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StripeInvoiceListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.InvoiceStatus.IsSet() {
		toSerialize["invoiceStatus"] = o.InvoiceStatus.Get()
	}
	if o.ChargeStatus.IsSet() {
		toSerialize["chargeStatus"] = o.ChargeStatus.Get()
	}
	if o.ChargeReason.IsSet() {
		toSerialize["chargeReason"] = o.ChargeReason.Get()
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	return toSerialize, nil
}

type NullableStripeInvoiceListDto struct {
	value *StripeInvoiceListDto
	isSet bool
}

func (v NullableStripeInvoiceListDto) Get() *StripeInvoiceListDto {
	return v.value
}

func (v *NullableStripeInvoiceListDto) Set(val *StripeInvoiceListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableStripeInvoiceListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableStripeInvoiceListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStripeInvoiceListDto(val *StripeInvoiceListDto) *NullableStripeInvoiceListDto {
	return &NullableStripeInvoiceListDto{value: val, isSet: true}
}

func (v NullableStripeInvoiceListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStripeInvoiceListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


