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

// V1GetPackageResponse struct for V1GetPackageResponse
type V1GetPackageResponse struct {
	Package *V1Package `json:"package,omitempty"`
}

// NewV1GetPackageResponse instantiates a new V1GetPackageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1GetPackageResponse() *V1GetPackageResponse {
	this := V1GetPackageResponse{}
	return &this
}

// NewV1GetPackageResponseWithDefaults instantiates a new V1GetPackageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1GetPackageResponseWithDefaults() *V1GetPackageResponse {
	this := V1GetPackageResponse{}
	return &this
}

// GetPackage returns the Package field value if set, zero value otherwise.
func (o *V1GetPackageResponse) GetPackage() V1Package {
	if o == nil || o.Package == nil {
		var ret V1Package
		return ret
	}
	return *o.Package
}

// GetPackageOk returns a tuple with the Package field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GetPackageResponse) GetPackageOk() (*V1Package, bool) {
	if o == nil || o.Package == nil {
		return nil, false
	}
	return o.Package, true
}

// HasPackage returns a boolean if a field has been set.
func (o *V1GetPackageResponse) HasPackage() bool {
	if o != nil && o.Package != nil {
		return true
	}

	return false
}

// SetPackage gets a reference to the given V1Package and assigns it to the Package field.
func (o *V1GetPackageResponse) SetPackage(v V1Package) {
	o.Package = &v
}

func (o V1GetPackageResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Package != nil {
		toSerialize["package"] = o.Package
	}
	return json.Marshal(toSerialize)
}

type NullableV1GetPackageResponse struct {
	value *V1GetPackageResponse
	isSet bool
}

func (v NullableV1GetPackageResponse) Get() *V1GetPackageResponse {
	return v.value
}

func (v *NullableV1GetPackageResponse) Set(val *V1GetPackageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV1GetPackageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV1GetPackageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1GetPackageResponse(val *V1GetPackageResponse) *NullableV1GetPackageResponse {
	return &NullableV1GetPackageResponse{value: val, isSet: true}
}

func (v NullableV1GetPackageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1GetPackageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


