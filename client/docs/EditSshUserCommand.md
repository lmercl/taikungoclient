# EditSshUserCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**SshPublicKey** | **string** |  | 
**AccessProfileId** | **int32** |  | 

## Methods

### NewEditSshUserCommand

`func NewEditSshUserCommand(id int32, name string, sshPublicKey string, accessProfileId int32, ) *EditSshUserCommand`

NewEditSshUserCommand instantiates a new EditSshUserCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditSshUserCommandWithDefaults

`func NewEditSshUserCommandWithDefaults() *EditSshUserCommand`

NewEditSshUserCommandWithDefaults instantiates a new EditSshUserCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EditSshUserCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EditSshUserCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EditSshUserCommand) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *EditSshUserCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditSshUserCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditSshUserCommand) SetName(v string)`

SetName sets Name field to given value.


### GetSshPublicKey

`func (o *EditSshUserCommand) GetSshPublicKey() string`

GetSshPublicKey returns the SshPublicKey field if non-nil, zero value otherwise.

### GetSshPublicKeyOk

`func (o *EditSshUserCommand) GetSshPublicKeyOk() (*string, bool)`

GetSshPublicKeyOk returns a tuple with the SshPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPublicKey

`func (o *EditSshUserCommand) SetSshPublicKey(v string)`

SetSshPublicKey sets SshPublicKey field to given value.


### GetAccessProfileId

`func (o *EditSshUserCommand) GetAccessProfileId() int32`

GetAccessProfileId returns the AccessProfileId field if non-nil, zero value otherwise.

### GetAccessProfileIdOk

`func (o *EditSshUserCommand) GetAccessProfileIdOk() (*int32, bool)`

GetAccessProfileIdOk returns a tuple with the AccessProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfileId

`func (o *EditSshUserCommand) SetAccessProfileId(v int32)`

SetAccessProfileId sets AccessProfileId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


