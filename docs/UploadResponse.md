# UploadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | Pointer to **bool** |  | [optional] [default to true]
**Value** | Pointer to [**NFT**](NFT.md) |  | [optional] 

## Methods

### NewUploadResponse

`func NewUploadResponse() *UploadResponse`

NewUploadResponse instantiates a new UploadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadResponseWithDefaults

`func NewUploadResponseWithDefaults() *UploadResponse`

NewUploadResponseWithDefaults instantiates a new UploadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *UploadResponse) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *UploadResponse) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *UploadResponse) SetOk(v bool)`

SetOk sets Ok field to given value.

### HasOk

`func (o *UploadResponse) HasOk() bool`

HasOk returns a boolean if a field has been set.

### GetValue

`func (o *UploadResponse) GetValue() NFT`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UploadResponse) GetValueOk() (*NFT, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UploadResponse) SetValue(v NFT)`

SetValue sets Value field to given value.

### HasValue

`func (o *UploadResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


