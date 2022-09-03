/*
 * peridot/proto/v1/batch.proto
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: version not set
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package peridotopenapi

import (
	"encoding/json"
)

// V1ListImportBatchesResponse struct for V1ListImportBatchesResponse
type V1ListImportBatchesResponse struct {
	ImportBatches *[]V1ImportBatch `json:"importBatches,omitempty"`
	Total *string `json:"total,omitempty"`
	Size *int32 `json:"size,omitempty"`
	Page *int32 `json:"page,omitempty"`
}

// NewV1ListImportBatchesResponse instantiates a new V1ListImportBatchesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ListImportBatchesResponse() *V1ListImportBatchesResponse {
	this := V1ListImportBatchesResponse{}
	return &this
}

// NewV1ListImportBatchesResponseWithDefaults instantiates a new V1ListImportBatchesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ListImportBatchesResponseWithDefaults() *V1ListImportBatchesResponse {
	this := V1ListImportBatchesResponse{}
	return &this
}

// GetImportBatches returns the ImportBatches field value if set, zero value otherwise.
func (o *V1ListImportBatchesResponse) GetImportBatches() []V1ImportBatch {
	if o == nil || o.ImportBatches == nil {
		var ret []V1ImportBatch
		return ret
	}
	return *o.ImportBatches
}

// GetImportBatchesOk returns a tuple with the ImportBatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ListImportBatchesResponse) GetImportBatchesOk() (*[]V1ImportBatch, bool) {
	if o == nil || o.ImportBatches == nil {
		return nil, false
	}
	return o.ImportBatches, true
}

// HasImportBatches returns a boolean if a field has been set.
func (o *V1ListImportBatchesResponse) HasImportBatches() bool {
	if o != nil && o.ImportBatches != nil {
		return true
	}

	return false
}

// SetImportBatches gets a reference to the given []V1ImportBatch and assigns it to the ImportBatches field.
func (o *V1ListImportBatchesResponse) SetImportBatches(v []V1ImportBatch) {
	o.ImportBatches = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *V1ListImportBatchesResponse) GetTotal() string {
	if o == nil || o.Total == nil {
		var ret string
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ListImportBatchesResponse) GetTotalOk() (*string, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *V1ListImportBatchesResponse) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given string and assigns it to the Total field.
func (o *V1ListImportBatchesResponse) SetTotal(v string) {
	o.Total = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *V1ListImportBatchesResponse) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ListImportBatchesResponse) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *V1ListImportBatchesResponse) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *V1ListImportBatchesResponse) SetSize(v int32) {
	o.Size = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *V1ListImportBatchesResponse) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ListImportBatchesResponse) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *V1ListImportBatchesResponse) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *V1ListImportBatchesResponse) SetPage(v int32) {
	o.Page = &v
}

func (o V1ListImportBatchesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImportBatches != nil {
		toSerialize["importBatches"] = o.ImportBatches
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	return json.Marshal(toSerialize)
}

type NullableV1ListImportBatchesResponse struct {
	value *V1ListImportBatchesResponse
	isSet bool
}

func (v NullableV1ListImportBatchesResponse) Get() *V1ListImportBatchesResponse {
	return v.value
}

func (v *NullableV1ListImportBatchesResponse) Set(val *V1ListImportBatchesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ListImportBatchesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ListImportBatchesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ListImportBatchesResponse(val *V1ListImportBatchesResponse) *NullableV1ListImportBatchesResponse {
	return &NullableV1ListImportBatchesResponse{value: val, isSet: true}
}

func (v NullableV1ListImportBatchesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ListImportBatchesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


