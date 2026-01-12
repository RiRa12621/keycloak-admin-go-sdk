# OrganizationInvitationRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**SentDate** | Pointer to **int32** |  | [optional] 
**ExpiresAt** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**InviteLink** | Pointer to **string** |  | [optional] 

## Methods

### NewOrganizationInvitationRepresentation

`func NewOrganizationInvitationRepresentation() *OrganizationInvitationRepresentation`

NewOrganizationInvitationRepresentation instantiates a new OrganizationInvitationRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationInvitationRepresentationWithDefaults

`func NewOrganizationInvitationRepresentationWithDefaults() *OrganizationInvitationRepresentation`

NewOrganizationInvitationRepresentationWithDefaults instantiates a new OrganizationInvitationRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationInvitationRepresentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationInvitationRepresentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationInvitationRepresentation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationInvitationRepresentation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *OrganizationInvitationRepresentation) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OrganizationInvitationRepresentation) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OrganizationInvitationRepresentation) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *OrganizationInvitationRepresentation) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetEmail

`func (o *OrganizationInvitationRepresentation) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationInvitationRepresentation) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationInvitationRepresentation) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrganizationInvitationRepresentation) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *OrganizationInvitationRepresentation) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *OrganizationInvitationRepresentation) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *OrganizationInvitationRepresentation) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *OrganizationInvitationRepresentation) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *OrganizationInvitationRepresentation) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *OrganizationInvitationRepresentation) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *OrganizationInvitationRepresentation) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *OrganizationInvitationRepresentation) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetSentDate

`func (o *OrganizationInvitationRepresentation) GetSentDate() int32`

GetSentDate returns the SentDate field if non-nil, zero value otherwise.

### GetSentDateOk

`func (o *OrganizationInvitationRepresentation) GetSentDateOk() (*int32, bool)`

GetSentDateOk returns a tuple with the SentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentDate

`func (o *OrganizationInvitationRepresentation) SetSentDate(v int32)`

SetSentDate sets SentDate field to given value.

### HasSentDate

`func (o *OrganizationInvitationRepresentation) HasSentDate() bool`

HasSentDate returns a boolean if a field has been set.

### GetExpiresAt

`func (o *OrganizationInvitationRepresentation) GetExpiresAt() int32`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *OrganizationInvitationRepresentation) GetExpiresAtOk() (*int32, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *OrganizationInvitationRepresentation) SetExpiresAt(v int32)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *OrganizationInvitationRepresentation) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetStatus

`func (o *OrganizationInvitationRepresentation) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrganizationInvitationRepresentation) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrganizationInvitationRepresentation) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrganizationInvitationRepresentation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetInviteLink

`func (o *OrganizationInvitationRepresentation) GetInviteLink() string`

GetInviteLink returns the InviteLink field if non-nil, zero value otherwise.

### GetInviteLinkOk

`func (o *OrganizationInvitationRepresentation) GetInviteLinkOk() (*string, bool)`

GetInviteLinkOk returns a tuple with the InviteLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteLink

`func (o *OrganizationInvitationRepresentation) SetInviteLink(v string)`

SetInviteLink sets InviteLink field to given value.

### HasInviteLink

`func (o *OrganizationInvitationRepresentation) HasInviteLink() bool`

HasInviteLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


