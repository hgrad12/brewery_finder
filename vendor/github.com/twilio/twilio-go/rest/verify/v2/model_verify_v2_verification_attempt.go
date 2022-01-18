/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.25.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// VerifyV2VerificationAttempt struct for VerifyV2VerificationAttempt
type VerifyV2VerificationAttempt struct {
	// Account Sid
	AccountSid *string `json:"account_sid,omitempty"`
	// Channel used for the attempt
	Channel *string `json:"channel,omitempty"`
	// Object with the channel information for an attempt
	ChannelData *map[string]interface{} `json:"channel_data,omitempty"`
	// Status of a conversion
	ConversionStatus *string `json:"conversion_status,omitempty"`
	// The date this Attempt was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date this Attempt was updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	ServiceSid  *string    `json:"service_sid,omitempty"`
	// A string that uniquely identifies this Verification Attempt
	Sid *string `json:"sid,omitempty"`
	Url *string `json:"url,omitempty"`
}
