/*
 * Twilio - Conversations
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

// ConversationsV1ServiceConversationMessage struct for ConversationsV1ServiceConversationMessage
type ConversationsV1ServiceConversationMessage struct {
	// The unique ID of the Account responsible for this message.
	AccountSid *string `json:"account_sid,omitempty"`
	// A string metadata field you can use to store any data you wish.
	Attributes *string `json:"attributes,omitempty"`
	// The channel specific identifier of the message's author.
	Author *string `json:"author,omitempty"`
	// The content of the message.
	Body *string `json:"body,omitempty"`
	// The SID of the Conversation Service that the resource is associated with.
	ChatServiceSid *string `json:"chat_service_sid,omitempty"`
	// The unique ID of the Conversation for this message.
	ConversationSid *string `json:"conversation_sid,omitempty"`
	// The date that this resource was created.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date that this resource was last updated.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// An object that contains the summary of delivery statuses for the message to non-chat participants.
	Delivery *map[string]interface{} `json:"delivery,omitempty"`
	// The index of the message within the Conversation.
	Index *int `json:"index,omitempty"`
	// Absolute URL to access the receipts of this message.
	Links *map[string]interface{} `json:"links,omitempty"`
	// An array of objects that describe the Message's media if attached, otherwise, null.
	Media *[]map[string]interface{} `json:"media,omitempty"`
	// The unique ID of messages's author participant.
	ParticipantSid *string `json:"participant_sid,omitempty"`
	// A 34 character string that uniquely identifies this resource.
	Sid *string `json:"sid,omitempty"`
	// An absolute URL for this message.
	Url *string `json:"url,omitempty"`
}