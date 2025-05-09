/*
REST API

REST API

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the ArticleAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArticleAttributes{}

// ArticleAttributes struct for ArticleAttributes
type ArticleAttributes struct {
	// Article title
	Title string `json:"title"`
	// Article status
	Status string `json:"status"`
	// Article link
	Icon *string `json:"icon,omitempty"`
	// Article description
	Desc *string `json:"desc,omitempty"`
	Content []Section `json:"content,omitempty"`
	// Published at
	PublishedAt *time.Time `json:"published_at,omitempty"`
	// Updated at
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Created at
	CreatedAt time.Time `json:"created_at"`
}

type _ArticleAttributes ArticleAttributes

// NewArticleAttributes instantiates a new ArticleAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArticleAttributes(title string, status string, createdAt time.Time) *ArticleAttributes {
	this := ArticleAttributes{}
	this.Title = title
	this.Status = status
	this.CreatedAt = createdAt
	return &this
}

// NewArticleAttributesWithDefaults instantiates a new ArticleAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArticleAttributesWithDefaults() *ArticleAttributes {
	this := ArticleAttributes{}
	return &this
}

// GetTitle returns the Title field value
func (o *ArticleAttributes) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ArticleAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ArticleAttributes) SetTitle(v string) {
	o.Title = v
}

// GetStatus returns the Status field value
func (o *ArticleAttributes) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ArticleAttributes) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ArticleAttributes) SetStatus(v string) {
	o.Status = v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *ArticleAttributes) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArticleAttributes) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *ArticleAttributes) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *ArticleAttributes) SetIcon(v string) {
	o.Icon = &v
}

// GetDesc returns the Desc field value if set, zero value otherwise.
func (o *ArticleAttributes) GetDesc() string {
	if o == nil || IsNil(o.Desc) {
		var ret string
		return ret
	}
	return *o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArticleAttributes) GetDescOk() (*string, bool) {
	if o == nil || IsNil(o.Desc) {
		return nil, false
	}
	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *ArticleAttributes) HasDesc() bool {
	if o != nil && !IsNil(o.Desc) {
		return true
	}

	return false
}

// SetDesc gets a reference to the given string and assigns it to the Desc field.
func (o *ArticleAttributes) SetDesc(v string) {
	o.Desc = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *ArticleAttributes) GetContent() []Section {
	if o == nil || IsNil(o.Content) {
		var ret []Section
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArticleAttributes) GetContentOk() ([]Section, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *ArticleAttributes) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given []Section and assigns it to the Content field.
func (o *ArticleAttributes) SetContent(v []Section) {
	o.Content = v
}

// GetPublishedAt returns the PublishedAt field value if set, zero value otherwise.
func (o *ArticleAttributes) GetPublishedAt() time.Time {
	if o == nil || IsNil(o.PublishedAt) {
		var ret time.Time
		return ret
	}
	return *o.PublishedAt
}

// GetPublishedAtOk returns a tuple with the PublishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArticleAttributes) GetPublishedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PublishedAt) {
		return nil, false
	}
	return o.PublishedAt, true
}

// HasPublishedAt returns a boolean if a field has been set.
func (o *ArticleAttributes) HasPublishedAt() bool {
	if o != nil && !IsNil(o.PublishedAt) {
		return true
	}

	return false
}

// SetPublishedAt gets a reference to the given time.Time and assigns it to the PublishedAt field.
func (o *ArticleAttributes) SetPublishedAt(v time.Time) {
	o.PublishedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ArticleAttributes) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArticleAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ArticleAttributes) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ArticleAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ArticleAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ArticleAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ArticleAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o ArticleAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArticleAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	toSerialize["status"] = o.Status
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !IsNil(o.Desc) {
		toSerialize["desc"] = o.Desc
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.PublishedAt) {
		toSerialize["published_at"] = o.PublishedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

func (o *ArticleAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
		"status",
		"created_at",
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

	varArticleAttributes := _ArticleAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varArticleAttributes)

	if err != nil {
		return err
	}

	*o = ArticleAttributes(varArticleAttributes)

	return err
}

type NullableArticleAttributes struct {
	value *ArticleAttributes
	isSet bool
}

func (v NullableArticleAttributes) Get() *ArticleAttributes {
	return v.value
}

func (v *NullableArticleAttributes) Set(val *ArticleAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableArticleAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableArticleAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArticleAttributes(val *ArticleAttributes) *NullableArticleAttributes {
	return &NullableArticleAttributes{value: val, isSet: true}
}

func (v NullableArticleAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArticleAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


