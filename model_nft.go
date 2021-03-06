/*
 * NFT Storage API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nftstorage

import (
	"encoding/json"
	"time"
)

// NFT struct for NFT
type NFT struct {
	// Self-describing content-addressed identifiers for distributed systems. Check [spec](https://github.com/multiformats/cid) for more info.
	Cid *string `json:"cid,omitempty"`
	// Size in bytes of the NFT data.
	Size *float32 `json:"size,omitempty"`
	// This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: YYYY-MM-DDTHH:MM:SSZ.
	Created *time.Time `json:"created,omitempty"`
	// MIME type of the upload file or 'directory' when uploading multiple files.
	Type *string `json:"type,omitempty"`
	// Name of the JWT token used to create this NFT.
	Scope *string `json:"scope,omitempty"`
	Pin *Pin `json:"pin,omitempty"`
	// Files in the directory (only if this NFT is a directory).
	Files *[]map[string]interface{} `json:"files,omitempty"`
	Deals *[]Deal `json:"deals,omitempty"`
}

// NewNFT instantiates a new NFT object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNFT() *NFT {
	this := NFT{}
	var size float32 = 132614
	this.Size = &size
	var scope string = "default"
	this.Scope = &scope
	return &this
}

// NewNFTWithDefaults instantiates a new NFT object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNFTWithDefaults() *NFT {
	this := NFT{}
	var size float32 = 132614
	this.Size = &size
	var scope string = "default"
	this.Scope = &scope
	return &this
}

// GetCid returns the Cid field value if set, zero value otherwise.
func (o *NFT) GetCid() string {
	if o == nil || o.Cid == nil {
		var ret string
		return ret
	}
	return *o.Cid
}

// GetCidOk returns a tuple with the Cid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFT) GetCidOk() (*string, bool) {
	if o == nil || o.Cid == nil {
		return nil, false
	}
	return o.Cid, true
}

// HasCid returns a boolean if a field has been set.
func (o *NFT) HasCid() bool {
	if o != nil && o.Cid != nil {
		return true
	}

	return false
}

// SetCid gets a reference to the given string and assigns it to the Cid field.
func (o *NFT) SetCid(v string) {
	o.Cid = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *NFT) GetSize() float32 {
	if o == nil || o.Size == nil {
		var ret float32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFT) GetSizeOk() (*float32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *NFT) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given float32 and assigns it to the Size field.
func (o *NFT) SetSize(v float32) {
	o.Size = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *NFT) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFT) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *NFT) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *NFT) SetCreated(v time.Time) {
	o.Created = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NFT) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFT) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NFT) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NFT) SetType(v string) {
	o.Type = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *NFT) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFT) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *NFT) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *NFT) SetScope(v string) {
	o.Scope = &v
}

// GetPin returns the Pin field value if set, zero value otherwise.
func (o *NFT) GetPin() Pin {
	if o == nil || o.Pin == nil {
		var ret Pin
		return ret
	}
	return *o.Pin
}

// GetPinOk returns a tuple with the Pin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFT) GetPinOk() (*Pin, bool) {
	if o == nil || o.Pin == nil {
		return nil, false
	}
	return o.Pin, true
}

// HasPin returns a boolean if a field has been set.
func (o *NFT) HasPin() bool {
	if o != nil && o.Pin != nil {
		return true
	}

	return false
}

// SetPin gets a reference to the given Pin and assigns it to the Pin field.
func (o *NFT) SetPin(v Pin) {
	o.Pin = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *NFT) GetFiles() []map[string]interface{} {
	if o == nil || o.Files == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFT) GetFilesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Files == nil {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *NFT) HasFiles() bool {
	if o != nil && o.Files != nil {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []map[string]interface{} and assigns it to the Files field.
func (o *NFT) SetFiles(v []map[string]interface{}) {
	o.Files = &v
}

// GetDeals returns the Deals field value if set, zero value otherwise.
func (o *NFT) GetDeals() []Deal {
	if o == nil || o.Deals == nil {
		var ret []Deal
		return ret
	}
	return *o.Deals
}

// GetDealsOk returns a tuple with the Deals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFT) GetDealsOk() (*[]Deal, bool) {
	if o == nil || o.Deals == nil {
		return nil, false
	}
	return o.Deals, true
}

// HasDeals returns a boolean if a field has been set.
func (o *NFT) HasDeals() bool {
	if o != nil && o.Deals != nil {
		return true
	}

	return false
}

// SetDeals gets a reference to the given []Deal and assigns it to the Deals field.
func (o *NFT) SetDeals(v []Deal) {
	o.Deals = &v
}

func (o NFT) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cid != nil {
		toSerialize["cid"] = o.Cid
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.Pin != nil {
		toSerialize["pin"] = o.Pin
	}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	if o.Deals != nil {
		toSerialize["deals"] = o.Deals
	}
	return json.Marshal(toSerialize)
}

type NullableNFT struct {
	value *NFT
	isSet bool
}

func (v NullableNFT) Get() *NFT {
	return v.value
}

func (v *NullableNFT) Set(val *NFT) {
	v.value = val
	v.isSet = true
}

func (v NullableNFT) IsSet() bool {
	return v.isSet
}

func (v *NullableNFT) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFT(val *NFT) *NullableNFT {
	return &NullableNFT{value: val, isSet: true}
}

func (v NullableNFT) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFT) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


