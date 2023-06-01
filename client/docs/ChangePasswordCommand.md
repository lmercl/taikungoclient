# ChangePasswordCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **interface{}** |  | 
**NewPassword** | **interface{}** |  | 

## Methods

### NewChangePasswordCommand

`func NewChangePasswordCommand(password interface{}, newPassword interface{}, ) *ChangePasswordCommand`

NewChangePasswordCommand instantiates a new ChangePasswordCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangePasswordCommandWithDefaults

`func NewChangePasswordCommandWithDefaults() *ChangePasswordCommand`

NewChangePasswordCommandWithDefaults instantiates a new ChangePasswordCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *ChangePasswordCommand) GetPassword() interface{}`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ChangePasswordCommand) GetPasswordOk() (*interface{}, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ChangePasswordCommand) SetPassword(v interface{})`

SetPassword sets Password field to given value.


### GetNewPassword

`func (o *ChangePasswordCommand) GetNewPassword() interface{}`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *ChangePasswordCommand) GetNewPasswordOk() (*interface{}, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *ChangePasswordCommand) SetNewPassword(v interface{})`

SetNewPassword sets NewPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


