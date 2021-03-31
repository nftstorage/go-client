# Deal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchRootCid** | Pointer to **string** |  | [optional] 
**LastChange** | **string** | This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: YYYY-MM-DDTHH:MM:SSZ. | 
**Miner** | Pointer to **string** | Miner ID | [optional] 
**Network** | Pointer to **string** | Filecoin network for this Deal | [optional] 
**PieceCid** | Pointer to **string** | Piece CID string | [optional] 
**Status** | **string** | Deal Status | 
**StatusText** | Pointer to **string** | Deal Status Description | [optional] 
**DealActivation** | Pointer to **string** | This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: YYYY-MM-DDTHH:MM:SSZ. | [optional] 
**DealExpiration** | Pointer to **string** | This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: YYYY-MM-DDTHH:MM:SSZ. | [optional] 

## Methods

### NewDeal

`func NewDeal(lastChange string, status string, ) *Deal`

NewDeal instantiates a new Deal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealWithDefaults

`func NewDealWithDefaults() *Deal`

NewDealWithDefaults instantiates a new Deal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchRootCid

`func (o *Deal) GetBatchRootCid() string`

GetBatchRootCid returns the BatchRootCid field if non-nil, zero value otherwise.

### GetBatchRootCidOk

`func (o *Deal) GetBatchRootCidOk() (*string, bool)`

GetBatchRootCidOk returns a tuple with the BatchRootCid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchRootCid

`func (o *Deal) SetBatchRootCid(v string)`

SetBatchRootCid sets BatchRootCid field to given value.

### HasBatchRootCid

`func (o *Deal) HasBatchRootCid() bool`

HasBatchRootCid returns a boolean if a field has been set.

### GetLastChange

`func (o *Deal) GetLastChange() string`

GetLastChange returns the LastChange field if non-nil, zero value otherwise.

### GetLastChangeOk

`func (o *Deal) GetLastChangeOk() (*string, bool)`

GetLastChangeOk returns a tuple with the LastChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChange

`func (o *Deal) SetLastChange(v string)`

SetLastChange sets LastChange field to given value.


### GetMiner

`func (o *Deal) GetMiner() string`

GetMiner returns the Miner field if non-nil, zero value otherwise.

### GetMinerOk

`func (o *Deal) GetMinerOk() (*string, bool)`

GetMinerOk returns a tuple with the Miner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiner

`func (o *Deal) SetMiner(v string)`

SetMiner sets Miner field to given value.

### HasMiner

`func (o *Deal) HasMiner() bool`

HasMiner returns a boolean if a field has been set.

### GetNetwork

`func (o *Deal) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Deal) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Deal) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Deal) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetPieceCid

`func (o *Deal) GetPieceCid() string`

GetPieceCid returns the PieceCid field if non-nil, zero value otherwise.

### GetPieceCidOk

`func (o *Deal) GetPieceCidOk() (*string, bool)`

GetPieceCidOk returns a tuple with the PieceCid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPieceCid

`func (o *Deal) SetPieceCid(v string)`

SetPieceCid sets PieceCid field to given value.

### HasPieceCid

`func (o *Deal) HasPieceCid() bool`

HasPieceCid returns a boolean if a field has been set.

### GetStatus

`func (o *Deal) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Deal) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Deal) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusText

`func (o *Deal) GetStatusText() string`

GetStatusText returns the StatusText field if non-nil, zero value otherwise.

### GetStatusTextOk

`func (o *Deal) GetStatusTextOk() (*string, bool)`

GetStatusTextOk returns a tuple with the StatusText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusText

`func (o *Deal) SetStatusText(v string)`

SetStatusText sets StatusText field to given value.

### HasStatusText

`func (o *Deal) HasStatusText() bool`

HasStatusText returns a boolean if a field has been set.

### GetDealActivation

`func (o *Deal) GetDealActivation() string`

GetDealActivation returns the DealActivation field if non-nil, zero value otherwise.

### GetDealActivationOk

`func (o *Deal) GetDealActivationOk() (*string, bool)`

GetDealActivationOk returns a tuple with the DealActivation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealActivation

`func (o *Deal) SetDealActivation(v string)`

SetDealActivation sets DealActivation field to given value.

### HasDealActivation

`func (o *Deal) HasDealActivation() bool`

HasDealActivation returns a boolean if a field has been set.

### GetDealExpiration

`func (o *Deal) GetDealExpiration() string`

GetDealExpiration returns the DealExpiration field if non-nil, zero value otherwise.

### GetDealExpirationOk

`func (o *Deal) GetDealExpirationOk() (*string, bool)`

GetDealExpirationOk returns a tuple with the DealExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealExpiration

`func (o *Deal) SetDealExpiration(v string)`

SetDealExpiration sets DealExpiration field to given value.

### HasDealExpiration

`func (o *Deal) HasDealExpiration() bool`

HasDealExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


