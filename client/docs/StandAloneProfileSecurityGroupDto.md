# StandAloneProfileSecurityGroupDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Protocol** | Pointer to [**SecurityGroupProtocol**](SecurityGroupProtocol.md) |  | [optional] 
**PortMinRange** | Pointer to **int32** |  | [optional] 
**PortMaxRange** | Pointer to **int32** |  | [optional] 
**RemoteIpPrefix** | **string** |  | 

## Methods

### NewStandAloneProfileSecurityGroupDto

`func NewStandAloneProfileSecurityGroupDto(name string, remoteIpPrefix string, ) *StandAloneProfileSecurityGroupDto`

NewStandAloneProfileSecurityGroupDto instantiates a new StandAloneProfileSecurityGroupDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandAloneProfileSecurityGroupDtoWithDefaults

`func NewStandAloneProfileSecurityGroupDtoWithDefaults() *StandAloneProfileSecurityGroupDto`

NewStandAloneProfileSecurityGroupDtoWithDefaults instantiates a new StandAloneProfileSecurityGroupDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StandAloneProfileSecurityGroupDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StandAloneProfileSecurityGroupDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StandAloneProfileSecurityGroupDto) SetName(v string)`

SetName sets Name field to given value.


### GetProtocol

`func (o *StandAloneProfileSecurityGroupDto) GetProtocol() SecurityGroupProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *StandAloneProfileSecurityGroupDto) GetProtocolOk() (*SecurityGroupProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *StandAloneProfileSecurityGroupDto) SetProtocol(v SecurityGroupProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *StandAloneProfileSecurityGroupDto) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPortMinRange

`func (o *StandAloneProfileSecurityGroupDto) GetPortMinRange() int32`

GetPortMinRange returns the PortMinRange field if non-nil, zero value otherwise.

### GetPortMinRangeOk

`func (o *StandAloneProfileSecurityGroupDto) GetPortMinRangeOk() (*int32, bool)`

GetPortMinRangeOk returns a tuple with the PortMinRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMinRange

`func (o *StandAloneProfileSecurityGroupDto) SetPortMinRange(v int32)`

SetPortMinRange sets PortMinRange field to given value.

### HasPortMinRange

`func (o *StandAloneProfileSecurityGroupDto) HasPortMinRange() bool`

HasPortMinRange returns a boolean if a field has been set.

### GetPortMaxRange

`func (o *StandAloneProfileSecurityGroupDto) GetPortMaxRange() int32`

GetPortMaxRange returns the PortMaxRange field if non-nil, zero value otherwise.

### GetPortMaxRangeOk

`func (o *StandAloneProfileSecurityGroupDto) GetPortMaxRangeOk() (*int32, bool)`

GetPortMaxRangeOk returns a tuple with the PortMaxRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMaxRange

`func (o *StandAloneProfileSecurityGroupDto) SetPortMaxRange(v int32)`

SetPortMaxRange sets PortMaxRange field to given value.

### HasPortMaxRange

`func (o *StandAloneProfileSecurityGroupDto) HasPortMaxRange() bool`

HasPortMaxRange returns a boolean if a field has been set.

### GetRemoteIpPrefix

`func (o *StandAloneProfileSecurityGroupDto) GetRemoteIpPrefix() string`

GetRemoteIpPrefix returns the RemoteIpPrefix field if non-nil, zero value otherwise.

### GetRemoteIpPrefixOk

`func (o *StandAloneProfileSecurityGroupDto) GetRemoteIpPrefixOk() (*string, bool)`

GetRemoteIpPrefixOk returns a tuple with the RemoteIpPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpPrefix

`func (o *StandAloneProfileSecurityGroupDto) SetRemoteIpPrefix(v string)`

SetRemoteIpPrefix sets RemoteIpPrefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


