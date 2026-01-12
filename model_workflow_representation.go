/*
Keycloak Admin REST API

This is a REST API reference for the Keycloak Admin REST API.

API version: 1.0
*/

package openapi

import (
	"encoding/json"
)

// checks if the WorkflowRepresentation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowRepresentation{}

// WorkflowRepresentation struct for WorkflowRepresentation
type WorkflowRepresentation struct {
	Id                *string                            `json:"id,omitempty"`
	Name              *string                            `json:"name,omitempty"`
	Enabled           *bool                              `json:"enabled,omitempty"`
	On                *string                            `json:"on,omitempty"`
	Schedule          *WorkflowScheduleRepresentation    `json:"schedule,omitempty"`
	Concurrency       *WorkflowConcurrencyRepresentation `json:"concurrency,omitempty"`
	If                *string                            `json:"if,omitempty"`
	Steps             []WorkflowStepRepresentation       `json:"steps,omitempty"`
	State             *WorkflowStateRepresentation       `json:"state,omitempty"`
	With              *map[string][]string               `json:"with,omitempty"`
	CancelInProgress  *string                            `json:"cancelInProgress,omitempty"`
	RestartInProgress *string                            `json:"restartInProgress,omitempty"`
}

// NewWorkflowRepresentation instantiates a new WorkflowRepresentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowRepresentation() *WorkflowRepresentation {
	this := WorkflowRepresentation{}
	return &this
}

// NewWorkflowRepresentationWithDefaults instantiates a new WorkflowRepresentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowRepresentationWithDefaults() *WorkflowRepresentation {
	this := WorkflowRepresentation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkflowRepresentation) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRepresentation) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkflowRepresentation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkflowRepresentation) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowRepresentation) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRepresentation) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowRepresentation) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowRepresentation) SetName(v string) {
	o.Name = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *WorkflowRepresentation) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRepresentation) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *WorkflowRepresentation) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *WorkflowRepresentation) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetOn returns the On field value if set, zero value otherwise.
func (o *WorkflowRepresentation) GetOn() string {
	if o == nil || IsNil(o.On) {
		var ret string
		return ret
	}
	return *o.On
}

// GetOnOk returns a tuple with the On field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRepresentation) GetOnOk() (*string, bool) {
	if o == nil || IsNil(o.On) {
		return nil, false
	}
	return o.On, true
}

// HasOn returns a boolean if a field has been set.
func (o *WorkflowRepresentation) HasOn() bool {
	if o != nil && !IsNil(o.On) {
		return true
	}

	return false
}

// SetOn gets a reference to the given string and assigns it to the On field.
func (o *WorkflowRepresentation) SetOn(v string) {
	o.On = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *WorkflowRepresentation) GetSchedule() WorkflowScheduleRepresentation {
	if o == nil || IsNil(o.Schedule) {
		var ret WorkflowScheduleRepresentation
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRepresentation) GetScheduleOk() (*WorkflowScheduleRepresentation, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *WorkflowRepresentation) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given WorkflowScheduleRepresentation and assigns it to the Schedule field.
func (o *WorkflowRepresentation) SetSchedule(v WorkflowScheduleRepresentation) {
	o.Schedule = &v
}

// GetConcurrency returns the Concurrency field value if set, zero value otherwise.
func (o *WorkflowRepresentation) GetConcurrency() WorkflowConcurrencyRepresentation {
	if o == nil || IsNil(o.Concurrency) {
		var ret WorkflowConcurrencyRepresentation
		return ret
	}
	return *o.Concurrency
}

// GetConcurrencyOk returns a tuple with the Concurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRepresentation) GetConcurrencyOk() (*WorkflowConcurrencyRepresentation, bool) {
	if o == nil || IsNil(o.Concurrency) {
		return nil, false
	}
	return o.Concurrency, true
}

// HasConcurrency returns a boolean if a field has been set.
func (o *WorkflowRepresentation) HasConcurrency() bool {
	if o != nil && !IsNil(o.Concurrency) {
		return true
	}

	return false
}

// SetConcurrency gets a reference to the given WorkflowConcurrencyRepresentation and assigns it to the Concurrency field.
func (o *WorkflowRepresentation) SetConcurrency(v WorkflowConcurrencyRepresentation) {
	o.Concurrency = &v
}

// GetIf returns the If field value if set, zero value otherwise.
func (o *WorkflowRepresentation) GetIf() string {
	if o == nil || IsNil(o.If) {
		var ret string
		return ret
	}
	return *o.If
}

// GetIfOk returns a tuple with the If field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRepresentation) GetIfOk() (*string, bool) {
	if o == nil || IsNil(o.If) {
		return nil, false
	}
	return o.If, true
}

// HasIf returns a boolean if a field has been set.
func (o *WorkflowRepresentation) HasIf() bool {
	if o != nil && !IsNil(o.If) {
		return true
	}

	return false
}

// SetIf gets a reference to the given string and assigns it to the If field.
func (o *WorkflowRepresentation) SetIf(v string) {
	o.If = &v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *WorkflowRepresentation) GetSteps() []WorkflowStepRepresentation {
	if o == nil || IsNil(o.Steps) {
		var ret []WorkflowStepRepresentation
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRepresentation) GetStepsOk() ([]WorkflowStepRepresentation, bool) {
	if o == nil || IsNil(o.Steps) {
		return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *WorkflowRepresentation) HasSteps() bool {
	if o != nil && !IsNil(o.Steps) {
		return true
	}

	return false
}

// SetSteps gets a reference to the given []WorkflowStepRepresentation and assigns it to the Steps field.
func (o *WorkflowRepresentation) SetSteps(v []WorkflowStepRepresentation) {
	o.Steps = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *WorkflowRepresentation) GetState() WorkflowStateRepresentation {
	if o == nil || IsNil(o.State) {
		var ret WorkflowStateRepresentation
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRepresentation) GetStateOk() (*WorkflowStateRepresentation, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *WorkflowRepresentation) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given WorkflowStateRepresentation and assigns it to the State field.
func (o *WorkflowRepresentation) SetState(v WorkflowStateRepresentation) {
	o.State = &v
}

// GetWith returns the With field value if set, zero value otherwise.
func (o *WorkflowRepresentation) GetWith() map[string][]string {
	if o == nil || IsNil(o.With) {
		var ret map[string][]string
		return ret
	}
	return *o.With
}

// GetWithOk returns a tuple with the With field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRepresentation) GetWithOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.With) {
		return nil, false
	}
	return o.With, true
}

// HasWith returns a boolean if a field has been set.
func (o *WorkflowRepresentation) HasWith() bool {
	if o != nil && !IsNil(o.With) {
		return true
	}

	return false
}

// SetWith gets a reference to the given map[string][]string and assigns it to the With field.
func (o *WorkflowRepresentation) SetWith(v map[string][]string) {
	o.With = &v
}

// GetCancelInProgress returns the CancelInProgress field value if set, zero value otherwise.
func (o *WorkflowRepresentation) GetCancelInProgress() string {
	if o == nil || IsNil(o.CancelInProgress) {
		var ret string
		return ret
	}
	return *o.CancelInProgress
}

// GetCancelInProgressOk returns a tuple with the CancelInProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRepresentation) GetCancelInProgressOk() (*string, bool) {
	if o == nil || IsNil(o.CancelInProgress) {
		return nil, false
	}
	return o.CancelInProgress, true
}

// HasCancelInProgress returns a boolean if a field has been set.
func (o *WorkflowRepresentation) HasCancelInProgress() bool {
	if o != nil && !IsNil(o.CancelInProgress) {
		return true
	}

	return false
}

// SetCancelInProgress gets a reference to the given string and assigns it to the CancelInProgress field.
func (o *WorkflowRepresentation) SetCancelInProgress(v string) {
	o.CancelInProgress = &v
}

// GetRestartInProgress returns the RestartInProgress field value if set, zero value otherwise.
func (o *WorkflowRepresentation) GetRestartInProgress() string {
	if o == nil || IsNil(o.RestartInProgress) {
		var ret string
		return ret
	}
	return *o.RestartInProgress
}

// GetRestartInProgressOk returns a tuple with the RestartInProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRepresentation) GetRestartInProgressOk() (*string, bool) {
	if o == nil || IsNil(o.RestartInProgress) {
		return nil, false
	}
	return o.RestartInProgress, true
}

// HasRestartInProgress returns a boolean if a field has been set.
func (o *WorkflowRepresentation) HasRestartInProgress() bool {
	if o != nil && !IsNil(o.RestartInProgress) {
		return true
	}

	return false
}

// SetRestartInProgress gets a reference to the given string and assigns it to the RestartInProgress field.
func (o *WorkflowRepresentation) SetRestartInProgress(v string) {
	o.RestartInProgress = &v
}

func (o WorkflowRepresentation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowRepresentation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.On) {
		toSerialize["on"] = o.On
	}
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	if !IsNil(o.Concurrency) {
		toSerialize["concurrency"] = o.Concurrency
	}
	if !IsNil(o.If) {
		toSerialize["if"] = o.If
	}
	if !IsNil(o.Steps) {
		toSerialize["steps"] = o.Steps
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.With) {
		toSerialize["with"] = o.With
	}
	if !IsNil(o.CancelInProgress) {
		toSerialize["cancelInProgress"] = o.CancelInProgress
	}
	if !IsNil(o.RestartInProgress) {
		toSerialize["restartInProgress"] = o.RestartInProgress
	}
	return toSerialize, nil
}

type NullableWorkflowRepresentation struct {
	value *WorkflowRepresentation
	isSet bool
}

func (v NullableWorkflowRepresentation) Get() *WorkflowRepresentation {
	return v.value
}

func (v *NullableWorkflowRepresentation) Set(val *WorkflowRepresentation) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowRepresentation) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowRepresentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowRepresentation(val *WorkflowRepresentation) *NullableWorkflowRepresentation {
	return &NullableWorkflowRepresentation{value: val, isSet: true}
}

func (v NullableWorkflowRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowRepresentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
