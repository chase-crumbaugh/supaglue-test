// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type EmailAddressesEmailAddressTypeEnum string

const (
	EmailAddressesEmailAddressTypeEnumPrimary EmailAddressesEmailAddressTypeEnum = "primary"
	EmailAddressesEmailAddressTypeEnumWork    EmailAddressesEmailAddressTypeEnum = "work"
)

func (e *EmailAddressesEmailAddressTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "primary":
		fallthrough
	case "work":
		*e = EmailAddressesEmailAddressTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for EmailAddressesEmailAddressTypeEnum: %s", s)
	}
}

type EmailAddresses struct {
	EmailAddress     string                             `json:"email_address"`
	EmailAddressType EmailAddressesEmailAddressTypeEnum `json:"email_address_type"`
}
