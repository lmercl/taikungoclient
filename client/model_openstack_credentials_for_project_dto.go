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

// checks if the OpenstackCredentialsForProjectDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenstackCredentialsForProjectDto{}

// OpenstackCredentialsForProjectDto struct for OpenstackCredentialsForProjectDto
type OpenstackCredentialsForProjectDto struct {
	OpenStackUser NullableString `json:"openStackUser,omitempty"`
	OpenStackDomain NullableString `json:"openStackDomain,omitempty"`
	OpenStackPassword NullableString `json:"openStackPassword,omitempty"`
	OpenStackProject NullableString `json:"openStackProject,omitempty"`
	OpenStackRegion NullableString `json:"openStackRegion,omitempty"`
	OpenStackUrl NullableString `json:"openStackUrl,omitempty"`
	OpenStackPublicNetwork NullableString `json:"openStackPublicNetwork,omitempty"`
	OpenStackAvailabilityZone NullableString `json:"openStackAvailabilityZone,omitempty"`
	OpenStackVolumeType NullableString `json:"openStackVolumeType,omitempty"`
	OpenStackTenantId NullableString `json:"openStackTenantId,omitempty"`
	OpenStackImportNetwork *bool `json:"openStackImportNetwork,omitempty"`
	OpenStackInternalSubnetId NullableString `json:"openStackInternalSubnetId,omitempty"`
	ApplicationCredEnabled *bool `json:"applicationCredEnabled,omitempty"`
}

// NewOpenstackCredentialsForProjectDto instantiates a new OpenstackCredentialsForProjectDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenstackCredentialsForProjectDto() *OpenstackCredentialsForProjectDto {
	this := OpenstackCredentialsForProjectDto{}
	return &this
}

// NewOpenstackCredentialsForProjectDtoWithDefaults instantiates a new OpenstackCredentialsForProjectDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenstackCredentialsForProjectDtoWithDefaults() *OpenstackCredentialsForProjectDto {
	this := OpenstackCredentialsForProjectDto{}
	return &this
}

// GetOpenStackUser returns the OpenStackUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackCredentialsForProjectDto) GetOpenStackUser() string {
	if o == nil || IsNil(o.OpenStackUser.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackUser.Get()
}

// GetOpenStackUserOk returns a tuple with the OpenStackUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackCredentialsForProjectDto) GetOpenStackUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackUser.Get(), o.OpenStackUser.IsSet()
}

// HasOpenStackUser returns a boolean if a field has been set.
func (o *OpenstackCredentialsForProjectDto) HasOpenStackUser() bool {
	if o != nil && o.OpenStackUser.IsSet() {
		return true
	}

	return false
}

// SetOpenStackUser gets a reference to the given NullableString and assigns it to the OpenStackUser field.
func (o *OpenstackCredentialsForProjectDto) SetOpenStackUser(v string) {
	o.OpenStackUser.Set(&v)
}
// SetOpenStackUserNil sets the value for OpenStackUser to be an explicit nil
func (o *OpenstackCredentialsForProjectDto) SetOpenStackUserNil() {
	o.OpenStackUser.Set(nil)
}

// UnsetOpenStackUser ensures that no value is present for OpenStackUser, not even an explicit nil
func (o *OpenstackCredentialsForProjectDto) UnsetOpenStackUser() {
	o.OpenStackUser.Unset()
}

// GetOpenStackDomain returns the OpenStackDomain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackCredentialsForProjectDto) GetOpenStackDomain() string {
	if o == nil || IsNil(o.OpenStackDomain.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackDomain.Get()
}

// GetOpenStackDomainOk returns a tuple with the OpenStackDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackCredentialsForProjectDto) GetOpenStackDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackDomain.Get(), o.OpenStackDomain.IsSet()
}

// HasOpenStackDomain returns a boolean if a field has been set.
func (o *OpenstackCredentialsForProjectDto) HasOpenStackDomain() bool {
	if o != nil && o.OpenStackDomain.IsSet() {
		return true
	}

	return false
}

// SetOpenStackDomain gets a reference to the given NullableString and assigns it to the OpenStackDomain field.
func (o *OpenstackCredentialsForProjectDto) SetOpenStackDomain(v string) {
	o.OpenStackDomain.Set(&v)
}
// SetOpenStackDomainNil sets the value for OpenStackDomain to be an explicit nil
func (o *OpenstackCredentialsForProjectDto) SetOpenStackDomainNil() {
	o.OpenStackDomain.Set(nil)
}

// UnsetOpenStackDomain ensures that no value is present for OpenStackDomain, not even an explicit nil
func (o *OpenstackCredentialsForProjectDto) UnsetOpenStackDomain() {
	o.OpenStackDomain.Unset()
}

// GetOpenStackPassword returns the OpenStackPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackCredentialsForProjectDto) GetOpenStackPassword() string {
	if o == nil || IsNil(o.OpenStackPassword.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackPassword.Get()
}

// GetOpenStackPasswordOk returns a tuple with the OpenStackPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackCredentialsForProjectDto) GetOpenStackPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackPassword.Get(), o.OpenStackPassword.IsSet()
}

// HasOpenStackPassword returns a boolean if a field has been set.
func (o *OpenstackCredentialsForProjectDto) HasOpenStackPassword() bool {
	if o != nil && o.OpenStackPassword.IsSet() {
		return true
	}

	return false
}

// SetOpenStackPassword gets a reference to the given NullableString and assigns it to the OpenStackPassword field.
func (o *OpenstackCredentialsForProjectDto) SetOpenStackPassword(v string) {
	o.OpenStackPassword.Set(&v)
}
// SetOpenStackPasswordNil sets the value for OpenStackPassword to be an explicit nil
func (o *OpenstackCredentialsForProjectDto) SetOpenStackPasswordNil() {
	o.OpenStackPassword.Set(nil)
}

// UnsetOpenStackPassword ensures that no value is present for OpenStackPassword, not even an explicit nil
func (o *OpenstackCredentialsForProjectDto) UnsetOpenStackPassword() {
	o.OpenStackPassword.Unset()
}

// GetOpenStackProject returns the OpenStackProject field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackCredentialsForProjectDto) GetOpenStackProject() string {
	if o == nil || IsNil(o.OpenStackProject.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackProject.Get()
}

// GetOpenStackProjectOk returns a tuple with the OpenStackProject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackCredentialsForProjectDto) GetOpenStackProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackProject.Get(), o.OpenStackProject.IsSet()
}

// HasOpenStackProject returns a boolean if a field has been set.
func (o *OpenstackCredentialsForProjectDto) HasOpenStackProject() bool {
	if o != nil && o.OpenStackProject.IsSet() {
		return true
	}

	return false
}

// SetOpenStackProject gets a reference to the given NullableString and assigns it to the OpenStackProject field.
func (o *OpenstackCredentialsForProjectDto) SetOpenStackProject(v string) {
	o.OpenStackProject.Set(&v)
}
// SetOpenStackProjectNil sets the value for OpenStackProject to be an explicit nil
func (o *OpenstackCredentialsForProjectDto) SetOpenStackProjectNil() {
	o.OpenStackProject.Set(nil)
}

// UnsetOpenStackProject ensures that no value is present for OpenStackProject, not even an explicit nil
func (o *OpenstackCredentialsForProjectDto) UnsetOpenStackProject() {
	o.OpenStackProject.Unset()
}

// GetOpenStackRegion returns the OpenStackRegion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackCredentialsForProjectDto) GetOpenStackRegion() string {
	if o == nil || IsNil(o.OpenStackRegion.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackRegion.Get()
}

// GetOpenStackRegionOk returns a tuple with the OpenStackRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackCredentialsForProjectDto) GetOpenStackRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackRegion.Get(), o.OpenStackRegion.IsSet()
}

// HasOpenStackRegion returns a boolean if a field has been set.
func (o *OpenstackCredentialsForProjectDto) HasOpenStackRegion() bool {
	if o != nil && o.OpenStackRegion.IsSet() {
		return true
	}

	return false
}

// SetOpenStackRegion gets a reference to the given NullableString and assigns it to the OpenStackRegion field.
func (o *OpenstackCredentialsForProjectDto) SetOpenStackRegion(v string) {
	o.OpenStackRegion.Set(&v)
}
// SetOpenStackRegionNil sets the value for OpenStackRegion to be an explicit nil
func (o *OpenstackCredentialsForProjectDto) SetOpenStackRegionNil() {
	o.OpenStackRegion.Set(nil)
}

// UnsetOpenStackRegion ensures that no value is present for OpenStackRegion, not even an explicit nil
func (o *OpenstackCredentialsForProjectDto) UnsetOpenStackRegion() {
	o.OpenStackRegion.Unset()
}

// GetOpenStackUrl returns the OpenStackUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackCredentialsForProjectDto) GetOpenStackUrl() string {
	if o == nil || IsNil(o.OpenStackUrl.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackUrl.Get()
}

// GetOpenStackUrlOk returns a tuple with the OpenStackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackCredentialsForProjectDto) GetOpenStackUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackUrl.Get(), o.OpenStackUrl.IsSet()
}

// HasOpenStackUrl returns a boolean if a field has been set.
func (o *OpenstackCredentialsForProjectDto) HasOpenStackUrl() bool {
	if o != nil && o.OpenStackUrl.IsSet() {
		return true
	}

	return false
}

// SetOpenStackUrl gets a reference to the given NullableString and assigns it to the OpenStackUrl field.
func (o *OpenstackCredentialsForProjectDto) SetOpenStackUrl(v string) {
	o.OpenStackUrl.Set(&v)
}
// SetOpenStackUrlNil sets the value for OpenStackUrl to be an explicit nil
func (o *OpenstackCredentialsForProjectDto) SetOpenStackUrlNil() {
	o.OpenStackUrl.Set(nil)
}

// UnsetOpenStackUrl ensures that no value is present for OpenStackUrl, not even an explicit nil
func (o *OpenstackCredentialsForProjectDto) UnsetOpenStackUrl() {
	o.OpenStackUrl.Unset()
}

// GetOpenStackPublicNetwork returns the OpenStackPublicNetwork field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackCredentialsForProjectDto) GetOpenStackPublicNetwork() string {
	if o == nil || IsNil(o.OpenStackPublicNetwork.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackPublicNetwork.Get()
}

// GetOpenStackPublicNetworkOk returns a tuple with the OpenStackPublicNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackCredentialsForProjectDto) GetOpenStackPublicNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackPublicNetwork.Get(), o.OpenStackPublicNetwork.IsSet()
}

// HasOpenStackPublicNetwork returns a boolean if a field has been set.
func (o *OpenstackCredentialsForProjectDto) HasOpenStackPublicNetwork() bool {
	if o != nil && o.OpenStackPublicNetwork.IsSet() {
		return true
	}

	return false
}

// SetOpenStackPublicNetwork gets a reference to the given NullableString and assigns it to the OpenStackPublicNetwork field.
func (o *OpenstackCredentialsForProjectDto) SetOpenStackPublicNetwork(v string) {
	o.OpenStackPublicNetwork.Set(&v)
}
// SetOpenStackPublicNetworkNil sets the value for OpenStackPublicNetwork to be an explicit nil
func (o *OpenstackCredentialsForProjectDto) SetOpenStackPublicNetworkNil() {
	o.OpenStackPublicNetwork.Set(nil)
}

// UnsetOpenStackPublicNetwork ensures that no value is present for OpenStackPublicNetwork, not even an explicit nil
func (o *OpenstackCredentialsForProjectDto) UnsetOpenStackPublicNetwork() {
	o.OpenStackPublicNetwork.Unset()
}

// GetOpenStackAvailabilityZone returns the OpenStackAvailabilityZone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackCredentialsForProjectDto) GetOpenStackAvailabilityZone() string {
	if o == nil || IsNil(o.OpenStackAvailabilityZone.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackAvailabilityZone.Get()
}

// GetOpenStackAvailabilityZoneOk returns a tuple with the OpenStackAvailabilityZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackCredentialsForProjectDto) GetOpenStackAvailabilityZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackAvailabilityZone.Get(), o.OpenStackAvailabilityZone.IsSet()
}

// HasOpenStackAvailabilityZone returns a boolean if a field has been set.
func (o *OpenstackCredentialsForProjectDto) HasOpenStackAvailabilityZone() bool {
	if o != nil && o.OpenStackAvailabilityZone.IsSet() {
		return true
	}

	return false
}

// SetOpenStackAvailabilityZone gets a reference to the given NullableString and assigns it to the OpenStackAvailabilityZone field.
func (o *OpenstackCredentialsForProjectDto) SetOpenStackAvailabilityZone(v string) {
	o.OpenStackAvailabilityZone.Set(&v)
}
// SetOpenStackAvailabilityZoneNil sets the value for OpenStackAvailabilityZone to be an explicit nil
func (o *OpenstackCredentialsForProjectDto) SetOpenStackAvailabilityZoneNil() {
	o.OpenStackAvailabilityZone.Set(nil)
}

// UnsetOpenStackAvailabilityZone ensures that no value is present for OpenStackAvailabilityZone, not even an explicit nil
func (o *OpenstackCredentialsForProjectDto) UnsetOpenStackAvailabilityZone() {
	o.OpenStackAvailabilityZone.Unset()
}

// GetOpenStackVolumeType returns the OpenStackVolumeType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackCredentialsForProjectDto) GetOpenStackVolumeType() string {
	if o == nil || IsNil(o.OpenStackVolumeType.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackVolumeType.Get()
}

// GetOpenStackVolumeTypeOk returns a tuple with the OpenStackVolumeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackCredentialsForProjectDto) GetOpenStackVolumeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackVolumeType.Get(), o.OpenStackVolumeType.IsSet()
}

// HasOpenStackVolumeType returns a boolean if a field has been set.
func (o *OpenstackCredentialsForProjectDto) HasOpenStackVolumeType() bool {
	if o != nil && o.OpenStackVolumeType.IsSet() {
		return true
	}

	return false
}

// SetOpenStackVolumeType gets a reference to the given NullableString and assigns it to the OpenStackVolumeType field.
func (o *OpenstackCredentialsForProjectDto) SetOpenStackVolumeType(v string) {
	o.OpenStackVolumeType.Set(&v)
}
// SetOpenStackVolumeTypeNil sets the value for OpenStackVolumeType to be an explicit nil
func (o *OpenstackCredentialsForProjectDto) SetOpenStackVolumeTypeNil() {
	o.OpenStackVolumeType.Set(nil)
}

// UnsetOpenStackVolumeType ensures that no value is present for OpenStackVolumeType, not even an explicit nil
func (o *OpenstackCredentialsForProjectDto) UnsetOpenStackVolumeType() {
	o.OpenStackVolumeType.Unset()
}

// GetOpenStackTenantId returns the OpenStackTenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackCredentialsForProjectDto) GetOpenStackTenantId() string {
	if o == nil || IsNil(o.OpenStackTenantId.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackTenantId.Get()
}

// GetOpenStackTenantIdOk returns a tuple with the OpenStackTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackCredentialsForProjectDto) GetOpenStackTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackTenantId.Get(), o.OpenStackTenantId.IsSet()
}

// HasOpenStackTenantId returns a boolean if a field has been set.
func (o *OpenstackCredentialsForProjectDto) HasOpenStackTenantId() bool {
	if o != nil && o.OpenStackTenantId.IsSet() {
		return true
	}

	return false
}

// SetOpenStackTenantId gets a reference to the given NullableString and assigns it to the OpenStackTenantId field.
func (o *OpenstackCredentialsForProjectDto) SetOpenStackTenantId(v string) {
	o.OpenStackTenantId.Set(&v)
}
// SetOpenStackTenantIdNil sets the value for OpenStackTenantId to be an explicit nil
func (o *OpenstackCredentialsForProjectDto) SetOpenStackTenantIdNil() {
	o.OpenStackTenantId.Set(nil)
}

// UnsetOpenStackTenantId ensures that no value is present for OpenStackTenantId, not even an explicit nil
func (o *OpenstackCredentialsForProjectDto) UnsetOpenStackTenantId() {
	o.OpenStackTenantId.Unset()
}

// GetOpenStackImportNetwork returns the OpenStackImportNetwork field value if set, zero value otherwise.
func (o *OpenstackCredentialsForProjectDto) GetOpenStackImportNetwork() bool {
	if o == nil || IsNil(o.OpenStackImportNetwork) {
		var ret bool
		return ret
	}
	return *o.OpenStackImportNetwork
}

// GetOpenStackImportNetworkOk returns a tuple with the OpenStackImportNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackCredentialsForProjectDto) GetOpenStackImportNetworkOk() (*bool, bool) {
	if o == nil || IsNil(o.OpenStackImportNetwork) {
		return nil, false
	}
	return o.OpenStackImportNetwork, true
}

// HasOpenStackImportNetwork returns a boolean if a field has been set.
func (o *OpenstackCredentialsForProjectDto) HasOpenStackImportNetwork() bool {
	if o != nil && !IsNil(o.OpenStackImportNetwork) {
		return true
	}

	return false
}

// SetOpenStackImportNetwork gets a reference to the given bool and assigns it to the OpenStackImportNetwork field.
func (o *OpenstackCredentialsForProjectDto) SetOpenStackImportNetwork(v bool) {
	o.OpenStackImportNetwork = &v
}

// GetOpenStackInternalSubnetId returns the OpenStackInternalSubnetId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenstackCredentialsForProjectDto) GetOpenStackInternalSubnetId() string {
	if o == nil || IsNil(o.OpenStackInternalSubnetId.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackInternalSubnetId.Get()
}

// GetOpenStackInternalSubnetIdOk returns a tuple with the OpenStackInternalSubnetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenstackCredentialsForProjectDto) GetOpenStackInternalSubnetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackInternalSubnetId.Get(), o.OpenStackInternalSubnetId.IsSet()
}

// HasOpenStackInternalSubnetId returns a boolean if a field has been set.
func (o *OpenstackCredentialsForProjectDto) HasOpenStackInternalSubnetId() bool {
	if o != nil && o.OpenStackInternalSubnetId.IsSet() {
		return true
	}

	return false
}

// SetOpenStackInternalSubnetId gets a reference to the given NullableString and assigns it to the OpenStackInternalSubnetId field.
func (o *OpenstackCredentialsForProjectDto) SetOpenStackInternalSubnetId(v string) {
	o.OpenStackInternalSubnetId.Set(&v)
}
// SetOpenStackInternalSubnetIdNil sets the value for OpenStackInternalSubnetId to be an explicit nil
func (o *OpenstackCredentialsForProjectDto) SetOpenStackInternalSubnetIdNil() {
	o.OpenStackInternalSubnetId.Set(nil)
}

// UnsetOpenStackInternalSubnetId ensures that no value is present for OpenStackInternalSubnetId, not even an explicit nil
func (o *OpenstackCredentialsForProjectDto) UnsetOpenStackInternalSubnetId() {
	o.OpenStackInternalSubnetId.Unset()
}

// GetApplicationCredEnabled returns the ApplicationCredEnabled field value if set, zero value otherwise.
func (o *OpenstackCredentialsForProjectDto) GetApplicationCredEnabled() bool {
	if o == nil || IsNil(o.ApplicationCredEnabled) {
		var ret bool
		return ret
	}
	return *o.ApplicationCredEnabled
}

// GetApplicationCredEnabledOk returns a tuple with the ApplicationCredEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackCredentialsForProjectDto) GetApplicationCredEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ApplicationCredEnabled) {
		return nil, false
	}
	return o.ApplicationCredEnabled, true
}

// HasApplicationCredEnabled returns a boolean if a field has been set.
func (o *OpenstackCredentialsForProjectDto) HasApplicationCredEnabled() bool {
	if o != nil && !IsNil(o.ApplicationCredEnabled) {
		return true
	}

	return false
}

// SetApplicationCredEnabled gets a reference to the given bool and assigns it to the ApplicationCredEnabled field.
func (o *OpenstackCredentialsForProjectDto) SetApplicationCredEnabled(v bool) {
	o.ApplicationCredEnabled = &v
}

func (o OpenstackCredentialsForProjectDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenstackCredentialsForProjectDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.OpenStackUser.IsSet() {
		toSerialize["openStackUser"] = o.OpenStackUser.Get()
	}
	if o.OpenStackDomain.IsSet() {
		toSerialize["openStackDomain"] = o.OpenStackDomain.Get()
	}
	if o.OpenStackPassword.IsSet() {
		toSerialize["openStackPassword"] = o.OpenStackPassword.Get()
	}
	if o.OpenStackProject.IsSet() {
		toSerialize["openStackProject"] = o.OpenStackProject.Get()
	}
	if o.OpenStackRegion.IsSet() {
		toSerialize["openStackRegion"] = o.OpenStackRegion.Get()
	}
	if o.OpenStackUrl.IsSet() {
		toSerialize["openStackUrl"] = o.OpenStackUrl.Get()
	}
	if o.OpenStackPublicNetwork.IsSet() {
		toSerialize["openStackPublicNetwork"] = o.OpenStackPublicNetwork.Get()
	}
	if o.OpenStackAvailabilityZone.IsSet() {
		toSerialize["openStackAvailabilityZone"] = o.OpenStackAvailabilityZone.Get()
	}
	if o.OpenStackVolumeType.IsSet() {
		toSerialize["openStackVolumeType"] = o.OpenStackVolumeType.Get()
	}
	if o.OpenStackTenantId.IsSet() {
		toSerialize["openStackTenantId"] = o.OpenStackTenantId.Get()
	}
	if !IsNil(o.OpenStackImportNetwork) {
		toSerialize["openStackImportNetwork"] = o.OpenStackImportNetwork
	}
	if o.OpenStackInternalSubnetId.IsSet() {
		toSerialize["openStackInternalSubnetId"] = o.OpenStackInternalSubnetId.Get()
	}
	if !IsNil(o.ApplicationCredEnabled) {
		toSerialize["applicationCredEnabled"] = o.ApplicationCredEnabled
	}
	return toSerialize, nil
}

type NullableOpenstackCredentialsForProjectDto struct {
	value *OpenstackCredentialsForProjectDto
	isSet bool
}

func (v NullableOpenstackCredentialsForProjectDto) Get() *OpenstackCredentialsForProjectDto {
	return v.value
}

func (v *NullableOpenstackCredentialsForProjectDto) Set(val *OpenstackCredentialsForProjectDto) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenstackCredentialsForProjectDto) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenstackCredentialsForProjectDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenstackCredentialsForProjectDto(val *OpenstackCredentialsForProjectDto) *NullableOpenstackCredentialsForProjectDto {
	return &NullableOpenstackCredentialsForProjectDto{value: val, isSet: true}
}

func (v NullableOpenstackCredentialsForProjectDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenstackCredentialsForProjectDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


