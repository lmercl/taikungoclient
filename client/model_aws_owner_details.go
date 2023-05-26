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

// checks if the AwsOwnerDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsOwnerDetails{}

// AwsOwnerDetails struct for AwsOwnerDetails
type AwsOwnerDetails struct {
	OwnerId NullableString `json:"ownerId,omitempty"`
	OwnerName NullableString `json:"ownerName,omitempty"`
	Image *AwsCommonImages `json:"image,omitempty"`
}

// NewAwsOwnerDetails instantiates a new AwsOwnerDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsOwnerDetails() *AwsOwnerDetails {
	this := AwsOwnerDetails{}
	return &this
}

// NewAwsOwnerDetailsWithDefaults instantiates a new AwsOwnerDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsOwnerDetailsWithDefaults() *AwsOwnerDetails {
	this := AwsOwnerDetails{}
	return &this
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsOwnerDetails) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId.Get()) {
		var ret string
		return ret
	}
	return *o.OwnerId.Get()
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsOwnerDetails) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerId.Get(), o.OwnerId.IsSet()
}

// HasOwnerId returns a boolean if a field has been set.
func (o *AwsOwnerDetails) HasOwnerId() bool {
	if o != nil && o.OwnerId.IsSet() {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given NullableString and assigns it to the OwnerId field.
func (o *AwsOwnerDetails) SetOwnerId(v string) {
	o.OwnerId.Set(&v)
}
// SetOwnerIdNil sets the value for OwnerId to be an explicit nil
func (o *AwsOwnerDetails) SetOwnerIdNil() {
	o.OwnerId.Set(nil)
}

// UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
func (o *AwsOwnerDetails) UnsetOwnerId() {
	o.OwnerId.Unset()
}

// GetOwnerName returns the OwnerName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsOwnerDetails) GetOwnerName() string {
	if o == nil || IsNil(o.OwnerName.Get()) {
		var ret string
		return ret
	}
	return *o.OwnerName.Get()
}

// GetOwnerNameOk returns a tuple with the OwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsOwnerDetails) GetOwnerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerName.Get(), o.OwnerName.IsSet()
}

// HasOwnerName returns a boolean if a field has been set.
func (o *AwsOwnerDetails) HasOwnerName() bool {
	if o != nil && o.OwnerName.IsSet() {
		return true
	}

	return false
}

// SetOwnerName gets a reference to the given NullableString and assigns it to the OwnerName field.
func (o *AwsOwnerDetails) SetOwnerName(v string) {
	o.OwnerName.Set(&v)
}
// SetOwnerNameNil sets the value for OwnerName to be an explicit nil
func (o *AwsOwnerDetails) SetOwnerNameNil() {
	o.OwnerName.Set(nil)
}

// UnsetOwnerName ensures that no value is present for OwnerName, not even an explicit nil
func (o *AwsOwnerDetails) UnsetOwnerName() {
	o.OwnerName.Unset()
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *AwsOwnerDetails) GetImage() AwsCommonImages {
	if o == nil || IsNil(o.Image) {
		var ret AwsCommonImages
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsOwnerDetails) GetImageOk() (*AwsCommonImages, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *AwsOwnerDetails) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given AwsCommonImages and assigns it to the Image field.
func (o *AwsOwnerDetails) SetImage(v AwsCommonImages) {
	o.Image = &v
}

func (o AwsOwnerDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsOwnerDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.OwnerId.IsSet() {
		toSerialize["ownerId"] = o.OwnerId.Get()
	}
	if o.OwnerName.IsSet() {
		toSerialize["ownerName"] = o.OwnerName.Get()
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	return toSerialize, nil
}

type NullableAwsOwnerDetails struct {
	value *AwsOwnerDetails
	isSet bool
}

func (v NullableAwsOwnerDetails) Get() *AwsOwnerDetails {
	return v.value
}

func (v *NullableAwsOwnerDetails) Set(val *AwsOwnerDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsOwnerDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsOwnerDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsOwnerDetails(val *AwsOwnerDetails) *NullableAwsOwnerDetails {
	return &NullableAwsOwnerDetails{value: val, isSet: true}
}

func (v NullableAwsOwnerDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsOwnerDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


