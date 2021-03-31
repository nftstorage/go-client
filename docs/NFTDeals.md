# NFTDeals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Overall deal status | [optional] 
**Deals** | Pointer to [**[]Deal**](Deal.md) |  | [optional] 

## Methods

### NewNFTDeals

`func NewNFTDeals() *NFTDeals`

NewNFTDeals instantiates a new NFTDeals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNFTDealsWithDefaults

`func NewNFTDealsWithDefaults() *NFTDeals`

NewNFTDealsWithDefaults instantiates a new NFTDeals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *NFTDeals) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NFTDeals) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NFTDeals) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NFTDeals) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDeals

`func (o *NFTDeals) GetDeals() []Deal`

GetDeals returns the Deals field if non-nil, zero value otherwise.

### GetDealsOk

`func (o *NFTDeals) GetDealsOk() (*[]Deal, bool)`

GetDealsOk returns a tuple with the Deals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeals

`func (o *NFTDeals) SetDeals(v []Deal)`

SetDeals sets Deals field to given value.

### HasDeals

`func (o *NFTDeals) HasDeals() bool`

HasDeals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


