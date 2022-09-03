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

// V1ListProjectsResponse struct for V1ListProjectsResponse
type V1ListProjectsResponse struct {
	Projects *[]V1Project `json:"projects,omitempty"`
	Total *string `json:"total,omitempty"`
	Current *string `json:"current,omitempty"`
	Page *string `json:"page,omitempty"`
}

// NewV1ListProjectsResponse instantiates a new V1ListProjectsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ListProjectsResponse() *V1ListProjectsResponse {
	this := V1ListProjectsResponse{}
	return &this
}

// NewV1ListProjectsResponseWithDefaults instantiates a new V1ListProjectsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ListProjectsResponseWithDefaults() *V1ListProjectsResponse {
	this := V1ListProjectsResponse{}
	return &this
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *V1ListProjectsResponse) GetProjects() []V1Project {
	if o == nil || o.Projects == nil {
		var ret []V1Project
		return ret
	}
	return *o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ListProjectsResponse) GetProjectsOk() (*[]V1Project, bool) {
	if o == nil || o.Projects == nil {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *V1ListProjectsResponse) HasProjects() bool {
	if o != nil && o.Projects != nil {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []V1Project and assigns it to the Projects field.
func (o *V1ListProjectsResponse) SetProjects(v []V1Project) {
	o.Projects = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *V1ListProjectsResponse) GetTotal() string {
	if o == nil || o.Total == nil {
		var ret string
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ListProjectsResponse) GetTotalOk() (*string, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *V1ListProjectsResponse) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given string and assigns it to the Total field.
func (o *V1ListProjectsResponse) SetTotal(v string) {
	o.Total = &v
}

// GetCurrent returns the Current field value if set, zero value otherwise.
func (o *V1ListProjectsResponse) GetCurrent() string {
	if o == nil || o.Current == nil {
		var ret string
		return ret
	}
	return *o.Current
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ListProjectsResponse) GetCurrentOk() (*string, bool) {
	if o == nil || o.Current == nil {
		return nil, false
	}
	return o.Current, true
}

// HasCurrent returns a boolean if a field has been set.
func (o *V1ListProjectsResponse) HasCurrent() bool {
	if o != nil && o.Current != nil {
		return true
	}

	return false
}

// SetCurrent gets a reference to the given string and assigns it to the Current field.
func (o *V1ListProjectsResponse) SetCurrent(v string) {
	o.Current = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *V1ListProjectsResponse) GetPage() string {
	if o == nil || o.Page == nil {
		var ret string
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ListProjectsResponse) GetPageOk() (*string, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *V1ListProjectsResponse) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given string and assigns it to the Page field.
func (o *V1ListProjectsResponse) SetPage(v string) {
	o.Page = &v
}

func (o V1ListProjectsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.Current != nil {
		toSerialize["current"] = o.Current
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	return json.Marshal(toSerialize)
}

type NullableV1ListProjectsResponse struct {
	value *V1ListProjectsResponse
	isSet bool
}

func (v NullableV1ListProjectsResponse) Get() *V1ListProjectsResponse {
	return v.value
}

func (v *NullableV1ListProjectsResponse) Set(val *V1ListProjectsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ListProjectsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ListProjectsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ListProjectsResponse(val *V1ListProjectsResponse) *NullableV1ListProjectsResponse {
	return &NullableV1ListProjectsResponse{value: val, isSet: true}
}

func (v NullableV1ListProjectsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ListProjectsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


