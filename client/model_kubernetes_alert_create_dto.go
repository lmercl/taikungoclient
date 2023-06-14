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
	"time"
)

// checks if the KubernetesAlertCreateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesAlertCreateDto{}

// KubernetesAlertCreateDto struct for KubernetesAlertCreateDto
type KubernetesAlertCreateDto struct {
	Status NullableString `json:"status,omitempty"`
	Labels interface{} `json:"labels,omitempty"`
	Annotations *Annotations `json:"annotations,omitempty"`
	StartsAt *time.Time `json:"startsAt,omitempty"`
	EndsAt *time.Time `json:"endsAt,omitempty"`
	Fingerprint NullableString `json:"fingerprint,omitempty"`
}

// NewKubernetesAlertCreateDto instantiates a new KubernetesAlertCreateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesAlertCreateDto() *KubernetesAlertCreateDto {
	this := KubernetesAlertCreateDto{}
	return &this
}

// NewKubernetesAlertCreateDtoWithDefaults instantiates a new KubernetesAlertCreateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesAlertCreateDtoWithDefaults() *KubernetesAlertCreateDto {
	this := KubernetesAlertCreateDto{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAlertCreateDto) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAlertCreateDto) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *KubernetesAlertCreateDto) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *KubernetesAlertCreateDto) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *KubernetesAlertCreateDto) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *KubernetesAlertCreateDto) UnsetStatus() {
	o.Status.Unset()
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAlertCreateDto) GetLabels() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAlertCreateDto) GetLabelsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *KubernetesAlertCreateDto) HasLabels() bool {
	if o != nil && IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given interface{} and assigns it to the Labels field.
func (o *KubernetesAlertCreateDto) SetLabels(v interface{}) {
	o.Labels = v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *KubernetesAlertCreateDto) GetAnnotations() Annotations {
	if o == nil || IsNil(o.Annotations) {
		var ret Annotations
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAlertCreateDto) GetAnnotationsOk() (*Annotations, bool) {
	if o == nil || IsNil(o.Annotations) {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *KubernetesAlertCreateDto) HasAnnotations() bool {
	if o != nil && !IsNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given Annotations and assigns it to the Annotations field.
func (o *KubernetesAlertCreateDto) SetAnnotations(v Annotations) {
	o.Annotations = &v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise.
func (o *KubernetesAlertCreateDto) GetStartsAt() time.Time {
	if o == nil || IsNil(o.StartsAt) {
		var ret time.Time
		return ret
	}
	return *o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAlertCreateDto) GetStartsAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartsAt) {
		return nil, false
	}
	return o.StartsAt, true
}

// HasStartsAt returns a boolean if a field has been set.
func (o *KubernetesAlertCreateDto) HasStartsAt() bool {
	if o != nil && !IsNil(o.StartsAt) {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given time.Time and assigns it to the StartsAt field.
func (o *KubernetesAlertCreateDto) SetStartsAt(v time.Time) {
	o.StartsAt = &v
}

// GetEndsAt returns the EndsAt field value if set, zero value otherwise.
func (o *KubernetesAlertCreateDto) GetEndsAt() time.Time {
	if o == nil || IsNil(o.EndsAt) {
		var ret time.Time
		return ret
	}
	return *o.EndsAt
}

// GetEndsAtOk returns a tuple with the EndsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAlertCreateDto) GetEndsAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndsAt) {
		return nil, false
	}
	return o.EndsAt, true
}

// HasEndsAt returns a boolean if a field has been set.
func (o *KubernetesAlertCreateDto) HasEndsAt() bool {
	if o != nil && !IsNil(o.EndsAt) {
		return true
	}

	return false
}

// SetEndsAt gets a reference to the given time.Time and assigns it to the EndsAt field.
func (o *KubernetesAlertCreateDto) SetEndsAt(v time.Time) {
	o.EndsAt = &v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAlertCreateDto) GetFingerprint() string {
	if o == nil || IsNil(o.Fingerprint.Get()) {
		var ret string
		return ret
	}
	return *o.Fingerprint.Get()
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAlertCreateDto) GetFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fingerprint.Get(), o.Fingerprint.IsSet()
}

// HasFingerprint returns a boolean if a field has been set.
func (o *KubernetesAlertCreateDto) HasFingerprint() bool {
	if o != nil && o.Fingerprint.IsSet() {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given NullableString and assigns it to the Fingerprint field.
func (o *KubernetesAlertCreateDto) SetFingerprint(v string) {
	o.Fingerprint.Set(&v)
}
// SetFingerprintNil sets the value for Fingerprint to be an explicit nil
func (o *KubernetesAlertCreateDto) SetFingerprintNil() {
	o.Fingerprint.Set(nil)
}

// UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
func (o *KubernetesAlertCreateDto) UnsetFingerprint() {
	o.Fingerprint.Unset()
}

func (o KubernetesAlertCreateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesAlertCreateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Annotations) {
		toSerialize["annotations"] = o.Annotations
	}
	if !IsNil(o.StartsAt) {
		toSerialize["startsAt"] = o.StartsAt
	}
	if !IsNil(o.EndsAt) {
		toSerialize["endsAt"] = o.EndsAt
	}
	if o.Fingerprint.IsSet() {
		toSerialize["fingerprint"] = o.Fingerprint.Get()
	}
	return toSerialize, nil
}

type NullableKubernetesAlertCreateDto struct {
	value *KubernetesAlertCreateDto
	isSet bool
}

func (v NullableKubernetesAlertCreateDto) Get() *KubernetesAlertCreateDto {
	return v.value
}

func (v *NullableKubernetesAlertCreateDto) Set(val *KubernetesAlertCreateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesAlertCreateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesAlertCreateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesAlertCreateDto(val *KubernetesAlertCreateDto) *NullableKubernetesAlertCreateDto {
	return &NullableKubernetesAlertCreateDto{value: val, isSet: true}
}

func (v NullableKubernetesAlertCreateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesAlertCreateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


