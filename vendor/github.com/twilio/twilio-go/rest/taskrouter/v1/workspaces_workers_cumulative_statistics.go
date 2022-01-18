/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.25.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"
	"time"
)

// Optional parameters for the method 'FetchWorkersCumulativeStatistics'
type FetchWorkersCumulativeStatisticsParams struct {
	// Only calculate statistics from this date and time and earlier, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EndDate *time.Time `json:"EndDate,omitempty"`
	// Only calculate statistics since this many minutes in the past. The default 15 minutes. This is helpful for displaying statistics for the last 15 minutes, 240 minutes (4 hours), and 480 minutes (8 hours) to see trends.
	Minutes *int `json:"Minutes,omitempty"`
	// Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	StartDate *time.Time `json:"StartDate,omitempty"`
	// Only calculate cumulative statistics on this TaskChannel. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`.
	TaskChannel *string `json:"TaskChannel,omitempty"`
}

func (params *FetchWorkersCumulativeStatisticsParams) SetEndDate(EndDate time.Time) *FetchWorkersCumulativeStatisticsParams {
	params.EndDate = &EndDate
	return params
}
func (params *FetchWorkersCumulativeStatisticsParams) SetMinutes(Minutes int) *FetchWorkersCumulativeStatisticsParams {
	params.Minutes = &Minutes
	return params
}
func (params *FetchWorkersCumulativeStatisticsParams) SetStartDate(StartDate time.Time) *FetchWorkersCumulativeStatisticsParams {
	params.StartDate = &StartDate
	return params
}
func (params *FetchWorkersCumulativeStatisticsParams) SetTaskChannel(TaskChannel string) *FetchWorkersCumulativeStatisticsParams {
	params.TaskChannel = &TaskChannel
	return params
}

func (c *ApiService) FetchWorkersCumulativeStatistics(WorkspaceSid string, params *FetchWorkersCumulativeStatisticsParams) (*TaskrouterV1WorkersCumulativeStatistics, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Workers/CumulativeStatistics"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.EndDate != nil {
		data.Set("EndDate", fmt.Sprint((*params.EndDate).Format(time.RFC3339)))
	}
	if params != nil && params.Minutes != nil {
		data.Set("Minutes", fmt.Sprint(*params.Minutes))
	}
	if params != nil && params.StartDate != nil {
		data.Set("StartDate", fmt.Sprint((*params.StartDate).Format(time.RFC3339)))
	}
	if params != nil && params.TaskChannel != nil {
		data.Set("TaskChannel", *params.TaskChannel)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1WorkersCumulativeStatistics{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}