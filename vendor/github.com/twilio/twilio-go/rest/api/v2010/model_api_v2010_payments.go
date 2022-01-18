/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.25.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010Payments struct for ApiV2010Payments
type ApiV2010Payments struct {
	// The SID of the Account that created the Payments resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Call the resource is associated with.
	CallSid *string `json:"call_sid,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated *string `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`
	// The SID of the Payments resource.
	Sid *string `json:"sid,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`
	Uri *string `json:"uri,omitempty"`
}
