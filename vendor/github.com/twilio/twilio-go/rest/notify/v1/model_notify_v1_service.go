/*
 * Twilio - Notify
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

// NotifyV1Service struct for NotifyV1Service
type NotifyV1Service struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// Deprecated
	AlexaSkillId *string `json:"alexa_skill_id,omitempty"`
	// The SID of the Credential to use for APN Bindings
	ApnCredentialSid *string `json:"apn_credential_sid,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// Deprecated
	DefaultAlexaNotificationProtocolVersion *string `json:"default_alexa_notification_protocol_version,omitempty"`
	// The protocol version to use for sending APNS notifications
	DefaultApnNotificationProtocolVersion *string `json:"default_apn_notification_protocol_version,omitempty"`
	// The protocol version to use for sending FCM notifications
	DefaultFcmNotificationProtocolVersion *string `json:"default_fcm_notification_protocol_version,omitempty"`
	// The protocol version to use for sending GCM notifications
	DefaultGcmNotificationProtocolVersion *string `json:"default_gcm_notification_protocol_version,omitempty"`
	// Enable delivery callbacks
	DeliveryCallbackEnabled *bool `json:"delivery_callback_enabled,omitempty"`
	// Webhook URL
	DeliveryCallbackUrl *string `json:"delivery_callback_url,omitempty"`
	// Deprecated
	FacebookMessengerPageId *string `json:"facebook_messenger_page_id,omitempty"`
	// The SID of the Credential to use for FCM Bindings
	FcmCredentialSid *string `json:"fcm_credential_sid,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The SID of the Credential to use for GCM Bindings
	GcmCredentialSid *string `json:"gcm_credential_sid,omitempty"`
	// The URLs of the resources related to the service
	Links *map[string]interface{} `json:"links,omitempty"`
	// Whether to log notifications
	LogEnabled *bool `json:"log_enabled,omitempty"`
	// The SID of the Messaging Service to use for SMS Bindings
	MessagingServiceSid *string `json:"messaging_service_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the Service resource
	Url *string `json:"url,omitempty"`
}
