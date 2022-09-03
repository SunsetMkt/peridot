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
	"fmt"
)

// V1TaskStatus the model 'V1TaskStatus'
type V1TaskStatus string

// List of v1TaskStatus
const (
	UNSPECIFIED V1TaskStatus = "TASK_STATUS_UNSPECIFIED"
	PENDING V1TaskStatus = "TASK_STATUS_PENDING"
	RUNNING V1TaskStatus = "TASK_STATUS_RUNNING"
	SUCCEEDED V1TaskStatus = "TASK_STATUS_SUCCEEDED"
	FAILED V1TaskStatus = "TASK_STATUS_FAILED"
	CANCELED V1TaskStatus = "TASK_STATUS_CANCELED"
)

func (v *V1TaskStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V1TaskStatus(value)
	for _, existing := range []V1TaskStatus{ "TASK_STATUS_UNSPECIFIED", "TASK_STATUS_PENDING", "TASK_STATUS_RUNNING", "TASK_STATUS_SUCCEEDED", "TASK_STATUS_FAILED", "TASK_STATUS_CANCELED",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V1TaskStatus", value)
}

// Ptr returns reference to v1TaskStatus value
func (v V1TaskStatus) Ptr() *V1TaskStatus {
	return &v
}

type NullableV1TaskStatus struct {
	value *V1TaskStatus
	isSet bool
}

func (v NullableV1TaskStatus) Get() *V1TaskStatus {
	return v.value
}

func (v *NullableV1TaskStatus) Set(val *V1TaskStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableV1TaskStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableV1TaskStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1TaskStatus(val *V1TaskStatus) *NullableV1TaskStatus {
	return &NullableV1TaskStatus{value: val, isSet: true}
}

func (v NullableV1TaskStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1TaskStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

