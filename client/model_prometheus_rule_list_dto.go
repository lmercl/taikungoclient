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

// checks if the PrometheusRuleListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrometheusRuleListDto{}

// PrometheusRuleListDto struct for PrometheusRuleListDto
type PrometheusRuleListDto struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Password NullableString `json:"password,omitempty"`
	UserName NullableString `json:"userName,omitempty"`
	Url NullableString `json:"url,omitempty"`
	MetricName NullableString `json:"metricName,omitempty"`
	Labels []PrometheusLabelUpdateDto `json:"labels,omitempty"`
	Type NullableString `json:"type,omitempty"`
	Price *float64 `json:"price,omitempty"`
	IsAll *bool `json:"isAll,omitempty"`
	BillingStartDate NullableString `json:"billingStartDate,omitempty"`
	CreatedAt NullableString `json:"createdAt,omitempty"`
	Partner *PartnerDetailsDto `json:"partner,omitempty"`
	BoundOrganizations []OrganizationForPrometheus `json:"boundOrganizations,omitempty"`
	OperationCredential *OperationCredentialsForOrganizationEntity `json:"operationCredential,omitempty"`
	CreatedBy NullableString `json:"createdBy,omitempty"`
	LastModified NullableString `json:"lastModified,omitempty"`
	LastModifiedBy NullableString `json:"lastModifiedBy,omitempty"`
}

// NewPrometheusRuleListDto instantiates a new PrometheusRuleListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrometheusRuleListDto() *PrometheusRuleListDto {
	this := PrometheusRuleListDto{}
	return &this
}

// NewPrometheusRuleListDtoWithDefaults instantiates a new PrometheusRuleListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrometheusRuleListDtoWithDefaults() *PrometheusRuleListDto {
	this := PrometheusRuleListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PrometheusRuleListDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusRuleListDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PrometheusRuleListDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusRuleListDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusRuleListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PrometheusRuleListDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *PrometheusRuleListDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PrometheusRuleListDto) UnsetName() {
	o.Name.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusRuleListDto) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusRuleListDto) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *PrometheusRuleListDto) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *PrometheusRuleListDto) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *PrometheusRuleListDto) UnsetPassword() {
	o.Password.Unset()
}

// GetUserName returns the UserName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusRuleListDto) GetUserName() string {
	if o == nil || IsNil(o.UserName.Get()) {
		var ret string
		return ret
	}
	return *o.UserName.Get()
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusRuleListDto) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserName.Get(), o.UserName.IsSet()
}

// HasUserName returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasUserName() bool {
	if o != nil && o.UserName.IsSet() {
		return true
	}

	return false
}

// SetUserName gets a reference to the given NullableString and assigns it to the UserName field.
func (o *PrometheusRuleListDto) SetUserName(v string) {
	o.UserName.Set(&v)
}
// SetUserNameNil sets the value for UserName to be an explicit nil
func (o *PrometheusRuleListDto) SetUserNameNil() {
	o.UserName.Set(nil)
}

// UnsetUserName ensures that no value is present for UserName, not even an explicit nil
func (o *PrometheusRuleListDto) UnsetUserName() {
	o.UserName.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusRuleListDto) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusRuleListDto) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *PrometheusRuleListDto) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *PrometheusRuleListDto) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *PrometheusRuleListDto) UnsetUrl() {
	o.Url.Unset()
}

// GetMetricName returns the MetricName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusRuleListDto) GetMetricName() string {
	if o == nil || IsNil(o.MetricName.Get()) {
		var ret string
		return ret
	}
	return *o.MetricName.Get()
}

// GetMetricNameOk returns a tuple with the MetricName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusRuleListDto) GetMetricNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetricName.Get(), o.MetricName.IsSet()
}

// HasMetricName returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasMetricName() bool {
	if o != nil && o.MetricName.IsSet() {
		return true
	}

	return false
}

// SetMetricName gets a reference to the given NullableString and assigns it to the MetricName field.
func (o *PrometheusRuleListDto) SetMetricName(v string) {
	o.MetricName.Set(&v)
}
// SetMetricNameNil sets the value for MetricName to be an explicit nil
func (o *PrometheusRuleListDto) SetMetricNameNil() {
	o.MetricName.Set(nil)
}

// UnsetMetricName ensures that no value is present for MetricName, not even an explicit nil
func (o *PrometheusRuleListDto) UnsetMetricName() {
	o.MetricName.Unset()
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusRuleListDto) GetLabels() []PrometheusLabelUpdateDto {
	if o == nil {
		var ret []PrometheusLabelUpdateDto
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusRuleListDto) GetLabelsOk() ([]PrometheusLabelUpdateDto, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasLabels() bool {
	if o != nil && IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []PrometheusLabelUpdateDto and assigns it to the Labels field.
func (o *PrometheusRuleListDto) SetLabels(v []PrometheusLabelUpdateDto) {
	o.Labels = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusRuleListDto) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusRuleListDto) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *PrometheusRuleListDto) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *PrometheusRuleListDto) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *PrometheusRuleListDto) UnsetType() {
	o.Type.Unset()
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *PrometheusRuleListDto) GetPrice() float64 {
	if o == nil || IsNil(o.Price) {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusRuleListDto) GetPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *PrometheusRuleListDto) SetPrice(v float64) {
	o.Price = &v
}

// GetIsAll returns the IsAll field value if set, zero value otherwise.
func (o *PrometheusRuleListDto) GetIsAll() bool {
	if o == nil || IsNil(o.IsAll) {
		var ret bool
		return ret
	}
	return *o.IsAll
}

// GetIsAllOk returns a tuple with the IsAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusRuleListDto) GetIsAllOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAll) {
		return nil, false
	}
	return o.IsAll, true
}

// HasIsAll returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasIsAll() bool {
	if o != nil && !IsNil(o.IsAll) {
		return true
	}

	return false
}

// SetIsAll gets a reference to the given bool and assigns it to the IsAll field.
func (o *PrometheusRuleListDto) SetIsAll(v bool) {
	o.IsAll = &v
}

// GetBillingStartDate returns the BillingStartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusRuleListDto) GetBillingStartDate() string {
	if o == nil || IsNil(o.BillingStartDate.Get()) {
		var ret string
		return ret
	}
	return *o.BillingStartDate.Get()
}

// GetBillingStartDateOk returns a tuple with the BillingStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusRuleListDto) GetBillingStartDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingStartDate.Get(), o.BillingStartDate.IsSet()
}

// HasBillingStartDate returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasBillingStartDate() bool {
	if o != nil && o.BillingStartDate.IsSet() {
		return true
	}

	return false
}

// SetBillingStartDate gets a reference to the given NullableString and assigns it to the BillingStartDate field.
func (o *PrometheusRuleListDto) SetBillingStartDate(v string) {
	o.BillingStartDate.Set(&v)
}
// SetBillingStartDateNil sets the value for BillingStartDate to be an explicit nil
func (o *PrometheusRuleListDto) SetBillingStartDateNil() {
	o.BillingStartDate.Set(nil)
}

// UnsetBillingStartDate ensures that no value is present for BillingStartDate, not even an explicit nil
func (o *PrometheusRuleListDto) UnsetBillingStartDate() {
	o.BillingStartDate.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusRuleListDto) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusRuleListDto) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableString and assigns it to the CreatedAt field.
func (o *PrometheusRuleListDto) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *PrometheusRuleListDto) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *PrometheusRuleListDto) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetPartner returns the Partner field value if set, zero value otherwise.
func (o *PrometheusRuleListDto) GetPartner() PartnerDetailsDto {
	if o == nil || IsNil(o.Partner) {
		var ret PartnerDetailsDto
		return ret
	}
	return *o.Partner
}

// GetPartnerOk returns a tuple with the Partner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusRuleListDto) GetPartnerOk() (*PartnerDetailsDto, bool) {
	if o == nil || IsNil(o.Partner) {
		return nil, false
	}
	return o.Partner, true
}

// HasPartner returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasPartner() bool {
	if o != nil && !IsNil(o.Partner) {
		return true
	}

	return false
}

// SetPartner gets a reference to the given PartnerDetailsDto and assigns it to the Partner field.
func (o *PrometheusRuleListDto) SetPartner(v PartnerDetailsDto) {
	o.Partner = &v
}

// GetBoundOrganizations returns the BoundOrganizations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusRuleListDto) GetBoundOrganizations() []OrganizationForPrometheus {
	if o == nil {
		var ret []OrganizationForPrometheus
		return ret
	}
	return o.BoundOrganizations
}

// GetBoundOrganizationsOk returns a tuple with the BoundOrganizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusRuleListDto) GetBoundOrganizationsOk() ([]OrganizationForPrometheus, bool) {
	if o == nil || IsNil(o.BoundOrganizations) {
		return nil, false
	}
	return o.BoundOrganizations, true
}

// HasBoundOrganizations returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasBoundOrganizations() bool {
	if o != nil && IsNil(o.BoundOrganizations) {
		return true
	}

	return false
}

// SetBoundOrganizations gets a reference to the given []OrganizationForPrometheus and assigns it to the BoundOrganizations field.
func (o *PrometheusRuleListDto) SetBoundOrganizations(v []OrganizationForPrometheus) {
	o.BoundOrganizations = v
}

// GetOperationCredential returns the OperationCredential field value if set, zero value otherwise.
func (o *PrometheusRuleListDto) GetOperationCredential() OperationCredentialsForOrganizationEntity {
	if o == nil || IsNil(o.OperationCredential) {
		var ret OperationCredentialsForOrganizationEntity
		return ret
	}
	return *o.OperationCredential
}

// GetOperationCredentialOk returns a tuple with the OperationCredential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusRuleListDto) GetOperationCredentialOk() (*OperationCredentialsForOrganizationEntity, bool) {
	if o == nil || IsNil(o.OperationCredential) {
		return nil, false
	}
	return o.OperationCredential, true
}

// HasOperationCredential returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasOperationCredential() bool {
	if o != nil && !IsNil(o.OperationCredential) {
		return true
	}

	return false
}

// SetOperationCredential gets a reference to the given OperationCredentialsForOrganizationEntity and assigns it to the OperationCredential field.
func (o *PrometheusRuleListDto) SetOperationCredential(v OperationCredentialsForOrganizationEntity) {
	o.OperationCredential = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusRuleListDto) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusRuleListDto) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *PrometheusRuleListDto) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}
// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *PrometheusRuleListDto) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *PrometheusRuleListDto) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetLastModified returns the LastModified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusRuleListDto) GetLastModified() string {
	if o == nil || IsNil(o.LastModified.Get()) {
		var ret string
		return ret
	}
	return *o.LastModified.Get()
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusRuleListDto) GetLastModifiedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModified.Get(), o.LastModified.IsSet()
}

// HasLastModified returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasLastModified() bool {
	if o != nil && o.LastModified.IsSet() {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given NullableString and assigns it to the LastModified field.
func (o *PrometheusRuleListDto) SetLastModified(v string) {
	o.LastModified.Set(&v)
}
// SetLastModifiedNil sets the value for LastModified to be an explicit nil
func (o *PrometheusRuleListDto) SetLastModifiedNil() {
	o.LastModified.Set(nil)
}

// UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
func (o *PrometheusRuleListDto) UnsetLastModified() {
	o.LastModified.Unset()
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusRuleListDto) GetLastModifiedBy() string {
	if o == nil || IsNil(o.LastModifiedBy.Get()) {
		var ret string
		return ret
	}
	return *o.LastModifiedBy.Get()
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusRuleListDto) GetLastModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModifiedBy.Get(), o.LastModifiedBy.IsSet()
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *PrometheusRuleListDto) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given NullableString and assigns it to the LastModifiedBy field.
func (o *PrometheusRuleListDto) SetLastModifiedBy(v string) {
	o.LastModifiedBy.Set(&v)
}
// SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil
func (o *PrometheusRuleListDto) SetLastModifiedByNil() {
	o.LastModifiedBy.Set(nil)
}

// UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
func (o *PrometheusRuleListDto) UnsetLastModifiedBy() {
	o.LastModifiedBy.Unset()
}

func (o PrometheusRuleListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrometheusRuleListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.UserName.IsSet() {
		toSerialize["userName"] = o.UserName.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.MetricName.IsSet() {
		toSerialize["metricName"] = o.MetricName.Get()
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.IsAll) {
		toSerialize["isAll"] = o.IsAll
	}
	if o.BillingStartDate.IsSet() {
		toSerialize["billingStartDate"] = o.BillingStartDate.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["createdAt"] = o.CreatedAt.Get()
	}
	if !IsNil(o.Partner) {
		toSerialize["partner"] = o.Partner
	}
	if o.BoundOrganizations != nil {
		toSerialize["boundOrganizations"] = o.BoundOrganizations
	}
	if !IsNil(o.OperationCredential) {
		toSerialize["operationCredential"] = o.OperationCredential
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

type NullablePrometheusRuleListDto struct {
	value *PrometheusRuleListDto
	isSet bool
}

func (v NullablePrometheusRuleListDto) Get() *PrometheusRuleListDto {
	return v.value
}

func (v *NullablePrometheusRuleListDto) Set(val *PrometheusRuleListDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePrometheusRuleListDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePrometheusRuleListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrometheusRuleListDto(val *PrometheusRuleListDto) *NullablePrometheusRuleListDto {
	return &NullablePrometheusRuleListDto{value: val, isSet: true}
}

func (v NullablePrometheusRuleListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrometheusRuleListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

