/*
Keycloak Admin REST API

This is a REST API reference for the Keycloak Admin REST API.

API version: 1.0
*/

package openapi

import (
	"encoding/json"
)

// checks if the WorkflowScheduleRepresentation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowScheduleRepresentation{}

// WorkflowScheduleRepresentation struct for WorkflowScheduleRepresentation
type WorkflowScheduleRepresentation struct {
	After     *string `json:"after,omitempty"`
	BatchSize *int32  `json:"batch-size,omitempty"`
}

// NewWorkflowScheduleRepresentation instantiates a new WorkflowScheduleRepresentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowScheduleRepresentation() *WorkflowScheduleRepresentation {
	this := WorkflowScheduleRepresentation{}
	return &this
}

// NewWorkflowScheduleRepresentationWithDefaults instantiates a new WorkflowScheduleRepresentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowScheduleRepresentationWithDefaults() *WorkflowScheduleRepresentation {
	this := WorkflowScheduleRepresentation{}
	return &this
}

// GetAfter returns the After field value if set, zero value otherwise.
func (o *WorkflowScheduleRepresentation) GetAfter() string {
	if o == nil || IsNil(o.After) {
		var ret string
		return ret
	}
	return *o.After
}

// GetAfterOk returns a tuple with the After field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowScheduleRepresentation) GetAfterOk() (*string, bool) {
	if o == nil || IsNil(o.After) {
		return nil, false
	}
	return o.After, true
}

// HasAfter returns a boolean if a field has been set.
func (o *WorkflowScheduleRepresentation) HasAfter() bool {
	if o != nil && !IsNil(o.After) {
		return true
	}

	return false
}

// SetAfter gets a reference to the given string and assigns it to the After field.
func (o *WorkflowScheduleRepresentation) SetAfter(v string) {
	o.After = &v
}

// GetBatchSize returns the BatchSize field value if set, zero value otherwise.
func (o *WorkflowScheduleRepresentation) GetBatchSize() int32 {
	if o == nil || IsNil(o.BatchSize) {
		var ret int32
		return ret
	}
	return *o.BatchSize
}

// GetBatchSizeOk returns a tuple with the BatchSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowScheduleRepresentation) GetBatchSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.BatchSize) {
		return nil, false
	}
	return o.BatchSize, true
}

// HasBatchSize returns a boolean if a field has been set.
func (o *WorkflowScheduleRepresentation) HasBatchSize() bool {
	if o != nil && !IsNil(o.BatchSize) {
		return true
	}

	return false
}

// SetBatchSize gets a reference to the given int32 and assigns it to the BatchSize field.
func (o *WorkflowScheduleRepresentation) SetBatchSize(v int32) {
	o.BatchSize = &v
}

func (o WorkflowScheduleRepresentation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowScheduleRepresentation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.After) {
		toSerialize["after"] = o.After
	}
	if !IsNil(o.BatchSize) {
		toSerialize["batch-size"] = o.BatchSize
	}
	return toSerialize, nil
}

type NullableWorkflowScheduleRepresentation struct {
	value *WorkflowScheduleRepresentation
	isSet bool
}

func (v NullableWorkflowScheduleRepresentation) Get() *WorkflowScheduleRepresentation {
	return v.value
}

func (v *NullableWorkflowScheduleRepresentation) Set(val *WorkflowScheduleRepresentation) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowScheduleRepresentation) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowScheduleRepresentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowScheduleRepresentation(val *WorkflowScheduleRepresentation) *NullableWorkflowScheduleRepresentation {
	return &NullableWorkflowScheduleRepresentation{value: val, isSet: true}
}

func (v NullableWorkflowScheduleRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowScheduleRepresentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
