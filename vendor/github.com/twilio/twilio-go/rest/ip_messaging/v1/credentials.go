/*
 * Twilio - Ip_messaging
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

// Optional parameters for the method 'CreateCredential'
type CreateCredentialParams struct {
	//
	ApiKey *string `json:"ApiKey,omitempty"`
	//
	Certificate *string `json:"Certificate,omitempty"`
	//
	FriendlyName *string `json:"FriendlyName,omitempty"`
	//
	PrivateKey *string `json:"PrivateKey,omitempty"`
	//
	Sandbox *bool `json:"Sandbox,omitempty"`
	//
	Secret *string `json:"Secret,omitempty"`
	//
	Type *string `json:"Type,omitempty"`
}

func (params *CreateCredentialParams) SetApiKey(ApiKey string) *CreateCredentialParams {
	params.ApiKey = &ApiKey
	return params
}
func (params *CreateCredentialParams) SetCertificate(Certificate string) *CreateCredentialParams {
	params.Certificate = &Certificate
	return params
}
func (params *CreateCredentialParams) SetFriendlyName(FriendlyName string) *CreateCredentialParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateCredentialParams) SetPrivateKey(PrivateKey string) *CreateCredentialParams {
	params.PrivateKey = &PrivateKey
	return params
}
func (params *CreateCredentialParams) SetSandbox(Sandbox bool) *CreateCredentialParams {
	params.Sandbox = &Sandbox
	return params
}
func (params *CreateCredentialParams) SetSecret(Secret string) *CreateCredentialParams {
	params.Secret = &Secret
	return params
}
func (params *CreateCredentialParams) SetType(Type string) *CreateCredentialParams {
	params.Type = &Type
	return params
}

func (c *ApiService) CreateCredential(params *CreateCredentialParams) (*IpMessagingV1Credential, error) {
	path := "/v1/Credentials"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ApiKey != nil {
		data.Set("ApiKey", *params.ApiKey)
	}
	if params != nil && params.Certificate != nil {
		data.Set("Certificate", *params.Certificate)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.PrivateKey != nil {
		data.Set("PrivateKey", *params.PrivateKey)
	}
	if params != nil && params.Sandbox != nil {
		data.Set("Sandbox", fmt.Sprint(*params.Sandbox))
	}
	if params != nil && params.Secret != nil {
		data.Set("Secret", *params.Secret)
	}
	if params != nil && params.Type != nil {
		data.Set("Type", *params.Type)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV1Credential{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteCredential(Sid string) error {
	path := "/v1/Credentials/{Sid}"
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

func (c *ApiService) FetchCredential(Sid string) (*IpMessagingV1Credential, error) {
	path := "/v1/Credentials/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV1Credential{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListCredential'
type ListCredentialParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListCredentialParams) SetPageSize(PageSize int) *ListCredentialParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListCredentialParams) SetLimit(Limit int) *ListCredentialParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Credential records from the API. Request is executed immediately.
func (c *ApiService) PageCredential(params *ListCredentialParams, pageToken, pageNumber string) (*ListCredentialResponse, error) {
	path := "/v1/Credentials"

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

	ps := &ListCredentialResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Credential records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListCredential(params *ListCredentialParams) ([]IpMessagingV1Credential, error) {
	if params == nil {
		params = &ListCredentialParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageCredential(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []IpMessagingV1Credential

	for response != nil {
		records = append(records, response.Credentials...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListCredentialResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListCredentialResponse)
	}

	return records, err
}

// Streams Credential records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamCredential(params *ListCredentialParams) (chan IpMessagingV1Credential, error) {
	if params == nil {
		params = &ListCredentialParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageCredential(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan IpMessagingV1Credential, 1)

	go func() {
		for response != nil {
			for item := range response.Credentials {
				channel <- response.Credentials[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListCredentialResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListCredentialResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListCredentialResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCredentialResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateCredential'
type UpdateCredentialParams struct {
	//
	ApiKey *string `json:"ApiKey,omitempty"`
	//
	Certificate *string `json:"Certificate,omitempty"`
	//
	FriendlyName *string `json:"FriendlyName,omitempty"`
	//
	PrivateKey *string `json:"PrivateKey,omitempty"`
	//
	Sandbox *bool `json:"Sandbox,omitempty"`
	//
	Secret *string `json:"Secret,omitempty"`
}

func (params *UpdateCredentialParams) SetApiKey(ApiKey string) *UpdateCredentialParams {
	params.ApiKey = &ApiKey
	return params
}
func (params *UpdateCredentialParams) SetCertificate(Certificate string) *UpdateCredentialParams {
	params.Certificate = &Certificate
	return params
}
func (params *UpdateCredentialParams) SetFriendlyName(FriendlyName string) *UpdateCredentialParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateCredentialParams) SetPrivateKey(PrivateKey string) *UpdateCredentialParams {
	params.PrivateKey = &PrivateKey
	return params
}
func (params *UpdateCredentialParams) SetSandbox(Sandbox bool) *UpdateCredentialParams {
	params.Sandbox = &Sandbox
	return params
}
func (params *UpdateCredentialParams) SetSecret(Secret string) *UpdateCredentialParams {
	params.Secret = &Secret
	return params
}

func (c *ApiService) UpdateCredential(Sid string, params *UpdateCredentialParams) (*IpMessagingV1Credential, error) {
	path := "/v1/Credentials/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ApiKey != nil {
		data.Set("ApiKey", *params.ApiKey)
	}
	if params != nil && params.Certificate != nil {
		data.Set("Certificate", *params.Certificate)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.PrivateKey != nil {
		data.Set("PrivateKey", *params.PrivateKey)
	}
	if params != nil && params.Sandbox != nil {
		data.Set("Sandbox", fmt.Sprint(*params.Sandbox))
	}
	if params != nil && params.Secret != nil {
		data.Set("Secret", *params.Secret)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV1Credential{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}