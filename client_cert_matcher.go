package client_cert_matcher

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"net/http"
)

func init() {
	caddy.RegisterModule(MatchClientCert{})
}

// MatchClientCert matches based on client certificate CN.
// Names in this list are allowed.
type MatchClientCert []string

func (MatchClientCert) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.matchers.client_cert",
		New: func() caddy.Module { return new(MatchClientCert) },
	}
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

// Match matches client certificate CN
func (m MatchClientCert) Match(r *http.Request) bool {
	cn := r.TLS.PeerCertificates[0].Subject
	return contains(m, cn.CommonName)
}

// Interface guard
var _ caddyhttp.RequestMatcher = (*MatchClientCert)(nil)
