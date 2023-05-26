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

// checks if the OpaProfileListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpaProfileListDto{}

// OpaProfileListDto struct for OpaProfileListDto
type OpaProfileListDto struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	ForbidNodePort *bool `json:"forbidNodePort,omitempty"`
	ForbidHttpIngress *bool `json:"forbidHttpIngress,omitempty"`
	RequireProbe *bool `json:"requireProbe,omitempty"`
	UniqueIngresses *bool `json:"uniqueIngresses,omitempty"`
	UniqueServiceSelector *bool `json:"uniqueServiceSelector,omitempty"`
	AllowedRepo []string `json:"allowedRepo,omitempty"`
	ForbidSpecificTags []string `json:"forbidSpecificTags,omitempty"`
	IngressWhitelist []string `json:"ingressWhitelist,omitempty"`
	IsLocked *bool `json:"isLocked,omitempty"`
	Revision *int32 `json:"revision,omitempty"`
	OrganizationId *int32 `json:"organizationId,omitempty"`
	OrganizationName NullableString `json:"organizationName,omitempty"`
	CreatedAt NullableString `json:"createdAt,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
	Projects []CommonDropdownDto `json:"projects,omitempty"`
}

// NewOpaProfileListDto instantiates a new OpaProfileListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpaProfileListDto() *OpaProfileListDto {
	this := OpaProfileListDto{}
	return &this
}

// NewOpaProfileListDtoWithDefaults instantiates a new OpaProfileListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpaProfileListDtoWithDefaults() *OpaProfileListDto {
	this := OpaProfileListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OpaProfileListDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpaProfileListDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OpaProfileListDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *OpaProfileListDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpaProfileListDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpaProfileListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *OpaProfileListDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *OpaProfileListDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *OpaProfileListDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *OpaProfileListDto) UnsetName() {
	o.Name.Unset()
}

// GetForbidNodePort returns the ForbidNodePort field value if set, zero value otherwise.
func (o *OpaProfileListDto) GetForbidNodePort() bool {
	if o == nil || IsNil(o.ForbidNodePort) {
		var ret bool
		return ret
	}
	return *o.ForbidNodePort
}

// GetForbidNodePortOk returns a tuple with the ForbidNodePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpaProfileListDto) GetForbidNodePortOk() (*bool, bool) {
	if o == nil || IsNil(o.ForbidNodePort) {
		return nil, false
	}
	return o.ForbidNodePort, true
}

// HasForbidNodePort returns a boolean if a field has been set.
func (o *OpaProfileListDto) HasForbidNodePort() bool {
	if o != nil && !IsNil(o.ForbidNodePort) {
		return true
	}

	return false
}

// SetForbidNodePort gets a reference to the given bool and assigns it to the ForbidNodePort field.
func (o *OpaProfileListDto) SetForbidNodePort(v bool) {
	o.ForbidNodePort = &v
}

// GetForbidHttpIngress returns the ForbidHttpIngress field value if set, zero value otherwise.
func (o *OpaProfileListDto) GetForbidHttpIngress() bool {
	if o == nil || IsNil(o.ForbidHttpIngress) {
		var ret bool
		return ret
	}
	return *o.ForbidHttpIngress
}

// GetForbidHttpIngressOk returns a tuple with the ForbidHttpIngress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpaProfileListDto) GetForbidHttpIngressOk() (*bool, bool) {
	if o == nil || IsNil(o.ForbidHttpIngress) {
		return nil, false
	}
	return o.ForbidHttpIngress, true
}

// HasForbidHttpIngress returns a boolean if a field has been set.
func (o *OpaProfileListDto) HasForbidHttpIngress() bool {
	if o != nil && !IsNil(o.ForbidHttpIngress) {
		return true
	}

	return false
}

// SetForbidHttpIngress gets a reference to the given bool and assigns it to the ForbidHttpIngress field.
func (o *OpaProfileListDto) SetForbidHttpIngress(v bool) {
	o.ForbidHttpIngress = &v
}

// GetRequireProbe returns the RequireProbe field value if set, zero value otherwise.
func (o *OpaProfileListDto) GetRequireProbe() bool {
	if o == nil || IsNil(o.RequireProbe) {
		var ret bool
		return ret
	}
	return *o.RequireProbe
}

// GetRequireProbeOk returns a tuple with the RequireProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpaProfileListDto) GetRequireProbeOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireProbe) {
		return nil, false
	}
	return o.RequireProbe, true
}

// HasRequireProbe returns a boolean if a field has been set.
func (o *OpaProfileListDto) HasRequireProbe() bool {
	if o != nil && !IsNil(o.RequireProbe) {
		return true
	}

	return false
}

// SetRequireProbe gets a reference to the given bool and assigns it to the RequireProbe field.
func (o *OpaProfileListDto) SetRequireProbe(v bool) {
	o.RequireProbe = &v
}

// GetUniqueIngresses returns the UniqueIngresses field value if set, zero value otherwise.
func (o *OpaProfileListDto) GetUniqueIngresses() bool {
	if o == nil || IsNil(o.UniqueIngresses) {
		var ret bool
		return ret
	}
	return *o.UniqueIngresses
}

// GetUniqueIngressesOk returns a tuple with the UniqueIngresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpaProfileListDto) GetUniqueIngressesOk() (*bool, bool) {
	if o == nil || IsNil(o.UniqueIngresses) {
		return nil, false
	}
	return o.UniqueIngresses, true
}

// HasUniqueIngresses returns a boolean if a field has been set.
func (o *OpaProfileListDto) HasUniqueIngresses() bool {
	if o != nil && !IsNil(o.UniqueIngresses) {
		return true
	}

	return false
}

// SetUniqueIngresses gets a reference to the given bool and assigns it to the UniqueIngresses field.
func (o *OpaProfileListDto) SetUniqueIngresses(v bool) {
	o.UniqueIngresses = &v
}

// GetUniqueServiceSelector returns the UniqueServiceSelector field value if set, zero value otherwise.
func (o *OpaProfileListDto) GetUniqueServiceSelector() bool {
	if o == nil || IsNil(o.UniqueServiceSelector) {
		var ret bool
		return ret
	}
	return *o.UniqueServiceSelector
}

// GetUniqueServiceSelectorOk returns a tuple with the UniqueServiceSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpaProfileListDto) GetUniqueServiceSelectorOk() (*bool, bool) {
	if o == nil || IsNil(o.UniqueServiceSelector) {
		return nil, false
	}
	return o.UniqueServiceSelector, true
}

// HasUniqueServiceSelector returns a boolean if a field has been set.
func (o *OpaProfileListDto) HasUniqueServiceSelector() bool {
	if o != nil && !IsNil(o.UniqueServiceSelector) {
		return true
	}

	return false
}

// SetUniqueServiceSelector gets a reference to the given bool and assigns it to the UniqueServiceSelector field.
func (o *OpaProfileListDto) SetUniqueServiceSelector(v bool) {
	o.UniqueServiceSelector = &v
}

// GetAllowedRepo returns the AllowedRepo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpaProfileListDto) GetAllowedRepo() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AllowedRepo
}

// GetAllowedRepoOk returns a tuple with the AllowedRepo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpaProfileListDto) GetAllowedRepoOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedRepo) {
		return nil, false
	}
	return o.AllowedRepo, true
}

// HasAllowedRepo returns a boolean if a field has been set.
func (o *OpaProfileListDto) HasAllowedRepo() bool {
	if o != nil && IsNil(o.AllowedRepo) {
		return true
	}

	return false
}

// SetAllowedRepo gets a reference to the given []string and assigns it to the AllowedRepo field.
func (o *OpaProfileListDto) SetAllowedRepo(v []string) {
	o.AllowedRepo = v
}

// GetForbidSpecificTags returns the ForbidSpecificTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpaProfileListDto) GetForbidSpecificTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ForbidSpecificTags
}

// GetForbidSpecificTagsOk returns a tuple with the ForbidSpecificTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpaProfileListDto) GetForbidSpecificTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.ForbidSpecificTags) {
		return nil, false
	}
	return o.ForbidSpecificTags, true
}

// HasForbidSpecificTags returns a boolean if a field has been set.
func (o *OpaProfileListDto) HasForbidSpecificTags() bool {
	if o != nil && IsNil(o.ForbidSpecificTags) {
		return true
	}

	return false
}

// SetForbidSpecificTags gets a reference to the given []string and assigns it to the ForbidSpecificTags field.
func (o *OpaProfileListDto) SetForbidSpecificTags(v []string) {
	o.ForbidSpecificTags = v
}

// GetIngressWhitelist returns the IngressWhitelist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpaProfileListDto) GetIngressWhitelist() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IngressWhitelist
}

// GetIngressWhitelistOk returns a tuple with the IngressWhitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpaProfileListDto) GetIngressWhitelistOk() ([]string, bool) {
	if o == nil || IsNil(o.IngressWhitelist) {
		return nil, false
	}
	return o.IngressWhitelist, true
}

// HasIngressWhitelist returns a boolean if a field has been set.
func (o *OpaProfileListDto) HasIngressWhitelist() bool {
	if o != nil && IsNil(o.IngressWhitelist) {
		return true
	}

	return false
}

// SetIngressWhitelist gets a reference to the given []string and assigns it to the IngressWhitelist field.
func (o *OpaProfileListDto) SetIngressWhitelist(v []string) {
	o.IngressWhitelist = v
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *OpaProfileListDto) GetIsLocked() bool {
	if o == nil || IsNil(o.IsLocked) {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpaProfileListDto) GetIsLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocked) {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *OpaProfileListDto) HasIsLocked() bool {
	if o != nil && !IsNil(o.IsLocked) {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *OpaProfileListDto) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *OpaProfileListDto) GetRevision() int32 {
	if o == nil || IsNil(o.Revision) {
		var ret int32
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpaProfileListDto) GetRevisionOk() (*int32, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *OpaProfileListDto) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given int32 and assigns it to the Revision field.
func (o *OpaProfileListDto) SetRevision(v int32) {
	o.Revision = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *OpaProfileListDto) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId) {
		var ret int32
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpaProfileListDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *OpaProfileListDto) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int32 and assigns it to the OrganizationId field.
func (o *OpaProfileListDto) SetOrganizationId(v int32) {
	o.OrganizationId = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpaProfileListDto) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationName.Get()
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpaProfileListDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationName.Get(), o.OrganizationName.IsSet()
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *OpaProfileListDto) HasOrganizationName() bool {
	if o != nil && o.OrganizationName.IsSet() {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given NullableString and assigns it to the OrganizationName field.
func (o *OpaProfileListDto) SetOrganizationName(v string) {
	o.OrganizationName.Set(&v)
}
// SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil
func (o *OpaProfileListDto) SetOrganizationNameNil() {
	o.OrganizationName.Set(nil)
}

// UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
func (o *OpaProfileListDto) UnsetOrganizationName() {
	o.OrganizationName.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpaProfileListDto) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpaProfileListDto) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OpaProfileListDto) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableString and assigns it to the CreatedAt field.
func (o *OpaProfileListDto) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *OpaProfileListDto) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *OpaProfileListDto) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *OpaProfileListDto) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpaProfileListDto) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *OpaProfileListDto) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *OpaProfileListDto) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpaProfileListDto) GetProjects() []CommonDropdownDto {
	if o == nil {
		var ret []CommonDropdownDto
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpaProfileListDto) GetProjectsOk() ([]CommonDropdownDto, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *OpaProfileListDto) HasProjects() bool {
	if o != nil && IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []CommonDropdownDto and assigns it to the Projects field.
func (o *OpaProfileListDto) SetProjects(v []CommonDropdownDto) {
	o.Projects = v
}

func (o OpaProfileListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpaProfileListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.ForbidNodePort) {
		toSerialize["forbidNodePort"] = o.ForbidNodePort
	}
	if !IsNil(o.ForbidHttpIngress) {
		toSerialize["forbidHttpIngress"] = o.ForbidHttpIngress
	}
	if !IsNil(o.RequireProbe) {
		toSerialize["requireProbe"] = o.RequireProbe
	}
	if !IsNil(o.UniqueIngresses) {
		toSerialize["uniqueIngresses"] = o.UniqueIngresses
	}
	if !IsNil(o.UniqueServiceSelector) {
		toSerialize["uniqueServiceSelector"] = o.UniqueServiceSelector
	}
	if o.AllowedRepo != nil {
		toSerialize["allowedRepo"] = o.AllowedRepo
	}
	if o.ForbidSpecificTags != nil {
		toSerialize["forbidSpecificTags"] = o.ForbidSpecificTags
	}
	if o.IngressWhitelist != nil {
		toSerialize["ingressWhitelist"] = o.IngressWhitelist
	}
	if !IsNil(o.IsLocked) {
		toSerialize["isLocked"] = o.IsLocked
	}
	if !IsNil(o.Revision) {
		toSerialize["revision"] = o.Revision
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if o.OrganizationName.IsSet() {
		toSerialize["organizationName"] = o.OrganizationName.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["createdAt"] = o.CreatedAt.Get()
	}
	if !IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	return toSerialize, nil
}

type NullableOpaProfileListDto struct {
	value *OpaProfileListDto
	isSet bool
}

func (v NullableOpaProfileListDto) Get() *OpaProfileListDto {
	return v.value
}

func (v *NullableOpaProfileListDto) Set(val *OpaProfileListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableOpaProfileListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableOpaProfileListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpaProfileListDto(val *OpaProfileListDto) *NullableOpaProfileListDto {
	return &NullableOpaProfileListDto{value: val, isSet: true}
}

func (v NullableOpaProfileListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpaProfileListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


