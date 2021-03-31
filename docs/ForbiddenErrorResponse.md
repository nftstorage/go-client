# ForbiddenErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | Pointer to **bool** |  | [optional] [default to false]
**Error** | Pointer to [**ForbiddenErrorResponseError**](ForbiddenErrorResponseError.md) |  | [optional] 

## Methods

### NewForbiddenErrorResponse

`func NewForbiddenErrorResponse() *ForbiddenErrorResponse`

NewForbiddenErrorResponse instantiates a new ForbiddenErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForbiddenErrorResponseWithDefaults

`func NewForbiddenErrorResponseWithDefaults() *ForbiddenErrorResponse`

NewForbiddenErrorResponseWithDefaults instantiates a new ForbiddenErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *ForbiddenErrorResponse) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *ForbiddenErrorResponse) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *ForbiddenErrorResponse) SetOk(v bool)`

SetOk sets Ok field to given value.

### HasOk

`func (o *ForbiddenErrorResponse) HasOk() bool`

HasOk returns a boolean if a field has been set.

### GetError

`func (o *ForbiddenErrorResponse) GetError() ForbiddenErrorResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ForbiddenErrorResponse) GetErrorOk() (*ForbiddenErrorResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ForbiddenErrorResponse) SetError(v ForbiddenErrorResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *ForbiddenErrorResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


