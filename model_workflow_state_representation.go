/*
Keycloak Admin REST API

This is a REST API reference for the Keycloak Admin REST API.

API version: 1.0
*/

package openapi

import (
	"encoding/json"
)

// checks if the WorkflowStateRepresentation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowStateRepresentation{}

// WorkflowStateRepresentation struct for WorkflowStateRepresentation
type WorkflowStateRepresentation struct {
	Errors []string `json:"errors,omitempty"`
}

// NewWorkflowStateRepresentation instantiates a new WorkflowStateRepresentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowStateRepresentation() *WorkflowStateRepresentation {
	this := WorkflowStateRepresentation{}
	return &this
}

// NewWorkflowStateRepresentationWithDefaults instantiates a new WorkflowStateRepresentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowStateRepresentationWithDefaults() *WorkflowStateRepresentation {
	this := WorkflowStateRepresentation{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *WorkflowStateRepresentation) GetErrors() []string {
	if o == nil || IsNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateRepresentation) GetErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *WorkflowStateRepresentation) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *WorkflowStateRepresentation) SetErrors(v []string) {
	o.Errors = v
}

func (o WorkflowStateRepresentation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowStateRepresentation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableWorkflowStateRepresentation struct {
	value *WorkflowStateRepresentation
	isSet bool
}

func (v NullableWorkflowStateRepresentation) Get() *WorkflowStateRepresentation {
	return v.value
}

func (v *NullableWorkflowStateRepresentation) Set(val *WorkflowStateRepresentation) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowStateRepresentation) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowStateRepresentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowStateRepresentation(val *WorkflowStateRepresentation) *NullableWorkflowStateRepresentation {
	return &NullableWorkflowStateRepresentation{value: val, isSet: true}
}

func (v NullableWorkflowStateRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowStateRepresentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
