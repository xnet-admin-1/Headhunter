// Package httpclient provides per-client HTTP identity policy.
package httpclient

import (
	"fmt"
	"net/http"
	"strings"
	"sync/atomic"
)

// DefaultUserAgent identifies Headhunter-controlled requests.
const DefaultUserAgent = "HEADHUNTER"

// defaultUserAgentVersion is the build version of the default identity.
// It is empty until the CLI records it.
var defaultUserAgentVersion atomic.Value

// SetDefaultUserAgentVersion records the build version of the default
// identity. Call it once at startup.
func SetDefaultUserAgentVersion(v string) {
	v = strings.TrimSpace(v)
	if v == "" {
		return
	}
	defaultUserAgentVersion.Store(v)
}

// DefaultUserAgentValue returns the User-Agent value for requests that carry
// no explicit identity. It is the product name and the build version.
//
// The version is not decoration. Some gateways, including the endpoint that
// freeinference.org runs, answer HTTP 403 to a bare product name. They read
// the User-Agent as a bot signature. A name and a version is what a command
// line tool normally sends, so the default identity carries the version.
func DefaultUserAgentValue() string {
	version := "dev"
	if v, ok := defaultUserAgentVersion.Load().(string); ok && v != "" {
		version = v
	}
	return DefaultUserAgent + "/" + version
}

// ResolveUserAgent validates identity without trimming valid input.
func ResolveUserAgent(value string) (string, error) {
	if value == "" {
		return DefaultUserAgentValue(), nil
	}
	for _, b := range []byte(value) {
		if b < 32 || b == 127 {
			return "", fmt.Errorf("user agent contains an HTTP control character")
		}
	}
	return value, nil
}

// WithUserAgent copies a client and enforces identity on each outgoing request.
// Callers must validate non-empty values with ResolveUserAgent first.
func WithUserAgent(client *http.Client, value string) *http.Client {
	if value == "" {
		value = DefaultUserAgentValue()
	}
	copy := http.Client{}
	if client != nil {
		copy = *client
	}
	base := copy.Transport
	if base == nil {
		base = http.DefaultTransport
	}
	copy.Transport = &userAgentTransport{base: base, value: value}
	return &copy
}

type userAgentTransport struct {
	base  http.RoundTripper
	value string
}

func (t *userAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	clone := req.Clone(req.Context())
	if clone.Header == nil {
		clone.Header = make(http.Header)
	}
	for key := range clone.Header {
		if strings.EqualFold(key, "User-Agent") {
			delete(clone.Header, key)
		}
	}
	clone.Header.Set("User-Agent", t.value)
	return t.base.RoundTrip(clone)
}

func (t *userAgentTransport) CloseIdleConnections() {
	if closer, ok := t.base.(interface{ CloseIdleConnections() }); ok {
		closer.CloseIdleConnections()
	}
}
