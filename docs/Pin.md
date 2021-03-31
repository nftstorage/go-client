# Pin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | Pointer to **string** | Self-describing content-addressed identifiers for distributed systems. Check [spec](https://github.com/multiformats/cid) for more info. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**PinStatus**](PinStatus.md) |  | [optional] 
**Created** | Pointer to **time.Time** | This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: YYYY-MM-DDTHH:MM:SSZ. | [optional] 
**Size** | Pointer to **float32** |  | [optional] 

## Methods

### NewPin

`func NewPin() *Pin`

NewPin instantiates a new Pin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPinWithDefaults

`func NewPinWithDefaults() *Pin`

NewPinWithDefaults instantiates a new Pin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *Pin) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *Pin) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *Pin) SetCid(v string)`

SetCid sets Cid field to given value.

### HasCid

`func (o *Pin) HasCid() bool`

HasCid returns a boolean if a field has been set.

### GetName

`func (o *Pin) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Pin) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Pin) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Pin) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *Pin) GetStatus() PinStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Pin) GetStatusOk() (*PinStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Pin) SetStatus(v PinStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Pin) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreated

`func (o *Pin) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Pin) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Pin) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Pin) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetSize

`func (o *Pin) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Pin) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Pin) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Pin) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


