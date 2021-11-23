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

// LegacySubscriptionDetail struct for LegacySubscriptionDetail
type LegacySubscriptionDetail struct {
	SiteCount *map[string]interface{} `json:"siteCount,omitempty"`
	Sites *[]LegacySite `json:"sites,omitempty"`
	OfferSelection *map[string]interface{} `json:"offerSelection,omitempty"`
	ServiceInstanceDetail *map[string]interface{} `json:"serviceInstanceDetail,omitempty"`
	PriceDetail *map[string]interface{} `json:"priceDetail,omitempty"`
	DealerCode *string `json:"dealerCode,omitempty"`
	PricePlanId *string `json:"pricePlanId,omitempty"`
	TermsAndConditionId *string `json:"termsAndConditionId,omitempty"`
	Configuration *map[string]string `json:"configuration,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LegacySubscriptionDetail LegacySubscriptionDetail

// NewLegacySubscriptionDetail instantiates a new LegacySubscriptionDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLegacySubscriptionDetail() *LegacySubscriptionDetail {
	this := LegacySubscriptionDetail{}
	return &this
}

// NewLegacySubscriptionDetailWithDefaults instantiates a new LegacySubscriptionDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLegacySubscriptionDetailWithDefaults() *LegacySubscriptionDetail {
	this := LegacySubscriptionDetail{}
	return &this
}

// GetSiteCount returns the SiteCount field value if set, zero value otherwise.
func (o *LegacySubscriptionDetail) GetSiteCount() map[string]interface{} {
	if o == nil || o.SiteCount == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.SiteCount
}

// GetSiteCountOk returns a tuple with the SiteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacySubscriptionDetail) GetSiteCountOk() (*map[string]interface{}, bool) {
	if o == nil || o.SiteCount == nil {
		return nil, false
	}
	return o.SiteCount, true
}

// HasSiteCount returns a boolean if a field has been set.
func (o *LegacySubscriptionDetail) HasSiteCount() bool {
	if o != nil && o.SiteCount != nil {
		return true
	}

	return false
}

// SetSiteCount gets a reference to the given map[string]interface{} and assigns it to the SiteCount field.
func (o *LegacySubscriptionDetail) SetSiteCount(v map[string]interface{}) {
	o.SiteCount = &v
}

// GetSites returns the Sites field value if set, zero value otherwise.
func (o *LegacySubscriptionDetail) GetSites() []LegacySite {
	if o == nil || o.Sites == nil {
		var ret []LegacySite
		return ret
	}
	return *o.Sites
}

// GetSitesOk returns a tuple with the Sites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacySubscriptionDetail) GetSitesOk() (*[]LegacySite, bool) {
	if o == nil || o.Sites == nil {
		return nil, false
	}
	return o.Sites, true
}

// HasSites returns a boolean if a field has been set.
func (o *LegacySubscriptionDetail) HasSites() bool {
	if o != nil && o.Sites != nil {
		return true
	}

	return false
}

// SetSites gets a reference to the given []LegacySite and assigns it to the Sites field.
func (o *LegacySubscriptionDetail) SetSites(v []LegacySite) {
	o.Sites = &v
}

// GetOfferSelection returns the OfferSelection field value if set, zero value otherwise.
func (o *LegacySubscriptionDetail) GetOfferSelection() map[string]interface{} {
	if o == nil || o.OfferSelection == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.OfferSelection
}

// GetOfferSelectionOk returns a tuple with the OfferSelection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacySubscriptionDetail) GetOfferSelectionOk() (*map[string]interface{}, bool) {
	if o == nil || o.OfferSelection == nil {
		return nil, false
	}
	return o.OfferSelection, true
}

// HasOfferSelection returns a boolean if a field has been set.
func (o *LegacySubscriptionDetail) HasOfferSelection() bool {
	if o != nil && o.OfferSelection != nil {
		return true
	}

	return false
}

// SetOfferSelection gets a reference to the given map[string]interface{} and assigns it to the OfferSelection field.
func (o *LegacySubscriptionDetail) SetOfferSelection(v map[string]interface{}) {
	o.OfferSelection = &v
}

// GetServiceInstanceDetail returns the ServiceInstanceDetail field value if set, zero value otherwise.
func (o *LegacySubscriptionDetail) GetServiceInstanceDetail() map[string]interface{} {
	if o == nil || o.ServiceInstanceDetail == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.ServiceInstanceDetail
}

// GetServiceInstanceDetailOk returns a tuple with the ServiceInstanceDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacySubscriptionDetail) GetServiceInstanceDetailOk() (*map[string]interface{}, bool) {
	if o == nil || o.ServiceInstanceDetail == nil {
		return nil, false
	}
	return o.ServiceInstanceDetail, true
}

// HasServiceInstanceDetail returns a boolean if a field has been set.
func (o *LegacySubscriptionDetail) HasServiceInstanceDetail() bool {
	if o != nil && o.ServiceInstanceDetail != nil {
		return true
	}

	return false
}

// SetServiceInstanceDetail gets a reference to the given map[string]interface{} and assigns it to the ServiceInstanceDetail field.
func (o *LegacySubscriptionDetail) SetServiceInstanceDetail(v map[string]interface{}) {
	o.ServiceInstanceDetail = &v
}

// GetPriceDetail returns the PriceDetail field value if set, zero value otherwise.
func (o *LegacySubscriptionDetail) GetPriceDetail() map[string]interface{} {
	if o == nil || o.PriceDetail == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.PriceDetail
}

// GetPriceDetailOk returns a tuple with the PriceDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacySubscriptionDetail) GetPriceDetailOk() (*map[string]interface{}, bool) {
	if o == nil || o.PriceDetail == nil {
		return nil, false
	}
	return o.PriceDetail, true
}

// HasPriceDetail returns a boolean if a field has been set.
func (o *LegacySubscriptionDetail) HasPriceDetail() bool {
	if o != nil && o.PriceDetail != nil {
		return true
	}

	return false
}

// SetPriceDetail gets a reference to the given map[string]interface{} and assigns it to the PriceDetail field.
func (o *LegacySubscriptionDetail) SetPriceDetail(v map[string]interface{}) {
	o.PriceDetail = &v
}

// GetDealerCode returns the DealerCode field value if set, zero value otherwise.
func (o *LegacySubscriptionDetail) GetDealerCode() string {
	if o == nil || o.DealerCode == nil {
		var ret string
		return ret
	}
	return *o.DealerCode
}

// GetDealerCodeOk returns a tuple with the DealerCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacySubscriptionDetail) GetDealerCodeOk() (*string, bool) {
	if o == nil || o.DealerCode == nil {
		return nil, false
	}
	return o.DealerCode, true
}

// HasDealerCode returns a boolean if a field has been set.
func (o *LegacySubscriptionDetail) HasDealerCode() bool {
	if o != nil && o.DealerCode != nil {
		return true
	}

	return false
}

// SetDealerCode gets a reference to the given string and assigns it to the DealerCode field.
func (o *LegacySubscriptionDetail) SetDealerCode(v string) {
	o.DealerCode = &v
}

// GetPricePlanId returns the PricePlanId field value if set, zero value otherwise.
func (o *LegacySubscriptionDetail) GetPricePlanId() string {
	if o == nil || o.PricePlanId == nil {
		var ret string
		return ret
	}
	return *o.PricePlanId
}

// GetPricePlanIdOk returns a tuple with the PricePlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacySubscriptionDetail) GetPricePlanIdOk() (*string, bool) {
	if o == nil || o.PricePlanId == nil {
		return nil, false
	}
	return o.PricePlanId, true
}

// HasPricePlanId returns a boolean if a field has been set.
func (o *LegacySubscriptionDetail) HasPricePlanId() bool {
	if o != nil && o.PricePlanId != nil {
		return true
	}

	return false
}

// SetPricePlanId gets a reference to the given string and assigns it to the PricePlanId field.
func (o *LegacySubscriptionDetail) SetPricePlanId(v string) {
	o.PricePlanId = &v
}

// GetTermsAndConditionId returns the TermsAndConditionId field value if set, zero value otherwise.
func (o *LegacySubscriptionDetail) GetTermsAndConditionId() string {
	if o == nil || o.TermsAndConditionId == nil {
		var ret string
		return ret
	}
	return *o.TermsAndConditionId
}

// GetTermsAndConditionIdOk returns a tuple with the TermsAndConditionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacySubscriptionDetail) GetTermsAndConditionIdOk() (*string, bool) {
	if o == nil || o.TermsAndConditionId == nil {
		return nil, false
	}
	return o.TermsAndConditionId, true
}

// HasTermsAndConditionId returns a boolean if a field has been set.
func (o *LegacySubscriptionDetail) HasTermsAndConditionId() bool {
	if o != nil && o.TermsAndConditionId != nil {
		return true
	}

	return false
}

// SetTermsAndConditionId gets a reference to the given string and assigns it to the TermsAndConditionId field.
func (o *LegacySubscriptionDetail) SetTermsAndConditionId(v string) {
	o.TermsAndConditionId = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *LegacySubscriptionDetail) GetConfiguration() map[string]string {
	if o == nil || o.Configuration == nil {
		var ret map[string]string
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacySubscriptionDetail) GetConfigurationOk() (*map[string]string, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *LegacySubscriptionDetail) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given map[string]string and assigns it to the Configuration field.
func (o *LegacySubscriptionDetail) SetConfiguration(v map[string]string) {
	o.Configuration = &v
}

func (o LegacySubscriptionDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SiteCount != nil {
		toSerialize["siteCount"] = o.SiteCount
	}
	if o.Sites != nil {
		toSerialize["sites"] = o.Sites
	}
	if o.OfferSelection != nil {
		toSerialize["offerSelection"] = o.OfferSelection
	}
	if o.ServiceInstanceDetail != nil {
		toSerialize["serviceInstanceDetail"] = o.ServiceInstanceDetail
	}
	if o.PriceDetail != nil {
		toSerialize["priceDetail"] = o.PriceDetail
	}
	if o.DealerCode != nil {
		toSerialize["dealerCode"] = o.DealerCode
	}
	if o.PricePlanId != nil {
		toSerialize["pricePlanId"] = o.PricePlanId
	}
	if o.TermsAndConditionId != nil {
		toSerialize["termsAndConditionId"] = o.TermsAndConditionId
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LegacySubscriptionDetail) UnmarshalJSON(bytes []byte) (err error) {
	varLegacySubscriptionDetail := _LegacySubscriptionDetail{}

	if err = json.Unmarshal(bytes, &varLegacySubscriptionDetail); err == nil {
		*o = LegacySubscriptionDetail(varLegacySubscriptionDetail)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "siteCount")
		delete(additionalProperties, "sites")
		delete(additionalProperties, "offerSelection")
		delete(additionalProperties, "serviceInstanceDetail")
		delete(additionalProperties, "priceDetail")
		delete(additionalProperties, "dealerCode")
		delete(additionalProperties, "pricePlanId")
		delete(additionalProperties, "termsAndConditionId")
		delete(additionalProperties, "configuration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLegacySubscriptionDetail struct {
	value *LegacySubscriptionDetail
	isSet bool
}

func (v NullableLegacySubscriptionDetail) Get() *LegacySubscriptionDetail {
	return v.value
}

func (v *NullableLegacySubscriptionDetail) Set(val *LegacySubscriptionDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableLegacySubscriptionDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableLegacySubscriptionDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLegacySubscriptionDetail(val *LegacySubscriptionDetail) *NullableLegacySubscriptionDetail {
	return &NullableLegacySubscriptionDetail{value: val, isSet: true}
}

func (v NullableLegacySubscriptionDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLegacySubscriptionDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


