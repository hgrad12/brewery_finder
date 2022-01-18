/*
 * Twilio - Sync
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

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateSyncStream'
type CreateSyncStreamParams struct {
	// How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Stream expires and is deleted (time-to-live).
	Ttl *int `json:"Ttl,omitempty"`
	// An application-defined string that uniquely identifies the resource. This value must be unique within its Service and it can be up to 320 characters long. The `unique_name` value can be used as an alternative to the `sid` in the URL path to address the resource.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *CreateSyncStreamParams) SetTtl(Ttl int) *CreateSyncStreamParams {
	params.Ttl = &Ttl
	return params
}
func (params *CreateSyncStreamParams) SetUniqueName(UniqueName string) *CreateSyncStreamParams {
	params.UniqueName = &UniqueName
	return params
}

// Create a new Stream.
func (c *ApiService) CreateSyncStream(ServiceSid string, params *CreateSyncStreamParams) (*SyncV1SyncStream, error) {
	path := "/v1/Services/{ServiceSid}/Streams"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Ttl != nil {
		data.Set("Ttl", fmt.Sprint(*params.Ttl))
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1SyncStream{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific Stream.
func (c *ApiService) DeleteSyncStream(ServiceSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/Streams/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch a specific Stream.
func (c *ApiService) FetchSyncStream(ServiceSid string, Sid string) (*SyncV1SyncStream, error) {
	path := "/v1/Services/{ServiceSid}/Streams/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1SyncStream{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSyncStream'
type ListSyncStreamParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSyncStreamParams) SetPageSize(PageSize int) *ListSyncStreamParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSyncStreamParams) SetLimit(Limit int) *ListSyncStreamParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of SyncStream records from the API. Request is executed immediately.
func (c *ApiService) PageSyncStream(ServiceSid string, params *ListSyncStreamParams, pageToken, pageNumber string) (*ListSyncStreamResponse, error) {
	path := "/v1/Services/{ServiceSid}/Streams"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSyncStreamResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists SyncStream records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSyncStream(ServiceSid string, params *ListSyncStreamParams) ([]SyncV1SyncStream, error) {
	if params == nil {
		params = &ListSyncStreamParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSyncStream(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []SyncV1SyncStream

	for response != nil {
		records = append(records, response.Streams...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListSyncStreamResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListSyncStreamResponse)
	}

	return records, err
}

// Streams SyncStream records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSyncStream(ServiceSid string, params *ListSyncStreamParams) (chan SyncV1SyncStream, error) {
	if params == nil {
		params = &ListSyncStreamParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSyncStream(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan SyncV1SyncStream, 1)

	go func() {
		for response != nil {
			for item := range response.Streams {
				channel <- response.Streams[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListSyncStreamResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListSyncStreamResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListSyncStreamResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSyncStreamResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateSyncStream'
type UpdateSyncStreamParams struct {
	// How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Stream expires and is deleted (time-to-live).
	Ttl *int `json:"Ttl,omitempty"`
}

func (params *UpdateSyncStreamParams) SetTtl(Ttl int) *UpdateSyncStreamParams {
	params.Ttl = &Ttl
	return params
}

// Update a specific Stream.
func (c *ApiService) UpdateSyncStream(ServiceSid string, Sid string, params *UpdateSyncStreamParams) (*SyncV1SyncStream, error) {
	path := "/v1/Services/{ServiceSid}/Streams/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Ttl != nil {
		data.Set("Ttl", fmt.Sprint(*params.Ttl))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1SyncStream{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
