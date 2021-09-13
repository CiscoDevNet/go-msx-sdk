/*
 * MSX SDK
 *
 * MSX SDK client.
 *
 * API version: 1.0.5
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msxsdk

import (
	"encoding/json"
	"fmt"
)

// GenericEventSeverity the model 'GenericEventSeverity'
type GenericEventSeverity string

// List of GenericEventSeverity
const (
	GENERICEVENTSEVERITY_CRITICAL GenericEventSeverity = "CRITICAL"
	GENERICEVENTSEVERITY_POOR GenericEventSeverity = "POOR"
	GENERICEVENTSEVERITY_FAIR GenericEventSeverity = "FAIR"
	GENERICEVENTSEVERITY_GOOD GenericEventSeverity = "GOOD"
	GENERICEVENTSEVERITY_UNKNOWN GenericEventSeverity = "UNKNOWN"
)

func (v *GenericEventSeverity) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GenericEventSeverity(value)
	for _, existing := range []GenericEventSeverity{ "CRITICAL", "POOR", "FAIR", "GOOD", "UNKNOWN",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GenericEventSeverity", value)
}

// Ptr returns reference to GenericEventSeverity value
func (v GenericEventSeverity) Ptr() *GenericEventSeverity {
	return &v
}

type NullableGenericEventSeverity struct {
	value *GenericEventSeverity
	isSet bool
}

func (v NullableGenericEventSeverity) Get() *GenericEventSeverity {
	return v.value
}

func (v *NullableGenericEventSeverity) Set(val *GenericEventSeverity) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericEventSeverity) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericEventSeverity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericEventSeverity(val *GenericEventSeverity) *NullableGenericEventSeverity {
	return &NullableGenericEventSeverity{value: val, isSet: true}
}

func (v NullableGenericEventSeverity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericEventSeverity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
