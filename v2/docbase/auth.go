package docbase

import "net/http"

// TokenTransport is an http.RoundTripper that authenticates all requests
// with the provided access token.
type TokenTransport struct {
	Token string

	// Transport is the underlying HTTP transport to use when making requests.
	// It will default to http.DefaultTransport if nil.
	Transport http.RoundTripper
}

// RoundTrip implements the RoundTripper interface.
func (t *TokenTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// To set extra headers, we must make a copy of the Request so
	// that we don't modify the Request we were given. This is required by the
	// specification of http.RoundTripper.
	//
	// Since we are going to modify only req.Header here, we only need a deep copy
	// of req.Header.
	req2 := new(http.Request)
	*req2 = *req
	req2.Header = make(http.Header, len(req.Header))
	for k, s := range req.Header {
		req2.Header[k] = append([]string(nil), s...)
	}

	req2.Header.Set(headerToken, t.Token)
	return t.transport().RoundTrip(req2)
}

// Client returns an *http.Client that makes requests that are authenticated
// with the provided access token.
func (t *TokenTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func (t *TokenTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

func NewAuthClient(domain, token string) *Client {
	transport := TokenTransport{Token: token}
	return NewClient(domain, transport.Client())
}
