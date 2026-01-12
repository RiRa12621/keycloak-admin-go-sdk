/*
Keycloak Admin REST API

This is a REST API reference for the Keycloak Admin REST API.

API version: 1.0
*/

package openapi

import (
	"encoding/json"
)

// checks if the KeyStoreConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyStoreConfig{}

// KeyStoreConfig struct for KeyStoreConfig
type KeyStoreConfig struct {
	RealmCertificate *bool   `json:"realmCertificate,omitempty"`
	StorePassword    *string `json:"storePassword,omitempty"`
	KeyPassword      *string `json:"keyPassword,omitempty"`
	KeyAlias         *string `json:"keyAlias,omitempty"`
	RealmAlias       *string `json:"realmAlias,omitempty"`
	Format           *string `json:"format,omitempty"`
	KeySize          *int32  `json:"keySize,omitempty"`
	Validity         *int32  `json:"validity,omitempty"`
}

// NewKeyStoreConfig instantiates a new KeyStoreConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyStoreConfig() *KeyStoreConfig {
	this := KeyStoreConfig{}
	return &this
}

// NewKeyStoreConfigWithDefaults instantiates a new KeyStoreConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyStoreConfigWithDefaults() *KeyStoreConfig {
	this := KeyStoreConfig{}
	return &this
}

// GetRealmCertificate returns the RealmCertificate field value if set, zero value otherwise.
func (o *KeyStoreConfig) GetRealmCertificate() bool {
	if o == nil || IsNil(o.RealmCertificate) {
		var ret bool
		return ret
	}
	return *o.RealmCertificate
}

// GetRealmCertificateOk returns a tuple with the RealmCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyStoreConfig) GetRealmCertificateOk() (*bool, bool) {
	if o == nil || IsNil(o.RealmCertificate) {
		return nil, false
	}
	return o.RealmCertificate, true
}

// HasRealmCertificate returns a boolean if a field has been set.
func (o *KeyStoreConfig) HasRealmCertificate() bool {
	if o != nil && !IsNil(o.RealmCertificate) {
		return true
	}

	return false
}

// SetRealmCertificate gets a reference to the given bool and assigns it to the RealmCertificate field.
func (o *KeyStoreConfig) SetRealmCertificate(v bool) {
	o.RealmCertificate = &v
}

// GetStorePassword returns the StorePassword field value if set, zero value otherwise.
func (o *KeyStoreConfig) GetStorePassword() string {
	if o == nil || IsNil(o.StorePassword) {
		var ret string
		return ret
	}
	return *o.StorePassword
}

// GetStorePasswordOk returns a tuple with the StorePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyStoreConfig) GetStorePasswordOk() (*string, bool) {
	if o == nil || IsNil(o.StorePassword) {
		return nil, false
	}
	return o.StorePassword, true
}

// HasStorePassword returns a boolean if a field has been set.
func (o *KeyStoreConfig) HasStorePassword() bool {
	if o != nil && !IsNil(o.StorePassword) {
		return true
	}

	return false
}

// SetStorePassword gets a reference to the given string and assigns it to the StorePassword field.
func (o *KeyStoreConfig) SetStorePassword(v string) {
	o.StorePassword = &v
}

// GetKeyPassword returns the KeyPassword field value if set, zero value otherwise.
func (o *KeyStoreConfig) GetKeyPassword() string {
	if o == nil || IsNil(o.KeyPassword) {
		var ret string
		return ret
	}
	return *o.KeyPassword
}

// GetKeyPasswordOk returns a tuple with the KeyPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyStoreConfig) GetKeyPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.KeyPassword) {
		return nil, false
	}
	return o.KeyPassword, true
}

// HasKeyPassword returns a boolean if a field has been set.
func (o *KeyStoreConfig) HasKeyPassword() bool {
	if o != nil && !IsNil(o.KeyPassword) {
		return true
	}

	return false
}

// SetKeyPassword gets a reference to the given string and assigns it to the KeyPassword field.
func (o *KeyStoreConfig) SetKeyPassword(v string) {
	o.KeyPassword = &v
}

// GetKeyAlias returns the KeyAlias field value if set, zero value otherwise.
func (o *KeyStoreConfig) GetKeyAlias() string {
	if o == nil || IsNil(o.KeyAlias) {
		var ret string
		return ret
	}
	return *o.KeyAlias
}

// GetKeyAliasOk returns a tuple with the KeyAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyStoreConfig) GetKeyAliasOk() (*string, bool) {
	if o == nil || IsNil(o.KeyAlias) {
		return nil, false
	}
	return o.KeyAlias, true
}

// HasKeyAlias returns a boolean if a field has been set.
func (o *KeyStoreConfig) HasKeyAlias() bool {
	if o != nil && !IsNil(o.KeyAlias) {
		return true
	}

	return false
}

// SetKeyAlias gets a reference to the given string and assigns it to the KeyAlias field.
func (o *KeyStoreConfig) SetKeyAlias(v string) {
	o.KeyAlias = &v
}

// GetRealmAlias returns the RealmAlias field value if set, zero value otherwise.
func (o *KeyStoreConfig) GetRealmAlias() string {
	if o == nil || IsNil(o.RealmAlias) {
		var ret string
		return ret
	}
	return *o.RealmAlias
}

// GetRealmAliasOk returns a tuple with the RealmAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyStoreConfig) GetRealmAliasOk() (*string, bool) {
	if o == nil || IsNil(o.RealmAlias) {
		return nil, false
	}
	return o.RealmAlias, true
}

// HasRealmAlias returns a boolean if a field has been set.
func (o *KeyStoreConfig) HasRealmAlias() bool {
	if o != nil && !IsNil(o.RealmAlias) {
		return true
	}

	return false
}

// SetRealmAlias gets a reference to the given string and assigns it to the RealmAlias field.
func (o *KeyStoreConfig) SetRealmAlias(v string) {
	o.RealmAlias = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *KeyStoreConfig) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyStoreConfig) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *KeyStoreConfig) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *KeyStoreConfig) SetFormat(v string) {
	o.Format = &v
}

// GetKeySize returns the KeySize field value if set, zero value otherwise.
func (o *KeyStoreConfig) GetKeySize() int32 {
	if o == nil || IsNil(o.KeySize) {
		var ret int32
		return ret
	}
	return *o.KeySize
}

// GetKeySizeOk returns a tuple with the KeySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyStoreConfig) GetKeySizeOk() (*int32, bool) {
	if o == nil || IsNil(o.KeySize) {
		return nil, false
	}
	return o.KeySize, true
}

// HasKeySize returns a boolean if a field has been set.
func (o *KeyStoreConfig) HasKeySize() bool {
	if o != nil && !IsNil(o.KeySize) {
		return true
	}

	return false
}

// SetKeySize gets a reference to the given int32 and assigns it to the KeySize field.
func (o *KeyStoreConfig) SetKeySize(v int32) {
	o.KeySize = &v
}

// GetValidity returns the Validity field value if set, zero value otherwise.
func (o *KeyStoreConfig) GetValidity() int32 {
	if o == nil || IsNil(o.Validity) {
		var ret int32
		return ret
	}
	return *o.Validity
}

// GetValidityOk returns a tuple with the Validity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyStoreConfig) GetValidityOk() (*int32, bool) {
	if o == nil || IsNil(o.Validity) {
		return nil, false
	}
	return o.Validity, true
}

// HasValidity returns a boolean if a field has been set.
func (o *KeyStoreConfig) HasValidity() bool {
	if o != nil && !IsNil(o.Validity) {
		return true
	}

	return false
}

// SetValidity gets a reference to the given int32 and assigns it to the Validity field.
func (o *KeyStoreConfig) SetValidity(v int32) {
	o.Validity = &v
}

func (o KeyStoreConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyStoreConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RealmCertificate) {
		toSerialize["realmCertificate"] = o.RealmCertificate
	}
	if !IsNil(o.StorePassword) {
		toSerialize["storePassword"] = o.StorePassword
	}
	if !IsNil(o.KeyPassword) {
		toSerialize["keyPassword"] = o.KeyPassword
	}
	if !IsNil(o.KeyAlias) {
		toSerialize["keyAlias"] = o.KeyAlias
	}
	if !IsNil(o.RealmAlias) {
		toSerialize["realmAlias"] = o.RealmAlias
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.KeySize) {
		toSerialize["keySize"] = o.KeySize
	}
	if !IsNil(o.Validity) {
		toSerialize["validity"] = o.Validity
	}
	return toSerialize, nil
}

type NullableKeyStoreConfig struct {
	value *KeyStoreConfig
	isSet bool
}

func (v NullableKeyStoreConfig) Get() *KeyStoreConfig {
	return v.value
}

func (v *NullableKeyStoreConfig) Set(val *KeyStoreConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyStoreConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyStoreConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyStoreConfig(val *KeyStoreConfig) *NullableKeyStoreConfig {
	return &NullableKeyStoreConfig{value: val, isSet: true}
}

func (v NullableKeyStoreConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyStoreConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
