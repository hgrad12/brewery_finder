/*
 * Twilio - Wireless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.25.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListAccountUsageRecordResponse struct for ListAccountUsageRecordResponse
type ListAccountUsageRecordResponse struct {
	Meta         ListCommandResponseMeta        `json:"meta,omitempty"`
	UsageRecords []WirelessV1AccountUsageRecord `json:"usage_records,omitempty"`
}
