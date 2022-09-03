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

// InlineObject3 struct for InlineObject3
type InlineObject3 struct {
	// Previously uploaded RPM tarball
	Rpms *string `json:"rpms,omitempty"`
	// Overwrite existing RPMs even if NVRA is locked Useful for secure boot scenarios for example
	ForceOverride *bool `json:"forceOverride,omitempty"`
}

// NewInlineObject3 instantiates a new InlineObject3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject3() *InlineObject3 {
	this := InlineObject3{}
	return &this
}

// NewInlineObject3WithDefaults instantiates a new InlineObject3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject3WithDefaults() *InlineObject3 {
	this := InlineObject3{}
	return &this
}

// GetRpms returns the Rpms field value if set, zero value otherwise.
func (o *InlineObject3) GetRpms() string {
	if o == nil || o.Rpms == nil {
		var ret string
		return ret
	}
	return *o.Rpms
}

// GetRpmsOk returns a tuple with the Rpms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetRpmsOk() (*string, bool) {
	if o == nil || o.Rpms == nil {
		return nil, false
	}
	return o.Rpms, true
}

// HasRpms returns a boolean if a field has been set.
func (o *InlineObject3) HasRpms() bool {
	if o != nil && o.Rpms != nil {
		return true
	}

	return false
}

// SetRpms gets a reference to the given string and assigns it to the Rpms field.
func (o *InlineObject3) SetRpms(v string) {
	o.Rpms = &v
}

// GetForceOverride returns the ForceOverride field value if set, zero value otherwise.
func (o *InlineObject3) GetForceOverride() bool {
	if o == nil || o.ForceOverride == nil {
		var ret bool
		return ret
	}
	return *o.ForceOverride
}

// GetForceOverrideOk returns a tuple with the ForceOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetForceOverrideOk() (*bool, bool) {
	if o == nil || o.ForceOverride == nil {
		return nil, false
	}
	return o.ForceOverride, true
}

// HasForceOverride returns a boolean if a field has been set.
func (o *InlineObject3) HasForceOverride() bool {
	if o != nil && o.ForceOverride != nil {
		return true
	}

	return false
}

// SetForceOverride gets a reference to the given bool and assigns it to the ForceOverride field.
func (o *InlineObject3) SetForceOverride(v bool) {
	o.ForceOverride = &v
}

func (o InlineObject3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rpms != nil {
		toSerialize["rpms"] = o.Rpms
	}
	if o.ForceOverride != nil {
		toSerialize["forceOverride"] = o.ForceOverride
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject3 struct {
	value *InlineObject3
	isSet bool
}

func (v NullableInlineObject3) Get() *InlineObject3 {
	return v.value
}

func (v *NullableInlineObject3) Set(val *InlineObject3) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject3) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject3(val *InlineObject3) *NullableInlineObject3 {
	return &NullableInlineObject3{value: val, isSet: true}
}

func (v NullableInlineObject3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


