/*
 * Ory Identities API
 *
 * This is the API specification for Ory Identities with features such as registration, login, recovery, account verification, profile settings, password reset, identity management, session management, email and sms delivery, and more.
 *
 * API version:
 * Contact: office@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// UiNodeAttributes - struct for UiNodeAttributes
type UiNodeAttributes struct {
	UiNodeAnchorAttributes *UiNodeAnchorAttributes
	UiNodeImageAttributes  *UiNodeImageAttributes
	UiNodeInputAttributes  *UiNodeInputAttributes
	UiNodeScriptAttributes *UiNodeScriptAttributes
	UiNodeTextAttributes   *UiNodeTextAttributes
}

// UiNodeAnchorAttributesAsUiNodeAttributes is a convenience function that returns UiNodeAnchorAttributes wrapped in UiNodeAttributes
func UiNodeAnchorAttributesAsUiNodeAttributes(v *UiNodeAnchorAttributes) UiNodeAttributes {
	return UiNodeAttributes{
		UiNodeAnchorAttributes: v,
	}
}

// UiNodeImageAttributesAsUiNodeAttributes is a convenience function that returns UiNodeImageAttributes wrapped in UiNodeAttributes
func UiNodeImageAttributesAsUiNodeAttributes(v *UiNodeImageAttributes) UiNodeAttributes {
	return UiNodeAttributes{
		UiNodeImageAttributes: v,
	}
}

// UiNodeInputAttributesAsUiNodeAttributes is a convenience function that returns UiNodeInputAttributes wrapped in UiNodeAttributes
func UiNodeInputAttributesAsUiNodeAttributes(v *UiNodeInputAttributes) UiNodeAttributes {
	return UiNodeAttributes{
		UiNodeInputAttributes: v,
	}
}

// UiNodeScriptAttributesAsUiNodeAttributes is a convenience function that returns UiNodeScriptAttributes wrapped in UiNodeAttributes
func UiNodeScriptAttributesAsUiNodeAttributes(v *UiNodeScriptAttributes) UiNodeAttributes {
	return UiNodeAttributes{
		UiNodeScriptAttributes: v,
	}
}

// UiNodeTextAttributesAsUiNodeAttributes is a convenience function that returns UiNodeTextAttributes wrapped in UiNodeAttributes
func UiNodeTextAttributesAsUiNodeAttributes(v *UiNodeTextAttributes) UiNodeAttributes {
	return UiNodeAttributes{
		UiNodeTextAttributes: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UiNodeAttributes) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'a'
	if jsonDict["node_type"] == "a" {
		// try to unmarshal JSON data into UiNodeAnchorAttributes
		err = json.Unmarshal(data, &dst.UiNodeAnchorAttributes)
		if err == nil {
			return nil // data stored in dst.UiNodeAnchorAttributes, return on the first match
		} else {
			dst.UiNodeAnchorAttributes = nil
			return fmt.Errorf("Failed to unmarshal UiNodeAttributes as UiNodeAnchorAttributes: %s", err.Error())
		}
	}

	// check if the discriminator value is 'img'
	if jsonDict["node_type"] == "img" {
		// try to unmarshal JSON data into UiNodeImageAttributes
		err = json.Unmarshal(data, &dst.UiNodeImageAttributes)
		if err == nil {
			return nil // data stored in dst.UiNodeImageAttributes, return on the first match
		} else {
			dst.UiNodeImageAttributes = nil
			return fmt.Errorf("Failed to unmarshal UiNodeAttributes as UiNodeImageAttributes: %s", err.Error())
		}
	}

	// check if the discriminator value is 'input'
	if jsonDict["node_type"] == "input" {
		// try to unmarshal JSON data into UiNodeInputAttributes
		err = json.Unmarshal(data, &dst.UiNodeInputAttributes)
		if err == nil {
			return nil // data stored in dst.UiNodeInputAttributes, return on the first match
		} else {
			dst.UiNodeInputAttributes = nil
			return fmt.Errorf("Failed to unmarshal UiNodeAttributes as UiNodeInputAttributes: %s", err.Error())
		}
	}

	// check if the discriminator value is 'script'
	if jsonDict["node_type"] == "script" {
		// try to unmarshal JSON data into UiNodeScriptAttributes
		err = json.Unmarshal(data, &dst.UiNodeScriptAttributes)
		if err == nil {
			return nil // data stored in dst.UiNodeScriptAttributes, return on the first match
		} else {
			dst.UiNodeScriptAttributes = nil
			return fmt.Errorf("Failed to unmarshal UiNodeAttributes as UiNodeScriptAttributes: %s", err.Error())
		}
	}

	// check if the discriminator value is 'text'
	if jsonDict["node_type"] == "text" {
		// try to unmarshal JSON data into UiNodeTextAttributes
		err = json.Unmarshal(data, &dst.UiNodeTextAttributes)
		if err == nil {
			return nil // data stored in dst.UiNodeTextAttributes, return on the first match
		} else {
			dst.UiNodeTextAttributes = nil
			return fmt.Errorf("Failed to unmarshal UiNodeAttributes as UiNodeTextAttributes: %s", err.Error())
		}
	}

	// check if the discriminator value is 'uiNodeAnchorAttributes'
	if jsonDict["node_type"] == "uiNodeAnchorAttributes" {
		// try to unmarshal JSON data into UiNodeAnchorAttributes
		err = json.Unmarshal(data, &dst.UiNodeAnchorAttributes)
		if err == nil {
			return nil // data stored in dst.UiNodeAnchorAttributes, return on the first match
		} else {
			dst.UiNodeAnchorAttributes = nil
			return fmt.Errorf("Failed to unmarshal UiNodeAttributes as UiNodeAnchorAttributes: %s", err.Error())
		}
	}

	// check if the discriminator value is 'uiNodeImageAttributes'
	if jsonDict["node_type"] == "uiNodeImageAttributes" {
		// try to unmarshal JSON data into UiNodeImageAttributes
		err = json.Unmarshal(data, &dst.UiNodeImageAttributes)
		if err == nil {
			return nil // data stored in dst.UiNodeImageAttributes, return on the first match
		} else {
			dst.UiNodeImageAttributes = nil
			return fmt.Errorf("Failed to unmarshal UiNodeAttributes as UiNodeImageAttributes: %s", err.Error())
		}
	}

	// check if the discriminator value is 'uiNodeInputAttributes'
	if jsonDict["node_type"] == "uiNodeInputAttributes" {
		// try to unmarshal JSON data into UiNodeInputAttributes
		err = json.Unmarshal(data, &dst.UiNodeInputAttributes)
		if err == nil {
			return nil // data stored in dst.UiNodeInputAttributes, return on the first match
		} else {
			dst.UiNodeInputAttributes = nil
			return fmt.Errorf("Failed to unmarshal UiNodeAttributes as UiNodeInputAttributes: %s", err.Error())
		}
	}

	// check if the discriminator value is 'uiNodeScriptAttributes'
	if jsonDict["node_type"] == "uiNodeScriptAttributes" {
		// try to unmarshal JSON data into UiNodeScriptAttributes
		err = json.Unmarshal(data, &dst.UiNodeScriptAttributes)
		if err == nil {
			return nil // data stored in dst.UiNodeScriptAttributes, return on the first match
		} else {
			dst.UiNodeScriptAttributes = nil
			return fmt.Errorf("Failed to unmarshal UiNodeAttributes as UiNodeScriptAttributes: %s", err.Error())
		}
	}

	// check if the discriminator value is 'uiNodeTextAttributes'
	if jsonDict["node_type"] == "uiNodeTextAttributes" {
		// try to unmarshal JSON data into UiNodeTextAttributes
		err = json.Unmarshal(data, &dst.UiNodeTextAttributes)
		if err == nil {
			return nil // data stored in dst.UiNodeTextAttributes, return on the first match
		} else {
			dst.UiNodeTextAttributes = nil
			return fmt.Errorf("Failed to unmarshal UiNodeAttributes as UiNodeTextAttributes: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UiNodeAttributes) MarshalJSON() ([]byte, error) {
	if src.UiNodeAnchorAttributes != nil {
		return json.Marshal(&src.UiNodeAnchorAttributes)
	}

	if src.UiNodeImageAttributes != nil {
		return json.Marshal(&src.UiNodeImageAttributes)
	}

	if src.UiNodeInputAttributes != nil {
		return json.Marshal(&src.UiNodeInputAttributes)
	}

	if src.UiNodeScriptAttributes != nil {
		return json.Marshal(&src.UiNodeScriptAttributes)
	}

	if src.UiNodeTextAttributes != nil {
		return json.Marshal(&src.UiNodeTextAttributes)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UiNodeAttributes) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.UiNodeAnchorAttributes != nil {
		return obj.UiNodeAnchorAttributes
	}

	if obj.UiNodeImageAttributes != nil {
		return obj.UiNodeImageAttributes
	}

	if obj.UiNodeInputAttributes != nil {
		return obj.UiNodeInputAttributes
	}

	if obj.UiNodeScriptAttributes != nil {
		return obj.UiNodeScriptAttributes
	}

	if obj.UiNodeTextAttributes != nil {
		return obj.UiNodeTextAttributes
	}

	// all schemas are nil
	return nil
}

type NullableUiNodeAttributes struct {
	value *UiNodeAttributes
	isSet bool
}

func (v NullableUiNodeAttributes) Get() *UiNodeAttributes {
	return v.value
}

func (v *NullableUiNodeAttributes) Set(val *UiNodeAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableUiNodeAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableUiNodeAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUiNodeAttributes(val *UiNodeAttributes) *NullableUiNodeAttributes {
	return &NullableUiNodeAttributes{value: val, isSet: true}
}

func (v NullableUiNodeAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUiNodeAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}