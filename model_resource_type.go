/*
Keycloak Admin REST API

This is a REST API reference for the Keycloak Admin REST API.

API version: 1.0
*/

package openapi

import (
	"encoding/json"
)

// checks if the ResourceType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceType{}

// ResourceType struct for ResourceType
type ResourceType struct {
	Type         *string              `json:"type,omitempty"`
	Scopes       []string             `json:"scopes,omitempty"`
	ScopeAliases *map[string][]string `json:"scopeAliases,omitempty"`
	GroupType    *string              `json:"groupType,omitempty"`
}

// NewResourceType instantiates a new ResourceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceType() *ResourceType {
	this := ResourceType{}
	return &this
}

// NewResourceTypeWithDefaults instantiates a new ResourceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceTypeWithDefaults() *ResourceType {
	this := ResourceType{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ResourceType) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceType) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ResourceType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ResourceType) SetType(v string) {
	o.Type = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *ResourceType) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceType) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *ResourceType) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *ResourceType) SetScopes(v []string) {
	o.Scopes = v
}

// GetScopeAliases returns the ScopeAliases field value if set, zero value otherwise.
func (o *ResourceType) GetScopeAliases() map[string][]string {
	if o == nil || IsNil(o.ScopeAliases) {
		var ret map[string][]string
		return ret
	}
	return *o.ScopeAliases
}

// GetScopeAliasesOk returns a tuple with the ScopeAliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceType) GetScopeAliasesOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.ScopeAliases) {
		return nil, false
	}
	return o.ScopeAliases, true
}

// HasScopeAliases returns a boolean if a field has been set.
func (o *ResourceType) HasScopeAliases() bool {
	if o != nil && !IsNil(o.ScopeAliases) {
		return true
	}

	return false
}

// SetScopeAliases gets a reference to the given map[string][]string and assigns it to the ScopeAliases field.
func (o *ResourceType) SetScopeAliases(v map[string][]string) {
	o.ScopeAliases = &v
}

// GetGroupType returns the GroupType field value if set, zero value otherwise.
func (o *ResourceType) GetGroupType() string {
	if o == nil || IsNil(o.GroupType) {
		var ret string
		return ret
	}
	return *o.GroupType
}

// GetGroupTypeOk returns a tuple with the GroupType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceType) GetGroupTypeOk() (*string, bool) {
	if o == nil || IsNil(o.GroupType) {
		return nil, false
	}
	return o.GroupType, true
}

// HasGroupType returns a boolean if a field has been set.
func (o *ResourceType) HasGroupType() bool {
	if o != nil && !IsNil(o.GroupType) {
		return true
	}

	return false
}

// SetGroupType gets a reference to the given string and assigns it to the GroupType field.
func (o *ResourceType) SetGroupType(v string) {
	o.GroupType = &v
}

func (o ResourceType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.ScopeAliases) {
		toSerialize["scopeAliases"] = o.ScopeAliases
	}
	if !IsNil(o.GroupType) {
		toSerialize["groupType"] = o.GroupType
	}
	return toSerialize, nil
}

type NullableResourceType struct {
	value *ResourceType
	isSet bool
}

func (v NullableResourceType) Get() *ResourceType {
	return v.value
}

func (v *NullableResourceType) Set(val *ResourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceType(val *ResourceType) *NullableResourceType {
	return &NullableResourceType{value: val, isSet: true}
}

func (v NullableResourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
