/*
 * Apicurio Registry API [v2]
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0.1-SNAPSHOT
 * Contact: apicurio@lists.jboss.org
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package srsdata

import (
	"encoding/json"
	"fmt"
)

// IfExists the model 'IfExists'
type IfExists string

// List of IfExists
const (
	FAIL             IfExists = "FAIL"
	UPDATE           IfExists = "UPDATE"
	RETURN           IfExists = "RETURN"
	RETURN_OR_UPDATE IfExists = "RETURN_OR_UPDATE"
)

var allowedIfExistsEnumValues = []IfExists{
	"FAIL",
	"UPDATE",
	"RETURN",
	"RETURN_OR_UPDATE",
}

func (v *IfExists) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IfExists(value)
	for _, existing := range allowedIfExistsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IfExists", value)
}

// NewIfExistsFromValue returns a pointer to a valid IfExists
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIfExistsFromValue(v string) (*IfExists, error) {
	ev := IfExists(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IfExists: valid values are %v", v, allowedIfExistsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IfExists) IsValid() bool {
	for _, existing := range allowedIfExistsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IfExists value
func (v IfExists) Ptr() *IfExists {
	return &v
}

type NullableIfExists struct {
	value *IfExists
	isSet bool
}

func (v NullableIfExists) Get() *IfExists {
	return v.value
}

func (v *NullableIfExists) Set(val *IfExists) {
	v.value = val
	v.isSet = true
}

func (v NullableIfExists) IsSet() bool {
	return v.isSet
}

func (v *NullableIfExists) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIfExists(val *IfExists) *NullableIfExists {
	return &NullableIfExists{value: val, isSet: true}
}

func (v NullableIfExists) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIfExists) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
