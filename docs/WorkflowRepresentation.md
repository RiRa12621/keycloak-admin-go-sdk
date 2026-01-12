# WorkflowRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**On** | Pointer to **string** |  | [optional] 
**Schedule** | Pointer to [**WorkflowScheduleRepresentation**](WorkflowScheduleRepresentation.md) |  | [optional] 
**Concurrency** | Pointer to [**WorkflowConcurrencyRepresentation**](WorkflowConcurrencyRepresentation.md) |  | [optional] 
**If** | Pointer to **string** |  | [optional] 
**Steps** | Pointer to [**[]WorkflowStepRepresentation**](WorkflowStepRepresentation.md) |  | [optional] 
**State** | Pointer to [**WorkflowStateRepresentation**](WorkflowStateRepresentation.md) |  | [optional] 
**With** | Pointer to **map[string][]string** |  | [optional] 
**CancelInProgress** | Pointer to **string** |  | [optional] 
**RestartInProgress** | Pointer to **string** |  | [optional] 

## Methods

### NewWorkflowRepresentation

`func NewWorkflowRepresentation() *WorkflowRepresentation`

NewWorkflowRepresentation instantiates a new WorkflowRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowRepresentationWithDefaults

`func NewWorkflowRepresentationWithDefaults() *WorkflowRepresentation`

NewWorkflowRepresentationWithDefaults instantiates a new WorkflowRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowRepresentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowRepresentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowRepresentation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowRepresentation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WorkflowRepresentation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowRepresentation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowRepresentation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowRepresentation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *WorkflowRepresentation) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WorkflowRepresentation) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WorkflowRepresentation) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WorkflowRepresentation) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOn

`func (o *WorkflowRepresentation) GetOn() string`

GetOn returns the On field if non-nil, zero value otherwise.

### GetOnOk

`func (o *WorkflowRepresentation) GetOnOk() (*string, bool)`

GetOnOk returns a tuple with the On field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOn

`func (o *WorkflowRepresentation) SetOn(v string)`

SetOn sets On field to given value.

### HasOn

`func (o *WorkflowRepresentation) HasOn() bool`

HasOn returns a boolean if a field has been set.

### GetSchedule

`func (o *WorkflowRepresentation) GetSchedule() WorkflowScheduleRepresentation`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *WorkflowRepresentation) GetScheduleOk() (*WorkflowScheduleRepresentation, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *WorkflowRepresentation) SetSchedule(v WorkflowScheduleRepresentation)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *WorkflowRepresentation) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetConcurrency

`func (o *WorkflowRepresentation) GetConcurrency() WorkflowConcurrencyRepresentation`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *WorkflowRepresentation) GetConcurrencyOk() (*WorkflowConcurrencyRepresentation, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *WorkflowRepresentation) SetConcurrency(v WorkflowConcurrencyRepresentation)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *WorkflowRepresentation) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### GetIf

`func (o *WorkflowRepresentation) GetIf() string`

GetIf returns the If field if non-nil, zero value otherwise.

### GetIfOk

`func (o *WorkflowRepresentation) GetIfOk() (*string, bool)`

GetIfOk returns a tuple with the If field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIf

`func (o *WorkflowRepresentation) SetIf(v string)`

SetIf sets If field to given value.

### HasIf

`func (o *WorkflowRepresentation) HasIf() bool`

HasIf returns a boolean if a field has been set.

### GetSteps

`func (o *WorkflowRepresentation) GetSteps() []WorkflowStepRepresentation`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *WorkflowRepresentation) GetStepsOk() (*[]WorkflowStepRepresentation, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *WorkflowRepresentation) SetSteps(v []WorkflowStepRepresentation)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *WorkflowRepresentation) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetState

`func (o *WorkflowRepresentation) GetState() WorkflowStateRepresentation`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *WorkflowRepresentation) GetStateOk() (*WorkflowStateRepresentation, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *WorkflowRepresentation) SetState(v WorkflowStateRepresentation)`

SetState sets State field to given value.

### HasState

`func (o *WorkflowRepresentation) HasState() bool`

HasState returns a boolean if a field has been set.

### GetWith

`func (o *WorkflowRepresentation) GetWith() map[string][]string`

GetWith returns the With field if non-nil, zero value otherwise.

### GetWithOk

`func (o *WorkflowRepresentation) GetWithOk() (*map[string][]string, bool)`

GetWithOk returns a tuple with the With field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWith

`func (o *WorkflowRepresentation) SetWith(v map[string][]string)`

SetWith sets With field to given value.

### HasWith

`func (o *WorkflowRepresentation) HasWith() bool`

HasWith returns a boolean if a field has been set.

### GetCancelInProgress

`func (o *WorkflowRepresentation) GetCancelInProgress() string`

GetCancelInProgress returns the CancelInProgress field if non-nil, zero value otherwise.

### GetCancelInProgressOk

`func (o *WorkflowRepresentation) GetCancelInProgressOk() (*string, bool)`

GetCancelInProgressOk returns a tuple with the CancelInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelInProgress

`func (o *WorkflowRepresentation) SetCancelInProgress(v string)`

SetCancelInProgress sets CancelInProgress field to given value.

### HasCancelInProgress

`func (o *WorkflowRepresentation) HasCancelInProgress() bool`

HasCancelInProgress returns a boolean if a field has been set.

### GetRestartInProgress

`func (o *WorkflowRepresentation) GetRestartInProgress() string`

GetRestartInProgress returns the RestartInProgress field if non-nil, zero value otherwise.

### GetRestartInProgressOk

`func (o *WorkflowRepresentation) GetRestartInProgressOk() (*string, bool)`

GetRestartInProgressOk returns a tuple with the RestartInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartInProgress

`func (o *WorkflowRepresentation) SetRestartInProgress(v string)`

SetRestartInProgress sets RestartInProgress field to given value.

### HasRestartInProgress

`func (o *WorkflowRepresentation) HasRestartInProgress() bool`

HasRestartInProgress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


