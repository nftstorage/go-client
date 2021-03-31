# UnauthorizedErrorResponseError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] [default to "HTTP Error"]
**Message** | Pointer to **string** |  | [optional] [default to "Unauthorized"]

## Methods

### NewUnauthorizedErrorResponseError

`func NewUnauthorizedErrorResponseError() *UnauthorizedErrorResponseError`

NewUnauthorizedErrorResponseError instantiates a new UnauthorizedErrorResponseError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnauthorizedErrorResponseErrorWithDefaults

`func NewUnauthorizedErrorResponseErrorWithDefaults() *UnauthorizedErrorResponseError`

NewUnauthorizedErrorResponseErrorWithDefaults instantiates a new UnauthorizedErrorResponseError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UnauthorizedErrorResponseError) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnauthorizedErrorResponseError) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnauthorizedErrorResponseError) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnauthorizedErrorResponseError) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMessage

`func (o *UnauthorizedErrorResponseError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UnauthorizedErrorResponseError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UnauthorizedErrorResponseError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UnauthorizedErrorResponseError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


