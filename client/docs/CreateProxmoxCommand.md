# CreateProxmoxCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Username** | **string** |  | 
**Url** | **string** |  | 
**Password** | **string** |  | 
**Storage** | **string** |  | 
**VmTemplateName** | **string** |  | 
**Continent** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **NullableInt32** |  | [optional] 
**Hypervisors** | Pointer to **[]string** |  | [optional] 
**PublicNetwork** | Pointer to [**CreateProxmoxNetworkDto**](CreateProxmoxNetworkDto.md) |  | [optional] 
**PrivateNetwork** | Pointer to [**CreateProxmoxNetworkDto**](CreateProxmoxNetworkDto.md) |  | [optional] 

## Methods

### NewCreateProxmoxCommand

`func NewCreateProxmoxCommand(name string, username string, url string, password string, storage string, vmTemplateName string, ) *CreateProxmoxCommand`

NewCreateProxmoxCommand instantiates a new CreateProxmoxCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProxmoxCommandWithDefaults

`func NewCreateProxmoxCommandWithDefaults() *CreateProxmoxCommand`

NewCreateProxmoxCommandWithDefaults instantiates a new CreateProxmoxCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateProxmoxCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateProxmoxCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateProxmoxCommand) SetName(v string)`

SetName sets Name field to given value.


### GetUsername

`func (o *CreateProxmoxCommand) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateProxmoxCommand) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateProxmoxCommand) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetUrl

`func (o *CreateProxmoxCommand) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateProxmoxCommand) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateProxmoxCommand) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetPassword

`func (o *CreateProxmoxCommand) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateProxmoxCommand) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateProxmoxCommand) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetStorage

`func (o *CreateProxmoxCommand) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *CreateProxmoxCommand) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *CreateProxmoxCommand) SetStorage(v string)`

SetStorage sets Storage field to given value.


### GetVmTemplateName

`func (o *CreateProxmoxCommand) GetVmTemplateName() string`

GetVmTemplateName returns the VmTemplateName field if non-nil, zero value otherwise.

### GetVmTemplateNameOk

`func (o *CreateProxmoxCommand) GetVmTemplateNameOk() (*string, bool)`

GetVmTemplateNameOk returns a tuple with the VmTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTemplateName

`func (o *CreateProxmoxCommand) SetVmTemplateName(v string)`

SetVmTemplateName sets VmTemplateName field to given value.


### GetContinent

`func (o *CreateProxmoxCommand) GetContinent() string`

GetContinent returns the Continent field if non-nil, zero value otherwise.

### GetContinentOk

`func (o *CreateProxmoxCommand) GetContinentOk() (*string, bool)`

GetContinentOk returns a tuple with the Continent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinent

`func (o *CreateProxmoxCommand) SetContinent(v string)`

SetContinent sets Continent field to given value.

### HasContinent

`func (o *CreateProxmoxCommand) HasContinent() bool`

HasContinent returns a boolean if a field has been set.

### SetContinentNil

`func (o *CreateProxmoxCommand) SetContinentNil(b bool)`

 SetContinentNil sets the value for Continent to be an explicit nil

### UnsetContinent
`func (o *CreateProxmoxCommand) UnsetContinent()`

UnsetContinent ensures that no value is present for Continent, not even an explicit nil
### GetOrganizationId

`func (o *CreateProxmoxCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateProxmoxCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateProxmoxCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CreateProxmoxCommand) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *CreateProxmoxCommand) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *CreateProxmoxCommand) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetHypervisors

`func (o *CreateProxmoxCommand) GetHypervisors() []string`

GetHypervisors returns the Hypervisors field if non-nil, zero value otherwise.

### GetHypervisorsOk

`func (o *CreateProxmoxCommand) GetHypervisorsOk() (*[]string, bool)`

GetHypervisorsOk returns a tuple with the Hypervisors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisors

`func (o *CreateProxmoxCommand) SetHypervisors(v []string)`

SetHypervisors sets Hypervisors field to given value.

### HasHypervisors

`func (o *CreateProxmoxCommand) HasHypervisors() bool`

HasHypervisors returns a boolean if a field has been set.

### SetHypervisorsNil

`func (o *CreateProxmoxCommand) SetHypervisorsNil(b bool)`

 SetHypervisorsNil sets the value for Hypervisors to be an explicit nil

### UnsetHypervisors
`func (o *CreateProxmoxCommand) UnsetHypervisors()`

UnsetHypervisors ensures that no value is present for Hypervisors, not even an explicit nil
### GetPublicNetwork

`func (o *CreateProxmoxCommand) GetPublicNetwork() CreateProxmoxNetworkDto`

GetPublicNetwork returns the PublicNetwork field if non-nil, zero value otherwise.

### GetPublicNetworkOk

`func (o *CreateProxmoxCommand) GetPublicNetworkOk() (*CreateProxmoxNetworkDto, bool)`

GetPublicNetworkOk returns a tuple with the PublicNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetwork

`func (o *CreateProxmoxCommand) SetPublicNetwork(v CreateProxmoxNetworkDto)`

SetPublicNetwork sets PublicNetwork field to given value.

### HasPublicNetwork

`func (o *CreateProxmoxCommand) HasPublicNetwork() bool`

HasPublicNetwork returns a boolean if a field has been set.

### GetPrivateNetwork

`func (o *CreateProxmoxCommand) GetPrivateNetwork() CreateProxmoxNetworkDto`

GetPrivateNetwork returns the PrivateNetwork field if non-nil, zero value otherwise.

### GetPrivateNetworkOk

`func (o *CreateProxmoxCommand) GetPrivateNetworkOk() (*CreateProxmoxNetworkDto, bool)`

GetPrivateNetworkOk returns a tuple with the PrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetwork

`func (o *CreateProxmoxCommand) SetPrivateNetwork(v CreateProxmoxNetworkDto)`

SetPrivateNetwork sets PrivateNetwork field to given value.

### HasPrivateNetwork

`func (o *CreateProxmoxCommand) HasPrivateNetwork() bool`

HasPrivateNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


