# WorkflowStepRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uses** | Pointer to **string** |  | [optional] 
**After** | Pointer to **string** |  | [optional] 
**ScheduledAt** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewWorkflowStepRepresentation

`func NewWorkflowStepRepresentation() *WorkflowStepRepresentation`

NewWorkflowStepRepresentation instantiates a new WorkflowStepRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowStepRepresentationWithDefaults

`func NewWorkflowStepRepresentationWithDefaults() *WorkflowStepRepresentation`

NewWorkflowStepRepresentationWithDefaults instantiates a new WorkflowStepRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUses

`func (o *WorkflowStepRepresentation) GetUses() string`

GetUses returns the Uses field if non-nil, zero value otherwise.

### GetUsesOk

`func (o *WorkflowStepRepresentation) GetUsesOk() (*string, bool)`

GetUsesOk returns a tuple with the Uses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUses

`func (o *WorkflowStepRepresentation) SetUses(v string)`

SetUses sets Uses field to given value.

### HasUses

`func (o *WorkflowStepRepresentation) HasUses() bool`

HasUses returns a boolean if a field has been set.

### GetAfter

`func (o *WorkflowStepRepresentation) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *WorkflowStepRepresentation) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *WorkflowStepRepresentation) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *WorkflowStepRepresentation) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### GetScheduledAt

`func (o *WorkflowStepRepresentation) GetScheduledAt() int64`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *WorkflowStepRepresentation) GetScheduledAtOk() (*int64, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *WorkflowStepRepresentation) SetScheduledAt(v int64)`

SetScheduledAt sets ScheduledAt field to given value.

### HasScheduledAt

`func (o *WorkflowStepRepresentation) HasScheduledAt() bool`

HasScheduledAt returns a boolean if a field has been set.

### GetId

`func (o *WorkflowStepRepresentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowStepRepresentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowStepRepresentation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowStepRepresentation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConfig

`func (o *WorkflowStepRepresentation) GetConfig() map[string][]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *WorkflowStepRepresentation) GetConfigOk() (*map[string][]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *WorkflowStepRepresentation) SetConfig(v map[string][]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *WorkflowStepRepresentation) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


