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

// checks if the ProjectsForBillingDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectsForBillingDto{}

// ProjectsForBillingDto struct for ProjectsForBillingDto
type ProjectsForBillingDto struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	CreatedAt NullableTime `json:"createdAt,omitempty"`
	BillingStartDate NullableTime `json:"billingStartDate,omitempty"`
	OrganizationName NullableString `json:"organizationName,omitempty"`
	Price *float64 `json:"price,omitempty"`
	Servers []ServersForBillingDto `json:"servers,omitempty"`
	StandaloneVms []StandaloneVmsForBillingDto `json:"standaloneVms,omitempty"`
	BillingEnabled *bool `json:"billingEnabled,omitempty"`
}

// NewProjectsForBillingDto instantiates a new ProjectsForBillingDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsForBillingDto() *ProjectsForBillingDto {
	this := ProjectsForBillingDto{}
	return &this
}

// NewProjectsForBillingDtoWithDefaults instantiates a new ProjectsForBillingDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsForBillingDtoWithDefaults() *ProjectsForBillingDto {
	this := ProjectsForBillingDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProjectsForBillingDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsForBillingDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProjectsForBillingDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ProjectsForBillingDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectsForBillingDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectsForBillingDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ProjectsForBillingDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ProjectsForBillingDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ProjectsForBillingDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ProjectsForBillingDto) UnsetName() {
	o.Name.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectsForBillingDto) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectsForBillingDto) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ProjectsForBillingDto) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *ProjectsForBillingDto) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *ProjectsForBillingDto) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *ProjectsForBillingDto) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetBillingStartDate returns the BillingStartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectsForBillingDto) GetBillingStartDate() time.Time {
	if o == nil || IsNil(o.BillingStartDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.BillingStartDate.Get()
}

// GetBillingStartDateOk returns a tuple with the BillingStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectsForBillingDto) GetBillingStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingStartDate.Get(), o.BillingStartDate.IsSet()
}

// HasBillingStartDate returns a boolean if a field has been set.
func (o *ProjectsForBillingDto) HasBillingStartDate() bool {
	if o != nil && o.BillingStartDate.IsSet() {
		return true
	}

	return false
}

// SetBillingStartDate gets a reference to the given NullableTime and assigns it to the BillingStartDate field.
func (o *ProjectsForBillingDto) SetBillingStartDate(v time.Time) {
	o.BillingStartDate.Set(&v)
}
// SetBillingStartDateNil sets the value for BillingStartDate to be an explicit nil
func (o *ProjectsForBillingDto) SetBillingStartDateNil() {
	o.BillingStartDate.Set(nil)
}

// UnsetBillingStartDate ensures that no value is present for BillingStartDate, not even an explicit nil
func (o *ProjectsForBillingDto) UnsetBillingStartDate() {
	o.BillingStartDate.Unset()
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectsForBillingDto) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationName.Get()
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectsForBillingDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationName.Get(), o.OrganizationName.IsSet()
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *ProjectsForBillingDto) HasOrganizationName() bool {
	if o != nil && o.OrganizationName.IsSet() {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given NullableString and assigns it to the OrganizationName field.
func (o *ProjectsForBillingDto) SetOrganizationName(v string) {
	o.OrganizationName.Set(&v)
}
// SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil
func (o *ProjectsForBillingDto) SetOrganizationNameNil() {
	o.OrganizationName.Set(nil)
}

// UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
func (o *ProjectsForBillingDto) UnsetOrganizationName() {
	o.OrganizationName.Unset()
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ProjectsForBillingDto) GetPrice() float64 {
	if o == nil || IsNil(o.Price) {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsForBillingDto) GetPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ProjectsForBillingDto) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *ProjectsForBillingDto) SetPrice(v float64) {
	o.Price = &v
}

// GetServers returns the Servers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectsForBillingDto) GetServers() []ServersForBillingDto {
	if o == nil {
		var ret []ServersForBillingDto
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectsForBillingDto) GetServersOk() ([]ServersForBillingDto, bool) {
	if o == nil || IsNil(o.Servers) {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *ProjectsForBillingDto) HasServers() bool {
	if o != nil && IsNil(o.Servers) {
		return true
	}

	return false
}

// SetServers gets a reference to the given []ServersForBillingDto and assigns it to the Servers field.
func (o *ProjectsForBillingDto) SetServers(v []ServersForBillingDto) {
	o.Servers = v
}

// GetStandaloneVms returns the StandaloneVms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectsForBillingDto) GetStandaloneVms() []StandaloneVmsForBillingDto {
	if o == nil {
		var ret []StandaloneVmsForBillingDto
		return ret
	}
	return o.StandaloneVms
}

// GetStandaloneVmsOk returns a tuple with the StandaloneVms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectsForBillingDto) GetStandaloneVmsOk() ([]StandaloneVmsForBillingDto, bool) {
	if o == nil || IsNil(o.StandaloneVms) {
		return nil, false
	}
	return o.StandaloneVms, true
}

// HasStandaloneVms returns a boolean if a field has been set.
func (o *ProjectsForBillingDto) HasStandaloneVms() bool {
	if o != nil && IsNil(o.StandaloneVms) {
		return true
	}

	return false
}

// SetStandaloneVms gets a reference to the given []StandaloneVmsForBillingDto and assigns it to the StandaloneVms field.
func (o *ProjectsForBillingDto) SetStandaloneVms(v []StandaloneVmsForBillingDto) {
	o.StandaloneVms = v
}

// GetBillingEnabled returns the BillingEnabled field value if set, zero value otherwise.
func (o *ProjectsForBillingDto) GetBillingEnabled() bool {
	if o == nil || IsNil(o.BillingEnabled) {
		var ret bool
		return ret
	}
	return *o.BillingEnabled
}

// GetBillingEnabledOk returns a tuple with the BillingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsForBillingDto) GetBillingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BillingEnabled) {
		return nil, false
	}
	return o.BillingEnabled, true
}

// HasBillingEnabled returns a boolean if a field has been set.
func (o *ProjectsForBillingDto) HasBillingEnabled() bool {
	if o != nil && !IsNil(o.BillingEnabled) {
		return true
	}

	return false
}

// SetBillingEnabled gets a reference to the given bool and assigns it to the BillingEnabled field.
func (o *ProjectsForBillingDto) SetBillingEnabled(v bool) {
	o.BillingEnabled = &v
}

func (o ProjectsForBillingDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectsForBillingDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["createdAt"] = o.CreatedAt.Get()
	}
	if o.BillingStartDate.IsSet() {
		toSerialize["billingStartDate"] = o.BillingStartDate.Get()
	}
	if o.OrganizationName.IsSet() {
		toSerialize["organizationName"] = o.OrganizationName.Get()
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if o.Servers != nil {
		toSerialize["servers"] = o.Servers
	}
	if o.StandaloneVms != nil {
		toSerialize["standaloneVms"] = o.StandaloneVms
	}
	if !IsNil(o.BillingEnabled) {
		toSerialize["billingEnabled"] = o.BillingEnabled
	}
	return toSerialize, nil
}

type NullableProjectsForBillingDto struct {
	value *ProjectsForBillingDto
	isSet bool
}

func (v NullableProjectsForBillingDto) Get() *ProjectsForBillingDto {
	return v.value
}

func (v *NullableProjectsForBillingDto) Set(val *ProjectsForBillingDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsForBillingDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsForBillingDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsForBillingDto(val *ProjectsForBillingDto) *NullableProjectsForBillingDto {
	return &NullableProjectsForBillingDto{value: val, isSet: true}
}

func (v NullableProjectsForBillingDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsForBillingDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


