package docbase

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

/*
An ErrorResponse reports one or more errors caused by an API request.

Docbase API docs: https://developer.github.com/v3/#client-errors
*/
type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Messages []string       `json:"messages"` // more detail on individual errors
	Value    string         `json:"error"`    // error message
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %+v %v",
		r.Response.Request.Method, sanitizeURL(r.Response.Request.URL),
		r.Response.StatusCode, r.Value, r.Messages)
}

// RateLimitError occurs when Docbase returns 429 Too many requests response
// with a rate limit remaining value of 0.
type RateLimitError struct {
	Rate     Rate           // Rate specifies last known rate limit for the client
	Response *http.Response // HTTP response that caused this error
	Messages []string       `json:"messages"` // error message
}

func (r *RateLimitError) Error() string {
	return fmt.Sprintf("%v %v: %d %+v %v",
		r.Response.Request.Method, sanitizeURL(r.Response.Request.URL),
		r.Response.StatusCode, r.Messages, formatRateReset(time.Until(r.Rate.Reset.Time)))
}

// sanitizeURL redacts the client_secret parameter from the URL which may be
// exposed to the user.
func sanitizeURL(uri *url.URL) *url.URL {
	if uri == nil {
		return nil
	}
	params := uri.Query()
	if len(params.Get("client_secret")) > 0 {
		params.Set("client_secret", "REDACTED")
		uri.RawQuery = params.Encode()
	}
	return uri
}

/*
An Error reports more details on an individual error in an ErrorResponse.
These are the possible validation error codes:

	missing:
	    resource does not exist
	missing_field:
	    a required field on a resource has not been set
	invalid:
	    the formatting of a field is invalid
	already_exists:
	    another resource has the same valid as this field
	custom:
	    some resources return this (e.g. github.User.CreateKey()), additional
	    information is set in the Message field of the Error

Docbase API docs: https://developer.github.com/v3/#client-errors
*/
type Error struct {
	Resource string `json:"resource"` // resource on which the error occurred
	Field    string `json:"field"`    // field on which the error occurred
	Code     string `json:"code"`     // validation error code
	Message  string `json:"message"`  // Message describing the error. Errors with Code == "custom" will always have this set.
}

func (e *Error) Error() string {
	return fmt.Sprintf("%v error caused by %v field on %v resource",
		e.Code, e.Field, e.Resource)
}

// checkResponse checks the API response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range.
// API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse. Any other
// response body will be silently ignored.
//
// The error type will be *RateLimitError for rate limit exceeded errors.
func checkResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
	data, err := io.ReadAll(r.Body)
	if err == nil && data != nil {
		if err := json.Unmarshal(data, errorResponse); err != nil {
			return err
		}
	}
	switch r.StatusCode {
	case http.StatusTooManyRequests:
		return &RateLimitError{
			Rate:     parseRate(r),
			Response: errorResponse.Response,
			Messages: errorResponse.Messages,
		}
	default:
		return errorResponse
	}
}

func sanitizeError(ctx context.Context, err error) error {
	// If we got an error, and the context has been canceled,
	// the context's error is probably more useful.
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	// If the error type is *url.Error, sanitize its URL before returning.
	if e, ok := err.(*url.Error); ok {
		if url, err := url.Parse(e.URL); err == nil {
			e.URL = sanitizeURL(url).String()
			return e
		}
	}

	return err
}
