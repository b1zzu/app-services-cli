/*
 * Service Registry Service - Fleet Manager - v1
 *
 * Main entry point for the system, responsible for all sorts of management operations for the whole service of managed service registry.
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kasclient

import (
	"encoding/json"
	"time"
)

// RegistryDeploymentStatus struct for RegistryDeploymentStatus
type RegistryDeploymentStatus struct {
	// ISO 8601 UTC timestamp.
	LastUpdated time.Time `json:"lastUpdated"`
	Value RegistryDeploymentStatusValue `json:"value"`
}

// NewRegistryDeploymentStatus instantiates a new RegistryDeploymentStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistryDeploymentStatus(lastUpdated time.Time, value RegistryDeploymentStatusValue) *RegistryDeploymentStatus {
	this := RegistryDeploymentStatus{}
	this.LastUpdated = lastUpdated
	this.Value = value
	return &this
}

// NewRegistryDeploymentStatusWithDefaults instantiates a new RegistryDeploymentStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistryDeploymentStatusWithDefaults() *RegistryDeploymentStatus {
	this := RegistryDeploymentStatus{}
	return &this
}

// GetLastUpdated returns the LastUpdated field value
func (o *RegistryDeploymentStatus) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *RegistryDeploymentStatus) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *RegistryDeploymentStatus) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetValue returns the Value field value
func (o *RegistryDeploymentStatus) GetValue() RegistryDeploymentStatusValue {
	if o == nil {
		var ret RegistryDeploymentStatusValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *RegistryDeploymentStatus) GetValueOk() (*RegistryDeploymentStatusValue, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *RegistryDeploymentStatus) SetValue(v RegistryDeploymentStatusValue) {
	o.Value = v
}

func (o RegistryDeploymentStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableRegistryDeploymentStatus struct {
	value *RegistryDeploymentStatus
	isSet bool
}

func (v NullableRegistryDeploymentStatus) Get() *RegistryDeploymentStatus {
	return v.value
}

func (v *NullableRegistryDeploymentStatus) Set(val *RegistryDeploymentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistryDeploymentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistryDeploymentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistryDeploymentStatus(val *RegistryDeploymentStatus) *NullableRegistryDeploymentStatus {
	return &NullableRegistryDeploymentStatus{value: val, isSet: true}
}

func (v NullableRegistryDeploymentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistryDeploymentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

