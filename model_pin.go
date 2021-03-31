/*
 * NFT Storage API
 *
 * # API Reference 
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nftstorage

import (
	"encoding/json"
	"time"
)

// Pin struct for Pin
type Pin struct {
	// Self-describing content-addressed identifiers for distributed systems. Check [spec](https://github.com/multiformats/cid) for more info.
	Cid *string `json:"cid,omitempty"`
	Name *string `json:"name,omitempty"`
	Status *PinStatus `json:"status,omitempty"`
	// This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: YYYY-MM-DDTHH:MM:SSZ.
	Created *time.Time `json:"created,omitempty"`
	Size *float32 `json:"size,omitempty"`
}

// NewPin instantiates a new Pin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPin() *Pin {
	this := Pin{}
	return &this
}

// NewPinWithDefaults instantiates a new Pin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPinWithDefaults() *Pin {
	this := Pin{}
	return &this
}

// GetCid returns the Cid field value if set, zero value otherwise.
func (o *Pin) GetCid() string {
	if o == nil || o.Cid == nil {
		var ret string
		return ret
	}
	return *o.Cid
}

// GetCidOk returns a tuple with the Cid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pin) GetCidOk() (*string, bool) {
	if o == nil || o.Cid == nil {
		return nil, false
	}
	return o.Cid, true
}

// HasCid returns a boolean if a field has been set.
func (o *Pin) HasCid() bool {
	if o != nil && o.Cid != nil {
		return true
	}

	return false
}

// SetCid gets a reference to the given string and assigns it to the Cid field.
func (o *Pin) SetCid(v string) {
	o.Cid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Pin) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pin) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Pin) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Pin) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Pin) GetStatus() PinStatus {
	if o == nil || o.Status == nil {
		var ret PinStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pin) GetStatusOk() (*PinStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Pin) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PinStatus and assigns it to the Status field.
func (o *Pin) SetStatus(v PinStatus) {
	o.Status = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Pin) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pin) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Pin) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Pin) SetCreated(v time.Time) {
	o.Created = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Pin) GetSize() float32 {
	if o == nil || o.Size == nil {
		var ret float32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pin) GetSizeOk() (*float32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *Pin) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given float32 and assigns it to the Size field.
func (o *Pin) SetSize(v float32) {
	o.Size = &v
}

func (o Pin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cid != nil {
		toSerialize["cid"] = o.Cid
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullablePin struct {
	value *Pin
	isSet bool
}

func (v NullablePin) Get() *Pin {
	return v.value
}

func (v *NullablePin) Set(val *Pin) {
	v.value = val
	v.isSet = true
}

func (v NullablePin) IsSet() bool {
	return v.isSet
}

func (v *NullablePin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePin(val *Pin) *NullablePin {
	return &NullablePin{value: val, isSet: true}
}

func (v NullablePin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

