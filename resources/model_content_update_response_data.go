/*
REST API

REST API

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ContentUpdateResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentUpdateResponseData{}

// ContentUpdateResponseData struct for ContentUpdateResponseData
type ContentUpdateResponseData struct {
	Type string `json:"type"`
	Attributes ContentUpdateResponseDataAttributes `json:"attributes"`
}

type _ContentUpdateResponseData ContentUpdateResponseData

// NewContentUpdateResponseData instantiates a new ContentUpdateResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentUpdateResponseData(type_ string, attributes ContentUpdateResponseDataAttributes) *ContentUpdateResponseData {
	this := ContentUpdateResponseData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewContentUpdateResponseDataWithDefaults instantiates a new ContentUpdateResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentUpdateResponseDataWithDefaults() *ContentUpdateResponseData {
	this := ContentUpdateResponseData{}
	return &this
}

// GetType returns the Type field value
func (o *ContentUpdateResponseData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ContentUpdateResponseData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ContentUpdateResponseData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ContentUpdateResponseData) GetAttributes() ContentUpdateResponseDataAttributes {
	if o == nil {
		var ret ContentUpdateResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ContentUpdateResponseData) GetAttributesOk() (*ContentUpdateResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ContentUpdateResponseData) SetAttributes(v ContentUpdateResponseDataAttributes) {
	o.Attributes = v
}

func (o ContentUpdateResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentUpdateResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *ContentUpdateResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"attributes",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varContentUpdateResponseData := _ContentUpdateResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContentUpdateResponseData)

	if err != nil {
		return err
	}

	*o = ContentUpdateResponseData(varContentUpdateResponseData)

	return err
}

type NullableContentUpdateResponseData struct {
	value *ContentUpdateResponseData
	isSet bool
}

func (v NullableContentUpdateResponseData) Get() *ContentUpdateResponseData {
	return v.value
}

func (v *NullableContentUpdateResponseData) Set(val *ContentUpdateResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableContentUpdateResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableContentUpdateResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentUpdateResponseData(val *ContentUpdateResponseData) *NullableContentUpdateResponseData {
	return &NullableContentUpdateResponseData{value: val, isSet: true}
}

func (v NullableContentUpdateResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentUpdateResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


