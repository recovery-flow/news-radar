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

// checks if the HashtagCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HashtagCreateData{}

// HashtagCreateData struct for HashtagCreateData
type HashtagCreateData struct {
	Type string `json:"type"`
	Attributes HashtagCreateDataAttributes `json:"attributes"`
}

type _HashtagCreateData HashtagCreateData

// NewHashtagCreateData instantiates a new HashtagCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHashtagCreateData(type_ string, attributes HashtagCreateDataAttributes) *HashtagCreateData {
	this := HashtagCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewHashtagCreateDataWithDefaults instantiates a new HashtagCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHashtagCreateDataWithDefaults() *HashtagCreateData {
	this := HashtagCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *HashtagCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HashtagCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *HashtagCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *HashtagCreateData) GetAttributes() HashtagCreateDataAttributes {
	if o == nil {
		var ret HashtagCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *HashtagCreateData) GetAttributesOk() (*HashtagCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *HashtagCreateData) SetAttributes(v HashtagCreateDataAttributes) {
	o.Attributes = v
}

func (o HashtagCreateData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HashtagCreateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *HashtagCreateData) UnmarshalJSON(data []byte) (err error) {
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

	varHashtagCreateData := _HashtagCreateData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHashtagCreateData)

	if err != nil {
		return err
	}

	*o = HashtagCreateData(varHashtagCreateData)

	return err
}

type NullableHashtagCreateData struct {
	value *HashtagCreateData
	isSet bool
}

func (v NullableHashtagCreateData) Get() *HashtagCreateData {
	return v.value
}

func (v *NullableHashtagCreateData) Set(val *HashtagCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableHashtagCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableHashtagCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHashtagCreateData(val *HashtagCreateData) *NullableHashtagCreateData {
	return &NullableHashtagCreateData{value: val, isSet: true}
}

func (v NullableHashtagCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHashtagCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


