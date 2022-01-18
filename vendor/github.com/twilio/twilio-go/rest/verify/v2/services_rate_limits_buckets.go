/*
 * Twilio - Verify
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

// Optional parameters for the method 'CreateBucket'
type CreateBucketParams struct {
	// Number of seconds that the rate limit will be enforced over.
	Interval *int `json:"Interval,omitempty"`
	// Maximum number of requests permitted in during the interval.
	Max *int `json:"Max,omitempty"`
}

func (params *CreateBucketParams) SetInterval(Interval int) *CreateBucketParams {
	params.Interval = &Interval
	return params
}
func (params *CreateBucketParams) SetMax(Max int) *CreateBucketParams {
	params.Max = &Max
	return params
}

// Create a new Bucket for a Rate Limit
func (c *ApiService) CreateBucket(ServiceSid string, RateLimitSid string, params *CreateBucketParams) (*VerifyV2Bucket, error) {
	path := "/v2/Services/{ServiceSid}/RateLimits/{RateLimitSid}/Buckets"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"RateLimitSid"+"}", RateLimitSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Interval != nil {
		data.Set("Interval", fmt.Sprint(*params.Interval))
	}
	if params != nil && params.Max != nil {
		data.Set("Max", fmt.Sprint(*params.Max))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VerifyV2Bucket{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific Bucket.
func (c *ApiService) DeleteBucket(ServiceSid string, RateLimitSid string, Sid string) error {
	path := "/v2/Services/{ServiceSid}/RateLimits/{RateLimitSid}/Buckets/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"RateLimitSid"+"}", RateLimitSid, -1)
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

// Fetch a specific Bucket.
func (c *ApiService) FetchBucket(ServiceSid string, RateLimitSid string, Sid string) (*VerifyV2Bucket, error) {
	path := "/v2/Services/{ServiceSid}/RateLimits/{RateLimitSid}/Buckets/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"RateLimitSid"+"}", RateLimitSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VerifyV2Bucket{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListBucket'
type ListBucketParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListBucketParams) SetPageSize(PageSize int) *ListBucketParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListBucketParams) SetLimit(Limit int) *ListBucketParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Bucket records from the API. Request is executed immediately.
func (c *ApiService) PageBucket(ServiceSid string, RateLimitSid string, params *ListBucketParams, pageToken, pageNumber string) (*ListBucketResponse, error) {
	path := "/v2/Services/{ServiceSid}/RateLimits/{RateLimitSid}/Buckets"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"RateLimitSid"+"}", RateLimitSid, -1)

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

	ps := &ListBucketResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Bucket records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListBucket(ServiceSid string, RateLimitSid string, params *ListBucketParams) ([]VerifyV2Bucket, error) {
	if params == nil {
		params = &ListBucketParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageBucket(ServiceSid, RateLimitSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []VerifyV2Bucket

	for response != nil {
		records = append(records, response.Buckets...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListBucketResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListBucketResponse)
	}

	return records, err
}

// Streams Bucket records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamBucket(ServiceSid string, RateLimitSid string, params *ListBucketParams) (chan VerifyV2Bucket, error) {
	if params == nil {
		params = &ListBucketParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageBucket(ServiceSid, RateLimitSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan VerifyV2Bucket, 1)

	go func() {
		for response != nil {
			for item := range response.Buckets {
				channel <- response.Buckets[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListBucketResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListBucketResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListBucketResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListBucketResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateBucket'
type UpdateBucketParams struct {
	// Number of seconds that the rate limit will be enforced over.
	Interval *int `json:"Interval,omitempty"`
	// Maximum number of requests permitted in during the interval.
	Max *int `json:"Max,omitempty"`
}

func (params *UpdateBucketParams) SetInterval(Interval int) *UpdateBucketParams {
	params.Interval = &Interval
	return params
}
func (params *UpdateBucketParams) SetMax(Max int) *UpdateBucketParams {
	params.Max = &Max
	return params
}

// Update a specific Bucket.
func (c *ApiService) UpdateBucket(ServiceSid string, RateLimitSid string, Sid string, params *UpdateBucketParams) (*VerifyV2Bucket, error) {
	path := "/v2/Services/{ServiceSid}/RateLimits/{RateLimitSid}/Buckets/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"RateLimitSid"+"}", RateLimitSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Interval != nil {
		data.Set("Interval", fmt.Sprint(*params.Interval))
	}
	if params != nil && params.Max != nil {
		data.Set("Max", fmt.Sprint(*params.Max))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VerifyV2Bucket{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
