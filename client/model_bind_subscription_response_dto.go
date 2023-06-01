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

// checks if the BindSubscriptionResponseDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BindSubscriptionResponseDto{}

// BindSubscriptionResponseDto struct for BindSubscriptionResponseDto
type BindSubscriptionResponseDto struct {
	Status NullableString `json:"status,omitempty"`
	PaymentIntentClientSecret NullableString `json:"paymentIntentClientSecret,omitempty"`
	InvoiceFailureCode NullableString `json:"invoiceFailureCode,omitempty"`
	InvoiceFailureMessage NullableString `json:"invoiceFailureMessage,omitempty"`
	InvoiceFailureReason NullableString `json:"invoiceFailureReason,omitempty"`
}

// NewBindSubscriptionResponseDto instantiates a new BindSubscriptionResponseDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBindSubscriptionResponseDto() *BindSubscriptionResponseDto {
	this := BindSubscriptionResponseDto{}
	return &this
}

// NewBindSubscriptionResponseDtoWithDefaults instantiates a new BindSubscriptionResponseDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBindSubscriptionResponseDtoWithDefaults() *BindSubscriptionResponseDto {
	this := BindSubscriptionResponseDto{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BindSubscriptionResponseDto) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BindSubscriptionResponseDto) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *BindSubscriptionResponseDto) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *BindSubscriptionResponseDto) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *BindSubscriptionResponseDto) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *BindSubscriptionResponseDto) UnsetStatus() {
	o.Status.Unset()
}

// GetPaymentIntentClientSecret returns the PaymentIntentClientSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BindSubscriptionResponseDto) GetPaymentIntentClientSecret() string {
	if o == nil || IsNil(o.PaymentIntentClientSecret.Get()) {
		var ret string
		return ret
	}
	return *o.PaymentIntentClientSecret.Get()
}

// GetPaymentIntentClientSecretOk returns a tuple with the PaymentIntentClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BindSubscriptionResponseDto) GetPaymentIntentClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentIntentClientSecret.Get(), o.PaymentIntentClientSecret.IsSet()
}

// HasPaymentIntentClientSecret returns a boolean if a field has been set.
func (o *BindSubscriptionResponseDto) HasPaymentIntentClientSecret() bool {
	if o != nil && o.PaymentIntentClientSecret.IsSet() {
		return true
	}

	return false
}

// SetPaymentIntentClientSecret gets a reference to the given NullableString and assigns it to the PaymentIntentClientSecret field.
func (o *BindSubscriptionResponseDto) SetPaymentIntentClientSecret(v string) {
	o.PaymentIntentClientSecret.Set(&v)
}
// SetPaymentIntentClientSecretNil sets the value for PaymentIntentClientSecret to be an explicit nil
func (o *BindSubscriptionResponseDto) SetPaymentIntentClientSecretNil() {
	o.PaymentIntentClientSecret.Set(nil)
}

// UnsetPaymentIntentClientSecret ensures that no value is present for PaymentIntentClientSecret, not even an explicit nil
func (o *BindSubscriptionResponseDto) UnsetPaymentIntentClientSecret() {
	o.PaymentIntentClientSecret.Unset()
}

// GetInvoiceFailureCode returns the InvoiceFailureCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BindSubscriptionResponseDto) GetInvoiceFailureCode() string {
	if o == nil || IsNil(o.InvoiceFailureCode.Get()) {
		var ret string
		return ret
	}
	return *o.InvoiceFailureCode.Get()
}

// GetInvoiceFailureCodeOk returns a tuple with the InvoiceFailureCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BindSubscriptionResponseDto) GetInvoiceFailureCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceFailureCode.Get(), o.InvoiceFailureCode.IsSet()
}

// HasInvoiceFailureCode returns a boolean if a field has been set.
func (o *BindSubscriptionResponseDto) HasInvoiceFailureCode() bool {
	if o != nil && o.InvoiceFailureCode.IsSet() {
		return true
	}

	return false
}

// SetInvoiceFailureCode gets a reference to the given NullableString and assigns it to the InvoiceFailureCode field.
func (o *BindSubscriptionResponseDto) SetInvoiceFailureCode(v string) {
	o.InvoiceFailureCode.Set(&v)
}
// SetInvoiceFailureCodeNil sets the value for InvoiceFailureCode to be an explicit nil
func (o *BindSubscriptionResponseDto) SetInvoiceFailureCodeNil() {
	o.InvoiceFailureCode.Set(nil)
}

// UnsetInvoiceFailureCode ensures that no value is present for InvoiceFailureCode, not even an explicit nil
func (o *BindSubscriptionResponseDto) UnsetInvoiceFailureCode() {
	o.InvoiceFailureCode.Unset()
}

// GetInvoiceFailureMessage returns the InvoiceFailureMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BindSubscriptionResponseDto) GetInvoiceFailureMessage() string {
	if o == nil || IsNil(o.InvoiceFailureMessage.Get()) {
		var ret string
		return ret
	}
	return *o.InvoiceFailureMessage.Get()
}

// GetInvoiceFailureMessageOk returns a tuple with the InvoiceFailureMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BindSubscriptionResponseDto) GetInvoiceFailureMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceFailureMessage.Get(), o.InvoiceFailureMessage.IsSet()
}

// HasInvoiceFailureMessage returns a boolean if a field has been set.
func (o *BindSubscriptionResponseDto) HasInvoiceFailureMessage() bool {
	if o != nil && o.InvoiceFailureMessage.IsSet() {
		return true
	}

	return false
}

// SetInvoiceFailureMessage gets a reference to the given NullableString and assigns it to the InvoiceFailureMessage field.
func (o *BindSubscriptionResponseDto) SetInvoiceFailureMessage(v string) {
	o.InvoiceFailureMessage.Set(&v)
}
// SetInvoiceFailureMessageNil sets the value for InvoiceFailureMessage to be an explicit nil
func (o *BindSubscriptionResponseDto) SetInvoiceFailureMessageNil() {
	o.InvoiceFailureMessage.Set(nil)
}

// UnsetInvoiceFailureMessage ensures that no value is present for InvoiceFailureMessage, not even an explicit nil
func (o *BindSubscriptionResponseDto) UnsetInvoiceFailureMessage() {
	o.InvoiceFailureMessage.Unset()
}

// GetInvoiceFailureReason returns the InvoiceFailureReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BindSubscriptionResponseDto) GetInvoiceFailureReason() string {
	if o == nil || IsNil(o.InvoiceFailureReason.Get()) {
		var ret string
		return ret
	}
	return *o.InvoiceFailureReason.Get()
}

// GetInvoiceFailureReasonOk returns a tuple with the InvoiceFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BindSubscriptionResponseDto) GetInvoiceFailureReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceFailureReason.Get(), o.InvoiceFailureReason.IsSet()
}

// HasInvoiceFailureReason returns a boolean if a field has been set.
func (o *BindSubscriptionResponseDto) HasInvoiceFailureReason() bool {
	if o != nil && o.InvoiceFailureReason.IsSet() {
		return true
	}

	return false
}

// SetInvoiceFailureReason gets a reference to the given NullableString and assigns it to the InvoiceFailureReason field.
func (o *BindSubscriptionResponseDto) SetInvoiceFailureReason(v string) {
	o.InvoiceFailureReason.Set(&v)
}
// SetInvoiceFailureReasonNil sets the value for InvoiceFailureReason to be an explicit nil
func (o *BindSubscriptionResponseDto) SetInvoiceFailureReasonNil() {
	o.InvoiceFailureReason.Set(nil)
}

// UnsetInvoiceFailureReason ensures that no value is present for InvoiceFailureReason, not even an explicit nil
func (o *BindSubscriptionResponseDto) UnsetInvoiceFailureReason() {
	o.InvoiceFailureReason.Unset()
}

func (o BindSubscriptionResponseDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BindSubscriptionResponseDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.PaymentIntentClientSecret.IsSet() {
		toSerialize["paymentIntentClientSecret"] = o.PaymentIntentClientSecret.Get()
	}
	if o.InvoiceFailureCode.IsSet() {
		toSerialize["invoiceFailureCode"] = o.InvoiceFailureCode.Get()
	}
	if o.InvoiceFailureMessage.IsSet() {
		toSerialize["invoiceFailureMessage"] = o.InvoiceFailureMessage.Get()
	}
	if o.InvoiceFailureReason.IsSet() {
		toSerialize["invoiceFailureReason"] = o.InvoiceFailureReason.Get()
	}
	return toSerialize, nil
}

type NullableBindSubscriptionResponseDto struct {
	value *BindSubscriptionResponseDto
	isSet bool
}

func (v NullableBindSubscriptionResponseDto) Get() *BindSubscriptionResponseDto {
	return v.value
}

func (v *NullableBindSubscriptionResponseDto) Set(val *BindSubscriptionResponseDto) {
	v.value = val
	v.isSet = true
}

func (v NullableBindSubscriptionResponseDto) IsSet() bool {
	return v.isSet
}

func (v *NullableBindSubscriptionResponseDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBindSubscriptionResponseDto(val *BindSubscriptionResponseDto) *NullableBindSubscriptionResponseDto {
	return &NullableBindSubscriptionResponseDto{value: val, isSet: true}
}

func (v NullableBindSubscriptionResponseDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBindSubscriptionResponseDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


