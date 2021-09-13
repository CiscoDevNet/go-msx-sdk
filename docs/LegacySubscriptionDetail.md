# LegacySubscriptionDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteCount** | Pointer to **map[string]interface{}** |  | [optional] 
**Sites** | Pointer to [**[]LegacySite**](LegacySite.md) |  | [optional] 
**OfferSelection** | Pointer to **map[string]interface{}** |  | [optional] 
**ServiceInstanceDetail** | Pointer to **map[string]interface{}** |  | [optional] 
**PriceDetail** | Pointer to **map[string]interface{}** |  | [optional] 
**DealerCode** | Pointer to **string** |  | [optional] 
**PricePlanId** | Pointer to **string** |  | [optional] 
**TermsAndConditionId** | Pointer to **string** |  | [optional] 
**Configuration** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewLegacySubscriptionDetail

`func NewLegacySubscriptionDetail() *LegacySubscriptionDetail`

NewLegacySubscriptionDetail instantiates a new LegacySubscriptionDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegacySubscriptionDetailWithDefaults

`func NewLegacySubscriptionDetailWithDefaults() *LegacySubscriptionDetail`

NewLegacySubscriptionDetailWithDefaults instantiates a new LegacySubscriptionDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteCount

`func (o *LegacySubscriptionDetail) GetSiteCount() map[string]interface{}`

GetSiteCount returns the SiteCount field if non-nil, zero value otherwise.

### GetSiteCountOk

`func (o *LegacySubscriptionDetail) GetSiteCountOk() (*map[string]interface{}, bool)`

GetSiteCountOk returns a tuple with the SiteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteCount

`func (o *LegacySubscriptionDetail) SetSiteCount(v map[string]interface{})`

SetSiteCount sets SiteCount field to given value.

### HasSiteCount

`func (o *LegacySubscriptionDetail) HasSiteCount() bool`

HasSiteCount returns a boolean if a field has been set.

### GetSites

`func (o *LegacySubscriptionDetail) GetSites() []LegacySite`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *LegacySubscriptionDetail) GetSitesOk() (*[]LegacySite, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *LegacySubscriptionDetail) SetSites(v []LegacySite)`

SetSites sets Sites field to given value.

### HasSites

`func (o *LegacySubscriptionDetail) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetOfferSelection

`func (o *LegacySubscriptionDetail) GetOfferSelection() map[string]interface{}`

GetOfferSelection returns the OfferSelection field if non-nil, zero value otherwise.

### GetOfferSelectionOk

`func (o *LegacySubscriptionDetail) GetOfferSelectionOk() (*map[string]interface{}, bool)`

GetOfferSelectionOk returns a tuple with the OfferSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferSelection

`func (o *LegacySubscriptionDetail) SetOfferSelection(v map[string]interface{})`

SetOfferSelection sets OfferSelection field to given value.

### HasOfferSelection

`func (o *LegacySubscriptionDetail) HasOfferSelection() bool`

HasOfferSelection returns a boolean if a field has been set.

### GetServiceInstanceDetail

`func (o *LegacySubscriptionDetail) GetServiceInstanceDetail() map[string]interface{}`

GetServiceInstanceDetail returns the ServiceInstanceDetail field if non-nil, zero value otherwise.

### GetServiceInstanceDetailOk

`func (o *LegacySubscriptionDetail) GetServiceInstanceDetailOk() (*map[string]interface{}, bool)`

GetServiceInstanceDetailOk returns a tuple with the ServiceInstanceDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceInstanceDetail

`func (o *LegacySubscriptionDetail) SetServiceInstanceDetail(v map[string]interface{})`

SetServiceInstanceDetail sets ServiceInstanceDetail field to given value.

### HasServiceInstanceDetail

`func (o *LegacySubscriptionDetail) HasServiceInstanceDetail() bool`

HasServiceInstanceDetail returns a boolean if a field has been set.

### GetPriceDetail

`func (o *LegacySubscriptionDetail) GetPriceDetail() map[string]interface{}`

GetPriceDetail returns the PriceDetail field if non-nil, zero value otherwise.

### GetPriceDetailOk

`func (o *LegacySubscriptionDetail) GetPriceDetailOk() (*map[string]interface{}, bool)`

GetPriceDetailOk returns a tuple with the PriceDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDetail

`func (o *LegacySubscriptionDetail) SetPriceDetail(v map[string]interface{})`

SetPriceDetail sets PriceDetail field to given value.

### HasPriceDetail

`func (o *LegacySubscriptionDetail) HasPriceDetail() bool`

HasPriceDetail returns a boolean if a field has been set.

### GetDealerCode

`func (o *LegacySubscriptionDetail) GetDealerCode() string`

GetDealerCode returns the DealerCode field if non-nil, zero value otherwise.

### GetDealerCodeOk

`func (o *LegacySubscriptionDetail) GetDealerCodeOk() (*string, bool)`

GetDealerCodeOk returns a tuple with the DealerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealerCode

`func (o *LegacySubscriptionDetail) SetDealerCode(v string)`

SetDealerCode sets DealerCode field to given value.

### HasDealerCode

`func (o *LegacySubscriptionDetail) HasDealerCode() bool`

HasDealerCode returns a boolean if a field has been set.

### GetPricePlanId

`func (o *LegacySubscriptionDetail) GetPricePlanId() string`

GetPricePlanId returns the PricePlanId field if non-nil, zero value otherwise.

### GetPricePlanIdOk

`func (o *LegacySubscriptionDetail) GetPricePlanIdOk() (*string, bool)`

GetPricePlanIdOk returns a tuple with the PricePlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePlanId

`func (o *LegacySubscriptionDetail) SetPricePlanId(v string)`

SetPricePlanId sets PricePlanId field to given value.

### HasPricePlanId

`func (o *LegacySubscriptionDetail) HasPricePlanId() bool`

HasPricePlanId returns a boolean if a field has been set.

### GetTermsAndConditionId

`func (o *LegacySubscriptionDetail) GetTermsAndConditionId() string`

GetTermsAndConditionId returns the TermsAndConditionId field if non-nil, zero value otherwise.

### GetTermsAndConditionIdOk

`func (o *LegacySubscriptionDetail) GetTermsAndConditionIdOk() (*string, bool)`

GetTermsAndConditionIdOk returns a tuple with the TermsAndConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsAndConditionId

`func (o *LegacySubscriptionDetail) SetTermsAndConditionId(v string)`

SetTermsAndConditionId sets TermsAndConditionId field to given value.

### HasTermsAndConditionId

`func (o *LegacySubscriptionDetail) HasTermsAndConditionId() bool`

HasTermsAndConditionId returns a boolean if a field has been set.

### GetConfiguration

`func (o *LegacySubscriptionDetail) GetConfiguration() map[string]string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *LegacySubscriptionDetail) GetConfigurationOk() (*map[string]string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *LegacySubscriptionDetail) SetConfiguration(v map[string]string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *LegacySubscriptionDetail) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


