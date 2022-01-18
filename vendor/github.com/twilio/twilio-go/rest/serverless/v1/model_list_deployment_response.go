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

// ListDeploymentResponse struct for ListDeploymentResponse
type ListDeploymentResponse struct {
	Deployments []ServerlessV1Deployment `json:"deployments,omitempty"`
	Meta        ListServiceResponseMeta  `json:"meta,omitempty"`
}
