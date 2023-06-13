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

// checks if the CreateAlertingProfileCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAlertingProfileCommand{}

// CreateAlertingProfileCommand struct for CreateAlertingProfileCommand
type CreateAlertingProfileCommand struct {
	Name NullableString `json:"name,omitempty"`
	SlackConfigurationId NullableInt32 `json:"slackConfigurationId,omitempty"`
	OrganizationId NullableInt32 `json:"organizationId,omitempty"`
	Emails []AlertingEmailDto `json:"emails,omitempty"`
	Webhooks []AlertingWebhookDto `json:"webhooks,omitempty"`
	AlertingIntegrations []AlertingIntegrationDto `json:"alertingIntegrations,omitempty"`
	Reminder *AlertingReminder `json:"reminder,omitempty"`
}

// NewCreateAlertingProfileCommand instantiates a new CreateAlertingProfileCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAlertingProfileCommand() *CreateAlertingProfileCommand {
	this := CreateAlertingProfileCommand{}
	return &this
}

// NewCreateAlertingProfileCommandWithDefaults instantiates a new CreateAlertingProfileCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAlertingProfileCommandWithDefaults() *CreateAlertingProfileCommand {
	this := CreateAlertingProfileCommand{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAlertingProfileCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAlertingProfileCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CreateAlertingProfileCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CreateAlertingProfileCommand) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CreateAlertingProfileCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CreateAlertingProfileCommand) UnsetName() {
	o.Name.Unset()
}

// GetSlackConfigurationId returns the SlackConfigurationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAlertingProfileCommand) GetSlackConfigurationId() int32 {
	if o == nil || IsNil(o.SlackConfigurationId.Get()) {
		var ret int32
		return ret
	}
	return *o.SlackConfigurationId.Get()
}

// GetSlackConfigurationIdOk returns a tuple with the SlackConfigurationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAlertingProfileCommand) GetSlackConfigurationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SlackConfigurationId.Get(), o.SlackConfigurationId.IsSet()
}

// HasSlackConfigurationId returns a boolean if a field has been set.
func (o *CreateAlertingProfileCommand) HasSlackConfigurationId() bool {
	if o != nil && o.SlackConfigurationId.IsSet() {
		return true
	}

	return false
}

// SetSlackConfigurationId gets a reference to the given NullableInt32 and assigns it to the SlackConfigurationId field.
func (o *CreateAlertingProfileCommand) SetSlackConfigurationId(v int32) {
	o.SlackConfigurationId.Set(&v)
}
// SetSlackConfigurationIdNil sets the value for SlackConfigurationId to be an explicit nil
func (o *CreateAlertingProfileCommand) SetSlackConfigurationIdNil() {
	o.SlackConfigurationId.Set(nil)
}

// UnsetSlackConfigurationId ensures that no value is present for SlackConfigurationId, not even an explicit nil
func (o *CreateAlertingProfileCommand) UnsetSlackConfigurationId() {
	o.SlackConfigurationId.Unset()
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAlertingProfileCommand) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret int32
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAlertingProfileCommand) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *CreateAlertingProfileCommand) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableInt32 and assigns it to the OrganizationId field.
func (o *CreateAlertingProfileCommand) SetOrganizationId(v int32) {
	o.OrganizationId.Set(&v)
}
// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *CreateAlertingProfileCommand) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *CreateAlertingProfileCommand) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

// GetEmails returns the Emails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAlertingProfileCommand) GetEmails() []AlertingEmailDto {
	if o == nil {
		var ret []AlertingEmailDto
		return ret
	}
	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAlertingProfileCommand) GetEmailsOk() ([]AlertingEmailDto, bool) {
	if o == nil || IsNil(o.Emails) {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *CreateAlertingProfileCommand) HasEmails() bool {
	if o != nil && IsNil(o.Emails) {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []AlertingEmailDto and assigns it to the Emails field.
func (o *CreateAlertingProfileCommand) SetEmails(v []AlertingEmailDto) {
	o.Emails = v
}

// GetWebhooks returns the Webhooks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAlertingProfileCommand) GetWebhooks() []AlertingWebhookDto {
	if o == nil {
		var ret []AlertingWebhookDto
		return ret
	}
	return o.Webhooks
}

// GetWebhooksOk returns a tuple with the Webhooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAlertingProfileCommand) GetWebhooksOk() ([]AlertingWebhookDto, bool) {
	if o == nil || IsNil(o.Webhooks) {
		return nil, false
	}
	return o.Webhooks, true
}

// HasWebhooks returns a boolean if a field has been set.
func (o *CreateAlertingProfileCommand) HasWebhooks() bool {
	if o != nil && IsNil(o.Webhooks) {
		return true
	}

	return false
}

// SetWebhooks gets a reference to the given []AlertingWebhookDto and assigns it to the Webhooks field.
func (o *CreateAlertingProfileCommand) SetWebhooks(v []AlertingWebhookDto) {
	o.Webhooks = v
}

// GetAlertingIntegrations returns the AlertingIntegrations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAlertingProfileCommand) GetAlertingIntegrations() []AlertingIntegrationDto {
	if o == nil {
		var ret []AlertingIntegrationDto
		return ret
	}
	return o.AlertingIntegrations
}

// GetAlertingIntegrationsOk returns a tuple with the AlertingIntegrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAlertingProfileCommand) GetAlertingIntegrationsOk() ([]AlertingIntegrationDto, bool) {
	if o == nil || IsNil(o.AlertingIntegrations) {
		return nil, false
	}
	return o.AlertingIntegrations, true
}

// HasAlertingIntegrations returns a boolean if a field has been set.
func (o *CreateAlertingProfileCommand) HasAlertingIntegrations() bool {
	if o != nil && IsNil(o.AlertingIntegrations) {
		return true
	}

	return false
}

// SetAlertingIntegrations gets a reference to the given []AlertingIntegrationDto and assigns it to the AlertingIntegrations field.
func (o *CreateAlertingProfileCommand) SetAlertingIntegrations(v []AlertingIntegrationDto) {
	o.AlertingIntegrations = v
}

// GetReminder returns the Reminder field value if set, zero value otherwise.
func (o *CreateAlertingProfileCommand) GetReminder() AlertingReminder {
	if o == nil || IsNil(o.Reminder) {
		var ret AlertingReminder
		return ret
	}
	return *o.Reminder
}

// GetReminderOk returns a tuple with the Reminder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAlertingProfileCommand) GetReminderOk() (*AlertingReminder, bool) {
	if o == nil || IsNil(o.Reminder) {
		return nil, false
	}
	return o.Reminder, true
}

// HasReminder returns a boolean if a field has been set.
func (o *CreateAlertingProfileCommand) HasReminder() bool {
	if o != nil && !IsNil(o.Reminder) {
		return true
	}

	return false
}

// SetReminder gets a reference to the given AlertingReminder and assigns it to the Reminder field.
func (o *CreateAlertingProfileCommand) SetReminder(v AlertingReminder) {
	o.Reminder = &v
}

func (o CreateAlertingProfileCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAlertingProfileCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.SlackConfigurationId.IsSet() {
		toSerialize["slackConfigurationId"] = o.SlackConfigurationId.Get()
	}
	if o.OrganizationId.IsSet() {
		toSerialize["organizationId"] = o.OrganizationId.Get()
	}
	if o.Emails != nil {
		toSerialize["emails"] = o.Emails
	}
	if o.Webhooks != nil {
		toSerialize["webhooks"] = o.Webhooks
	}
	if o.AlertingIntegrations != nil {
		toSerialize["alertingIntegrations"] = o.AlertingIntegrations
	}
	if !IsNil(o.Reminder) {
		toSerialize["reminder"] = o.Reminder
	}
	return toSerialize, nil
}

type NullableCreateAlertingProfileCommand struct {
	value *CreateAlertingProfileCommand
	isSet bool
}

func (v NullableCreateAlertingProfileCommand) Get() *CreateAlertingProfileCommand {
	return v.value
}

func (v *NullableCreateAlertingProfileCommand) Set(val *CreateAlertingProfileCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAlertingProfileCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAlertingProfileCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAlertingProfileCommand(val *CreateAlertingProfileCommand) *NullableCreateAlertingProfileCommand {
	return &NullableCreateAlertingProfileCommand{value: val, isSet: true}
}

func (v NullableCreateAlertingProfileCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAlertingProfileCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


