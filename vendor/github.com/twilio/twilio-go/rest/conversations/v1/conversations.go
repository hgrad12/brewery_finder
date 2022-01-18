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
	"encoding/json"
	"fmt"
	"net/url"

	"strings"
	"time"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateConversation'
type CreateConversationParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
	// An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned.
	Attributes *string `json:"Attributes,omitempty"`
	// The date that this resource was created.
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The date that this resource was last updated.
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The human-readable name of this conversation, limited to 256 characters. Optional.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The unique ID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) this conversation belongs to.
	MessagingServiceSid *string `json:"MessagingServiceSid,omitempty"`
	// Current state of this conversation. Can be either `active`, `inactive` or `closed` and defaults to `active`
	State *string `json:"State,omitempty"`
	// ISO8601 duration when conversation will be switched to `closed` state. Minimum value for this timer is 10 minutes.
	TimersClosed *string `json:"Timers.Closed,omitempty"`
	// ISO8601 duration when conversation will be switched to `inactive` state. Minimum value for this timer is 1 minute.
	TimersInactive *string `json:"Timers.Inactive,omitempty"`
	// An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource's `sid` in the URL.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *CreateConversationParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *CreateConversationParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}
func (params *CreateConversationParams) SetAttributes(Attributes string) *CreateConversationParams {
	params.Attributes = &Attributes
	return params
}
func (params *CreateConversationParams) SetDateCreated(DateCreated time.Time) *CreateConversationParams {
	params.DateCreated = &DateCreated
	return params
}
func (params *CreateConversationParams) SetDateUpdated(DateUpdated time.Time) *CreateConversationParams {
	params.DateUpdated = &DateUpdated
	return params
}
func (params *CreateConversationParams) SetFriendlyName(FriendlyName string) *CreateConversationParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateConversationParams) SetMessagingServiceSid(MessagingServiceSid string) *CreateConversationParams {
	params.MessagingServiceSid = &MessagingServiceSid
	return params
}
func (params *CreateConversationParams) SetState(State string) *CreateConversationParams {
	params.State = &State
	return params
}
func (params *CreateConversationParams) SetTimersClosed(TimersClosed string) *CreateConversationParams {
	params.TimersClosed = &TimersClosed
	return params
}
func (params *CreateConversationParams) SetTimersInactive(TimersInactive string) *CreateConversationParams {
	params.TimersInactive = &TimersInactive
	return params
}
func (params *CreateConversationParams) SetUniqueName(UniqueName string) *CreateConversationParams {
	params.UniqueName = &UniqueName
	return params
}

// Create a new conversation in your account&#39;s default service
func (c *ApiService) CreateConversation(params *CreateConversationParams) (*ConversationsV1Conversation, error) {
	path := "/v1/Conversations"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.DateCreated != nil {
		data.Set("DateCreated", fmt.Sprint((*params.DateCreated).Format(time.RFC3339)))
	}
	if params != nil && params.DateUpdated != nil {
		data.Set("DateUpdated", fmt.Sprint((*params.DateUpdated).Format(time.RFC3339)))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.MessagingServiceSid != nil {
		data.Set("MessagingServiceSid", *params.MessagingServiceSid)
	}
	if params != nil && params.State != nil {
		data.Set("State", *params.State)
	}
	if params != nil && params.TimersClosed != nil {
		data.Set("Timers.Closed", *params.TimersClosed)
	}
	if params != nil && params.TimersInactive != nil {
		data.Set("Timers.Inactive", *params.TimersInactive)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1Conversation{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteConversation'
type DeleteConversationParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
}

func (params *DeleteConversationParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *DeleteConversationParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}

// Remove a conversation from your account&#39;s default service
func (c *ApiService) DeleteConversation(Sid string, params *DeleteConversationParams) error {
	path := "/v1/Conversations/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch a conversation from your account&#39;s default service
func (c *ApiService) FetchConversation(Sid string) (*ConversationsV1Conversation, error) {
	path := "/v1/Conversations/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1Conversation{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListConversation'
type ListConversationParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListConversationParams) SetPageSize(PageSize int) *ListConversationParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListConversationParams) SetLimit(Limit int) *ListConversationParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Conversation records from the API. Request is executed immediately.
func (c *ApiService) PageConversation(params *ListConversationParams, pageToken, pageNumber string) (*ListConversationResponse, error) {
	path := "/v1/Conversations"

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

	ps := &ListConversationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Conversation records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListConversation(params *ListConversationParams) ([]ConversationsV1Conversation, error) {
	if params == nil {
		params = &ListConversationParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageConversation(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ConversationsV1Conversation

	for response != nil {
		records = append(records, response.Conversations...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListConversationResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListConversationResponse)
	}

	return records, err
}

// Streams Conversation records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamConversation(params *ListConversationParams) (chan ConversationsV1Conversation, error) {
	if params == nil {
		params = &ListConversationParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageConversation(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ConversationsV1Conversation, 1)

	go func() {
		for response != nil {
			for item := range response.Conversations {
				channel <- response.Conversations[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListConversationResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListConversationResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListConversationResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListConversationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateConversation'
type UpdateConversationParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
	// An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned.
	Attributes *string `json:"Attributes,omitempty"`
	// The date that this resource was created.
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The date that this resource was last updated.
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The human-readable name of this conversation, limited to 256 characters. Optional.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The unique ID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) this conversation belongs to.
	MessagingServiceSid *string `json:"MessagingServiceSid,omitempty"`
	// Current state of this conversation. Can be either `active`, `inactive` or `closed` and defaults to `active`
	State *string `json:"State,omitempty"`
	// ISO8601 duration when conversation will be switched to `closed` state. Minimum value for this timer is 10 minutes.
	TimersClosed *string `json:"Timers.Closed,omitempty"`
	// ISO8601 duration when conversation will be switched to `inactive` state. Minimum value for this timer is 1 minute.
	TimersInactive *string `json:"Timers.Inactive,omitempty"`
	// An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource's `sid` in the URL.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *UpdateConversationParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *UpdateConversationParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}
func (params *UpdateConversationParams) SetAttributes(Attributes string) *UpdateConversationParams {
	params.Attributes = &Attributes
	return params
}
func (params *UpdateConversationParams) SetDateCreated(DateCreated time.Time) *UpdateConversationParams {
	params.DateCreated = &DateCreated
	return params
}
func (params *UpdateConversationParams) SetDateUpdated(DateUpdated time.Time) *UpdateConversationParams {
	params.DateUpdated = &DateUpdated
	return params
}
func (params *UpdateConversationParams) SetFriendlyName(FriendlyName string) *UpdateConversationParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateConversationParams) SetMessagingServiceSid(MessagingServiceSid string) *UpdateConversationParams {
	params.MessagingServiceSid = &MessagingServiceSid
	return params
}
func (params *UpdateConversationParams) SetState(State string) *UpdateConversationParams {
	params.State = &State
	return params
}
func (params *UpdateConversationParams) SetTimersClosed(TimersClosed string) *UpdateConversationParams {
	params.TimersClosed = &TimersClosed
	return params
}
func (params *UpdateConversationParams) SetTimersInactive(TimersInactive string) *UpdateConversationParams {
	params.TimersInactive = &TimersInactive
	return params
}
func (params *UpdateConversationParams) SetUniqueName(UniqueName string) *UpdateConversationParams {
	params.UniqueName = &UniqueName
	return params
}

// Update an existing conversation in your account&#39;s default service
func (c *ApiService) UpdateConversation(Sid string, params *UpdateConversationParams) (*ConversationsV1Conversation, error) {
	path := "/v1/Conversations/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.DateCreated != nil {
		data.Set("DateCreated", fmt.Sprint((*params.DateCreated).Format(time.RFC3339)))
	}
	if params != nil && params.DateUpdated != nil {
		data.Set("DateUpdated", fmt.Sprint((*params.DateUpdated).Format(time.RFC3339)))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.MessagingServiceSid != nil {
		data.Set("MessagingServiceSid", *params.MessagingServiceSid)
	}
	if params != nil && params.State != nil {
		data.Set("State", *params.State)
	}
	if params != nil && params.TimersClosed != nil {
		data.Set("Timers.Closed", *params.TimersClosed)
	}
	if params != nil && params.TimersInactive != nil {
		data.Set("Timers.Inactive", *params.TimersInactive)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1Conversation{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
