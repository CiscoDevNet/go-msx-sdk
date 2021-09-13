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

// DeviceComplianceState the model 'DeviceComplianceState'
type DeviceComplianceState string

// List of DeviceComplianceState
const (
	DEVICECOMPLIANCESTATE_COMPLIANT DeviceComplianceState = "COMPLIANT"
	DEVICECOMPLIANCESTATE_NOT_COMPLIANT DeviceComplianceState = "NOT_COMPLIANT"
	DEVICECOMPLIANCESTATE_NOT_APPLICABLE DeviceComplianceState = "NOT_APPLICABLE"
	DEVICECOMPLIANCESTATE_UNKNOWN DeviceComplianceState = "UNKNOWN"
)

func (v *DeviceComplianceState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeviceComplianceState(value)
	for _, existing := range []DeviceComplianceState{ "COMPLIANT", "NOT_COMPLIANT", "NOT_APPLICABLE", "UNKNOWN",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeviceComplianceState", value)
}

// Ptr returns reference to DeviceComplianceState value
func (v DeviceComplianceState) Ptr() *DeviceComplianceState {
	return &v
}

type NullableDeviceComplianceState struct {
	value *DeviceComplianceState
	isSet bool
}

func (v NullableDeviceComplianceState) Get() *DeviceComplianceState {
	return v.value
}

func (v *NullableDeviceComplianceState) Set(val *DeviceComplianceState) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceComplianceState) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceComplianceState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceComplianceState(val *DeviceComplianceState) *NullableDeviceComplianceState {
	return &NullableDeviceComplianceState{value: val, isSet: true}
}

func (v NullableDeviceComplianceState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceComplianceState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

