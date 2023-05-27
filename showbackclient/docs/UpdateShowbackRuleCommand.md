# UpdateShowbackRuleCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ByLabel** | Pointer to **string** |  | [optional] 
**MetricName** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to [**EShowbackType**](EShowbackType.md) |  | [optional] 
**Type** | Pointer to [**EPrometheusType**](EPrometheusType.md) |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**ProjectAlertLimit** | Pointer to **int32** |  | [optional] 
**GlobalAlertLimit** | Pointer to **int32** |  | [optional] 
**Labels** | Pointer to [**[]ShowbackLabelCreateDto**](ShowbackLabelCreateDto.md) |  | [optional] 

## Methods

### NewUpdateShowbackRuleCommand

`func NewUpdateShowbackRuleCommand() *UpdateShowbackRuleCommand`

NewUpdateShowbackRuleCommand instantiates a new UpdateShowbackRuleCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateShowbackRuleCommandWithDefaults

`func NewUpdateShowbackRuleCommandWithDefaults() *UpdateShowbackRuleCommand`

NewUpdateShowbackRuleCommandWithDefaults instantiates a new UpdateShowbackRuleCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateShowbackRuleCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateShowbackRuleCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateShowbackRuleCommand) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateShowbackRuleCommand) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateShowbackRuleCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateShowbackRuleCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateShowbackRuleCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateShowbackRuleCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### GetByLabel

`func (o *UpdateShowbackRuleCommand) GetByLabel() string`

GetByLabel returns the ByLabel field if non-nil, zero value otherwise.

### GetByLabelOk

`func (o *UpdateShowbackRuleCommand) GetByLabelOk() (*string, bool)`

GetByLabelOk returns a tuple with the ByLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByLabel

`func (o *UpdateShowbackRuleCommand) SetByLabel(v string)`

SetByLabel sets ByLabel field to given value.

### HasByLabel

`func (o *UpdateShowbackRuleCommand) HasByLabel() bool`

HasByLabel returns a boolean if a field has been set.

### GetMetricName

`func (o *UpdateShowbackRuleCommand) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *UpdateShowbackRuleCommand) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *UpdateShowbackRuleCommand) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *UpdateShowbackRuleCommand) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetKind

`func (o *UpdateShowbackRuleCommand) GetKind() EShowbackType`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *UpdateShowbackRuleCommand) GetKindOk() (*EShowbackType, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *UpdateShowbackRuleCommand) SetKind(v EShowbackType)`

SetKind sets Kind field to given value.

### HasKind

`func (o *UpdateShowbackRuleCommand) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetType

`func (o *UpdateShowbackRuleCommand) GetType() EPrometheusType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateShowbackRuleCommand) GetTypeOk() (*EPrometheusType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateShowbackRuleCommand) SetType(v EPrometheusType)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateShowbackRuleCommand) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPrice

`func (o *UpdateShowbackRuleCommand) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *UpdateShowbackRuleCommand) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *UpdateShowbackRuleCommand) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *UpdateShowbackRuleCommand) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetProjectAlertLimit

`func (o *UpdateShowbackRuleCommand) GetProjectAlertLimit() int32`

GetProjectAlertLimit returns the ProjectAlertLimit field if non-nil, zero value otherwise.

### GetProjectAlertLimitOk

`func (o *UpdateShowbackRuleCommand) GetProjectAlertLimitOk() (*int32, bool)`

GetProjectAlertLimitOk returns a tuple with the ProjectAlertLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectAlertLimit

`func (o *UpdateShowbackRuleCommand) SetProjectAlertLimit(v int32)`

SetProjectAlertLimit sets ProjectAlertLimit field to given value.

### HasProjectAlertLimit

`func (o *UpdateShowbackRuleCommand) HasProjectAlertLimit() bool`

HasProjectAlertLimit returns a boolean if a field has been set.

### GetGlobalAlertLimit

`func (o *UpdateShowbackRuleCommand) GetGlobalAlertLimit() int32`

GetGlobalAlertLimit returns the GlobalAlertLimit field if non-nil, zero value otherwise.

### GetGlobalAlertLimitOk

`func (o *UpdateShowbackRuleCommand) GetGlobalAlertLimitOk() (*int32, bool)`

GetGlobalAlertLimitOk returns a tuple with the GlobalAlertLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalAlertLimit

`func (o *UpdateShowbackRuleCommand) SetGlobalAlertLimit(v int32)`

SetGlobalAlertLimit sets GlobalAlertLimit field to given value.

### HasGlobalAlertLimit

`func (o *UpdateShowbackRuleCommand) HasGlobalAlertLimit() bool`

HasGlobalAlertLimit returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateShowbackRuleCommand) GetLabels() []ShowbackLabelCreateDto`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateShowbackRuleCommand) GetLabelsOk() (*[]ShowbackLabelCreateDto, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateShowbackRuleCommand) SetLabels(v []ShowbackLabelCreateDto)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateShowbackRuleCommand) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


