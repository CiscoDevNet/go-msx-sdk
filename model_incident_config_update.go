/*
MSX SDK

MSX SDK client.

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msxsdk

import (
	"encoding/json"
)

// IncidentConfigUpdate struct for IncidentConfigUpdate
type IncidentConfigUpdate struct {
	ClientId string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	CriticalEvent *bool `json:"criticalEvent,omitempty"`
	Domain string `json:"domain"`
	Password string `json:"password"`
	UserName string `json:"userName"`
	Proxy NullableString `json:"proxy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IncidentConfigUpdate IncidentConfigUpdate

// NewIncidentConfigUpdate instantiates a new IncidentConfigUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentConfigUpdate(clientId string, clientSecret string, domain string, password string, userName string) *IncidentConfigUpdate {
	this := IncidentConfigUpdate{}
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	this.Domain = domain
	this.Password = password
	this.UserName = userName
	return &this
}

// NewIncidentConfigUpdateWithDefaults instantiates a new IncidentConfigUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentConfigUpdateWithDefaults() *IncidentConfigUpdate {
	this := IncidentConfigUpdate{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *IncidentConfigUpdate) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *IncidentConfigUpdate) GetClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *IncidentConfigUpdate) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *IncidentConfigUpdate) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *IncidentConfigUpdate) GetClientSecretOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *IncidentConfigUpdate) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetCriticalEvent returns the CriticalEvent field value if set, zero value otherwise.
func (o *IncidentConfigUpdate) GetCriticalEvent() bool {
	if o == nil || o.CriticalEvent == nil {
		var ret bool
		return ret
	}
	return *o.CriticalEvent
}

// GetCriticalEventOk returns a tuple with the CriticalEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentConfigUpdate) GetCriticalEventOk() (*bool, bool) {
	if o == nil || o.CriticalEvent == nil {
		return nil, false
	}
	return o.CriticalEvent, true
}

// HasCriticalEvent returns a boolean if a field has been set.
func (o *IncidentConfigUpdate) HasCriticalEvent() bool {
	if o != nil && o.CriticalEvent != nil {
		return true
	}

	return false
}

// SetCriticalEvent gets a reference to the given bool and assigns it to the CriticalEvent field.
func (o *IncidentConfigUpdate) SetCriticalEvent(v bool) {
	o.CriticalEvent = &v
}

// GetDomain returns the Domain field value
func (o *IncidentConfigUpdate) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *IncidentConfigUpdate) GetDomainOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *IncidentConfigUpdate) SetDomain(v string) {
	o.Domain = v
}

// GetPassword returns the Password field value
func (o *IncidentConfigUpdate) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *IncidentConfigUpdate) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *IncidentConfigUpdate) SetPassword(v string) {
	o.Password = v
}

// GetUserName returns the UserName field value
func (o *IncidentConfigUpdate) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *IncidentConfigUpdate) GetUserNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *IncidentConfigUpdate) SetUserName(v string) {
	o.UserName = v
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentConfigUpdate) GetProxy() string {
	if o == nil || o.Proxy.Get() == nil {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncidentConfigUpdate) GetProxyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *IncidentConfigUpdate) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *IncidentConfigUpdate) SetProxy(v string) {
	o.Proxy.Set(&v)
}
// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *IncidentConfigUpdate) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *IncidentConfigUpdate) UnsetProxy() {
	o.Proxy.Unset()
}

func (o IncidentConfigUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["clientId"] = o.ClientId
	}
	if true {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	if o.CriticalEvent != nil {
		toSerialize["criticalEvent"] = o.CriticalEvent
	}
	if true {
		toSerialize["domain"] = o.Domain
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["userName"] = o.UserName
	}
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IncidentConfigUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varIncidentConfigUpdate := _IncidentConfigUpdate{}

	if err = json.Unmarshal(bytes, &varIncidentConfigUpdate); err == nil {
		*o = IncidentConfigUpdate(varIncidentConfigUpdate)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "clientId")
		delete(additionalProperties, "clientSecret")
		delete(additionalProperties, "criticalEvent")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "password")
		delete(additionalProperties, "userName")
		delete(additionalProperties, "proxy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIncidentConfigUpdate struct {
	value *IncidentConfigUpdate
	isSet bool
}

func (v NullableIncidentConfigUpdate) Get() *IncidentConfigUpdate {
	return v.value
}

func (v *NullableIncidentConfigUpdate) Set(val *IncidentConfigUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentConfigUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentConfigUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentConfigUpdate(val *IncidentConfigUpdate) *NullableIncidentConfigUpdate {
	return &NullableIncidentConfigUpdate{value: val, isSet: true}
}

func (v NullableIncidentConfigUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentConfigUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


