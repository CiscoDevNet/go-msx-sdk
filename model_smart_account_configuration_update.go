/*
 * MSX SDK
 *
 * MSX SDK client.
 *
 * API version: 1.0.8
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msxsdk

import (
	"encoding/json"
)

// SmartAccountConfigurationUpdate struct for SmartAccountConfigurationUpdate
type SmartAccountConfigurationUpdate struct {
	BaseAuthUrl *string `json:"baseAuthUrl,omitempty"`
	BaseSmartUrl *string `json:"baseSmartUrl,omitempty"`
	ContentType *string `json:"contentType,omitempty"`
	GrantType *string `json:"grantType,omitempty"`
	ClientId *string `json:"clientId,omitempty"`
	ClientSecret *string `json:"clientSecret,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SmartAccountConfigurationUpdate SmartAccountConfigurationUpdate

// NewSmartAccountConfigurationUpdate instantiates a new SmartAccountConfigurationUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartAccountConfigurationUpdate() *SmartAccountConfigurationUpdate {
	this := SmartAccountConfigurationUpdate{}
	return &this
}

// NewSmartAccountConfigurationUpdateWithDefaults instantiates a new SmartAccountConfigurationUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartAccountConfigurationUpdateWithDefaults() *SmartAccountConfigurationUpdate {
	this := SmartAccountConfigurationUpdate{}
	return &this
}

// GetBaseAuthUrl returns the BaseAuthUrl field value if set, zero value otherwise.
func (o *SmartAccountConfigurationUpdate) GetBaseAuthUrl() string {
	if o == nil || o.BaseAuthUrl == nil {
		var ret string
		return ret
	}
	return *o.BaseAuthUrl
}

// GetBaseAuthUrlOk returns a tuple with the BaseAuthUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartAccountConfigurationUpdate) GetBaseAuthUrlOk() (*string, bool) {
	if o == nil || o.BaseAuthUrl == nil {
		return nil, false
	}
	return o.BaseAuthUrl, true
}

// HasBaseAuthUrl returns a boolean if a field has been set.
func (o *SmartAccountConfigurationUpdate) HasBaseAuthUrl() bool {
	if o != nil && o.BaseAuthUrl != nil {
		return true
	}

	return false
}

// SetBaseAuthUrl gets a reference to the given string and assigns it to the BaseAuthUrl field.
func (o *SmartAccountConfigurationUpdate) SetBaseAuthUrl(v string) {
	o.BaseAuthUrl = &v
}

// GetBaseSmartUrl returns the BaseSmartUrl field value if set, zero value otherwise.
func (o *SmartAccountConfigurationUpdate) GetBaseSmartUrl() string {
	if o == nil || o.BaseSmartUrl == nil {
		var ret string
		return ret
	}
	return *o.BaseSmartUrl
}

// GetBaseSmartUrlOk returns a tuple with the BaseSmartUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartAccountConfigurationUpdate) GetBaseSmartUrlOk() (*string, bool) {
	if o == nil || o.BaseSmartUrl == nil {
		return nil, false
	}
	return o.BaseSmartUrl, true
}

// HasBaseSmartUrl returns a boolean if a field has been set.
func (o *SmartAccountConfigurationUpdate) HasBaseSmartUrl() bool {
	if o != nil && o.BaseSmartUrl != nil {
		return true
	}

	return false
}

// SetBaseSmartUrl gets a reference to the given string and assigns it to the BaseSmartUrl field.
func (o *SmartAccountConfigurationUpdate) SetBaseSmartUrl(v string) {
	o.BaseSmartUrl = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *SmartAccountConfigurationUpdate) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartAccountConfigurationUpdate) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *SmartAccountConfigurationUpdate) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *SmartAccountConfigurationUpdate) SetContentType(v string) {
	o.ContentType = &v
}

// GetGrantType returns the GrantType field value if set, zero value otherwise.
func (o *SmartAccountConfigurationUpdate) GetGrantType() string {
	if o == nil || o.GrantType == nil {
		var ret string
		return ret
	}
	return *o.GrantType
}

// GetGrantTypeOk returns a tuple with the GrantType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartAccountConfigurationUpdate) GetGrantTypeOk() (*string, bool) {
	if o == nil || o.GrantType == nil {
		return nil, false
	}
	return o.GrantType, true
}

// HasGrantType returns a boolean if a field has been set.
func (o *SmartAccountConfigurationUpdate) HasGrantType() bool {
	if o != nil && o.GrantType != nil {
		return true
	}

	return false
}

// SetGrantType gets a reference to the given string and assigns it to the GrantType field.
func (o *SmartAccountConfigurationUpdate) SetGrantType(v string) {
	o.GrantType = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *SmartAccountConfigurationUpdate) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartAccountConfigurationUpdate) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *SmartAccountConfigurationUpdate) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *SmartAccountConfigurationUpdate) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *SmartAccountConfigurationUpdate) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartAccountConfigurationUpdate) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *SmartAccountConfigurationUpdate) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *SmartAccountConfigurationUpdate) SetClientSecret(v string) {
	o.ClientSecret = &v
}

func (o SmartAccountConfigurationUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BaseAuthUrl != nil {
		toSerialize["baseAuthUrl"] = o.BaseAuthUrl
	}
	if o.BaseSmartUrl != nil {
		toSerialize["baseSmartUrl"] = o.BaseSmartUrl
	}
	if o.ContentType != nil {
		toSerialize["contentType"] = o.ContentType
	}
	if o.GrantType != nil {
		toSerialize["grantType"] = o.GrantType
	}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["clientSecret"] = o.ClientSecret
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SmartAccountConfigurationUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varSmartAccountConfigurationUpdate := _SmartAccountConfigurationUpdate{}

	if err = json.Unmarshal(bytes, &varSmartAccountConfigurationUpdate); err == nil {
		*o = SmartAccountConfigurationUpdate(varSmartAccountConfigurationUpdate)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "baseAuthUrl")
		delete(additionalProperties, "baseSmartUrl")
		delete(additionalProperties, "contentType")
		delete(additionalProperties, "grantType")
		delete(additionalProperties, "clientId")
		delete(additionalProperties, "clientSecret")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSmartAccountConfigurationUpdate struct {
	value *SmartAccountConfigurationUpdate
	isSet bool
}

func (v NullableSmartAccountConfigurationUpdate) Get() *SmartAccountConfigurationUpdate {
	return v.value
}

func (v *NullableSmartAccountConfigurationUpdate) Set(val *SmartAccountConfigurationUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartAccountConfigurationUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartAccountConfigurationUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartAccountConfigurationUpdate(val *SmartAccountConfigurationUpdate) *NullableSmartAccountConfigurationUpdate {
	return &NullableSmartAccountConfigurationUpdate{value: val, isSet: true}
}

func (v NullableSmartAccountConfigurationUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartAccountConfigurationUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


