# NFT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | Pointer to **string** | Self-describing content-addressed identifiers for distributed systems. Check [spec](https://github.com/multiformats/cid) for more info. | [optional] 
**Size** | Pointer to **float32** | Size in bytes of the NFT data. | [optional] [default to 132614]
**Created** | Pointer to **time.Time** | This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: YYYY-MM-DDTHH:MM:SSZ. | [optional] 
**Type** | Pointer to **string** | MIME type of the upload file or &#39;directory&#39; when uploading multiple files. | [optional] 
**Scope** | Pointer to **string** | Name of the JWT token used to create this NFT. | [optional] [default to "default"]
**Pin** | Pointer to [**Pin**](Pin.md) |  | [optional] 
**Files** | Pointer to **[]map[string]interface{}** | Files in the directory (only if this NFT is a directory). | [optional] 
**Deals** | Pointer to [**[]Deal**](Deal.md) |  | [optional] 

## Methods

### NewNFT

`func NewNFT() *NFT`

NewNFT instantiates a new NFT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNFTWithDefaults

`func NewNFTWithDefaults() *NFT`

NewNFTWithDefaults instantiates a new NFT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *NFT) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *NFT) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *NFT) SetCid(v string)`

SetCid sets Cid field to given value.

### HasCid

`func (o *NFT) HasCid() bool`

HasCid returns a boolean if a field has been set.

### GetSize

`func (o *NFT) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *NFT) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *NFT) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *NFT) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetCreated

`func (o *NFT) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *NFT) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *NFT) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *NFT) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetType

`func (o *NFT) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NFT) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NFT) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NFT) HasType() bool`

HasType returns a boolean if a field has been set.

### GetScope

`func (o *NFT) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *NFT) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *NFT) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *NFT) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetPin

`func (o *NFT) GetPin() Pin`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *NFT) GetPinOk() (*Pin, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *NFT) SetPin(v Pin)`

SetPin sets Pin field to given value.

### HasPin

`func (o *NFT) HasPin() bool`

HasPin returns a boolean if a field has been set.

### GetFiles

`func (o *NFT) GetFiles() []map[string]interface{}`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *NFT) GetFilesOk() (*[]map[string]interface{}, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *NFT) SetFiles(v []map[string]interface{})`

SetFiles sets Files field to given value.

### HasFiles

`func (o *NFT) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetDeals

`func (o *NFT) GetDeals() []Deal`

GetDeals returns the Deals field if non-nil, zero value otherwise.

### GetDealsOk

`func (o *NFT) GetDealsOk() (*[]Deal, bool)`

GetDealsOk returns a tuple with the Deals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeals

`func (o *NFT) SetDeals(v []Deal)`

SetDeals sets Deals field to given value.

### HasDeals

`func (o *NFT) HasDeals() bool`

HasDeals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


