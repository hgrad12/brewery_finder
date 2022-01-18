/*
 * Twilio - Serverless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.25.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListEnvironmentResponse struct for ListEnvironmentResponse
type ListEnvironmentResponse struct {
	Environments []ServerlessV1Environment `json:"environments,omitempty"`
	Meta         ListServiceResponseMeta   `json:"meta,omitempty"`
}
