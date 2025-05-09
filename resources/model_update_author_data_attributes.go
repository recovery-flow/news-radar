/*
REST API

REST API

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
)

// checks if the UpdateAuthorDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAuthorDataAttributes{}

// UpdateAuthorDataAttributes struct for UpdateAuthorDataAttributes
type UpdateAuthorDataAttributes struct {
	Name *string `json:"name,omitempty"`
	Status *string `json:"status,omitempty"`
	Desc *string `json:"desc,omitempty"`
	Avatar *string `json:"avatar,omitempty"`
	Email *string `json:"email,omitempty"`
	Telegram *string `json:"telegram,omitempty"`
	Twitter *string `json:"twitter,omitempty"`
}

// NewUpdateAuthorDataAttributes instantiates a new UpdateAuthorDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAuthorDataAttributes() *UpdateAuthorDataAttributes {
	this := UpdateAuthorDataAttributes{}
	return &this
}

// NewUpdateAuthorDataAttributesWithDefaults instantiates a new UpdateAuthorDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAuthorDataAttributesWithDefaults() *UpdateAuthorDataAttributes {
	this := UpdateAuthorDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateAuthorDataAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthorDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateAuthorDataAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateAuthorDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UpdateAuthorDataAttributes) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthorDataAttributes) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UpdateAuthorDataAttributes) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UpdateAuthorDataAttributes) SetStatus(v string) {
	o.Status = &v
}

// GetDesc returns the Desc field value if set, zero value otherwise.
func (o *UpdateAuthorDataAttributes) GetDesc() string {
	if o == nil || IsNil(o.Desc) {
		var ret string
		return ret
	}
	return *o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthorDataAttributes) GetDescOk() (*string, bool) {
	if o == nil || IsNil(o.Desc) {
		return nil, false
	}
	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *UpdateAuthorDataAttributes) HasDesc() bool {
	if o != nil && !IsNil(o.Desc) {
		return true
	}

	return false
}

// SetDesc gets a reference to the given string and assigns it to the Desc field.
func (o *UpdateAuthorDataAttributes) SetDesc(v string) {
	o.Desc = &v
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *UpdateAuthorDataAttributes) GetAvatar() string {
	if o == nil || IsNil(o.Avatar) {
		var ret string
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthorDataAttributes) GetAvatarOk() (*string, bool) {
	if o == nil || IsNil(o.Avatar) {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *UpdateAuthorDataAttributes) HasAvatar() bool {
	if o != nil && !IsNil(o.Avatar) {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given string and assigns it to the Avatar field.
func (o *UpdateAuthorDataAttributes) SetAvatar(v string) {
	o.Avatar = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UpdateAuthorDataAttributes) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthorDataAttributes) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UpdateAuthorDataAttributes) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UpdateAuthorDataAttributes) SetEmail(v string) {
	o.Email = &v
}

// GetTelegram returns the Telegram field value if set, zero value otherwise.
func (o *UpdateAuthorDataAttributes) GetTelegram() string {
	if o == nil || IsNil(o.Telegram) {
		var ret string
		return ret
	}
	return *o.Telegram
}

// GetTelegramOk returns a tuple with the Telegram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthorDataAttributes) GetTelegramOk() (*string, bool) {
	if o == nil || IsNil(o.Telegram) {
		return nil, false
	}
	return o.Telegram, true
}

// HasTelegram returns a boolean if a field has been set.
func (o *UpdateAuthorDataAttributes) HasTelegram() bool {
	if o != nil && !IsNil(o.Telegram) {
		return true
	}

	return false
}

// SetTelegram gets a reference to the given string and assigns it to the Telegram field.
func (o *UpdateAuthorDataAttributes) SetTelegram(v string) {
	o.Telegram = &v
}

// GetTwitter returns the Twitter field value if set, zero value otherwise.
func (o *UpdateAuthorDataAttributes) GetTwitter() string {
	if o == nil || IsNil(o.Twitter) {
		var ret string
		return ret
	}
	return *o.Twitter
}

// GetTwitterOk returns a tuple with the Twitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAuthorDataAttributes) GetTwitterOk() (*string, bool) {
	if o == nil || IsNil(o.Twitter) {
		return nil, false
	}
	return o.Twitter, true
}

// HasTwitter returns a boolean if a field has been set.
func (o *UpdateAuthorDataAttributes) HasTwitter() bool {
	if o != nil && !IsNil(o.Twitter) {
		return true
	}

	return false
}

// SetTwitter gets a reference to the given string and assigns it to the Twitter field.
func (o *UpdateAuthorDataAttributes) SetTwitter(v string) {
	o.Twitter = &v
}

func (o UpdateAuthorDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAuthorDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Desc) {
		toSerialize["desc"] = o.Desc
	}
	if !IsNil(o.Avatar) {
		toSerialize["avatar"] = o.Avatar
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Telegram) {
		toSerialize["telegram"] = o.Telegram
	}
	if !IsNil(o.Twitter) {
		toSerialize["twitter"] = o.Twitter
	}
	return toSerialize, nil
}

type NullableUpdateAuthorDataAttributes struct {
	value *UpdateAuthorDataAttributes
	isSet bool
}

func (v NullableUpdateAuthorDataAttributes) Get() *UpdateAuthorDataAttributes {
	return v.value
}

func (v *NullableUpdateAuthorDataAttributes) Set(val *UpdateAuthorDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAuthorDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAuthorDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAuthorDataAttributes(val *UpdateAuthorDataAttributes) *NullableUpdateAuthorDataAttributes {
	return &NullableUpdateAuthorDataAttributes{value: val, isSet: true}
}

func (v NullableUpdateAuthorDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAuthorDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


