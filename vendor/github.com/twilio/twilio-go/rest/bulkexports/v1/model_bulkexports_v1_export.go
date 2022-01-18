/*
 * Twilio - Bulkexports
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.25.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// BulkexportsV1Export struct for BulkexportsV1Export
type BulkexportsV1Export struct {
	// Nested resource URLs.
	Links *map[string]interface{} `json:"links,omitempty"`
	// The type of communication – Messages, Calls, Conferences, and Participants
	ResourceType *string `json:"resource_type,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
}