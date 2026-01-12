# ErrorRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**Params** | Pointer to **[]interface{}** |  | [optional] 
**Errors** | Pointer to [**[]ErrorRepresentation**](ErrorRepresentation.md) |  | [optional] 

## Methods

### NewErrorRepresentation

`func NewErrorRepresentation() *ErrorRepresentation`

NewErrorRepresentation instantiates a new ErrorRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorRepresentationWithDefaults

`func NewErrorRepresentationWithDefaults() *ErrorRepresentation`

NewErrorRepresentationWithDefaults instantiates a new ErrorRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *ErrorRepresentation) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *ErrorRepresentation) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *ErrorRepresentation) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *ErrorRepresentation) HasField() bool`

HasField returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ErrorRepresentation) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ErrorRepresentation) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ErrorRepresentation) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ErrorRepresentation) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetParams

`func (o *ErrorRepresentation) GetParams() []interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ErrorRepresentation) GetParamsOk() (*[]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ErrorRepresentation) SetParams(v []interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *ErrorRepresentation) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetErrors

`func (o *ErrorRepresentation) GetErrors() []ErrorRepresentation`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorRepresentation) GetErrorsOk() (*[]ErrorRepresentation, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorRepresentation) SetErrors(v []ErrorRepresentation)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ErrorRepresentation) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


