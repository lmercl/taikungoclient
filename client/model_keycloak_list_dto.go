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

// checks if the KeycloakListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeycloakListDto{}

// KeycloakListDto struct for KeycloakListDto
type KeycloakListDto struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Url NullableString `json:"url,omitempty"`
	ClientId NullableString `json:"clientId,omitempty"`
	ClientSecret NullableString `json:"clientSecret,omitempty"`
	RealmsName NullableString `json:"realmsName,omitempty"`
	OrganizationId *int32 `json:"organizationId,omitempty"`
	OrganizationName NullableString `json:"organizationName,omitempty"`
	PartnerLogo NullableString `json:"partnerLogo,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
}

// NewKeycloakListDto instantiates a new KeycloakListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeycloakListDto() *KeycloakListDto {
	this := KeycloakListDto{}
	return &this
}

// NewKeycloakListDtoWithDefaults instantiates a new KeycloakListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeycloakListDtoWithDefaults() *KeycloakListDto {
	this := KeycloakListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KeycloakListDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeycloakListDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KeycloakListDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KeycloakListDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeycloakListDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeycloakListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *KeycloakListDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *KeycloakListDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *KeycloakListDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *KeycloakListDto) UnsetName() {
	o.Name.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeycloakListDto) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeycloakListDto) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *KeycloakListDto) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *KeycloakListDto) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *KeycloakListDto) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *KeycloakListDto) UnsetUrl() {
	o.Url.Unset()
}

// GetClientId returns the ClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeycloakListDto) GetClientId() string {
	if o == nil || IsNil(o.ClientId.Get()) {
		var ret string
		return ret
	}
	return *o.ClientId.Get()
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeycloakListDto) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientId.Get(), o.ClientId.IsSet()
}

// HasClientId returns a boolean if a field has been set.
func (o *KeycloakListDto) HasClientId() bool {
	if o != nil && o.ClientId.IsSet() {
		return true
	}

	return false
}

// SetClientId gets a reference to the given NullableString and assigns it to the ClientId field.
func (o *KeycloakListDto) SetClientId(v string) {
	o.ClientId.Set(&v)
}
// SetClientIdNil sets the value for ClientId to be an explicit nil
func (o *KeycloakListDto) SetClientIdNil() {
	o.ClientId.Set(nil)
}

// UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
func (o *KeycloakListDto) UnsetClientId() {
	o.ClientId.Unset()
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeycloakListDto) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret.Get()) {
		var ret string
		return ret
	}
	return *o.ClientSecret.Get()
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeycloakListDto) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientSecret.Get(), o.ClientSecret.IsSet()
}

// HasClientSecret returns a boolean if a field has been set.
func (o *KeycloakListDto) HasClientSecret() bool {
	if o != nil && o.ClientSecret.IsSet() {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given NullableString and assigns it to the ClientSecret field.
func (o *KeycloakListDto) SetClientSecret(v string) {
	o.ClientSecret.Set(&v)
}
// SetClientSecretNil sets the value for ClientSecret to be an explicit nil
func (o *KeycloakListDto) SetClientSecretNil() {
	o.ClientSecret.Set(nil)
}

// UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
func (o *KeycloakListDto) UnsetClientSecret() {
	o.ClientSecret.Unset()
}

// GetRealmsName returns the RealmsName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeycloakListDto) GetRealmsName() string {
	if o == nil || IsNil(o.RealmsName.Get()) {
		var ret string
		return ret
	}
	return *o.RealmsName.Get()
}

// GetRealmsNameOk returns a tuple with the RealmsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeycloakListDto) GetRealmsNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RealmsName.Get(), o.RealmsName.IsSet()
}

// HasRealmsName returns a boolean if a field has been set.
func (o *KeycloakListDto) HasRealmsName() bool {
	if o != nil && o.RealmsName.IsSet() {
		return true
	}

	return false
}

// SetRealmsName gets a reference to the given NullableString and assigns it to the RealmsName field.
func (o *KeycloakListDto) SetRealmsName(v string) {
	o.RealmsName.Set(&v)
}
// SetRealmsNameNil sets the value for RealmsName to be an explicit nil
func (o *KeycloakListDto) SetRealmsNameNil() {
	o.RealmsName.Set(nil)
}

// UnsetRealmsName ensures that no value is present for RealmsName, not even an explicit nil
func (o *KeycloakListDto) UnsetRealmsName() {
	o.RealmsName.Unset()
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *KeycloakListDto) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId) {
		var ret int32
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeycloakListDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *KeycloakListDto) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int32 and assigns it to the OrganizationId field.
func (o *KeycloakListDto) SetOrganizationId(v int32) {
	o.OrganizationId = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeycloakListDto) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationName.Get()
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeycloakListDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationName.Get(), o.OrganizationName.IsSet()
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *KeycloakListDto) HasOrganizationName() bool {
	if o != nil && o.OrganizationName.IsSet() {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given NullableString and assigns it to the OrganizationName field.
func (o *KeycloakListDto) SetOrganizationName(v string) {
	o.OrganizationName.Set(&v)
}
// SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil
func (o *KeycloakListDto) SetOrganizationNameNil() {
	o.OrganizationName.Set(nil)
}

// UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
func (o *KeycloakListDto) UnsetOrganizationName() {
	o.OrganizationName.Unset()
}

// GetPartnerLogo returns the PartnerLogo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeycloakListDto) GetPartnerLogo() string {
	if o == nil || IsNil(o.PartnerLogo.Get()) {
		var ret string
		return ret
	}
	return *o.PartnerLogo.Get()
}

// GetPartnerLogoOk returns a tuple with the PartnerLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeycloakListDto) GetPartnerLogoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PartnerLogo.Get(), o.PartnerLogo.IsSet()
}

// HasPartnerLogo returns a boolean if a field has been set.
func (o *KeycloakListDto) HasPartnerLogo() bool {
	if o != nil && o.PartnerLogo.IsSet() {
		return true
	}

	return false
}

// SetPartnerLogo gets a reference to the given NullableString and assigns it to the PartnerLogo field.
func (o *KeycloakListDto) SetPartnerLogo(v string) {
	o.PartnerLogo.Set(&v)
}
// SetPartnerLogoNil sets the value for PartnerLogo to be an explicit nil
func (o *KeycloakListDto) SetPartnerLogoNil() {
	o.PartnerLogo.Set(nil)
}

// UnsetPartnerLogo ensures that no value is present for PartnerLogo, not even an explicit nil
func (o *KeycloakListDto) UnsetPartnerLogo() {
	o.PartnerLogo.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *KeycloakListDto) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeycloakListDto) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *KeycloakListDto) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *KeycloakListDto) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o KeycloakListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeycloakListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.ClientId.IsSet() {
		toSerialize["clientId"] = o.ClientId.Get()
	}
	if o.ClientSecret.IsSet() {
		toSerialize["clientSecret"] = o.ClientSecret.Get()
	}
	if o.RealmsName.IsSet() {
		toSerialize["realmsName"] = o.RealmsName.Get()
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if o.OrganizationName.IsSet() {
		toSerialize["organizationName"] = o.OrganizationName.Get()
	}
	if o.PartnerLogo.IsSet() {
		toSerialize["partnerLogo"] = o.PartnerLogo.Get()
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableKeycloakListDto struct {
	value *KeycloakListDto
	isSet bool
}

func (v NullableKeycloakListDto) Get() *KeycloakListDto {
	return v.value
}

func (v *NullableKeycloakListDto) Set(val *KeycloakListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableKeycloakListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableKeycloakListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeycloakListDto(val *KeycloakListDto) *NullableKeycloakListDto {
	return &NullableKeycloakListDto{value: val, isSet: true}
}

func (v NullableKeycloakListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeycloakListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


