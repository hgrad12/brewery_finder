/*
 * Twilio - Media
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.25.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListMediaProcessorResponse struct for ListMediaProcessorResponse
type ListMediaProcessorResponse struct {
	MediaProcessors []MediaV1MediaProcessor        `json:"media_processors,omitempty"`
	Meta            ListMediaProcessorResponseMeta `json:"meta,omitempty"`
}
