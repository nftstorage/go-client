# Links

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipfs** | Pointer to **string** |  | [optional] 
**Http** | Pointer to **string** |  | [optional] 
**File** | Pointer to [**[]LinksFile**](LinksFile.md) |  | [optional] 

## Methods

### NewLinks

`func NewLinks() *Links`

NewLinks instantiates a new Links object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinksWithDefaults

`func NewLinksWithDefaults() *Links`

NewLinksWithDefaults instantiates a new Links object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpfs

`func (o *Links) GetIpfs() string`

GetIpfs returns the Ipfs field if non-nil, zero value otherwise.

### GetIpfsOk

`func (o *Links) GetIpfsOk() (*string, bool)`

GetIpfsOk returns a tuple with the Ipfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpfs

`func (o *Links) SetIpfs(v string)`

SetIpfs sets Ipfs field to given value.

### HasIpfs

`func (o *Links) HasIpfs() bool`

HasIpfs returns a boolean if a field has been set.

### GetHttp

`func (o *Links) GetHttp() string`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *Links) GetHttpOk() (*string, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *Links) SetHttp(v string)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *Links) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetFile

`func (o *Links) GetFile() []LinksFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *Links) GetFileOk() (*[]LinksFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *Links) SetFile(v []LinksFile)`

SetFile sets File field to given value.

### HasFile

`func (o *Links) HasFile() bool`

HasFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


