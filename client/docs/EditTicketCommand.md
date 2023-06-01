# EditTicketCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TicketId** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 

## Methods

### NewEditTicketCommand

`func NewEditTicketCommand(ticketId string, name string, description string, ) *EditTicketCommand`

NewEditTicketCommand instantiates a new EditTicketCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditTicketCommandWithDefaults

`func NewEditTicketCommandWithDefaults() *EditTicketCommand`

NewEditTicketCommandWithDefaults instantiates a new EditTicketCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTicketId

`func (o *EditTicketCommand) GetTicketId() string`

GetTicketId returns the TicketId field if non-nil, zero value otherwise.

### GetTicketIdOk

`func (o *EditTicketCommand) GetTicketIdOk() (*string, bool)`

GetTicketIdOk returns a tuple with the TicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketId

`func (o *EditTicketCommand) SetTicketId(v string)`

SetTicketId sets TicketId field to given value.


### GetName

`func (o *EditTicketCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditTicketCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditTicketCommand) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EditTicketCommand) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditTicketCommand) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditTicketCommand) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


