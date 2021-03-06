// Package trainer provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package trainer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestEditor RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = http.DefaultClient
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditor = fn
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetTrainerAvailableHours request
	GetTrainerAvailableHours(ctx context.Context, params *GetTrainerAvailableHoursParams) (*http.Response, error)

	// MakeHourAvailable request  with any body
	MakeHourAvailableWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	MakeHourAvailable(ctx context.Context, body MakeHourAvailableJSONRequestBody) (*http.Response, error)

	// MakeHourUnavailable request  with any body
	MakeHourUnavailableWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	MakeHourUnavailable(ctx context.Context, body MakeHourUnavailableJSONRequestBody) (*http.Response, error)
}

func (c *Client) GetTrainerAvailableHours(ctx context.Context, params *GetTrainerAvailableHoursParams) (*http.Response, error) {
	req, err := NewGetTrainerAvailableHoursRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) MakeHourAvailableWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewMakeHourAvailableRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) MakeHourAvailable(ctx context.Context, body MakeHourAvailableJSONRequestBody) (*http.Response, error) {
	req, err := NewMakeHourAvailableRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) MakeHourUnavailableWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewMakeHourUnavailableRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) MakeHourUnavailable(ctx context.Context, body MakeHourUnavailableJSONRequestBody) (*http.Response, error) {
	req, err := NewMakeHourUnavailableRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

// NewGetTrainerAvailableHoursRequest generates requests for GetTrainerAvailableHours
func NewGetTrainerAvailableHoursRequest(server string, params *GetTrainerAvailableHoursParams) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/trainer/calendar")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if queryFrag, err := runtime.StyleParam("form", true, "dateFrom", params.DateFrom); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if queryFrag, err := runtime.StyleParam("form", true, "dateTo", params.DateTo); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewMakeHourAvailableRequest calls the generic MakeHourAvailable builder with application/json body
func NewMakeHourAvailableRequest(server string, body MakeHourAvailableJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewMakeHourAvailableRequestWithBody(server, "application/json", bodyReader)
}

// NewMakeHourAvailableRequestWithBody generates requests for MakeHourAvailable with any type of body
func NewMakeHourAvailableRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/trainer/calendar/make-hour-available")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewMakeHourUnavailableRequest calls the generic MakeHourUnavailable builder with application/json body
func NewMakeHourUnavailableRequest(server string, body MakeHourUnavailableJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewMakeHourUnavailableRequestWithBody(server, "application/json", bodyReader)
}

// NewMakeHourUnavailableRequestWithBody generates requests for MakeHourUnavailable with any type of body
func NewMakeHourUnavailableRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/trainer/calendar/make-hour-unavailable")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetTrainerAvailableHours request
	GetTrainerAvailableHoursWithResponse(ctx context.Context, params *GetTrainerAvailableHoursParams) (*GetTrainerAvailableHoursResponse, error)

	// MakeHourAvailable request  with any body
	MakeHourAvailableWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*MakeHourAvailableResponse, error)

	MakeHourAvailableWithResponse(ctx context.Context, body MakeHourAvailableJSONRequestBody) (*MakeHourAvailableResponse, error)

	// MakeHourUnavailable request  with any body
	MakeHourUnavailableWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*MakeHourUnavailableResponse, error)

	MakeHourUnavailableWithResponse(ctx context.Context, body MakeHourUnavailableJSONRequestBody) (*MakeHourUnavailableResponse, error)
}

type GetTrainerAvailableHoursResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Date
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r GetTrainerAvailableHoursResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetTrainerAvailableHoursResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type MakeHourAvailableResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON204      *[]Date
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r MakeHourAvailableResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r MakeHourAvailableResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type MakeHourUnavailableResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON204      *[]Date
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r MakeHourUnavailableResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r MakeHourUnavailableResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetTrainerAvailableHoursWithResponse request returning *GetTrainerAvailableHoursResponse
func (c *ClientWithResponses) GetTrainerAvailableHoursWithResponse(ctx context.Context, params *GetTrainerAvailableHoursParams) (*GetTrainerAvailableHoursResponse, error) {
	rsp, err := c.GetTrainerAvailableHours(ctx, params)
	if err != nil {
		return nil, err
	}
	return ParseGetTrainerAvailableHoursResponse(rsp)
}

// MakeHourAvailableWithBodyWithResponse request with arbitrary body returning *MakeHourAvailableResponse
func (c *ClientWithResponses) MakeHourAvailableWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*MakeHourAvailableResponse, error) {
	rsp, err := c.MakeHourAvailableWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseMakeHourAvailableResponse(rsp)
}

func (c *ClientWithResponses) MakeHourAvailableWithResponse(ctx context.Context, body MakeHourAvailableJSONRequestBody) (*MakeHourAvailableResponse, error) {
	rsp, err := c.MakeHourAvailable(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParseMakeHourAvailableResponse(rsp)
}

// MakeHourUnavailableWithBodyWithResponse request with arbitrary body returning *MakeHourUnavailableResponse
func (c *ClientWithResponses) MakeHourUnavailableWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*MakeHourUnavailableResponse, error) {
	rsp, err := c.MakeHourUnavailableWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseMakeHourUnavailableResponse(rsp)
}

func (c *ClientWithResponses) MakeHourUnavailableWithResponse(ctx context.Context, body MakeHourUnavailableJSONRequestBody) (*MakeHourUnavailableResponse, error) {
	rsp, err := c.MakeHourUnavailable(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParseMakeHourUnavailableResponse(rsp)
}

// ParseGetTrainerAvailableHoursResponse parses an HTTP response from a GetTrainerAvailableHoursWithResponse call
func ParseGetTrainerAvailableHoursResponse(rsp *http.Response) (*GetTrainerAvailableHoursResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetTrainerAvailableHoursResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Date
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json"):
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseMakeHourAvailableResponse parses an HTTP response from a MakeHourAvailableWithResponse call
func ParseMakeHourAvailableResponse(rsp *http.Response) (*MakeHourAvailableResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &MakeHourAvailableResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 204:
		var dest []Date
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON204 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json"):
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseMakeHourUnavailableResponse parses an HTTP response from a MakeHourUnavailableWithResponse call
func ParseMakeHourUnavailableResponse(rsp *http.Response) (*MakeHourUnavailableResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &MakeHourUnavailableResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 204:
		var dest []Date
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON204 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json"):
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}
